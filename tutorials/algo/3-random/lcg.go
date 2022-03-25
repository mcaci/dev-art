package main

import "fmt"

// main for linear congruential generator
func main() {
	r := rand{A: 3, B: 5, M: 7}
	for i := 0; i < 10; i++ {
		fmt.Println(r.lcg())
	}
}

// rand struct implements a linear congruential generator for random numbers
// Conditions to apply for it to work (otherwise it gets stuck in a loop)
// B and M must be relatively prime
// A-1 is divisible by all prime factors of M
// A-1 and M are both or none of them multiple of 4 (if A-1 is M is and viceversa)
type rand struct {
	x, A, B, M int
}

// To compute the n-th random number n steps need to be done so the CC is O(n)
// Go Note: forgot the * in *rand which made rand be passed by value and not by reference (at the end it means it was not updated)
func (r *rand) lcg() int {
	// x(i) := (x(i-1) * A + B) mod M
	r.x = (r.x*r.A + r.B) % r.M
	return r.x
}
