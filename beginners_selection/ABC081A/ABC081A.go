package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	var splitStr []string = strings.Split(s, "")

	count := 0
	for _, char := range splitStr {
		if char == "1" {
			count++
		}
	}
	fmt.Println(count)
}
