# Reading stack traces in Go

In this guide I will show some useful information and tips on how to read a stack trace. Stack traces in Go are very common when a panic is triggered, they are not easy to read but the information contained in them may be useful to get some context and understand better what was happening in the code when the panic started.

## A generic stack trace

We start by generating a panic with the following [playground example](https://play.golang.org/p/upJekR68oyw) where we call recursively `i` times a function call before starting a panic and exiting.

```go
func main() {
    iPanic(5)
}

func iPanic(i int) {
    if i > 0 {
        iPanic(i - 1)
    }
    panic("I'm outta here")
}
```

The output, as shown below, is composed of the message passed to the panic built-in function, a line describing which goroutine was running at the moment of the panic and the stack trace containing all the calls from the beginning of the program execution to the line that panicked.

```txt
panic: I'm outta here

goroutine 1 [running]:
main.iPanic(0x0)
    /tmp/sandbox634668020/prog.go:11 +0x4f
main.iPanic(0x1)
    /tmp/sandbox634668020/prog.go:9 +0x33
main.iPanic(0x2)
    /tmp/sandbox634668020/prog.go:9 +0x33
main.iPanic(0x3)
    /tmp/sandbox634668020/prog.go:9 +0x33
main.iPanic(0x4)
    /tmp/sandbox634668020/prog.go:9 +0x33
main.iPanic(0x5)
    /tmp/sandbox634668020/prog.go:9 +0x33
main.main()
    /tmp/sandbox634668020/prog.go:4 +0x2a
```

`main.main() <br/> /tmp/sandbox634668020/prog.go:4 +0x2a` is an example of line that composes the stack trace. Each line contains:

- __the package name and the name and arguments of the caller function__: `main.main()`
- __the file where the function was called__ : `/tmp/sandbox634668020/prog.go`
- __the line in the file where the function is located__ : `4`
- __the relative position of the function in the stack frame__ : `+0x2a`

This picture/table shows an example of stack frame to illustrate what _the relative position of the function in the stack frame_ means:

| Function call       | Relative position |
|---------------------|:-----------------:|
| `main.iPanic(0x0)`  | +0x4f             |
| `main.iPanic(...)`  | +0x33             |
| `main.main()`       | +0x2a             |
| Bottom of the stack |  0x0              |

------------------------------------------

The code also features a stack traces with a call to a function with an `int` parameter that is represented for example in the line:

```txt
main.iPanic(0x3)
   /tmp/sandbox634668020/prog.go:9 +0x33
```

In this line we see the value `0x3` which is the hexadecimal encoding of the `int` value assigned to the parameter `i` in the call of `main.iPanic(0x3)`.

### Note on functions with unused or empty parameter list

Take a look at this example below and its correspondent output

```go
func main() {
    iPanic(5)
}

func iPanic(i int) {
    panic("I'm outta here")
}
```

```txt
panic: I'm outta here

goroutine 1 [running]:
main.iPanic(...)
    /tmp/sandbox137813596/prog.go:8
main.main()
    /tmp/sandbox137813596/prog.go:4 +0x39
```

Even though the `iPanic` function is called with a valid parameter the stack traces shows a __`(...)`__, in the line `main.iPanic(...)` instead.

This __empty parameter list symbol__ is returned either when the __parameters are not used__ in a meaningful computation, e.g. `fmt.Print(i)` would not change the stack trace content, or when the function has __no parameters__ at all.

## Bool and numeric types

Let's now take this [playground example](https://play.golang.org/p/-C0ov_WLkVN) with several parameters of type `bool` and numeric.

```go
import "fmt"

func main() {
    iPanic(true, 6, 'c', 6.7, complex(2, 1))
}

func iPanic(b bool, by byte, r rune, f float32, c complex64) {
    fmt.Println(by + 1)
    panic("I'm outta here")
}
```

Here's the output (excluding the `Println` line)

```txt
panic: I'm outta here

goroutine 1 [running]:
main.iPanic(0x6300000601, 0x4000000040d66666, 0xc03f800000)
    /tmp/sandbox564764099/prog.go:11 +0xad
main.main()
    /tmp/sandbox564764099/prog.go:6 +0x5a
```

This output is a bit messy to read as it appears to show 3 out 5 parameters but this is not the case. In fact this example is a way to show that parameters may be encoded to fit together into several _words_, sequences of 32 or 64 bits.

If `iPanic` is rerun with one parameter at a time we can see each value separately as follows:

- __bool__: is located in main.iPanic(0xc00005e0**01**)
  - or in main.iPanic(0x63000006**01**, 0x4000000040d66666, 0xc03f800000) in the original call
- __byte__: is located in main.iPanic(0xc00005e0**06**)
  - or in main.iPanic(0x630000**06**01, 0x4000000040d66666, 0xc03f800000) in the original call
- __rune__: is located in main.iPanic(0xc0000000**63**)
  - or in main.iPanic(0x**63**00000601, 0x4000000040d66666, 0xc03f800000) in the original call
- __float32__: is located in main.iPanic(0xc0**40d66666**)
  - or in main.iPanic(0x6300000601, 0x40000000**40d66666**, 0xc03f800000) in the original call
- __complex64__: is located in main.iPanic(0x**3f80000040000000**)
  - or in main.iPanic(0x6300000601, 0x**40000000**40d66666, 0xc0**3f800000**) in the original call

All of this values are hexadecimal encoding of the value that has been passed to the parameters in the example. For instance `0x01` stands for `true` for the `bool` parameter (as well as `0x00` stands for `false`), `0x06` is the encoding of the byte parameter with decimal value 6 and so on.

The encoding of the values is done in this way for `bool` and all numeric types, in which `rune` is included.

## Pointers and nil

Regarding pointers and nil values, we generate a panic with the following [playground example](https://play.golang.org/p/5Ica8GqHDNn) where we call recursively `i` and times a function call before starting a panic and exiting. In addition of this we add a `*int` parameter to read the value printed in the stack trace.

```go
func main() {
    a := 5
    iPanic(a, &a)
}

func iPanic(i int, j *int) {
    if i > 0 {
        iPanic(i-1, j)
    }
    panic("I'm outta here")
}
```

The output, as shown below, is composed of the message passed to the panic built-in function, a line describing which goroutine was running at the moment of the panic and the stack trace containing all the calls from the beginning of the program execution to the line that panicked.

```txt
panic: I'm outta here

goroutine 1 [running]:
main.iPanic(0x0, 0xc000032770)
    /tmp/sandbox341218196/prog.go:12 +0x59
main.iPanic(0x1, 0xc000032770)
    /tmp/sandbox341218196/prog.go:10 +0x3d
main.iPanic(0x2, 0xc000032770)
    /tmp/sandbox341218196/prog.go:10 +0x3d
main.iPanic(0x3, 0xc000032770)
    /tmp/sandbox341218196/prog.go:10 +0x3d
main.iPanic(0x4, 0xc000032770)
    /tmp/sandbox341218196/prog.go:10 +0x3d
main.iPanic(0x5, 0xc000032770)
    /tmp/sandbox341218196/prog.go:10 +0x3d
main.main()
    /tmp/sandbox341218196/prog.go:5 +0x3d
```

If we replace `&a` with `nil` we will see in the stack trace that the address changes to `0x0` as shown in the output of the appropriate run.

```txt
panic: I'm outta here

goroutine 1 [running]:
main.iPanic(0x0, 0x0)
    /tmp/sandbox325035618/prog.go:12 +0x59
main.iPanic(0x1, 0x0)
    /tmp/sandbox325035618/prog.go:10 +0x3d
main.iPanic(0x2, 0x0)
    /tmp/sandbox325035618/prog.go:10 +0x3d
main.iPanic(0x3, 0x0)
    /tmp/sandbox325035618/prog.go:10 +0x3d
main.iPanic(0x4, 0x0)
    /tmp/sandbox325035618/prog.go:10 +0x3d
main.iPanic(0x5, 0x0)
    /tmp/sandbox325035618/prog.go:10 +0x3d
main.main()
    /tmp/sandbox325035618/prog.go:5 +0x33
```

## Strings and slices

Let's now take a look at this [playground example](https://play.golang.org/p/258SGLRIgiF) featuring `strings` and `slices`  and its correspondent output.

```go
import "fmt"

func main() {
    iPanic("hello", make([]string, 3, 5))
}

func iPanic(s string, v []string) {
    fmt.Println(s + "world")
    panic("I'm outta here")
}
```

```txt
panic: I'm outta here

goroutine 1 [running]:
main.iPanic(0x4bcedc, 0x5, 0xc000068f28, 0x3, 0x5)
    /tmp/sandbox564764099/prog.go:11 +0xad
main.main()
    /tmp/sandbox564764099/prog.go:6 +0x5a
```

Now in the `iPanic` line we have more information than number of parameters. This is related to how strings and slices are represented in memory.

For the __string__ `hello` parameter the stack trace line shows its address `0x4bcedc` and its size `0x5`

- main.iPanic(**0x4bcedc, 0x5**, 0xc000068f28, 0x3, 0x5) -> this is the `hello` **string**

For the __[]string__ parameter the stack trace line shows again its address `0xc000068f28`, its size `0x3` and its capacity `0x5`

- main.iPanic(0x4bcedc, 0x5, **0xc000068f28, 0x3, 0x5**) -> this is the `make([]string, 3, 5)` **slice**

Note that strings and slices are _referenced by their address_ in the stack trace and not to the values they hold.

## Structs

Structs are collections of fields and/or embeded structs and interfaces. As we can see in this [playground example](https://play.golang.org/p/jpwBi1dLL2U) and its stack trace the struct fields are put in the paramters in the order in which they are defined in the struct when referenced by value.

```go
import "fmt"

type A struct {
    i int
    s string
}

func main() {
    iPanic(A{i:50})
}

func iPanic(a A) {
    fmt.Println(a.i + 1)
    panic("I'm outta here")
}
```

```txt
panic: I'm outta here

goroutine 1 [running]:
main.iPanic(0x32, 0x0, 0x0)
    /tmp/sandbox211242967/prog.go:16 +0xa5
main.main()
    /tmp/sandbox211242967/prog.go:11 +0x32
```

Changing the value parameter to a reference using `a *A` will change the stack trace which will show the reference to the parameter instead of its content.

### Methods

Changing the struct parameter to be a receiver in the `iPanic` function is shown in this [Playground example](https://play.golang.org/p/2C4qVjoKNK5). Whether the struct a parameter of the function or a receiver for the method, __there is no actual difference in display the parameters list__. However what was `main.iPanic(...)` before, becomes `main.A.iPanic(...)` if the method has a __value receiver__ and `main.(*A).iPanic(...)` if the method has a __pointer receiver__.

Below is the code and the stack trace.

```go
import "fmt"

type A struct {
    i int
    s string
}

func main() {
    A{i: 50}.iPanic(true)
}

func (a A) iPanic(b bool) {
    fmt.Println(a.i + 1)
    panic("I'm outta here")
}
```

```txt
panic: I'm outta here

goroutine 1 [running]:
main.A.iPanic(0x32, 0x0, 0x0, 0xc00005e001)
    /tmp/sandbox467574958/prog.go:16 +0xa5
main.main()
    /tmp/sandbox467574958/prog.go:11 +0x37
```

Note that even if `a` is a _receiver_ outside the parameters list, it is actually referenced in the parameter list as in "main.**A**.iPanic(**0x32, 0x0, 0x0**, 0xc00005e001)" with its content. When we change the receiver to a pointer, `a` is referenced by it's address in the parameter list as in "main.**(*A)**.iPanic(**0xc000068f60**, 0x1)", as shown in the example below.

```go
import "fmt"

type A struct {
    i int
    s string
}

func main() {
    (&A{i: 50}).iPanic(true)
}

func (a *A) iPanic(b bool) {
    fmt.Println(a.i + 1)
    panic("I'm outta here")
}
```

```txt
panic: I'm outta here

goroutine 1 [running]:
main.(*A).iPanic(0xc000068f60, 0x1)
    /tmp/sandbox772882731/prog.go:16 +0xa7
main.main()
    /tmp/sandbox772882731/prog.go:11 +0x4f
```

## Interface

Interfaces, whether they are held by a struct or a pointer to a struct are always represented in the parameter list as two words holding a pointer to the information about the type stored and one to the object that they hold.

The two examples below can be explored with the basis of this [playground example](https://play.golang.org/p/9B28yoAw3AD).

### Struct implementing interface example

```go
import "fmt"

type oper interface {
    op() string
}

type A struct {
    s string
}

func (a A) op() string { return a.s }

func main() {
    iPanic(A{"op"})
}

func iPanic(o oper) {
    fmt.Println(o.op() + "!")
    panic("I'm outta here")
}
```

```txt
panic: I'm outta here

goroutine 1 [running]:
main.iPanic(0x4dc000, 0xc000010200)
    /tmp/sandbox220169591/prog.go:21 +0xe5
main.main()
    /tmp/sandbox220169591/prog.go:16 +0x50
```

### Pointer to struct implementing interface example

```go
import "fmt"

type oper interface {
    op() string
}

type A struct {
    s string
}

func (a *A) op() string { return a.s }

func main() {
    iPanic(&A{"op"})
}

func iPanic(o oper) {
    fmt.Println(o.op() + "!")
    panic("I'm outta here")
}
```

```txt
panic: I'm outta here

goroutine 1 [running]:
main.iPanic(0x4dbfa0, 0xc0001021e0)
    /tmp/sandbox587239589/prog.go:21 +0xe5
main.main()
    /tmp/sandbox587239589/prog.go:16 +0x59
```

## Maps, channels and function types

For maps, channels and function types as parameters, the stack trace will always show the reference to the values passed to the function. Here is the correspondent [playground example](https://play.golang.org/p/j8K-QO-cKy3) for these types.

```go
func main() {
    iPanic(make(map[int]bool), make(chan int), func() {})
}

func iPanic(m map[int]bool, c chan int, f func()) {
    go func() { f() }()
    panic("I'm outta here")
}
```

```txt
panic: I'm outta here

goroutine 1 [running]:
main.iPanic(0xc000032748, 0xc000094060, 0x479288)
    /tmp/sandbox760033282/prog.go:9 +0x5b
main.main()
    /tmp/sandbox760033282/prog.go:4 +0xc9
```

## Parting thoughts

I hope this article helps you unlock the secrets of the Go stack and, even though they prove to difficult to read, you may gain a better understanding of the meaning and the information that the stack is providing before heading to potentially long debugging sessions.

You can find me up on twitter @nikiforos_frees if you have any questions or comments and follow me on dev.to @mcaci

**This was Michele and thanks for reading!**

## References

Here is a list of references that helped me in doing my research on the go stack:

- <https://www.ardanlabs.com/blog/2015/01/stack-traces-in-go.html>
- <https://go101.org/article/type-system-overview.html>
- [Go Data Structures: Interfaces](https://research.swtch.com/interfaces)
- [runtime/stack.go](https://golang.org/src/runtime/stack.go)
