// insertion point for imports
import { AsSplitAPI } from './assplit-api'
import { ButtonAPI } from './button-api'
import { CursorAPI } from './cursor-api'
import { DocAPI } from './doc-api'
import { FormAPI } from './form-api'
import { LoadAPI } from './load-api'
import { SliderAPI } from './slider-api'
import { SplitAPI } from './split-api'
import { SvgAPI } from './svg-api'
import { TableAPI } from './table-api'
import { ToneAPI } from './tone-api'
import { TreeAPI } from './tree-api'
import { XlsxAPI } from './xlsx-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AsSplitAreaAPI {

	static GONGSTRUCT_NAME = "AsSplitArea"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	ShowNameInHeader: boolean = false
	Size: number = 0
	IsAny: boolean = false
	HasDiv: boolean = false
	DivStyle: string = ""

	// insertion point for other decls

	AsSplitAreaPointersEncoding: AsSplitAreaPointersEncoding = new AsSplitAreaPointersEncoding
}

export class AsSplitAreaPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	AsSplitID: NullInt64 = new NullInt64 // if pointer is null, AsSplit.ID = 0

	ButtonID: NullInt64 = new NullInt64 // if pointer is null, Button.ID = 0

	CursorID: NullInt64 = new NullInt64 // if pointer is null, Cursor.ID = 0

	DocID: NullInt64 = new NullInt64 // if pointer is null, Doc.ID = 0

	FormID: NullInt64 = new NullInt64 // if pointer is null, Form.ID = 0

	LoadID: NullInt64 = new NullInt64 // if pointer is null, Load.ID = 0

	SliderID: NullInt64 = new NullInt64 // if pointer is null, Slider.ID = 0

	SplitID: NullInt64 = new NullInt64 // if pointer is null, Split.ID = 0

	SvgID: NullInt64 = new NullInt64 // if pointer is null, Svg.ID = 0

	TableID: NullInt64 = new NullInt64 // if pointer is null, Table.ID = 0

	ToneID: NullInt64 = new NullInt64 // if pointer is null, Tone.ID = 0

	TreeID: NullInt64 = new NullInt64 // if pointer is null, Tree.ID = 0

	XlsxID: NullInt64 = new NullInt64 // if pointer is null, Xlsx.ID = 0

}
