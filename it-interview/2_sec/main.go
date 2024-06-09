package main

func main() {
	v := 5
	p := &v

	println(*p) // 5, разыменование
	changePointer(p)
	println(*p) // 5
}

func changePointer(p *int) {
	v := 3
	p = &v // меняется локальная p (тут не ссылка на ссылку, а просто p - ссылка)
}
