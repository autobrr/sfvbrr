package validate

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

var (
	magenta   = color.New(color.FgMagenta).SprintFunc()
	yellow    = color.New(color.FgYellow).SprintFunc()
	success   = color.New(color.FgGreen).SprintFunc()
	label     = color.New(color.FgCyan).SprintFunc()
	errorColor = color.New(color.FgRed).SprintFunc()
)

// DisplayResult displays the validation results to the user
// Returns true if validation failed (has invalid rules)
func DisplayResult(result *ValidationResult, opts Options) bool {
	if opts.Quiet {
		// In quiet mode, only show errors
		if !result.Valid {
			fmt.Fprintf(os.Stderr, "%s: validation failed\n", result.FolderPath)
		}
		return !result.Valid
	}

	// Show folder path and category
	fmt.Fprintf(os.Stdout, "\n%s\n", magenta("Validating Release:"))
	fmt.Fprintf(os.Stdout, "  %-13s %s\n", label("Folder:"), result.FolderPath)

	if result.Category != "" {
		fmt.Fprintf(os.Stdout, "  %-13s %s\n", label("Category:"), result.Category)
	} else {
		fmt.Fprintf(os.Stdout, "  %-13s %s\n", label("Category:"), yellow("unknown"))
	}
	fmt.Fprintln(os.Stdout)

	// Show rule results
	if len(result.RuleResults) > 0 {
		fmt.Fprintf(os.Stdout, "%s\n", magenta("Rule Validation:"))

		validCount := 0
		invalidCount := 0

		for _, ruleResult := range result.RuleResults {
			if ruleResult.Valid {
				validCount++
				if opts.Verbose {
					fmt.Fprintf(os.Stdout, "  %s %s", success("✓"), ruleResult.Rule.Pattern)
					if ruleResult.Matched > 0 {
						fmt.Fprintf(os.Stdout, " (found %d)", ruleResult.Matched)
					}
					if ruleResult.Description != "" {
						fmt.Fprintf(os.Stdout, " - %s", ruleResult.Description)
					}
					fmt.Fprintln(os.Stdout)
				}
			} else {
				invalidCount++
				fmt.Fprintf(os.Stdout, "  %s %s", errorColor("✗"), ruleResult.Rule.Pattern)
				if ruleResult.Matched > 0 {
					fmt.Fprintf(os.Stdout, " (found %d)", ruleResult.Matched)
				}
				if ruleResult.Error != nil {
					fmt.Fprintf(os.Stdout, " - %s", errorColor(ruleResult.Error.Error()))
				} else if ruleResult.Description != "" {
					fmt.Fprintf(os.Stdout, " - %s", ruleResult.Description)
				}
				fmt.Fprintln(os.Stdout)
			}
		}

		fmt.Fprintln(os.Stdout)

		// Show summary
		fmt.Fprintf(os.Stdout, "%s\n", magenta("Summary:"))
		fmt.Fprintf(os.Stdout, "  %-15s %s\n", label("Valid rules:"), success(validCount))
		if invalidCount > 0 {
			fmt.Fprintf(os.Stdout, "  %-15s %s\n", label("Invalid rules:"), errorColor(invalidCount))
		}
		fmt.Fprintln(os.Stdout)
	} else {
		// No rules found for this category
		fmt.Fprintf(os.Stdout, "%s\n", yellow("No validation rules found for this category"))
		fmt.Fprintln(os.Stdout)
	}

	// Show unexpected files if any
	if len(result.UnexpectedFiles) > 0 {
		fmt.Fprintf(os.Stdout, "%s\n", errorColor("Unexpected Files/Directories:"))
		for _, file := range result.UnexpectedFiles {
			fmt.Fprintf(os.Stdout, "  %s %s\n", errorColor("✗"), file)
		}
		fmt.Fprintln(os.Stdout)
	}

	// Show errors if any
	if len(result.Errors) > 0 {
		fmt.Fprintf(os.Stdout, "%s\n", errorColor("Errors:"))
		for _, err := range result.Errors {
			fmt.Fprintf(os.Stdout, "  %s\n", errorColor(err.Error()))
		}
		fmt.Fprintln(os.Stdout)
	}

	// Return true if validation failed
	return !result.Valid
}

// FormatFolderPath formats a folder path for display (relative to current directory if possible)
func FormatFolderPath(path string) string {
	wd, err := os.Getwd()
	if err != nil {
		return path
	}

	relPath, err := filepath.Rel(wd, path)
	if err != nil {
		return path
	}

	// If relative path is shorter or more readable, use it
	if len(relPath) < len(path) && !filepath.IsAbs(relPath) {
		return relPath
	}

	return path
}
