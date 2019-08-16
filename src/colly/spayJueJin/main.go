package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"time"
)

func fetch(url string) *goquery.Document {
	fmt.Println("Fetch Url", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	req.Header.Set("X-Agent", "Juejin/Web")
	req.Header.Set("X-Legacy-Device-Id", "1562740326378")
	req.Header.Set("X-Legacy-Uid", "583295ce128fe100697c455f")
	req.Header.Set("X-Legacy-Token", "eyJhY2Nlc3NfdG9rZW4iOiJsUWZsd2tKQ2RKMml6OHR5IiwicmVmcmVzaF90b2tlbiI6Imp1RWJTVERoTVdRS1VFUzAiLCJ0b2tlbl90eXBlIjoibWFjIiwiZXhwaXJlX2luIjoyNTkyMDAwfQ==")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Http get err:", err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Http status code:", resp.StatusCode)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func parseUrls(url string) {
	doc := fetch(url)
	time.Sleep(10 * time.Second)
	doc.Find(".entry-list li").Each(func(index int, ele *goquery.Selection) {
		title := ele.Find(".title-row .title").Text()
		fmt.Println(title)
		// movieUrl, _ := ele.Find(".title-row a").Attr("href")
		//fmt.Println(strings.Split(movieUrl, "/")[4], ele.Find(".title").Eq(0).Text())
	})
	//time.Sleep(goquery * time.Second)
}

func main() {
	//start := time.Now()
	//ch := make(chan bool)
	parseUrls("https://juejin.im")
	//for i := 0; i < 10; i++ {
	//	go parseUrls("https://movie.douban.com/top250?start="+strconv.Itoa(25*i), ch)
	//}
	//
	//for i := 0; i < 10; i++ {
	//	<-ch
	//}
	//<-ch
	//elapsed := time.Since(start)
	//fmt.Printf("Took %s", elapsed)
}