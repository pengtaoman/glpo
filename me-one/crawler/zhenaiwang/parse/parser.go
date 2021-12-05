package parse

import (
	"fmt"
	"regexp"
)

func processReg(contents []byte) {

	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" data-v-1573aa7c>([^<]+)</a>`)

	match2 := re.FindAllSubmatch(contents, -1)
	for _, mm := range match2 {
		//fmt.Printf("%s\n", mm)
		fmt.Printf("%s   %s\n", mm[1], mm[2])
	}
	fmt.Printf("共有 %d 个城市", len(match2))
}

func main() {

}
