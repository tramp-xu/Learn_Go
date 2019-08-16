package main

import "fmt"

func f (from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	// 命令行 输入值
	fmt.Scanln(&input)
	fmt.Println("输入值是: ", input)
	fmt.Println("done")
}
