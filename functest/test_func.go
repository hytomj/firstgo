package main

import (
	"fmt"
	"math"
)

type Vars struct {
	i float64
	//j int
}

func (x Vars) Ceil() float64 {
	return math.Ceil(x.i)
}

func main() {
	f := Vars{10.05}
	fmt.Println(f.Ceil())
}
