// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FileToUploadAPI {

	static GONGSTRUCT_NAME = "FileToUpload"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Base64EncodedContent: string = ""

	// insertion point for other decls

	FileToUploadPointersEncoding: FileToUploadPointersEncoding = new FileToUploadPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class FileToUploadPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
