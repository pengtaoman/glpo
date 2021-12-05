package main

import (
	"encoding/json"
	"fmt"
	"github.com/gopl.io/me-one/crawler/engine"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, error := os.Open("me-one/crawler/finejson.json")
	defer file.Close()
	if error != nil {
		log.Fatalln(error)
	}

	//var buf bytes.Buffer
	all, err := ioutil.ReadAll(file)
	//io.Copy(&buf, file)
	if err != nil {
		log.Fatalln(err)
	}
	//print(string(all))
	//b := buf.Bytes()
	//rarr := []rune{}
	//for len(b) > 0 {
	//	r, size := utf8.DecodeRune(b)
	//	rarr = append(rarr, r)
	//	b = b[size:]
	//}

	//ioutil.WriteFile("000001.txt", []byte(string(all)), 0644)

	var mem engine.Member
	err1 := json.Unmarshal(all, &mem)
	if err1 != nil {
		log.Fatalln(err1)
	}

	//mem := engine.Member{}
	fmt.Printf("############################## %d", len(mem.MemberListData.MemberList))

	//io.ByteWriter()
	//io.w(atom.String(buf.Bytes()))

}
