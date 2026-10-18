// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/gantt/go/models"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func StageSetFillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *StageSetProbe,
) {
	switch inst := instance.(type) {
	case *models.Arrow:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("From", inst.From, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Bar](), probe.formStage)
		StageSetAssociationFieldToForm("To", inst.To, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Bar](), probe.formStage)
		StageSetBasicFieldtoForm("OptionnalColor", inst.OptionnalColor, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OptionnalStroke", inst.OptionnalStroke, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Gantts {
				for _, target := range src.Arrows {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Gantt", "Arrows", refNames, formGroup, probe.formStage)
		}
	case *models.Bar:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Start", inst.Start, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("End", inst.End, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedDuration", inst.ComputedDuration, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OptionnalColor", inst.OptionnalColor, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OptionnalStroke", inst.OptionnalStroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FillOpacity", inst.FillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeWidth", inst.StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArray", inst.StrokeDashArray, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Arrows {
				if src.From == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Arrow", "From", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Arrows {
				if src.To == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Arrow", "To", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Lanes {
				for _, target := range src.Bars {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Lane", "Bars", refNames, formGroup, probe.formStage)
		}
	case *models.Gantt:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedStart", inst.ComputedStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedEnd", inst.ComputedEnd, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedDuration", inst.ComputedDuration, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("UseManualStartAndEndDates", inst.UseManualStartAndEndDates, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ManualStart", inst.ManualStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ManualEnd", inst.ManualEnd, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LaneHeight", inst.LaneHeight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RatioBarToLaneHeight", inst.RatioBarToLaneHeight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("YTopMargin", inst.YTopMargin, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("XLeftText", inst.XLeftText, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TextHeight", inst.TextHeight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("XLeftLanes", inst.XLeftLanes, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("XRightMargin", inst.XRightMargin, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ArrowLengthToTheRightOfStartBar", inst.ArrowLengthToTheRightOfStartBar, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ArrowTipLenght", inst.ArrowTipLenght, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TimeLine_Color", inst.TimeLine_Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TimeLine_FillOpacity", inst.TimeLine_FillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TimeLine_Stroke", inst.TimeLine_Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TimeLine_StrokeWidth", inst.TimeLine_StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Group_Stroke", inst.Group_Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Group_StrokeWidth", inst.Group_StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Group_StrokeDashArray", inst.Group_StrokeDashArray, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DateYOffset", inst.DateYOffset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("AlignOnStartEndOnYearStart", inst.AlignOnStartEndOnYearStart, probe.formStage, formGroup)

		{
			// Slice of pointers: Lanes
			div := (&form.FormDiv{Name: "Lanes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Lanes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Lanes",
				Label: "Lanes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Milestones
			div := (&form.FormDiv{Name: "Milestones"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Milestones {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Milestones",
				Label: "Milestones",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Groups
			div := (&form.FormDiv{Name: "Groups"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Groups {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Groups",
				Label: "Groups",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Arrows
			div := (&form.FormDiv{Name: "Arrows"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Arrows {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Arrows",
				Label: "Arrows",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
	case *models.Group:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: GroupLanes
			div := (&form.FormDiv{Name: "GroupLanes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GroupLanes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GroupLanes",
				Label: "GroupLanes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Gantts {
				for _, target := range src.Groups {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Gantt", "Groups", refNames, formGroup, probe.formStage)
		}
	case *models.Lane:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Order", inst.Order, probe.formStage, formGroup)

		{
			// Slice of pointers: Bars
			div := (&form.FormDiv{Name: "Bars"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Bars {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Bars",
				Label: "Bars",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Gantts {
				for _, target := range src.Lanes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Gantt", "Lanes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Groups {
				for _, target := range src.GroupLanes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Group", "GroupLanes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.LaneUses {
				if src.Lane == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.LaneUse", "Lane", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Milestones {
				for _, target := range src.LanesToDisplay {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Milestone", "LanesToDisplay", refNames, formGroup, probe.formStage)
		}
	case *models.LaneUse:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Lane", inst.Lane, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Lane](), probe.formStage)
	case *models.Milestone:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Date", inst.Date, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DisplayVerticalBar", inst.DisplayVerticalBar, probe.formStage, formGroup)

		{
			// Slice of pointers: LanesToDisplay
			div := (&form.FormDiv{Name: "LanesToDisplay"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.LanesToDisplay {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "LanesToDisplay",
				Label: "LanesToDisplay",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Gantts {
				for _, target := range src.Milestones {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Gantt", "Milestones", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
