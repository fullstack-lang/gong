// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/lib/gantt/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Arrow:
		arrowInstance := any(concreteInstance).(*models.Arrow)
		ret2 := backRepo.BackRepoArrow.GetArrowDBFromArrowPtr(arrowInstance)
		ret = any(ret2).(*T2)
	case *models.Bar:
		barInstance := any(concreteInstance).(*models.Bar)
		ret2 := backRepo.BackRepoBar.GetBarDBFromBarPtr(barInstance)
		ret = any(ret2).(*T2)
	case *models.Gantt:
		ganttInstance := any(concreteInstance).(*models.Gantt)
		ret2 := backRepo.BackRepoGantt.GetGanttDBFromGanttPtr(ganttInstance)
		ret = any(ret2).(*T2)
	case *models.Group:
		groupInstance := any(concreteInstance).(*models.Group)
		ret2 := backRepo.BackRepoGroup.GetGroupDBFromGroupPtr(groupInstance)
		ret = any(ret2).(*T2)
	case *models.Lane:
		laneInstance := any(concreteInstance).(*models.Lane)
		ret2 := backRepo.BackRepoLane.GetLaneDBFromLanePtr(laneInstance)
		ret = any(ret2).(*T2)
	case *models.LaneUse:
		laneuseInstance := any(concreteInstance).(*models.LaneUse)
		ret2 := backRepo.BackRepoLaneUse.GetLaneUseDBFromLaneUsePtr(laneuseInstance)
		ret = any(ret2).(*T2)
	case *models.Milestone:
		milestoneInstance := any(concreteInstance).(*models.Milestone)
		ret2 := backRepo.BackRepoMilestone.GetMilestoneDBFromMilestonePtr(milestoneInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Arrow:
		tmp := GetInstanceDBFromInstance[models.Arrow, ArrowDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Bar:
		tmp := GetInstanceDBFromInstance[models.Bar, BarDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Gantt:
		tmp := GetInstanceDBFromInstance[models.Gantt, GanttDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Group:
		tmp := GetInstanceDBFromInstance[models.Group, GroupDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Lane:
		tmp := GetInstanceDBFromInstance[models.Lane, LaneDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LaneUse:
		tmp := GetInstanceDBFromInstance[models.LaneUse, LaneUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Milestone:
		tmp := GetInstanceDBFromInstance[models.Milestone, MilestoneDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Arrow:
		tmp := GetInstanceDBFromInstance[models.Arrow, ArrowDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Bar:
		tmp := GetInstanceDBFromInstance[models.Bar, BarDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Gantt:
		tmp := GetInstanceDBFromInstance[models.Gantt, GanttDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Group:
		tmp := GetInstanceDBFromInstance[models.Group, GroupDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Lane:
		tmp := GetInstanceDBFromInstance[models.Lane, LaneDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LaneUse:
		tmp := GetInstanceDBFromInstance[models.LaneUse, LaneUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Milestone:
		tmp := GetInstanceDBFromInstance[models.Milestone, MilestoneDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
