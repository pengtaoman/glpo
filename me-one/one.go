package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gopl.io/me-one/bysical"
	iface "github.com/gopl.io/me-one/iface"
)

// type Driver interface {
// 	Run() string
// }

func runDriver(dri iface.Driver) {
	fmt.Println(dri.Run(244))
}

func main() {
	bby := bysical.Bysical{Name: "BUSICAL"}
	runDriver(bby)
	defer fmt.Println("defer main") // will this be called when panic?
	var user = os.Getenv("USER_")
	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {
				fmt.Println("recover success.", err)
			}
		}()
		func() {
			defer func() {
				fmt.Println("defer here")
			}()

			if user == "" {
				panic("should set user env.")
			}
			fmt.Println("after panic")
		}()
	}()

	time.Sleep(1 * time.Second)
	fmt.Printf("get result %d\r\n", 1111)

}
