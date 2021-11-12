package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	for i := 1; i <= num; i++ {
		fmt.Printf("Case %d: %d\n", i, maximumKnight())
	}
}

var rows, columns, maximumNumKnight int
func maximumKnight() int {
	fmt.Scanln(&rows, &columns)
	if columns==1 || columns==2 {
		columns, rows = rows, columns
	}
	if rows == 1 {
		return columns
	} else if rows == 2 {
		maximumNumKnight = ((columns / 4) + 1) * 4
		if (columns % 4) == 1{
			maximumNumKnight -= 2
		} else if (columns % 4) ==0{
			maximumNumKnight -= 4
		}
	}else {
		maximumNumKnight = ((rows * columns) + 1) / 2
	}

	return maximumNumKnight
}