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
				Name: "test",
				Run:  cleanStepRun("/usr/bin/echo $item"),
			},
			{
				Name: "test",
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
				Name: "test",
				Run:  cleanStepRun("/usr/bin/echo $index"),
			},
			{
				Name: "test",
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
/usr/bin/mkdir -p "$HOME/.config/enroot/"
/usr/bin/cat << 'EOFnetrc' > "$HOME/.config/enroot/.credentials"
machine "registry" login "username" password "password"
EOFnetrc
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' test='value' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  --no-container-remap-root \
  --container-mounts="${MOUNTS}" \
  --container-image='registry#image' \
  /bin/sh -c '/usr/bin/echo $item'
/usr/bin/echo 'Running: ''test'
/usr/bin/mkdir -p "$HOME/.config/enroot/"
/usr/bin/cat << 'EOFnetrc' > "$HOME/.config/enroot/.credentials"
machine "registry" login "username" password "password"
EOFnetrc
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' test='value' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  --no-container-remap-root \
  --container-mounts="${MOUNTS}" \
  --container-image='registry#image' \
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
/usr/bin/mkdir -p "$HOME/.config/enroot/"
/usr/bin/cat << 'EOFnetrc' > "$HOME/.config/enroot/.credentials"
machine "registry" login "username" password "password"
EOFnetrc
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' test='value' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  --no-container-remap-root \
  --container-mounts="${MOUNTS}" \
  --container-image='registry#image' \
  /bin/sh -c '/usr/bin/echo $index'
/usr/bin/echo 'Running: ''test'
/usr/bin/mkdir -p "$HOME/.config/enroot/"
/usr/bin/cat << 'EOFnetrc' > "$HOME/.config/enroot/.credentials"
machine "registry" login "username" password "password"
EOFnetrc
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' test='value' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  --no-container-remap-root \
  --container-mounts="${MOUNTS}" \
  --container-image='registry#image' \
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
