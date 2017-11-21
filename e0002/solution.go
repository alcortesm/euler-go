package e0002

import (
	"github.com/alcortesm/euler-go/filter"
	"github.com/alcortesm/euler-go/sink"
	"github.com/alcortesm/euler-go/source"
)

// Solution returns the solution to problem 2
func Solution() int {
	return solution(4000000)
}

func solution(ceil int) int {
	fibFrom2 := filter.Drop(1, source.Fib(ceil))
	return sink.Sum(keepEven(fibFrom2))
}

func keepEven(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		for n := range input {
			if n%2 == 0 {
				output <- n
			}
		}
		close(output)
	}()
	return output
}
