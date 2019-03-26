package main

import "fmt"

func main() {
	if 4%2 == 0 {
		fmt.Println("4 is even number")
	} else {
		fmt.Println("4 is odd number")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
