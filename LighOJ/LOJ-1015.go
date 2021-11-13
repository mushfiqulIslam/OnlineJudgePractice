package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	for i := 1; i <= num; i++ {
		fmt.Printf("Case %d: %d\n", i, requiredDust())
	}
}

func requiredDust() int {
	var numOfStudent, num, requiredDustNum int
	fmt.Scan(&numOfStudent)
	for i:=0; i<numOfStudent; i++ {
		fmt.Scan(&num)
		if num > 0 {
			requiredDustNum += num
		}
	}
	return requiredDustNum
}
