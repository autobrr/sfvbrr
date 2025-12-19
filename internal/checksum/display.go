package checksum

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	progressbar "github.com/schollz/progressbar/v3"
)

type Display struct {
	output    io.Writer
	formatter *Formatter
	bar       *progressbar.ProgressBar
	isBatch   bool
	quiet     bool
}

func NewDisplay(formatter *Formatter) *Display {
	return &Display{
		formatter: formatter,
		quiet:     false,
		output:    os.Stdout,
	}
}

// SetQuiet enables/disables quiet mode (output redirected to io.Discard)
func (d *Display) SetQuiet(quiet bool) {
	d.quiet = quiet
	if quiet {
		d.output = io.Discard
	} else {
		d.output = os.Stdout
	}
}

func (d *Display) ShowProgress(total int) {
	// Progress bar needs explicit quiet check because it writes directly to the terminal,
	// bypassing our d.output writer
	if d.quiet {
		return
	}
	fmt.Fprintln(d.output)
	d.bar = progressbar.NewOptions(total,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetDescription("[cyan][bold]Validating files...[reset]"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)
}

func (d *Display) UpdateProgress(completed int, rate float64) {
	// Progress bar needs explicit quiet check because it writes directly to the terminal,
	// bypassing our d.output writer
	if d.quiet {
		return
	}
	// Allow progress updates even in batch mode - batch mode only suppresses file listings
	if d.bar != nil {
		if err := d.bar.Set(completed); err != nil {
			// Silently ignore progress bar errors
			_ = err
		}

		if rate > 0 {
			rateStr := d.formatter.FormatBytes(int64(rate))
			description := fmt.Sprintf("[cyan][bold]Validating files...[reset] [%s/s]", rateStr)
			d.bar.Describe(description)
		}
	}
}

// ShowFiles displays the list of files being validated and the number of workers used.
func (d *Display) ShowFiles(entries []SFVEntry, numWorkers int) {
	if d.quiet || d.isBatch {
		return
	}

	workerMsg := fmt.Sprintf("Using %d worker(s)", numWorkers)
	if numWorkers == 0 {
		workerMsg = "Using automatic worker count"
	}
	fmt.Fprintf(d.output, "\n%s %s\n", label("Concurrency:"), workerMsg)

	if !d.formatter.verbose && len(entries) > 20 {
		fmt.Fprintf(d.output, "%s suppressed file output (limit 20, found %d), use --verbose to show all\n", yellow("Note:"), len(entries))
		fmt.Fprintf(d.output, "%s\n", magenta("Files being validated:"))
		return
	}
	fmt.Fprintf(d.output, "\n%s\n", magenta("Files being validated:"))

	if len(entries) == 0 {
		return
	}

	// Find the common base directory
	commonBase := filepath.Dir(entries[0].Path)
	for _, entry := range entries[1:] {
		dir := filepath.Dir(entry.Path)
		for !strings.HasPrefix(dir, commonBase) {
			commonBase = filepath.Dir(commonBase)
			if commonBase == "/" || commonBase == "." {
				break
			}
		}
	}

	// Build a tree structure
	type fileNode struct {
		name     string
		isDir    bool
		children map[string]*fileNode
	}

	root := &fileNode{
		name:     filepath.Base(commonBase),
		isDir:    true,
		children: make(map[string]*fileNode),
	}

	// Add files to the tree
	for _, entry := range entries {
		relPath, _ := filepath.Rel(commonBase, entry.Path)
		parts := strings.Split(relPath, string(filepath.Separator))

		current := root
		for _, part := range parts[:len(parts)-1] {
			if _, exists := current.children[part]; !exists {
				current.children[part] = &fileNode{
					name:     part,
					isDir:    true,
					children: make(map[string]*fileNode),
				}
			}
			current = current.children[part]
		}

		// Add the file
		fileName := parts[len(parts)-1]
		current.children[fileName] = &fileNode{
			name:  fileName,
			isDir: false,
		}
	}

	// Display the tree
	var displayTree func(node *fileNode, prefix string, isLast bool)
	displayTree = func(node *fileNode, prefix string, isLast bool) {
		connector := "├─"
		if isLast {
			connector = "└─"
		}

		if prefix == "" {
			// Root node
			fmt.Fprintf(d.output, "%s %s\n", connector, success(node.name))
		} else {
			if node.isDir {
				fmt.Fprintf(d.output, "%s%s %s\n", prefix, connector, success(node.name))
			} else {
				fmt.Fprintf(d.output, "%s%s %s\n", prefix, connector, success(node.name))
			}
		}

		// Get sorted children
		childNames := make([]string, 0, len(node.children))
		for name := range node.children {
			childNames = append(childNames, name)
		}
		sort.Strings(childNames)

		// Display children
		for i, childName := range childNames {
			child := node.children[childName]
			childPrefix := prefix
			if prefix == "" {
				childPrefix = "  "
			} else {
				if isLast {
					childPrefix = prefix + "  "
				} else {
					childPrefix = prefix + "│ "
				}
			}
			displayTree(child, childPrefix, i == len(childNames)-1)
		}
	}

	displayTree(root, "", true)
}

func (d *Display) FinishProgress() {
	// Progress bar needs explicit quiet check because it writes directly to the terminal,
	// bypassing our d.output writer
	if d.quiet {
		return
	}
	if d.bar != nil {
		if err := d.bar.Finish(); err != nil {
			// Silently ignore progress bar errors
			_ = err
		}
		fmt.Fprintln(d.output)
	}
}

func (d *Display) IsBatch() bool {
	return d.isBatch
}

func (d *Display) SetBatch(isBatch bool) {
	d.isBatch = isBatch
}

var (
	magenta    = color.New(color.FgMagenta).SprintFunc()
	yellow     = color.New(color.FgYellow).SprintFunc()
	success    = color.New(color.FgGreen).SprintFunc()
	label      = color.New(color.FgCyan).SprintFunc()
	errorColor = color.New(color.FgRed).SprintFunc()
)

func (d *Display) ShowMessage(msg string) {
	fmt.Fprintf(d.output, "%s %s\n", success("\nInfo:"), msg)
}

func (d *Display) ShowError(msg string) {
	fmt.Fprintln(d.output, errorColor(msg))
}

func (d *Display) ShowWarning(msg string) {
	fmt.Fprintf(d.output, "%s %s\n", yellow("Warning:"), msg)
}

// DisplayResult displays the validation results to the user
// Returns true if validation failed (has invalid or missing files)
func DisplayResult(result *ValidationResult, opts Options) bool {
	// Handle JSON/YAML output
	if opts.OutputFormat != OutputFormatText {
		if err := OutputValidationResult(result, opts.OutputFormat); err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to output result: %v\n", err)
			return true
		}
		return result.InvalidFiles > 0 || result.MissingFiles > 0
	}

	formatter := NewFormatter(opts.Verbose)
	display := NewDisplay(formatter)
	display.SetQuiet(opts.Quiet)
	display.SetBatch(opts.Recursive) // Batch mode for recursive operations

	if opts.Quiet {
		// In quiet mode, only show summary if there are errors
		if result.InvalidFiles > 0 || result.MissingFiles > 0 {
			fmt.Fprintf(os.Stderr, "%s: %d invalid, %d missing\n",
				result.SFVFile.Path,
				result.InvalidFiles,
				result.MissingFiles)
		}
		return result.InvalidFiles > 0 || result.MissingFiles > 0
	}

	// Show SFV file path
	fmt.Fprintf(display.output, "\n%s\n", magenta("Validating SFV:"))
	fmt.Fprintf(display.output, "  %-13s %s\n", label("SFV file:"), result.SFVFile.Path)
	fmt.Fprintf(display.output, "  %-13s %d\n", label("Total files:"), result.TotalFiles)
	fmt.Fprintln(display.output)

	// Show individual results if verbose
	if opts.Verbose {
		fmt.Fprintf(display.output, "%s\n", magenta("Validation results:"))
		for _, res := range result.Results {
			if res.Valid {
				fmt.Fprintf(display.output, "  %s %s\n", success("✓"), res.Entry.Filename)
			} else {
				if res.Error != nil {
					if strings.Contains(res.Error.Error(), "file not found") {
						fmt.Fprintf(display.output, "  %s %s %s\n", errorColor("✗"), res.Entry.Filename, errorColor("(MISSING)"))
					} else {
						fmt.Fprintf(display.output, "  %s %s %s\n", errorColor("✗"), res.Entry.Filename, errorColor(fmt.Sprintf("(%s)", res.Error.Error())))
					}
				}
			}
		}
		fmt.Fprintln(display.output)
	}

	// Show summary
	fmt.Fprintf(display.output, "%s\n", magenta("Summary:"))
	fmt.Fprintf(display.output, "  %-15s %s\n", label("Valid:"), success(result.ValidFiles))
	if result.InvalidFiles > 0 {
		fmt.Fprintf(display.output, "  %-15s %s\n", label("Invalid:"), errorColor(result.InvalidFiles))
	}
	if result.MissingFiles > 0 {
		fmt.Fprintf(display.output, "  %-15s %s\n", label("Missing:"), errorColor(result.MissingFiles))
	}
	fmt.Fprintln(display.output)

	// Return true if validation failed
	return result.InvalidFiles > 0 || result.MissingFiles > 0
}

