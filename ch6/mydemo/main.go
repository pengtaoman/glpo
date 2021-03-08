package main

import (
	"fmt"
	"reflect"
)

type p struct {
	name string
	age int
	write func()
}

type MMap map[string]string

func (m MMap) Showkey() {
	if m["ali"] == "baba" {
		println("ali ======================================== baba")
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++", reflect.TypeOf(m))
	}
}

func (p *p) showInfo() {
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++", reflect.TypeOf(p))
	p.name = "alibabba"

	fmt.Fprintf()
}

func (p p) showIn() {
	println("==============================================")
	println(":::::::::::::::::::::::::::::" + p.name)
	p.name = "163163"
}

func main() {

	p := &p{name:"aa", age:23}
	p.showInfo()
	p.showIn()
	println("?????????????????" + p.name)

	var mmap = MMap{}
	mmap["asdf"] = "asdf"
	mmap["ali"] = "baba"
	mmap.Showkey()

}