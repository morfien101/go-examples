package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var (
	workers      = flag.Int("w", 4, "How many workers")
	jobs         = flag.Int("j", 100, "How many jobs")
	sleepEnabled = flag.Bool("s", false, "Should we sleep to slow down the demo")
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	flag.Parse()

	workq := make(chan int, *workers)
	logq := make(chan string, 1000)

	workersWg := new(sync.WaitGroup)
	loggerWg := new(sync.WaitGroup)

	go logWorker(logq, loggerWg)

	for i := 1; i <= *workers; i++ {
		go work(i, workq, logq, workersWg)
	}

	for i := 0; i < *jobs; i++ {
		workq <- i + 1
	}

	close(workq)
	workersWg.Wait()
	close(logq)
	loggerWg.Wait()
	fmt.Println("Finished!")

}

func work(id int, num chan int, logq chan string, wg *sync.WaitGroup) {
	wg.Add(1)
	for {
		select {
		case job, ok := <-num:
			if !ok {
				wg.Done()
				return
			}
			sleeper()
			logq <- fmt.Sprintf("Worker: %d, Job: %d", id, job)
		}
	}
}

func logWorker(logq chan string, wg *sync.WaitGroup) {
	wg.Add(1)
	for {
		select {
		case msg, ok := <-logq:
			if !ok {
				wg.Done()
				return
			}
			fmt.Println(msg)
		}
	}
}

func sleeper() {
	if *sleepEnabled {
		time.Sleep(time.Millisecond * time.Duration(rand.Int63n(100)))
	}
}
