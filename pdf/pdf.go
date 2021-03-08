package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {

   files, _ := ioutil.ReadDir("/Volumes/860EVO/Downloads/天宫系列初级培训认证课件")
   for _, fi := range files {
   	  println(fi.Name())
   	  println(fi.IsDir())
	}
    println("=============================================================")

	filess, _ := filepath.Glob("*")
	fmt.Println(filess) //
	
}
