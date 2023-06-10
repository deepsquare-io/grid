package status

import (
	"context"

	"github.com/deepsquare-io/the-grid/cli"
	"github.com/deepsquare-io/the-grid/cli/internal/log"
	"github.com/deepsquare-io/the-grid/cli/tui/channel"
	"github.com/deepsquare-io/the-grid/cli/tui/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

type transitionMsg cli.Job

func makeWatchJobsModel(
	ctx context.Context,
	userAddress common.Address,
	logs chan types.Log,
	eventSubscriber cli.EventSubscriber,
	jobFilterer cli.JobFilterer,
	jobFetcher cli.JobFetcher,
) channel.Model[transitionMsg] {
	return channel.Model[transitionMsg]{
		Channel: make(chan transitionMsg, 100),
		OnInit: func(c chan transitionMsg) func() error {
			sub, err := eventSubscriber.SubscribeEvents(ctx, logs)
			if err != nil {
				log.I.Fatal(err.Error())
			}

			go func() {
				defer sub.Unsubscribe()
				transitions, rest := jobFilterer.FilterJobTransition(logs)
				newJobs, rest := jobFilterer.FilterNewJobRequests(rest)
				go util.IgnoreElements(rest)

				for {
					select {
					case transition := <-transitions:
						go func() {
							job, err := jobFetcher.GetJob(ctx, transition.JobId)
							if err != nil {
								log.I.Error(
									"failed to get job from transition, ignoring...",
									zap.Error(err),
								)
								return
							}
							if job.CustomerAddr != userAddress {
								return
							}
							select {
							case c <- transitionMsg(*job):
							case <-ctx.Done():
								// Context canceled. This is not an error.
								return
							}
						}()

					case newJob := <-newJobs:
						if newJob.CustomerAddr != userAddress {
							continue
						}
						go func() {
							job, err := jobFetcher.GetJob(ctx, newJob.JobId)
							if err != nil {
								log.I.Error(
									"failed to get new job request event, ignoring...",
									zap.Error(err),
								)
								return
							}
							select {
							case c <- transitionMsg(*job):
							case <-ctx.Done():
								// Context canceled. This is not an error.
								return
							}
						}()
					case <-ctx.Done():
						return
					}
				}
			}()

			return func() error {
				sub.Unsubscribe()
				return nil
			}
		},
	}
}
