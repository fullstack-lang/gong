package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// countLinesBySymbol counts the occurrences of a specific line pattern in a file
// and aggregates the counts by a captured symbol within that pattern.
// The target pattern is: __A__000000_00000 := (&models.A{}).Stage(stage)
// where 'A' is the symbol to be captured.
func countLinesBySymbol(filename string) (map[string]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	counts := make(map[string]int)
	// Regex to match the line and capture the symbol (e.g., "A")
	re := regexp.MustCompile(`__([A-Z]+)__\d{6}_\d{5} := \(&models\.([A-Z]+)\{\}\)\.Stage\(stage\)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			// The first submatch (index 1) is the captured symbol
			symbol := matches[1]
			counts[symbol]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %w", err)
	}

	return counts, nil
}
