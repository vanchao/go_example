package main

import "fmt"

func main() {
	if 0 == 7%2 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 0 == 8%4 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has i digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
