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

type Service struct {
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
) *Service {
	pk, err := base64.StdEncoding.DecodeString(pkB64)
	if err != nil {
		logger.I.Panic("failed to decode key", zap.Error(err))
	}

	signer, err := ssh.ParsePrivateKey(pk)
	if err != nil {
		logger.I.Panic("couldn't parse private key", zap.Error(err))
	}

	return &Service{
		address:    address,
		authMethod: ssh.PublicKeys(signer),
		adminUser:  adminUser,
		scancel:    scancel,
		sbatch:     sbatch,
		squeue:     squeue,
		scontrol:   scontrol,
	}
}

func (s *Service) establish(user string) (session *ssh.Session, close func(), err error) {
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
		if err := client.Close(); err != nil && err != io.EOF {
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

type CancelJobRequest struct {
	// Name of the job
	Name string
	// User is a UNIX User used for impersonation.
	User string
}

// CancelJob kils a job using scancel command.
func (s *Service) CancelJob(req *CancelJobRequest) error {
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

type SubmitJobRequest struct {
	// Name of the job
	Name string
	// User is a UNIX User used for impersonation.
	User string
	*JobDefinition
}

// SubmitJob submits sbatch definition script to the SLURM controller using the sbatch command.
func (s *Service) SubmitJob(req *SubmitJobRequest) (int, error) {
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
		req.MemoryPerNode,
		req.GPUsPerNode,
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

type TopUpRequest struct {
	// Name of the job
	Name string
	// User is a UNIX User used for impersonation.
	User string
	// AdditionalTime is the number of minutes to be added
	AdditionalTime uint64
}

// TopUp add additional time to a SLURM job
func (s *Service) TopUp(req *TopUpRequest) error {
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

type FindRunningJobByNameRequest struct {
	// Name of the job
	Name string
	// User is a UNIX User used for impersonation. This user should be SLURM admin.
	User string
}

// FindRunningJobByName find a running job using squeue.
func (s *Service) FindRunningJobByName(req *FindRunningJobByNameRequest) (int, error) {
	sess, close, err := s.establish(req.User)
	if err != nil {
		return 0, err
	}
	defer close()

	cmd := fmt.Sprintf("%s --name %s -O JobId:256 --noheader", s.squeue, req.Name)
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
