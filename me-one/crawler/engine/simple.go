package engine

import (
	"github.com/gopl.io/me-one/crawler/fetch"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {

	var request []Request
	for _, r := range seeds {
		request = append(request, r)
	}

	itemNum := 0
	for len(request) > 0 {
		r := request[0]
		request = request[1:]
		log.Printf("Fetching url  %s", r.Url)

		parseResult, err := e.worker(r)
		if err != nil {
			continue
		}

		//此处不断for循环， len(request) > 0
		request = append(request, parseResult.Requests...)

		for _, item := range parseResult.Items {
			itemNum++
			log.Printf("Got item  %d ====== %s", itemNum, item)
		}

	}

}
func (SimpleEngine) worker(r Request) (ParseResult, error) {
	body, err := fetch.Fetcher(r.Url)
	if err != nil {
		log.Printf("Fetcher error : fetching url %s : %v", r.Url, err)
		return ParseResult{}, err
	}

	if r.Url == "https://www.zhenai.com/zhenghun/fuling" {
	}
	return r.Parser(body, r.Url), nil
}
