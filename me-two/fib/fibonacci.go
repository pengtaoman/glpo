package fib

import "strconv"

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		println("???????????????? " + strconv.Itoa(a))
		a, b = b, a+b
		return a
	}
}

func main() {

}
