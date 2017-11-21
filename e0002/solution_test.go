package e0002

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alcortesm/euler-go/sink"
	"github.com/alcortesm/euler-go/source"
)

func TestSolution(t *testing.T) {
	// Even fibonacci numbers smaller than 4000000:
	//   2, 8, 34, 144, 610, 2584, 10946, 46368, 196418, 832040, 3524578
	// their sum is
	//   4613732.
	expected := 4613732
	obtained := Solution()
	if obtained != expected {
		t.Errorf("expected %d, got %d", expected, obtained)
	}
}

func TestGenericSolution(t *testing.T) {
	for _, tt := range []struct {
		ceil, expected int
	}{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 2},   // 2
		{4, 2},   // 2
		{5, 2},   // 2
		{6, 2},   // 2
		{7, 2},   // 2
		{8, 2},   // 2
		{9, 10},  // 2 + 8
		{10, 10}, // 2 + 8
		{13, 10}, // 2 + 8
		{14, 10}, // 2 + 8
		{21, 10}, // 2 + 8
		{22, 10}, // 2 + 8
		{34, 10}, // 2 + 8
		{35, 44}, // 2 + 8 + 34
	} {
		name := fmt.Sprintf("ceil=%d", tt.ceil)
		t.Run(name, func(t *testing.T) {
			obtained := solution(tt.ceil)
			if obtained != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, obtained)
			}
		})
	}
}

func TestKeepEven(t *testing.T) {
	for _, tt := range []struct {
		name     string
		input    []int
		expected []int
	}{
		{"nil", nil, []int{}},
		{"empty", []int{}, []int{}},
		{"one odd", []int{43}, []int{}},
		{"one even", []int{42}, []int{42}},
		{"zero", []int{0}, []int{0}},
		{"negative odd", []int{-3}, []int{}},
		{"negative even", []int{-4}, []int{-4}},
		{"mixed evens and odds", []int{1, 5, 8, 12, 5, 2, 0, 43}, []int{8, 12, 2, 0}},
		{"all together", []int{-2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 1000, 1001, 5, 4, 3, 2, 1}, []int{-2, 0, 2, 4, 6, 1000, 4, 2}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := source.FromSlice(tt.input)
			obtained := sink.ToSlice(keepEven(input))
			if !reflect.DeepEqual(tt.expected, obtained) {
				t.Errorf("obtained and expected differ:\nexpected = %v\nobtained = %v",
					tt.expected, obtained)
			}
		})
	}
}
