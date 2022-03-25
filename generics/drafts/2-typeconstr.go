package drafts

import (
	"fmt"
	"io"
	"os"
)

// Type constraints are interfaces

type TypeConstr interface {
	Write([]byte) (int, error) // could also write "io.Writer"
}

// Struct implementing type constraint and usage

type NoWriter struct{}

func (*NoWriter) Write([]byte) (int, error) { return 0, nil }

// func GenericPrint[T1 io.Writer, T2 any](t1 T1, t2 T2) {
func GenericPrint[T1 TypeConstr, T2 any](t1 T1, t2 T2) {
	fmt.Fprintln(t1, t2)
}

// Type constraints can also list of types
// Ints includes all int* type and implements all the common operations for these types
type Ints interface {
	int | int8 | int16 | int32 | int64
}

func SumInts[I Ints](i1, i2 I) I {
	return i1 + i2
}

// This works too but is not very generic
type MyInt int

func SumMyInt[I MyInt](i ...I) I {
	var n I
	for _, x := range i {
		n += x
	}
	return n
}

// Example to show that + cannot be used here because bool doesn't support it
type IntBool interface {
	int | bool
}

func Do[I IntBool](i1, i2 I) I {
	return i1 //+ i2
}

// constraints.Ordered includes all integers. floats and string types

func Run2() {
	// calling generic function with different types inside the type constraints
	GenericPrint[TypeConstr, string](&NoWriter{}, "Hello World")
	GenericPrint[TypeConstr, string](os.Stdout, "TypeConstr: Hello World")
	GenericPrint[io.Writer, string](os.Stdout, "io.Writer: Hello World")
	GenericPrint[*os.File, string](os.Stdout, "*os.File: Hello World")
}
