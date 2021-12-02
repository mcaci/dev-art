package bench

import (
	"sort"
	"testing"
)

func Benchmark3Sort(b *testing.B) {
	benchData := map[string]struct {
		size int
	}{
		"with size 1000":    {size: 1000},
		"with size 10000":   {size: 10000},
		"with size 100000":  {size: 100000},
		"with size 1000000": {size: 1000000},
	}
	b.ResetTimer()
	for benchName, data := range benchData {
		b.StopTimer()
		s := generateSlice(data.size)
		b.StartTimer()
		b.Run(benchName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.Ints(s)
			}
		})
	}
}
