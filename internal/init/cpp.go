package init

import "fmt"

// CppProj initializes a c++ project
func CppProj() {
	fmt.Println("Initializing cpp project")
	makeProj()
	// create the .proj folder
	// populate the directory with files from ~/.proj_config

	// Creates the foillowing within the project
	/*
		| bin/
		| cmd/
		`- | build.sh
		| data/
		| src/
		| Makefile
		| path.sh
	*/
}
