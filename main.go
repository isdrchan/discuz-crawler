package main

import (
	"dicuz-crawler/engine"
	"dicuz-crawler/model"
	"dicuz-crawler/parser"
)

func main() {
	e := engine.Simple{}
	e.Run(model.Request{
		Url:       "https://bbs.kafan.cn/forum.php",
		ParseFunc: parser.ParseForum,
	})
}
