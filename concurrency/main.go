package main

import (
	"fmt"
	"sync"
)

var mt sync.Mutex

func increment(num *int, wg *sync.WaitGroup) {
	defer wg.Done()
	mt.Lock()
	*num++
	mt.Unlock()
}

func main() {
	value := 1

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&value, &wg)
	}

	wg.Wait()

	fmt.Println("value incremented", value)
}
