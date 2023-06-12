// insertion point for imports
import { CircleDB } from './circle-db'
import { EllipseDB } from './ellipse-db'
import { LineDB } from './line-db'
import { LinkAnchoredTextDB } from './linkanchoredtext-db'
import { PathDB } from './path-db'
import { PolygoneDB } from './polygone-db'
import { PolylineDB } from './polyline-db'
import { RectDB } from './rect-db'
import { RectAnchoredTextDB } from './rectanchoredtext-db'
import { TextDB } from './text-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AnimateDB {

	static GONGSTRUCT_NAME = "Animate"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	AttributeName: string = ""
	Values: string = ""
	Dur: string = ""
	RepeatCount: string = ""

	// insertion point for other declarations
	Circle_AnimationsDBID: NullInt64 = new NullInt64
	Circle_AnimationsDBID_Index: NullInt64  = new NullInt64 // store the index of the animate instance in Circle.Animations
	Circle_Animations_reverse?: CircleDB 

	Ellipse_AnimatesDBID: NullInt64 = new NullInt64
	Ellipse_AnimatesDBID_Index: NullInt64  = new NullInt64 // store the index of the animate instance in Ellipse.Animates
	Ellipse_Animates_reverse?: EllipseDB 

	Line_AnimatesDBID: NullInt64 = new NullInt64
	Line_AnimatesDBID_Index: NullInt64  = new NullInt64 // store the index of the animate instance in Line.Animates
	Line_Animates_reverse?: LineDB 

	LinkAnchoredText_AnimatesDBID: NullInt64 = new NullInt64
	LinkAnchoredText_AnimatesDBID_Index: NullInt64  = new NullInt64 // store the index of the animate instance in LinkAnchoredText.Animates
	LinkAnchoredText_Animates_reverse?: LinkAnchoredTextDB 

	Path_AnimatesDBID: NullInt64 = new NullInt64
	Path_AnimatesDBID_Index: NullInt64  = new NullInt64 // store the index of the animate instance in Path.Animates
	Path_Animates_reverse?: PathDB 

	Polygone_AnimatesDBID: NullInt64 = new NullInt64
	Polygone_AnimatesDBID_Index: NullInt64  = new NullInt64 // store the index of the animate instance in Polygone.Animates
	Polygone_Animates_reverse?: PolygoneDB 

	Polyline_AnimatesDBID: NullInt64 = new NullInt64
	Polyline_AnimatesDBID_Index: NullInt64  = new NullInt64 // store the index of the animate instance in Polyline.Animates
	Polyline_Animates_reverse?: PolylineDB 

	Rect_AnimationsDBID: NullInt64 = new NullInt64
	Rect_AnimationsDBID_Index: NullInt64  = new NullInt64 // store the index of the animate instance in Rect.Animations
	Rect_Animations_reverse?: RectDB 

	RectAnchoredText_AnimatesDBID: NullInt64 = new NullInt64
	RectAnchoredText_AnimatesDBID_Index: NullInt64  = new NullInt64 // store the index of the animate instance in RectAnchoredText.Animates
	RectAnchoredText_Animates_reverse?: RectAnchoredTextDB 

	Text_AnimatesDBID: NullInt64 = new NullInt64
	Text_AnimatesDBID_Index: NullInt64  = new NullInt64 // store the index of the animate instance in Text.Animates
	Text_Animates_reverse?: TextDB 

}
