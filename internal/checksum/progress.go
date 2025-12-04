package checksum

import (
	"sync"
	"time"
)

// Displayer defines the interface for displaying progress during SFV validation
type Displayer interface {
	ShowProgress(total int)
	UpdateProgress(completed int, rate float64)
	ShowFiles(entries []SFVEntry, numWorkers int)
	FinishProgress()
	IsBatch() bool
	SetQuiet(quiet bool)
}

// ProgressTracker tracks progress of SFV validation with timing information
type ProgressTracker struct {
	Total      int
	Current    int
	StartTime  time.Time
	mu         sync.Mutex
	lastUpdate time.Time
	lastCount  int
}

// NewProgressTracker creates a new progress tracker
func NewProgressTracker(total int) *ProgressTracker {
	now := time.Now()
	return &ProgressTracker{
		Total:      total,
		Current:    0,
		StartTime:  now,
		lastUpdate: now,
		lastCount:  0,
	}
}

// Update updates the current progress
func (pt *ProgressTracker) Update(current int) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.Current = current
	pt.lastUpdate = time.Now()
}

// GetProgress returns the current progress percentage
func (pt *ProgressTracker) GetProgress() float64 {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	if pt.Total == 0 {
		return 0
	}
	return float64(pt.Current) / float64(pt.Total) * 100
}

// GetElapsed returns the elapsed time since the tracker was created
func (pt *ProgressTracker) GetElapsed() time.Duration {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	return time.Since(pt.StartTime)
}

// GetRate calculates the current processing rate (items per second)
func (pt *ProgressTracker) GetRate() float64 {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	elapsed := time.Since(pt.lastUpdate).Seconds()
	if elapsed < 0.1 {
		// Use overall rate if update was too recent
		elapsed = time.Since(pt.StartTime).Seconds()
		if elapsed < 0.1 {
			return 0
		}
		return float64(pt.Current) / elapsed
	}
	delta := pt.Current - pt.lastCount
	if delta <= 0 {
		return 0
	}
	return float64(delta) / elapsed
}

// GetETA estimates the time remaining based on current rate
func (pt *ProgressTracker) GetETA() time.Duration {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	remaining := pt.Total - pt.Current
	if remaining <= 0 {
		return 0
	}
	rate := pt.GetRate()
	if rate <= 0 {
		return 0
	}
	seconds := float64(remaining) / rate
	return time.Duration(seconds) * time.Second
}
