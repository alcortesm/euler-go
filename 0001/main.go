package main

import (
	"fmt"
	"log"
)

func main() {
	s, err := sumUniqMultiples([]int{3, 5}, 1000)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
}

// Returns the sum of all non repeted multiples of bases lower than max.
//
// For example: sumUniqMultiples([]int{3, 5}, 18) would be 60, this is,
// the sum of:
//
// - 3, 6, 9, 12, 15 (multiples of 3 lower than 18)
//
// - and 5, 10 (multiples of 5 lower than 18, not counting 15, as it is
// already a multiple of 3).
//
// Returns an error if any of the bases is less than 1 or if max is
// negative.
func sumUniqMultiples(bases []int, max int) (int, error) {
	if err := checkBases(bases); err != nil {
		return 0, err
	}
	if err := checkMax(max); err != nil {
		return 0, err
	}

	ms := multiplesForAll(bases, max)

	s := 0
	for n := range uniqOfSorted(ms) {
		s += n
	}

	return s, nil
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

// Returns a channel of uniq and sorted integers that are the result of
// merging the contents of a slice of channels with sorted integers.
func uniqOfSorted(cs []<-chan int) <-chan int {
	uniq := make(chan int)
	go func() {
		heads := make([]int, len(cs))
		for {
			min := 0
			imin := -1
			for i := range heads {
				if heads[i] == 0 {
					heads[i] = <-cs[i]
					if heads[i] == 0 {
						continue
					}
				}

				if imin == -1 {
					min = heads[i]
					imin = i
					continue
				}

				if heads[i] == min {
					heads[i] = 0 // remove repeated
				} else if heads[i] < min {
					min = heads[i]
					imin = i
				}
			}
			if imin == -1 {
				break
			}
			uniq <- heads[imin]
			heads[imin] = 0
		}
		close(uniq)
	}()
	return uniq
}

func multiplesForAll(countings []int, max int) []<-chan int {
	ms := make([]<-chan int, 0, len(countings))
	for _, b := range countings {
		c := multiples(b, max)
		ms = append(ms, c)
	}
	return ms
}

// Returns a channel to receive the multiples of counting, up to,
// and not including, max.  The numbers are received over the channel in
// increasing order.
//
// Counting must be a counting number and max must be positive.
//
// Example: multiples(3, 12) will return a channel with the numbers 3,
// 6, 9, in this same order.
func multiples(counting int, max int) <-chan int {
	multiples := make(chan int)
	go func() {
		i := 1
		for {
			m := i * counting
			if m >= max {
				break
			}
			multiples <- m
			i++
		}
		close(multiples)
	}()
	return multiples
}
