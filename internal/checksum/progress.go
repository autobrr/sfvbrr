package checksum

// ProgressTracker tracks progress of SFV validation
// This is a placeholder for future progress tracking implementation
type ProgressTracker struct {
	Total   int
	Current int
}

// NewProgressTracker creates a new progress tracker
func NewProgressTracker(total int) *ProgressTracker {
	return &ProgressTracker{
		Total:   total,
		Current: 0,
	}
}

// Update updates the current progress
func (pt *ProgressTracker) Update(current int) {
	pt.Current = current
}

