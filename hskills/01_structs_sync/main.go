package main

import "fmt"

func add(s []string) {
	s = append(s, "x")
}

func main() {
	s := []string{"a", "b", "c"}
	r := s[1:2]
	add(r)
	fmt.Println(s)
}
