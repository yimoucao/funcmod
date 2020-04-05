package funcmod

type chainiter struct {
	its []Iterable
	i   int
}

func (it chainiter) Next() (ele interface{}, ok bool) {
	for !ok && it.i < len(it.its) {
		ele, ok = it.its[it.i].Next()
		if !ok {
			it.i++
		}
	}
	return ele, ok
}

// Chain chain Iterables into one
func Chain(its ...Iterable) Iterable {
	return chainiter{its: its}
}
