package main

import (
	"fmt"
	"os"
	"strings"

	com "awesomeProject/common"
)

const (
	len1 = 2
	len4 = 4
	len7 = 3

	lenGroup6 = 6
	lenGroup5 = 5
)

const (
	T  = 0
	TL = 1
	TR = 2
	M  = 3
	BL = 4
	BR = 5
	B  = 6
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(f), "\n")

	var t int
	for _, l := range lines {
		parts := strings.Split(l, " | ")
		entriesPart := parts[0]
		entries := strings.Split(entriesPart, " ")
		outputsPart := parts[1]
		outputs := strings.Split(outputsPart, " ")

		wires := getWires(entries)
		fmt.Printf("%s\n", string(wires))
		t += getCode(wires, outputs)
	}
	fmt.Println(t)
}

func getCode(wires []rune, outputs []string) int {
	var numberStr string
	for _, o := range outputs {
		var n int
		switch {
		case isZero(wires, o):
			n = 0
		case isOne(wires, o):
			n = 1
		case isTwo(wires, o):
			n = 2
		case isThree(wires, o):
			n = 3
		case isFour(wires, o):
			n = 4
		case isFive(wires, o):
			n = 5
		case isSix(wires, o):
			n = 6
		case isSeven(wires, o):
			n = 7
		case isHeight(wires, o):
			n = 8
		case isNine(wires, o):
			n = 9
		}
		numberStr = fmt.Sprintf("%s%d", numberStr, n)
	}
	return com.ToInt(numberStr)
}

func getWires(entries []string) (wires []rune) {
	var (
		one, four, seven string
		group5           []string
		group6           []string
	)
	wires = make([]rune, 7)
	for _, o := range entries {
		switch len(o) {
		case len1:
			one = o
		case len4:
			four = o
		case len7:
			seven = o
		case lenGroup5:
			group5 = append(group5, o)
		case lenGroup6:
			group6 = append(group6, o)
		}
	}
	wires[T] = common(append(group5, seven)...)
	wires[M] = common(append(group5, four)...)
	wires[B] = commonBut([]rune{wires[T], wires[M]}, group5...)
	wires[TL] = uniqueBut([]rune{wires[M]}, append([]string{}, four, one)...)
	wires[BL] = uniqueBut([]rune{wires[TL]}, group5...)
	wires[BR] = commonBut([]rune{wires[T], wires[B], wires[TL]}, group6...)
	wires[TR] = commonBut([]rune{wires[BR]}, []string{one}...)
	return
}

func common(str ...string) rune {
	var wires = make(map[rune]int)
	for _, s := range str {
		for _, l := range s {
			wires[l]++
		}
	}

	for k, w := range wires {
		if w == len(str) {
			return k
		}
	}
	return 0
}

func commonBut(rr []rune, str ...string) rune {
	var wires = make(map[rune]int)
	for _, s := range str {
		for _, l := range s {
			if inRune(l, rr) {
				continue
			}
			wires[l]++
		}
	}

	for k, w := range wires {
		if w == len(str) {
			return k
		}
	}
	return 0
}

func unique(str ...string) rune {
	var wires = make(map[rune]int)
	for _, s := range str {
		for _, l := range s {
			wires[l]++
		}
	}

	for k, w := range wires {
		if w == 1 {
			return k
		}
	}
	return 0
}

func uniqueBut(rr []rune, str ...string) rune {
	var wires = make(map[rune]int)
	for _, s := range str {
		for _, l := range s {
			if inRune(l, rr) {
				continue
			}
			wires[l]++
		}
	}

	for k, w := range wires {
		if w == 1 {
			return k
		}
	}
	return 0
}

func inRune(ru rune, rr []rune) bool {
	for _, r := range rr {
		if ru == r {
			return true
		}
	}
	return false
}

func isZero(wires []rune, t string) bool {
	return len(t) == 6 && strings.ContainsRune(t, wires[T]) &&
		strings.ContainsRune(t, wires[B]) &&
		strings.ContainsRune(t, wires[TR]) &&
		strings.ContainsRune(t, wires[TL]) &&
		strings.ContainsRune(t, wires[BR]) &&
		strings.ContainsRune(t, wires[B]) &&
		strings.ContainsRune(t, wires[BL])

}

func isOne(wires []rune, t string) bool {
	return len(t) == 2 && strings.ContainsRune(t, wires[TR]) &&
		strings.ContainsRune(t, wires[BR])

}

func isTwo(wires []rune, t string) bool {
	return len(t) == 5 && strings.ContainsRune(t, wires[T]) &&
		strings.ContainsRune(t, wires[M]) &&
		strings.ContainsRune(t, wires[B]) &&
		strings.ContainsRune(t, wires[TR]) &&
		strings.ContainsRune(t, wires[BL])
}

func isThree(wires []rune, t string) bool {
	return len(t) == 5 && strings.ContainsRune(t, wires[T]) &&
		strings.ContainsRune(t, wires[M]) &&
		strings.ContainsRune(t, wires[B]) &&
		strings.ContainsRune(t, wires[TR]) &&
		strings.ContainsRune(t, wires[BR])
}

func isFour(wires []rune, t string) bool {
	return len(t) == 4 && strings.ContainsRune(t, wires[TR]) &&
		strings.ContainsRune(t, wires[TL]) &&
		strings.ContainsRune(t, wires[M]) &&
		strings.ContainsRune(t, wires[BR])
}

func isFive(wires []rune, t string) bool {
	return len(t) == 5 && strings.ContainsRune(t, wires[T]) &&
		strings.ContainsRune(t, wires[M]) &&
		strings.ContainsRune(t, wires[B]) &&
		strings.ContainsRune(t, wires[TL]) &&
		strings.ContainsRune(t, wires[BR])
}

func isSix(wires []rune, t string) bool {
	return len(t) == 6 && strings.ContainsRune(t, wires[T]) &&
		strings.ContainsRune(t, wires[M]) &&
		strings.ContainsRune(t, wires[B]) &&
		strings.ContainsRune(t, wires[TL]) &&
		strings.ContainsRune(t, wires[BL]) &&
		strings.ContainsRune(t, wires[BR])
}

func isSeven(wires []rune, t string) bool {
	return len(t) == 3 && strings.ContainsRune(t, wires[TR]) &&
		strings.ContainsRune(t, wires[BR]) &&
		strings.ContainsRune(t, wires[T])
}

func isHeight(wires []rune, t string) bool {
	return len(t) == 7 && strings.ContainsRune(t, wires[T]) &&
		strings.ContainsRune(t, wires[B]) &&
		strings.ContainsRune(t, wires[M]) &&
		strings.ContainsRune(t, wires[TR]) &&
		strings.ContainsRune(t, wires[TL]) &&
		strings.ContainsRune(t, wires[BR]) &&
		strings.ContainsRune(t, wires[BL])
}
func isNine(wires []rune, t string) bool {
	return len(t) == 6 && strings.ContainsRune(t, wires[T]) &&
		strings.ContainsRune(t, wires[B]) &&
		strings.ContainsRune(t, wires[M]) &&
		strings.ContainsRune(t, wires[TR]) &&
		strings.ContainsRune(t, wires[TL]) &&
		strings.ContainsRune(t, wires[BR])
}
