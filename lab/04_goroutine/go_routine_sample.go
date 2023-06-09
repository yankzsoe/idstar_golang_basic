package goroutine_sample

import (
	"fmt"
	"sync"
	"time"
)

func Do5SecAsync(p int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Mulai: ", p)
	time.Sleep(time.Second * 5)
	fmt.Println("Selesai")
}

func Do5Sec(p int) {
	fmt.Println("Mulai: ", p)
	time.Sleep(time.Second * 5)
	fmt.Println("Selesai")
}
