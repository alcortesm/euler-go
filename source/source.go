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
