package preset

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Rule represents a single validation rule
type Rule struct {
	Pattern     string `yaml:"pattern"`
	Type        string `yaml:"type,omitempty"` // "file" (default) or "dir"
	Min         int    `yaml:"min,omitempty"`
	Max         int    `yaml:"max,omitempty"`
	Description string `yaml:"description,omitempty"`
	Regex       bool   `yaml:"regex,omitempty"` // If true, pattern is treated as regex instead of glob
}

// PresetConfig represents the entire preset configuration
type PresetConfig struct {
	SchemaVersion int               `yaml:"schema_version"`
	Rules         map[string][]Rule `yaml:"rules"`
}

// LoadPresets loads the preset configuration from a YAML file
func LoadPresets(presetPath string) (*PresetConfig, error) {
	// If no path provided, try default location
	if presetPath == "" {
		// Try to find presets.yaml in common locations
		possiblePaths := []string{
			"~/.config/sfvbrr/presets.yaml",
			"docs/presets.yaml",
			"presets.yaml",
			"./presets.yaml",
		}

		for _, path := range possiblePaths {
			if _, err := os.Stat(path); err == nil {
				presetPath = path
				break
			}
		}

		if presetPath == "" {
			return nil, fmt.Errorf("preset file not found in default locations")
		}
	}

	// Resolve absolute path
	absPath, err := filepath.Abs(presetPath)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve preset path: %w", err)
	}

	// Read file
	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read preset file: %w", err)
	}

	// Parse YAML
	var config PresetConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse preset file: %w", err)
	}

	return &config, nil
}

// GetRulesForCategory returns the rules for a specific category
func (c *PresetConfig) GetRulesForCategory(category string) ([]Rule, error) {
	rules, exists := c.Rules[category]
	if !exists {
		return nil, fmt.Errorf("no rules found for category: %s", category)
	}
	return rules, nil
}
