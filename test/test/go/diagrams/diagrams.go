package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/test/test/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration
var _ ref_models.Stage

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T02_22_27Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_Gstruct := (&models.GongStructShape{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = ``
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T02_22_27Z.Name = `Diagram Package created the 2025-06-03T02:22:27Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T02_22_27Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T02_22_27Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T02_22_27Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_Gstruct.Name = `Default-Gstruct`
	__GongStructShape__000000_Default_Gstruct.X = 47.000000
	__GongStructShape__000000_Default_Gstruct.Y = 20.000000

	//gong:ident [] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Gstruct.Identifier = ``
	__GongStructShape__000000_Default_Gstruct.IdentifierMeta = ref_models.Gstruct{}
	__GongStructShape__000000_Default_Gstruct.ShowNbInstances = false
	__GongStructShape__000000_Default_Gstruct.NbInstances = 0
	__GongStructShape__000000_Default_Gstruct.Width = 240.000000
	__GongStructShape__000000_Default_Gstruct.Height = 63.000000
	__GongStructShape__000000_Default_Gstruct.IsSelected = false

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Gstruct)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T02_22_27Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T02_22_27Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T02_22_27Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
}
