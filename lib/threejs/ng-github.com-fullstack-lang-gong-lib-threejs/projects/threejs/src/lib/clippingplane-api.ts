// generated code - do not edit
// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ClippingPlaneAPI {

	static GONGSTRUCT_NAME = "ClippingPlane"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsEnabled: boolean = false
	NormalX: number = 0
	NormalY: number = 0
	NormalZ: number = 0
	Constant: number = 0
	IsHelperVisible: boolean = false
	HelperSize: number = 0
	HelperCenterX: number = 0
	HelperCenterY: number = 0
	HelperCenterZ: number = 0

	// insertion point for other decls

	ClippingPlanePointersEncoding: ClippingPlanePointersEncoding = new ClippingPlanePointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class ClippingPlanePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
