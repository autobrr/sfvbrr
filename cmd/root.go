package cmd

import (
	"github.com/spf13/cobra"
)

const banner = `
  ________________________   __________________________________
 /   _____/\_   _____/\   \ /   /\______   \______   \______   \
 \_____  \  |    __)   \   Y   /  |    |  _/|       _/|       _/
 /        \ |     \     \     /   |    |   \|    |   \|    |   \
/_______  / \___  /      \___/    |______  /|____|_  /|____|_  /
        \/      \/                       \/        \/        \/
`

var rootCmd = &cobra.Command{
	Use:   "sfvbrr",
	Short: "Scene release validation tool",
	Long:  banner + "sfvbrr is a high-performance scene release validation tool.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add subcommands here
}
