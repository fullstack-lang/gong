// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/button/go/models"
)

// updateFillUpForm updates the current form if there is one
func (probe *Probe) updateFillUpForm() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *ButtonFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Button", true)
			} else {
				FillUpFormFromGongstruct(onSave.button, probe)
			}
		case *ButtonToggleFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ButtonToggle", true)
			} else {
				FillUpFormFromGongstruct(onSave.buttontoggle, probe)
			}
		case *GroupFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Group", true)
			} else {
				FillUpFormFromGongstruct(onSave.group, probe)
			}
		case *GroupToogleFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "GroupToogle", true)
			} else {
				FillUpFormFromGongstruct(onSave.grouptoogle, probe)
			}
		case *LayoutFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Layout", true)
			} else {
				FillUpFormFromGongstruct(onSave.layout, probe)
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
	case "Button":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Button Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ButtonFormCallback(
			nil,
			probe,
			formGroup,
		)
		button := new(models.Button)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(button, formGroup, probe)
	case "ButtonToggle":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ButtonToggle Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ButtonToggleFormCallback(
			nil,
			probe,
			formGroup,
		)
		buttontoggle := new(models.ButtonToggle)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(buttontoggle, formGroup, probe)
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
	case "GroupToogle":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GroupToogle Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GroupToogleFormCallback(
			nil,
			probe,
			formGroup,
		)
		grouptoogle := new(models.GroupToogle)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(grouptoogle, formGroup, probe)
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
	}
	formStage.Commit()
}
