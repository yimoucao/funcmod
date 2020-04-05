package funcmod

// ReduceFunc for reduce
type ReduceFunc func(x, y interface{}) interface{}

// Reduce reduce from left
func Reduce(it Iterable, f ReduceFunc, acc interface{}) interface{} {
	for {
		ele, ok := it.Next()
		if !ok {
			return acc
		}
		acc = f(acc, ele)
	}
}
