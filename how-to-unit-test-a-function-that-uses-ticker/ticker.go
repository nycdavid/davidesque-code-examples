package ticker

import (
	"context"
	"fmt"
	"time"
)

type (
	Printer struct {
		Iteration int
	}
)

func (p *Printer) Start(ctx context.Context, done chan struct{}) {
	defer p.Clock.Stop()

	for {
		select {
		case <-t.C:
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
