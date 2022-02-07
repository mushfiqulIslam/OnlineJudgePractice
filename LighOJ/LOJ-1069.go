package main

import (
	"fmt"
)

func main() {
	var numOfCase, myPosition, liftPosition, takenTime int
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&myPosition, &liftPosition)
		// Entering, exiting, lift opening 2 times and closing 1 time
		takenTime = 5 + 5 + 2*3 + 3
		if myPosition == liftPosition {
			takenTime += myPosition * 4
		} else if myPosition > liftPosition {
			takenTime += (myPosition + (myPosition - liftPosition)) * 4
		} else {
			takenTime += liftPosition * 4
		}
		fmt.Printf("Case %d: %d\n", i, takenTime)
	}
}
