package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

func main_day12() {
	day12()
}

func day12() {
	f, err := os.Open("day12")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	caves := cavernGraph(r)
	var start *Cave
	for _, c := range caves {
		if c.name != "start" {
			continue
		}
		start = c
		break
	}
	cPart1 := nPathsDFS(start, nil)
	for _, c := range caves {
		c.nVisits = 0
	}
	cPart2 := nPathsDFS(start, caves)
	fmt.Println(cPart1, cPart2)
}

type OneVisitRestriction struct{}

func (OneVisitRestriction) deny() bool { return true }

type Caves []*Cave

func (caves Caves) deny() bool {
	for _, c := range caves {
		if !(c.isSmall() && c.nVisits > 1) {
			continue
		}
		return true
	}
	return false
}

type Cave struct {
	name    string
	nVisits int
	connTo  []*Cave
}

func (c Cave) visited() bool { return c.nVisits > 0 }
func (c Cave) isSmall() bool { return strings.IndexFunc(c.name, unicode.IsUpper) == -1 }
func (c Cave) deny() bool    { return c.isSmall() && c.visited() }

func (c Cave) String() string {
	var connNames []string
	for _, conn := range c.connTo {
		connNames = append(connNames, conn.name)
	}
	return fmt.Sprintf("[Name: %s, visited: %t, isSmall: %t, connected to (%s)]", c.name, c.visited(), c.isSmall(), strings.Join(connNames, ","))
}

func nPathsDFS(c *Cave, restricter interface{ deny() bool }) (count int) {
	if restricter == nil {
		restricter = OneVisitRestriction{}
	}
	c.nVisits++
	defer func() { c.nVisits-- }()
	for i := range c.connTo {
		nextCave := c.connTo[i]
		switch {
		case nextCave.name == "end":
			count++
		// errors: did not include start and nextCave.isSmall (missing many skip cases this way)
		case nextCave.name == "start",
			nextCave.isSmall() && nextCave.visited() && restricter.deny():
			continue
		default:
			count += nPathsDFS(nextCave, restricter)
		}
	}
	return count
}

func countVisitDFS(c *Cave) (count int) {
	if c.name == "end" {
		count++
		return count
	}
	if c.visited() {
		return count
	}
	c.nVisits++
	defer func() { c.nVisits-- }()
	for i := range c.connTo {
		count += countVisitDFS(c.connTo[i])
	}
	return count
}

func cavernGraph(r *bufio.Reader) Caves {
	// error: didnt use []*Cave but just []Cave... all modifications to Cave (adj list) were not saved
	var caves []*Cave
	var lastLine bool
	for !lastLine {
		line, err := r.ReadString('\n')
		switch err {
		case nil:
			line = line[:len(line)-1]
		case io.EOF:
			lastLine = true
		default:
			log.Fatal(err)
		}
		i := strings.Index(line, "-")
		nn1, nn2 := line[:i], line[i+1:]
		var cave1, cave2 *Cave
		for _, c := range caves {
			switch c.name {
			case nn1:
				cave1 = c
			case nn2:
				cave2 = c
			default:
				continue
			}
		}
		if cave1 == nil {
			cave1 = &Cave{name: nn1}
			caves = append(caves, cave1)
		}
		if cave2 == nil {
			cave2 = &Cave{name: nn2}
			caves = append(caves, cave2)
		}
		cave1.connTo = append(cave1.connTo, cave2)
		cave2.connTo = append(cave2.connTo, cave1)
	}
	return caves
}

func dfsVisit(c *Cave) []*Cave {
	if c.visited() {
		return nil
	}
	c.nVisits++
	defer func() { c.nVisits-- }()
	if c.name == "end" {
		return []*Cave{c}
	}
	var path []*Cave
	path = append(path, c)
	for i := range c.connTo {
		pathVisit := dfsVisit(c.connTo[i])
		if len(pathVisit) == 0 {
			continue
		}
		if pathVisit[len(pathVisit)-1].name != "end" {
			continue
		}
		return append(path, pathVisit...)
	}
	path = path[:len(path)-1]
	return path
}
