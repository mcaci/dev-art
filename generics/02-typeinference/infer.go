package infer

func LoopOn[T any](t ...T) {
	for range t {
	}
}

func Run() {
	LoopOn(1, 2, 3)
	// LoopOn() // Cannot infer type
	LoopOn[any]()
}
