// generated code - do not edit

import { RectDB } from './rect-db'
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

export function CopyRectToRectDB(rect: Rect, rectDB: RectDB) {

	rectDB.CreatedAt = rect.CreatedAt
	rectDB.DeletedAt = rect.DeletedAt
	rectDB.ID = rect.ID

	// insertion point for basic fields copy operations
	rectDB.Name = rect.Name
	rectDB.X = rect.X
	rectDB.Y = rect.Y
	rectDB.Width = rect.Width
	rectDB.Height = rect.Height
	rectDB.RX = rect.RX
	rectDB.Color = rect.Color
	rectDB.FillOpacity = rect.FillOpacity
	rectDB.Stroke = rect.Stroke
	rectDB.StrokeWidth = rect.StrokeWidth
	rectDB.StrokeDashArray = rect.StrokeDashArray
	rectDB.StrokeDashArrayWhenSelected = rect.StrokeDashArrayWhenSelected
	rectDB.Transform = rect.Transform
	rectDB.IsSelectable = rect.IsSelectable
	rectDB.IsSelected = rect.IsSelected
	rectDB.CanHaveLeftHandle = rect.CanHaveLeftHandle
	rectDB.HasLeftHandle = rect.HasLeftHandle
	rectDB.CanHaveRightHandle = rect.CanHaveRightHandle
	rectDB.HasRightHandle = rect.HasRightHandle
	rectDB.CanHaveTopHandle = rect.CanHaveTopHandle
	rectDB.HasTopHandle = rect.HasTopHandle
	rectDB.IsScalingProportionally = rect.IsScalingProportionally
	rectDB.CanHaveBottomHandle = rect.CanHaveBottomHandle
	rectDB.HasBottomHandle = rect.HasBottomHandle
	rectDB.CanMoveHorizontaly = rect.CanMoveHorizontaly
	rectDB.CanMoveVerticaly = rect.CanMoveVerticaly

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	rectDB.RectPointersEncoding.Animations = []
	for (let _animate of rect.Animations) {
		rectDB.RectPointersEncoding.Animations.push(_animate.ID)
	}

	rectDB.RectPointersEncoding.RectAnchoredTexts = []
	for (let _rectanchoredtext of rect.RectAnchoredTexts) {
		rectDB.RectPointersEncoding.RectAnchoredTexts.push(_rectanchoredtext.ID)
	}

	rectDB.RectPointersEncoding.RectAnchoredRects = []
	for (let _rectanchoredrect of rect.RectAnchoredRects) {
		rectDB.RectPointersEncoding.RectAnchoredRects.push(_rectanchoredrect.ID)
	}

	rectDB.RectPointersEncoding.RectAnchoredPaths = []
	for (let _rectanchoredpath of rect.RectAnchoredPaths) {
		rectDB.RectPointersEncoding.RectAnchoredPaths.push(_rectanchoredpath.ID)
	}

}

// CopyRectDBToRect update basic, pointers and slice of pointers fields of rect
// from respectively the basic fields and encoded fields of pointers and slices of pointers of rectDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyRectDBToRect(rectDB: RectDB, rect: Rect, frontRepo: FrontRepo) {

	rect.CreatedAt = rectDB.CreatedAt
	rect.DeletedAt = rectDB.DeletedAt
	rect.ID = rectDB.ID

	// insertion point for basic fields copy operations
	rect.Name = rectDB.Name
	rect.X = rectDB.X
	rect.Y = rectDB.Y
	rect.Width = rectDB.Width
	rect.Height = rectDB.Height
	rect.RX = rectDB.RX
	rect.Color = rectDB.Color
	rect.FillOpacity = rectDB.FillOpacity
	rect.Stroke = rectDB.Stroke
	rect.StrokeWidth = rectDB.StrokeWidth
	rect.StrokeDashArray = rectDB.StrokeDashArray
	rect.StrokeDashArrayWhenSelected = rectDB.StrokeDashArrayWhenSelected
	rect.Transform = rectDB.Transform
	rect.IsSelectable = rectDB.IsSelectable
	rect.IsSelected = rectDB.IsSelected
	rect.CanHaveLeftHandle = rectDB.CanHaveLeftHandle
	rect.HasLeftHandle = rectDB.HasLeftHandle
	rect.CanHaveRightHandle = rectDB.CanHaveRightHandle
	rect.HasRightHandle = rectDB.HasRightHandle
	rect.CanHaveTopHandle = rectDB.CanHaveTopHandle
	rect.HasTopHandle = rectDB.HasTopHandle
	rect.IsScalingProportionally = rectDB.IsScalingProportionally
	rect.CanHaveBottomHandle = rectDB.CanHaveBottomHandle
	rect.HasBottomHandle = rectDB.HasBottomHandle
	rect.CanMoveHorizontaly = rectDB.CanMoveHorizontaly
	rect.CanMoveVerticaly = rectDB.CanMoveVerticaly

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	rect.Animations = new Array<Animate>()
	for (let _id of rectDB.RectPointersEncoding.Animations) {
		let _animate = frontRepo.map_ID_Animate.get(_id)
		if (_animate != undefined) {
			rect.Animations.push(_animate!)
		}
	}
	rect.RectAnchoredTexts = new Array<RectAnchoredText>()
	for (let _id of rectDB.RectPointersEncoding.RectAnchoredTexts) {
		let _rectanchoredtext = frontRepo.map_ID_RectAnchoredText.get(_id)
		if (_rectanchoredtext != undefined) {
			rect.RectAnchoredTexts.push(_rectanchoredtext!)
		}
	}
	rect.RectAnchoredRects = new Array<RectAnchoredRect>()
	for (let _id of rectDB.RectPointersEncoding.RectAnchoredRects) {
		let _rectanchoredrect = frontRepo.map_ID_RectAnchoredRect.get(_id)
		if (_rectanchoredrect != undefined) {
			rect.RectAnchoredRects.push(_rectanchoredrect!)
		}
	}
	rect.RectAnchoredPaths = new Array<RectAnchoredPath>()
	for (let _id of rectDB.RectPointersEncoding.RectAnchoredPaths) {
		let _rectanchoredpath = frontRepo.map_ID_RectAnchoredPath.get(_id)
		if (_rectanchoredpath != undefined) {
			rect.RectAnchoredPaths.push(_rectanchoredpath!)
		}
	}
}
