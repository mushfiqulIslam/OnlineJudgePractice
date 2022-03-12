package main

import (
	"fmt"
	"strings"
)

func main() {
	var numOfCase int
	var url string
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&url)
		splitUrl := strings.Split(url, "//")
		if splitUrl[0] == "http:" {
			splitUrl[0] = "https:"
		}
		fmt.Printf("Case %d: %s\n", i, splitUrl[0]+"//"+splitUrl[1])
	}
}
