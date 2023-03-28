package scheduler

import "context"

type Executor interface {
	ExecAs(ctx context.Context, user string, cmd string) (string, error)
}
