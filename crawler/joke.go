package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func GetJokes() {
	doc, err := goquery.NewDocument("http://www.qiushibaike.com")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".content").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}

func main() {
	GetJokes()
}
