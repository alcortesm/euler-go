package main

import (
	"reflect"
	"testing"
)

func TestMultiplesOK(t *testing.T) {
	for _, test := range [...]struct {
		whole    int
		max      int
		expected []int
	}{
		{
			whole: 1, max: 9,
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		}, {
			whole: 2, max: 9,
			expected: []int{2, 4, 6, 8},
		}, {
			whole: 3, max: 9,
			expected: []int{3, 6},
		}, {
			whole: 4, max: 9,
			expected: []int{4, 8},
		}, {
			whole: 5, max: 9,
			expected: []int{5},
		}, {
			whole: 6, max: 9,
			expected: []int{6},
		}, {
			whole: 9, max: 9,
			expected: []int{},
		}, {
			whole: 10, max: 9,
			expected: []int{},
		}, {
			whole: 1, max: 0,
			expected: []int{},
		},
	} {
		obtained := []int{}
		ms, err := multiples(test.whole, test.max)
		if err != nil {
			t.Fatal(err)
		}

		for m := range ms {
			obtained = append(obtained, m)
		}

		if !reflect.DeepEqual(obtained, test.expected) {
			t.Errorf("whole = %d, max = %d\nexpected = %v\nobtained = %v\n",
				test.whole, test.max, test.expected, obtained)
		}
	}
}

func TestMultiplesErrorNotCointingNumber(t *testing.T) {
	for _, counting := range []int{0, -1, -2, -3, -4} {
		if _, err := multiples(0, 3); err == nil {
			t.Errorf("err was nil with counting = %d", counting)
		}
	}
}

func TestMultiplesErrorNegativeMax(t *testing.T) {
	max := -1
	if _, err := multiples(1, -1); err == nil {
		t.Errorf("err was nil with max = %d", max)
	}
}

func TestSortedMultiplesOK(t *testing.T) {
	for _, test := range [...]struct {
		wholes   []int
		max      int
		expected []int
	}{
		{
			wholes: []int{}, max: 20,
			expected: []int{},
		}, {
			wholes: []int{3}, max: 20,
			expected: []int{3, 6, 9, 12, 15, 18},
		}, {
			wholes: []int{3, 7}, max: 20,
			expected: []int{3, 6, 7, 9, 12, 14, 18},
		},
	} {
		obtained := []int{}
		ms, err := sortedMultiples(test.wholes, test.max)
		if err != nil {
			t.Fatal(err)
		}

		for m := range ms {
			obtained = append(obtained, m)
		}

		if !reflect.DeepEqual(obtained, test.expected) {
			t.Errorf("wholes = %v, max = %d\nexpected = %v\nobtained = %v\n",
				test.wholes, test.max, test.expected, obtained)
		}
	}
}

func TestSortedMultiplesOKNoRepetitions(t *testing.T) {
	for _, test := range [...]struct {
		wholes   []int
		max      int
		expected []int
	}{
		{
			wholes: []int{1, 2}, max: 6,
			expected: []int{1, 2, 3, 4, 5},
		}, {
			wholes: []int{3, 5}, max: 20,
			expected: []int{3, 5, 6, 7, 9, 10, 12, 15, 18},
		},
	} {
		obtained := []int{}
		ms, err := sortedMultiples(test.wholes, test.max)
		if err != nil {
			t.Fatal(err)
		}

		for m := range ms {
			obtained = append(obtained, m)
		}

		if !reflect.DeepEqual(obtained, test.expected) {
			t.Errorf("wholes = %v, max = %d\nexpected = %v\nobtained = %v\n",
				test.wholes, test.max, test.expected, obtained)
		}
	}
}
