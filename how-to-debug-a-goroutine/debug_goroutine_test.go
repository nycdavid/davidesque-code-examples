package how_to_debug_a_goroutine_test

import (
	"runtime"
	"sync"
	"testing"

	"github.com/nycdavid/davidesque-code-examples/how-to-debug-a-goroutine"
)

func Test_HeavyProcess(t *testing.T) {
	runtime.GOMAXPROCS(1)
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)

		// Insert breakpoint here
		go how_to_debug_a_goroutine.HeavyProcess(wg, i)
	}

	wg.Wait()
}
