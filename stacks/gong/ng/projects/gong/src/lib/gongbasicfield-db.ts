// insertion point for imports
import { GongEnumDB } from './gongenum-db'
import { GongStructDB } from './gongstruct-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class GongBasicFieldDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	BasicKindName?: string
	DeclaredType?: string

	// insertion point for other declarations
	GongEnum?: GongEnumDB
	GongEnumID?: NullInt64
	GongEnumName?: string

	GongStruct_GongBasicFieldsDBID?: NullInt64
	GongStruct_GongBasicFields_reverse?: GongStructDB

}
