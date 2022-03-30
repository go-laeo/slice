package slice

// Unique removes duplicate values from src.
func Unique[T comparable](src []T) []T {
	dst := make([]T, 0)

	if len(src) < 2 {
		return src
	}

	for i := 0; i < len(src); i++ {
		exists := false
	LOOKING:
		for j := 0; j < len(dst); j++ {
			if src[i] == dst[j] {
				exists = true
				break LOOKING
			}
		}

		if !exists {
			dst = append(dst, src[i])
		}
	}

	return dst
}
