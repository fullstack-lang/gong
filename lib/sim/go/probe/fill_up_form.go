// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/sim/go/models"
	"github.com/fullstack-lang/gong/lib/sim/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Command:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Command", instanceWithInferedType.Command, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CommandDate", instanceWithInferedType.CommandDate, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Engine", instanceWithInferedType.Engine, formGroup, probe)

	case *models.DummyAgent:
		// insertion point
		BasicFieldtoForm("TechName", instanceWithInferedType.TechName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Engine:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndTime", instanceWithInferedType.EndTime, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CurrentTime", instanceWithInferedType.CurrentTime, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DisplayFormat", instanceWithInferedType.DisplayFormat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SecondsSinceStart", instanceWithInferedType.SecondsSinceStart, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Fired", instanceWithInferedType.Fired, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("ControlMode", instanceWithInferedType.ControlMode, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("State", instanceWithInferedType.State, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Speed", instanceWithInferedType.Speed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Event:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Duration", instanceWithInferedType.Duration, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Status:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("CurrentCommand", instanceWithInferedType.CurrentCommand, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CompletionDate", instanceWithInferedType.CompletionDate, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("CurrentSpeedCommand", instanceWithInferedType.CurrentSpeedCommand, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("SpeedCommandCompletionDate", instanceWithInferedType.SpeedCommandCompletionDate, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.UpdateState:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Duration", instanceWithInferedType.Duration, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Period", instanceWithInferedType.Period, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	default:
		_ = instanceWithInferedType
	}
}
