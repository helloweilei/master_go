package main

import "fmt"

const (
	a = 4
	b = 6
	c = iota
)

func main() {
	fmt.Println(a, b, c) // 4 6 2
}