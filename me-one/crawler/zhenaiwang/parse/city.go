package parse

import (
	"encoding/json"
	"fmt"
	"github.com/gopl.io/me-one/crawler/engine"
	"log"
	"regexp"
	"strconv"
)

func ParseCity(contents []byte, url string) engine.ParseResult {
	log.Printf("$$$$$$$$$$$$$$$$$$$$ ParseCity URL : %s", url)
	//re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" data-v-1573aa7c>([^<]+)</a>`)
	re := regexp.MustCompile(`window.__INITIAL_STATE__\=([^;]+);\(function\(\){var s;\(s=document.currentScript`)
	match2 := re.FindAllSubmatch(contents, -1)
	if len(match2) <= 0 {
		log.Printf("ERROR ERROR ERROR ERROR : %s", url)
		return engine.ParseResult{}
	}
	//fmt.Printf("!!!!!!!!!!!!!!!!!!!!!!! match2 :::: %d \n\n", match2)
	//fmt.Printf("!!!!!!!!!!!!!!!!!!!!!!! find num :::: %d \n\n", len(match2[0]))
	//fmt.Println("\n\n\n")
	//fmt.Printf("!!!!!!!!!!!!!!!!!!!!!!! city member :::: %s", match2[0])
	//fmt.Println("\n\n\n")

	var mem engine.Member
	err1 := json.Unmarshal(match2[0][1], &mem)
	if err1 != nil {
		log.Fatalln(err1)
	}

	//ioutil.WriteFile("000001.txt", match2[0][1], 0644)
	result := engine.ParseResult{}

	reurl := regexp.MustCompile(`(http[s]{0,1}://www.zhenai.com/zhenghun/[a-z]+)/([0-9]+)`)
	matchurl := reurl.FindAllSubmatch([]byte(url), -1)
	pageNum := 1
	requesturl := ""
	if len(matchurl) <= 0 {
		//fmt.Printf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!没有页面索引 \n\n\n", requesturl)
		pageNum = 2
		requesturl = url + "/" + strconv.Itoa(pageNum)
	} else {
		inx, _ := strconv.Atoi(string(matchurl[0][2]))
		pageNum = inx + 1
		requesturl = string(matchurl[0][1]) + "/" + strconv.Itoa(pageNum)
		//fmt.Printf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!%s \n\n\n" , matchurl)
	}

	fmt.Printf("################################# Next Page Request URL : %s \n", requesturl)
	for _, mm := range mem.MemberListData.MemberList {
		//fmt.Printf("########################### member: %v \n", mm)
		result.Items = append(result.Items, mm)
	}

	result.Requests = append(result.Requests, engine.Request{
		Url:    string(requesturl),
		Parser: ParseCity,
	})
	//for _, mm := range match2 {
	//	url := strings.Replace(string(mm[2]), "\\u002F", "/", -1)
	//	//fmt.Printf("%s ---- %s\n", mm[1], url)
	//	result.Items = append(result.Items, mm[1])
	//	result.Requests = append(result.Requests, engine.Request{
	//		Url: string(url),
	//		Parser: engine.NilParser,
	//	})
	//}
	return result
}
