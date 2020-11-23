package main

import "fmt"

type Person struct {
	Name string
	Sex  string
	Age  int
}

func main() {
	var p1 = Person{
		Name: "张三",
		Sex:  "男",
		Age:  30,
	}
	fmt.Printf("%#v", p1)
}
