package init

// init.go
// Eric Steadman - Copyright 2019
// shared init functions for all languages

import (
	"fmt"
	"io"
	"os"

	fp "path/filepath"

	"github.com/es3649/proj/internal/config"
	"github.com/es3649/proj/pkg/log"
	"github.com/sirupsen/logrus"
)

// projCommon starts the project:
// it begins by creating the .proj folder
// it returns a partially filled out config.ProjData
// object with the following fields:
// GitExists, Name, Root
func projCommon() (*config.ProjData, error) {
	// check is a project already exists
	path, err := config.LocateProject()

	// if we got an error that wasn't a "Proj doesn't exists"
	if err != nil && !config.IsEmptyProj(err) {
		// then that's a problem
		return nil, fmt.Errorf("Error during project lookup: %v", err)
	} else if err == nil {
		// then a project exists
		return nil, fmt.Errorf("Project already exists at: %s", path)
	}

	// then no project exists yet, create one
	err = os.Mkdir(".proj", 0755)
	if err != nil {
		return nil, fmt.Errorf("Failed to create a project data folder: %v", err)
	}

	f, err := os.Create("./.proj/projData.json")
	f.Chmod(0666)
	if err != nil {
		return nil, fmt.Errorf("Failed to create a project data file: %v", err)
	}

	// get project name
	hereAbs, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Failed to stat '.' : %v", err)
	}

	_, here := fp.Split(hereAbs)

	projData := &config.ProjData{
		GitExists: gitExists(hereAbs),
		Name:      here,
		Root:      hereAbs,
	}

	// create the .proj/projData.json file

	return projData, nil
}

// newFolder creates a folder with the given name, if it does not
// exist already
func newFolder(name string) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		if err = os.Mkdir(name, 0755); err != nil {
			log.Log.WithField("where", "init.populateCppProj").Warn(fmt.Sprintf("Failed to create `%s/'", name))
			log.Log.WithField("error", err).Info(fmt.Sprintf("Failed to create `%s/'", name))
		}
	} else {
		log.Log.WithFields(logrus.Fields{"where": "init.populateCppProj", "error": err}).Info(fmt.Sprintf("Not creating `%s/'. Already exists (or other error)", name))
	}
}

// copyResource copies a resource pointed to by pathToResource
// to destination
func copyResource(pathToResource, destination string) {
	// this should be similar to newFolder
	_, err := os.Stat(destination)
	if os.IsNotExist(err) {
		// open the resource
		r, err := os.Open(pathToResource)
		defer r.Close()
		if err != nil {
			log.Log.WithFields(logrus.Fields{"where": "init.copyResource", "error": err}).Info(fmt.Sprintf("Failed to open resource `%s/'", pathToResource))
			return
		}

		// create the resource here
		w, err := os.Create(destination)
		defer w.Close()
		if err != nil {
			log.Log.WithFields(logrus.Fields{"where": "init.copyResource", "error": err}).Info(fmt.Sprintf("Failed to create `%s/'", destination))
			return
		}

		// copy the resource over
		if _, err = io.Copy(w, r); err != nil {
			log.Log.WithFields(logrus.Fields{"where": "init.copyResource", "error": err}).Info(fmt.Sprintf("Failed to copy resource `%s/'", pathToResource))
		}

	} else {
		log.Log.WithFields(logrus.Fields{"where": "init.copyResource", "error": err}).Info(fmt.Sprintf("Not copying `%s/'. Already exists (or other error)", pathToResource))
	}
}
