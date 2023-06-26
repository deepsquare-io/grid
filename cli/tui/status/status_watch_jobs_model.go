package status

import (
	"context"

	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	"github.com/deepsquare-io/the-grid/cli/internal/log"
	"github.com/deepsquare-io/the-grid/cli/tui/channel"
	"github.com/deepsquare-io/the-grid/cli/tui/util"
	"github.com/deepsquare-io/the-grid/cli/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

type transitionMsg types.Job

func makeWatchJobsModel(
	ctx context.Context,
	userAddress common.Address,
	logs chan ethtypes.Log,
	watcher deepsquare.Watcher,
	client deepsquare.Client,
) channel.Model[transitionMsg] {
	return channel.Model[transitionMsg]{
		Channel: make(chan transitionMsg, 100),
		OnInit: func(c chan transitionMsg) func() error {
			sub, err := watcher.SubscribeEvents(ctx, logs)
			if err != nil {
				log.I.Fatal(err.Error())
			}

			go func() {
				defer sub.Unsubscribe()
				transitions, rest := watcher.FilterJobTransition(logs)
				newJobs, rest := watcher.FilterNewJobRequests(rest)
				go util.IgnoreElements(rest)

				for {
					select {
					case transition := <-transitions:
						go func() {
							job, err := client.GetJob(ctx, transition.JobId)
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
							job, err := client.GetJob(ctx, newJob.JobId)
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
