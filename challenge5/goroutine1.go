package main

import (
	"fmt"
	"sync"
)

func goroutine1() {
	bisa := []string{"bisa1", "bisa2", "bisa3"}
	coba := []string{"coba1", "coba2", "coba3"}

	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go execGoroutine1(i, bisa, &wg)
		go execGoroutine1(i, coba, &wg)
	}

	wg.Wait()
}

func execGoroutine1(index int, data []string, wg *sync.WaitGroup) {
	fmt.Println(data, index)
	wg.Done()
}
