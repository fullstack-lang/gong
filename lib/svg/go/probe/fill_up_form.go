// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Animate:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("AttributeName", instanceWithInferedType.AttributeName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Values", instanceWithInferedType.Values, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("From", instanceWithInferedType.From, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("To", instanceWithInferedType.To, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Dur", instanceWithInferedType.Dur, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("RepeatCount", instanceWithInferedType.RepeatCount, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Circle, *models.Animate](
				"Circle",
				"Animations",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Circle) []*models.Animate {
					return owner.Animations
				})
		}
		{
			AssociationReverseSliceToForm[*models.Ellipse, *models.Animate](
				"Ellipse",
				"Animates",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Ellipse) []*models.Animate {
					return owner.Animates
				})
		}
		{
			AssociationReverseSliceToForm[*models.Line, *models.Animate](
				"Line",
				"Animates",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Line) []*models.Animate {
					return owner.Animates
				})
		}
		{
			AssociationReverseSliceToForm[*models.LinkAnchoredText, *models.Animate](
				"LinkAnchoredText",
				"Animates",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.LinkAnchoredText) []*models.Animate {
					return owner.Animates
				})
		}
		{
			AssociationReverseSliceToForm[*models.Path, *models.Animate](
				"Path",
				"Animates",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Path) []*models.Animate {
					return owner.Animates
				})
		}
		{
			AssociationReverseSliceToForm[*models.Polygone, *models.Animate](
				"Polygone",
				"Animates",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Polygone) []*models.Animate {
					return owner.Animates
				})
		}
		{
			AssociationReverseSliceToForm[*models.Polyline, *models.Animate](
				"Polyline",
				"Animates",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Polyline) []*models.Animate {
					return owner.Animates
				})
		}
		{
			AssociationReverseSliceToForm[*models.Rect, *models.Animate](
				"Rect",
				"Animations",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Rect) []*models.Animate {
					return owner.Animations
				})
		}
		{
			AssociationReverseSliceToForm[*models.RectAnchoredText, *models.Animate](
				"RectAnchoredText",
				"Animates",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.RectAnchoredText) []*models.Animate {
					return owner.Animates
				})
		}
		{
			AssociationReverseSliceToForm[*models.Text, *models.Animate](
				"Text",
				"Animates",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Text) []*models.Animate {
					return owner.Animates
				})
		}

	case *models.Circle:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CX", instanceWithInferedType.CX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CY", instanceWithInferedType.CY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Radius", instanceWithInferedType.Radius, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		AssociationSliceToForm("Animations", instanceWithInferedType, &instanceWithInferedType.Animations, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Layer, *models.Circle](
				"Layer",
				"Circles",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Layer) []*models.Circle {
					return owner.Circles
				})
		}

	case *models.Condition:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Rect, *models.Condition](
				"Rect",
				"HoveringTrigger",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Rect) []*models.Condition {
					return owner.HoveringTrigger
				})
		}
		{
			AssociationReverseSliceToForm[*models.Rect, *models.Condition](
				"Rect",
				"DisplayConditions",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Rect) []*models.Condition {
					return owner.DisplayConditions
				})
		}

	case *models.ControlPoint:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X_Relative", instanceWithInferedType.X_Relative, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y_Relative", instanceWithInferedType.Y_Relative, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ClosestRect", instanceWithInferedType.ClosestRect, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Link, *models.ControlPoint](
				"Link",
				"ControlPoints",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Link) []*models.ControlPoint {
					return owner.ControlPoints
				})
		}

	case *models.Ellipse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CX", instanceWithInferedType.CX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CY", instanceWithInferedType.CY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("RX", instanceWithInferedType.RX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("RY", instanceWithInferedType.RY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Layer, *models.Ellipse](
				"Layer",
				"Ellipses",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Layer) []*models.Ellipse {
					return owner.Ellipses
				})
		}

	case *models.FileToDownload:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, false, 0)
		BasicFieldtoForm("Base64EncodedContent", instanceWithInferedType.Base64EncodedContent, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 2000)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Layer:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Rects", instanceWithInferedType, &instanceWithInferedType.Rects, formGroup, probe)
		AssociationSliceToForm("Texts", instanceWithInferedType, &instanceWithInferedType.Texts, formGroup, probe)
		AssociationSliceToForm("Circles", instanceWithInferedType, &instanceWithInferedType.Circles, formGroup, probe)
		AssociationSliceToForm("Lines", instanceWithInferedType, &instanceWithInferedType.Lines, formGroup, probe)
		AssociationSliceToForm("Ellipses", instanceWithInferedType, &instanceWithInferedType.Ellipses, formGroup, probe)
		AssociationSliceToForm("Polylines", instanceWithInferedType, &instanceWithInferedType.Polylines, formGroup, probe)
		AssociationSliceToForm("Polygones", instanceWithInferedType, &instanceWithInferedType.Polygones, formGroup, probe)
		AssociationSliceToForm("Paths", instanceWithInferedType, &instanceWithInferedType.Paths, formGroup, probe)
		AssociationSliceToForm("Links", instanceWithInferedType, &instanceWithInferedType.Links, formGroup, probe)
		AssociationSliceToForm("RectLinkLinks", instanceWithInferedType, &instanceWithInferedType.RectLinkLinks, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.SVG, *models.Layer](
				"SVG",
				"Layers",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.SVG) []*models.Layer {
					return owner.Layers
				})
		}

	case *models.Line:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X1", instanceWithInferedType.X1, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y1", instanceWithInferedType.Y1, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X2", instanceWithInferedType.X2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y2", instanceWithInferedType.Y2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, probe)
		BasicFieldtoForm("MouseClickX", instanceWithInferedType.MouseClickX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MouseClickY", instanceWithInferedType.MouseClickY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Layer, *models.Line](
				"Layer",
				"Lines",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Layer) []*models.Line {
					return owner.Lines
				})
		}

	case *models.Link:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("IsBezierCurve", instanceWithInferedType.IsBezierCurve, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Start", instanceWithInferedType.Start, formGroup, probe)
		EnumTypeStringToForm("StartAnchorType", instanceWithInferedType.StartAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		AssociationFieldToForm("End", instanceWithInferedType.End, formGroup, probe)
		EnumTypeStringToForm("EndAnchorType", instanceWithInferedType.EndAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CornerRadius", instanceWithInferedType.CornerRadius, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasEndArrow", instanceWithInferedType.HasEndArrow, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndArrowSize", instanceWithInferedType.EndArrowSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndArrowOffset", instanceWithInferedType.EndArrowOffset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasStartArrow", instanceWithInferedType.HasStartArrow, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StartArrowSize", instanceWithInferedType.StartArrowSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StartArrowOffset", instanceWithInferedType.StartArrowOffset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("TextAtArrowStart", instanceWithInferedType, &instanceWithInferedType.TextAtArrowStart, formGroup, probe)
		AssociationSliceToForm("TextAtArrowEnd", instanceWithInferedType, &instanceWithInferedType.TextAtArrowEnd, formGroup, probe)
		AssociationSliceToForm("TextAtCorner", instanceWithInferedType, &instanceWithInferedType.TextAtCorner, formGroup, probe)
		AssociationSliceToForm("PathAtArrowStart", instanceWithInferedType, &instanceWithInferedType.PathAtArrowStart, formGroup, probe)
		AssociationSliceToForm("PathAtArrowEnd", instanceWithInferedType, &instanceWithInferedType.PathAtArrowEnd, formGroup, probe)
		AssociationSliceToForm("PathAtCorner", instanceWithInferedType, &instanceWithInferedType.PathAtCorner, formGroup, probe)
		AssociationSliceToForm("ControlPoints", instanceWithInferedType, &instanceWithInferedType.ControlPoints, formGroup, probe)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		BasicFieldtoForm("MouseX", instanceWithInferedType.MouseX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MouseY", instanceWithInferedType.MouseY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("MouseEventKey", instanceWithInferedType.MouseEventKey, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Layer, *models.Link](
				"Layer",
				"Links",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Layer) []*models.Link {
					return owner.Links
				})
		}

	case *models.LinkAnchoredPath:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Definition", instanceWithInferedType.Definition, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X_Offset", instanceWithInferedType.X_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y_Offset", instanceWithInferedType.Y_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ScalePropotionnally", instanceWithInferedType.ScalePropotionnally, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("AppliedScaling", instanceWithInferedType.AppliedScaling, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Link, *models.LinkAnchoredPath](
				"Link",
				"PathAtArrowStart",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Link) []*models.LinkAnchoredPath {
					return owner.PathAtArrowStart
				})
		}
		{
			AssociationReverseSliceToForm[*models.Link, *models.LinkAnchoredPath](
				"Link",
				"PathAtArrowEnd",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Link) []*models.LinkAnchoredPath {
					return owner.PathAtArrowEnd
				})
		}
		{
			AssociationReverseSliceToForm[*models.Link, *models.LinkAnchoredPath](
				"Link",
				"PathAtCorner",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Link) []*models.LinkAnchoredPath {
					return owner.PathAtCorner
				})
		}

	case *models.LinkAnchoredText:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		BasicFieldtoForm("AutomaticLayout", instanceWithInferedType.AutomaticLayout, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("LinkAnchorType", instanceWithInferedType.LinkAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("X_Offset", instanceWithInferedType.X_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y_Offset", instanceWithInferedType.Y_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FontWeight", instanceWithInferedType.FontWeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FontSize", instanceWithInferedType.FontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FontStyle", instanceWithInferedType.FontStyle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LetterSpacing", instanceWithInferedType.LetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FontFamily", instanceWithInferedType.FontFamily, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("WhiteSpace", instanceWithInferedType.WhiteSpace, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Link, *models.LinkAnchoredText](
				"Link",
				"TextAtArrowStart",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Link) []*models.LinkAnchoredText {
					return owner.TextAtArrowStart
				})
		}
		{
			AssociationReverseSliceToForm[*models.Link, *models.LinkAnchoredText](
				"Link",
				"TextAtArrowEnd",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Link) []*models.LinkAnchoredText {
					return owner.TextAtArrowEnd
				})
		}
		{
			AssociationReverseSliceToForm[*models.Link, *models.LinkAnchoredText](
				"Link",
				"TextAtCorner",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Link) []*models.LinkAnchoredText {
					return owner.TextAtCorner
				})
		}

	case *models.Path:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Definition", instanceWithInferedType.Definition, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Layer, *models.Path](
				"Layer",
				"Paths",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Layer) []*models.Path {
					return owner.Paths
				})
		}

	case *models.Point:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Polygone:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Points", instanceWithInferedType.Points, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Layer, *models.Polygone](
				"Layer",
				"Polygones",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Layer) []*models.Polygone {
					return owner.Polygones
				})
		}

	case *models.Polyline:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Points", instanceWithInferedType.Points, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Layer, *models.Polyline](
				"Layer",
				"Polylines",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Layer) []*models.Polyline {
					return owner.Polylines
				})
		}

	case *models.Rect:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("RX", instanceWithInferedType.RX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Peers", instanceWithInferedType, &instanceWithInferedType.Peers, formGroup, probe)
		AssociationFieldToForm("EnclosingRect", instanceWithInferedType.EnclosingRect, formGroup, probe)
		AssociationSliceToForm("Obstacles", instanceWithInferedType, &instanceWithInferedType.Obstacles, formGroup, probe)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		AssociationSliceToForm("HoveringTrigger", instanceWithInferedType, &instanceWithInferedType.HoveringTrigger, formGroup, probe)
		AssociationSliceToForm("DisplayConditions", instanceWithInferedType, &instanceWithInferedType.DisplayConditions, formGroup, probe)
		AssociationSliceToForm("Animations", instanceWithInferedType, &instanceWithInferedType.Animations, formGroup, probe)
		BasicFieldtoForm("IsSelectable", instanceWithInferedType.IsSelectable, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsSelected", instanceWithInferedType.IsSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CanHaveLeftHandle", instanceWithInferedType.CanHaveLeftHandle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasLeftHandle", instanceWithInferedType.HasLeftHandle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CanHaveRightHandle", instanceWithInferedType.CanHaveRightHandle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasRightHandle", instanceWithInferedType.HasRightHandle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CanHaveTopHandle", instanceWithInferedType.CanHaveTopHandle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasTopHandle", instanceWithInferedType.HasTopHandle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsScalingProportionally", instanceWithInferedType.IsScalingProportionally, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CanHaveBottomHandle", instanceWithInferedType.CanHaveBottomHandle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasBottomHandle", instanceWithInferedType.HasBottomHandle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CanMoveHorizontaly", instanceWithInferedType.CanMoveHorizontaly, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CanMoveVerticaly", instanceWithInferedType.CanMoveVerticaly, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("RectAnchoredTexts", instanceWithInferedType, &instanceWithInferedType.RectAnchoredTexts, formGroup, probe)
		AssociationSliceToForm("RectAnchoredRects", instanceWithInferedType, &instanceWithInferedType.RectAnchoredRects, formGroup, probe)
		AssociationSliceToForm("RectAnchoredPaths", instanceWithInferedType, &instanceWithInferedType.RectAnchoredPaths, formGroup, probe)
		AssociationSliceToForm("RectAnchoredPngImages", instanceWithInferedType, &instanceWithInferedType.RectAnchoredPngImages, formGroup, probe)
		BasicFieldtoForm("ChangeColorWhenHovered", instanceWithInferedType.ChangeColorWhenHovered, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ColorWhenHovered", instanceWithInferedType.ColorWhenHovered, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OriginalColor", instanceWithInferedType.OriginalColor, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacityWhenHovered", instanceWithInferedType.FillOpacityWhenHovered, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OriginalFillOpacity", instanceWithInferedType.OriginalFillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasToolTip", instanceWithInferedType.HasToolTip, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ToolTipText", instanceWithInferedType.ToolTipText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("ToolTipPosition", instanceWithInferedType.ToolTipPosition, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("MouseX", instanceWithInferedType.MouseX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MouseY", instanceWithInferedType.MouseY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("MouseEventKey", instanceWithInferedType.MouseEventKey, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("URLPath", instanceWithInferedType.URLPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("URLTarget", instanceWithInferedType.URLTarget, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Layer, *models.Rect](
				"Layer",
				"Rects",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Layer) []*models.Rect {
					return owner.Rects
				})
		}
		{
			AssociationReverseSliceToForm[*models.Rect, *models.Rect](
				"Rect",
				"Peers",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Rect) []*models.Rect {
					return owner.Peers
				})
		}
		{
			AssociationReverseSliceToForm[*models.Rect, *models.Rect](
				"Rect",
				"Obstacles",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Rect) []*models.Rect {
					return owner.Obstacles
				})
		}

	case *models.RectAnchoredPath:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Definition", instanceWithInferedType.Definition, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X_Offset", instanceWithInferedType.X_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y_Offset", instanceWithInferedType.Y_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("RectAnchorType", instanceWithInferedType.RectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("ScalePropotionnally", instanceWithInferedType.ScalePropotionnally, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("AppliedScaling", instanceWithInferedType.AppliedScaling, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Rect, *models.RectAnchoredPath](
				"Rect",
				"RectAnchoredPaths",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Rect) []*models.RectAnchoredPath {
					return owner.RectAnchoredPaths
				})
		}

	case *models.RectAnchoredPngImage:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("RX", instanceWithInferedType.RX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X_Offset", instanceWithInferedType.X_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y_Offset", instanceWithInferedType.Y_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("RectAnchorType", instanceWithInferedType.RectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Base64Content", instanceWithInferedType.Base64Content, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Rect, *models.RectAnchoredPngImage](
				"Rect",
				"RectAnchoredPngImages",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Rect) []*models.RectAnchoredPngImage {
					return owner.RectAnchoredPngImages
				})
		}

	case *models.RectAnchoredRect:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("RX", instanceWithInferedType.RX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X_Offset", instanceWithInferedType.X_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y_Offset", instanceWithInferedType.Y_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("RectAnchorType", instanceWithInferedType.RectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("WidthFollowRect", instanceWithInferedType.WidthFollowRect, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HeightFollowRect", instanceWithInferedType.HeightFollowRect, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasToolTip", instanceWithInferedType.HasToolTip, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ToolTipText", instanceWithInferedType.ToolTipText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Rect, *models.RectAnchoredRect](
				"Rect",
				"RectAnchoredRects",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Rect) []*models.RectAnchoredRect {
					return owner.RectAnchoredRects
				})
		}

	case *models.RectAnchoredText:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		BasicFieldtoForm("FontWeight", instanceWithInferedType.FontWeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FontSize", instanceWithInferedType.FontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FontStyle", instanceWithInferedType.FontStyle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LetterSpacing", instanceWithInferedType.LetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FontFamily", instanceWithInferedType.FontFamily, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("WhiteSpace", instanceWithInferedType.WhiteSpace, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("X_Offset", instanceWithInferedType.X_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y_Offset", instanceWithInferedType.Y_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("RectAnchorType", instanceWithInferedType.RectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("TextAnchorType", instanceWithInferedType.TextAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("DominantBaseline", instanceWithInferedType.DominantBaseline, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("WritingMode", instanceWithInferedType.WritingMode, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, probe)
		BasicFieldtoForm("URLPath", instanceWithInferedType.URLPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("URLTarget", instanceWithInferedType.URLTarget, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Rect, *models.RectAnchoredText](
				"Rect",
				"RectAnchoredTexts",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Rect) []*models.RectAnchoredText {
					return owner.RectAnchoredTexts
				})
		}

	case *models.RectLinkLink:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Start", instanceWithInferedType.Start, formGroup, probe)
		AssociationFieldToForm("End", instanceWithInferedType.End, formGroup, probe)
		BasicFieldtoForm("TargetAnchorPosition", instanceWithInferedType.TargetAnchorPosition, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Layer, *models.RectLinkLink](
				"Layer",
				"RectLinkLinks",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Layer) []*models.RectLinkLink {
					return owner.RectLinkLinks
				})
		}

	case *models.SVG:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Layers", instanceWithInferedType, &instanceWithInferedType.Layers, formGroup, probe)
		EnumTypeStringToForm("DrawingState", instanceWithInferedType.DrawingState, instanceWithInferedType, probe.formStage, formGroup)
		AssociationFieldToForm("StartRect", instanceWithInferedType.StartRect, formGroup, probe)
		AssociationFieldToForm("EndRect", instanceWithInferedType.EndRect, formGroup, probe)
		BasicFieldtoForm("IsEditable", instanceWithInferedType.IsEditable, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsSVGFrontEndFileGenerated", instanceWithInferedType.IsSVGFrontEndFileGenerated, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsSVGBackEndFileGenerated", instanceWithInferedType.IsSVGBackEndFileGenerated, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DefaultDirectoryForGeneratedImages", instanceWithInferedType.DefaultDirectoryForGeneratedImages, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsControlBannerHidden", instanceWithInferedType.IsControlBannerHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OverrideWidth", instanceWithInferedType.OverrideWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OverriddenWidth", instanceWithInferedType.OverriddenWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OverrideHeight", instanceWithInferedType.OverrideHeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OverriddenHeight", instanceWithInferedType.OverriddenHeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.SvgText:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 300)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		BasicFieldtoForm("FontWeight", instanceWithInferedType.FontWeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FontSize", instanceWithInferedType.FontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FontStyle", instanceWithInferedType.FontStyle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LetterSpacing", instanceWithInferedType.LetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FontFamily", instanceWithInferedType.FontFamily, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("WhiteSpace", instanceWithInferedType.WhiteSpace, instanceWithInferedType, probe.formStage, formGroup)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Layer, *models.Text](
				"Layer",
				"Texts",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Layer) []*models.Text {
					return owner.Texts
				})
		}

	default:
		_ = instanceWithInferedType
	}
}
