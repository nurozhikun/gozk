package zutils

import (
	"sync/atomic"
	"time"
)

//example
// var RunAdds int32
// defer zutils.WaitingForRunning(&RunAdds)
// g1 := &Goloop {
//     FnLoop: func() {},
//	   Timeout: 100*time.Millisecond,
//     RunAdds: &RunAdds,
// }
// g1.asynRun()
// defer g1.asynRun()
//

type Goloop struct {
	FnLoop  func() (waittimes int) //necessary
	RunAdds *int32                 //necessary
	Timeout time.Duration
	chquit  chan int32
}

func WaitingForRunning(RunAdds *int32) {
	for *RunAdds != 0 {
		time.Sleep(25 * time.Millisecond)
	}
}

func Go(fn func() int, RunAdds *int32, Timeout time.Duration) *Goloop {
	glp := &Goloop{
		FnLoop:  fn,
		RunAdds: RunAdds,
		Timeout: Timeout,
	}
	go glp.asynRun()
	return glp
}

func (g *Goloop) Close() {
	g.asynEnd()
}

func (g *Goloop) asynRun() {
	atomic.AddInt32(g.RunAdds, 1)
	g.chquit = make(chan int32)
	go func() {
		var ret int32
		if g.Timeout == 0 {
			g.Timeout = 25
		}
	LOOP_OUT:
		for {
			n := g.FnLoop()
			if n > 0 {
				n = n * int(g.Timeout)
			} else {
				n = int(g.Timeout)
			}
			select {
			case <-g.chquit:
				break LOOP_OUT
			case <-time.After(time.Duration(n)):
			}
		}
		g.chquit <- ret
	}()
}

func (g *Goloop) asynEnd() {
	go func() {
		g.chquit <- 1
		<-g.chquit
		atomic.AddInt32(g.RunAdds, -1)
	}()
}
