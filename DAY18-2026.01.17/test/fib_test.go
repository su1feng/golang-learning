// fib_test.go
package split

import (
	"testing"
)

func benchmarkFib(b *testing.B, n int) {
	Fib(n)
	b.ResetTimer() //重置时间
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

// 性能比较函数
func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }
