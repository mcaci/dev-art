package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

// complexity O(n*m) where n is the number of lines in the file and m is the avg length of the lines
// read n lines plus one more for the counts
func day3Part1() {
	f, err := os.Open("day3")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	var lastCall bool
	var count []uint
	var total uint
	for !lastCall {
		line, err := r.ReadString('\n')
		switch err {
		case nil:
			line = line[:len(line)-1]
		case io.EOF:
			lastCall = true
		default:
			log.Fatal(err)
		}
		if count == nil { // if before err we get the extra space for the '\n' char
			count = make([]uint, len(line))
		}
		for j := range line {
			switch line[j] {
			case '0':
				count[j]++
			default:
				continue
			}
		}
		total++
	}
	var gamma []byte
	for i := range count {
		if count[i] >= (total / 2) {
			gamma = append(gamma, '0')
			continue
		}
		gamma = append(gamma, '1')
	}
	g, err := strconv.ParseUint(string(gamma), 2, 0)
	if err != nil {
		log.Println(err)
	}
	mask := uint64(math.Exp2(float64(len(count)))) - 1
	e := ^g & mask
	// fmt.Println(total, count, string(gamma), g, mask, ^g&mask)
	fmt.Println(g * e)
}
