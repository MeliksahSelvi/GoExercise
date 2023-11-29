package main

import "testing"

func BenchmarkFib(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = Fib(15)
	}
}

func BenchmarkFibNaive(b *testing.B) {

	memo := make([]int, 15)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = FibNaive(15, memo)
	}
}

func BenchmarkFibBottomUp(b *testing.B) {
	n := 15
	k := make([]int, n+1)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = FibBottomUp(n, k)
	}
}
