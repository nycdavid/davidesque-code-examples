package ticker_test

import (
	"testing"

	"github.com/nycdavid/davidesque-code-examples/how-to-unit-test-a-function-that-uses-ticker"
	_assert "github.com/stretchr/testify/assert"
)

func Test_PeriodicPublish(t *testing.T) {
	assert := _assert.New(t)

	printer := &ticker.Printer{}
	printer.Start()

	assert.Equal(printer.Iteration, 5)
}
