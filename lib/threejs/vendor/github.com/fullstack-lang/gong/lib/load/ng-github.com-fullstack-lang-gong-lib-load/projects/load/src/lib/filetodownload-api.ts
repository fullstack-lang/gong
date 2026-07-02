// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FileToDownloadAPI {

	static GONGSTRUCT_NAME = "FileToDownload"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Base64EncodedContent: string = ""

	// insertion point for other decls

	FileToDownloadPointersEncoding: FileToDownloadPointersEncoding = new FileToDownloadPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class FileToDownloadPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
