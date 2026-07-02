// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MarkdownAPI {

	static GONGSTRUCT_NAME = "Markdown"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for other decls

	MarkdownPointersEncoding: MarkdownPointersEncoding = new MarkdownPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class MarkdownPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
