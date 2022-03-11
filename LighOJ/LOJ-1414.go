package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numOfCase, multipleOf4, multipleOf100, multipleOf400, leapYearCount int
	reader := bufio.NewReader(os.Stdin)
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		date1, _ := reader.ReadString('\n')
		date2, _ := reader.ReadString('\n')
		date1Slice := strings.Split(date1, " ")
		date2Slice := strings.Split(date2, " ")
		date1Slice[1] = strings.TrimSuffix(date1Slice[1], ",")
		date2Slice[1] = strings.TrimSuffix(date2Slice[1], ",")
		date1Slice[2] = date1Slice[2][:len(date1Slice[2])-1]
		date2Slice[2] = date2Slice[2][:len(date2Slice[2])-1]
		day2, _ := strconv.Atoi(date2Slice[1])
		year1, _ := strconv.Atoi(date1Slice[2])
		year2, _ := strconv.Atoi(date2Slice[2])
		leapYearCount = 0
		if !(date1Slice[0] == "January" || date1Slice[0] == "February") {
			year1++
		}
		if date2Slice[0] == "January" || (date2Slice[0] == "February" && day2 < 29) {
			year2--
		}

		multipleOf4 = year2/4 - (year1-1)/4
		multipleOf100 = year2/100 - (year1-1)/100
		multipleOf400 = year2/400 - (year1-1)/400
		leapYearCount = multipleOf4 + multipleOf400 - multipleOf100

		fmt.Printf("Case %d: %d\n", i, leapYearCount)
	}
}
