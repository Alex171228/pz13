package work

import "testing"

// BenchmarkFib — бенчмарк для медленной рекурсивной реализации.
func BenchmarkFib(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = Fib(30) // поменьше, чтобы бенчмарк не был слишком долгим
    }
}

// BenchmarkFibFast — бенчмарк для оптимизированной итеративной версии.
func BenchmarkFibFast(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = FibFast(30)
    }
}
