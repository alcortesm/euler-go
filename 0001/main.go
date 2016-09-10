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

// Returns the sum of all unique multiples, lower than max,  of the numbers in bases.
// Example: sumUniqMultiples([]int{3, 5}, 18) would be 60 (3+5+6+9+10+12+15).
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
	u := uniq(ms)
	s := sum(u)
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

// Returns a channel and sends over it all the multiples of base, lower
// than max, in ascending order.  Closes the channel when done.
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

// Returns a channel of unique integers comming form a slice of channels
// of sorted integers.
func uniq(cs []<-chan int) <-chan int {
	ps := peekersFromChannels(cs)
	ret := make(chan int)
	go func() {
		first := true
		var last int
		for {
			ps = removeEmpties(ps)
			if len(ps) == 0 {
				break
			}
			i := indexOfMin(ps)
			n, _ := ps[i].Recv()
			if first {
				first = false
				last = n
				ret <- n
				continue
			}
			if n != last {
				last = n
				ret <- n
			}
		}
		close(ret)
	}()
	return ret
}

func peekersFromChannels(cs []<-chan int) []peek.Peeker {
	ps := make([]peek.Peeker, len(cs))
	for i, c := range cs {
		ps[i] = peek.NewChannel(c)
	}
	return ps
}

func removeEmpties(ps []peek.Peeker) []peek.Peeker {
	ret := make([]peek.Peeker, 0, len(ps))
	for _, p := range ps {
		if _, ok := p.Peek(); ok {
			ret = append(ret, p)
		}
	}
	return ret
}

// take a non empty slice of non empty peekers and returns the index of
// the peeker with the smallest integer.
func indexOfMin(ps []peek.Peeker) int {
	min, _ := ps[0].Peek()
	imin := 0
	for i := 1; i < len(ps); i++ {
		n, _ := ps[i].Peek()
		if n < min {
			min = n
			imin = i
		}
	}
	return imin
}

func sum(c <-chan int) int {
	s := 0
	for n := range c {
		s += n
	}
	return s
}
