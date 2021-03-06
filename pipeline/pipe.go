package pipeline

import "github.com/goreleaser/releaser/config"

// Pipe interface
type Pipe interface {
	// Name of the pipe
	Name() string

	// Run the pipe
	Run(config config.ProjectConfig) error
}
