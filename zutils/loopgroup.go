package zutils

import (
	"fmt"
	"sync"
	"time"
)

func WaitForEnter(s string) {
	var sRead string
	for {
		fmt.Printf("enter '%s' to exit the process...\n", s)
		fmt.Scanln(&sRead)
		if sRead == s {
			break
		}
	}
}

type loop struct {
	// fn      func() int //necessary
	// timeout time.Duration
	chquit chan int32
}

type LoopGroup struct {
	sync.WaitGroup
	mtx   sync.Mutex
	loops map[string]*loop
}

func (lg *LoopGroup) Go(key string, fn func() int, timeout time.Duration, fnCancel func()) error {
	lg.mtx.Lock()
	defer lg.mtx.Unlock()
	_, ok := lg.loops[key]
	if ok {
		return fmt.Errorf("the key %s has existed", key)
	}
	lg.Add(1)
	l := &loop{chquit: make(chan int32)}
	lg.loops[key] = l
	go func() {
		if timeout == 0 {
			timeout = 25
		}
	LOOP_OUT:
		for {
			n := fn()
			if n > 0 {
				n = n * int(timeout)
			} else {
				n = int(timeout)
			}
			select {
			case <-l.chquit:
				break LOOP_OUT
			case <-time.After(time.Duration(n)):
			}
		}
		lg.Done()
	}()
	return nil
}

func (lg *LoopGroup) Exit(key string) {
	lg.mtx.Lock()
	defer lg.mtx.Unlock()
	v, ok := lg.loops[key]
	if !ok {
		return
	}
	v.chquit <- 1
	delete(lg.loops, key)
}
