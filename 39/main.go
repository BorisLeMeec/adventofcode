package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const timeOK = "8h05"

func main() {
	f, _ := os.ReadFile("input.txt")

	part01(string(f))
}

func part01(lines string) {
	parts := strings.Split(lines, "\n\n")

	algo := getAlgo(parts[0])
	input, lenLine := getInput(parts[1])

	for i := 0; i < 2; i++ {
		input = enhance(algo, input, lenLine+(i*4), i)
	}
	sum := 0

	for i := range input {
		if input[i] {
			sum++
		}
	}
	fmt.Println(sum)
}

func getAlgo(input string) []bool {
	source := make([]bool, len(input))

	for i := range input {
		source[i] = input[i] == '#'
	}
	return source
}

func getInput(str string) ([]bool, int) {
	inputStr := strings.Split(str, "\n")
	lenLine := len(inputStr[0])
	input := make([]bool, len(inputStr)*lenLine)
	for i := range inputStr {
		for j := range inputStr[i] {
			input[i*lenLine+j] = inputStr[i][j] == '#'
		}
	}
	return input, lenLine
}

func enhance(algo, input []bool, lenLine, iteration int) []bool {
	output := make([]bool, (len(input)/lenLine+4)*(lenLine+4)) // two more each side so line full of algo[0] appears
	for i := 0; i < (len(input)/lenLine)+4; i++ {
		for j := 0; j < lenLine+4; j++ {
			pix := getPixel(algo, input, lenLine, i-2, j-2, iteration) // -1 because enhanced image is larger by 2 (so start two before)
			output[i*(lenLine+4)+j] = pix
		}
	}
	return output
}

func getPixel(algo, input []bool, lenLine, outputY, outputX, iteration int) bool {
	source := make([]bool, 9)
	for y := -1; y < 2; y++ {
		for x := -1; x < 2; x++ {
			inputY := outputY + y
			inputX := outputX + x
			source[(y+1)*3+x+1] = getInputPixel(input, lenLine, inputY, inputX, iteration)
		}
	}
	return algo[toIndex(source)]
}

func getInputPixel(input []bool, lenLine, y, x, iteration int) bool {
	if x < 0 || y < 0 || x >= lenLine || y >= len(input)/lenLine {
		if iteration == 0 || iteration%2 == 0 {
			return false
		}
		return input[0]
	}
	return input[y*lenLine+x]
}

func toIndex(source []bool) uint64 {
	if len(source) != 9 {
		panic("wrong len for source")
	}
	binary := ""
	for i := range source {
		if source[i] {
			binary += fmt.Sprint("1")
		} else {
			binary += fmt.Sprint("0")
		}
	}
	index, _ := strconv.ParseUint(binary, 2, 64)
	return index
}
