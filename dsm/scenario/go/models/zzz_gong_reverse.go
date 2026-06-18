// generated code - do not edit
package models

// insertion point
func (inst *ActorState) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ActorStatesWhoseNodeIsExpanded":
			if _diagram, ok := stage.Diagram_ActorStatesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	case "Scenario":
		switch reverseField.Fieldname {
		case "ActorStates":
			if _scenario, ok := stage.Scenario_ActorStates_reverseMap[inst]; ok {
				res = _scenario.Name
			}
		}
	}
	return
}

func (inst *ActorStateShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ActorStateShapes":
			if _diagram, ok := stage.Diagram_ActorStateShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *ActorStateTransition) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ActorStateTransitionsWhoseNodeIsExpanded":
			if _diagram, ok := stage.Diagram_ActorStateTransitionsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	case "Scenario":
		switch reverseField.Fieldname {
		case "ActorStateTransitions":
			if _scenario, ok := stage.Scenario_ActorStateTransitions_reverseMap[inst]; ok {
				res = _scenario.Name
			}
		}
	}
	return
}

func (inst *ActorStateTransitionShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ActorStateTransitionShapes":
			if _diagram, ok := stage.Diagram_ActorStateTransitionShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *Analysis) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "Analyses":
			if _library, ok := stage.Library_Analyses_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *ControlPointShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ActorStateTransitionShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			if _actorstatetransitionshape, ok := stage.ActorStateTransitionShape_ControlPointShapes_reverseMap[inst]; ok {
				res = _actorstatetransitionshape.Name
			}
		}
	}
	return
}

func (inst *Diagram) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Scenario":
		switch reverseField.Fieldname {
		case "Diagrams":
			if _scenario, ok := stage.Scenario_Diagrams_reverseMap[inst]; ok {
				res = _scenario.Name
			}
		}
	}
	return
}

func (inst *Document) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *DocumentUse) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Parameter":
		switch reverseField.Fieldname {
		case "DocumentUse":
			if _parameter, ok := stage.Parameter_DocumentUse_reverseMap[inst]; ok {
				res = _parameter.Name
			}
		}
	}
	return
}

func (inst *EvolutionDirection) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "EvolutionDirectionsWhoseNodeIsExpanded":
			if _diagram, ok := stage.Diagram_EvolutionDirectionsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	case "Scenario":
		switch reverseField.Fieldname {
		case "EvolutionDirections":
			if _scenario, ok := stage.Scenario_EvolutionDirections_reverseMap[inst]; ok {
				res = _scenario.Name
			}
		}
	}
	return
}

func (inst *EvolutionDirectionShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "EvolutionDirectionShapes":
			if _diagram, ok := stage.Diagram_EvolutionDirectionShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *Foo) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *GeoObject) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *GeoObjectUse) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Analysis":
		switch reverseField.Fieldname {
		case "GeoObjectUse":
			if _analysis, ok := stage.Analysis_GeoObjectUse_reverseMap[inst]; ok {
				res = _analysis.Name
			}
		}
	case "Document":
		switch reverseField.Fieldname {
		case "GeoObjectUse":
			if _document, ok := stage.Document_GeoObjectUse_reverseMap[inst]; ok {
				res = _document.Name
			}
		}
	case "Parameter":
		switch reverseField.Fieldname {
		case "GeoObjectUse":
			if _parameter, ok := stage.Parameter_GeoObjectUse_reverseMap[inst]; ok {
				res = _parameter.Name
			}
		}
	}
	return
}

func (inst *Group) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *GroupUse) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Analysis":
		switch reverseField.Fieldname {
		case "GroupUse":
			if _analysis, ok := stage.Analysis_GroupUse_reverseMap[inst]; ok {
				res = _analysis.Name
			}
		}
	case "Parameter":
		switch reverseField.Fieldname {
		case "GroupUse":
			if _parameter, ok := stage.Parameter_GroupUse_reverseMap[inst]; ok {
				res = _parameter.Name
			}
		}
	case "Repository":
		switch reverseField.Fieldname {
		case "GroupUse":
			if _repository, ok := stage.Repository_GroupUse_reverseMap[inst]; ok {
				res = _repository.Name
			}
		}
	}
	return
}

