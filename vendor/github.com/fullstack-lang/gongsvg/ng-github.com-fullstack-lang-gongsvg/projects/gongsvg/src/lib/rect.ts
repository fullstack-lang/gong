// generated code - do not edit

import { RectAPI } from './rect-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Animate } from './animate'
import { RectAnchoredText } from './rectanchoredtext'
import { RectAnchoredRect } from './rectanchoredrect'
import { RectAnchoredPath } from './rectanchoredpath'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Rect {

	static GONGSTRUCT_NAME = "Rect"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0
	Width: number = 0
	Height: number = 0
	RX: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""
	IsSelectable: boolean = false
	IsSelected: boolean = false
	CanHaveLeftHandle: boolean = false
	HasLeftHandle: boolean = false
	CanHaveRightHandle: boolean = false
	HasRightHandle: boolean = false
	CanHaveTopHandle: boolean = false
	HasTopHandle: boolean = false
	IsScalingProportionally: boolean = false
	CanHaveBottomHandle: boolean = false
	HasBottomHandle: boolean = false
	CanMoveHorizontaly: boolean = false
	CanMoveVerticaly: boolean = false

	// insertion point for pointers and slices of pointers declarations
	Animations: Array<Animate> = []
	RectAnchoredTexts: Array<RectAnchoredText> = []
	RectAnchoredRects: Array<RectAnchoredRect> = []
	RectAnchoredPaths: Array<RectAnchoredPath> = []
}

export function CopyRectToRectAPI(rect: Rect, rectAPI: RectAPI) {

	rectAPI.CreatedAt = rect.CreatedAt
	rectAPI.DeletedAt = rect.DeletedAt
	rectAPI.ID = rect.ID

	// insertion point for basic fields copy operations
	rectAPI.Name = rect.Name
	rectAPI.X = rect.X
	rectAPI.Y = rect.Y
	rectAPI.Width = rect.Width
	rectAPI.Height = rect.Height
	rectAPI.RX = rect.RX
	rectAPI.Color = rect.Color
	rectAPI.FillOpacity = rect.FillOpacity
	rectAPI.Stroke = rect.Stroke
	rectAPI.StrokeOpacity = rect.StrokeOpacity
	rectAPI.StrokeWidth = rect.StrokeWidth
	rectAPI.StrokeDashArray = rect.StrokeDashArray
	rectAPI.StrokeDashArrayWhenSelected = rect.StrokeDashArrayWhenSelected
	rectAPI.Transform = rect.Transform
	rectAPI.IsSelectable = rect.IsSelectable
	rectAPI.IsSelected = rect.IsSelected
	rectAPI.CanHaveLeftHandle = rect.CanHaveLeftHandle
	rectAPI.HasLeftHandle = rect.HasLeftHandle
	rectAPI.CanHaveRightHandle = rect.CanHaveRightHandle
	rectAPI.HasRightHandle = rect.HasRightHandle
	rectAPI.CanHaveTopHandle = rect.CanHaveTopHandle
	rectAPI.HasTopHandle = rect.HasTopHandle
	rectAPI.IsScalingProportionally = rect.IsScalingProportionally
	rectAPI.CanHaveBottomHandle = rect.CanHaveBottomHandle
	rectAPI.HasBottomHandle = rect.HasBottomHandle
	rectAPI.CanMoveHorizontaly = rect.CanMoveHorizontaly
	rectAPI.CanMoveVerticaly = rect.CanMoveVerticaly

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	rectAPI.RectPointersEncoding.Animations = []
	for (let _animate of rect.Animations) {
		rectAPI.RectPointersEncoding.Animations.push(_animate.ID)
	}

	rectAPI.RectPointersEncoding.RectAnchoredTexts = []
	for (let _rectanchoredtext of rect.RectAnchoredTexts) {
		rectAPI.RectPointersEncoding.RectAnchoredTexts.push(_rectanchoredtext.ID)
	}

	rectAPI.RectPointersEncoding.RectAnchoredRects = []
	for (let _rectanchoredrect of rect.RectAnchoredRects) {
		rectAPI.RectPointersEncoding.RectAnchoredRects.push(_rectanchoredrect.ID)
	}

	rectAPI.RectPointersEncoding.RectAnchoredPaths = []
	for (let _rectanchoredpath of rect.RectAnchoredPaths) {
		rectAPI.RectPointersEncoding.RectAnchoredPaths.push(_rectanchoredpath.ID)
	}

}

