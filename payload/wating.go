package payload

import (
	"fmt"
	"sync"
	"time"
)

func Wait() {
	var wg sync.WaitGroup
	// var mut sync.Mutex
	for j := 0; j < 500; j++ {
		i := 2
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			start := time.Now()
			time.Sleep(time.Duration(val) * time.Second)
			end := time.Now()
			count := end.Sub(start)
			if count >= time.Duration(3)*time.Second {
				fmt.Printf("task %d is more than 3s\n", val)
			} else {
				fmt.Printf("task %d finished at %v, took %v\n", val, end, count)
			}
		}(i)
	}
	wg.Wait()
}
