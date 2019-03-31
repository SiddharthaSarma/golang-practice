package main

import "fmt"

func closureFunc() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func main() {
	innerFunc := closureFunc()

	fmt.Println(innerFunc())
	fmt.Println(innerFunc())

	otherInnerFunc := closureFunc()
	fmt.Println(otherInnerFunc())
}
