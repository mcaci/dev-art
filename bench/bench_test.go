package bench

import (
	"sort"
	"testing"
)

func BenchmarkAnything(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Ints(generateSlice(1000))
	}
}

func BenchmarkAnything2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := generateSlice(1000)
		b.StartTimer()
		sort.Ints(s)
	}
}

func BenchmarkAnything3(b *testing.B) {
	s := generateSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(s)
	}
}

func BenchmarkAnything4(b *testing.B) {
	benchData := map[string]int{
		"Run with 1000":   1000,
		"Run with 10000":  10000,
		"Run with 100000": 100000,
	}
	for name, bc := range benchData {
		s := generateSlice(bc)
		b.ResetTimer()
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.Ints(s)
			}
		})
	}
}
