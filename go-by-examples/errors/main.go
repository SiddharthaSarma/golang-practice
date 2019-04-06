package main

import "fmt"
import "errors"

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

func main() {
	result, error := f1(42)
	fmt.Print(result, error)
}
