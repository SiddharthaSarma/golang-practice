package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 8
	m["k2"] = 5
	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println(v1)

	a, b := m["k1"]
	fmt.Println(a, b)

	_, c := m["k1"]
	fmt.Println(c)

	delete(m, "k1")
	_, d := m["k1"]
	fmt.Println(d)

	n := map[string]int{"foo": 1, "bar": 2}
	n["sid"] = 5
	fmt.Println("map:", n)
}
