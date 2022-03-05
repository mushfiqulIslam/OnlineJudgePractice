package main

import (
	"fmt"
)

func main() {
	var numOfCase, ax, ay, bx, by, cx, cy, dx, dy, area int
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&ax, &ay, &bx, &by, &cx, &cy)
		dx = ax + cx - bx
		dy = ay + cy - by
		area = int(0.5 * float64(((ax*by)+(bx*cy)+(cx*dy)+(dx*ay))-((ay*bx)+(by*cx)+(cy*dx)+(dy*ax))))
		if area < 0 {
			area *= -1
		}
		fmt.Printf("Case %d: %d %d %d\n", i, dx, dy, area)
	}
}
