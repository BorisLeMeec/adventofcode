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
	for i := range ss {
		if i == len(ss)-2{
			break
		}
		n, _ := strconv.Atoi(ss[i])
		n1, _ := strconv.Atoi(ss[i+1])
		n2, _ := strconv.Atoi(ss[i+2])
		t := n+n1+n2
		if previous != -1 {
			if t > previous{
				c++
			}
		}
		previous = t
	}

	fmt.Println(c)
}
