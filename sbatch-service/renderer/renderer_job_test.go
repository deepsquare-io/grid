package renderer_test

import (
	"testing"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/renderer"
	"github.com/deepsquare-io/the-grid/sbatch-service/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var cleanJob = model.Job{
	EnableLogging: utils.Ptr(true),
	Resources: &model.Resources{
		Tasks:       1,
		CpusPerTask: 4,
		MemPerCPU:   4096,
		GpusPerTask: 1,
	},
	Env: []*model.EnvVar{
		{
			Key:   "key",
			Value: "test'test",
		},
	},
	Steps: []*model.Step{
		cleanStepWithRun("echo 'hello world'"),
	},
}

func TestRenderJob(t *testing.T) {
	tests := []struct {
		input         model.Job
		isError       bool
		errorContains []string
		expected      string
		title         string
	}{
		{
			input: cleanJob,
			expected: `#!/bin/bash -l

set -e

export NTASKS='1'
export CPUS_PER_TASK='4'
export MEM_PER_CPU='4096'
export GPUS_PER_TASK='1'
export GPUS='1'
export CPUS='4'
export MEM='16384'
/usr/local/bin/grid-logger-writer \
  --server.tls \
  --server.tls.ca=/etc/ssl/certs/ca-certificates.crt \
  --server.tls.server-host-override=grid-logger.deepsquare.run \
  --server.endpoint=grid-logger.deepsquare.run:443 \
  --pipe.path="/tmp/$SLURM_JOB_NAME-pipe" \
  --log-name="$SLURM_JOB_NAME" \
  --user="$USER" \
  >/dev/stdout 2>/dev/stderr &
LOGGER_PID="$!"
sleep 1
exec &>>"/tmp/$SLURM_JOB_NAME-pipe"
export STORAGE_PATH="/opt/cache/shared/$UID/$SLURM_JOB_NAME"
/usr/bin/mkdir -p "$STORAGE_PATH" "$STORAGE_PATH/output/" "$STORAGE_PATH/input/"
/usr/bin/chmod -R 700 "$STORAGE_PATH"
/usr/bin/chown -R "$UID:cluster-users" "$STORAGE_PATH"
export 'key'='test'\''test'
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro"
/usr/bin/srun --job-name='test' \
  --export=ALL,'STORAGE_PATH=/deepsquare','test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  --container-mounts="${MOUNTS}" \
  --container-image='image' \
  /bin/sh -c 'echo '\''hello world'\'''
kill $LOGGER_PID
wait $LOGGER_PID
`,
			title: "Positive test 'hello world'",
		},
		{
			input: model.Job{
				Env:           cleanJob.Env,
				EnableLogging: cleanJob.EnableLogging,
				Resources:     cleanJob.Resources,
				Input: &model.TransportData{
					S3: &model.S3Data{
						Region:          "us‑east‑2",
						BucketURL:       "s3://test",
						Path:            "/in",
						AccessKeyID:     "AccessKeyID",
						SecretAccessKey: "SecretAccessKey",
						EndpointURL:     "https://s3.us‑east‑2.amazonaws.com",
					},
				},
				Steps: cleanJob.Steps,
				Output: &model.TransportData{
					S3: &model.S3Data{
						Region:          "us‑east‑2",
						BucketURL:       "s3://test",
						Path:            "/out",
						AccessKeyID:     "AccessKeyID",
						SecretAccessKey: "SecretAccessKey",
						EndpointURL:     "https://s3.us‑east‑2.amazonaws.com",
					},
				},
				ContinuousOutputSync: utils.Ptr(true),
			},
			expected: `#!/bin/bash -l

set -e

export NTASKS='1'
export CPUS_PER_TASK='4'
export MEM_PER_CPU='4096'
export GPUS_PER_TASK='1'
export GPUS='1'
export CPUS='4'
export MEM='16384'
/usr/local/bin/grid-logger-writer \
  --server.tls \
  --server.tls.ca=/etc/ssl/certs/ca-certificates.crt \
  --server.tls.server-host-override=grid-logger.deepsquare.run \
  --server.endpoint=grid-logger.deepsquare.run:443 \
  --pipe.path="/tmp/$SLURM_JOB_NAME-pipe" \
  --log-name="$SLURM_JOB_NAME" \
  --user="$USER" \
  >/dev/stdout 2>/dev/stderr &
LOGGER_PID="$!"
sleep 1
exec &>>"/tmp/$SLURM_JOB_NAME-pipe"
export STORAGE_PATH="/opt/cache/shared/$UID/$SLURM_JOB_NAME"
/usr/bin/mkdir -p "$STORAGE_PATH" "$STORAGE_PATH/output/" "$STORAGE_PATH/input/"
/usr/bin/chmod -R 700 "$STORAGE_PATH"
/usr/bin/chown -R "$UID:cluster-users" "$STORAGE_PATH"
export AWS_ACCESS_KEY_ID='AccessKeyID'
export AWS_SECRET_ACCESS_KEY='SecretAccessKey'
export S3_ENDPOINT_URL='https://s3.us‑east‑2.amazonaws.com'

s5cmd sync --source-region 'us‑east‑2' 's3://test''/in' "$STORAGE_PATH/input/"
echo "Input contains:"
find "$STORAGE_PATH/input/"
ContinuousOutputSync() {
  export AWS_ACCESS_KEY_ID='AccessKeyID'
  export AWS_SECRET_ACCESS_KEY='SecretAccessKey'
  export S3_ENDPOINT_URL='https://s3.us‑east‑2.amazonaws.com'
  while true; do
    s5cmd sync --destination-region 'us‑east‑2' "$STORAGE_PATH/outputs/" 's3://test''/out'
    sleep 5
  done
}
ContinuousOutputSync &
CONTINUOUS_SYNC_PID="$!"
export 'key'='test'\''test'
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro"
/usr/bin/srun --job-name='test' \
  --export=ALL,'STORAGE_PATH=/deepsquare','test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  --container-mounts="${MOUNTS}" \
  --container-image='image' \
  /bin/sh -c 'echo '\''hello world'\'''
kill $CONTINUOUS_SYNC_PID
wait $CONTINUOUS_SYNC_PID
echo "Output contains:"
find "$STORAGE_PATH/output/"
export AWS_ACCESS_KEY_ID='AccessKeyID'
export AWS_SECRET_ACCESS_KEY='SecretAccessKey'
export S3_ENDPOINT_URL='https://s3.us‑east‑2.amazonaws.com'

s5cmd sync --destination-region 'us‑east‑2' "$STORAGE_PATH/outputs/" 's3://test''/out'
kill $LOGGER_PID
wait $LOGGER_PID
`,
			title: "Positive test with S3 input output",
		},
		{
			input: model.Job{
				Env:           cleanJob.Env,
				EnableLogging: cleanJob.EnableLogging,
				Resources:     cleanJob.Resources,
				Input: &model.TransportData{
					S3: &model.S3Data{
						Region:          "us‑east‑2",
						BucketURL:       "s3://test",
						Path:            "/in",
						AccessKeyID:     "AccessKeyID",
						SecretAccessKey: "SecretAccessKey",
						EndpointURL:     "https://s3.us‑east‑2.amazonaws.com",
					},
				},
				Steps: cleanJob.Steps,
				Output: &model.TransportData{
					S3: &model.S3Data{
						Region:          "us‑east‑2",
						BucketURL:       "s3://test",
						Path:            "/out",
						AccessKeyID:     "AccessKeyID",
						SecretAccessKey: "SecretAccessKey",
						EndpointURL:     "https://s3.us‑east‑2.amazonaws.com",
					},
				},
			},
			expected: `#!/bin/bash -l

set -e

export NTASKS='1'
export CPUS_PER_TASK='4'
export MEM_PER_CPU='4096'
export GPUS_PER_TASK='1'
export GPUS='1'
export CPUS='4'
export MEM='16384'
/usr/local/bin/grid-logger-writer \
  --server.tls \
  --server.tls.ca=/etc/ssl/certs/ca-certificates.crt \
  --server.tls.server-host-override=grid-logger.deepsquare.run \
  --server.endpoint=grid-logger.deepsquare.run:443 \
  --pipe.path="/tmp/$SLURM_JOB_NAME-pipe" \
  --log-name="$SLURM_JOB_NAME" \
  --user="$USER" \
  >/dev/stdout 2>/dev/stderr &
LOGGER_PID="$!"
sleep 1
exec &>>"/tmp/$SLURM_JOB_NAME-pipe"
export STORAGE_PATH="/opt/cache/shared/$UID/$SLURM_JOB_NAME"
/usr/bin/mkdir -p "$STORAGE_PATH" "$STORAGE_PATH/output/" "$STORAGE_PATH/input/"
/usr/bin/chmod -R 700 "$STORAGE_PATH"
/usr/bin/chown -R "$UID:cluster-users" "$STORAGE_PATH"
export AWS_ACCESS_KEY_ID='AccessKeyID'
export AWS_SECRET_ACCESS_KEY='SecretAccessKey'
export S3_ENDPOINT_URL='https://s3.us‑east‑2.amazonaws.com'

s5cmd sync --source-region 'us‑east‑2' 's3://test''/in' "$STORAGE_PATH/input/"
echo "Input contains:"
find "$STORAGE_PATH/input/"
export 'key'='test'\''test'
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro"
/usr/bin/srun --job-name='test' \
  --export=ALL,'STORAGE_PATH=/deepsquare','test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  --container-mounts="${MOUNTS}" \
  --container-image='image' \
  /bin/sh -c 'echo '\''hello world'\'''
echo "Output contains:"
find "$STORAGE_PATH/output/"
export AWS_ACCESS_KEY_ID='AccessKeyID'
export AWS_SECRET_ACCESS_KEY='SecretAccessKey'
export S3_ENDPOINT_URL='https://s3.us‑east‑2.amazonaws.com'

s5cmd sync --destination-region 'us‑east‑2' "$STORAGE_PATH/outputs/" 's3://test''/out'
kill $LOGGER_PID
wait $LOGGER_PID
`,
			title: "Positive test with S3 input output and continuous sync",
		},
		{
			input: model.Job{
				Env:           cleanJob.Env,
				EnableLogging: cleanJob.EnableLogging,
				Resources:     cleanJob.Resources,
				Input: &model.TransportData{
					HTTP: &model.HTTPData{
						URL: "https://test/in",
					},
				},
				Steps: cleanJob.Steps,
				Output: &model.TransportData{
					HTTP: &model.HTTPData{
						URL: "https://test/out",
					},
				},
			},
			expected: `#!/bin/bash -l

set -e

export NTASKS='1'
export CPUS_PER_TASK='4'
export MEM_PER_CPU='4096'
export GPUS_PER_TASK='1'
export GPUS='1'
export CPUS='4'
export MEM='16384'
/usr/local/bin/grid-logger-writer \
  --server.tls \
  --server.tls.ca=/etc/ssl/certs/ca-certificates.crt \
  --server.tls.server-host-override=grid-logger.deepsquare.run \
  --server.endpoint=grid-logger.deepsquare.run:443 \
  --pipe.path="/tmp/$SLURM_JOB_NAME-pipe" \
  --log-name="$SLURM_JOB_NAME" \
  --user="$USER" \
  >/dev/stdout 2>/dev/stderr &
LOGGER_PID="$!"
sleep 1
exec &>>"/tmp/$SLURM_JOB_NAME-pipe"
export STORAGE_PATH="/opt/cache/shared/$UID/$SLURM_JOB_NAME"
/usr/bin/mkdir -p "$STORAGE_PATH" "$STORAGE_PATH/output/" "$STORAGE_PATH/input/"
/usr/bin/chmod -R 700 "$STORAGE_PATH"
/usr/bin/chown -R "$UID:cluster-users" "$STORAGE_PATH"
curl -JORSL 'https://test/in' -o "$STORAGE_PATH/input/"
for filepath in "$STORAGE_PATH/input/"*; do
  tar -xvaf "$filepath" 2>/dev/null && continue
  case $(file "$filepath") in
      *bzip2*) bzip2 -dk "$filepath";;
      *gzip*) gunzip "$filepath";;
      *zip*) ;&
      *Zip*) unzip "$filepath";;
      *xz*) ;&
      *XZ*) unxz  "$filepath";;
      *'7-zip'*) 7z x "$filepath";;
      *) 1>&2 echo "Unknown archive '$filepath'";;
  esac
done
echo "Input contains:"
find "$STORAGE_PATH/input/"
export 'key'='test'\''test'
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro"
/usr/bin/srun --job-name='test' \
  --export=ALL,'STORAGE_PATH=/deepsquare','test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  --container-mounts="${MOUNTS}" \
  --container-image='image' \
  /bin/sh -c 'echo '\''hello world'\'''
echo "Output contains:"
find "$STORAGE_PATH/output/"
tar -cvf "$STORAGE_PATH/output.tar" "$STORAGE_PATH/output/"
curl --upload-file "$STORAGE_PATH/output.tar" 'https://test/out'
kill $LOGGER_PID
wait $LOGGER_PID
`,
			title: "Positive test with HTTP input output",
		},
		{
			input: model.Job{
				Env:           cleanJob.Env,
				Resources:     cleanJob.Resources,
				EnableLogging: utils.Ptr(false),
				Steps:         cleanJob.Steps,
			},
			expected: `#!/bin/bash -l

set -e

export NTASKS='1'
export CPUS_PER_TASK='4'
export MEM_PER_CPU='4096'
export GPUS_PER_TASK='1'
export GPUS='1'
export CPUS='4'
export MEM='16384'
export STORAGE_PATH="/opt/cache/shared/$UID/$SLURM_JOB_NAME"
/usr/bin/mkdir -p "$STORAGE_PATH" "$STORAGE_PATH/output/" "$STORAGE_PATH/input/"
/usr/bin/chmod -R 700 "$STORAGE_PATH"
/usr/bin/chown -R "$UID:cluster-users" "$STORAGE_PATH"
export 'key'='test'\''test'
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro"
/usr/bin/srun --job-name='test' \
  --export=ALL,'STORAGE_PATH=/deepsquare','test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  --container-mounts="${MOUNTS}" \
  --container-image='image' \
  /bin/sh -c 'echo '\''hello world'\'''
`,
			title: "Positive test with no logs",
		},
		{
			input: model.Job{
				Env: []*model.EnvVar{
					{
						Key:   "aze'aze",
						Value: "test'test",
					},
				},
			},
			isError:       true,
			errorContains: []string{"valid_envvar_name", "Key"},
			title:         "Negative test invalid env var name",
		},
		{
			input: model.Job{
				Env: []*model.EnvVar{
					{
						Key:   "PATH",
						Value: "test'test",
					},
				},
			},
			isError:       true,
			errorContains: []string{"ne", "Key"},
			title:         "Negative test invalid env var Key name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			actual, err := renderer.RenderJob(&tt.input)

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, actual)
				shellcheck(t, actual)
			}
		})
	}
}
