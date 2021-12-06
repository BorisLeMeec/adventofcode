package main

import (
	"awesomeProject/common"

	"fmt"
	"os"
	"strings"
)

const X = 0
const Y = 1

type pts [2]int

func main() {
	var arr [1000][1000]uint8

	f, _ := os.ReadFile("input.txt")

	ss := strings.Split(string(f), "\n")

	for _, s := range ss {
		points := strings.Split(s, " -> ")
		astr := strings.Split(points[0], ",")
		bstr := strings.Split(points[1], ",")
		a, b := pts{common.ToInt(astr[X]), common.ToInt(astr[Y])}, pts{common.ToInt(bstr[X]), common.ToInt(bstr[Y])}
		ptsb := ptsBetween(a, b)
		for _, p := range ptsb {
			arr[p[X]][p[Y]]++
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

func ptsBetween(a, b pts) (o []pts) {
	if a[X] == b[X] {
		points := straight(a[Y], b[Y])
		for _, p := range points {
			o = append(o, pts{a[X], p})
		}
		return
	} else if a[Y] == b[Y] {
		points := straight(a[X], b[X])
		for _, p := range points {
			o = append(o, pts{p, a[Y]})
		}
		return
	} else {
		o = append(o, diag(a, b)...)
	}
	return
}

func diag(a, b pts) (o []pts) {
	px := straight(a[X], b[X])
	py := straight(a[Y], b[Y])
	for i, x := range px {
		o = append(o, pts{x, py[i]})
	}
	return
}

func straight(a, b int) (o []int) {
	if a > b {
		for ; b <= a; b++ {
			o = append(o, b)
		}
		return
	} else {
		for ; b >= a; b-- {
			o = append(o, b)
		}
		return
	}
}
