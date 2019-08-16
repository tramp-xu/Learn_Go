package main

import "fmt"

func tryDefer()  {
	// 在栈内 先进后出
	defer fmt.Println(1)
	fmt.Println(2)
}

func main() {
	tryDefer()
}
