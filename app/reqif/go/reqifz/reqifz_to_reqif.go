package reqifz

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ConvertReQIFzToReqIF converts a .reqifz file to .reqif by unzipping it
// Returns true if conversion was performed, false if no action was taken
func ConvertReQIFzToReqIF(filePath string) (bool, error) {
	// Get the file extension
	ext := strings.ToLower(filepath.Ext(filePath))

	// If the file has .reqif extension, do nothing
	if ext == ".reqif" {
		return false, nil
	}

	// If the file has .reqifz extension, unzip it
	if ext == ".reqifz" {
		// Create the output file path by replacing .reqifz with .reqif
		outputPath := strings.TrimSuffix(filePath, ".reqifz") + ".reqif"

		err := unzipReqifz(filePath, outputPath)
		if err != nil {
			return false, fmt.Errorf("failed to unzip reqifz file: %w", err)
		}

		return true, nil
	}

	// If neither .reqif nor .reqifz, return error
	return false, fmt.Errorf("unsupported file extension: %s (expected .reqif or .reqifz)", ext)
}

// unzipReqifz extracts the content from a .reqifz file and writes it to the output path
func unzipReqifz(reqifzPath, outputPath string) error {
	// Open the zip file
	reader, err := zip.OpenReader(reqifzPath)
	if err != nil {
		return fmt.Errorf("failed to open reqifz file: %w", err)
	}
	defer reader.Close()

	// ReQIFz files typically contain a single .reqif file
	// Find the first .reqif file in the archive
	var reqifFile *zip.File
	for _, file := range reader.File {
		if strings.ToLower(filepath.Ext(file.Name)) == ".reqif" {
			reqifFile = file
			break
		}
	}

	if reqifFile == nil {
		return fmt.Errorf("no .reqif file found in the reqifz archive")
	}

	// Open the file inside the zip
	rc, err := reqifFile.Open()
	if err != nil {
		return fmt.Errorf("failed to open reqif file inside archive: %w", err)
	}
	defer rc.Close()

	// Create the output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	// Copy the content
	_, err = io.Copy(outFile, rc)
	if err != nil {
		return fmt.Errorf("failed to copy file content: %w", err)
	}

	return nil
}
