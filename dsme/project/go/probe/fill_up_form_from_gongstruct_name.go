// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/dsme/project/go/models"
)

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
	case "Diagram":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Diagram Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramFormCallback(
			nil,
			probe,
			formGroup,
		)
		diagram := new(models.Diagram)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(diagram, formGroup, probe)
	case "Product":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Product Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProductFormCallback(
			nil,
			probe,
			formGroup,
		)
		product := new(models.Product)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(product, formGroup, probe)
	case "ProductShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ProductShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProductShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		productshape := new(models.ProductShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(productshape, formGroup, probe)
	case "Project":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Project Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProjectFormCallback(
			nil,
			probe,
			formGroup,
		)
		project := new(models.Project)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(project, formGroup, probe)
	case "Root":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Root Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RootFormCallback(
			nil,
			probe,
			formGroup,
		)
		root := new(models.Root)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(root, formGroup, probe)
	case "Task":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Task Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskFormCallback(
			nil,
			probe,
			formGroup,
		)
		task := new(models.Task)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(task, formGroup, probe)
	}
	formStage.Commit()
}
