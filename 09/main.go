package main

import (
	"awesomeProject/common"
	"fmt"
	"os"
	"strings"
)

func main() {
	var arr [1000][1000]uint8

	f, _ := os.ReadFile("input.txt")

	ss := strings.Split(string(f), "\n")

	for _, s := range ss {
		pts := strings.Split(s, " -> ")
		a := strings.Split(pts[0], ",")
		b := strings.Split(pts[1], ",")
		xa, ya := common.ToInt(a[0]), common.ToInt(a[1])
		xb, yb := common.ToInt(b[0]), common.ToInt(b[1])

		if xa == xb {
			ps := ptsBetween(ya, yb)
			for _, p := range ps {
				arr[xa][p]++
			}
		} else if ya == yb {
			ps := ptsBetween(xa, xb)
			for _, p := range ps {
				arr[p][ya]++
			}
		}
	}

	var t = 0
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] >= 2 {
				t++
			}
		}
	}
	fmt.Println(t)
	return
}

func ptsBetween(a, b int) (o []int) {
	if a > b {
		for ; b <= a; b++ {
			o = append(o, b)
		}
		return
	} else {
		for ; a <= b; a++ {
			o = append(o, a)
		}
		return
	}
}
