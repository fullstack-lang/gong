// insertion point for imports
import { GongBasicFieldDB } from './gongbasicfield-db'
import { PointerToGongStructFieldDB } from './pointertogongstructfield-db'
import { SliceOfPointerToGongStructFieldDB } from './sliceofpointertogongstructfield-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class GongStructDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string

	// insertion point for other declarations
	GongBasicFields?: Array<GongBasicFieldDB>
	PointerToGongStructFields?: Array<PointerToGongStructFieldDB>
	SliceOfPointerToGongStructFields?: Array<SliceOfPointerToGongStructFieldDB>
}
