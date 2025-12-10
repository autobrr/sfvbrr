package checksum

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// ZIPEntry represents a single entry in a ZIP file
type ZIPEntry struct {
	Name string // Name of the file inside the ZIP
	Path string // Full path to the ZIP file
}

// ZIPResult represents the result of validating a single ZIP entry
type ZIPResult struct {
	Entry    ZIPEntry
	Valid    bool
	Error    error
}

// ZIPFile represents a ZIP file being validated
type ZIPFile struct {
	Path    string     // Path to the ZIP file
	Entries []ZIPEntry // All entries in the ZIP file
}

// ZIPValidationResult represents the overall result of ZIP validation
type ZIPValidationResult struct {
	ZIPFile      ZIPFile
	Results      []ZIPResult
	TotalEntries int
	ValidEntries int
	InvalidEntries int
	Errors       []error
}

// FindZIPFiles finds all ZIP files in the given directory (case insensitive)
func FindZIPFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var zipFiles []string
	for _, entry := range entries {
		if !entry.IsDir() {
			filename := entry.Name()
			ext := strings.ToLower(filepath.Ext(filename))
			if ext == ".zip" {
				zipFiles = append(zipFiles, filepath.Join(dir, filename))
			}
		}
	}

	if len(zipFiles) == 0 {
		return nil, fmt.Errorf("no ZIP files found in directory: %s", dir)
	}

	return zipFiles, nil
}

// FindZIPFilesRecursive finds all ZIP files recursively in the given directory
func FindZIPFilesRecursive(dir string) ([]string, error) {
	var zipFiles []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Continue on errors (e.g., permission denied)
			return nil
		}

		if !info.IsDir() {
			// Check if file has .zip extension (case insensitive)
			ext := strings.ToLower(filepath.Ext(path))
			if ext == ".zip" {
				zipFiles = append(zipFiles, path)
			}
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %w", err)
	}

	return zipFiles, nil
}

// ParseZIPFile parses a ZIP file and returns all entries
func ParseZIPFile(zipPath string) (*ZIPFile, error) {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open ZIP file: %w", err)
	}
	defer r.Close()

	zipFile := &ZIPFile{
		Path:    zipPath,
		Entries: make([]ZIPEntry, 0, len(r.File)),
	}

	for _, f := range r.File {
		// Skip directory entries
		if f.FileInfo().IsDir() {
			continue
		}
		entry := ZIPEntry{
			Name: f.Name,
			Path: zipPath,
		}
		zipFile.Entries = append(zipFile.Entries, entry)
	}

	if len(zipFile.Entries) == 0 {
		return nil, fmt.Errorf("no entries found in ZIP file")
	}

	return zipFile, nil
}

// validateZIPEntry validates a single entry in a ZIP file by reading it
// This is equivalent to `zip -T` which tests the integrity of ZIP entries
func validateZIPEntry(zipPath string, entryName string) ZIPResult {
	result := ZIPResult{
		Entry: ZIPEntry{
			Name: entryName,
			Path: zipPath,
		},
	}

	// Open the ZIP file
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		result.Valid = false
		result.Error = fmt.Errorf("failed to open ZIP file: %w", err)
		return result
	}
	defer r.Close()

	// Find the entry
	var file *zip.File
	for _, f := range r.File {
		if f.Name == entryName {
			file = f
			break
		}
	}

	if file == nil {
		result.Valid = false
		result.Error = fmt.Errorf("entry not found: %s", entryName)
		return result
	}

	// Open and read the entry to verify its CRC-32
	// The zip package automatically verifies CRC-32 when reading
	rc, err := file.Open()
	if err != nil {
		result.Valid = false
		result.Error = fmt.Errorf("failed to open entry: %w", err)
		return result
	}
	defer rc.Close()

	// Read the entire entry to trigger CRC-32 verification
	_, err = io.Copy(io.Discard, rc)
	if err != nil {
		result.Valid = false
		result.Error = fmt.Errorf("failed to read entry (CRC-32 mismatch or corrupted): %w", err)
		return result
	}

	result.Valid = true
	return result
}

