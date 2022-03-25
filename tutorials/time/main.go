package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// time.Sleep
	fmt.Println("Task 1: with time.Sleep")
	time.Sleep(time.Second)
	do()

	// time.After
	fmt.Println("Task 2: with time.After")
	now := time.Now()
	t := <-time.After(time.Second)
	do()
	fmt.Println("Duration of Task 2", t.Sub(now))

	// time.AfterFunc
	fmt.Println("Task 3: with time.AfterFunc")
	done := make(chan struct{})
	// time.AfterFunc(time.Second, do)
	time.AfterFunc(time.Second, func() { doAndNotify(done) })
	<-done

	// time.NewTimer
	fmt.Println("Task 4: concurrent task with time.NewTimer")
	tmr := time.NewTimer(time.Second)
	<-tmr.C
	done = make(chan struct{})
	go doAndNotify(done)
	<-done

	// time.NewTimer seq
	fmt.Println("Task 5: sequential task with time.NewTimer")
	tmr = time.NewTimer(time.Second)
	<-tmr.C
	do()

	// timer's time.AfterFunc
	fmt.Println("Task 6: time.AfterFunc's timer")
	tmr = time.AfterFunc(time.Second, do)
	select {
	case t := <-tmr.C:
		fmt.Println("ended successfully at", t)
	case <-time.After(2 * time.Second):
		fmt.Println("timer abandoned")
	}
	time.Sleep(2 * time.Second)

	// cancel time.AfterFunc
	fmt.Println("Task 7: stop time.AfterFunc's timer")
	done = make(chan struct{})
	tmr = time.AfterFunc(2*time.Second, func() { doAndNotify(done) })
	select {
	case <-done:
		fmt.Println("ended successfully")
	case <-time.After(time.Second):
		tmr.Stop()
		fmt.Println("timer abandoned")
	}

	// sequence of scheduled tasks
	fmt.Println("Task 8: schedule different tasks in different times")
	wg := sync.WaitGroup{}
	wg.Add(2)
	time.AfterFunc(200*time.Millisecond, func() { fmt.Println("first"); wg.Done() })
	time.AfterFunc(500*time.Millisecond, func() { fmt.Println("second"); wg.Done() })
	wg.Wait()

	// schedule task with NewTicker
	fmt.Println("Task 9: schedule repetitive tasks with time.Ticker")
	tckr := time.NewTicker(300 * time.Millisecond)
	defer tckr.Stop()
	doneTmr := time.NewTimer(2 * time.Second)
	now = time.Now()
scheduler:
	for {
		select {
		case t := <-doneTmr.C:
			fmt.Println("Stopping after", t.Sub(now))
			break scheduler
		case t := <-tckr.C:
			do()
			fmt.Println("task done after", t.Sub(now))
		}
	}

	// renew task with Timer and After
	fmt.Println("Task 10: schedule repetitive tasks with time.Timer and time.Ticker renewal")
	d := time.Second
	tmr = time.NewTimer(d)
	defer tmr.Stop()
	var i int
tmr_renew:
	for {
		i += 150
		dr := time.Duration(i) * time.Millisecond
		now = time.Now()
		select {
		case t := <-tmr.C:
			fmt.Println("timeout: took more than", d, ":", t.Sub(now))
			break tmr_renew
		case t := <-time.After(dr):
			do()
			fmt.Println("task done after", dr, ":", t.Sub(now))
			if !tmr.Stop() {
				<-tmr.C
			}
			tmr.Reset(d)
		}
	}

	tckr = time.NewTicker(d)
	defer tckr.Stop()
	i = 0
tckr_renew:
	for {
		i += 150
		dr := time.Duration(i) * time.Millisecond
		now = time.Now()
		select {
		case t := <-tckr.C:
			fmt.Println("timeout: took more than", d, ":", t.Sub(now))
			break tckr_renew
		case t := <-time.After(dr):
			do()
			fmt.Println("task done after", dr, ":", t.Sub(now))
			tckr.Reset(d)
		}
	}
}

func do() { fmt.Println("done") }
func doAndNotify(done chan<- struct{}) {
	fmt.Println("done")
	if done != nil {
		done <- struct{}{}
	}
}
