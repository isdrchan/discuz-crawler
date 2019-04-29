package parser

import (
	"dicuz-crawler/model"
	"github.com/PuerkitoBio/goquery"
)

func ParseForum(doc *goquery.Document) model.ParseResult {
	parseResult := model.ParseResult{}
	doc.Find(".fl_row .fl_icn+td h2 a").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		content := selection.Text()
		//log.Printf("url: %s, title: %s", url, content)
		parseResult.Items = append(parseResult.Items, content)
		parseResult.Requests = append(parseResult.Requests, model.Request{
			Url:       url,
			ParseFunc: ParseList,
		})
	})
	return parseResult
}