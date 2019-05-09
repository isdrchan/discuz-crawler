package engine

import (
	"dicuz-crawler/fetcher"
	"dicuz-crawler/model"
	"dicuz-crawler/persist"
	"log"
)

type Simple struct {
	Saver persist.Saver
}

func (e Simple) Run(seeds ...model.Request) {
	var requests []model.Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	e.Saver.Init()

	count := 0
	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		parseResult, err := e.worker(request)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("#%d-item: %+v", count, item)
			item, ok := item.(model.Item)
			if ok {
				err := e.Saver.Save(item)
				log.Printf("数据 %v 保存出错: %s", item, err)
			}
			count++
		}
	}

	e.Saver.Close()
}

func (e Simple) worker(request model.Request) (model.ParseResult, error) {
	doc, err := fetcher.Fetch(request.Url)
	if err != nil {
		return model.ParseResult{}, err
	}
	return request.ParseFunc(doc, request.Deliver), nil
}
