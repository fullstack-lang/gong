// generated code - do not edit
package models

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// ComputeDiff generates a git-style unified diff from string a to string b
func ComputeDiff(a, b string) string {
	linesA := strings.Split(a, "\n")
	linesB := strings.Split(b, "\n")

	// Simple LCS-based diff algorithm
	diff := computeLCS(linesA, linesB)
	return formatUnifiedDiff(diff)
}

// ApplyDiff applies a git-style unified diff c to string b to reconstruct string a
func ApplyDiff(b, c string) (string, error) {
	linesB := strings.Split(b, "\n")

	// Parse the diff
	hunks, err := parseDiff(c)
	if err != nil {
		return "", err
	}

	// Apply hunks in reverse order to avoid index shifting issues
	result := make([]string, len(linesB))
	copy(result, linesB)

	for i := len(hunks) - 1; i >= 0; i-- {
		hunk := hunks[i]
		err := applyHunk(result, hunk)
		if err != nil {
			return "", err
		}
		// Update result slice if length changed
		if hunk.NewLines != hunk.OldLines {
			newResult := make([]string, len(result)+(hunk.OldLines-hunk.NewLines))
			copy(newResult, result)
			result = newResult
		}
	}

	return strings.Join(result, "\n"), nil
}

// DiffHunk represents a single hunk in a unified diff
type DiffHunk struct {
	OldStart int
	OldLines int
	NewStart int
	NewLines int
	Lines    []string
}

// DiffLine represents a line in a diff with its operation
type DiffOp struct {
	Type string // "context", "delete", "add"
	Text string
}

func computeLCS(a, b []string) []DiffOp {
	m, n := len(a), len(b)

	// Create LCS table
	lcs := make([][]int, m+1)
	for i := range lcs {
		lcs[i] = make([]int, n+1)
	}

	// Fill LCS table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = diff_max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}

	// Backtrack to find diff
	var result []DiffOp
	i, j := m, n

	for i > 0 || j > 0 {
		if i > 0 && j > 0 && a[i-1] == b[j-1] {
			result = append([]DiffOp{{Type: "context", Text: a[i-1]}}, result...)
			i--
			j--
		} else if i > 0 && (j == 0 || lcs[i-1][j] >= lcs[i][j-1]) {
			result = append([]DiffOp{{Type: "delete", Text: a[i-1]}}, result...)
			i--
		} else if j > 0 {
			result = append([]DiffOp{{Type: "add", Text: b[j-1]}}, result...)
			j--
		}
	}

	return result
}

func formatUnifiedDiff(ops []DiffOp) string {
	if len(ops) == 0 {
		return ""
	}

	var result strings.Builder
	result.WriteString("--- a\n")
	result.WriteString("+++ b\n")

	// Group operations into hunks
	hunks := groupIntoHunks(ops)

	for _, hunk := range hunks {
		// Write hunk header
		result.WriteString(fmt.Sprintf("@@ -%d,%d +%d,%d @@\n",
			hunk.OldStart, hunk.OldLines, hunk.NewStart, hunk.NewLines))

		// Write hunk lines
		for _, line := range hunk.Lines {
			result.WriteString(line + "\n")
		}
	}

	return result.String()
}

