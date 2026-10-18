// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
)

type functionalStageSetFormCallback struct {
	onSave func()
}

func (cb *functionalStageSetFormCallback) OnSave() {
	if cb.onSave != nil {
		cb.onSave()
	}
}

func StageSetFillUpFormFromGongstruct(
	instance any,
	probe *StageSetProbe,
) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name: "Form",
	}).Stage(probe.formStage)

	switch inst := instance.(type) {
	case *models.Animate:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Animate_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Animate_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Circle:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Circle_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Circle_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Condition:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Condition_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Condition_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ControlPoint:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ControlPoint_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ControlPoint_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Ellipse:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Ellipse_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Ellipse_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.FileToDownload:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FileToDownload_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_FileToDownload_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Layer:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Layer_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Layer_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Line:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Line_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Line_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Link:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Link_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Link_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.LinkAnchoredPath:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LinkAnchoredPath_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_LinkAnchoredPath_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.LinkAnchoredText:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LinkAnchoredText_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_LinkAnchoredText_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Path:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Path_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Path_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Point:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Point_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Point_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Polygone:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Polygone_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Polygone_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Polyline:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Polyline_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Polyline_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Rect:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Rect_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Rect_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.RectAnchoredPath:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RectAnchoredPath_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_RectAnchoredPath_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.RectAnchoredPngImage:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RectAnchoredPngImage_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_RectAnchoredPngImage_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.RectAnchoredRect:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RectAnchoredRect_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_RectAnchoredRect_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.RectAnchoredText:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RectAnchoredText_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_RectAnchoredText_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.RectLinkLink:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RectLinkLink_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_RectLinkLink_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.SVG:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SVG_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_SVG_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.SvgText:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SvgText_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_SvgText_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Text:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Text_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Text_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	default:
		_ = inst
	}

	probe.formStage.Commit()
}

func saveStageSet_Animate_Stage(
	inst *models.Animate,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "AttributeName":
			FormDivBasicFieldToField(&inst.AttributeName, formDiv)
		case "Values":
			FormDivBasicFieldToField(&inst.Values, formDiv)
		case "From":
			FormDivBasicFieldToField(&inst.From, formDiv)
		case "To":
			FormDivBasicFieldToField(&inst.To, formDiv)
		case "Dur":
			FormDivBasicFieldToField(&inst.Dur, formDiv)
		case "RepeatCount":
			FormDivBasicFieldToField(&inst.RepeatCount, formDiv)
		}
	}
}

func saveStageSet_Circle_Stage(
	inst *models.Circle,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "CX":
			FormDivBasicFieldToField(&inst.CX, formDiv)
		case "CY":
			FormDivBasicFieldToField(&inst.CY, formDiv)
		case "Radius":
			FormDivBasicFieldToField(&inst.Radius, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&inst.FillOpacity, formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&inst.Stroke, formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&inst.StrokeOpacity, formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&inst.StrokeWidth, formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&inst.StrokeDashArray, formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&inst.StrokeDashArrayWhenSelected, formDiv)
		case "Transform":
			FormDivBasicFieldToField(&inst.Transform, formDiv)
		}
	}
}

func saveStageSet_Condition_Stage(
	inst *models.Condition,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		}
	}
}

func saveStageSet_ControlPoint_Stage(
	inst *models.ControlPoint,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "X_Relative":
			FormDivBasicFieldToField(&inst.X_Relative, formDiv)
		case "Y_Relative":
			FormDivBasicFieldToField(&inst.Y_Relative, formDiv)
		case "ClosestRect":
			StageSetFormDivSelectFieldToField(&inst.ClosestRect, probe.stageSet.Stage.GetInstancesSet[*models.Rect](), formDiv)
		}
	}
}

