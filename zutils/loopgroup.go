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

///example
///

type FnRunner = func()

type LoopGroup struct {
	sync.WaitGroup
	mtx   sync.Mutex
	loops map[string]*loop
	stops []FnRunner
	// eventLoops map[string]*eventLoop
}

func NewLoopGroup() *LoopGroup {
	return &LoopGroup{
		loops: make(map[string]*loop),
	}
}

type loop struct {
	chquit chan int32
}

func (lg *LoopGroup) GoOnce(fnProc func()) {
	lg.mtx.Lock()
	defer lg.mtx.Unlock()
	lg.Add(1)
	go func() {
		defer lg.Done()
		fnProc()
	}()
}

// /针对自己有Stop函数的调用，比如Http server
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

func (lg *LoopGroup) createLoop(key string) (l *loop, bExist bool) {
	lg.mtx.Lock()
	defer lg.mtx.Unlock()
	l, bExist = lg.loops[key]
	if bExist {
		return
	}
	l = &loop{chquit: make(chan int32)}
	lg.loops[key] = l
	return
}

func (lg *LoopGroup) deleteLoop(key string) {
	lg.mtx.Lock()
	defer lg.mtx.Unlock()
	delete(lg.loops, key) //结束的时候从map中删除
}

// /添加一个loop函数，用key来命名，可以用ExitLoop来结束，程序结束时也会自动结束
func (lg *LoopGroup) GoLoop(key string, fn func() int, timeout time.Duration, fnCancel func()) error {
	l, bExist := lg.createLoop(key)
	if bExist {
		return fmt.Errorf("the key %s has existed", key)
	}
	lg.Add(1)
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
		lg.deleteLoop(key)
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
}

func (lg *LoopGroup) WaitForEnter(enter string) {
	WaitForEnter(enter)
	lg.Wait()
}

func (lg *LoopGroup) Wait() {
	// lg.mtx.Lock()
	// defer lg.mtx.Unlock()
	for _, stop := range lg.stops {
		go stop()
	}
	lg.stops = nil
	// zlogger.Info(lg.loops)
	for k, l := range lg.loops {
		go func(key string, lp *loop) {
			lp.chquit <- 1
		}(k, l)
	}
	lg.loops = nil
	lg.WaitGroup.Wait()
}

func (lg *LoopGroup) Done() {
	lg.mtx.Lock()
	defer lg.mtx.Unlock()
	lg.WaitGroup.Done()
}
