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

	hpos, depth, aim := 0,0,0

	for _, s := range ss {
		input := strings.Split(s, " ")
		move, xstr := input[0], input[1]
		x, _ := strconv.Atoi(xstr)

		switch move {
		case "up":
			aim -= x
		case "down":
			aim += x
		case "forward":
			hpos += x
			depth += x*aim
		}
	}

	fmt.Println(hpos*depth)
}
