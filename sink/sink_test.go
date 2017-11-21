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
