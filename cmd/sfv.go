package cmd

import (
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/autobrr/sfvbrr/internal/checksum"
	"github.com/spf13/cobra"
)

var (
	sfvWorkers    int
	sfvBufferSize int
	sfvVerbose    bool
	sfvQuiet      bool
	sfvRecursive  bool
	sfvCPUProfile string
)

var sfvCmd = &cobra.Command{
	Use:   "sfv [folder...]",
	Short: "Validate SFV CRC-32 checksums",
	Long: `Validate SFV (Simple File Verification) CRC-32 checksums for files in the specified folder(s).

The command will search for an SFV file (case insensitive) in each specified folder
and validate all files listed in the SFV file against their CRC-32 checksums.

When the recursive option (-r) is used, the command will search for SFV files in all
subdirectories of the specified folder(s).

Examples:
  # Validate a single folder
  sfvbrr sfv /path/to/release

  # Validate multiple folders
  sfvbrr sfv /path/to/release1 /path/to/release2

  # Validate recursively
  sfvbrr sfv -r /path/to/releases`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cleanup, err := setupProfiling(sfvCPUProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		defer cleanup()

		opts := checksum.Options{
			Workers:    sfvWorkers,
			BufferSize: sfvBufferSize,
			Verbose:    sfvVerbose,
			Quiet:      sfvQuiet,
			Recursive:  sfvRecursive,
		}

		if err := checksum.ValidateFolders(args, opts); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(sfvCmd)

	sfvCmd.Flags().IntVarP(&sfvWorkers, "workers", "w", 0, "Number of parallel workers (0 = auto-detect)")
	sfvCmd.Flags().IntVarP(&sfvBufferSize, "buffer-size", "b", 0, "Buffer size for file reading in bytes (0 = auto, default 64KB)")
	sfvCmd.Flags().BoolVarP(&sfvVerbose, "verbose", "v", false, "Show detailed validation results for each file")
	sfvCmd.Flags().BoolVarP(&sfvQuiet, "quiet", "q", false, "Quiet mode - only show errors")
	sfvCmd.Flags().BoolVarP(&sfvRecursive, "recursive", "r", false, "Recursively search for SFV files in subdirectories")
	sfvCmd.Flags().StringVar(&sfvCPUProfile, "cpuprofile", "", "Write CPU profile to file")
}

// setupProfiling sets up CPU profiling if the cpuprofile path is provided.
// It returns a cleanup function that should be deferred by the caller.
func setupProfiling(cpuprofile string) (func(), error) {
	if cpuprofile == "" {
		return func() {}, nil
	}

	f, err := os.Create(cpuprofile)
	if err != nil {
		return nil, fmt.Errorf("failed to create CPU profile file: %w", err)
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		f.Close()
		return nil, fmt.Errorf("failed to start CPU profile: %w", err)
	}

	return func() {
		pprof.StopCPUProfile()
		f.Close()
	}, nil
}
