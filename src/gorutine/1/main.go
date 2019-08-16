package main

import "fmt"
// gorutine 是 协程
func main() {
	for i :=0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from gorutine %d", i)
			}
		}(i)
	}
}
