package slice

// Reduce iteratively reduce the array `values` to a single value using a callback function `fn`.
func Reduce[T any, Y any](values []T, fn func(carry Y, item T) Y, initial Y) Y {
	for i := 0; i < len(values); i++ {
		initial = fn(initial, values[i])
	}
	return initial
}
