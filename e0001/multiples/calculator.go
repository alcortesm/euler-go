// Package multiples defines a calculator that returns all the multiples to a set
// of positive integers that are smaller than a certain maximum value, in ascending
// order and with no duplicates.
package multiples

import (
	"fmt"
	"sort"
)

// Calculator runs a gorutine that will calculate multiples as explained
// in the package description.  It returns a channel where the multiples
// are sent.
//
// The numbers in bases will be used as the originators
// of the multiples to calculate.
//
// The calculator will stop and close the channel once it reaches
// a number bigger or equal to the max value, that will not be
// included in the results.
//
// If the numbers in bases or max are not positive it will return
// a nil channel and an error.
func Calculator(bases []int, max int) (<-chan int, error) {
	if err := checkBases(bases); err != nil {
		return nil, err
	}
	if err := checkMax(max); err != nil {
		return nil, err
	}
	clean := cleanBases(bases)
	c := calculator{
		bases:     clean,
		max:       max,
		factors:   ones(uint(len(clean))),
		multiples: make([]int, len(clean)),
	}
	ret := make(chan int)
	go func() {
		defer close(ret)
		for {
			n, ok := c.next()
			if !ok {
				break
			}
			ret <- n
		}
	}()
	return ret, nil
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

func cleanBases(bases []int) []int {
	if len(bases) < 2 {
		return bases
	}
	sorted := dup(bases)
	sort.Sort(sort.IntSlice(sorted))
	ret := make([]int, 1)
	ret[0] = sorted[0]
	for _, e := range sorted[1:] {
		if isMultiple(e, ret) {
			continue
		}
		ret = append(ret, e)
	}
	return ret
}

func dup(s []int) []int {
	d := make([]int, len(s))
	copy(d, s)
	return d
}

func isMultiple(x int, s []int) bool {
	for _, e := range s {
		if x%e == 0 {
			return true
		}
	}
	return false
}

type calculator struct {
	bases []int
	max   int
	// the factors to use to generate the next multiples from each base.
	factors []int
	// the smallest multiples for each base that has not yet being send
	// over the channel.  The special value `toBeCalculated` will be used
	// when a multiple has just been sent but the next one has not jet been
	// calculated.  When all multiples here are >= max, the calculation
	// has to end.
	multiples []int
}

const toBeCalculated = 0

// Ones returns a slice with n ones.
func ones(n uint) []int {
	ret := make([]int, int(n))
	for i := range ret {
		ret[i] = 1
	}
	return ret
}

// Next return the next multiple and true.  If there are no more multiples
// smaller than the calculator maximum, it will return 0 and false.
func (c *calculator) next() (int, bool) {
	// are there still numbers to return or have we reached the max on with all bases?
	pending := false
	for i, m := range c.multiples {
		switch {
		case m >= c.max:
			// ignore this multiple as it is already >= max
			continue
		case m == toBeCalculated:
			// a multiple that has just been sent or was a duplicated
			// of the one that has just been sent.  In either case,
			// the next multiple for its corresponding base and factor
			// has to be calculated.
			c.multiples[i] = c.bases[i] * c.factors[i]
			// if the generated multiple is below max, increment
			// its corresponding factor for the next round and
			// set the pending flag so we know there are new
			// data to be returned.
			if c.multiples[i] < c.max {
				c.factors[i]++
				pending = true
			}
		default:
			// a multiple that was generated in the previous runs
			// and is waiting its turn to be returned.
			pending = true
		}
	}
	if !pending {
		return 0, false
	}
	// find smallest multiple and its duplicates.
	// return its value and set them all to `toBeCalculated`
	// so we can generate the next multiples for each curresponding
	// base in the next iteration.
	indexIncludingDuplicates, value := mins(c.multiples)
	for _, i := range indexIncludingDuplicates {
		c.multiples[i] = toBeCalculated
	}
	return value, pending
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
