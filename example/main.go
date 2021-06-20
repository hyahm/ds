package main

import "fmt"

type zouchang struct {
	w, h int
}

type adder interface {
	add(a, b int)
}

func (zc *zouchang) Add(add adder) {
	add.add(zc.h, zc.w)
}

type x func(a, b int) int

func (x x) add(a, b int) int {
	return x(a, b)
}

func main() {
	var temp x
	temp = func(a, b int) int {
		return a + b + 100
	}
	result := temp.add(10, 20)
	fmt.Println(result)
}
