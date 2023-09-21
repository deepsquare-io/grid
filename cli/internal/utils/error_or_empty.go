package utils

import "fmt"

func ErrorfOrEmpty(format string, err error, va ...any) string {
	if err != nil {
		a := append([]any{err}, va...)
		return fmt.Sprintf(format, a...)
	}
	return ""
}
