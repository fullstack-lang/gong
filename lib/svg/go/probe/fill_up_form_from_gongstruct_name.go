// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
)

// updateFillUpForm updates the current form if there is one
func (probe *Probe) updateFillUpForm() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *AnimateFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Animate", true)
			} else {
				FillUpFormFromGongstruct(onSave.animate, probe)
			}
		case *CircleFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Circle", true)
			} else {
				FillUpFormFromGongstruct(onSave.circle, probe)
			}
		case *ConditionFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Condition", true)
			} else {
				FillUpFormFromGongstruct(onSave.condition, probe)
			}
		case *ControlPointFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ControlPoint", true)
			} else {
				FillUpFormFromGongstruct(onSave.controlpoint, probe)
			}
		case *EllipseFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Ellipse", true)
			} else {
				FillUpFormFromGongstruct(onSave.ellipse, probe)
			}
		case *LayerFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Layer", true)
			} else {
				FillUpFormFromGongstruct(onSave.layer, probe)
			}
		case *LineFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Line", true)
			} else {
				FillUpFormFromGongstruct(onSave.line, probe)
			}
		case *LinkFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Link", true)
			} else {
				FillUpFormFromGongstruct(onSave.link, probe)
			}
		case *LinkAnchoredTextFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "LinkAnchoredText", true)
			} else {
				FillUpFormFromGongstruct(onSave.linkanchoredtext, probe)
			}
		case *PathFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Path", true)
			} else {
				FillUpFormFromGongstruct(onSave.path, probe)
			}
		case *PointFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Point", true)
			} else {
				FillUpFormFromGongstruct(onSave.point, probe)
			}
		case *PolygoneFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Polygone", true)
			} else {
				FillUpFormFromGongstruct(onSave.polygone, probe)
			}
		case *PolylineFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Polyline", true)
			} else {
				FillUpFormFromGongstruct(onSave.polyline, probe)
			}
		case *RectFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Rect", true)
			} else {
				FillUpFormFromGongstruct(onSave.rect, probe)
			}
		case *RectAnchoredPathFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "RectAnchoredPath", true)
			} else {
				FillUpFormFromGongstruct(onSave.rectanchoredpath, probe)
			}
		case *RectAnchoredRectFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "RectAnchoredRect", true)
			} else {
				FillUpFormFromGongstruct(onSave.rectanchoredrect, probe)
			}
		case *RectAnchoredTextFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "RectAnchoredText", true)
			} else {
				FillUpFormFromGongstruct(onSave.rectanchoredtext, probe)
			}
		case *RectLinkLinkFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "RectLinkLink", true)
			} else {
				FillUpFormFromGongstruct(onSave.rectlinklink, probe)
			}
		case *SVGFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SVG", true)
			} else {
				FillUpFormFromGongstruct(onSave.svg, probe)
			}
		case *SvgTextFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SvgText", true)
			} else {
				FillUpFormFromGongstruct(onSave.svgtext, probe)
			}
		case *TextFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Text", true)
			} else {
				FillUpFormFromGongstruct(onSave.text, probe)
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
	case "Animate":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Animate Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AnimateFormCallback(
			nil,
			probe,
			formGroup,
		)
		animate := new(models.Animate)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(animate, formGroup, probe)
	case "Circle":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Circle Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CircleFormCallback(
			nil,
			probe,
			formGroup,
		)
		circle := new(models.Circle)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(circle, formGroup, probe)
	case "Condition":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Condition Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ConditionFormCallback(
			nil,
			probe,
			formGroup,
		)
		condition := new(models.Condition)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(condition, formGroup, probe)
	case "ControlPoint":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ControlPoint Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ControlPointFormCallback(
			nil,
			probe,
			formGroup,
		)
		controlpoint := new(models.ControlPoint)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(controlpoint, formGroup, probe)
	case "Ellipse":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Ellipse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EllipseFormCallback(
			nil,
			probe,
			formGroup,
		)
		ellipse := new(models.Ellipse)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(ellipse, formGroup, probe)
	case "Layer":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Layer Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LayerFormCallback(
			nil,
			probe,
			formGroup,
		)
		layer := new(models.Layer)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(layer, formGroup, probe)
	case "Line":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Line Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LineFormCallback(
			nil,
			probe,
			formGroup,
		)
		line := new(models.Line)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(line, formGroup, probe)
	case "Link":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Link Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkFormCallback(
			nil,
			probe,
			formGroup,
		)
		link := new(models.Link)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(link, formGroup, probe)
	case "LinkAnchoredText":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "LinkAnchoredText Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkAnchoredTextFormCallback(
			nil,
			probe,
			formGroup,
		)
		linkanchoredtext := new(models.LinkAnchoredText)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(linkanchoredtext, formGroup, probe)
	case "Path":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Path Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PathFormCallback(
			nil,
			probe,
			formGroup,
		)
		path := new(models.Path)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(path, formGroup, probe)
	case "Point":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Point Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PointFormCallback(
			nil,
			probe,
			formGroup,
		)
		point := new(models.Point)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(point, formGroup, probe)
	case "Polygone":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Polygone Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PolygoneFormCallback(
			nil,
			probe,
			formGroup,
		)
		polygone := new(models.Polygone)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(polygone, formGroup, probe)
	case "Polyline":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Polyline Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PolylineFormCallback(
			nil,
			probe,
			formGroup,
		)
		polyline := new(models.Polyline)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(polyline, formGroup, probe)
	case "Rect":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Rect Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RectFormCallback(
			nil,
			probe,
			formGroup,
		)
		rect := new(models.Rect)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rect, formGroup, probe)
	case "RectAnchoredPath":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "RectAnchoredPath Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RectAnchoredPathFormCallback(
			nil,
			probe,
			formGroup,
		)
		rectanchoredpath := new(models.RectAnchoredPath)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rectanchoredpath, formGroup, probe)
	case "RectAnchoredRect":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "RectAnchoredRect Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RectAnchoredRectFormCallback(
			nil,
			probe,
			formGroup,
		)
		rectanchoredrect := new(models.RectAnchoredRect)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rectanchoredrect, formGroup, probe)
	case "RectAnchoredText":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "RectAnchoredText Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RectAnchoredTextFormCallback(
			nil,
			probe,
			formGroup,
		)
		rectanchoredtext := new(models.RectAnchoredText)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rectanchoredtext, formGroup, probe)
	case "RectLinkLink":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "RectLinkLink Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RectLinkLinkFormCallback(
			nil,
			probe,
			formGroup,
		)
		rectlinklink := new(models.RectLinkLink)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rectlinklink, formGroup, probe)
	case "SVG":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SVG Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SVGFormCallback(
			nil,
			probe,
			formGroup,
		)
		svg := new(models.SVG)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(svg, formGroup, probe)
	case "SvgText":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SvgText Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SvgTextFormCallback(
			nil,
			probe,
			formGroup,
		)
		svgtext := new(models.SvgText)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(svgtext, formGroup, probe)
	case "Text":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Text Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TextFormCallback(
			nil,
			probe,
			formGroup,
		)
		text := new(models.Text)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(text, formGroup, probe)
	}
	formStage.Commit()
}
