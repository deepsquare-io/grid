package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"fmt"
	"time"

	"github.com/deepsquare-io/grid/smart-contracts-exporter/graph/model"
	"github.com/deepsquare-io/grid/smart-contracts-exporter/logger"
	metricsv1 "github.com/deepsquare-io/grid/smart-contracts-exporter/metrics/v1"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	pmodel "github.com/prometheus/common/model"
	"go.uber.org/zap"
)

// Max is the resolver for the max field.
func (r *jobDurationMetricsResolver) Max(ctx context.Context, obj *model.JobDurationMetrics, days int) (float64, error) {
	query := fmt.Sprintf("scalar(max(max by (wallet_address) (rate(metascheduler_jobs_duration_total_minutes{metascheduler_address=\"%s\"}[%dd]) * 86400 * %d)) or vector(0))", r.metaschedulerAddress, days, days)
	val, warnings, err := r.PromAPI.Query(ctx, query, time.Now(), v1.WithTimeout(10*time.Second))
	if len(warnings) > 0 {
		logger.I.Warn("Metrics thrown warnings", zap.Any("warnings", warnings))
	}
	if err != nil {
		logger.I.Error("Metrics thrown an error", zap.Error(err))
		return 0, err
	}

	scalarVal, ok := val.(*pmodel.Scalar)
	if !ok {
		logger.I.Panic("prometheus didn't return a scalar with Max", zap.Any("val", val))
	}

	return float64(scalarVal.Value), nil
}

// Average is the resolver for the average field.
func (r *jobDurationMetricsResolver) Average(ctx context.Context, obj *model.JobDurationMetrics, days int) (float64, error) {
	query := fmt.Sprintf("scalar(avg(max by (wallet_address) (rate(metascheduler_jobs_duration_total_minutes{metascheduler_address=\"%s\"}[%dd]) * 86400 * %d)) > 0 or vector(0))", r.metaschedulerAddress, days, days)
	val, warnings, err := r.PromAPI.Query(ctx, query, time.Now(), v1.WithTimeout(10*time.Second))
	if len(warnings) > 0 {
		logger.I.Warn("Metrics thrown warnings", zap.Any("warnings", warnings))
	}
	if err != nil {
		logger.I.Error("Metrics thrown an error", zap.Error(err))
		return 0, err
	}

	scalarVal, ok := val.(*pmodel.Scalar)
	if !ok {
		logger.I.Panic("prometheus didn't return a scalar with Average", zap.Any("val", val))
	}

	return float64(scalarVal.Value), nil
}

// Total is the resolver for the total field.
func (r *jobMetricsResolver) Total(ctx context.Context, obj *model.JobMetrics) (float64, error) {
	query := fmt.Sprintf("scalar(sum(avg by (wallet_address) (metascheduler_jobs_total{metascheduler_address=\"%s\"})) or vector(0))", r.metaschedulerAddress)
	val, warnings, err := r.PromAPI.Query(ctx, query, time.Now(), v1.WithTimeout(10*time.Second))
	if len(warnings) > 0 {
		logger.I.Warn("Metrics thrown warnings", zap.Any("warnings", warnings))
	}
	if err != nil {
		logger.I.Error("Metrics thrown an error", zap.Error(err))
		return 0, err
	}

	scalarVal, ok := val.(*pmodel.Scalar)
	if !ok {
		logger.I.Panic("prometheus didn't return a scalar with Total", zap.Any("val", val))
	}

	return float64(scalarVal.Value), nil
}

