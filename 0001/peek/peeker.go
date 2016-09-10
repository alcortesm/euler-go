package peek

// Peeker allows to extract or to peek the first value from a colletion.
//
// In both cases, the second returned value will be false
// if there are no more values in the collection.
type Peeker interface {
	Recv() (int, bool)
	Peek() (int, bool)
}
