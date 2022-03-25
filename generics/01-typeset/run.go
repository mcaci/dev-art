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

func Contains[S ~[]E, E comparable](s S, e E) bool {
	for i := range s {
		if s[i] != e {
			continue
		}
		return true
	}
	return false
}

func PrintAnyValue[S any](s ...S) {
	for i := range s {
		fmt.Println(s[i])
	}
}

func Run() {
	// Print[A](2)
	// Print[B](2)
	// Print[C](2)
	// Print[D](2)
	// Print[int](2)
	// Print[int32](2)
}
