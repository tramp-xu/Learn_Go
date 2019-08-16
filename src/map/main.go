package main

import "fmt"

func main() {
	m1 := make(map[string] int)
	m2 := map[string] string {
		"name": "cc",
		"course": "golang",
		"site": "icc",
		"qul": "not bad",
	}
	fmt.Println(m1)
	fmt.Println(m2)
}
