package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(f), "\n")

	line := lines[0]

	toReplace := map[string]string{}
	for _, ll := range lines[2:] {
		l := strings.Split(ll, " -> ")
		toReplace[l[0]] = string(l[0][0]) + string(l[1][0]) + string(l[0][1])
	}
	for x := 0; x < 10; x++ {
		replace(&line, toReplace)
	}
	min, max := findMinMax(line)
	fmt.Println(max - min)
}

func findMinMax(line string) (min, max int) {
	min = len(line)

	letters := make(map[byte]int)
	for _, r := range line {
		letters[byte(r)]++
	}
	for _, v := range letters {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return
}

func replace(line *string, replacer map[string]string) {
	for x := 0; x < len(*line)-1; x++ {
		l := *line
		for old, new := range replacer {
			if old == l[x:x+2] {
				l = l[:x] + strings.Replace(l[x:], old, new, 1)
				x++
				break
			}
		}
		*line = l
	}
	return
}
