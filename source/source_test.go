package source_test

import (
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
