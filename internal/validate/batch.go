package validate

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/autobrr/sfvbrr/internal/preset"
)

// FindFoldersRecursive finds all folders recursively in the given directory
// that can be validated (i.e., have a detectable category)
func FindFoldersRecursive(dir string) ([]string, error) {
	var folders []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Continue on errors (e.g., permission denied)
			return nil
		}

		if info.IsDir() {
			// Try to detect category for this folder
			category, err := DetectCategory(path)
			if err == nil && category != "" {
				folders = append(folders, path)
			}
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %w", err)
	}

	return folders, nil
}

// validateSingleFolder validates a single folder and displays results
func validateSingleFolder(folderPath string, presetConfig *preset.PresetConfig, opts Options) (bool, error) {
	// Detect category
	category, err := DetectCategory(folderPath)
	if err != nil {
		return false, fmt.Errorf("failed to detect category for %s: %w", folderPath, err)
	}

	// If category is unknown, skip or report
	if category == "" {
		if !opts.Quiet {
			fmt.Fprintf(os.Stderr, "Warning: %s - unknown or unsupported release category\n", folderPath)
		}
		return false, nil
	}

	// Validate folder
	result, err := ValidateFolder(folderPath, presetConfig, category)
	if err != nil {
		return false, fmt.Errorf("failed to validate folder: %w", err)
	}

	// Display results and return validation status
	failed := DisplayResult(result, opts)
	return !failed, nil
}

// ValidateFolders validates multiple folders
func ValidateFolders(folders []string, opts Options) error {
	// Load preset configuration
	presetConfig, err := preset.LoadPresets(opts.PresetPath)
	if err != nil {
		return fmt.Errorf("failed to load presets: %w", err)
	}

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
			// Find all folders recursively
			subFolders, err := FindFoldersRecursive(absPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: failed to find folders recursively in %s: %v\n", folder, err)
				hasErrors = true
				continue
			}

			if len(subFolders) == 0 {
				if !opts.Quiet {
					fmt.Fprintf(os.Stderr, "No valid release folders found in %s\n", folder)
				}
				hasErrors = true
				continue
			}

			// Validate each folder found
			for _, subFolder := range subFolders {
				valid, err := validateSingleFolder(subFolder, presetConfig, opts)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error: %v\n", err)
					hasErrors = true
				} else if !valid {
					hasErrors = true
				}
			}
		} else {
			// Validate single folder
			valid, err := validateSingleFolder(absPath, presetConfig, opts)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				hasErrors = true
			} else if !valid {
				hasErrors = true
			}
		}
	}

	if hasErrors {
		return fmt.Errorf("one or more folders had errors")
	}

	return nil
}
