package scheduler

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils"
	"go.uber.org/zap"
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
	out, err := s.ExecAs(ctx, req.User, cmd)
	if err != nil {
		logger.I.Error("CancelJob failed with error", zap.Error(err), zap.String("out", out))
	}
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
	out = strings.TrimSpace(strings.TrimRight(string(out), "\n"))
	if err != nil {
		logger.I.Error("Submit failed with error", zap.Error(err), zap.String("out", out))
		return out, fmt.Errorf(
			"failed to submit: %w, %s",
			err,
			out,
		)
	}

	return out, nil
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
	out, err := s.ExecAs(ctx, s.adminUser, cmd)
	out = strings.TrimSpace(strings.TrimRight(string(out), "\n"))
	if err != nil {
		logger.I.Error("TopUp failed with error", zap.Error(err), zap.String("out", out))
		return fmt.Errorf("failed to top up: %w, %s", err, out)
	}
	return err
}

// HealthCheck runs squeue to check if the queue is running
func (s *Slurm) HealthCheck(ctx context.Context) error {
	out, err := s.ExecAs(ctx, s.adminUser, s.squeue)
	if err != nil {
		logger.I.Error("HealthCheck failed with error", zap.Error(err), zap.String("out", out))
	}
	return err
}

// FindRunningJobByName find a running job using squeue.
func (s *Slurm) FindRunningJobByName(
	ctx context.Context,
	req *FindRunningJobByNameRequest,
) (int, error) {
	cmd := fmt.Sprintf("%s --name %s -O JobId:256 --noheader", s.squeue, req.Name)
	out, err := s.ExecAs(ctx, req.User, cmd)
	out = strings.TrimSpace(strings.TrimRight(string(out), "\n"))
	if err != nil {
		logger.I.Error(
			"FindRunningJobByName failed with error",
			zap.Error(err),
			zap.String("out", out),
		)
		return 0, fmt.Errorf("failed to find job: %w, %s", err, out)
	}

	return strconv.Atoi(out)
}
