package validate

// OutputFormat represents the output format type
type OutputFormat string

const (
	OutputFormatText OutputFormat = "text"
	OutputFormatJSON OutputFormat = "json"
	OutputFormatYAML OutputFormat = "yaml"
)

// RuleResult represents the result of validating a single rule
type RuleResult struct {
	Rule        Rule
	Matched     int
	Valid       bool
	Error       error
	Description string
}

// ValidationResult represents the overall result of folder validation
type ValidationResult struct {
	FolderPath      string
	Category        string
	Valid           bool
	RuleResults     []RuleResult
	Errors          []error
	UnexpectedFiles []string // Files/directories that don't match any rule pattern
}

// Options contains configuration options for validation
type Options struct {
	PresetPath        string       // Path to preset YAML file (empty = auto-detect)
	Verbose           bool         // Verbose output
	Quiet             bool         // Quiet mode (minimal output)
	Recursive         bool         // Recursive mode - search subdirectories
	OverwriteCategory string       // Override category detection (empty = use auto-detection)
	OutputFormat      OutputFormat // Output format: text, json, or yaml
}

// DefaultOptions returns default options for validation
func DefaultOptions() Options {
	return Options{
		PresetPath:       "",
		Verbose:          false,
		Quiet:            false,
		Recursive:        false,
		OverwriteCategory: "",
	}
}

// Rule represents a validation rule (imported from preset package)
type Rule struct {
	Pattern     string
	Type        string // "file" (default) or "dir"
	Min         int
	Max         int
	Description string
	Regex       bool // If true, pattern is treated as regex instead of glob
}
