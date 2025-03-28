// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/gantt/go/models"
	"github.com/fullstack-lang/gong/lib/gantt/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Arrow:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("From", instanceWithInferedType.From, formGroup, probe)
		AssociationFieldToForm("To", instanceWithInferedType.To, formGroup, probe)
		BasicFieldtoForm("OptionnalColor", instanceWithInferedType.OptionnalColor, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OptionnalStroke", instanceWithInferedType.OptionnalStroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Gantt"
			rf.Fieldname = "Arrows"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Gantt),
					"Arrows",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Gantt, *models.Arrow](
					nil,
					"Arrows",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Bar:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Start", instanceWithInferedType.Start, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("End", instanceWithInferedType.End, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedDuration", instanceWithInferedType.ComputedDuration, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OptionnalColor", instanceWithInferedType.OptionnalColor, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OptionnalStroke", instanceWithInferedType.OptionnalStroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Lane"
			rf.Fieldname = "Bars"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Lane),
					"Bars",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Lane, *models.Bar](
					nil,
					"Bars",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Gantt:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedStart", instanceWithInferedType.ComputedStart, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedEnd", instanceWithInferedType.ComputedEnd, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedDuration", instanceWithInferedType.ComputedDuration, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("UseManualStartAndEndDates", instanceWithInferedType.UseManualStartAndEndDates, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ManualStart", instanceWithInferedType.ManualStart, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ManualEnd", instanceWithInferedType.ManualEnd, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LaneHeight", instanceWithInferedType.LaneHeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("RatioBarToLaneHeight", instanceWithInferedType.RatioBarToLaneHeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("YTopMargin", instanceWithInferedType.YTopMargin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("XLeftText", instanceWithInferedType.XLeftText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TextHeight", instanceWithInferedType.TextHeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("XLeftLanes", instanceWithInferedType.XLeftLanes, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("XRightMargin", instanceWithInferedType.XRightMargin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ArrowLengthToTheRightOfStartBar", instanceWithInferedType.ArrowLengthToTheRightOfStartBar, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ArrowTipLenght", instanceWithInferedType.ArrowTipLenght, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TimeLine_Color", instanceWithInferedType.TimeLine_Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TimeLine_FillOpacity", instanceWithInferedType.TimeLine_FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TimeLine_Stroke", instanceWithInferedType.TimeLine_Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TimeLine_StrokeWidth", instanceWithInferedType.TimeLine_StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Group_Stroke", instanceWithInferedType.Group_Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Group_StrokeWidth", instanceWithInferedType.Group_StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Group_StrokeDashArray", instanceWithInferedType.Group_StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DateYOffset", instanceWithInferedType.DateYOffset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("AlignOnStartEndOnYearStart", instanceWithInferedType.AlignOnStartEndOnYearStart, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Lanes", instanceWithInferedType, &instanceWithInferedType.Lanes, formGroup, probe)
		AssociationSliceToForm("Milestones", instanceWithInferedType, &instanceWithInferedType.Milestones, formGroup, probe)
		AssociationSliceToForm("Groups", instanceWithInferedType, &instanceWithInferedType.Groups, formGroup, probe)
		AssociationSliceToForm("Arrows", instanceWithInferedType, &instanceWithInferedType.Arrows, formGroup, probe)

	case *models.Group:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("GroupLanes", instanceWithInferedType, &instanceWithInferedType.GroupLanes, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Gantt"
			rf.Fieldname = "Groups"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Gantt),
					"Groups",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Gantt, *models.Group](
					nil,
					"Groups",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Lane:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Order", instanceWithInferedType.Order, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Bars", instanceWithInferedType, &instanceWithInferedType.Bars, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Gantt"
			rf.Fieldname = "Lanes"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Gantt),
					"Lanes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Gantt, *models.Lane](
					nil,
					"Lanes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "GroupLanes"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Group),
					"GroupLanes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Group, *models.Lane](
					nil,
					"GroupLanes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.LaneUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Lane", instanceWithInferedType.Lane, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Milestone"
			rf.Fieldname = "LanesToDisplayMilestoneUse"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Milestone),
					"LanesToDisplayMilestoneUse",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Milestone, *models.LaneUse](
					nil,
					"LanesToDisplayMilestoneUse",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Milestone:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Date", instanceWithInferedType.Date, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DisplayVerticalBar", instanceWithInferedType.DisplayVerticalBar, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("LanesToDisplayMilestoneUse", instanceWithInferedType, &instanceWithInferedType.LanesToDisplayMilestoneUse, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Gantt"
			rf.Fieldname = "Milestones"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Gantt),
					"Milestones",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Gantt, *models.Milestone](
					nil,
					"Milestones",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	default:
		_ = instanceWithInferedType
	}
}
