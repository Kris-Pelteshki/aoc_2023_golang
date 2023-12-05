package collections

func Chunks[T any](items []T, chunkSize int) (chunks [][]T) {
	itemsLen := len(items)

	for i := 0; i < itemsLen; i += chunkSize {
		end := i + chunkSize
		if end > itemsLen {
			end = itemsLen
		}
		chunks = append(chunks, items[i:end])
	}
	return chunks
}
