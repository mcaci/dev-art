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
	day4Part1()
}

type Seq []uint
type Board []uint

type BoardInfo struct {
	hitIdx      []bool
	found       bool
	lastNCalled uint
}

func (s Seq) Win(b Board) BoardInfo {
	hitIdx := make([]bool, 25)
	for i := range b {
		for _, n := range s {
			if n != b[i] {
				continue
			}
			hitIdx[i] = true
		}
	}
	for i := 0; i < len(hitIdx); i += 5 {
		found := true
		for _, hit := range hitIdx[i : i+5] {
			found = found && hit
		}
		if !found {
			continue
		}
		return BoardInfo{hitIdx: hitIdx, found: true, lastNCalled: s[len(s)-1]}
	}
	for i := 0; i < len(hitIdx)/5; i++ {
		found := true
		localHits := []bool{hitIdx[i], hitIdx[i+5], hitIdx[i+10], hitIdx[i+15], hitIdx[i+20]}
		for _, hit := range localHits {
			found = found && hit
		}
		if !found {
			continue
		}
		return BoardInfo{hitIdx: hitIdx, found: true, lastNCalled: s[len(s)-1]}
	}
	return BoardInfo{}
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
	boardBuf := make([]uint, 25)
	var i int
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
		}
		for _, boardElStr := range strings.Fields(boardLine) {
			boardEl, err := strconv.Atoi(boardElStr)
			if err != nil {
				log.Fatal(err)
			}
			boardBuf[i] = uint(boardEl)
			i++
		}
	}
	// fmt.Println(boards)
	for i := range draw {
		if i < 5 {
			continue // min to have a full line
		}
		for j, b := range boards {
			info := draw[:i].Win(b)
			if !info.found {
				continue
			}
			fmt.Println(i, j, len(boards), info.lastNCalled)
			fmt.Println(info.hitIdx)
			fmt.Println(b)
			var sum uint
			for k := range b {
				if info.hitIdx[k] {
					continue
				}
				sum += b[k]
			}
			fmt.Println(sum * info.lastNCalled)
			return
		}
	}
}
