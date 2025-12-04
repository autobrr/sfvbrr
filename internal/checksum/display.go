package checksum

import (
	"fmt"
	"os"
	"strings"
)

// DisplayResult displays the validation results to the user
// Returns true if validation failed (has invalid or missing files)
func DisplayResult(result *ValidationResult, opts Options) bool {
	if opts.Quiet {
		// In quiet mode, only show summary if there are errors
		if result.InvalidFiles > 0 || result.MissingFiles > 0 {
			fmt.Fprintf(os.Stderr, "%s: %d invalid, %d missing\n",
				result.SFVFile.Path,
				result.InvalidFiles,
				result.MissingFiles)
		}
		return result.InvalidFiles > 0 || result.MissingFiles > 0
	}

	// Show SFV file path
	fmt.Printf("Validating SFV: %s\n", result.SFVFile.Path)
	fmt.Printf("Total files: %d\n\n", result.TotalFiles)

	// Show individual results if verbose
	if opts.Verbose {
		for _, res := range result.Results {
			if res.Valid {
				fmt.Printf("✓ %s\n", res.Entry.Filename)
			} else {
				if res.Error != nil {
					if strings.Contains(res.Error.Error(), "file not found") {
						fmt.Printf("✗ %s (MISSING)\n", res.Entry.Filename)
					} else {
						fmt.Printf("✗ %s (%s)\n", res.Entry.Filename, res.Error.Error())
					}
				}
			}
		}
		fmt.Println()
	}

	// Show summary
	fmt.Printf("Summary:\n")
	fmt.Printf("  Valid:   %d\n", result.ValidFiles)
	if result.InvalidFiles > 0 {
		fmt.Printf("  Invalid: %d\n", result.InvalidFiles)
	}
	if result.MissingFiles > 0 {
		fmt.Printf("  Missing: %d\n", result.MissingFiles)
	}

	// Return true if validation failed
	return result.InvalidFiles > 0 || result.MissingFiles > 0
}
