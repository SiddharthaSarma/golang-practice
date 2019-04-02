package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"sid", 26})
	fmt.Println(person{name: "Alice", age: 30})
	s := person{name: "Bob", age: 50}
	fmt.Println(s.age, s.name)
	s.name = "Michael"
	fmt.Println(s)
	newPerson := &s
	newPerson.age = 100
	fmt.Println(*newPerson)
	fmt.Println(s)
}
