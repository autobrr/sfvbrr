package checksum

import (
	"bufio"
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

const (
	// Default buffer size for reading files (64KB)
	defaultBufferSize = 64 * 1024
	// Minimum buffer size (4KB)
	minBufferSize = 4 * 1024
	// Maximum buffer size (1MB)
	maxBufferSize = 1024 * 1024
)

// bufferPool is a pool of reusable buffers for file reading
var bufferPool = sync.Pool{
	New: func() interface{} {
		buf := make([]byte, defaultBufferSize)
		return &buf
	},
}

// FindSFVFile searches for an SFV file in the given directory (case insensitive)
func FindSFVFile(dir string) (string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", fmt.Errorf("failed to read directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			filename := entry.Name()
			if strings.EqualFold(filepath.Ext(filename), ".sfv") {
				return filepath.Join(dir, filename), nil
			}
		}
	}

	return "", fmt.Errorf("no SFV file found in directory: %s", dir)
}

// ParseSFVFile parses an SFV file and returns all entries
func ParseSFVFile(sfvPath string) (*SFVFile, error) {
	file, err := os.Open(sfvPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open SFV file: %w", err)
	}
	defer file.Close()

	dir := filepath.Dir(sfvPath)
	sfv := &SFVFile{
		Path:    sfvPath,
		Dir:     dir,
		Entries: make([]SFVEntry, 0),
	}

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, ";") {
			continue
		}

		// Parse line: filename checksum
		parts := strings.Fields(line)
		if len(parts) < 2 {
			// Skip malformed lines
			continue
		}

		// Last part is the checksum, everything before is the filename
		checksum := parts[len(parts)-1]
		filename := strings.Join(parts[:len(parts)-1], " ")

		// Validate checksum format (should be 8 hex characters)
		if len(checksum) != 8 {
			continue
		}

		entry := SFVEntry{
			Filename: filename,
			Checksum: strings.ToUpper(checksum),
		}
		entry.JoinPath(dir)

		sfv.Entries = append(sfv.Entries, entry)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading SFV file: %w", err)
	}

	if len(sfv.Entries) == 0 {
		return nil, fmt.Errorf("no valid entries found in SFV file")
	}

	return sfv, nil
}

// computeCRC32 computes the CRC-32 checksum of a file
func computeCRC32(filePath string, buffer []byte) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	hash := crc32.NewIEEE()
	_, err = io.CopyBuffer(hash, file, buffer)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	// Format as 8-character uppercase hexadecimal
	return strings.ToUpper(fmt.Sprintf("%08x", hash.Sum32())), nil
}

// validateFile validates a single file against its expected checksum
func validateFile(entry SFVEntry, buffer []byte) SFVResult {
	result := SFVResult{
		Entry: entry,
	}

	// Check if file exists
	if _, err := os.Stat(entry.Path); os.IsNotExist(err) {
		result.Valid = false
		result.Error = fmt.Errorf("file not found: %s", entry.Filename)
		return result
	}

	// Compute CRC-32
	computed, err := computeCRC32(entry.Path, buffer)
	if err != nil {
		result.Valid = false
		result.Error = err
		return result
	}

	result.Computed = computed
	result.Valid = strings.EqualFold(computed, entry.Checksum)

	if !result.Valid {
		result.Error = fmt.Errorf("checksum mismatch: expected %s, got %s", entry.Checksum, computed)
	}

	return result
}

// calculateOptimalWorkers calculates the optimal number of workers based on file count
func calculateOptimalWorkers(fileCount int, requested int) int {
	if requested > 0 {
		return requested
	}

	// Use number of CPU cores as base
	numCPU := runtime.NumCPU()

	// For small file counts, use fewer workers
	if fileCount < numCPU {
		return fileCount
	}

	// For larger file counts, use 2x CPU cores for better parallelism
	// but cap at a reasonable maximum
	maxWorkers := numCPU * 2
	if maxWorkers > 16 {
		maxWorkers = 16
	}

	if fileCount < maxWorkers {
		return fileCount
	}

	return maxWorkers
}

// ValidateSFV validates all files in an SFV file using parallel processing
func ValidateSFV(sfv *SFVFile, opts Options) (*ValidationResult, error) {
	if len(sfv.Entries) == 0 {
		return nil, fmt.Errorf("no entries to validate")
	}

	// Calculate optimal worker count
	workers := calculateOptimalWorkers(len(sfv.Entries), opts.Workers)

	// Determine buffer size
	bufferSize := opts.BufferSize
	if bufferSize == 0 {
		bufferSize = defaultBufferSize
	}
	if bufferSize < minBufferSize {
		bufferSize = minBufferSize
	}
	if bufferSize > maxBufferSize {
		bufferSize = maxBufferSize
	}

	// Create displayer for progress tracking
	formatter := NewFormatter(opts.Verbose)
	displayer := NewDisplay(formatter)
	displayer.SetQuiet(opts.Quiet)
	displayer.SetBatch(opts.Recursive) // Batch mode for recursive operations

	// Show files and initialize progress bar
	if !opts.Quiet && !opts.Recursive {
		displayer.ShowFiles(sfv.Entries, workers)
		displayer.ShowProgress(len(sfv.Entries))
	}
	defer func() {
		if !opts.Quiet && !opts.Recursive {
			displayer.FinishProgress()
		}
	}()

	result := &ValidationResult{
		SFVFile:    *sfv,
		Results:    make([]SFVResult, len(sfv.Entries)),
		TotalFiles: len(sfv.Entries),
		Errors:     make([]error, 0),
	}

	// Create channels for work distribution
	entryChan := make(chan int, workers)
	resultChan := make(chan struct {
		index  int
		result SFVResult
	}, len(sfv.Entries))

	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Get buffer from pool or allocate new one
			var buffer []byte
			if bufPtr := bufferPool.Get(); bufPtr != nil {
				buf := bufPtr.(*[]byte)
				if len(*buf) >= bufferSize {
					buffer = (*buf)[:bufferSize]
					defer bufferPool.Put(buf)
				} else {
					buffer = make([]byte, bufferSize)
					defer func() {
						newBuf := buffer
						bufferPool.Put(&newBuf)
					}()
				}
			} else {
				buffer = make([]byte, bufferSize)
				defer func() {
					newBuf := buffer
					bufferPool.Put(&newBuf)
				}()
			}

			for idx := range entryChan {
				entry := sfv.Entries[idx]
				validationResult := validateFile(entry, buffer)
				resultChan <- struct {
					index  int
					result SFVResult
				}{idx, validationResult}
			}
		}()
	}

	// Send work to workers
	go func() {
		for i := range sfv.Entries {
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
	tracker := NewProgressTracker(len(sfv.Entries))
	completed := 0

	// Collect results and update progress
	for res := range resultChan {
		result.Results[res.index] = res.result
		if res.result.Valid {
			result.ValidFiles++
		} else {
			if res.result.Error != nil {
				if strings.Contains(res.result.Error.Error(), "file not found") {
					result.MissingFiles++
				} else {
					result.InvalidFiles++
				}
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
