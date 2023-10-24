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
	"github.com/deepsquare-io/grid/sbatch-service/utils/base58"
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
set +e
/usr/bin/enroot import -o "$IMAGE_PATH" -- 'docker://registry-1.docker.io#library/busybox:latest' &> "/tmp/enroot.import.$SLURM_JOB_ID.log"
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
MOUNTS="$STORAGE_PATH:/deepsquare:rw,$DEEPSQUARE_SHARED_TMP:/deepsquare/tmp:rw,$DEEPSQUARE_SHARED_WORLD_TMP:/deepsquare/world-tmp:rw,$DEEPSQUARE_DISK_TMP:/deepsquare/disk/tmp:rw,$DEEPSQUARE_DISK_WORLD_TMP:/deepsquare/disk/world-tmp:rw"
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' \
DEEPSQUARE_TMP='/deepsquare/tmp' \
DEEPSQUARE_SHARED_TMP='/deepsquare/tmp' \
DEEPSQUARE_SHARED_WORLD_TMP='/deepsquare/world-tmp' \
DEEPSQUARE_DISK_TMP='/deepsquare/disk/tmp' \
DEEPSQUARE_DISK_WORLD_TMP='/deepsquare/disk/world-tmp' \
DEEPSQUARE_INPUT='/deepsquare/input' \
DEEPSQUARE_OUTPUT='/deepsquare/output' \
DEEPSQUARE_ENV="/deepsquare/$(basename $DEEPSQUARE_ENV)" /usr/bin/srun --job-name='Say hello World"' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=4 \
  --mem-per-cpu=4096M \
  --gpus-per-task=1 \
  --ntasks=1 \
  --gpu-bind=none \
  --no-container-remap-root \
  --container-writable \
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
			actual, err := renderer.NewStepUseRenderer(base58.FakeEncoder{}).
				Render(&cleanJob, cleanStepWithUse(&tt.input), &tt.input)

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
