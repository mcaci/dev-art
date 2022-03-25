package typeset

import "fmt"

// MethodTypeSet is an example of a set of types and methods
type MethodTypeSet interface {
	~int | int32
	String() string
}

func Print[T MethodTypeSet](t T) { fmt.Println(t) }

func RunMethodTypeSet() {
	// Print[int](2)
	// Print[int32](2)
	// Print[A](A{2})
	// Print[B](2)
	// Print[C](2)
	// Print[D](2)
	// Print[E](E{2})
}

type A struct{ int }
type B int
type C struct{}
type D int32
type E struct{ int32 }

func (A) String() string { return "A" }
func (B) String() string { return "B" }
func (C) String() string { return "C" }
func (D) String() string { return "D" }
func (E) String() string { return "E" }
