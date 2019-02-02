package init

import (
	"os"

	"github.com/es3649/proj/internal/config"
)

// makeProj creates the
func makeProj() (*config.ProjData, error) {
	err := os.Mkdir(".proj", 755)

	if err != nil {
		return nil, err
	}
	// project name

	// save metadata about the project
	// language

	return nil, nil
}
