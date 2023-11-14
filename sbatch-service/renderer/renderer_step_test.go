// Copyright (C) 2023 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package renderer_test

import (
	"testing"

	"github.com/deepsquare-io/grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/grid/sbatch-service/renderer"
	"github.com/deepsquare-io/grid/sbatch-service/utils"
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
/usr/bin/enroot import -o "$IMAGE_PATH" -- 'docker://registry#image' &> "/tmp/enroot.import.$SLURM_JOB_ID.log"
if [ $? -ne 0 ]; then
  cat "/tmp/enroot.import.$SLURM_JOB_ID.log"
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
# shellcheck disable=SC2097,SC2098,SC1078
/usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  /bin/sh -c '/usr/bin/enroot create --name "container-$SLURM_JOB_ID.$SLURM_STEP_ID" -- "$IMAGE_PATH"
enrootClean() {
  /usr/bin/enroot remove -f "container-$SLURM_JOB_ID.$SLURM_STEP_ID"
}
trap enrootClean EXIT INT TERM
''/usr/bin/cat <<'"'"'EOFenroot'"'"' >"$STORAGE_PATH/enroot.conf"
#ENROOT_REMAP_ROOT=n
#ENROOT_ROOTFS_WRITABLE=y
#ENROOT_MOUNT_HOME=n

environ() {
  # Keep all the environment from the host
  /usr/bin/env

  /usr/bin/cat "${ENROOT_ROOTFS}/etc/environment"

  /usr/bin/echo "STORAGE_PATH=/deepsquare"
  /usr/bin/echo "DEEPSQUARE_TMP=/deepsquare/tmp"
  /usr/bin/echo "DEEPSQUARE_SHARED_TMP=/deepsquare/tmp"
  /usr/bin/echo "DEEPSQUARE_SHARED_WORLD_TMP=/deepsquare/world-tmp"
  /usr/bin/echo "DEEPSQUARE_DISK_TMP=/deepsquare/disk/tmp"
  /usr/bin/echo "DEEPSQUARE_DISK_WORLD_TMP=/deepsquare/disk/world-tmp"
  /usr/bin/echo "DEEPSQUARE_INPUT=/deepsquare/input"
  /usr/bin/echo "DEEPSQUARE_OUTPUT=/deepsquare/output"
  /usr/bin/echo "DEEPSQUARE_ENV=/deepsquare/$(basename $DEEPSQUARE_ENV)"
  /usr/bin/echo "test='"'"'value'"'"'"
}

mounts() {
  /usr/bin/echo "$STORAGE_PATH /deepsquare none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_SHARED_TMP /deepsquare/tmp none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_SHARED_WORLD_TMP /deepsquare/world-tmp none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_DISK_TMP /deepsquare/disk/tmp none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_DISK_WORLD_TMP /deepsquare/disk/world-tmp none x-create=dir,bind,rw"
  /usr/bin/echo "/tmp/.X11-unix /tmp/.X11-unix none x-create=dir,bind,ro"
  /usr/bin/echo '"'"'/host /container none x-create=auto,bind,ro'"'"'
}

hooks() {
  /usr/bin/cat << '"'"'EOFrclocal'"'"' > "${ENROOT_ROOTFS}/etc/rc.local"
cd "/deepsquare" || { echo "change dir to working directory failed"; exit 1; }
exec "$@"
EOFrclocal
}
EOFenroot
/usr/bin/enroot start \
  --conf "$STORAGE_PATH/enroot.conf" \
  "container-$SLURM_JOB_ID.$SLURM_STEP_ID" \
  /bin/sh -c '"'"'hostname'"'"''`,
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
/usr/bin/enroot import -o "$IMAGE_PATH" -- 'docker://registry#image' &> "/tmp/enroot.import.$SLURM_JOB_ID.log"
if [ $? -ne 0 ]; then
  cat "/tmp/enroot.import.$SLURM_JOB_ID.log"
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
# shellcheck disable=SC2097,SC2098,SC1078
/usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  /bin/sh -c '/usr/bin/enroot create --name "container-$SLURM_JOB_ID.$SLURM_STEP_ID" -- "$IMAGE_PATH"
enrootClean() {
  /usr/bin/enroot remove -f "container-$SLURM_JOB_ID.$SLURM_STEP_ID"
}
trap enrootClean EXIT INT TERM
''/usr/bin/cat <<'"'"'EOFenroot'"'"' >"$STORAGE_PATH/enroot.conf"
#ENROOT_REMAP_ROOT=n
#ENROOT_ROOTFS_WRITABLE=y
#ENROOT_MOUNT_HOME=n

environ() {
  # Keep all the environment from the host
  /usr/bin/env

  /usr/bin/cat "${ENROOT_ROOTFS}/etc/environment"

  /usr/bin/echo "STORAGE_PATH=/deepsquare"
  /usr/bin/echo "DEEPSQUARE_TMP=/deepsquare/tmp"
  /usr/bin/echo "DEEPSQUARE_SHARED_TMP=/deepsquare/tmp"
  /usr/bin/echo "DEEPSQUARE_SHARED_WORLD_TMP=/deepsquare/world-tmp"
  /usr/bin/echo "DEEPSQUARE_DISK_TMP=/deepsquare/disk/tmp"
  /usr/bin/echo "DEEPSQUARE_DISK_WORLD_TMP=/deepsquare/disk/world-tmp"
  /usr/bin/echo "DEEPSQUARE_INPUT=/deepsquare/input"
  /usr/bin/echo "DEEPSQUARE_OUTPUT=/deepsquare/output"
  /usr/bin/echo "DEEPSQUARE_ENV=/deepsquare/$(basename $DEEPSQUARE_ENV)"
  /usr/bin/echo "test='"'"'value'"'"'"
}

mounts() {
  /usr/bin/echo "$STORAGE_PATH /deepsquare none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_SHARED_TMP /deepsquare/tmp none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_SHARED_WORLD_TMP /deepsquare/world-tmp none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_DISK_TMP /deepsquare/disk/tmp none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_DISK_WORLD_TMP /deepsquare/disk/world-tmp none x-create=dir,bind,rw"
  /usr/bin/echo "/tmp/.X11-unix /tmp/.X11-unix none x-create=dir,bind,ro"
  /usr/bin/echo '"'"'/host /container none x-create=auto,bind,ro'"'"'
}

hooks() {
  /usr/bin/cat << '"'"'EOFrclocal'"'"' > "${ENROOT_ROOTFS}/etc/rc.local"
cd "/deepsquare" || { echo "change dir to working directory failed"; exit 1; }
exec "$@"
EOFrclocal
}
EOFenroot
/usr/bin/enroot start \
  --conf "$STORAGE_PATH/enroot.conf" \
  "container-$SLURM_JOB_ID.$SLURM_STEP_ID" \
  /bin/sh -c '"'"'/usr/bin/echo $item'"'"''
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
/usr/bin/enroot import -o "$IMAGE_PATH" -- 'docker://registry#image' &> "/tmp/enroot.import.$SLURM_JOB_ID.log"
if [ $? -ne 0 ]; then
  cat "/tmp/enroot.import.$SLURM_JOB_ID.log"
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
# shellcheck disable=SC2097,SC2098,SC1078
/usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  /bin/sh -c '/usr/bin/enroot create --name "container-$SLURM_JOB_ID.$SLURM_STEP_ID" -- "$IMAGE_PATH"
enrootClean() {
  /usr/bin/enroot remove -f "container-$SLURM_JOB_ID.$SLURM_STEP_ID"
}
trap enrootClean EXIT INT TERM
''/usr/bin/cat <<'"'"'EOFenroot'"'"' >"$STORAGE_PATH/enroot.conf"
#ENROOT_REMAP_ROOT=n
#ENROOT_ROOTFS_WRITABLE=y
#ENROOT_MOUNT_HOME=n

environ() {
  # Keep all the environment from the host
  /usr/bin/env

  /usr/bin/cat "${ENROOT_ROOTFS}/etc/environment"

  /usr/bin/echo "STORAGE_PATH=/deepsquare"
  /usr/bin/echo "DEEPSQUARE_TMP=/deepsquare/tmp"
  /usr/bin/echo "DEEPSQUARE_SHARED_TMP=/deepsquare/tmp"
  /usr/bin/echo "DEEPSQUARE_SHARED_WORLD_TMP=/deepsquare/world-tmp"
  /usr/bin/echo "DEEPSQUARE_DISK_TMP=/deepsquare/disk/tmp"
  /usr/bin/echo "DEEPSQUARE_DISK_WORLD_TMP=/deepsquare/disk/world-tmp"
  /usr/bin/echo "DEEPSQUARE_INPUT=/deepsquare/input"
  /usr/bin/echo "DEEPSQUARE_OUTPUT=/deepsquare/output"
  /usr/bin/echo "DEEPSQUARE_ENV=/deepsquare/$(basename $DEEPSQUARE_ENV)"
  /usr/bin/echo "test='"'"'value'"'"'"
}

mounts() {
  /usr/bin/echo "$STORAGE_PATH /deepsquare none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_SHARED_TMP /deepsquare/tmp none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_SHARED_WORLD_TMP /deepsquare/world-tmp none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_DISK_TMP /deepsquare/disk/tmp none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_DISK_WORLD_TMP /deepsquare/disk/world-tmp none x-create=dir,bind,rw"
  /usr/bin/echo "/tmp/.X11-unix /tmp/.X11-unix none x-create=dir,bind,ro"
  /usr/bin/echo '"'"'/host /container none x-create=auto,bind,ro'"'"'
}

hooks() {
  /usr/bin/cat << '"'"'EOFrclocal'"'"' > "${ENROOT_ROOTFS}/etc/rc.local"
cd "/deepsquare" || { echo "change dir to working directory failed"; exit 1; }
exec "$@"
EOFrclocal
}
EOFenroot
/usr/bin/enroot start \
  --conf "$STORAGE_PATH/enroot.conf" \
  "container-$SLURM_JOB_ID.$SLURM_STEP_ID" \
  /bin/sh -c '"'"'/usr/bin/echo $item'"'"''
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
