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
	lines := strings.Split(string(f), "\n")

	result := reduce(lines[0])
	for _, l := range lines[1:] {
		result = reduce("[" + result + "," + l + "]")
	}

	fmt.Println(calcMagn(result))
}

func calcMagn(str string) uint64 {
	compile, _ := regexp.Compile(`([[][0-9]*,[0-9]*[]])`)
	for {
		subs := compile.FindAllStringIndex(str, -1)
		if len(subs) == 0 {
			break
		}
		numbers := strings.Split(str[subs[0][0]+1:subs[0][1]-1], ",")
		l, _ := strconv.Atoi(numbers[0])
		r, _ := strconv.Atoi(numbers[1])
		str = str[:subs[0][0]] + strconv.Itoa(3*l+2*r) + str[subs[0][1]:]
	}
	output, _ := strconv.ParseUint(str, 10, 64)
	return output
}

func reduce(str string) string {
	for {
		for explode(&str) {
		}
		if !split(&str) {
			break
		}
	}
	return str
}

func explode(str *string) (did bool) {
	compile, _ := regexp.Compile(`([[][0-9]*,[0-9]*[]])`)
	keep := true
	for keep {
		keep = false
		subs := compile.FindAllStringIndex(*str, -1)
		if len(subs) == 0 {
			break
		}
		for _, subI := range subs {
			if isDepth4(*str, subI[0]) {
				numbers := strings.Split((*str)[subI[0]+1:subI[1]-1], ",")
				l, _ := strconv.Atoi(numbers[0])
				r, _ := strconv.Atoi(numbers[1])
				*str = (*str)[:subI[0]] + "0" + (*str)[subI[1]:]
				addLeft(str, subI[0], l)
				addRight(str, subI[0]+2, r)

				keep = true
				break
			}
		}
	}
	return
}

func addLeft(str *string, i int, nb int) {
	compile, _ := regexp.Compile(`[0-9]+`)
	subs := compile.FindAllStringIndex((*str)[:i], -1)
	if len(subs) == 0 {
		return
	}
	baseNb, _ := strconv.Atoi((*str)[subs[len(subs)-1][0]:subs[len(subs)-1][1]])
	*str = (*str)[:subs[len(subs)-1][0]] + strconv.Itoa(baseNb+nb) + (*str)[subs[len(subs)-1][1]:]
}

func addRight(str *string, i int, nb int) {
	compile, _ := regexp.Compile(`[0-9]+`)
	subs := compile.FindAllStringIndex((*str)[i:], -1)
	if len(subs) == 0 {
		return
	}
	baseNb, _ := strconv.Atoi((*str)[i+subs[0][0] : i+subs[0][1]])
	*str = (*str)[:i+subs[0][0]] + strconv.Itoa(baseNb+nb) + (*str)[i+subs[0][1]:]
}

func split(str *string) bool {
	compile, _ := regexp.Compile(`[1-9][0-9]+`)
	subs := compile.FindAllStringIndex(*str, -1)
	if len(subs) == 0 {
		return false
	}
	subI := subs[0]
	big, _ := strconv.Atoi((*str)[subI[0]:subI[1]])
	newPair := fmt.Sprintf("[%d,%d]", big/2, int(math.Round(float64(big)/2.0)))
	*str = (*str)[:subI[0]] + newPair + (*str)[subI[1]:]
	return true
}

func isDepth4(str string, i int) bool {
	return strings.Count(str[:i], "[")-strings.Count(str[:i], "]") >= 4
}
