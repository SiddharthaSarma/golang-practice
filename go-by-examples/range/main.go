package main

import "fmt"

func main() {
	for i, k := range "go" {
		if string(k) == "g" {
			fmt.Println("hey")
		}
		fmt.Printf("%d -> %c\n", i, k)
	}

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
