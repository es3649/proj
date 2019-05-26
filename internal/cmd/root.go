package cmd

import (
	"github.com/es3649/proj/pkg/log"
	"github.com/spf13/cobra"
)

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

var verbosity int

var rootCmd = &cobra.Command{
	Use:   "proj",
	Short: "command line tool for managing projects from the shell",
	Long: `The purpose of this tool is to enhance programming work
within the shell.`,
	PersistentPreRun: func(_ *cobra.Command, _ []string) {
		log.SetVerbosity(verbosity)
	},
	// Run: run,
}

func init() {
	// add and initialize subcommands
	rootCmd.AddCommand(initCmd)
	initInit()

	// rootCmd.AddCommand(classCmd)
	// initClass() this will create new class files

	rootCmd.AddCommand(versionCmd)

	// add flags for the root
	rootCmd.PersistentFlags().IntVarP(&verbosity, "verbosity", "v", 0, "verbosity level [less 0-4 loud]")
}

// func run(cmd *cobra.Command, args []string) {

// }
