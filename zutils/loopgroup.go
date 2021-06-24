package zutils

import (
	"fmt"
	"sync"
	"time"

	"gitee.com/sienectagv/gozk/zlogger"
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

///example
///

type FnRunner = func()

type LoopGroup struct {
	sync.WaitGroup
	mtx   sync.Mutex
	loops map[string]*loop
	stops []FnRunner
}

func NewLoopGroup() *LoopGroup {
	return &LoopGroup{
		loops: make(map[string]*loop),
	}
}

type loop struct {
	chquit chan int32
}

func (lg *LoopGroup) AddAsyncBlock(fnBlock func(), fnStop func()) {
	lg.mtx.Lock()
	defer lg.mtx.Unlock()
	lg.Add(1)
	lg.stops = append(lg.stops, fnStop)
	go func() {
		defer lg.Done()
		fnBlock()
	}()
}

func (lg *LoopGroup) GoLoop(key string, fn func() int, timeout time.Duration, fnCancel func()) error {
	lg.mtx.Lock()
	defer lg.mtx.Unlock()
	_, ok := lg.loops[key]
	if ok {
		return fmt.Errorf("the key %s has existed", key)
	}
	lg.Add(1)
	l := &loop{chquit: make(chan int32)}
	// if nil == lg.loops {
	// 	lg.loops = make(map[string]*loop)
	// }
	zlogger.Info(lg.loops)
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
		if nil != fnCancel {
			fnCancel()
		}
		lg.Done()
	}()
	return nil
}

func (lg *LoopGroup) ExitLoop(key string) {
	lg.mtx.Lock()
	defer lg.mtx.Unlock()
	v, ok := lg.loops[key]
	if !ok {
		return
	}
	v.chquit <- 1
	delete(lg.loops, key)
}

func (lg *LoopGroup) WaitForEnter(enter string) {
	WaitForEnter(enter)
	lg.Wait()
}

func (lg *LoopGroup) Wait() {
	lg.mtx.Lock()
	defer lg.mtx.Unlock()
	for _, stop := range lg.stops {
		go stop()
	}
	lg.stops = nil
	for _, v := range lg.loops {
		go func() {
			v.chquit <- 1
		}()
	}
	lg.loops = nil
	lg.WaitGroup.Wait()
}
