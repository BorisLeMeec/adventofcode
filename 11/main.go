package main

import (
	"awesomeProject/common"
	"fmt"
	"os"
	"strings"
)

const (
	maxDays      = 80
	defaultTimer = 8
	childTimer   = 0
)

func main() {
	f, _ := os.ReadFile("input.txt")
	startstr := strings.Split(string(f), ",")

	var start []int
	for _, sfStr := range startstr {
		start = append(start, common.ToInt(sfStr))
	}

	fmt.Println(simulate(start, 0))

}

func simulate(sf []int, day int) int {
	if day == maxDays {
		return len(sf)
	}
	for i := range sf {
		if sf[i] == childTimer {
			sf[i] = 6
			sf = append(sf, defaultTimer)
		} else {
			sf[i]--
		}
	}
	return simulate(sf, day+1)
}
