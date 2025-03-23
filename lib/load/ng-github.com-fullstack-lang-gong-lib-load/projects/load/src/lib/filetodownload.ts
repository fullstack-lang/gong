// generated code - do not edit

import { FileToDownloadAPI } from './filetodownload-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FileToDownload {

	static GONGSTRUCT_NAME = "FileToDownload"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyFileToDownloadToFileToDownloadAPI(filetodownload: FileToDownload, filetodownloadAPI: FileToDownloadAPI) {

	filetodownloadAPI.CreatedAt = filetodownload.CreatedAt
	filetodownloadAPI.DeletedAt = filetodownload.DeletedAt
	filetodownloadAPI.ID = filetodownload.ID

	// insertion point for basic fields copy operations
	filetodownloadAPI.Name = filetodownload.Name
	filetodownloadAPI.Content = filetodownload.Content

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFileToDownloadAPIToFileToDownload update basic, pointers and slice of pointers fields of filetodownload
// from respectively the basic fields and encoded fields of pointers and slices of pointers of filetodownloadAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFileToDownloadAPIToFileToDownload(filetodownloadAPI: FileToDownloadAPI, filetodownload: FileToDownload, frontRepo: FrontRepo) {

	filetodownload.CreatedAt = filetodownloadAPI.CreatedAt
	filetodownload.DeletedAt = filetodownloadAPI.DeletedAt
	filetodownload.ID = filetodownloadAPI.ID

	// insertion point for basic fields copy operations
	filetodownload.Name = filetodownloadAPI.Name
	filetodownload.Content = filetodownloadAPI.Content

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
