package funcmod

// IterChan return a iterable for channel
func IterChan(ch <-chan interface{}) Iterable {
	return iterable(ch)
}
