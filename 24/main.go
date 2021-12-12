package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type (
	cave struct {
		name  string
		t     typeCave
		paths []*cave
	}

	typeCave uint8
)

const (
	typeCaveBIG typeCave = iota
	typeCaveSMALL
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(f), "\n")

	startCave := getCaves(lines)

	fmt.Println(findExit(startCave, []string{}, false))
}

func findExit(c *cave, visited []string, twice bool) int {
	ex := 0
	visited = append(visited, c.name)
	for _, p := range c.paths {
		if p.t == typeCaveSMALL && isVisited(p.name, visited) {
			continue
		}
		if p.name == "end" {
			ex++
		} else {
			ex += findExit(p, visited, twice)
		}
	}
	for _, p := range c.paths {
		if p.name != "start" && p.name != "end" && p.t == typeCaveSMALL && isVisited(p.name, visited) && !twice {
			ex += findExit(p, visited, true)
		}
	}
	return ex
}

func isVisited(name string, visited []string) bool {
	for _, v := range visited {
		if v == name {
			return true
		}
	}
	return false
}

func getCaves(lines []string) *cave {
	c := initCaves(lines)
	createPaths(c, lines)
	return c["start"]
}

func createPaths(caves map[string]*cave, lines []string) {
	for _, l := range lines {
		cavesNames := strings.Split(l, "-")
		c1, c2 := caves[cavesNames[0]], caves[cavesNames[1]]
		c1.paths = append(c1.paths, c2)
		c2.paths = append(c2.paths, c1)
	}
}

func initCaves(lines []string) map[string]*cave {
	caves := map[string]*cave{}
	for _, l := range lines {
		cavesNames := strings.Split(l, "-")
		for _, c := range cavesNames {
			if _, ok := caves[c]; !ok {
				t := typeCaveSMALL
				if isUpper(c) {
					t = typeCaveBIG
				}
				caves[c] = &cave{
					name:  c,
					t:     t,
					paths: nil,
				}
			}
		}
	}
	return caves
}

func isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
