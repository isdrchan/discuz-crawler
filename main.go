package main

import (
	"dicuz-crawler/config"
	"dicuz-crawler/engine"
	"dicuz-crawler/model"
	"dicuz-crawler/parser"
	"dicuz-crawler/persist"
)

func main() {
	e := engine.Simple{
		Saver: &persist.FileSaver{},
	}
	//e := engine.Concurrent{
	//	Saver: &persist.FileSaver{},
	//	WorkerCount: 1,
	//}
	e.Run(model.Request{
		Url:       config.Crawler.Seed.Url,
		ParseFunc: parser.StrToFuncOfParser(config.Crawler.Seed.Parser),
	})
}
