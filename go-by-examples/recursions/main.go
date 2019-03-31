package main

import "fmt"

func fact(num int) int {
	if num == 1 {
		return num
	}
	return num * fact(num-1)
}
func main() {
	fmt.Println(fact(7))
}
