package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// Record the start time
	start := time.Now()

	// Open the file and defer its closure
	file, err := openFile("test.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Count the words in the file
	wordCount, err := countWords(file)
	if err != nil {
		log.Fatalf("Failed to count words: %v", err)
	}

	// Print the total number of words and the time taken to execute the program
	fmt.Printf("Number of words: %d\n", wordCount)
	fmt.Printf("Time taken: %v\n", time.Since(start))
}

// openFile opens the specified file and returns a file pointer
func openFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	return file, nil
}

// countWords counts the number of words in the given file
func countWords(file *os.File) (int, error) {
	// Create a new scanner
	scanner := bufio.NewScanner(file)

	// Split the input into words
	scanner.Split(bufio.ScanWords)

	count := 0

	// Count the words
	for scanner.Scan() {
		count++
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error scanning file: %w", err)
	}

	// Return the count
	return count, nil
}