func (inst *Library) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "SubLibraries":
			if _library, ok := stage.Library_SubLibraries_reverseMap[inst]; ok {
				res = _library.Name
			}
		case "SubLibrariesWhoseNodeIsExpanded":
			if _library, ok := stage.Library_SubLibrariesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *MapObject) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *MapObjectUse) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Analysis":
		switch reverseField.Fieldname {
		case "MapUse":
			if _analysis, ok := stage.Analysis_MapUse_reverseMap[inst]; ok {
				res = _analysis.Name
			}
		}
	}
	return
}

func (inst *Parameter) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ActorStateTransition":
		switch reverseField.Fieldname {
		case "Justifications":
			if _actorstatetransition, ok := stage.ActorStateTransition_Justifications_reverseMap[inst]; ok {
				res = _actorstatetransition.Name
			}
		}
	case "Diagram":
		switch reverseField.Fieldname {
		case "ParametersWhoseNodeIsExpanded":
			if _diagram, ok := stage.Diagram_ParametersWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	case "ParametersAggregate":
		switch reverseField.Fieldname {
		case "Parameters":
			if _parametersaggregate, ok := stage.ParametersAggregate_Parameters_reverseMap[inst]; ok {
				res = _parametersaggregate.Name
			}
		}
	case "Scenario":
		switch reverseField.Fieldname {
		case "Parameters":
			if _scenario, ok := stage.Scenario_Parameters_reverseMap[inst]; ok {
				res = _scenario.Name
			}
		}
	}
	return
}

func (inst *ParameterCategory) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *ParameterCategoryUse) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *ParameterShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ParameterShapes":
			if _diagram, ok := stage.Diagram_ParameterShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	case "ParameterCategory":
		switch reverseField.Fieldname {
		case "ParameterUse":
			if _parametercategory, ok := stage.ParameterCategory_ParameterUse_reverseMap[inst]; ok {
				res = _parametercategory.Name
			}
		}
	case "Repository":
		switch reverseField.Fieldname {
		case "ParameterUse":
			if _repository, ok := stage.Repository_ParameterUse_reverseMap[inst]; ok {
				res = _repository.Name
			}
		}
	}
	return
}

func (inst *ParametersAggregate) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ParametersAggregatesWhoseNodeIsExpanded":
			if _diagram, ok := stage.Diagram_ParametersAggregatesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	case "Scenario":
		switch reverseField.Fieldname {
		case "ParametersAggretates":
			if _scenario, ok := stage.Scenario_ParametersAggretates_reverseMap[inst]; ok {
				res = _scenario.Name
			}
		}
	}
	return
}

func (inst *ParametersAggregateShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ScenarioParameterShapes":
			if _diagram, ok := stage.Diagram_ScenarioParameterShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *Position) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Repository) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Scenario) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Analysis":
		switch reverseField.Fieldname {
		case "Scenarios":
			if _analysis, ok := stage.Analysis_Scenarios_reverseMap[inst]; ok {
				res = _analysis.Name
			}
		}
	}
	return
}

func (inst *User) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *UserUse) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Group":
		switch reverseField.Fieldname {
		case "UserUse":
			if _group, ok := stage.Group_UserUse_reverseMap[inst]; ok {
				res = _group.Name
			}
		}
	}
	return
}

func (inst *Workspace) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

// insertion point
func (inst *ActorState) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ActorStatesWhoseNodeIsExpanded":
			res = stage.Diagram_ActorStatesWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Scenario":
		switch reverseField.Fieldname {
		case "ActorStates":
			res = stage.Scenario_ActorStates_reverseMap[inst]
		}
	}
	return res
}

func (inst *ActorStateShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ActorStateShapes":
			res = stage.Diagram_ActorStateShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *ActorStateTransition) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ActorStateTransitionsWhoseNodeIsExpanded":
			res = stage.Diagram_ActorStateTransitionsWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Scenario":
		switch reverseField.Fieldname {
		case "ActorStateTransitions":
			res = stage.Scenario_ActorStateTransitions_reverseMap[inst]
		}
	}
	return res
}

func (inst *ActorStateTransitionShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ActorStateTransitionShapes":
			res = stage.Diagram_ActorStateTransitionShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Analysis) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "Analyses":
			res = stage.Library_Analyses_reverseMap[inst]
		}
	}
	return res
}

func (inst *ControlPointShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "ActorStateTransitionShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			res = stage.ActorStateTransitionShape_ControlPointShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Diagram) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Scenario":
		switch reverseField.Fieldname {
		case "Diagrams":
			res = stage.Scenario_Diagrams_reverseMap[inst]
		}
	}
	return res
}

