package main

import (
	"fmt"
	"math"
)

func main() {
	var numOfCase int
	var ox, oy, ax, ay, bx, by, oa, ob, ab, angle float64
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&ox, &oy, &ax, &ay, &bx, &by)

		oa = math.Sqrt(math.Pow(ax-ox, 2) + math.Pow(ay-oy, 2))
		ob = math.Sqrt(math.Pow(bx-ox, 2) + math.Pow(by-oy, 2))
		ab = math.Sqrt(math.Pow(ax-bx, 2) + math.Pow(ay-by, 2))
		angle = math.Acos((math.Pow(oa, 2) + math.Pow(ob, 2) - math.Pow(ab, 2)) / (2 * oa * ob))
		fmt.Printf("Case %d: %.3f\n", i, oa*angle)
	}
}
