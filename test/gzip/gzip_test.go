package gzip

import (
	"testing"
)

func BenchmarkGzipWithoutPoll(b *testing.B) {
	gzip := NewGzipWithoutPool()

	b.ResetTimer()
	b.ReportAllocs()
	//b.N birim zamanda çalıştırılabilecek sayı
	for n := 0; n < b.N; n++ {
		err := gzip.Zip("Go Eğitim")

		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGzipWithPoll(b *testing.B) {
	gzip := NewGzipWithPool()

	b.ResetTimer()
	b.ReportAllocs()
	//b.N birim zamanda çalıştırılabilecek sayı
	for n := 0; n < b.N; n++ {
		err := gzip.Zip("Go Eğitim")

		if err != nil {
			b.Fatal(err)
		}
	}
}
