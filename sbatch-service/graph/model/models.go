package model

import (
	"github.com/deepsquare-io/the-grid/sbatch-service/validate"
)

// ForRange describes the parameter for a range loop.
type ForRange struct {
	Begin     int `json:"Begin"`
	End       int `json:"End"`
	Increment int `json:"Increment"`
}

func (s *ForRange) Validate() error {
	return validate.I.Struct(s)
}

// A Job is a finite sequence of instructions.
type Job struct {
	Steps []*Step `json:"steps" validate:"dive"`
}

func (s *Job) Validate() error {
	return validate.I.Struct(s)
}

// Resources are the allocated resources for a command.
type Resources struct {
	// Number of tasks which are run in parallel.
	//
	// Can be greater or equal to 1.
	Tasks int `json:"tasks" validate:"gte=1"`
	// Allocated CPUs per task.
	//
	// Can be greater or equal to 1.
	CpusPerTask int `json:"cpusPerTask" validate:"gte=1"`
	// Allocated memory (MB) per task.
	//
	// Can be greater or equal to 1.
	MemPerCPU int `json:"memPerCpu" validate:"gte=1"`
	// Allocated GPUs per task.
	//
	// Can be greater or equal to 0.
	GpusPerTask int `json:"gpusPerTask" validate:"gte=0"`
}

func (s *Resources) Validate() error {
	return validate.I.Struct(s)
}

// Step is one instruction.
type Step struct {
	// Name of the instruction.
	Name string `json:"name"`
	// Run a command if not null.
	//
	// Is exclusive with "for".
	Run *StepRun `json:"run"`
	// Run a for loop if not null.
	//
	// Is exclusive with "run".
	For *StepFor `json:"for"`
}

func (s *Step) Validate() error {
	return validate.I.Struct(s)
}

// StepFor describes a for loop.
type StepFor struct {
	// Do a parallel for loop. Each iteration is run in parallel.
	Parallel bool `json:"parallel"`
	// Item accessible via the {{ .Item }} variable. Index accessible via the {{ .Index }} variable.
	//
	// Exclusive with "range".
	Items []string `json:"items"`
	// Index accessible via the {{ .Index }} variable.
	//
	// Exclusive with "items".
	Range *ForRange `json:"range"`
	// Steps are run sequentially in one iteration.
	Steps []*Step `json:"steps"`
}

func (s *StepFor) Validate() error {
	return validate.I.Struct(s)
}

// StepRun is one script executed with the shell.
//
// Shared storage is accessible through the $STORAGE_PATH environment variable.
//
// echo "KEY=value" >> "$STORAGE_PATH/env" can be used to share environment variables.
type StepRun struct {
	// Allocated resources for the command.
	Resources *Resources `json:"resources"`
	// Run the command inside a container.
	//
	// Format [user:password]@<host>#<image>:<tag>.
	// If null or empty, run on the host.
	Image *string `json:"image" validate:"omitempty,valid_container_image_url"`
	// Command specifies a shell script.
	Command string `json:"command"`
}

func (s *StepRun) Validate() error {
	return validate.I.Struct(s)
}
