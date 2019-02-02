package cmd

import (
	"github.com/spf13/cobra"

	it "github.com/es3649/proj/internal/init"
)

var initJavaCmd = &cobra.Command{
	Use:   "java",
	Short: "Initializes a new java project",
	Long:  `Initializes a new java project with a Makefile and a path.sh modifier`,
	Run:   runJava,
}

func initJava() {

}

func runJava(cmd *cobra.Command, args []string) {
	it.JavaProj()
}
