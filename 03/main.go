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

	x, y := 0,0
	for _, s := range ss {
		input := strings.Split(s, " ")
		move, l := input[0], input[1]
		ll, _ := strconv.Atoi(l)

		switch move {
		case "up":
			y -= ll
		case "down":
			y += ll
		case "forward":
			x += ll
		}
	}

	fmt.Println(x*y)
}
