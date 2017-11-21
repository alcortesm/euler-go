package sink_test

import (
	"reflect"
	"testing"

	"github.com/alcortesm/euler-go/sink"
	"github.com/alcortesm/euler-go/source"
)

func TestToSlice(t *testing.T) {
	for _, tt := range []struct {
		name string
		data []int
	}{
		{name: "empty", data: []int{}},
		{name: "one element", data: []int{1}},
		{name: "two elements", data: []int{1, 2}},
		{name: "many elements", data: []int{1, 2, 1000, -12, 42}},
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

func TestSum(t *testing.T) {
	for _, tt := range []struct {
		name     string
		input    []int
		expected int
	}{
		{name: "empty", input: []int{}, expected: 0},
		{name: "one element", input: []int{1}, expected: 1},
		{name: "two elements", input: []int{1, 2}, expected: 3},
		{name: "two negated elements", input: []int{42, -42}, expected: 0},
		{name: "many elements", input: []int{100, -12, 42}, expected: 130},
	} {
		t.Run(tt.name, func(t *testing.T) {
			obtained := sink.Sum(source.FromSlice(tt.input))
			if obtained != tt.expected {
				t.Errorf("expected: %v, obtained: %v", tt.expected, obtained)
			}
		})
	}
}
