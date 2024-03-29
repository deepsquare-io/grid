// Copyright (C) 2024 DeepSquare Association
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

func cleanEnrootStepRun(command string) *model.StepRun {
	return &model.StepRun{
		Resources: &cleanStepRunResources,
		Container: &model.ContainerRun{
			Image:    "image",
			Registry: utils.Ptr("registry"),
			Username: utils.Ptr("username"),
			Password: utils.Ptr("password"),
			X11:      utils.Ptr(true),
			Mounts: []*model.Mount{
				{
					HostDir:      "/host",
					ContainerDir: "/container",
					Options:      "ro",
				},
			},
		},
		Env: []*model.EnvVar{
			{
				Key:   "test",
				Value: "value",
			},
		},
		Command: command,
	}
}

func TestRenderEnrootCommand(t *testing.T) {
	tests := []struct {
		input         model.StepRun
		isError       bool
		errorContains []string
		expected      string
		title         string
	}{
		{
			input: *cleanEnrootStepRun("hostname"),
			expected: `/usr/bin/cat <<'EOFenroot' >"$STORAGE_PATH/enroot-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID.conf"
#ENROOT_REMAP_ROOT=n
#ENROOT_ROOTFS_WRITABLE=y
#ENROOT_MOUNT_HOME=n

environ() {
  # Keep all the environment from the host
  env

  cat "${ENROOT_ROOTFS}/etc/environment"

  echo "STORAGE_PATH=/deepsquare"
  echo "DEEPSQUARE_TMP=/deepsquare/tmp"
  echo "DEEPSQUARE_SHARED_TMP=/deepsquare/tmp"
  echo "DEEPSQUARE_SHARED_WORLD_TMP=/deepsquare/world-tmp"
  echo "DEEPSQUARE_DISK_TMP=/deepsquare/disk/tmp"
  echo "DEEPSQUARE_DISK_WORLD_TMP=/deepsquare/disk/world-tmp"
  echo "DEEPSQUARE_INPUT=/deepsquare/input"
  echo "DEEPSQUARE_OUTPUT=/deepsquare/output"
  echo "DEEPSQUARE_ENV=/deepsquare/$(basename $DEEPSQUARE_ENV)"
}

mounts() {
  echo "$STORAGE_PATH /deepsquare none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_SHARED_TMP /deepsquare/tmp none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_SHARED_WORLD_TMP /deepsquare/world-tmp none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_DISK_TMP /deepsquare/disk/tmp none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_DISK_WORLD_TMP /deepsquare/disk/world-tmp none x-create=dir,bind,rw"
  echo "/tmp/.X11-unix /tmp/.X11-unix none x-create=dir,bind,ro"
  echo '/host /container none x-create=auto,bind,ro'
}

hooks() {
  cat << 'EOFrclocal' > "${ENROOT_ROOTFS}/etc/rc.local"
test="$(cat << 'EOF[random_string]'
value
EOF[random_string]
)"
export test
exec "$@"
EOFrclocal
}
EOFenroot
/usr/bin/enroot start \
  --conf "$STORAGE_PATH/enroot-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID.conf" \
  "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID" \
  /bin/sh -c 'hostname'`,
			title: "Positive test with enroot image",
		},
		{
			input: func() model.StepRun {
				s := cleanEnrootStepRun("hostname")
				s.Env = append(s.Env, &model.EnvVar{
					Key: "test2",
					Value: `value1
value2`,
				})
				return *s
			}(),
			expected: `/usr/bin/cat <<'EOFenroot' >"$STORAGE_PATH/enroot-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID.conf"
#ENROOT_REMAP_ROOT=n
#ENROOT_ROOTFS_WRITABLE=y
#ENROOT_MOUNT_HOME=n

environ() {
  # Keep all the environment from the host
  env

  cat "${ENROOT_ROOTFS}/etc/environment"

  echo "STORAGE_PATH=/deepsquare"
  echo "DEEPSQUARE_TMP=/deepsquare/tmp"
  echo "DEEPSQUARE_SHARED_TMP=/deepsquare/tmp"
  echo "DEEPSQUARE_SHARED_WORLD_TMP=/deepsquare/world-tmp"
  echo "DEEPSQUARE_DISK_TMP=/deepsquare/disk/tmp"
  echo "DEEPSQUARE_DISK_WORLD_TMP=/deepsquare/disk/world-tmp"
  echo "DEEPSQUARE_INPUT=/deepsquare/input"
  echo "DEEPSQUARE_OUTPUT=/deepsquare/output"
  echo "DEEPSQUARE_ENV=/deepsquare/$(basename $DEEPSQUARE_ENV)"
}

mounts() {
  echo "$STORAGE_PATH /deepsquare none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_SHARED_TMP /deepsquare/tmp none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_SHARED_WORLD_TMP /deepsquare/world-tmp none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_DISK_TMP /deepsquare/disk/tmp none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_DISK_WORLD_TMP /deepsquare/disk/world-tmp none x-create=dir,bind,rw"
  echo "/tmp/.X11-unix /tmp/.X11-unix none x-create=dir,bind,ro"
  echo '/host /container none x-create=auto,bind,ro'
}

hooks() {
  cat << 'EOFrclocal' > "${ENROOT_ROOTFS}/etc/rc.local"
test="$(cat << 'EOF[random_string]'
value
EOF[random_string]
)"
export test
test2="$(cat << 'EOF[random_string]'
value1
value2
EOF[random_string]
)"
export test2
exec "$@"
EOFrclocal
}
EOFenroot
/usr/bin/enroot start \
  --conf "$STORAGE_PATH/enroot-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID.conf" \
  "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID" \
  /bin/sh -c 'hostname'`,
			title: "Positive test with multiline env",
		},
		{
			input: func() model.StepRun {
				r := *cleanEnrootStepRun("hostname")
				r.Container.Image = "/test/my.sqshfs"
				return r
			}(),
			expected: `/usr/bin/cat <<'EOFenroot' >"$STORAGE_PATH/enroot-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID.conf"
#ENROOT_REMAP_ROOT=n
#ENROOT_ROOTFS_WRITABLE=y
#ENROOT_MOUNT_HOME=n

environ() {
  # Keep all the environment from the host
  env

  cat "${ENROOT_ROOTFS}/etc/environment"

  echo "STORAGE_PATH=/deepsquare"
  echo "DEEPSQUARE_TMP=/deepsquare/tmp"
  echo "DEEPSQUARE_SHARED_TMP=/deepsquare/tmp"
  echo "DEEPSQUARE_SHARED_WORLD_TMP=/deepsquare/world-tmp"
  echo "DEEPSQUARE_DISK_TMP=/deepsquare/disk/tmp"
  echo "DEEPSQUARE_DISK_WORLD_TMP=/deepsquare/disk/world-tmp"
  echo "DEEPSQUARE_INPUT=/deepsquare/input"
  echo "DEEPSQUARE_OUTPUT=/deepsquare/output"
  echo "DEEPSQUARE_ENV=/deepsquare/$(basename $DEEPSQUARE_ENV)"
}

mounts() {
  echo "$STORAGE_PATH /deepsquare none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_SHARED_TMP /deepsquare/tmp none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_SHARED_WORLD_TMP /deepsquare/world-tmp none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_DISK_TMP /deepsquare/disk/tmp none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_DISK_WORLD_TMP /deepsquare/disk/world-tmp none x-create=dir,bind,rw"
  echo "/tmp/.X11-unix /tmp/.X11-unix none x-create=dir,bind,ro"
  echo '/host /container none x-create=auto,bind,ro'
}

hooks() {
  cat << 'EOFrclocal' > "${ENROOT_ROOTFS}/etc/rc.local"
test="$(cat << 'EOF[random_string]'
value
EOF[random_string]
)"
export test
exec "$@"
EOFrclocal
}
EOFenroot
/usr/bin/enroot start \
  --conf "$STORAGE_PATH/enroot-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID.conf" \
  "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID" \
  /bin/sh -c 'hostname'`,
			title: "Positive test with enroot absolute path image",
		},
		{
			input: func() model.StepRun {
				r := *cleanEnrootStepRun("hostname")
				r.WorkDir = utils.Ptr("/dir")
				return r
			}(),
			expected: `/usr/bin/cat <<'EOFenroot' >"$STORAGE_PATH/enroot-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID.conf"
#ENROOT_REMAP_ROOT=n
#ENROOT_ROOTFS_WRITABLE=y
#ENROOT_MOUNT_HOME=n

environ() {
  # Keep all the environment from the host
  env

  cat "${ENROOT_ROOTFS}/etc/environment"

  echo "STORAGE_PATH=/deepsquare"
  echo "DEEPSQUARE_TMP=/deepsquare/tmp"
  echo "DEEPSQUARE_SHARED_TMP=/deepsquare/tmp"
  echo "DEEPSQUARE_SHARED_WORLD_TMP=/deepsquare/world-tmp"
  echo "DEEPSQUARE_DISK_TMP=/deepsquare/disk/tmp"
  echo "DEEPSQUARE_DISK_WORLD_TMP=/deepsquare/disk/world-tmp"
  echo "DEEPSQUARE_INPUT=/deepsquare/input"
  echo "DEEPSQUARE_OUTPUT=/deepsquare/output"
  echo "DEEPSQUARE_ENV=/deepsquare/$(basename $DEEPSQUARE_ENV)"
}

mounts() {
  echo "$STORAGE_PATH /deepsquare none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_SHARED_TMP /deepsquare/tmp none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_SHARED_WORLD_TMP /deepsquare/world-tmp none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_DISK_TMP /deepsquare/disk/tmp none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_DISK_WORLD_TMP /deepsquare/disk/world-tmp none x-create=dir,bind,rw"
  echo "/tmp/.X11-unix /tmp/.X11-unix none x-create=dir,bind,ro"
  echo '/host /container none x-create=auto,bind,ro'
}

hooks() {
  cat << 'EOFrclocal' > "${ENROOT_ROOTFS}/etc/rc.local"
test="$(cat << 'EOF[random_string]'
value
EOF[random_string]
)"
export test
mkdir -p '/dir' && cd '/dir' || { echo "change dir to working directory failed"; exit 1; }
exec "$@"
EOFrclocal
}
EOFenroot
/usr/bin/enroot start \
  --conf "$STORAGE_PATH/enroot-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID.conf" \
  "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID" \
  /bin/sh -c 'hostname'`,
			title: "Positive test with workdir",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			actual, err := renderer.RenderEnrootCommand(&tt.input)

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
