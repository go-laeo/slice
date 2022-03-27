package slice

func Contain[T comparable](items []T, one T) bool {
	for i := 0; i < len(items); i++ {
		if items[i] == one {
			return true
		}
	}
	return false
}
