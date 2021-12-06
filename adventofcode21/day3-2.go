package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func day3Part2() {
	day3_2_algo2()
}

// To be implemented: test with sorting
func day3_2_algo4() {}

// too expensive with respect to algo2
// memory is not not gained and reading the file multiple times increase the cpu time of an order of magnitude
func day3_2_algo3() {
	oxy := oxyInfo.build()
	oNum, err := strconv.ParseUint(string(oxy), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	co2 := co2Info.build()
	cNum, err := strconv.ParseUint(string(co2), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(oNum, cNum, oNum*cNum)
}

type builderInfo struct{ moreOnesByte, moreZeroesByte byte }

var oxyInfo = builderInfo{moreOnesByte: '1', moreZeroesByte: '0'}
var co2Info = builderInfo{moreOnesByte: '0', moreZeroesByte: '1'}

func (bInf builderInfo) build() (out []byte) {
	for i := 0; i < 12; i++ {
		f, err := os.Open("day3")
		if err != nil {
			log.Fatal(err)
		}
		r := bufio.NewReader(f)
		var lastLine bool
		var zeroes, ones int
		for !lastLine {
			line, err := r.ReadBytes('\n')
			switch err {
			case nil:
				line = line[:len(line)-1]
			case io.EOF:
				lastLine = true
			default:
				log.Fatal(err)
			}
			switch string(line[:i+1]) {
			case string(out[:i]) + "0":
				zeroes++
			case string(out[:i]) + "1":
				ones++
			default:
				// skipping
			}
		}
		switch {
		case zeroes == 0:
			out = append(out, '1')
		case ones == 0:
			out = append(out, '0')
		case ones >= zeroes:
			out = append(out, bInf.moreOnesByte)
		case zeroes > ones:
			out = append(out, bInf.moreZeroesByte)
		default:
			// skipping
		}
	}
	return out
}

// concurrent version of algo1: no special gains but a bit more performant on the benchmark
func day3_2_algo2() {
	f, err := ioutil.ReadFile("day3")
	if err != nil {
		log.Fatal(err)
	}
	in := bytes.Split(f, []byte{'\n'})

	oxyChan := make(chan int)
	go func() {
		o := string(find(in, oxy))
		oNum, err := strconv.ParseUint(o, 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		oxyChan <- int(oNum)
	}()
	co2Chan := make(chan int)
	go func() {
		c := string(find(in, co2))
		cNum, err := strconv.ParseUint(c, 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		co2Chan <- int(cNum)
	}()
	oNum, cNum := <-oxyChan, <-co2Chan
	fmt.Println(oNum, cNum, oNum*cNum)
}

// Complexity is O(n * m) where n is the length of the file and m the length of the line
// One read of the file to store all in memory (210kb for this file), this could be a problem for bigger sizes potentially
func day3_2_algo1() {
	f, err := ioutil.ReadFile("day3")
	if err != nil {
		log.Fatal(err)
	}
	in := bytes.Split(f, []byte{'\n'})

	o := string(find(in, oxy))
	oNum, err := strconv.ParseUint(o, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	c := string(find(in, co2))
	cNum, err := strconv.ParseUint(c, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(oNum, cNum, oNum*cNum)
}

func find(in [][]byte, f interface{ filter([][]byte, int) [][]byte }) (out []byte) {
	data := make([][]byte, len(in))
	copy(data, in)
	for i := 0; i < len(data[0]); i++ {
		data = f.filter(data, i)
		if len(data) == 1 {
			return data[0]
		}
	}
	return nil
}

type filterByInfo struct{ moreOnesByte, moreZeroesByte byte }

var oxy = filterByInfo{moreOnesByte: '1', moreZeroesByte: '0'}
var co2 = filterByInfo{moreOnesByte: '0', moreZeroesByte: '1'}

type filterFuncAtPos func([][]byte, int) [][]byte

func (fbinfo filterByInfo) filter(data [][]byte, pos int) [][]byte {
	zeroes, ones := counts(data, pos)
	var f filterFuncAtPos
	switch {
	case ones >= zeroes:
		f = using(fbinfo.moreOnesByte)
	case zeroes > ones:
		f = using(fbinfo.moreZeroesByte)
	default:
		log.Fatal("Unexpected situation with zero and one counts")
	}
	return f(data, pos)
}

// Error I did before: (count0, count1 byte)
// Always use int, not byte as return type (too short for order of magnitude 1000)
func counts(in [][]byte, j int) (count0, count1 int) {
	for i := range in {
		switch in[i][j] {
		case '0':
			count0++
		case '1':
			count1++
		default:
			log.Fatal(in[i], ":", in[i][j], "was not expected")
		}
	}
	return count0, count1
}

func using(n byte) filterFuncAtPos {
	return func(data [][]byte, j int) (out [][]byte) {
		for i := range data {
			if data[i][j] != n {
				continue
			}
			out = append(out, data[i])
		}
		return out
	}
}
