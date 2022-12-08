package utils

func Ptr[T any](it T) *T {
	return &it
}
