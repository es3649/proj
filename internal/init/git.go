package init

// git.go - initializes git repository
// Eric Steadman - Copyright 2019

import (
	"os"
	fp "path/filepath"
)

// gitExists stats for a .git folder
func gitExists(base string) bool {
	_, err := os.Stat(fp.Join(base, ".git"))
	return err == nil
}

// GitInit initializes a git repository
func GitInit() {
	// check if a github repository already exists here
	// (Consider checking in lower directories as well,
	// not sure what we would do if we find one)

	// Create the repo locally

	// Create the remote repo

	// set the upstream for the local repo,
	// make an initial commit
	// push
}
