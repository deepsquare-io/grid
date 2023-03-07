package renderer_test

import (
	"testing"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/renderer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func cleanStepWithRun(r *model.StepRun) *model.Step {
	return &model.Step{
		Name: "test",
		Run:  r,
	}
}

func cleanStepWithFor(f *model.StepFor) *model.Step {
	return &model.Step{
		Name: "test",
		For:  f,
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
/usr/bin/mkdir -p "$HOME/.config/enroot/"
/usr/bin/cat << 'EOFnetrc' > "$HOME/.config/enroot/.credentials"
machine "registry" login "username" password "password"
EOFnetrc
IMAGE_PATH="$STORAGE_PATH/$SLURM_JOB_ID-$(echo $RANDOM | md5sum | head -c 20).sqsh"
export IMAGE_PATH
/usr/bin/echo "Importing image..."
/usr/bin/enroot import -o "$IMAGE_PATH" -- 'docker://registry#image'
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
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' test='value' /usr/bin/srun --job-name='test' \
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
			input: *cleanStepWithFor(&cleanStepForItems),
			expected: `/usr/bin/echo 'Running: ''test'
doFor() {
export item="$1"
/usr/bin/echo 'Running: ''test'
/usr/bin/mkdir -p "$HOME/.config/enroot/"
/usr/bin/cat << 'EOFnetrc' > "$HOME/.config/enroot/.credentials"
machine "registry" login "username" password "password"
EOFnetrc
IMAGE_PATH="$STORAGE_PATH/$SLURM_JOB_ID-$(echo $RANDOM | md5sum | head -c 20).sqsh"
export IMAGE_PATH
/usr/bin/echo "Importing image..."
/usr/bin/enroot import -o "$IMAGE_PATH" -- 'docker://registry#image'
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
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' test='value' /usr/bin/srun --job-name='test' \
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
/usr/bin/mkdir -p "$HOME/.config/enroot/"
/usr/bin/cat << 'EOFnetrc' > "$HOME/.config/enroot/.credentials"
machine "registry" login "username" password "password"
EOFnetrc
IMAGE_PATH="$STORAGE_PATH/$SLURM_JOB_ID-$(echo $RANDOM | md5sum | head -c 20).sqsh"
export IMAGE_PATH
/usr/bin/echo "Importing image..."
/usr/bin/enroot import -o "$IMAGE_PATH" -- 'docker://registry#image'
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
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' test='value' /usr/bin/srun --job-name='test' \
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
				Name: "test",
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
