package main

import "testing"

func BenchmarkDay3Part2Algo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day3_2_algo1()
	}
}

func BenchmarkDay3Part2Algo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day3_2_algo2()
	}
}

func BenchmarkDay3Part2Algo3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day3_2_algo3()
	}
}
