package main

import "fmt"

type Parent struct{}

func (c *Parent) Print() {
	fmt.Println("parent")
}

type Child struct {
	Parent
}

func (p *Child) Print() {
	fmt.Println("child")
}

func main() {
	var child Child
	child.Print() // child. Но внутри будет parent с методом print. А сама структура child будет использовать свой метод
	// при этом если бы его не было, то код был бы валиден.
}
