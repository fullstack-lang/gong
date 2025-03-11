// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/sim/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Command":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Command Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CommandFormCallback(
			nil,
			probe,
			formGroup,
		)
		command := new(models.Command)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(command, formGroup, probe)
	case "DummyAgent":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DummyAgent Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DummyAgentFormCallback(
			nil,
			probe,
			formGroup,
		)
		dummyagent := new(models.DummyAgent)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(dummyagent, formGroup, probe)
	case "Engine":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Engine Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EngineFormCallback(
			nil,
			probe,
			formGroup,
		)
		engine := new(models.Engine)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(engine, formGroup, probe)
	case "Event":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Event Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EventFormCallback(
			nil,
			probe,
			formGroup,
		)
		event := new(models.Event)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(event, formGroup, probe)
	case "Status":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Status Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StatusFormCallback(
			nil,
			probe,
			formGroup,
		)
		status := new(models.Status)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(status, formGroup, probe)
	case "UpdateState":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "UpdateState Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UpdateStateFormCallback(
			nil,
			probe,
			formGroup,
		)
		updatestate := new(models.UpdateState)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(updatestate, formGroup, probe)
	}
	formStage.Commit()
}
