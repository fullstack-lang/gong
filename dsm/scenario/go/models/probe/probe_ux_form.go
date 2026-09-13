// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/scenario/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *ActorStateFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ActorState", true)
			} else {
				FillUpFormFromGongstruct(onSave.actorstate, probe)
			}
		case *ActorStateShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ActorStateShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.actorstateshape, probe)
			}
		case *ActorStateTransitionFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ActorStateTransition", true)
			} else {
				FillUpFormFromGongstruct(onSave.actorstatetransition, probe)
			}
		case *ActorStateTransitionShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ActorStateTransitionShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.actorstatetransitionshape, probe)
			}
		case *AnalysisFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Analysis", true)
			} else {
				FillUpFormFromGongstruct(onSave.analysis, probe)
			}
		case *ControlPointShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ControlPointShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.controlpointshape, probe)
			}
		case *DiagramFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Diagram", true)
			} else {
				FillUpFormFromGongstruct(onSave.diagram, probe)
			}
		case *DocumentFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Document", true)
			} else {
				FillUpFormFromGongstruct(onSave.document, probe)
			}
		case *DocumentUseFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "DocumentUse", true)
			} else {
				FillUpFormFromGongstruct(onSave.documentuse, probe)
			}
		case *EvolutionDirectionFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "EvolutionDirection", true)
			} else {
				FillUpFormFromGongstruct(onSave.evolutiondirection, probe)
			}
		case *EvolutionDirectionShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "EvolutionDirectionShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.evolutiondirectionshape, probe)
			}
		case *FooFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Foo", true)
			} else {
				FillUpFormFromGongstruct(onSave.foo, probe)
			}
		case *GeoObjectFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "GeoObject", true)
			} else {
				FillUpFormFromGongstruct(onSave.geoobject, probe)
			}
		case *GeoObjectUseFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "GeoObjectUse", true)
			} else {
				FillUpFormFromGongstruct(onSave.geoobjectuse, probe)
			}
		case *GroupFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Group", true)
			} else {
				FillUpFormFromGongstruct(onSave.group, probe)
			}
		case *GroupUseFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "GroupUse", true)
			} else {
				FillUpFormFromGongstruct(onSave.groupuse, probe)
			}
		case *LibraryFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Library", true)
			} else {
				FillUpFormFromGongstruct(onSave.library, probe)
			}
		case *MapObjectFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "MapObject", true)
			} else {
				FillUpFormFromGongstruct(onSave.mapobject, probe)
			}
		case *MapObjectUseFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "MapObjectUse", true)
			} else {
				FillUpFormFromGongstruct(onSave.mapobjectuse, probe)
			}
		case *ParameterFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Parameter", true)
			} else {
				FillUpFormFromGongstruct(onSave.parameter, probe)
			}
		case *ParameterCategoryFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ParameterCategory", true)
			} else {
				FillUpFormFromGongstruct(onSave.parametercategory, probe)
			}
		case *ParameterCategoryUseFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ParameterCategoryUse", true)
			} else {
				FillUpFormFromGongstruct(onSave.parametercategoryuse, probe)
			}
		case *ParameterShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ParameterShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.parametershape, probe)
			}
		case *ParametersAggregateFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ParametersAggregate", true)
			} else {
				FillUpFormFromGongstruct(onSave.parametersaggregate, probe)
			}
		case *ParametersAggregateShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ParametersAggregateShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.parametersaggregateshape, probe)
			}
		case *PositionFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Position", true)
			} else {
				FillUpFormFromGongstruct(onSave.position, probe)
			}
		case *RepositoryFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Repository", true)
			} else {
				FillUpFormFromGongstruct(onSave.repository, probe)
			}
		case *ScenarioFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Scenario", true)
			} else {
				FillUpFormFromGongstruct(onSave.scenario, probe)
			}
		case *UserFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "User", true)
			} else {
				FillUpFormFromGongstruct(onSave.user, probe)
			}
		case *UserUseFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "UserUse", true)
			} else {
				FillUpFormFromGongstruct(onSave.useruse, probe)
			}
		case *WorkspaceFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Workspace", true)
			} else {
				FillUpFormFromGongstruct(onSave.workspace, probe)
			}
		}
	}
}

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
	case "ActorState":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ActorState Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ActorStateFormCallback(
			nil,
			probe,
			formGroup,
		)
		actorstate := new(models.ActorState)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(actorstate, formGroup, probe)
	case "ActorStateShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ActorStateShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ActorStateShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		actorstateshape := new(models.ActorStateShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(actorstateshape, formGroup, probe)
	case "ActorStateTransition":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ActorStateTransition Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ActorStateTransitionFormCallback(
			nil,
			probe,
			formGroup,
		)
		actorstatetransition := new(models.ActorStateTransition)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(actorstatetransition, formGroup, probe)
	case "ActorStateTransitionShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ActorStateTransitionShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ActorStateTransitionShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		actorstatetransitionshape := new(models.ActorStateTransitionShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(actorstatetransitionshape, formGroup, probe)
	case "Analysis":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Analysis Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AnalysisFormCallback(
			nil,
			probe,
			formGroup,
		)
		analysis := new(models.Analysis)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(analysis, formGroup, probe)
	case "ControlPointShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ControlPointShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ControlPointShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		controlpointshape := new(models.ControlPointShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(controlpointshape, formGroup, probe)
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
	case "Document":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Document Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DocumentFormCallback(
			nil,
			probe,
			formGroup,
		)
		document := new(models.Document)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(document, formGroup, probe)
	case "DocumentUse":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DocumentUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DocumentUseFormCallback(
			nil,
			probe,
			formGroup,
		)
		documentuse := new(models.DocumentUse)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(documentuse, formGroup, probe)
	case "EvolutionDirection":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "EvolutionDirection Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EvolutionDirectionFormCallback(
			nil,
			probe,
			formGroup,
		)
		evolutiondirection := new(models.EvolutionDirection)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(evolutiondirection, formGroup, probe)
	case "EvolutionDirectionShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "EvolutionDirectionShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EvolutionDirectionShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		evolutiondirectionshape := new(models.EvolutionDirectionShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(evolutiondirectionshape, formGroup, probe)
	case "Foo":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Foo Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FooFormCallback(
			nil,
			probe,
			formGroup,
		)
		foo := new(models.Foo)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(foo, formGroup, probe)
	case "GeoObject":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GeoObject Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GeoObjectFormCallback(
			nil,
			probe,
			formGroup,
		)
		geoobject := new(models.GeoObject)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(geoobject, formGroup, probe)
	case "GeoObjectUse":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GeoObjectUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GeoObjectUseFormCallback(
			nil,
			probe,
			formGroup,
		)
		geoobjectuse := new(models.GeoObjectUse)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(geoobjectuse, formGroup, probe)
	case "Group":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Group Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GroupFormCallback(
			nil,
			probe,
			formGroup,
		)
		group := new(models.Group)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(group, formGroup, probe)
	case "GroupUse":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GroupUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GroupUseFormCallback(
			nil,
			probe,
			formGroup,
		)
		groupuse := new(models.GroupUse)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(groupuse, formGroup, probe)
	case "Library":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Library Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LibraryFormCallback(
			nil,
			probe,
			formGroup,
		)
		library := new(models.Library)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(library, formGroup, probe)
	case "MapObject":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "MapObject Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MapObjectFormCallback(
			nil,
			probe,
			formGroup,
		)
		mapobject := new(models.MapObject)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(mapobject, formGroup, probe)
	case "MapObjectUse":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "MapObjectUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MapObjectUseFormCallback(
			nil,
			probe,
			formGroup,
		)
		mapobjectuse := new(models.MapObjectUse)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(mapobjectuse, formGroup, probe)
	case "Parameter":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Parameter Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParameterFormCallback(
			nil,
			probe,
			formGroup,
		)
		parameter := new(models.Parameter)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(parameter, formGroup, probe)
	case "ParameterCategory":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ParameterCategory Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParameterCategoryFormCallback(
			nil,
			probe,
			formGroup,
		)
		parametercategory := new(models.ParameterCategory)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(parametercategory, formGroup, probe)
	case "ParameterCategoryUse":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ParameterCategoryUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParameterCategoryUseFormCallback(
			nil,
			probe,
			formGroup,
		)
		parametercategoryuse := new(models.ParameterCategoryUse)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(parametercategoryuse, formGroup, probe)
	case "ParameterShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ParameterShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParameterShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		parametershape := new(models.ParameterShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(parametershape, formGroup, probe)
	case "ParametersAggregate":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ParametersAggregate Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParametersAggregateFormCallback(
			nil,
			probe,
			formGroup,
		)
		parametersaggregate := new(models.ParametersAggregate)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(parametersaggregate, formGroup, probe)
	case "ParametersAggregateShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ParametersAggregateShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParametersAggregateShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		parametersaggregateshape := new(models.ParametersAggregateShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(parametersaggregateshape, formGroup, probe)
	case "Position":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Position Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PositionFormCallback(
			nil,
			probe,
			formGroup,
		)
		position := new(models.Position)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(position, formGroup, probe)
	case "Repository":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Repository Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RepositoryFormCallback(
			nil,
			probe,
			formGroup,
		)
		repository := new(models.Repository)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(repository, formGroup, probe)
	case "Scenario":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Scenario Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ScenarioFormCallback(
			nil,
			probe,
			formGroup,
		)
		scenario := new(models.Scenario)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(scenario, formGroup, probe)
	case "User":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "User Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UserFormCallback(
			nil,
			probe,
			formGroup,
		)
		user := new(models.User)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(user, formGroup, probe)
	case "UserUse":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "UserUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UserUseFormCallback(
			nil,
			probe,
			formGroup,
		)
		useruse := new(models.UserUse)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(useruse, formGroup, probe)
	case "Workspace":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Workspace Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WorkspaceFormCallback(
			nil,
			probe,
			formGroup,
		)
		workspace := new(models.Workspace)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(workspace, formGroup, probe)
	}
	formStage.Commit()
}
