package main

import (
	"fmt"
	"sort"
)

type NumberPair struct {
	number       int
	divisorCount int
}

// NumberPairSlice Sorts NumberPair by divisorCount.
type NumberPairSlice []NumberPair

func (a NumberPairSlice) Len() int      { return len(a) }
func (a NumberPairSlice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a NumberPairSlice) Less(i, j int) bool {
	if a[i].divisorCount != a[j].divisorCount {
		return a[i].divisorCount < a[j].divisorCount
	} else {
		return a[i].number > a[j].number
	}
}

func main() {
	var numOfCase, n int
	var divisorCountPair = make([]NumberPair, 1001)
	// Divisor Sieve O(n log n)
	for i := 1; i < 1001; i++ {
		for j := i; j < 1001; j += i {
			divisorCountPair[j].number = j
			divisorCountPair[j].divisorCount++
		}
	}
	sort.Sort(NumberPairSlice(divisorCountPair))
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&n)
		fmt.Printf("Case %d: %d\n", i, divisorCountPair[n].number)
	}
}
