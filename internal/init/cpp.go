package init

import (
	"fmt"

	"github.com/es3649/proj/internal/config"
	"github.com/es3649/proj/pkg/log"
	"github.com/sirupsen/logrus"
)

// CppProj initializes a c++ project
func CppProj() error {
	fmt.Println("Initializing cpp project")
	projData, err := projCommon()
	if err != nil {
		return err
	}
	projData.Language = "c++"

	log.Log.WithFields(logrus.Fields{"where": "makeProj", "projData": fmt.Sprintf("%#v", *projData)}).Info("Initializing project")

	// populate the directory with files from ~/.proj_config
	err = populateCppProj()
	if err != nil {

	}

	err = config.SaveProjData(projData)
	if err != nil {

	}

	return nil
}

// populateCppProj populates the project with files if they don't
// already exist
/*
	| bin/
	| cmd/
	`- | build.sh
	| data/
	| src/
	| Makefile
	| path.sh
*/
func populateCppProj() error {
	// create each of bin, cmd, data, src
	newFolder("bin")
	newFolder("cmd")
	newFolder("data")
	newFolder("src")

	// copy the c++ resources over

	return nil
}
