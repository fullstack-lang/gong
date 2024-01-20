// generated code - do not edit

import { RectAnchoredRectDB } from './rectanchoredrect-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

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
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyRectAnchoredRectToRectAnchoredRectDB(rectanchoredrect: RectAnchoredRect, rectanchoredrectDB: RectAnchoredRectDB) {

	rectanchoredrectDB.CreatedAt = rectanchoredrect.CreatedAt
	rectanchoredrectDB.DeletedAt = rectanchoredrect.DeletedAt
	rectanchoredrectDB.ID = rectanchoredrect.ID
	
	// insertion point for basic fields copy operations
	rectanchoredrectDB.Name = rectanchoredrect.Name
	rectanchoredrectDB.X = rectanchoredrect.X
	rectanchoredrectDB.Y = rectanchoredrect.Y
	rectanchoredrectDB.Width = rectanchoredrect.Width
	rectanchoredrectDB.Height = rectanchoredrect.Height
	rectanchoredrectDB.RX = rectanchoredrect.RX
	rectanchoredrectDB.X_Offset = rectanchoredrect.X_Offset
	rectanchoredrectDB.Y_Offset = rectanchoredrect.Y_Offset
	rectanchoredrectDB.RectAnchorType = rectanchoredrect.RectAnchorType
	rectanchoredrectDB.WidthFollowRect = rectanchoredrect.WidthFollowRect
	rectanchoredrectDB.HeightFollowRect = rectanchoredrect.HeightFollowRect
	rectanchoredrectDB.Color = rectanchoredrect.Color
	rectanchoredrectDB.FillOpacity = rectanchoredrect.FillOpacity
	rectanchoredrectDB.Stroke = rectanchoredrect.Stroke
	rectanchoredrectDB.StrokeWidth = rectanchoredrect.StrokeWidth
	rectanchoredrectDB.StrokeDashArray = rectanchoredrect.StrokeDashArray
	rectanchoredrectDB.StrokeDashArrayWhenSelected = rectanchoredrect.StrokeDashArrayWhenSelected
	rectanchoredrectDB.Transform = rectanchoredrect.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

export function CopyRectAnchoredRectDBToRectAnchoredRect(rectanchoredrectDB: RectAnchoredRectDB, rectanchoredrect: RectAnchoredRect, frontRepo: FrontRepo) {

	rectanchoredrect.CreatedAt = rectanchoredrectDB.CreatedAt
	rectanchoredrect.DeletedAt = rectanchoredrectDB.DeletedAt
	rectanchoredrect.ID = rectanchoredrectDB.ID
	
	// insertion point for basic fields copy operations
	rectanchoredrect.Name = rectanchoredrectDB.Name
	rectanchoredrect.X = rectanchoredrectDB.X
	rectanchoredrect.Y = rectanchoredrectDB.Y
	rectanchoredrect.Width = rectanchoredrectDB.Width
	rectanchoredrect.Height = rectanchoredrectDB.Height
	rectanchoredrect.RX = rectanchoredrectDB.RX
	rectanchoredrect.X_Offset = rectanchoredrectDB.X_Offset
	rectanchoredrect.Y_Offset = rectanchoredrectDB.Y_Offset
	rectanchoredrect.RectAnchorType = rectanchoredrectDB.RectAnchorType
	rectanchoredrect.WidthFollowRect = rectanchoredrectDB.WidthFollowRect
	rectanchoredrect.HeightFollowRect = rectanchoredrectDB.HeightFollowRect
	rectanchoredrect.Color = rectanchoredrectDB.Color
	rectanchoredrect.FillOpacity = rectanchoredrectDB.FillOpacity
	rectanchoredrect.Stroke = rectanchoredrectDB.Stroke
	rectanchoredrect.StrokeWidth = rectanchoredrectDB.StrokeWidth
	rectanchoredrect.StrokeDashArray = rectanchoredrectDB.StrokeDashArray
	rectanchoredrect.StrokeDashArrayWhenSelected = rectanchoredrectDB.StrokeDashArrayWhenSelected
	rectanchoredrect.Transform = rectanchoredrectDB.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
