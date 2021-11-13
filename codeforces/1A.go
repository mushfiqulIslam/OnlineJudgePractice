package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m, a, num, mum float64
	fmt.Scanln(&n, &m, &a)
	num = math.Ceil(n / a)
	mum = math.Ceil(m / a)
	fmt.Println(int64(num * mum))
}

