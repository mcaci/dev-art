package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 20)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one again"
	}()
	go func() {
		time.Sleep(20 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			go func() {
				fmt.Println("received", msg1)
				time.Sleep(5 * time.Second)
				fmt.Println(msg1, "has slept")
			}()
		case msg2 := <-c2:
			fmt.Println("received", msg2)
			// default:
			// 	fmt.Println("dunno")
		}
	}
}
