package ticker_test

import (
	"context"
	"testing"
	"time"

	"github.com/nycdavid/davidesque-code-examples/how-to-unit-test-a-function-that-uses-ticker"
	_assert "github.com/stretchr/testify/assert"
)

type (
	mockClock struct {
		c chan time.Time
	}
)

func (m *mockClock) C() <-chan time.Time {
	return m.c
}

func (m *mockClock) Stop() {
}

func Test_PeriodicPublish(t *testing.T) {
	assert := _assert.New(t)

	timeChan := make(chan time.Time)

	printer := &ticker.Printer{
		Clock: &mockClock{
			c: timeChan,
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go printer.Start(ctx, done)

	timeChan <- time.Now()
	timeChan <- time.Now()
	timeChan <- time.Now()
	timeChan <- time.Now()
	timeChan <- time.Now()

	cancel()
	<-done

	assert.Equal(5, printer.Iteration)
}
