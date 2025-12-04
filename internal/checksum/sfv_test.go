package checksum

import (
	"fmt"
	"hash/crc32"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFindSFVFile(t *testing.T) {
	// Create a temporary directory
	tmpDir := t.TempDir()

	// Test case: no SFV file
	_, err := FindSFVFile(tmpDir)
	if err == nil {
		t.Error("Expected error when no SFV file exists")
	}

	// Create an SFV file
	sfvPath := filepath.Join(tmpDir, "test.sfv")
	err = os.WriteFile(sfvPath, []byte("test.txt 12345678\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test SFV file: %v", err)
	}

	// Test case: SFV file exists
	found, err := FindSFVFile(tmpDir)
	if err != nil {
		t.Fatalf("Expected to find SFV file: %v", err)
	}
	if found != sfvPath {
		t.Errorf("Expected %s, got %s", sfvPath, found)
	}

	// Test case: case insensitive
	sfvPath2 := filepath.Join(tmpDir, "TEST.SFV")
	err = os.WriteFile(sfvPath2, []byte("test.txt 12345678\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test SFV file: %v", err)
	}

	found, err = FindSFVFile(tmpDir)
	if err != nil {
		t.Fatalf("Expected to find SFV file: %v", err)
	}
	// Should find one of them (implementation may vary)
	if found != sfvPath && found != sfvPath2 {
		t.Errorf("Expected to find either %s or %s, got %s", sfvPath, sfvPath2, found)
	}
}

func TestParseSFVFile(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a test SFV file
	sfvContent := `; This is a comment
file1.txt 12345678
file2.bin ABCDEF00
file with spaces.dat 00000000
; Another comment
file3.txt DEADBEEF
`
	sfvPath := filepath.Join(tmpDir, "test.sfv")
	err := os.WriteFile(sfvPath, []byte(sfvContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test SFV file: %v", err)
	}

	sfv, err := ParseSFVFile(sfvPath)
	if err != nil {
		t.Fatalf("Failed to parse SFV file: %v", err)
	}

	if len(sfv.Entries) != 4 {
		t.Errorf("Expected 4 entries, got %d", len(sfv.Entries))
	}

	// Check first entry
	if sfv.Entries[0].Filename != "file1.txt" {
		t.Errorf("Expected filename 'file1.txt', got '%s'", sfv.Entries[0].Filename)
	}
	if sfv.Entries[0].Checksum != "12345678" {
		t.Errorf("Expected checksum '12345678', got '%s'", sfv.Entries[0].Checksum)
	}

	// Check entry with spaces
	if sfv.Entries[2].Filename != "file with spaces.dat" {
		t.Errorf("Expected filename 'file with spaces.dat', got '%s'", sfv.Entries[2].Filename)
	}
}

func TestValidateSFV_FileNotFound(t *testing.T) {
	tmpDir := t.TempDir()

	// Create an SFV file with a non-existent file
	sfvContent := "nonexistent.txt 12345678\n"
	sfvPath := filepath.Join(tmpDir, "test.sfv")
	err := os.WriteFile(sfvPath, []byte(sfvContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test SFV file: %v", err)
	}

	sfv, err := ParseSFVFile(sfvPath)
	if err != nil {
		t.Fatalf("Failed to parse SFV file: %v", err)
	}

	result, err := ValidateSFV(sfv, DefaultOptions())
	if err != nil {
		t.Fatalf("Failed to validate SFV: %v", err)
	}

	if result.MissingFiles != 1 {
		t.Errorf("Expected 1 missing file, got %d", result.MissingFiles)
	}
	if result.ValidFiles != 0 {
		t.Errorf("Expected 0 valid files, got %d", result.ValidFiles)
	}
}

func TestValidateSFV_ValidFile(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a test file
	testFile := filepath.Join(tmpDir, "test.txt")
	testContent := []byte("Hello, World!")
	err := os.WriteFile(testFile, testContent, 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Compute expected CRC-32
	expectedCRC := computeCRC32ForContent(testContent)

	// Create SFV file with correct checksum
	sfvContent := "test.txt " + expectedCRC + "\n"
	sfvPath := filepath.Join(tmpDir, "test.sfv")
	err = os.WriteFile(sfvPath, []byte(sfvContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test SFV file: %v", err)
	}

	sfv, err := ParseSFVFile(sfvPath)
	if err != nil {
		t.Fatalf("Failed to parse SFV file: %v", err)
	}

	result, err := ValidateSFV(sfv, DefaultOptions())
	if err != nil {
		t.Fatalf("Failed to validate SFV: %v", err)
	}

	if result.ValidFiles != 1 {
		t.Errorf("Expected 1 valid file, got %d", result.ValidFiles)
	}
	if result.InvalidFiles != 0 {
		t.Errorf("Expected 0 invalid files, got %d", result.InvalidFiles)
	}
	if result.MissingFiles != 0 {
		t.Errorf("Expected 0 missing files, got %d", result.MissingFiles)
	}
}

// Helper function to compute CRC-32 for content
func computeCRC32ForContent(content []byte) string {
	hash := crc32.NewIEEE()
	hash.Write(content)
	return strings.ToUpper(fmt.Sprintf("%08x", hash.Sum32()))
}

