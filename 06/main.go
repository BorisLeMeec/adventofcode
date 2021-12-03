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

	bitoxy := findOxy(ss)
	oxy, _ := strconv.ParseInt(bitoxy, 2, 64)

	bitco2 := findCo2(ss)
	co2, _ := strconv.ParseInt(bitco2, 2, 64)

	fmt.Println(oxy * co2)
}

func findOxy(ss []string) string {
	return findTheOne(ss, 0, maxAtPos)
}

func findCo2(ss []string) string {
	return findTheOne(ss, 0, minAtPos)
}

func findTheOne(ss []string, index uint8, f func([]string, int) uint8) string {
	if len(ss) <= 1 {
		return ss[0]
	}
	var newss []string

	bitWanted := f(ss, int(index))
	for _, s := range ss {
		if s[index] == bitWanted {
			newss = append(newss, s)
		}
	}

	return findTheOne(newss, index+1, f)
}

func maxAtPos(ss []string, pos int) uint8 {
	t := 0

	for _, s := range ss {
		t += common.ToInt(string(s[pos]))
	}
	if float64(t) >= float64(len(ss))/2.0 {
		return '1'
	}
	return '0'
}

func minAtPos(ss []string, pos int) uint8 {
	t := 0

	for _, s := range ss {
		t += common.ToInt(string(s[pos]))
	}
	if float64(t) >= float64(len(ss))/2.0 {
		return '0'
	}
	return '1'
}
