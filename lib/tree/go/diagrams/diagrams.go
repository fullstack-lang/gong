package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/tree/go/models"
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
	__AttributeShape__000001_FontStyle := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000002_BackgroundColor := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000003_IsExpanded := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000004_HasCheckboxButton := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000005_IsChecked := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000006_IsCheckboxDisabled := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000007_CheckboxHasToolTip := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000008_CheckboxToolTipText := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000009_CheckboxToolTipPosition := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000010_HasSecondCheckboxButton := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000011_IsSecondCheckboxChecked := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000012_IsSecondCheckboxDisabled := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000013_TextAfterSecondCheckbox := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000014_IsInEditMode := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000015_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000016_Icon := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000017_IsDisabled := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000018_HasToolTip := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000019_ToolTipText := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000020_ToolTipPosition := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_11T03_39_06Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_Tree := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_SVGIcon := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_Node := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000003_Default_Button := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_RootNodes := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000001_Buttons := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000002_Children := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000003_SVGIcon := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_Name.Name = `Name`
	__AttributeShape__000000_Name.IdentifierMeta = ref_models.Node{}.Name
	__AttributeShape__000000_Name.FieldTypeAsString = ``
	__AttributeShape__000000_Name.Structname = `Node`
	__AttributeShape__000000_Name.Fieldtypename = `string`

	__AttributeShape__000001_FontStyle.Name = `FontStyle`
	__AttributeShape__000001_FontStyle.IdentifierMeta = ref_models.Node{}.FontStyle
	__AttributeShape__000001_FontStyle.FieldTypeAsString = ``
	__AttributeShape__000001_FontStyle.Structname = `Node`
	__AttributeShape__000001_FontStyle.Fieldtypename = `FontStyleEnum`

	__AttributeShape__000002_BackgroundColor.Name = `BackgroundColor`
	__AttributeShape__000002_BackgroundColor.IdentifierMeta = ref_models.Node{}.BackgroundColor
	__AttributeShape__000002_BackgroundColor.FieldTypeAsString = ``
	__AttributeShape__000002_BackgroundColor.Structname = `Node`
	__AttributeShape__000002_BackgroundColor.Fieldtypename = `string`

	__AttributeShape__000003_IsExpanded.Name = `IsExpanded`
	__AttributeShape__000003_IsExpanded.IdentifierMeta = ref_models.Node{}.IsExpanded
	__AttributeShape__000003_IsExpanded.FieldTypeAsString = ``
	__AttributeShape__000003_IsExpanded.Structname = `Node`
	__AttributeShape__000003_IsExpanded.Fieldtypename = `bool`

	__AttributeShape__000004_HasCheckboxButton.Name = `HasCheckboxButton`
	__AttributeShape__000004_HasCheckboxButton.IdentifierMeta = ref_models.Node{}.HasCheckboxButton
	__AttributeShape__000004_HasCheckboxButton.FieldTypeAsString = ``
	__AttributeShape__000004_HasCheckboxButton.Structname = `Node`
	__AttributeShape__000004_HasCheckboxButton.Fieldtypename = `bool`

	__AttributeShape__000005_IsChecked.Name = `IsChecked`
	__AttributeShape__000005_IsChecked.IdentifierMeta = ref_models.Node{}.IsChecked
	__AttributeShape__000005_IsChecked.FieldTypeAsString = ``
	__AttributeShape__000005_IsChecked.Structname = `Node`
	__AttributeShape__000005_IsChecked.Fieldtypename = `bool`

	__AttributeShape__000006_IsCheckboxDisabled.Name = `IsCheckboxDisabled`
	__AttributeShape__000006_IsCheckboxDisabled.IdentifierMeta = ref_models.Node{}.IsCheckboxDisabled
	__AttributeShape__000006_IsCheckboxDisabled.FieldTypeAsString = ``
	__AttributeShape__000006_IsCheckboxDisabled.Structname = `Node`
	__AttributeShape__000006_IsCheckboxDisabled.Fieldtypename = `bool`

	__AttributeShape__000007_CheckboxHasToolTip.Name = `CheckboxHasToolTip`
	__AttributeShape__000007_CheckboxHasToolTip.IdentifierMeta = ref_models.Node{}.CheckboxHasToolTip
	__AttributeShape__000007_CheckboxHasToolTip.FieldTypeAsString = ``
	__AttributeShape__000007_CheckboxHasToolTip.Structname = `Node`
	__AttributeShape__000007_CheckboxHasToolTip.Fieldtypename = `bool`

	__AttributeShape__000008_CheckboxToolTipText.Name = `CheckboxToolTipText`
	__AttributeShape__000008_CheckboxToolTipText.IdentifierMeta = ref_models.Node{}.CheckboxToolTipText
	__AttributeShape__000008_CheckboxToolTipText.FieldTypeAsString = ``
	__AttributeShape__000008_CheckboxToolTipText.Structname = `Node`
	__AttributeShape__000008_CheckboxToolTipText.Fieldtypename = `string`

	__AttributeShape__000009_CheckboxToolTipPosition.Name = `CheckboxToolTipPosition`
	__AttributeShape__000009_CheckboxToolTipPosition.IdentifierMeta = ref_models.Node{}.CheckboxToolTipPosition
	__AttributeShape__000009_CheckboxToolTipPosition.FieldTypeAsString = ``
	__AttributeShape__000009_CheckboxToolTipPosition.Structname = `Node`
	__AttributeShape__000009_CheckboxToolTipPosition.Fieldtypename = `ToolTipPositionEnum`

	__AttributeShape__000010_HasSecondCheckboxButton.Name = `HasSecondCheckboxButton`
	__AttributeShape__000010_HasSecondCheckboxButton.IdentifierMeta = ref_models.Node{}.HasSecondCheckboxButton
	__AttributeShape__000010_HasSecondCheckboxButton.FieldTypeAsString = ``
	__AttributeShape__000010_HasSecondCheckboxButton.Structname = `Node`
	__AttributeShape__000010_HasSecondCheckboxButton.Fieldtypename = `bool`

	__AttributeShape__000011_IsSecondCheckboxChecked.Name = `IsSecondCheckboxChecked`
	__AttributeShape__000011_IsSecondCheckboxChecked.IdentifierMeta = ref_models.Node{}.IsSecondCheckboxChecked
	__AttributeShape__000011_IsSecondCheckboxChecked.FieldTypeAsString = ``
	__AttributeShape__000011_IsSecondCheckboxChecked.Structname = `Node`
	__AttributeShape__000011_IsSecondCheckboxChecked.Fieldtypename = `bool`

	__AttributeShape__000012_IsSecondCheckboxDisabled.Name = `IsSecondCheckboxDisabled`
	__AttributeShape__000012_IsSecondCheckboxDisabled.IdentifierMeta = ref_models.Node{}.IsSecondCheckboxDisabled
	__AttributeShape__000012_IsSecondCheckboxDisabled.FieldTypeAsString = ``
	__AttributeShape__000012_IsSecondCheckboxDisabled.Structname = `Node`
	__AttributeShape__000012_IsSecondCheckboxDisabled.Fieldtypename = `bool`

	__AttributeShape__000013_TextAfterSecondCheckbox.Name = `TextAfterSecondCheckbox`
	__AttributeShape__000013_TextAfterSecondCheckbox.IdentifierMeta = ref_models.Node{}.TextAfterSecondCheckbox
	__AttributeShape__000013_TextAfterSecondCheckbox.FieldTypeAsString = ``
	__AttributeShape__000013_TextAfterSecondCheckbox.Structname = `Node`
	__AttributeShape__000013_TextAfterSecondCheckbox.Fieldtypename = `string`

	__AttributeShape__000014_IsInEditMode.Name = `IsInEditMode`
	__AttributeShape__000014_IsInEditMode.IdentifierMeta = ref_models.Node{}.IsInEditMode
	__AttributeShape__000014_IsInEditMode.FieldTypeAsString = ``
	__AttributeShape__000014_IsInEditMode.Structname = `Node`
	__AttributeShape__000014_IsInEditMode.Fieldtypename = `bool`

	__AttributeShape__000015_Name.Name = `Name`
	__AttributeShape__000015_Name.IdentifierMeta = ref_models.Button{}.Name
	__AttributeShape__000015_Name.FieldTypeAsString = ``
	__AttributeShape__000015_Name.Structname = `Button`
	__AttributeShape__000015_Name.Fieldtypename = `string`

	__AttributeShape__000016_Icon.Name = `Icon`
	__AttributeShape__000016_Icon.IdentifierMeta = ref_models.Button{}.Icon
	__AttributeShape__000016_Icon.FieldTypeAsString = ``
	__AttributeShape__000016_Icon.Structname = `Button`
	__AttributeShape__000016_Icon.Fieldtypename = `string`

	__AttributeShape__000017_IsDisabled.Name = `IsDisabled`
	__AttributeShape__000017_IsDisabled.IdentifierMeta = ref_models.Button{}.IsDisabled
	__AttributeShape__000017_IsDisabled.FieldTypeAsString = ``
	__AttributeShape__000017_IsDisabled.Structname = `Button`
	__AttributeShape__000017_IsDisabled.Fieldtypename = `bool`

	__AttributeShape__000018_HasToolTip.Name = `HasToolTip`
	__AttributeShape__000018_HasToolTip.IdentifierMeta = ref_models.Button{}.HasToolTip
	__AttributeShape__000018_HasToolTip.FieldTypeAsString = ``
	__AttributeShape__000018_HasToolTip.Structname = `Button`
	__AttributeShape__000018_HasToolTip.Fieldtypename = `bool`

	__AttributeShape__000019_ToolTipText.Name = `ToolTipText`
	__AttributeShape__000019_ToolTipText.IdentifierMeta = ref_models.Button{}.ToolTipText
	__AttributeShape__000019_ToolTipText.FieldTypeAsString = ``
	__AttributeShape__000019_ToolTipText.Structname = `Button`
	__AttributeShape__000019_ToolTipText.Fieldtypename = `string`

	__AttributeShape__000020_ToolTipPosition.Name = `ToolTipPosition`
	__AttributeShape__000020_ToolTipPosition.IdentifierMeta = ref_models.Button{}.ToolTipPosition
	__AttributeShape__000020_ToolTipPosition.FieldTypeAsString = ``
	__AttributeShape__000020_ToolTipPosition.Structname = `Button`
	__AttributeShape__000020_ToolTipPosition.Fieldtypename = `ToolTipPositionEnum`

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = `[true,true,false,true]`
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_11T03_39_06Z.Name = `Diagram Package created the 2025-06-11T03:39:06Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_11T03_39_06Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_11T03_39_06Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_11T03_39_06Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_Tree.Name = `Default-Tree`
	__GongStructShape__000000_Default_Tree.X = 7.000000
	__GongStructShape__000000_Default_Tree.Y = 196.000000
	__GongStructShape__000000_Default_Tree.IdentifierMeta = ref_models.Tree{}

	__GongStructShape__000000_Default_Tree.Width = 240.000000
	__GongStructShape__000000_Default_Tree.Height = 63.000000
	__GongStructShape__000000_Default_Tree.IsSelected = false

	__GongStructShape__000001_Default_SVGIcon.Name = `Default-SVGIcon`
	__GongStructShape__000001_Default_SVGIcon.X = 1014.000000
	__GongStructShape__000001_Default_SVGIcon.Y = 331.000000
	__GongStructShape__000001_Default_SVGIcon.IdentifierMeta = ref_models.SVGIcon{}

	__GongStructShape__000001_Default_SVGIcon.Width = 240.000000
	__GongStructShape__000001_Default_SVGIcon.Height = 63.000000
	__GongStructShape__000001_Default_SVGIcon.IsSelected = false

	__GongStructShape__000002_Default_Node.Name = `Default-Node`
	__GongStructShape__000002_Default_Node.X = 426.000000
	__GongStructShape__000002_Default_Node.Y = 48.000000
	__GongStructShape__000002_Default_Node.IdentifierMeta = ref_models.Node{}

	__GongStructShape__000002_Default_Node.Width = 377.000000
	__GongStructShape__000002_Default_Node.Height = 363.000000
	__GongStructShape__000002_Default_Node.IsSelected = false

	__GongStructShape__000003_Default_Button.Name = `Default-Button`
	__GongStructShape__000003_Default_Button.X = 985.000000
	__GongStructShape__000003_Default_Button.Y = 43.000000
	__GongStructShape__000003_Default_Button.IdentifierMeta = ref_models.Button{}

	__GongStructShape__000003_Default_Button.Width = 316.000000
	__GongStructShape__000003_Default_Button.Height = 183.000000
	__GongStructShape__000003_Default_Button.IsSelected = false

	__LinkShape__000000_RootNodes.Name = `RootNodes`
	__LinkShape__000000_RootNodes.IdentifierMeta = ref_models.Tree{}.RootNodes
	__LinkShape__000000_RootNodes.FieldTypeIdentifierMeta = ref_models.Node{}
	__LinkShape__000000_RootNodes.FieldOffsetX = 0.000000
	__LinkShape__000000_RootNodes.FieldOffsetY = 0.000000
	__LinkShape__000000_RootNodes.TargetMultiplicity = models.MANY
	__LinkShape__000000_RootNodes.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_RootNodes.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_RootNodes.SourceMultiplicity = models.MANY
	__LinkShape__000000_RootNodes.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_RootNodes.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_RootNodes.X = 421.000000
	__LinkShape__000000_RootNodes.Y = 87.500000
	__LinkShape__000000_RootNodes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_RootNodes.StartRatio = 0.500000
	__LinkShape__000000_RootNodes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_RootNodes.EndRatio = 0.500000
	__LinkShape__000000_RootNodes.CornerOffsetRatio = 1.380000

	__LinkShape__000001_Buttons.Name = `Buttons`
	__LinkShape__000001_Buttons.IdentifierMeta = ref_models.Node{}.Buttons
	__LinkShape__000001_Buttons.FieldTypeIdentifierMeta = ref_models.Button{}
	__LinkShape__000001_Buttons.FieldOffsetX = 0.000000
	__LinkShape__000001_Buttons.FieldOffsetY = 0.000000
	__LinkShape__000001_Buttons.TargetMultiplicity = models.MANY
	__LinkShape__000001_Buttons.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Buttons.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Buttons.SourceMultiplicity = models.MANY
	__LinkShape__000001_Buttons.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Buttons.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Buttons.X = 1198.000000
	__LinkShape__000001_Buttons.Y = 228.500000
	__LinkShape__000001_Buttons.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Buttons.StartRatio = 0.243519
	__LinkShape__000001_Buttons.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Buttons.EndRatio = 0.500000
	__LinkShape__000001_Buttons.CornerOffsetRatio = 1.380000

	__LinkShape__000002_Children.Name = `Children`
	__LinkShape__000002_Children.IdentifierMeta = ref_models.Node{}.Children
	__LinkShape__000002_Children.FieldTypeIdentifierMeta = ref_models.Node{}
	__LinkShape__000002_Children.FieldOffsetX = 0.000000
	__LinkShape__000002_Children.FieldOffsetY = 0.000000
	__LinkShape__000002_Children.TargetMultiplicity = models.MANY
	__LinkShape__000002_Children.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000002_Children.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000002_Children.SourceMultiplicity = models.MANY
	__LinkShape__000002_Children.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000002_Children.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000002_Children.X = 991.500000
	__LinkShape__000002_Children.Y = 229.500000
	__LinkShape__000002_Children.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_Children.StartRatio = 0.764181
	__LinkShape__000002_Children.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_Children.EndRatio = 0.965282
	__LinkShape__000002_Children.CornerOffsetRatio = 1.380000

	__LinkShape__000003_SVGIcon.Name = `SVGIcon`
	__LinkShape__000003_SVGIcon.IdentifierMeta = ref_models.Button{}.SVGIcon
	__LinkShape__000003_SVGIcon.FieldTypeIdentifierMeta = ref_models.SVGIcon{}
	__LinkShape__000003_SVGIcon.FieldOffsetX = 0.000000
	__LinkShape__000003_SVGIcon.FieldOffsetY = 0.000000
	__LinkShape__000003_SVGIcon.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000003_SVGIcon.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000003_SVGIcon.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000003_SVGIcon.SourceMultiplicity = models.MANY
	__LinkShape__000003_SVGIcon.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000003_SVGIcon.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000003_SVGIcon.X = 1433.500000
	__LinkShape__000003_SVGIcon.Y = 323.500000
	__LinkShape__000003_SVGIcon.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000003_SVGIcon.StartRatio = 0.500000
	__LinkShape__000003_SVGIcon.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000003_SVGIcon.EndRatio = 0.500000
	__LinkShape__000003_SVGIcon.CornerOffsetRatio = 1.380000

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Tree)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_SVGIcon)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Node)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_Button)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_11T03_39_06Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_06_11T03_39_06Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_11T03_39_06Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Tree.LinkShapes = append(__GongStructShape__000000_Default_Tree.LinkShapes, __LinkShape__000000_RootNodes)
	__GongStructShape__000002_Default_Node.AttributeShapes = append(__GongStructShape__000002_Default_Node.AttributeShapes, __AttributeShape__000000_Name)
	__GongStructShape__000002_Default_Node.AttributeShapes = append(__GongStructShape__000002_Default_Node.AttributeShapes, __AttributeShape__000001_FontStyle)
	__GongStructShape__000002_Default_Node.AttributeShapes = append(__GongStructShape__000002_Default_Node.AttributeShapes, __AttributeShape__000002_BackgroundColor)
	__GongStructShape__000002_Default_Node.AttributeShapes = append(__GongStructShape__000002_Default_Node.AttributeShapes, __AttributeShape__000003_IsExpanded)
	__GongStructShape__000002_Default_Node.AttributeShapes = append(__GongStructShape__000002_Default_Node.AttributeShapes, __AttributeShape__000004_HasCheckboxButton)
	__GongStructShape__000002_Default_Node.AttributeShapes = append(__GongStructShape__000002_Default_Node.AttributeShapes, __AttributeShape__000005_IsChecked)
	__GongStructShape__000002_Default_Node.AttributeShapes = append(__GongStructShape__000002_Default_Node.AttributeShapes, __AttributeShape__000006_IsCheckboxDisabled)
	__GongStructShape__000002_Default_Node.AttributeShapes = append(__GongStructShape__000002_Default_Node.AttributeShapes, __AttributeShape__000007_CheckboxHasToolTip)
	__GongStructShape__000002_Default_Node.AttributeShapes = append(__GongStructShape__000002_Default_Node.AttributeShapes, __AttributeShape__000008_CheckboxToolTipText)
	__GongStructShape__000002_Default_Node.AttributeShapes = append(__GongStructShape__000002_Default_Node.AttributeShapes, __AttributeShape__000009_CheckboxToolTipPosition)
	__GongStructShape__000002_Default_Node.AttributeShapes = append(__GongStructShape__000002_Default_Node.AttributeShapes, __AttributeShape__000010_HasSecondCheckboxButton)
	__GongStructShape__000002_Default_Node.AttributeShapes = append(__GongStructShape__000002_Default_Node.AttributeShapes, __AttributeShape__000011_IsSecondCheckboxChecked)
	__GongStructShape__000002_Default_Node.AttributeShapes = append(__GongStructShape__000002_Default_Node.AttributeShapes, __AttributeShape__000012_IsSecondCheckboxDisabled)
	__GongStructShape__000002_Default_Node.AttributeShapes = append(__GongStructShape__000002_Default_Node.AttributeShapes, __AttributeShape__000013_TextAfterSecondCheckbox)
	__GongStructShape__000002_Default_Node.AttributeShapes = append(__GongStructShape__000002_Default_Node.AttributeShapes, __AttributeShape__000014_IsInEditMode)
	__GongStructShape__000002_Default_Node.LinkShapes = append(__GongStructShape__000002_Default_Node.LinkShapes, __LinkShape__000001_Buttons)
	__GongStructShape__000002_Default_Node.LinkShapes = append(__GongStructShape__000002_Default_Node.LinkShapes, __LinkShape__000002_Children)
	__GongStructShape__000003_Default_Button.AttributeShapes = append(__GongStructShape__000003_Default_Button.AttributeShapes, __AttributeShape__000015_Name)
	__GongStructShape__000003_Default_Button.AttributeShapes = append(__GongStructShape__000003_Default_Button.AttributeShapes, __AttributeShape__000016_Icon)
	__GongStructShape__000003_Default_Button.AttributeShapes = append(__GongStructShape__000003_Default_Button.AttributeShapes, __AttributeShape__000017_IsDisabled)
	__GongStructShape__000003_Default_Button.AttributeShapes = append(__GongStructShape__000003_Default_Button.AttributeShapes, __AttributeShape__000018_HasToolTip)
	__GongStructShape__000003_Default_Button.AttributeShapes = append(__GongStructShape__000003_Default_Button.AttributeShapes, __AttributeShape__000019_ToolTipText)
	__GongStructShape__000003_Default_Button.AttributeShapes = append(__GongStructShape__000003_Default_Button.AttributeShapes, __AttributeShape__000020_ToolTipPosition)
	__GongStructShape__000003_Default_Button.LinkShapes = append(__GongStructShape__000003_Default_Button.LinkShapes, __LinkShape__000003_SVGIcon)
	// setup of LinkShape instances pointers
}
