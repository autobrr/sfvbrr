package checksum

import (
	"path/filepath"
)

// OutputFormat represents the output format type
type OutputFormat string

const (
	OutputFormatText OutputFormat = "text"
	OutputFormatJSON OutputFormat = "json"
	OutputFormatYAML OutputFormat = "yaml"
)

// SFVEntry represents a single entry in an SFV file
type SFVEntry struct {
	Filename string
	Checksum string // CRC-32 checksum in hexadecimal format
	Path     string // Full path to the file
}

// SFVResult represents the result of validating a single file
type SFVResult struct {
	Entry    SFVEntry
	Valid    bool
	Error    error
	Computed string // The computed CRC-32 checksum
}

// SFVFile represents a parsed SFV file
type SFVFile struct {
	Path    string     // Path to the SFV file
	Entries []SFVEntry // All entries in the SFV file
	Dir     string     // Directory containing the SFV file
}

// ValidationResult represents the overall result of SFV validation
type ValidationResult struct {
	SFVFile    SFVFile
	Results    []SFVResult
	TotalFiles int
	ValidFiles int
	InvalidFiles int
	MissingFiles int
	Errors     []error
}

// Options contains configuration options for SFV validation
type Options struct {
	Workers      int          // Number of parallel workers (0 = auto)
	BufferSize   int          // Buffer size for file reading (0 = auto)
	Verbose      bool         // Verbose output
	Quiet        bool         // Quiet mode (minimal output)
	Recursive    bool         // Recursive mode - search subdirectories
	OutputFormat OutputFormat // Output format: text, json, or yaml
}

// DefaultOptions returns default options for SFV validation
func DefaultOptions() Options {
	return Options{
		Workers:    0, // Auto-detect
		BufferSize: 0, // Auto-detect
		Verbose:    false,
		Quiet:      false,
	}
}

// JoinPath safely joins the directory path with the filename
func (e *SFVEntry) JoinPath(dir string) {
	e.Path = filepath.Join(dir, e.Filename)
}
