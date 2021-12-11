package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main_day11() {
	isPart2 := flag.Bool("part2", false, "activate part 2")
	flag.Parse()
	day11(*isPart2)
}

func day11(isPart2 bool) {
	f, err := os.ReadFile("day11")
	if err != nil {
		log.Fatal(err)
	}
	var octs [][]uint
	var octRow []uint
	for _, b := range f {
		switch b {
		case '\n':
			octs = append(octs, octRow)
			octRow = nil
		default:
			octRow = append(octRow, uint(b-'0'))
		}
	}
	octs = append(octs, octRow)
	if isPart2 {
		part2(octs)
		return
	}
	part1(octs)

}

func part2(octs [][]uint) {
	var step uint
	var stop bool
	for !stop {
		// step 0: raise step
		step++
		// step 1: raise
		for i := range octs {
			for j := range octs[i] {
				octs[i][j]++
			}
		}
		// step 2: flash
		var flashed [][]bool
		for i := range octs {
			flashed = append(flashed, make([]bool, len(octs[i])))
		}
		for i := range octs {
			for j := range octs[i] {
				if flashed[i][j] {
					continue
				}
				if octs[i][j] <= 9 {
					continue
				}
				flash(octs, flashed, i, j)
			}
		}
		// step 3: reset flashed
		inSync := true
		for i := range octs {
			for j := range octs[i] {
				if !flashed[i][j] {
					inSync = false
					continue
				}
				octs[i][j] = 0
			}
		}
		if inSync {
			stop = true
		}
	}
	fmt.Println(step, octs)
}

func part1(octs [][]uint) {
	var count int
	fmt.Println(octs)
	for i := 0; i < 100; i++ {
		// step 1: raise
		for i := range octs {
			for j := range octs[i] {
				octs[i][j]++
			}
		}
		// step 2: flash
		var flashed [][]bool
		for i := range octs {
			flashed = append(flashed, make([]bool, len(octs[i])))
		}
		for i := range octs {
			for j := range octs[i] {
				if flashed[i][j] {
					continue
				}
				if octs[i][j] <= 9 {
					continue
				}
				flash(octs, flashed, i, j)
			}
		}
		// step 3: reset flashed
		for i := range octs {
			for j := range octs[i] {
				if !flashed[i][j] {
					continue
				}
				count++
				octs[i][j] = 0
			}
		}
		fmt.Println(count, octs)
	}
}

func flash(octs [][]uint, flashed [][]bool, i, j int) {
	flashed[i][j] = true
	flashAdj := func(octs [][]uint, flashed [][]bool, i, j int, canGoToPos func(int, int) bool) {
		if !canGoToPos(i, j) {
			return
		}
		if flashed[i][j] {
			return
		}
		octs[i][j]++
		if octs[i][j] > 9 {
			flash(octs, flashed, i, j)
		}
	}
	flashAdj(octs, flashed, i-1, j, func(i, j int) bool { return i >= 0 })
	flashAdj(octs, flashed, i-1, j-1, func(i, j int) bool { return i >= 0 && j >= 0 })
	flashAdj(octs, flashed, i, j-1, func(i, j int) bool { return j >= 0 })
	flashAdj(octs, flashed, i+1, j-1, func(i, j int) bool { return i <= 9 && j >= 0 })
	flashAdj(octs, flashed, i+1, j, func(i, j int) bool { return i <= 9 })
	flashAdj(octs, flashed, i+1, j+1, func(i, j int) bool { return i <= 9 && j <= 9 })
	flashAdj(octs, flashed, i, j+1, func(i, j int) bool { return j <= 9 })
	flashAdj(octs, flashed, i-1, j+1, func(i, j int) bool { return i >= 0 && j <= 9 })
}
