package checksum

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindSFVFilesRecursive(t *testing.T) {
	tmpDir := t.TempDir()

	// Create directory structure
	subDir1 := filepath.Join(tmpDir, "subdir1")
	subDir2 := filepath.Join(tmpDir, "subdir2")
	nestedDir := filepath.Join(subDir1, "nested")

	err := os.MkdirAll(nestedDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create directories: %v", err)
	}

	err = os.MkdirAll(subDir2, 0755)
	if err != nil {
		t.Fatalf("Failed to create subdir2: %v", err)
	}

	// Create SFV files at different levels
	sfv1 := filepath.Join(tmpDir, "root.sfv")
	sfv2 := filepath.Join(subDir1, "sub1.sfv")
	sfv3 := filepath.Join(subDir2, "sub2.sfv")
	sfv4 := filepath.Join(nestedDir, "nested.sfv")

	sfvFiles := []string{sfv1, sfv2, sfv3, sfv4}
	for _, sfv := range sfvFiles {
		err = os.WriteFile(sfv, []byte("test.txt 12345678\n"), 0644)
		if err != nil {
			t.Fatalf("Failed to create SFV file %s: %v", sfv, err)
		}
	}

	// Test recursive search
	found, err := FindSFVFilesRecursive(tmpDir)
	if err != nil {
		t.Fatalf("Failed to find SFV files recursively: %v", err)
	}

	if len(found) != 4 {
		t.Errorf("Expected 4 SFV files, got %d", len(found))
	}

	// Verify all files are found
	foundMap := make(map[string]bool)
	for _, f := range found {
		foundMap[f] = true
	}

	for _, expected := range sfvFiles {
		if !foundMap[expected] {
			t.Errorf("Expected to find SFV file: %s", expected)
		}
	}
}

func TestFindSFVFilesRecursive_CaseInsensitive(t *testing.T) {
	tmpDir := t.TempDir()

	// Create SFV files with different case extensions
	sfv1 := filepath.Join(tmpDir, "test.SFV")
	sfv2 := filepath.Join(tmpDir, "test2.sfv")
	sfv3 := filepath.Join(tmpDir, "test3.Sfv")

	sfvFiles := []string{sfv1, sfv2, sfv3}
	for _, sfv := range sfvFiles {
		err := os.WriteFile(sfv, []byte("test.txt 12345678\n"), 0644)
		if err != nil {
			t.Fatalf("Failed to create SFV file %s: %v", sfv, err)
		}
	}

	// Test recursive search
	found, err := FindSFVFilesRecursive(tmpDir)
	if err != nil {
		t.Fatalf("Failed to find SFV files recursively: %v", err)
	}

	if len(found) != 3 {
		t.Errorf("Expected 3 SFV files (case insensitive), got %d", len(found))
	}
}

func TestValidateFolders_Recursive(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a subdirectory with an SFV file
	subDir := filepath.Join(tmpDir, "subdir")
	err := os.MkdirAll(subDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}

	// Create a test file in subdirectory
	testFile := filepath.Join(subDir, "test.txt")
	testContent := []byte("Hello, World!")
	err = os.WriteFile(testFile, testContent, 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Compute expected CRC-32
	expectedCRC := computeCRC32ForContent(testContent)

	// Create SFV file in subdirectory
	sfvPath := filepath.Join(subDir, "test.sfv")
	sfvContent := "test.txt " + expectedCRC + "\n"
	err = os.WriteFile(sfvPath, []byte(sfvContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create SFV file: %v", err)
	}

	// Test recursive validation
	opts := Options{
		Recursive:    true,
		Quiet:        true, // Use quiet mode to avoid output in tests
		OutputFormat: OutputFormatText,
	}

	err = ValidateFolders([]string{tmpDir}, opts)
	if err != nil {
		t.Errorf("Expected validation to succeed, got error: %v", err)
	}
}
