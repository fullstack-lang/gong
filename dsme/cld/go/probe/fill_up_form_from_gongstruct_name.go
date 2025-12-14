// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/dsme/cld/go/models"
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
	case "Category1":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Category1 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Category1FormCallback(
			nil,
			probe,
			formGroup,
		)
		category1 := new(models.Category1)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(category1, formGroup, probe)
	case "Category1Shape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Category1Shape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Category1ShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		category1shape := new(models.Category1Shape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(category1shape, formGroup, probe)
	case "Category2":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Category2 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Category2FormCallback(
			nil,
			probe,
			formGroup,
		)
		category2 := new(models.Category2)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(category2, formGroup, probe)
	case "Category2Shape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Category2Shape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Category2ShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		category2shape := new(models.Category2Shape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(category2shape, formGroup, probe)
	case "Category3":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Category3 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Category3FormCallback(
			nil,
			probe,
			formGroup,
		)
		category3 := new(models.Category3)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(category3, formGroup, probe)
	case "Category3Shape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Category3Shape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Category3ShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		category3shape := new(models.Category3Shape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(category3shape, formGroup, probe)
	case "ControlPointShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ControlPointShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ControlPointShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		controlpointshape := new(models.ControlPointShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(controlpointshape, formGroup, probe)
	case "Desk":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Desk Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DeskFormCallback(
			nil,
			probe,
			formGroup,
		)
		desk := new(models.Desk)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(desk, formGroup, probe)
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
	case "Influence":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Influence Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InfluenceFormCallback(
			nil,
			probe,
			formGroup,
		)
		influence := new(models.Influence)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(influence, formGroup, probe)
	case "InfluenceShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "InfluenceShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InfluenceShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		influenceshape := new(models.InfluenceShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(influenceshape, formGroup, probe)
	}
	formStage.Commit()
}
