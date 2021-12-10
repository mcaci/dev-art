package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"os"
	"sort"
)

func main_day9() {
	day9Part2()
	draw()
}

func draw() {
	height := flag.Int("h", 400, "height of the output image in pixels")
	width := flag.Int("w", 400, "width of the output image in pixels")
	flag.Parse()

	const out = "map.png"
	f, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}
	c := make([][]color.RGBA, *height)
	for i := range c {
		c[i] = make([]color.RGBA, *width)
	}
	heightmap := heightmapF()
	for i, row := range c {
		for j := range row {
			n := uint8(255.0 * (1.0 - float64(heightmap[i/4][j/4])/9.0))
			c[i][j] = color.RGBA{R: uint8(float64(n) * 0.6), G: uint8(float64(n) * 0.5), B: uint8(float64(n) * 0.4), A: 255}
		}
	}
	img := &img{
		h: *height,
		w: *width,
		m: c,
	}

	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}

type img struct {
	h, w int
	m    [][]color.RGBA
}

func (m *img) At(x, y int) color.Color { return m.m[x][y] }
func (m *img) ColorModel() color.Model { return color.RGBAModel }
func (m *img) Bounds() image.Rectangle { return image.Rect(0, 0, m.h, m.w) }

type HMCoord struct{ i, j int }

func day9Part2() {
	heightmap := heightmapF()
	var lowpoints []HMCoord
	for i := range heightmap {
		for j := range heightmap[i] {
			lessThanTopVal := i == 0 || heightmap[i][j] < heightmap[i-1][j]
			lessThanBotmVal := i == len(heightmap)-1 || heightmap[i][j] < heightmap[i+1][j]
			lessThanLeftVal := j == 0 || heightmap[i][j] < heightmap[i][j-1]
			lessThanRightVal := j == len(heightmap)-1 || heightmap[i][j] < heightmap[i][j+1]
			if !(lessThanTopVal && lessThanBotmVal && lessThanLeftVal && lessThanRightVal) {
				continue
			}
			lowpoints = append(lowpoints, HMCoord{i, j})
		}
	}
	var areas []int
	visited := make([][]bool, len(heightmap))
	for i := range heightmap {
		visited[i] = make([]bool, len(heightmap[i]))
	}
	for i := range lowpoints {
		area := explore(heightmap, visited, lowpoints[i].i, lowpoints[i].j)
		areas = append(areas, area)
	}
	sort.Ints(areas)
	mult := 1
	for _, area := range areas[len(areas)-3:] {
		mult *= area
	}
	fmt.Println(mult)
}

func explore(heightmap [][]uint8, visited [][]bool, i, j int) int {
	// fmt.Println(i, j)
	if visited[i][j] {
		return 0
	}
	if heightmap[i][j] == 9 {
		visited[i][j] = true
		return 0
	}
	area := 1
	visited[i][j] = true
	if i != 0 {
		area += explore(heightmap, visited, i-1, j)
	}
	if j != 0 {
		area += explore(heightmap, visited, i, j-1)
	}
	// error: forgot -1
	if j != len(heightmap[i])-1 {
		area += explore(heightmap, visited, i, j+1)
	}
	if i != len(heightmap)-1 {
		area += explore(heightmap, visited, i+1, j)
	}
	return area
}

func day9Part1() {
	heightmap := heightmapF()
	var lowpoints []uint8
	for i := range heightmap {
		for j := range heightmap[i] {
			lessThanTopVal := i == 0 || heightmap[i][j] < heightmap[i-1][j]
			lessThanBotmVal := i == len(heightmap)-1 || heightmap[i][j] < heightmap[i+1][j]
			lessThanLeftVal := j == 0 || heightmap[i][j] < heightmap[i][j-1]
			lessThanRightVal := j == len(heightmap)-1 || heightmap[i][j] < heightmap[i][j+1]
			if !(lessThanTopVal && lessThanBotmVal && lessThanLeftVal && lessThanRightVal) {
				continue
			}
			lowpoints = append(lowpoints, heightmap[i][j])
		}
	}
	var sum int
	for _, lowP := range lowpoints {
		sum += int(lowP) + 1
	}
	fmt.Println(sum)
}

func heightmapF() [][]uint8 {
	f, err := os.Open("day9")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	var lastline bool
	var heightmap [][]uint8
	var heightmaprow []uint8
	for !lastline {
		b, err := r.ReadByte()
		switch err {
		case nil:
		case io.EOF:
			heightmap = append(heightmap, heightmaprow)
			heightmaprow = nil
			lastline = true
			continue
		default:
			log.Fatal(err)
		}
		switch b {
		case '\n':
			heightmap = append(heightmap, heightmaprow)
			heightmaprow = nil
		default:
			heightmaprow = append(heightmaprow, b-'0')
		}
	}
	return heightmap
}
