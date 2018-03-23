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
	length        = flag.Int64("l", 1024, "How many bytes that you want to count.")
	numberOflists = flag.Int64("n", 2, "How many lists to run through")
	verbose       = flag.Bool("verbose", false, "See the tables.")
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	flag.Parse()

	// We craft a thing to count
	countMe := makeSlice(*length)

	// single Threaded
	t := time.Now()
	table := countSingleThreaded(countMe)
	endTime := time.Since(t) / time.Millisecond
	fmt.Printf("Single Threaded: %d ms\n", endTime)
	printTable(table)

	// Multi Threaded test
	t = time.Now()
	table = countMultiThreaded(countMe)
	endTime = time.Since(t) / time.Millisecond
	fmt.Printf("Multi Threaded: %d ms\n", endTime)
	printTable(table)

	// Ah what why is it slower?
	// let's try this again by with multiple lists.
	lists := make([][]byte, *numberOflists)
	var counter int64
	for counter = 0; counter < *numberOflists; counter++ {
		lists[counter] = makeSlice(*length)
	}

	// single threaded test
	t = time.Now()
	listToMerge := make([]map[byte]int64, len(lists))
	for index, value := range lists {
		listToMerge[index] = countSingleThreaded(value)
	}

	completeList := make(map[byte]int64)
	for _, v := range listToMerge {
		for key, value := range v {
			completeList[key] = completeList[key] + value
		}
	}
	endTime = time.Since(t) / time.Millisecond
	fmt.Printf("Single Threaded: %d ms\n", endTime)
	printTable(completeList)

	// multi threaded test
	t = time.Now()
	completeList = countParrallel(lists)
	endTime = time.Since(t) / time.Millisecond
	fmt.Printf("Multi Threaded: %d ms\n", endTime)
	printTable(completeList)
}

func makeSlice(length int64) []byte {
	s := make([]byte, length)

	for i := 0; i < len(s); i++ {
		s[i] = byte(rand.Intn(26) + 65)
	}

	return s
}

func countSingleThreaded(s []byte) map[byte]int64 {
	countTable := make(map[byte]int64)
	for _, v := range s {
		countTable[v]++
	}

	return countTable
}

func printTable(table map[byte]int64) {
	if *verbose {
		for i := 0; i < 26; i++ {
			key := byte(65 + i)
			fmt.Printf("%s: %d\n", string(key), table[key])
		}
	}
}

func countMultiThreaded(byteSlice []byte) map[byte]int64 {
	// as many workers as there is CPU -1 (one for os and controllers)
	workers := runtime.NumCPU()
	wg := new(sync.WaitGroup)
	byteChan := make(chan byte, 10)
	mapChan := make(chan map[byte]int64, workers)

	// Start the workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go multiCounterfunc(wg, byteChan, mapChan)
	}

	// feed the workers
	for _, v := range byteSlice {
		byteChan <- v
	}
	// Close the channel to signal there is no more work
	close(byteChan)
	// Wait for the workers to finish
	wg.Wait()
	// Close the results channel
	close(mapChan)
	// We need to collect the results and merge them
	// make a new table and push in the results.
	mergeTable := make(map[byte]int64)
	for wTable := range mapChan {
		for k, v := range wTable {
			mergeTable[k] = mergeTable[k] + v
		}
	}
	return mergeTable
}

func multiCounterfunc(
	wg *sync.WaitGroup,
	byteChan chan byte,
	mapChan chan map[byte]int64,
) {
	defer wg.Done()
	workerTable := make(map[byte]int64)
	for {
		select {
		case v, ok := <-byteChan:
			if !ok {
				// When we return the defer statement kicks in.
				mapChan <- workerTable
				return
			}
			workerTable[v]++
		}
	}
}

func countParrallel(lists [][]byte) map[byte]int64 {
	workers := len(lists)
	wg := new(sync.WaitGroup)
	mapChan := make(chan map[byte]int64, workers)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go parrallelCounterfunc(wg, lists[i], mapChan)
	}
	wg.Wait()
	close(mapChan)
	// We need to collect the results and merge them
	// make a new table and push in the results.
	mergeTable := make(map[byte]int64)
	for wTable := range mapChan {
		for k, v := range wTable {
			mergeTable[k] = mergeTable[k] + v
		}
	}
	return mergeTable
}

func parrallelCounterfunc(
	wg *sync.WaitGroup,
	list []byte,
	mapChan chan map[byte]int64,
) {
	defer wg.Done()
	workerTable := make(map[byte]int64)
	for _, v := range list {
		workerTable[v]++
	}
	mapChan <- workerTable
}
