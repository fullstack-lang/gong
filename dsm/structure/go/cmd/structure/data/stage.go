package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/structure/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Data__00000000_ := (&models.Data{Name: `D1`}).Stage(stage)

	__DataFlow__00000000_ := (&models.DataFlow{Name: `"Port 1" to "Port 2"`}).Stage(stage)

	__DataFlowShape__00000000_ := (&models.DataFlowShape{Name: `"Port 1" to "Port 2"`}).Stage(stage)

	__DataShape__00000000_ := (&models.DataShape{Name: `D1`}).Stage(stage)

	__DiagramStructure__00000000_ := (&models.DiagramStructure{Name: `DiagramStructure`}).Stage(stage)
	__DiagramStructure__00000001_ := (&models.DiagramStructure{Name: `DiagramStructure`}).Stage(stage)
	__DiagramStructure__00000002_ := (&models.DiagramStructure{Name: `sdqsdqsd`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: ``}).Stage(stage)

	__Note__00000000_ := (&models.Note{Name: `Note on the diagram that show that there is wrapping`}).Stage(stage)

	__NotePartShape__00000001_ := (&models.NotePartShape{Name: `P1 shape`}).Stage(stage)

	__NotePortShape__00000001_ := (&models.NotePortShape{Name: `Port 1 shape`}).Stage(stage)

	__NoteShape__00000000_ := (&models.NoteShape{Name: `Note on the diagram that show that there is wrapping shape`}).Stage(stage)

	__Part__00000000_ := (&models.Part{Name: `P1`}).Stage(stage)
	__Part__00000001_ := (&models.Part{Name: `P2`}).Stage(stage)

	__PartAnchoredPath__00000000_ := (&models.PartAnchoredPath{Name: `arrow`}).Stage(stage)

	__PartShape__00000007_ := (&models.PartShape{Name: `P2-DiagramStructure`}).Stage(stage)
	__PartShape__00000009_ := (&models.PartShape{Name: `P1-DiagramStructure`}).Stage(stage)

	__Port__00000000_ := (&models.Port{Name: `Port 1`}).Stage(stage)
	__Port__00000001_ := (&models.Port{Name: `Port 2`}).Stage(stage)
	__Port__00000002_ := (&models.Port{Name: `Port 3`}).Stage(stage)
	__Port__00000003_ := (&models.Port{Name: ``}).Stage(stage)

	__PortShape__00000009_ := (&models.PortShape{Name: `Port 1-DiagramStructure`}).Stage(stage)
	__PortShape__00000010_ := (&models.PortShape{Name: `-DiagramStructure`}).Stage(stage)
	__PortShape__00000013_ := (&models.PortShape{Name: `Port 3-DiagramStructure`}).Stage(stage)
	__PortShape__00000016_ := (&models.PortShape{Name: `-DiagramStructure`}).Stage(stage)

	__System__00000000_ := (&models.System{Name: `System 1`}).Stage(stage)
	__System__00000001_ := (&models.System{Name: ``}).Stage(stage)

	__SystemShape__00000000_ := (&models.SystemShape{Name: `SystemShape`}).Stage(stage)
	__SystemShape__00000001_ := (&models.SystemShape{Name: `SystemShape`}).Stage(stage)
	__SystemShape__00000002_ := (&models.SystemShape{Name: `SystemShape`}).Stage(stage)

	// insertion point for initialization of values

	__Data__00000000_.Name = `D1`
	__Data__00000000_.Acronym = ``
	__Data__00000000_.Description = ``
	__Data__00000000_.ComputedPrefix = ``
	__Data__00000000_.IsExpanded = false
	__Data__00000000_.LayoutDirection = models.Vertical
	__Data__00000000_.SVG_Path = ``
	__Data__00000000_.InverseAppliedScaling = 0.000000

	__DataFlow__00000000_.Name = `"Port 1" to "Port 2"`
	__DataFlow__00000000_.Description = ``
	__DataFlow__00000000_.Direction = models.Unspecified
	__DataFlow__00000000_.IsDatasNodeExpanded = false
	__DataFlow__00000000_.ComputedPrefix = ``
	__DataFlow__00000000_.IsExpanded = false
	__DataFlow__00000000_.LayoutDirection = models.Vertical
	__DataFlow__00000000_.Type = models.DataFlow_Port2Port

	__DataFlowShape__00000000_.Name = `"Port 1" to "Port 2"`
	__DataFlowShape__00000000_.StartRatio = 0.500000
	__DataFlowShape__00000000_.EndRatio = 0.539216
	__DataFlowShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__DataFlowShape__00000000_.CornerOffsetRatio = 2.803797
	__DataFlowShape__00000000_.IsHidden = false

	__DataShape__00000000_.Name = `D1`

	__DiagramStructure__00000000_.Name = `DiagramStructure`
	__DiagramStructure__00000000_.Description = ``
	__DiagramStructure__00000000_.ComputedPrefix = ``
	__DiagramStructure__00000000_.IsExpanded = false
	__DiagramStructure__00000000_.LayoutDirection = models.Vertical
	__DiagramStructure__00000000_.IsChecked = true
	__DiagramStructure__00000000_.IsEditable_ = true
	__DiagramStructure__00000000_.IsShowPrefix = false
	__DiagramStructure__00000000_.DefaultBoxWidth = 250.000000
	__DiagramStructure__00000000_.DefaultBoxHeigth = 70.000000
	__DiagramStructure__00000000_.Width = 1652.999939
	__DiagramStructure__00000000_.Height = 1326.999985
	__DiagramStructure__00000000_.IsSystemsNodeExpanded = false
	__DiagramStructure__00000000_.IsPartsNodeExpanded = true
	__DiagramStructure__00000000_.IsExternalPartsNodeExpanded = false
	__DiagramStructure__00000000_.IsNotesNodeExpanded = true

	__DiagramStructure__00000001_.Name = `DiagramStructure`
	__DiagramStructure__00000001_.Description = ``
	__DiagramStructure__00000001_.ComputedPrefix = ``
	__DiagramStructure__00000001_.IsExpanded = false
	__DiagramStructure__00000001_.LayoutDirection = models.Vertical
	__DiagramStructure__00000001_.IsChecked = false
	__DiagramStructure__00000001_.IsEditable_ = true
	__DiagramStructure__00000001_.IsShowPrefix = false
	__DiagramStructure__00000001_.DefaultBoxWidth = 250.000000
	__DiagramStructure__00000001_.DefaultBoxHeigth = 70.000000
	__DiagramStructure__00000001_.Width = 900.000000
	__DiagramStructure__00000001_.Height = 1350.000000
	__DiagramStructure__00000001_.IsSystemsNodeExpanded = false
	__DiagramStructure__00000001_.IsPartsNodeExpanded = false
	__DiagramStructure__00000001_.IsExternalPartsNodeExpanded = false
	__DiagramStructure__00000001_.IsNotesNodeExpanded = false

	__DiagramStructure__00000002_.Name = `sdqsdqsd`
	__DiagramStructure__00000002_.Description = ``
	__DiagramStructure__00000002_.ComputedPrefix = ``
	__DiagramStructure__00000002_.IsExpanded = true
	__DiagramStructure__00000002_.LayoutDirection = models.Vertical
	__DiagramStructure__00000002_.IsChecked = false
	__DiagramStructure__00000002_.IsEditable_ = true
	__DiagramStructure__00000002_.IsShowPrefix = false
	__DiagramStructure__00000002_.DefaultBoxWidth = 250.000000
	__DiagramStructure__00000002_.DefaultBoxHeigth = 70.000000
	__DiagramStructure__00000002_.Width = 900.000000
	__DiagramStructure__00000002_.Height = 1350.000000
	__DiagramStructure__00000002_.IsSystemsNodeExpanded = false
	__DiagramStructure__00000002_.IsPartsNodeExpanded = false
	__DiagramStructure__00000002_.IsExternalPartsNodeExpanded = false
	__DiagramStructure__00000002_.IsNotesNodeExpanded = false

	__Library__00000000_.Name = ``
	__Library__00000000_.Description = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsExpanded = true
	__Library__00000000_.LayoutDirection = models.Vertical
	__Library__00000000_.IsRootLibrary = true
	__Library__00000000_.IsSubLibrariesNodeExpanded = false
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.IsSystemesNodeExpanded = true
	__Library__00000000_.IsDataFlowsNodeExpanded = false
	__Library__00000000_.IsDatasNodeExpanded = true
	__Library__00000000_.IsResourcesNodeExpanded = false
	__Library__00000000_.IsNotesNodeExpanded = true
	__Library__00000000_.IsExpandedTmp = true

	__Note__00000000_.Name = `Note on the diagram that show that there is wrapping`
	__Note__00000000_.Description = ``
	__Note__00000000_.ComputedPrefix = ``
	__Note__00000000_.IsExpanded = false
	__Note__00000000_.LayoutDirection = models.Vertical
	__Note__00000000_.IsPartsNodeExpanded = true
	__Note__00000000_.IsPortsNodeExpanded = true

	__NotePartShape__00000001_.Name = `P1 shape`
	__NotePartShape__00000001_.StartRatio = 0.000000
	__NotePartShape__00000001_.EndRatio = 0.000000
	__NotePartShape__00000001_.StartOrientation = ""
	__NotePartShape__00000001_.EndOrientation = ""
	__NotePartShape__00000001_.CornerOffsetRatio = 0.000000
	__NotePartShape__00000001_.IsHidden = false

	__NotePortShape__00000001_.Name = `Port 1 shape`
	__NotePortShape__00000001_.StartRatio = 0.000000
	__NotePortShape__00000001_.EndRatio = 0.000000
	__NotePortShape__00000001_.StartOrientation = ""
	__NotePortShape__00000001_.EndOrientation = ""
	__NotePortShape__00000001_.CornerOffsetRatio = 0.000000
	__NotePortShape__00000001_.IsHidden = false

	__NoteShape__00000000_.Name = `Note on the diagram that show that there is wrapping shape`
	__NoteShape__00000000_.X = 183.203176
	__NoteShape__00000000_.Y = 132.642074
	__NoteShape__00000000_.Width = 202.000000
	__NoteShape__00000000_.Height = 70.000000
	__NoteShape__00000000_.IsHidden = false

	__Part__00000000_.Name = `P1`
	__Part__00000000_.Description = ``
	__Part__00000000_.IsPartNameNotSystemName = false
	__Part__00000000_.IsControlFlowsNodeExpanded = false
	__Part__00000000_.IsDataFlowsNodeExpanded = false
	__Part__00000000_.ComputedPrefix = ``
	__Part__00000000_.IsExpanded = false
	__Part__00000000_.LayoutDirection = models.Vertical
	__Part__00000000_.IsPortsNodeExpanded = true

	__Part__00000001_.Name = `P2`
	__Part__00000001_.Description = ``
	__Part__00000001_.IsPartNameNotSystemName = false
	__Part__00000001_.IsControlFlowsNodeExpanded = false
	__Part__00000001_.IsDataFlowsNodeExpanded = false
	__Part__00000001_.ComputedPrefix = ``
	__Part__00000001_.IsExpanded = false
	__Part__00000001_.LayoutDirection = models.Vertical
	__Part__00000001_.IsPortsNodeExpanded = true

	__PartAnchoredPath__00000000_.Name = `arrow`
	__PartAnchoredPath__00000000_.Definition = `M 0,41 30,11 20,1 l 30,0 V 31 L 40,21 10,51 Z`
	__PartAnchoredPath__00000000_.X_Offset = 0.000000
	__PartAnchoredPath__00000000_.Y_Offset = 0.000000
	__PartAnchoredPath__00000000_.RectAnchorType = models.RECT_CENTER
	__PartAnchoredPath__00000000_.ScalePropotionnally = false
	__PartAnchoredPath__00000000_.AppliedScaling = 0.000000
	__PartAnchoredPath__00000000_.Color = `blue`
	__PartAnchoredPath__00000000_.FillOpacity = 1.000000
	__PartAnchoredPath__00000000_.Stroke = `blue`
	__PartAnchoredPath__00000000_.StrokeOpacity = 1.000000
	__PartAnchoredPath__00000000_.StrokeWidth = 1.000000
	__PartAnchoredPath__00000000_.StrokeDashArray = ``
	__PartAnchoredPath__00000000_.StrokeDashArrayWhenSelected = ``
	__PartAnchoredPath__00000000_.Transform = ``

	__PartShape__00000007_.Name = `P2-DiagramStructure`
	__PartShape__00000007_.IsExpanded = false
	__PartShape__00000007_.X = 767.000000
	__PartShape__00000007_.Y = 438.999985
	__PartShape__00000007_.Width = 329.000000
	__PartShape__00000007_.Height = 282.000000
	__PartShape__00000007_.IsHidden = false

	__PartShape__00000009_.Name = `P1-DiagramStructure`
	__PartShape__00000009_.IsExpanded = false
	__PartShape__00000009_.X = 203.000000
	__PartShape__00000009_.Y = 332.897868
	__PartShape__00000009_.Width = 250.000000
	__PartShape__00000009_.Height = 144.999985
	__PartShape__00000009_.IsHidden = false

	__Port__00000000_.Name = `Port 1`
	__Port__00000000_.Description = ``
	__Port__00000000_.ComputedPrefix = ``
	__Port__00000000_.IsExpanded = false
	__Port__00000000_.LayoutDirection = models.Vertical

	__Port__00000001_.Name = `Port 2`
	__Port__00000001_.Description = ``
	__Port__00000001_.ComputedPrefix = ``
	__Port__00000001_.IsExpanded = false
	__Port__00000001_.LayoutDirection = models.Vertical

	__Port__00000002_.Name = `Port 3`
	__Port__00000002_.Description = ``
	__Port__00000002_.ComputedPrefix = ``
	__Port__00000002_.IsExpanded = false
	__Port__00000002_.LayoutDirection = models.Vertical

	__Port__00000003_.Name = ``
	__Port__00000003_.Description = ``
	__Port__00000003_.ComputedPrefix = ``
	__Port__00000003_.IsExpanded = false
	__Port__00000003_.LayoutDirection = models.Vertical

	__PortShape__00000009_.Name = `Port 1-DiagramStructure`
	__PortShape__00000009_.IsExpanded = false
	__PortShape__00000009_.X = 419.000000
	__PortShape__00000009_.Y = 407.897837
	__PortShape__00000009_.Width = 68.000000
	__PortShape__00000009_.Height = 40.000000
	__PortShape__00000009_.IsHidden = false

	__PortShape__00000010_.Name = `-DiagramStructure`
	__PortShape__00000010_.IsExpanded = false
	__PortShape__00000010_.X = 711.000000
	__PortShape__00000010_.Y = 493.999985
	__PortShape__00000010_.Width = 112.000000
	__PortShape__00000010_.Height = 34.000000
	__PortShape__00000010_.IsHidden = false

	__PortShape__00000013_.Name = `Port 3-DiagramStructure`
	__PortShape__00000013_.IsExpanded = false
	__PortShape__00000013_.X = 183.000000
	__PortShape__00000013_.Y = 340.397861
	__PortShape__00000013_.Width = 40.000000
	__PortShape__00000013_.Height = 40.000000
	__PortShape__00000013_.IsHidden = false

	__PortShape__00000016_.Name = `-DiagramStructure`
	__PortShape__00000016_.IsExpanded = false
	__PortShape__00000016_.X = 183.000000
	__PortShape__00000016_.Y = 405.397861
	__PortShape__00000016_.Width = 40.000000
	__PortShape__00000016_.Height = 40.000000
	__PortShape__00000016_.IsHidden = false

	__System__00000000_.Name = `System 1`
	__System__00000000_.Description = ``
	__System__00000000_.ComputedPrefix = ``
	__System__00000000_.IsExpanded = true
	__System__00000000_.LayoutDirection = models.Vertical
	__System__00000000_.SVG_Path = ``
	__System__00000000_.InverseAppliedScaling = 0.000000
	__System__00000000_.IsSubSystemNodeExpanded = false
	__System__00000000_.IsDataFlowsNodeExpanded = true

	__System__00000001_.Name = ``
	__System__00000001_.Description = ``
	__System__00000001_.ComputedPrefix = ``
	__System__00000001_.IsExpanded = false
	__System__00000001_.LayoutDirection = models.Vertical
	__System__00000001_.SVG_Path = ``
	__System__00000001_.InverseAppliedScaling = 0.000000
	__System__00000001_.IsSubSystemNodeExpanded = false
	__System__00000001_.IsDataFlowsNodeExpanded = false

	__SystemShape__00000000_.Name = `SystemShape`
	__SystemShape__00000000_.IsExpanded = false
	__SystemShape__00000000_.X = 94.000000
	__SystemShape__00000000_.Y = -1.000015
	__SystemShape__00000000_.Width = 1258.999939
	__SystemShape__00000000_.Height = 1028.000000
	__SystemShape__00000000_.IsHidden = false

	__SystemShape__00000001_.Name = `SystemShape`
	__SystemShape__00000001_.IsExpanded = false
	__SystemShape__00000001_.X = 100.000000
	__SystemShape__00000001_.Y = 50.000000
	__SystemShape__00000001_.Width = 500.000000
	__SystemShape__00000001_.Height = 1000.000000
	__SystemShape__00000001_.IsHidden = false

	__SystemShape__00000002_.Name = `SystemShape`
	__SystemShape__00000002_.IsExpanded = false
	__SystemShape__00000002_.X = 100.000000
	__SystemShape__00000002_.Y = 50.000000
	__SystemShape__00000002_.Width = 500.000000
	__SystemShape__00000002_.Height = 1000.000000
	__SystemShape__00000002_.IsHidden = false

	// insertion point for setup of pointers
	__DataFlow__00000000_.StartPort = __Port__00000000_
	__DataFlow__00000000_.EndPort = __Port__00000001_
	__DataFlow__00000000_.StartExternalPart = nil
	__DataFlow__00000000_.EndExternalPart = nil
	__DataFlow__00000000_.Datas = append(__DataFlow__00000000_.Datas, __Data__00000000_)
	__DataFlowShape__00000000_.DataFlow = __DataFlow__00000000_
	__DataShape__00000000_.Data = __Data__00000000_
	__DataShape__00000000_.DataFlow = __DataFlow__00000000_
	__DiagramStructure__00000000_.System_Shapes = append(__DiagramStructure__00000000_.System_Shapes, __SystemShape__00000000_)
	__DiagramStructure__00000000_.Part_Shapes = append(__DiagramStructure__00000000_.Part_Shapes, __PartShape__00000007_)
	__DiagramStructure__00000000_.Part_Shapes = append(__DiagramStructure__00000000_.Part_Shapes, __PartShape__00000009_)
	__DiagramStructure__00000000_.PartWhoseNodeIsExpanded = append(__DiagramStructure__00000000_.PartWhoseNodeIsExpanded, __Part__00000000_)
	__DiagramStructure__00000000_.PortsWhoseNodeIsExpanded = append(__DiagramStructure__00000000_.PortsWhoseNodeIsExpanded, __Port__00000000_)
	__DiagramStructure__00000000_.Port_Shapes = append(__DiagramStructure__00000000_.Port_Shapes, __PortShape__00000009_)
	__DiagramStructure__00000000_.Port_Shapes = append(__DiagramStructure__00000000_.Port_Shapes, __PortShape__00000010_)
	__DiagramStructure__00000000_.Port_Shapes = append(__DiagramStructure__00000000_.Port_Shapes, __PortShape__00000013_)
	__DiagramStructure__00000000_.Port_Shapes = append(__DiagramStructure__00000000_.Port_Shapes, __PortShape__00000016_)
	__DiagramStructure__00000000_.DataFlow_Shapes = append(__DiagramStructure__00000000_.DataFlow_Shapes, __DataFlowShape__00000000_)
	__DiagramStructure__00000000_.Data_Shapes = append(__DiagramStructure__00000000_.Data_Shapes, __DataShape__00000000_)
	__DiagramStructure__00000000_.Note_Shapes = append(__DiagramStructure__00000000_.Note_Shapes, __NoteShape__00000000_)
	__DiagramStructure__00000000_.NotesWhoseNodeIsExpanded = append(__DiagramStructure__00000000_.NotesWhoseNodeIsExpanded, __Note__00000000_)
	__DiagramStructure__00000000_.NotePortShapes = append(__DiagramStructure__00000000_.NotePortShapes, __NotePortShape__00000001_)
	__DiagramStructure__00000000_.NotePartShapes = append(__DiagramStructure__00000000_.NotePartShapes, __NotePartShape__00000001_)
	__DiagramStructure__00000001_.System_Shapes = append(__DiagramStructure__00000001_.System_Shapes, __SystemShape__00000001_)
	__DiagramStructure__00000002_.System_Shapes = append(__DiagramStructure__00000002_.System_Shapes, __SystemShape__00000002_)
	__Library__00000000_.RootSystemes = append(__Library__00000000_.RootSystemes, __System__00000000_)
	__Library__00000000_.SystemsWhoseNodeIsExpanded = append(__Library__00000000_.SystemsWhoseNodeIsExpanded, __System__00000000_)
	__Library__00000000_.RootDataFlows = append(__Library__00000000_.RootDataFlows, __DataFlow__00000000_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000000_)
	__Library__00000000_.RootNotes = append(__Library__00000000_.RootNotes, __Note__00000000_)
	__Note__00000000_.Parts = append(__Note__00000000_.Parts, __Part__00000000_)
	__Note__00000000_.Ports = append(__Note__00000000_.Ports, __Port__00000000_)
	__NotePartShape__00000001_.Note = __Note__00000000_
	__NotePartShape__00000001_.Part = __Part__00000000_
	__NotePortShape__00000001_.Note = __Note__00000000_
	__NotePortShape__00000001_.Port = __Port__00000000_
	__NoteShape__00000000_.Note = __Note__00000000_
	__Part__00000000_.Ports = append(__Part__00000000_.Ports, __Port__00000000_)
	__Part__00000000_.Ports = append(__Part__00000000_.Ports, __Port__00000002_)
	__Part__00000000_.Ports = append(__Part__00000000_.Ports, __Port__00000003_)
	__Part__00000000_.TypeOfPart = nil
	__Part__00000000_.PortWhoseOutDataFlowsNodeIsExpanded = append(__Part__00000000_.PortWhoseOutDataFlowsNodeIsExpanded, __Port__00000000_)
	__Part__00000001_.Ports = append(__Part__00000001_.Ports, __Port__00000001_)
	__Part__00000001_.TypeOfPart = nil
	__Part__00000001_.PartAnchoredPath = append(__Part__00000001_.PartAnchoredPath, __PartAnchoredPath__00000000_)
	__PartShape__00000007_.Part = __Part__00000001_
	__PartShape__00000009_.Part = __Part__00000000_
	__PortShape__00000009_.Port = __Port__00000000_
	__PortShape__00000010_.Port = __Port__00000001_
	__PortShape__00000013_.Port = __Port__00000002_
	__PortShape__00000016_.Port = __Port__00000003_
	__System__00000000_.DiagramStructures = append(__System__00000000_.DiagramStructures, __DiagramStructure__00000000_)
	__System__00000000_.DiagramStructures = append(__System__00000000_.DiagramStructures, __DiagramStructure__00000002_)
	__System__00000000_.DiagramStructureWhoseNodeIsExpanded = append(__System__00000000_.DiagramStructureWhoseNodeIsExpanded, __DiagramStructure__00000000_)
	__System__00000000_.SubSystemes = append(__System__00000000_.SubSystemes, __System__00000001_)
	__System__00000000_.Parts = append(__System__00000000_.Parts, __Part__00000000_)
	__System__00000000_.Parts = append(__System__00000000_.Parts, __Part__00000001_)
	__System__00000000_.DataFlows = append(__System__00000000_.DataFlows, __DataFlow__00000000_)
	__System__00000001_.DiagramStructures = append(__System__00000001_.DiagramStructures, __DiagramStructure__00000001_)
	__SystemShape__00000000_.System = __System__00000000_
	__SystemShape__00000001_.System = __System__00000001_
	__SystemShape__00000002_.System = __System__00000000_
}
