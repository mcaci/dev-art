package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

// ASCII values: 40,41:(), 60,62:<>, 91,93:[], 123,125:{}
var errScoreMap = map[byte]int{')': 3, ']': 57, '}': 1197, '>': 25137}
var autoComplScoreMap = map[byte]int{')': 1, ']': 2, '}': 3, '>': 4}
var autoComplMap = map[byte]byte{'(': ')', '[': ']', '{': '}', '<': '>'}

func main_day10() {
	f, err := os.Open("day10")
	if err != nil {
		log.Print(err)
	}
	r := bufio.NewReader(f)
	var errScore int
	var autoComplScores []int
	var lastline bool
	for !lastline {
		l, err := r.ReadBytes('\n')
		switch err {
		case nil:
			l = l[:len(l)-1]
		case io.EOF:
			lastline = true
		default:
			log.Print(err)
		}
		var stack []byte
		var invalid bool
	linecheck:
		for _, b := range l {
			switch b {
			case '(', '<', '[', '{':
				stack = append(stack, b)
			case ')', '>', ']', '}':
				popID := len(stack) - 1
				popB := stack[popID]
				if autoComplMap[popB] == b {
					stack = stack[:popID]
					continue
				}
				errScore += errScoreMap[b]
				invalid = true
				break linecheck
			}
		}
		if invalid {
			continue
		}
		var autoComplScore int
		for j := len(stack) - 1; j >= 0; j-- {
			autoComplScore = autoComplScore*5 + autoComplScoreMap[autoComplMap[stack[j]]]
		}
		autoComplScores = append(autoComplScores, autoComplScore)
	}
	sort.Ints(autoComplScores)
	fmt.Println(errScore, autoComplScores[len(autoComplScores)/2])
}
