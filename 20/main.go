package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(f), "\n")

	var scores []int
	for _, l := range lines {
		if nl, err := simplify(l); err == nil {
			scores = append(scores, score(nl))
		}
	}

	sort.Ints(scores)

	fmt.Println(scores[len(scores)/2])
}

func score(str string) int {
	sc := 0
	for i := len(str) - 1; i >= 0; i-- {
		sc *= 5
		sc += invertValue(str[i])
	}
	return sc
}

func invertValue(c byte) int {
	switch c {
	case '(':
		return 1
	case '[':
		return 2
	case '{':
		return 3
	case '<':
		return 4
	}
	return 0
}

func simplify(line string) (string, error) {
	index := 0
	for index < len(line) {
		if isOpener(line[index]) {
			if len(line[index:]) < 2 {
				return line, nil
			}
			if !isOpener(line[index+1]) {
				if pt := grade(line[index+1], line[index]); pt != 0 {
					return "", errors.New("")
				}
				line = line[:index] + line[index+2:]
				if index > 0 {
					index--
				}
				continue
			}
		} else {
			return "", errors.New("")
		}
		index++
	}
	return line, nil
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
