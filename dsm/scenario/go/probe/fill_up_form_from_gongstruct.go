// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/scenario/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.ActorState:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ActorState",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ActorStateFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ActorStateShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ActorStateShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ActorStateShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ActorStateTransition:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ActorStateTransition",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ActorStateTransitionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ActorStateTransitionShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ActorStateTransitionShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ActorStateTransitionShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Analysis:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Analysis",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AnalysisFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ControlPointShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ControlPointShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ControlPointShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Diagram:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Diagram",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Document:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Document",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DocumentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DocumentUse:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "DocumentUse",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DocumentUseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.EvolutionDirection:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "EvolutionDirection",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EvolutionDirectionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.EvolutionDirectionShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "EvolutionDirectionShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EvolutionDirectionShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Foo:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Foo",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FooFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GeoObject:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "GeoObject",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GeoObjectFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GeoObjectUse:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "GeoObjectUse",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GeoObjectUseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Group:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Group",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GroupFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GroupUse:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "GroupUse",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GroupUseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Library:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Library",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LibraryFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MapObject:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "MapObject",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MapObjectFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MapObjectUse:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "MapObjectUse",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MapObjectUseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Parameter:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Parameter",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParameterFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ParameterCategory:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ParameterCategory",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParameterCategoryFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ParameterCategoryUse:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ParameterCategoryUse",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParameterCategoryUseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ParameterShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ParameterShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParameterShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ParametersAggregate:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ParametersAggregate",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParametersAggregateFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ParametersAggregateShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ParametersAggregateShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParametersAggregateShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Position:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Position",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PositionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Repository:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Repository",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RepositoryFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Scenario:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Scenario",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ScenarioFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.User:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "User",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UserFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.UserUse:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "UserUse",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UserUseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Workspace:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Workspace",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WorkspaceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}

	if probe.GetCommitMode() {
		formStage.Commit()
	}
}
