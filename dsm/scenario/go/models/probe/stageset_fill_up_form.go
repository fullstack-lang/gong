// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/scenario/go/models"
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
	case *models.ActorState:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsWithProbaility", inst.IsWithProbaility, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Probability", inst.Probability, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ActorStateShapes {
				if src.ActorState == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ActorStateShape", "ActorState", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ActorStateTransitions {
				if src.StartState == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ActorStateTransition", "StartState", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ActorStateTransitions {
				if src.EndState == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ActorStateTransition", "EndState", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ActorStatesWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ActorStatesWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Scenarios {
				for _, target := range src.ActorStates {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Scenario", "ActorStates", refNames, formGroup, probe.formStage)
		}
	case *models.ActorStateShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ActorState", inst.ActorState, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ActorState](), probe.formStage)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ActorStateTransitionShapes {
				if src.Start == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ActorStateTransitionShape", "Start", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ActorStateTransitionShapes {
				if src.End == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ActorStateTransitionShape", "End", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ActorStateShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ActorStateShapes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Workspaces {
				if src.Default_ActorStateShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Workspace", "Default_ActorStateShape", refNames, formGroup, probe.formStage)
		}
	case *models.ActorStateTransition:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("StartState", inst.StartState, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ActorState](), probe.formStage)
		StageSetAssociationFieldToForm("EndState", inst.EndState, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ActorState](), probe.formStage)

		{
			// Slice of pointers: Justifications
			div := (&form.FormDiv{Name: "Justifications"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Justifications {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Justifications",
				Label: "Justifications",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ActorStateTransitionShapes {
				if src.ActorStateTransition == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ActorStateTransitionShape", "ActorStateTransition", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ActorStateTransitionsWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ActorStateTransitionsWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Scenarios {
				for _, target := range src.ActorStateTransitions {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Scenario", "ActorStateTransitions", refNames, formGroup, probe.formStage)
		}
	case *models.ActorStateTransitionShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ActorStateTransition", inst.ActorStateTransition, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ActorStateTransition](), probe.formStage)
		StageSetAssociationFieldToForm("Start", inst.Start, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ActorStateShape](), probe.formStage)
		StageSetAssociationFieldToForm("End", inst.End, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ActorStateShape](), probe.formStage)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			// Slice of pointers: ControlPointShapes
			div := (&form.FormDiv{Name: "ControlPointShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ControlPointShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ControlPointShapes",
				Label: "ControlPointShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ActorStateTransitionShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ActorStateTransitionShapes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Workspaces {
				if src.Default_ActorStateTransitionShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Workspace", "Default_ActorStateTransitionShape", refNames, formGroup, probe.formStage)
		}
	case *models.Analysis:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)

		{
			// Slice of pointers: Scenarios
			div := (&form.FormDiv{Name: "Scenarios"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Scenarios {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Scenarios",
				Label: "Scenarios",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsScenariosNodeExpanded", inst.IsScenariosNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: GroupUse
			div := (&form.FormDiv{Name: "GroupUse"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GroupUse {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GroupUse",
				Label: "GroupUse",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsGroupUseNodeExpanded", inst.IsGroupUseNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: GeoObjectUse
			div := (&form.FormDiv{Name: "GeoObjectUse"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GeoObjectUse {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GeoObjectUse",
				Label: "GeoObjectUse",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsGeoObjectUseNodeExpanded", inst.IsGeoObjectUseNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: MapUse
			div := (&form.FormDiv{Name: "MapUse"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.MapUse {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "MapUse",
				Label: "MapUse",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsMapUseNodeExpanded", inst.IsMapUseNodeExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.Analyses {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "Analyses", refNames, formGroup, probe.formStage)
		}
	case *models.ControlPointShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X_Relative", inst.X_Relative, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y_Relative", inst.Y_Relative, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsStartShapeTheClosestShape", inst.IsStartShapeTheClosestShape, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ActorStateTransitionShapes {
				for _, target := range src.ControlPointShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ActorStateTransitionShape", "ControlPointShapes", refNames, formGroup, probe.formStage)
		}
	case *models.Diagram:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsChecked", inst.IsChecked, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsShowPrefix", inst.IsShowPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)

		{
			// Slice of pointers: EvolutionDirectionShapes
			div := (&form.FormDiv{Name: "EvolutionDirectionShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.EvolutionDirectionShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "EvolutionDirectionShapes",
				Label: "EvolutionDirectionShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: EvolutionDirectionsWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "EvolutionDirectionsWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.EvolutionDirectionsWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "EvolutionDirectionsWhoseNodeIsExpanded",
				Label: "EvolutionDirectionsWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsEvolutionDirectionsNodeExpanded", inst.IsEvolutionDirectionsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: ActorStateShapes
			div := (&form.FormDiv{Name: "ActorStateShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ActorStateShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ActorStateShapes",
				Label: "ActorStateShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ActorStatesWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "ActorStatesWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ActorStatesWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ActorStatesWhoseNodeIsExpanded",
				Label: "ActorStatesWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsActorStatesNodeExpanded", inst.IsActorStatesNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: ParameterShapes
			div := (&form.FormDiv{Name: "ParameterShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ParameterShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ParameterShapes",
				Label: "ParameterShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ParametersWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "ParametersWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ParametersWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ParametersWhoseNodeIsExpanded",
				Label: "ParametersWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsParametersNodeExpanded", inst.IsParametersNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: ScenarioParameterShapes
			div := (&form.FormDiv{Name: "ScenarioParameterShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ScenarioParameterShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ScenarioParameterShapes",
				Label: "ScenarioParameterShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ParametersAggregatesWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "ParametersAggregatesWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ParametersAggregatesWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ParametersAggregatesWhoseNodeIsExpanded",
				Label: "ParametersAggregatesWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsParametersAggregatesNodeExpanded", inst.IsParametersAggregatesNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: ActorStateTransitionShapes
			div := (&form.FormDiv{Name: "ActorStateTransitionShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ActorStateTransitionShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ActorStateTransitionShapes",
				Label: "ActorStateTransitionShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ActorStateTransitionsWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "ActorStateTransitionsWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ActorStateTransitionsWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ActorStateTransitionsWhoseNodeIsExpanded",
				Label: "ActorStateTransitionsWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsActorStateTransitionsNodeExpanded", inst.IsActorStateTransitionsNodeExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("AxisOrign_X", inst.AxisOrign_X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("AxisOrign_Y", inst.AxisOrign_Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("VerticalAxis_Top_Y", inst.VerticalAxis_Top_Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("VerticalAxis_Bottom_Y", inst.VerticalAxis_Bottom_Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("VerticalAxis_StrokeWidth", inst.VerticalAxis_StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HorizontalAxis_Right_X", inst.HorizontalAxis_Right_X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Start", inst.Start, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("End", inst.End, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumberOfYearsBetweenTicks", inst.NumberOfYearsBetweenTicks, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsInDrawMode", inst.IsInDrawMode, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Scenarios {
				for _, target := range src.Diagrams {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Scenario", "Diagrams", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Workspaces {
				if src.SelectedDiagram == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Workspace", "SelectedDiagram", refNames, formGroup, probe.formStage)
		}
	case *models.Document:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: GeoObjectUse
			div := (&form.FormDiv{Name: "GeoObjectUse"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GeoObjectUse {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GeoObjectUse",
				Label: "GeoObjectUse",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DocumentUses {
				if src.Document == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DocumentUse", "Document", refNames, formGroup, probe.formStage)
		}
	case *models.DocumentUse:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Document", inst.Document, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Document](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Parameters {
				for _, target := range src.DocumentUse {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Parameter", "DocumentUse", refNames, formGroup, probe.formStage)
		}
	case *models.EvolutionDirection:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.EvolutionDirectionsWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "EvolutionDirectionsWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.EvolutionDirectionShapes {
				if src.EvolutionDirection == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.EvolutionDirectionShape", "EvolutionDirection", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Scenarios {
				for _, target := range src.EvolutionDirections {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Scenario", "EvolutionDirections", refNames, formGroup, probe.formStage)
		}
	case *models.EvolutionDirectionShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("EvolutionDirection", inst.EvolutionDirection, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.EvolutionDirection](), probe.formStage)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.EvolutionDirectionShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "EvolutionDirectionShapes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Workspaces {
				if src.Default_EvolutionDirectionShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Workspace", "Default_EvolutionDirectionShape", refNames, formGroup, probe.formStage)
		}
	case *models.Foo:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
	case *models.GeoObject:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.GeoObjectUses {
				if src.GeoObject == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.GeoObjectUse", "GeoObject", refNames, formGroup, probe.formStage)
		}
	case *models.GeoObjectUse:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("GeoObject", inst.GeoObject, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.GeoObject](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Analysiss {
				for _, target := range src.GeoObjectUse {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Analysis", "GeoObjectUse", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Documents {
				for _, target := range src.GeoObjectUse {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Document", "GeoObjectUse", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Parameters {
				for _, target := range src.GeoObjectUse {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Parameter", "GeoObjectUse", refNames, formGroup, probe.formStage)
		}
	case *models.Group:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: UserUse
			div := (&form.FormDiv{Name: "UserUse"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.UserUse {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "UserUse",
				Label: "UserUse",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.GroupUses {
				if src.Group == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.GroupUse", "Group", refNames, formGroup, probe.formStage)
		}
	case *models.GroupUse:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Group", inst.Group, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Group](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Analysiss {
				for _, target := range src.GroupUse {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Analysis", "GroupUse", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Parameters {
				for _, target := range src.GroupUse {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Parameter", "GroupUse", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Repositorys {
				for _, target := range src.GroupUse {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Repository", "GroupUse", refNames, formGroup, probe.formStage)
		}
	case *models.Library:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsRootLibrary", inst.IsRootLibrary, probe.formStage, formGroup)

		{
			// Slice of pointers: Analyses
			div := (&form.FormDiv{Name: "Analyses"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Analyses {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Analyses",
				Label: "Analyses",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsAnalysesNodeExpanded", inst.IsAnalysesNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: SubLibraries
			div := (&form.FormDiv{Name: "SubLibraries"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SubLibraries {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SubLibraries",
				Label: "SubLibraries",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsSubLibrariesNodeExpanded", inst.IsSubLibrariesNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: SubLibrariesWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "SubLibrariesWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SubLibrariesWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SubLibrariesWhoseNodeIsExpanded",
				Label: "SubLibrariesWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("NbPixPerCharacter", inst.NbPixPerCharacter, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LogoSVGFile", inst.LogoSVGFile, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpandedTmp", inst.IsExpandedTmp, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.SubLibraries {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "SubLibraries", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.SubLibrariesWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "SubLibrariesWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}
	case *models.MapObject:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.MapObjectUses {
				if src.Map == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.MapObjectUse", "Map", refNames, formGroup, probe.formStage)
		}
	case *models.MapObjectUse:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Map", inst.Map, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.MapObject](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Analysiss {
				for _, target := range src.MapUse {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Analysis", "MapUse", refNames, formGroup, probe.formStage)
		}
	case *models.Parameter:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsResponse", inst.IsResponse, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Start", inst.Start, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("End", inst.End, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Force", inst.Force, probe.formStage, formGroup)

		{
			// Slice of pointers: GroupUse
			div := (&form.FormDiv{Name: "GroupUse"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GroupUse {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GroupUse",
				Label: "GroupUse",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: DocumentUse
			div := (&form.FormDiv{Name: "DocumentUse"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DocumentUse {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DocumentUse",
				Label: "DocumentUse",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: GeoObjectUse
			div := (&form.FormDiv{Name: "GeoObjectUse"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GeoObjectUse {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GeoObjectUse",
				Label: "GeoObjectUse",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Tag", inst.Tag, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ActorStateTransitions {
				for _, target := range src.Justifications {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ActorStateTransition", "Justifications", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ParametersWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ParametersWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ParameterShapes {
				if src.Parameter == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ParameterShape", "Parameter", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ParametersAggregates {
				for _, target := range src.Parameters {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ParametersAggregate", "Parameters", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Scenarios {
				for _, target := range src.Parameters {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Scenario", "Parameters", refNames, formGroup, probe.formStage)
		}
	case *models.ParameterCategory:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: ParameterUse
			div := (&form.FormDiv{Name: "ParameterUse"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ParameterUse {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ParameterUse",
				Label: "ParameterUse",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ParameterCategoryUses {
				if src.ParameterCategory == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ParameterCategoryUse", "ParameterCategory", refNames, formGroup, probe.formStage)
		}
	case *models.ParameterCategoryUse:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ParameterCategory", inst.ParameterCategory, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ParameterCategory](), probe.formStage)
	case *models.ParameterShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Parameter", inst.Parameter, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Parameter](), probe.formStage)
		StageSetEnumStringFieldToForm("Direction", inst.Direction, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("ShapeIsComputedFromModel", inst.ShapeIsComputedFromModel, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ParameterShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ParameterShapes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ParameterCategorys {
				for _, target := range src.ParameterUse {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ParameterCategory", "ParameterUse", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Repositorys {
				for _, target := range src.ParameterUse {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Repository", "ParameterUse", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Workspaces {
				if src.Default_ParameterShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Workspace", "Default_ParameterShape", refNames, formGroup, probe.formStage)
		}
	case *models.ParametersAggregate:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Tag", inst.Tag, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)

		{
			// Slice of pointers: Parameters
			div := (&form.FormDiv{Name: "Parameters"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Parameters {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Parameters",
				Label: "Parameters",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ParametersAggregatesWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ParametersAggregatesWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ParametersAggregateShapes {
				if src.ScenarioParameter == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ParametersAggregateShape", "ScenarioParameter", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Scenarios {
				for _, target := range src.ParametersAggretates {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Scenario", "ParametersAggretates", refNames, formGroup, probe.formStage)
		}
	case *models.ParametersAggregateShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ScenarioParameter", inst.ScenarioParameter, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ParametersAggregate](), probe.formStage)
		StageSetEnumStringFieldToForm("Direction", inst.Direction, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ScenarioParameterShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ScenarioParameterShapes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Workspaces {
				if src.Default_ScenarioParameterShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Workspace", "Default_ScenarioParameterShape", refNames, formGroup, probe.formStage)
		}
	case *models.Position:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Date", inst.Date, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Ordinate", inst.Ordinate, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
	case *models.Repository:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: ParameterUse
			div := (&form.FormDiv{Name: "ParameterUse"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ParameterUse {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ParameterUse",
				Label: "ParameterUse",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: GroupUse
			div := (&form.FormDiv{Name: "GroupUse"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GroupUse {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GroupUse",
				Label: "GroupUse",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
	case *models.Scenario:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)

		{
			// Slice of pointers: Diagrams
			div := (&form.FormDiv{Name: "Diagrams"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Diagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Diagrams",
				Label: "Diagrams",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsDiagramsNodeExpanded", inst.IsDiagramsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: ActorStates
			div := (&form.FormDiv{Name: "ActorStates"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ActorStates {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ActorStates",
				Label: "ActorStates",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsActorStatesNodeExpanded", inst.IsActorStatesNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: ActorStateTransitions
			div := (&form.FormDiv{Name: "ActorStateTransitions"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ActorStateTransitions {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ActorStateTransitions",
				Label: "ActorStateTransitions",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsActorStateTransitionsNodeExpanded", inst.IsActorStateTransitionsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: EvolutionDirections
			div := (&form.FormDiv{Name: "EvolutionDirections"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.EvolutionDirections {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "EvolutionDirections",
				Label: "EvolutionDirections",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsEvolutionDirectionsNodeExpanded", inst.IsEvolutionDirectionsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: Parameters
			div := (&form.FormDiv{Name: "Parameters"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Parameters {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Parameters",
				Label: "Parameters",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsParametersNodeExpanded", inst.IsParametersNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: ParametersAggretates
			div := (&form.FormDiv{Name: "ParametersAggretates"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ParametersAggretates {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ParametersAggretates",
				Label: "ParametersAggretates",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsParametersAggretatesNodeExpanded", inst.IsParametersAggretatesNodeExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Analysiss {
				for _, target := range src.Scenarios {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Analysis", "Scenarios", refNames, formGroup, probe.formStage)
		}
	case *models.User:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.UserUses {
				if src.User == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.UserUse", "User", refNames, formGroup, probe.formStage)
		}
	case *models.UserUse:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("User", inst.User, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.User](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Groups {
				for _, target := range src.UserUse {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Group", "UserUse", refNames, formGroup, probe.formStage)
		}
	case *models.Workspace:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SelectedDiagram", inst.SelectedDiagram, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Diagram](), probe.formStage)
		StageSetAssociationFieldToForm("Default_EvolutionDirectionShape", inst.Default_EvolutionDirectionShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.EvolutionDirectionShape](), probe.formStage)
		StageSetAssociationFieldToForm("Default_ParameterShape", inst.Default_ParameterShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ParameterShape](), probe.formStage)
		StageSetAssociationFieldToForm("Default_ScenarioParameterShape", inst.Default_ScenarioParameterShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ParametersAggregateShape](), probe.formStage)
		StageSetAssociationFieldToForm("Default_ActorStateShape", inst.Default_ActorStateShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ActorStateShape](), probe.formStage)
		StageSetAssociationFieldToForm("Default_ActorStateTransitionShape", inst.Default_ActorStateTransitionShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ActorStateTransitionShape](), probe.formStage)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
	default:
		_ = inst
	}
}
