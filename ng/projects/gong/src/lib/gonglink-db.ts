// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongLinkDB {

	static GONGSTRUCT_NAME = "GongLink"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Recv: string = ""
	ImportPath: string = ""

	// insertion point for pointers and slices of pointers declarations

	GongLinkPointersEncoding: GongLinkPointersEncoding = new GongLinkPointersEncoding
}

export class GongLinkPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
