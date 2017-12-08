package main

import "fmt"

func zeroptr(iptr *int) {
	*iptr = *iptr
}

func main() {
	i := 1
	zeroptr(&i)
	fmt.Println(i)
	fmt.Println(&i)
}
