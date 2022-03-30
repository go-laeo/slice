package slice

// Diff compares array `a` against `b` and returns the values in array `a` that are not present in `b`.
func Diff[T comparable](a, b []T) []T {
	n := len(a)
	for i := 0; i < n; i++ {
	MATCH:
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				copy(a[i:], a[i+1:])
				n--
				i--
				break MATCH
			}
		}
	}
	return a[:n]
}
