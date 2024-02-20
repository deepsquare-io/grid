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

package log

import (
	"context"

	"github.com/deepsquare-io/grid/cli/deepsquare"
	"github.com/deepsquare-io/grid/cli/internal/log"
	"github.com/deepsquare-io/grid/cli/metascheduler"
	"github.com/deepsquare-io/grid/cli/tui/channel"
	"github.com/deepsquare-io/grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
	"github.com/deepsquare-io/grid/cli/types/event"
)

type transitionMsg types.JobTransition

func makeWatchTransitionModel(
	ctx context.Context,
	jobID [32]byte,
	watcher *deepsquare.Watcher,
	client *deepsquare.Client,
) channel.Model[transitionMsg] {
	return channel.Model[transitionMsg]{
		Channel: make(chan transitionMsg, 100),
		OnInit: func(c chan transitionMsg) func() error {
			transitions := make(chan types.JobTransition, 1)
			sub, err := watcher.SubscribeEvents(ctx,
				event.FilterJobTransition(transitions),
			)
			if err != nil {
				log.I.Fatal(err.Error())
			}

			go func() {
				// Send initial state
				func() {
					job, err := client.GetJob(ctx, jobID)
					if err != nil {
						return
					}
					if metascheduler.JobStatus(job.Status) != metascheduler.JobStatusPending {
						c <- transitionMsg(&metaschedulerabi.IJobRepositoryJobTransitionEvent{
							JobId: jobID,
							From:  uint8(metascheduler.JobStatusUnknown),
							To:    job.Status,
						})
					}

				}()

				defer sub.Unsubscribe()

				for {
					select {
					case transition := <-transitions:
						go func() {
							select {
							case c <- transitionMsg(transition):
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
