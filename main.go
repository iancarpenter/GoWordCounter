package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	file := openFile()
	defer file.Close()

	wordCount := countWords(file)

	fmt.Printf("Number of words: %d\n", wordCount)
	fmt.Printf("Time taken: %v", time.Since(start))

}

func openFile() *os.File {
	file, err := os.Open("test.txt")

	if err != nil {
		fmt.Println("Error opening file")
		return nil
	}
	return file
}

func countWords(file *os.File) int {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}
