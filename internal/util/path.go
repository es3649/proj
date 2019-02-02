package util

// ProjectData contains metadata about the current project
type ProjectData struct {
	Root string
	Language string
}

// LoadProject walks up the directory tree until it finds the
// .proj file, then loads that data into a ProjectData object
// and returns it
func LoadProject() *ProjectData {

	// TODO stub return
	return nil
}