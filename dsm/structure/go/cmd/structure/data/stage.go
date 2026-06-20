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

	__DiagramStructure__00000000_ := (&models.DiagramStructure{Name: `DiagramStructure`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: ``}).Stage(stage)

	__Part__00000000_ := (&models.Part{Name: `P1`}).Stage(stage)
	__Part__00000001_ := (&models.Part{Name: `P2`}).Stage(stage)

	__PartShape__00000007_ := (&models.PartShape{Name: `P2-DiagramStructure`}).Stage(stage)
	__PartShape__00000008_ := (&models.PartShape{Name: `P1-DiagramStructure`}).Stage(stage)

	__Port__00000000_ := (&models.Port{Name: `Port 1`}).Stage(stage)

	__PortShape__00000002_ := (&models.PortShape{Name: `Port 1-DiagramStructure`}).Stage(stage)

	__System__00000000_ := (&models.System{Name: `System 1`}).Stage(stage)

	__SystemShape__00000000_ := (&models.SystemShape{Name: `SystemShape`}).Stage(stage)

	// insertion point for initialization of values

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
	__DiagramStructure__00000000_.Width = 1618.999939
	__DiagramStructure__00000000_.Height = 1327.000000
	__DiagramStructure__00000000_.IsSystemsNodeExpanded = false
	__DiagramStructure__00000000_.IsPartsNodeExpanded = true
	__DiagramStructure__00000000_.IsExternalPartsNodeExpanded = false
	__DiagramStructure__00000000_.IsNotesNodeExpanded = false

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
	__Library__00000000_.IsDatasNodeExpanded = false
	__Library__00000000_.IsResourcesNodeExpanded = false
	__Library__00000000_.IsNotesNodeExpanded = false
	__Library__00000000_.IsExpandedTmp = true

	__Part__00000000_.Name = `P1`
	__Part__00000000_.Description = ``
	__Part__00000000_.ComputedPrefix = ``
	__Part__00000000_.IsExpanded = false
	__Part__00000000_.LayoutDirection = models.Vertical
	__Part__00000000_.IsPortsNodeExpanded = true
	__Part__00000000_.IsControlFlowsNodeExpanded = false
	__Part__00000000_.IsDataFlowsNodeExpanded = false

	__Part__00000001_.Name = `P2`
	__Part__00000001_.Description = ``
	__Part__00000001_.ComputedPrefix = ``
	__Part__00000001_.IsExpanded = false
	__Part__00000001_.LayoutDirection = models.Vertical
	__Part__00000001_.IsPortsNodeExpanded = false
	__Part__00000001_.IsControlFlowsNodeExpanded = false
	__Part__00000001_.IsDataFlowsNodeExpanded = false

	__PartShape__00000007_.Name = `P2-DiagramStructure`
	__PartShape__00000007_.IsExpanded = false
	__PartShape__00000007_.X = 779.000000
	__PartShape__00000007_.Y = 279.999985
	__PartShape__00000007_.Width = 411.000000
	__PartShape__00000007_.Height = 282.000000
	__PartShape__00000007_.IsHidden = false

	__PartShape__00000008_.Name = `P1-DiagramStructure`
	__PartShape__00000008_.IsExpanded = false
	__PartShape__00000008_.X = 186.999939
	__PartShape__00000008_.Y = 347.999985
	__PartShape__00000008_.Width = 415.000000
	__PartShape__00000008_.Height = 387.000000
	__PartShape__00000008_.IsHidden = false

	__Port__00000000_.Name = `Port 1`
	__Port__00000000_.Description = ``
	__Port__00000000_.ComputedPrefix = ``
	__Port__00000000_.IsExpanded = false
	__Port__00000000_.LayoutDirection = models.Vertical
	__Port__00000000_.IsStartPort = false
	__Port__00000000_.IsEndPort = false
	__Port__00000000_.IsPortNameNotSystemName = false

	__PortShape__00000002_.Name = `Port 1-DiagramStructure`
	__PortShape__00000002_.IsExpanded = false
	__PortShape__00000002_.X = 561.999939
	__PortShape__00000002_.Y = 404.499985
	__PortShape__00000002_.Width = 40.000000
	__PortShape__00000002_.Height = 40.000000
	__PortShape__00000002_.IsHidden = false

	__System__00000000_.Name = `System 1`
	__System__00000000_.Description = ``
	__System__00000000_.ComputedPrefix = ``
	__System__00000000_.IsExpanded = false
	__System__00000000_.LayoutDirection = models.Vertical
	__System__00000000_.SVG_Path = ``
	__System__00000000_.InverseAppliedScaling = 0.000000
	__System__00000000_.IsSubSystemNodeExpanded = false
	__System__00000000_.IsDataFlowsNodeExpanded = false

	__SystemShape__00000000_.Name = `SystemShape`
	__SystemShape__00000000_.IsExpanded = false
	__SystemShape__00000000_.X = 60.000000
	__SystemShape__00000000_.Y = 42.000000
	__SystemShape__00000000_.Width = 1258.999939
	__SystemShape__00000000_.Height = 985.000000
	__SystemShape__00000000_.IsHidden = false

	// insertion point for setup of pointers
	__DiagramStructure__00000000_.System_Shapes = append(__DiagramStructure__00000000_.System_Shapes, __SystemShape__00000000_)
	__DiagramStructure__00000000_.Part_Shapes = append(__DiagramStructure__00000000_.Part_Shapes, __PartShape__00000007_)
	__DiagramStructure__00000000_.Part_Shapes = append(__DiagramStructure__00000000_.Part_Shapes, __PartShape__00000008_)
	__DiagramStructure__00000000_.PartWhoseNodeIsExpanded = append(__DiagramStructure__00000000_.PartWhoseNodeIsExpanded, __Part__00000000_)
	__DiagramStructure__00000000_.Port_Shapes = append(__DiagramStructure__00000000_.Port_Shapes, __PortShape__00000002_)
	__Library__00000000_.RootSystemes = append(__Library__00000000_.RootSystemes, __System__00000000_)
	__Library__00000000_.SystemsWhoseNodeIsExpanded = append(__Library__00000000_.SystemsWhoseNodeIsExpanded, __System__00000000_)
	__Part__00000000_.Ports = append(__Part__00000000_.Ports, __Port__00000000_)
	__PartShape__00000007_.Part = __Part__00000001_
	__PartShape__00000008_.Part = __Part__00000000_
	__Port__00000000_.Type = nil
	__PortShape__00000002_.Port = __Port__00000000_
	__System__00000000_.DiagramStructures = append(__System__00000000_.DiagramStructures, __DiagramStructure__00000000_)
	__System__00000000_.DiagramStructureWhoseNodeIsExpanded = append(__System__00000000_.DiagramStructureWhoseNodeIsExpanded, __DiagramStructure__00000000_)
	__System__00000000_.Parts = append(__System__00000000_.Parts, __Part__00000000_)
	__System__00000000_.Parts = append(__System__00000000_.Parts, __Part__00000001_)
	__SystemShape__00000000_.System = __System__00000000_
}