func saveStageSet_Ellipse_Stage(
	inst *models.Ellipse,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "CX":
			FormDivBasicFieldToField(&inst.CX, formDiv)
		case "CY":
			FormDivBasicFieldToField(&inst.CY, formDiv)
		case "RX":
			FormDivBasicFieldToField(&inst.RX, formDiv)
		case "RY":
			FormDivBasicFieldToField(&inst.RY, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&inst.FillOpacity, formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&inst.Stroke, formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&inst.StrokeOpacity, formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&inst.StrokeWidth, formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&inst.StrokeDashArray, formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&inst.StrokeDashArrayWhenSelected, formDiv)
		case "Transform":
			FormDivBasicFieldToField(&inst.Transform, formDiv)
		}
	}
}

func saveStageSet_FileToDownload_Stage(
	inst *models.FileToDownload,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Base64EncodedContent":
			FormDivBasicFieldToField(&inst.Base64EncodedContent, formDiv)
		}
	}
}

func saveStageSet_Layer_Stage(
	inst *models.Layer,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		}
	}
}

func saveStageSet_Line_Stage(
	inst *models.Line,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "X1":
			FormDivBasicFieldToField(&inst.X1, formDiv)
		case "Y1":
			FormDivBasicFieldToField(&inst.Y1, formDiv)
		case "X2":
			FormDivBasicFieldToField(&inst.X2, formDiv)
		case "Y2":
			FormDivBasicFieldToField(&inst.Y2, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&inst.FillOpacity, formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&inst.Stroke, formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&inst.StrokeOpacity, formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&inst.StrokeWidth, formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&inst.StrokeDashArray, formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&inst.StrokeDashArrayWhenSelected, formDiv)
		case "Transform":
			FormDivBasicFieldToField(&inst.Transform, formDiv)
		case "MouseClickX":
			FormDivBasicFieldToField(&inst.MouseClickX, formDiv)
		case "MouseClickY":
			FormDivBasicFieldToField(&inst.MouseClickY, formDiv)
		}
	}
}

func saveStageSet_Link_Stage(
	inst *models.Link,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "IsBezierCurve":
			FormDivBasicFieldToField(&inst.IsBezierCurve, formDiv)
		case "Start":
			StageSetFormDivSelectFieldToField(&inst.Start, probe.stageSet.Stage.GetInstancesSet[*models.Rect](), formDiv)
		case "StartAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.StartAnchorType, formDiv)
		case "End":
			StageSetFormDivSelectFieldToField(&inst.End, probe.stageSet.Stage.GetInstancesSet[*models.Rect](), formDiv)
		case "EndAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.EndAnchorType, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "CornerRadius":
			FormDivBasicFieldToField(&inst.CornerRadius, formDiv)
		case "HasEndArrow":
			FormDivBasicFieldToField(&inst.HasEndArrow, formDiv)
		case "EndArrowSize":
			FormDivBasicFieldToField(&inst.EndArrowSize, formDiv)
		case "EndArrowOffset":
			FormDivBasicFieldToField(&inst.EndArrowOffset, formDiv)
		case "HasStartArrow":
			FormDivBasicFieldToField(&inst.HasStartArrow, formDiv)
		case "StartArrowSize":
			FormDivBasicFieldToField(&inst.StartArrowSize, formDiv)
		case "StartArrowOffset":
			FormDivBasicFieldToField(&inst.StartArrowOffset, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&inst.FillOpacity, formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&inst.Stroke, formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&inst.StrokeOpacity, formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&inst.StrokeWidth, formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&inst.StrokeDashArray, formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&inst.StrokeDashArrayWhenSelected, formDiv)
		case "Transform":
			FormDivBasicFieldToField(&inst.Transform, formDiv)
		case "MouseX":
			FormDivBasicFieldToField(&inst.MouseX, formDiv)
		case "MouseY":
			FormDivBasicFieldToField(&inst.MouseY, formDiv)
		case "MouseEventKey":
			StageSetFormDivEnumStringFieldToField(&inst.MouseEventKey, formDiv)
		}
	}
}

