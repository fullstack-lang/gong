package ignore

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestGitignore(t *testing.T) {
	// Define the test data
	gitignoreContent := `
	# Comment
	*.go
	!a.go
	`

	// Create a temporary file to store the gitignore content
	tmpFile, err := ioutil.TempFile("", "test_gitignore")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write the gitignore content to the temporary file
	if _, err := tmpFile.WriteString(gitignoreContent); err != nil {
		t.Fatalf("Failed to write gitignore content: %v", err)
	}

	// Parse the gitignore file
	gitignoreEntries, err := ParseGoGitignore(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to parse gitignore: %v", err)
	}

	// Parse the directory
	actualFiles, err := NotIngoredGoFiles(".", gitignoreEntries)
	if err != nil {
		t.Fatalf("Failed to parse directory: %v", err)
	}

	// Define the test data
	matchedFiles := []string{"a.go"}

	// Check that the expected files were matched
	if !reflect.DeepEqual(actualFiles, matchedFiles) {
		t.Errorf("Incorrect matched files: expected %v, got %v", matchedFiles, actualFiles)
	}

}
