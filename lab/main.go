package main

import (
	"fmt"
	goroutine_sample "lab/04_goroutine"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	startTime := time.Now()
	wg.Add(4)

	go goroutine_sample.Do5SecAsync(1, &wg)
	go goroutine_sample.Do5SecAsync(2, &wg)
	go goroutine_sample.Do5SecAsync(3, &wg)
	goroutine_sample.Do5Sec(4)
	goroutine_sample.Do5Sec(5)

	fmt.Println("Menunggu")
	wg.Wait()

	endTime := time.Now()
	fmt.Println("rest: ", endTime.Sub(startTime))
}
