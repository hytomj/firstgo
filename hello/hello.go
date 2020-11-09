package main

import (
	"fmt"
)

func showArray() {
	s := []string{"java", "python", "golang"}
	//s := [...]int{11, 22, 99, 8, 6}
	for i, v := range s {
		fmt.Println(i, "->", v)
	}
}

func main() {
	//fmt.Println("hello golang")
	//var s = []int{11, 8, 99, 0, 2}
	//sort.Sort(sort.Reverse(sort.IntSlice(s)))
	//fmt.Println(s)
	showArray()
}
