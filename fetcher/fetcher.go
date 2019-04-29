package fetcher

import (
	"bufio"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"log"
	"net/http"
	"strconv"
)

func Fetch(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("请求错误: %s", err.Error())
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("status code : %d %s", res.StatusCode, res.Status)
		return nil, errors.New("status code : " + strconv.Itoa(res.StatusCode) + " " + res.Status)
	}
	utf8Reader := transform.NewReader(res.Body, DetermineEncoding(res.Body).NewDecoder())
	doc, err := goquery.NewDocumentFromReader(utf8Reader)
	if err != nil {
		log.Printf("解析dom出错: %s", err.Error())
		return nil, err
	}
	return doc, nil
}

func DetermineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
