package graphql_test

import (
	"encoding/json"
	"testing"

	"github.com/deepsquare-io/grid/cli/internal/graphql"
	"github.com/stretchr/testify/require"
)

type submitResponseData struct {
	Submit string `json:"submit"`
}

const fixture = `{"errors":[{"message":"template: gotpl:168:3: executing \"gotpl\" at \u003crenderStep $.Job $step\u003e: error calling renderStep: template: gotpl:50:3: executing \"gotpl\" at \u003crenderStepAsyncLaunch .Job .Step.Launch\u003e: error calling renderStepAsyncLaunch: template: gotpl:5:3: executing \"gotpl\" at \u003crenderStep $.Job $step\u003e: error calling renderStep: template: gotpl:46:3: executing \"gotpl\" at \u003crenderStepRun .Job .Step\u003e: error calling renderStepRun: template: gotpl:136:82: executing \"gotpl\" at \u003crenderSlirp4NetNS .Step.Run .Job (renderEnrootCommand .Step.Run)\u003e: error calling renderSlirp4NetNS: template: gotpl:92:3: executing \"gotpl\" at \u003crenderVNet $nic.VNet (printf \"vnet%d\" $i) $.Job\u003e: error calling renderVNet: missing virtual network: network","path":["submit"],"extensions":{"type":"internal"}}],"data":{"submit":""}}`

func TestParseResponse(t *testing.T) {
	var result graphql.Response[submitResponseData]

	err := json.Unmarshal([]byte(fixture), &result)
	require.NoError(t, err)
	require.Greater(t, len(result.Errors), 0)
}
