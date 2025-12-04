package validate

import (
	"path/filepath"
	"strings"

	"github.com/moistari/rls"
)

// DetectCategory detects the release category from a folder path
func DetectCategory(folderPath string) (string, error) {
	// Extract folder name from path
	folderName := filepath.Base(folderPath)

	// Remove trailing slash if present
	folderName = strings.TrimSuffix(folderName, "/")
	folderName = strings.TrimSuffix(folderName, "\\")

	// Parse using rls library
	release := rls.ParseString(folderName)

	// Get category type as string
	category := release.Type.String()

	// Handle empty/unknown categories
	if category == "" || category == "unknown" {
		return "", nil // Return empty string for unknown categories
	}

	return category, nil
}
