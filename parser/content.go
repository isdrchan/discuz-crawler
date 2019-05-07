package parser

import (
	"dicuz-crawler/config"
	"dicuz-crawler/model"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func ParseContent(doc *goquery.Document, item model.Item) model.ParseResult {
	parseResult := model.ParseResult{}
	doc.Find(config.Crawler.Selector.Content).Each(func(i int, selection *goquery.Selection) {
		content, _ := selection.Html()
		content = strings.Replace(content, "\n", "", -1)
		//log.Printf("content: %s", content)
		item.Content = content
		parseResult.Items = append(parseResult.Items, item)
	})
	return parseResult
}
