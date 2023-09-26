package module_test

import (
	"testing"

	"github.com/deepsquare-io/grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/grid/sbatch-service/module"
	"github.com/deepsquare-io/grid/sbatch-service/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRender(t *testing.T) {
	tests := []struct {
		input struct {
			j        *model.Job
			s        *model.Step
			template string
		}
		isError       bool
		errorContains []string
		expected      string
		title         string
	}{
		{
			title: "step of steps",
			input: struct {
				j        *model.Job
				s        *model.Step
				template string
			}{
				s: &model.Step{
					Use: &model.StepUse{
						Steps: []*model.Step{
							{
								Name: utils.Ptr("step by user"),
							},
							{
								Name: utils.Ptr("another step by user"),
							},
						},
					},
				},
				template: `# module.yaml
steps:
  - name: my step
  {{- .Step.Use.Steps | toYaml | nindent 2 }}
  - name: my other step
`,
			},
			expected: `# module.yaml
steps:
  - name: my step
  - name: step by user
  - name: another step by user
  - name: my other step
`,
		},
		{
			title: "step of steps: empty",
			input: struct {
				j        *model.Job
				s        *model.Step
				template string
			}{
				s: &model.Step{
					Use: &model.StepUse{},
				},
				template: `# module.yaml
steps:
  - name: my step
  {{- .Step.Use.Steps | toYaml | nindent 2 }}
  - name: my other step
`,
			},
			expected: "# module.yaml\nsteps:\n  - name: my step\n  \n  - name: my other step\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			actual, err := module.Render(tt.input.j, tt.input.s, tt.input.template)

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, actual)
			}
		})
	}
}
