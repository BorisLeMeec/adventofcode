package main

import (
	"awesomeProject/common"
	"fmt"
	"os"
	"strings"
)

type step struct {
	danger uint
	weight uint
	paths  paths
	isEnd  bool
	name   string
}
type heap []*step
type paths []*step

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(f), "\n")

	start := getCavern(lines)

	fmt.Println(getShortest(start, heap{}))
}

func getCavern(lines []string) *step {
	cavern := make([]*step, len(lines)*len(lines[0]))

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			cavern[y*len(lines[0])+x] = &step{danger: uint(common.ToInt(string(lines[y][x]))), weight: ^uint(0), name: fmt.Sprintf("%d-%d", y, x)}
			if y == len(lines)-1 && x == len(lines[y])-1 {
				cavern[y*len(lines[0])+x].isEnd = true
			}
		}
	}
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			s := cavern[y*len(lines[0])+x]
			if y > 0 {
				s.paths = append(s.paths, cavern[(y-1)*len(lines[0])+x])
			}
			if y < len(lines)-1 {
				s.paths = append(s.paths, cavern[(y+1)*len(lines[0])+x])
			}
			if x > 0 {
				s.paths = append(s.paths, cavern[(y)*len(lines[0])+x-1])
			}
			if y < len(lines[0])-1 {
				s.paths = append(s.paths, cavern[(y)*len(lines[0])+x+1])
			}
		}
	}

	return cavern[0]
}

func getShortest(current *step, h heap) uint {
	for !current.isEnd {
		fmt.Printf("current is %s\n", current.name)
		if !h.in(current) {
			h = append(h, current)
		}
		min := struct {
			min  uint
			h, p int
		}{^uint(0), 0, 0}
		for j, step := range h {
			for i := 0; i < len(step.paths); i++ {
				if step.paths[i].weight != 0 && step.weight+step.paths[i].danger > step.paths[i].weight {
					removePath(step.paths[i], step)
					i--
					continue
				} else if step.weight+step.paths[i].danger < min.min {
					min.min = step.weight + step.paths[i].danger
					min.h = j
					min.p = i
					current = step.paths[i]
				}
			}
			if len(step.paths) == 0 {
				h = h.delete(step)
			}
		}
		current.weight = h[min.h].weight + current.danger
		removePath(current, h[min.h])
	}

	return current.weight + current.danger
}

func removePath(a *step, b *step) {
	for i := range a.paths {
		if a.paths[i] == b {
			a.paths = append(a.paths[:i], a.paths[i+1:]...)
			break
		}
	}
	for i := range b.paths {
		if b.paths[i] == a {
			b.paths = append(b.paths[:i], b.paths[i+1:]...)
			break
		}
	}
}

func (h heap) delete(s *step) (n heap) {
	for i := range h {
		if h[i] == s {
			return append(h[:i], h[i+1:]...)
		}
	}
	return h
}

func (h heap) in(s *step) bool {
	for _, hs := range h {
		if hs == s {
			return true
		}
	}

	return false
}

func (h heap) print() {
	fmt.Printf("[")
	for i := range h {
		fmt.Printf("%s,", h[i].name)
	}
	fmt.Println("]")
}
