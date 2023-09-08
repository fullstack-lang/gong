package main

import (
	"bufio"
	"os"
	"strings"
)

// removeDuplicateLines takes a Typescript file path as input and rewrites the file without duplicate lines
// if the line contains patterns
func removeDuplicateLinesInFile(filePath string, pattern string) error {
	// open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// create a set to store the lines already seen
	seen := make(map[string]bool)

	// create a slice to store the lines without duplicates
	result := []string{}

	// loop through the file line by line
	for scanner.Scan() {
		// get the current line
		line := scanner.Text()

		// add non duplicate lines (check if the line is in the set)
		if !seen[line] {
			// if not, add the line to the set and the slice
			seen[line] = true
			result = append(result, line)
		} else {
			// if it is in the set, do not add it if does not have the pattern
			if !strings.Contains(line, pattern) {
				seen[line] = true
				result = append(result, line)
			}
		}
	}

	// check if there is an error during scanning
	if err := scanner.Err(); err != nil {
		return err
	}

	// open the file for writing (truncate it)
	file, err = os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// create a writer to write to the file
	writer := bufio.NewWriter(file)

	// loop through the slice of lines without duplicates
	for _, line := range result {
		// write the line to the file with a newline character
		writer.WriteString(line + "\n")
	}

	// flush the writer buffer
	writer.Flush()

	// return nil to indicate no error
	return nil
}
