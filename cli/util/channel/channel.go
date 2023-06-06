package channel

func IgnoreElements[T any](ch <-chan T) {
	for range ch {
	}
}
