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
		if stage := sfs.GetShapeFormStage(); stage != nil {
			return stage
		}
	}
	if stager.probeForm != nil {
		if sfs, ok := any(stager.probeForm).(interface{ GetShapeFormStage() *form.Stage }); ok {
			return sfs.GetShapeFormStage()
		}
	}
	return nil
}

type rectShapeAppearanceCallback struct {
	stager    *Stager
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
	if cb.stager != nil {
		if gs, ok := cb.shape.(GongstructIF); ok {
			fillUpShapeAppearanceForm(cb.stager, gs)
		}
	}
}

type linkShapeAppearanceCallback struct {
	stager    *Stager
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
	if cb.stager != nil {
		if gs, ok := cb.shape.(GongstructIF); ok {
			fillUpShapeAppearanceForm(cb.stager, gs)
		}
	}
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
			stager:    stager,
			stage:     stager.stage,
			formStage: formStage,
			formGroup: formGroup,
			shape:     rectShape,
		}

		gongprobe.BasicFieldtoForm("IsPresentationOverridden", rectShape.GetIsPresentationOverridden(), formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("Color", rectShape.GetColor(), formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("FillOpacity", rectShape.GetFillOpacity(), formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("Stroke", rectShape.GetStroke(), formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("StrokeWidth", rectShape.GetStrokeWidth(), formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("StrokeOpacity", rectShape.GetStrokeOpacity(), formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("StrokeDashArray", rectShape.GetStrokeDashArray(), formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("RX", rectShape.GetRX(), formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("FontSize", rectShape.GetFontSize(), formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("FontWeight", rectShape.GetFontWeight(), formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("FontColor", rectShape.GetFontColor(), formStage, formGroup, false, false, 0, false, 0, false)

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
			stager:    stager,
			stage:     stager.stage,
			formStage: formStage,
			formGroup: formGroup,
			shape:     linkShape,
		}

		gongprobe.BasicFieldtoForm("IsPresentationOverridden", linkShape.GetIsPresentationOverridden(), formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("Stroke", linkShape.GetStroke(), formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("StrokeWidth", linkShape.GetStrokeWidth(), formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("StrokeOpacity", linkShape.GetStrokeOpacity(), formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("StrokeDashArray", linkShape.GetStrokeDashArray(), formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("FontSize", linkShape.GetFontSize(), formStage, formGroup, false, false, 0, false, 0, false)
		gongprobe.BasicFieldtoForm("FontColor", linkShape.GetFontColor(), formStage, formGroup, false, false, 0, false, 0, false)

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
