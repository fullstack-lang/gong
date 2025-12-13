package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/markdown/go/models"
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

	const __write__local_time = "2025-07-07 05:02:30.201795 CEST"
	const __write__utc_time__ = "2025-07-07 03:02:30.201795 UTC"

	const __commitId__ = "0000000010"

	// Declaration of instances to stage

	__AttributeShape__000000_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000001_Content := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_07_07T03_02_10Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_Content := (&models.GongStructShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_Name.Name = `Name`
	__AttributeShape__000000_Name.IdentifierMeta = ref_models.Content{}.Name
	__AttributeShape__000000_Name.FieldTypeAsString = ``
	__AttributeShape__000000_Name.Structname = `Content`
	__AttributeShape__000000_Name.Fieldtypename = `string`

	__AttributeShape__000001_Content.Name = `Content`
	__AttributeShape__000001_Content.IdentifierMeta = ref_models.Content{}.Content
	__AttributeShape__000001_Content.FieldTypeAsString = ``
	__AttributeShape__000001_Content.Structname = `Content`
	__AttributeShape__000001_Content.Fieldtypename = `string`

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

	__DiagramPackage__000000_Diagram_Package_created_the_2025_07_07T03_02_10Z.Name = `Diagram Package created the 2025-07-07T03:02:10Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_07_07T03_02_10Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_07_07T03_02_10Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_07_07T03_02_10Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_Content.Name = `Default-Content`
	__GongStructShape__000000_Default_Content.X = 33.000000
	__GongStructShape__000000_Default_Content.Y = 88.000000
	__GongStructShape__000000_Default_Content.IdentifierMeta = ref_models.Content{}
	__GongStructShape__000000_Default_Content.Width = 240.000000
	__GongStructShape__000000_Default_Content.Height = 103.000000
	__GongStructShape__000000_Default_Content.IsSelected = false

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Content)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_07_07T03_02_10Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_07_07T03_02_10Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_07_07T03_02_10Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Content.AttributeShapes = append(__GongStructShape__000000_Default_Content.AttributeShapes, __AttributeShape__000000_Name)
	__GongStructShape__000000_Default_Content.AttributeShapes = append(__GongStructShape__000000_Default_Content.AttributeShapes, __AttributeShape__000001_Content)
}
