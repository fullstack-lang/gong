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

	__AttributeShape__000000_NoName_yet := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_01T08_59_07Z := (&models.DiagramPackage{}).Stage(stage)

	__GongNoteLinkShape__000000_NoName_yet := (&models.GongNoteLinkShape{}).Stage(stage)

	__GongStructShape__000000_Default_Group := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_Button := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_Buttons := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_NoName_yet.Name = `NoName yet`

	__AttributeShape__000000_NoName_yet.IdentifierMeta = ref_models.Button{}.Name
	__AttributeShape__000000_NoName_yet.FieldTypeAsString = ``
	__AttributeShape__000000_NoName_yet.Structname = ``
	__AttributeShape__000000_NoName_yet.Fieldtypename = ``

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = `[false,true]`
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_01T08_59_07Z.Name = `Diagram Package created the 2025-06-01T08:59:07Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_01T08_59_07Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_01T08_59_07Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_01T08_59_07Z.AbsolutePathToDiagramPackage = ``

	__GongNoteLinkShape__000000_NoName_yet.Name = `NoName yet`

	__GongNoteLinkShape__000000_NoName_yet.Type = models.NOTE_SHAPE_LINK_TO_GONG_FIELD

	__GongStructShape__000000_Default_Group.Name = `Default-Group`
	__GongStructShape__000000_Default_Group.X = 17.000000
	__GongStructShape__000000_Default_Group.Y = 52.000000

	__GongStructShape__000000_Default_Group.IdentifierMeta = ref_models.Group{}
	__GongStructShape__000000_Default_Group.ShowNbInstances = false
	__GongStructShape__000000_Default_Group.NbInstances = 0
	__GongStructShape__000000_Default_Group.Width = 240.000000
	__GongStructShape__000000_Default_Group.Height = 63.000000
	__GongStructShape__000000_Default_Group.IsSelected = false

	__GongStructShape__000001_Default_Button.Name = `Default-Button`
	__GongStructShape__000001_Default_Button.X = 554.000000
	__GongStructShape__000001_Default_Button.Y = 128.000000

	__GongStructShape__000001_Default_Button.IdentifierMeta = ref_models.Button{}
	__GongStructShape__000001_Default_Button.ShowNbInstances = false
	__GongStructShape__000001_Default_Button.NbInstances = 0
	__GongStructShape__000001_Default_Button.Width = 240.000000
	__GongStructShape__000001_Default_Button.Height = 63.000000
	__GongStructShape__000001_Default_Button.IsSelected = false

	__LinkShape__000000_Buttons.Name = `Buttons`

	__LinkShape__000000_Buttons.IdentifierMeta = ref_models.Group{}.Buttons

	__LinkShape__000000_Buttons.FieldTypeIdentifierMeta = ref_models.Button{}
	__LinkShape__000000_Buttons.FieldOffsetX = 0.000000
	__LinkShape__000000_Buttons.FieldOffsetY = 0.000000
	__LinkShape__000000_Buttons.TargetMultiplicity = models.MANY
	__LinkShape__000000_Buttons.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Buttons.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Buttons.SourceMultiplicity = models.MANY
	__LinkShape__000000_Buttons.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Buttons.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Buttons.X = 562.000000
	__LinkShape__000000_Buttons.Y = 106.000000
	__LinkShape__000000_Buttons.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Buttons.StartRatio = 0.500000
	__LinkShape__000000_Buttons.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Buttons.EndRatio = 0.500000
	__LinkShape__000000_Buttons.CornerOffsetRatio = 1.380000

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Group)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Button)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_01T08_59_07Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_06_01T08_59_07Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_01T08_59_07Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongNoteLinkShape instances pointers
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Group.LinkShapes = append(__GongStructShape__000000_Default_Group.LinkShapes, __LinkShape__000000_Buttons)
	// setup of LinkShape instances pointers
}
