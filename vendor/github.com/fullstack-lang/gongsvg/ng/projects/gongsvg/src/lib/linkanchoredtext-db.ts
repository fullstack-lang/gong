// insertion point for imports
import { AnimateDB } from './animate-db'
import { LinkDB } from './link-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LinkAnchoredTextDB {

	static GONGSTRUCT_NAME = "LinkAnchoredText"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""
	X_Offset: number = 0
	Y_Offset: number = 0
	FontWeight: string = ""
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
	Animates: Array<AnimateDB> = []

	LinkAnchoredTextPointersEncoding: LinkAnchoredTextPointersEncoding = new LinkAnchoredTextPointersEncoding
}

export class LinkAnchoredTextPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Animates: number[] = []
	// reverse pointers encoding (to be removed)
	Link_TextAtArrowEndDBID: NullInt64 = new NullInt64
	Link_TextAtArrowEndDBID_Index: NullInt64  = new NullInt64 // store the index of the linkanchoredtext instance in Link.TextAtArrowEnd
	Link_TextAtArrowEnd_reverse?: LinkDB 

	// reverse pointers encoding (to be removed)
	Link_TextAtArrowStartDBID: NullInt64 = new NullInt64
	Link_TextAtArrowStartDBID_Index: NullInt64  = new NullInt64 // store the index of the linkanchoredtext instance in Link.TextAtArrowStart
	Link_TextAtArrowStart_reverse?: LinkDB 

}
