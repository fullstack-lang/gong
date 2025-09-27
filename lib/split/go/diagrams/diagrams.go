package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/split/go/models"
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

	const __write__local_time = "2025-09-27 11:14:10.557646 CEST"
	const __write__utc_time__ = "2025-09-27 09:14:10.557646 UTC"

	const __commitId__ = "0000000209"

	// Declaration of instances to stage

	__AttributeShape__000000_StackName := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000001_IsSelectedView := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000002_ShowViewName := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000003_SVG := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000004_Height := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000005_Width := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000006_Width := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000007_Height := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000008_SVG := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000009_ShowNameInHeader := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000010_Size := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000011_IsAny := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000012_HasDiv := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000013_DivStyle := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000014_NoName_yet := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_12T00_16_22Z := (&models.DiagramPackage{}).Stage(stage)

	__GongEnumShape__000000_Default_Direction := (&models.GongEnumShape{}).Stage(stage)

	__GongEnumValueShape__000000_Horizontal := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000001_Vertical := (&models.GongEnumValueShape{}).Stage(stage)

	__GongStructShape__000000_Default_Split := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_Xlsx := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_View := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000003_Default_Tree := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000004_Default_Tone := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000005_Default_Title := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000006_Default_Table := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000007_Default_Svg := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000008_Default_Slider := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000009_Default_Markdown := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000010_Default_LogoOnTheRight := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000011_Default_LogoOnTheLeft := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000012_Default_Load := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000013_Default_Form := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000014_Default_FavIcon := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000015_Default_Cursor := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000016_Default_Button := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000017_Default_AsSplitArea := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000018_NoName_yet := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000019_Default_AsSplit := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_NoName_yet := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000001_AsSplit := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000002_Button := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000003_Cursor := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000004_Form := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000005_Load := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000006_Markdown := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000007_Slider := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000008_Split := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000009_Svg := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000010_Table := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000011_Tone := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000012_Tree := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000013_Xlsx := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000014_NoName_yet := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000015_AsSplitAreas := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000016_RootAsSplitAreas := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_StackName.Name = `StackName`
	__AttributeShape__000000_StackName.IdentifierMeta = ref_models.Split{}.StackName
	__AttributeShape__000000_StackName.FieldTypeAsString = ``
	__AttributeShape__000000_StackName.Structname = `Split`
	__AttributeShape__000000_StackName.Fieldtypename = `string`

	__AttributeShape__000001_IsSelectedView.Name = `IsSelectedView`
	__AttributeShape__000001_IsSelectedView.IdentifierMeta = ref_models.View{}.IsSelectedView
	__AttributeShape__000001_IsSelectedView.FieldTypeAsString = ``
	__AttributeShape__000001_IsSelectedView.Structname = `View`
	__AttributeShape__000001_IsSelectedView.Fieldtypename = `bool`

	__AttributeShape__000002_ShowViewName.Name = `ShowViewName`
	__AttributeShape__000002_ShowViewName.IdentifierMeta = ref_models.View{}.ShowViewName
	__AttributeShape__000002_ShowViewName.FieldTypeAsString = ``
	__AttributeShape__000002_ShowViewName.Structname = `View`
	__AttributeShape__000002_ShowViewName.Fieldtypename = `bool`

	__AttributeShape__000003_SVG.Name = `SVG`
	__AttributeShape__000003_SVG.IdentifierMeta = ref_models.LogoOnTheRight{}.SVG
	__AttributeShape__000003_SVG.FieldTypeAsString = ``
	__AttributeShape__000003_SVG.Structname = `LogoOnTheRight`
	__AttributeShape__000003_SVG.Fieldtypename = `string`

	__AttributeShape__000004_Height.Name = `Height`
	__AttributeShape__000004_Height.IdentifierMeta = ref_models.LogoOnTheRight{}.Height
	__AttributeShape__000004_Height.FieldTypeAsString = ``
	__AttributeShape__000004_Height.Structname = `LogoOnTheRight`
	__AttributeShape__000004_Height.Fieldtypename = `int`

	__AttributeShape__000005_Width.Name = `Width`
	__AttributeShape__000005_Width.IdentifierMeta = ref_models.LogoOnTheRight{}.Width
	__AttributeShape__000005_Width.FieldTypeAsString = ``
	__AttributeShape__000005_Width.Structname = `LogoOnTheRight`
	__AttributeShape__000005_Width.Fieldtypename = `int`

	__AttributeShape__000006_Width.Name = `Width`
	__AttributeShape__000006_Width.IdentifierMeta = ref_models.LogoOnTheLeft{}.Width
	__AttributeShape__000006_Width.FieldTypeAsString = ``
	__AttributeShape__000006_Width.Structname = `LogoOnTheLeft`
	__AttributeShape__000006_Width.Fieldtypename = `int`

	__AttributeShape__000007_Height.Name = `Height`
	__AttributeShape__000007_Height.IdentifierMeta = ref_models.LogoOnTheLeft{}.Height
	__AttributeShape__000007_Height.FieldTypeAsString = ``
	__AttributeShape__000007_Height.Structname = `LogoOnTheLeft`
	__AttributeShape__000007_Height.Fieldtypename = `int`

	__AttributeShape__000008_SVG.Name = `SVG`
	__AttributeShape__000008_SVG.IdentifierMeta = ref_models.LogoOnTheLeft{}.SVG
	__AttributeShape__000008_SVG.FieldTypeAsString = ``
	__AttributeShape__000008_SVG.Structname = `LogoOnTheLeft`
	__AttributeShape__000008_SVG.Fieldtypename = `string`

	__AttributeShape__000009_ShowNameInHeader.Name = `ShowNameInHeader`
	__AttributeShape__000009_ShowNameInHeader.IdentifierMeta = ref_models.AsSplitArea{}.ShowNameInHeader
	__AttributeShape__000009_ShowNameInHeader.FieldTypeAsString = ``
	__AttributeShape__000009_ShowNameInHeader.Structname = `AsSplitArea`
	__AttributeShape__000009_ShowNameInHeader.Fieldtypename = `bool`

	__AttributeShape__000010_Size.Name = `Size`
	__AttributeShape__000010_Size.IdentifierMeta = ref_models.AsSplitArea{}.Size
	__AttributeShape__000010_Size.FieldTypeAsString = ``
	__AttributeShape__000010_Size.Structname = `AsSplitArea`
	__AttributeShape__000010_Size.Fieldtypename = `float64`

	__AttributeShape__000011_IsAny.Name = `IsAny`
	__AttributeShape__000011_IsAny.IdentifierMeta = ref_models.AsSplitArea{}.IsAny
	__AttributeShape__000011_IsAny.FieldTypeAsString = ``
	__AttributeShape__000011_IsAny.Structname = `AsSplitArea`
	__AttributeShape__000011_IsAny.Fieldtypename = `bool`

	__AttributeShape__000012_HasDiv.Name = `HasDiv`
	__AttributeShape__000012_HasDiv.IdentifierMeta = ref_models.AsSplitArea{}.HasDiv
	__AttributeShape__000012_HasDiv.FieldTypeAsString = ``
	__AttributeShape__000012_HasDiv.Structname = `AsSplitArea`
	__AttributeShape__000012_HasDiv.Fieldtypename = `bool`

	__AttributeShape__000013_DivStyle.Name = `DivStyle`
	__AttributeShape__000013_DivStyle.IdentifierMeta = ref_models.AsSplitArea{}.DivStyle
	__AttributeShape__000013_DivStyle.FieldTypeAsString = ``
	__AttributeShape__000013_DivStyle.Structname = `AsSplitArea`
	__AttributeShape__000013_DivStyle.Fieldtypename = `string`

	__AttributeShape__000014_NoName_yet.Name = `NoName yet`
	__AttributeShape__000014_NoName_yet.FieldTypeAsString = ``
	__AttributeShape__000014_NoName_yet.Structname = ``
	__AttributeShape__000014_NoName_yet.Fieldtypename = ``

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = `[true,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,true]`
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = `[false]`
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_12T00_16_22Z.Name = `Diagram Package created the 2025-06-12T00:16:22Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_12T00_16_22Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_12T00_16_22Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_12T00_16_22Z.AbsolutePathToDiagramPackage = ``

	__GongEnumShape__000000_Default_Direction.Name = `Default-Direction`
	__GongEnumShape__000000_Default_Direction.X = 935.000000
	__GongEnumShape__000000_Default_Direction.Y = 88.000000
	__GongEnumShape__000000_Default_Direction.IdentifierMeta = new(ref_models.Direction)
	__GongEnumShape__000000_Default_Direction.Width = 240.000000
	__GongEnumShape__000000_Default_Direction.Height = 103.000000
	__GongEnumShape__000000_Default_Direction.IsExpanded = false

	__GongEnumValueShape__000000_Horizontal.Name = `Horizontal`
	__GongEnumValueShape__000000_Horizontal.IdentifierMeta = ref_models.Horizontal

	__GongEnumValueShape__000001_Vertical.Name = `Vertical`
	__GongEnumValueShape__000001_Vertical.IdentifierMeta = ref_models.Vertical

	__GongStructShape__000000_Default_Split.Name = `Default-Split`
	__GongStructShape__000000_Default_Split.X = 579.000000
	__GongStructShape__000000_Default_Split.Y = 994.000000
	__GongStructShape__000000_Default_Split.IdentifierMeta = ref_models.Split{}
	__GongStructShape__000000_Default_Split.Width = 240.000000
	__GongStructShape__000000_Default_Split.Height = 83.000000
	__GongStructShape__000000_Default_Split.IsSelected = false

	__GongStructShape__000001_Default_Xlsx.Name = `Default-Xlsx`
	__GongStructShape__000001_Default_Xlsx.X = 579.000000
	__GongStructShape__000001_Default_Xlsx.Y = 908.000000
	__GongStructShape__000001_Default_Xlsx.IdentifierMeta = ref_models.Xlsx{}
	__GongStructShape__000001_Default_Xlsx.Width = 240.000000
	__GongStructShape__000001_Default_Xlsx.Height = 63.000000
	__GongStructShape__000001_Default_Xlsx.IsSelected = false

	__GongStructShape__000002_Default_View.Name = `Default-View`
	__GongStructShape__000002_Default_View.X = 49.000000
	__GongStructShape__000002_Default_View.Y = 155.000000
	__GongStructShape__000002_Default_View.IdentifierMeta = ref_models.View{}
	__GongStructShape__000002_Default_View.Width = 240.000000
	__GongStructShape__000002_Default_View.Height = 123.000000
	__GongStructShape__000002_Default_View.IsSelected = false

	__GongStructShape__000003_Default_Tree.Name = `Default-Tree`
	__GongStructShape__000003_Default_Tree.X = 581.000000
	__GongStructShape__000003_Default_Tree.Y = 840.000000
	__GongStructShape__000003_Default_Tree.IdentifierMeta = ref_models.Tree{}
	__GongStructShape__000003_Default_Tree.Width = 240.000000
	__GongStructShape__000003_Default_Tree.Height = 63.000000
	__GongStructShape__000003_Default_Tree.IsSelected = false

	__GongStructShape__000004_Default_Tone.Name = `Default-Tone`
	__GongStructShape__000004_Default_Tone.X = 581.000000
	__GongStructShape__000004_Default_Tone.Y = 749.000000
	__GongStructShape__000004_Default_Tone.IdentifierMeta = ref_models.Tone{}
	__GongStructShape__000004_Default_Tone.Width = 240.000000
	__GongStructShape__000004_Default_Tone.Height = 63.000000
	__GongStructShape__000004_Default_Tone.IsSelected = false

	__GongStructShape__000005_Default_Title.Name = `Default-Title`
	__GongStructShape__000005_Default_Title.X = 1275.000000
	__GongStructShape__000005_Default_Title.Y = 497.000000
	__GongStructShape__000005_Default_Title.IdentifierMeta = ref_models.Title{}
	__GongStructShape__000005_Default_Title.Width = 240.000000
	__GongStructShape__000005_Default_Title.Height = 63.000000
	__GongStructShape__000005_Default_Title.IsSelected = false

	__GongStructShape__000006_Default_Table.Name = `Default-Table`
	__GongStructShape__000006_Default_Table.X = 571.000000
	__GongStructShape__000006_Default_Table.Y = 489.000000
	__GongStructShape__000006_Default_Table.IdentifierMeta = ref_models.Table{}
	__GongStructShape__000006_Default_Table.Width = 240.000000
	__GongStructShape__000006_Default_Table.Height = 63.000000
	__GongStructShape__000006_Default_Table.IsSelected = false

	__GongStructShape__000007_Default_Svg.Name = `Default-Svg`
	__GongStructShape__000007_Default_Svg.X = 582.000000
	__GongStructShape__000007_Default_Svg.Y = 667.000000
	__GongStructShape__000007_Default_Svg.IdentifierMeta = ref_models.Svg{}
	__GongStructShape__000007_Default_Svg.Width = 240.000000
	__GongStructShape__000007_Default_Svg.Height = 63.000000
	__GongStructShape__000007_Default_Svg.IsSelected = false

	__GongStructShape__000008_Default_Slider.Name = `Default-Slider`
	__GongStructShape__000008_Default_Slider.X = 557.000000
	__GongStructShape__000008_Default_Slider.Y = 400.000000
	__GongStructShape__000008_Default_Slider.IdentifierMeta = ref_models.Slider{}
	__GongStructShape__000008_Default_Slider.Width = 240.000000
	__GongStructShape__000008_Default_Slider.Height = 63.000000
	__GongStructShape__000008_Default_Slider.IsSelected = false

	__GongStructShape__000009_Default_Markdown.Name = `Default-Markdown`
	__GongStructShape__000009_Default_Markdown.X = 576.000000
	__GongStructShape__000009_Default_Markdown.Y = 575.000000
	__GongStructShape__000009_Default_Markdown.IdentifierMeta = ref_models.Markdown{}
	__GongStructShape__000009_Default_Markdown.Width = 240.000000
	__GongStructShape__000009_Default_Markdown.Height = 63.000000
	__GongStructShape__000009_Default_Markdown.IsSelected = false

	__GongStructShape__000010_Default_LogoOnTheRight.Name = `Default-LogoOnTheRight`
	__GongStructShape__000010_Default_LogoOnTheRight.X = 1292.000000
	__GongStructShape__000010_Default_LogoOnTheRight.Y = 229.000000
	__GongStructShape__000010_Default_LogoOnTheRight.IdentifierMeta = ref_models.LogoOnTheRight{}
	__GongStructShape__000010_Default_LogoOnTheRight.Width = 240.000000
	__GongStructShape__000010_Default_LogoOnTheRight.Height = 143.000000
	__GongStructShape__000010_Default_LogoOnTheRight.IsSelected = false

	__GongStructShape__000011_Default_LogoOnTheLeft.Name = `Default-LogoOnTheLeft`
	__GongStructShape__000011_Default_LogoOnTheLeft.X = 1306.000000
	__GongStructShape__000011_Default_LogoOnTheLeft.Y = 51.000000
	__GongStructShape__000011_Default_LogoOnTheLeft.IdentifierMeta = ref_models.LogoOnTheLeft{}
	__GongStructShape__000011_Default_LogoOnTheLeft.Width = 240.000000
	__GongStructShape__000011_Default_LogoOnTheLeft.Height = 143.000000
	__GongStructShape__000011_Default_LogoOnTheLeft.IsSelected = false

	__GongStructShape__000012_Default_Load.Name = `Default-Load`
	__GongStructShape__000012_Default_Load.X = 555.000000
	__GongStructShape__000012_Default_Load.Y = 308.000000
	__GongStructShape__000012_Default_Load.IdentifierMeta = ref_models.Load{}
	__GongStructShape__000012_Default_Load.Width = 240.000000
	__GongStructShape__000012_Default_Load.Height = 63.000000
	__GongStructShape__000012_Default_Load.IsSelected = false

	__GongStructShape__000013_Default_Form.Name = `Default-Form`
	__GongStructShape__000013_Default_Form.X = 562.000000
	__GongStructShape__000013_Default_Form.Y = 229.000000
	__GongStructShape__000013_Default_Form.IdentifierMeta = ref_models.Form{}
	__GongStructShape__000013_Default_Form.Width = 240.000000
	__GongStructShape__000013_Default_Form.Height = 63.000000
	__GongStructShape__000013_Default_Form.IsSelected = false

	__GongStructShape__000014_Default_FavIcon.Name = `Default-FavIcon`
	__GongStructShape__000014_Default_FavIcon.X = 1283.000000
	__GongStructShape__000014_Default_FavIcon.Y = 402.000000
	__GongStructShape__000014_Default_FavIcon.IdentifierMeta = ref_models.FavIcon{}
	__GongStructShape__000014_Default_FavIcon.Width = 240.000000
	__GongStructShape__000014_Default_FavIcon.Height = 63.000000
	__GongStructShape__000014_Default_FavIcon.IsSelected = false

	__GongStructShape__000015_Default_Cursor.Name = `Default-Cursor`
	__GongStructShape__000015_Default_Cursor.X = 553.000000
	__GongStructShape__000015_Default_Cursor.Y = 123.000000
	__GongStructShape__000015_Default_Cursor.IdentifierMeta = ref_models.Cursor{}
	__GongStructShape__000015_Default_Cursor.Width = 240.000000
	__GongStructShape__000015_Default_Cursor.Height = 63.000000
	__GongStructShape__000015_Default_Cursor.IsSelected = false

	__GongStructShape__000016_Default_Button.Name = `Default-Button`
	__GongStructShape__000016_Default_Button.X = 551.000000
	__GongStructShape__000016_Default_Button.Y = 36.000000
	__GongStructShape__000016_Default_Button.IdentifierMeta = ref_models.Button{}
	__GongStructShape__000016_Default_Button.Width = 240.000000
	__GongStructShape__000016_Default_Button.Height = 63.000000
	__GongStructShape__000016_Default_Button.IsSelected = false

	__GongStructShape__000017_Default_AsSplitArea.Name = `Default-AsSplitArea`
	__GongStructShape__000017_Default_AsSplitArea.X = 35.000000
	__GongStructShape__000017_Default_AsSplitArea.Y = 512.000000
	__GongStructShape__000017_Default_AsSplitArea.IdentifierMeta = ref_models.AsSplitArea{}
	__GongStructShape__000017_Default_AsSplitArea.Width = 240.000000
	__GongStructShape__000017_Default_AsSplitArea.Height = 183.000000
	__GongStructShape__000017_Default_AsSplitArea.IsSelected = false

	__GongStructShape__000018_NoName_yet.Name = `NoName yet`
	__GongStructShape__000018_NoName_yet.X = 0.000000
	__GongStructShape__000018_NoName_yet.Y = 0.000000
	__GongStructShape__000018_NoName_yet.Width = 0.000000
	__GongStructShape__000018_NoName_yet.Height = 0.000000
	__GongStructShape__000018_NoName_yet.IsSelected = false

	__GongStructShape__000019_Default_AsSplit.Name = `Default-AsSplit`
	__GongStructShape__000019_Default_AsSplit.X = 38.000000
	__GongStructShape__000019_Default_AsSplit.Y = 830.000000
	__GongStructShape__000019_Default_AsSplit.IdentifierMeta = ref_models.AsSplit{}
	__GongStructShape__000019_Default_AsSplit.Width = 240.000000
	__GongStructShape__000019_Default_AsSplit.Height = 83.000000
	__GongStructShape__000019_Default_AsSplit.IsSelected = false

	__LinkShape__000000_NoName_yet.Name = `NoName yet`
	__LinkShape__000000_NoName_yet.FieldOffsetX = 0.000000
	__LinkShape__000000_NoName_yet.FieldOffsetY = 0.000000
	__LinkShape__000000_NoName_yet.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_NoName_yet.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_NoName_yet.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_NoName_yet.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_NoName_yet.X = 0.000000
	__LinkShape__000000_NoName_yet.Y = 0.000000
	__LinkShape__000000_NoName_yet.StartRatio = 0.000000
	__LinkShape__000000_NoName_yet.EndRatio = 0.000000
	__LinkShape__000000_NoName_yet.CornerOffsetRatio = 0.000000

	__LinkShape__000001_AsSplit.Name = `AsSplit`
	__LinkShape__000001_AsSplit.IdentifierMeta = ref_models.AsSplitArea{}.AsSplit
	__LinkShape__000001_AsSplit.FieldTypeIdentifierMeta = ref_models.AsSplit{}
	__LinkShape__000001_AsSplit.FieldOffsetX = 0.000000
	__LinkShape__000001_AsSplit.FieldOffsetY = 0.000000
	__LinkShape__000001_AsSplit.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000001_AsSplit.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000001_AsSplit.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000001_AsSplit.SourceMultiplicity = models.MANY
	__LinkShape__000001_AsSplit.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000001_AsSplit.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000001_AsSplit.X = 631.500000
	__LinkShape__000001_AsSplit.Y = 436.500000
	__LinkShape__000001_AsSplit.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_AsSplit.StartRatio = 0.500000
	__LinkShape__000001_AsSplit.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_AsSplit.EndRatio = 0.500000
	__LinkShape__000001_AsSplit.CornerOffsetRatio = 1.380000

	__LinkShape__000002_Button.Name = `Button`
	__LinkShape__000002_Button.IdentifierMeta = ref_models.AsSplitArea{}.Button
	__LinkShape__000002_Button.FieldTypeIdentifierMeta = ref_models.Button{}
	__LinkShape__000002_Button.FieldOffsetX = 0.000000
	__LinkShape__000002_Button.FieldOffsetY = 0.000000
	__LinkShape__000002_Button.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000002_Button.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000002_Button.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000002_Button.SourceMultiplicity = models.MANY
	__LinkShape__000002_Button.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000002_Button.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000002_Button.X = 571.000000
	__LinkShape__000002_Button.Y = 354.000000
	__LinkShape__000002_Button.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_Button.StartRatio = 0.500000
	__LinkShape__000002_Button.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_Button.EndRatio = 0.500000
	__LinkShape__000002_Button.CornerOffsetRatio = 1.380000

	__LinkShape__000003_Cursor.Name = `Cursor`
	__LinkShape__000003_Cursor.IdentifierMeta = ref_models.AsSplitArea{}.Cursor
	__LinkShape__000003_Cursor.FieldTypeIdentifierMeta = ref_models.Cursor{}
	__LinkShape__000003_Cursor.FieldOffsetX = 0.000000
	__LinkShape__000003_Cursor.FieldOffsetY = 0.000000
	__LinkShape__000003_Cursor.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000003_Cursor.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000003_Cursor.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000003_Cursor.SourceMultiplicity = models.MANY
	__LinkShape__000003_Cursor.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000003_Cursor.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000003_Cursor.X = 672.500000
	__LinkShape__000003_Cursor.Y = 381.000000
	__LinkShape__000003_Cursor.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000003_Cursor.StartRatio = 0.500000
	__LinkShape__000003_Cursor.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000003_Cursor.EndRatio = 0.500000
	__LinkShape__000003_Cursor.CornerOffsetRatio = 1.380000

	__LinkShape__000004_Form.Name = `Form`
	__LinkShape__000004_Form.IdentifierMeta = ref_models.AsSplitArea{}.Form
	__LinkShape__000004_Form.FieldTypeIdentifierMeta = ref_models.Form{}
	__LinkShape__000004_Form.FieldOffsetX = 0.000000
	__LinkShape__000004_Form.FieldOffsetY = 0.000000
	__LinkShape__000004_Form.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000004_Form.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000004_Form.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000004_Form.SourceMultiplicity = models.MANY
	__LinkShape__000004_Form.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000004_Form.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000004_Form.X = 540.000000
	__LinkShape__000004_Form.Y = 439.000000
	__LinkShape__000004_Form.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000004_Form.StartRatio = 0.500000
	__LinkShape__000004_Form.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000004_Form.EndRatio = 0.500000
	__LinkShape__000004_Form.CornerOffsetRatio = 1.380000

	__LinkShape__000005_Load.Name = `Load`
	__LinkShape__000005_Load.IdentifierMeta = ref_models.AsSplitArea{}.Load
	__LinkShape__000005_Load.FieldTypeIdentifierMeta = ref_models.Load{}
	__LinkShape__000005_Load.FieldOffsetX = 0.000000
	__LinkShape__000005_Load.FieldOffsetY = 0.000000
	__LinkShape__000005_Load.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000005_Load.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000005_Load.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000005_Load.SourceMultiplicity = models.MANY
	__LinkShape__000005_Load.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000005_Load.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000005_Load.X = 434.000000
	__LinkShape__000005_Load.Y = 444.500000
	__LinkShape__000005_Load.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000005_Load.StartRatio = 0.500000
	__LinkShape__000005_Load.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000005_Load.EndRatio = 0.500000
	__LinkShape__000005_Load.CornerOffsetRatio = 1.380000

	__LinkShape__000006_Markdown.Name = `Markdown`
	__LinkShape__000006_Markdown.IdentifierMeta = ref_models.AsSplitArea{}.Markdown
	__LinkShape__000006_Markdown.FieldTypeIdentifierMeta = ref_models.Markdown{}
	__LinkShape__000006_Markdown.FieldOffsetX = 0.000000
	__LinkShape__000006_Markdown.FieldOffsetY = 0.000000
	__LinkShape__000006_Markdown.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000006_Markdown.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000006_Markdown.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000006_Markdown.SourceMultiplicity = models.MANY
	__LinkShape__000006_Markdown.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000006_Markdown.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000006_Markdown.X = 631.500000
	__LinkShape__000006_Markdown.Y = 520.500000
	__LinkShape__000006_Markdown.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000006_Markdown.StartRatio = 0.500000
	__LinkShape__000006_Markdown.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000006_Markdown.EndRatio = 0.500000
	__LinkShape__000006_Markdown.CornerOffsetRatio = 1.380000

	__LinkShape__000007_Slider.Name = `Slider`
	__LinkShape__000007_Slider.IdentifierMeta = ref_models.AsSplitArea{}.Slider
	__LinkShape__000007_Slider.FieldTypeIdentifierMeta = ref_models.Slider{}
	__LinkShape__000007_Slider.FieldOffsetX = 0.000000
	__LinkShape__000007_Slider.FieldOffsetY = 0.000000
	__LinkShape__000007_Slider.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000007_Slider.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000007_Slider.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000007_Slider.SourceMultiplicity = models.MANY
	__LinkShape__000007_Slider.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000007_Slider.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000007_Slider.X = 666.500000
	__LinkShape__000007_Slider.Y = 572.000000
	__LinkShape__000007_Slider.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000007_Slider.StartRatio = 0.500000
	__LinkShape__000007_Slider.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000007_Slider.EndRatio = 0.500000
	__LinkShape__000007_Slider.CornerOffsetRatio = 1.380000

	__LinkShape__000008_Split.Name = `Split`
	__LinkShape__000008_Split.IdentifierMeta = ref_models.AsSplitArea{}.Split
	__LinkShape__000008_Split.FieldTypeIdentifierMeta = ref_models.Split{}
	__LinkShape__000008_Split.FieldOffsetX = 0.000000
	__LinkShape__000008_Split.FieldOffsetY = 0.000000
	__LinkShape__000008_Split.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000008_Split.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000008_Split.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000008_Split.SourceMultiplicity = models.MANY
	__LinkShape__000008_Split.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000008_Split.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000008_Split.X = 390.500000
	__LinkShape__000008_Split.Y = 351.500000
	__LinkShape__000008_Split.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000008_Split.StartRatio = 0.500000
	__LinkShape__000008_Split.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000008_Split.EndRatio = 0.500000
	__LinkShape__000008_Split.CornerOffsetRatio = 1.380000

	__LinkShape__000009_Svg.Name = `Svg`
	__LinkShape__000009_Svg.IdentifierMeta = ref_models.AsSplitArea{}.Svg
	__LinkShape__000009_Svg.FieldTypeIdentifierMeta = ref_models.Svg{}
	__LinkShape__000009_Svg.FieldOffsetX = 0.000000
	__LinkShape__000009_Svg.FieldOffsetY = 0.000000
	__LinkShape__000009_Svg.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000009_Svg.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000009_Svg.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000009_Svg.SourceMultiplicity = models.MANY
	__LinkShape__000009_Svg.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000009_Svg.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000009_Svg.X = 805.500000
	__LinkShape__000009_Svg.Y = 519.000000
	__LinkShape__000009_Svg.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000009_Svg.StartRatio = 0.500000
	__LinkShape__000009_Svg.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000009_Svg.EndRatio = 0.500000
	__LinkShape__000009_Svg.CornerOffsetRatio = 1.380000

	__LinkShape__000010_Table.Name = `Table`
	__LinkShape__000010_Table.IdentifierMeta = ref_models.AsSplitArea{}.Table
	__LinkShape__000010_Table.FieldTypeIdentifierMeta = ref_models.Table{}
	__LinkShape__000010_Table.FieldOffsetX = 0.000000
	__LinkShape__000010_Table.FieldOffsetY = 0.000000
	__LinkShape__000010_Table.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000010_Table.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000010_Table.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000010_Table.SourceMultiplicity = models.MANY
	__LinkShape__000010_Table.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000010_Table.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000010_Table.X = 695.000000
	__LinkShape__000010_Table.Y = 520.500000
	__LinkShape__000010_Table.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000010_Table.StartRatio = 0.500000
	__LinkShape__000010_Table.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000010_Table.EndRatio = 0.500000
	__LinkShape__000010_Table.CornerOffsetRatio = 1.380000

	__LinkShape__000011_Tone.Name = `Tone`
	__LinkShape__000011_Tone.IdentifierMeta = ref_models.AsSplitArea{}.Tone
	__LinkShape__000011_Tone.FieldTypeIdentifierMeta = ref_models.Tone{}
	__LinkShape__000011_Tone.FieldOffsetX = 0.000000
	__LinkShape__000011_Tone.FieldOffsetY = 0.000000
	__LinkShape__000011_Tone.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000011_Tone.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000011_Tone.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000011_Tone.SourceMultiplicity = models.MANY
	__LinkShape__000011_Tone.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000011_Tone.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000011_Tone.X = 853.500000
	__LinkShape__000011_Tone.Y = 456.500000
	__LinkShape__000011_Tone.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000011_Tone.StartRatio = 0.500000
	__LinkShape__000011_Tone.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000011_Tone.EndRatio = 0.500000
	__LinkShape__000011_Tone.CornerOffsetRatio = 1.380000

	__LinkShape__000012_Tree.Name = `Tree`
	__LinkShape__000012_Tree.IdentifierMeta = ref_models.AsSplitArea{}.Tree
	__LinkShape__000012_Tree.FieldTypeIdentifierMeta = ref_models.Tree{}
	__LinkShape__000012_Tree.FieldOffsetX = 0.000000
	__LinkShape__000012_Tree.FieldOffsetY = 0.000000
	__LinkShape__000012_Tree.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000012_Tree.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000012_Tree.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000012_Tree.SourceMultiplicity = models.MANY
	__LinkShape__000012_Tree.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000012_Tree.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000012_Tree.X = 894.000000
	__LinkShape__000012_Tree.Y = 562.000000
	__LinkShape__000012_Tree.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000012_Tree.StartRatio = 0.500000
	__LinkShape__000012_Tree.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000012_Tree.EndRatio = 0.500000
	__LinkShape__000012_Tree.CornerOffsetRatio = 1.380000

	__LinkShape__000013_Xlsx.Name = `Xlsx`
	__LinkShape__000013_Xlsx.IdentifierMeta = ref_models.AsSplitArea{}.Xlsx
	__LinkShape__000013_Xlsx.FieldTypeIdentifierMeta = ref_models.Xlsx{}
	__LinkShape__000013_Xlsx.FieldOffsetX = 0.000000
	__LinkShape__000013_Xlsx.FieldOffsetY = 0.000000
	__LinkShape__000013_Xlsx.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000013_Xlsx.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000013_Xlsx.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000013_Xlsx.SourceMultiplicity = models.MANY
	__LinkShape__000013_Xlsx.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000013_Xlsx.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000013_Xlsx.X = 848.500000
	__LinkShape__000013_Xlsx.Y = 656.000000
	__LinkShape__000013_Xlsx.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000013_Xlsx.StartRatio = 0.498709
	__LinkShape__000013_Xlsx.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000013_Xlsx.EndRatio = 0.500000
	__LinkShape__000013_Xlsx.CornerOffsetRatio = 1.382324

	__LinkShape__000014_NoName_yet.Name = `NoName yet`
	__LinkShape__000014_NoName_yet.FieldOffsetX = 0.000000
	__LinkShape__000014_NoName_yet.FieldOffsetY = 0.000000
	__LinkShape__000014_NoName_yet.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000014_NoName_yet.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000014_NoName_yet.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000014_NoName_yet.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000014_NoName_yet.X = 0.000000
	__LinkShape__000014_NoName_yet.Y = 0.000000
	__LinkShape__000014_NoName_yet.StartRatio = 0.000000
	__LinkShape__000014_NoName_yet.EndRatio = 0.000000
	__LinkShape__000014_NoName_yet.CornerOffsetRatio = 0.000000

	__LinkShape__000015_AsSplitAreas.Name = `AsSplitAreas`
	__LinkShape__000015_AsSplitAreas.IdentifierMeta = ref_models.AsSplit{}.AsSplitAreas
	__LinkShape__000015_AsSplitAreas.FieldTypeIdentifierMeta = ref_models.AsSplitArea{}
	__LinkShape__000015_AsSplitAreas.FieldOffsetX = 0.000000
	__LinkShape__000015_AsSplitAreas.FieldOffsetY = 0.000000
	__LinkShape__000015_AsSplitAreas.TargetMultiplicity = models.MANY
	__LinkShape__000015_AsSplitAreas.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000015_AsSplitAreas.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000015_AsSplitAreas.SourceMultiplicity = models.MANY
	__LinkShape__000015_AsSplitAreas.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000015_AsSplitAreas.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000015_AsSplitAreas.X = 396.500000
	__LinkShape__000015_AsSplitAreas.Y = 712.500000
	__LinkShape__000015_AsSplitAreas.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000015_AsSplitAreas.StartRatio = 0.786491
	__LinkShape__000015_AsSplitAreas.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000015_AsSplitAreas.EndRatio = 0.798991
	__LinkShape__000015_AsSplitAreas.CornerOffsetRatio = -0.623329

	__LinkShape__000016_RootAsSplitAreas.Name = `RootAsSplitAreas`
	__LinkShape__000016_RootAsSplitAreas.IdentifierMeta = ref_models.View{}.RootAsSplitAreas
	__LinkShape__000016_RootAsSplitAreas.FieldTypeIdentifierMeta = ref_models.AsSplitArea{}
	__LinkShape__000016_RootAsSplitAreas.FieldOffsetX = 0.000000
	__LinkShape__000016_RootAsSplitAreas.FieldOffsetY = 0.000000
	__LinkShape__000016_RootAsSplitAreas.TargetMultiplicity = models.MANY
	__LinkShape__000016_RootAsSplitAreas.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000016_RootAsSplitAreas.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000016_RootAsSplitAreas.SourceMultiplicity = models.MANY
	__LinkShape__000016_RootAsSplitAreas.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000016_RootAsSplitAreas.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000016_RootAsSplitAreas.X = 402.000000
	__LinkShape__000016_RootAsSplitAreas.Y = 395.000000
	__LinkShape__000016_RootAsSplitAreas.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000016_RootAsSplitAreas.StartRatio = 0.757324
	__LinkShape__000016_RootAsSplitAreas.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000016_RootAsSplitAreas.EndRatio = 0.844824
	__LinkShape__000016_RootAsSplitAreas.CornerOffsetRatio = 2.026534

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Split)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Xlsx)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_View)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_Tree)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000004_Default_Tone)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000005_Default_Title)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000006_Default_Table)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000007_Default_Svg)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000008_Default_Slider)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000009_Default_Markdown)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000010_Default_LogoOnTheRight)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000011_Default_LogoOnTheLeft)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000012_Default_Load)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000013_Default_Form)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000014_Default_FavIcon)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000015_Default_Cursor)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000016_Default_Button)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000017_Default_AsSplitArea)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000018_NoName_yet)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000019_Default_AsSplit)
	__Classdiagram__000000_Default.GongEnumShapes = append(__Classdiagram__000000_Default.GongEnumShapes, __GongEnumShape__000000_Default_Direction)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_12T00_16_22Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_06_12T00_16_22Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_12T00_16_22Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongEnumShape instances pointers
	__GongEnumShape__000000_Default_Direction.GongEnumValueShapes = append(__GongEnumShape__000000_Default_Direction.GongEnumValueShapes, __GongEnumValueShape__000001_Vertical)
	__GongEnumShape__000000_Default_Direction.GongEnumValueShapes = append(__GongEnumShape__000000_Default_Direction.GongEnumValueShapes, __GongEnumValueShape__000000_Horizontal)
	// setup of GongEnumValueShape instances pointers
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Split.AttributeShapes = append(__GongStructShape__000000_Default_Split.AttributeShapes, __AttributeShape__000000_StackName)
	__GongStructShape__000002_Default_View.AttributeShapes = append(__GongStructShape__000002_Default_View.AttributeShapes, __AttributeShape__000002_ShowViewName)
	__GongStructShape__000002_Default_View.AttributeShapes = append(__GongStructShape__000002_Default_View.AttributeShapes, __AttributeShape__000001_IsSelectedView)
	__GongStructShape__000002_Default_View.LinkShapes = append(__GongStructShape__000002_Default_View.LinkShapes, __LinkShape__000000_NoName_yet)
	__GongStructShape__000002_Default_View.LinkShapes = append(__GongStructShape__000002_Default_View.LinkShapes, __LinkShape__000016_RootAsSplitAreas)
	__GongStructShape__000010_Default_LogoOnTheRight.AttributeShapes = append(__GongStructShape__000010_Default_LogoOnTheRight.AttributeShapes, __AttributeShape__000005_Width)
	__GongStructShape__000010_Default_LogoOnTheRight.AttributeShapes = append(__GongStructShape__000010_Default_LogoOnTheRight.AttributeShapes, __AttributeShape__000004_Height)
	__GongStructShape__000010_Default_LogoOnTheRight.AttributeShapes = append(__GongStructShape__000010_Default_LogoOnTheRight.AttributeShapes, __AttributeShape__000003_SVG)
	__GongStructShape__000011_Default_LogoOnTheLeft.AttributeShapes = append(__GongStructShape__000011_Default_LogoOnTheLeft.AttributeShapes, __AttributeShape__000006_Width)
	__GongStructShape__000011_Default_LogoOnTheLeft.AttributeShapes = append(__GongStructShape__000011_Default_LogoOnTheLeft.AttributeShapes, __AttributeShape__000007_Height)
	__GongStructShape__000011_Default_LogoOnTheLeft.AttributeShapes = append(__GongStructShape__000011_Default_LogoOnTheLeft.AttributeShapes, __AttributeShape__000008_SVG)
	__GongStructShape__000017_Default_AsSplitArea.AttributeShapes = append(__GongStructShape__000017_Default_AsSplitArea.AttributeShapes, __AttributeShape__000009_ShowNameInHeader)
	__GongStructShape__000017_Default_AsSplitArea.AttributeShapes = append(__GongStructShape__000017_Default_AsSplitArea.AttributeShapes, __AttributeShape__000010_Size)
	__GongStructShape__000017_Default_AsSplitArea.AttributeShapes = append(__GongStructShape__000017_Default_AsSplitArea.AttributeShapes, __AttributeShape__000011_IsAny)
	__GongStructShape__000017_Default_AsSplitArea.AttributeShapes = append(__GongStructShape__000017_Default_AsSplitArea.AttributeShapes, __AttributeShape__000012_HasDiv)
	__GongStructShape__000017_Default_AsSplitArea.AttributeShapes = append(__GongStructShape__000017_Default_AsSplitArea.AttributeShapes, __AttributeShape__000013_DivStyle)
	__GongStructShape__000017_Default_AsSplitArea.LinkShapes = append(__GongStructShape__000017_Default_AsSplitArea.LinkShapes, __LinkShape__000001_AsSplit)
	__GongStructShape__000017_Default_AsSplitArea.LinkShapes = append(__GongStructShape__000017_Default_AsSplitArea.LinkShapes, __LinkShape__000002_Button)
	__GongStructShape__000017_Default_AsSplitArea.LinkShapes = append(__GongStructShape__000017_Default_AsSplitArea.LinkShapes, __LinkShape__000003_Cursor)
	__GongStructShape__000017_Default_AsSplitArea.LinkShapes = append(__GongStructShape__000017_Default_AsSplitArea.LinkShapes, __LinkShape__000004_Form)
	__GongStructShape__000017_Default_AsSplitArea.LinkShapes = append(__GongStructShape__000017_Default_AsSplitArea.LinkShapes, __LinkShape__000005_Load)
	__GongStructShape__000017_Default_AsSplitArea.LinkShapes = append(__GongStructShape__000017_Default_AsSplitArea.LinkShapes, __LinkShape__000006_Markdown)
	__GongStructShape__000017_Default_AsSplitArea.LinkShapes = append(__GongStructShape__000017_Default_AsSplitArea.LinkShapes, __LinkShape__000007_Slider)
	__GongStructShape__000017_Default_AsSplitArea.LinkShapes = append(__GongStructShape__000017_Default_AsSplitArea.LinkShapes, __LinkShape__000008_Split)
	__GongStructShape__000017_Default_AsSplitArea.LinkShapes = append(__GongStructShape__000017_Default_AsSplitArea.LinkShapes, __LinkShape__000009_Svg)
	__GongStructShape__000017_Default_AsSplitArea.LinkShapes = append(__GongStructShape__000017_Default_AsSplitArea.LinkShapes, __LinkShape__000010_Table)
	__GongStructShape__000017_Default_AsSplitArea.LinkShapes = append(__GongStructShape__000017_Default_AsSplitArea.LinkShapes, __LinkShape__000011_Tone)
	__GongStructShape__000017_Default_AsSplitArea.LinkShapes = append(__GongStructShape__000017_Default_AsSplitArea.LinkShapes, __LinkShape__000012_Tree)
	__GongStructShape__000017_Default_AsSplitArea.LinkShapes = append(__GongStructShape__000017_Default_AsSplitArea.LinkShapes, __LinkShape__000013_Xlsx)
	__GongStructShape__000018_NoName_yet.AttributeShapes = append(__GongStructShape__000018_NoName_yet.AttributeShapes, __AttributeShape__000014_NoName_yet)
	__GongStructShape__000018_NoName_yet.LinkShapes = append(__GongStructShape__000018_NoName_yet.LinkShapes, __LinkShape__000014_NoName_yet)
	__GongStructShape__000019_Default_AsSplit.LinkShapes = append(__GongStructShape__000019_Default_AsSplit.LinkShapes, __LinkShape__000015_AsSplitAreas)
	// setup of LinkShape instances pointers
}

