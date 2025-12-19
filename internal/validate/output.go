package validate

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// OutputResult represents the JSON/YAML output structure for validation
type OutputResult struct {
	FolderPath      string            `json:"folder_path" yaml:"folder_path"`
	Category        string            `json:"category" yaml:"category"`
	Valid           bool              `json:"valid" yaml:"valid"`
	RuleResults     []RuleResultOutput `json:"rule_results,omitempty" yaml:"rule_results,omitempty"`
	UnexpectedFiles []string          `json:"unexpected_files,omitempty" yaml:"unexpected_files,omitempty"`
	Errors          []string          `json:"errors,omitempty" yaml:"errors,omitempty"`
}

type RuleResultOutput struct {
	Pattern     string `json:"pattern" yaml:"pattern"`
	Type        string `json:"type" yaml:"type"`
	Matched     int    `json:"matched" yaml:"matched"`
	Valid       bool   `json:"valid" yaml:"valid"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Error       string `json:"error,omitempty" yaml:"error,omitempty"`
}

// convertValidationResult converts ValidationResult to OutputResult
func convertValidationResult(result *ValidationResult) *OutputResult {
	output := &OutputResult{
		FolderPath:      result.FolderPath,
		Category:        result.Category,
		Valid:           result.Valid,
		UnexpectedFiles: result.UnexpectedFiles,
	}

	if len(result.RuleResults) > 0 {
		output.RuleResults = make([]RuleResultOutput, len(result.RuleResults))
		for i, res := range result.RuleResults {
			output.RuleResults[i] = RuleResultOutput{
				Pattern:     res.Rule.Pattern,
				Type:        res.Rule.Type,
				Matched:     res.Matched,
				Valid:       res.Valid,
				Description: res.Description,
			}
			if res.Error != nil {
				output.RuleResults[i].Error = res.Error.Error()
			}
		}
	}

	if len(result.Errors) > 0 {
		output.Errors = make([]string, len(result.Errors))
		for i, err := range result.Errors {
			output.Errors[i] = err.Error()
		}
	}

	return output
}

// OutputValidationResult outputs the validation result in the specified format
func OutputValidationResult(result *ValidationResult, format OutputFormat) error {
	if format == OutputFormatText {
		return nil // Use regular display
	}

	output := convertValidationResult(result)

	switch format {
	case OutputFormatJSON:
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		return encoder.Encode(output)
	case OutputFormatYAML:
		encoder := yaml.NewEncoder(os.Stdout)
		defer encoder.Close()
		return encoder.Encode(output)
	default:
		return fmt.Errorf("unknown output format: %s", format)
	}
}
