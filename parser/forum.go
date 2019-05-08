package parser

import (
	"dicuz-crawler/config"
	"dicuz-crawler/model"
	"github.com/PuerkitoBio/goquery"
)

func ParseForum(doc *goquery.Document, _ model.Item) model.ParseResult {
	parseResult := model.ParseResult{}
	doc.Find(config.Crawler.Selector.Section).Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		content := selection.Text()
		//log.Printf("url: %s, title: %s", url, content)
		parseResult.Items = append(parseResult.Items, content)
		parseResult.Requests = append(parseResult.Requests, model.Request{
			Url:       url,
			ParseFunc: ParseSection,
			Deliver: model.Item{
				Section: content,
			},
		})
	})
	doc.Find(config.Crawler.Selector.SubSection).Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		content := selection.Text()
		//log.Printf("url: %s, title: %s", url, content)
		parseResult.Items = append(parseResult.Items, content)
		parseResult.Requests = append(parseResult.Requests, model.Request{
			Url:       url,
			ParseFunc: ParseSection,
			Deliver: model.Item{
				Section: content,
			},
		})
	})
	return parseResult
}
