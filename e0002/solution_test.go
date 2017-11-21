package e0002

import (
	"fmt"
	"reflect"
	"testing"
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

func TestFib(t *testing.T) {
	for _, tt := range []struct {
		ceil     int
		expected []int
	}{
		{-1, []int{}},
		{0, []int{}},
		{1, []int{}},
		{2, []int{1}},
		{3, []int{1, 2}},
		{4, []int{1, 2, 3}},
		{5, []int{1, 2, 3}},
		{6, []int{1, 2, 3, 5}},
		{7, []int{1, 2, 3, 5}},
		{8, []int{1, 2, 3, 5}},
		{9, []int{1, 2, 3, 5, 8}},
		{10, []int{1, 2, 3, 5, 8}},
		{11, []int{1, 2, 3, 5, 8}},
		{12, []int{1, 2, 3, 5, 8}},
		{13, []int{1, 2, 3, 5, 8}},
		{14, []int{1, 2, 3, 5, 8, 13}},
		{20, []int{1, 2, 3, 5, 8, 13}},
		{21, []int{1, 2, 3, 5, 8, 13}},
		{22, []int{1, 2, 3, 5, 8, 13, 21}},
		{33, []int{1, 2, 3, 5, 8, 13, 21}},
		{34, []int{1, 2, 3, 5, 8, 13, 21}},
		{35, []int{1, 2, 3, 5, 8, 13, 21, 34}},
		{54, []int{1, 2, 3, 5, 8, 13, 21, 34}},
		{55, []int{1, 2, 3, 5, 8, 13, 21, 34}},
		{56, []int{1, 2, 3, 5, 8, 13, 21, 34, 55}},
	} {
		name := fmt.Sprintf("fib smaller than %d", tt.ceil)
		t.Run(name, func(t *testing.T) {
			testFib(t, tt.ceil, tt.expected)
		})
	}
}

func testFib(t *testing.T, ceil int, expected []int) {
	obtained := []int{}
	for n := range fib(ceil) {
		obtained = append(obtained, n)
	}
	if !reflect.DeepEqual(obtained, expected) {
		t.Errorf("expected %v\nobtained %v", expected, obtained)
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
			obtained := sink(keepEven(source(tt.input)))
			if !reflect.DeepEqual(tt.expected, obtained) {
				t.Errorf("obtained and expected differ:\nexpected = %v\nobtained = %v",
					tt.expected, obtained)
			}
		})
	}
}

func source(s []int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, n := range s {
			ch <- n
		}
		close(ch)
	}()
	return ch
}

func sink(ch <-chan int) []int {
	s := []int{}
	for n := range ch {
		s = append(s, n)
	}
	return s
}
