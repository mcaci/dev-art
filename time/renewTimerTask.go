package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// renew task
	d := time.Second
	tmr := time.NewTimer(d)
	done := make(chan struct{})

	for {
		go doTask(done)

		select {
		case <-tmr.C:
			fmt.Println("timeout: took more than", d)
			return
		case <-done:
			fmt.Println("done")
			if !tmr.Stop() {
				<-tmr.C
			}
			tmr.Reset(d)
		}
	}
}

func doTask(done chan<- struct{}) {
	d := rand.Intn(1500)
	dr := time.Duration(d) * time.Millisecond
	fmt.Println("working for", dr)
	time.Sleep(dr)
	done <- struct{}{}
}

