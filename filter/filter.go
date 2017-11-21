// Package filter provides tools to modify the stream of data
// sent by a source.
package filter

// Drop removes the first n elements from the channel.
func Drop(n int, ch <-chan int) <-chan int {
	for i := 0; i < n; i++ {
		<-ch
	}
	return ch
}
