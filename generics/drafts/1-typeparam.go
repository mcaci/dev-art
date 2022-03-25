package drafts

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func IntMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Type parameter list
// [P, Q constraint1, R constraint2, ...]
// Capital letter is a convention
// constraint1, constraint2 are interface types

// constraints.Ordered includes numeric types and string
func Ex1[T constraints.Ordered](a, b T) T {
	if a < b {
		return b
	}
	return a
}

// any is equivalent to an empty interface
func Ex2[P any, Q constraints.Ordered](a P, b Q) string {
	return fmt.Sprint(a, " ", b)
}

// type parameters can be applied to functions and methods, structures and interfaces
type Tree[T any] struct {
	p, l, r *Tree[T]
	e       T
}

func (t *Tree[T]) Next() *Tree[T] { return nil }

var intTree Tree[int]

type Minr[T any] interface {
	Min(a, b T) T
}

var minFunc8 Minr[int8]

func Run1() {
	// Direct call
	fmt.Println(Ex1[int](1, 2))

	// Call via generic function instantiation
	f := Ex1[int]
	fmt.Println(f(1, 2))

	// Direct call with 2 type parameters
	fmt.Println(Ex2[string, float64]("2.0", 3.14))
}
