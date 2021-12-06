package main

import (
	"awesomeProject/common"
	"fmt"
	"os"
	"strings"
)

const (
	maxDays      = 256
	defaultTimer = 8
	childTimer   = 0
)

func main() {
	f, _ := os.ReadFile("input.txt")
	startstr := strings.Split(string(f), ",")

	var start [defaultTimer + 1]int
	for _, sfStr := range startstr {
		start[common.ToInt(sfStr)]++
	}
	sf := simulate(start, 0)
	var t = 0
	for i := defaultTimer; i >= 0; i-- {
		t += sf[i]
	}
	fmt.Println(t)
}

func simulate(sf [defaultTimer + 1]int, day int) [defaultTimer + 1]int {
	if day == maxDays {
		return sf
	}
	var tmp int
	var tmp2 int
	for i := defaultTimer; i >= childTimer; i-- {
		if i == childTimer {
			sf[defaultTimer] += sf[i]
			sf[6] += sf[i]
		}
		tmp2 = sf[i]
		sf[i] = tmp
		tmp = tmp2
	}
	fmt.Println(day)
	return simulate(sf, day+1)
}