// DisplayZIPResult displays the ZIP validation results to the user
// Returns true if validation failed (has invalid entries)
func DisplayZIPResult(result *ZIPValidationResult, opts Options) bool {
	// Handle JSON/YAML output
	if opts.OutputFormat != OutputFormatText {
		if err := OutputZIPValidationResult(result, opts.OutputFormat); err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to output result: %v\n", err)
			return true
		}
		return result.InvalidEntries > 0
	}

	formatter := NewFormatter(opts.Verbose)
	display := NewDisplay(formatter)
	display.SetQuiet(opts.Quiet)
	display.SetBatch(opts.Recursive) // Batch mode for recursive operations

	if opts.Quiet {
		// In quiet mode, only show summary if there are errors
		if result.InvalidEntries > 0 {
			fmt.Fprintf(os.Stderr, "%s: %d invalid\n",
				result.ZIPFile.Path,
				result.InvalidEntries)
		}
		return result.InvalidEntries > 0
	}

	// Show ZIP file path
	fmt.Fprintf(display.output, "\n%s\n", magenta("Validating ZIP:"))
	fmt.Fprintf(display.output, "  %-13s %s\n", label("ZIP file:"), result.ZIPFile.Path)
	fmt.Fprintf(display.output, "  %-13s %d\n", label("Files in archive:"), result.TotalEntries)

	// If ZIP file couldn't be parsed, show the error
	if result.TotalEntries == 0 && len(result.Errors) > 0 {
		fmt.Fprintf(display.output, "  %-13s %s\n", label("Error:"), errorColor(result.Errors[0].Error()))
	}
	fmt.Fprintln(display.output)

	// Show individual results if verbose
	if opts.Verbose {
		fmt.Fprintf(display.output, "%s\n", magenta("Validation results:"))
		for _, res := range result.Results {
			if res.Valid {
				fmt.Fprintf(display.output, "  %s %s\n", success("✓"), res.Entry.Name)
			} else {
				if res.Error != nil {
					fmt.Fprintf(display.output, "  %s %s %s\n", errorColor("✗"), res.Entry.Name, errorColor(fmt.Sprintf("(%s)", res.Error.Error())))
				}
			}
		}
		fmt.Fprintln(display.output)
	}

	// Show summary
	fmt.Fprintf(display.output, "%s\n", magenta("Summary:"))
	fmt.Fprintf(display.output, "  %-15s %s\n", label("Valid:"), success(result.ValidEntries))
	if result.InvalidEntries > 0 {
		fmt.Fprintf(display.output, "  %-15s %s\n", label("Invalid:"), errorColor(result.InvalidEntries))
	}
	fmt.Fprintln(display.output)

	// Return true if validation failed
	return result.InvalidEntries > 0
}

type Formatter struct {
	verbose bool
}

func NewFormatter(verbose bool) *Formatter {
	return &Formatter{verbose: verbose}
}

func (f *Formatter) FormatBytes(bytes int64) string {
	return humanize.IBytes(uint64(bytes))
}

func (f *Formatter) FormatDuration(dur time.Duration) string {
	switch {
	case dur < time.Second:
		return fmt.Sprintf("%dms", dur.Milliseconds())
	case dur < time.Minute:
		return fmt.Sprintf("%.1fs", dur.Seconds())
	case dur < time.Hour:
		minutes := int(dur.Minutes())
		seconds := int(dur.Seconds()) % 60
		return fmt.Sprintf("%dm %ds", minutes, seconds)
	default:
		hours := int(dur.Hours())
		minutes := int(dur.Minutes()) % 60
		seconds := int(dur.Seconds()) % 60
		return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
	}
}
