package flags

import "github.com/es3649/proj/internal/git"

// Init holds all the command line flags associated with the init command
var Init = struct {
	IncludeGit bool
	Git        git.Config
}{}
