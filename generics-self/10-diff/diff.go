package diff

import (
	"fmt"
	"io"
)

func genericPrint[T io.Writer](t T) {
	fmt.Fprintln(t, "ciao")
}

func interfacePrint(t io.Writer) {
	fmt.Fprintln(t, "ciao")
}

func genericPrintVec[T io.Writer](t ...T) {
	fmt.Fprintln(t[0], "ciao")
}

func interfacePrintVec(t ...io.Writer) {
	fmt.Fprintln(t[0], "ciao")
}

// any can be used as regular parameter (as interface{})
func Ex3(a any) bool { return a == nil }

