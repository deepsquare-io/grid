package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"errors"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/logger"
	"github.com/deepsquare-io/the-grid/sbatch-service/validate"
	validator "github.com/go-playground/validator/v10"
	shortuuid "github.com/lithammer/shortuuid/v4"
	redis "github.com/redis/go-redis/v9"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
)

// Submit is the resolver for the submit field.
func (r *mutationResolver) Submit(ctx context.Context, job model.Job) (string, error) {
	script, err := r.JobRenderer.RenderJob(&job)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			for _, err := range errors {
				graphql.AddError(ctx, &gqlerror.Error{
					Path:    graphql.GetPath(ctx),
					Message: validate.Format(err),
					Extensions: map[string]interface{}{
						"type":             "validation",
						"tag":              err.Tag(),
						"namespace":        err.Namespace(),
						"field":            err.Field(),
						"original_message": err.Error(),
					},
				})
			}
			return "", nil
		} else {
			graphql.AddError(ctx, &gqlerror.Error{
				Path:    graphql.GetPath(ctx),
				Message: err.Error(),
				Extensions: map[string]interface{}{
					"type": "internal",
				},
			})
			return "", nil
		}
	}
	u := shortuuid.New()
	logger.I.Info("set", zap.String("uuid", u), zap.String("script", script))
	_, err = r.RedisClient.Set(ctx, u, script, 1*time.Hour).Result()
	if err != nil {
		graphql.AddError(ctx, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: err.Error(),
			Extensions: map[string]interface{}{
				"type": "internal",
			},
		})
		return "", nil
	}
	return u, nil
}

// Job is the resolver for the job field.
func (r *queryResolver) Job(ctx context.Context, batchLocationHash string) (string, error) {
	if r.Debug {
		logger.I.Info("get", zap.String("batchLocationHash", batchLocationHash))
		resp, err := r.RedisClient.Get(ctx, batchLocationHash).Result()
		if err != nil {
			if err == redis.Nil {
				return "", errors.New("no entry exists under this name")
			}
			return "", err
		}
		return resp, nil
	} else {
		return "", errors.New("debug mode is disabled and is not allowing query from graphql")
	}
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
