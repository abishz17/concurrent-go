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
var currentNumber int32 = 2

func checkPrime(number int) {
	if number%2 == 0 {
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

	for i := 0; i < batches; i++ {
		wg.Add(1)
		go performTask(strconv.Itoa(i), &wg)
	}
	wg.Wait()

	fmt.Println("Total no of prime numbers till", maxLimit, "is", totalNumberofPrimes, "and the process completed in ", time.Since(start))

}

func performTask(name string, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()

	for {
		x := atomic.AddInt32(&currentNumber, 1)
		if x > maxLimit {
			break
		}
		checkPrime(int(x))
	}
	fmt.Println("Thread no.", name, " completed in ", time.Since(start))

}
