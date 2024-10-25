package orm

import (
	"bufio"
	"os"
	"strings"
)

const removeKeyWord = "THIS IS REMOVED BY GONG COMPILER IF TARGET IS"
const Gorm = " gorm"
const Lite = " lite"
const firstCommentLinePrefix = "/* "
const lastCommentLinePostfix = " */"

const gormFirstLineToBeRemoved = firstCommentLinePrefix +
	removeKeyWord +
	Gorm

const gormLastLineToBeRemoved = removeKeyWord +
	Gorm +
	lastCommentLinePostfix

const liteFirstLineToBeRemoved = firstCommentLinePrefix +
	removeKeyWord +
	Lite

const liteLastLineToBeRemoved = removeKeyWord +
	Lite +
	lastCommentLinePostfix

// RemoveTargetedLines removes lines between specific comment markers based on the target.
func RemoveTargetedLines(path string, target string) error {
	// Open the file for reading
	inputFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// Prepare the start and end markers
	startMarker := firstCommentLinePrefix + removeKeyWord + target
	endMarker := removeKeyWord + target + lastCommentLinePostfix

	// Read the file line by line
	scanner := bufio.NewScanner(inputFile)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, startMarker) {
			continue // Skip the line containing the marker
		}
		if strings.Contains(line, endMarker) {
			continue // Skip the line containing the marker
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// Write the updated content back to the file
	outputContent := strings.Join(lines, "\n")
	err = os.WriteFile(path, []byte(outputContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
