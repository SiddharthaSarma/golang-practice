package main

import "fmt"
import "time"

func main() {
	num := 2
	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its weekend")
	default:
		fmt.Println("Its weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whoamI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("integer")
		case bool:
			fmt.Println("boolean")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whoamI(true)
	whoamI("hello")
	whoamI(5)

}
