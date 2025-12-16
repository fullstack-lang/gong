// insertion point for imports
import { LayerAPI } from './layer-api'
import { RectAPI } from './rect-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SVGAPI {

	static GONGSTRUCT_NAME = "SVG"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	DrawingState: string = ""
	IsEditable: boolean = false
	IsSVGFrontEndFileGenerated: boolean = false
	IsSVGBackEndFileGenerated: boolean = false
	DefaultDirectoryForGeneratedImages: string = ""
	IsControlBannerHidden: boolean = false
	OverrideWidth: boolean = false
	OverriddenWidth: number = 0
	OverrideHeight: boolean = false
	OverriddenHeight: number = 0

	// insertion point for other decls

	SVGPointersEncoding: SVGPointersEncoding = new SVGPointersEncoding

	CreatedAt?: string
	DeletedAt?: string
}

export class SVGPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Layers: number[] = []
	StartRectID: NullInt64 = new NullInt64 // if pointer is null, StartRect.ID = 0

	EndRectID: NullInt64 = new NullInt64 // if pointer is null, EndRect.ID = 0

}