func saveStageSet_LinkAnchoredPath_Stage(
	inst *models.LinkAnchoredPath,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Definition":
			FormDivBasicFieldToField(&inst.Definition, formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&inst.X_Offset, formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&inst.Y_Offset, formDiv)
		case "ScalePropotionnally":
			FormDivBasicFieldToField(&inst.ScalePropotionnally, formDiv)
		case "AppliedScaling":
			FormDivBasicFieldToField(&inst.AppliedScaling, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&inst.FillOpacity, formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&inst.Stroke, formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&inst.StrokeOpacity, formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&inst.StrokeWidth, formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&inst.StrokeDashArray, formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&inst.StrokeDashArrayWhenSelected, formDiv)
		case "Transform":
			FormDivBasicFieldToField(&inst.Transform, formDiv)
		}
	}
}

func saveStageSet_LinkAnchoredText_Stage(
	inst *models.LinkAnchoredText,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Content":
			FormDivBasicFieldToField(&inst.Content, formDiv)
		case "AutomaticLayout":
			FormDivBasicFieldToField(&inst.AutomaticLayout, formDiv)
		case "LinkAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.LinkAnchorType, formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&inst.X_Offset, formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&inst.Y_Offset, formDiv)
		case "FontWeight":
			FormDivBasicFieldToField(&inst.FontWeight, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontStyle":
			FormDivBasicFieldToField(&inst.FontStyle, formDiv)
		case "LetterSpacing":
			FormDivBasicFieldToField(&inst.LetterSpacing, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "WhiteSpace":
			StageSetFormDivEnumStringFieldToField(&inst.WhiteSpace, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&inst.FillOpacity, formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&inst.Stroke, formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&inst.StrokeOpacity, formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&inst.StrokeWidth, formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&inst.StrokeDashArray, formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&inst.StrokeDashArrayWhenSelected, formDiv)
		case "Transform":
			FormDivBasicFieldToField(&inst.Transform, formDiv)
		}
	}
}

func saveStageSet_Path_Stage(
	inst *models.Path,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Definition":
			FormDivBasicFieldToField(&inst.Definition, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&inst.FillOpacity, formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&inst.Stroke, formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&inst.StrokeOpacity, formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&inst.StrokeWidth, formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&inst.StrokeDashArray, formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&inst.StrokeDashArrayWhenSelected, formDiv)
		case "Transform":
			FormDivBasicFieldToField(&inst.Transform, formDiv)
		}
	}
}

func saveStageSet_Point_Stage(
	inst *models.Point,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "X":
			FormDivBasicFieldToField(&inst.X, formDiv)
		case "Y":
			FormDivBasicFieldToField(&inst.Y, formDiv)
		}
	}
}

func saveStageSet_Polygone_Stage(
	inst *models.Polygone,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Points":
			FormDivBasicFieldToField(&inst.Points, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&inst.FillOpacity, formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&inst.Stroke, formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&inst.StrokeOpacity, formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&inst.StrokeWidth, formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&inst.StrokeDashArray, formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&inst.StrokeDashArrayWhenSelected, formDiv)
		case "Transform":
			FormDivBasicFieldToField(&inst.Transform, formDiv)
		}
	}
}

func saveStageSet_Polyline_Stage(
	inst *models.Polyline,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Points":
			FormDivBasicFieldToField(&inst.Points, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&inst.FillOpacity, formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&inst.Stroke, formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&inst.StrokeOpacity, formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&inst.StrokeWidth, formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&inst.StrokeDashArray, formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&inst.StrokeDashArrayWhenSelected, formDiv)
		case "Transform":
			FormDivBasicFieldToField(&inst.Transform, formDiv)
		}
	}
}

