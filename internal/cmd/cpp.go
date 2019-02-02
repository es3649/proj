package cmd

import (
	"github.com/spf13/cobra"

	it "github.com/es3649/proj/internal/init"
)

var initCppCmd = &cobra.Command{
	Use:   "cpp",
	Short: "Initializes a new c++ project",
	Long:  `Initializes a new c++ project with a Makefile (and some accompanying shell scripts in a cmd folder)`,
	Run:   runCpp,
}

func initCpp() {

}

func runCpp(cmd *cobra.Command, args []string) {
	it.CppProj()
}
