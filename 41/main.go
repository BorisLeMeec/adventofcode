package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type player struct {
	pos   uint64
	score uint64
}

func main() {
	f, _ := os.ReadFile("input.txt")

	part01(string(f))
}

func part01(s string) uint64 {
	lines := strings.Split(s, "\n")
	players := make([]player, len(lines))
	for i := range lines {
		str := strings.Split(lines[i], ": ")[1]
		players[i].pos, _ = strconv.ParseUint(str, 10, 64)
	}
	dice := play(players)
	if players[0].score > players[1].score {
		fmt.Println(players[1].score * dice)
	} else {
		fmt.Println(players[0].score * dice)
	}
	return 0
}

func play(players []player) (dice uint64) {
	turn := uint64(0)
	for !won(players) {
		sum := uint64(0)
		for x := 0; x < 3; x++ {
			sum += dice%100 + 1
			dice++
		}
		nPos := (players[turn].pos + sum) % 10
		if nPos == 0 {
			nPos = 10
		}
		players[turn].pos = nPos
		players[turn].score += players[turn].pos
		turn = (turn + 1) % uint64(len(players))
	}
	return dice
}

func won(players []player) bool {
	for i := range players {
		if players[i].score >= 1000 {
			return true
		}
	}
	return false
}
