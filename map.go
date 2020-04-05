package funcmod

// MapFunc for map
type MapFunc func(interface{}) interface{}

// mapiter implements Iterable
type mapiter struct {
	Iterable
	f MapFunc
}

func (it *mapiter) Next() (interface{}, bool) {
	ele, ok := it.Iterable.Next()
	if !ok {
		return nil, false
	}
	return it.f(ele), true
}

// Map map x to y
func Map(f MapFunc, it Iterable) Iterable {
	return &mapiter{it, f}
}
