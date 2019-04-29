package parser

import (
	"dicuz-crawler/model"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func ParseContent(doc *goquery.Document) model.ParseResult {
	parseResult := model.ParseResult{}
	doc.Find("div[id^=post_]:nth-child(2) td[id^=postmessage_]").Each(func(i int, selection *goquery.Selection) {
		content, _ := selection.Html()
		content = strings.Replace(content, "\n", "", -1)
		//log.Printf("content: %s", content)
		parseResult.Items = append(parseResult.Items, model.Article{
			Content: content,
		})
	})
	return parseResult
}
