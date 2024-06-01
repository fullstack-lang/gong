// generated code - do not edit

import { RectLinkLinkAPI } from './rectlinklink-api'
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
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
	Start?: Rect

	End?: Link

}

export function CopyRectLinkLinkToRectLinkLinkAPI(rectlinklink: RectLinkLink, rectlinklinkAPI: RectLinkLinkAPI) {

	rectlinklinkAPI.CreatedAt = rectlinklink.CreatedAt
	rectlinklinkAPI.DeletedAt = rectlinklink.DeletedAt
	rectlinklinkAPI.ID = rectlinklink.ID

	// insertion point for basic fields copy operations
	rectlinklinkAPI.Name = rectlinklink.Name
	rectlinklinkAPI.TargetAnchorPosition = rectlinklink.TargetAnchorPosition
	rectlinklinkAPI.Color = rectlinklink.Color
	rectlinklinkAPI.FillOpacity = rectlinklink.FillOpacity
	rectlinklinkAPI.Stroke = rectlinklink.Stroke
	rectlinklinkAPI.StrokeOpacity = rectlinklink.StrokeOpacity
	rectlinklinkAPI.StrokeWidth = rectlinklink.StrokeWidth
	rectlinklinkAPI.StrokeDashArray = rectlinklink.StrokeDashArray
	rectlinklinkAPI.StrokeDashArrayWhenSelected = rectlinklink.StrokeDashArrayWhenSelected
	rectlinklinkAPI.Transform = rectlinklink.Transform

	// insertion point for pointer fields encoding
	rectlinklinkAPI.RectLinkLinkPointersEncoding.StartID.Valid = true
	if (rectlinklink.Start != undefined) {
		rectlinklinkAPI.RectLinkLinkPointersEncoding.StartID.Int64 = rectlinklink.Start.ID  
	} else {
		rectlinklinkAPI.RectLinkLinkPointersEncoding.StartID.Int64 = 0 		
	}

	rectlinklinkAPI.RectLinkLinkPointersEncoding.EndID.Valid = true
	if (rectlinklink.End != undefined) {
		rectlinklinkAPI.RectLinkLinkPointersEncoding.EndID.Int64 = rectlinklink.End.ID  
	} else {
		rectlinklinkAPI.RectLinkLinkPointersEncoding.EndID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyRectLinkLinkAPIToRectLinkLink update basic, pointers and slice of pointers fields of rectlinklink
// from respectively the basic fields and encoded fields of pointers and slices of pointers of rectlinklinkAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyRectLinkLinkAPIToRectLinkLink(rectlinklinkAPI: RectLinkLinkAPI, rectlinklink: RectLinkLink, frontRepo: FrontRepo) {

	rectlinklink.CreatedAt = rectlinklinkAPI.CreatedAt
	rectlinklink.DeletedAt = rectlinklinkAPI.DeletedAt
	rectlinklink.ID = rectlinklinkAPI.ID

	// insertion point for basic fields copy operations
	rectlinklink.Name = rectlinklinkAPI.Name
	rectlinklink.TargetAnchorPosition = rectlinklinkAPI.TargetAnchorPosition
	rectlinklink.Color = rectlinklinkAPI.Color
	rectlinklink.FillOpacity = rectlinklinkAPI.FillOpacity
	rectlinklink.Stroke = rectlinklinkAPI.Stroke
	rectlinklink.StrokeOpacity = rectlinklinkAPI.StrokeOpacity
	rectlinklink.StrokeWidth = rectlinklinkAPI.StrokeWidth
	rectlinklink.StrokeDashArray = rectlinklinkAPI.StrokeDashArray
	rectlinklink.StrokeDashArrayWhenSelected = rectlinklinkAPI.StrokeDashArrayWhenSelected
	rectlinklink.Transform = rectlinklinkAPI.Transform

	// insertion point for pointer fields encoding
	rectlinklink.Start = frontRepo.map_ID_Rect.get(rectlinklinkAPI.RectLinkLinkPointersEncoding.StartID.Int64)
	rectlinklink.End = frontRepo.map_ID_Link.get(rectlinklinkAPI.RectLinkLinkPointersEncoding.EndID.Int64)

	// insertion point for slice of pointers fields encoding
}
