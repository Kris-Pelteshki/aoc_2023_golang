package collections

func CountBy[T any, K comparable](items []T, keyFunc func(T) K) map[K]int {
	counts := make(map[K]int)
	for _, item := range items {
		key := keyFunc(item)
		counts[key]++
	}
	return counts
}
