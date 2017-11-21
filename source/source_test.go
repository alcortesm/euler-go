package source_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alcortesm/euler-go/sink"
	"github.com/alcortesm/euler-go/source"
)

func TestFromSlice(t *testing.T) {
	for _, tt := range []struct {
		name string
		data []int
	}{
		{"empty", []int{}},
		{"one element", []int{0}},
		{"two elements", []int{0, 1}},
		{"many elements", []int{0, 1, 1000, -12, 42}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			obtained := sink.ToSlice(source.FromSlice(tt.data))
			if !reflect.DeepEqual(obtained, tt.data) {
				t.Errorf("expected: %v\nobtained: %v",
					tt.data, obtained)
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
		{2, []int{1, 1}},
		{3, []int{1, 1, 2}},
		{4, []int{1, 1, 2, 3}},
		{5, []int{1, 1, 2, 3}},
		{6, []int{1, 1, 2, 3, 5}},
		{7, []int{1, 1, 2, 3, 5}},
		{8, []int{1, 1, 2, 3, 5}},
		{9, []int{1, 1, 2, 3, 5, 8}},
		{10, []int{1, 1, 2, 3, 5, 8}},
		{11, []int{1, 1, 2, 3, 5, 8}},
		{12, []int{1, 1, 2, 3, 5, 8}},
		{13, []int{1, 1, 2, 3, 5, 8}},
		{14, []int{1, 1, 2, 3, 5, 8, 13}},
		{20, []int{1, 1, 2, 3, 5, 8, 13}},
		{21, []int{1, 1, 2, 3, 5, 8, 13}},
		{22, []int{1, 1, 2, 3, 5, 8, 13, 21}},
		{33, []int{1, 1, 2, 3, 5, 8, 13, 21}},
		{34, []int{1, 1, 2, 3, 5, 8, 13, 21}},
		{35, []int{1, 1, 2, 3, 5, 8, 13, 21, 34}},
		{54, []int{1, 1, 2, 3, 5, 8, 13, 21, 34}},
		{55, []int{1, 1, 2, 3, 5, 8, 13, 21, 34}},
		{56, []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}},
	} {
		name := fmt.Sprintf("fib smaller than %d", tt.ceil)
		t.Run(name, func(t *testing.T) {
			obtained := sink.ToSlice(source.Fib(tt.ceil))
			if !reflect.DeepEqual(obtained, tt.expected) {
				t.Errorf("expected %v\nobtained %v", tt.expected, obtained)
			}
		})
	}
}

func ExampleFib() {
	for n := range source.Fib(15) {
		fmt.Println(n)
	}
	// Output:
	// 1
	// 1
	// 2
	// 3
	// 5
	// 8
	// 13
}
