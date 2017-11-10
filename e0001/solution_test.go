package e0001_test

import (
	"testing"

	"github.com/alcortesm/euler-go/e0001"
)

func TestSolution(t *testing.T) {
	for _, test := range [...]struct {
		wholes   []int
		max      int
		expected int
	}{
		{
			wholes:   []int{1},
			max:      1,
			expected: 0,
		}, {
			wholes:   []int{1},
			max:      10,
			expected: 45,
		}, {
			wholes:   []int{1, 2},
			max:      5,
			expected: 10,
		}, {
			wholes:   []int{3, 5},
			max:      10,
			expected: 23,
		}, {
			wholes:   []int{5, 3},
			max:      10,
			expected: 23,
		}, {
			wholes:   []int{5, 3, 11},
			max:      100,
			expected: 2560,
		}, {
			wholes:   []int{3, 5},
			max:      1000,
			expected: 233168,
		}, {
			wholes:   []int{3, 5, 11, 2},
			max:      7,
			expected: 20,
		},
	} {
		obtained, err := e0001.Solution(test.wholes, test.max)
		if err != nil {
			t.Errorf("wholes = %v, max = %d\nexpected = %v\nerror = %v\n",
				test.wholes, test.max, test.expected, err)
		}
		if obtained != test.expected {
			t.Errorf("wholes = %v, max = %d\nexpected = %d\nobtained = %d\n",
				test.wholes, test.max, test.expected, obtained)
		}
	}
}

func TestSolutionErrorSmallMax(t *testing.T) {
	_, err := e0001.Solution([]int{}, -1)
	if err == nil {
		t.Errorf("err is nil")
	}
}

func TestSolutionErrorInvalidBases(t *testing.T) {
	for _, bases := range [...][]int{
		{0},
		{-1},
		{0, 1},
		{1, 0},
		{-3, 1},
		{1, -3},
		{1, 2, -1, 3, 4},
	} {
		_, err := e0001.Solution(bases, 14)
		if err == nil {
			t.Errorf("bases = %v, no error detected", bases)
		}
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
