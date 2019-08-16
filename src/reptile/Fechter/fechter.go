package main

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

// printCityList 获取所有城市的 a 链接

func printCityList (contents []byte)  {
	// ^ not
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	fmt.Printf("%s", matches)
	for _, m := range matches {
		fmt.Printf("City: %s , Url: %s\n", m[2], m[1])
	}
	fmt.Printf("Matches found: %d\n", len(matches))

}


func determineEncoding(r io.Reader) encoding.Encoding  {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func Fechter(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errText := "happened error, code:  " + string(resp.StatusCode)
		return nil, errors.New(errText)
	}

	e := determineEncoding(resp.Body)
	// 转换成使用 UTF-8 的 reader
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s/n", all)
	printCityList(all)
}

func main() {

}