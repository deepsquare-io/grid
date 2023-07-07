package renderer_test

import (
	"testing"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/renderer"
	"github.com/deepsquare-io/the-grid/sbatch-service/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func cleanStepWithRun(r *model.StepRun) *model.Step {
	return &model.Step{
		Name: utils.Ptr("test"),
		Run:  r,
	}
}

func cleanStepWithFor(f *model.StepFor) *model.Step {
	return &model.Step{
		Name: utils.Ptr("test"),
		For:  f,
	}
}

func cleanStepWithAsyncLaunch(s *model.StepAsyncLaunch) *model.Step {
	return &model.Step{
		Name:   utils.Ptr("test"),
		Launch: s,
	}
}

func cleanStepWithUse(s *model.StepUse) *model.Step {
	return &model.Step{
		Name: utils.Ptr("test"),
		Use:  s,
	}
}

func TestRenderStep(t *testing.T) {
	tests := []struct {
		input         model.Step
		isError       bool
		errorContains []string
		expected      string
		title         string
	}{
		{
			input: *cleanStepWithRun(cleanStepRun("hostname")),
			expected: `/usr/bin/echo 'Running: ''test'
/usr/bin/cat << 'EOFmounterror'
WARNING: Mounts is now deprecated.
If you need a cache (disk, shared, per-user or global), please read https://docs.deepsquare.run/workflow/guides/environment-variables.
The cache is cleared periodically and only persists on the site.
EOFmounterror
/usr/bin/mkdir -p "$HOME/.config/enroot/"
/usr/bin/cat << 'EOFnetrc' > "$HOME/.config/enroot/.credentials"
machine registry login "username" password "password"
EOFnetrc
IMAGE_PATH="$STORAGE_PATH/$SLURM_JOB_ID-$(echo $RANDOM | md5sum | head -c 20).sqsh"
export IMAGE_PATH
/usr/bin/echo "Importing image..."
set +e
/usr/bin/enroot import -o "$IMAGE_PATH" -- 'docker://registry#image' &> /tmp/enroot.import.$SLURM_JOB_ID.log
if [ $? -ne 0 ]; then
  cat /tmp/enroot.import.$SLURM_JOB_ID.log
fi
set -e
tries=1; while [ "$tries" -lt 10 ]; do
  if /usr/bin/file "$IMAGE_PATH" | /usr/bin/grep -q "Squashfs filesystem"; then
    break
  fi
  /usr/bin/echo "Image is not complete. Wait a few seconds... ($tries/10)"
  /usr/bin/sleep 10
  tries=$((tries+1))
done
if [ "$tries" -ge 10 ]; then
  /usr/bin/echo "Image import failure (corrupted image). Please try again."
  exit 1
fi
/usr/bin/echo "Image successfully imported!"
MOUNTS="$STORAGE_PATH:/deepsquare:rw,$DEEPSQUARE_SHARED_TMP:/deepsquare/tmp:rw,$DEEPSQUARE_SHARED_WORLD_TMP:/deepsquare/world-tmp:rw,$DEEPSQUARE_DISK_TMP:/deepsquare/disk/tmp:rw,$DEEPSQUARE_DISK_WORLD_TMP:/deepsquare/disk/world-tmp:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' \
DEEPSQUARE_TMP='/deepsquare/tmp' \
DEEPSQUARE_SHARED_TMP='/deepsquare/tmp' \
DEEPSQUARE_SHARED_WORLD_TMP='/deepsquare/world-tmp' \
DEEPSQUARE_DISK_TMP='/deepsquare/disk/tmp' \
DEEPSQUARE_DISK_WORLD_TMP='/deepsquare/disk/world-tmp' \
DEEPSQUARE_INPUT='/deepsquare/input' \
DEEPSQUARE_OUTPUT='/deepsquare/output' \
DEEPSQUARE_ENV="/deepsquare/$(basename $DEEPSQUARE_ENV)" test='value' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  --no-container-remap-root \
  --no-container-mount-home \
  --container-mounts="${MOUNTS}" \
  --container-workdir=/deepsquare \
  --container-image="$IMAGE_PATH" \
  /bin/sh -c 'hostname'`,
			title: "Positive test with run",
		},
		{
			input: *cleanStepWithAsyncLaunch(cleanStepAsyncLaunchWithStep("async_launch", &model.Step{})),
			expected: `/usr/bin/echo 'Running: ''test'
(
declare -A EXIT_SIGNALS


for pid in "${!EXIT_SIGNALS[@]}"; do
  /usr/bin/echo "Process $$ sending signal ${EXIT_SIGNALS[$pid]} to $pid"
  kill -s "${EXIT_SIGNALS[$pid]}" "$pid" || echo "Sending signal ${EXIT_SIGNALS[$pid]} to $pid failed, continuing..."
done
) &
asynclaunchpid="$!"
export PID_ASYNC_LAUNCH="$asynclaunchpid"
EXIT_SIGNALS[$asynclaunchpid]=15`,
			title: "Positive test with async launch",
		},
		{
			input: model.Step{
				Name:      utils.Ptr("test"),
				DependsOn: []string{"test2"},
			},
			expected: `if [ -n "${PID_TEST2+x}" ]; then
/usr/bin/echo 'Waiting for: TEST2.'
wait "${PID_TEST2}"
else
/usr/bin/echo 'Cannot await: TEST2 is undefined. Exiting to avoid undefined behavior.'
/usr/bin/echo 'Is the TEST2 handleName set and is defined at the same scope ?'
exit 1
fi
/usr/bin/echo 'Running: ''test'`,
			title: "Positive test with await",
		},
		{
			input: model.Step{
				Name: utils.Ptr("test"),
				If:   utils.Ptr(`$CONDITION_RESULT = "false"`),
			},
			expected: `
# shellcheck disable=SC2016,SC2089
CONDITION='$CONDITION_RESULT = "false"'
# shellcheck disable=SC2090
export CONDITION
CONDITION_RESULT="$(eval "if [[ $CONDITION ]]; then echo 'true' ; else echo 'false' ; fi")"
if [ $CONDITION_RESULT = "true" ]; then
/usr/bin/echo 'Running: ''test'
:
fi
`,
			title: "Positive test with if",
		},
		{
			input: *cleanStepWithFor(&cleanStepForItems),
			expected: `/usr/bin/echo 'Running: ''test'
doFor() {
export item="$1"
/usr/bin/echo 'Running: ''test'
/usr/bin/cat << 'EOFmounterror'
WARNING: Mounts is now deprecated.
If you need a cache (disk, shared, per-user or global), please read https://docs.deepsquare.run/workflow/guides/environment-variables.
The cache is cleared periodically and only persists on the site.
EOFmounterror
/usr/bin/mkdir -p "$HOME/.config/enroot/"
/usr/bin/cat << 'EOFnetrc' > "$HOME/.config/enroot/.credentials"
machine registry login "username" password "password"
EOFnetrc
IMAGE_PATH="$STORAGE_PATH/$SLURM_JOB_ID-$(echo $RANDOM | md5sum | head -c 20).sqsh"
export IMAGE_PATH
/usr/bin/echo "Importing image..."
set +e
/usr/bin/enroot import -o "$IMAGE_PATH" -- 'docker://registry#image' &> /tmp/enroot.import.$SLURM_JOB_ID.log
if [ $? -ne 0 ]; then
  cat /tmp/enroot.import.$SLURM_JOB_ID.log
fi
set -e
tries=1; while [ "$tries" -lt 10 ]; do
  if /usr/bin/file "$IMAGE_PATH" | /usr/bin/grep -q "Squashfs filesystem"; then
    break
  fi
  /usr/bin/echo "Image is not complete. Wait a few seconds... ($tries/10)"
  /usr/bin/sleep 10
  tries=$((tries+1))
done
if [ "$tries" -ge 10 ]; then
  /usr/bin/echo "Image import failure (corrupted image). Please try again."
  exit 1
fi
/usr/bin/echo "Image successfully imported!"
MOUNTS="$STORAGE_PATH:/deepsquare:rw,$DEEPSQUARE_SHARED_TMP:/deepsquare/tmp:rw,$DEEPSQUARE_SHARED_WORLD_TMP:/deepsquare/world-tmp:rw,$DEEPSQUARE_DISK_TMP:/deepsquare/disk/tmp:rw,$DEEPSQUARE_DISK_WORLD_TMP:/deepsquare/disk/world-tmp:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' \
DEEPSQUARE_TMP='/deepsquare/tmp' \
DEEPSQUARE_SHARED_TMP='/deepsquare/tmp' \
DEEPSQUARE_SHARED_WORLD_TMP='/deepsquare/world-tmp' \
DEEPSQUARE_DISK_TMP='/deepsquare/disk/tmp' \
DEEPSQUARE_DISK_WORLD_TMP='/deepsquare/disk/world-tmp' \
DEEPSQUARE_INPUT='/deepsquare/input' \
DEEPSQUARE_OUTPUT='/deepsquare/output' \
DEEPSQUARE_ENV="/deepsquare/$(basename $DEEPSQUARE_ENV)" test='value' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  --no-container-remap-root \
  --no-container-mount-home \
  --container-mounts="${MOUNTS}" \
  --container-workdir=/deepsquare \
  --container-image="$IMAGE_PATH" \
  /bin/sh -c '/usr/bin/echo $item'
/usr/bin/echo 'Running: ''test'
/usr/bin/cat << 'EOFmounterror'
WARNING: Mounts is now deprecated.
If you need a cache (disk, shared, per-user or global), please read https://docs.deepsquare.run/workflow/guides/environment-variables.
The cache is cleared periodically and only persists on the site.
EOFmounterror
/usr/bin/mkdir -p "$HOME/.config/enroot/"
/usr/bin/cat << 'EOFnetrc' > "$HOME/.config/enroot/.credentials"
machine registry login "username" password "password"
EOFnetrc
IMAGE_PATH="$STORAGE_PATH/$SLURM_JOB_ID-$(echo $RANDOM | md5sum | head -c 20).sqsh"
export IMAGE_PATH
/usr/bin/echo "Importing image..."
set +e
/usr/bin/enroot import -o "$IMAGE_PATH" -- 'docker://registry#image' &> /tmp/enroot.import.$SLURM_JOB_ID.log
if [ $? -ne 0 ]; then
  cat /tmp/enroot.import.$SLURM_JOB_ID.log
fi
set -e
tries=1; while [ "$tries" -lt 10 ]; do
  if /usr/bin/file "$IMAGE_PATH" | /usr/bin/grep -q "Squashfs filesystem"; then
    break
  fi
  /usr/bin/echo "Image is not complete. Wait a few seconds... ($tries/10)"
  /usr/bin/sleep 10
  tries=$((tries+1))
done
if [ "$tries" -ge 10 ]; then
  /usr/bin/echo "Image import failure (corrupted image). Please try again."
  exit 1
fi
/usr/bin/echo "Image successfully imported!"
MOUNTS="$STORAGE_PATH:/deepsquare:rw,$DEEPSQUARE_SHARED_TMP:/deepsquare/tmp:rw,$DEEPSQUARE_SHARED_WORLD_TMP:/deepsquare/world-tmp:rw,$DEEPSQUARE_DISK_TMP:/deepsquare/disk/tmp:rw,$DEEPSQUARE_DISK_WORLD_TMP:/deepsquare/disk/world-tmp:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' \
DEEPSQUARE_TMP='/deepsquare/tmp' \
DEEPSQUARE_SHARED_TMP='/deepsquare/tmp' \
DEEPSQUARE_SHARED_WORLD_TMP='/deepsquare/world-tmp' \
DEEPSQUARE_DISK_TMP='/deepsquare/disk/tmp' \
DEEPSQUARE_DISK_WORLD_TMP='/deepsquare/disk/world-tmp' \
DEEPSQUARE_INPUT='/deepsquare/input' \
DEEPSQUARE_OUTPUT='/deepsquare/output' \
DEEPSQUARE_ENV="/deepsquare/$(basename $DEEPSQUARE_ENV)" test='value' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  --no-container-remap-root \
  --no-container-mount-home \
  --container-mounts="${MOUNTS}" \
  --container-workdir=/deepsquare \
  --container-image="$IMAGE_PATH" \
  /bin/sh -c '/usr/bin/echo $item'
}
pids=()
items=('a' 'b' 'c' )
for item in "${items[@]}"; do
  doFor "$item" &
  pids+=("$!")
done
for pid in "${pids[@]}"; do
  wait "$pid"
done`,
			title: "Positive test with for",
		},
		{
			input: model.Step{
				Name: utils.Ptr("test"),
				Steps: []*model.Step{
					{
						Name: utils.Ptr("test 2"),
					},
				},
			},
			expected: `/usr/bin/echo 'Running: ''test'
/usr/bin/echo 'Running: ''test 2'`,
			title: "Positive test with step of steps",
		},
		{
			input: model.Step{
				Name: utils.Ptr("catch"),
				Run: &model.StepRun{
					Command: "exit 1",
				},
				Catch: []*model.Step{
					{
						Run: &model.StepRun{
							Command: "echo $DEEPSQUARE_ERROR_CODE",
						},
					},
				},
			},
			expected: `
( # CATCH FINALLY
set +e
/usr/bin/echo 'Running: ''catch'
( # CATCH
set -e
/usr/bin/srun --job-name='catch' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=4 \
  --mem-per-cpu=4096M \
  --gpus-per-task=1 \
  --ntasks=1 \
  /bin/sh -c 'exit 1'
) # CATCH
DEEPSQUARE_ERROR_CODE=$?
export DEEPSQUARE_ERROR_CODE
if [ $DEEPSQUARE_ERROR_CODE -ne 0 ]; then
set -e

/usr/bin/srun  \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=4 \
  --mem-per-cpu=4096M \
  --gpus-per-task=1 \
  --ntasks=1 \
  /bin/sh -c 'echo $DEEPSQUARE_ERROR_CODE'
fi
) # CATCH FINALLY`,
			title: "Positive test with catch",
		},
		{
			input: model.Step{
				Name: utils.Ptr("catch"),
				Run: &model.StepRun{
					Command: "exit 1",
				},
				Finally: []*model.Step{
					{
						Run: &model.StepRun{
							Command: "cleaning up",
						},
					},
				},
			},
			expected: `
( # CATCH FINALLY
finally() {
set -e

/usr/bin/srun  \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=4 \
  --mem-per-cpu=4096M \
  --gpus-per-task=1 \
  --ntasks=1 \
  /bin/sh -c 'cleaning up'
}
trap finally EXIT INT TERM
/usr/bin/echo 'Running: ''catch'
/usr/bin/srun --job-name='catch' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=4 \
  --mem-per-cpu=4096M \
  --gpus-per-task=1 \
  --ntasks=1 \
  /bin/sh -c 'exit 1'
) # CATCH FINALLY`,
			title: "Positive test with finally",
		},
		{
			input: model.Step{
				Name: utils.Ptr("catch"),
				Run: &model.StepRun{
					Command: "exit 1",
				},
				Catch: []*model.Step{
					{
						Run: &model.StepRun{
							Command: "echo $DEEPSQUARE_ERROR_CODE",
						},
					},
				},
				Finally: []*model.Step{
					{
						Run: &model.StepRun{
							Command: "cleaning up",
						},
					},
				},
			},
			expected: `
( # CATCH FINALLY
set +e

finally() {
set -e

/usr/bin/srun  \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=4 \
  --mem-per-cpu=4096M \
  --gpus-per-task=1 \
  --ntasks=1 \
  /bin/sh -c 'cleaning up'
}
trap finally EXIT INT TERM
/usr/bin/echo 'Running: ''catch'
( # CATCH
set -e
/usr/bin/srun --job-name='catch' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=4 \
  --mem-per-cpu=4096M \
  --gpus-per-task=1 \
  --ntasks=1 \
  /bin/sh -c 'exit 1'
) # CATCH
DEEPSQUARE_ERROR_CODE=$?
export DEEPSQUARE_ERROR_CODE
if [ $DEEPSQUARE_ERROR_CODE -ne 0 ]; then
set -e

/usr/bin/srun  \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=4 \
  --mem-per-cpu=4096M \
  --gpus-per-task=1 \
  --ntasks=1 \
  /bin/sh -c 'echo $DEEPSQUARE_ERROR_CODE'
fi
) # CATCH FINALLY`,
			title: "Positive test with catch-finally",
		},
		{
			input: model.Step{
				Name: utils.Ptr("test"),
			},
			expected: "/usr/bin/echo 'Running: ''test'",
			title:    "Positive test with none",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			actual, err := renderer.RenderStep(&cleanJob, &tt.input)

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, actual)
				require.NoError(t, renderer.Shellcheck(actual))
			}
		})
	}
}
