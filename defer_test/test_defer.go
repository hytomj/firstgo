package main

import "fmt"

func main() {
	s := make([]int, 1, 5)
	fmt.Printf("%v\n", s)

	x, j := 22, 1003
	//var p *int
	p := &x
	fmt.Println(*p)
	p = &j
	fmt.Println(*p)
	fmt.Println("start...")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("end...")
}
