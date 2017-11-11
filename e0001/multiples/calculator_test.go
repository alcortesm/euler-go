package multiples_test

import (
	"fmt"
	"testing"

	"github.com/alcortesm/euler-go/e0001/multiples"
)

func TestNewErrorInvalidMax(t *testing.T) {
	for _, max := range []int{0, -1, -10} {
		name := fmt.Sprintf("max=%v", max)
		t.Run(name, func(t *testing.T) {
			if _, err := multiples.Calculator([]int{1, 2}, 0); err == nil {
				t.Errorf("expected an error, but no error found")
			}
		})
	}
}

func TestNewErrorInvalidBases(t *testing.T) {
	for _, tt := range []struct {
		name  string
		bases []int
	}{
		{name: "a zero", bases: []int{0}},
		{name: "a negative", bases: []int{-1}},
		{name: "a zero among many positives",
			bases: []int{1, 5, 0, 3}},
		{name: "a negative among many positives",
			bases: []int{1, 5, -1, 3}},
		{name: "a zero at the begining",
			bases: []int{0, 5, 1, 3}},
		{name: "a negative at the begining",
			bases: []int{-1, 5, 1, 3}},
		{name: "a zero at the end",
			bases: []int{1, 5, 3, 0}},
		{name: "a negative at the end",
			bases: []int{1, 5, 3, -1}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := multiples.Calculator(tt.bases, 10); err == nil {
				t.Errorf("expected an error, but no error found")
			}
		})
	}
}

func TestCalculator(t *testing.T) {
	for _, tt := range [...]struct {
		name     string
		bases    []int
		max      int
		expected []int
	}{
		{
			name:     "multiples of 1 smaller than 1",
			bases:    []int{1},
			max:      1,
			expected: []int{},
		}, {
			name:     "multiples of 1 smaller than 2",
			bases:    []int{1},
			max:      2,
			expected: []int{1},
		}, {
			name:     "multiples of 1 smaller than 5",
			bases:    []int{1},
			max:      5,
			expected: []int{1, 2, 3, 4},
		}, {
			name:     "multiples of 3 smaller than itself",
			bases:    []int{3},
			max:      3,
			expected: []int{},
		}, {
			name:     "one multiple of 3",
			bases:    []int{3},
			max:      5,
			expected: []int{3},
		}, {
			name:     "3 multiples of 3",
			bases:    []int{3},
			max:      10,
			expected: []int{3, 6, 9},
		}, {
			name:     "multiples of 1 and 2 smaller than 2",
			bases:    []int{1, 2},
			max:      2,
			expected: []int{1},
		}, {
			name:     "first multiple of a pair",
			bases:    []int{2, 3},
			max:      4,
			expected: []int{2, 3},
		}, {
			name:     "there should not be duplicates",
			bases:    []int{1, 2},
			max:      5,
			expected: []int{1, 2, 3, 4},
		}, {
			name:     "multiples of first but not of second",
			bases:    []int{3, 5},
			max:      10,
			expected: []int{3, 5, 6, 9},
		}, {
			name:     "multiples of second but not of first",
			bases:    []int{5, 3},
			max:      10,
			expected: []int{3, 5, 6, 9},
		}, {
			name:  "lots of multiples",
			bases: []int{5, 3, 11},
			max:   50,
			expected: []int{
				3, 5, 6, 9, 10,
				11, 12, 15, 18,
				20, 21, 22, 24, 25, 27,
				30, 33, 35, 36, 39,
				40, 42, 44, 45, 48},
		}, {
			name:     "lots of bases but small max",
			bases:    []int{3, 5, 11, 2},
			max:      7,
			expected: []int{2, 3, 4, 5, 6},
		}, {
			name:     "lots of multiples in the bases",
			bases:    []int{2, 4, 8},
			max:      13,
			expected: []int{2, 4, 6, 8, 10, 12},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			c, err := multiples.Calculator(tt.bases, tt.max)
			if err != nil {
				t.Fatalf("cannot create calculator: %v", err)
			}
			for i, e := range tt.expected {
				obtained, ok := <-c
				if !ok {
					t.Errorf("calculator is empty prematurely at `next` call #%d", i+1)
				}
				if obtained != e {
					t.Errorf("expected #%d value was %d, but got %d", i+1, e, obtained)
				}
			}
			if obtained, ok := <-c; ok {
				t.Errorf("calculator should be empty, but it return a number (%d)", obtained)
			}
		})
	}
}
