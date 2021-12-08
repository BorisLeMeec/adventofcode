package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	len1 = 2
	len4 = 4
	len7 = 3
	len8 = 7
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(f), "\n")

	var t int
	for _, l := range lines {
		parts := strings.Split(l, " | ")

		outputsPart := parts[1]
		outputs := strings.Split(outputsPart, " ")
		for _, o := range outputs {
			switch len(o) {
			case len1, len4, len7, len8:
				t++
			}
		}
	}
	fmt.Println(t)
}
