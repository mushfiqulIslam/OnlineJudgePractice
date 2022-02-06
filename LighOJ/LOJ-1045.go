package main

import (
	"fmt"
	"math"
)

var arr [1000006]float64

func main() {
	var numOfCase, number, base int
	arr[0] = 0
	for i := 1; i < 1000006; i++ {
		arr[i] = arr[i-1] + math.Log10(1.0*float64(i))
	}
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&number, &base)
		fmt.Printf("Case %d: %d\n", i, int(arr[number]/math.Log10(float64(base)))+1)
	}
}
