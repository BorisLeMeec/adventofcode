package main

import (
	"awesomeProject/common"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	startPosStr := strings.Split(string(f), ",")

	var startPos []int
	for _, posStr := range startPosStr {
		startPos = append(startPos, common.ToInt(posStr))
	}

	sort.Ints(startPos)

	min := startPos[0]
	max := startPos[len(startPos)-1]

	fuel := ^uint(0)
	for x := min; x <= max; x++ {
		var t uint

		for _, crabe := range startPos {
			t += consecutiveSum(uint(math.Abs(float64(crabe - x))))
		}
		if t < fuel {
			fuel = t
		}
	}

	fmt.Println(fuel)
}

func consecutiveSum(u uint) (t uint) {
	for x := uint(1); x <= u; x++ {
		t += x
	}
	return
}
