// generated code - do not edit
package models

// Clean computes the reverse map, for all intances, for all clean to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) Clean() {
	// insertion point per named struct
	// Compute reverse map for named struct FileToDownload
	for filetodownload := range stage.FileToDownloads {
		_ = filetodownload
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct FileToUpload
	for filetoupload := range stage.FileToUploads {
		_ = filetoupload
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Message
	for message := range stage.Messages {
		_ = message
		// insertion point per field
		// insertion point per field
	}

}
