// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FileToDownloadAPI {

	static GONGSTRUCT_NAME = "FileToDownload"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""

	// insertion point for other decls

	FileToDownloadPointersEncoding: FileToDownloadPointersEncoding = new FileToDownloadPointersEncoding
}

export class FileToDownloadPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
