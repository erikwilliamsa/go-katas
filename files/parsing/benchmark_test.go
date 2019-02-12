package parsing

import "testing"

const path = "./_resources/bm"

func BenchmarkBlockingFileStats(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		BlockingAllFileStats(path)
	}
}

func BenchmarkConcurrentFileStats(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ConcurrentAllFileStats(path)
	}
}

func BenchmarkPipeFileStats(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		PipedGetAllFileStats(path)
	}
}
