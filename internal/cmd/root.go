package cmd

import "github.com/spf13/cobra"

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

var rootCmd = &cobra.Command {
	Use: "proj",
	Short: "command line tool for managing projects from the shell",
	Long: `The purpose of this tool is to enhance programming work
within the shell.`,
	// Run: run,
}

func init() {
	// add and initialize subcommands
	rootCmd.AddCommand(initCmd)
	initInit()

	// initClass() this will create new class files

	// add flags for the root
}

// func run(cmd *cobra.Command, args []string) {

// }