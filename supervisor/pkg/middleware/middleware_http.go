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
