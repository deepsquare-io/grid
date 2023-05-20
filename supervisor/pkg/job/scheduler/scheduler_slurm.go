package scheduler

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils"
)

type Slurm struct {
	Executor
	adminUser               string
	scancel                 string
	sbatch                  string
	squeue                  string
	scontrol                string
	supervisorPublicAddress string
}

func NewSlurm(
	executor Executor,
	adminUser string,
	scancel string,
	sbatch string,
	squeue string,
	scontrol string,
	supervisorPublicAddress string,
) *Slurm {
	return &Slurm{
		Executor:                executor,
		adminUser:               adminUser,
		scancel:                 scancel,
		sbatch:                  sbatch,
		squeue:                  squeue,
		scontrol:                scontrol,
		supervisorPublicAddress: supervisorPublicAddress,
	}
}

// CancelJob kills a job using scancel command.
func (s *Slurm) CancelJob(ctx context.Context, req *CancelRequest) error {
	cmd := fmt.Sprintf("%s --name=%s --me", s.scancel, req.Name)
	_, err := s.ExecAs(ctx, req.User, cmd)
	return err
}

// Submit a sbatch definition script to the SLURM controller using the sbatch command.
func (s *Slurm) Submit(ctx context.Context, req *SubmitRequest) (string, error) {
	eof := utils.GenerateRandomString(10)

	cmd := fmt.Sprintf(`%s \
  --parsable \
  --job-name=%s \
	--comment="supervisor %s" \
  --time=%d \
  --ntasks=%d \
  --cpus-per-task=%d \
  --mem-per-cpu=%dM \
  --gpus-per-task=%d \
  --output=/tmp/supervisor-%s-%s.log << '%s'
#!/bin/bash -l
true
%s
%s`,
		s.sbatch,
		req.Name,
		s.supervisorPublicAddress,
		req.TimeLimit,
		req.NTasks,
		req.CPUsPerTask,
		req.MemoryPerCPU,
		req.GPUsPerTask,
		req.Name,
		utils.GenerateRandomString(10),
		eof,
		req.Body,
		eof,
	)
	out, err := s.ExecAs(ctx, req.User, cmd)
	if err != nil {
		return strings.TrimSpace(strings.TrimRight(string(out), "\n")), err
	}

	return strings.TrimSpace(strings.TrimRight(string(out), "\n")), nil
}

// TopUp add additional time to a SLURM job
func (s *Slurm) TopUp(ctx context.Context, req *TopUpRequest) error {
	// Fetch jobID
	jobID, err := s.FindRunningJobByName(ctx, &FindRunningJobByNameRequest{
		Name: req.Name,
		User: s.adminUser,
	})
	if err != nil {
		return err
	}

	cmd := fmt.Sprintf("%s update job %d TimeLimit+=%d", s.scontrol, jobID, req.AdditionalTime)
	_, err = s.ExecAs(ctx, s.adminUser, cmd)
	return err
}

// HealthCheck runs squeue to check if the queue is running
func (s *Slurm) HealthCheck(ctx context.Context) error {
	_, err := s.ExecAs(ctx, s.adminUser, s.squeue)
	return err
}

// FindRunningJobByName find a running job using squeue.
func (s *Slurm) FindRunningJobByName(
	ctx context.Context,
	req *FindRunningJobByNameRequest,
) (int, error) {
	cmd := fmt.Sprintf("%s --name %s -O JobId:256 --noheader", s.squeue, req.Name)
	out, err := s.ExecAs(ctx, req.User, cmd)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(strings.TrimSpace(strings.TrimRight(out, "\n")))
}
