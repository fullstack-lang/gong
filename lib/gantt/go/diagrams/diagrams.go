package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/gantt/go/models"
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
	__AttributeShape__000025_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000026_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000027_Order := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000028_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000029_Date := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000030_DisplayVerticalBar := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000031_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000032_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000033_OptionnalColor := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000034_OptionnalStroke := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000035_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000036_Start := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000037_End := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000038_ComputedDuration := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000039_OptionnalColor := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000040_OptionnalStroke := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T20_48_55Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_Milestone := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_LaneUse := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_Lane := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000003_Default_Group := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000004_Default_Gantt := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000005_Default_Bar := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000006_Default_Arrow := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_Lanes := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000001_Milestones := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000002_Groups := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000003_Arrows := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000004_GroupLanes := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000005_Bars := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000006_LanesToDisplayMilestoneUse := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000007_Lane := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000008_From := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000009_To := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_Name.Name = `Name`

	//gong:ident [ref_models.Gantt.Name] comment added to overcome the problem with the comment map association
	__AttributeShape__000000_Name.Identifier = `ref_models.Gantt.Name`
	__AttributeShape__000000_Name.IdentifierMeta = ref_models.Gantt{}.Name
	__AttributeShape__000000_Name.FieldTypeAsString = ``
	__AttributeShape__000000_Name.Structname = `Gantt`
	__AttributeShape__000000_Name.Fieldtypename = `string`

	__AttributeShape__000001_ComputedStart.Name = `ComputedStart`

	//gong:ident [ref_models.Gantt.ComputedStart] comment added to overcome the problem with the comment map association
	__AttributeShape__000001_ComputedStart.Identifier = `ref_models.Gantt.ComputedStart`
	__AttributeShape__000001_ComputedStart.IdentifierMeta = ref_models.Gantt{}.ComputedStart
	__AttributeShape__000001_ComputedStart.FieldTypeAsString = ``
	__AttributeShape__000001_ComputedStart.Structname = `Gantt`
	__AttributeShape__000001_ComputedStart.Fieldtypename = `Time`

	__AttributeShape__000002_ComputedEnd.Name = `ComputedEnd`

	//gong:ident [ref_models.Gantt.ComputedEnd] comment added to overcome the problem with the comment map association
	__AttributeShape__000002_ComputedEnd.Identifier = `ref_models.Gantt.ComputedEnd`
	__AttributeShape__000002_ComputedEnd.IdentifierMeta = ref_models.Gantt{}.ComputedEnd
	__AttributeShape__000002_ComputedEnd.FieldTypeAsString = ``
	__AttributeShape__000002_ComputedEnd.Structname = `Gantt`
	__AttributeShape__000002_ComputedEnd.Fieldtypename = `Time`

	__AttributeShape__000003_ComputedDuration.Name = `ComputedDuration`

	//gong:ident [ref_models.Gantt.ComputedDuration] comment added to overcome the problem with the comment map association
	__AttributeShape__000003_ComputedDuration.Identifier = `ref_models.Gantt.ComputedDuration`
	__AttributeShape__000003_ComputedDuration.IdentifierMeta = ref_models.Gantt{}.ComputedDuration
	__AttributeShape__000003_ComputedDuration.FieldTypeAsString = ``
	__AttributeShape__000003_ComputedDuration.Structname = `Gantt`
	__AttributeShape__000003_ComputedDuration.Fieldtypename = `Duration`

	__AttributeShape__000004_UseManualStartAndEndDates.Name = `UseManualStartAndEndDates`

	//gong:ident [ref_models.Gantt.UseManualStartAndEndDates] comment added to overcome the problem with the comment map association
	__AttributeShape__000004_UseManualStartAndEndDates.Identifier = `ref_models.Gantt.UseManualStartAndEndDates`
	__AttributeShape__000004_UseManualStartAndEndDates.IdentifierMeta = ref_models.Gantt{}.UseManualStartAndEndDates
	__AttributeShape__000004_UseManualStartAndEndDates.FieldTypeAsString = ``
	__AttributeShape__000004_UseManualStartAndEndDates.Structname = `Gantt`
	__AttributeShape__000004_UseManualStartAndEndDates.Fieldtypename = `bool`

	__AttributeShape__000005_ManualStart.Name = `ManualStart`

	//gong:ident [ref_models.Gantt.ManualStart] comment added to overcome the problem with the comment map association
	__AttributeShape__000005_ManualStart.Identifier = `ref_models.Gantt.ManualStart`
	__AttributeShape__000005_ManualStart.IdentifierMeta = ref_models.Gantt{}.ManualStart
	__AttributeShape__000005_ManualStart.FieldTypeAsString = ``
	__AttributeShape__000005_ManualStart.Structname = `Gantt`
	__AttributeShape__000005_ManualStart.Fieldtypename = `Time`

	__AttributeShape__000006_ManualEnd.Name = `ManualEnd`

	//gong:ident [ref_models.Gantt.ManualEnd] comment added to overcome the problem with the comment map association
	__AttributeShape__000006_ManualEnd.Identifier = `ref_models.Gantt.ManualEnd`
	__AttributeShape__000006_ManualEnd.IdentifierMeta = ref_models.Gantt{}.ManualEnd
	__AttributeShape__000006_ManualEnd.FieldTypeAsString = ``
	__AttributeShape__000006_ManualEnd.Structname = `Gantt`
	__AttributeShape__000006_ManualEnd.Fieldtypename = `Time`

	__AttributeShape__000007_LaneHeight.Name = `LaneHeight`

	//gong:ident [ref_models.Gantt.LaneHeight] comment added to overcome the problem with the comment map association
	__AttributeShape__000007_LaneHeight.Identifier = `ref_models.Gantt.LaneHeight`
	__AttributeShape__000007_LaneHeight.IdentifierMeta = ref_models.Gantt{}.LaneHeight
	__AttributeShape__000007_LaneHeight.FieldTypeAsString = ``
	__AttributeShape__000007_LaneHeight.Structname = `Gantt`
	__AttributeShape__000007_LaneHeight.Fieldtypename = `float64`

	__AttributeShape__000008_RatioBarToLaneHeight.Name = `RatioBarToLaneHeight`

	//gong:ident [ref_models.Gantt.RatioBarToLaneHeight] comment added to overcome the problem with the comment map association
	__AttributeShape__000008_RatioBarToLaneHeight.Identifier = `ref_models.Gantt.RatioBarToLaneHeight`
	__AttributeShape__000008_RatioBarToLaneHeight.IdentifierMeta = ref_models.Gantt{}.RatioBarToLaneHeight
	__AttributeShape__000008_RatioBarToLaneHeight.FieldTypeAsString = ``
	__AttributeShape__000008_RatioBarToLaneHeight.Structname = `Gantt`
	__AttributeShape__000008_RatioBarToLaneHeight.Fieldtypename = `float64`

	__AttributeShape__000009_YTopMargin.Name = `YTopMargin`

	//gong:ident [ref_models.Gantt.YTopMargin] comment added to overcome the problem with the comment map association
	__AttributeShape__000009_YTopMargin.Identifier = `ref_models.Gantt.YTopMargin`
	__AttributeShape__000009_YTopMargin.IdentifierMeta = ref_models.Gantt{}.YTopMargin
	__AttributeShape__000009_YTopMargin.FieldTypeAsString = ``
	__AttributeShape__000009_YTopMargin.Structname = `Gantt`
	__AttributeShape__000009_YTopMargin.Fieldtypename = `float64`

	__AttributeShape__000010_XLeftText.Name = `XLeftText`

	//gong:ident [ref_models.Gantt.XLeftText] comment added to overcome the problem with the comment map association
	__AttributeShape__000010_XLeftText.Identifier = `ref_models.Gantt.XLeftText`
	__AttributeShape__000010_XLeftText.IdentifierMeta = ref_models.Gantt{}.XLeftText
	__AttributeShape__000010_XLeftText.FieldTypeAsString = ``
	__AttributeShape__000010_XLeftText.Structname = `Gantt`
	__AttributeShape__000010_XLeftText.Fieldtypename = `float64`

	__AttributeShape__000011_TextHeight.Name = `TextHeight`

	//gong:ident [ref_models.Gantt.TextHeight] comment added to overcome the problem with the comment map association
	__AttributeShape__000011_TextHeight.Identifier = `ref_models.Gantt.TextHeight`
	__AttributeShape__000011_TextHeight.IdentifierMeta = ref_models.Gantt{}.TextHeight
	__AttributeShape__000011_TextHeight.FieldTypeAsString = ``
	__AttributeShape__000011_TextHeight.Structname = `Gantt`
	__AttributeShape__000011_TextHeight.Fieldtypename = `float64`

	__AttributeShape__000012_XLeftLanes.Name = `XLeftLanes`

	//gong:ident [ref_models.Gantt.XLeftLanes] comment added to overcome the problem with the comment map association
	__AttributeShape__000012_XLeftLanes.Identifier = `ref_models.Gantt.XLeftLanes`
	__AttributeShape__000012_XLeftLanes.IdentifierMeta = ref_models.Gantt{}.XLeftLanes
	__AttributeShape__000012_XLeftLanes.FieldTypeAsString = ``
	__AttributeShape__000012_XLeftLanes.Structname = `Gantt`
	__AttributeShape__000012_XLeftLanes.Fieldtypename = `float64`

	__AttributeShape__000013_XRightMargin.Name = `XRightMargin`

	//gong:ident [ref_models.Gantt.XRightMargin] comment added to overcome the problem with the comment map association
	__AttributeShape__000013_XRightMargin.Identifier = `ref_models.Gantt.XRightMargin`
	__AttributeShape__000013_XRightMargin.IdentifierMeta = ref_models.Gantt{}.XRightMargin
	__AttributeShape__000013_XRightMargin.FieldTypeAsString = ``
	__AttributeShape__000013_XRightMargin.Structname = `Gantt`
	__AttributeShape__000013_XRightMargin.Fieldtypename = `float64`

	__AttributeShape__000014_ArrowLengthToTheRightOfStartBar.Name = `ArrowLengthToTheRightOfStartBar`

	//gong:ident [ref_models.Gantt.ArrowLengthToTheRightOfStartBar] comment added to overcome the problem with the comment map association
	__AttributeShape__000014_ArrowLengthToTheRightOfStartBar.Identifier = `ref_models.Gantt.ArrowLengthToTheRightOfStartBar`
	__AttributeShape__000014_ArrowLengthToTheRightOfStartBar.IdentifierMeta = ref_models.Gantt{}.ArrowLengthToTheRightOfStartBar
	__AttributeShape__000014_ArrowLengthToTheRightOfStartBar.FieldTypeAsString = ``
	__AttributeShape__000014_ArrowLengthToTheRightOfStartBar.Structname = `Gantt`
	__AttributeShape__000014_ArrowLengthToTheRightOfStartBar.Fieldtypename = `float64`

	__AttributeShape__000015_ArrowTipLenght.Name = `ArrowTipLenght`

	//gong:ident [ref_models.Gantt.ArrowTipLenght] comment added to overcome the problem with the comment map association
	__AttributeShape__000015_ArrowTipLenght.Identifier = `ref_models.Gantt.ArrowTipLenght`
	__AttributeShape__000015_ArrowTipLenght.IdentifierMeta = ref_models.Gantt{}.ArrowTipLenght
	__AttributeShape__000015_ArrowTipLenght.FieldTypeAsString = ``
	__AttributeShape__000015_ArrowTipLenght.Structname = `Gantt`
	__AttributeShape__000015_ArrowTipLenght.Fieldtypename = `float64`

	__AttributeShape__000016_TimeLine_Color.Name = `TimeLine_Color`

	//gong:ident [ref_models.Gantt.TimeLine_Color] comment added to overcome the problem with the comment map association
	__AttributeShape__000016_TimeLine_Color.Identifier = `ref_models.Gantt.TimeLine_Color`
	__AttributeShape__000016_TimeLine_Color.IdentifierMeta = ref_models.Gantt{}.TimeLine_Color
	__AttributeShape__000016_TimeLine_Color.FieldTypeAsString = ``
	__AttributeShape__000016_TimeLine_Color.Structname = `Gantt`
	__AttributeShape__000016_TimeLine_Color.Fieldtypename = `string`

	__AttributeShape__000017_TimeLine_FillOpacity.Name = `TimeLine_FillOpacity`

	//gong:ident [ref_models.Gantt.TimeLine_FillOpacity] comment added to overcome the problem with the comment map association
	__AttributeShape__000017_TimeLine_FillOpacity.Identifier = `ref_models.Gantt.TimeLine_FillOpacity`
	__AttributeShape__000017_TimeLine_FillOpacity.IdentifierMeta = ref_models.Gantt{}.TimeLine_FillOpacity
	__AttributeShape__000017_TimeLine_FillOpacity.FieldTypeAsString = ``
	__AttributeShape__000017_TimeLine_FillOpacity.Structname = `Gantt`
	__AttributeShape__000017_TimeLine_FillOpacity.Fieldtypename = `float64`

	__AttributeShape__000018_TimeLine_Stroke.Name = `TimeLine_Stroke`

	//gong:ident [ref_models.Gantt.TimeLine_Stroke] comment added to overcome the problem with the comment map association
	__AttributeShape__000018_TimeLine_Stroke.Identifier = `ref_models.Gantt.TimeLine_Stroke`
	__AttributeShape__000018_TimeLine_Stroke.IdentifierMeta = ref_models.Gantt{}.TimeLine_Stroke
	__AttributeShape__000018_TimeLine_Stroke.FieldTypeAsString = ``
	__AttributeShape__000018_TimeLine_Stroke.Structname = `Gantt`
	__AttributeShape__000018_TimeLine_Stroke.Fieldtypename = `string`

	__AttributeShape__000019_TimeLine_StrokeWidth.Name = `TimeLine_StrokeWidth`

	//gong:ident [ref_models.Gantt.TimeLine_StrokeWidth] comment added to overcome the problem with the comment map association
	__AttributeShape__000019_TimeLine_StrokeWidth.Identifier = `ref_models.Gantt.TimeLine_StrokeWidth`
	__AttributeShape__000019_TimeLine_StrokeWidth.IdentifierMeta = ref_models.Gantt{}.TimeLine_StrokeWidth
	__AttributeShape__000019_TimeLine_StrokeWidth.FieldTypeAsString = ``
	__AttributeShape__000019_TimeLine_StrokeWidth.Structname = `Gantt`
	__AttributeShape__000019_TimeLine_StrokeWidth.Fieldtypename = `float64`

	__AttributeShape__000020_Group_Stroke.Name = `Group_Stroke`

	//gong:ident [ref_models.Gantt.Group_Stroke] comment added to overcome the problem with the comment map association
	__AttributeShape__000020_Group_Stroke.Identifier = `ref_models.Gantt.Group_Stroke`
	__AttributeShape__000020_Group_Stroke.IdentifierMeta = ref_models.Gantt{}.Group_Stroke
	__AttributeShape__000020_Group_Stroke.FieldTypeAsString = ``
	__AttributeShape__000020_Group_Stroke.Structname = `Gantt`
	__AttributeShape__000020_Group_Stroke.Fieldtypename = `string`

	__AttributeShape__000021_Group_StrokeWidth.Name = `Group_StrokeWidth`

	//gong:ident [ref_models.Gantt.Group_StrokeWidth] comment added to overcome the problem with the comment map association
	__AttributeShape__000021_Group_StrokeWidth.Identifier = `ref_models.Gantt.Group_StrokeWidth`
	__AttributeShape__000021_Group_StrokeWidth.IdentifierMeta = ref_models.Gantt{}.Group_StrokeWidth
	__AttributeShape__000021_Group_StrokeWidth.FieldTypeAsString = ``
	__AttributeShape__000021_Group_StrokeWidth.Structname = `Gantt`
	__AttributeShape__000021_Group_StrokeWidth.Fieldtypename = `float64`

	__AttributeShape__000022_Group_StrokeDashArray.Name = `Group_StrokeDashArray`

	//gong:ident [ref_models.Gantt.Group_StrokeDashArray] comment added to overcome the problem with the comment map association
	__AttributeShape__000022_Group_StrokeDashArray.Identifier = `ref_models.Gantt.Group_StrokeDashArray`
	__AttributeShape__000022_Group_StrokeDashArray.IdentifierMeta = ref_models.Gantt{}.Group_StrokeDashArray
	__AttributeShape__000022_Group_StrokeDashArray.FieldTypeAsString = ``
	__AttributeShape__000022_Group_StrokeDashArray.Structname = `Gantt`
	__AttributeShape__000022_Group_StrokeDashArray.Fieldtypename = `string`

	__AttributeShape__000023_DateYOffset.Name = `DateYOffset`

	//gong:ident [ref_models.Gantt.DateYOffset] comment added to overcome the problem with the comment map association
	__AttributeShape__000023_DateYOffset.Identifier = `ref_models.Gantt.DateYOffset`
	__AttributeShape__000023_DateYOffset.IdentifierMeta = ref_models.Gantt{}.DateYOffset
	__AttributeShape__000023_DateYOffset.FieldTypeAsString = ``
	__AttributeShape__000023_DateYOffset.Structname = `Gantt`
	__AttributeShape__000023_DateYOffset.Fieldtypename = `float64`

	__AttributeShape__000024_AlignOnStartEndOnYearStart.Name = `AlignOnStartEndOnYearStart`

	//gong:ident [ref_models.Gantt.AlignOnStartEndOnYearStart] comment added to overcome the problem with the comment map association
	__AttributeShape__000024_AlignOnStartEndOnYearStart.Identifier = `ref_models.Gantt.AlignOnStartEndOnYearStart`
	__AttributeShape__000024_AlignOnStartEndOnYearStart.IdentifierMeta = ref_models.Gantt{}.AlignOnStartEndOnYearStart
	__AttributeShape__000024_AlignOnStartEndOnYearStart.FieldTypeAsString = ``
	__AttributeShape__000024_AlignOnStartEndOnYearStart.Structname = `Gantt`
	__AttributeShape__000024_AlignOnStartEndOnYearStart.Fieldtypename = `bool`

	__AttributeShape__000025_Name.Name = `Name`

	//gong:ident [ref_models.Group.Name] comment added to overcome the problem with the comment map association
	__AttributeShape__000025_Name.Identifier = `ref_models.Group.Name`
	__AttributeShape__000025_Name.IdentifierMeta = ref_models.Group{}.Name
	__AttributeShape__000025_Name.FieldTypeAsString = ``
	__AttributeShape__000025_Name.Structname = `Group`
	__AttributeShape__000025_Name.Fieldtypename = `string`

	__AttributeShape__000026_Name.Name = `Name`

	//gong:ident [ref_models.Lane.Name] comment added to overcome the problem with the comment map association
	__AttributeShape__000026_Name.Identifier = `ref_models.Lane.Name`
	__AttributeShape__000026_Name.IdentifierMeta = ref_models.Lane{}.Name
	__AttributeShape__000026_Name.FieldTypeAsString = ``
	__AttributeShape__000026_Name.Structname = `Lane`
	__AttributeShape__000026_Name.Fieldtypename = `string`

	__AttributeShape__000027_Order.Name = `Order`

	//gong:ident [ref_models.Lane.Order] comment added to overcome the problem with the comment map association
	__AttributeShape__000027_Order.Identifier = `ref_models.Lane.Order`
	__AttributeShape__000027_Order.IdentifierMeta = ref_models.Lane{}.Order
	__AttributeShape__000027_Order.FieldTypeAsString = ``
	__AttributeShape__000027_Order.Structname = `Lane`
	__AttributeShape__000027_Order.Fieldtypename = `int`

	__AttributeShape__000028_Name.Name = `Name`

	//gong:ident [ref_models.Milestone.Name] comment added to overcome the problem with the comment map association
	__AttributeShape__000028_Name.Identifier = `ref_models.Milestone.Name`
	__AttributeShape__000028_Name.IdentifierMeta = ref_models.Milestone{}.Name
	__AttributeShape__000028_Name.FieldTypeAsString = ``
	__AttributeShape__000028_Name.Structname = `Milestone`
	__AttributeShape__000028_Name.Fieldtypename = `string`

	__AttributeShape__000029_Date.Name = `Date`

	//gong:ident [ref_models.Milestone.Date] comment added to overcome the problem with the comment map association
	__AttributeShape__000029_Date.Identifier = `ref_models.Milestone.Date`
	__AttributeShape__000029_Date.IdentifierMeta = ref_models.Milestone{}.Date
	__AttributeShape__000029_Date.FieldTypeAsString = ``
	__AttributeShape__000029_Date.Structname = `Milestone`
	__AttributeShape__000029_Date.Fieldtypename = `Time`

	__AttributeShape__000030_DisplayVerticalBar.Name = `DisplayVerticalBar`

	//gong:ident [ref_models.Milestone.DisplayVerticalBar] comment added to overcome the problem with the comment map association
	__AttributeShape__000030_DisplayVerticalBar.Identifier = `ref_models.Milestone.DisplayVerticalBar`
	__AttributeShape__000030_DisplayVerticalBar.IdentifierMeta = ref_models.Milestone{}.DisplayVerticalBar
	__AttributeShape__000030_DisplayVerticalBar.FieldTypeAsString = ``
	__AttributeShape__000030_DisplayVerticalBar.Structname = `Milestone`
	__AttributeShape__000030_DisplayVerticalBar.Fieldtypename = `bool`

	__AttributeShape__000031_Name.Name = `Name`

	//gong:ident [ref_models.LaneUse.Name] comment added to overcome the problem with the comment map association
	__AttributeShape__000031_Name.Identifier = `ref_models.LaneUse.Name`
	__AttributeShape__000031_Name.IdentifierMeta = ref_models.LaneUse{}.Name
	__AttributeShape__000031_Name.FieldTypeAsString = ``
	__AttributeShape__000031_Name.Structname = `LaneUse`
	__AttributeShape__000031_Name.Fieldtypename = `string`

	__AttributeShape__000032_Name.Name = `Name`

	//gong:ident [ref_models.Arrow.Name] comment added to overcome the problem with the comment map association
	__AttributeShape__000032_Name.Identifier = `ref_models.Arrow.Name`
	__AttributeShape__000032_Name.IdentifierMeta = ref_models.Arrow{}.Name
	__AttributeShape__000032_Name.FieldTypeAsString = ``
	__AttributeShape__000032_Name.Structname = `Arrow`
	__AttributeShape__000032_Name.Fieldtypename = `string`

	__AttributeShape__000033_OptionnalColor.Name = `OptionnalColor`

	//gong:ident [ref_models.Arrow.OptionnalColor] comment added to overcome the problem with the comment map association
	__AttributeShape__000033_OptionnalColor.Identifier = `ref_models.Arrow.OptionnalColor`
	__AttributeShape__000033_OptionnalColor.IdentifierMeta = ref_models.Arrow{}.OptionnalColor
	__AttributeShape__000033_OptionnalColor.FieldTypeAsString = ``
	__AttributeShape__000033_OptionnalColor.Structname = `Arrow`
	__AttributeShape__000033_OptionnalColor.Fieldtypename = `string`

	__AttributeShape__000034_OptionnalStroke.Name = `OptionnalStroke`

	//gong:ident [ref_models.Arrow.OptionnalStroke] comment added to overcome the problem with the comment map association
	__AttributeShape__000034_OptionnalStroke.Identifier = `ref_models.Arrow.OptionnalStroke`
	__AttributeShape__000034_OptionnalStroke.IdentifierMeta = ref_models.Arrow{}.OptionnalStroke
	__AttributeShape__000034_OptionnalStroke.FieldTypeAsString = ``
	__AttributeShape__000034_OptionnalStroke.Structname = `Arrow`
	__AttributeShape__000034_OptionnalStroke.Fieldtypename = `string`

	__AttributeShape__000035_Name.Name = `Name`

	//gong:ident [ref_models.Bar.Name] comment added to overcome the problem with the comment map association
	__AttributeShape__000035_Name.Identifier = `ref_models.Bar.Name`
	__AttributeShape__000035_Name.IdentifierMeta = ref_models.Bar{}.Name
	__AttributeShape__000035_Name.FieldTypeAsString = ``
	__AttributeShape__000035_Name.Structname = `Bar`
	__AttributeShape__000035_Name.Fieldtypename = `string`

	__AttributeShape__000036_Start.Name = `Start`

	//gong:ident [ref_models.Bar.Start] comment added to overcome the problem with the comment map association
	__AttributeShape__000036_Start.Identifier = `ref_models.Bar.Start`
	__AttributeShape__000036_Start.IdentifierMeta = ref_models.Bar{}.Start
	__AttributeShape__000036_Start.FieldTypeAsString = ``
	__AttributeShape__000036_Start.Structname = `Bar`
	__AttributeShape__000036_Start.Fieldtypename = `Time`

	__AttributeShape__000037_End.Name = `End`

	//gong:ident [ref_models.Bar.End] comment added to overcome the problem with the comment map association
	__AttributeShape__000037_End.Identifier = `ref_models.Bar.End`
	__AttributeShape__000037_End.IdentifierMeta = ref_models.Bar{}.End
	__AttributeShape__000037_End.FieldTypeAsString = ``
	__AttributeShape__000037_End.Structname = `Bar`
	__AttributeShape__000037_End.Fieldtypename = `Time`

	__AttributeShape__000038_ComputedDuration.Name = `ComputedDuration`

	//gong:ident [ref_models.Bar.ComputedDuration] comment added to overcome the problem with the comment map association
	__AttributeShape__000038_ComputedDuration.Identifier = `ref_models.Bar.ComputedDuration`
	__AttributeShape__000038_ComputedDuration.IdentifierMeta = ref_models.Bar{}.ComputedDuration
	__AttributeShape__000038_ComputedDuration.FieldTypeAsString = ``
	__AttributeShape__000038_ComputedDuration.Structname = `Bar`
	__AttributeShape__000038_ComputedDuration.Fieldtypename = `Duration`

	__AttributeShape__000039_OptionnalColor.Name = `OptionnalColor`

	//gong:ident [ref_models.Bar.OptionnalColor] comment added to overcome the problem with the comment map association
	__AttributeShape__000039_OptionnalColor.Identifier = `ref_models.Bar.OptionnalColor`
	__AttributeShape__000039_OptionnalColor.IdentifierMeta = ref_models.Bar{}.OptionnalColor
	__AttributeShape__000039_OptionnalColor.FieldTypeAsString = ``
	__AttributeShape__000039_OptionnalColor.Structname = `Bar`
	__AttributeShape__000039_OptionnalColor.Fieldtypename = `string`

	__AttributeShape__000040_OptionnalStroke.Name = `OptionnalStroke`

	//gong:ident [ref_models.Bar.OptionnalStroke] comment added to overcome the problem with the comment map association
	__AttributeShape__000040_OptionnalStroke.Identifier = `ref_models.Bar.OptionnalStroke`
	__AttributeShape__000040_OptionnalStroke.IdentifierMeta = ref_models.Bar{}.OptionnalStroke
	__AttributeShape__000040_OptionnalStroke.FieldTypeAsString = ``
	__AttributeShape__000040_OptionnalStroke.Structname = `Bar`
	__AttributeShape__000040_OptionnalStroke.Fieldtypename = `string`

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = `[true,true,true,true,true,true,true]`
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T20_48_55Z.Name = `Diagram Package created the 2025-06-04T20:48:55Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T20_48_55Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T20_48_55Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T20_48_55Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_Milestone.Name = `Default-Milestone`
	__GongStructShape__000000_Default_Milestone.X = 655.000000
	__GongStructShape__000000_Default_Milestone.Y = 516.000000

	//gong:ident [] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Milestone.Identifier = ``
	__GongStructShape__000000_Default_Milestone.IdentifierMeta = ref_models.Milestone{}
	__GongStructShape__000000_Default_Milestone.ShowNbInstances = false
	__GongStructShape__000000_Default_Milestone.NbInstances = 0
	__GongStructShape__000000_Default_Milestone.Width = 240.000000
	__GongStructShape__000000_Default_Milestone.Height = 123.000000
	__GongStructShape__000000_Default_Milestone.IsSelected = false

	__GongStructShape__000001_Default_LaneUse.Name = `Default-LaneUse`
	__GongStructShape__000001_Default_LaneUse.X = 1157.000000
	__GongStructShape__000001_Default_LaneUse.Y = 545.000000

	//gong:ident [] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_LaneUse.Identifier = ``
	__GongStructShape__000001_Default_LaneUse.IdentifierMeta = ref_models.LaneUse{}
	__GongStructShape__000001_Default_LaneUse.ShowNbInstances = false
	__GongStructShape__000001_Default_LaneUse.NbInstances = 0
	__GongStructShape__000001_Default_LaneUse.Width = 240.000000
	__GongStructShape__000001_Default_LaneUse.Height = 83.000000
	__GongStructShape__000001_Default_LaneUse.IsSelected = false

	__GongStructShape__000002_Default_Lane.Name = `Default-Lane`
	__GongStructShape__000002_Default_Lane.X = 656.000000
	__GongStructShape__000002_Default_Lane.Y = 164.000000

	//gong:ident [] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_Lane.Identifier = ``
	__GongStructShape__000002_Default_Lane.IdentifierMeta = ref_models.Lane{}
	__GongStructShape__000002_Default_Lane.ShowNbInstances = false
	__GongStructShape__000002_Default_Lane.NbInstances = 0
	__GongStructShape__000002_Default_Lane.Width = 240.000000
	__GongStructShape__000002_Default_Lane.Height = 138.000000
	__GongStructShape__000002_Default_Lane.IsSelected = false

	__GongStructShape__000003_Default_Group.Name = `Default-Group`
	__GongStructShape__000003_Default_Group.X = 659.000000
	__GongStructShape__000003_Default_Group.Y = 68.000000

	//gong:ident [] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_Default_Group.Identifier = ``
	__GongStructShape__000003_Default_Group.IdentifierMeta = ref_models.Group{}
	__GongStructShape__000003_Default_Group.ShowNbInstances = false
	__GongStructShape__000003_Default_Group.NbInstances = 0
	__GongStructShape__000003_Default_Group.Width = 240.000000
	__GongStructShape__000003_Default_Group.Height = 83.000000
	__GongStructShape__000003_Default_Group.IsSelected = false

	__GongStructShape__000004_Default_Gantt.Name = `Default-Gantt`
	__GongStructShape__000004_Default_Gantt.X = 48.000000
	__GongStructShape__000004_Default_Gantt.Y = 81.000000

	//gong:ident [] comment added to overcome the problem with the comment map association
	__GongStructShape__000004_Default_Gantt.Identifier = ``
	__GongStructShape__000004_Default_Gantt.IdentifierMeta = ref_models.Gantt{}
	__GongStructShape__000004_Default_Gantt.ShowNbInstances = false
	__GongStructShape__000004_Default_Gantt.NbInstances = 0
	__GongStructShape__000004_Default_Gantt.Width = 240.000000
	__GongStructShape__000004_Default_Gantt.Height = 563.000000
	__GongStructShape__000004_Default_Gantt.IsSelected = false

	__GongStructShape__000005_Default_Bar.Name = `Default-Bar`
	__GongStructShape__000005_Default_Bar.X = 657.000000
	__GongStructShape__000005_Default_Bar.Y = 322.000000

	//gong:ident [] comment added to overcome the problem with the comment map association
	__GongStructShape__000005_Default_Bar.Identifier = ``
	__GongStructShape__000005_Default_Bar.IdentifierMeta = ref_models.Bar{}
	__GongStructShape__000005_Default_Bar.ShowNbInstances = false
	__GongStructShape__000005_Default_Bar.NbInstances = 0
	__GongStructShape__000005_Default_Bar.Width = 240.000000
	__GongStructShape__000005_Default_Bar.Height = 170.000000
	__GongStructShape__000005_Default_Bar.IsSelected = false

	__GongStructShape__000006_Default_Arrow.Name = `Default-Arrow`
	__GongStructShape__000006_Default_Arrow.X = 656.000000
	__GongStructShape__000006_Default_Arrow.Y = 677.000000

	//gong:ident [] comment added to overcome the problem with the comment map association
	__GongStructShape__000006_Default_Arrow.Identifier = ``
	__GongStructShape__000006_Default_Arrow.IdentifierMeta = ref_models.Arrow{}
	__GongStructShape__000006_Default_Arrow.ShowNbInstances = false
	__GongStructShape__000006_Default_Arrow.NbInstances = 0
	__GongStructShape__000006_Default_Arrow.Width = 240.000000
	__GongStructShape__000006_Default_Arrow.Height = 123.000000
	__GongStructShape__000006_Default_Arrow.IsSelected = false

	__LinkShape__000000_Lanes.Name = `Lanes`

	//gong:ident [ref_models.Gantt.Lanes] comment added to overcome the problem with the comment map association
	__LinkShape__000000_Lanes.Identifier = `ref_models.Gantt.Lanes`
	__LinkShape__000000_Lanes.IdentifierMeta = ref_models.Gantt{}.Lanes

	//gong:ident [ref_models.Lane] comment added to overcome the problem with the comment map association
	__LinkShape__000000_Lanes.Fieldtypename = `ref_models.Lane`
	__LinkShape__000000_Lanes.FieldTypeIdentifierMeta = ref_models.Lane{}
	__LinkShape__000000_Lanes.FieldOffsetX = 0.000000
	__LinkShape__000000_Lanes.FieldOffsetY = 0.000000
	__LinkShape__000000_Lanes.TargetMultiplicity = models.MANY
	__LinkShape__000000_Lanes.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Lanes.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Lanes.SourceMultiplicity = models.MANY
	__LinkShape__000000_Lanes.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Lanes.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Lanes.X = 711.000000
	__LinkShape__000000_Lanes.Y = 454.000000
	__LinkShape__000000_Lanes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Lanes.StartRatio = 0.500000
	__LinkShape__000000_Lanes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Lanes.EndRatio = 0.500000
	__LinkShape__000000_Lanes.CornerOffsetRatio = 1.380000

	__LinkShape__000001_Milestones.Name = `Milestones`

	//gong:ident [ref_models.Gantt.Milestones] comment added to overcome the problem with the comment map association
	__LinkShape__000001_Milestones.Identifier = `ref_models.Gantt.Milestones`
	__LinkShape__000001_Milestones.IdentifierMeta = ref_models.Gantt{}.Milestones

	//gong:ident [ref_models.Milestone] comment added to overcome the problem with the comment map association
	__LinkShape__000001_Milestones.Fieldtypename = `ref_models.Milestone`
	__LinkShape__000001_Milestones.FieldTypeIdentifierMeta = ref_models.Milestone{}
	__LinkShape__000001_Milestones.FieldOffsetX = 0.000000
	__LinkShape__000001_Milestones.FieldOffsetY = 0.000000
	__LinkShape__000001_Milestones.TargetMultiplicity = models.MANY
	__LinkShape__000001_Milestones.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Milestones.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Milestones.SourceMultiplicity = models.MANY
	__LinkShape__000001_Milestones.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Milestones.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Milestones.X = 585.500000
	__LinkShape__000001_Milestones.Y = 532.500000
	__LinkShape__000001_Milestones.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Milestones.StartRatio = 0.500000
	__LinkShape__000001_Milestones.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Milestones.EndRatio = 0.500000
	__LinkShape__000001_Milestones.CornerOffsetRatio = 1.380000

	__LinkShape__000002_Groups.Name = `Groups`

	//gong:ident [ref_models.Gantt.Groups] comment added to overcome the problem with the comment map association
	__LinkShape__000002_Groups.Identifier = `ref_models.Gantt.Groups`
	__LinkShape__000002_Groups.IdentifierMeta = ref_models.Gantt{}.Groups

	//gong:ident [ref_models.Group] comment added to overcome the problem with the comment map association
	__LinkShape__000002_Groups.Fieldtypename = `ref_models.Group`
	__LinkShape__000002_Groups.FieldTypeIdentifierMeta = ref_models.Group{}
	__LinkShape__000002_Groups.FieldOffsetX = 0.000000
	__LinkShape__000002_Groups.FieldOffsetY = 0.000000
	__LinkShape__000002_Groups.TargetMultiplicity = models.MANY
	__LinkShape__000002_Groups.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000002_Groups.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000002_Groups.SourceMultiplicity = models.MANY
	__LinkShape__000002_Groups.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000002_Groups.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000002_Groups.X = 713.500000
	__LinkShape__000002_Groups.Y = 356.000000
	__LinkShape__000002_Groups.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_Groups.StartRatio = 0.500000
	__LinkShape__000002_Groups.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_Groups.EndRatio = 0.500000
	__LinkShape__000002_Groups.CornerOffsetRatio = 1.380000

	__LinkShape__000003_Arrows.Name = `Arrows`

	//gong:ident [ref_models.Gantt.Arrows] comment added to overcome the problem with the comment map association
	__LinkShape__000003_Arrows.Identifier = `ref_models.Gantt.Arrows`
	__LinkShape__000003_Arrows.IdentifierMeta = ref_models.Gantt{}.Arrows

	//gong:ident [ref_models.Arrow] comment added to overcome the problem with the comment map association
	__LinkShape__000003_Arrows.Fieldtypename = `ref_models.Arrow`
	__LinkShape__000003_Arrows.FieldTypeIdentifierMeta = ref_models.Arrow{}
	__LinkShape__000003_Arrows.FieldOffsetX = 0.000000
	__LinkShape__000003_Arrows.FieldOffsetY = 0.000000
	__LinkShape__000003_Arrows.TargetMultiplicity = models.MANY
	__LinkShape__000003_Arrows.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000003_Arrows.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000003_Arrows.SourceMultiplicity = models.MANY
	__LinkShape__000003_Arrows.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000003_Arrows.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000003_Arrows.X = 407.500000
	__LinkShape__000003_Arrows.Y = 556.500000
	__LinkShape__000003_Arrows.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000003_Arrows.StartRatio = 0.500000
	__LinkShape__000003_Arrows.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000003_Arrows.EndRatio = 0.911947
	__LinkShape__000003_Arrows.CornerOffsetRatio = 1.380000

	__LinkShape__000004_GroupLanes.Name = `GroupLanes`

	//gong:ident [ref_models.Group.GroupLanes] comment added to overcome the problem with the comment map association
	__LinkShape__000004_GroupLanes.Identifier = `ref_models.Group.GroupLanes`
	__LinkShape__000004_GroupLanes.IdentifierMeta = ref_models.Group{}.GroupLanes

	//gong:ident [ref_models.Lane] comment added to overcome the problem with the comment map association
	__LinkShape__000004_GroupLanes.Fieldtypename = `ref_models.Lane`
	__LinkShape__000004_GroupLanes.FieldTypeIdentifierMeta = ref_models.Lane{}
	__LinkShape__000004_GroupLanes.FieldOffsetX = 0.000000
	__LinkShape__000004_GroupLanes.FieldOffsetY = 0.000000
	__LinkShape__000004_GroupLanes.TargetMultiplicity = models.MANY
	__LinkShape__000004_GroupLanes.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000004_GroupLanes.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000004_GroupLanes.SourceMultiplicity = models.MANY
	__LinkShape__000004_GroupLanes.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000004_GroupLanes.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000004_GroupLanes.X = 1016.500000
	__LinkShape__000004_GroupLanes.Y = 207.500000
	__LinkShape__000004_GroupLanes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000004_GroupLanes.StartRatio = 0.500000
	__LinkShape__000004_GroupLanes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000004_GroupLanes.EndRatio = 0.164117
	__LinkShape__000004_GroupLanes.CornerOffsetRatio = 1.723047

	__LinkShape__000005_Bars.Name = `Bars`

	//gong:ident [ref_models.Lane.Bars] comment added to overcome the problem with the comment map association
	__LinkShape__000005_Bars.Identifier = `ref_models.Lane.Bars`
	__LinkShape__000005_Bars.IdentifierMeta = ref_models.Lane{}.Bars

	//gong:ident [ref_models.Bar] comment added to overcome the problem with the comment map association
	__LinkShape__000005_Bars.Fieldtypename = `ref_models.Bar`
	__LinkShape__000005_Bars.FieldTypeIdentifierMeta = ref_models.Bar{}
	__LinkShape__000005_Bars.FieldOffsetX = 0.000000
	__LinkShape__000005_Bars.FieldOffsetY = 0.000000
	__LinkShape__000005_Bars.TargetMultiplicity = models.MANY
	__LinkShape__000005_Bars.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000005_Bars.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000005_Bars.SourceMultiplicity = models.MANY
	__LinkShape__000005_Bars.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000005_Bars.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000005_Bars.X = 1010.500000
	__LinkShape__000005_Bars.Y = 300.000000
	__LinkShape__000005_Bars.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000005_Bars.StartRatio = 0.797029
	__LinkShape__000005_Bars.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000005_Bars.EndRatio = 0.500000
	__LinkShape__000005_Bars.CornerOffsetRatio = 1.380000

	__LinkShape__000006_LanesToDisplayMilestoneUse.Name = `LanesToDisplayMilestoneUse`

	//gong:ident [ref_models.Milestone.LanesToDisplayMilestoneUse] comment added to overcome the problem with the comment map association
	__LinkShape__000006_LanesToDisplayMilestoneUse.Identifier = `ref_models.Milestone.LanesToDisplayMilestoneUse`
	__LinkShape__000006_LanesToDisplayMilestoneUse.IdentifierMeta = ref_models.Milestone{}.LanesToDisplayMilestoneUse

	//gong:ident [ref_models.LaneUse] comment added to overcome the problem with the comment map association
	__LinkShape__000006_LanesToDisplayMilestoneUse.Fieldtypename = `ref_models.LaneUse`
	__LinkShape__000006_LanesToDisplayMilestoneUse.FieldTypeIdentifierMeta = ref_models.LaneUse{}
	__LinkShape__000006_LanesToDisplayMilestoneUse.FieldOffsetX = 0.000000
	__LinkShape__000006_LanesToDisplayMilestoneUse.FieldOffsetY = 0.000000
	__LinkShape__000006_LanesToDisplayMilestoneUse.TargetMultiplicity = models.MANY
	__LinkShape__000006_LanesToDisplayMilestoneUse.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000006_LanesToDisplayMilestoneUse.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000006_LanesToDisplayMilestoneUse.SourceMultiplicity = models.MANY
	__LinkShape__000006_LanesToDisplayMilestoneUse.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000006_LanesToDisplayMilestoneUse.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000006_LanesToDisplayMilestoneUse.X = 1204.500000
	__LinkShape__000006_LanesToDisplayMilestoneUse.Y = 534.500000
	__LinkShape__000006_LanesToDisplayMilestoneUse.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000006_LanesToDisplayMilestoneUse.StartRatio = 0.500000
	__LinkShape__000006_LanesToDisplayMilestoneUse.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000006_LanesToDisplayMilestoneUse.EndRatio = 0.500000
	__LinkShape__000006_LanesToDisplayMilestoneUse.CornerOffsetRatio = 1.380000

	__LinkShape__000007_Lane.Name = `Lane`

	//gong:ident [ref_models.LaneUse.Lane] comment added to overcome the problem with the comment map association
	__LinkShape__000007_Lane.Identifier = `ref_models.LaneUse.Lane`
	__LinkShape__000007_Lane.IdentifierMeta = ref_models.LaneUse{}.Lane

	//gong:ident [ref_models.Lane] comment added to overcome the problem with the comment map association
	__LinkShape__000007_Lane.Fieldtypename = `ref_models.Lane`
	__LinkShape__000007_Lane.FieldTypeIdentifierMeta = ref_models.Lane{}
	__LinkShape__000007_Lane.FieldOffsetX = 0.000000
	__LinkShape__000007_Lane.FieldOffsetY = 0.000000
	__LinkShape__000007_Lane.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000007_Lane.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000007_Lane.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000007_Lane.SourceMultiplicity = models.MANY
	__LinkShape__000007_Lane.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000007_Lane.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000007_Lane.X = 1267.000000
	__LinkShape__000007_Lane.Y = 400.000000
	__LinkShape__000007_Lane.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000007_Lane.StartRatio = 0.500000
	__LinkShape__000007_Lane.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000007_Lane.EndRatio = 0.500000
	__LinkShape__000007_Lane.CornerOffsetRatio = 1.380000

	__LinkShape__000008_From.Name = `From`

	//gong:ident [ref_models.Arrow.From] comment added to overcome the problem with the comment map association
	__LinkShape__000008_From.Identifier = `ref_models.Arrow.From`
	__LinkShape__000008_From.IdentifierMeta = ref_models.Arrow{}.From

	//gong:ident [ref_models.Bar] comment added to overcome the problem with the comment map association
	__LinkShape__000008_From.Fieldtypename = `ref_models.Bar`
	__LinkShape__000008_From.FieldTypeIdentifierMeta = ref_models.Bar{}
	__LinkShape__000008_From.FieldOffsetX = 0.000000
	__LinkShape__000008_From.FieldOffsetY = 0.000000
	__LinkShape__000008_From.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000008_From.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000008_From.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000008_From.SourceMultiplicity = models.MANY
	__LinkShape__000008_From.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000008_From.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000008_From.X = 1014.500000
	__LinkShape__000008_From.Y = 573.500000
	__LinkShape__000008_From.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000008_From.StartRatio = 0.261462
	__LinkShape__000008_From.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000008_From.EndRatio = 0.273962
	__LinkShape__000008_From.CornerOffsetRatio = -0.864453

	__LinkShape__000009_To.Name = `To`

	//gong:ident [ref_models.Arrow.To] comment added to overcome the problem with the comment map association
	__LinkShape__000009_To.Identifier = `ref_models.Arrow.To`
	__LinkShape__000009_To.IdentifierMeta = ref_models.Arrow{}.To

	//gong:ident [ref_models.Bar] comment added to overcome the problem with the comment map association
	__LinkShape__000009_To.Fieldtypename = `ref_models.Bar`
	__LinkShape__000009_To.FieldTypeIdentifierMeta = ref_models.Bar{}
	__LinkShape__000009_To.FieldOffsetX = 0.000000
	__LinkShape__000009_To.FieldOffsetY = 0.000000
	__LinkShape__000009_To.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000009_To.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000009_To.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000009_To.SourceMultiplicity = models.MANY
	__LinkShape__000009_To.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000009_To.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000009_To.X = 1014.500000
	__LinkShape__000009_To.Y = 573.500000
	__LinkShape__000009_To.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000009_To.StartRatio = 0.251753
	__LinkShape__000009_To.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000009_To.EndRatio = 0.823413
	__LinkShape__000009_To.CornerOffsetRatio = -0.647786

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Milestone)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_LaneUse)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Lane)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_Group)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000004_Default_Gantt)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000005_Default_Bar)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000006_Default_Arrow)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T20_48_55Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T20_48_55Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T20_48_55Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Milestone.AttributeShapes = append(__GongStructShape__000000_Default_Milestone.AttributeShapes, __AttributeShape__000028_Name)
	__GongStructShape__000000_Default_Milestone.AttributeShapes = append(__GongStructShape__000000_Default_Milestone.AttributeShapes, __AttributeShape__000029_Date)
	__GongStructShape__000000_Default_Milestone.AttributeShapes = append(__GongStructShape__000000_Default_Milestone.AttributeShapes, __AttributeShape__000030_DisplayVerticalBar)
	__GongStructShape__000000_Default_Milestone.LinkShapes = append(__GongStructShape__000000_Default_Milestone.LinkShapes, __LinkShape__000006_LanesToDisplayMilestoneUse)
	__GongStructShape__000001_Default_LaneUse.AttributeShapes = append(__GongStructShape__000001_Default_LaneUse.AttributeShapes, __AttributeShape__000031_Name)
	__GongStructShape__000001_Default_LaneUse.LinkShapes = append(__GongStructShape__000001_Default_LaneUse.LinkShapes, __LinkShape__000007_Lane)
	__GongStructShape__000002_Default_Lane.AttributeShapes = append(__GongStructShape__000002_Default_Lane.AttributeShapes, __AttributeShape__000026_Name)
	__GongStructShape__000002_Default_Lane.AttributeShapes = append(__GongStructShape__000002_Default_Lane.AttributeShapes, __AttributeShape__000027_Order)
	__GongStructShape__000002_Default_Lane.LinkShapes = append(__GongStructShape__000002_Default_Lane.LinkShapes, __LinkShape__000005_Bars)
	__GongStructShape__000003_Default_Group.AttributeShapes = append(__GongStructShape__000003_Default_Group.AttributeShapes, __AttributeShape__000025_Name)
	__GongStructShape__000003_Default_Group.LinkShapes = append(__GongStructShape__000003_Default_Group.LinkShapes, __LinkShape__000004_GroupLanes)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000000_Name)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000001_ComputedStart)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000002_ComputedEnd)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000003_ComputedDuration)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000004_UseManualStartAndEndDates)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000005_ManualStart)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000006_ManualEnd)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000007_LaneHeight)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000008_RatioBarToLaneHeight)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000009_YTopMargin)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000010_XLeftText)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000011_TextHeight)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000012_XLeftLanes)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000013_XRightMargin)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000014_ArrowLengthToTheRightOfStartBar)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000015_ArrowTipLenght)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000016_TimeLine_Color)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000017_TimeLine_FillOpacity)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000018_TimeLine_Stroke)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000019_TimeLine_StrokeWidth)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000020_Group_Stroke)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000021_Group_StrokeWidth)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000022_Group_StrokeDashArray)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000023_DateYOffset)
	__GongStructShape__000004_Default_Gantt.AttributeShapes = append(__GongStructShape__000004_Default_Gantt.AttributeShapes, __AttributeShape__000024_AlignOnStartEndOnYearStart)
	__GongStructShape__000004_Default_Gantt.LinkShapes = append(__GongStructShape__000004_Default_Gantt.LinkShapes, __LinkShape__000000_Lanes)
	__GongStructShape__000004_Default_Gantt.LinkShapes = append(__GongStructShape__000004_Default_Gantt.LinkShapes, __LinkShape__000001_Milestones)
	__GongStructShape__000004_Default_Gantt.LinkShapes = append(__GongStructShape__000004_Default_Gantt.LinkShapes, __LinkShape__000002_Groups)
	__GongStructShape__000004_Default_Gantt.LinkShapes = append(__GongStructShape__000004_Default_Gantt.LinkShapes, __LinkShape__000003_Arrows)
	__GongStructShape__000005_Default_Bar.AttributeShapes = append(__GongStructShape__000005_Default_Bar.AttributeShapes, __AttributeShape__000035_Name)
	__GongStructShape__000005_Default_Bar.AttributeShapes = append(__GongStructShape__000005_Default_Bar.AttributeShapes, __AttributeShape__000036_Start)
	__GongStructShape__000005_Default_Bar.AttributeShapes = append(__GongStructShape__000005_Default_Bar.AttributeShapes, __AttributeShape__000037_End)
	__GongStructShape__000005_Default_Bar.AttributeShapes = append(__GongStructShape__000005_Default_Bar.AttributeShapes, __AttributeShape__000038_ComputedDuration)
	__GongStructShape__000005_Default_Bar.AttributeShapes = append(__GongStructShape__000005_Default_Bar.AttributeShapes, __AttributeShape__000039_OptionnalColor)
	__GongStructShape__000005_Default_Bar.AttributeShapes = append(__GongStructShape__000005_Default_Bar.AttributeShapes, __AttributeShape__000040_OptionnalStroke)
	__GongStructShape__000006_Default_Arrow.AttributeShapes = append(__GongStructShape__000006_Default_Arrow.AttributeShapes, __AttributeShape__000032_Name)
	__GongStructShape__000006_Default_Arrow.AttributeShapes = append(__GongStructShape__000006_Default_Arrow.AttributeShapes, __AttributeShape__000033_OptionnalColor)
	__GongStructShape__000006_Default_Arrow.AttributeShapes = append(__GongStructShape__000006_Default_Arrow.AttributeShapes, __AttributeShape__000034_OptionnalStroke)
	__GongStructShape__000006_Default_Arrow.LinkShapes = append(__GongStructShape__000006_Default_Arrow.LinkShapes, __LinkShape__000008_From)
	__GongStructShape__000006_Default_Arrow.LinkShapes = append(__GongStructShape__000006_Default_Arrow.LinkShapes, __LinkShape__000009_To)
	// setup of LinkShape instances pointers
}
