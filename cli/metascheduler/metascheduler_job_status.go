// Copyright (C) 2023 DeepSquare Association
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

package metascheduler

// JobStatus is the job status stored in the smart-contract.
type JobStatus uint8

const (
	// JobStatusPending means the job has been submitted, but the meta-scheduler
	// hasn't considered to schedule it yet.
	JobStatusPending JobStatus = 0
	// JobStatusMetaScheduled means that the job has been scheduled to a cluster, but that cluster has not yet handled it.
	JobStatusMetaScheduled JobStatus = 1
	// JobStatusScheduled means that the job has been scheduled to a worker node.
	JobStatusScheduled JobStatus = 2
	// JobStatusRunning means that the job is running on a worker node.
	JobStatusRunning JobStatus = 3
	// JobStatusCancelled means that the job has been cancelled by the user or admin.
	JobStatusCancelled JobStatus = 4
	// JobStatusFinished means that the job has finished with success.
	JobStatusFinished JobStatus = 5
	// JobStatusFailed means that the job has failed.
	JobStatusFailed JobStatus = 6
	// JobStatusOutOfCredits means that the job has timed out.
	JobStatusOutOfCredits JobStatus = 7
	// JobStatusPanicked means that the job has failed from an unexpected error.
	JobStatusPanicked JobStatus = 8
	// JobStatusUnknown is a unknown state.
	JobStatusUnknown JobStatus = 255
)

func (s JobStatus) String() string {
	switch s {
	case JobStatusPending:
		return "Pending"
	case JobStatusMetaScheduled:
		return "MetaScheduled"
	case JobStatusScheduled:
		return "Scheduled"
	case JobStatusRunning:
		return "Running"
	case JobStatusCancelled:
		return "Cancelled"
	case JobStatusFinished:
		return "Finished"
	case JobStatusFailed:
		return "Failed"
	case JobStatusOutOfCredits:
		return "OutOfCredits"
	case JobStatusPanicked:
		return "Panicked"
	case JobStatusUnknown:
		return "Unknown"
	}
	return "Unknown"
}
