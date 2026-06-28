package specifications

import (
	"bufio"
	"fmt"
	"strings"
)

// AddHeaderNumbering takes raw markdown content and returns it with recursive numbering.
func AddHeaderNumbering(content string) string {
	var result strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(content))

	// Slice to hold counters for levels 1 through 6
	// counts[0] corresponds to H1, counts[1] to H2, etc.
	counts := make([]int, 6)

	// State to track if we are inside a code block
	inCodeBlock := false

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		// 1. Handle Code Blocks
		// We must ignore lines starting with # inside code blocks
		if strings.HasPrefix(trimmed, "```") {
			inCodeBlock = !inCodeBlock
		}

		// If we are in a code block or the line is empty, just write it and continue
		if inCodeBlock || trimmed == "" {
			result.WriteString(line + "\n")
			continue
		}

		// 2. Detect Header
		// Check if line starts with '#' and has a space afterwards
		if strings.HasPrefix(line, "#") {
			level := 0
			// Count the number of hashes
			for _, char := range line {
				if char == '#' {
					level++
				} else {
					break
				}
			}

			// Valid headers must be followed by a space (e.g. "# Title")
			// We also ensure we don't exceed H6
			if level > 0 && level <= 6 && len(line) > level && line[level] == ' ' {

				// Increment the counter for the current level
				// Note: level is 1-based, counts index is 0-based
				counts[level-1]++

				// Reset all sublevels to 0
				// e.g. If we just hit an H2, we reset H3, H4, H5, H6
				for i := level; i < 6; i++ {
					counts[i] = 0
				}

				// Build the number string (e.g., "1.2.1")
				var numStrBuilder strings.Builder
				for i := 0; i < level; i++ {
					if i > 0 {
						numStrBuilder.WriteString(".")
					}
					numStrBuilder.WriteString(fmt.Sprintf("%d", counts[i]))
				}

				// Reconstruct the line:
				// [Hashes] + [Space] + [Number] + [Original Text without Hashes]
				// Note: line[level:] includes the leading space, so we trim it slightly to normalize
				titleText := strings.TrimSpace(line[level:])
				newLine := fmt.Sprintf("%s %s %s", strings.Repeat("#", level), numStrBuilder.String(), titleText)

				result.WriteString(newLine + "\n")
				continue
			}
		}

		// 3. Default: write the line unchanged
		result.WriteString(line + "\n")
	}

	return result.String()
}
