package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/button/go/models"
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

	__AttributeShape__000000_Name := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_31T10_57_01Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_Button := (&models.GongStructShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_Name.Name = `Name`

	//gong:ident [ref_models.Button.Name] comment added to overcome the problem with the comment map association
	__AttributeShape__000000_Name.Identifier = `ref_models.Button.Name`
	__AttributeShape__000000_Name.IdentifierMeta = ref_models.Button{}.Name
	__AttributeShape__000000_Name.FieldTypeAsString = ``
	__AttributeShape__000000_Name.Structname = `Button`
	__AttributeShape__000000_Name.Fieldtypename = `string`

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = `[true]`
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_31T10_57_01Z.Name = `Diagram Package created the 2025-05-31T10:57:01Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_31T10_57_01Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_31T10_57_01Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_31T10_57_01Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_Button.Name = `Default-Button`
	__GongStructShape__000000_Default_Button.X = 42.000000
	__GongStructShape__000000_Default_Button.Y = 24.000000

	//gong:ident [ref_models.Button] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Button.Identifier = `ref_models.Button`
	__GongStructShape__000000_Default_Button.ShowNbInstances = false
	__GongStructShape__000000_Default_Button.NbInstances = 0
	__GongStructShape__000000_Default_Button.Width = 240.000000
	__GongStructShape__000000_Default_Button.Height = 83.000000
	__GongStructShape__000000_Default_Button.IsSelected = false

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Button)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_31T10_57_01Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_31T10_57_01Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_31T10_57_01Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Button.AttributeShapes = append(__GongStructShape__000000_Default_Button.AttributeShapes, __AttributeShape__000000_Name)
}
