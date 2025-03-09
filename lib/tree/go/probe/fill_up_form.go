// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/models"
	"github.com/fullstack-lang/gong/lib/tree/go/orm"
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
	case *models.Button:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Icon", instanceWithInferedType.Icon, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("SVGIcon", instanceWithInferedType.SVGIcon, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Node"
			rf.Fieldname = "Buttons"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Node),
					"Buttons",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Node, *models.Button](
					nil,
					"Buttons",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Node:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("FontStyle", instanceWithInferedType.FontStyle, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("BackgroundColor", instanceWithInferedType.BackgroundColor, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasCheckboxButton", instanceWithInferedType.HasCheckboxButton, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsCheckboxDisabled", instanceWithInferedType.IsCheckboxDisabled, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasSecondCheckboxButton", instanceWithInferedType.HasSecondCheckboxButton, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsSecondCheckboxChecked", instanceWithInferedType.IsSecondCheckboxChecked, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsSecondCheckboxDisabled", instanceWithInferedType.IsSecondCheckboxDisabled, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TextAfterSecondCheckbox", instanceWithInferedType.TextAfterSecondCheckbox, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsInEditMode", instanceWithInferedType.IsInEditMode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsNodeClickable", instanceWithInferedType.IsNodeClickable, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsWithPreceedingIcon", instanceWithInferedType.IsWithPreceedingIcon, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("PreceedingIcon", instanceWithInferedType.PreceedingIcon, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("PreceedingSVGIcon", instanceWithInferedType.PreceedingSVGIcon, formGroup, probe)
		AssociationSliceToForm("Children", instanceWithInferedType, &instanceWithInferedType.Children, formGroup, probe)
		AssociationSliceToForm("Buttons", instanceWithInferedType, &instanceWithInferedType.Buttons, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Node"
			rf.Fieldname = "Children"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Node),
					"Children",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Node, *models.Node](
					nil,
					"Children",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Tree"
			rf.Fieldname = "RootNodes"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Tree),
					"RootNodes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Tree, *models.Node](
					nil,
					"RootNodes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SVGIcon:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, false, 0)
		BasicFieldtoForm("SVG", instanceWithInferedType.SVG, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, true, 300)

	case *models.Tree:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("RootNodes", instanceWithInferedType, &instanceWithInferedType.RootNodes, formGroup, probe)

	default:
		_ = instanceWithInferedType
	}
}
