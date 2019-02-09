package init

import (
	"fmt"

	"github.com/es3649/proj/internal/config"
)

// JavaProj initializes a c++ project
// Creates the following within the project
// 	| bin/
// 	| cmd/
// 	| data/
// 	| src/
// 	| Makefile
// 	| path.sh
func JavaProj() error {
	fmt.Println("Initializing java project")
	projData, err := projCommon()
	if err != nil {
		return err
	}

	newFolder("bin")
	newFolder("cmd")
	newFolder("data")
	newFolder("src")

	copyResource("~/.proj_config/java/default_Makefile", "Makefile")
	copyResource("~/.proj_config/java/default_path.sh", "path.sh")

	projData.Language = "java"

	err = config.SaveProjData(projData)
	if err != nil {
		return err
	}

	return nil
}
