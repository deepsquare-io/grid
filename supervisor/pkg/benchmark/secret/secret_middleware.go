package secret

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func Guard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secretB64 := r.Header.Get("X-Secret")
		data, err := base64.StdEncoding.DecodeString(secretB64)
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf("bad request: %s", err),
				http.StatusBadRequest,
			)
			return
		}
		if !Validate(data) {
			http.Error(
				w,
				fmt.Sprintf("invalid secret: %s", err),
				http.StatusBadRequest,
			)
			return
		}
		next.ServeHTTP(w, r)
	})
}
