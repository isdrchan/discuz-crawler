package fetcher

import (
	"bufio"
	"dicuz-crawler/config"
	"dicuz-crawler/parser"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"log"
	"net/http"
	"strconv"
)

func Fetch(url string) (*goquery.Document, error) {
	client := &http.Client{CheckRedirect: redirect}
	req, err := http.NewRequest("GET", url, nil)
	for k, v := range config.Crawler.Header {
		req.Header.Add(k, v)
	}
	//log.Printf("请求url: %s", url)
	res, err := client.Do(req)
	if err != nil {
		log.Printf("请求错误: %s", err.Error())
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("status code: %d %s", res.StatusCode, res.Status)
		return nil, errors.New("status code: " + strconv.Itoa(res.StatusCode) + " " + res.Status)
	}
	utf8Reader := transform.NewReader(res.Body, DetermineEncoding(res.Body).NewDecoder())
	doc, err := goquery.NewDocumentFromReader(utf8Reader)
	if err != nil {
		log.Printf("解析dom出错: %s", err.Error())
		return nil, err
	}
	return doc, nil
}

func redirect(req *http.Request, via []*http.Request) (e error) {
	if req.URL.String() != parser.SeedUrlParse.Host {
		return errors.New("host不同源: " + req.URL.String())
	}
	return nil
}

func DetermineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("DetermineEncoding出错: %s", err.Error())
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
