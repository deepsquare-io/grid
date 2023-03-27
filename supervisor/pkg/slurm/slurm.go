package slurm

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils"
)

type Service struct {
	executor                Executor
	adminUser               string
	scancel                 string
	sbatch                  string
	squeue                  string
	scontrol                string
	supervisorPublicAddress string
}

func New(
	executor Executor,
	adminUser string,
	scancel string,
	sbatch string,
	squeue string,
	scontrol string,
	supervisorPublicAddress string,
) *Service {
	return &Service{
		executor:                executor,
		adminUser:               adminUser,
		scancel:                 scancel,
		sbatch:                  sbatch,
		squeue:                  squeue,
		scontrol:                scontrol,
		supervisorPublicAddress: supervisorPublicAddress,
	}
}

type CancelJobRequest struct {
	// Name of the job
	Name string
	// User is a UNIX User used for impersonation.
	User string
}

// CancelJob kills a job using scancel command.
func (s *Service) CancelJob(ctx context.Context, req *CancelJobRequest) error {
	cmd := fmt.Sprintf("%s --name=%s --me", s.scancel, req.Name)
	_, err := s.executor.ExecAs(ctx, req.User, cmd)
	return err
}

type SubmitJobRequest struct {
	// Name of the job
	Name string
	// User is a UNIX User used for impersonation.
	User string
	*JobDefinition
}

// Submit a sbatch definition script to the SLURM controller using the sbatch command.
func (s *Service) Submit(ctx context.Context, req *SubmitJobRequest) (string, error) {
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
	out, err := s.executor.ExecAs(ctx, req.User, cmd)
	if err != nil {
		return strings.TrimSpace(strings.TrimRight(string(out), "\n")), err
	}

	return strings.TrimSpace(strings.TrimRight(string(out), "\n")), nil
}

type TopUpRequest struct {
	// Name of the job
	Name string
	// User is a UNIX User used for impersonation.
	User string
	// AdditionalTime is the number of minutes to be added
	AdditionalTime uint64
}

// TopUp add additional time to a SLURM job
func (s *Service) TopUp(ctx context.Context, req *TopUpRequest) error {
	// Fetch jobID
	jobID, err := s.FindRunningJobByName(ctx, &FindRunningJobByNameRequest{
		Name: req.Name,
		User: s.adminUser,
	})
	if err != nil {
		return err
	}

	cmd := fmt.Sprintf("%s update job %d TimeLimit+=%d", s.scontrol, jobID, req.AdditionalTime)
	_, err = s.executor.ExecAs(ctx, s.adminUser, cmd)
	return err
}

// HealthCheck runs squeue to check if the queue is running
func (s *Service) HealthCheck(ctx context.Context) error {
	_, err := s.executor.ExecAs(ctx, s.adminUser, s.squeue)
	return err
}

type FindRunningJobByNameRequest struct {
	// Name of the job
	Name string
	// User is a UNIX User used for impersonation. This user should be SLURM admin.
	User string
}

// FindRunningJobByName find a running job using squeue.
func (s *Service) FindRunningJobByName(ctx context.Context, req *FindRunningJobByNameRequest) (int, error) {
	cmd := fmt.Sprintf("%s --name %s -O JobId:256 --noheader", s.squeue, req.Name)
	out, err := s.executor.ExecAs(ctx, req.User, cmd)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(strings.TrimSpace(strings.TrimRight(out, "\n")))
}
