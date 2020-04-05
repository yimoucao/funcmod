package funcmod

// IterList return a iterable for list
func IterList(l []interface{}) Iterable {
	res := make(chan interface{})
	go func() {
		for _, ele := range l {
			res <- ele
		}
		close(res)
	}()
	return iterable(res)
}

// List lists out all elements in the iterable
func List(it Iterable) []interface{} {
	res := make([]interface{}, 0)
	for {
		ele, ok := it.Next()
		if !ok {
			break
		}
		res = append(res, ele)
	}
	return res
}