// ValidateZIP validates all entries in a ZIP file using parallel processing
func ValidateZIP(zip *ZIPFile, opts Options) (*ZIPValidationResult, error) {
	if len(zip.Entries) == 0 {
		return nil, fmt.Errorf("no entries to validate")
	}

	// Calculate optimal worker count
	workers := calculateOptimalWorkers(len(zip.Entries), opts.Workers)

	// Create displayer for progress tracking
	formatter := NewFormatter(opts.Verbose)
	displayer := NewDisplay(formatter)
	displayer.SetQuiet(opts.Quiet)

	// Show files and initialize progress bar
	if !opts.Quiet {
		// For ZIP files, we don't show the file tree (entries are inside ZIP files)
		// Just show the progress bar which is the main reporting mechanism
		displayer.ShowProgress(len(zip.Entries))
	}
	defer func() {
		if !opts.Quiet {
			displayer.FinishProgress()
		}
	}()

	result := &ZIPValidationResult{
		ZIPFile:      *zip,
		Results:      make([]ZIPResult, len(zip.Entries)),
		TotalEntries: len(zip.Entries),
		Errors:       make([]error, 0),
	}

	// Create channels for work distribution
	entryChan := make(chan int, workers)
	resultChan := make(chan struct {
		index  int
		result ZIPResult
	}, len(zip.Entries))

	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range entryChan {
				entry := zip.Entries[idx]
				validationResult := validateZIPEntry(entry.Path, entry.Name)
				resultChan <- struct {
					index  int
					result ZIPResult
				}{idx, validationResult}
			}
		}()
	}

	// Send work to workers
	go func() {
		for i := range zip.Entries {
			entryChan <- i
		}
		close(entryChan)
	}()

	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Create progress tracker
	tracker := NewProgressTracker(len(zip.Entries))
	completed := 0

	// Collect results and update progress
	for res := range resultChan {
		result.Results[res.index] = res.result
		if res.result.Valid {
			result.ValidEntries++
		} else {
			result.InvalidEntries++
			if res.result.Error != nil {
				result.Errors = append(result.Errors, res.result.Error)
			}
		}

		// Update progress
		completed++
		tracker.Update(completed)
		rate := tracker.GetRate()
		displayer.UpdateProgress(completed, rate)
	}

	return result, nil
}

// validateSingleZIP validates a single ZIP file and displays results
// Returns true if validation failed (has invalid entries)
func validateSingleZIP(zipPath string, opts Options) (bool, error) {
	// Parse ZIP file
	zip, err := ParseZIPFile(zipPath)
	if err != nil {
		return false, fmt.Errorf("failed to parse ZIP file %s: %w", zipPath, err)
	}

	// Validate ZIP
	result, err := ValidateZIP(zip, opts)
	if err != nil {
		return false, fmt.Errorf("failed to validate ZIP: %w", err)
	}

	// Display results and return validation status
	failed := DisplayZIPResult(result, opts)
	return failed, nil
}

// ValidateZIPFolders validates ZIP files in multiple folders
func ValidateZIPFolders(folders []string, opts Options) error {
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
			// Find all ZIP files recursively
			zipFiles, err := FindZIPFilesRecursive(absPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: failed to find ZIP files recursively in %s: %v\n", folder, err)
				hasErrors = true
				continue
			}

			if len(zipFiles) == 0 {
				if !opts.Quiet {
					fmt.Fprintf(os.Stderr, "No ZIP files found in %s\n", folder)
				}
				hasErrors = true
				continue
			}

			// Validate each ZIP file found
			for _, zipPath := range zipFiles {
				failed, err := validateSingleZIP(zipPath, opts)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error: %v\n", err)
					hasErrors = true
				} else if failed {
					hasErrors = true
				}
			}
		} else {
			// Find all ZIP files in current directory only
			zipFiles, err := FindZIPFiles(absPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				hasErrors = true
				continue
			}

			// Validate each ZIP file found
			for _, zipPath := range zipFiles {
				failed, err := validateSingleZIP(zipPath, opts)
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
