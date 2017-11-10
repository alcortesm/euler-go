package e0001

import (
	"fmt"
)

func Solution(bases []int, max int) (int, error) {
	if err := checkBases(bases); err != nil {
		return 0, err
	}
	if err := checkMax(max); err != nil {
		return 0, err
	}

	s := 0
	multiples := make([]int, len(bases))
	factors := ones(uint(len(bases)))
	const unknown = 0

	//fmt.Println(" - - - - - - SOLUTION", bases, max)
	for {
		pending := false
		//fmt.Println("before", multiples, bases, factors)
		for i, m := range multiples {
			//fmt.Println("inner", i, m)
			switch {
			case m >= max:
				continue
			case m == unknown:
				multiples[i] = bases[i] * factors[i]
				if multiples[i] < max {
					factors[i]++
					pending = true
				}
			default:
				pending = true
			}
		}
		//fmt.Println("after", multiples, bases, pending)
		if !pending {
			break
		}
		ixs, v := mins(multiples)
		//fmt.Println("mins", ixs, v)
		s += v
		for _, i := range ixs {
			multiples[i] = unknown
		}
	}
	return s, nil
}

// Ones returns a slice with n ones.
func ones(n uint) []int {
	ret := make([]int, int(n))
	for i := range ret {
		ret[i] = 1
	}
	return ret
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

func checkBases(bases []int) error {
	for i, b := range bases {
		if b < 1 {
			return fmt.Errorf("invalid bases: base %d is < 1 (%d)", i, b)
		}
	}
	return nil
}

func checkMax(m int) error {
	if m < 0 {
		return fmt.Errorf("invalid max %d: cannot be negative", m)
	}
	return nil
}
