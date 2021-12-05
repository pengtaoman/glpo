package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {

	for n := range c {
		time.Sleep(5 * time.Second) //模拟接收延迟的情况，消耗速度延迟
		fmt.Printf("Worker ID:%d recieved:%d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	/***
		w1 := createWorker(1001)
		w2 := createWorker(2002)
		for {
			select {
			case n := <-c1:
				//fmt.Println("从C1获取：", n)
				w1 <- n
			case n := <-c2:
				//fmt.Println("从C2获取：", n)
				w2 <- n
			}
		}
	    **/

	var worker = createWorker(1111)

	//延迟接收消耗的情况下，需要把收到的存下来
	var values []int

	tm := time.After(10 * time.Second)
	//n := 0
	//hasValue := false //c1 c2是否已经产生值的判断
	for {
		var activeWorker chan<- int
		//if hasValue {
		//	activeWorker = worker
		//}
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			//hasValue = true
			//延迟接收情况下，用values缓存消息
			values = append(values, n)
			fmt.Printf("recieve FROM C1 :%d\n", n)
		case n := <-c2:
			//hasValue = true
			values = append(values, n)
			fmt.Printf("recieve FROM C2 :%d\n", n)
		case activeWorker <- activeValue: //如果activeWorker是nil，那么这个case不会走到
			//hasValue = false
			values = values[1:]
			fmt.Println("从c1 和 c2收到数据，并传给worker！！已缓存数据长度：：：", len(values))
		case <-tm:
			fmt.Println("ByeBye!!!")
			return
		}
	}
}
