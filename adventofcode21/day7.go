package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	day7()
}

func day7() {
	b, err := os.ReadFile("day7")
	if err != nil {
		log.Fatal(err)
	}
	f := make(map[int]uint)
	v := make([]int, 2000)
	for _, num := range bytes.Split(b, []byte{','}) {
		n, err := strconv.Atoi(string(num))
		if err != nil {
			log.Fatal(err)
		}
		f[n]++
		v[n]++
	}
	fmt.Println(f)
	fmt.Println(v)
}
