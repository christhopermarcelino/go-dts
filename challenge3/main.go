package main

import "fmt"

func main() {
	textMap := make(map[string]int)
	text := "selamat malam"

	for _, valRune := range text {
		textMap[string(valRune)]++
	}

	fmt.Println(textMap)
}
