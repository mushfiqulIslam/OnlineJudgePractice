package main

import (
	"fmt"
	"math"
)

var pi float64 = 2 * math.Acos(0)

func main() {
	var numOfCase int
	var radius, areaOfCircle, areaOfSquare, shadedArea float64
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&radius)
		areaOfCircle = radius * radius * pi
		areaOfSquare = 2 * radius * 2 * radius
		shadedArea = areaOfSquare - areaOfCircle
		fmt.Printf("Case %d: %.2f\n", i, shadedArea)
	}
}
