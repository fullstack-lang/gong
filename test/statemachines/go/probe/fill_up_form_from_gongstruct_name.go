// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/test/statemachines/go/models"
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
	case "Architecture":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Architecture Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArchitectureFormCallback(
			nil,
			probe,
			formGroup,
		)
		architecture := new(models.Architecture)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(architecture, formGroup, probe)
	case "Diagram":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Diagram Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramFormCallback(
			nil,
			probe,
			formGroup,
		)
		diagram := new(models.Diagram)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(diagram, formGroup, probe)
	case "DoAction":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DoAction Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DoActionFormCallback(
			nil,
			probe,
			formGroup,
		)
		doaction := new(models.DoAction)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(doaction, formGroup, probe)
	case "Kill":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Kill Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KillFormCallback(
			nil,
			probe,
			formGroup,
		)
		kill := new(models.Kill)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(kill, formGroup, probe)
	case "Message":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Message Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MessageFormCallback(
			nil,
			probe,
			formGroup,
		)
		message := new(models.Message)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(message, formGroup, probe)
	case "MessageType":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "MessageType Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MessageTypeFormCallback(
			nil,
			probe,
			formGroup,
		)
		messagetype := new(models.MessageType)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(messagetype, formGroup, probe)
	case "Object":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Object Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ObjectFormCallback(
			nil,
			probe,
			formGroup,
		)
		object := new(models.Object)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(object, formGroup, probe)
	case "Role":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Role Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RoleFormCallback(
			nil,
			probe,
			formGroup,
		)
		role := new(models.Role)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(role, formGroup, probe)
	case "State":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "State Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StateFormCallback(
			nil,
			probe,
			formGroup,
		)
		state := new(models.State)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(state, formGroup, probe)
	case "StateMachine":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StateMachine Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StateMachineFormCallback(
			nil,
			probe,
			formGroup,
		)
		statemachine := new(models.StateMachine)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(statemachine, formGroup, probe)
	case "StateShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StateShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StateShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		stateshape := new(models.StateShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stateshape, formGroup, probe)
	case "Transition":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Transition Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TransitionFormCallback(
			nil,
			probe,
			formGroup,
		)
		transition := new(models.Transition)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(transition, formGroup, probe)
	case "Transition_Shape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Transition_Shape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Transition_ShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		transition_shape := new(models.Transition_Shape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(transition_shape, formGroup, probe)
	}
	formStage.Commit()
}
