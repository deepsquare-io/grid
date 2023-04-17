package singleton

// Map creates a instance on the left and a getter on the right
func Map[T any](m map[string]T, constructor func(key string) T) func(key string) T {
	return func(key string) T {
		if _, ok := m[key]; !ok {
			m[key] = constructor(key)
		}

		return m[key]
	}
}