// RateRange is the resolver for the rateRange field.
func (r *jobMetricsResolver) RateRange(ctx context.Context, obj *model.JobMetrics, days int, startTime time.Time, endTime time.Time) ([]*model.TimestampValue, error) {
	query := fmt.Sprintf("sum(max by (wallet_address) (rate(metascheduler_jobs_total{metascheduler_address=\"%s\"}[%dd]) * 86400 * %d) or vector(0))", r.metaschedulerAddress, days, days)
	val, warnings, err := r.PromAPI.QueryRange(ctx, query, v1.Range{
		Start: startTime,
		End:   endTime,
		Step:  time.Duration(days) * 24 * time.Hour,
	}, v1.WithTimeout(10*time.Second))
	if len(warnings) > 0 {
		logger.I.Warn("CreditsSpent thrown warnings", zap.Any("warnings", warnings))
	}
	if err != nil {
		logger.I.Error("CreditsSpent thrown an error", zap.Error(err))
		return nil, err
	}

	matrixVal, ok := val.(pmodel.Matrix)
	if !ok {
		logger.I.Panic("prometheus didn't return a matrix with RateRange", zap.Any("val", val))
	}
	if len(matrixVal) < 1 {
		return []*model.TimestampValue{}, nil
	}

	out := make([]*model.TimestampValue, 0, len(matrixVal[0].Values))
	for _, v := range matrixVal[0].Values {
		out = append(out, &model.TimestampValue{
			Timestamp: int(v.Timestamp.Unix()),
			Value:     float64(v.Value),
		})
	}
	return out, nil
}

// CreditsMetrics is the resolver for the creditsMetrics field.
func (r *queryResolver) CreditsMetrics(ctx context.Context) (*model.CreditsMetrics, error) {
	query := fmt.Sprintf("scalar(sum(avg by (wallet_address) (metascheduler_credits_total{metascheduler_address=\"%s\"})) or vector(0))", r.metaschedulerAddress)
	val, warnings, err := r.PromAPI.Query(ctx, query, time.Now(), v1.WithTimeout(10*time.Second))
	if len(warnings) > 0 {
		logger.I.Warn("CreditsSpent thrown warnings", zap.Any("warnings", warnings))
	}
	if err != nil {
		logger.I.Error("CreditsSpent thrown an error", zap.Error(err))
		return nil, err
	}
	scalarVal, ok := val.(*pmodel.Scalar)
	if !ok {
		logger.I.Panic("prometheus didn't return a scalar with CreditsMetrics", zap.Any("val", val))
	}
	return &model.CreditsMetrics{
		SpentTotal: float64(scalarVal.Value),
	}, nil
}

// GpuTimeMetrics is the resolver for the gpuTimeMetrics field.
func (r *queryResolver) GpuTimeMetrics(ctx context.Context) (*model.GpuTimeMetrics, error) {
	query := fmt.Sprintf("scalar(sum(avg by (wallet_address) (metascheduler_gpu_total_minutes{metascheduler_address=\"%s\"})) or vector(0))", r.metaschedulerAddress)
	val, warnings, err := r.PromAPI.Query(ctx, query, time.Now(), v1.WithTimeout(10*time.Second))
	if len(warnings) > 0 {
		logger.I.Warn("GpuTime thrown warnings", zap.Any("warnings", warnings))
	}
	if err != nil {
		logger.I.Error("GpuTime thrown an error", zap.Error(err))
		return nil, err
	}
	scalarVal, ok := val.(*pmodel.Scalar)
	if !ok {
		logger.I.Panic("prometheus didn't return a scalar with GpuTimeMetrics", zap.Any("val", val))
	}
	return &model.GpuTimeMetrics{
		Total: float64(scalarVal.Value),
	}, nil
}

// CPUTimeMetrics is the resolver for the cpuTimeMetrics field.
func (r *queryResolver) CPUTimeMetrics(ctx context.Context) (*model.CPUTimeMetrics, error) {
	query := fmt.Sprintf("scalar(sum(avg by (wallet_address) (metascheduler_cpu_total_minutes{metascheduler_address=\"%s\"})) or vector(0))", r.metaschedulerAddress)
	val, warnings, err := r.PromAPI.Query(ctx, query, time.Now(), v1.WithTimeout(10*time.Second))
	if len(warnings) > 0 {
		logger.I.Warn("CPUTime thrown warnings", zap.Any("warnings", warnings))
	}
	if err != nil {
		logger.I.Error("CPUTime thrown an error", zap.Error(err))
		return nil, err
	}
	scalarVal, ok := val.(*pmodel.Scalar)
	if !ok {
		logger.I.Panic("prometheus didn't return a scalar with CPUTimeMetrics", zap.Any("val", val))
	}
	return &model.CPUTimeMetrics{
		Total: float64(scalarVal.Value),
	}, nil
}

