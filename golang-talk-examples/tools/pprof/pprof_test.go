package pprof_test

import (
	"testing"
)

func calculate(a, b int) (d int) {
	d = a + b
	for i := 0; i < d; i++ {
		tab := make([]string, d+1)
		tab[0] = "aaa"
	}
	return
}

func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculate(2, 1000)
	}
}
