// generated code - do not edit

import { FileToUploadAPI } from './filetoupload-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FileToUpload {

	static GONGSTRUCT_NAME = "FileToUpload"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Base64EncodedContent: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyFileToUploadToFileToUploadAPI(filetoupload: FileToUpload, filetouploadAPI: FileToUploadAPI) {

	filetouploadAPI.CreatedAt = filetoupload.CreatedAt
	filetouploadAPI.DeletedAt = filetoupload.DeletedAt
	filetouploadAPI.ID = filetoupload.ID

	// insertion point for basic fields copy operations
	filetouploadAPI.Name = filetoupload.Name
	filetouploadAPI.Base64EncodedContent = filetoupload.Base64EncodedContent

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFileToUploadAPIToFileToUpload update basic, pointers and slice of pointers fields of filetoupload
// from respectively the basic fields and encoded fields of pointers and slices of pointers of filetouploadAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFileToUploadAPIToFileToUpload(filetouploadAPI: FileToUploadAPI, filetoupload: FileToUpload, frontRepo: FrontRepo) {

	filetoupload.CreatedAt = filetouploadAPI.CreatedAt
	filetoupload.DeletedAt = filetouploadAPI.DeletedAt
	filetoupload.ID = filetouploadAPI.ID

	// insertion point for basic fields copy operations
	filetoupload.Name = filetouploadAPI.Name
	filetoupload.Base64EncodedContent = filetouploadAPI.Base64EncodedContent

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
