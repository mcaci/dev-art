package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func day1Part1() {
	b, err := os.ReadFile("day1")
	if err != nil {
		log.Fatal(err)
	}
	var start int
	var depths []int
	for i, c := range b {
		if c != '\n' {
			continue
		}
		s := string(b[start:i])
		fmt.Println(s)
		d, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, d)
		start = i + 1
	}
	s := string(b[start:])
	fmt.Println(s)
	d, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	depths = append(depths, d)
	fmt.Println(depths)
	var count int
	for i := range depths {
		if i+1 == len(depths) {
			break
		}
		if depths[i+1] <= depths[i] {
			continue
		}
		count++
	}
	fmt.Println(count)
}
