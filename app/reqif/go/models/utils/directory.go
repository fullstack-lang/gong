package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// RemoveAllContents clears all files and subdirectories within a given directory,
// but leaves the directory itself intact.
func RemoveAllContents(dir string) error {
	// Open the directory to read its entries.
	d, err := os.Open(dir)
	if err != nil {
		return fmt.Errorf("failed to open directory: %w", err)
	}
	defer d.Close()

	// Read all the file and subdirectory names.
	names, err := d.Readdirnames(-1)
	if err != nil {
		return fmt.Errorf("failed to read directory contents: %w", err)
	}

	// Iterate over the names and remove each entry.
	for _, name := range names {
		// Construct the full path to the entry.
		fullPath := filepath.Join(dir, name)

		// os.RemoveAll can be used here to recursively delete subdirectories.
		err = os.RemoveAll(fullPath)
		if err != nil {
			// Continue trying to remove other files even if one fails.
			// You might want to collect these errors and return them.
			fmt.Printf("Failed to remove %s: %v\n", fullPath, err)
		}
	}
	return nil
}
