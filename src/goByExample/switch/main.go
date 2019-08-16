package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is the weekend")
	default:
		fmt.Println("it is a weekday")
	}
}
