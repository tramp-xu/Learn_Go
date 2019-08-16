package main

import (
	"fmt"
	"time"
)

func calcSquares(number int, squareop chan int)  {
	sum := 0
	for number != 0  {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	time.Sleep(4 * time.Second)
	squareop <- sum
}

func calcCubes(number int, cubeop chan int)  {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares := <-sqrch
	fmt.Printf("square: %d\n", squares)
	cubes := <-cubech
	fmt.Println("Final output ", squares + cubes)
}
