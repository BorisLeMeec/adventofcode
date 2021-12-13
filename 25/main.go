package main

import (
	"awesomeProject/common"
	"fmt"
	"os"
	"strings"
)

type dot struct {
	x, y int
}

func main() {
	f, _ := os.ReadFile("input.txt")
	parts := strings.Split(string(f), "\n\n")

	dotsStr, instructions := parts[0], parts[1]
	dotsL := strings.Split(dotsStr, "\n")

	dots := getDots(dotsL)
	maxX, maxY := findMaxs(dots)

	sheet := getSheet(maxX, maxY)
	for _, d := range dots {
		sheet[d.y][d.x] = true
	}

	fold(strings.Split(instructions[:1], "\n"), sheet)

	sheet.print()

	fmt.Println(sheet.countDots())
}

func (s paperSheet) countDots() (d int) {
	for i := range s {
		for j := range s[i] {
			if s[i][j] {
				d++
			}
		}
	}
	return
}

func getSheet(x, y int) paperSheet {
	ret := make(paperSheet, y+1)
	for i := 0; i <= y; i++ {
		ret[i] = make([]bool, x+1)
	}

	return ret
}

type paperSheet [][]bool

func (s paperSheet) print() {
	for i := range s {
		for j := range s[i] {
			var c byte
			if s[i][j] {
				c = '#'
			} else {
				c = '.'
			}
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

func fold(instructions []string, sheet [][]bool) {
	for _, i := range instructions {
		i = i[11:]
		val := common.ToInt(i[2:])
		if i[0] == 'x' {
			foldVert(sheet, val)
		} else {
			foldHor(sheet, val)
		}
	}
}

func foldVert(sheet [][]bool, val int) {
	for i, l := range sheet {
		for x := val + 1; x < len(l); x++ {
			nval := val - (x - val)
			if nval < 0 {
				break
			}
			if sheet[i][x] == true {
				sheet[i][x] = false
				sheet[i][nval] = true
			}
		}
	}
}

func foldHor(sheet [][]bool, val int) {
	for j := range sheet[0] {
		for y := val + 1; y < len(sheet); y++ {
			nval := val - (y - val)
			if nval < 0 {
				break
			}
			if sheet[y][j] == true {
				sheet[y][j] = false
				sheet[nval][j] = true
			}
		}
	}
}

func getDots(lines []string) (arr []dot) {
	for _, d := range lines {
		coords := strings.Split(d, ",")
		arr = append(arr, dot{x: common.ToInt(coords[0]), y: common.ToInt(coords[1])})
	}
	return
}

func findMaxs(dots []dot) (maxX int, maxY int) {
	for _, d := range dots {
		if d.x > maxX {
			maxX = d.x
		}
		if d.y > maxY {
			maxY = d.y
		}
	}

	return
}
