package e0001

import (
	"testing"
)

func TestSolution(t *testing.T) {
	const expected = 233168
	obtained, err := Solution()
	if err != nil {
		t.Error(err)
	}
	if obtained != expected {
		t.Errorf("expected 100, got %d", obtained)
	}
}

func BenchmarkSolutionMaxIs1k(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solution([]int{3, 5}, 1000)
	}
}

func BenchmarkSolutionMaxIs10k(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solution([]int{3, 5}, 10000)
	}
}

func BenchmarkSolutionMaxIs100k(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solution([]int{3, 5}, 100000)
	}
}

func BenchmarkSolutionManyMultiplesInBases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solution([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, 10000)
	}
}

func BenchmarkSolutionNOMultiplesInBases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solution([]int{1}, 10000)
	}
}
