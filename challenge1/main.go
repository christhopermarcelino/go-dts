package main

import "fmt"

func main() {
	i := 21

	fmt.Printf("%v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%%\n")

	j := true
	fmt.Printf("%t\n\n", j)

	fmt.Printf("%b\n", i)

	Я := '\u042F'
	fmt.Printf("%c\n", Я)

	fmt.Printf("%d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%x\n", 15)
	fmt.Printf("%X\n", 15)

	fmt.Printf("%U\n\n", Я)

	var k float64 = 123.456
	fmt.Printf("%f\n", k)
	fmt.Printf("%E", k)

}
