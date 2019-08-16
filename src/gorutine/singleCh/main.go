package main

import (
	"fmt"
	"time"
)

func sendData(sendch chan  int)  {
	sendch <- 10
	time.Sleep(1 * time.Second)
}

func main() {
	cha1 := make(chan int)
	go sendData(cha1)
	fmt.Println(<-cha1)
}
