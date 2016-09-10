package main

import (
	"fmt"
	"log"

	"github.com/alcortesm/euler-go/0001/peek"
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

	ms := allMultiples(bases, max)

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

func allMultiples(bases []int, max int) []<-chan int {
	ms := make([]<-chan int, 0, len(bases))
	for _, b := range bases {
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
func multiples(base int, max int) <-chan int {
	multiples := make(chan int)
	go func() {
		i := 1
		for {
			m := i * base
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

// Returns a channel of unique and sorted integers that are the result of
// merging the contents of a slice of channels holding sorted integers.
func uniqOfSorted(cs []<-chan int) <-chan int {
	ps := peekersFromChannels(cs)
	uniq := make(chan int)
	go func() {
		first := true
		var last int
		for {
			i, ok := min(ps)
			if !ok {
				break
			}
			n, _ := ps[i].Recv()
			if first {
				first = false
				last = n
				uniq <- n
				continue
			}
			if n != last {
				last = n
				uniq <- n
			}
		}
		close(uniq)
	}()
	return uniq
}

func peekersFromChannels(cs []<-chan int) []peek.Peeker {
	ps := make([]peek.Peeker, len(cs))
	for i, c := range cs {
		ps[i] = peek.NewChannel(c)
	}
	return ps
}

// Returns the index of the peeker with the smallest number and
// true if a there was any number at all in the peekers.
func min(ps []peek.Peeker) (int, bool) {
	var found bool
	var min int
	var imin int
	for i, p := range ps {
		n, ok := p.Peek()
		if !ok {
			continue
		}
		if !found {
			found = true
			min = n
			imin = i
			continue
		}
		if n < min {
			min = n
			imin = i
		}
	}
	return imin, found
}
