package init

// git.go - initialized git repository
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

}
