package utils

func ErrorOrEmpty(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
