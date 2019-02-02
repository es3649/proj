package init

import "fmt"

// JavaProj initializes a c++ project
func JavaProj() {
	fmt.Println("Initializing java project")
	makeProj()
	// create the .proj folder
	// populate the directory with files from ~/.proj_config

	// Creates the following within the project
	/*
		| bin/
		| cmd/
		| data/
		| src/
		| Makefile
		| path.sh
	*/
}
