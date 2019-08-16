package main

import (
	"fmt"
	"github.com/gocolly/colly"
)
func main() {
	c := colly.NewCollector()
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %\n", e.Text, link)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r  *colly.Request) {
		fmt.Println("Visting", r.URL)
	})

	c.Visit("http://go-colly.org/")
}
