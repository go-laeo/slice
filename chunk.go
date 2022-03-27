package slice

// Chunk chunks an array into arrays with length elements. The last chunk may contain less than length elements.
func Chunk[T any](values []T, length int) [][]T {
	if length < 1 {
		panic("chunk: length must not less than 1")
	}

	chunks := make([][]T, (len(values)+length-1)/length)
	for i, j := 0, 0; i < len(values); i++ {
		if i > 0 && i%length == 0 {
			j++
		}

		if chunks[j] == nil {
			chunks[j] = make([]T, 0)
		}

		chunks[j] = append(chunks[j], values[i])
	}

	return chunks
}