func saveStageSet_Rect_Stage(
	inst *models.Rect,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "X":
			FormDivBasicFieldToField(&inst.X, formDiv)
		case "Y":
			FormDivBasicFieldToField(&inst.Y, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "RX":
			FormDivBasicFieldToField(&inst.RX, formDiv)
		case "EnclosingRect":
			StageSetFormDivSelectFieldToField(&inst.EnclosingRect, probe.stageSet.Stage.GetInstancesSet[*models.Rect](), formDiv)
		case "AnchoredTo":
			StageSetFormDivSelectFieldToField(&inst.AnchoredTo, probe.stageSet.Stage.GetInstancesSet[*models.Rect](), formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&inst.FillOpacity, formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&inst.Stroke, formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&inst.StrokeOpacity, formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&inst.StrokeWidth, formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&inst.StrokeDashArray, formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&inst.StrokeDashArrayWhenSelected, formDiv)
		case "Transform":
			FormDivBasicFieldToField(&inst.Transform, formDiv)
		case "IsSelectable":
			FormDivBasicFieldToField(&inst.IsSelectable, formDiv)
		case "IsSelected":
			FormDivBasicFieldToField(&inst.IsSelected, formDiv)
		case "CanHaveLeftHandle":
			FormDivBasicFieldToField(&inst.CanHaveLeftHandle, formDiv)
		case "HasLeftHandle":
			FormDivBasicFieldToField(&inst.HasLeftHandle, formDiv)
		case "CanHaveRightHandle":
			FormDivBasicFieldToField(&inst.CanHaveRightHandle, formDiv)
		case "HasRightHandle":
			FormDivBasicFieldToField(&inst.HasRightHandle, formDiv)
		case "CanHaveTopHandle":
			FormDivBasicFieldToField(&inst.CanHaveTopHandle, formDiv)
		case "HasTopHandle":
			FormDivBasicFieldToField(&inst.HasTopHandle, formDiv)
		case "IsScalingProportionally":
			FormDivBasicFieldToField(&inst.IsScalingProportionally, formDiv)
		case "CanHaveBottomHandle":
			FormDivBasicFieldToField(&inst.CanHaveBottomHandle, formDiv)
		case "HasBottomHandle":
			FormDivBasicFieldToField(&inst.HasBottomHandle, formDiv)
		case "CanMoveHorizontaly":
			FormDivBasicFieldToField(&inst.CanMoveHorizontaly, formDiv)
		case "CanMoveVerticaly":
			FormDivBasicFieldToField(&inst.CanMoveVerticaly, formDiv)
		case "ChangeColorWhenHovered":
			FormDivBasicFieldToField(&inst.ChangeColorWhenHovered, formDiv)
		case "ColorWhenHovered":
			FormDivBasicFieldToField(&inst.ColorWhenHovered, formDiv)
		case "OriginalColor":
			FormDivBasicFieldToField(&inst.OriginalColor, formDiv)
		case "FillOpacityWhenHovered":
			FormDivBasicFieldToField(&inst.FillOpacityWhenHovered, formDiv)
		case "OriginalFillOpacity":
			FormDivBasicFieldToField(&inst.OriginalFillOpacity, formDiv)
		case "HasToolTip":
			FormDivBasicFieldToField(&inst.HasToolTip, formDiv)
		case "ToolTipText":
			FormDivBasicFieldToField(&inst.ToolTipText, formDiv)
		case "ToolTipPosition":
			StageSetFormDivEnumStringFieldToField(&inst.ToolTipPosition, formDiv)
		case "MouseX":
			FormDivBasicFieldToField(&inst.MouseX, formDiv)
		case "MouseY":
			FormDivBasicFieldToField(&inst.MouseY, formDiv)
		case "MouseEventKey":
			StageSetFormDivEnumStringFieldToField(&inst.MouseEventKey, formDiv)
		case "URLPath":
			FormDivBasicFieldToField(&inst.URLPath, formDiv)
		case "URLTarget":
			StageSetFormDivEnumStringFieldToField(&inst.URLTarget, formDiv)
		}
	}
}

