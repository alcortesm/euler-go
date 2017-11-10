package e0001

import "github.com/alcortesm/euler-go/e0001/multiples"

// Solution solve exercise 0001.
func Solution(bases []int, max int) (int, error) {
	c, err := multiples.Calculator(bases, max)
	if err != nil {
		return 0, err
	}
	s := 0
	for m := range c {
		s += m
	}
	return s, nil
}
