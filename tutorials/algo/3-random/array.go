package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	randArray(v)
	fmt.Println(v)
	fmt.Println(randPartArray(v, 3))
	m := [][]int{v, v}
	fmt.Println(rand2DArray(m))
}

// randArray func
// For i from 0 to N=len(v)-1
// (1) pick an item for position
// j = random number between i and N
// (2) swap items at i and j
// O(n) where n is the lenght of the array
func randArray(v []int) {
	rand.Seed(time.Now().UnixNano())
	for i := range v {
		// i + (start from i)
		// len(v)-i are the items that are left from i to len(v)
		// Intn selects an interval [0,n)
		j := i + rand.Intn(len(v)-i)
		v[i], v[j] = v[j], v[i]
	}
}

// randPartArray func is used to pick the first n random elements
// For i from 0 to N=n
// (1) pick an item for position
// j = random number between i and N
// (2) swap items at i and j
// O(n) where n is the lenght of the subarray (param)
func randPartArray(v []int, n int) []int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		// i + (start from i)
		// len(v)-i are the items that are left from i to len(v)
		// Intn selects an interval [0,n)
		j := i + rand.Intn(len(v)-i)
		v[i], v[j] = v[j], v[i]
	}
	return v[:n]
}

// rand2DArray randomizes a 2D array by copying it to 1D, randomizing and copying back
// O(n) where n is the number of the 2D array
func rand2DArray(v [][]int) [][]int {
	rand.Seed(time.Now().UnixNano())
	var copy []int
	for i := range v {
		for j := range v[i] {
			copy = append(copy, v[i][j])
		}
	}
	randArray(copy)
	var m [][]int
	for i := range v {
		m = append(m, []int{})
		for j := range v[i] {
			// strange behavior v[i][j] are equal for all [i] because slices are used by reference
			// v[i][j] = copy[i*len(v)+j]
			m[i] = append(m[i], copy[i*len(v)+j])
		}
	}
	return m
}
