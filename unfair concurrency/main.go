package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const maxLimit = 100000000

var batches int = 10
var totalNumberofPrimes int32 = 0

func checkPrime(number int) {
	if number&1 == 0 {
		return
	}
	for i := 3; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return
		}
	}
	atomic.AddInt32(&totalNumberofPrimes, 1)
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	beginId := 3
	batchSize := maxLimit / batches
	for i := 0; i < batches; i++ {
		wg.Add(1)
		endId := beginId + batchSize
		go calculateBatches(strconv.Itoa(i), &wg, beginId, endId)
		beginId = endId + 1
	}
	wg.Wait()
	fmt.Println("Total no of prime numbers till", maxLimit, "is", totalNumberofPrimes, "and the process completed in ", time.Since(start))

}

func calculateBatches(name string, wg *sync.WaitGroup, beginId int, endId int) {
	defer wg.Done()
	start := time.Now()
	for i := beginId; i < endId; i++ {
		checkPrime(i)
	}
	fmt.Println("Time taken for thread ", name, "completed in ", time.Since(start))
}
