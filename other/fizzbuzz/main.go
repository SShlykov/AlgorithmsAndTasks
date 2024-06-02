package main

import "fmt"

func fizzbuzz(n int) {
	for i := 1; i <= n; i++ {
		t, f := i%3, i%5
		switch {
		case t == 0 && f == 0:
			fmt.Printf("FizzBuzz")
		case t == 0:
			fmt.Printf("Fizz")
		case f == 0:
			fmt.Printf("Buzz")
		default:
			fmt.Printf("%d", i)
		}

		fmt.Printf(", ")
	}
}
