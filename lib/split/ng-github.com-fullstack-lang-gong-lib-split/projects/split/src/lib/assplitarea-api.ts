// insertion point for imports
import { AsSplitAPI } from './assplit-api'
import { TreeAPI } from './tree-api'
import { TableAPI } from './table-api'
import { FormAPI } from './form-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AsSplitAreaAPI {

	static GONGSTRUCT_NAME = "AsSplitArea"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Size: number = 0
	IsAny: boolean = false

	// insertion point for other decls

	AsSplitAreaPointersEncoding: AsSplitAreaPointersEncoding = new AsSplitAreaPointersEncoding
}

export class AsSplitAreaPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	AsSplits: number[] = []
	TreeID: NullInt64 = new NullInt64 // if pointer is null, Tree.ID = 0

	TableID: NullInt64 = new NullInt64 // if pointer is null, Table.ID = 0

	FormID: NullInt64 = new NullInt64 // if pointer is null, Form.ID = 0

}
