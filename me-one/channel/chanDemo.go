package main

import (
	"fmt"
	"sync"
)

func worker(id int, c chan int, group *sync.WaitGroup) {
	for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("!!!!!!! Worker 的 ID 是:::%d ,收到:::%d\n", id, n)
		group.Done()
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go func() {
		for {
			//n := <-c
			fmt.Printf("!!!!!!! 创建的 Worker 的 ID 是:::%d ,收到:::%d\n", id, <-c)
		}
	}()
	return c
}

func createWorkerRec(id int, group *sync.WaitGroup) chan<- int {
	c := make(chan int)
	go func() {
		for {
			//n := <-c
			fmt.Printf("!!!!!!! createWorkerRec 创建的 Worker 的 ID 是:::%d ,收到:::%d\n", id, <-c)
			group.Done()
		}
	}()
	//runtime.Gosched()
	return c
}

func bufferChan(group *sync.WaitGroup) {
	c := make(chan int, 30)
	go worker(555555, c, group)

	for i := 0; i < 10; i++ {
		c <- i
	}
}

func closeChan(group *sync.WaitGroup) {
	c := make(chan int, 30)
	go worker(333333, c, group)

	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func chanDemo(wg *sync.WaitGroup) {

	var chans [10]chan<- int
	//for i := 0; i < 10; i++ {
	//	chans[i] = make(chan int)
	//	go worker(100+i,chans[i])
	//}

	for i := 0; i < 10; i++ {
		chans[i] = createWorkerRec(700+i, wg)

	}
	for i := 0; i < 10; i++ {
		chans[i] <- 3000 + i
	}

	for i := 0; i < 10; i++ {
		chans[i] <- 9000 + i
	}

}

func main() {

	var group sync.WaitGroup
	group.Add(40) //40 是写入channel的次数
	fmt.Println("channel 是一等公民")
	chanDemo(&group)
	//runtime.Gosched()

	fmt.Println("channel 缓存")
	bufferChan(&group)

	//runtime.Gosched()

	fmt.Println("channel 关闭")
	closeChan(&group)

	group.Wait()
	//runtime.Gosched()
	//cores := runtime.GOMAXPROCS(-1)
	//println("CORES !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!", cores)
	//time.Sleep(time.Millisecond)
	//
}
