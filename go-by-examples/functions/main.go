package main

import "fmt"

func addTwoNumbers(a int, b int) int {
	return a + b
}
func addThreeNumbers(a, b, c int) int {
	return a + b + c
}

func main() {
	fmt.Println(addTwoNumbers(2, 3))
	sum := addThreeNumbers(2, 3, 4)
	fmt.Println(sum)
}
