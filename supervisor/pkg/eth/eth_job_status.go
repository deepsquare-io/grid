package eth

type JobStatus uint8

const (
	JobStatusPending JobStatus = iota
	JobStatusMetaScheduled
	JobStatusScheduled
	JobStatusRunning
	JobStatusCancelling
	JobStatusCancelled
	JobStatusFinished
	JobStatusFailed
	JobStatusOutOfCredits
)
