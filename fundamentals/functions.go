package main

import "fmt"

func hanoi(a, b, c string, n int) {
	if n == 1 {
		fmt.Println(a + " --> " + c)
		return
	}
	hanoi(a, c, b, n - 1)
	fmt.Println(a + " --> " + c)
	hanoi(b, a, c, n - 1)
}

func main() {
	hanoi("A", "B", "C", 5)
}