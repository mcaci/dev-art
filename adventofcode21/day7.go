package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"math"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main_day7() {
	day7(func(posA, posB float64) float64 { return math.Abs(posA - posB) })
	// error: did not use Abs here (negative numbers impact the calculation)
	day7(func(posA, posB float64) float64 {
		dist := math.Abs(posA - posB)
		return math.Abs(dist * (dist + 1) / 2)
	})
	plotFunc(func(posA, posB float64) float64 {
		dist := math.Abs(posA - posB)
		return math.Abs(dist * (dist + 1) / 2)
	})
}

// complexity cost for both memory and time linear in the size of the input
func day7(costF func(pos, dist float64) float64) {
	b, err := os.ReadFile("day7")
	if err != nil {
		log.Fatal(err)
	}
	var v []int
	for _, num := range bytes.Split(b, []byte{','}) {
		n, err := strconv.Atoi(string(num))
		if err != nil {
			log.Fatal(err)
		}
		v = append(v, n)
	}
	var sum float64
	const maxPos = 2000.0
	for i := 0.0; i < maxPos; i++ {
		var tmpSum float64
		for j := range v {
			// error: did not use Abs here (negative numbers impact the calculation)
			tmpSum += costF(float64(v[j]), i)
		}
		if sum == 0 {
			sum = tmpSum
			continue
		}
		if tmpSum < sum {
			sum = tmpSum
			continue
		}
		fmt.Printf("%.0f, %.0f\n", i, sum)
		break
	}
}

// complexity cost for both memory and time linear in the size of the input
func plotFunc(costF func(pos, dist float64) float64) {
	b, err := os.ReadFile("day7")
	if err != nil {
		log.Fatal(err)
	}
	var v []int
	for _, num := range bytes.Split(b, []byte{','}) {
		n, err := strconv.Atoi(string(num))
		if err != nil {
			log.Fatal(err)
		}
		v = append(v, n)
	}
	var sum float64
	const maxPos = 2000.0
	var values plotter.XYs
	for i := 0.0; i < maxPos; i++ {
		var tmpSum float64
		for j := range v {
			// error: did not use Abs here (negative numbers impact the calculation)
			tmpSum += costF(float64(v[j]), i)
		}
		if sum == 0 {
			sum = tmpSum
			continue
		}
		sum = tmpSum
		values = append(values, plotter.XY{X: i, Y: sum})
	}
	l, err := plotter.NewLine(values)
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(1), vg.Points(1)}
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	p := plot.New()
	if p == nil {
		log.Fatal("nil Plot")
	}
	p.Title.Text = "histogram plot"
	p.Add(plotter.NewGrid())
	p.Add(l)

	if err := p.Save(5*vg.Inch, 5*vg.Inch, "hist.png"); err != nil {
		log.Fatal(err)
	}
}
