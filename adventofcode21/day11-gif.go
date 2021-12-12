package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"log"
	"os"
)

func main_day11draw() {
	height := flag.Int("h", 160, "height of the output image in pixels")
	width := flag.Int("w", 160, "width of the output image in pixels")
	flag.Parse()

	f, err := os.ReadFile("day11")
	var images []*image.Paletted
	var delays []int
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
	for i := 0; i < 500; i++ {
		log.Println("Step count:", i)
		// step 1: raise
		for i := range octs {
			for j := range octs[i] {
				octs[i][j]++
			}
		}
		// create and append image
		images = append(images, octopusImg(octs, *height, *width))
		delays = append(delays, 15)
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
				flashGif(octs, flashed, i, j, images, delays, *height, *width)
			}
		}
		// step 3: reset flashed
		for i := range octs {
			for j := range octs[i] {
				if !flashed[i][j] {
					continue
				}
				octs[i][j] = 0
			}
		}
	}
	const out = "octs.gif"
	outGif, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}
	defer outGif.Close()
	fmt.Println("Encoding started", len(images))
	err = gif.EncodeAll(outGif, &gif.GIF{
		Image: images,
		Delay: delays,
	})
	if err != nil {
		log.Fatal(err)
	}
}

var imgCount int

func octopusImg(octs [][]uint, h, w int) *image.Paletted {
	log.Println("Image count: ", imgCount)
	imgCount++
	img := image.NewPaletted(image.Rect(0, 0, h, w), palette.Plan9)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			n := uint8(255.0 * (1.0 - float64(octs[i/16][j/16])/9.0))
			img.Set(i, j, color.RGBA{R: uint8(float64(n) * 0.5), G: uint8(float64(n) * 0.3), B: uint8(float64(n) * 0.8), A: 255})
		}
	}
	return img
}

func flashGif(octs [][]uint, flashed [][]bool, i, j int, images []*image.Paletted, delays []int, h, w int) {
	flashed[i][j] = true
	flashAdj := func(octs [][]uint, flashed [][]bool, i, j int, canGoToPos func(int, int) bool) {
		if !canGoToPos(i, j) {
			return
		}
		if flashed[i][j] {
			return
		}
		// create and append image
		images = append(images, octopusImg(octs, h, w))
		delays = append(delays, 15)
		octs[i][j]++
		if octs[i][j] > 9 {
			flashGif(octs, flashed, i, j, images, delays, h, w)
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
