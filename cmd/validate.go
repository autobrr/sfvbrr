package cmd

import (
	"fmt"
	"os"

	"github.com/autobrr/sfvbrr/internal/validate"
	"github.com/spf13/cobra"
)

var (
	validatePresetPath        string
	validateVerbose           bool
	validateQuiet             bool
	validateRecursive         bool
	validateOverwriteCategory string
	validateCPUProfile        string
)

var validateCmd = &cobra.Command{
	Use:   "validate [folder...]",
	Short: "Validate scene release folders",
	Long: `Validate scene release folders against category-specific rules.

The command detects the release category from the folder name and validates
the folder contents against the rules defined in the preset configuration file.

When the recursive option (-r) is used, the command will search for valid
release folders in all subdirectories of the specified folder(s).

The --overwrite flag allows you to bypass automatic category detection and
manually specify a category for validation.

Examples:
  # Validate a single folder
  sfvbrr validate /path/to/release

  # Validate multiple folders
  sfvbrr validate /path/to/release1 /path/to/release2

  # Validate recursively
  sfvbrr validate -r /path/to/releases

  # Override category detection
  sfvbrr validate --overwrite app /path/to/release`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cleanup, err := setupProfiling(validateCPUProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		defer cleanup()

		opts := validate.Options{
			PresetPath:        validatePresetPath,
			Verbose:           validateVerbose,
			Quiet:             validateQuiet,
			Recursive:         validateRecursive,
			OverwriteCategory: validateOverwriteCategory,
		}

		if err := validate.ValidateFolders(args, opts); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	validateCmd.Flags().StringVarP(&validatePresetPath, "preset", "p", "", "Path to preset YAML file (default: auto-detect)")
	validateCmd.Flags().BoolVarP(&validateVerbose, "verbose", "v", false, "Show detailed validation results for each rule")
	validateCmd.Flags().BoolVarP(&validateQuiet, "quiet", "q", false, "Quiet mode - only show errors")
	validateCmd.Flags().BoolVarP(&validateRecursive, "recursive", "r", false, "Recursively search for release folders in subdirectories")
	validateCmd.Flags().StringVar(&validateOverwriteCategory, "overwrite", "", "Override category detection with specified category (bypasses automatic detection)")
	validateCmd.Flags().StringVar(&validateCPUProfile, "cpuprofile", "", "Write CPU profile to file")
}