func groupIntoHunks(ops []DiffOp) []DiffHunk {
	var hunks []DiffHunk
	var currentHunk *DiffHunk

	oldLine, newLine := 1, 1
	contextBefore := 3
	contextAfter := 3

	for i, op := range ops {
		if op.Type == "context" {
			// Check if we need to start a new hunk
			if currentHunk == nil {
				// Look ahead to see if there are changes coming
				hasChanges := false
				for j := i + 1; j < len(ops) && j < i+contextBefore+1; j++ {
					if ops[j].Type != "context" {
						hasChanges = true
						break
					}
				}

				if hasChanges {
					currentHunk = &DiffHunk{
						OldStart: oldLine,
						NewStart: newLine,
					}
				}
			}

			if currentHunk != nil {
				currentHunk.Lines = append(currentHunk.Lines, " "+op.Text)
				currentHunk.OldLines++
				currentHunk.NewLines++
			}

			oldLine++
			newLine++
		} else {
			// Ensure we have a current hunk for changes
			if currentHunk == nil {
				// Include some context before
				contextStart := diff_max(0, i-contextBefore)
				currentHunk = &DiffHunk{
					OldStart: oldLine - (i - contextStart),
					NewStart: newLine - (i - contextStart),
				}

				// Add context lines
				for j := contextStart; j < i; j++ {
					if ops[j].Type == "context" {
						currentHunk.Lines = append(currentHunk.Lines, " "+ops[j].Text)
						currentHunk.OldLines++
						currentHunk.NewLines++
					}
				}
			}

			switch op.Type {
			case "delete":
				currentHunk.Lines = append(currentHunk.Lines, "-"+op.Text)
				currentHunk.OldLines++
				oldLine++
			case "add":
				currentHunk.Lines = append(currentHunk.Lines, "+"+op.Text)
				currentHunk.NewLines++
				newLine++
			}
		}

		// Check if we should close the current hunk
		if currentHunk != nil {
			contextCount := 0
			for j := i + 1; j < len(ops) && j < i+contextAfter+1; j++ {
				if ops[j].Type == "context" {
					contextCount++
				} else {
					break
				}
			}

			if contextCount == contextAfter || i == len(ops)-1 {
				// Add trailing context
				for j := i + 1; j < len(ops) && j < i+contextAfter+1; j++ {
					if ops[j].Type == "context" {
						currentHunk.Lines = append(currentHunk.Lines, " "+ops[j].Text)
						currentHunk.OldLines++
						currentHunk.NewLines++
					}
				}

				hunks = append(hunks, *currentHunk)
				currentHunk = nil
			}
		}
	}

	return hunks
}

func parseDiff(diff string) ([]DiffHunk, error) {
	var hunks []DiffHunk
	scanner := bufio.NewScanner(strings.NewReader(diff))

	// Skip header lines
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "@@") {
			// Parse hunk header
			hunk, err := parseHunkHeader(line)
			if err != nil {
				return nil, err
			}

			// Read hunk lines
			for scanner.Scan() {
				line := scanner.Text()
				if strings.HasPrefix(line, "@@") {
					// Next hunk, put back the line
					scanner = bufio.NewScanner(strings.NewReader(line + "\n" + readRemainingLines(scanner)))
					break
				}
				hunk.Lines = append(hunk.Lines, line)
			}

			hunks = append(hunks, hunk)
		}
	}

	return hunks, nil
}

func parseHunkHeader(line string) (DiffHunk, error) {
	// Parse "@@" -oldStart,oldLines +newStart,newLines "@@"
	parts := strings.Fields(line)
	if len(parts) < 4 {
		return DiffHunk{}, fmt.Errorf("invalid hunk header: %s", line)
	}

	oldPart := strings.TrimPrefix(parts[1], "-")
	newPart := strings.TrimPrefix(parts[2], "+")

	oldStart, oldLines, err := parseRange(oldPart)
	if err != nil {
		return DiffHunk{}, err
	}

	newStart, newLines, err := parseRange(newPart)
	if err != nil {
		return DiffHunk{}, err
	}

	return DiffHunk{
		OldStart: oldStart,
		OldLines: oldLines,
		NewStart: newStart,
		NewLines: newLines,
	}, nil
}

func parseRange(rangeStr string) (int, int, error) {
	if strings.Contains(rangeStr, ",") {
		parts := strings.Split(rangeStr, ",")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, 0, err
		}
		lines, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, 0, err
		}
		return start, lines, nil
	} else {
		start, err := strconv.Atoi(rangeStr)
		if err != nil {
			return 0, 0, err
		}
		return start, 1, nil
	}
}

func applyHunk(lines []string, hunk DiffHunk) error {
	// This is a simplified reverse application
	// In practice, you'd want more robust handling
	oldIdx := hunk.OldStart - 1

	for _, line := range hunk.Lines {
		if len(line) == 0 {
			continue
		}

		switch line[0] {
		case ' ': // context line
			oldIdx++
		case '+': // line to remove (reverse of add)
			if oldIdx >= len(lines) {
				return fmt.Errorf("index out of range")
			}
			lines = append(lines[:oldIdx], lines[oldIdx+1:]...)
		case '-': // line to add (reverse of delete)
			text := line[1:]
			if oldIdx >= len(lines) {
				lines = append(lines, text)
			} else {
				lines = append(lines[:oldIdx], append([]string{text}, lines[oldIdx:]...)...)
			}
			oldIdx++
		}
	}

	return nil
}

func readRemainingLines(scanner *bufio.Scanner) string {
	var result strings.Builder
	for scanner.Scan() {
		result.WriteString(scanner.Text() + "\n")
	}
	return result.String()
}

func diff_max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
