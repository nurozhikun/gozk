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
// g1.AsynRun()
// defer g1.AsynRun()
//

type Goloop struct {
	FnLoop  func()
	Timeout time.Duration
	RunAdds *int32
	chquit  chan int32
}

func WaitingForRunning(RunAdds *int32) {
	for *RunAdds != 0 {
		time.Sleep(25 * time.Millisecond)
	}
}

func (g *Goloop) AsynRun() {
	atomic.AddInt32(g.RunAdds, 1)
	g.chquit = make(chan int32)
	go func() {
		var ret int32
	LOOP_OUT:
		for {
			g.FnLoop()
			select {
			case <-g.chquit:
				break LOOP_OUT
			case <-time.After(g.Timeout):
			}
		}
		g.chquit <- ret
	}()
}

func (g *Goloop) AsynEnd() {
	go func() {
		g.chquit <- 1
		<-g.chquit
		atomic.AddInt32(g.RunAdds, -1)
	}()
}
