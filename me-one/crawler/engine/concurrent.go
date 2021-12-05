package engine

import (
	"fmt"
	"github.com/gopl.io/me-one/crawler/fetch"
	"log"
)

type ConcurrentEngine struct {
	scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
}

func (e ConcurrentEngine) run(seeds ...Request) {
	for _, r := range seeds {
		e.scheduler.Submit(r)
	}

	in := make(chan Request)
	out := make(chan ParseResult)
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(in, out)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			fmt.Printf("Get item:::: %v", item)
		}

		for _, request := range result.Requests {
			e.scheduler.Submit(request)
		}
	}
}

func (e ConcurrentEngine) createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := e.worker(request)
			if err != nil {

			}
			out <- result
		}
	}()
}

func (ConcurrentEngine) worker(r Request) (ParseResult, error) {
	body, err := fetch.Fetcher(r.Url)
	if err != nil {
		log.Printf("Fetcher error : fetching url %s : %v", r.Url, err)
		return ParseResult{}, err
	}

	if r.Url == "https://www.zhenai.com/zhenghun/fuling" {
	}
	return r.Parser(body, r.Url), nil
}
