package main

import "fmt"

func main() {
	c := make([]string, 3)
	fmt.Println(c)

	c[0] = "a"
	c[1] = "b"
	c[2] = "c"
	fmt.Println("set:", c)
	fmt.Println("get:", c[2])
	// append
	c = append(c, "e", "f", "g")
	fmt.Println(c)
	fmt.Println(len(c))

	// copy
	s := make([]string, len(c))
	copy(s, c)
	fmt.Println("cpy:", s)
	fmt.Println("splice", s[:2])
	fmt.Println("splice", s[2:])

}
