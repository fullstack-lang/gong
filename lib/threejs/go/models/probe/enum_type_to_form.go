// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func EnumTypeStringToForm[T models.PointerToGongstruct, TF models.GongstructEnumStringField](
	fieldName string, field TF, instance T, formStage *form.Stage, formGroup *form.FormGroup,
) {
	gongprobe.EnumTypeStringToForm(fieldName, field, formStage, formGroup)
}

func EnumTypeIntToForm[T models.PointerToGongstruct, TF models.GongstructEnumIntField](
	fieldName string, field TF, instance T, formStage *form.Stage, formGroup *form.FormGroup,
) {
	gongprobe.EnumTypeIntToForm(fieldName, field, formStage, formGroup)
}
