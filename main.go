package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// Start measuring time
	start := time.Now()

	// Open the file
	file := openFile()
	// Close the file when the function returns
	defer file.Close()

	// Count the words in the file
	wordCount := countWords(file)

	// Print the results
	fmt.Printf("Number of words: %d\n", wordCount)
	fmt.Printf("Time taken: %v", time.Since(start))

}

func openFile() *os.File {
	// Open the file
	file, err := os.Open("test.txt")

	// Check for errors
	if err != nil {
		fmt.Println("Error opening file")
		return nil
	}
	// Return the file
	return file
}

func countWords(file *os.File) int {
	// Create a new scanner
	scanner := bufio.NewScanner(file)

	// Split the input into words
	scanner.Split(bufio.ScanWords)

	count := 0

	// Count the words
	for scanner.Scan() {
		count++
	}
	// Return the count
	return count
}
