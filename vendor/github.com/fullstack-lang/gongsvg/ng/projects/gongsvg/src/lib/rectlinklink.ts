// generated code - do not edit

import { RectLinkLinkDB } from './rectlinklink-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Rect } from './rect'
import { Link } from './link'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RectLinkLink {

	static GONGSTRUCT_NAME = "RectLinkLink"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	TargetAnchorPosition: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
	Start?: Rect

	End?: Link

}

export function CopyRectLinkLinkToRectLinkLinkDB(rectlinklink: RectLinkLink, rectlinklinkDB: RectLinkLinkDB) {

	rectlinklinkDB.CreatedAt = rectlinklink.CreatedAt
	rectlinklinkDB.DeletedAt = rectlinklink.DeletedAt
	rectlinklinkDB.ID = rectlinklink.ID

	// insertion point for basic fields copy operations
	rectlinklinkDB.Name = rectlinklink.Name
	rectlinklinkDB.TargetAnchorPosition = rectlinklink.TargetAnchorPosition
	rectlinklinkDB.Color = rectlinklink.Color
	rectlinklinkDB.FillOpacity = rectlinklink.FillOpacity
	rectlinklinkDB.Stroke = rectlinklink.Stroke
	rectlinklinkDB.StrokeWidth = rectlinklink.StrokeWidth
	rectlinklinkDB.StrokeDashArray = rectlinklink.StrokeDashArray
	rectlinklinkDB.StrokeDashArrayWhenSelected = rectlinklink.StrokeDashArrayWhenSelected
	rectlinklinkDB.Transform = rectlinklink.Transform

	// insertion point for pointer fields encoding
	rectlinklinkDB.RectLinkLinkPointersEncoding.StartID.Valid = true
	if (rectlinklink.Start != undefined) {
		rectlinklinkDB.RectLinkLinkPointersEncoding.StartID.Int64 = rectlinklink.Start.ID  
	} else {
		rectlinklinkDB.RectLinkLinkPointersEncoding.StartID.Int64 = 0 		
	}

	rectlinklinkDB.RectLinkLinkPointersEncoding.EndID.Valid = true
	if (rectlinklink.End != undefined) {
		rectlinklinkDB.RectLinkLinkPointersEncoding.EndID.Int64 = rectlinklink.End.ID  
	} else {
		rectlinklinkDB.RectLinkLinkPointersEncoding.EndID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyRectLinkLinkDBToRectLinkLink update basic, pointers and slice of pointers fields of rectlinklink
// from respectively the basic fields and encoded fields of pointers and slices of pointers of rectlinklinkDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyRectLinkLinkDBToRectLinkLink(rectlinklinkDB: RectLinkLinkDB, rectlinklink: RectLinkLink, frontRepo: FrontRepo) {

	rectlinklink.CreatedAt = rectlinklinkDB.CreatedAt
	rectlinklink.DeletedAt = rectlinklinkDB.DeletedAt
	rectlinklink.ID = rectlinklinkDB.ID

	// insertion point for basic fields copy operations
	rectlinklink.Name = rectlinklinkDB.Name
	rectlinklink.TargetAnchorPosition = rectlinklinkDB.TargetAnchorPosition
	rectlinklink.Color = rectlinklinkDB.Color
	rectlinklink.FillOpacity = rectlinklinkDB.FillOpacity
	rectlinklink.Stroke = rectlinklinkDB.Stroke
	rectlinklink.StrokeWidth = rectlinklinkDB.StrokeWidth
	rectlinklink.StrokeDashArray = rectlinklinkDB.StrokeDashArray
	rectlinklink.StrokeDashArrayWhenSelected = rectlinklinkDB.StrokeDashArrayWhenSelected
	rectlinklink.Transform = rectlinklinkDB.Transform

	// insertion point for pointer fields encoding
	rectlinklink.Start = frontRepo.map_ID_Rect.get(rectlinklinkDB.RectLinkLinkPointersEncoding.StartID.Int64)
	rectlinklink.End = frontRepo.map_ID_Link.get(rectlinklinkDB.RectLinkLinkPointersEncoding.EndID.Int64)

	// insertion point for slice of pointers fields encoding
}
