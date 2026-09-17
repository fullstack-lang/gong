// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		if onSave, ok := formGroup.OnSave.(FormCallbackIF); ok {
			if onSave.GetCreationMode() {
				FillUpFormFromGongstructName(probe, onSave.GetGongstructName(), true)
			} else {
				FillUpFormFromGongstruct(onSave.GetInstance(), probe)
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
	case "Menu":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Menu Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MenuFormCallback(
			nil,
			probe,
			formGroup,
		)
		menu := new(models.Menu)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(menu, formGroup, probe)
	case "Node":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Node Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NodeFormCallback(
			nil,
			probe,
			formGroup,
		)
		node := new(models.Node)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(node, formGroup, probe)
	case "SVGIcon":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SVGIcon Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SVGIconFormCallback(
			nil,
			probe,
			formGroup,
		)
		svgicon := new(models.SVGIcon)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(svgicon, formGroup, probe)
	case "Tree":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Tree Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TreeFormCallback(
			nil,
			probe,
			formGroup,
		)
		tree := new(models.Tree)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tree, formGroup, probe)
	}
	formStage.Commit()
}
