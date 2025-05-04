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

	__AttributeShape__000000_Name := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)
	__Classdiagram__000001_Default_1 := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T09_52_44Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_Astruct := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_AstructBstruct2Use := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_AstructBstructUse := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000003_Default_Bstruct := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000004_Default_1_AstructBstruct2Use := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_Associationtob := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_Name.Name = `Name`

	//gong:ident [ref_models.Astruct.Name] comment added to overcome the problem with the comment map association
	__AttributeShape__000000_Name.Identifier = `ref_models.Astruct.Name`
	__AttributeShape__000000_Name.FieldTypeAsString = ``
	__AttributeShape__000000_Name.Structname = `Astruct`
	__AttributeShape__000000_Name.Fieldtypename = `string`

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = false
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansionBinaryEncoding = 9
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansionBinaryEncoding = 0
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansionBinaryEncoding = 0

	__Classdiagram__000001_Default_1.Name = `Default_1`
	__Classdiagram__000001_Default_1.IsInRenameMode = false
	__Classdiagram__000001_Default_1.IsExpanded = true
	__Classdiagram__000001_Default_1.NodeGongStructsIsExpanded = true
	__Classdiagram__000001_Default_1.NodeGongStructNodeExpansionBinaryEncoding = 0
	__Classdiagram__000001_Default_1.NodeGongEnumsIsExpanded = false
	__Classdiagram__000001_Default_1.NodeGongEnumNodeExpansionBinaryEncoding = 0
	__Classdiagram__000001_Default_1.NodeGongNotesIsExpanded = false
	__Classdiagram__000001_Default_1.NodeGongNoteNodeExpansionBinaryEncoding = 0

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T09_52_44Z.Name = `Diagram Package created the 2025-05-04T09:52:44Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T09_52_44Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T09_52_44Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T09_52_44Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_Astruct.Name = `Default-Astruct`
	__GongStructShape__000000_Default_Astruct.X = 76.000000
	__GongStructShape__000000_Default_Astruct.Y = 36.000000

	//gong:ident [ref_models.Astruct] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Astruct.Identifier = `ref_models.Astruct`
	__GongStructShape__000000_Default_Astruct.ShowNbInstances = false
	__GongStructShape__000000_Default_Astruct.NbInstances = 0
	__GongStructShape__000000_Default_Astruct.Width = 240.000000
	__GongStructShape__000000_Default_Astruct.Height = 83.000000
	__GongStructShape__000000_Default_Astruct.IsSelected = false

	__GongStructShape__000001_Default_AstructBstruct2Use.Name = `Default-AstructBstruct2Use`
	__GongStructShape__000001_Default_AstructBstruct2Use.X = 55.000000
	__GongStructShape__000001_Default_AstructBstruct2Use.Y = 292.000000

	//gong:ident [ref_models.AstructBstruct2Use] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_AstructBstruct2Use.Identifier = `ref_models.AstructBstruct2Use`
	__GongStructShape__000001_Default_AstructBstruct2Use.ShowNbInstances = false
	__GongStructShape__000001_Default_AstructBstruct2Use.NbInstances = 0
	__GongStructShape__000001_Default_AstructBstruct2Use.Width = 240.000000
	__GongStructShape__000001_Default_AstructBstruct2Use.Height = 63.000000
	__GongStructShape__000001_Default_AstructBstruct2Use.IsSelected = false

	__GongStructShape__000002_Default_AstructBstructUse.Name = `Default-AstructBstructUse`
	__GongStructShape__000002_Default_AstructBstructUse.X = 247.000000
	__GongStructShape__000002_Default_AstructBstructUse.Y = 227.000000

	//gong:ident [ref_models.AstructBstructUse] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_AstructBstructUse.Identifier = `ref_models.AstructBstructUse`
	__GongStructShape__000002_Default_AstructBstructUse.ShowNbInstances = false
	__GongStructShape__000002_Default_AstructBstructUse.NbInstances = 0
	__GongStructShape__000002_Default_AstructBstructUse.Width = 240.000000
	__GongStructShape__000002_Default_AstructBstructUse.Height = 63.000000
	__GongStructShape__000002_Default_AstructBstructUse.IsSelected = false

	__GongStructShape__000003_Default_Bstruct.Name = `Default-Bstruct`
	__GongStructShape__000003_Default_Bstruct.X = 736.000000
	__GongStructShape__000003_Default_Bstruct.Y = 106.000000

	//gong:ident [ref_models.Bstruct] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_Default_Bstruct.Identifier = `ref_models.Bstruct`
	__GongStructShape__000003_Default_Bstruct.ShowNbInstances = false
	__GongStructShape__000003_Default_Bstruct.NbInstances = 0
	__GongStructShape__000003_Default_Bstruct.Width = 240.000000
	__GongStructShape__000003_Default_Bstruct.Height = 63.000000
	__GongStructShape__000003_Default_Bstruct.IsSelected = false

	__GongStructShape__000004_Default_1_AstructBstruct2Use.Name = `Default_1-AstructBstruct2Use`
	__GongStructShape__000004_Default_1_AstructBstruct2Use.X = 34.000000
	__GongStructShape__000004_Default_1_AstructBstruct2Use.Y = 30.000000

	//gong:ident [ref_models.AstructBstruct2Use] comment added to overcome the problem with the comment map association
	__GongStructShape__000004_Default_1_AstructBstruct2Use.Identifier = `ref_models.AstructBstruct2Use`
	__GongStructShape__000004_Default_1_AstructBstruct2Use.ShowNbInstances = false
	__GongStructShape__000004_Default_1_AstructBstruct2Use.NbInstances = 0
	__GongStructShape__000004_Default_1_AstructBstruct2Use.Width = 240.000000
	__GongStructShape__000004_Default_1_AstructBstruct2Use.Height = 63.000000
	__GongStructShape__000004_Default_1_AstructBstruct2Use.IsSelected = false

	__LinkShape__000000_Associationtob.Name = `Associationtob`

	//gong:ident [ref_models.Astruct.Associationtob] comment added to overcome the problem with the comment map association
	__LinkShape__000000_Associationtob.Identifier = `ref_models.Astruct.Associationtob`

	//gong:ident [ref_models.Bstruct] comment added to overcome the problem with the comment map association
	__LinkShape__000000_Associationtob.Fieldtypename = `ref_models.Bstruct`
	__LinkShape__000000_Associationtob.FieldOffsetX = 0.000000
	__LinkShape__000000_Associationtob.FieldOffsetY = 0.000000
	__LinkShape__000000_Associationtob.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000000_Associationtob.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Associationtob.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Associationtob.SourceMultiplicity = models.MANY
	__LinkShape__000000_Associationtob.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Associationtob.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Associationtob.X = 717.500000
	__LinkShape__000000_Associationtob.Y = 134.500000
	__LinkShape__000000_Associationtob.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Associationtob.StartRatio = 0.500000
	__LinkShape__000000_Associationtob.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Associationtob.EndRatio = 0.500000
	__LinkShape__000000_Associationtob.CornerOffsetRatio = 1.380000

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Astruct)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_AstructBstruct2Use)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_AstructBstructUse)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_Bstruct)
	__Classdiagram__000001_Default_1.GongStructShapes = append(__Classdiagram__000001_Default_1.GongStructShapes, __GongStructShape__000004_Default_1_AstructBstruct2Use)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T09_52_44Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T09_52_44Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T09_52_44Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T09_52_44Z.Classdiagrams, __Classdiagram__000001_Default_1)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T09_52_44Z.SelectedClassdiagram = __Classdiagram__000001_Default_1
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Astruct.AttributeShapes = append(__GongStructShape__000000_Default_Astruct.AttributeShapes, __AttributeShape__000000_Name)
	__GongStructShape__000000_Default_Astruct.LinkShapes = append(__GongStructShape__000000_Default_Astruct.LinkShapes, __LinkShape__000000_Associationtob)
	// setup of LinkShape instances pointers
}
