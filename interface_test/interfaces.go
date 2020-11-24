package main

import "fmt"

//定义接口
type Mymath interface {
	Add() int
	Sub() int
}

//定义结构体
type One struct {
	X, Y int
}

//实现方法Add
func (i One) Add() int {
	return i.X + i.Y
}

//实现方法Sub
func (j One) Sub() int {
	return j.X - j.Y
}

//主程序
func main() {
	var r Mymath
	r = One{10, 100}
	fmt.Println(r.Add())
	fmt.Println(r.Sub())
}
