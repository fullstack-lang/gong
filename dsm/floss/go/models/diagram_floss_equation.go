package models

import svg "github.com/fullstack-lang/gong/lib/svg/go/models"

type DiagramFlossEquation struct {
	Name string

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields

	IsChecked   bool
	IsEditable_ bool

	IsInDelta3ColumnsMode bool

	AreQuantitativeElementsVisible bool
	AreSubsystemsVisible           bool

	Width float64

	Height float64

	Scale float64 // pixels per unit

	DefaultBoxWidth  float64
	DefaultBoxHeigth float64

	// Note
	Note_Shapes                   []*NoteShape
	NoteComplexityShapes          []*NoteComplexityShape
	NotePerformanceShapes         []*NotePerformanceShape
	NoteEffortShapes              []*NoteEffortShape
	map_Note_NoteShape            map[*Note]*NoteShape
	map_Note_Rect                 map[*Note]*svg.Rect
	map_Note_NoteComplexityShape  map[noteComplexityKey]*NoteComplexityShape
	map_Note_NotePerformanceShape map[notePerformanceKey]*NotePerformanceShape
	map_Note_NoteEffortShape      map[noteEffortKey]*NoteEffortShape
	IsNotesNodeExpanded           bool
	NotesWhoseNodeIsExpanded      []*Note

	// CPE node expansion
	IsComplexitysNodeExpanded       bool
	ComplexitysWhoseNodeIsExpanded  []*Complexity
	IsPerformancesNodeExpanded      bool
	PerformancesWhoseNodeIsExpanded []*Performance
	IsEffortsNodeExpanded           bool
	EffortsWhoseNodeIsExpanded      []*Effort

	map_SvgRect_NoteShape   map[*svg.Rect]*NoteShape
	map_SvgRect_Complexity  map[*svg.Rect]*Complexity
	map_SvgRect_Performance map[*svg.Rect]*Performance
	map_SvgRect_Effort      map[*svg.Rect]*Effort

	owningCompareAnalysis *CompareAnalysis
	owningSystem          *System
}

func (d *DiagramFlossEquation) IsEditable() bool {
	return d.IsEditable_
}

func (d *DiagramFlossEquation) SetEditable(v bool) {
	d.IsEditable_ = v
}

func (d *DiagramFlossEquation) GetIsChecked() bool {
	return d.IsChecked
}

func (d *DiagramFlossEquation) SetIsChecked(v bool) {
	d.IsChecked = v
}

func (d *DiagramFlossEquation) GetDefaultBoxHeigth() float64 {
	return d.DefaultBoxHeigth
}

func (d *DiagramFlossEquation) GetDefaultBoxWidth() float64 {
	return d.DefaultBoxWidth
}

func (d *DiagramFlossEquation) GetOwningCompareAnalysis() *CompareAnalysis {
	return d.owningCompareAnalysis
}

func (d *DiagramFlossEquation) SetOwningCompareAnalysis(ca *CompareAnalysis) {
	d.owningCompareAnalysis = ca
}

func (d *DiagramFlossEquation) GetOwningSystem() *System {
	return d.owningSystem
}

func (d *DiagramFlossEquation) SetOwningSystem(s *System) {
	d.owningSystem = s
}
