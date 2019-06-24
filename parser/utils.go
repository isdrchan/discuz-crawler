package parser

import (
	"dicuz-crawler/config"
	"dicuz-crawler/model"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
)

var seedUrl string
var SeedUrlParse *url.URL

func init() {
	seedUrl = config.Crawler.Seed.Url
	SeedUrlParse, _ = url.Parse(seedUrl)
}

func RelativeToAbsoluteOfUrl(relativeUrl string) (absoluteUrl string, err error) {
	u, err := url.Parse(relativeUrl)
	if err != nil {
		return relativeUrl, errors.New("relativeUrl: " + err.Error())
	}
	if u.Scheme == "http" || u.Scheme == "https" {
		return relativeUrl, errors.New("unknown scheme")
	}
	return fmt.Sprintf("%s://%s/%s", SeedUrlParse.Scheme, SeedUrlParse.Host, relativeUrl), nil
}

func StrToFuncOfParser(parserStr string) func(*goquery.Document, model.Item) model.ParseResult {
	switch parserStr {
	case "forum":
		return ParseForum
	case "section":
		return ParseSection
	case "article":
		return ParseArticle
	default:
		return ParseForum
	}
}
