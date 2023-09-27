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

package secret_test

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/secret"
	"github.com/stretchr/testify/assert"
)

func TestGuard_ValidSecret(t *testing.T) {
	// Create a test request with a valid secret in the X-Secret header.
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-Secret", base64.StdEncoding.EncodeToString(secret.Get()))

	// Create a ResponseRecorder to capture the response.
	rr := httptest.NewRecorder()

	// Create a dummy handler to use as the 'next' handler.
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Create a middleware instance.
	middleware := secret.Guard(dummyHandler)

	// Call the middleware.
	middleware.ServeHTTP(rr, req)

	// Check the response.
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGuard_InvalidSecret(t *testing.T) {
	// Create a test request with an invalid secret in the X-Secret header.
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-Secret", base64.StdEncoding.EncodeToString([]byte("invalid-secret")))

	// Create a ResponseRecorder to capture the response.
	rr := httptest.NewRecorder()

	// Create a dummy handler to use as the 'next' handler.
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("Next handler should not be called in case of an invalid secret")
	})

	// Create a middleware instance.
	middleware := secret.Guard(dummyHandler)

	// Call the middleware.
	middleware.ServeHTTP(rr, req)

	// Check the response.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "invalid secret")
}

func TestGuard_BadRequest(t *testing.T) {
	// Create a test request with a badly encoded secret in the X-Secret header.
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-Secret", "invalid-base64-data")

	// Create a ResponseRecorder to capture the response.
	rr := httptest.NewRecorder()

	// Create a dummy handler to use as the 'next' handler.
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("Next handler should not be called in case of a bad request")
	})

	// Create a middleware instance.
	middleware := secret.Guard(dummyHandler)

	// Call the middleware.
	middleware.ServeHTTP(rr, req)

	// Check the response.
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "bad request")
}
