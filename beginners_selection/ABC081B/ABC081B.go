package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n string
	var A string
	var aStrs []string
	var aInts []int
	var a int
	var err error

	fmt.Scanln(&n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	A = scanner.Text()

	aStrs = strings.Split(A, " ")
	for i := 0; i < len(aStrs); i++ {
		a, err = strconv.Atoi(aStrs[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		aInts = append(aInts, a)
	}

	count := 0
	for {
		for i := 0; i < len(aInts); i++ {
			if aInts[i]%2 == 0 {
				aInts[i] = aInts[i] / 2
			} else {
				fmt.Println(count)
				return
			}
		}
		count++
	}
}
