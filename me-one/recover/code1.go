package main

import (
	"fmt"
	"time"
)

func catchErr(num int) {
	// Try removing the following code of defer block
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[recover]", err, num)
		}
	}()

	fmt.Println("goroutine", num)
	//这里的panic 直接返回到方法内的defer中，被recover了
	panic("这里的panic 直接返回到方法内的defer中，被recover了 panic occurred ...")
}

func main() {

	//这个defer始终没有执行，
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("这个defer始终没有执行 main recover: ", err)
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println("每个goroutine打一个main goroutine", i)
		go catchErr(i)
		time.Sleep(time.Second * 1)
	}

start:
	goto start

}
