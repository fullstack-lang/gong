// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/sim/go/models"
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
	case *models.Command:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Command", inst.Command, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CommandDate", inst.CommandDate, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Engine", inst.Engine, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Engine](), probe.formStage)
	case *models.DummyAgent:
		StageSetBasicFieldtoForm("TechName", inst.TechName, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Engine", inst.Engine, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Engine](), probe.formStage)
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
	case *models.Engine:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndTime", inst.EndTime, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CurrentTime", inst.CurrentTime, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DisplayFormat", inst.DisplayFormat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SecondsSinceStart", inst.SecondsSinceStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Fired", inst.Fired, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("ControlMode", inst.ControlMode, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("State", inst.State, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Speed", inst.Speed, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Commands {
				if src.Engine == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Command", "Engine", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DummyAgents {
				if src.Engine == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DummyAgent", "Engine", refNames, formGroup, probe.formStage)
		}
	case *models.Event:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Duration", inst.Duration, probe.formStage, formGroup)
	case *models.Status:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("CurrentCommand", inst.CurrentCommand, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CompletionDate", inst.CompletionDate, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("CurrentSpeedCommand", inst.CurrentSpeedCommand, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("SpeedCommandCompletionDate", inst.SpeedCommandCompletionDate, probe.formStage, formGroup)
	case *models.UpdateState:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Duration", inst.Duration, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Period", inst.Period, probe.formStage, formGroup)
	default:
		_ = inst
	}
}
