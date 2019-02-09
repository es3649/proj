package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	it "github.com/es3649/proj/internal/init"
	"github.com/es3649/proj/pkg/log"
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
	err := it.JavaProj()
	if err != nil {
		fmt.Println("Could not create project")
		log.Log.WithFields(logrus.Fields{"where": "cmd.runJava", "error": err}).Info("Failed to create java project")
	}
}
