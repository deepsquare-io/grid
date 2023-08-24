package scheduler

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils"
	"go.uber.org/zap"
)

const (
	defaultSInfo     = "sinfo"
	defaultSCancel   = "scancel"
	defaultSBatch    = "sbatch"
	defaultSQueue    = "squeue"
	defaultSControl  = "scontrol"
	defaultNVidiaSMI = "nvidia-smi"
)

type SlurmOption func(*Slurm)

func WithSInfo(path string) SlurmOption {
	return func(s *Slurm) {
		s.sinfo = path
	}
}

func WithSCancel(path string) SlurmOption {
	return func(s *Slurm) {
		s.scancel = path
	}
}

func WithSQueue(path string) SlurmOption {
	return func(s *Slurm) {
		s.squeue = path
	}
}

func WithSBatch(path string) SlurmOption {
	return func(s *Slurm) {
		s.sbatch = path
	}
}

func WithSControl(path string) SlurmOption {
	return func(s *Slurm) {
		s.scontrol = path
	}
}

func WithNVidiaSMI(path string) SlurmOption {
	return func(s *Slurm) {
		s.nvidiaSMI = path
	}
}

type Slurm struct {
	Executor
	adminUser               string
	sinfo                   string
	scancel                 string
	sbatch                  string
	squeue                  string
	scontrol                string
	supervisorPublicAddress string
	nvidiaSMI               string
	partition               string
}

func NewSlurm(
	executor Executor,
	adminUser string,
	supervisorPublicAddress string,
	partition string,
	opts ...SlurmOption,
) Scheduler {
	s := &Slurm{
		Executor:                executor,
		adminUser:               adminUser,
		sinfo:                   defaultSInfo,
		scancel:                 defaultSCancel,
		sbatch:                  defaultSBatch,
		squeue:                  defaultSQueue,
		scontrol:                defaultSControl,
		supervisorPublicAddress: supervisorPublicAddress,
		nvidiaSMI:               defaultNVidiaSMI,
		partition:               partition,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// CancelJob kills a job using scancel command.
func (s *Slurm) CancelJob(ctx context.Context, name string, user string) error {
	cmd := fmt.Sprintf("%s --name=%s --me", s.scancel, name)
	out, err := s.ExecAs(ctx, user, cmd)
	if err != nil {
		logger.I.Error("CancelJob failed with error", zap.Error(err), zap.String("out", out))
	}
	return err
}

// Submit a sbatch definition script to the SLURM controller using the sbatch command.
func (s *Slurm) Submit(ctx context.Context, req *SubmitRequest) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	eof := utils.GenerateRandomString(10)

	cmd := fmt.Sprintf(`%s \
  --parsable \
  --job-name=%s \
  --comment="%s %s" \
  --partition="%s" \
  --output=/tmp/%s-%s-%s.log`,
		s.sbatch,
		req.Name,
		req.Prefix,
		s.supervisorPublicAddress,
		s.partition,
		req.Prefix,
		req.Name,
		utils.GenerateRandomString(10),
	)

	if req.NTasks > 0 {
		cmd = fmt.Sprintf(`%s \
  --ntasks=%d`,
			cmd,
			req.NTasks,
		)
	}

	if req.NTasksPerNode > 0 {
		cmd = fmt.Sprintf(`%s \
  --ntasks-per-node=%d`,
			cmd,
			req.NTasksPerNode,
		)
	}

	if req.CPUsPerTask > 0 {
		cmd = fmt.Sprintf(`%s \
  --cpus-per-task=%d`,
			cmd,
			req.CPUsPerTask,
		)
	}

	if req.CPUsPerNode > 0 {
		cmd = fmt.Sprintf(`%s \
  --mincpus=%d`,
			cmd,
			req.CPUsPerNode,
		)
	}

	if req.MinNodes > 0 {
		cmd = fmt.Sprintf(`%s \
  --nodes=%d`,
			cmd,
			req.MinNodes,
		)
		if req.MaxNodes > 0 {
			cmd = fmt.Sprintf(`%s-%d`, cmd, req.MaxNodes)
		}
	}

	if req.Memory != nil {
		cmd = fmt.Sprintf(`%s \
  --mem=%d`,
			cmd,
			*req.Memory,
		)
	}

	if req.TimeLimit > 0 {
		cmd = fmt.Sprintf(`%s \
  --time=%d`,
			cmd,
			req.TimeLimit,
		)
	}

	if req.MemoryPerCPU > 0 {
		cmd = fmt.Sprintf(`%s \
  --mem-per-cpu=%d`,
			cmd,
			req.MemoryPerCPU,
		)
	}

	if req.GPUsPerTask != nil {
		cmd = fmt.Sprintf(`%s \
  --gpus-per-task=%d`,
			cmd,
			*req.GPUsPerTask,
		)
	}

	if req.GPUsPerNode > 0 {
		cmd = fmt.Sprintf(`%s \
  --gpus-per-node=%d`,
			cmd,
			req.GPUsPerNode,
		)
	}
	if req.Wait {
		cmd = fmt.Sprintf(`%s \
  --wait`,
			cmd,
		)
	}
	cmd = fmt.Sprintf(`%s \
<< '%s'
#!/bin/bash -l
true
%s
%s`,
		cmd,
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

func (s *Slurm) TopUp(ctx context.Context, name string, additionalTime uint64) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	// Fetch jobID
	jobID, err := s.FindRunningJobByName(ctx, name, s.adminUser)
	if err != nil {
		return err
	}

	cmd := fmt.Sprintf("%s update job %d TimeLimit+=%d", s.scontrol, jobID, additionalTime)
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
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	out, err := s.ExecAs(ctx, s.adminUser, fmt.Sprintf("timeout 10 %s", s.squeue))
	if err != nil {
		logger.I.Error("HealthCheck failed with error", zap.Error(err), zap.String("out", out))
	}
	return err
}

// FindRunningJobByName find a running job using squeue.
func (s *Slurm) FindRunningJobByName(
	ctx context.Context,
	name string,
	user string,
) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	cmd := fmt.Sprintf("%s --name %s -O JobId:256 --noheader", s.squeue, name)
	out, err := s.ExecAs(ctx, user, cmd)
	out = strings.TrimSpace(strings.TrimRight(string(out), "\n"))
	if err != nil {
		logger.I.Error(
			"FindRunningJobByName failed with error",
			zap.Error(err),
			zap.String("out", out),
		)
		return 0, fmt.Errorf("failed to find job: %w, %s", err, out)
	}

	if out == "" {
		return 0, nil
	}

	return strconv.Atoi(out)
}

func (s *Slurm) FindMemPerNode(ctx context.Context, opts ...FindSpecOption) ([]uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	o := newFindSpecOptions(opts...)
	cmd := fmt.Sprintf(
		`%s show nodes --oneliner | grep 'Partitions=[^ ]*%s'`,
		s.scontrol,
		s.partition,
	)
	if o.onlyResponding {
		cmd = fmt.Sprintf(`%s | grep -v NOT_RESPONDING`, cmd)
	}
	cmd = fmt.Sprintf(
		`%s | sed -E 's/.*CfgTRES=[^ ]*mem=([0-9]+)[^0-9].*/\1/'`,
		cmd,
	)
	out, err := s.ExecAs(ctx, s.adminUser, cmd)
	if err != nil {
		logger.I.Error(
			"FindMemPerNode failed with error",
			zap.Error(err),
			zap.String("out", out),
		)
		return []uint64{}, err
	}

	lines := strings.Split(strings.TrimSpace(out), "\n")
	memPerN := make([]uint64, len(lines))
	for i, line := range lines {
		mem, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			logger.I.Error(
				"failed to convert string to int",
				zap.Error(err),
				zap.String("in", strings.TrimSpace(line)),
			)
		}
		memPerN[i] = mem
	}

	return memPerN, nil
}

