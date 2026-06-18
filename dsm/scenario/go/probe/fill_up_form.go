// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/scenario/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.ActorState:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 500, true, 500)
		BasicFieldtoForm("IsWithProbaility", instanceWithInferedType.IsWithProbaility, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Probability", instanceWithInferedType.Probability, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.ActorState](
				"Diagram",
				"ActorStatesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.ActorState {
					return owner.ActorStatesWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Scenario, *models.ActorState](
				"Scenario",
				"ActorStates",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Scenario) []*models.ActorState {
					return owner.ActorStates
				})
		}

	case *models.ActorStateShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ActorState", instanceWithInferedType.ActorState, formGroup, probe)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.ActorStateShape](
				"Diagram",
				"ActorStateShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.ActorStateShape {
					return owner.ActorStateShapes
				})
		}

	case *models.ActorStateTransition:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("StartState", instanceWithInferedType.StartState, formGroup, probe)
		AssociationFieldToForm("EndState", instanceWithInferedType.EndState, formGroup, probe)
		AssociationSliceToForm("Justifications", instanceWithInferedType, &instanceWithInferedType.Justifications, formGroup, probe)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.ActorStateTransition](
				"Diagram",
				"ActorStateTransitionsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.ActorStateTransition {
					return owner.ActorStateTransitionsWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Scenario, *models.ActorStateTransition](
				"Scenario",
				"ActorStateTransitions",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Scenario) []*models.ActorStateTransition {
					return owner.ActorStateTransitions
				})
		}

	case *models.ActorStateTransitionShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ActorStateTransition", instanceWithInferedType.ActorStateTransition, formGroup, probe)
		AssociationFieldToForm("Start", instanceWithInferedType.Start, formGroup, probe)
		AssociationFieldToForm("End", instanceWithInferedType.End, formGroup, probe)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ControlPointShapes", instanceWithInferedType, &instanceWithInferedType.ControlPointShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.ActorStateTransitionShape](
				"Diagram",
				"ActorStateTransitionShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.ActorStateTransitionShape {
					return owner.ActorStateTransitionShapes
				})
		}

	case *models.Analysis:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 500, true, 500)
		AssociationSliceToForm("Scenarios", instanceWithInferedType, &instanceWithInferedType.Scenarios, formGroup, probe)
		BasicFieldtoForm("IsScenariosNodeExpanded", instanceWithInferedType.IsScenariosNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("GroupUse", instanceWithInferedType, &instanceWithInferedType.GroupUse, formGroup, probe)
		BasicFieldtoForm("IsGroupUseNodeExpanded", instanceWithInferedType.IsGroupUseNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("GeoObjectUse", instanceWithInferedType, &instanceWithInferedType.GeoObjectUse, formGroup, probe)
		BasicFieldtoForm("IsGeoObjectUseNodeExpanded", instanceWithInferedType.IsGeoObjectUseNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("MapUse", instanceWithInferedType, &instanceWithInferedType.MapUse, formGroup, probe)
		BasicFieldtoForm("IsMapUseNodeExpanded", instanceWithInferedType.IsMapUseNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Library, *models.Analysis](
				"Library",
				"Analyses",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Analysis {
					return owner.Analyses
				})
		}

	case *models.ControlPointShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X_Relative", instanceWithInferedType.X_Relative, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y_Relative", instanceWithInferedType.Y_Relative, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsStartShapeTheClosestShape", instanceWithInferedType.IsStartShapeTheClosestShape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.ActorStateTransitionShape, *models.ControlPointShape](
				"ActorStateTransitionShape",
				"ControlPointShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.ActorStateTransitionShape) []*models.ControlPointShape {
					return owner.ControlPointShapes
				})
		}

	case *models.Diagram:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsShowPrefix", instanceWithInferedType.IsShowPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		AssociationSliceToForm("EvolutionDirectionShapes", instanceWithInferedType, &instanceWithInferedType.EvolutionDirectionShapes, formGroup, probe)
		AssociationSliceToForm("EvolutionDirectionsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.EvolutionDirectionsWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsEvolutionDirectionsNodeExpanded", instanceWithInferedType.IsEvolutionDirectionsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ActorStateShapes", instanceWithInferedType, &instanceWithInferedType.ActorStateShapes, formGroup, probe)
		AssociationSliceToForm("ActorStatesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ActorStatesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsActorStatesNodeExpanded", instanceWithInferedType.IsActorStatesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ParameterShapes", instanceWithInferedType, &instanceWithInferedType.ParameterShapes, formGroup, probe)
		AssociationSliceToForm("ParametersWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ParametersWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsParametersNodeExpanded", instanceWithInferedType.IsParametersNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ScenarioParameterShapes", instanceWithInferedType, &instanceWithInferedType.ScenarioParameterShapes, formGroup, probe)
		AssociationSliceToForm("ParametersAggregatesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ParametersAggregatesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsParametersAggregatesNodeExpanded", instanceWithInferedType.IsParametersAggregatesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ActorStateTransitionShapes", instanceWithInferedType, &instanceWithInferedType.ActorStateTransitionShapes, formGroup, probe)
		AssociationSliceToForm("ActorStateTransitionsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ActorStateTransitionsWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsActorStateTransitionsNodeExpanded", instanceWithInferedType.IsActorStateTransitionsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("AxisOrign_X", instanceWithInferedType.AxisOrign_X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("AxisOrign_Y", instanceWithInferedType.AxisOrign_Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("VerticalAxis_Top_Y", instanceWithInferedType.VerticalAxis_Top_Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("VerticalAxis_Bottom_Y", instanceWithInferedType.VerticalAxis_Bottom_Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("VerticalAxis_StrokeWidth", instanceWithInferedType.VerticalAxis_StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HorizontalAxis_Right_X", instanceWithInferedType.HorizontalAxis_Right_X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Start", instanceWithInferedType.Start, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("End", instanceWithInferedType.End, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NumberOfYearsBetweenTicks", instanceWithInferedType.NumberOfYearsBetweenTicks, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Scenario, *models.Diagram](
				"Scenario",
				"Diagrams",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Scenario) []*models.Diagram {
					return owner.Diagrams
				})
		}

	case *models.Document:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("GeoObjectUse", instanceWithInferedType, &instanceWithInferedType.GeoObjectUse, formGroup, probe)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.DocumentUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Document", instanceWithInferedType.Document, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Parameter, *models.DocumentUse](
				"Parameter",
				"DocumentUse",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Parameter) []*models.DocumentUse {
					return owner.DocumentUse
				})
		}

	case *models.EvolutionDirection:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.EvolutionDirection](
				"Diagram",
				"EvolutionDirectionsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.EvolutionDirection {
					return owner.EvolutionDirectionsWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Scenario, *models.EvolutionDirection](
				"Scenario",
				"EvolutionDirections",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Scenario) []*models.EvolutionDirection {
					return owner.EvolutionDirections
				})
		}

	case *models.EvolutionDirectionShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("EvolutionDirection", instanceWithInferedType.EvolutionDirection, formGroup, probe)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.EvolutionDirectionShape](
				"Diagram",
				"EvolutionDirectionShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.EvolutionDirectionShape {
					return owner.EvolutionDirectionShapes
				})
		}

	case *models.Foo:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.GeoObject:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.GeoObjectUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("GeoObject", instanceWithInferedType.GeoObject, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Analysis, *models.GeoObjectUse](
				"Analysis",
				"GeoObjectUse",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Analysis) []*models.GeoObjectUse {
					return owner.GeoObjectUse
				})
		}
		{
			AssociationReverseSliceToForm[*models.Document, *models.GeoObjectUse](
				"Document",
				"GeoObjectUse",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Document) []*models.GeoObjectUse {
					return owner.GeoObjectUse
				})
		}
		{
			AssociationReverseSliceToForm[*models.Parameter, *models.GeoObjectUse](
				"Parameter",
				"GeoObjectUse",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Parameter) []*models.GeoObjectUse {
					return owner.GeoObjectUse
				})
		}

	case *models.Group:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("UserUse", instanceWithInferedType, &instanceWithInferedType.UserUse, formGroup, probe)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.GroupUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Group", instanceWithInferedType.Group, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Analysis, *models.GroupUse](
				"Analysis",
				"GroupUse",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Analysis) []*models.GroupUse {
					return owner.GroupUse
				})
		}
		{
			AssociationReverseSliceToForm[*models.Parameter, *models.GroupUse](
				"Parameter",
				"GroupUse",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Parameter) []*models.GroupUse {
					return owner.GroupUse
				})
		}
		{
			AssociationReverseSliceToForm[*models.Repository, *models.GroupUse](
				"Repository",
				"GroupUse",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Repository) []*models.GroupUse {
					return owner.GroupUse
				})
		}

	case *models.Library:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("IsRootLibrary", instanceWithInferedType.IsRootLibrary, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Analyses", instanceWithInferedType, &instanceWithInferedType.Analyses, formGroup, probe)
		BasicFieldtoForm("IsAnalysesNodeExpanded", instanceWithInferedType.IsAnalysesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SubLibraries", instanceWithInferedType, &instanceWithInferedType.SubLibraries, formGroup, probe)
		BasicFieldtoForm("IsSubLibrariesNodeExpanded", instanceWithInferedType.IsSubLibrariesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SubLibrariesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.SubLibrariesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("NbPixPerCharacter", instanceWithInferedType.NbPixPerCharacter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LogoSVGFile", instanceWithInferedType.LogoSVGFile, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, true, 300)
		BasicFieldtoForm("IsExpandedTmp", instanceWithInferedType.IsExpandedTmp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Library, *models.Library](
				"Library",
				"SubLibraries",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Library {
					return owner.SubLibraries
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Library](
				"Library",
				"SubLibrariesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Library {
					return owner.SubLibrariesWhoseNodeIsExpanded
				})
		}

	case *models.MapObject:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.MapObjectUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Map", instanceWithInferedType.Map, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Analysis, *models.MapObjectUse](
				"Analysis",
				"MapUse",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Analysis) []*models.MapObjectUse {
					return owner.MapUse
				})
		}

	case *models.Parameter:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 250, true, 500)
		BasicFieldtoForm("IsResponse", instanceWithInferedType.IsResponse, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Start", instanceWithInferedType.Start, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("End", instanceWithInferedType.End, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Force", instanceWithInferedType.Force, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("GroupUse", instanceWithInferedType, &instanceWithInferedType.GroupUse, formGroup, probe)
		AssociationSliceToForm("DocumentUse", instanceWithInferedType, &instanceWithInferedType.DocumentUse, formGroup, probe)
		AssociationSliceToForm("GeoObjectUse", instanceWithInferedType, &instanceWithInferedType.GeoObjectUse, formGroup, probe)
		BasicFieldtoForm("Tag", instanceWithInferedType.Tag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.ActorStateTransition, *models.Parameter](
				"ActorStateTransition",
				"Justifications",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.ActorStateTransition) []*models.Parameter {
					return owner.Justifications
				})
		}
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Parameter](
				"Diagram",
				"ParametersWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Parameter {
					return owner.ParametersWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.ParametersAggregate, *models.Parameter](
				"ParametersAggregate",
				"Parameters",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.ParametersAggregate) []*models.Parameter {
					return owner.Parameters
				})
		}
		{
			AssociationReverseSliceToForm[*models.Scenario, *models.Parameter](
				"Scenario",
				"Parameters",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Scenario) []*models.Parameter {
					return owner.Parameters
				})
		}

	case *models.ParameterCategory:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ParameterUse", instanceWithInferedType, &instanceWithInferedType.ParameterUse, formGroup, probe)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.ParameterCategoryUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ParameterCategory", instanceWithInferedType.ParameterCategory, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.ParameterShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Parameter", instanceWithInferedType.Parameter, formGroup, probe)
		EnumTypeStringToForm("Direction", instanceWithInferedType.Direction, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("ShapeIsComputedFromModel", instanceWithInferedType.ShapeIsComputedFromModel, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.ParameterShape](
				"Diagram",
				"ParameterShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.ParameterShape {
					return owner.ParameterShapes
				})
		}
		{
			AssociationReverseSliceToForm[*models.ParameterCategory, *models.ParameterShape](
				"ParameterCategory",
				"ParameterUse",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.ParameterCategory) []*models.ParameterShape {
					return owner.ParameterUse
				})
		}
		{
			AssociationReverseSliceToForm[*models.Repository, *models.ParameterShape](
				"Repository",
				"ParameterUse",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Repository) []*models.ParameterShape {
					return owner.ParameterUse
				})
		}

	case *models.ParametersAggregate:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Tag", instanceWithInferedType.Tag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 250, true, 500)
		AssociationSliceToForm("Parameters", instanceWithInferedType, &instanceWithInferedType.Parameters, formGroup, probe)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.ParametersAggregate](
				"Diagram",
				"ParametersAggregatesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.ParametersAggregate {
					return owner.ParametersAggregatesWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Scenario, *models.ParametersAggregate](
				"Scenario",
				"ParametersAggretates",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Scenario) []*models.ParametersAggregate {
					return owner.ParametersAggretates
				})
		}

	case *models.ParametersAggregateShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ScenarioParameter", instanceWithInferedType.ScenarioParameter, formGroup, probe)
		EnumTypeStringToForm("Direction", instanceWithInferedType.Direction, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.ParametersAggregateShape](
				"Diagram",
				"ScenarioParameterShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.ParametersAggregateShape {
					return owner.ScenarioParameterShapes
				})
		}

	case *models.Position:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Date", instanceWithInferedType.Date, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Ordinate", instanceWithInferedType.Ordinate, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Repository:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ParameterUse", instanceWithInferedType, &instanceWithInferedType.ParameterUse, formGroup, probe)
		AssociationSliceToForm("GroupUse", instanceWithInferedType, &instanceWithInferedType.GroupUse, formGroup, probe)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Scenario:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		AssociationSliceToForm("Diagrams", instanceWithInferedType, &instanceWithInferedType.Diagrams, formGroup, probe)
		BasicFieldtoForm("IsDiagramsNodeExpanded", instanceWithInferedType.IsDiagramsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ActorStates", instanceWithInferedType, &instanceWithInferedType.ActorStates, formGroup, probe)
		BasicFieldtoForm("IsActorStatesNodeExpanded", instanceWithInferedType.IsActorStatesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ActorStateTransitions", instanceWithInferedType, &instanceWithInferedType.ActorStateTransitions, formGroup, probe)
		BasicFieldtoForm("IsActorStateTransitionsNodeExpanded", instanceWithInferedType.IsActorStateTransitionsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("EvolutionDirections", instanceWithInferedType, &instanceWithInferedType.EvolutionDirections, formGroup, probe)
		BasicFieldtoForm("IsEvolutionDirectionsNodeExpanded", instanceWithInferedType.IsEvolutionDirectionsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Parameters", instanceWithInferedType, &instanceWithInferedType.Parameters, formGroup, probe)
		BasicFieldtoForm("IsParametersNodeExpanded", instanceWithInferedType.IsParametersNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ParametersAggretates", instanceWithInferedType, &instanceWithInferedType.ParametersAggretates, formGroup, probe)
		BasicFieldtoForm("IsParametersAggretatesNodeExpanded", instanceWithInferedType.IsParametersAggretatesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Analysis, *models.Scenario](
				"Analysis",
				"Scenarios",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Analysis) []*models.Scenario {
					return owner.Scenarios
				})
		}

	case *models.User:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.UserUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("User", instanceWithInferedType.User, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Group, *models.UserUse](
				"Group",
				"UserUse",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Group) []*models.UserUse {
					return owner.UserUse
				})
		}

	case *models.Workspace:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("SelectedDiagram", instanceWithInferedType.SelectedDiagram, formGroup, probe)
		AssociationFieldToForm("Default_EvolutionDirectionShape", instanceWithInferedType.Default_EvolutionDirectionShape, formGroup, probe)
		AssociationFieldToForm("Default_ParameterShape", instanceWithInferedType.Default_ParameterShape, formGroup, probe)
		AssociationFieldToForm("Default_ScenarioParameterShape", instanceWithInferedType.Default_ScenarioParameterShape, formGroup, probe)
		AssociationFieldToForm("Default_ActorStateShape", instanceWithInferedType.Default_ActorStateShape, formGroup, probe)
		AssociationFieldToForm("Default_ActorStateTransitionShape", instanceWithInferedType.Default_ActorStateTransitionShape, formGroup, probe)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	default:
		_ = instanceWithInferedType
	}
}
