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

	countWords(file)

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

func countWords(file *os.File) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	fmt.Printf("Number of words: %d\n", count)
}
