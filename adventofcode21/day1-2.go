package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func day1Part2() {
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
		if i+3 == len(depths) {
			break
		}
		if sumOfThree(depths[i+1:i+4]) <= sumOfThree(depths[i:i+3]) {
			continue
		}
		count++
	}
	fmt.Println(count)
}

func sumOfThree(v []int) (sum int) {
	for i := range v {
		sum += v[i]
	}
	return sum
}
