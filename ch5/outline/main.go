// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 123.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

type ss struct {
	aa string
	bb int
}

//!+
func main() {
	fmt.Println(strings.Map(add1, "HAL-9000"))

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func init() {
	log.SetPrefix("HTML :: ")

}

func add1(r rune)(rune) {return r + 1}

func outline(stack []string, n *html.Node) {

	println("::::::::::::::::::::::::::" + n.Data )

	log.Printf("asdfasdfasdfasd")
	//log.Fatalf("asdfasdf rgfkrog : ", n)
	if n.Type == html.ElementNode {


		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

//!-
