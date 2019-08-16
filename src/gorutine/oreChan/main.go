package main

import (
	"fmt"
	"time"
)

func main() {
	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	fmt.Println(theMine, len(theMine))
	oreChan := make(chan string)
	mineOreChan := make(chan string)
	// Finder
	go func(mine [5]string) {
		for _, item := range mine {
			if item == "ore" {
				oreChan <- item
			}
		}
	}(theMine)
	// Ore Breaker
	go func() {
		for i := 0; i < 3; i++ {
			fundOre := <-oreChan
			fmt.Println("Form Finder", fundOre)
			mineOreChan <- "minedOre"
		}
	}()
	// Smelter
	go func() {
		for i := 0; i < 3; i++ {
			minedOre := <-mineOreChan //从 minedOreChan 读取
			fmt.Println("From Miner: ", minedOre)
			fmt.Println("From Smelter: Ore is smelted")
		}
	}()
	<-time.After(time.Second * 5)

}
