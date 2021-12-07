package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main_day6() {
	// const nDays = 80
	day6algo1(80)
	const nDays = 256
	day6algo2(nDays)
	day6algo2(nDays)
}

// complexity of algo1 is exponential of O(2^n) in space and with respect to the input size
// because in the second for loop each element may be doubled in the slice which in turn
// takes a double the time to read fully
func day6algo1(nDays int) {
	b, err := os.ReadFile("day6")
	if err != nil {
		log.Fatal(err)
	}
	var fishes []uint
	{
		for _, b := range bytes.Split(b, []byte{','}) {
			f, err := strconv.Atoi(string(b))
			if err != nil {
				log.Fatal(err)
			}
			fishes = append(fishes, uint(f))
		}
	}
	for i := 0; i < nDays; i++ {
		var newFishes int
		for fID := range fishes {
			switch fishes[fID] {
			case 0:
				fishes[fID] = 6
				newFishes++
			default:
				fishes[fID]--
			}
		}
		for j := 0; j < newFishes; j++ {
			fishes = append(fishes, 8)
		}
	}
	fmt.Println(len(fishes))
}

// complexity of algo2 is O(n + m) in time where n is the initial number of fishes and m is the number of days
// and in space is constant as we need only two maps (or slices) with size the max number of days needed for a new spawning (which is 8)
func day6algo2(nDays int) {
	b, err := os.ReadFile("day6")
	if err != nil {
		log.Fatal(err)
	}
	fishMapDayNMinus1 := make(map[uint8]uint)
	{
		for _, b := range bytes.Split(b, []byte{','}) {
			f, err := strconv.Atoi(string(b))
			if err != nil {
				log.Fatal(err)
			}
			fishMapDayNMinus1[uint8(f)]++
		}
	}
	for i := 0; i < nDays; i++ {
		fishMapDayN := make(map[uint8]uint)
		for k, v := range fishMapDayNMinus1 {
			if k == 0 {
				fishMapDayN[6] += v
				fishMapDayN[8] += v
				continue
			}
			fishMapDayN[k-1] += v
		}
		fishMapDayNMinus1 = fishMapDayN
	}
	var sum uint
	for _, v := range fishMapDayNMinus1 {
		sum += v
	}
	fmt.Println(sum)
}
