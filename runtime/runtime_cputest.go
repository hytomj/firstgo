package main

import (
	"fmt"
	"runtime"
)

func main() {
	cupNum := runtime.NumCPU()
	fmt.Println(cupNum)

	s := runtime.GOMAXPROCS(cupNum - 5)
	fmt.Println(s)
}
