// insertion point for imports
import { SVGIconAPI } from './svgicon-api'
import { ButtonAPI } from './button-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class NodeAPI {

	static GONGSTRUCT_NAME = "Node"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	FontStyle: string = ""
	BackgroundColor: string = ""
	IsExpanded: boolean = false
	HasCheckboxButton: boolean = false
	IsChecked: boolean = false
	IsCheckboxDisabled: boolean = false
	HasSecondCheckboxButton: boolean = false
	IsSecondCheckboxChecked: boolean = false
	IsSecondCheckboxDisabled: boolean = false
	TextAfterSecondCheckbox: string = ""
	IsInEditMode: boolean = false
	IsNodeClickable: boolean = false
	IsWithPreceedingIcon: boolean = false
	PreceedingIcon: string = ""

	// insertion point for other decls

	NodePointersEncoding: NodePointersEncoding = new NodePointersEncoding
}

export class NodePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	PreceedingSVGIconID: NullInt64 = new NullInt64 // if pointer is null, PreceedingSVGIcon.ID = 0

	Children: number[] = []
	Buttons: number[] = []
}
