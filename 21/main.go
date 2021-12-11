package main

import (
	"awesomeProject/common"
	"fmt"
	"os"
	"strings"
)

type octo struct {
	p int
	f bool
}

type octopuses [][]octo

func (o octopuses) print() {
	for _, octoLine := range o {
		for _, thisocto := range octoLine {
			fmt.Printf("%d", thisocto.p)
		}
		fmt.Println()
	}
}

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(f), "\n")

	arr := make(octopuses, len(lines))
	for i, l := range lines {
		arr[i] = make([]octo, len(l))
		for j, c := range l {
			arr[i][j].p = common.ToInt(string(c))
		}
	}

	t := 0
	for x := 0; x < 100; x++ {
		up(arr)
		nf := checkFlash(arr)
		t += nf
		for nf != 0 {
			nf = checkFlash(arr)
			t += nf
		}
		reset(arr)
	}

	fmt.Println(t)
}

func reset(arr [][]octo) {
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j].f {
				arr[i][j].p = 0
				arr[i][j].f = false
			}
		}
	}
}

func up(arr octopuses) {
	for i := range arr {
		for j := range arr[i] {
			arr[i][j].p++
		}
	}
}

func checkFlash(arr octopuses) (f int) {
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j].p > 9 && !arr[i][j].f {
				f++
				flash(arr, i, j)
			}
		}
	}
	return
}

func flash(arr octopuses, i, j int) {
	if j > 0 {
		if i > 0 {
			arr[i-1][j-1].p++
		}
		arr[i][j-1].p++
		if i < len(arr)-1 {
			arr[i+1][j-1].p++
		}
	}
	if i > 0 {
		arr[i-1][j].p++
	}
	arr[i][j].f = true
	arr[i][j].p++
	if i < len(arr)-1 {
		arr[i+1][j].p++
	}
	if j < len(arr[i])-1 {
		if i > 0 {
			arr[i-1][j+1].p++
		}
		arr[i][j+1].p++
		if i < len(arr)-1 {
			arr[i+1][j+1].p++
		}
	}
}
