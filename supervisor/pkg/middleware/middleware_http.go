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

package middleware

import (
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/deepsquare-io/grid/supervisor/logger"
)

type LoggingTransport struct {
	Transport http.RoundTripper
}

func (t *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := t.Transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		return nil, err
	}

	dumpStr := string(dump)
	if strings.Contains(dumpStr, "error") {
		logger.I.Error(dumpStr)
	} else {
		logger.I.Debug(dumpStr)
	}

	return res, nil
}