func (s *Slurm) FindGPUsPerNode(ctx context.Context, opts ...FindSpecOption) ([]uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	o := newFindSpecOptions(opts...)
	cmd := fmt.Sprintf(
		`%s show nodes --oneliner | grep 'Partitions=[^ ]*%s'`,
		s.scontrol,
		s.partition,
	)
	if o.onlyResponding {
		cmd = fmt.Sprintf(`%s | grep -v NOT_RESPONDING`, cmd)
	}
	cmd = fmt.Sprintf(
		`%s | sed -E 's|.*CfgTRES=[^ ]*gres/gpu=([0-9]+)[^0-9].*|\1|g'`,
		cmd,
	)
	out, err := s.ExecAs(ctx, s.adminUser, cmd)
	if err != nil {
		logger.I.Error(
			"FindGPUPerNode failed with error",
			zap.Error(err),
			zap.String("out", out),
		)
		return []uint64{}, err
	}

	lines := strings.Split(strings.TrimSpace(out), "\n")
	gpusPerN := make([]uint64, len(lines))
	for i, line := range lines {
		gpus, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			logger.I.Error(
				"failed to convert string to int",
				zap.Error(err),
				zap.String("in", strings.TrimSpace(line)),
			)
		}
		gpusPerN[i] = gpus
	}

	return gpusPerN, nil
}

