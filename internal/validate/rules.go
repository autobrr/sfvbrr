package validate

import (
	"fmt"
	"os"
	"path/filepath"
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
		},
		Description: rule.Description,
	}

	// Determine if we're matching files or directories
	isDirRule := rule.Type == "dir"

	// Count matches
	matched, err := countMatches(folderPath, rule.Pattern, isDirRule)
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
func countMatches(folderPath string, pattern string, isDir bool) (int, error) {
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

			matched, err := matchPattern(entry.Name(), dirPattern)
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

					matched, err := matchPattern(subEntry.Name(), filePattern)
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

			matched, err := matchPattern(entry.Name(), pattern)
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
// Supports glob patterns like *.nfo, *.r???, and brace expansion like {mkv,mp4}
func matchPattern(filename string, pattern string) (bool, error) {
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
