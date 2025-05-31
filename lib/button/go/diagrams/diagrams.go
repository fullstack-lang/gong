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

	__GongEnumShape__000000_Default_StacksNames := (&models.GongEnumShape{}).Stage(stage)

	__GongEnumValueShape__000000_ButtonStackName := (&models.GongEnumValueShape{}).Stage(stage)

	__GongStructShape__000000_Default_Button := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_Group := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_Buttons := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_Name.Name = `Name`

	//gong:ident [ref_models.Button.Name] comment added to overcome the problem with the comment map association
	__AttributeShape__000000_Name.Identifier = `ref_models.Button.Name`
	__AttributeShape__000000_Name.FieldTypeAsString = ``
	__AttributeShape__000000_Name.Structname = `Button`
	__AttributeShape__000000_Name.Fieldtypename = `string`

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = `[false,true]`
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = `[true]`
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_31T10_57_01Z.Name = `Diagram Package created the 2025-05-31T10:57:01Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_31T10_57_01Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_31T10_57_01Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_31T10_57_01Z.AbsolutePathToDiagramPackage = ``

	__GongEnumShape__000000_Default_StacksNames.Name = `Default-StacksNames`
	__GongEnumShape__000000_Default_StacksNames.X = 834.000000
	__GongEnumShape__000000_Default_StacksNames.Y = 88.000000

	//gong:ident [ref_models.StacksNames] comment added to overcome the problem with the comment map association
	__GongEnumShape__000000_Default_StacksNames.Identifier = `ref_models.StacksNames`
	__GongEnumShape__000000_Default_StacksNames.Width = 240.000000
	__GongEnumShape__000000_Default_StacksNames.Height = 83.000000
	__GongEnumShape__000000_Default_StacksNames.IsExpanded = false

	__GongEnumValueShape__000000_ButtonStackName.Name = `ButtonStackName`

	//gong:ident [ref_models.StacksNames.ButtonStackName] comment added to overcome the problem with the comment map association
	__GongEnumValueShape__000000_ButtonStackName.Identifier = `ref_models.StacksNames.ButtonStackName`

	__GongStructShape__000000_Default_Button.Name = `Default-Button`
	__GongStructShape__000000_Default_Button.X = 513.000000
	__GongStructShape__000000_Default_Button.Y = 45.000000

	//gong:ident [ref_models.Button] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Button.Identifier = `ref_models.Button`
	__GongStructShape__000000_Default_Button.ShowNbInstances = false
	__GongStructShape__000000_Default_Button.NbInstances = 0
	__GongStructShape__000000_Default_Button.Width = 240.000000
	__GongStructShape__000000_Default_Button.Height = 83.000000
	__GongStructShape__000000_Default_Button.IsSelected = false

	__GongStructShape__000001_Default_Group.Name = `Default-Group`
	__GongStructShape__000001_Default_Group.X = 108.000000
	__GongStructShape__000001_Default_Group.Y = 63.000000

	//gong:ident [ref_models.Group] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_Group.Identifier = `ref_models.Group`
	__GongStructShape__000001_Default_Group.IdentifierMeta = ref_models.Group{}
	__GongStructShape__000001_Default_Group.ShowNbInstances = false
	__GongStructShape__000001_Default_Group.NbInstances = 0
	__GongStructShape__000001_Default_Group.Width = 240.000000
	__GongStructShape__000001_Default_Group.Height = 63.000000
	__GongStructShape__000001_Default_Group.IsSelected = false

	__LinkShape__000000_Buttons.Name = `Buttons`

	//gong:ident [ref_models.Group.Buttons] comment added to overcome the problem with the comment map association
	__LinkShape__000000_Buttons.Identifier = `ref_models.Group.Buttons`
	__LinkShape__000000_Buttons.IdentifierMeta = ref_models.Group{}.Buttons

	//gong:ident [ref_models.Button] comment added to overcome the problem with the comment map association
	__LinkShape__000000_Buttons.Fieldtypename = `ref_models.Button`
	__LinkShape__000000_Buttons.FieldOffsetX = 0.000000
	__LinkShape__000000_Buttons.FieldOffsetY = 0.000000
	__LinkShape__000000_Buttons.TargetMultiplicity = models.MANY
	__LinkShape__000000_Buttons.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Buttons.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Buttons.SourceMultiplicity = models.MANY
	__LinkShape__000000_Buttons.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Buttons.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Buttons.X = 670.500000
	__LinkShape__000000_Buttons.Y = 85.500000
	__LinkShape__000000_Buttons.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Buttons.StartRatio = 0.500000
	__LinkShape__000000_Buttons.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Buttons.EndRatio = 0.500000
	__LinkShape__000000_Buttons.CornerOffsetRatio = 1.380000

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Button)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Group)
	__Classdiagram__000000_Default.GongEnumShapes = append(__Classdiagram__000000_Default.GongEnumShapes, __GongEnumShape__000000_Default_StacksNames)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_31T10_57_01Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_31T10_57_01Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_31T10_57_01Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongEnumShape instances pointers
	__GongEnumShape__000000_Default_StacksNames.GongEnumValueShapes = append(__GongEnumShape__000000_Default_StacksNames.GongEnumValueShapes, __GongEnumValueShape__000000_ButtonStackName)
	// setup of GongEnumValueShape instances pointers
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Button.AttributeShapes = append(__GongStructShape__000000_Default_Button.AttributeShapes, __AttributeShape__000000_Name)
	__GongStructShape__000001_Default_Group.LinkShapes = append(__GongStructShape__000001_Default_Group.LinkShapes, __LinkShape__000000_Buttons)
	// setup of LinkShape instances pointers
}
