package bench

import (
	"math/rand"
	"sort"
	"testing"
)

func generateSlice(n int) []int {
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, rand.Intn(1e9))
	}
	return s
}

func Benchmark1Sort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Ints(generateSlice(1000))
	}
}

func Test1Sort(t *testing.T) {
	slice := generateSlice(1000)
	if len(slice) != 1000 {
		t.Errorf("unexpected slice size: %d", len(slice))
	}
}
