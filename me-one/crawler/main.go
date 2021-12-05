package main

import (
	"github.com/gopl.io/me-one/crawler/engine"
	"github.com/gopl.io/me-one/crawler/zhenaiwang/parse"
)

/**************
$('.city-list').children[1].children[1].innerText



*/

//const text = "My email is pengtaoman@163.com"
//
//const text1 = `My email is pengtaoman@163.com
//email 2 is pengtao1000@gmail.com
//
//email 3 is pengtaobatman@gmail.com
//email 4  is pengtaoironman@ggg.com.cn@gmail.com`
//
//func reg() {
//
//	re := regexp.MustCompile(`[a-zA-Z0-9]+@.+\..+`)
//	match := re.FindString(text)
//	fmt.Println(match)
//
//	fmt.Println("==========================================================")
//
//	re1 := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\..+`)
//	match1 := re1.FindAllString(text1, -1)
//	fmt.Println(match1)
//
//	fmt.Println("==========================================================")
//
//	re2 := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.(.+)`)
//	match2 := re2.FindAllStringSubmatch(text1, -1)
//	fmt.Println(match2)
//	for _, mm := range match2 {
//		fmt.Println(mm)
//	}
//}
//
//func processReg(contents []byte) {
//
//	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" data-v-1573aa7c>([^<]+)</a>`)
//
//	match2 := re.FindAllSubmatch(contents, -1)
//	for _, mm := range match2 {
//		//fmt.Printf("%s\n", mm)
//		fmt.Printf("%s   %s\n", mm[1], mm[2])
//	}
//	fmt.Printf("共有 %d 个城市", len(match2))
//}

func main() {
	engine.SimpleEngine{}.Run(engine.Request{
		Url:    "https://www.zhenai.com/zhenghun",
		Parser: parse.ParseCityList,
	})

	//engine.Run(engine.Request{
	//	Url: "https://www.zhenai.com/zhenghun/aba",
	//	Parser: parse.ParseCity,
	//})
}
