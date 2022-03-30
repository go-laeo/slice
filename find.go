package slice

// Find finds by using function `fn` and returns its value.
func Find[T any](values []T, fn func(v T) bool) (v T, found bool) {
	for i := 0; i < len(values); i++ {
		if fn(values[i]) {
			return values[i], true
		}
	}
	return
}
