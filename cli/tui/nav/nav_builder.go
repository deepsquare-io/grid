// Copyright (C) 2023 DeepSquare Association
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

	tea "github.com/charmbracelet/bubbletea"
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

// ModelBuilder contains the dependencies used to build the bubbletea Model for the navigator.
type ModelBuilder struct {
	userAddress          common.Address
	client               deepsquare.Client
	watcher              deepsquare.Watcher
	version              string
	availableVersion     string
	metaschedulerAddress string
}

// NewModelBuilder create a ModelBuilder.
func NewModelBuilder() *ModelBuilder {
	return &ModelBuilder{}
}

// WithUserAddress sets the user public address.
func (b *ModelBuilder) WithUserAddress(userAddress common.Address) *ModelBuilder {
	b.userAddress = userAddress
	return b
}

// WithClient sets the DeepSquare Client.
func (b *ModelBuilder) WithClient(client deepsquare.Client) *ModelBuilder {
	b.client = client
	return b
}

// WithWatcher sets the DeepSquare Watcher.
func (b *ModelBuilder) WithWatcher(watcher deepsquare.Watcher) *ModelBuilder {
	b.watcher = watcher
	return b
}

// WithVersion sets the version of the application.
func (b *ModelBuilder) WithVersion(version string) *ModelBuilder {
	b.version = version
	return b
}

func (b *ModelBuilder) WithAvailableVersion(version string) *ModelBuilder {
	b.availableVersion = version
	return b
}

// WithMetaschedulerAddress sets the Meta-Scheduler smart-contract address.
func (b *ModelBuilder) WithMetaschedulerAddress(metaschedulerAddress string) *ModelBuilder {
	b.metaschedulerAddress = metaschedulerAddress
	return b
}

// Build the bubbletea Model for the navigator.
func (b *ModelBuilder) Build(ctx context.Context) tea.Model {
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
		availableVersion:     b.availableVersion,
		metaschedulerAddress: b.metaschedulerAddress,
	}
}
