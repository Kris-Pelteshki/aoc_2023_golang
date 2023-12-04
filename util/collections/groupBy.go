package collections

func GroupBy[T any, K comparable](items []T, keyFunc func(T) K) map[K][]T {
	groups := make(map[K][]T)
	for _, item := range items {
		key := keyFunc(item)
		groups[key] = append(groups[key], item)
	}
	return groups
}
