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

func cleanStepAsyncLaunchWithStep(handleName string, s *model.Step) *model.StepAsyncLaunch {
	return &model.StepAsyncLaunch{
		HandleName: &handleName,
		Steps:      []*model.Step{s},
	}
}

func TestRenderStepAsyncLaunch(t *testing.T) {
	tests := []struct {
		input struct {
			Job             model.Job
			StepAsyncLaunch model.StepAsyncLaunch
		}
		isError       bool
		errorContains []string
		expected      string
		title         string
	}{
		{
			input: struct {
				Job             model.Job
				StepAsyncLaunch model.StepAsyncLaunch
			}{
				Job: model.Job{Resources: cleanJob.Resources},
				StepAsyncLaunch: *cleanStepAsyncLaunchWithStep("async_launch", &model.Step{
					Name: utils.Ptr("test"),
					Launch: cleanStepAsyncLaunchWithStep("async_sub", &model.Step{
						Name: utils.Ptr("subtest"),
					}),
				}),
			},
			expected: `(
declare -A EXIT_SIGNALS
/usr/bin/echo 'Running: ''test'
(
declare -A EXIT_SIGNALS
/usr/bin/echo 'Running: ''subtest'

for pid in "${!EXIT_SIGNALS[@]}"; do
  /usr/bin/echo "Process $$ sending signal ${EXIT_SIGNALS[$pid]} to $pid"
  kill -s "${EXIT_SIGNALS[$pid]}" "$pid" || echo "Sending signal ${EXIT_SIGNALS[$pid]} to $pid failed, continuing..."
done
) &
asynclaunchpid="$!"
export PID_ASYNC_SUB="$asynclaunchpid"
EXIT_SIGNALS[$asynclaunchpid]=15

for pid in "${!EXIT_SIGNALS[@]}"; do
  /usr/bin/echo "Process $$ sending signal ${EXIT_SIGNALS[$pid]} to $pid"
  kill -s "${EXIT_SIGNALS[$pid]}" "$pid" || echo "Sending signal ${EXIT_SIGNALS[$pid]} to $pid failed, continuing..."
done
) &
asynclaunchpid="$!"
export PID_ASYNC_LAUNCH="$asynclaunchpid"
EXIT_SIGNALS[$asynclaunchpid]=15`,
			title: "Positive test",
		},
		{
			input: struct {
				Job             model.Job
				StepAsyncLaunch model.StepAsyncLaunch
			}{
				Job: model.Job{Resources: cleanJob.Resources},
				StepAsyncLaunch: func() model.StepAsyncLaunch {
					m := *cleanStepAsyncLaunchWithStep("async_launch", &model.Step{
						Name: utils.Ptr("test"),
						Launch: cleanStepAsyncLaunchWithStep("async_sub", &model.Step{
							Name: utils.Ptr("subtest"),
						}),
					})
					m.SignalOnParentStepExit = utils.Ptr(9)
					return m
				}(),
			},
			expected: `(
declare -A EXIT_SIGNALS
/usr/bin/echo 'Running: ''test'
(
declare -A EXIT_SIGNALS
/usr/bin/echo 'Running: ''subtest'

for pid in "${!EXIT_SIGNALS[@]}"; do
  /usr/bin/echo "Process $$ sending signal ${EXIT_SIGNALS[$pid]} to $pid"
  kill -s "${EXIT_SIGNALS[$pid]}" "$pid" || echo "Sending signal ${EXIT_SIGNALS[$pid]} to $pid failed, continuing..."
done
) &
asynclaunchpid="$!"
export PID_ASYNC_SUB="$asynclaunchpid"
EXIT_SIGNALS[$asynclaunchpid]=15

for pid in "${!EXIT_SIGNALS[@]}"; do
  /usr/bin/echo "Process $$ sending signal ${EXIT_SIGNALS[$pid]} to $pid"
  kill -s "${EXIT_SIGNALS[$pid]}" "$pid" || echo "Sending signal ${EXIT_SIGNALS[$pid]} to $pid failed, continuing..."
done
) &
asynclaunchpid="$!"
export PID_ASYNC_LAUNCH="$asynclaunchpid"
EXIT_SIGNALS[$asynclaunchpid]=9`,
			title: "Positive test: explicit signal",
		},
		{
			input: struct {
				Job             model.Job
				StepAsyncLaunch model.StepAsyncLaunch
			}{
				Job: model.Job{Resources: cleanJob.Resources},
				StepAsyncLaunch: func() model.StepAsyncLaunch {
					m := *cleanStepAsyncLaunchWithStep("async_launch", &model.Step{
						Name: utils.Ptr("test"),
						Launch: cleanStepAsyncLaunchWithStep("async_sub", &model.Step{
							Name: utils.Ptr("subtest"),
						}),
					})
					m.SignalOnParentStepExit = utils.Ptr(0)
					return m
				}(),
			},
			expected: `(
declare -A EXIT_SIGNALS
/usr/bin/echo 'Running: ''test'
(
declare -A EXIT_SIGNALS
/usr/bin/echo 'Running: ''subtest'

for pid in "${!EXIT_SIGNALS[@]}"; do
  /usr/bin/echo "Process $$ sending signal ${EXIT_SIGNALS[$pid]} to $pid"
  kill -s "${EXIT_SIGNALS[$pid]}" "$pid" || echo "Sending signal ${EXIT_SIGNALS[$pid]} to $pid failed, continuing..."
done
) &
asynclaunchpid="$!"
export PID_ASYNC_SUB="$asynclaunchpid"
EXIT_SIGNALS[$asynclaunchpid]=15

for pid in "${!EXIT_SIGNALS[@]}"; do
  /usr/bin/echo "Process $$ sending signal ${EXIT_SIGNALS[$pid]} to $pid"
  kill -s "${EXIT_SIGNALS[$pid]}" "$pid" || echo "Sending signal ${EXIT_SIGNALS[$pid]} to $pid failed, continuing..."
done
) &
asynclaunchpid="$!"
export PID_ASYNC_LAUNCH="$asynclaunchpid"`,
			title: "Positive test: explicit unsafe",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			actual, err := renderer.RenderStepAsyncLaunch(&tt.input.Job, &tt.input.StepAsyncLaunch)
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
