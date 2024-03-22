// Copyright (C) 2024 DeepSquare Association
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

package auth

import (
	"context"

	internalauth "github.com/deepsquare-io/grid/sbatch-service/auth"
	sbatchapiv1alpha1 "github.com/deepsquare-io/grid/sbatch-service/gen/go/sbatchapi/v1alpha1"
	"github.com/deepsquare-io/grid/sbatch-service/logger"
)

type API struct {
	sbatchapiv1alpha1.UnimplementedAuthAPIServer
	auth *internalauth.Auth
}

func NewAPI(auth *internalauth.Auth) *API {
	if auth == nil {
		logger.I.Panic("auth is nil")
	}
	return &API{
		auth: auth,
	}
}

func (a *API) Challenge(
	ctx context.Context,
	_ *sbatchapiv1alpha1.ChallengeRequest,
) (*sbatchapiv1alpha1.ChallengeResponse, error) {
	return &sbatchapiv1alpha1.ChallengeResponse{
		Challenge: a.auth.Challenge(ctx, "login"),
	}, nil
}
