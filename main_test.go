package main

import (
	"os"
	"testing"
)

func TestOpenFile(t *testing.T) {
	// Test case: File exists
	t.Run("File exists", func(t *testing.T) {
		// Create a temporary file
		tmpFile, err := os.CreateTemp("", "test.txt")
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name())

		// Override the file name in openFile function
		file, err := os.Open(tmpFile.Name())
		if err != nil {
			t.Fatalf("Expected to open file, got error: %v", err)
		}
		defer file.Close()

		if file == nil {
			t.Error("Expected file to be opened, got nil")
		}
	})

	// Test case: File does not exist
	t.Run("File does not exist", func(t *testing.T) {
		file, err := os.Open("nonexistent.txt")
		if err == nil {
			t.Error("Expected error when opening nonexistent file, got nil")
		}

		if file != nil {
			t.Error("Expected file to be nil, got non-nil")
		}
	})
}

func TestCountWords(t *testing.T) {
	// Test case: File with multiple words
	t.Run("Multiple words", func(t *testing.T) {
		content := "Hello world this is a test"
		tmpFile, err := os.CreateTemp("", "test.txt")
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name())

		if _, err := tmpFile.WriteString(content); err != nil {
			t.Fatalf("Failed to write to temp file: %v", err)
		}
		tmpFile.Seek(0, 0)

		wordCount := countWords(tmpFile)
		expectedCount := 6
		if wordCount != expectedCount {
			t.Errorf("Expected %d words, got %d", expectedCount, wordCount)
		}
	})

	// Test case: Empty file
	t.Run("Empty file", func(t *testing.T) {
		tmpFile, err := os.CreateTemp("", "test.txt")
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name())

		wordCount := countWords(tmpFile)
		expectedCount := 0
		if wordCount != expectedCount {
			t.Errorf("Expected %d words, got %d", expectedCount, wordCount)
		}
	})

	// Test case: File with one word
	t.Run("One word", func(t *testing.T) {
		content := "Hello"
		tmpFile, err := os.CreateTemp("", "test.txt")
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name())

		if _, err := tmpFile.WriteString(content); err != nil {
			t.Fatalf("Failed to write to temp file: %v", err)
		}
		tmpFile.Seek(0, 0)

		wordCount := countWords(tmpFile)
		expectedCount := 1
		if wordCount != expectedCount {
			t.Errorf("Expected %d words, got %d", expectedCount, wordCount)
		}
	})

	// Test case: File with multiple lines
	t.Run("Multiple lines", func(t *testing.T) {
		content := "Hello world\nthis is a test"
		tmpFile, err := os.CreateTemp("", "test.txt")
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name())

		if _, err := tmpFile.WriteString(content); err != nil {
			t.Fatalf("Failed to write to temp file: %v", err)
		}
		tmpFile.Seek(0, 0)

		wordCount := countWords(tmpFile)
		expectedCount := 6
		if wordCount != expectedCount {
			t.Errorf("Expected %d words, got %d", expectedCount, wordCount)
		}
	})
}
