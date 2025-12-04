package validate

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
	FolderPath string
	Category   string
	Valid      bool
	RuleResults []RuleResult
	Errors     []error
}

// Options contains configuration options for validation
type Options struct {
	PresetPath string // Path to preset YAML file (empty = auto-detect)
	Verbose    bool   // Verbose output
	Quiet      bool   // Quiet mode (minimal output)
	Recursive  bool   // Recursive mode - search subdirectories
}

// DefaultOptions returns default options for validation
func DefaultOptions() Options {
	return Options{
		PresetPath: "",
		Verbose:    false,
		Quiet:      false,
		Recursive:  false,
	}
}

// Rule represents a validation rule (imported from preset package)
type Rule struct {
	Pattern     string
	Type        string // "file" (default) or "dir"
	Min         int
	Max         int
	Description string
}
