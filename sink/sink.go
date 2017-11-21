// Package sink provides ways to completely drain a channel of integers,
// returning the data gathered in different ways.
package sink

// ToSlice returns a slice with all the numbers in the given channel.
func ToSlice(ch <-chan int) []int {
	ret := []int{}
	for n := range ch {
		ret = append(ret, n)
	}
	return ret
}

// Sum returns the sum of the numbers in the given channel.
func Sum(ch <-chan int) int {
	sum := 0
	for n := range ch {
		sum += n
	}
	return sum
}
