package main

import (
	"dicuz-crawler/engine"
	"dicuz-crawler/model"
	"dicuz-crawler/parser"
	"dicuz-crawler/persist"
)

func main() {
	e := engine.Simple{
		Saver: &persist.FileSaver{},
	}
	e.Run(model.Request{
		Url:       "https://bbs.kafan.cn/forum.php",
		ParseFunc: parser.ParseForum,
	})
}
