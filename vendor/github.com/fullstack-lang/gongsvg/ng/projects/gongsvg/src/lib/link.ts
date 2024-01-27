// generated code - do not edit

import { LinkDB } from './link-db'
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

export function CopyLinkToLinkDB(link: Link, linkDB: LinkDB) {

	linkDB.CreatedAt = link.CreatedAt
	linkDB.DeletedAt = link.DeletedAt
	linkDB.ID = link.ID

	// insertion point for basic fields copy operations
	linkDB.Name = link.Name
	linkDB.Type = link.Type
	linkDB.IsBezierCurve = link.IsBezierCurve
	linkDB.StartAnchorType = link.StartAnchorType
	linkDB.EndAnchorType = link.EndAnchorType
	linkDB.StartOrientation = link.StartOrientation
	linkDB.StartRatio = link.StartRatio
	linkDB.EndOrientation = link.EndOrientation
	linkDB.EndRatio = link.EndRatio
	linkDB.CornerOffsetRatio = link.CornerOffsetRatio
	linkDB.CornerRadius = link.CornerRadius
	linkDB.HasEndArrow = link.HasEndArrow
	linkDB.EndArrowSize = link.EndArrowSize
	linkDB.HasStartArrow = link.HasStartArrow
	linkDB.StartArrowSize = link.StartArrowSize
	linkDB.Color = link.Color
	linkDB.FillOpacity = link.FillOpacity
	linkDB.Stroke = link.Stroke
	linkDB.StrokeWidth = link.StrokeWidth
	linkDB.StrokeDashArray = link.StrokeDashArray
	linkDB.StrokeDashArrayWhenSelected = link.StrokeDashArrayWhenSelected
	linkDB.Transform = link.Transform

	// insertion point for pointer fields encoding
	linkDB.LinkPointersEncoding.StartID.Valid = true
	if (link.Start != undefined) {
		linkDB.LinkPointersEncoding.StartID.Int64 = link.Start.ID  
	} else {
		linkDB.LinkPointersEncoding.StartID.Int64 = 0 		
	}

	linkDB.LinkPointersEncoding.EndID.Valid = true
	if (link.End != undefined) {
		linkDB.LinkPointersEncoding.EndID.Int64 = link.End.ID  
	} else {
		linkDB.LinkPointersEncoding.EndID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	linkDB.LinkPointersEncoding.TextAtArrowEnd = []
	for (let _linkanchoredtext of link.TextAtArrowEnd) {
		linkDB.LinkPointersEncoding.TextAtArrowEnd.push(_linkanchoredtext.ID)
	}

	linkDB.LinkPointersEncoding.TextAtArrowStart = []
	for (let _linkanchoredtext of link.TextAtArrowStart) {
		linkDB.LinkPointersEncoding.TextAtArrowStart.push(_linkanchoredtext.ID)
	}

	linkDB.LinkPointersEncoding.ControlPoints = []
	for (let _point of link.ControlPoints) {
		linkDB.LinkPointersEncoding.ControlPoints.push(_point.ID)
	}

}

// CopyLinkDBToLink update basic, pointers and slice of pointers fields of link
// from respectively the basic fields and encoded fields of pointers and slices of pointers of linkDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyLinkDBToLink(linkDB: LinkDB, link: Link, frontRepo: FrontRepo) {

	link.CreatedAt = linkDB.CreatedAt
	link.DeletedAt = linkDB.DeletedAt
	link.ID = linkDB.ID

	// insertion point for basic fields copy operations
	link.Name = linkDB.Name
	link.Type = linkDB.Type
	link.IsBezierCurve = linkDB.IsBezierCurve
	link.StartAnchorType = linkDB.StartAnchorType
	link.EndAnchorType = linkDB.EndAnchorType
	link.StartOrientation = linkDB.StartOrientation
	link.StartRatio = linkDB.StartRatio
	link.EndOrientation = linkDB.EndOrientation
	link.EndRatio = linkDB.EndRatio
	link.CornerOffsetRatio = linkDB.CornerOffsetRatio
	link.CornerRadius = linkDB.CornerRadius
	link.HasEndArrow = linkDB.HasEndArrow
	link.EndArrowSize = linkDB.EndArrowSize
	link.HasStartArrow = linkDB.HasStartArrow
	link.StartArrowSize = linkDB.StartArrowSize
	link.Color = linkDB.Color
	link.FillOpacity = linkDB.FillOpacity
	link.Stroke = linkDB.Stroke
	link.StrokeWidth = linkDB.StrokeWidth
	link.StrokeDashArray = linkDB.StrokeDashArray
	link.StrokeDashArrayWhenSelected = linkDB.StrokeDashArrayWhenSelected
	link.Transform = linkDB.Transform

	// insertion point for pointer fields encoding
	link.Start = frontRepo.map_ID_Rect.get(linkDB.LinkPointersEncoding.StartID.Int64)
	link.End = frontRepo.map_ID_Rect.get(linkDB.LinkPointersEncoding.EndID.Int64)

	// insertion point for slice of pointers fields encoding
	link.TextAtArrowEnd = new Array<LinkAnchoredText>()
	for (let _id of linkDB.LinkPointersEncoding.TextAtArrowEnd) {
		let _linkanchoredtext = frontRepo.map_ID_LinkAnchoredText.get(_id)
		if (_linkanchoredtext != undefined) {
			link.TextAtArrowEnd.push(_linkanchoredtext!)
		}
	}
	link.TextAtArrowStart = new Array<LinkAnchoredText>()
	for (let _id of linkDB.LinkPointersEncoding.TextAtArrowStart) {
		let _linkanchoredtext = frontRepo.map_ID_LinkAnchoredText.get(_id)
		if (_linkanchoredtext != undefined) {
			link.TextAtArrowStart.push(_linkanchoredtext!)
		}
	}
	link.ControlPoints = new Array<Point>()
	for (let _id of linkDB.LinkPointersEncoding.ControlPoints) {
		let _point = frontRepo.map_ID_Point.get(_id)
		if (_point != undefined) {
			link.ControlPoints.push(_point!)
		}
	}
}
