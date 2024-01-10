// Copyright (C) 2024 DeepSquare Association
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
