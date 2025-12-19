package cmd

import (
	"fmt"
	"os"

	"github.com/autobrr/sfvbrr/internal/checksum"
	"github.com/spf13/cobra"
)

var (
	zipWorkers     int
	zipBufferSize  int
	zipVerbose     bool
	zipQuiet       bool
	zipRecursive   bool
	zipCPUProfile  string
	zipOutputJSON  bool
	zipOutputYAML  bool
)

var zipCmd = &cobra.Command{
	Use:   "zip [folder...]",
	Short: "Validate ZIP file integrity",
	Long: `Validate ZIP file integrity by testing all entries in ZIP files (equivalent to zip -T).

The command will search for ZIP files (case insensitive) in each specified folder
and validate all entries in each ZIP file by reading them and verifying their CRC-32 checksums.

When the recursive option (-r) is used, the command will search for ZIP files in all
subdirectories of the specified folder(s).

Examples:
  # Validate ZIP files in a single folder
  sfvbrr zip /path/to/release

  # Validate ZIP files in multiple folders
  sfvbrr zip /path/to/release1 /path/to/release2

  # Validate ZIP files recursively
  sfvbrr zip -r /path/to/releases`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cleanup, err := setupProfiling(zipCPUProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		defer cleanup()

		outputFormat := checksum.OutputFormatText
		if zipOutputJSON {
			outputFormat = checksum.OutputFormatJSON
		} else if zipOutputYAML {
			outputFormat = checksum.OutputFormatYAML
		}

		opts := checksum.Options{
			Workers:      zipWorkers,
			BufferSize:   zipBufferSize,
			Verbose:      zipVerbose,
			Quiet:        zipQuiet,
			Recursive:    zipRecursive,
			OutputFormat: outputFormat,
		}

		if err := checksum.ValidateZIPFolders(args, opts); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(zipCmd)

	zipCmd.Flags().IntVarP(&zipWorkers, "workers", "w", 0, "Number of parallel workers (0 = auto-detect)")
	zipCmd.Flags().IntVarP(&zipBufferSize, "buffer-size", "b", 0, "Buffer size for file reading in bytes (0 = auto, default 64KB)")
	zipCmd.Flags().BoolVarP(&zipVerbose, "verbose", "v", false, "Show detailed validation results for each entry")
	zipCmd.Flags().BoolVarP(&zipQuiet, "quiet", "q", false, "Quiet mode - only show errors")
	zipCmd.Flags().BoolVarP(&zipRecursive, "recursive", "r", false, "Recursively search for ZIP files in subdirectories")
	zipCmd.Flags().StringVar(&zipCPUProfile, "cpuprofile", "", "Write CPU profile to file")
	zipCmd.Flags().BoolVar(&zipOutputJSON, "json", false, "Output results in JSON format")
	zipCmd.Flags().BoolVar(&zipOutputYAML, "yaml", false, "Output results in YAML format")
	zipCmd.MarkFlagsMutuallyExclusive("json", "yaml")
}
