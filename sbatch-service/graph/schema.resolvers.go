package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.21 DO NOT EDIT.

import (
	"context"
	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/logger"
	"github.com/deepsquare-io/the-grid/sbatch-service/renderer"
	"go.uber.org/zap"
	shortuuid "github.com/lithammer/shortuuid/v4"
)

// Submit is the resolver for the submit field.
func (r *mutationResolver) Submit(ctx context.Context, job model.Job) (string, error) {
	script, err := renderer.RenderJob(&job)
	if err != nil {
		return "", err
	}
	u := shortuuid.New()
	logger.I.Info("set", zap.String("uuid", u), zap.String("script", script))
	_, err = r.RedisClient.Set(ctx, u, script, 0).Result()
	if err != nil {
		return "", err
	}
	return u, nil
}

// Job is the resolver for the job field.
func (r *queryResolver) Job(ctx context.Context, batchLocationHash string) (string, error) {
	logger.I.Info("get", zap.String("batchLocationHash", batchLocationHash))
	resp, err := r.RedisClient.Get(ctx, batchLocationHash).Result()
	if err != nil {
		return "", err
	}
	return resp, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
