package models

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// CopyDirectory recursively copies a directory from src to dst.
//   - If dst exists, it will be REMOVED entirely before the copy starts.
//   - It attempts to preserve file permissions.
//   - It does NOT handle symbolic links specially; it will copy the target
//     of the link if it's a file, or traverse into it if it's a directory.
func CopyDirectory(src, dst string) error {
	// 1. Get source directory info
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to stat source directory %q: %w", src, err)
	}
	if !srcInfo.IsDir() {
		return fmt.Errorf("source %q is not a directory", src)
	}

	// 2. Check if destination exists. If it does, remove it first.
	// Use os.Lstat to correctly identify symlinks to directories if needed,
	// but os.Stat works fine here as os.RemoveAll handles both files and dirs.
	_, err = os.Stat(dst)
	if err == nil {
		// Destination exists, remove it recursively
		fmt.Printf("Destination %q exists, removing it first.\n", dst)
		if err = os.RemoveAll(dst); err != nil {
			return fmt.Errorf("failed to remove existing destination %q: %w", dst, err)
		}
	} else if !os.IsNotExist(err) {
		// An error occurred trying to stat dst, and it's *not* because it doesn't exist.
		// This could be a permission error, etc. Report it.
		return fmt.Errorf("failed to check destination %q: %w", dst, err)
	}
	// If err was os.IsNotExist, we proceed without removing anything.

	// 3. Create the top-level destination directory.
	// At this point, dst either never existed or was just removed.
	// Use the source directory's permissions.
	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		// Check if the error is because the target path exists as a file
		// This might happen in rare race conditions or if os.RemoveAll failed silently
		// (though unlikely). Double-check if needed, but MkdirAll should usually be robust.
		return fmt.Errorf("failed to create destination directory %q: %w", dst, err)
	}

	// 4. Walk through the source directory and copy contents
	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		// Handle walk errors (e.g., permission denied reading a subdir)
		if err != nil {
			// Skip errors related to the source root itself if needed,
			// although usually indicates a problem reading source content.
			return fmt.Errorf("error accessing path %q during walk: %w", path, err)
		}

		// Calculate the relative path from the source root
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path for %q: %w", path, err)
		}

		// Calculate the full destination path
		destPath := filepath.Join(dst, relPath)

		// --- Handle Directories ---
		if d.IsDir() {
			// Get info to retrieve permissions
			// Using d.Info() is generally efficient within WalkDir
			info, err := d.Info()
			if err != nil {
				return fmt.Errorf("failed to get info for directory %q: %w", path, err)
			}
			// Create the corresponding directory in the destination.
			// MkdirAll handles nested paths and permissions correctly here.
			// Skip creating the root dst again, though MkdirAll handles it.
			if destPath == dst {
				return nil // Already created the top-level dest directory
			}
			if err := os.MkdirAll(destPath, info.Mode()); err != nil {
				return fmt.Errorf("failed to create directory %q: %w", destPath, err)
			}
			return nil // Continue walking into the directory
		}

		// --- Handle Files ---
		// Skip if it's not a regular file
		if d.Type()&fs.ModeType != 0 {
			fmt.Printf("Skipping non-regular file: %s (type: %s)\n", path, d.Type())
			return nil
		}

		// Open the source file
		srcFile, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open source file %q: %w", path, err)
		}
		defer srcFile.Close() // Ensure file is closed

		// Create the destination file (truncates if somehow exists, but shouldn't)
		// Use os.OpenFile with flags for more control if needed, but Create is fine.
		destFile, err := os.Create(destPath)
		if err != nil {
			return fmt.Errorf("failed to create destination file %q: %w", destPath, err)
		}
		defer destFile.Close() // Ensure file is closed

		// Copy the contents
		bytesCopied, err := io.Copy(destFile, srcFile)
		if err != nil {
			return fmt.Errorf("failed to copy data from %q to %q: %w", path, destPath, err)
		}

		// Get source file info for permissions and size check
		info, err := d.Info()
		if err != nil {
			return fmt.Errorf("failed to get source file info %q after copy: %w", path, err)
		}

		// Verify all bytes were copied
		if bytesCopied != info.Size() {
			return fmt.Errorf("copy incomplete: copied %d bytes, expected %d for %q", bytesCopied, info.Size(), path)
		}

		// Set permissions on the destination file
		if err := os.Chmod(destPath, info.Mode()); err != nil {
			// Consider logging this as a warning instead of returning a fatal error
			// if partial permission setting failure is acceptable.
			return fmt.Errorf("failed to set permissions on %q: %w", destPath, err)
		}

		return nil // Continue walking
	})
}

// CopyFile copies a file from srcPath to destPath.
// It returns an error if any step of the process fails.
func CopyFile(srcPath, destPath string) error {
	// Open the source file for reading.
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file for writing.
	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Copy the contents of the source file to the destination file.
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	// Ensure all data is written to disk.
	err = destFile.Sync()
	if err != nil {
		return err
	}

	return nil
}
