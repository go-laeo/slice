package slice

// Search searches by using function `fn` and returns its index.
func Search[T any](values []T, fn func(v T) bool) int {
	for i := 0; i < len(values); i++ {
		if fn(values[i]) {
			return i
		}
	}
	return -1
}
