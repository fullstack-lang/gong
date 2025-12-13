package models

import "strings"

// WrapString wraps a string at a specified character limit.
// It inserts newlines (\n) to ensure no line exceeds the cutoff.
// It attempts to break lines at spaces (word boundaries).
func WrapString(s string, cutoff int) string {
	if cutoff <= 0 {
		return s // Return original string if cutoff is invalid
	}

	var b strings.Builder
	b.Grow(len(s)) // Pre-allocate capacity for efficiency

	currentLineLen := 0
	words := strings.Fields(s) // Split the string into words

	for _, word := range words {
		wordLen := len(word)

		// If this word alone is longer than the cutoff,
		// we have to break it, or just let it overflow.
		// For simplicity, this version just lets it overflow.
		// A more complex version might break mid-word.

		// Check if adding this word (plus a space if needed) exceeds the cutoff
		if currentLineLen > 0 && currentLineLen+wordLen+1 > cutoff {
			// Time to wrap. Add a newline.
			b.WriteRune('\n')
			currentLineLen = 0
		}

		// Add a space before the word if it's not the start of a new line
		if currentLineLen > 0 {
			b.WriteRune(' ')
			currentLineLen++
		}

		// Add the word itself
		b.WriteString(word)
		currentLineLen += wordLen

		// This handles the edge case where the original string
		// might have multiple spaces. strings.Fields() removes them,
		// but we might want to handle trailing newlines from the input.
		// For this implementation, we just rebuild based on words.
	}

	return b.String()
}
