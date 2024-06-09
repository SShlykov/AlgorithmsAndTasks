package main

import (
	"fmt"
	"runtime"
)

// все выведет и не заблокируется если го > 1.13, runtime вытеснит main в пользу локальной рутины
func main() {
	runtime.GOMAXPROCS(1)

	done := false

	go func() {
		done = true
	}()

	for !done {
	}

	fmt.Println("finished")
}
