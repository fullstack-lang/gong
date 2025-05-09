// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

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
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Circle"
			rf.Fieldname = "Animations"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Circle),
					"Animations",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Circle, *models.Animate](
					nil,
					"Animations",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ellipse"
			rf.Fieldname = "Animates"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Ellipse),
					"Animates",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Ellipse, *models.Animate](
					nil,
					"Animates",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Line"
			rf.Fieldname = "Animates"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Line),
					"Animates",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Line, *models.Animate](
					nil,
					"Animates",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "LinkAnchoredText"
			rf.Fieldname = "Animates"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.LinkAnchoredText),
					"Animates",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.LinkAnchoredText, *models.Animate](
					nil,
					"Animates",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Path"
			rf.Fieldname = "Animates"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Path),
					"Animates",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Path, *models.Animate](
					nil,
					"Animates",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Polygone"
			rf.Fieldname = "Animates"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Polygone),
					"Animates",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Polygone, *models.Animate](
					nil,
					"Animates",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Polyline"
			rf.Fieldname = "Animates"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Polyline),
					"Animates",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Polyline, *models.Animate](
					nil,
					"Animates",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Rect"
			rf.Fieldname = "Animations"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Rect),
					"Animations",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Rect, *models.Animate](
					nil,
					"Animations",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RectAnchoredText"
			rf.Fieldname = "Animates"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.RectAnchoredText),
					"Animates",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.RectAnchoredText, *models.Animate](
					nil,
					"Animates",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Text"
			rf.Fieldname = "Animates"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Text),
					"Animates",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Text, *models.Animate](
					nil,
					"Animates",
					instanceWithInferedType,
					formGroup,
					probe)
			}
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
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Circles"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"Circles",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.Circle](
					nil,
					"Circles",
					instanceWithInferedType,
					formGroup,
					probe)
			}
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
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Ellipses"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"Ellipses",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.Ellipse](
					nil,
					"Ellipses",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Layer:
		// insertion point
		BasicFieldtoForm("Display", instanceWithInferedType.Display, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
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
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SVG"
			rf.Fieldname = "Layers"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SVG),
					"Layers",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SVG, *models.Layer](
					nil,
					"Layers",
					instanceWithInferedType,
					formGroup,
					probe)
			}
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
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Lines"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"Lines",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.Line](
					nil,
					"Lines",
					instanceWithInferedType,
					formGroup,
					probe)
			}
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
		BasicFieldtoForm("HasStartArrow", instanceWithInferedType.HasStartArrow, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StartArrowSize", instanceWithInferedType.StartArrowSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("TextAtArrowStart", instanceWithInferedType, &instanceWithInferedType.TextAtArrowStart, formGroup, probe)
		AssociationSliceToForm("TextAtArrowEnd", instanceWithInferedType, &instanceWithInferedType.TextAtArrowEnd, formGroup, probe)
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
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Links"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"Links",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.Link](
					nil,
					"Links",
					instanceWithInferedType,
					formGroup,
					probe)
			}
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
		BasicFieldtoForm("LetterSpacing", instanceWithInferedType.LetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
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
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Link"
			rf.Fieldname = "TextAtArrowStart"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Link),
					"TextAtArrowStart",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Link, *models.LinkAnchoredText](
					nil,
					"TextAtArrowStart",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Link"
			rf.Fieldname = "TextAtArrowEnd"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Link),
					"TextAtArrowEnd",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Link, *models.LinkAnchoredText](
					nil,
					"TextAtArrowEnd",
					instanceWithInferedType,
					formGroup,
					probe)
			}
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
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Paths"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"Paths",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.Path](
					nil,
					"Paths",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Point:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Link"
			rf.Fieldname = "ControlPoints"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Link),
					"ControlPoints",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Link, *models.Point](
					nil,
					"ControlPoints",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

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
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Polygones"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"Polygones",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.Polygone](
					nil,
					"Polygones",
					instanceWithInferedType,
					formGroup,
					probe)
			}
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
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Polylines"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"Polylines",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.Polyline](
					nil,
					"Polylines",
					instanceWithInferedType,
					formGroup,
					probe)
			}
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
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Rects"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"Rects",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.Rect](
					nil,
					"Rects",
					instanceWithInferedType,
					formGroup,
					probe)
			}
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
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Rect"
			rf.Fieldname = "RectAnchoredPaths"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Rect),
					"RectAnchoredPaths",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Rect, *models.RectAnchoredPath](
					nil,
					"RectAnchoredPaths",
					instanceWithInferedType,
					formGroup,
					probe)
			}
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
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Rect"
			rf.Fieldname = "RectAnchoredRects"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Rect),
					"RectAnchoredRects",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Rect, *models.RectAnchoredRect](
					nil,
					"RectAnchoredRects",
					instanceWithInferedType,
					formGroup,
					probe)
			}
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
		BasicFieldtoForm("X_Offset", instanceWithInferedType.X_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y_Offset", instanceWithInferedType.Y_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("RectAnchorType", instanceWithInferedType.RectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("TextAnchorType", instanceWithInferedType.TextAnchorType, instanceWithInferedType, probe.formStage, formGroup)
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
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Rect"
			rf.Fieldname = "RectAnchoredTexts"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Rect),
					"RectAnchoredTexts",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Rect, *models.RectAnchoredText](
					nil,
					"RectAnchoredTexts",
					instanceWithInferedType,
					formGroup,
					probe)
			}
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
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "RectLinkLinks"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"RectLinkLinks",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.RectLinkLink](
					nil,
					"RectLinkLinks",
					instanceWithInferedType,
					formGroup,
					probe)
			}
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
		BasicFieldtoForm("IsSVGFileGenerated", instanceWithInferedType.IsSVGFileGenerated, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.SvgText:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 300)

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
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Texts"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"Texts",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.Text](
					nil,
					"Texts",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	default:
		_ = instanceWithInferedType
	}
}
