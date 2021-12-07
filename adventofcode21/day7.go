package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
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
	f := make([]int, len(b))
	for i, num := range bytes.Split(b, []byte{','}) {
		n, err := strconv.Atoi(string(num))
		if err != nil {
			log.Fatal(err)
		}
		f[i] = n
	}
	sort.Ints(f)
	fmt.Println(f)
}
