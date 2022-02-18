package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numOfCase, number int
	var output string
	var factArray [21]int
	factArray[0] = 1
	for i := 1; i < 21; i++ {
		factArray[i] = factArray[i-1] * i
	}
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		factStack := make([]int, 0)
		fmt.Scan(&number)
		for j := 20; j >= 0 && number > 0; j-- {
			if number >= factArray[j] {
				factStack = append(factStack, j)
				number -= factArray[j]
			}
		}
		if number != 0 {
			output = "impossible"
		} else {
			//                1!+3!
			for j := len(factStack) - 1; j >= 0; j-- {
				if j == len(factStack)-1 {
					output = strconv.Itoa(factStack[j]) + "!"
				} else {
					output += "+" + strconv.Itoa(factStack[j]) + "!"
				}
			}
		}
		fmt.Printf("Case %d: %s\n", i, output)
		output = ""
	}
}
