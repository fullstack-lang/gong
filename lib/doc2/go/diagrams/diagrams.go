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

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_AttributeShape := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_DiagramPackage := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_GongStructShape := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000003_Default_Classdiagram := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_Classdiagrams := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000001_GongStructShapes := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000002_AttributeShapes := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansionBinaryEncoding = 32
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansionBinaryEncoding = 0
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansionBinaryEncoding = 0

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.Name = `Diagram Package created the 2025-05-04T22:53:27Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_AttributeShape.Name = `Default-AttributeShape`
	__GongStructShape__000000_Default_AttributeShape.X = 1444.000000
	__GongStructShape__000000_Default_AttributeShape.Y = 26.000000

	//gong:ident [ref_models.AttributeShape] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_AttributeShape.Identifier = `ref_models.AttributeShape`
	__GongStructShape__000000_Default_AttributeShape.ShowNbInstances = false
	__GongStructShape__000000_Default_AttributeShape.NbInstances = 0
	__GongStructShape__000000_Default_AttributeShape.Width = 240.000000
	__GongStructShape__000000_Default_AttributeShape.Height = 63.000000
	__GongStructShape__000000_Default_AttributeShape.IsSelected = false

	__GongStructShape__000001_Default_DiagramPackage.Name = `Default-DiagramPackage`
	__GongStructShape__000001_Default_DiagramPackage.X = 58.000000
	__GongStructShape__000001_Default_DiagramPackage.Y = 36.000000

	//gong:ident [ref_models.DiagramPackage] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_DiagramPackage.Identifier = `ref_models.DiagramPackage`
	__GongStructShape__000001_Default_DiagramPackage.ShowNbInstances = false
	__GongStructShape__000001_Default_DiagramPackage.NbInstances = 0
	__GongStructShape__000001_Default_DiagramPackage.Width = 240.000000
	__GongStructShape__000001_Default_DiagramPackage.Height = 63.000000
	__GongStructShape__000001_Default_DiagramPackage.IsSelected = false

	__GongStructShape__000002_Default_GongStructShape.Name = `Default-GongStructShape`
	__GongStructShape__000002_Default_GongStructShape.X = 985.000000
	__GongStructShape__000002_Default_GongStructShape.Y = 32.000000

	//gong:ident [ref_models.GongStructShape] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_GongStructShape.Identifier = `ref_models.GongStructShape`
	__GongStructShape__000002_Default_GongStructShape.ShowNbInstances = false
	__GongStructShape__000002_Default_GongStructShape.NbInstances = 0
	__GongStructShape__000002_Default_GongStructShape.Width = 240.000000
	__GongStructShape__000002_Default_GongStructShape.Height = 63.000000
	__GongStructShape__000002_Default_GongStructShape.IsSelected = false

	__GongStructShape__000003_Default_Classdiagram.Name = `Default-Classdiagram`
	__GongStructShape__000003_Default_Classdiagram.X = 517.000000
	__GongStructShape__000003_Default_Classdiagram.Y = 35.000000

	//gong:ident [ref_models.Classdiagram] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_Default_Classdiagram.Identifier = `ref_models.Classdiagram`
	__GongStructShape__000003_Default_Classdiagram.ShowNbInstances = false
	__GongStructShape__000003_Default_Classdiagram.NbInstances = 0
	__GongStructShape__000003_Default_Classdiagram.Width = 240.000000
	__GongStructShape__000003_Default_Classdiagram.Height = 63.000000
	__GongStructShape__000003_Default_Classdiagram.IsSelected = false

	__LinkShape__000000_Classdiagrams.Name = `Classdiagrams`

	//gong:ident [ref_models.DiagramPackage.Classdiagrams] comment added to overcome the problem with the comment map association
	__LinkShape__000000_Classdiagrams.Identifier = `ref_models.DiagramPackage.Classdiagrams`

	//gong:ident [ref_models.Classdiagram] comment added to overcome the problem with the comment map association
	__LinkShape__000000_Classdiagrams.Fieldtypename = `ref_models.Classdiagram`
	__LinkShape__000000_Classdiagrams.FieldOffsetX = 0.000000
	__LinkShape__000000_Classdiagrams.FieldOffsetY = 0.000000
	__LinkShape__000000_Classdiagrams.TargetMultiplicity = models.MANY
	__LinkShape__000000_Classdiagrams.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Classdiagrams.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Classdiagrams.SourceMultiplicity = models.MANY
	__LinkShape__000000_Classdiagrams.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Classdiagrams.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Classdiagrams.X = 441.500000
	__LinkShape__000000_Classdiagrams.Y = 162.000000
	__LinkShape__000000_Classdiagrams.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Classdiagrams.StartRatio = 0.500000
	__LinkShape__000000_Classdiagrams.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Classdiagrams.EndRatio = 0.500000
	__LinkShape__000000_Classdiagrams.CornerOffsetRatio = 1.380000

	__LinkShape__000001_GongStructShapes.Name = `GongStructShapes`

	//gong:ident [ref_models.Classdiagram.GongStructShapes] comment added to overcome the problem with the comment map association
	__LinkShape__000001_GongStructShapes.Identifier = `ref_models.Classdiagram.GongStructShapes`

	//gong:ident [ref_models.GongStructShape] comment added to overcome the problem with the comment map association
	__LinkShape__000001_GongStructShapes.Fieldtypename = `ref_models.GongStructShape`
	__LinkShape__000001_GongStructShapes.FieldOffsetX = 0.000000
	__LinkShape__000001_GongStructShapes.FieldOffsetY = 0.000000
	__LinkShape__000001_GongStructShapes.TargetMultiplicity = models.MANY
	__LinkShape__000001_GongStructShapes.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000001_GongStructShapes.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000001_GongStructShapes.SourceMultiplicity = models.MANY
	__LinkShape__000001_GongStructShapes.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000001_GongStructShapes.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000001_GongStructShapes.X = 1140.500000
	__LinkShape__000001_GongStructShapes.Y = 197.000000
	__LinkShape__000001_GongStructShapes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_GongStructShapes.StartRatio = 0.500000
	__LinkShape__000001_GongStructShapes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_GongStructShapes.EndRatio = 0.500000
	__LinkShape__000001_GongStructShapes.CornerOffsetRatio = 1.380000

	__LinkShape__000002_AttributeShapes.Name = `AttributeShapes`

	//gong:ident [ref_models.GongStructShape.AttributeShapes] comment added to overcome the problem with the comment map association
	__LinkShape__000002_AttributeShapes.Identifier = `ref_models.GongStructShape.AttributeShapes`

	//gong:ident [ref_models.AttributeShape] comment added to overcome the problem with the comment map association
	__LinkShape__000002_AttributeShapes.Fieldtypename = `ref_models.AttributeShape`
	__LinkShape__000002_AttributeShapes.FieldOffsetX = 0.000000
	__LinkShape__000002_AttributeShapes.FieldOffsetY = 0.000000
	__LinkShape__000002_AttributeShapes.TargetMultiplicity = models.MANY
	__LinkShape__000002_AttributeShapes.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000002_AttributeShapes.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000002_AttributeShapes.SourceMultiplicity = models.MANY
	__LinkShape__000002_AttributeShapes.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000002_AttributeShapes.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000002_AttributeShapes.X = 1575.500000
	__LinkShape__000002_AttributeShapes.Y = 191.500000
	__LinkShape__000002_AttributeShapes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_AttributeShapes.StartRatio = 0.500000
	__LinkShape__000002_AttributeShapes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_AttributeShapes.EndRatio = 0.500000
	__LinkShape__000002_AttributeShapes.CornerOffsetRatio = 1.380000

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_AttributeShape)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_DiagramPackage)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_GongStructShape)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_Classdiagram)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__000001_Default_DiagramPackage.LinkShapes = append(__GongStructShape__000001_Default_DiagramPackage.LinkShapes, __LinkShape__000000_Classdiagrams)
	__GongStructShape__000002_Default_GongStructShape.LinkShapes = append(__GongStructShape__000002_Default_GongStructShape.LinkShapes, __LinkShape__000002_AttributeShapes)
	__GongStructShape__000003_Default_Classdiagram.LinkShapes = append(__GongStructShape__000003_Default_Classdiagram.LinkShapes, __LinkShape__000001_GongStructShapes)
	// setup of LinkShape instances pointers
}
