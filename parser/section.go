package parser

import (
	"dicuz-crawler/config"
	"dicuz-crawler/model"
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

var IdRe0 = regexp.MustCompile(`/thread-([\d]+)-[\d]+-[\d]+.html`)
var IdRe1 = regexp.MustCompile(`tid=([\d]+)&`)

func ParseSection(doc *goquery.Document, item model.Item) model.ParseResult {
	parseResult := model.ParseResult{}
	doc.Find(config.Crawler.Selector.Title).Each(func(i int, selection *goquery.Selection) {
		content, _ := selection.Html()
		url, _ := selection.Attr("href")
		//log.Printf("url: %s, content: %s", url, content)
		url, _ = RelativeToAbsoluteOfUrl(url)
		item.Title = content
		item.Url = url
		match := IdRe0.FindSubmatch([]byte(url))
		var matchResult string
		if len(match) >= 2 {
			matchResult = string(match[1])
			item.Id = matchResult
		}
		if len(matchResult) == 0 {
			match = IdRe1.FindSubmatch([]byte(url))
			if len(match) >= 2 {
				matchResult = string(match[1])
				item.Id = matchResult
			} else {
				item.Id = "-"
			}
		}
		parseResult.Items = append(parseResult.Items, content)
		parseResult.Requests = append(parseResult.Requests, model.Request{
			Url:       url,
			ParseFunc: ParseArticle,
			Deliver:   item,
		})
	})
	count := 0
	doc.Find(config.Crawler.Selector.NextPage).Each(func(i int, selection *goquery.Selection) {
		if count > 0 {
			return
		}
		url, _ := selection.Attr("href")
		content := selection.Text()
		url, _ = RelativeToAbsoluteOfUrl(url)
		//log.Printf("url: %s, title: %s", url, content)
		parseResult.Items = append(parseResult.Items, content)
		parseResult.Requests = append(parseResult.Requests, model.Request{
			Url:       url,
			ParseFunc: ParseSection,
			Deliver:   item,
		})
		count++
	})
	return parseResult
}
