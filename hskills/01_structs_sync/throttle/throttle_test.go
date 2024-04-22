package throttle

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestThrottle(t *testing.T) {
	var counter int
	f := func() {
		counter++
	}

	throttled := Throttle(f, 1000)
	throttled()
	throttled()
	throttled()
	time.Sleep(1000 * time.Millisecond)
	throttled()

	assert.Equal(t, 2, counter)
}
