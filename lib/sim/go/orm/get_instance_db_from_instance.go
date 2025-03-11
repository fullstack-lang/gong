// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/lib/sim/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Command:
		commandInstance := any(concreteInstance).(*models.Command)
		ret2 := backRepo.BackRepoCommand.GetCommandDBFromCommandPtr(commandInstance)
		ret = any(ret2).(*T2)
	case *models.DummyAgent:
		dummyagentInstance := any(concreteInstance).(*models.DummyAgent)
		ret2 := backRepo.BackRepoDummyAgent.GetDummyAgentDBFromDummyAgentPtr(dummyagentInstance)
		ret = any(ret2).(*T2)
	case *models.Engine:
		engineInstance := any(concreteInstance).(*models.Engine)
		ret2 := backRepo.BackRepoEngine.GetEngineDBFromEnginePtr(engineInstance)
		ret = any(ret2).(*T2)
	case *models.Event:
		eventInstance := any(concreteInstance).(*models.Event)
		ret2 := backRepo.BackRepoEvent.GetEventDBFromEventPtr(eventInstance)
		ret = any(ret2).(*T2)
	case *models.Status:
		statusInstance := any(concreteInstance).(*models.Status)
		ret2 := backRepo.BackRepoStatus.GetStatusDBFromStatusPtr(statusInstance)
		ret = any(ret2).(*T2)
	case *models.UpdateState:
		updatestateInstance := any(concreteInstance).(*models.UpdateState)
		ret2 := backRepo.BackRepoUpdateState.GetUpdateStateDBFromUpdateStatePtr(updatestateInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Command:
		tmp := GetInstanceDBFromInstance[models.Command, CommandDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DummyAgent:
		tmp := GetInstanceDBFromInstance[models.DummyAgent, DummyAgentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Engine:
		tmp := GetInstanceDBFromInstance[models.Engine, EngineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Event:
		tmp := GetInstanceDBFromInstance[models.Event, EventDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Status:
		tmp := GetInstanceDBFromInstance[models.Status, StatusDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.UpdateState:
		tmp := GetInstanceDBFromInstance[models.UpdateState, UpdateStateDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Command:
		tmp := GetInstanceDBFromInstance[models.Command, CommandDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DummyAgent:
		tmp := GetInstanceDBFromInstance[models.DummyAgent, DummyAgentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Engine:
		tmp := GetInstanceDBFromInstance[models.Engine, EngineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Event:
		tmp := GetInstanceDBFromInstance[models.Event, EventDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Status:
		tmp := GetInstanceDBFromInstance[models.Status, StatusDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.UpdateState:
		tmp := GetInstanceDBFromInstance[models.UpdateState, UpdateStateDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
