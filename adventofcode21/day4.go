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

func day4() {
	day4Part1()
	day4Part2()
}

type Seq []uint
type Board []Cell
type ExtraBoard struct {
	b   Board
	won bool
}

type Cell struct {
	n      uint
	called bool
}

func (s Seq) Win(b Board) bool {
	for i := range b {
		for _, n := range s {
			if n != b[i].n {
				continue
			}
			b[i].called = true
		}
	}
	for i := 0; i < len(b); i += 5 {
		found := true
		for _, cell := range b[i : i+5] {
			if cell.called {
				continue
			}
			found = false
			break
		}
		if !found {
			continue
		}
		return true
	}
	for i := 0; i < len(b)/5; i++ {
		found := true
		localHits := []bool{b[i].called, b[i+5].called, b[i+10].called, b[i+15].called, b[i+20].called}
		for _, hit := range localHits {
			if hit {
				continue
			}
			found = false
			break
		}
		if !found {
			continue
		}
		return true
	}
	return false
}

func day4Part1() {
	f, err := os.Open("day4")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	drawSeq, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	var draw Seq
	for _, item := range strings.Split(drawSeq[:len(drawSeq)-1], ",") {
		n, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal(err)
		}
		draw = append(draw, uint(n))
	}
	var boards []Board
	var lastLine bool
	var i int
	boardBuf := make([]Cell, 25)
	for !lastLine {
		boardLine, err := r.ReadString('\n')
		switch err {
		case nil:
			boardLine = boardLine[:len(boardLine)-1]
		case io.EOF:
			lastLine = true
		default:
			log.Fatal(err)
		}
		if len(boardLine) == 0 {
			continue
		}
		if i == 25 {
			boards = append(boards, boardBuf)
			i = 0
			// Error: board defined in line 91: before the for loop
			// the values of the slice were overwritten each time by the next iteration
			// must be reinitialized with a new slice
			boardBuf = make([]Cell, 25)
		}
		for _, boardElStr := range strings.Fields(boardLine) {
			boardEl, err := strconv.Atoi(boardElStr)
			if err != nil {
				log.Fatal(err)
			}
			boardBuf[i] = Cell{n: uint(boardEl)}
			i++
		}
	}
	// fmt.Println(boards)
	for i := range draw {
		if i < 4 {
			continue // min to have a full line
		}
		for j, b := range boards {
			found := draw[:i].Win(b)
			if !found {
				continue
			}
			fmt.Println(i, j, len(boards), draw[i-1])
			fmt.Println(draw[:i])
			fmt.Println(b)
			var sum uint
			for k := range b {
				if b[k].called {
					continue
				}
				sum += b[k].n
			}
			fmt.Println(sum * draw[i-1])
			return
		}
	}
}

func day4Part2() {
	f, err := os.Open("day4")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	drawSeq, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	var draw Seq
	for _, item := range strings.Split(drawSeq[:len(drawSeq)-1], ",") {
		n, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal(err)
		}
		draw = append(draw, uint(n))
	}
	var boards []ExtraBoard
	var lastLine bool
	var i int
	boardBuf := make([]Cell, 25)
	for !lastLine {
		boardLine, err := r.ReadString('\n')
		switch err {
		case nil:
			boardLine = boardLine[:len(boardLine)-1]
		case io.EOF:
			lastLine = true
		default:
			log.Fatal(err)
		}
		if len(boardLine) == 0 {
			continue
		}
		if i == 25 {
			boards = append(boards, ExtraBoard{b: boardBuf})
			i = 0
			// Error: board defined in line 91: before the for loop
			// the values of the slice were overwritten each time by the next iteration
			// must be reinitialized with a new slice
			boardBuf = make([]Cell, 25)
		}
		for _, boardElStr := range strings.Fields(boardLine) {
			boardEl, err := strconv.Atoi(boardElStr)
			if err != nil {
				log.Fatal(err)
			}
			boardBuf[i] = Cell{n: uint(boardEl)}
			i++
		}
	}
	var lastBoard Board
	var lastDraw uint
	for i := range draw {
		if i < 4 {
			continue // min to have a full line
		}
		for j, b := range boards {
			if b.won {
				continue
			}
			win := draw[:i].Win(b.b)
			if !win {
				continue
			}
			lastBoard = b.b
			lastDraw = draw[i-1]
			// error: b.won = true is not truly set/updated (b is local variable)
			// use boards[j].won = true instead
			boards[j].won = true
		}
	}
	var sum uint
	for k := range lastBoard {
		if lastBoard[k].called {
			continue
		}
		sum += lastBoard[k].n
	}
	fmt.Println(sum, lastDraw, sum*lastDraw)
}
