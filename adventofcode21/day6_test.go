package main

import "testing"

func BenchmarkDay6Algo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day6algo1(100)
	}
}

func BenchmarkDay6Algo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day6algo2(100)
	}
}
