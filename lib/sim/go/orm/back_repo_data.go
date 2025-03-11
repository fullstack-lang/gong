// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	CommandAPIs []*CommandAPI

	DummyAgentAPIs []*DummyAgentAPI

	EngineAPIs []*EngineAPI

	EventAPIs []*EventAPI

	StatusAPIs []*StatusAPI

	UpdateStateAPIs []*UpdateStateAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, commandDB := range backRepo.BackRepoCommand.Map_CommandDBID_CommandDB {

		var commandAPI CommandAPI
		commandAPI.ID = commandDB.ID
		commandAPI.CommandPointersEncoding = commandDB.CommandPointersEncoding
		commandDB.CopyBasicFieldsToCommand_WOP(&commandAPI.Command_WOP)

		backRepoData.CommandAPIs = append(backRepoData.CommandAPIs, &commandAPI)
	}

	for _, dummyagentDB := range backRepo.BackRepoDummyAgent.Map_DummyAgentDBID_DummyAgentDB {

		var dummyagentAPI DummyAgentAPI
		dummyagentAPI.ID = dummyagentDB.ID
		dummyagentAPI.DummyAgentPointersEncoding = dummyagentDB.DummyAgentPointersEncoding
		dummyagentDB.CopyBasicFieldsToDummyAgent_WOP(&dummyagentAPI.DummyAgent_WOP)

		backRepoData.DummyAgentAPIs = append(backRepoData.DummyAgentAPIs, &dummyagentAPI)
	}

	for _, engineDB := range backRepo.BackRepoEngine.Map_EngineDBID_EngineDB {

		var engineAPI EngineAPI
		engineAPI.ID = engineDB.ID
		engineAPI.EnginePointersEncoding = engineDB.EnginePointersEncoding
		engineDB.CopyBasicFieldsToEngine_WOP(&engineAPI.Engine_WOP)

		backRepoData.EngineAPIs = append(backRepoData.EngineAPIs, &engineAPI)
	}

	for _, eventDB := range backRepo.BackRepoEvent.Map_EventDBID_EventDB {

		var eventAPI EventAPI
		eventAPI.ID = eventDB.ID
		eventAPI.EventPointersEncoding = eventDB.EventPointersEncoding
		eventDB.CopyBasicFieldsToEvent_WOP(&eventAPI.Event_WOP)

		backRepoData.EventAPIs = append(backRepoData.EventAPIs, &eventAPI)
	}

	for _, statusDB := range backRepo.BackRepoStatus.Map_StatusDBID_StatusDB {

		var statusAPI StatusAPI
		statusAPI.ID = statusDB.ID
		statusAPI.StatusPointersEncoding = statusDB.StatusPointersEncoding
		statusDB.CopyBasicFieldsToStatus_WOP(&statusAPI.Status_WOP)

		backRepoData.StatusAPIs = append(backRepoData.StatusAPIs, &statusAPI)
	}

	for _, updatestateDB := range backRepo.BackRepoUpdateState.Map_UpdateStateDBID_UpdateStateDB {

		var updatestateAPI UpdateStateAPI
		updatestateAPI.ID = updatestateDB.ID
		updatestateAPI.UpdateStatePointersEncoding = updatestateDB.UpdateStatePointersEncoding
		updatestateDB.CopyBasicFieldsToUpdateState_WOP(&updatestateAPI.UpdateState_WOP)

		backRepoData.UpdateStateAPIs = append(backRepoData.UpdateStateAPIs, &updatestateAPI)
	}

}
