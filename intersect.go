package slice

// Intersect computes the intersection of array `a` and `b`.
func Intersect[T comparable](a, b []T) []T {
	if len(a) == 0 {
		return a
	}
	if len(b) == 0 {
		return b
	}

	n := len(a)
	for i := 0; i < n; i++ {
		found := false
	MATCH:
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				found = true
				break MATCH
			}
		}

		if !found {
			copy(a[i:], a[i+1:])
			n--
			i--
		}
	}
	return a[:n]
}
