package main

import "fmt"

func divisible(a int) int {
	/*
		If n % 3 = 0 then exactly (n / 3) integers are not divisible by 3.
		If n % 3 = 1 then exactly floor(n / 3) + 1 integers are not divisible by 3.
		If n % 3 = 2 then exactly floor(n / 3) + 1 integers are not divisible by 3.
	*/
	var less int
	if a%3 == 0 {
		less = a / 3
	} else {
		less = (a / 3) + 1
	}
	return a - less
}

func main() {
	var numOfCase, start, end int
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&start, &end)
		fmt.Printf("Case %d: %d\n", i, divisible(end)-divisible(start-1))
	}
}
