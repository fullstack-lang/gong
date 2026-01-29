// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/slider/go/models"
)

// updateFillUpForm updates the current form if there is one
func (probe *Probe) updateFillUpForm() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *CheckboxFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Checkbox", true)
			} else {
				FillUpFormFromGongstruct(onSave.checkbox, probe)
			}
		case *GroupFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Group", true)
			} else {
				FillUpFormFromGongstruct(onSave.group, probe)
			}
		case *LayoutFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Layout", true)
			} else {
				FillUpFormFromGongstruct(onSave.layout, probe)
			}
		case *SliderFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Slider", true)
			} else {
				FillUpFormFromGongstruct(onSave.slider, probe)
			}
		}
	}
}

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Checkbox":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Checkbox Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CheckboxFormCallback(
			nil,
			probe,
			formGroup,
		)
		checkbox := new(models.Checkbox)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(checkbox, formGroup, probe)
	case "Group":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Group Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GroupFormCallback(
			nil,
			probe,
			formGroup,
		)
		group := new(models.Group)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(group, formGroup, probe)
	case "Layout":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Layout Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LayoutFormCallback(
			nil,
			probe,
			formGroup,
		)
		layout := new(models.Layout)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(layout, formGroup, probe)
	case "Slider":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Slider Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SliderFormCallback(
			nil,
			probe,
			formGroup,
		)
		slider := new(models.Slider)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(slider, formGroup, probe)
	}
	formStage.Commit()
}
