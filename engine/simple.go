package engine

import (
	"dicuz-crawler/fetcher"
	"dicuz-crawler/model"
	"log"
)

type Simple struct{}

func (e Simple) Run(seeds ...model.Request) {
	var requests []model.Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		parseResult, err := e.worker(request)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("url: %s, item: %s", request.Url, item)
		}
	}
}

func (e Simple) worker(request model.Request) (model.ParseResult, error) {
	doc, err := fetcher.Fetch(request.Url)
	if err != nil {
		return model.ParseResult{}, err
	}
	return request.ParseFunc(doc), nil
}
