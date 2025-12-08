package validate

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/autobrr/sfvbrr/internal/preset"
)

// ValidateFolder validates a folder against rules for its category
func ValidateFolder(folderPath string, presetConfig *preset.PresetConfig, category string) (*ValidationResult, error) {
	result := &ValidationResult{
		FolderPath:  folderPath,
		Category:    category,
		Valid:       true,
		RuleResults: make([]RuleResult, 0),
		Errors:      make([]error, 0),
	}

	// If category is empty/unknown, return early
	if category == "" {
		result.Valid = false
		result.Errors = append(result.Errors, fmt.Errorf("unknown or unsupported release category"))
		return result, nil
	}

	// Get rules for this category
	rules, err := presetConfig.GetRulesForCategory(category)
	if err != nil {
		result.Valid = false
		result.Errors = append(result.Errors, err)
		return result, nil
	}

	// Validate each rule
	for _, rule := range rules {
		ruleResult := validateRule(folderPath, rule)
		result.RuleResults = append(result.RuleResults, ruleResult)

		if !ruleResult.Valid {
			result.Valid = false
			if ruleResult.Error != nil {
				result.Errors = append(result.Errors, ruleResult.Error)
			}
		}
	}

	// Check for unexpected files/directories if deny_unexpected is enabled
	if presetConfig.GetDenyUnexpected(category) {
		unexpected, err := findUnexpectedFiles(folderPath, rules)
		if err != nil {
			result.Valid = false
			result.Errors = append(result.Errors, fmt.Errorf("failed to check for unexpected files: %w", err))
		} else if len(unexpected) > 0 {
			result.Valid = false
			result.UnexpectedFiles = unexpected
			result.Errors = append(result.Errors, fmt.Errorf("found %d unexpected file(s)/directory(ies)", len(unexpected)))
		}
	}

	return result, nil
}

// validateRule validates a single rule against a folder
func validateRule(folderPath string, rule preset.Rule) RuleResult {
	result := RuleResult{
		Rule: Rule{
			Pattern:     rule.Pattern,
			Type:        rule.Type,
			Min:         rule.Min,
			Max:         rule.Max,
			Description: rule.Description,
			Regex:       rule.Regex,
		},
		Description: rule.Description,
	}

	// Determine if we're matching files or directories
	isDirRule := rule.Type == "dir"

	// Count matches
	matched, err := countMatches(folderPath, rule.Pattern, isDirRule, rule.Regex)
	if err != nil {
		result.Valid = false
		result.Error = err
		return result
	}

	result.Matched = matched

	// Check min constraint
	if rule.Min > 0 && matched < rule.Min {
		result.Valid = false
		result.Error = fmt.Errorf("found %d matches, but minimum required is %d", matched, rule.Min)
		return result
	}

	// Check max constraint
	if rule.Max > 0 && matched > rule.Max {
		result.Valid = false
		result.Error = fmt.Errorf("found %d matches, but maximum allowed is %d", matched, rule.Max)
		return result
	}

	result.Valid = true
	return result
}

// countMatches counts how many files or directories match the pattern
func countMatches(folderPath string, pattern string, isDir bool, useRegex bool) (int, error) {
	// Read directory entries
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return 0, fmt.Errorf("failed to read directory: %w", err)
	}

	count := 0

	// Handle special patterns like "Sample/*.{mkv,mp4}"
	if strings.Contains(pattern, "/") && strings.Contains(pattern, "{") {
		// This is a nested pattern like "Sample/*.{mkv,mp4}"
		parts := strings.SplitN(pattern, "/", 2)
		dirPattern := parts[0]
		filePattern := parts[1]

		// First find matching directories
		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}

			matched, err := matchPattern(entry.Name(), dirPattern, useRegex)
			if err != nil {
				continue
			}

			if matched {
				// Check files inside this directory
				subDirPath := filepath.Join(folderPath, entry.Name())
				subEntries, err := os.ReadDir(subDirPath)
				if err != nil {
					continue
				}

				for _, subEntry := range subEntries {
					if subEntry.IsDir() {
						continue
					}

					matched, err := matchPattern(subEntry.Name(), filePattern, useRegex)
					if err != nil {
						continue
					}

					if matched {
						count++
					}
				}
			}
		}
	} else {
		// Regular pattern matching
		for _, entry := range entries {
			// Filter by type
			if isDir && !entry.IsDir() {
				continue
			}
			if !isDir && entry.IsDir() {
				continue
			}

			matched, err := matchPattern(entry.Name(), pattern, useRegex)
			if err != nil {
				continue
			}

			if matched {
				count++
			}
		}
	}

	return count, nil
}

