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
	req *sbatchapiv1alpha1.ChallengeRequest,
) (*sbatchapiv1alpha1.ChallengeResponse, error) {
	return &sbatchapiv1alpha1.ChallengeResponse{
		Challenge: a.auth.Challenge(ctx, "login"),
	}, nil
}
