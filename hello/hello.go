package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hello golang")
	var s = []int{11, 8, 99, 0, 2}
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}
