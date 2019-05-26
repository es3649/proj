package config

// ProjData holds configuration data for the project
type ProjData struct {
	GitExists bool   `json:"git_exists"`       // whether or not a git repository exists here
	Language  string `json:"project_language"` // language in which the project is written
	Name      string `json:"name"`             // name of the project
	Root      string `json:"root"`             // full path of the roof of this project
}

// Manifest holds information about the initial structure of
// a project. It contains a list of files to be copied, as well
// as directories to make
type Manifest struct {
	Folders   []string   `json:"folders"`
	CopyFiles []CopyFile `json:"copyFiles"`
}

// CopyFile represents a file to be copied, including
// the file mode of the destination files
type CopyFile struct {
	Name string `json:"name"`
	Mode int    `json:"mode"`
}
