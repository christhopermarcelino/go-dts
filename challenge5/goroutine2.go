package main

import (
	"fmt"
	"sync"
)

var (
	flag = 0
	bisa = []string{"bisa1", "bisa2", "bisa3"}
	coba = []string{"coba1", "coba2", "coba3"}
	wg   sync.WaitGroup
	mut  sync.Mutex
)

func goroutine2() {

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go execGoroutine2(i, bisa)
		go execGoroutine2(i, coba)
	}

	wg.Wait()
}

func execGoroutine2(index int, currentData []string) {
	mut.Lock()
	defer mut.Unlock()

	if flag == 0 {
		fmt.Println(bisa, index)
		flag = 1
	} else {
		fmt.Println(coba, index)
		flag = 0
	}

	wg.Done()
}
