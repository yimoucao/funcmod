package funcmod

// IterChan return a iterable for channel
// NOTE: channle could be blocking
func IterChan(ch <-chan interface{}) Iterable {
	return iterable(ch)
}
