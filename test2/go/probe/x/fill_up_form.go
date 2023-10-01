// generated code - do not edit
package x

import (
	"github.com/fullstack-lang/gong/test2/go/models/x"
	orm_x "github.com/fullstack-lang/gong/test2/go/orm/x"

	form "github.com/fullstack-lang/gongtable/go/models"
)

var __dummy_orm_fillup_form = orm_x.BackRepoStruct{}

func FillUpForm[T x.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *x.A:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)

	default:
		_ = instanceWithInferedType
	}
}
