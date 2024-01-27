// generated code - do not edit

import { RectAnchoredPathDB } from './rectanchoredpath-db'
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
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyRectAnchoredPathToRectAnchoredPathDB(rectanchoredpath: RectAnchoredPath, rectanchoredpathDB: RectAnchoredPathDB) {

	rectanchoredpathDB.CreatedAt = rectanchoredpath.CreatedAt
	rectanchoredpathDB.DeletedAt = rectanchoredpath.DeletedAt
	rectanchoredpathDB.ID = rectanchoredpath.ID

	// insertion point for basic fields copy operations
	rectanchoredpathDB.Name = rectanchoredpath.Name
	rectanchoredpathDB.Definition = rectanchoredpath.Definition
	rectanchoredpathDB.X_Offset = rectanchoredpath.X_Offset
	rectanchoredpathDB.Y_Offset = rectanchoredpath.Y_Offset
	rectanchoredpathDB.RectAnchorType = rectanchoredpath.RectAnchorType
	rectanchoredpathDB.ScalePropotionnally = rectanchoredpath.ScalePropotionnally
	rectanchoredpathDB.AppliedScaling = rectanchoredpath.AppliedScaling
	rectanchoredpathDB.Color = rectanchoredpath.Color
	rectanchoredpathDB.FillOpacity = rectanchoredpath.FillOpacity
	rectanchoredpathDB.Stroke = rectanchoredpath.Stroke
	rectanchoredpathDB.StrokeWidth = rectanchoredpath.StrokeWidth
	rectanchoredpathDB.StrokeDashArray = rectanchoredpath.StrokeDashArray
	rectanchoredpathDB.StrokeDashArrayWhenSelected = rectanchoredpath.StrokeDashArrayWhenSelected
	rectanchoredpathDB.Transform = rectanchoredpath.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyRectAnchoredPathDBToRectAnchoredPath update basic, pointers and slice of pointers fields of rectanchoredpath
// from respectively the basic fields and encoded fields of pointers and slices of pointers of rectanchoredpathDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyRectAnchoredPathDBToRectAnchoredPath(rectanchoredpathDB: RectAnchoredPathDB, rectanchoredpath: RectAnchoredPath, frontRepo: FrontRepo) {

	rectanchoredpath.CreatedAt = rectanchoredpathDB.CreatedAt
	rectanchoredpath.DeletedAt = rectanchoredpathDB.DeletedAt
	rectanchoredpath.ID = rectanchoredpathDB.ID

	// insertion point for basic fields copy operations
	rectanchoredpath.Name = rectanchoredpathDB.Name
	rectanchoredpath.Definition = rectanchoredpathDB.Definition
	rectanchoredpath.X_Offset = rectanchoredpathDB.X_Offset
	rectanchoredpath.Y_Offset = rectanchoredpathDB.Y_Offset
	rectanchoredpath.RectAnchorType = rectanchoredpathDB.RectAnchorType
	rectanchoredpath.ScalePropotionnally = rectanchoredpathDB.ScalePropotionnally
	rectanchoredpath.AppliedScaling = rectanchoredpathDB.AppliedScaling
	rectanchoredpath.Color = rectanchoredpathDB.Color
	rectanchoredpath.FillOpacity = rectanchoredpathDB.FillOpacity
	rectanchoredpath.Stroke = rectanchoredpathDB.Stroke
	rectanchoredpath.StrokeWidth = rectanchoredpathDB.StrokeWidth
	rectanchoredpath.StrokeDashArray = rectanchoredpathDB.StrokeDashArray
	rectanchoredpath.StrokeDashArrayWhenSelected = rectanchoredpathDB.StrokeDashArrayWhenSelected
	rectanchoredpath.Transform = rectanchoredpathDB.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
