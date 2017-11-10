package e0001

import "github.com/alcortesm/euler-go/e0001/multiples"

func Solution(bases []int, max int) (int, error) {
	c, err := multiples.NewCalculator(bases, max)
	if err != nil {
		return 0, err
	}

	s := 0
	for {
		m, ok := c.Next()
		if !ok {
			break
		}
		s += m
	}
	return s, nil
}
