package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(f), "\n")

	totalPoints := 0
	for i, l := range lines {
		fmt.Printf("line %d: \n", i)
		totalPoints += pointsLine(l)
		fmt.Println("points are ", totalPoints)
	}

	fmt.Println(totalPoints)
}

func pointsLine(line string) int {
	index := 0
	for index < len(line) {
		if isOpener(line[index]) {
			if len(line[index:]) < 2 {
				return 0
			}
			if !isOpener(line[index+1]) {
				if pt := grade(line[index+1], line[index]); pt != 0 {
					return pt
				}
				line = line[:index] + line[index+2:]
				if index > 0 {
					index--
				}
				continue
			}
		} else {
			return grade(line[index], '0')
		}
		index++
	}
	return 0
}

func isOpener(b byte) bool {
	return b == '[' || b == '{' || b == '<' || b == '('
}

func grade(closer, opener byte) int {
	switch closer {
	case ')':
		if opener == '(' {
			return 0
		}
		return 3
	case ']':
		if opener == '[' {
			return 0
		}
		return 57

	case '}':
		if opener == '{' {
			return 0
		}
		return 1197
	case '>':
		if opener == '<' {
			return 0
		}
		return 25137
	}
	return 0
}