func saveStageSet_RectAnchoredPath_Stage(
	inst *models.RectAnchoredPath,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Definition":
			FormDivBasicFieldToField(&inst.Definition, formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&inst.X_Offset, formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&inst.Y_Offset, formDiv)
		case "RectAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.RectAnchorType, formDiv)
		case "ScalePropotionnally":
			FormDivBasicFieldToField(&inst.ScalePropotionnally, formDiv)
		case "AppliedScaling":
			FormDivBasicFieldToField(&inst.AppliedScaling, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&inst.FillOpacity, formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&inst.Stroke, formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&inst.StrokeOpacity, formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&inst.StrokeWidth, formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&inst.StrokeDashArray, formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&inst.StrokeDashArrayWhenSelected, formDiv)
		case "Transform":
			FormDivBasicFieldToField(&inst.Transform, formDiv)
		}
	}
}

func saveStageSet_RectAnchoredPngImage_Stage(
	inst *models.RectAnchoredPngImage,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "X":
			FormDivBasicFieldToField(&inst.X, formDiv)
		case "Y":
			FormDivBasicFieldToField(&inst.Y, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "RX":
			FormDivBasicFieldToField(&inst.RX, formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&inst.X_Offset, formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&inst.Y_Offset, formDiv)
		case "RectAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.RectAnchorType, formDiv)
		case "Base64Content":
			FormDivBasicFieldToField(&inst.Base64Content, formDiv)
		}
	}
}

func saveStageSet_RectAnchoredRect_Stage(
	inst *models.RectAnchoredRect,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "X":
			FormDivBasicFieldToField(&inst.X, formDiv)
		case "Y":
			FormDivBasicFieldToField(&inst.Y, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "RX":
			FormDivBasicFieldToField(&inst.RX, formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&inst.X_Offset, formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&inst.Y_Offset, formDiv)
		case "RectAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.RectAnchorType, formDiv)
		case "WidthFollowRect":
			FormDivBasicFieldToField(&inst.WidthFollowRect, formDiv)
		case "HeightFollowRect":
			FormDivBasicFieldToField(&inst.HeightFollowRect, formDiv)
		case "HasToolTip":
			FormDivBasicFieldToField(&inst.HasToolTip, formDiv)
		case "ToolTipText":
			FormDivBasicFieldToField(&inst.ToolTipText, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&inst.FillOpacity, formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&inst.Stroke, formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&inst.StrokeOpacity, formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&inst.StrokeWidth, formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&inst.StrokeDashArray, formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&inst.StrokeDashArrayWhenSelected, formDiv)
		case "Transform":
			FormDivBasicFieldToField(&inst.Transform, formDiv)
		}
	}
}

func saveStageSet_RectAnchoredText_Stage(
	inst *models.RectAnchoredText,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Content":
			FormDivBasicFieldToField(&inst.Content, formDiv)
		case "FontWeight":
			FormDivBasicFieldToField(&inst.FontWeight, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontStyle":
			FormDivBasicFieldToField(&inst.FontStyle, formDiv)
		case "LetterSpacing":
			FormDivBasicFieldToField(&inst.LetterSpacing, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "WhiteSpace":
			StageSetFormDivEnumStringFieldToField(&inst.WhiteSpace, formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&inst.X_Offset, formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&inst.Y_Offset, formDiv)
		case "RectAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.RectAnchorType, formDiv)
		case "TextAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.TextAnchorType, formDiv)
		case "DominantBaseline":
			StageSetFormDivEnumStringFieldToField(&inst.DominantBaseline, formDiv)
		case "WritingMode":
			StageSetFormDivEnumStringFieldToField(&inst.WritingMode, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&inst.FillOpacity, formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&inst.Stroke, formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&inst.StrokeOpacity, formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&inst.StrokeWidth, formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&inst.StrokeDashArray, formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&inst.StrokeDashArrayWhenSelected, formDiv)
		case "Transform":
			FormDivBasicFieldToField(&inst.Transform, formDiv)
		case "URLPath":
			FormDivBasicFieldToField(&inst.URLPath, formDiv)
		case "URLTarget":
			StageSetFormDivEnumStringFieldToField(&inst.URLTarget, formDiv)
		}
	}
}

