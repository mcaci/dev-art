package main

import "testing"

// go test -bench Day2 -run ^$ -benchmem

func BenchmarkDay2Part2Algo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algo1()
	}
}

func BenchmarkDay2Part2Algo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algo2()
	}
}

func BenchmarkDay2Part2Algo3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algo3()
	}
}

func BenchmarkDay2Part2Algo4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algo4()
	}
}
