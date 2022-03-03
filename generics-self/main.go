package main

import (
	"fmt"
	"io"
	"os"

	"mcaci/dev-art/generics-self/gen"
)

type NoWriter struct{}

func (*NoWriter) Write([]byte) (int, error) { return 0, nil }

func main() {
	// Direct call
	fmt.Println(gen.Ex1[int](1, 2))

	// Call via generic function instantiation
	f := gen.Ex1[int]
	fmt.Println(f(1, 2))

	// Direct call with 2 type parameters
	fmt.Println(gen.Ex2[string, float64]("2.0", 3.14))
	// ERR invalid type
	// fmt.Println(gen.Ex2[int, float64]("2.0", 3.14))

	// Usage of "type parameter" type inside a func call
	// fmt.Println(gen.Ex3((*os.File)(nil)))

	// calling generic function with different types inside the type constraints
	gen.GenericPrint[gen.TypeConstr, string](&gen.NoWriter{}, "Hello World")
	gen.GenericPrint[gen.TypeConstr, string](os.Stdout, "gen.TypeConstr: Hello World")
	gen.GenericPrint[io.Writer, string](os.Stdout, "io.Writer: Hello World")
	gen.GenericPrint[*os.File, string](os.Stdout, "*os.File: Hello World")

}
