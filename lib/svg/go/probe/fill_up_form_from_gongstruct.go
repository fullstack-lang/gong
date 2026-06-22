// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Animate:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Animate",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AnimateFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Circle:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Circle",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CircleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Condition:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Condition",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ConditionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ControlPoint:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ControlPoint",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ControlPointFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Ellipse:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Ellipse",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EllipseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FileToDownload:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "FileToDownload",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FileToDownloadFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Layer:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Layer",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LayerFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Line:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Line",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Link:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Link",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.LinkAnchoredPath:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "LinkAnchoredPath",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkAnchoredPathFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.LinkAnchoredText:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "LinkAnchoredText",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkAnchoredTextFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Path:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Path",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PathFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Point:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Point",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PointFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Polygone:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Polygone",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PolygoneFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Polyline:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Polyline",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PolylineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Rect:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Rect",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RectFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RectAnchoredPath:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "RectAnchoredPath",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RectAnchoredPathFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RectAnchoredPngImage:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "RectAnchoredPngImage",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RectAnchoredPngImageFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RectAnchoredRect:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "RectAnchoredRect",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RectAnchoredRectFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RectAnchoredText:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "RectAnchoredText",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RectAnchoredTextFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RectLinkLink:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "RectLinkLink",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RectLinkLinkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SVG:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "SVG",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SVGFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SvgText:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "SvgText",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SvgTextFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Text:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Text",
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

	if probe.GetCommitMode() {
		formStage.Commit()
	}
}
