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

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_30_30Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_Astruct := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_Bstruct := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_Associationtob := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansionBinaryEncoding = 1
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansionBinaryEncoding = 0
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansionBinaryEncoding = 0

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_30_30Z.Name = `Diagram Package created the 2025-05-04T22:30:30Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_30_30Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_30_30Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_30_30Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_Astruct.Name = `Default-Astruct`
	__GongStructShape__000000_Default_Astruct.X = 31.000000
	__GongStructShape__000000_Default_Astruct.Y = 96.000000

	//gong:ident [ref_models.Astruct] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Astruct.Identifier = `ref_models.Astruct`
	__GongStructShape__000000_Default_Astruct.ShowNbInstances = false
	__GongStructShape__000000_Default_Astruct.NbInstances = 0
	__GongStructShape__000000_Default_Astruct.Width = 240.000000
	__GongStructShape__000000_Default_Astruct.Height = 63.000000
	__GongStructShape__000000_Default_Astruct.IsSelected = false

	__GongStructShape__000001_Default_Bstruct.Name = `Default-Bstruct`
	__GongStructShape__000001_Default_Bstruct.X = 634.000000
	__GongStructShape__000001_Default_Bstruct.Y = 99.000000

	//gong:ident [ref_models.Bstruct] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_Bstruct.Identifier = `ref_models.Bstruct`
	__GongStructShape__000001_Default_Bstruct.ShowNbInstances = false
	__GongStructShape__000001_Default_Bstruct.NbInstances = 0
	__GongStructShape__000001_Default_Bstruct.Width = 240.000000
	__GongStructShape__000001_Default_Bstruct.Height = 63.000000
	__GongStructShape__000001_Default_Bstruct.IsSelected = false

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
	__LinkShape__000000_Associationtob.X = 428.000000
	__LinkShape__000000_Associationtob.Y = 112.500000
	__LinkShape__000000_Associationtob.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Associationtob.StartRatio = 0.500000
	__LinkShape__000000_Associationtob.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Associationtob.EndRatio = 0.500000
	__LinkShape__000000_Associationtob.CornerOffsetRatio = 1.380000

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Astruct)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Bstruct)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_30_30Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_30_30Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_30_30Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Astruct.LinkShapes = append(__GongStructShape__000000_Default_Astruct.LinkShapes, __LinkShape__000000_Associationtob)
	// setup of LinkShape instances pointers
}
