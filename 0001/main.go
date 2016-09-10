package main

import (
	"fmt"
	"log"
)

func main() {
	s, err := sum([]int{3, 5}, 1000)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
}

func sum(countings []int, max int) (int, error) {
	c, err := sortedMultiples(countings, max)
	if err != nil {
		return 0, err
	}

	s := 0
	for n := range c {
		s += n
	}

	return s, nil
}

// Returns a channel to receive the multiples of every element of
// countings, up to, and not including, max.  The numbers are received in
// increasing order and there will be no repetitions.  The channel is
// closed at the end.
//
// Example: sliceMultiples([]int{3, 5}, 20) will return a channel with
// the numbers 3, 5, 6, 9, 10, 12, 15 and 18 in this same order.  Note
// how only one 15 is received even though 15 is multiple of both 3 and
// 5.
func sortedMultiples(countings []int, max int) (<-chan int, error) {
	ms, err := multiplesForAll(countings, max)
	if err != nil {
		return nil, err
	}

	sorted := make(chan int)
	go func() {
		heads := make([]int, len(countings))
		for {
			min := 0
			imin := -1
			for i, _ := range heads {
				if heads[i] == 0 {
					heads[i] = <-ms[i]
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
			sorted <- heads[imin]
			heads[imin] = 0
		}
		close(sorted)
	}()
	return sorted, nil
}

func multiplesForAll(countings []int, max int) ([]<-chan int, error) {
	ms := make([]<-chan int, 0, len(countings))
	for _, b := range countings {
		c, err := multiples(b, max)
		if err != nil {
			// TODO close already opened channels
			return nil, err
		}
		ms = append(ms, c)
	}
	return ms, nil
}

// Returns a channel to receive the multiples of counting, up to,
// and not including, max.  The numbers are received over the channel in
// increasing order.
//
// Counting must be a counting number and max must be positive.
//
// Example: multiples(3, 12) will return a channel with the numbers 3,
// 6, 9, in this same order.
func multiples(counting int, max int) (<-chan int, error) {
	if counting < 1 {
		return nil, fmt.Errorf("multiples: invalid counting number: %d", counting)
	}
	if max < 0 {
		return nil, fmt.Errorf("multiples: max cannot be < 1, was %d", max)
	}
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
	return multiples, nil
}
