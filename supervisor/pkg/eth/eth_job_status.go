package eth

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
	case JobStatusUnknown:
		return "Unknown"
	}
	return "Unknown"
}
