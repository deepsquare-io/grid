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

var (
	cleanStepForItems = model.StepFor{
		Parallel: true,
		Items:    []string{"a", "b", "c"},
		Steps: []*model.Step{
			{
				Name: utils.Ptr("test"),
				Run:  cleanStepRun("/usr/bin/echo $item"),
			},
			{
				Name: utils.Ptr("test"),
				Run:  cleanStepRun("/usr/bin/echo $item"),
			},
		},
	}
	cleanStepForRange = model.StepFor{
		Parallel: true,
		Range: &model.ForRange{
			Begin:     0,
			End:       -10,
			Increment: utils.Ptr(-2),
		},
		Steps: []*model.Step{
			{
				Name: utils.Ptr("test"),
				Run:  cleanStepRun("/usr/bin/echo $index"),
			},
			{
				Name: utils.Ptr("test"),
				Run:  cleanStepRun("/usr/bin/echo $index"),
			},
		},
	}
)

func TestRenderStepFor(t *testing.T) {
	tests := []struct {
		input         model.StepFor
		isError       bool
		errorContains []string
		expected      string
		title         string
	}{
		{
			input: cleanStepForItems,
			expected: `doFor() {
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
  /bin/sh -c '/usr/bin/enroot create --name "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID" -- "$IMAGE_PATH" >/dev/null 2>&1
enrootClean() {
  /usr/bin/enroot remove -f "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID"
}
trap enrootClean EXIT INT TERM
''/usr/bin/cat <<'"'"'EOFenroot'"'"' >"$STORAGE_PATH/enroot-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID.conf"
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
  --conf "$STORAGE_PATH/enroot-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID.conf" \
  "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID" \
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
  /bin/sh -c '/usr/bin/enroot create --name "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID" -- "$IMAGE_PATH" >/dev/null 2>&1
enrootClean() {
  /usr/bin/enroot remove -f "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID"
}
trap enrootClean EXIT INT TERM
''/usr/bin/cat <<'"'"'EOFenroot'"'"' >"$STORAGE_PATH/enroot-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID.conf"
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
  --conf "$STORAGE_PATH/enroot-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID.conf" \
  "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID" \
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
			title: "Positive test with items",
		},
		{
			input: cleanStepForRange,
			expected: `doFor() {
export index="$1"
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
  /bin/sh -c '/usr/bin/enroot create --name "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID" -- "$IMAGE_PATH" >/dev/null 2>&1
enrootClean() {
  /usr/bin/enroot remove -f "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID"
}
trap enrootClean EXIT INT TERM
''/usr/bin/cat <<'"'"'EOFenroot'"'"' >"$STORAGE_PATH/enroot-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID.conf"
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
  --conf "$STORAGE_PATH/enroot-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID.conf" \
  "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID" \
  /bin/sh -c '"'"'/usr/bin/echo $index'"'"''
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
  /bin/sh -c '/usr/bin/enroot create --name "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID" -- "$IMAGE_PATH" >/dev/null 2>&1
enrootClean() {
  /usr/bin/enroot remove -f "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID"
}
trap enrootClean EXIT INT TERM
''/usr/bin/cat <<'"'"'EOFenroot'"'"' >"$STORAGE_PATH/enroot-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID.conf"
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
  --conf "$STORAGE_PATH/enroot-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID.conf" \
  "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID" \
  /bin/sh -c '"'"'/usr/bin/echo $index'"'"''
}
pids=()
for index in $(seq 0 -2 -10); do
  doFor "$index" &
  pids+=("$!")
done
for pid in "${pids[@]}"; do
  wait "$pid"
done`,
			title: "Positive test with range",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			actual, err := renderer.RenderStepFor(&cleanJob, &tt.input)

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
