package main

import (
	"fmt"
)

type iAdder3 func(int) int

func adder3() iAdder3 {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	a := adder3()
	var s int
	for i := 1; i <= 100; i++ {
		s = a(i)
	}
	fmt.Println(s)
}
