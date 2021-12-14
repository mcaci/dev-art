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

func main() {
	day13()
}

func day13() {
	f, err := os.Open("day13")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	pointsMap := make(map[origamiCoord]struct{})
	var instructions []origamiInstruction
	var lastLine bool
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
		case strings.Contains(l, ","):
			i := strings.Index(l, ",")
			x, err := strconv.Atoi(l[:i])
			if err != nil {
				log.Fatal(err)
			}
			y, err := strconv.Atoi(l[i+1:])
			if err != nil {
				log.Fatal(err)
			}
			pointsMap[origamiCoord{x: x, y: y}] = struct{}{}
		case strings.Contains(l, "="):
			i := strings.Index(l, "=")
			n, err := strconv.Atoi(l[i+1:])
			if err != nil {
				log.Fatal(err)
			}
			switch l[i-1] {
			case 'x':
				instructions = append(instructions, origamiInstruction{coord: n, fold: foldOnX})
			case 'y':
				instructions = append(instructions, origamiInstruction{coord: n, fold: foldOnY})
			default:
				log.Fatal("unexpected letter", l[i-1])
			}
		}
	}
	for i, instr := range instructions {
		tmpPointsMap := instr.fold(instr.coord, pointsMap)
		fmt.Println(i, len(pointsMap), len(tmpPointsMap))
		pointsMap = tmpPointsMap
	}
	var h, w int
	for k := range pointsMap {
		if h < k.x {
			h = k.x
		}
		if w < k.y {
			w = k.y
		}
	}
	code := make([][]byte, w+1)
	for i := range code {
		code[i] = make([]byte, h+1)
		for j := range code[i] {
			code[i][j] = ' '
		}
	}
	for k := range pointsMap {
		code[k.y][k.x] = '*'
	}
	for i := range code {
		for j := range code[i] {
			fmt.Print(string(code[i][j]))
		}
		fmt.Println()
	}
}

type origamiCoord struct{ x, y int }

type origamiInstruction struct {
	coord int
	fold  func(int, map[origamiCoord]struct{}) map[origamiCoord]struct{}
}

func foldOnX(x int, in map[origamiCoord]struct{}) map[origamiCoord]struct{} {
	out := make(map[origamiCoord]struct{})
	for k := range in {
		if k.x < x {
			out[k] = struct{}{}
			continue
		}
		out[origamiCoord{x: 2*x - k.x, y: k.y}] = struct{}{}
	}
	return out
}

func foldOnY(y int, in map[origamiCoord]struct{}) map[origamiCoord]struct{} {
	out := make(map[origamiCoord]struct{})
	for k := range in {
		if k.y < y {
			out[k] = struct{}{}
			continue
		}
		out[origamiCoord{x: k.x, y: 2*y - k.y}] = struct{}{}
	}
	return out
}
