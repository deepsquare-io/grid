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

type JobStatus uint8

const (
	JobStatusPending       JobStatus = 0
	JobStatusMetaScheduled JobStatus = 1
	JobStatusScheduled     JobStatus = 2
	JobStatusRunning       JobStatus = 3
	JobStatusCancelled     JobStatus = 4
	JobStatusFinished      JobStatus = 5
	JobStatusFailed        JobStatus = 6
	JobStatusOutOfCredits  JobStatus = 7
	JobStatusPanicked      JobStatus = 8
	JobStatusUnknown       JobStatus = 255
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
