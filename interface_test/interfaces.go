package main

import "fmt"

type Mymath interface {
	Add()
	Sub()
}

type One struct {
	X, Y int
}

func (i One) Add() {
	fmt.Println(i.X + i.Y)
}

func (j One) Sub() {
	fmt.Println(j.X - j.Y)
}

func main() {
	var r Mymath
	r = One{10, 100}
	r.Add()
	r.Sub()
}
