package engine

import (
	"dicuz-crawler/fetcher"
	"dicuz-crawler/model"
	"dicuz-crawler/persist"
	"log"
)

type Concurrent struct {
	Saver       persist.Saver
	WorkerCount int
}

func (e *Concurrent) Run(seeds ...model.Request) {
	in := make(chan model.Request)
	out := make(chan model.ParseResult)
	e.Saver.Init()

	workerNum := 0
	for i := 0; i < e.WorkerCount; i++ {
		workerNum++
		createWorker(in, out)
	}

	for _, request := range seeds {
		go func() { in <- request }()
	}

	count := 0
	for {
		result := <-out
		//log.Printf("result:%v", result)
		for _, request := range result.Requests {
			log.Printf("request:%v", request)
			go func() { in <- request }()
		}
		for _, item := range result.Items {
			log.Printf("#%d-item: %+v", count, item)
			item, ok := item.(model.Item)
			if ok {
				err := e.Saver.Save(item)
				if err != nil {
					log.Printf("数据 %v 保存出错: %s", item, err)
				}
			}
			count++
		}
	}

	e.Saver.Close()
}

func Worker(request model.Request) (model.ParseResult, error) {
	doc, err := fetcher.Fetch(request.Url)
	if err != nil {
		return model.ParseResult{}, err
	}
	return request.ParseFunc(doc, request.Deliver), nil
}

func createWorker(in chan model.Request, out chan model.ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
