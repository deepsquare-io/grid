package renderer_test

import (
	"testing"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/renderer"
	"github.com/deepsquare-io/the-grid/sbatch-service/utils"
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
If you need a persistent cache, use the environment variable $DEEPSQUARE_TMP which is the cache location.
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
MOUNTS="$STORAGE_PATH:/deepsquare:rw,$DEEPSQUARE_TMP:/deepsquare/tmp:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' DEEPSQUARE_TMP='/deepsquare/tmp' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV="/deepsquare/$(basename $DEEPSQUARE_ENV)" test='value' /usr/bin/srun --job-name='test' \
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
If you need a persistent cache, use the environment variable $DEEPSQUARE_TMP which is the cache location.
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
MOUNTS="$STORAGE_PATH:/deepsquare:rw,$DEEPSQUARE_TMP:/deepsquare/tmp:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' DEEPSQUARE_TMP='/deepsquare/tmp' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV="/deepsquare/$(basename $DEEPSQUARE_ENV)" test='value' /usr/bin/srun --job-name='test' \
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
			title: "Positive test with items",
		},
		{
			input: cleanStepForRange,
			expected: `doFor() {
export index="$1"
/usr/bin/echo 'Running: ''test'
/usr/bin/cat << 'EOFmounterror'
WARNING: Mounts is now deprecated.
If you need a persistent cache, use the environment variable $DEEPSQUARE_TMP which is the cache location.
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
MOUNTS="$STORAGE_PATH:/deepsquare:rw,$DEEPSQUARE_TMP:/deepsquare/tmp:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' DEEPSQUARE_TMP='/deepsquare/tmp' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV="/deepsquare/$(basename $DEEPSQUARE_ENV)" test='value' /usr/bin/srun --job-name='test' \
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
  /bin/sh -c '/usr/bin/echo $index'
/usr/bin/echo 'Running: ''test'
/usr/bin/cat << 'EOFmounterror'
WARNING: Mounts is now deprecated.
If you need a persistent cache, use the environment variable $DEEPSQUARE_TMP which is the cache location.
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
MOUNTS="$STORAGE_PATH:/deepsquare:rw,$DEEPSQUARE_TMP:/deepsquare/tmp:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' DEEPSQUARE_TMP='/deepsquare/tmp' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV="/deepsquare/$(basename $DEEPSQUARE_ENV)" test='value' /usr/bin/srun --job-name='test' \
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
  /bin/sh -c '/usr/bin/echo $index'
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
