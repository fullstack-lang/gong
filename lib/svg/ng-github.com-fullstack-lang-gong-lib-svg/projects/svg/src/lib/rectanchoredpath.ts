// generated code - do not edit

import { RectAnchoredPathAPI } from './rectanchoredpath-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RectAnchoredPath {

	static GONGSTRUCT_NAME = "RectAnchoredPath"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Definition: string = ""
	X_Offset: number = 0
	Y_Offset: number = 0
	RectAnchorType: string = ""
	ScalePropotionnally: boolean = false
	AppliedScaling: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyRectAnchoredPathToRectAnchoredPathAPI(rectanchoredpath: RectAnchoredPath, rectanchoredpathAPI: RectAnchoredPathAPI) {

	rectanchoredpathAPI.CreatedAt = rectanchoredpath.CreatedAt
	rectanchoredpathAPI.DeletedAt = rectanchoredpath.DeletedAt
	rectanchoredpathAPI.ID = rectanchoredpath.ID

	// insertion point for basic fields copy operations
	rectanchoredpathAPI.Name = rectanchoredpath.Name
	rectanchoredpathAPI.Definition = rectanchoredpath.Definition
	rectanchoredpathAPI.X_Offset = rectanchoredpath.X_Offset
	rectanchoredpathAPI.Y_Offset = rectanchoredpath.Y_Offset
	rectanchoredpathAPI.RectAnchorType = rectanchoredpath.RectAnchorType
	rectanchoredpathAPI.ScalePropotionnally = rectanchoredpath.ScalePropotionnally
	rectanchoredpathAPI.AppliedScaling = rectanchoredpath.AppliedScaling
	rectanchoredpathAPI.Color = rectanchoredpath.Color
	rectanchoredpathAPI.FillOpacity = rectanchoredpath.FillOpacity
	rectanchoredpathAPI.Stroke = rectanchoredpath.Stroke
	rectanchoredpathAPI.StrokeOpacity = rectanchoredpath.StrokeOpacity
	rectanchoredpathAPI.StrokeWidth = rectanchoredpath.StrokeWidth
	rectanchoredpathAPI.StrokeDashArray = rectanchoredpath.StrokeDashArray
	rectanchoredpathAPI.StrokeDashArrayWhenSelected = rectanchoredpath.StrokeDashArrayWhenSelected
	rectanchoredpathAPI.Transform = rectanchoredpath.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyRectAnchoredPathAPIToRectAnchoredPath update basic, pointers and slice of pointers fields of rectanchoredpath
// from respectively the basic fields and encoded fields of pointers and slices of pointers of rectanchoredpathAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyRectAnchoredPathAPIToRectAnchoredPath(rectanchoredpathAPI: RectAnchoredPathAPI, rectanchoredpath: RectAnchoredPath, frontRepo: FrontRepo) {

	rectanchoredpath.CreatedAt = rectanchoredpathAPI.CreatedAt
	rectanchoredpath.DeletedAt = rectanchoredpathAPI.DeletedAt
	rectanchoredpath.ID = rectanchoredpathAPI.ID

	// insertion point for basic fields copy operations
	rectanchoredpath.Name = rectanchoredpathAPI.Name
	rectanchoredpath.Definition = rectanchoredpathAPI.Definition
	rectanchoredpath.X_Offset = rectanchoredpathAPI.X_Offset
	rectanchoredpath.Y_Offset = rectanchoredpathAPI.Y_Offset
	rectanchoredpath.RectAnchorType = rectanchoredpathAPI.RectAnchorType
	rectanchoredpath.ScalePropotionnally = rectanchoredpathAPI.ScalePropotionnally
	rectanchoredpath.AppliedScaling = rectanchoredpathAPI.AppliedScaling
	rectanchoredpath.Color = rectanchoredpathAPI.Color
	rectanchoredpath.FillOpacity = rectanchoredpathAPI.FillOpacity
	rectanchoredpath.Stroke = rectanchoredpathAPI.Stroke
	rectanchoredpath.StrokeOpacity = rectanchoredpathAPI.StrokeOpacity
	rectanchoredpath.StrokeWidth = rectanchoredpathAPI.StrokeWidth
	rectanchoredpath.StrokeDashArray = rectanchoredpathAPI.StrokeDashArray
	rectanchoredpath.StrokeDashArrayWhenSelected = rectanchoredpathAPI.StrokeDashArrayWhenSelected
	rectanchoredpath.Transform = rectanchoredpathAPI.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
