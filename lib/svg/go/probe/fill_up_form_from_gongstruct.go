// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *gongtable.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Animate:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Animate Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AnimateFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Circle:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Circle Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CircleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Ellipse:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Ellipse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EllipseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Layer:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Layer Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LayerFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Line:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Line Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Link:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Link Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.LinkAnchoredText:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "LinkAnchoredText Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkAnchoredTextFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Path:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Path Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PathFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Point:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Point Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PointFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Polygone:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Polygone Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PolygoneFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Polyline:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Polyline Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PolylineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Rect:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Rect Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RectFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RectAnchoredPath:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "RectAnchoredPath Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RectAnchoredPathFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RectAnchoredRect:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "RectAnchoredRect Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RectAnchoredRectFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RectAnchoredText:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "RectAnchoredText Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RectAnchoredTextFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RectLinkLink:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "RectLinkLink Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RectLinkLinkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SVG:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SVG Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SVGFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SvgText:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SvgText Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SvgTextFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Text:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Text Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TextFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
