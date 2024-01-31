// Copyright (C) 2024 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Package types provides the main types of the library.
package types

import (
	"context"

	loggerv1alpha1 "github.com/deepsquare-io/grid/cli/internal/logger/v1alpha1"
	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
)

// LogStream is a readable stream of logs.
type LogStream loggerv1alpha1.LoggerAPI_ReadClient

// Logger fetches the logs of a job.
type Logger interface {
	// Watch the logs of a job
	WatchLogs(ctx context.Context, jobID [32]byte) (LogStream, error)
}

// Job is the object stored in the smart-contract for accounting.
type Job *metaschedulerabi.Job

// Label is a key-value object used for filtering and annotating clusters.
type Label metaschedulerabi.Label

// Affinity is a key-value object with an operator for filtering clusters.
type Affinity metaschedulerabi.Affinity

// NewJobRequest is an event that happens when a user submit a job.
type NewJobRequest *metaschedulerabi.MetaSchedulerNewJobRequestEvent

// JobTransition is an event that happens when the status of a job changes.
type JobTransition *metaschedulerabi.IJobRepositoryJobTransitionEvent

// Transfer is an event that happens when there is a ERC20 transaction.
type Transfer *metaschedulerabi.IERC20Transfer

// Approval is an event that happens when an user sets a new allowance.
type Approval *metaschedulerabi.IERC20Approval
