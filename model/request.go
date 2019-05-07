package model

import "github.com/PuerkitoBio/goquery"

type Request struct {
	Url       string
	ParseFunc func(*goquery.Document, Item) ParseResult
	Deliver   Item
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}
