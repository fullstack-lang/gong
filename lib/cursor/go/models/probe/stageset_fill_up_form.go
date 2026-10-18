// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/cursor/go/models"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func StageSetFillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *StageSetProbe,
) {
	switch inst := instance.(type) {
	case *models.Cursor:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StartX", inst.StartX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndX", inst.EndX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y1", inst.Y1, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y2", inst.Y2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DurationSeconds", inst.DurationSeconds, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FillOpacity", inst.FillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Stroke", inst.Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeOpacity", inst.StrokeOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeWidth", inst.StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArray", inst.StrokeDashArray, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArrayWhenSelected", inst.StrokeDashArrayWhenSelected, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transform", inst.Transform, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsPlaying", inst.IsPlaying, probe.formStage, formGroup)
	default:
		_ = inst
	}
}
