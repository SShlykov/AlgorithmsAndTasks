package main

import "fmt"

func main() {
	var sl []int

	for i := range 10 {
		sl = append(sl, i)
		fmt.Printf("sl: %v, len: %d, cap: %d\n", sl, len(sl), cap(sl))
	}
}
