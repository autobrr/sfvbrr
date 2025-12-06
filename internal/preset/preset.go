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

// CategoryRules represents rules and settings for a category
type CategoryRules struct {
	DenyUnexpected bool   `yaml:"deny_unexpected"`
	Rules          []Rule `yaml:"rules"`
}

// PresetConfig represents the entire preset configuration
type PresetConfig struct {
	SchemaVersion int                       `yaml:"schema_version"`
	Rules         map[string]*CategoryRules `yaml:"rules"`
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

	// First, parse as raw structure to validate required fields
	var rawConfig struct {
		Rules map[string]interface{} `yaml:"rules"`
	}
	if err := yaml.Unmarshal(data, &rawConfig); err != nil {
		return nil, fmt.Errorf("failed to parse preset file: %w", err)
	}

	// Validate that deny_unexpected is present in all categories
	for category, rawValue := range rawConfig.Rules {
		if rawValue == nil {
			return nil, fmt.Errorf("category %q has no configuration", category)
		}

		// Check if it's a map (new format)
		if categoryMap, ok := rawValue.(map[string]interface{}); ok {
			if _, exists := categoryMap["deny_unexpected"]; !exists {
				return nil, fmt.Errorf("category %q is missing required field 'deny_unexpected'", category)
			}
		} else {
			// If it's not a map, it's invalid format
			return nil, fmt.Errorf("category %q has invalid format: expected map with 'deny_unexpected' and 'rules'", category)
		}
	}

	// Now parse into the actual config structure
	var config PresetConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse preset file: %w", err)
	}

	return &config, nil
}

// GetRulesForCategory returns the rules for a specific category
func (c *PresetConfig) GetRulesForCategory(category string) ([]Rule, error) {
	catRules, exists := c.Rules[category]
	if !exists {
		return nil, fmt.Errorf("no rules found for category: %s", category)
	}
	return catRules.Rules, nil
}

// GetDenyUnexpected returns whether unexpected files should be denied for a category
func (c *PresetConfig) GetDenyUnexpected(category string) bool {
	catRules, exists := c.Rules[category]
	if !exists {
		return false
	}
	return catRules.DenyUnexpected
}
