package main

import "fmt"

func main() {
	//var m map[string]int
	m := make(map[string]int)
	m["tom"] = 100
	m["john"] = 120
	for i, v := range m {
		fmt.Println(i, v)
	}

	e, ok := m["tom"]
	fmt.Println(e, ok)
}
