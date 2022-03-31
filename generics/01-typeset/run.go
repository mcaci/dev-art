package typeset

import "fmt"

// MyInterface is an example of a set of methods
type MyInterface interface {
	MyInterfaceF()
}

// MyTypeSet is an example of a set of types
type MyTypeSet interface {
	int | bool | string
}

// Contains uses comparable keyword to check
// if container of type slice of E contains element of type E
func Contains[S []E, E comparable](s S, e E) bool {
	for i := range s {
		if s[i] != e {
			continue
		}
		return true
	}
	return false
}

// Print all is an example of usage of any
// func PrintAll(s ...any) { // is also good
func PrintAll[S any](s ...S) {
	for i := range s {
		fmt.Println(s[i])
	}
}

func Run() {
	fmt.Println(Contains[[]string, string]([]string{"a", "b", "c"}, "c"))
	fmt.Println(Contains[[]string, string]([]string{"a", "b", "c"}, "d"))
	PrintAll("a", "b", "c", "d")
}
