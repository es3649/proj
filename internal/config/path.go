package config

import (
	"fmt"
	"os"
	"strings"

	fp "path/filepath"

	"github.com/es3649/proj/pkg/log"
)

// LoadProject walks up the directory tree until it finds the
// .proj file, then loads that data into a ProjData object
// and returns it
// if no project exists, we return no error and an empty
// ProjData object
func LoadProject() (*ProjData, error) {

	projLocation, err := LocateProject()
	// TODO possibly make this a little more descriptive
	if err != nil {
		return nil, err
	}

	fmt.Println(projLocation)

	// open and unmarshall the project metadata

	// check for error and return the ProjectData object

	// TODO stub return
	return nil, err
}

// IsEmptyProj determines whether the given error indicates
// the nonexistance of a project in this directory
func IsEmptyProj(err error) bool {
	if fmt.Sprintf("%v", err) == "No project exists" {
		return true
	}
	return false
}

// LocateProject looks up the directory tree until it finds
// a .proj folder. It returns the path of the first .proj
// folder it finds and an error od type "IsEmptyProj" if no
// .proj folder is found
func LocateProject() (string, error) {
	// check if the .proj folder is here
	path, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("Failed to get working directory: %v", err)
	}

	// stat .proj to see if it's here
	file, err := os.Stat(".proj")

	if err == nil {
		// if no error, then we found it
		return fp.Join(path, file.Name()), nil
	} else if err != nil && !os.IsNotExist(err) {
		// if some error besides a not exist, it's a legit error
		return "", fmt.Errorf("Failed to stat .proj locally: %v", err)
	} // else it's a IsNotExist error, whcih we kinda expected. Keep looking

	// split it into a list
	hierarchy := strings.Split(path, "/")
	// assign a position marker to the current directory
	pos := len(hierarchy) - 1

	// while the position marker is not 0
	fmt.Println(hierarchy)
	for ; pos > 0; pos-- {
		// descend
		err = os.Chdir("..")
		log.Log.WithField("path", "/"+fp.Join(hierarchy[:pos]...)).Debug("Searching for .proj")
		if err != nil {
			return "", fmt.Errorf("Failed to move to enclosing folder: %v", err)
		}

		// look for the .proj folder
		file, err := os.Stat(".proj")

		if err == nil {
			// if no error, then we found it
			return "/" + fp.Join(fp.Join(hierarchy[:pos]...), file.Name()), nil
			// if some error besides a not exist, it's a legit error
		} else if err != nil && !os.IsNotExist(err) {
			return "", fmt.Errorf("Failed to stat .proj at %s: %v", "/"+fp.Join(hierarchy[:pos]...), err)
		} // else it's a IsNotExist error, whcih we kinda expected. Keep looking

	}

	// chdir back to where we started
	err = os.Chdir(path)
	if err != nil {

	}

	return "", fmt.Errorf("No project exists")
}
