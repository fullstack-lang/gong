// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/dsme/project/go/models"
)

// updateFillUpForm updates the current form if there is one
func (probe *Probe) updateFillUpForm() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *DiagramFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Diagram", true)
			} else {
				FillUpFormFromGongstruct(onSave.diagram, probe)
			}
		case *NoteFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Note", true)
			} else {
				FillUpFormFromGongstruct(onSave.note, probe)
			}
		case *NoteProductShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "NoteProductShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.noteproductshape, probe)
			}
		case *NoteShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "NoteShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.noteshape, probe)
			}
		case *NoteTaskShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "NoteTaskShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.notetaskshape, probe)
			}
		case *ProductFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Product", true)
			} else {
				FillUpFormFromGongstruct(onSave.product, probe)
			}
		case *ProductCompositionShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ProductCompositionShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.productcompositionshape, probe)
			}
		case *ProductShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ProductShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.productshape, probe)
			}
		case *ProjectFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Project", true)
			} else {
				FillUpFormFromGongstruct(onSave.project, probe)
			}
		case *ResourceFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Resource", true)
			} else {
				FillUpFormFromGongstruct(onSave.resource, probe)
			}
		case *ResourceShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ResourceShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.resourceshape, probe)
			}
		case *ResourceTaskShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ResourceTaskShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.resourcetaskshape, probe)
			}
		case *RootFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Root", true)
			} else {
				FillUpFormFromGongstruct(onSave.root, probe)
			}
		case *TaskFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Task", true)
			} else {
				FillUpFormFromGongstruct(onSave.task, probe)
			}
		case *TaskCompositionShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TaskCompositionShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.taskcompositionshape, probe)
			}
		case *TaskInputShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TaskInputShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.taskinputshape, probe)
			}
		case *TaskOutputShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TaskOutputShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.taskoutputshape, probe)
			}
		case *TaskShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TaskShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.taskshape, probe)
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
	case "Note":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Note Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteFormCallback(
			nil,
			probe,
			formGroup,
		)
		note := new(models.Note)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(note, formGroup, probe)
	case "NoteProductShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "NoteProductShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteProductShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		noteproductshape := new(models.NoteProductShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(noteproductshape, formGroup, probe)
	case "NoteShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "NoteShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		noteshape := new(models.NoteShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(noteshape, formGroup, probe)
	case "NoteTaskShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "NoteTaskShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteTaskShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		notetaskshape := new(models.NoteTaskShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(notetaskshape, formGroup, probe)
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
	case "ProductCompositionShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ProductCompositionShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProductCompositionShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		productcompositionshape := new(models.ProductCompositionShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(productcompositionshape, formGroup, probe)
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
	case "Resource":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Resource Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ResourceFormCallback(
			nil,
			probe,
			formGroup,
		)
		resource := new(models.Resource)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(resource, formGroup, probe)
	case "ResourceShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ResourceShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ResourceShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		resourceshape := new(models.ResourceShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(resourceshape, formGroup, probe)
	case "ResourceTaskShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ResourceTaskShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ResourceTaskShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		resourcetaskshape := new(models.ResourceTaskShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(resourcetaskshape, formGroup, probe)
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
	case "TaskCompositionShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TaskCompositionShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskCompositionShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		taskcompositionshape := new(models.TaskCompositionShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(taskcompositionshape, formGroup, probe)
	case "TaskInputShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TaskInputShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskInputShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		taskinputshape := new(models.TaskInputShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(taskinputshape, formGroup, probe)
	case "TaskOutputShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TaskOutputShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskOutputShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		taskoutputshape := new(models.TaskOutputShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(taskoutputshape, formGroup, probe)
	case "TaskShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TaskShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		taskshape := new(models.TaskShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(taskshape, formGroup, probe)
	}
	formStage.Commit()
}
