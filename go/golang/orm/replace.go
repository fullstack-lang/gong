package orm

import (
	"bytes"
	"io"
	"os"
)

func ReplaceInFile(filePath string, toReplace string, replacement string) error {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the content from the file
	var content bytes.Buffer
	if _, err := io.Copy(&content, file); err != nil {
		return err
	}

	// Replace the string
	newContent := bytes.ReplaceAll(content.Bytes(), []byte(toReplace), []byte(replacement))

	// Open the file for writing
	file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the new content back to the file
	if _, err := file.Write(newContent); err != nil {
		return err
	}

	return nil
}
