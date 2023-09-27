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

package status

import (
	"context"

	"github.com/deepsquare-io/grid/cli/deepsquare"
	"github.com/deepsquare-io/grid/cli/internal/log"
	"github.com/deepsquare-io/grid/cli/tui/channel"
	"github.com/deepsquare-io/grid/cli/types"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type transitionMsg types.Job

func makeWatchJobsModel(
	ctx context.Context,
	userAddress common.Address,
	watcher deepsquare.Watcher,
	client deepsquare.Client,
) channel.Model[transitionMsg] {
	return channel.Model[transitionMsg]{
		Channel: make(chan transitionMsg, 100),
		OnInit: func(c chan transitionMsg) func() error {
			newJobs := make(chan types.NewJobRequest, 1)
			transitions := make(chan types.JobTransition, 1)
			sub, err := watcher.SubscribeEvents(ctx,
				types.FilterNewJobRequest(newJobs),
				types.FilterJobTransition(transitions),
			)
			if err != nil {
				log.I.Fatal(err.Error())
			}

			go func() {
				defer sub.Unsubscribe()

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
							case c <- transitionMsg(job):
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
							case c <- transitionMsg(job):
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
