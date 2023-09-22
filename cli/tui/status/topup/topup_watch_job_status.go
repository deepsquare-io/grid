package topup

import (
	"context"

	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	"github.com/deepsquare-io/the-grid/cli/internal/log"
	"github.com/deepsquare-io/the-grid/cli/tui/channel"
	"github.com/deepsquare-io/the-grid/cli/types"
	"go.uber.org/zap"
)

type transitionMsg types.Job

func makeWatchJobModel(
	ctx context.Context,
	jobID [32]byte,
	watcher deepsquare.Watcher,
	client deepsquare.Client,
) channel.Model[transitionMsg] {
	return channel.Model[transitionMsg]{
		Channel: make(chan transitionMsg, 1),
		OnInit: func(c chan transitionMsg) func() error {
			transitions := make(chan types.JobTransition, 1)
			sub, err := watcher.SubscribeEvents(ctx,
				types.FilterJobTransition(transitions),
			)
			if err != nil {
				log.I.Fatal(err.Error())
			}

			// Get initial value
			func() {
				job, err := client.GetJob(ctx, jobID)
				if err != nil {
					log.I.Error(
						"failed to get job from transition, ignoring...",
						zap.Error(err),
					)
					return
				}
				select {
				case c <- transitionMsg(job):
				default:
					// Skip when buffer full
				}
			}()

			go func() {
				defer sub.Unsubscribe()

				for {
					select {
					case transition := <-transitions:
						if jobID != transition.JobId {
							continue
						}
						go func() {
							job, err := client.GetJob(ctx, jobID)
							if err != nil {
								log.I.Error(
									"failed to get job from transition, ignoring...",
									zap.Error(err),
								)
								return
							}
							select {
							case c <- transitionMsg(job):
							default:
								// Skip when buffer full
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
