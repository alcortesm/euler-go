package main

import "fmt"

func main() {
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
		for _, c := range ms {
			sorted <- <-c
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
