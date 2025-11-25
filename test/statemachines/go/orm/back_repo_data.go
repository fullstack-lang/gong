// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	ArchitectureAPIs []*ArchitectureAPI

	DiagramAPIs []*DiagramAPI

	KillAPIs []*KillAPI

	MessageAPIs []*MessageAPI

	MessageTypeAPIs []*MessageTypeAPI

	ObjectAPIs []*ObjectAPI

	RoleAPIs []*RoleAPI

	StateAPIs []*StateAPI

	StateMachineAPIs []*StateMachineAPI

	StateShapeAPIs []*StateShapeAPI

	TransitionAPIs []*TransitionAPI

	Transition_ShapeAPIs []*Transition_ShapeAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, architectureDB := range backRepo.BackRepoArchitecture.Map_ArchitectureDBID_ArchitectureDB {

		var architectureAPI ArchitectureAPI
		architectureAPI.ID = architectureDB.ID
		architectureAPI.ArchitecturePointersEncoding = architectureDB.ArchitecturePointersEncoding
		architectureDB.CopyBasicFieldsToArchitecture_WOP(&architectureAPI.Architecture_WOP)

		backRepoData.ArchitectureAPIs = append(backRepoData.ArchitectureAPIs, &architectureAPI)
	}

	for _, diagramDB := range backRepo.BackRepoDiagram.Map_DiagramDBID_DiagramDB {

		var diagramAPI DiagramAPI
		diagramAPI.ID = diagramDB.ID
		diagramAPI.DiagramPointersEncoding = diagramDB.DiagramPointersEncoding
		diagramDB.CopyBasicFieldsToDiagram_WOP(&diagramAPI.Diagram_WOP)

		backRepoData.DiagramAPIs = append(backRepoData.DiagramAPIs, &diagramAPI)
	}

	for _, killDB := range backRepo.BackRepoKill.Map_KillDBID_KillDB {

		var killAPI KillAPI
		killAPI.ID = killDB.ID
		killAPI.KillPointersEncoding = killDB.KillPointersEncoding
		killDB.CopyBasicFieldsToKill_WOP(&killAPI.Kill_WOP)

		backRepoData.KillAPIs = append(backRepoData.KillAPIs, &killAPI)
	}

	for _, messageDB := range backRepo.BackRepoMessage.Map_MessageDBID_MessageDB {

		var messageAPI MessageAPI
		messageAPI.ID = messageDB.ID
		messageAPI.MessagePointersEncoding = messageDB.MessagePointersEncoding
		messageDB.CopyBasicFieldsToMessage_WOP(&messageAPI.Message_WOP)

		backRepoData.MessageAPIs = append(backRepoData.MessageAPIs, &messageAPI)
	}

	for _, messagetypeDB := range backRepo.BackRepoMessageType.Map_MessageTypeDBID_MessageTypeDB {

		var messagetypeAPI MessageTypeAPI
		messagetypeAPI.ID = messagetypeDB.ID
		messagetypeAPI.MessageTypePointersEncoding = messagetypeDB.MessageTypePointersEncoding
		messagetypeDB.CopyBasicFieldsToMessageType_WOP(&messagetypeAPI.MessageType_WOP)

		backRepoData.MessageTypeAPIs = append(backRepoData.MessageTypeAPIs, &messagetypeAPI)
	}

	for _, objectDB := range backRepo.BackRepoObject.Map_ObjectDBID_ObjectDB {

		var objectAPI ObjectAPI
		objectAPI.ID = objectDB.ID
		objectAPI.ObjectPointersEncoding = objectDB.ObjectPointersEncoding
		objectDB.CopyBasicFieldsToObject_WOP(&objectAPI.Object_WOP)

		backRepoData.ObjectAPIs = append(backRepoData.ObjectAPIs, &objectAPI)
	}

	for _, roleDB := range backRepo.BackRepoRole.Map_RoleDBID_RoleDB {

		var roleAPI RoleAPI
		roleAPI.ID = roleDB.ID
		roleAPI.RolePointersEncoding = roleDB.RolePointersEncoding
		roleDB.CopyBasicFieldsToRole_WOP(&roleAPI.Role_WOP)

		backRepoData.RoleAPIs = append(backRepoData.RoleAPIs, &roleAPI)
	}

	for _, stateDB := range backRepo.BackRepoState.Map_StateDBID_StateDB {

		var stateAPI StateAPI
		stateAPI.ID = stateDB.ID
		stateAPI.StatePointersEncoding = stateDB.StatePointersEncoding
		stateDB.CopyBasicFieldsToState_WOP(&stateAPI.State_WOP)

		backRepoData.StateAPIs = append(backRepoData.StateAPIs, &stateAPI)
	}

	for _, statemachineDB := range backRepo.BackRepoStateMachine.Map_StateMachineDBID_StateMachineDB {

		var statemachineAPI StateMachineAPI
		statemachineAPI.ID = statemachineDB.ID
		statemachineAPI.StateMachinePointersEncoding = statemachineDB.StateMachinePointersEncoding
		statemachineDB.CopyBasicFieldsToStateMachine_WOP(&statemachineAPI.StateMachine_WOP)

		backRepoData.StateMachineAPIs = append(backRepoData.StateMachineAPIs, &statemachineAPI)
	}

	for _, stateshapeDB := range backRepo.BackRepoStateShape.Map_StateShapeDBID_StateShapeDB {

		var stateshapeAPI StateShapeAPI
		stateshapeAPI.ID = stateshapeDB.ID
		stateshapeAPI.StateShapePointersEncoding = stateshapeDB.StateShapePointersEncoding
		stateshapeDB.CopyBasicFieldsToStateShape_WOP(&stateshapeAPI.StateShape_WOP)

		backRepoData.StateShapeAPIs = append(backRepoData.StateShapeAPIs, &stateshapeAPI)
	}

	for _, transitionDB := range backRepo.BackRepoTransition.Map_TransitionDBID_TransitionDB {

		var transitionAPI TransitionAPI
		transitionAPI.ID = transitionDB.ID
		transitionAPI.TransitionPointersEncoding = transitionDB.TransitionPointersEncoding
		transitionDB.CopyBasicFieldsToTransition_WOP(&transitionAPI.Transition_WOP)

		backRepoData.TransitionAPIs = append(backRepoData.TransitionAPIs, &transitionAPI)
	}

	for _, transition_shapeDB := range backRepo.BackRepoTransition_Shape.Map_Transition_ShapeDBID_Transition_ShapeDB {

		var transition_shapeAPI Transition_ShapeAPI
		transition_shapeAPI.ID = transition_shapeDB.ID
		transition_shapeAPI.Transition_ShapePointersEncoding = transition_shapeDB.Transition_ShapePointersEncoding
		transition_shapeDB.CopyBasicFieldsToTransition_Shape_WOP(&transition_shapeAPI.Transition_Shape_WOP)

		backRepoData.Transition_ShapeAPIs = append(backRepoData.Transition_ShapeAPIs, &transition_shapeAPI)
	}

}
