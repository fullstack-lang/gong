package models

type Plant3DDiagram struct {
	Name string

	IsHiddenStemCylinder3DShape bool
	StemCylinder3DShape         *StemCylinder3DShape

	IsHiddenParastichyNCurves3DShape bool
	ParastichyNCurves3DShape         *ParastichyNCurves3DShape

	IsHiddenParastichyMCurves3DShape bool
	ParastichyMCurves3DShape         *ParastichyMCurves3DShape

	IsHiddenCutLine3DShape bool
	CutLine3DShape         *CutLine3DShape

	IsHiddenCircumference3DShape bool
	Circumference3DShape         *Circumference3DShape

	IsHiddenTiledFloor3DShape bool
	TiledFloor3DShape         *TiledFloor3DShape

	IsHiddenLeaves3DShape bool
	Leaves3DShape         *Leaves3DShape

	Rendered3DShape *Rendered3DShape

	IsChecked bool
	AbstractTypeFields
}

type StemCylinder3DShape struct {
	Name string
}

type ParastichyNCurves3DShape struct {
	Name string
}

type ParastichyMCurves3DShape struct {
	Name string
}

type CutLine3DShape struct {
	Name string
}

type Circumference3DShape struct {
	Name string
}

type Leaves3DShape struct {
	Name string
}
