package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")

	ss := strings.Split(string(f), "\n")

	previous := -1
	c:= 0
	for _, s := range ss {
		n, _ := strconv.Atoi(s)
		if previous != -1 {
			if n > previous{
				c++
			}
		}
		previous = n
	}

	fmt.Println(c)
}
