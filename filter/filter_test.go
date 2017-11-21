package filter_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alcortesm/euler-go/filter"
	"github.com/alcortesm/euler-go/sink"
	"github.com/alcortesm/euler-go/source"
)

func TestDrop(t *testing.T) {
	for _, tt := range []struct {
		name     string
		n        int
		data     []int
		expected []int
	}{
		{"drop 0 of empty", 0, []int{}, []int{}},
		{"drop 0 of one", 0, []int{1}, []int{1}},
		{"drop 0 of two", 0, []int{1, 2}, []int{1, 2}},
		{"drop 1 of one", 1, []int{1}, []int{}},
		{"drop 1 of two", 1, []int{1, 2}, []int{2}},
		{"drop 2 of two", 2, []int{1, 2}, []int{}},
		{"drop 5 of seven", 5,
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{6, 7}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := source.FromSlice(tt.data)
			obtained := sink.ToSlice(filter.Drop(tt.n, input))
			if !reflect.DeepEqual(tt.expected, obtained) {
				t.Errorf("expected = %v\nobtained = %v",
					tt.expected, obtained)
			}

		})
	}
}

func ExampleDrop() {
	input := source.FromSlice([]int{10, 20, 30, 40})
	for n := range filter.Drop(2, input) {
		fmt.Println(n)
	}
	// Output:
	// 30
	// 40
}
