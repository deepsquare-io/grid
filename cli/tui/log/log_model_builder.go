// Copyright (C) 2023 DeepSquare
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

package log

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/the-grid/cli/types"
	"github.com/ethereum/go-ethereum/common"
)

type ModelBuilder struct {
	Logger      types.Logger
	UserAddress common.Address
}

func (b *ModelBuilder) Build(ctx context.Context, jobID [32]byte) tea.Model {
	return Model(ctx, b.Logger, b.UserAddress, jobID)
}
