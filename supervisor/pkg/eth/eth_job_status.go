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
