package ignore

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type GitignoreEntry struct {
	Pattern   string
	IsNegated bool
}

func ParseGoGitignore(filename string) ([]GitignoreEntry, error) {
	// Read the contents of the gitignore file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Split the file contents by newline characters
	entries := strings.Split(string(data), "\n")

	// Parse each entry
	gitignoreEntries := []GitignoreEntry{}
	for _, entry := range entries {
		// Remove any leading or trailing whitespace from the entry
		entry = strings.TrimSpace(entry)
		var isNegated bool

		// handle empty line
		if len(entry) == 0 {
			continue
		}

		// Handle comment
		if entry[0] == '#' {
			continue
		}

		// Handle negation
		if entry[0] == '!' {
			entry = entry[1:]
			isNegated = true
		} else {
			isNegated = false
		}

		// Parse the pattern
		pattern := entry
		re := regexp.MustCompile(`\.go$`)
		if !re.MatchString(pattern) {
			return nil, fmt.Errorf("invalid gitignore pattern: %s", entry)
		}

		// Add the entry to the list
		gitignoreEntries = append(gitignoreEntries, GitignoreEntry{Pattern: pattern, IsNegated: isNegated})
	}

	return gitignoreEntries, nil
}

// NotIngoredGoFiles returns the list of files that are not ignored
func NotIngoredGoFiles(dirPath string, gitignoreEntries []GitignoreEntry) ([]string, error) {
	// Get the list of files in the directory
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	// Filter out files that should be ignored
	matchedFiles := []string{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePath := filepath.Join(dirPath, file.Name())

		if isFileIgnored := CheckFileMatches(filePath, gitignoreEntries); !isFileIgnored {
			matchedFiles = append(matchedFiles, filePath)
		}
	}

	return matchedFiles, nil
}

func CheckFileMatches(filePath string, gitignoreEntries []GitignoreEntry) bool {
	pattern := `*.go`

	if match, _ := filepath.Match(pattern, filePath); !match {
		return true // Not a go file, so ignore
	}

	isFileIgnored := false
	for _, entry := range gitignoreEntries {
		if entry.IsNegated {
			match, _ := filepath.Match(entry.Pattern, filePath)
			if match {
				isFileIgnored = false // Ignored pattern matched, so not ignored
			}
		} else {
			match, _ := filepath.Match(entry.Pattern, filePath)
			if match {
				isFileIgnored = true // Ignored pattern matched, so file is ignored
			}
		}
	}

	return isFileIgnored
}
