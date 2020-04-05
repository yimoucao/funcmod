package funcmod

// TODO: Predicate must handle multiple args. Otherwise call it FilterFunc

// Predicate return true or false
type Predicate func(interface{}) bool

type filteriter struct {
	Iterable
	f Predicate
}

func (it *filteriter) Next() (interface{}, bool) {
	for {
		ele, ok := it.Iterable.Next()
		if !ok {
			return nil, false
		}
		if it.f(ele) {
			return ele, true
		}
	}
}

// Filter filters elements with given func
func Filter(pre Predicate, it Iterable) Iterable {
	return &filteriter{it, pre}
}
