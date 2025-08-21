package ticker

import (
	"context"
	"fmt"
	"time"
)

type (
	Printer struct {
		Iteration int
		Clock     Ticker
	}

	Ticker interface {
		C() <-chan time.Time
		Stop()
	}

	realClock struct {
		t *time.Ticker
	}
)

func (r *realClock) C() <-chan time.Time {
	return r.t.C
}

func (r *realClock) Stop() {
	r.t.Stop()
}

func (p *Printer) Start(ctx context.Context, done chan struct{}) {
	defer p.Clock.Stop()

	for {
		select {
		case <-p.Clock.C():
			// Directly tied to wall clock, hard to control in tests
			fmt.Println("tick at", time.Now())
			p.Iteration++
		case <-ctx.Done():
			fmt.Println("exiting...")
			close(done)
			return
		}
	}
}
