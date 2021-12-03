package main

import (
	"awesomeProject/common"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")

	ss := strings.Split(string(f), "\n")

	nbBits := len(ss[0])

	bitsCount := make([]int, nbBits)

	for _, s := range ss {
		for i := 0; i < nbBits; i++ {
			bitsCount[i] += common.ToInt(string(s[i]))
		}
	}
	var binaryG string
	var binaryE string
	for _, t := range bitsCount {
		var newbitG string
		var newbitE string
		if t > len(ss)/2 {
			newbitG = "1"
			newbitE = "0"
		} else {
			newbitG = "0"
			newbitE = "1"
		}
		binaryG = fmt.Sprintf("%s%s", binaryG, newbitG)
		binaryE = fmt.Sprintf("%s%s", binaryE, newbitE)
	}
	g, _ := strconv.ParseInt(binaryG, 2, 64)
	e, _ := strconv.ParseInt(binaryE, 2, 64)

	fmt.Println(g * e)
}
