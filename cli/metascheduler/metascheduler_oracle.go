// Copyright (C) 2023 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package metascheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deepsquare-io/grid/cli/types"
)

// DefaultOracleURL is the default oracle URL.
const DefaultOracleURL = "https://meta-scheduler.deepsquare.run"

type oracle struct {
	url  string
	opts OracleOptions
}

// OracleOptions are options for the oracle.
type OracleOptions struct {
	Client *http.Client
}

// NewOracle instanciates an Oracle.
func NewOracle(url string, opts OracleOptions) types.MetaScheduledJobsIdsFetcher {
	if url == "" {
		url = DefaultOracleURL
	}
	if opts.Client == nil {
		opts.Client = http.DefaultClient
	}
	return &oracle{
		url:  url,
		opts: opts,
	}
}

func (o *oracle) GetMetaScheduledJobIDs(ctx context.Context) ([][32]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", o.url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := o.opts.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(
			"failed to fetch running job ids: non-ok status code %d",
			resp.StatusCode,
		)
	}

	// Parse data
	var parsed map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, err
	}

	keys := make([][32]byte, 0, len(parsed))
	for k := range parsed {
		keys = append(keys, JobIDFromHex(k))
	}
	return keys, nil
}
