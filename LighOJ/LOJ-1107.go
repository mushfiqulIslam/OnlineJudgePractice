package main

import (
	"fmt"
)

func main() {
	var numOfCase, numOfCow int
	var coordinate [4]int
	var cowPosition [2]int
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&coordinate[0], &coordinate[1], &coordinate[2], &coordinate[3])
		fmt.Scan(&numOfCow)
		fmt.Printf("Case %d:\n", i)
		for j := 1; j <= numOfCow; j++ {
			fmt.Scan(&cowPosition[0], &cowPosition[1])
			if (coordinate[0] <= cowPosition[0]) && (cowPosition[0] <= coordinate[2]) &&
				(coordinate[1] <= cowPosition[1]) && (cowPosition[1] <= coordinate[3]) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}
