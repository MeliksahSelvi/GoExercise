package main

import "testing"

func BenchmarkPrintTree(b *testing.B) {
	root := InitializeNodes()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PrintTree(root)
	}
}

func BenchmarkAddNode(b *testing.B) {
	root := &Tree{value: 5}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		root = AddNode(root, &Tree{value: 3})
	}
}
