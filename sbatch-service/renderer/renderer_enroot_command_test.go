package renderer_test

import (
	"testing"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/renderer"
	"github.com/deepsquare-io/the-grid/sbatch-service/utils"
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
			expected: `/usr/bin/cat << 'EOFenroot' > "$STORAGE_PATH/enroot.conf"
#ENROOT_REMAP_ROOT=n
#ENROOT_ROOTFS_WRITABLE=y
#ENROOT_MOUNT_HOME=y

environ() {
  # Keep all the environment from the host
  /usr/bin/env

  /usr/bin/cat "${ENROOT_ROOTFS}/etc/environment"

  /usr/bin/echo "STORAGE_PATH=/deepsquare"
  /usr/bin/echo "DEEPSQUARE_INPUT=/deepsquare/input"
  /usr/bin/echo "DEEPSQUARE_OUTPUT=/deepsquare/output"
  /usr/bin/echo "DEEPSQUARE_ENV=/deepsquare/env"
  /usr/bin/echo "test='value'"
}

mounts() {
  /usr/bin/echo "$STORAGE_PATH /deepsquare none x-create=dir,bind,rw"
  /usr/bin/echo "/tmp/.X11-unix /tmp/.X11-unix none x-create=dir,bind,ro"
  /usr/bin/echo '/host /container bind,ro'
}

hooks() {
  /usr/bin/cat << 'EOFrclocal' > "${ENROOT_ROOTFS}/etc/rc.local"
  exec "$@"
  EOFrclocal
}
EOFenroot
/usr/bin/enroot start \
  --conf "$STORAGE_PATH/enroot.conf" \
  "container-$SLURM_JOB_ID" \
  /bin/sh -c 'hostname'`,
			title: "Positive test with enroot image",
		},
		{
			input: func() model.StepRun {
				r := *cleanEnrootStepRun("hostname")
				r.MapRoot = utils.Ptr(true)
				return r
			}(),
			expected: `/usr/bin/cat << 'EOFenroot' > "$STORAGE_PATH/enroot.conf"
#ENROOT_REMAP_ROOT=y
#ENROOT_ROOTFS_WRITABLE=y
#ENROOT_MOUNT_HOME=y

environ() {
  # Keep all the environment from the host
  /usr/bin/env

  /usr/bin/cat "${ENROOT_ROOTFS}/etc/environment"

  /usr/bin/echo "STORAGE_PATH=/deepsquare"
  /usr/bin/echo "DEEPSQUARE_INPUT=/deepsquare/input"
  /usr/bin/echo "DEEPSQUARE_OUTPUT=/deepsquare/output"
  /usr/bin/echo "DEEPSQUARE_ENV=/deepsquare/env"
  /usr/bin/echo "test='value'"
}

mounts() {
  /usr/bin/echo "$STORAGE_PATH /deepsquare none x-create=dir,bind,rw"
  /usr/bin/echo "/tmp/.X11-unix /tmp/.X11-unix none x-create=dir,bind,ro"
  /usr/bin/echo '/host /container bind,ro'
}

hooks() {
  /usr/bin/cat << 'EOFrclocal' > "${ENROOT_ROOTFS}/etc/rc.local"
  exec "$@"
  EOFrclocal
}
EOFenroot
/usr/bin/enroot start \
  --conf "$STORAGE_PATH/enroot.conf" \
  "container-$SLURM_JOB_ID" \
  /bin/sh -c 'hostname'`,
			title: "Positive test with maproot",
		},
		{
			input: func() model.StepRun {
				r := *cleanEnrootStepRun("hostname")
				r.Container.Image = "/test/my.sqshfs"
				return r
			}(),
			expected: `/usr/bin/cat << 'EOFenroot' > "$STORAGE_PATH/enroot.conf"
#ENROOT_REMAP_ROOT=n
#ENROOT_ROOTFS_WRITABLE=y
#ENROOT_MOUNT_HOME=y

environ() {
  # Keep all the environment from the host
  /usr/bin/env

  /usr/bin/cat "${ENROOT_ROOTFS}/etc/environment"

  /usr/bin/echo "STORAGE_PATH=/deepsquare"
  /usr/bin/echo "DEEPSQUARE_INPUT=/deepsquare/input"
  /usr/bin/echo "DEEPSQUARE_OUTPUT=/deepsquare/output"
  /usr/bin/echo "DEEPSQUARE_ENV=/deepsquare/env"
  /usr/bin/echo "test='value'"
}

mounts() {
  /usr/bin/echo "$STORAGE_PATH /deepsquare none x-create=dir,bind,rw"
  /usr/bin/echo "/tmp/.X11-unix /tmp/.X11-unix none x-create=dir,bind,ro"
  /usr/bin/echo '/host /container bind,ro'
}

hooks() {
  /usr/bin/cat << 'EOFrclocal' > "${ENROOT_ROOTFS}/etc/rc.local"
  exec "$@"
  EOFrclocal
}
EOFenroot
/usr/bin/enroot start \
  --conf "$STORAGE_PATH/enroot.conf" \
  "container-$SLURM_JOB_ID" \
  /bin/sh -c 'hostname'`,
			title: "Positive test with enroot absolute path image",
		},
		{
			input: func() model.StepRun {
				r := *cleanEnrootStepRun("hostname")
				r.WorkDir = utils.Ptr("/dir")
				return r
			}(),
			expected: `/usr/bin/cat << 'EOFenroot' > "$STORAGE_PATH/enroot.conf"
#ENROOT_REMAP_ROOT=n
#ENROOT_ROOTFS_WRITABLE=y
#ENROOT_MOUNT_HOME=y

environ() {
  # Keep all the environment from the host
  /usr/bin/env

  /usr/bin/cat "${ENROOT_ROOTFS}/etc/environment"

  /usr/bin/echo "STORAGE_PATH=/deepsquare"
  /usr/bin/echo "DEEPSQUARE_INPUT=/deepsquare/input"
  /usr/bin/echo "DEEPSQUARE_OUTPUT=/deepsquare/output"
  /usr/bin/echo "DEEPSQUARE_ENV=/deepsquare/env"
  /usr/bin/echo "test='value'"
}

mounts() {
  /usr/bin/echo "$STORAGE_PATH /deepsquare none x-create=dir,bind,rw"
  /usr/bin/echo "/tmp/.X11-unix /tmp/.X11-unix none x-create=dir,bind,ro"
  /usr/bin/echo '/host /container bind,ro'
}

hooks() {
  /usr/bin/cat << 'EOFrclocal' > "${ENROOT_ROOTFS}/etc/rc.local"
  cd '/dir' || { echo "change dir to working directory failed"; exit 1; }
  exec "$@"
  EOFrclocal
}
EOFenroot
/usr/bin/enroot start \
  --conf "$STORAGE_PATH/enroot.conf" \
  "container-$SLURM_JOB_ID" \
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