func (s *Slurm) FindCPUsPerNode(ctx context.Context, opts ...FindSpecOption) ([]uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	o := newFindSpecOptions(opts...)
	cmd := fmt.Sprintf(
		`%s show nodes --oneliner | grep 'Partitions=[^ ]*%s'`,
		s.scontrol,
		s.partition,
	)
	if o.onlyResponding {
		cmd = fmt.Sprintf(`%s | grep -v NOT_RESPONDING`, cmd)
	}
	cmd = fmt.Sprintf(
		`%s | sed -E 's|.*CfgTRES=[^ ]*cpu=([0-9]+)[^0-9].*|\1|g'`,
		cmd,
	)
	out, err := s.ExecAs(ctx, s.adminUser, cmd)
	if err != nil {
		logger.I.Error(
			"FindCPUPerNode failed with error",
			zap.Error(err),
			zap.String("out", out),
		)
		return []uint64{}, err
	}

	lines := strings.Split(strings.TrimSpace(out), "\n")
	cpusPerN := make([]uint64, len(lines))
	for i, line := range lines {
		cpus, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			logger.I.Error(
				"failed to convert string to int",
				zap.Error(err),
				zap.String("in", strings.TrimSpace(line)),
			)
		}
		cpusPerN[i] = cpus
	}

	return cpusPerN, nil
}

func (s *Slurm) FindTotalCPUs(ctx context.Context) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	cmd := fmt.Sprintf(
		`%s show partition '%s' --oneliner | sed -E 's|.*TRES=[^ ]*cpu=([0-9]+)[^0-9].*|\1|g'`,
		s.scontrol,
		s.partition,
	)
	out, err := s.ExecAs(ctx, s.adminUser, cmd)
	if err != nil {
		logger.I.Error(
			"FindTotalCPUs failed with error",
			zap.Error(err),
			zap.String("out", out),
		)
		return 0, err
	}

	cpus, err := strconv.ParseUint(strings.TrimSpace(out), 10, 64)
	if err != nil {
		logger.I.Error(
			"FindTotalCPUs failed with error",
			zap.Error(err),
			zap.String("out", out),
		)
		return 0, err
	}

	return cpus, nil
}

func (s *Slurm) FindTotalMem(ctx context.Context) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	cmd := fmt.Sprintf(
		`%s show partition '%s' --oneliner | sed -E 's|.*TRES=[^ ]*mem=([0-9]+)[^0-9].*|\1|g'`,
		s.scontrol,
		s.partition,
	)
	out, err := s.ExecAs(ctx, s.adminUser, cmd)
	if err != nil {
		logger.I.Error(
			"FindTotalMem failed with error",
			zap.Error(err),
			zap.String("out", out),
		)
		return 0, err
	}

	mem, err := strconv.ParseUint(strings.TrimSpace(out), 10, 64)
	if err != nil {
		logger.I.Error(
			"FindTotalMem failed with error",
			zap.Error(err),
			zap.String("out", out),
		)
		return 0, err
	}

	return mem, nil
}

func (s *Slurm) FindTotalGPUs(ctx context.Context) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	cmd := fmt.Sprintf(
		`%s show partition '%s' --oneliner | sed -E 's|.*TRES=[^ ]*gpu=([0-9]+)[^0-9].*|\1|g'`,
		s.scontrol,
		s.partition,
	)
	out, err := s.ExecAs(ctx, s.adminUser, cmd)
	if err != nil {
		logger.I.Error(
			"FindTotalGPUs failed with error",
			zap.Error(err),
			zap.String("out", out),
		)
		return 0, err
	}

	gpu, err := strconv.ParseUint(strings.TrimSpace(out), 10, 64)
	if err != nil {
		logger.I.Error(
			"FindTotalGPUs failed with error",
			zap.Error(err),
			zap.String("out", out),
		)
		return 0, err
	}

	return gpu, nil
}

func (s *Slurm) FindTotalNodes(ctx context.Context, opts ...FindSpecOption) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	o := newFindSpecOptions(opts...)
	var cmd string
	if o.onlyResponding {
		cmd = fmt.Sprintf(
			`%s --partition='%s' --responding --Format=Nodes:10 --noheader`,
			s.sinfo,
			s.partition,
		)
	} else {
		cmd = fmt.Sprintf(
			`%s show partition '%s' --oneliner | sed -E 's|.*TRES=[^ ]*node=([0-9]+)[^0-9].*|\1|g'`,
			s.scontrol,
			s.partition,
		)
	}
	out, err := s.ExecAs(ctx, s.adminUser, cmd)
	if err != nil {
		logger.I.Error(
			"FindTotalNodes failed with error",
			zap.Error(err),
			zap.String("out", out),
		)
		return 0, err
	}

	ret, err := strconv.ParseUint(strings.TrimSpace(out), 10, 64)
	if err != nil {
		logger.I.Error(
			"FindTotalNodes failed with error",
			zap.Error(err),
			zap.String("out", out),
		)
		return 0, err
	}

	return ret, nil
}
