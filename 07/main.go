package main

import (
	"awesomeProject/common"
	"fmt"
	"os"
	"strings"
)

const lenLine = 5

func main() {
	f, _ := os.ReadFile("input.txt")

	ss := strings.Split(string(f), "\n")

	bingoNumberstr := strings.Split(ss[0], ",")
	var bingoNumber []int
	for _, n := range bingoNumberstr {
		bingoNumber = append(bingoNumber, common.ToInt(n))
	}
	bingoSheets := getBingoShets(ss[2:])
	for _, n := range bingoNumber {
		for bsi := range bingoSheets {
			bingoSheets[bsi].mark(n)
			if bingoSheets[bsi].won() {
				fmt.Printf("%d\n", n*bingoSheets[bsi].sumNonMarked())
				return
			}
		}
	}
	fmt.Println("error")
}

func getBingoShets(ss []string) (bs []bingoSheet) {
	nbSheets := 0
	for nbSheets*6 < len(ss) {
		var b = bingoSheet{}
		for i := 0; i < lenLine; i++ {
			line := strings.Replace(ss[(nbSheets*6)+i], "  ", " ", -1)
			line = strings.Trim(line, " ")
			nbrs := strings.Split(line, " ")
			for j, l := range nbrs {
				b[i][j] = bingoCase{
					n: common.ToInt(l),
					d: false,
				}
			}
		}
		nbSheets++
		bs = append(bs, b)
	}
	return
}

type bingoSheet [lenLine][lenLine]bingoCase

type bingoCase struct {
	n int
	d bool
}

func (s bingoSheet) won() bool {
	for i := 0; i < lenLine; i++ {
		l := 0
		for j := 0; j < lenLine; j++ {
			if s[i][j].d == true {
				l++
			}
			if l == lenLine {
				return true
			}
		}
	}
	for j := 0; j < lenLine; j++ {
		h := 0
		for i := 0; i < lenLine; i++ {
			if s[i][j].d == true {
				h++
			}
			if h == lenLine {
				return true
			}
		}
	}

	return false
}

func (s *bingoSheet) mark(n int) {
	for i := 0; i < lenLine; i++ {
		for j := 0; j < lenLine; j++ {
			if s[i][j].n == n {
				s[i][j].d = true
			}
		}
	}
}

func (s bingoSheet) sumNonMarked() int {
	var t = 0
	for i := 0; i < lenLine; i++ {
		for j := 0; j < lenLine; j++ {
			if s[i][j].d == false {
				t += s[i][j].n
			}
		}
	}
	return t
}
