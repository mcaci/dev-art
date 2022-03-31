package typeparam

import "fmt"

type Pair[T any] struct {
	a, b T
}

func (p Pair[T]) String() string {
	return fmt.Sprintf("[%v, %v]", p.a, p.b)
}

func (p *Pair[T]) Rotate(t T) T {
	t, p.a, p.b = p.b, t, p.a
	return t
}
