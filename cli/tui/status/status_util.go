package status

func ignoreElements[T any](ch <-chan T) {
	for range ch {
	}
}
