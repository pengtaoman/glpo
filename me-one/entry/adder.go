package main

import (
	"fmt"
	"strconv"
)

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

/******************************************************/
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		println("!!!!!!!!!!!!! " + strconv.Itoa(v))
		return base + v, adder2(base + v)
	}
}

/******************************************************/
func main() {
	add := adder()
	add2 := adder2(0)
	fmt.Printf("Type:::: %T, value:::: %v", add, add)
	for i := 0; i < 10; i++ {
		var vv int
		vv, add2 = add2(i)
		fmt.Printf("0 + 1 + ... %d = %d \n", i, vv)
	}
}
