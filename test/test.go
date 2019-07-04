package main

import "fmt"

func main() {
	var a,b uint8
	a, b = 255, 85 // 11111111, 01010101
	fmt.Printf("%08b\n", a&b) // 01010101
	fmt.Printf("Hello World\n")
	}
