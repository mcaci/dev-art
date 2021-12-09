package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main_day8() {
	day8part2Conc()
}

func day8part2Conc() {
	f, err := os.Open("day8")
	if err != nil {
		log.Print(err)
	}
	r := bufio.NewReader(f)
	var lastLine bool
	outC := make(chan int)
	for !lastLine {
		line, err := r.ReadString('\n')
		switch err {
		case nil:
			go processLine(line, outC)
		case io.EOF:
			go processLine(line, outC)
			lastLine = true
			close(outC)
		default:
			log.Print(err)
		}
	}
	for n := range outC {
		fmt.Print(n, "-")
	}
}

func processLine(line string, outC chan<- int) {
	outC <- len(line)
}

// complexity is linear on the size of the input file
func day8part2() {
	f, err := os.Open("day8")
	if err != nil {
		log.Print(err)
	}
	r := bufio.NewReader(f)
	var lastLine bool
	var sum int
	for !lastLine {
		line, err := r.ReadString('\n')
		var code, digits string
		switch err {
		case nil:
			pipeID := strings.Index(line, "|")
			code = line[:pipeID]
			digits = line[pipeID+1 : len(line)-1]
		case io.EOF:
			pipeID := strings.Index(line, "|")
			code = line[:pipeID]
			digits = line[pipeID+1:]
			lastLine = true
		default:
			log.Print(err)
		}
		chart := createBarPositionToLetterChart(code)
		// nMap = createDigitToLetterMap(codes, chart)
		nStr := mapDigitScreenNumberToString(digits, chart)
		n, err := strconv.Atoi(nStr)
		if err != nil {
			log.Print(err)
		}
		// fmt.Println(code, digits, n)
		sum += n
	}
	fmt.Println(sum)
}

// complexity is linear on the size of the input "code"
func createBarPositionToLetterChart(code string) map[uint8]rune {
	codes := strings.Fields(code)
	counts := make(map[rune]uint)
	for _, c := range codes {
		for _, l := range c {
			counts[l]++
		}
	}
	chart := make(map[uint8]rune)
	for k, v := range counts {
		switch v {
		case 4:
			// isolate lower left bar of the digit display (id 4)
			chart[4] = k
		case 6:
			// isolate top left bar of the digit display (id 1)
			chart[1] = k
		case 9:
			// isolate lower right bar of the digit display (id 5) // first part of 0
			chart[5] = k
		}
	}
	nMap := make(map[uint8]string)
	// list easy numbers
	for _, c := range codes {
		switch len(c) {
		case 2:
			nMap[1] = c
		case 3:
			nMap[7] = c
		case 4:
			nMap[4] = c
		case 7:
			nMap[8] = c
		}
	}

	// isolate upper right bar of the digit display (id 2) // second part of 0
	upRightBarID := strings.IndexFunc(nMap[1], func(r rune) bool { return r != chart[5] })
	chart[2] = rune(nMap[1][upRightBarID])
	// isolate upper bar of the digit display (id 0) // last part of 7 which includes the parts of 0
	upBarID := strings.IndexFunc(nMap[7], func(r rune) bool { return r != chart[5] && r != chart[2] })
	chart[0] = rune(nMap[7][upBarID])
	// isolate middle bar of the digit display (id 3) // last part of 4
	midBarID := strings.IndexFunc(nMap[4], func(r rune) bool { return r != chart[5] && r != chart[2] && r != chart[1] })
	chart[3] = rune(nMap[4][midBarID])
	// isolate lower bar of the digit display (id 6)
	lowBarID := strings.IndexFunc(nMap[8], func(r rune) bool {
		return r != chart[0] && r != chart[1] && r != chart[2] && r != chart[3] && r != chart[4] && r != chart[5]
	})
	chart[6] = rune(nMap[8][lowBarID])
	return chart
}

// complexity is linear on the size of the input "digit"
func mapDigitScreenNumberToString(digit string, chart map[uint8]rune) string {
	var s []string
	for _, n := range strings.Fields(digit) {
		switch {
		case len(n) == 6 && !strings.ContainsRune(n, chart[3]):
			s = append(s, "0")
		case len(n) == 2:
			s = append(s, "1")
		case len(n) == 5 && !strings.ContainsRune(n, chart[1]) && !strings.ContainsRune(n, chart[5]):
			s = append(s, "2")
		case len(n) == 5 && !strings.ContainsRune(n, chart[1]) && !strings.ContainsRune(n, chart[4]):
			s = append(s, "3")
		case len(n) == 4:
			s = append(s, "4")
		case len(n) == 5 && !strings.ContainsRune(n, chart[2]) && !strings.ContainsRune(n, chart[4]):
			s = append(s, "5")
		case len(n) == 6 && !strings.ContainsRune(n, chart[2]):
			s = append(s, "6")
		case len(n) == 3:
			s = append(s, "7")
		case len(n) == 7:
			s = append(s, "8")
		case len(n) == 6 && !strings.ContainsRune(n, chart[4]):
			s = append(s, "9")
		}
	}
	return strings.Join(s, "")
}

// nice for information but not useful
func createDigitToLetterMap(codes []string, chart map[uint8]rune) map[uint8]string {
	nMap := make(map[uint8]string)
	for _, c := range codes {
		switch {
		case len(c) == 6 && !strings.ContainsRune(c, chart[3]):
			nMap[0] = c
		case len(c) == 2:
			nMap[1] = c
		case len(c) == 5 && !strings.ContainsRune(c, chart[1]) && !strings.ContainsRune(c, chart[5]):
			nMap[2] = c
		case len(c) == 5 && !strings.ContainsRune(c, chart[1]) && !strings.ContainsRune(c, chart[4]):
			nMap[3] = c
		case len(c) == 4:
			nMap[4] = c
		case len(c) == 5 && !strings.ContainsRune(c, chart[2]) && !strings.ContainsRune(c, chart[4]):
			nMap[5] = c
		case len(c) == 6 && !strings.ContainsRune(c, chart[2]):
			nMap[6] = c
		case len(c) == 3:
			nMap[7] = c
		case len(c) == 7:
			nMap[8] = c
		case len(c) == 6 && !strings.ContainsRune(c, chart[4]):
			nMap[9] = c
		}
	}
	return nMap
}

func day8part1() {
	f, err := os.Open("day8")
	if err != nil {
		log.Print(err)
	}
	r := bufio.NewReader(f)
	var lastLine bool
	var count int
	for !lastLine {
		line, err := r.ReadString('\n')
		switch err {
		case nil:
			pipeID := strings.Index(line, "|")
			line = line[pipeID+1 : len(line)-1]
		case io.EOF:
			pipeID := strings.Index(line, "|")
			line = line[pipeID+1:]
			lastLine = true
		default:
			log.Print(err)
		}
		fields := strings.Fields(line)
		fmt.Println(fields)
		for i := range fields {
			switch len(fields[i]) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	fmt.Println(count)
}
