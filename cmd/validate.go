package cmd

import (
	"fmt"

	"github.com/moistari/rls"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate [folder]",
	Short: "Show the release type of a folder",
	Long:  `Detects and displays the type of scene release for the given folder name.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		folderName := args[0]

		// Parse the release using the rls library
		release := rls.ParseString(folderName)

		// Display the release type
		fmt.Printf("Release Type: %s\n", release.Type.String())
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
