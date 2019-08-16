package main

import (
	"fmt"
	"interface/duckTyping/mock"
	"interface/duckTyping/real"
)

// duck typing 描述事物的外部行为而非内部结构
type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return  r.Get("http://www.baidu.com/s?ie=UTF-8&wd=imooc")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake mook"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
