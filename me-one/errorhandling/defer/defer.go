package main

import (
	"bufio"
	"fmt"
	"github.com/gopl.io/me-two/fib"
	"os"
)

func tryDefer() {
	//defer 是栈结构 先进后出
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	//panic("error occurred !!")
	defer fmt.Println(4)
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 30; i++ {
		fmt.Fprintln(writer, f())
	}

}

func main() {
	//tryDefer()
	writeFile("fib111.txt")
}
