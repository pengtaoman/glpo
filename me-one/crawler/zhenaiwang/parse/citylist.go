package parse

import (
	"github.com/gopl.io/me-one/crawler/engine"
	"log"
	"regexp"
	"strings"
)

func ParseCityList(contents []byte, url string) engine.ParseResult {

	log.Print("$$$$$$$$$$$$$$$$$$$$ ParseCityList URL : %s", url)
	//re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" data-v-1573aa7c>([^<]+)</a>`)
	re := regexp.MustCompile(`{"linkContent":"([^\"]+)","linkURL":"(http:\\u002F\\u002Fwww.zhenai.com\\u002Fzhenghun\\u002F[a-z0-9]+)"}`)
	match2 := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, mm := range match2 {
		url := strings.Replace(string(mm[2]), "\\u002F", "/", -1)
		//fmt.Printf("%s ---- %s\n", mm[1], url)
		result.Items = append(result.Items, mm[1])
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(url),
			Parser: ParseCity,
		})
	}
	return result
}
