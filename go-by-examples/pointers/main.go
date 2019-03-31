package main

import "fmt"

func zeroVal(ival int) {
	fmt.Println("pass by value", ival)
	ival = 0
}

func zeroPtr(ival *int) {
	fmt.Println("pass by reference", ival)
	*ival = 0
}

func main() {
	i := 2
	zeroVal(i)
	fmt.Println("i value:", i)

	zeroPtr(&i)
	fmt.Println("i value", i)
}