func saveStageSet_RectLinkLink_Stage(
	inst *models.RectLinkLink,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Start":
			StageSetFormDivSelectFieldToField(&inst.Start, probe.stageSet.Stage.GetInstancesSet[*models.Rect](), formDiv)
		case "End":
			StageSetFormDivSelectFieldToField(&inst.End, probe.stageSet.Stage.GetInstancesSet[*models.Link](), formDiv)
		case "TargetAnchorPosition":
			FormDivBasicFieldToField(&inst.TargetAnchorPosition, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&inst.FillOpacity, formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&inst.Stroke, formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&inst.StrokeOpacity, formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&inst.StrokeWidth, formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&inst.StrokeDashArray, formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&inst.StrokeDashArrayWhenSelected, formDiv)
		case "Transform":
			FormDivBasicFieldToField(&inst.Transform, formDiv)
		}
	}
}

func saveStageSet_SVG_Stage(
	inst *models.SVG,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DrawingState":
			StageSetFormDivEnumStringFieldToField(&inst.DrawingState, formDiv)
		case "StartRect":
			StageSetFormDivSelectFieldToField(&inst.StartRect, probe.stageSet.Stage.GetInstancesSet[*models.Rect](), formDiv)
		case "EndRect":
			StageSetFormDivSelectFieldToField(&inst.EndRect, probe.stageSet.Stage.GetInstancesSet[*models.Rect](), formDiv)
		case "IsEditable":
			FormDivBasicFieldToField(&inst.IsEditable, formDiv)
		case "IsSVGFrontEndFileGenerated":
			FormDivBasicFieldToField(&inst.IsSVGFrontEndFileGenerated, formDiv)
		case "IsSVGBackEndFileGenerated":
			FormDivBasicFieldToField(&inst.IsSVGBackEndFileGenerated, formDiv)
		case "DefaultDirectoryForGeneratedImages":
			FormDivBasicFieldToField(&inst.DefaultDirectoryForGeneratedImages, formDiv)
		case "IsControlBannerHidden":
			FormDivBasicFieldToField(&inst.IsControlBannerHidden, formDiv)
		case "PanX":
			FormDivBasicFieldToField(&inst.PanX, formDiv)
		case "PanY":
			FormDivBasicFieldToField(&inst.PanY, formDiv)
		case "Zoom":
			FormDivBasicFieldToField(&inst.Zoom, formDiv)
		case "OverrideWidth":
			FormDivBasicFieldToField(&inst.OverrideWidth, formDiv)
		case "OverriddenWidth":
			FormDivBasicFieldToField(&inst.OverriddenWidth, formDiv)
		case "OverrideHeight":
			FormDivBasicFieldToField(&inst.OverrideHeight, formDiv)
		case "OverriddenHeight":
			FormDivBasicFieldToField(&inst.OverriddenHeight, formDiv)
		}
	}
}

func saveStageSet_SvgText_Stage(
	inst *models.SvgText,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		}
	}
}

func saveStageSet_Text_Stage(
	inst *models.Text,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "X":
			FormDivBasicFieldToField(&inst.X, formDiv)
		case "Y":
			FormDivBasicFieldToField(&inst.Y, formDiv)
		case "Content":
			FormDivBasicFieldToField(&inst.Content, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&inst.FillOpacity, formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&inst.Stroke, formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&inst.StrokeOpacity, formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&inst.StrokeWidth, formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&inst.StrokeDashArray, formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&inst.StrokeDashArrayWhenSelected, formDiv)
		case "Transform":
			FormDivBasicFieldToField(&inst.Transform, formDiv)
		case "FontWeight":
			FormDivBasicFieldToField(&inst.FontWeight, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontStyle":
			FormDivBasicFieldToField(&inst.FontStyle, formDiv)
		case "LetterSpacing":
			FormDivBasicFieldToField(&inst.LetterSpacing, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "WhiteSpace":
			StageSetFormDivEnumStringFieldToField(&inst.WhiteSpace, formDiv)
		}
	}
}


