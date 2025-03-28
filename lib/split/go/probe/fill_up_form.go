// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/split/go/models"
	"github.com/fullstack-lang/gong/lib/split/go/orm"
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
	case *models.AsSplit:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Direction", instanceWithInferedType.Direction, instanceWithInferedType, probe.formStage, formGroup)
		AssociationSliceToForm("AsSplitAreas", instanceWithInferedType, &instanceWithInferedType.AsSplitAreas, formGroup, probe)

	case *models.AsSplitArea:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ShowNameInHeader", instanceWithInferedType.ShowNameInHeader, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Size", instanceWithInferedType.Size, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsAny", instanceWithInferedType.IsAny, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("AsSplit", instanceWithInferedType.AsSplit, formGroup, probe)
		AssociationFieldToForm("Button", instanceWithInferedType.Button, formGroup, probe)
		AssociationFieldToForm("Cursor", instanceWithInferedType.Cursor, formGroup, probe)
		AssociationFieldToForm("Doc", instanceWithInferedType.Doc, formGroup, probe)
		AssociationFieldToForm("Form", instanceWithInferedType.Form, formGroup, probe)
		AssociationFieldToForm("Load", instanceWithInferedType.Load, formGroup, probe)
		AssociationFieldToForm("Slider", instanceWithInferedType.Slider, formGroup, probe)
		AssociationFieldToForm("Split", instanceWithInferedType.Split, formGroup, probe)
		AssociationFieldToForm("Svg", instanceWithInferedType.Svg, formGroup, probe)
		AssociationFieldToForm("Table", instanceWithInferedType.Table, formGroup, probe)
		AssociationFieldToForm("Tone", instanceWithInferedType.Tone, formGroup, probe)
		AssociationFieldToForm("Tree", instanceWithInferedType.Tree, formGroup, probe)
		BasicFieldtoForm("HasDiv", instanceWithInferedType.HasDiv, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DivStyle", instanceWithInferedType.DivStyle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "AsSplit"
			rf.Fieldname = "AsSplitAreas"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.AsSplit),
					"AsSplitAreas",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.AsSplit, *models.AsSplitArea](
					nil,
					"AsSplitAreas",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "View"
			rf.Fieldname = "RootAsSplitAreas"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.View),
					"RootAsSplitAreas",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.View, *models.AsSplitArea](
					nil,
					"RootAsSplitAreas",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Button:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StackName", instanceWithInferedType.StackName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Cursor:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StackName", instanceWithInferedType.StackName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Style", instanceWithInferedType.Style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Doc:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StackName", instanceWithInferedType.StackName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Form:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StackName", instanceWithInferedType.StackName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FormName", instanceWithInferedType.FormName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Load:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StackName", instanceWithInferedType.StackName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Slider:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StackName", instanceWithInferedType.StackName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Split:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StackName", instanceWithInferedType.StackName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Svg:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StackName", instanceWithInferedType.StackName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Style", instanceWithInferedType.Style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Table:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StackName", instanceWithInferedType.StackName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TableName", instanceWithInferedType.TableName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Tone:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StackName", instanceWithInferedType.StackName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Tree:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StackName", instanceWithInferedType.StackName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TreeName", instanceWithInferedType.TreeName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.View:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("RootAsSplitAreas", instanceWithInferedType, &instanceWithInferedType.RootAsSplitAreas, formGroup, probe)

	default:
		_ = instanceWithInferedType
	}
}
