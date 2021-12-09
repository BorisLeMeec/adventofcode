package main

import (
	"awesomeProject/common"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(f), "\n")

	arr := make([][]int, len(lines))
	for i, l := range lines {
		arr[i] = make([]int, len(l))
		for j, c := range l {
			arr[i][j] = common.ToInt(string(c))
		}
	}

	var t int

	for i := range arr {
		for j := range arr[i] {
			toTest := arr[i][j]
			if i > 0 && toTest >= arr[i-1][j] {
				continue
			}
			if j > 0 && toTest >= arr[i][j-1] {
				continue
			}
			if j < len(arr[i])-1 && toTest >= arr[i][j+1] {
				continue
			}
			if i < len(arr)-1 && toTest >= arr[i+1][j] {
				continue
			}
			fmt.Println(toTest)
			t += toTest + 1
		}
	}
	fmt.Println(t)
}
