package main

import (
	"fmt"
	"sync"
	"time"
)

//type atomicInt int
//为了防止读写冲突，因此加锁
type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increament() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return int(a.value)
}

func main() {
	var a atomicInt
	a.increament()
	go func() {
		a.increament()
	}()
	time.Sleep(time.Millisecond * 100)
	fmt.Println(a.get())

}

/**
pengtaodeMacBook2019:atomic-demo pengtao$ go run -race atomic.go
==================
WARNING: DATA RACE
#####################在读的时候
Read at 0x00c00013c018 by main goroutine:
  main.main()
      /Users/pengtao/go-path/src/github.com/gopl.io/me-one/channel/atomic-demo/atomic.go:25 +0xee

#####################有两个地方正在写， 因此冲突
Previous write at 0x00c00013c018 by goroutine 7:
  main.(*atomicInt).increament()
      /Users/pengtao/go-path/src/github.com/gopl.io/me-one/channel/atomic-demo/atomic.go:11 +0x45
  main.main.func1()
      /Users/pengtao/go-path/src/github.com/gopl.io/me-one/channel/atomic-demo/atomic.go:22 +0x2e

Goroutine 7 (finished) created at:
  main.main()
      /Users/pengtao/go-path/src/github.com/gopl.io/me-one/channel/atomic-demo/atomic.go:21 +0xd8
==================
2
Found 1 data race(s)


**/
