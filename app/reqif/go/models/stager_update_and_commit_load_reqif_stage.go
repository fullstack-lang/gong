package models

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
)

type FileToUploadProxy struct {
	stager *Stager
}

func (fileToUploadProxy *FileToUploadProxy) OnFileUpload(uploadedFile *load.FileToUpload) error {
	fileName := uploadedFile.GetName()
	fileExt := strings.ToLower(filepath.Ext(fileName))

	var contentToProcess []byte
	var svgImages []*EmbeddedSvgImage
	var jpgImages []*EmbeddedJpgImage
	var pngImages []*EmbeddedPngImage
	var err error

	// Direct processing for .reqif files
	decodedBytes, err := base64.StdEncoding.DecodeString(uploadedFile.Base64EncodedContent)
	if err != nil {
		return fmt.Errorf("base64.StdEncoding.DecodeString failed: %w", err)
	}

	switch fileExt {
	case ".reqif":

		contentToProcess = decodedBytes

	case ".reqifz":
		// Unzip and extract .reqif content for .reqifz files
		contentToProcess, svgImages, jpgImages, pngImages, err = extractReqifFromZip(decodedBytes)
		if err != nil {
			return fmt.Errorf("failed to extract REQIF from zip: %w", err)
		}
	default:
		return fmt.Errorf("unsupported file extension: %s", fileExt)
	}

	// Process the content (either direct .reqif or extracted from .reqifz)
	fileToUploadProxy.stager.processReqifData(contentToProcess, svgImages, jpgImages, pngImages, fileName)
	return nil
}

// extractReqifFromZip extracts the first .reqif file found in the zip archive
// and all svg files.
func extractReqifFromZip(zipData []byte) ([]byte, []*EmbeddedSvgImage, []*EmbeddedJpgImage, []*EmbeddedPngImage, error) {

	var svgImages []*EmbeddedSvgImage
	var jpgImages []*EmbeddedJpgImage
	var pngImages []*EmbeddedPngImage
	var reqifContent []byte

	// Create a reader from the zip data
	zipReader, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if err != nil {
		return nil, svgImages, jpgImages, pngImages, fmt.Errorf("failed to create zip reader: %w", err)
	}

	// Look for .reqif files in the archive
	for _, file := range zipReader.File {
		switch strings.ToLower(filepath.Ext(file.Name)) {
		case ".reqif":
			// Open the file
			rc, err := file.Open()
			if err != nil {
				return nil, svgImages, jpgImages, pngImages, fmt.Errorf("failed to open file %s in zip: %w", file.Name, err)
			}
			defer rc.Close()

			// Read the content
			content, err := io.ReadAll(rc)
			if err != nil {
				return nil, svgImages, jpgImages, pngImages, fmt.Errorf("failed to read file %s from zip: %w", file.Name, err)
			}
			reqifContent = content
		case ".svg":
			// Open the file
			rc, err := file.Open()
			if err != nil {
				return nil, svgImages, jpgImages, pngImages, fmt.Errorf("failed to open file %s in zip: %w", file.Name, err)
			}
			defer rc.Close()

			// Read the content
			content, err := io.ReadAll(rc)
			if err != nil {
				return nil, svgImages, jpgImages, pngImages, fmt.Errorf("failed to read file %s from zip: %w", file.Name, err)
			}

			svgImage := &EmbeddedSvgImage{
				Name:    file.Name,
				Content: string(content),
			}
			svgImages = append(svgImages, svgImage)
		case ".jpg":
			// Open the file
			rc, err := file.Open()
			if err != nil {
				return nil, svgImages, jpgImages, pngImages, fmt.Errorf("failed to open file %s in zip: %w", file.Name, err)
			}
			defer rc.Close()

			// Read the content
			content, err := io.ReadAll(rc)
			if err != nil {
				return nil, svgImages, jpgImages, pngImages, fmt.Errorf("failed to read file %s from zip: %w", file.Name, err)
			}

			jpgImage := &EmbeddedJpgImage{
				Name: file.Name,
				// Encode the raw byte content to a Base64 string
				Base64Content: base64.StdEncoding.EncodeToString(content),
			}
			jpgImages = append(jpgImages, jpgImage)
		case ".png":
			// Open the file
			rc, err := file.Open()
			if err != nil {
				return nil, svgImages, jpgImages, pngImages, fmt.Errorf("failed to open file %s in zip: %w", file.Name, err)
			}
			defer rc.Close()

			// Read the content
			content, err := io.ReadAll(rc)
			if err != nil {
				return nil, svgImages, jpgImages, pngImages, fmt.Errorf("failed to read file %s from zip: %w", file.Name, err)
			}

			pngImage := &EmbeddedPngImage{
				Name: file.Name,
				// Encode the raw byte content to a Base64 string
				Base64Content: base64.StdEncoding.EncodeToString(content),
			}
			pngImages = append(pngImages, pngImage)
		}
	}

	if reqifContent == nil {
		return nil, svgImages, jpgImages, pngImages, fmt.Errorf("no .reqif file found in the zip archive")
	}

	return reqifContent, svgImages, jpgImages, pngImages, nil
}

func (stager *Stager) UpdateAndCommitLoadReqifStage() {

	stager.loadReqifStage.Reset()

	var fileToUpload = &load.FileToUpload{
		Name: "Name of file",
		FileToUploadProxy: &FileToUploadProxy{
			stager: stager,
		},
	}

	load.StageBranch(stager.loadReqifStage,
		fileToUpload,
	)

	message := &load.Message{
		Name: "Drop your .reqif or .reqifz file here or ",
	}

	message.Stage(stager.loadReqifStage)

	stager.loadReqifStage.Commit()
}
