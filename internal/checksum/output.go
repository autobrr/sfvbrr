package checksum

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// OutputResult represents the JSON/YAML output structure for SFV validation
type OutputResult struct {
	SFVFile     SFVFileOutput `json:"sfv_file" yaml:"sfv_file"`
	TotalFiles  int           `json:"total_files" yaml:"total_files"`
	ValidFiles  int           `json:"valid_files" yaml:"valid_files"`
	InvalidFiles int          `json:"invalid_files" yaml:"invalid_files"`
	MissingFiles int          `json:"missing_files" yaml:"missing_files"`
	Results     []SFVResultOutput `json:"results,omitempty" yaml:"results,omitempty"`
	Errors      []string      `json:"errors,omitempty" yaml:"errors,omitempty"`
}

type SFVFileOutput struct {
	Path    string        `json:"path" yaml:"path"`
	Dir     string        `json:"dir" yaml:"dir"`
	Entries []SFVEntry    `json:"entries" yaml:"entries"`
}

type SFVResultOutput struct {
	Filename string `json:"filename" yaml:"filename"`
	Path     string `json:"path" yaml:"path"`
	Valid    bool   `json:"valid" yaml:"valid"`
	Computed string `json:"computed,omitempty" yaml:"computed,omitempty"`
	Error    string `json:"error,omitempty" yaml:"error,omitempty"`
}

// ZIPOutputResult represents the JSON/YAML output structure for ZIP validation
type ZIPOutputResult struct {
	ZIPFile       string            `json:"zip_file" yaml:"zip_file"`
	TotalEntries  int               `json:"total_entries" yaml:"total_entries"`
	ValidEntries  int               `json:"valid_entries" yaml:"valid_entries"`
	InvalidEntries int              `json:"invalid_entries" yaml:"invalid_entries"`
	Results       []ZIPResultOutput `json:"results,omitempty" yaml:"results,omitempty"`
	Errors        []string           `json:"errors,omitempty" yaml:"errors,omitempty"`
}

type ZIPResultOutput struct {
	Name   string `json:"name" yaml:"name"`
	Valid  bool   `json:"valid" yaml:"valid"`
	Error  string `json:"error,omitempty" yaml:"error,omitempty"`
}

// convertValidationResult converts ValidationResult to OutputResult
func convertValidationResult(result *ValidationResult) *OutputResult {
	output := &OutputResult{
		SFVFile: SFVFileOutput{
			Path:    result.SFVFile.Path,
			Dir:     result.SFVFile.Dir,
			Entries: result.SFVFile.Entries,
		},
		TotalFiles:   result.TotalFiles,
		ValidFiles:   result.ValidFiles,
		InvalidFiles: result.InvalidFiles,
		MissingFiles: result.MissingFiles,
	}

	if len(result.Results) > 0 {
		output.Results = make([]SFVResultOutput, len(result.Results))
		for i, res := range result.Results {
			output.Results[i] = SFVResultOutput{
				Filename: res.Entry.Filename,
				Path:     res.Entry.Path,
				Valid:    res.Valid,
				Computed: res.Computed,
			}
			if res.Error != nil {
				output.Results[i].Error = res.Error.Error()
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

// convertZIPValidationResult converts ZIPValidationResult to ZIPOutputResult
func convertZIPValidationResult(result *ZIPValidationResult) *ZIPOutputResult {
	output := &ZIPOutputResult{
		ZIPFile:        result.ZIPFile.Path,
		TotalEntries:   result.TotalEntries,
		ValidEntries:   result.ValidEntries,
		InvalidEntries: result.InvalidEntries,
	}

	if len(result.Results) > 0 {
		output.Results = make([]ZIPResultOutput, len(result.Results))
		for i, res := range result.Results {
			output.Results[i] = ZIPResultOutput{
				Name:  res.Entry.Name,
				Valid: res.Valid,
			}
			if res.Error != nil {
				output.Results[i].Error = res.Error.Error()
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

// OutputZIPValidationResult outputs the ZIP validation result in the specified format
func OutputZIPValidationResult(result *ZIPValidationResult, format OutputFormat) error {
	if format == OutputFormatText {
		return nil // Use regular display
	}

	output := convertZIPValidationResult(result)

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
