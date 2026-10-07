// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"github.com/fullstack-lang/gong/go/models"
)

func BasicFieldtoForm[TF models.GongtructBasicField](
	fieldName string, field TF, instance models.GongstructIF, formStage *form.Stage, formGroup *form.FormGroup,
	isTextArea bool, isBespokeWidth bool, bespokeWidth int, isBespokeHeight bool, bespokeHeight int, isTimeFormOnly bool,
) {
	gongprobe.BasicFieldtoForm(fieldName, field, formStage, formGroup,
		isTextArea, isBespokeWidth, bespokeWidth, isBespokeHeight, bespokeHeight, isTimeFormOnly)
}
