package renderer_test

import (
	"fmt"
	"testing"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/renderer"
	"github.com/deepsquare-io/the-grid/sbatch-service/utils"
	"github.com/deepsquare-io/the-grid/sbatch-service/utils/base58"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var cleanStepUse = model.StepUse{
	Source: "github.com/deepsquare-io/workflow-module-example@6abe5d5",
	Args: []*model.EnvVar{
		{
			Key:   "WHO",
			Value: "me",
		},
	},
	ExportEnvAs: utils.Ptr("HELLO_WORLD"),
}

func TestStepUse(t *testing.T) {
	tests := []struct {
		input         model.StepUse
		isError       bool
		errorContains []string
		expected      string
		title         string
	}{
		{
			input: cleanStepUse,
			expected: `(
export WHO='World'
export WHO='me'
DEEPSQUARE_BhNwmz1fC9zVZ8im94bLbw_OLD_ENV="$DEEPSQUARE_ENV"
export DEEPSQUARE_ENV="$STORAGE_PATH/DEEPSQUARE_BhNwmz1fC9zVZ8im94bLbw_env"
/usr/bin/touch $DEEPSQUARE_ENV
/usr/bin/echo 'Running: ''Say hello World"'

IMAGE_PATH="$STORAGE_PATH/$SLURM_JOB_ID-$(echo $RANDOM | md5sum | head -c 20).sqsh"
export IMAGE_PATH
/usr/bin/echo "Importing image..."
/usr/bin/enroot import -o "$IMAGE_PATH" -- 'docker://registry-1.docker.io#library/busybox:latest' > /dev/null
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
MOUNTS="$STORAGE_PATH:/deepsquare:rw"
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV="/deepsquare/$(basename $DEEPSQUARE_ENV)" /usr/bin/srun --job-name='Say hello World"' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=4 \
  --mem-per-cpu=4096M \
  --gpus-per-task=1 \
  --ntasks=1 \
  --gpu-bind=none \
  --no-container-remap-root \
  --no-container-mount-home \
  --container-mounts="${MOUNTS}" \
  --container-workdir=/deepsquare \
  --container-image="$IMAGE_PATH" \
  /bin/sh -c 'echo "Hello ${WHO}"
echo "RESULT=Hello ${WHO}" >> ${DEEPSQUARE_ENV}
'
echo "HELLO_WORLD_$(/usr/bin/grep "^RESULT" "$DEEPSQUARE_ENV")" >> "${DEEPSQUARE_BhNwmz1fC9zVZ8im94bLbw_OLD_ENV}"
export DEEPSQUARE_ENV="${DEEPSQUARE_BhNwmz1fC9zVZ8im94bLbw_OLD_ENV}"
)
export "HELLO_WORLD_$(/usr/bin/grep "^RESULT" "$DEEPSQUARE_ENV")"
`,
			title: "Positive test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			actual, err := renderer.NewStepUseRenderer(base58.FakeEncoder{}).Render(&cleanJob, cleanStepWithUse(&tt.input), &tt.input)
			fmt.Println(actual)

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
