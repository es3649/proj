package config

import (
	"fmt"
	"os"
)

func locateRoot() (string, error) {

	// check if the .proj is here
	here, err := os.Stat(".proj")

	// TODO fix this up
	if err != nil {
		if !os.IsNotExist(err) {
			// if we got some error besides "file doesn't exist"
			return "", fmt.Errorf("Error searching for project root: %s", err)
		}
		// we kinda expect it to not exist
	} else {
		// we found it
		return here.Name(), nil
	}

	initialRoot, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("Error fetching working directory: %s", err)
	}

	// TODO descend here
	// we need a way to stop descending once we get to root.
	// we look for .proj at each level

	err = os.Chdir(initialRoot)
	if err != nil {
		return "", fmt.Errorf("Failed to return to the initial directory: %s", err)
	}

	return "", nil
}

// LoadConfig loads the configuration files for the projects
// they are marshalled into a ProjData object
func LoadConfig() (*ProjData, error) {

	return nil, nil
}
