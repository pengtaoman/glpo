package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 4; i++ {
		//协程 非抢占式的多任务处理。由协程主动交出控制权
		go func(i int) {
			for {
				fmt.Printf("Hello from goroutine 编号：：：： %d \n", i)
			}
		}(i)
	}

	//如果不sleep，则main方法直接执行退出，不会给goroutine运行的机会
	//main退掉，goroutine全部杀掉
	time.Sleep(time.Second)
}
