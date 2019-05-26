package cmd

import (
	"github.com/spf13/cobra"

	"github.com/es3649/proj/internal/flags"
)

var initCmd = &cobra.Command{
	Use:               "init",
	Short:             "Initializes a workspace",
	PersistentPostRun: postRunInit,
}

func initInit() {
	// add subcommands to the root
	initCmd.AddCommand(initCppCmd)
	initCmd.AddCommand(initJavaCmd)
	// initCmd.AddCommand(initPythonCmd)
	// initCmd.AddCommand(initGoCmd)

	// configure the children
	initCpp()
	initJava()
	// initPython()
	// initGo()

	// add flags
	initCmd.Flags().BoolVarP(&flags.Init.IncludeGit, "git", "g", false, "initialize a git repository with this project")
	initCmd.Flags().BoolVarP(&flags.Init.Git.IsPrivate, "private", "p", false, "used with -g, indicates that the new repo should be private")
	initCmd.Flags().BoolVarP(&flags.Init.Git.OmitReadme, "omit_readme", "r", false, "used with -g, skips inclusion of a README file")
}

func postRunInit(_ *cobra.Command, args []string) {
	// initialize a git repository, including remote
	// COMMIT README ONLY!! The user may want to privatize the repo
	// we might be able to handle this^^ with flags

	// whether or not a git repository exists at this point
	// whether we created it, or it already existed
	// add .proj to the .gitignore
}
