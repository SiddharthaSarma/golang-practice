package main

import "fmt"

func addTwoNumbers(a int, b int) int {
	return a + b
}
func addThreeNumbers(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 2, 4
}

func variadicFunc(nums ...int) int {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(addTwoNumbers(2, 3))
	sum := addThreeNumbers(2, 3, 4)
	fmt.Println(sum)
	fmt.Println(vals())
	fmt.Println(variadicFunc(2, 3, 4))
}
