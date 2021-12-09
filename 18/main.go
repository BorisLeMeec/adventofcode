package main

import (
	"awesomeProject/common"
	"fmt"
	"os"
	"sort"
	"strings"
)

type point struct {
	weight  int
	checked bool
}

type basinList []int

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(f), "\n")

	arr := make([][]point, len(lines))
	for i, l := range lines {
		arr[i] = make([]point, len(l))
		for j, c := range l {
			arr[i][j].weight = common.ToInt(string(c))
		}
	}
	basins := make(basinList, 3)

	for i := range arr {
		for j := range arr[i] {
			b := explore(&arr, i, j)
			if b > basins[0] {
				basins[0] = b
			} else if b > basins[1] {
				basins[1] = b
			} else if b > basins[2] {
				basins[2] = b
			}
			sort.Ints(basins)
		}
	}

	fmt.Println(basins[0] * basins[1] * basins[2])
}

func explore(local *[][]point, i, j int) (size int) {
	if !isLowest(*local, i, j) {
		return
	}
	(*local)[i][j].checked = true
	size++
	type candidate struct {
		i, j int
	}
	var candidates []candidate
	if i > 0 {
		size += explore(local, i-1, j)
		candidates = append(candidates, candidate{i - 1, j})
	}
	if i < len(*local)-1 {
		candidates = append(candidates, candidate{i + 1, j})
	}
	if j < len((*local)[i])-1 {
		candidates = append(candidates, candidate{i, j + 1})
	}
	if j > 0 {
		size += explore(local, i, j-1)
		candidates = append(candidates, candidate{i, j - 1})
	}
	sort.Slice(candidates, func(a, b int) bool {
		return (*local)[candidates[a].i][candidates[a].j].weight < (*local)[candidates[b].i][candidates[b].j].weight
	})
	for _, c := range candidates {
		size += explore(local, c.i, c.j)
	}
	return size
}

func isLowest(arr [][]point, i, j int) bool {
	if arr[i][j].checked {
		return false
	}
	toTest := arr[i][j].weight
	if toTest == 9 {
		return false
	}
	if i > 0 && !arr[i-1][j].checked {
		if toTest > arr[i-1][j].weight {
			return false
		}
	}
	if j > 0 && !arr[i][j-1].checked {
		if toTest > arr[i][j-1].weight {
			return false
		}
	}
	if j < len(arr[i])-1 && !arr[i][j+1].checked {
		if toTest > arr[i][j+1].weight {
			return false
		}
	}
	if i < len(arr)-1 && !arr[i+1][j].checked {
		if toTest > arr[i+1][j].weight {
			return false
		}
	}
	return true
}