// matchPattern matches a filename against a pattern
// Supports glob patterns like *.nfo, *.r???, brace expansion like {mkv,mp4}, and regex patterns
func matchPattern(filename string, pattern string, useRegex bool) (bool, error) {
	// If regex is enabled, use regex matching
	if useRegex {
		matched, err := regexp.MatchString(pattern, filename)
		if err != nil {
			return false, fmt.Errorf("invalid regex pattern: %w", err)
		}
		return matched, nil
	}

	// Handle brace expansion like {mkv,mp4} or *.{mkv,mp4}
	if strings.Contains(pattern, "{") && strings.Contains(pattern, "}") {
		// Extract the base pattern and brace content
		startIdx := strings.Index(pattern, "{")
		endIdx := strings.Index(pattern, "}")
		if startIdx >= 0 && endIdx > startIdx {
			prefix := pattern[:startIdx]
			suffix := pattern[endIdx+1:]
			braceContent := pattern[startIdx+1 : endIdx]

			// Split brace content by comma
			options := strings.Split(braceContent, ",")
			for _, option := range options {
				option = strings.TrimSpace(option)
				// Reconstruct pattern with expanded option
				testPattern := prefix + option + suffix
				matched, err := filepath.Match(testPattern, filename)
				if err != nil {
					continue
				}
				if matched {
					return true, nil
				}
			}
			return false, nil
		}
	}

	// Use filepath.Match for glob patterns
	matched, err := filepath.Match(pattern, filename)
	if err != nil {
		return false, err
	}

	return matched, nil
}

// findUnexpectedFiles finds all files and directories that don't match any rule pattern
func findUnexpectedFiles(folderPath string, rules []preset.Rule) ([]string, error) {
	// Read all directory entries (including hidden files)
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	// Track allowed entries
	allowedRootEntries := make(map[string]bool)            // Root level files/dirs
	allowedDirsForNested := make(map[string]bool)          // Dirs that have nested patterns
	allowedNestedFiles := make(map[string]map[string]bool) // dir -> set of allowed files

	// First pass: identify which entries match rules
	for _, rule := range rules {
		isDirRule := rule.Type == "dir"

		// Handle nested patterns like "Sample/*.{mkv,mp4}"
		if strings.Contains(rule.Pattern, "/") {
			parts := strings.SplitN(rule.Pattern, "/", 2)
			dirPattern := parts[0]
			filePattern := parts[1]

			// Find matching directories
			for _, entry := range entries {
				if !entry.IsDir() {
					continue
				}

				matched, err := matchPattern(entry.Name(), dirPattern, rule.Regex)
				if err != nil || !matched {
					continue
				}

				// This directory is allowed for nested patterns
				allowedDirsForNested[entry.Name()] = true
				if allowedNestedFiles[entry.Name()] == nil {
					allowedNestedFiles[entry.Name()] = make(map[string]bool)
				}

				// Check files inside this directory
				subDirPath := filepath.Join(folderPath, entry.Name())
				subEntries, err := os.ReadDir(subDirPath)
				if err != nil {
					continue
				}

				for _, subEntry := range subEntries {
					if subEntry.IsDir() {
						continue // Subdirectories are not allowed
					}

					matched, err := matchPattern(subEntry.Name(), filePattern, rule.Regex)
					if err == nil && matched {
						allowedNestedFiles[entry.Name()][subEntry.Name()] = true
					}
				}
			}
		} else {
			// Regular pattern matching
			for _, entry := range entries {
				// Filter by type
				if isDirRule && !entry.IsDir() {
					continue
				}
				if !isDirRule && entry.IsDir() {
					continue
				}

				matched, err := matchPattern(entry.Name(), rule.Pattern, rule.Regex)
				if err == nil && matched {
					allowedRootEntries[entry.Name()] = true
				}
			}
		}
	}

	// Second pass: find unexpected entries
	unexpected := make([]string, 0)

	for _, entry := range entries {
		entryName := entry.Name()

		if entry.IsDir() {
			// Check if directory is allowed (either explicitly or via nested patterns)
			if allowedRootEntries[entryName] || allowedDirsForNested[entryName] {
				// Only check directory contents if it has nested patterns
				// If it's just a simple directory rule, we don't validate its contents
				if allowedDirsForNested[entryName] {
					// Directory is allowed for nested patterns - check its contents
					subDirPath := filepath.Join(folderPath, entryName)
					subEntries, err := os.ReadDir(subDirPath)
					if err == nil {
						allowedFiles := allowedNestedFiles[entryName]
						for _, subEntry := range subEntries {
							if subEntry.IsDir() {
								// Subdirectories are not allowed
								unexpected = append(unexpected, filepath.Join(entryName, subEntry.Name()))
							} else if allowedFiles == nil || !allowedFiles[subEntry.Name()] {
								// File doesn't match any nested pattern
								unexpected = append(unexpected, filepath.Join(entryName, subEntry.Name()))
							}
						}
					}
				}
				// If directory is only in allowedRootEntries (simple dir rule), don't check contents
			} else {
				// Directory is not allowed at all
				unexpected = append(unexpected, entryName)
			}
		} else {
			// File is not allowed
			if !allowedRootEntries[entryName] {
				unexpected = append(unexpected, entryName)
			}
		}
	}

	return unexpected, nil
}
