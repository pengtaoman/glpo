package main

import (
	"fmt"
	"strconv"
)

type treeNode struct {
	value int
	left  *treeNode
}

func (tn *treeNode) setValue(v int) {
	tn.value = v
}

func main() {

	root := &treeNode{}
	root.setValue(233)
	fmt.Println(root.value)
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

	const (
		a = 1 << (10 * iota)
		b
		c
		d
	)
	fmt.Println(a, b, c, d)
	aaa := 23
	bb := strconv.Itoa(aaa)
	fmt.Println(bb)
	fmt.Print()

}
