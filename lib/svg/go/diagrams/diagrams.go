package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_SVG := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_Layer := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_Circle := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_Circles := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000001_Layers := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansionBinaryEncoding = 65544
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansionBinaryEncoding = 0
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansionBinaryEncoding = 0

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.Name = `Diagram Package created the 2025-05-05T06:12:28Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_SVG.Name = `Default-SVG`
	__GongStructShape__000000_Default_SVG.X = 26.000000
	__GongStructShape__000000_Default_SVG.Y = 83.000000

	//gong:ident [ref_models.SVG] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_SVG.Identifier = `ref_models.SVG`
	__GongStructShape__000000_Default_SVG.ShowNbInstances = false
	__GongStructShape__000000_Default_SVG.NbInstances = 0
	__GongStructShape__000000_Default_SVG.Width = 240.000000
	__GongStructShape__000000_Default_SVG.Height = 63.000000
	__GongStructShape__000000_Default_SVG.IsSelected = false

	__GongStructShape__000001_Default_Layer.Name = `Default-Layer`
	__GongStructShape__000001_Default_Layer.X = 425.000000
	__GongStructShape__000001_Default_Layer.Y = 81.000000

	//gong:ident [ref_models.Layer] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_Layer.Identifier = `ref_models.Layer`
	__GongStructShape__000001_Default_Layer.ShowNbInstances = false
	__GongStructShape__000001_Default_Layer.NbInstances = 0
	__GongStructShape__000001_Default_Layer.Width = 240.000000
	__GongStructShape__000001_Default_Layer.Height = 63.000000
	__GongStructShape__000001_Default_Layer.IsSelected = false

	__GongStructShape__000002_Default_Circle.Name = `Default-Circle`
	__GongStructShape__000002_Default_Circle.X = 819.000000
	__GongStructShape__000002_Default_Circle.Y = 90.000000

	//gong:ident [ref_models.Circle] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_Circle.Identifier = `ref_models.Circle`
	__GongStructShape__000002_Default_Circle.ShowNbInstances = false
	__GongStructShape__000002_Default_Circle.NbInstances = 0
	__GongStructShape__000002_Default_Circle.Width = 240.000000
	__GongStructShape__000002_Default_Circle.Height = 63.000000
	__GongStructShape__000002_Default_Circle.IsSelected = false

	__LinkShape__000000_Circles.Name = `Circles`

	//gong:ident [ref_models.Layer.Circles] comment added to overcome the problem with the comment map association
	__LinkShape__000000_Circles.Identifier = `ref_models.Layer.Circles`

	//gong:ident [ref_models.Circle] comment added to overcome the problem with the comment map association
	__LinkShape__000000_Circles.Fieldtypename = `ref_models.Circle`
	__LinkShape__000000_Circles.FieldOffsetX = 0.000000
	__LinkShape__000000_Circles.FieldOffsetY = 0.000000
	__LinkShape__000000_Circles.TargetMultiplicity = models.MANY
	__LinkShape__000000_Circles.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Circles.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Circles.SourceMultiplicity = models.MANY
	__LinkShape__000000_Circles.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Circles.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Circles.X = 394.500000
	__LinkShape__000000_Circles.Y = 101.500000
	__LinkShape__000000_Circles.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Circles.StartRatio = 0.500000
	__LinkShape__000000_Circles.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Circles.EndRatio = 0.500000
	__LinkShape__000000_Circles.CornerOffsetRatio = 1.380000

	__LinkShape__000001_Layers.Name = `Layers`

	//gong:ident [ref_models.SVG.Layers] comment added to overcome the problem with the comment map association
	__LinkShape__000001_Layers.Identifier = `ref_models.SVG.Layers`

	//gong:ident [ref_models.Layer] comment added to overcome the problem with the comment map association
	__LinkShape__000001_Layers.Fieldtypename = `ref_models.Layer`
	__LinkShape__000001_Layers.FieldOffsetX = 0.000000
	__LinkShape__000001_Layers.FieldOffsetY = 0.000000
	__LinkShape__000001_Layers.TargetMultiplicity = models.MANY
	__LinkShape__000001_Layers.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Layers.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Layers.SourceMultiplicity = models.MANY
	__LinkShape__000001_Layers.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Layers.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Layers.X = 569.000000
	__LinkShape__000001_Layers.Y = 115.000000
	__LinkShape__000001_Layers.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Layers.StartRatio = 0.500000
	__LinkShape__000001_Layers.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Layers.EndRatio = 0.500000
	__LinkShape__000001_Layers.CornerOffsetRatio = 1.380000

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_SVG)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Layer)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Circle)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_SVG.LinkShapes = append(__GongStructShape__000000_Default_SVG.LinkShapes, __LinkShape__000001_Layers)
	__GongStructShape__000001_Default_Layer.LinkShapes = append(__GongStructShape__000001_Default_Layer.LinkShapes, __LinkShape__000000_Circles)
	// setup of LinkShape instances pointers
}
