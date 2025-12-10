package checksum

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FindSFVFiles finds all SFV files in the given directory (case insensitive)
func FindSFVFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var sfvFiles []string
	for _, entry := range entries {
		if !entry.IsDir() {
			filename := entry.Name()
			if strings.EqualFold(filepath.Ext(filename), ".sfv") {
				sfvFiles = append(sfvFiles, filepath.Join(dir, filename))
			}
		}
	}

	if len(sfvFiles) == 0 {
		return nil, fmt.Errorf("no SFV files found in directory: %s", dir)
	}

	return sfvFiles, nil
}

// FindSFVFilesRecursive finds all SFV files recursively in the given directory
func FindSFVFilesRecursive(dir string) ([]string, error) {
	var sfvFiles []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Continue on errors (e.g., permission denied)
			return nil
		}

		if !info.IsDir() {
			// Check if file has .sfv extension (case insensitive)
			if strings.EqualFold(filepath.Ext(path), ".sfv") {
				sfvFiles = append(sfvFiles, path)
			}
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %w", err)
	}

	return sfvFiles, nil
}

// validateSingleSFV validates a single SFV file and displays results
// Returns true if validation failed (has invalid or missing files)
func validateSingleSFV(sfvPath string, opts Options) (bool, error) {
	// Parse SFV file
	sfv, err := ParseSFVFile(sfvPath)
	if err != nil {
		return false, fmt.Errorf("failed to parse SFV file %s: %w", sfvPath, err)
	}

	// Validate SFV
	result, err := ValidateSFV(sfv, opts)
	if err != nil {
		return false, fmt.Errorf("failed to validate SFV: %w", err)
	}

	// Display results and return validation status
	failed := DisplayResult(result, opts)
	return failed, nil
}

// ValidateFolders validates SFV files in multiple folders
func ValidateFolders(folders []string, opts Options) error {
	var hasErrors bool

	for _, folder := range folders {
		// Resolve absolute path
		absPath, err := filepath.Abs(folder)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to resolve path %s: %v\n", folder, err)
			hasErrors = true
			continue
		}

		// Check if directory exists
		info, err := os.Stat(absPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s does not exist: %v\n", folder, err)
			hasErrors = true
			continue
		}

		if !info.IsDir() {
			fmt.Fprintf(os.Stderr, "Error: %s is not a directory\n", folder)
			hasErrors = true
			continue
		}

		if opts.Recursive {
			// Find all SFV files recursively
			sfvFiles, err := FindSFVFilesRecursive(absPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: failed to find SFV files recursively in %s: %v\n", folder, err)
				hasErrors = true
				continue
			}

			if len(sfvFiles) == 0 {
				if !opts.Quiet {
					fmt.Fprintf(os.Stderr, "No SFV files found in %s\n", folder)
				}
				hasErrors = true
				continue
			}

			// Validate each SFV file found
			for _, sfvPath := range sfvFiles {
				failed, err := validateSingleSFV(sfvPath, opts)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error: %v\n", err)
					hasErrors = true
				} else if failed {
					hasErrors = true
				}
			}
		} else {
			// Find all SFV files in current directory only
			sfvFiles, err := FindSFVFiles(absPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				hasErrors = true
				continue
			}

			// Validate each SFV file found
			for _, sfvPath := range sfvFiles {
				failed, err := validateSingleSFV(sfvPath, opts)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error: %v\n", err)
					hasErrors = true
				} else if failed {
					hasErrors = true
				}
			}
		}
	}

	if hasErrors {
		return fmt.Errorf("one or more folders had errors")
	}

	return nil
}

