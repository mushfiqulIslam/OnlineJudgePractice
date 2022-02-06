package main

import (
	"fmt"
	"sort"
)

func main() {
	var numOfCase int
	var arr [3]int
	output := "no"
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		output = "no"
		fmt.Scan(&arr[0], &arr[1], &arr[2])
		sort.Ints(arr[:])
		if arr[2]*arr[2] == ((arr[0] * arr[0]) + (arr[1] * arr[1])) {
			output = "yes"
		}
		fmt.Printf("Case %d: %s\n", i, output)
	}
}
