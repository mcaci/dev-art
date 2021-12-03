package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func day2Part2() {
	algo4()
}

// improved on algo3 using indexes instead of splitting the line (faster and less memory used)
func algo4() {
	f, err := os.Open("day2")
	if err != nil {
		log.Fatal(err)
	}
	var hPos, depth, aim int
	r := bufio.NewReader(f)
	var lastStep bool
	for !lastStep {
		line, err := r.ReadString('\n')
		switch err {
		case nil:
		case io.EOF:
			// log.Print(err)
			lastStep = true
		default:
			log.Fatal(err)
		}
		sepIdx := strings.Index(line, " ")
		lastDigitIdx := strings.LastIndexFunc(line, unicode.IsDigit)
		n, err := strconv.Atoi(line[sepIdx+1 : lastDigitIdx+1])
		if err != nil {
			log.Fatal(err)
		}
		s := line[:sepIdx]
		switch s {
		case "forward":
			hPos += n
			depth = depth + aim*n
		case "down":
			aim += n
		case "up":
			aim -= n
		default:
			log.Fatal(s, "is not recognised")
		}
	}
	fmt.Println(hPos, depth, aim, hPos*depth)
}

func algo3() {
	f, err := os.Open("day2")
	if err != nil {
		log.Fatal(err)
	}
	var hPos, depth, aim int
	r := bufio.NewReader(f)
	var lastStep bool
	for !lastStep {
		line, err := r.ReadString('\n')
		switch err {
		case nil:
		case io.EOF:
			// log.Print(err)
			lastStep = true
		default:
			log.Fatal(err)
		}
		fields := strings.Fields(line)
		n, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		s := fields[0]
		switch s {
		case "forward":
			hPos += n
			depth = depth + aim*n
		case "down":
			aim += n
		case "up":
			aim -= n
		default:
			log.Fatal(s, "is not recognised")
		}
	}
	fmt.Println(hPos, depth, aim, hPos*depth)
}

func algo2() {
	b, err := os.ReadFile("day2")
	if err != nil {
		log.Fatal(err)
	}
	var hPos, depth, aim int
	lines := bytes.Split(b, []byte{'\n'})
	for _, line := range lines {
		fields := bytes.Split(line, []byte{' '})
		s := string(fields[0])
		n, err := strconv.Atoi(string(fields[1]))
		if err != nil {
			log.Fatal(err)
		}
		switch s {
		case "forward":
			hPos += n
			depth = depth + aim*n
		case "down":
			aim += n
		case "up":
			aim -= n
		default:
			log.Fatal(s, "is not recognised")
		}
	}
	fmt.Println(hPos, depth, aim, hPos*depth)
}

func algo1() {
	b, err := os.ReadFile("day2")
	if err != nil {
		log.Fatal(err)
	}
	var hPos, depth, aim, start int
	for i, c := range b {
		if c != '\n' {
			continue
		}
		line := bytes.Split(b[start:i], []byte{' '})
		s := string(line[0])
		n, err := strconv.Atoi(string(line[1]))
		if err != nil {
			log.Fatal(err)
		}
		switch s {
		case "forward":
			hPos += n
			depth = depth + aim*n
		case "down":
			aim += n
		case "up":
			aim -= n
		default:
			log.Fatal(s, "is not recognised")
		}
		start = i + 1
	}
	line := bytes.Split(b[start:], []byte{' '})
	s := string(line[0])
	n, err := strconv.Atoi(string(line[1]))
	if err != nil {
		log.Fatal(err)
	}
	switch s {
	case "forward":
		hPos += n
		depth = depth + aim*n
	case "down":
		aim += n
	case "up":
		aim -= n
	default:
		log.Fatal(s, "is not recognised")
	}
	fmt.Println(hPos, depth, aim, hPos*depth)
}
