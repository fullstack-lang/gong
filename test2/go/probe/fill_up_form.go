// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test2/go/models"
)

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Dummy:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)

	default:
		_ = instanceWithInferedType
	}
}
