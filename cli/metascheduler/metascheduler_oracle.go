package metascheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deepsquare-io/grid/cli/types"
)

const DefaultOracleURL = "https://meta-scheduler.deepsquare.run"

type oracle struct {
	url  string
	opts OracleOptions
}

type OracleOptions struct {
	Client *http.Client
}

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
