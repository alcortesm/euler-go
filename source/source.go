// Package source provides functions that returns channels
// over which finite sequences of numbers will be sent.
package source

func FromSlice(s []int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, n := range s {
			ch <- n
		}
	}()
	return ch
}

// Fib returns a channel with the numbers of the Fibonacci sequence
// starting from 1 and smaller than ceil.
func Fib(ceil int) <-chan int {
	ch := make(chan int)
	go func() {
		for a, b := 0, 1; b < ceil; a, b = b, a+b {
			ch <- b
		}
		close(ch)
	}()
	return ch
}
