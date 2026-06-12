// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/gantt/go/models"
)

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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Gantt, *models.Arrow](
				"Gantt",
				"Arrows",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Gantt) []*models.Arrow {
					return owner.Arrows
				})
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Lane, *models.Bar](
				"Lane",
				"Bars",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Lane) []*models.Bar {
					return owner.Bars
				})
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Group:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("GroupLanes", instanceWithInferedType, &instanceWithInferedType.GroupLanes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Gantt, *models.Group](
				"Gantt",
				"Groups",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Gantt) []*models.Group {
					return owner.Groups
				})
		}

	case *models.Lane:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Order", instanceWithInferedType.Order, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Bars", instanceWithInferedType, &instanceWithInferedType.Bars, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Gantt, *models.Lane](
				"Gantt",
				"Lanes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Gantt) []*models.Lane {
					return owner.Lanes
				})
		}
		{
			AssociationReverseSliceToForm[*models.Group, *models.Lane](
				"Group",
				"GroupLanes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Group) []*models.Lane {
					return owner.GroupLanes
				})
		}
		{
			AssociationReverseSliceToForm[*models.Milestone, *models.Lane](
				"Milestone",
				"LanesToDisplay",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Milestone) []*models.Lane {
					return owner.LanesToDisplay
				})
		}

	case *models.LaneUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Lane", instanceWithInferedType.Lane, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Milestone:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Date", instanceWithInferedType.Date, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DisplayVerticalBar", instanceWithInferedType.DisplayVerticalBar, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("LanesToDisplay", instanceWithInferedType, &instanceWithInferedType.LanesToDisplay, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Gantt, *models.Milestone](
				"Gantt",
				"Milestones",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Gantt) []*models.Milestone {
					return owner.Milestones
				})
		}

	default:
		_ = instanceWithInferedType
	}
}
