package slurm

import (
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

type JobDefinition struct {
	// TimeLimit is a time allocation which at the end kills the running job.
	//
	// TimeLimit is in minutes.
	TimeLimit uint64
	// NTasks indicates the number
	NTasks int64
	// GPUs indicates the number of requested GPU.
	GPUs int64
	// CPUs indicates the number of requested CPU.
	CPUsPerTask int64
	// Memory indicates the number of requested MB of memory.
	Memory int64
	// Body of the job, in a sbatch script.
	Body string
}

type CancelJobRequest struct {
	// Name of the job
	Name string
	// User is a UNIX User used for impersonation.
	User string
}

type SubmitJobRequest struct {
	// Name of the job
	Name string
	// User is a UNIX User used for impersonation.
	User string
	JobDefinition
}

type TopUpRequest struct {
	// Name of the job
	Name string
	// User is a UNIX User used for impersonation.
	User string
	// AdditionalTime is the number of minutes to be added
	AdditionalTime uint64
}

type FindRunningJobByNameRequest struct {
	// Name of the job
	Name string
	// User is a UNIX User used for impersonation. This user should be SLURM admin.
	User string
}

type JobService interface {
	// CancelJob kils a job using scancel command.
	CancelJob(req *CancelJobRequest) error
	// SubmitJob submits sbatch definition script to the SLURM controller using the sbatch command.
	SubmitJob(req *SubmitJobRequest) (int, error)
	// TopUp add additional time to a SLURM job
	TopUp(req *TopUpRequest) error
	// FindRunningJobByName find a running job using squeue.
	FindRunningJobByName(req *FindRunningJobByNameRequest) (int, error)
}

type jobService struct {
	address    string
	authMethod ssh.AuthMethod
	adminUser  string
	scancel    string
	sbatch     string
	squeue     string
	scontrol   string
}

func New(
	address string,
	pkB64 string,
	adminUser string,
	scancel string,
	sbatch string,
	squeue string,
	scontrol string,
) *jobService {
	pk, err := base64.StdEncoding.DecodeString(pkB64)
	if err != nil {
		logger.I.Panic("failed to decode key", zap.Error(err))
	}

	signer, err := ssh.ParsePrivateKey(pk)
	if err != nil {
		logger.I.Panic("couldn't parse private key", zap.Error(err))
	}

	return &jobService{
		address:    address,
		authMethod: ssh.PublicKeys(signer),
		adminUser:  adminUser,
		scancel:    scancel,
		sbatch:     sbatch,
		squeue:     squeue,
		scontrol:   scontrol,
	}
}

func (s *jobService) establish(user string) (session *ssh.Session, close func(), err error) {
	config := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
		Auth:            []ssh.AuthMethod{s.authMethod},
	}
	client, err := ssh.Dial("tcp", s.address, config)
	if err != nil {
		return nil, nil, err
	}
	session, err = client.NewSession()
	if err != nil {
		if err := client.Close(); err != nil {
			logger.I.Warn("closing SSH client thrown an error", zap.Error(err))
		}
		return nil, nil, err
	}

	return session, func() {
		if err := session.Close(); err != nil && err != io.EOF {
			logger.I.Warn("closing SSH session thrown an error", zap.Error(err))
		}
		if err := client.Close(); err != nil && err != io.EOF {
			logger.I.Warn("closing SSH client thrown an error", zap.Error(err))
		}
	}, nil
}

func (s *jobService) CancelJob(req *CancelJobRequest) error {
	sess, close, err := s.establish(req.User)
	if err != nil {
		return err
	}
	defer close()

	cmd := fmt.Sprintf("%s --name=%s --me", s.scancel, req.Name)
	out, err := sess.CombinedOutput(cmd)
	if err != nil {
		logger.I.Error(
			"scancel failed",
			zap.Error(err),
			zap.Any("params", req),
			zap.String("output", string(out)),
		)
	}

	return nil
}

func (s *jobService) SubmitJob(req *SubmitJobRequest) (int, error) {
	sess, close, err := s.establish(req.User)
	if err != nil {
		return 0, err
	}
	defer close()

	eof := utils.GenerateRandomString(10)

	cmd := fmt.Sprintf(`%s \
	--parsable \
  --job-name=%s \
	--time=%d \
  --ntasks=%d \
  --cpus-per-task=%d \
  --mem=%dM \
  --gpus-per-node=%d << %s
%s

%s`,
		s.sbatch,
		req.Name,
		req.TimeLimit,
		req.NTasks,
		req.CPUsPerTask,
		req.Memory,
		req.GPUs,
		eof,
		req.Body,
		eof,
	)
	out, err := sess.CombinedOutput(cmd)
	if err != nil {
		logger.I.Error(
			"sbatch command failed",
			zap.Error(err),
			zap.Any("params", req),
			zap.String("output", string(out)),
		)
		return 0, err
	}

	return strconv.Atoi(strings.TrimSpace(strings.TrimRight(string(out), "\n")))
}

func (s *jobService) TopUp(req *TopUpRequest) error {
	// Fetch jobID
	jobID, err := s.FindRunningJobByName(&FindRunningJobByNameRequest{
		Name: req.Name,
		User: req.User,
	})
	if err != nil {
		return err
	}

	sess, close, err := s.establish(s.adminUser)
	if err != nil {
		return err
	}
	defer close()

	cmd := fmt.Sprintf("%s update job %d TimeLimit+=%d", s.scontrol, jobID, req.AdditionalTime)
	out, err := sess.CombinedOutput(cmd)
	if err != nil {
		logger.I.Error(
			"top up command failed",
			zap.Error(err),
			zap.Any("params", req),
			zap.String("output", string(out)),
		)
	}

	return err
}

func (s *jobService) FindRunningJobByName(req *FindRunningJobByNameRequest) (int, error) {
	sess, close, err := s.establish(req.User)
	if err != nil {
		return 0, err
	}
	defer close()

	cmd := fmt.Sprintf("%s --name %s -O JobId:1256 --noheader", s.squeue, req.Name)
	out, err := sess.CombinedOutput(cmd)
	if err != nil {
		logger.I.Error(
			"squeue command failed",
			zap.Error(err),
			zap.Any("params", req),
			zap.String("output", string(out)),
		)
		return 0, err
	}

	return strconv.Atoi(strings.TrimSpace(strings.TrimRight(string(out), "\n")))
}
