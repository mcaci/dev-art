package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

func main_day14() {
	day14Algo1(10)
	// day14Algo2(40) // can't do... too time and mem consuming and gets killed
	day14Algo2(40)
}

// algo1 just manipulate the string
func day14Algo1(steps int) {
	f, err := os.Open("day14")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	var rules []PairInsertionRule
	var polyTmpl []byte
	var lastLine bool
	for !lastLine {
		l, err := r.ReadBytes('\n')
		switch err {
		case nil:
			l = l[:len(l)-1]
		case io.EOF:
			lastLine = true
		default:
			log.Fatal(err)
		}
		switch {
		case bytes.Contains(l, []byte{'-'}):
			rules = append(rules, PairInsertionRule{pair: l[:2], letter: l[len(l)-1]})
		case len(l) == 0:
		default:
			polyTmpl = l
		}
	}

	for n := 0; n < steps; n++ {
		var toInsert []byte
		for i := range polyTmpl {
			if i == 0 {
				continue
			}
			b, err := findLetter(polyTmpl[i-1:i+1], rules)
			if err != nil {
				log.Fatal(err)
			}
			toInsert = append(toInsert, b)
		}
		poly := make([]byte, len(polyTmpl)+len(toInsert))
		for i := range poly {
			switch i % 2 {
			case 0:
				poly[i] = polyTmpl[i/2]
			case 1:
				poly[i] = toInsert[i/2]
			}
		}
		polyTmpl = poly
	}
	count := make(map[byte]int)
	for _, b := range polyTmpl {
		count[b]++
	}
	// MaxInt32 is ok for algo1 but after a certain step must use bigger max: MaxInt64
	min, max := math.MaxInt32, 0
	for _, v := range count {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	fmt.Println(max-min, count)
}

// algo2 maybe count the pairs and modify the count map as they form?
func day14Algo2(steps int) {
	f, err := os.Open("day14")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	var rules []PairInsertionRule
	polyCount := make(map[string]int)
	var lastLine bool
	var startPolyTmpl string
	for !lastLine {
		l, err := r.ReadString('\n')
		switch err {
		case nil:
			l = l[:len(l)-1]
		case io.EOF:
			lastLine = true
		default:
			log.Fatal(err)
		}
		switch {
		case strings.Contains(l, "-"):
			rules = append(rules, PairInsertionRule{pair: []byte(l[:2]), letter: l[len(l)-1]})
		case len(l) == 0:
		default:
			startPolyTmpl = l
			for i := range l {
				if i == 0 {
					continue
				}
				polyCount[l[i-1:i+1]]++
			}
		}
	}
	for n := 0; n < steps; n++ {
		finalPolyCount := make(map[string]int)
		for k, v := range polyCount {
			b, err := findLetter([]byte(k), rules)
			if err != nil {
				log.Fatal(err)
			}
			finalPolyCount[string([]byte{k[0], b})] += v
			finalPolyCount[string([]byte{b, k[1]})] += v
		}
		polyCount = finalPolyCount
	}
	count := make(map[byte]int)
	for k, v := range polyCount {
		count[k[0]] += v
		count[k[1]] += v
	}
	// error: used MaxInt32 instead of MaxInt64 (min didn't work)
	min, max := math.MaxInt64, 0
	for k := range count {
		count[k] /= 2
		switch k {
		case startPolyTmpl[0], startPolyTmpl[len(startPolyTmpl)-1]:
			count[k]++
		}
		if count[k] < min {
			min = count[k]
		}
		if count[k] > max {
			max = count[k]
		}
	}
	fmt.Println(max-min, count)
	// fmt.Println(polyCount)
}

func findLetter(pair []byte, rules []PairInsertionRule) (byte, error) {
	for _, rule := range rules {
		if !(rule.pair[0] == pair[0] && rule.pair[1] == pair[1]) {
			continue
		}
		return rule.letter, nil
	}
	return 0, fmt.Errorf("no letter found for pair %s and rules %v", pair, rules)
}

type PairInsertionRule struct {
	pair   []byte
	letter byte
}

func (r PairInsertionRule) String() string {
	return fmt.Sprintf("{[%s] %s}", r.pair, string(r.letter))
}
