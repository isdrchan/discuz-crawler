package parser

import (
	"dicuz-crawler/config"
	"dicuz-crawler/model"
	"github.com/PuerkitoBio/goquery"
)

func ParseList(doc *goquery.Document) model.ParseResult {
	parseResult := model.ParseResult{}
	doc.Find(config.Crawler.Selector.List).Each(func(i int, selection *goquery.Selection) {
		content, _ := selection.Html()
		url, _ := selection.Attr("href")
		//log.Printf("url: %s, content: %s", url, content)
		parseResult.Items = append(parseResult.Items, content)
		parseResult.Requests = append(parseResult.Requests, model.Request{
			Url:       url,
			ParseFunc: ParseContent,
		})
	})
	return parseResult
}