// JobMetrics is the resolver for the jobMetrics field.
func (r *queryResolver) JobMetrics(ctx context.Context) (*model.JobMetrics, error) {
	return &model.JobMetrics{
		Duration: &model.JobDurationMetrics{},
	}, nil
}

// WalletMetrics is the resolver for the walletMetrics field.
func (r *queryResolver) WalletMetrics(ctx context.Context) (*model.WalletMetrics, error) {
	query := fmt.Sprintf("scalar(count(avg by (wallet_address) (metascheduler_jobs_total{metascheduler_address=\"%s\"})) or vector(0))", r.metaschedulerAddress)
	val, warnings, err := r.PromAPI.Query(ctx, query, time.Now(), v1.WithTimeout(10*time.Second))
	if len(warnings) > 0 {
		logger.I.Warn("WalletMetrics thrown warnings", zap.Any("warnings", warnings))
	}
	if err != nil {
		logger.I.Error("WalletMetrics thrown an error", zap.Error(err))
		return nil, err
	}
	scalarVal, ok := val.(*pmodel.Scalar)
	if !ok {
		logger.I.Panic("prometheus didn't return a scalar with WalletMetrics", zap.Any("val", val))
	}
	return &model.WalletMetrics{
		Count: int(scalarVal.Value),
	}, nil
}

// Job is the resolver for the job field.
func (r *subscriptionResolver) Job(ctx context.Context) (<-chan *model.Job, error) {
	return ConvertJobs(metricsv1.WatchLatest(ctx)), nil
}

// Top10 is the resolver for the top10 field.
func (r *walletMetricsResolver) Top10(ctx context.Context, obj *model.WalletMetrics, orderBy model.WalletOrderBy) ([]*model.Metric, error) {
	var query string
	switch orderBy {
	case model.WalletOrderByJobsSubmitted:
		query = fmt.Sprintf("topk(10, avg by (wallet_address) (metascheduler_jobs_total{metascheduler_address=\"%s\"}))", r.metaschedulerAddress)

	case model.WalletOrderByCreditSpent:
		query = fmt.Sprintf("topk(10, avg by (wallet_address) (metascheduler_credits_total{metascheduler_address=\"%s\"}))", r.metaschedulerAddress)
	}

	val, warnings, err := r.PromAPI.Query(ctx, query, time.Now(), v1.WithTimeout(10*time.Second))

	if len(warnings) > 0 {
		logger.I.Warn("Wallet thrown warnings", zap.Any("warnings", warnings))
	}
	if err != nil {
		logger.I.Error("Wallet thrown an error", zap.Error(err))
		return nil, err
	}

	vectorVal, ok := val.(pmodel.Vector)
	if !ok {
		logger.I.Panic("prometheus didn't return vector with Top10", zap.Any("val", val))
	}

	out := make([]*model.Metric, 0, 10)
	for _, val := range vectorVal {
		if w, ok := val.Metric["wallet_address"]; ok {
			out = append(out, &model.Metric{
				Key:   string(w),
				Value: float64(val.Value),
			})
		}
	}
	return out, nil
}

// JobDurationMetrics returns JobDurationMetricsResolver implementation.
func (r *Resolver) JobDurationMetrics() JobDurationMetricsResolver {
	return &jobDurationMetricsResolver{r}
}

// JobMetrics returns JobMetricsResolver implementation.
func (r *Resolver) JobMetrics() JobMetricsResolver { return &jobMetricsResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

// WalletMetrics returns WalletMetricsResolver implementation.
func (r *Resolver) WalletMetrics() WalletMetricsResolver { return &walletMetricsResolver{r} }

type jobDurationMetricsResolver struct{ *Resolver }
type jobMetricsResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type walletMetricsResolver struct{ *Resolver }
