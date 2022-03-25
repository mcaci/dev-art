package typeparam

import "fmt"

type PairF32 struct {
	a, b float32
}

func (p PairF32) String() string {
	return fmt.Sprintf("[%v, %v]", p.a, p.b)
}

func (p *PairF32) Rotate(t float32) float32 {
	t, p.a, p.b = p.b, t, p.a
	return t
}
