package bench

import (
	"sort"
	"testing"
)

func Benchmark2aSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := generateSlice(1000)
		b.StartTimer()
		sort.Ints(s)
	}
}

func Benchmark2bSort(b *testing.B) {
	s := generateSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(s)
	}
}
