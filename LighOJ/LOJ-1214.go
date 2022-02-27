package main

import (
	"fmt"
	"math/big"
)

func main() {
	var numOfCase int
	var stra, strb string
	zero := big.NewInt(0)
	a := new(big.Int)
	b := new(big.Int)
	modOut := new(big.Int)
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&stra, &strb)
		fmt.Sscan(stra, a)
		fmt.Sscan(strb, b)
		modOut = modOut.Mod(a, b)
		if modOut.Cmp(zero) == 0 {
			fmt.Printf("Case %d: divisible\n", i)
		} else {
			fmt.Printf("Case %d: not divisible\n", i)
		}
	}
}
