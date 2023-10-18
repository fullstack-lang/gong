// insertion point for imports
import { NodeDB } from './node-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ButtonDB {

	static GONGSTRUCT_NAME = "Button"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Icon: string = ""

	// insertion point for pointers and slices of pointers declarations

	ButtonPointersEncoding: ButtonPointersEncoding = new ButtonPointersEncoding
}

export class ButtonPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	// reverse pointers encoding (to be removed)
	Node_ButtonsDBID: NullInt64 = new NullInt64
	Node_ButtonsDBID_Index: NullInt64  = new NullInt64 // store the index of the button instance in Node.Buttons
	Node_Buttons_reverse?: NodeDB 

}
