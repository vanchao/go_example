package main

import "fmt"

func vals(a int, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 7, 10
	a, b = vals(a, b)
	fmt.Println(a)
	fmt.Println(b)
}
