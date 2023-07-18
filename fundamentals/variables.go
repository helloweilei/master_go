package main

import "fmt"

func main() {
	var a, b, c int = 122, 2333, 343
	var adams int = 88
	z := 42

	fmt.Println(a, b, c, z)
	fmt.Println(`
		我自横刀向天笑，
		去留肝胆两昆仑。
	`)
	fmt.Printf("%v \t %b \t %X", adams, adams, adams)

	var x float32 = 42.42
	fmt.Println(x)
	num, str := 43, "hello world"
	fmt.Println(num, str)

	s, i, f := "hello", 42, 42.42
	fmt.Printf("%v is of type %T\n", s, s)
	fmt.Printf("%v is of type %T\n", i, i)
	fmt.Printf("%v is of type %T\n", f, f)
}