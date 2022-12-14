package utils

// Ptr is a helper for setting const variable as pointer
func Ptr[T any](it T) *T {
	return &it
}
