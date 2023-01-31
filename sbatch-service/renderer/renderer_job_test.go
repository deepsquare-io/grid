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
	InputMode:     utils.Ptr(493),
	Resources: &model.JobResources{
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
		{
			Name: "test",
			Run:  cleanStepRun("/usr/bin/echo 'hello world'"),
		},
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
/usr/bin/sleep 1
exec 3>&1
exec 4>&2
exec 1>>"/tmp/$SLURM_JOB_NAME-pipe"
exec 2>&1

disposeLogs() {
  echo cleaning up
  /usr/bin/sleep 15
  exec 1>&3
  exec 2>&4
  echo killing logger
  kill $LOGGER_PID || true
  wait $LOGGER_PID || true
  echo cleaned
}
trap disposeLogs EXIT
(
export STORAGE_PATH="/opt/cache/shared/$UID/$SLURM_JOB_NAME"
export DEEPSQUARE_INPUT="$STORAGE_PATH/input"
export DEEPSQUARE_OUTPUT="$STORAGE_PATH/output"
export DEEPSQUARE_ENV="$STORAGE_PATH/env"
/usr/bin/mkdir -p "$STORAGE_PATH" "$DEEPSQUARE_OUTPUT" "$DEEPSQUARE_INPUT"
/usr/bin/touch "$DEEPSQUARE_ENV"
/usr/bin/chmod -R 700 "$STORAGE_PATH"
/usr/bin/chown -R "$UID:cluster-users" "$STORAGE_PATH"
loadDeepsquareEnv() {
  /usr/bin/grep -v '^#' "$DEEPSQUARE_ENV" | /usr/bin/grep '=' | /usr/bin/sed -Ez '$ s/\n+$//' | tr '\n' ','
}
/usr/bin/chmod -R 755 "$DEEPSQUARE_INPUT/"
export 'key'='test'\''test'
/usr/bin/echo 'Running: ''test'
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro"
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' /usr/bin/srun --job-name='test' \
  --export=ALL,"$(loadDeepsquareEnv)",'test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  --container-mounts="${MOUNTS}" \
  --container-image='image' \
  /bin/sh -c '/usr/bin/echo '\''hello world'\'''
)
`,
			title: "Positive test 'hello world'",
		},
		{
			input: model.Job{
				Env:           cleanJob.Env,
				EnableLogging: cleanJob.EnableLogging,
				Resources:     cleanJob.Resources,
				InputMode:     cleanJob.InputMode,
				Input: &model.TransportData{
					S3: &model.S3Data{
						Region:          "us‑east‑2",
						BucketURL:       "s3://test",
						Path:            "/in",
						AccessKeyID:     "AccessKeyID",
						SecretAccessKey: "SecretAccessKey",
						EndpointURL:     "https://s3.us‑east‑2.amazonaws.com",
						DeleteSync:      utils.Ptr(true),
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
						DeleteSync:      utils.Ptr(true),
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
/usr/bin/sleep 1
exec 3>&1
exec 4>&2
exec 1>>"/tmp/$SLURM_JOB_NAME-pipe"
exec 2>&1

disposeLogs() {
  echo cleaning up
  /usr/bin/sleep 15
  exec 1>&3
  exec 2>&4
  echo killing logger
  kill $LOGGER_PID || true
  wait $LOGGER_PID || true
  echo cleaned
}
trap disposeLogs EXIT
(
export STORAGE_PATH="/opt/cache/shared/$UID/$SLURM_JOB_NAME"
export DEEPSQUARE_INPUT="$STORAGE_PATH/input"
export DEEPSQUARE_OUTPUT="$STORAGE_PATH/output"
export DEEPSQUARE_ENV="$STORAGE_PATH/env"
/usr/bin/mkdir -p "$STORAGE_PATH" "$DEEPSQUARE_OUTPUT" "$DEEPSQUARE_INPUT"
/usr/bin/touch "$DEEPSQUARE_ENV"
/usr/bin/chmod -R 700 "$STORAGE_PATH"
/usr/bin/chown -R "$UID:cluster-users" "$STORAGE_PATH"
loadDeepsquareEnv() {
  /usr/bin/grep -v '^#' "$DEEPSQUARE_ENV" | /usr/bin/grep '=' | /usr/bin/sed -Ez '$ s/\n+$//' | tr '\n' ','
}
export AWS_ACCESS_KEY_ID='AccessKeyID'
export AWS_SECRET_ACCESS_KEY='SecretAccessKey'
export S3_ENDPOINT_URL='https://s3.us‑east‑2.amazonaws.com'

s5cmd cp --source-region 'us‑east‑2' 's3://test''/in''*' "$DEEPSQUARE_INPUT/"
/usr/bin/chmod -R 700 "$DEEPSQUARE_INPUT/" || echo "chmod failed, but we are ignoring it"
/usr/bin/echo "Input contains:"
/usr/bin/find "$DEEPSQUARE_INPUT/" -exec realpath --relative-to "$DEEPSQUARE_INPUT/" {} \;
/usr/bin/chmod -R 755 "$DEEPSQUARE_INPUT/"
ContinuousOutputSync() {
  export AWS_ACCESS_KEY_ID='AccessKeyID'
  export AWS_SECRET_ACCESS_KEY='SecretAccessKey'
  export S3_ENDPOINT_URL='https://s3.us‑east‑2.amazonaws.com'
  while true; do
    s5cmd sync --delete --destination-region 'us‑east‑2' "$DEEPSQUARE_OUTPUT/" 's3://test''/out'
    /usr/bin/sleep 5
  done
}
ContinuousOutputSync &
CONTINUOUS_SYNC_PID="$!"
(
export 'key'='test'\''test'
/usr/bin/echo 'Running: ''test'
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro"
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' /usr/bin/srun --job-name='test' \
  --export=ALL,"$(loadDeepsquareEnv)",'test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  --container-mounts="${MOUNTS}" \
  --container-image='image' \
  /bin/sh -c '/usr/bin/echo '\''hello world'\'''
)
kill $CONTINUOUS_SYNC_PID || true
wait $CONTINUOUS_SYNC_PID || true
/usr/bin/echo "Output contains:"
/usr/bin/find "$DEEPSQUARE_OUTPUT/" -exec realpath --relative-to "$DEEPSQUARE_OUTPUT/" {} \;
export AWS_ACCESS_KEY_ID='AccessKeyID'
export AWS_SECRET_ACCESS_KEY='SecretAccessKey'
export S3_ENDPOINT_URL='https://s3.us‑east‑2.amazonaws.com'

s5cmd sync --delete --destination-region 'us‑east‑2' "$DEEPSQUARE_OUTPUT/" 's3://test''/out'
)
`,
			title: "Positive test with S3 input output",
		},
		{
			input: model.Job{
				Env:           cleanJob.Env,
				EnableLogging: cleanJob.EnableLogging,
				InputMode:     cleanJob.InputMode,
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
/usr/bin/sleep 1
exec 3>&1
exec 4>&2
exec 1>>"/tmp/$SLURM_JOB_NAME-pipe"
exec 2>&1

disposeLogs() {
  echo cleaning up
  /usr/bin/sleep 15
  exec 1>&3
  exec 2>&4
  echo killing logger
  kill $LOGGER_PID || true
  wait $LOGGER_PID || true
  echo cleaned
}
trap disposeLogs EXIT
(
export STORAGE_PATH="/opt/cache/shared/$UID/$SLURM_JOB_NAME"
export DEEPSQUARE_INPUT="$STORAGE_PATH/input"
export DEEPSQUARE_OUTPUT="$STORAGE_PATH/output"
export DEEPSQUARE_ENV="$STORAGE_PATH/env"
/usr/bin/mkdir -p "$STORAGE_PATH" "$DEEPSQUARE_OUTPUT" "$DEEPSQUARE_INPUT"
/usr/bin/touch "$DEEPSQUARE_ENV"
/usr/bin/chmod -R 700 "$STORAGE_PATH"
/usr/bin/chown -R "$UID:cluster-users" "$STORAGE_PATH"
loadDeepsquareEnv() {
  /usr/bin/grep -v '^#' "$DEEPSQUARE_ENV" | /usr/bin/grep '=' | /usr/bin/sed -Ez '$ s/\n+$//' | tr '\n' ','
}
export AWS_ACCESS_KEY_ID='AccessKeyID'
export AWS_SECRET_ACCESS_KEY='SecretAccessKey'
export S3_ENDPOINT_URL='https://s3.us‑east‑2.amazonaws.com'

s5cmd cp --source-region 'us‑east‑2' 's3://test''/in''*' "$DEEPSQUARE_INPUT/"
/usr/bin/chmod -R 700 "$DEEPSQUARE_INPUT/" || echo "chmod failed, but we are ignoring it"
/usr/bin/echo "Input contains:"
/usr/bin/find "$DEEPSQUARE_INPUT/" -exec realpath --relative-to "$DEEPSQUARE_INPUT/" {} \;
/usr/bin/chmod -R 755 "$DEEPSQUARE_INPUT/"
export 'key'='test'\''test'
/usr/bin/echo 'Running: ''test'
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro"
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' /usr/bin/srun --job-name='test' \
  --export=ALL,"$(loadDeepsquareEnv)",'test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  --container-mounts="${MOUNTS}" \
  --container-image='image' \
  /bin/sh -c '/usr/bin/echo '\''hello world'\'''
/usr/bin/echo "Output contains:"
/usr/bin/find "$DEEPSQUARE_OUTPUT/" -exec realpath --relative-to "$DEEPSQUARE_OUTPUT/" {} \;
export AWS_ACCESS_KEY_ID='AccessKeyID'
export AWS_SECRET_ACCESS_KEY='SecretAccessKey'
export S3_ENDPOINT_URL='https://s3.us‑east‑2.amazonaws.com'

s5cmd sync --destination-region 'us‑east‑2' "$DEEPSQUARE_OUTPUT/" 's3://test''/out'
)
`,
			title: "Positive test with S3 input output and continuous sync",
		},
		{
			input: model.Job{
				Env:           cleanJob.Env,
				EnableLogging: cleanJob.EnableLogging,
				Resources:     cleanJob.Resources,
				InputMode:     cleanJob.InputMode,
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
/usr/bin/sleep 1
exec 3>&1
exec 4>&2
exec 1>>"/tmp/$SLURM_JOB_NAME-pipe"
exec 2>&1

disposeLogs() {
  echo cleaning up
  /usr/bin/sleep 15
  exec 1>&3
  exec 2>&4
  echo killing logger
  kill $LOGGER_PID || true
  wait $LOGGER_PID || true
  echo cleaned
}
trap disposeLogs EXIT
(
export STORAGE_PATH="/opt/cache/shared/$UID/$SLURM_JOB_NAME"
export DEEPSQUARE_INPUT="$STORAGE_PATH/input"
export DEEPSQUARE_OUTPUT="$STORAGE_PATH/output"
export DEEPSQUARE_ENV="$STORAGE_PATH/env"
/usr/bin/mkdir -p "$STORAGE_PATH" "$DEEPSQUARE_OUTPUT" "$DEEPSQUARE_INPUT"
/usr/bin/touch "$DEEPSQUARE_ENV"
/usr/bin/chmod -R 700 "$STORAGE_PATH"
/usr/bin/chown -R "$UID:cluster-users" "$STORAGE_PATH"
loadDeepsquareEnv() {
  /usr/bin/grep -v '^#' "$DEEPSQUARE_ENV" | /usr/bin/grep '=' | /usr/bin/sed -Ez '$ s/\n+$//' | tr '\n' ','
}
cd $DEEPSQUARE_INPUT/
/usr/bin/curl -fsORSL 'https://test/in'
for filepath in "$DEEPSQUARE_INPUT/"*; do
  /usr/bin/tar -xvaf "$filepath" 2>/dev/null && continue
  case $(file "$filepath") in
      *bzip2*) /usr/bin/bzip2 -fdk "$filepath";;
      *gzip*) /usr/bin/gunzip -df "$filepath";;
      *zip*) ;&
      *Zip*) /usr/bin/unzip -o "$filepath";;
      *RAR*) /usr/bin/unrar x -o+ "$filepath";;
      *xz*) ;&
      *XZ*) /usr/bin/unxz -f "$filepath";;
      *'7-zip'*) /usr/bin/7z x "$filepath" -aoa;;
      *) 1>&2 /usr/bin/echo "Unknown archive '$filepath'";;
  esac
done
cd -
/usr/bin/chmod -R 700 "$DEEPSQUARE_INPUT/" || echo "chmod failed, but we are ignoring it"
/usr/bin/echo "Input contains:"
/usr/bin/find "$DEEPSQUARE_INPUT/" -exec realpath --relative-to "$DEEPSQUARE_INPUT/" {} \;
/usr/bin/chmod -R 755 "$DEEPSQUARE_INPUT/"
export 'key'='test'\''test'
/usr/bin/echo 'Running: ''test'
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro"
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' /usr/bin/srun --job-name='test' \
  --export=ALL,"$(loadDeepsquareEnv)",'test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  --container-mounts="${MOUNTS}" \
  --container-image='image' \
  /bin/sh -c '/usr/bin/echo '\''hello world'\'''
/usr/bin/echo "Output contains:"
/usr/bin/find "$DEEPSQUARE_OUTPUT/" -exec realpath --relative-to "$DEEPSQUARE_OUTPUT/" {} \;
cd $DEEPSQUARE_OUTPUT/..
if [ "$(find output/ -type f | wc -l)" -eq 1 ]; then
/usr/bin/curl --upload-file "$(find output/ -type f | wc -l)" 'https://test/out'
else
/usr/bin/zip -r "output.zip" "output/"
/usr/bin/curl --upload-file "output.zip" 'https://test/out'
fi
/usr/bin/echo
cd -
)
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
export DEEPSQUARE_INPUT="$STORAGE_PATH/input"
export DEEPSQUARE_OUTPUT="$STORAGE_PATH/output"
export DEEPSQUARE_ENV="$STORAGE_PATH/env"
/usr/bin/mkdir -p "$STORAGE_PATH" "$DEEPSQUARE_OUTPUT" "$DEEPSQUARE_INPUT"
/usr/bin/touch "$DEEPSQUARE_ENV"
/usr/bin/chmod -R 700 "$STORAGE_PATH"
/usr/bin/chown -R "$UID:cluster-users" "$STORAGE_PATH"
loadDeepsquareEnv() {
  /usr/bin/grep -v '^#' "$DEEPSQUARE_ENV" | /usr/bin/grep '=' | /usr/bin/sed -Ez '$ s/\n+$//' | tr '\n' ','
}
export 'key'='test'\''test'
/usr/bin/echo 'Running: ''test'
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro"
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' /usr/bin/srun --job-name='test' \
  --export=ALL,"$(loadDeepsquareEnv)",'test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  --container-mounts="${MOUNTS}" \
  --container-image='image' \
  /bin/sh -c '/usr/bin/echo '\''hello world'\'''
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