func StageSetNewInstance_Animate_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Animate",
	}).Stage(probe.formStage)
	inst := new(models.Animate)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Animate_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Animate_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Circle_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Circle",
	}).Stage(probe.formStage)
	inst := new(models.Circle)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Circle_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Circle_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Condition_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Condition",
	}).Stage(probe.formStage)
	inst := new(models.Condition)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Condition_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Condition_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ControlPoint_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ControlPoint",
	}).Stage(probe.formStage)
	inst := new(models.ControlPoint)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ControlPoint_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ControlPoint_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Ellipse_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Ellipse",
	}).Stage(probe.formStage)
	inst := new(models.Ellipse)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Ellipse_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Ellipse_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FileToDownload_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New FileToDownload",
	}).Stage(probe.formStage)
	inst := new(models.FileToDownload)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_FileToDownload_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_FileToDownload_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Layer_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Layer",
	}).Stage(probe.formStage)
	inst := new(models.Layer)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Layer_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Layer_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Line_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Line",
	}).Stage(probe.formStage)
	inst := new(models.Line)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Line_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Line_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Link_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Link",
	}).Stage(probe.formStage)
	inst := new(models.Link)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Link_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Link_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LinkAnchoredPath_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New LinkAnchoredPath",
	}).Stage(probe.formStage)
	inst := new(models.LinkAnchoredPath)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_LinkAnchoredPath_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_LinkAnchoredPath_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LinkAnchoredText_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New LinkAnchoredText",
	}).Stage(probe.formStage)
	inst := new(models.LinkAnchoredText)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_LinkAnchoredText_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_LinkAnchoredText_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Path_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Path",
	}).Stage(probe.formStage)
	inst := new(models.Path)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Path_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Path_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Point_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Point",
	}).Stage(probe.formStage)
	inst := new(models.Point)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Point_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Point_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Polygone_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Polygone",
	}).Stage(probe.formStage)
	inst := new(models.Polygone)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Polygone_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Polygone_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Polyline_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Polyline",
	}).Stage(probe.formStage)
	inst := new(models.Polyline)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Polyline_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Polyline_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Rect_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Rect",
	}).Stage(probe.formStage)
	inst := new(models.Rect)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Rect_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Rect_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RectAnchoredPath_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New RectAnchoredPath",
	}).Stage(probe.formStage)
	inst := new(models.RectAnchoredPath)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_RectAnchoredPath_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_RectAnchoredPath_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RectAnchoredPngImage_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New RectAnchoredPngImage",
	}).Stage(probe.formStage)
	inst := new(models.RectAnchoredPngImage)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_RectAnchoredPngImage_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_RectAnchoredPngImage_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RectAnchoredRect_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New RectAnchoredRect",
	}).Stage(probe.formStage)
	inst := new(models.RectAnchoredRect)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_RectAnchoredRect_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_RectAnchoredRect_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RectAnchoredText_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New RectAnchoredText",
	}).Stage(probe.formStage)
	inst := new(models.RectAnchoredText)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_RectAnchoredText_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_RectAnchoredText_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RectLinkLink_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New RectLinkLink",
	}).Stage(probe.formStage)
	inst := new(models.RectLinkLink)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_RectLinkLink_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_RectLinkLink_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SVG_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New SVG",
	}).Stage(probe.formStage)
	inst := new(models.SVG)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_SVG_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_SVG_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SvgText_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New SvgText",
	}).Stage(probe.formStage)
	inst := new(models.SvgText)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_SvgText_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_SvgText_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Text_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Text",
	}).Stage(probe.formStage)
	inst := new(models.Text)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Text_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Text_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

