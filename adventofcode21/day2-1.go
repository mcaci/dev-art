package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func day2Part1() {
	b, err := os.ReadFile("day2")
	if err != nil {
		log.Fatal(err)
	}
	var hPos, depth, start int
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
		case "down":
			depth += n
		case "up":
			depth -= n
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
	case "down":
		depth += n
	case "up":
		depth -= n
	default:
		log.Fatal(s, "is not recognised")
	}
	fmt.Println(hPos, depth, hPos*depth)
}
