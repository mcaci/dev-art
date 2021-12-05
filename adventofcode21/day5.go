package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main_day5() {
	part2 := flag.Bool("part2", false, "trigger part2")
	flag.Parse()
	day5(*part2)
}

const side = 1000

func day5(isPart2 bool) {
	f, err := os.Open("day5")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	field := make([]uint, side*side)
	var lastLine bool
	for !lastLine {
		line, err := r.ReadString('\n')
		switch err {
		case nil:
		case io.EOF:
			lastLine = true
		default:
			log.Fatal(err)
		}
		ns := strings.FieldsFunc(line, func(r rune) bool { return !unicode.IsNumber(r) })
		if !isPart2 && ns[0] != ns[2] && ns[1] != ns[3] {
			// log.Print("skipping ", ns)
			continue
		}
		c0 := NewCoord(ns[0], ns[1])
		c1 := NewCoord(ns[2], ns[3])
		var idxs []int
		switch {
		case c0.x == c1.x, c0.y == c1.y:
			idxs = fieldIndexes(c0, c1)
		default:
			idxs = fieldIndexesDiag(c0, c1)
		}
		// big error: took first index i instead of actual value idx
		// which resulted in filling all first fields values
		for _, idx := range idxs {
			if field[idx] > 1 {
				continue
			}
			field[idx]++
		}
	}
	var count uint
	for _, f := range field {
		if f <= 1 {
			continue
		}
		count++
	}
	fmt.Println(count)
}

type Coord struct{ x, y int }

func NewCoord(x, y string) Coord {
	xN, err := strconv.Atoi(x)
	if err != nil {
		log.Fatal(err)
	}
	yN, err := strconv.Atoi(y)
	if err != nil {
		log.Fatal(err)
	}
	return Coord{xN, yN}
}

func fieldIndexes(c1, c2 Coord) (idxs []int) {
	switch {
	case c1.x == c2.x:
		start, end := c1.y, c2.y
		if c2.y < c1.y {
			start, end = end, start
		}
		for i := start; i <= end; i++ {
			idxs = append(idxs, i*side+c1.x)
		}
		return idxs
	case c1.y == c2.y:
		start, end := c1.x, c2.x
		if c2.x < c1.x {
			start, end = end, start
		}
		for i := start; i <= end; i++ {
			idxs = append(idxs, c1.y*side+i)
		}
		return idxs
	default:
		log.Fatal("unexpected situation")
		return idxs
	}
}

func fieldIndexesDiag(c1, c2 Coord) (idxs []int) {
	start, end := c1, c2
	if c2.y < c1.y {
		start, end = end, start
	}
	switch {
	case start.x < end.x:
		for i := start.x; i <= end.x; i++ {
			point := start.x + (i - start.x) + (start.y+i-start.x)*side
			idxs = append(idxs, point)
		}
		return idxs
	case start.x > end.x:
		// another error here : "for i := start.x; i <= end.x; i++"
		// while start.x > end.x... wrong direction altogether
		for i := start.x; i >= end.x; i-- {
			point := start.x + (i - start.x) + (start.y+start.x-i)*side
			idxs = append(idxs, point)
		}
		return idxs
	default:
		log.Fatal("unexpected situation for part 2")
		return idxs
	}
}
