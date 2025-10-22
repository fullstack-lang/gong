// generated code - do not edit

import { RectAnchoredRectAPI } from './rectanchoredrect-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Condition } from './condition'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RectAnchoredRect {

	static GONGSTRUCT_NAME = "RectAnchoredRect"

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
	X_Offset: number = 0
	Y_Offset: number = 0
	RectAnchorType: string = ""
	WidthFollowRect: boolean = false
	HeightFollowRect: boolean = false
	HasToolTip: boolean = false
	ToolTipText: string = ""
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
	HoveringTrigger: Array<Condition> = []
	DisplayConditions: Array<Condition> = []
}

export function CopyRectAnchoredRectToRectAnchoredRectAPI(rectanchoredrect: RectAnchoredRect, rectanchoredrectAPI: RectAnchoredRectAPI) {

	rectanchoredrectAPI.CreatedAt = rectanchoredrect.CreatedAt
	rectanchoredrectAPI.DeletedAt = rectanchoredrect.DeletedAt
	rectanchoredrectAPI.ID = rectanchoredrect.ID

	// insertion point for basic fields copy operations
	rectanchoredrectAPI.Name = rectanchoredrect.Name
	rectanchoredrectAPI.X = rectanchoredrect.X
	rectanchoredrectAPI.Y = rectanchoredrect.Y
	rectanchoredrectAPI.Width = rectanchoredrect.Width
	rectanchoredrectAPI.Height = rectanchoredrect.Height
	rectanchoredrectAPI.RX = rectanchoredrect.RX
	rectanchoredrectAPI.X_Offset = rectanchoredrect.X_Offset
	rectanchoredrectAPI.Y_Offset = rectanchoredrect.Y_Offset
	rectanchoredrectAPI.RectAnchorType = rectanchoredrect.RectAnchorType
	rectanchoredrectAPI.WidthFollowRect = rectanchoredrect.WidthFollowRect
	rectanchoredrectAPI.HeightFollowRect = rectanchoredrect.HeightFollowRect
	rectanchoredrectAPI.HasToolTip = rectanchoredrect.HasToolTip
	rectanchoredrectAPI.ToolTipText = rectanchoredrect.ToolTipText
	rectanchoredrectAPI.Color = rectanchoredrect.Color
	rectanchoredrectAPI.FillOpacity = rectanchoredrect.FillOpacity
	rectanchoredrectAPI.Stroke = rectanchoredrect.Stroke
	rectanchoredrectAPI.StrokeOpacity = rectanchoredrect.StrokeOpacity
	rectanchoredrectAPI.StrokeWidth = rectanchoredrect.StrokeWidth
	rectanchoredrectAPI.StrokeDashArray = rectanchoredrect.StrokeDashArray
	rectanchoredrectAPI.StrokeDashArrayWhenSelected = rectanchoredrect.StrokeDashArrayWhenSelected
	rectanchoredrectAPI.Transform = rectanchoredrect.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	rectanchoredrectAPI.RectAnchoredRectPointersEncoding.HoveringTrigger = []
	for (let _condition of rectanchoredrect.HoveringTrigger) {
		rectanchoredrectAPI.RectAnchoredRectPointersEncoding.HoveringTrigger.push(_condition.ID)
	}

	rectanchoredrectAPI.RectAnchoredRectPointersEncoding.DisplayConditions = []
	for (let _condition of rectanchoredrect.DisplayConditions) {
		rectanchoredrectAPI.RectAnchoredRectPointersEncoding.DisplayConditions.push(_condition.ID)
	}

}

// CopyRectAnchoredRectAPIToRectAnchoredRect update basic, pointers and slice of pointers fields of rectanchoredrect
// from respectively the basic fields and encoded fields of pointers and slices of pointers of rectanchoredrectAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyRectAnchoredRectAPIToRectAnchoredRect(rectanchoredrectAPI: RectAnchoredRectAPI, rectanchoredrect: RectAnchoredRect, frontRepo: FrontRepo) {

	rectanchoredrect.CreatedAt = rectanchoredrectAPI.CreatedAt
	rectanchoredrect.DeletedAt = rectanchoredrectAPI.DeletedAt
	rectanchoredrect.ID = rectanchoredrectAPI.ID

	// insertion point for basic fields copy operations
	rectanchoredrect.Name = rectanchoredrectAPI.Name
	rectanchoredrect.X = rectanchoredrectAPI.X
	rectanchoredrect.Y = rectanchoredrectAPI.Y
	rectanchoredrect.Width = rectanchoredrectAPI.Width
	rectanchoredrect.Height = rectanchoredrectAPI.Height
	rectanchoredrect.RX = rectanchoredrectAPI.RX
	rectanchoredrect.X_Offset = rectanchoredrectAPI.X_Offset
	rectanchoredrect.Y_Offset = rectanchoredrectAPI.Y_Offset
	rectanchoredrect.RectAnchorType = rectanchoredrectAPI.RectAnchorType
	rectanchoredrect.WidthFollowRect = rectanchoredrectAPI.WidthFollowRect
	rectanchoredrect.HeightFollowRect = rectanchoredrectAPI.HeightFollowRect
	rectanchoredrect.HasToolTip = rectanchoredrectAPI.HasToolTip
	rectanchoredrect.ToolTipText = rectanchoredrectAPI.ToolTipText
	rectanchoredrect.Color = rectanchoredrectAPI.Color
	rectanchoredrect.FillOpacity = rectanchoredrectAPI.FillOpacity
	rectanchoredrect.Stroke = rectanchoredrectAPI.Stroke
	rectanchoredrect.StrokeOpacity = rectanchoredrectAPI.StrokeOpacity
	rectanchoredrect.StrokeWidth = rectanchoredrectAPI.StrokeWidth
	rectanchoredrect.StrokeDashArray = rectanchoredrectAPI.StrokeDashArray
	rectanchoredrect.StrokeDashArrayWhenSelected = rectanchoredrectAPI.StrokeDashArrayWhenSelected
	rectanchoredrect.Transform = rectanchoredrectAPI.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(rectanchoredrectAPI.RectAnchoredRectPointersEncoding.HoveringTrigger)) {
		console.error('Rects is not an array:', rectanchoredrectAPI.RectAnchoredRectPointersEncoding.HoveringTrigger);
		return;
	}

	rectanchoredrect.HoveringTrigger = new Array<Condition>()
	for (let _id of rectanchoredrectAPI.RectAnchoredRectPointersEncoding.HoveringTrigger) {
		let _condition = frontRepo.map_ID_Condition.get(_id)
		if (_condition != undefined) {
			rectanchoredrect.HoveringTrigger.push(_condition!)
		}
	}
	if (!Array.isArray(rectanchoredrectAPI.RectAnchoredRectPointersEncoding.DisplayConditions)) {
		console.error('Rects is not an array:', rectanchoredrectAPI.RectAnchoredRectPointersEncoding.DisplayConditions);
		return;
	}

	rectanchoredrect.DisplayConditions = new Array<Condition>()
	for (let _id of rectanchoredrectAPI.RectAnchoredRectPointersEncoding.DisplayConditions) {
		let _condition = frontRepo.map_ID_Condition.get(_id)
		if (_condition != undefined) {
			rectanchoredrect.DisplayConditions.push(_condition!)
		}
	}
}
