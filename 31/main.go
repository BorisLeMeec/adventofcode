package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	lenVersion = 3
	lenType    = 3

	typeLiteral = 4

	lenghtType  = 6
	bigLenth    = 15
	shortLength = 11
)

type result struct {
	version uint64
	t       uint64
	content interface{}
}

func main() {
	f, _ := os.ReadFile("input.txt")
	hexStr := strings.Split(string(f), "\n")[0]

	fmt.Println(getSumVersion(asBits(hexStr)))
}

func getSumVersion(bits string) uint64 {
	result, _ := parse(bits)

	return sumV(result)
}

func sumV(r result) (res uint64) {
	res += r.version
	if c, ok := r.content.([]interface{}); ok {
		for i := range c {
			r, _ := c[i].(result)
			res += sumV(r)
		}
	}
	return
}

func parse(bits string) (r result, read uint64) {
	var cursor = uint64(0)
	r.version = getVersion(bits)
	r.t = getType(bits)
	if r.t == typeLiteral {
		r.content, cursor = getLiteral(bits)
		read += cursor
	} else {
		if bits[lenghtType] == '1' {
			content, c := readSubPacketByNumbers(bits)
			r.content = content
			read += c
		} else {
			content, c := readSubPacketByLength(bits)
			r.content = content
			read += c
		}
	}
	return r, read
}

func readSubPacketByLength(bits string) (c []interface{}, cursor uint64) {
	l, st := getSubPacketOrLengthAndStart(bigLenth, bits)
	subBits := bits[st:]
	for x := 0; cursor < l; x++ {
		e, readed := parse(subBits[cursor:])
		c = append(c, e)
		cursor += readed
	}
	return c, cursor + st
}

func readSubPacketByNumbers(bits string) (c []interface{}, cursor uint64) {
	nb, st := getSubPacketOrLengthAndStart(shortLength, bits)
	c = make([]interface{}, nb)
	for x := 0; uint64(x) < nb; x++ {
		c[x], cursor = parse(bits[st:])
		st += cursor
	}
	return c, st
}

func getSubPacketOrLengthAndStart(lengthT int, bits string) (uint64, uint64) {
	b, _ := strconv.ParseUint(bits[lenghtType+1:lenghtType+lengthT+1], 2, 64)
	return b, uint64(lenghtType + lengthT + 1)
}

func getLiteral(bits string) (s string, i uint64) {
	over := false
	i = lenType + lenVersion
	nb := ""
	for !over {
		if bits[i] == '0' {
			over = true
		}
		nb += bits[i+1 : i+5]
		i += 5
	}
	return s, i
}

func getVersion(bits string) uint64 {
	v, _ := strconv.ParseUint(bits[:lenVersion], 2, 64)
	return v
}

func getType(bits string) uint64 {
	v, _ := strconv.ParseUint(bits[lenVersion:lenVersion+lenType], 2, 64)
	return v
}

func asBits(valHex string) string {
	str := ""
	for _, b := range valHex {
		switch b {
		case '0':
			str += "0000"
		case '1':
			str += "0001"
		case '2':
			str += "0010"
		case '3':
			str += "0011"
		case '4':
			str += "0100"
		case '5':
			str += "0101"
		case '6':
			str += "0110"
		case '7':
			str += "0111"
		case '8':
			str += "1000"
		case '9':
			str += "1001"
		case 'A', 'a':
			str += "1010"
		case 'B', 'b':
			str += "1011"
		case 'C', 'c':
			str += "1100"
		case 'D', 'd':
			str += "1101"
		case 'E', 'e':
			str += "1110"
		case 'F', 'f':
			str += "1111"
		}
	}
	return str
}
