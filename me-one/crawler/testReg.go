package main

import (
	"fmt"
	"regexp"
)

func main() {

	url := " http://www.zhenai.com/zhenghun/yingkou/2"
	reurl := regexp.MustCompile(`(http[s]{0,1}://www.zhenai.com/zhenghun/[a-z]+)/([0-9]+)`)
	matchurl := reurl.FindAllSubmatch([]byte(url), -1)
	fmt.Printf("???????????????????????????? %s", matchurl)

}