// CopyRectAPIToRect update basic, pointers and slice of pointers fields of rect
// from respectively the basic fields and encoded fields of pointers and slices of pointers of rectAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyRectAPIToRect(rectAPI: RectAPI, rect: Rect, frontRepo: FrontRepo) {

	rect.CreatedAt = rectAPI.CreatedAt
	rect.DeletedAt = rectAPI.DeletedAt
	rect.ID = rectAPI.ID

	// insertion point for basic fields copy operations
	rect.Name = rectAPI.Name
	rect.X = rectAPI.X
	rect.Y = rectAPI.Y
	rect.Width = rectAPI.Width
	rect.Height = rectAPI.Height
	rect.RX = rectAPI.RX
	rect.Color = rectAPI.Color
	rect.FillOpacity = rectAPI.FillOpacity
	rect.Stroke = rectAPI.Stroke
	rect.StrokeOpacity = rectAPI.StrokeOpacity
	rect.StrokeWidth = rectAPI.StrokeWidth
	rect.StrokeDashArray = rectAPI.StrokeDashArray
	rect.StrokeDashArrayWhenSelected = rectAPI.StrokeDashArrayWhenSelected
	rect.Transform = rectAPI.Transform
	rect.IsSelectable = rectAPI.IsSelectable
	rect.IsSelected = rectAPI.IsSelected
	rect.CanHaveLeftHandle = rectAPI.CanHaveLeftHandle
	rect.HasLeftHandle = rectAPI.HasLeftHandle
	rect.CanHaveRightHandle = rectAPI.CanHaveRightHandle
	rect.HasRightHandle = rectAPI.HasRightHandle
	rect.CanHaveTopHandle = rectAPI.CanHaveTopHandle
	rect.HasTopHandle = rectAPI.HasTopHandle
	rect.IsScalingProportionally = rectAPI.IsScalingProportionally
	rect.CanHaveBottomHandle = rectAPI.CanHaveBottomHandle
	rect.HasBottomHandle = rectAPI.HasBottomHandle
	rect.CanMoveHorizontaly = rectAPI.CanMoveHorizontaly
	rect.CanMoveVerticaly = rectAPI.CanMoveVerticaly

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(rectAPI.RectPointersEncoding.Animations)) {
		console.error('Rects is not an array:', rectAPI.RectPointersEncoding.Animations);
		return;
	}

	rect.Animations = new Array<Animate>()
	for (let _id of rectAPI.RectPointersEncoding.Animations) {
		let _animate = frontRepo.map_ID_Animate.get(_id)
		if (_animate != undefined) {
			rect.Animations.push(_animate!)
		}
	}
	if (!Array.isArray(rectAPI.RectPointersEncoding.RectAnchoredTexts)) {
		console.error('Rects is not an array:', rectAPI.RectPointersEncoding.RectAnchoredTexts);
		return;
	}

	rect.RectAnchoredTexts = new Array<RectAnchoredText>()
	for (let _id of rectAPI.RectPointersEncoding.RectAnchoredTexts) {
		let _rectanchoredtext = frontRepo.map_ID_RectAnchoredText.get(_id)
		if (_rectanchoredtext != undefined) {
			rect.RectAnchoredTexts.push(_rectanchoredtext!)
		}
	}
	if (!Array.isArray(rectAPI.RectPointersEncoding.RectAnchoredRects)) {
		console.error('Rects is not an array:', rectAPI.RectPointersEncoding.RectAnchoredRects);
		return;
	}

	rect.RectAnchoredRects = new Array<RectAnchoredRect>()
	for (let _id of rectAPI.RectPointersEncoding.RectAnchoredRects) {
		let _rectanchoredrect = frontRepo.map_ID_RectAnchoredRect.get(_id)
		if (_rectanchoredrect != undefined) {
			rect.RectAnchoredRects.push(_rectanchoredrect!)
		}
	}
	if (!Array.isArray(rectAPI.RectPointersEncoding.RectAnchoredPaths)) {
		console.error('Rects is not an array:', rectAPI.RectPointersEncoding.RectAnchoredPaths);
		return;
	}

	rect.RectAnchoredPaths = new Array<RectAnchoredPath>()
	for (let _id of rectAPI.RectPointersEncoding.RectAnchoredPaths) {
		let _rectanchoredpath = frontRepo.map_ID_RectAnchoredPath.get(_id)
		if (_rectanchoredpath != undefined) {
			rect.RectAnchoredPaths.push(_rectanchoredpath!)
		}
	}
}