func (inst *Document) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *DocumentUse) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Parameter":
		switch reverseField.Fieldname {
		case "DocumentUse":
			res = stage.Parameter_DocumentUse_reverseMap[inst]
		}
	}
	return res
}

func (inst *EvolutionDirection) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "EvolutionDirectionsWhoseNodeIsExpanded":
			res = stage.Diagram_EvolutionDirectionsWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Scenario":
		switch reverseField.Fieldname {
		case "EvolutionDirections":
			res = stage.Scenario_EvolutionDirections_reverseMap[inst]
		}
	}
	return res
}

func (inst *EvolutionDirectionShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "EvolutionDirectionShapes":
			res = stage.Diagram_EvolutionDirectionShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Foo) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *GeoObject) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *GeoObjectUse) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Analysis":
		switch reverseField.Fieldname {
		case "GeoObjectUse":
			res = stage.Analysis_GeoObjectUse_reverseMap[inst]
		}
	case "Document":
		switch reverseField.Fieldname {
		case "GeoObjectUse":
			res = stage.Document_GeoObjectUse_reverseMap[inst]
		}
	case "Parameter":
		switch reverseField.Fieldname {
		case "GeoObjectUse":
			res = stage.Parameter_GeoObjectUse_reverseMap[inst]
		}
	}
	return res
}

func (inst *Group) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *GroupUse) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Analysis":
		switch reverseField.Fieldname {
		case "GroupUse":
			res = stage.Analysis_GroupUse_reverseMap[inst]
		}
	case "Parameter":
		switch reverseField.Fieldname {
		case "GroupUse":
			res = stage.Parameter_GroupUse_reverseMap[inst]
		}
	case "Repository":
		switch reverseField.Fieldname {
		case "GroupUse":
			res = stage.Repository_GroupUse_reverseMap[inst]
		}
	}
	return res
}

func (inst *Library) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "SubLibraries":
			res = stage.Library_SubLibraries_reverseMap[inst]
		case "SubLibrariesWhoseNodeIsExpanded":
			res = stage.Library_SubLibrariesWhoseNodeIsExpanded_reverseMap[inst]
		}
	}
	return res
}

func (inst *MapObject) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *MapObjectUse) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Analysis":
		switch reverseField.Fieldname {
		case "MapUse":
			res = stage.Analysis_MapUse_reverseMap[inst]
		}
	}
	return res
}

func (inst *Parameter) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "ActorStateTransition":
		switch reverseField.Fieldname {
		case "Justifications":
			res = stage.ActorStateTransition_Justifications_reverseMap[inst]
		}
	case "Diagram":
		switch reverseField.Fieldname {
		case "ParametersWhoseNodeIsExpanded":
			res = stage.Diagram_ParametersWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "ParametersAggregate":
		switch reverseField.Fieldname {
		case "Parameters":
			res = stage.ParametersAggregate_Parameters_reverseMap[inst]
		}
	case "Scenario":
		switch reverseField.Fieldname {
		case "Parameters":
			res = stage.Scenario_Parameters_reverseMap[inst]
		}
	}
	return res
}

func (inst *ParameterCategory) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *ParameterCategoryUse) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *ParameterShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ParameterShapes":
			res = stage.Diagram_ParameterShapes_reverseMap[inst]
		}
	case "ParameterCategory":
		switch reverseField.Fieldname {
		case "ParameterUse":
			res = stage.ParameterCategory_ParameterUse_reverseMap[inst]
		}
	case "Repository":
		switch reverseField.Fieldname {
		case "ParameterUse":
			res = stage.Repository_ParameterUse_reverseMap[inst]
		}
	}
	return res
}

func (inst *ParametersAggregate) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ParametersAggregatesWhoseNodeIsExpanded":
			res = stage.Diagram_ParametersAggregatesWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Scenario":
		switch reverseField.Fieldname {
		case "ParametersAggretates":
			res = stage.Scenario_ParametersAggretates_reverseMap[inst]
		}
	}
	return res
}

func (inst *ParametersAggregateShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ScenarioParameterShapes":
			res = stage.Diagram_ScenarioParameterShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Position) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Repository) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Scenario) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Analysis":
		switch reverseField.Fieldname {
		case "Scenarios":
			res = stage.Analysis_Scenarios_reverseMap[inst]
		}
	}
	return res
}

func (inst *User) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *UserUse) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Group":
		switch reverseField.Fieldname {
		case "UserUse":
			res = stage.Group_UserUse_reverseMap[inst]
		}
	}
	return res
}

func (inst *Workspace) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}
