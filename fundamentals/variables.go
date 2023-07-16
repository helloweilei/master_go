package main

import "fmt"

func main() {
	var a, b, c int = 122, 2333, 343
	var adams int = 88

	fmt.Println(a, b, c)
	fmt.Println(`
		我自横刀向天笑，
		去留肝胆两昆仑。
	`)
	fmt.Printf("%v \t %b \t %X", adams, adams, adams)
}