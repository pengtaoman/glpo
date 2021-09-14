// Go program which illustrates
// recover in a goroutine
package main

import (
	"fmt"
	"time"
)

// For recovery
func handlepanic() {
	if a := recover(); a != nil {
		fmt.Println("RECOVER", a)
	}
}

/* Here, this panic is not handled by the recover function because of the recover function is not called in the same
goroutine in which the panic occurs */

// Function 1
func myfun1() {

	defer handlepanic()
	fmt.Println("Welcome to Function 1")

	panic("这个panic被recover，同一个goroutine里面 Panicked!!")

}

// Function 2
func myfun2() {

	fmt.Println("Welcome to Function 2")
	time.Sleep(3 * time.Second)
	println("####################################### 3秒之后")
	//这个panic没有被recover，因为不在同一个goroutine里面
	//panic("这个panic没有被recover，因为不在同一个goroutine里面 Panicked!!")
}

// Main function
func main() {
	// 这样func1异常后，被recover，不会影响fun2继续运行
	myfun1()
	myfun2()
	fmt.Println("Return successfully from the main function")
}
