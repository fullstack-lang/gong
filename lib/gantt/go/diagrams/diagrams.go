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
	__AttributeShape__000001_ComputedStart := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000002_ComputedEnd := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000003_ComputedDuration := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000004_UseManualStartAndEndDates := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000005_ManualStart := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000006_ManualEnd := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000007_LaneHeight := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000008_RatioBarToLaneHeight := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000009_YTopMargin := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000010_XLeftText := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000011_TextHeight := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000012_XLeftLanes := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000013_XRightMargin := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000014_ArrowLengthToTheRightOfStartBar := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000015_ArrowTipLenght := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000016_TimeLine_Color := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000017_TimeLine_FillOpacity := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000018_TimeLine_Stroke := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000019_TimeLine_StrokeWidth := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000020_Group_Stroke := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000021_Group_StrokeWidth := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000022_Group_StrokeDashArray := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000023_DateYOffset := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000024_AlignOnStartEndOnYearStart := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000025_Order := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000026_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000027_Date := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000028_DisplayVerticalBar := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_19T06_00_19Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_Gantt := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_Group := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_Lane := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000003_Default_Milestone := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000004_Default_Arrow := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000005_Default_Bar := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_Lanes := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000001_Milestones := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000002_Groups := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000003_GroupLanes := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000004_Bars := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_Name.Name = `Name`

	//gong:ident [ref_models.Gantt.Name] comment added to overcome the problem with the comment map association
	__AttributeShape__000000_Name.Identifier = `ref_models.Gantt.Name`
	__AttributeShape__000000_Name.FieldTypeAsString = ``
	__AttributeShape__000000_Name.Structname = `Gantt`
	__AttributeShape__000000_Name.Fieldtypename = `string`

	__AttributeShape__000001_ComputedStart.Name = `ComputedStart`

	//gong:ident [ref_models.Gantt.ComputedStart] comment added to overcome the problem with the comment map association
	__AttributeShape__000001_ComputedStart.Identifier = `ref_models.Gantt.ComputedStart`
	__AttributeShape__000001_ComputedStart.FieldTypeAsString = ``
	__AttributeShape__000001_ComputedStart.Structname = `Gantt`
	__AttributeShape__000001_ComputedStart.Fieldtypename = `Time`

	__AttributeShape__000002_ComputedEnd.Name = `ComputedEnd`

	//gong:ident [ref_models.Gantt.ComputedEnd] comment added to overcome the problem with the comment map association
	__AttributeShape__000002_ComputedEnd.Identifier = `ref_models.Gantt.ComputedEnd`
	__AttributeShape__000002_ComputedEnd.FieldTypeAsString = ``
	__AttributeShape__000002_ComputedEnd.Structname = `Gantt`
	__AttributeShape__000002_ComputedEnd.Fieldtypename = `Time`

	__AttributeShape__000003_ComputedDuration.Name = `ComputedDuration`

	//gong:ident [ref_models.Gantt.ComputedDuration] comment added to overcome the problem with the comment map association
	__AttributeShape__000003_ComputedDuration.Identifier = `ref_models.Gantt.ComputedDuration`
	__AttributeShape__000003_ComputedDuration.FieldTypeAsString = ``
	__AttributeShape__000003_ComputedDuration.Structname = `Gantt`
	__AttributeShape__000003_ComputedDuration.Fieldtypename = `Duration`

	__AttributeShape__000004_UseManualStartAndEndDates.Name = `UseManualStartAndEndDates`

	//gong:ident [ref_models.Gantt.UseManualStartAndEndDates] comment added to overcome the problem with the comment map association
	__AttributeShape__000004_UseManualStartAndEndDates.Identifier = `ref_models.Gantt.UseManualStartAndEndDates`
	__AttributeShape__000004_UseManualStartAndEndDates.FieldTypeAsString = ``
	__AttributeShape__000004_UseManualStartAndEndDates.Structname = `Gantt`
	__AttributeShape__000004_UseManualStartAndEndDates.Fieldtypename = `bool`

	__AttributeShape__000005_ManualStart.Name = `ManualStart`

	//gong:ident [ref_models.Gantt.ManualStart] comment added to overcome the problem with the comment map association
	__AttributeShape__000005_ManualStart.Identifier = `ref_models.Gantt.ManualStart`
	__AttributeShape__000005_ManualStart.FieldTypeAsString = ``
	__AttributeShape__000005_ManualStart.Structname = `Gantt`
	__AttributeShape__000005_ManualStart.Fieldtypename = `Time`

	__AttributeShape__000006_ManualEnd.Name = `ManualEnd`

	//gong:ident [ref_models.Gantt.ManualEnd] comment added to overcome the problem with the comment map association
	__AttributeShape__000006_ManualEnd.Identifier = `ref_models.Gantt.ManualEnd`
	__AttributeShape__000006_ManualEnd.FieldTypeAsString = ``
	__AttributeShape__000006_ManualEnd.Structname = `Gantt`
	__AttributeShape__000006_ManualEnd.Fieldtypename = `Time`

	__AttributeShape__000007_LaneHeight.Name = `LaneHeight`

	//gong:ident [ref_models.Gantt.LaneHeight] comment added to overcome the problem with the comment map association
	__AttributeShape__000007_LaneHeight.Identifier = `ref_models.Gantt.LaneHeight`
	__AttributeShape__000007_LaneHeight.FieldTypeAsString = ``
	__AttributeShape__000007_LaneHeight.Structname = `Gantt`
	__AttributeShape__000007_LaneHeight.Fieldtypename = `float64`

	__AttributeShape__000008_RatioBarToLaneHeight.Name = `RatioBarToLaneHeight`

	//gong:ident [ref_models.Gantt.RatioBarToLaneHeight] comment added to overcome the problem with the comment map association
	__AttributeShape__000008_RatioBarToLaneHeight.Identifier = `ref_models.Gantt.RatioBarToLaneHeight`
	__AttributeShape__000008_RatioBarToLaneHeight.FieldTypeAsString = ``
	__AttributeShape__000008_RatioBarToLaneHeight.Structname = `Gantt`
	__AttributeShape__000008_RatioBarToLaneHeight.Fieldtypename = `float64`

	__AttributeShape__000009_YTopMargin.Name = `YTopMargin`

	//gong:ident [ref_models.Gantt.YTopMargin] comment added to overcome the problem with the comment map association
	__AttributeShape__000009_YTopMargin.Identifier = `ref_models.Gantt.YTopMargin`
	__AttributeShape__000009_YTopMargin.FieldTypeAsString = ``
	__AttributeShape__000009_YTopMargin.Structname = `Gantt`
	__AttributeShape__000009_YTopMargin.Fieldtypename = `float64`

	__AttributeShape__000010_XLeftText.Name = `XLeftText`

	//gong:ident [ref_models.Gantt.XLeftText] comment added to overcome the problem with the comment map association
	__AttributeShape__000010_XLeftText.Identifier = `ref_models.Gantt.XLeftText`
	__AttributeShape__000010_XLeftText.FieldTypeAsString = ``
	__AttributeShape__000010_XLeftText.Structname = `Gantt`
	__AttributeShape__000010_XLeftText.Fieldtypename = `float64`

	__AttributeShape__000011_TextHeight.Name = `TextHeight`

	//gong:ident [ref_models.Gantt.TextHeight] comment added to overcome the problem with the comment map association
	__AttributeShape__000011_TextHeight.Identifier = `ref_models.Gantt.TextHeight`
	__AttributeShape__000011_TextHeight.FieldTypeAsString = ``
	__AttributeShape__000011_TextHeight.Structname = `Gantt`
	__AttributeShape__000011_TextHeight.Fieldtypename = `float64`

	__AttributeShape__000012_XLeftLanes.Name = `XLeftLanes`

	//gong:ident [ref_models.Gantt.XLeftLanes] comment added to overcome the problem with the comment map association
	__AttributeShape__000012_XLeftLanes.Identifier = `ref_models.Gantt.XLeftLanes`
	__AttributeShape__000012_XLeftLanes.FieldTypeAsString = ``
	__AttributeShape__000012_XLeftLanes.Structname = `Gantt`
	__AttributeShape__000012_XLeftLanes.Fieldtypename = `float64`

	__AttributeShape__000013_XRightMargin.Name = `XRightMargin`

	//gong:ident [ref_models.Gantt.XRightMargin] comment added to overcome the problem with the comment map association
	__AttributeShape__000013_XRightMargin.Identifier = `ref_models.Gantt.XRightMargin`
	__AttributeShape__000013_XRightMargin.FieldTypeAsString = ``
	__AttributeShape__000013_XRightMargin.Structname = `Gantt`
	__AttributeShape__000013_XRightMargin.Fieldtypename = `float64`

	__AttributeShape__000014_ArrowLengthToTheRightOfStartBar.Name = `ArrowLengthToTheRightOfStartBar`

	//gong:ident [ref_models.Gantt.ArrowLengthToTheRightOfStartBar] comment added to overcome the problem with the comment map association
	__AttributeShape__000014_ArrowLengthToTheRightOfStartBar.Identifier = `ref_models.Gantt.ArrowLengthToTheRightOfStartBar`
	__AttributeShape__000014_ArrowLengthToTheRightOfStartBar.FieldTypeAsString = ``
	__AttributeShape__000014_ArrowLengthToTheRightOfStartBar.Structname = `Gantt`
	__AttributeShape__000014_ArrowLengthToTheRightOfStartBar.Fieldtypename = `float64`

	__AttributeShape__000015_ArrowTipLenght.Name = `ArrowTipLenght`

	//gong:ident [ref_models.Gantt.ArrowTipLenght] comment added to overcome the problem with the comment map association
	__AttributeShape__000015_ArrowTipLenght.Identifier = `ref_models.Gantt.ArrowTipLenght`
	__AttributeShape__000015_ArrowTipLenght.FieldTypeAsString = ``
	__AttributeShape__000015_ArrowTipLenght.Structname = `Gantt`
	__AttributeShape__000015_ArrowTipLenght.Fieldtypename = `float64`

	__AttributeShape__000016_TimeLine_Color.Name = `TimeLine_Color`

	//gong:ident [ref_models.Gantt.TimeLine_Color] comment added to overcome the problem with the comment map association
	__AttributeShape__000016_TimeLine_Color.Identifier = `ref_models.Gantt.TimeLine_Color`
	__AttributeShape__000016_TimeLine_Color.FieldTypeAsString = ``
	__AttributeShape__000016_TimeLine_Color.Structname = `Gantt`
	__AttributeShape__000016_TimeLine_Color.Fieldtypename = `string`

	__AttributeShape__000017_TimeLine_FillOpacity.Name = `TimeLine_FillOpacity`

	//gong:ident [ref_models.Gantt.TimeLine_FillOpacity] comment added to overcome the problem with the comment map association
	__AttributeShape__000017_TimeLine_FillOpacity.Identifier = `ref_models.Gantt.TimeLine_FillOpacity`
	__AttributeShape__000017_TimeLine_FillOpacity.FieldTypeAsString = ``
	__AttributeShape__000017_TimeLine_FillOpacity.Structname = `Gantt`
	__AttributeShape__000017_TimeLine_FillOpacity.Fieldtypename = `float64`

	__AttributeShape__000018_TimeLine_Stroke.Name = `TimeLine_Stroke`

	//gong:ident [ref_models.Gantt.TimeLine_Stroke] comment added to overcome the problem with the comment map association
	__AttributeShape__000018_TimeLine_Stroke.Identifier = `ref_models.Gantt.TimeLine_Stroke`
	__AttributeShape__000018_TimeLine_Stroke.FieldTypeAsString = ``
	__AttributeShape__000018_TimeLine_Stroke.Structname = `Gantt`
	__AttributeShape__000018_TimeLine_Stroke.Fieldtypename = `string`

	__AttributeShape__000019_TimeLine_StrokeWidth.Name = `TimeLine_StrokeWidth`

	//gong:ident [ref_models.Gantt.TimeLine_StrokeWidth] comment added to overcome the problem with the comment map association
	__AttributeShape__000019_TimeLine_StrokeWidth.Identifier = `ref_models.Gantt.TimeLine_StrokeWidth`
	__AttributeShape__000019_TimeLine_StrokeWidth.FieldTypeAsString = ``
	__AttributeShape__000019_TimeLine_StrokeWidth.Structname = `Gantt`
	__AttributeShape__000019_TimeLine_StrokeWidth.Fieldtypename = `float64`

	__AttributeShape__000020_Group_Stroke.Name = `Group_Stroke`

	//gong:ident [ref_models.Gantt.Group_Stroke] comment added to overcome the problem with the comment map association
	__AttributeShape__000020_Group_Stroke.Identifier = `ref_models.Gantt.Group_Stroke`
	__AttributeShape__000020_Group_Stroke.FieldTypeAsString = ``
	__AttributeShape__000020_Group_Stroke.Structname = `Gantt`
	__AttributeShape__000020_Group_Stroke.Fieldtypename = `string`

	__AttributeShape__000021_Group_StrokeWidth.Name = `Group_StrokeWidth`

	//gong:ident [ref_models.Gantt.Group_StrokeWidth] comment added to overcome the problem with the comment map association
	__AttributeShape__000021_Group_StrokeWidth.Identifier = `ref_models.Gantt.Group_StrokeWidth`
	__AttributeShape__000021_Group_StrokeWidth.FieldTypeAsString = ``
	__AttributeShape__000021_Group_StrokeWidth.Structname = `Gantt`
	__AttributeShape__000021_Group_StrokeWidth.Fieldtypename = `float64`

	__AttributeShape__000022_Group_StrokeDashArray.Name = `Group_StrokeDashArray`

	//gong:ident [ref_models.Gantt.Group_StrokeDashArray] comment added to overcome the problem with the comment map association
	__AttributeShape__000022_Group_StrokeDashArray.Identifier = `ref_models.Gantt.Group_StrokeDashArray`
	__AttributeShape__000022_Group_StrokeDashArray.FieldTypeAsString = ``
	__AttributeShape__000022_Group_StrokeDashArray.Structname = `Gantt`
	__AttributeShape__000022_Group_StrokeDashArray.Fieldtypename = `string`

	__AttributeShape__000023_DateYOffset.Name = `DateYOffset`

	//gong:ident [ref_models.Gantt.DateYOffset] comment added to overcome the problem with the comment map association
	__AttributeShape__000023_DateYOffset.Identifier = `ref_models.Gantt.DateYOffset`
	__AttributeShape__000023_DateYOffset.FieldTypeAsString = ``
	__AttributeShape__000023_DateYOffset.Structname = `Gantt`
	__AttributeShape__000023_DateYOffset.Fieldtypename = `float64`

	__AttributeShape__000024_AlignOnStartEndOnYearStart.Name = `AlignOnStartEndOnYearStart`

	//gong:ident [ref_models.Gantt.AlignOnStartEndOnYearStart] comment added to overcome the problem with the comment map association
	__AttributeShape__000024_AlignOnStartEndOnYearStart.Identifier = `ref_models.Gantt.AlignOnStartEndOnYearStart`
	__AttributeShape__000024_AlignOnStartEndOnYearStart.FieldTypeAsString = ``
	__AttributeShape__000024_AlignOnStartEndOnYearStart.Structname = `Gantt`
	__AttributeShape__000024_AlignOnStartEndOnYearStart.Fieldtypename = `bool`

	__AttributeShape__000025_Order.Name = `Order`

	//gong:ident [ref_models.Lane.Order] comment added to overcome the problem with the comment map association
	__AttributeShape__000025_Order.Identifier = `ref_models.Lane.Order`
	__AttributeShape__000025_Order.FieldTypeAsString = ``
	__AttributeShape__000025_Order.Structname = `Lane`
	__AttributeShape__000025_Order.Fieldtypename = `int`

	__AttributeShape__000026_Name.Name = `Name`

	//gong:ident [ref_models.Milestone.Name] comment added to overcome the problem with the comment map association
	__AttributeShape__000026_Name.Identifier = `ref_models.Milestone.Name`
	__AttributeShape__000026_Name.FieldTypeAsString = ``
	__AttributeShape__000026_Name.Structname = `Milestone`
	__AttributeShape__000026_Name.Fieldtypename = `string`

	__AttributeShape__000027_Date.Name = `Date`

	//gong:ident [ref_models.Milestone.Date] comment added to overcome the problem with the comment map association
	__AttributeShape__000027_Date.Identifier = `ref_models.Milestone.Date`
	__AttributeShape__000027_Date.FieldTypeAsString = ``
	__AttributeShape__000027_Date.Structname = `Milestone`
	__AttributeShape__000027_Date.Fieldtypename = `Time`

	__AttributeShape__000028_DisplayVerticalBar.Name = `DisplayVerticalBar`

	//gong:ident [ref_models.Milestone.DisplayVerticalBar] comment added to overcome the problem with the comment map association
	__AttributeShape__000028_DisplayVerticalBar.Identifier = `ref_models.Milestone.DisplayVerticalBar`
	__AttributeShape__000028_DisplayVerticalBar.FieldTypeAsString = ``
	__AttributeShape__000028_DisplayVerticalBar.Structname = `Milestone`
	__AttributeShape__000028_DisplayVerticalBar.Fieldtypename = `bool`

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = `[false,false,false,true,true,true,true]`
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_19T06_00_19Z.Name = `Diagram Package created the 2025-05-19T06:00:19Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_19T06_00_19Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_19T06_00_19Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_19T06_00_19Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_Gantt.Name = `Default-Gantt`
	__GongStructShape__000000_Default_Gantt.X = 13.000000
	__GongStructShape__000000_Default_Gantt.Y = 89.000000

	//gong:ident [ref_models.Gantt] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Gantt.Identifier = `ref_models.Gantt`
	__GongStructShape__000000_Default_Gantt.ShowNbInstances = false
	__GongStructShape__000000_Default_Gantt.NbInstances = 0
	__GongStructShape__000000_Default_Gantt.Width = 345.000000
	__GongStructShape__000000_Default_Gantt.Height = 563.000000
	__GongStructShape__000000_Default_Gantt.IsSelected = false

	__GongStructShape__000001_Default_Group.Name = `Default-Group`
	__GongStructShape__000001_Default_Group.X = 592.000000
	__GongStructShape__000001_Default_Group.Y = 205.000000

	//gong:ident [ref_models.Group] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_Group.Identifier = `ref_models.Group`
	__GongStructShape__000001_Default_Group.ShowNbInstances = false
	__GongStructShape__000001_Default_Group.NbInstances = 0
	__GongStructShape__000001_Default_Group.Width = 240.000000
	__GongStructShape__000001_Default_Group.Height = 63.000000
	__GongStructShape__000001_Default_Group.IsSelected = false

	__GongStructShape__000002_Default_Lane.Name = `Default-Lane`
	__GongStructShape__000002_Default_Lane.X = 594.000000
	__GongStructShape__000002_Default_Lane.Y = 391.000000

	//gong:ident [ref_models.Lane] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_Lane.Identifier = `ref_models.Lane`
	__GongStructShape__000002_Default_Lane.ShowNbInstances = false
	__GongStructShape__000002_Default_Lane.NbInstances = 0
	__GongStructShape__000002_Default_Lane.Width = 240.000000
	__GongStructShape__000002_Default_Lane.Height = 83.000000
	__GongStructShape__000002_Default_Lane.IsSelected = false

	__GongStructShape__000003_Default_Milestone.Name = `Default-Milestone`
	__GongStructShape__000003_Default_Milestone.X = 588.000000
	__GongStructShape__000003_Default_Milestone.Y = 57.000000

	//gong:ident [ref_models.Milestone] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_Default_Milestone.Identifier = `ref_models.Milestone`
	__GongStructShape__000003_Default_Milestone.ShowNbInstances = false
	__GongStructShape__000003_Default_Milestone.NbInstances = 0
	__GongStructShape__000003_Default_Milestone.Width = 240.000000
	__GongStructShape__000003_Default_Milestone.Height = 123.000000
	__GongStructShape__000003_Default_Milestone.IsSelected = false

	__GongStructShape__000004_Default_Arrow.Name = `Default-Arrow`
	__GongStructShape__000004_Default_Arrow.X = 601.000000
	__GongStructShape__000004_Default_Arrow.Y = 522.999969

	//gong:ident [ref_models.Arrow] comment added to overcome the problem with the comment map association
	__GongStructShape__000004_Default_Arrow.Identifier = `ref_models.Arrow`
	__GongStructShape__000004_Default_Arrow.ShowNbInstances = false
	__GongStructShape__000004_Default_Arrow.NbInstances = 0
	__GongStructShape__000004_Default_Arrow.Width = 240.000000
	__GongStructShape__000004_Default_Arrow.Height = 63.000000
	__GongStructShape__000004_Default_Arrow.IsSelected = false

	__GongStructShape__000005_Default_Bar.Name = `Default-Bar`
	__GongStructShape__000005_Default_Bar.X = 1105.000000
	__GongStructShape__000005_Default_Bar.Y = 402.000000

	//gong:ident [ref_models.Bar] comment added to overcome the problem with the comment map association
	__GongStructShape__000005_Default_Bar.Identifier = `ref_models.Bar`
	__GongStructShape__000005_Default_Bar.ShowNbInstances = false
	__GongStructShape__000005_Default_Bar.NbInstances = 0
	__GongStructShape__000005_Default_Bar.Width = 240.000000
	__GongStructShape__000005_Default_Bar.Height = 63.000000
	__GongStructShape__000005_Default_Bar.IsSelected = false

	__LinkShape__000000_Lanes.Name = `Lanes`

	//gong:ident [ref_models.Gantt.Lanes] comment added to overcome the problem with the comment map association
	__LinkShape__000000_Lanes.Identifier = `ref_models.Gantt.Lanes`

	//gong:ident [ref_models.Lane] comment added to overcome the problem with the comment map association
	__LinkShape__000000_Lanes.Fieldtypename = `ref_models.Lane`
	__LinkShape__000000_Lanes.FieldOffsetX = 0.000000
	__LinkShape__000000_Lanes.FieldOffsetY = 0.000000
	__LinkShape__000000_Lanes.TargetMultiplicity = models.MANY
	__LinkShape__000000_Lanes.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Lanes.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Lanes.SourceMultiplicity = models.MANY
	__LinkShape__000000_Lanes.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Lanes.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Lanes.X = 821.000000
	__LinkShape__000000_Lanes.Y = 521.500000
	__LinkShape__000000_Lanes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Lanes.StartRatio = 0.589420
	__LinkShape__000000_Lanes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Lanes.EndRatio = 0.500000
	__LinkShape__000000_Lanes.CornerOffsetRatio = 1.380000

	__LinkShape__000001_Milestones.Name = `Milestones`

	//gong:ident [ref_models.Gantt.Milestones] comment added to overcome the problem with the comment map association
	__LinkShape__000001_Milestones.Identifier = `ref_models.Gantt.Milestones`

	//gong:ident [ref_models.Milestone] comment added to overcome the problem with the comment map association
	__LinkShape__000001_Milestones.Fieldtypename = `ref_models.Milestone`
	__LinkShape__000001_Milestones.FieldOffsetX = 0.000000
	__LinkShape__000001_Milestones.FieldOffsetY = 0.000000
	__LinkShape__000001_Milestones.TargetMultiplicity = models.MANY
	__LinkShape__000001_Milestones.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Milestones.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Milestones.SourceMultiplicity = models.MANY
	__LinkShape__000001_Milestones.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Milestones.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Milestones.X = 819.500000
	__LinkShape__000001_Milestones.Y = 366.500000
	__LinkShape__000001_Milestones.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Milestones.StartRatio = 0.047680
	__LinkShape__000001_Milestones.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Milestones.EndRatio = 0.500000
	__LinkShape__000001_Milestones.CornerOffsetRatio = 1.380000

	__LinkShape__000002_Groups.Name = `Groups`

	//gong:ident [ref_models.Gantt.Groups] comment added to overcome the problem with the comment map association
	__LinkShape__000002_Groups.Identifier = `ref_models.Gantt.Groups`

	//gong:ident [ref_models.Group] comment added to overcome the problem with the comment map association
	__LinkShape__000002_Groups.Fieldtypename = `ref_models.Group`
	__LinkShape__000002_Groups.FieldOffsetX = 0.000000
	__LinkShape__000002_Groups.FieldOffsetY = 0.000000
	__LinkShape__000002_Groups.TargetMultiplicity = models.MANY
	__LinkShape__000002_Groups.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000002_Groups.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000002_Groups.SourceMultiplicity = models.MANY
	__LinkShape__000002_Groups.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000002_Groups.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000002_Groups.X = 820.000000
	__LinkShape__000002_Groups.Y = 447.000000
	__LinkShape__000002_Groups.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_Groups.StartRatio = 0.260824
	__LinkShape__000002_Groups.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_Groups.EndRatio = 0.500000
	__LinkShape__000002_Groups.CornerOffsetRatio = 1.380000

	__LinkShape__000003_GroupLanes.Name = `GroupLanes`

	//gong:ident [ref_models.Group.GroupLanes] comment added to overcome the problem with the comment map association
	__LinkShape__000003_GroupLanes.Identifier = `ref_models.Group.GroupLanes`

	//gong:ident [ref_models.Lane] comment added to overcome the problem with the comment map association
	__LinkShape__000003_GroupLanes.Fieldtypename = `ref_models.Lane`
	__LinkShape__000003_GroupLanes.FieldOffsetX = 0.000000
	__LinkShape__000003_GroupLanes.FieldOffsetY = 0.000000
	__LinkShape__000003_GroupLanes.TargetMultiplicity = models.MANY
	__LinkShape__000003_GroupLanes.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000003_GroupLanes.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000003_GroupLanes.SourceMultiplicity = models.MANY
	__LinkShape__000003_GroupLanes.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000003_GroupLanes.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000003_GroupLanes.X = 953.000000
	__LinkShape__000003_GroupLanes.Y = 329.500000
	__LinkShape__000003_GroupLanes.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000003_GroupLanes.StartRatio = 0.814714
	__LinkShape__000003_GroupLanes.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000003_GroupLanes.EndRatio = 0.806380
	__LinkShape__000003_GroupLanes.CornerOffsetRatio = 1.140376

	__LinkShape__000004_Bars.Name = `Bars`

	//gong:ident [ref_models.Lane.Bars] comment added to overcome the problem with the comment map association
	__LinkShape__000004_Bars.Identifier = `ref_models.Lane.Bars`

	//gong:ident [ref_models.Bar] comment added to overcome the problem with the comment map association
	__LinkShape__000004_Bars.Fieldtypename = `ref_models.Bar`
	__LinkShape__000004_Bars.FieldOffsetX = 0.000000
	__LinkShape__000004_Bars.FieldOffsetY = 0.000000
	__LinkShape__000004_Bars.TargetMultiplicity = models.MANY
	__LinkShape__000004_Bars.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000004_Bars.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000004_Bars.SourceMultiplicity = models.MANY
	__LinkShape__000004_Bars.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000004_Bars.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000004_Bars.X = 1185.000000
	__LinkShape__000004_Bars.Y = 423.500000
	__LinkShape__000004_Bars.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000004_Bars.StartRatio = 0.711805
	__LinkShape__000004_Bars.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000004_Bars.EndRatio = 0.775297
	__LinkShape__000004_Bars.CornerOffsetRatio = 1.380000

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Gantt)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Group)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Lane)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_Milestone)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000004_Default_Arrow)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000005_Default_Bar)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_19T06_00_19Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_19T06_00_19Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_19T06_00_19Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000000_Name)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000001_ComputedStart)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000002_ComputedEnd)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000003_ComputedDuration)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000004_UseManualStartAndEndDates)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000005_ManualStart)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000006_ManualEnd)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000007_LaneHeight)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000008_RatioBarToLaneHeight)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000009_YTopMargin)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000010_XLeftText)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000011_TextHeight)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000012_XLeftLanes)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000013_XRightMargin)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000014_ArrowLengthToTheRightOfStartBar)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000015_ArrowTipLenght)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000016_TimeLine_Color)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000017_TimeLine_FillOpacity)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000018_TimeLine_Stroke)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000019_TimeLine_StrokeWidth)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000020_Group_Stroke)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000021_Group_StrokeWidth)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000022_Group_StrokeDashArray)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000023_DateYOffset)
	__GongStructShape__000000_Default_Gantt.AttributeShapes = append(__GongStructShape__000000_Default_Gantt.AttributeShapes, __AttributeShape__000024_AlignOnStartEndOnYearStart)
	__GongStructShape__000000_Default_Gantt.LinkShapes = append(__GongStructShape__000000_Default_Gantt.LinkShapes, __LinkShape__000000_Lanes)
	__GongStructShape__000000_Default_Gantt.LinkShapes = append(__GongStructShape__000000_Default_Gantt.LinkShapes, __LinkShape__000001_Milestones)
	__GongStructShape__000000_Default_Gantt.LinkShapes = append(__GongStructShape__000000_Default_Gantt.LinkShapes, __LinkShape__000002_Groups)
	__GongStructShape__000001_Default_Group.LinkShapes = append(__GongStructShape__000001_Default_Group.LinkShapes, __LinkShape__000003_GroupLanes)
	__GongStructShape__000002_Default_Lane.AttributeShapes = append(__GongStructShape__000002_Default_Lane.AttributeShapes, __AttributeShape__000025_Order)
	__GongStructShape__000002_Default_Lane.LinkShapes = append(__GongStructShape__000002_Default_Lane.LinkShapes, __LinkShape__000004_Bars)
	__GongStructShape__000003_Default_Milestone.AttributeShapes = append(__GongStructShape__000003_Default_Milestone.AttributeShapes, __AttributeShape__000026_Name)
	__GongStructShape__000003_Default_Milestone.AttributeShapes = append(__GongStructShape__000003_Default_Milestone.AttributeShapes, __AttributeShape__000027_Date)
	__GongStructShape__000003_Default_Milestone.AttributeShapes = append(__GongStructShape__000003_Default_Milestone.AttributeShapes, __AttributeShape__000028_DisplayVerticalBar)
	// setup of LinkShape instances pointers
}
