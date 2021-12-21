package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	initStr := strings.Split(string(f), "\n")[0]
	reg, _ := regexp.Compile(`x=(-?\d*)..(-?\d*), y=(-?\d*)..(-?\d*)`)
	nbStr := reg.FindStringSubmatch(initStr)[1:]
	minX, _ := strconv.ParseInt(nbStr[0], 10, 64)
	maxX, _ := strconv.ParseInt(nbStr[1], 10, 64)
	minY, _ := strconv.ParseInt(nbStr[2], 10, 64)
	maxY, _ := strconv.ParseInt(nbStr[3], 10, 64)

	maxH := int64(0)
	for velX := int64(math.Round(0.5*math.Sqrt(float64(8*minX+1)) - 0.5)); velX < int64(math.Round(0.5*math.Sqrt(float64(8*maxX+1))-0.5)); velX++ {
		fmt.Println("essaie velx ", velX)
		for velY := int64(0); velY < 1000; velY++ {
			maxHeight := (velY*velY + velY) / 2

			fmt.Println("essaie vely ", velY)
			for step := uint(0); ; step++ {
				x, y := pos(velX, velY, step)
				if y < minY || x > maxX {
					fmt.Println("too")
					break
				} else if x >= minX && x <= maxX && y >= minY && y <= maxY {
					if maxHeight > maxH {
						maxH = maxHeight
					}
				}
			}
		}
	}

	fmt.Println(maxH)
}

func pos(initXVel, initYVel int64, step uint) (x, y int64) {
	for i := uint(0); i <= step; i++ {
		x += initXVel
		y += initYVel
		initYVel--
		if initXVel > 0 {
			initXVel--
		} else if initXVel < 0 {
			initXVel++
		}
	}
	return
}
