package main

import "fmt"

func main() {
	fmt.Println("Go routine 1")
	goroutine1()
	fmt.Println("Go routine 2")
	goroutine2()
}
