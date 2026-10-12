// generated code (do not edit)
package models

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"
)

func getShapeFormStage(stager *Stager) *form.Stage {
	if stager == nil {
		return nil
	}
	if sfs, ok := any(stager).(interface{ GetShapeFormStage() *form.Stage }); ok {
		return sfs.GetShapeFormStage()
	}
	return nil
}

type rectShapeAppearanceCallback struct {
	stage     *Stage
	formStage *form.Stage
	formGroup *form.FormGroup
	shape     RectShapeInterface
}

func (cb *rectShapeAppearanceCallback) OnSave() {
	if cb.formStage == nil || cb.shape == nil || cb.stage == nil {
		return
	}
	cb.formStage.Checkout()
	for _, formDiv := range cb.formGroup.FormDivs {
		switch formDiv.Name {
		case "IsPresentationOverridden":
			val := cb.shape.GetIsPresentationOverridden()
			gongprobe.FormDivBasicFieldToField(&val, formDiv)
			cb.shape.SetIsPresentationOverridden(val)
		case "Color":
			val := cb.shape.GetColor()
			gongprobe.FormDivBasicFieldToField(&val, formDiv)
			cb.shape.SetColor(val)
		case "FillOpacity":
			val := cb.shape.GetFillOpacity()
			gongprobe.FormDivBasicFieldToField(&val, formDiv)
			cb.shape.SetFillOpacity(val)
		case "Stroke":
			val := cb.shape.GetStroke()
			gongprobe.FormDivBasicFieldToField(&val, formDiv)
			cb.shape.SetStroke(val)
		case "StrokeWidth":
			val := cb.shape.GetStrokeWidth()
			gongprobe.FormDivBasicFieldToField(&val, formDiv)
			cb.shape.SetStrokeWidth(val)
		case "StrokeOpacity":
			val := cb.shape.GetStrokeOpacity()
			gongprobe.FormDivBasicFieldToField(&val, formDiv)
			cb.shape.SetStrokeOpacity(val)
		case "StrokeDashArray":
			val := cb.shape.GetStrokeDashArray()
			gongprobe.FormDivBasicFieldToField(&val, formDiv)
			cb.shape.SetStrokeDashArray(val)
		case "RX":
			val := cb.shape.GetRX()
			gongprobe.FormDivBasicFieldToField(&val, formDiv)
			cb.shape.SetRX(val)
		case "FontSize":
			val := cb.shape.GetFontSize()
			gongprobe.FormDivBasicFieldToField(&val, formDiv)
			cb.shape.SetFontSize(val)
		case "FontWeight":
			val := cb.shape.GetFontWeight()
			gongprobe.FormDivBasicFieldToField(&val, formDiv)
			cb.shape.SetFontWeight(val)
		case "FontColor":
			val := cb.shape.GetFontColor()
			gongprobe.FormDivBasicFieldToField(&val, formDiv)
			cb.shape.SetFontColor(val)
		}
	}
	cb.stage.Commit()
}

type linkShapeAppearanceCallback struct {
	stage     *Stage
	formStage *form.Stage
	formGroup *form.FormGroup
	shape     LinkShapeInterface
}

func (cb *linkShapeAppearanceCallback) OnSave() {
	if cb.formStage == nil || cb.shape == nil || cb.stage == nil {
		return
	}
	cb.formStage.Checkout()
	for _, formDiv := range cb.formGroup.FormDivs {
		switch formDiv.Name {
		case "IsPresentationOverridden":
			val := cb.shape.GetIsPresentationOverridden()
			gongprobe.FormDivBasicFieldToField(&val, formDiv)
			cb.shape.SetIsPresentationOverridden(val)
		case "Stroke":
			val := cb.shape.GetStroke()
			gongprobe.FormDivBasicFieldToField(&val, formDiv)
			cb.shape.SetStroke(val)
		case "StrokeWidth":
			val := cb.shape.GetStrokeWidth()
			gongprobe.FormDivBasicFieldToField(&val, formDiv)
			cb.shape.SetStrokeWidth(val)
		case "StrokeOpacity":
			val := cb.shape.GetStrokeOpacity()
			gongprobe.FormDivBasicFieldToField(&val, formDiv)
			cb.shape.SetStrokeOpacity(val)
		case "StrokeDashArray":
			val := cb.shape.GetStrokeDashArray()
			gongprobe.FormDivBasicFieldToField(&val, formDiv)
			cb.shape.SetStrokeDashArray(val)
		case "FontSize":
			val := cb.shape.GetFontSize()
			gongprobe.FormDivBasicFieldToField(&val, formDiv)
			cb.shape.SetFontSize(val)
		case "FontColor":
			val := cb.shape.GetFontColor()
			gongprobe.FormDivBasicFieldToField(&val, formDiv)
			cb.shape.SetFontColor(val)
		}
	}
	cb.stage.Commit()
}

