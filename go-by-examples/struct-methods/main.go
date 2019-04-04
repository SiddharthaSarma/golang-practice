package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}
func main() {

	s := rect{width: 20, height: 5}
	fmt.Println(s.area())
	fmt.Println(s.perimeter())

	rp := &s
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perimeter())
}
