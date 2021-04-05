// insertion point for imports
import { GongStructDB } from './gongstruct-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class PointerToGongStructFieldDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string

	// insertion point for other declarations
	GongStruct?: GongStructDB
	GongStructID?: NullInt64
	GongStructName?: string

	GongStruct_PointerToGongStructFieldsDBID?: NullInt64
	GongStruct_PointerToGongStructFields_reverse?: GongStructDB

}
