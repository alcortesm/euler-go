package e0001_test

import (
	"testing"

	"github.com/alcortesm/euler-go/e0001"
)

func TestSolution(t *testing.T) {
	for _, test := range [...]struct {
		bases    []int
		max      int
		expected int
	}{
		{
			bases:    []int{1},
			max:      1,
			expected: 0,
		}, {
			bases:    []int{1},
			max:      10,
			expected: 45,
		}, {
			bases:    []int{1, 2},
			max:      5,
			expected: 10,
		}, {
			bases:    []int{3, 5},
			max:      10,
			expected: 23,
		}, {
			bases:    []int{5, 3},
			max:      10,
			expected: 23,
		}, {
			bases:    []int{5, 3, 11},
			max:      100,
			expected: 2560,
		}, {
			bases:    []int{3, 5},
			max:      1000,
			expected: 233168,
		}, {
			bases:    []int{3, 5, 11, 2},
			max:      7,
			expected: 20,
		},
	} {
		obtained, err := e0001.Solution(test.bases, test.max)
		if err != nil {
			t.Errorf("bases = %v, max = %d\nexpected = %v\nerror = %v\n",
				test.bases, test.max, test.expected, err)
		}
		if obtained != test.expected {
			t.Errorf("bases = %v, max = %d\nexpected = %d\nobtained = %d\n",
				test.bases, test.max, test.expected, obtained)
		}
	}
}

func TestSolutionError(t *testing.T) {
	_, err := e0001.Solution([]int{}, -1)
	if err == nil {
		t.Errorf("expected an error")
	}
}

func BenchmarkSolutionMaxIs1k(b *testing.B) {
	for n := 0; n < b.N; n++ {
		e0001.Solution([]int{3, 5}, 1000)
	}
}

func BenchmarkSolutionMaxIs10k(b *testing.B) {
	for n := 0; n < b.N; n++ {
		e0001.Solution([]int{3, 5}, 10000)
	}
}

func BenchmarkSolutionMaxIs100k(b *testing.B) {
	for n := 0; n < b.N; n++ {
		e0001.Solution([]int{3, 5}, 100000)
	}
}
