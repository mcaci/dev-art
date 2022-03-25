package main

import (
	"fmt"
	"time"
)

func main() {
	// sleeping
	fmt.Println("sleeping for one second")
	time.Sleep(time.Second)
	fmt.Println("sleeping done")
	
	// sleeping with After
	fmt.Println("waiting one second")
	t := <-time.After(time.Second)
	fmt.Println("waiting done", t)
	
	// schedule task with AfterFunc
	wait := make(chan struct{})
	time.AfterFunc(time.Second, func() { fmt.Println("work in progress"); wait <- struct{}{} })
	fmt.Println("waiting one more second")	
	<-wait
	fmt.Println("work done")
	
	// cancel scheduled task with AfterFunc
	tmr := time.AfterFunc(time.Second, func() { fmt.Println("work in progress"); wait <- struct{}{}; fmt.Println("work done") })
	fmt.Println("waiting one more more second")	
	time.AfterFunc(500*time.Millisecond, func() { tmr.Stop(); fmt.Println("work timer stopped"); wait <- struct{}{} })
	<-wait

	// sequence of scheduled tasks
	wait = make(chan struct{}, 2)
	time.AfterFunc(200*time.Millisecond, func() { fmt.Println("first"); wait <- struct{}{} })
	time.AfterFunc(500*time.Millisecond, func() { fmt.Println("second"); wait <- struct{}{} })
	for i := 0; i < cap(wait); i++ {
		<-wait
	}

	// schedule task with NewTicker
	tckr := time.NewTicker(300 * time.Millisecond)
	defer tckr.Stop()
	tmr = time.NewTimer(5*time.Second)
	for {
		select {
		case t := <-tmr.C:
			fmt.Println("done", t)
			return
		case t := <-tckr.C:
			fmt.Println(t)
		}

	}

	// timeout running program
	for range time.Tick(time.Second) {
		fmt.Println("Hello, playground")
	}

	// renew task with After
	d := time.Second
	tmr = time.NewTimer(d)
	var i int

	for {
		i += 150
		dr := time.Duration(i) * time.Millisecond
		select {
		case <-tmr.C:
			fmt.Println("timeout: took more than", d)
			return
		case <-time.After(dr):
			fmt.Println("task done after", dr)
			tmr.Reset(d)
		}
	}

	
}
