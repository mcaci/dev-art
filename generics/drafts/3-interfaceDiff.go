package drafts

import (
	"fmt"
	"io"
	"os"
)

func InterfacePrint(t io.Writer, s any) {
	fmt.Fprintln(t, s)
}

func GenericPrintVec[T io.Writer](s any, t ...T) {
	fmt.Fprintln(t[0], s)
}

func InterfacePrintVec(s any, t ...io.Writer) {
	fmt.Fprintln(t[0], s)
}

func Run3() {
	GenericPrint[TypeConstr, string](os.Stdout, "with Generics: Hello World")
	InterfacePrint(os.Stdout, "with Interface: Hello World")
	GenericPrintVec[TypeConstr]("slice with Generics: Hello World", os.Stdout)
	InterfacePrintVec("slice with Interface: Hello World", os.Stdout)
}
