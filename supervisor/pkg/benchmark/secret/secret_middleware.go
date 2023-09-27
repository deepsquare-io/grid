// Copyright (C) 2023 DeepSquare Asociation
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
