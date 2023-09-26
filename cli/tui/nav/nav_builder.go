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

package nav

import (
	"context"
	"math/big"

	_ "embed"

	"github.com/deepsquare-io/grid/cli/deepsquare"
	"github.com/deepsquare-io/grid/cli/tui/editor"
	"github.com/deepsquare-io/grid/cli/tui/provider"
	"github.com/deepsquare-io/grid/cli/tui/status"
	"github.com/deepsquare-io/grid/cli/tui/status/log"
	"github.com/deepsquare-io/grid/cli/tui/status/topup"
	"github.com/deepsquare-io/grid/cli/tui/transfer"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type ModelBuilder struct {
	userAddress          common.Address
	client               deepsquare.Client
	watcher              deepsquare.Watcher
	version              string
	metaschedulerAddress string
}

func NewModelBuilder() *ModelBuilder {
	return &ModelBuilder{}
}

func (b *ModelBuilder) WithUserAddress(userAddress common.Address) *ModelBuilder {
	b.userAddress = userAddress
	return b
}

func (b *ModelBuilder) WithClient(client deepsquare.Client) *ModelBuilder {
	b.client = client
	return b
}

func (b *ModelBuilder) WithWatcher(watcher deepsquare.Watcher) *ModelBuilder {
	b.watcher = watcher
	return b
}

func (b *ModelBuilder) WithVersion(version string) *ModelBuilder {
	b.version = version
	return b
}

func (b *ModelBuilder) WithMetaschedulerAddress(metaschedulerAddress string) *ModelBuilder {
	b.metaschedulerAddress = metaschedulerAddress
	return b
}

func (b *ModelBuilder) Build(ctx context.Context) *model {
	if b.client == nil {
		panic("client is nil")
	}
	if b.watcher == nil {
		panic("watcher is nil")
	}
	return &model{
		logs:          make(chan ethtypes.Log, 100),
		balanceChan:   make(chan *big.Int, 10),
		balance:       new(big.Int),
		allowanceChan: make(chan *big.Int, 10),
		allowance:     new(big.Int),

		statusModel: status.Model(ctx, b.client, b.watcher, b.userAddress),
		logModelBuilder: log.ModelBuilder{
			Logger: b.client,
		},
		editorModelBuilder: editor.ModelBuilder{
			Client: b.client,
		},
		transferModelBuilder: transfer.ModelBuilder{
			Client: b.client,
		},
		providerModelBuilder: provider.ModelBuilder{
			Client: b.client,
		},
		topupModelBuilder: topup.ModelBuilder{
			Client:  b.client,
			Watcher: b.watcher,
		},
		userAddress:          b.userAddress,
		client:               b.client,
		watcher:              b.watcher,
		version:              b.version,
		metaschedulerAddress: b.metaschedulerAddress,
	}
}
