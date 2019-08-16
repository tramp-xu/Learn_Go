package main

import (
	"fmt"
	"regexp"
)

const text = `
	My email is ccmouse@gamil.com@abc.com
	email is abc@.com
	email2 is kkd@qq.com
`

func main() {
	// 正则中``中不需要转义\
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)(\.[a-zA-Z0-9]+)`)
	// -1 代表 所有
	//match := re.FindAllString(text, -1)
	// 根据括号返回所需要的值
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
}
