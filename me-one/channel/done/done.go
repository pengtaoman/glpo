package main

import (
	"fmt"
)

func doWork(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("!!!!!!! Worker 的 ID 是:::%d ,收到:::%d\n", id, n)
		go func() { done <- true }()
	}
}

type worker struct {
	c    chan int
	done chan bool
}

func createworker(id int) worker {
	w := worker{
		c:    make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.c, w.done)
	return w
}

func chandemo() {
	var worker [10]worker
	for i := 0; i < 10; i++ {
		worker[i] = createworker(i)
	}

	for i := 0; i < 10; i++ {
		worker[i].c <- 'a' + i
		//这里怎么能保证运行完呢，是因为channel是goroutine可能的切换点这个原理么？
		//<-worker[i].done
		//println(bb)
	}

	for i := 0; i < 10; i++ {
		worker[i].c <- 'A' + i
		//这里怎么能保证运行完呢，是因为channel是goroutine可能的切换点这个原理么？
		//<-worker[i].done
	}

	/** <-worker[i].done 放在每个循环里，实际上是每个循环顺序执行，不是并发，
		    所以把<-worker[i].done 放到最后
	        但是把两个<-worker[i].done 放在一起，在上面的doWork方法会引发死锁
		所有channel的发都是阻塞式的，发出去必须另一方有人收，
		    所以doWork 方法中的 done <- true 也需要改成go routine执行
	*/
	for _, work := range worker {
		<-work.done
		<-work.done
	}

}

func main() {

	chandemo()
}
