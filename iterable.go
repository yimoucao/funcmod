package funcmod

// Iterable is iterable
type Iterable interface {
	// when no next element, reutrn false, end of iterable
	Next() (interface{}, bool)
}

type iterable <-chan interface{} // take care of read and write
func (it iterable) Next() (interface{}, bool) {
	ele, ok := <-it
	return ele, ok
}
