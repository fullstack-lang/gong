// generated code - do not edit

import { LinkAPI } from './link-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Rect } from './rect'
import { LinkAnchoredText } from './linkanchoredtext'
import { Point } from './point'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Link {

	static GONGSTRUCT_NAME = "Link"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Type: string = ""
	IsBezierCurve: boolean = false
	StartAnchorType: string = ""
	EndAnchorType: string = ""
	StartOrientation: string = ""
	StartRatio: number = 0
	EndOrientation: string = ""
	EndRatio: number = 0
	CornerOffsetRatio: number = 0
	CornerRadius: number = 0
	HasEndArrow: boolean = false
	EndArrowSize: number = 0
	HasStartArrow: boolean = false
	StartArrowSize: number = 0
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

	End?: Rect

	TextAtArrowEnd: Array<LinkAnchoredText> = []
	TextAtArrowStart: Array<LinkAnchoredText> = []
	ControlPoints: Array<Point> = []
}

export function CopyLinkToLinkAPI(link: Link, linkAPI: LinkAPI) {

	linkAPI.CreatedAt = link.CreatedAt
	linkAPI.DeletedAt = link.DeletedAt
	linkAPI.ID = link.ID

	// insertion point for basic fields copy operations
	linkAPI.Name = link.Name
	linkAPI.Type = link.Type
	linkAPI.IsBezierCurve = link.IsBezierCurve
	linkAPI.StartAnchorType = link.StartAnchorType
	linkAPI.EndAnchorType = link.EndAnchorType
	linkAPI.StartOrientation = link.StartOrientation
	linkAPI.StartRatio = link.StartRatio
	linkAPI.EndOrientation = link.EndOrientation
	linkAPI.EndRatio = link.EndRatio
	linkAPI.CornerOffsetRatio = link.CornerOffsetRatio
	linkAPI.CornerRadius = link.CornerRadius
	linkAPI.HasEndArrow = link.HasEndArrow
	linkAPI.EndArrowSize = link.EndArrowSize
	linkAPI.HasStartArrow = link.HasStartArrow
	linkAPI.StartArrowSize = link.StartArrowSize
	linkAPI.Color = link.Color
	linkAPI.FillOpacity = link.FillOpacity
	linkAPI.Stroke = link.Stroke
	linkAPI.StrokeOpacity = link.StrokeOpacity
	linkAPI.StrokeWidth = link.StrokeWidth
	linkAPI.StrokeDashArray = link.StrokeDashArray
	linkAPI.StrokeDashArrayWhenSelected = link.StrokeDashArrayWhenSelected
	linkAPI.Transform = link.Transform

	// insertion point for pointer fields encoding
	linkAPI.LinkPointersEncoding.StartID.Valid = true
	if (link.Start != undefined) {
		linkAPI.LinkPointersEncoding.StartID.Int64 = link.Start.ID  
	} else {
		linkAPI.LinkPointersEncoding.StartID.Int64 = 0 		
	}

	linkAPI.LinkPointersEncoding.EndID.Valid = true
	if (link.End != undefined) {
		linkAPI.LinkPointersEncoding.EndID.Int64 = link.End.ID  
	} else {
		linkAPI.LinkPointersEncoding.EndID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	linkAPI.LinkPointersEncoding.TextAtArrowEnd = []
	for (let _linkanchoredtext of link.TextAtArrowEnd) {
		linkAPI.LinkPointersEncoding.TextAtArrowEnd.push(_linkanchoredtext.ID)
	}

	linkAPI.LinkPointersEncoding.TextAtArrowStart = []
	for (let _linkanchoredtext of link.TextAtArrowStart) {
		linkAPI.LinkPointersEncoding.TextAtArrowStart.push(_linkanchoredtext.ID)
	}

	linkAPI.LinkPointersEncoding.ControlPoints = []
	for (let _point of link.ControlPoints) {
		linkAPI.LinkPointersEncoding.ControlPoints.push(_point.ID)
	}

}

// CopyLinkAPIToLink update basic, pointers and slice of pointers fields of link
// from respectively the basic fields and encoded fields of pointers and slices of pointers of linkAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyLinkAPIToLink(linkAPI: LinkAPI, link: Link, frontRepo: FrontRepo) {

	link.CreatedAt = linkAPI.CreatedAt
	link.DeletedAt = linkAPI.DeletedAt
	link.ID = linkAPI.ID

	// insertion point for basic fields copy operations
	link.Name = linkAPI.Name
	link.Type = linkAPI.Type
	link.IsBezierCurve = linkAPI.IsBezierCurve
	link.StartAnchorType = linkAPI.StartAnchorType
	link.EndAnchorType = linkAPI.EndAnchorType
	link.StartOrientation = linkAPI.StartOrientation
	link.StartRatio = linkAPI.StartRatio
	link.EndOrientation = linkAPI.EndOrientation
	link.EndRatio = linkAPI.EndRatio
	link.CornerOffsetRatio = linkAPI.CornerOffsetRatio
	link.CornerRadius = linkAPI.CornerRadius
	link.HasEndArrow = linkAPI.HasEndArrow
	link.EndArrowSize = linkAPI.EndArrowSize
	link.HasStartArrow = linkAPI.HasStartArrow
	link.StartArrowSize = linkAPI.StartArrowSize
	link.Color = linkAPI.Color
	link.FillOpacity = linkAPI.FillOpacity
	link.Stroke = linkAPI.Stroke
	link.StrokeOpacity = linkAPI.StrokeOpacity
	link.StrokeWidth = linkAPI.StrokeWidth
	link.StrokeDashArray = linkAPI.StrokeDashArray
	link.StrokeDashArrayWhenSelected = linkAPI.StrokeDashArrayWhenSelected
	link.Transform = linkAPI.Transform

	// insertion point for pointer fields encoding
	link.Start = frontRepo.map_ID_Rect.get(linkAPI.LinkPointersEncoding.StartID.Int64)
	link.End = frontRepo.map_ID_Rect.get(linkAPI.LinkPointersEncoding.EndID.Int64)

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(linkAPI.LinkPointersEncoding.TextAtArrowEnd)) {
		console.error('Rects is not an array:', linkAPI.LinkPointersEncoding.TextAtArrowEnd);
		return;
	}

	link.TextAtArrowEnd = new Array<LinkAnchoredText>()
	for (let _id of linkAPI.LinkPointersEncoding.TextAtArrowEnd) {
		let _linkanchoredtext = frontRepo.map_ID_LinkAnchoredText.get(_id)
		if (_linkanchoredtext != undefined) {
			link.TextAtArrowEnd.push(_linkanchoredtext!)
		}
	}
	if (!Array.isArray(linkAPI.LinkPointersEncoding.TextAtArrowStart)) {
		console.error('Rects is not an array:', linkAPI.LinkPointersEncoding.TextAtArrowStart);
		return;
	}

	link.TextAtArrowStart = new Array<LinkAnchoredText>()
	for (let _id of linkAPI.LinkPointersEncoding.TextAtArrowStart) {
		let _linkanchoredtext = frontRepo.map_ID_LinkAnchoredText.get(_id)
		if (_linkanchoredtext != undefined) {
			link.TextAtArrowStart.push(_linkanchoredtext!)
		}
	}
	if (!Array.isArray(linkAPI.LinkPointersEncoding.ControlPoints)) {
		console.error('Rects is not an array:', linkAPI.LinkPointersEncoding.ControlPoints);
		return;
	}

	link.ControlPoints = new Array<Point>()
	for (let _id of linkAPI.LinkPointersEncoding.ControlPoints) {
		let _point = frontRepo.map_ID_Point.get(_id)
		if (_point != undefined) {
			link.ControlPoints.push(_point!)
		}
	}
}
