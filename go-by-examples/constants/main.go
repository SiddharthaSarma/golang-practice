package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	const n = 500000000
	fmt.Println(s)
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
