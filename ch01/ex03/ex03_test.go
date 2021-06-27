package ex03

import "testing"

func BenchmarkEchoUsingLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoUsingLoop()
	}
}

func BenchmarkEchoUsingStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoUsingStringJoin()
	}
}