func fillUpShapeAppearanceForm(stager *Stager, shape GongstructIF) {
	formStage := getShapeFormStage(stager)
	if formStage == nil {
		return
	}
	formStage.Reset()

	if shape == nil {
		formStage.Commit()
		return
	}

	if rectShape, ok := shape.(RectShapeInterface); ok {
		label := shape.GetName()
		typeLabel := "Box Appearance"
		if gs, ok := shape.(GongstructIF); ok {
			typeLabel = gs.GongGetGongstructName()
		}
		formGroup := (&form.FormGroup{
			Name:      "Appearance Override",
			Label:     label,
			TypeLabel: typeLabel,
		}).Stage(formStage)
		formGroup.OnSave = &rectShapeAppearanceCallback{
			stage:     stager.stage,
			formStage: formStage,
			formGroup: formGroup,
			shape:     rectShape,
		}

		color := rectShape.GetColor()
		fillOpacity := rectShape.GetFillOpacity()
		stroke := rectShape.GetStroke()
		strokeWidth := rectShape.GetStrokeWidth()
		strokeOpacity := rectShape.GetStrokeOpacity()
		strokeDashArray := rectShape.GetStrokeDashArray()
		rx := rectShape.GetRX()
		fontSize := rectShape.GetFontSize()
		fontWeight := rectShape.GetFontWeight()
		fontColor := rectShape.GetFontColor()

		if color == "" {
			color = "whitesmoke"
		}
		if stroke == "" {
			stroke = "black"
		}
		if strokeWidth == 0 {
			strokeWidth = 1.0
		}
		if strokeOpacity == 0 {
			strokeOpacity = 1.0
		}
		if fillOpacity == 0 {
			fillOpacity = 1.0
		}
		if strokeDashArray == "" {
			strokeDashArray = "none"
		}
		if fontSize == "" {
			fontSize = "14px"
		}
		if fontWeight == "" {
			fontWeight = "normal"
		}
		if fontColor == "" {
			fontColor = "black"
		}

		isOverridden := rectShape.GetIsPresentationOverridden()
		if !isOverridden {
			isOverridden = true
		}
		gongprobe.BasicFieldtoForm("IsPresentationOverridden", isOverridden, formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("Color", color, formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("FillOpacity", fillOpacity, formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("Stroke", stroke, formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("StrokeWidth", strokeWidth, formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("StrokeOpacity", strokeOpacity, formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("StrokeDashArray", strokeDashArray, formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("RX", rx, formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("FontSize", fontSize, formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("FontWeight", fontWeight, formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("FontColor", fontColor, formStage, formGroup, false, false, 0, false, 0, false)

		formStage.Commit()
		return
	}

	if linkShape, ok := shape.(LinkShapeInterface); ok {
		label := shape.GetName()
		typeLabel := "Link Appearance"
		if gs, ok := shape.(GongstructIF); ok {
			typeLabel = gs.GongGetGongstructName()
		}
		formGroup := (&form.FormGroup{
			Name:      "Appearance Override",
			Label:     label,
			TypeLabel: typeLabel,
		}).Stage(formStage)
		formGroup.OnSave = &linkShapeAppearanceCallback{
			stage:     stager.stage,
			formStage: formStage,
			formGroup: formGroup,
			shape:     linkShape,
		}

		stroke := linkShape.GetStroke()
		strokeWidth := linkShape.GetStrokeWidth()
		strokeOpacity := linkShape.GetStrokeOpacity()
		strokeDashArray := linkShape.GetStrokeDashArray()
		fontSize := linkShape.GetFontSize()
		fontColor := linkShape.GetFontColor()

		if stroke == "" {
			stroke = "black"
		}
		if strokeWidth == 0 {
			strokeWidth = 1.5
		}
		if strokeOpacity == 0 {
			strokeOpacity = 1.0
		}
		if strokeDashArray == "" {
			strokeDashArray = "none"
		}
		if fontSize == "" {
			fontSize = "12px"
		}
		if fontColor == "" {
			fontColor = "black"
		}

		isOverridden := linkShape.GetIsPresentationOverridden()
		if !isOverridden {
			isOverridden = true
		}
		gongprobe.BasicFieldtoForm("IsPresentationOverridden", isOverridden, formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("Stroke", stroke, formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("StrokeWidth", strokeWidth, formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("StrokeOpacity", strokeOpacity, formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("StrokeDashArray", strokeDashArray, formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("FontSize", fontSize, formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("FontColor", fontColor, formStage, formGroup, false, false, 0, false, 0, false)

		formStage.Commit()
		return
	}

	formStage.Commit()
}

func clearShapeAppearanceForm(stager *Stager) {
	formStage := getShapeFormStage(stager)
	if formStage == nil {
		return
	}
	formStage.Reset()
	formStage.Commit()
}
