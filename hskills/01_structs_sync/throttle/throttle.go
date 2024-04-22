package throttle

import (
	"sync"
	"time"
)

// Throttle - HOF и замыкание
func Throttle(f func(), duration time.Duration) func() {
	var once sync.Once
	var mu sync.Mutex
	var lastTime time.Time

	return func() {
		mu.Lock()
		defer mu.Unlock()

		once.Do(func() {
			lastTime = time.Now().Add(-duration)
		})

		if time.Since(lastTime) < duration {
			return
		}

		f()

		lastTime = time.Now()
	}
}
