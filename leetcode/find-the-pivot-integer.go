package main

import "math"

func pivotInteger(n int) int {
	if n <= 1 {
		return n
	}

	x := math.Sqrt(float64(n * (n + 1) / 2))

	if math.Mod(x, 1) == 0 {
		return int(x)
	}
	return -1
}
