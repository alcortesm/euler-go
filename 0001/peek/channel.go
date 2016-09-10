package peek

// A wrapper of go read channels with peek support.
type Channel struct {
	validHead bool
	head      int
	tail      <-chan int
}

// NewChannel returns a new Channel from a read channel.
func NewChannel(c <-chan int) *Channel {
	return &Channel{
		tail: c,
	}
}

// Recv extracts and returns the first value from the channel and true if it is valid.
// It returns 0 and false if there were no more values in the channel.
func (pc *Channel) Recv() (int, bool) {
	r, ok := pc.Peek()
	pc.validHead = false
	return r, ok
}

// Peek returns the first value in the channel and true if the value is
// valid.  It does not consume the value, therefore, consecutive calls
// to peek will return the same value.
//
// If there are no more data in the channel, it returns 0 and false.
func (pc *Channel) Peek() (int, bool) {
	if pc.validHead {
		return pc.head, true
	}

	var ok bool
	pc.head, ok = <-pc.tail
	if !ok {
		return 0, false
	}
	pc.validHead = true
	return pc.head, true
}
