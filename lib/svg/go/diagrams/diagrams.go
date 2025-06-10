package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/svg/go/models"
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
	__AttributeShape__000001_Content := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000002_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000003_StartAnchorType := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z := (&models.DiagramPackage{}).Stage(stage)

	__GongEnumShape__000000_Default_AnchorType := (&models.GongEnumShape{}).Stage(stage)
	__GongEnumShape__000001_Default_RectAnchorType := (&models.GongEnumShape{}).Stage(stage)

	__GongEnumValueShape__000000_ANCHOR_TOP := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000001_ANCHOR_BOTTOM := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000002_ANCHOR_LEFT := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000003_ANCHOR_RIGHT := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000004_ANCHOR_CENTER := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000005_RECT_TOP := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000006_RECT_TOP_LEFT := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000007_RECT_TOP_RIGHT := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000008_RECT_BOTTOM := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000009_RECT_BOTTOM_LEFT := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000010_RECT_BOTTOM_LEFT_LEFT := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000011_RECT_BOTTOM_BOTTOM_LEFT := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000012_RECT_BOTTOM_RIGHT := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000013_RECT_BOTTOM_INSIDE_RIGHT := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000014_RECT_LEFT := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000015_RECT_RIGHT := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000016_RECT_CENTER := (&models.GongEnumValueShape{}).Stage(stage)

	__GongStructShape__000000_Default_LinkAnchoredText := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_Link := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_Rect := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_End := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000001_Start := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000002_TextAtArrowStart := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000003_TextAtArrowEnd := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_Name.Name = `Name`
	__AttributeShape__000000_Name.IdentifierMeta = ref_models.LinkAnchoredText{}.Name
	__AttributeShape__000000_Name.FieldTypeAsString = ``
	__AttributeShape__000000_Name.Structname = `LinkAnchoredText`
	__AttributeShape__000000_Name.Fieldtypename = `string`

	__AttributeShape__000001_Content.Name = `Content`
	__AttributeShape__000001_Content.IdentifierMeta = ref_models.LinkAnchoredText{}.Content
	__AttributeShape__000001_Content.FieldTypeAsString = ``
	__AttributeShape__000001_Content.Structname = `LinkAnchoredText`
	__AttributeShape__000001_Content.Fieldtypename = `string`

	__AttributeShape__000002_Name.Name = `Name`
	__AttributeShape__000002_Name.IdentifierMeta = ref_models.Link{}.Name
	__AttributeShape__000002_Name.FieldTypeAsString = ``
	__AttributeShape__000002_Name.Structname = `Link`
	__AttributeShape__000002_Name.Fieldtypename = `string`

	__AttributeShape__000003_StartAnchorType.Name = `StartAnchorType`
	__AttributeShape__000003_StartAnchorType.IdentifierMeta = ref_models.Link{}.StartAnchorType
	__AttributeShape__000003_StartAnchorType.FieldTypeAsString = ``
	__AttributeShape__000003_StartAnchorType.Structname = `Link`
	__AttributeShape__000003_StartAnchorType.Fieldtypename = `AnchorType`

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = `[false,false,false,false,false,true,false]`
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = `[false,false,false,false,false,false,false,true]`
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.Name = `Diagram Package created the 2025-05-05T06:12:28Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.AbsolutePathToDiagramPackage = ``

	__GongEnumShape__000000_Default_AnchorType.Name = `Default-AnchorType`
	__GongEnumShape__000000_Default_AnchorType.X = 1421.000000
	__GongEnumShape__000000_Default_AnchorType.Y = 49.000000
	__GongEnumShape__000000_Default_AnchorType.IdentifierMeta = new(ref_models.AnchorType)
	__GongEnumShape__000000_Default_AnchorType.Width = 240.000000
	__GongEnumShape__000000_Default_AnchorType.Height = 163.000000
	__GongEnumShape__000000_Default_AnchorType.IsExpanded = false

	__GongEnumShape__000001_Default_RectAnchorType.Name = `Default-RectAnchorType`
	__GongEnumShape__000001_Default_RectAnchorType.X = 1147.000000
	__GongEnumShape__000001_Default_RectAnchorType.Y = 49.000000
	__GongEnumShape__000001_Default_RectAnchorType.IdentifierMeta = new(ref_models.RectAnchorType)
	__GongEnumShape__000001_Default_RectAnchorType.Width = 253.000000
	__GongEnumShape__000001_Default_RectAnchorType.Height = 303.000000
	__GongEnumShape__000001_Default_RectAnchorType.IsExpanded = false

	__GongEnumValueShape__000000_ANCHOR_TOP.Name = `ANCHOR_TOP`
	__GongEnumValueShape__000000_ANCHOR_TOP.IdentifierMeta = ref_models.ANCHOR_TOP

	__GongEnumValueShape__000001_ANCHOR_BOTTOM.Name = `ANCHOR_BOTTOM`
	__GongEnumValueShape__000001_ANCHOR_BOTTOM.IdentifierMeta = ref_models.ANCHOR_BOTTOM

	__GongEnumValueShape__000002_ANCHOR_LEFT.Name = `ANCHOR_LEFT`
	__GongEnumValueShape__000002_ANCHOR_LEFT.IdentifierMeta = ref_models.ANCHOR_LEFT

	__GongEnumValueShape__000003_ANCHOR_RIGHT.Name = `ANCHOR_RIGHT`
	__GongEnumValueShape__000003_ANCHOR_RIGHT.IdentifierMeta = ref_models.ANCHOR_RIGHT

	__GongEnumValueShape__000004_ANCHOR_CENTER.Name = `ANCHOR_CENTER`
	__GongEnumValueShape__000004_ANCHOR_CENTER.IdentifierMeta = ref_models.ANCHOR_CENTER

	__GongEnumValueShape__000005_RECT_TOP.Name = `RECT_TOP`
	__GongEnumValueShape__000005_RECT_TOP.IdentifierMeta = ref_models.RECT_TOP

	__GongEnumValueShape__000006_RECT_TOP_LEFT.Name = `RECT_TOP_LEFT`
	__GongEnumValueShape__000006_RECT_TOP_LEFT.IdentifierMeta = ref_models.RECT_TOP_LEFT

	__GongEnumValueShape__000007_RECT_TOP_RIGHT.Name = `RECT_TOP_RIGHT`
	__GongEnumValueShape__000007_RECT_TOP_RIGHT.IdentifierMeta = ref_models.RECT_TOP_RIGHT

	__GongEnumValueShape__000008_RECT_BOTTOM.Name = `RECT_BOTTOM`
	__GongEnumValueShape__000008_RECT_BOTTOM.IdentifierMeta = ref_models.RECT_BOTTOM

	__GongEnumValueShape__000009_RECT_BOTTOM_LEFT.Name = `RECT_BOTTOM_LEFT`
	__GongEnumValueShape__000009_RECT_BOTTOM_LEFT.IdentifierMeta = ref_models.RECT_BOTTOM_LEFT

	__GongEnumValueShape__000010_RECT_BOTTOM_LEFT_LEFT.Name = `RECT_BOTTOM_LEFT_LEFT`
	__GongEnumValueShape__000010_RECT_BOTTOM_LEFT_LEFT.IdentifierMeta = ref_models.RECT_BOTTOM_LEFT_LEFT

	__GongEnumValueShape__000011_RECT_BOTTOM_BOTTOM_LEFT.Name = `RECT_BOTTOM_BOTTOM_LEFT`
	__GongEnumValueShape__000011_RECT_BOTTOM_BOTTOM_LEFT.IdentifierMeta = ref_models.RECT_BOTTOM_BOTTOM_LEFT

	__GongEnumValueShape__000012_RECT_BOTTOM_RIGHT.Name = `RECT_BOTTOM_RIGHT`
	__GongEnumValueShape__000012_RECT_BOTTOM_RIGHT.IdentifierMeta = ref_models.RECT_BOTTOM_RIGHT

	__GongEnumValueShape__000013_RECT_BOTTOM_INSIDE_RIGHT.Name = `RECT_BOTTOM_INSIDE_RIGHT`
	__GongEnumValueShape__000013_RECT_BOTTOM_INSIDE_RIGHT.IdentifierMeta = ref_models.RECT_BOTTOM_INSIDE_RIGHT

	__GongEnumValueShape__000014_RECT_LEFT.Name = `RECT_LEFT`
	__GongEnumValueShape__000014_RECT_LEFT.IdentifierMeta = ref_models.RECT_LEFT

	__GongEnumValueShape__000015_RECT_RIGHT.Name = `RECT_RIGHT`
	__GongEnumValueShape__000015_RECT_RIGHT.IdentifierMeta = ref_models.RECT_RIGHT

	__GongEnumValueShape__000016_RECT_CENTER.Name = `RECT_CENTER`
	__GongEnumValueShape__000016_RECT_CENTER.IdentifierMeta = ref_models.RECT_CENTER

	__GongStructShape__000000_Default_LinkAnchoredText.Name = `Default-LinkAnchoredText`
	__GongStructShape__000000_Default_LinkAnchoredText.X = 854.000000
	__GongStructShape__000000_Default_LinkAnchoredText.Y = 259.000000
	__GongStructShape__000000_Default_LinkAnchoredText.IdentifierMeta = ref_models.LinkAnchoredText{}
	__GongStructShape__000000_Default_LinkAnchoredText.ShowNbInstances = false
	__GongStructShape__000000_Default_LinkAnchoredText.NbInstances = 0
	__GongStructShape__000000_Default_LinkAnchoredText.Width = 240.000000
	__GongStructShape__000000_Default_LinkAnchoredText.Height = 103.000000
	__GongStructShape__000000_Default_LinkAnchoredText.IsSelected = false

	__GongStructShape__000001_Default_Link.Name = `Default-Link`
	__GongStructShape__000001_Default_Link.X = 75.000000
	__GongStructShape__000001_Default_Link.Y = 77.000000
	__GongStructShape__000001_Default_Link.IdentifierMeta = ref_models.Link{}
	__GongStructShape__000001_Default_Link.ShowNbInstances = false
	__GongStructShape__000001_Default_Link.NbInstances = 0
	__GongStructShape__000001_Default_Link.Width = 240.000000
	__GongStructShape__000001_Default_Link.Height = 103.000000
	__GongStructShape__000001_Default_Link.IsSelected = false

	__GongStructShape__000002_Default_Rect.Name = `Default-Rect`
	__GongStructShape__000002_Default_Rect.X = 842.000000
	__GongStructShape__000002_Default_Rect.Y = 46.000000
	__GongStructShape__000002_Default_Rect.IdentifierMeta = ref_models.Rect{}
	__GongStructShape__000002_Default_Rect.ShowNbInstances = false
	__GongStructShape__000002_Default_Rect.NbInstances = 0
	__GongStructShape__000002_Default_Rect.Width = 240.000000
	__GongStructShape__000002_Default_Rect.Height = 182.000000
	__GongStructShape__000002_Default_Rect.IsSelected = false

	__LinkShape__000000_End.Name = `End`
	__LinkShape__000000_End.IdentifierMeta = ref_models.Link{}.End
	__LinkShape__000000_End.FieldTypeIdentifierMeta = ref_models.Rect{}
	__LinkShape__000000_End.FieldOffsetX = 0.000000
	__LinkShape__000000_End.FieldOffsetY = 0.000000
	__LinkShape__000000_End.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000000_End.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_End.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_End.SourceMultiplicity = models.MANY
	__LinkShape__000000_End.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_End.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_End.X = 537.500000
	__LinkShape__000000_End.Y = 238.500000
	__LinkShape__000000_End.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_End.StartRatio = 0.500000
	__LinkShape__000000_End.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_End.EndRatio = 0.640261
	__LinkShape__000000_End.CornerOffsetRatio = 1.380000

	__LinkShape__000001_Start.Name = `Start`
	__LinkShape__000001_Start.IdentifierMeta = ref_models.Link{}.Start
	__LinkShape__000001_Start.FieldTypeIdentifierMeta = ref_models.Rect{}
	__LinkShape__000001_Start.FieldOffsetX = 0.000000
	__LinkShape__000001_Start.FieldOffsetY = 0.000000
	__LinkShape__000001_Start.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000001_Start.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Start.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Start.SourceMultiplicity = models.MANY
	__LinkShape__000001_Start.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Start.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Start.X = 537.500000
	__LinkShape__000001_Start.Y = 238.500000
	__LinkShape__000001_Start.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Start.StartRatio = 0.513821
	__LinkShape__000001_Start.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Start.EndRatio = 0.169910
	__LinkShape__000001_Start.CornerOffsetRatio = 1.381380

	__LinkShape__000002_TextAtArrowStart.Name = `TextAtArrowStart`
	__LinkShape__000002_TextAtArrowStart.IdentifierMeta = ref_models.Link{}.TextAtArrowStart
	__LinkShape__000002_TextAtArrowStart.FieldTypeIdentifierMeta = ref_models.LinkAnchoredText{}
	__LinkShape__000002_TextAtArrowStart.FieldOffsetX = 0.000000
	__LinkShape__000002_TextAtArrowStart.FieldOffsetY = 0.000000
	__LinkShape__000002_TextAtArrowStart.TargetMultiplicity = models.MANY
	__LinkShape__000002_TextAtArrowStart.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000002_TextAtArrowStart.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000002_TextAtArrowStart.SourceMultiplicity = models.MANY
	__LinkShape__000002_TextAtArrowStart.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000002_TextAtArrowStart.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000002_TextAtArrowStart.X = 821.500000
	__LinkShape__000002_TextAtArrowStart.Y = 110.000000
	__LinkShape__000002_TextAtArrowStart.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_TextAtArrowStart.StartRatio = 0.500000
	__LinkShape__000002_TextAtArrowStart.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_TextAtArrowStart.EndRatio = 0.698287
	__LinkShape__000002_TextAtArrowStart.CornerOffsetRatio = 1.380000

	__LinkShape__000003_TextAtArrowEnd.Name = `TextAtArrowEnd`
	__LinkShape__000003_TextAtArrowEnd.IdentifierMeta = ref_models.Link{}.TextAtArrowEnd
	__LinkShape__000003_TextAtArrowEnd.FieldTypeIdentifierMeta = ref_models.LinkAnchoredText{}
	__LinkShape__000003_TextAtArrowEnd.FieldOffsetX = 0.000000
	__LinkShape__000003_TextAtArrowEnd.FieldOffsetY = 0.000000
	__LinkShape__000003_TextAtArrowEnd.TargetMultiplicity = models.MANY
	__LinkShape__000003_TextAtArrowEnd.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000003_TextAtArrowEnd.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000003_TextAtArrowEnd.SourceMultiplicity = models.MANY
	__LinkShape__000003_TextAtArrowEnd.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000003_TextAtArrowEnd.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000003_TextAtArrowEnd.X = 821.500000
	__LinkShape__000003_TextAtArrowEnd.Y = 110.000000
	__LinkShape__000003_TextAtArrowEnd.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000003_TextAtArrowEnd.StartRatio = 0.504112
	__LinkShape__000003_TextAtArrowEnd.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000003_TextAtArrowEnd.EndRatio = 0.193433
	__LinkShape__000003_TextAtArrowEnd.CornerOffsetRatio = 1.380000

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_LinkAnchoredText)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Link)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Rect)
	__Classdiagram__000000_Default.GongEnumShapes = append(__Classdiagram__000000_Default.GongEnumShapes, __GongEnumShape__000000_Default_AnchorType)
	__Classdiagram__000000_Default.GongEnumShapes = append(__Classdiagram__000000_Default.GongEnumShapes, __GongEnumShape__000001_Default_RectAnchorType)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_05T06_12_28Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongEnumShape instances pointers
	__GongEnumShape__000000_Default_AnchorType.GongEnumValueShapes = append(__GongEnumShape__000000_Default_AnchorType.GongEnumValueShapes, __GongEnumValueShape__000000_ANCHOR_TOP)
	__GongEnumShape__000000_Default_AnchorType.GongEnumValueShapes = append(__GongEnumShape__000000_Default_AnchorType.GongEnumValueShapes, __GongEnumValueShape__000001_ANCHOR_BOTTOM)
	__GongEnumShape__000000_Default_AnchorType.GongEnumValueShapes = append(__GongEnumShape__000000_Default_AnchorType.GongEnumValueShapes, __GongEnumValueShape__000002_ANCHOR_LEFT)
	__GongEnumShape__000000_Default_AnchorType.GongEnumValueShapes = append(__GongEnumShape__000000_Default_AnchorType.GongEnumValueShapes, __GongEnumValueShape__000003_ANCHOR_RIGHT)
	__GongEnumShape__000000_Default_AnchorType.GongEnumValueShapes = append(__GongEnumShape__000000_Default_AnchorType.GongEnumValueShapes, __GongEnumValueShape__000004_ANCHOR_CENTER)
	__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes = append(__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes, __GongEnumValueShape__000005_RECT_TOP)
	__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes = append(__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes, __GongEnumValueShape__000006_RECT_TOP_LEFT)
	__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes = append(__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes, __GongEnumValueShape__000007_RECT_TOP_RIGHT)
	__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes = append(__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes, __GongEnumValueShape__000008_RECT_BOTTOM)
	__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes = append(__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes, __GongEnumValueShape__000009_RECT_BOTTOM_LEFT)
	__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes = append(__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes, __GongEnumValueShape__000010_RECT_BOTTOM_LEFT_LEFT)
	__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes = append(__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes, __GongEnumValueShape__000011_RECT_BOTTOM_BOTTOM_LEFT)
	__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes = append(__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes, __GongEnumValueShape__000012_RECT_BOTTOM_RIGHT)
	__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes = append(__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes, __GongEnumValueShape__000013_RECT_BOTTOM_INSIDE_RIGHT)
	__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes = append(__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes, __GongEnumValueShape__000014_RECT_LEFT)
	__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes = append(__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes, __GongEnumValueShape__000015_RECT_RIGHT)
	__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes = append(__GongEnumShape__000001_Default_RectAnchorType.GongEnumValueShapes, __GongEnumValueShape__000016_RECT_CENTER)
	// setup of GongEnumValueShape instances pointers
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_LinkAnchoredText.AttributeShapes = append(__GongStructShape__000000_Default_LinkAnchoredText.AttributeShapes, __AttributeShape__000000_Name)
	__GongStructShape__000000_Default_LinkAnchoredText.AttributeShapes = append(__GongStructShape__000000_Default_LinkAnchoredText.AttributeShapes, __AttributeShape__000001_Content)
	__GongStructShape__000001_Default_Link.AttributeShapes = append(__GongStructShape__000001_Default_Link.AttributeShapes, __AttributeShape__000002_Name)
	__GongStructShape__000001_Default_Link.AttributeShapes = append(__GongStructShape__000001_Default_Link.AttributeShapes, __AttributeShape__000003_StartAnchorType)
	__GongStructShape__000001_Default_Link.LinkShapes = append(__GongStructShape__000001_Default_Link.LinkShapes, __LinkShape__000000_End)
	__GongStructShape__000001_Default_Link.LinkShapes = append(__GongStructShape__000001_Default_Link.LinkShapes, __LinkShape__000001_Start)
	__GongStructShape__000001_Default_Link.LinkShapes = append(__GongStructShape__000001_Default_Link.LinkShapes, __LinkShape__000002_TextAtArrowStart)
	__GongStructShape__000001_Default_Link.LinkShapes = append(__GongStructShape__000001_Default_Link.LinkShapes, __LinkShape__000003_TextAtArrowEnd)
	// setup of LinkShape instances pointers
}
