// generated code - do not edit

import { LinkAnchoredPathAPI } from './linkanchoredpath-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LinkAnchoredPath {

	static GONGSTRUCT_NAME = "LinkAnchoredPath"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Definition: string = ""
	X_Offset: number = 0
	Y_Offset: number = 0
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

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyLinkAnchoredPathToLinkAnchoredPathAPI(linkanchoredpath: LinkAnchoredPath, linkanchoredpathAPI: LinkAnchoredPathAPI) {

	linkanchoredpathAPI.CreatedAt = linkanchoredpath.CreatedAt
	linkanchoredpathAPI.DeletedAt = linkanchoredpath.DeletedAt
	linkanchoredpathAPI.ID = linkanchoredpath.ID

	// insertion point for basic fields copy operations
	linkanchoredpathAPI.Name = linkanchoredpath.Name
	linkanchoredpathAPI.Definition = linkanchoredpath.Definition
	linkanchoredpathAPI.X_Offset = linkanchoredpath.X_Offset
	linkanchoredpathAPI.Y_Offset = linkanchoredpath.Y_Offset
	linkanchoredpathAPI.ScalePropotionnally = linkanchoredpath.ScalePropotionnally
	linkanchoredpathAPI.AppliedScaling = linkanchoredpath.AppliedScaling
	linkanchoredpathAPI.Color = linkanchoredpath.Color
	linkanchoredpathAPI.FillOpacity = linkanchoredpath.FillOpacity
	linkanchoredpathAPI.Stroke = linkanchoredpath.Stroke
	linkanchoredpathAPI.StrokeOpacity = linkanchoredpath.StrokeOpacity
	linkanchoredpathAPI.StrokeWidth = linkanchoredpath.StrokeWidth
	linkanchoredpathAPI.StrokeDashArray = linkanchoredpath.StrokeDashArray
	linkanchoredpathAPI.StrokeDashArrayWhenSelected = linkanchoredpath.StrokeDashArrayWhenSelected
	linkanchoredpathAPI.Transform = linkanchoredpath.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyLinkAnchoredPathAPIToLinkAnchoredPath update basic, pointers and slice of pointers fields of linkanchoredpath
// from respectively the basic fields and encoded fields of pointers and slices of pointers of linkanchoredpathAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyLinkAnchoredPathAPIToLinkAnchoredPath(linkanchoredpathAPI: LinkAnchoredPathAPI, linkanchoredpath: LinkAnchoredPath, frontRepo: FrontRepo) {

	linkanchoredpath.CreatedAt = linkanchoredpathAPI.CreatedAt
	linkanchoredpath.DeletedAt = linkanchoredpathAPI.DeletedAt
	linkanchoredpath.ID = linkanchoredpathAPI.ID

	// insertion point for basic fields copy operations
	linkanchoredpath.Name = linkanchoredpathAPI.Name
	linkanchoredpath.Definition = linkanchoredpathAPI.Definition
	linkanchoredpath.X_Offset = linkanchoredpathAPI.X_Offset
	linkanchoredpath.Y_Offset = linkanchoredpathAPI.Y_Offset
	linkanchoredpath.ScalePropotionnally = linkanchoredpathAPI.ScalePropotionnally
	linkanchoredpath.AppliedScaling = linkanchoredpathAPI.AppliedScaling
	linkanchoredpath.Color = linkanchoredpathAPI.Color
	linkanchoredpath.FillOpacity = linkanchoredpathAPI.FillOpacity
	linkanchoredpath.Stroke = linkanchoredpathAPI.Stroke
	linkanchoredpath.StrokeOpacity = linkanchoredpathAPI.StrokeOpacity
	linkanchoredpath.StrokeWidth = linkanchoredpathAPI.StrokeWidth
	linkanchoredpath.StrokeDashArray = linkanchoredpathAPI.StrokeDashArray
	linkanchoredpath.StrokeDashArrayWhenSelected = linkanchoredpathAPI.StrokeDashArrayWhenSelected
	linkanchoredpath.Transform = linkanchoredpathAPI.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
