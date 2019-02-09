package config

// ProjData holds configuration data for the project
type ProjData struct {
	GitExists bool   // whether or not a git repository exists here
	Language  string // language in which the project is written
	Name      string // name of the project
	Root      string // full path of the roof of this project
}
