// Package multiples defines a calculator that returns all the multiples to a set
// of positive integers that are smaller than a certain maximum value, in ascending
// order and with no duplicates.
package multiples

import "fmt"

// Calculator generates multiples according to the package description.
// Its zero value is not safe, use the New function as a constructor.
type Calculator struct {
	bases     []int
	max       int
	factors   []int
	multiples []int
}

// NewCalculator returns a new calculator.  The numbers in bases will be used as the originators
// of the multiples to calculate.  The calculator will stop calculating multiples when
// it reaches the max value, not including it in the results.
// If the numbers in bases or max are not positive it will return an error.
func NewCalculator(bases []int, max int) (*Calculator, error) {
	if err := checkBases(bases); err != nil {
		return nil, err
	}
	if err := checkMax(max); err != nil {
		return nil, err
	}
	return &Calculator{
		bases:     bases,
		max:       max,
		factors:   ones(uint(len(bases))),
		multiples: make([]int, len(bases)),
	}, nil
}

func checkBases(bases []int) error {
	for i, b := range bases {
		if b < 1 {
			return fmt.Errorf("invalid base: base #%d is not positive (%d)", i+1, b)
		}
	}
	return nil
}

func checkMax(m int) error {
	if m < 1 {
		return fmt.Errorf("invalid max (%d): it is not positive", m)
	}
	return nil
}

// Ones returns a slice with n ones.
func ones(n uint) []int {
	ret := make([]int, int(n))
	for i := range ret {
		ret[i] = 1
	}
	return ret
}

const unknown = 0

// Next return the next multiple and true.  If there are no more multiples
// smaller than the calculator maximum, it will return 0 and false.
func (c *Calculator) Next() (int, bool) {
	pending := false
	for i, m := range c.multiples {
		switch {
		case m >= c.max:
			continue
		case m == unknown:
			c.multiples[i] = c.bases[i] * c.factors[i]
			if c.multiples[i] < c.max {
				c.factors[i]++
				pending = true
			}
		default:
			pending = true
		}
	}
	if !pending {
		return 0, false
	}
	ixs, v := mins(c.multiples)
	for _, i := range ixs {
		c.multiples[i] = unknown
	}
	return v, pending
}

// Mins returns the indexes and the value of the minimum number in s.
// If s has zero length, it will return nil and 0.
//
// For example: if s = {5, 3, 7, 12, 3, 5}, min(s) will return
// the indexes {1, 4} and the value 3.
func mins(s []int) ([]int, int) {
	if len(s) == 0 {
		return nil, 0
	}
	is := []int{0}
	v := s[0]
	for i, e := range s[1:] {
		if e == v {
			is = append(is, i+1)
			continue
		}
		if e < v {
			is = []int{i + 1}
			v = e
		}
	}
	return is, v
}
