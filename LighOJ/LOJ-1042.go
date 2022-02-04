package main

import (
	"fmt"
)

func main() {
	var numOfCase, number, rightMostNumber, nextHigherNumber, rightOnesPatternNumber int

	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&number)
		// Bitwise AND to find Right most patterns number like: 100, 10, 1
		rightMostNumber = number & -number
		nextHigherNumber = number + rightMostNumber
		// Bitwise XOR to find bigger pattern's number of right side
		rightOnesPatternNumber = number ^ nextHigherNumber
		// Find little number of the patter
		rightOnesPatternNumber /= rightMostNumber
		//  2 bit Right shift
		rightOnesPatternNumber = rightOnesPatternNumber >> 2
		// Bitwise OR to find the next higher number
		number = nextHigherNumber | rightOnesPatternNumber
		fmt.Printf("Case %d: %d\n", i, number)
	}
}
