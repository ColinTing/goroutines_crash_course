package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu     sync.Mutex
	NumMap map[string]int
}

func (sc *SafeCounter) Inc(num int) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.NumMap["key"] = num
}

func main() {
	s := SafeCounter{NumMap: make(map[string]int)}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Inc(i)
		}(i)
	}
	wg.Wait()
	fmt.Println(s.NumMap["key"])
}
