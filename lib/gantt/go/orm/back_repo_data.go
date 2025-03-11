// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	ArrowAPIs []*ArrowAPI

	BarAPIs []*BarAPI

	GanttAPIs []*GanttAPI

	GroupAPIs []*GroupAPI

	LaneAPIs []*LaneAPI

	LaneUseAPIs []*LaneUseAPI

	MilestoneAPIs []*MilestoneAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, arrowDB := range backRepo.BackRepoArrow.Map_ArrowDBID_ArrowDB {

		var arrowAPI ArrowAPI
		arrowAPI.ID = arrowDB.ID
		arrowAPI.ArrowPointersEncoding = arrowDB.ArrowPointersEncoding
		arrowDB.CopyBasicFieldsToArrow_WOP(&arrowAPI.Arrow_WOP)

		backRepoData.ArrowAPIs = append(backRepoData.ArrowAPIs, &arrowAPI)
	}

	for _, barDB := range backRepo.BackRepoBar.Map_BarDBID_BarDB {

		var barAPI BarAPI
		barAPI.ID = barDB.ID
		barAPI.BarPointersEncoding = barDB.BarPointersEncoding
		barDB.CopyBasicFieldsToBar_WOP(&barAPI.Bar_WOP)

		backRepoData.BarAPIs = append(backRepoData.BarAPIs, &barAPI)
	}

	for _, ganttDB := range backRepo.BackRepoGantt.Map_GanttDBID_GanttDB {

		var ganttAPI GanttAPI
		ganttAPI.ID = ganttDB.ID
		ganttAPI.GanttPointersEncoding = ganttDB.GanttPointersEncoding
		ganttDB.CopyBasicFieldsToGantt_WOP(&ganttAPI.Gantt_WOP)

		backRepoData.GanttAPIs = append(backRepoData.GanttAPIs, &ganttAPI)
	}

	for _, groupDB := range backRepo.BackRepoGroup.Map_GroupDBID_GroupDB {

		var groupAPI GroupAPI
		groupAPI.ID = groupDB.ID
		groupAPI.GroupPointersEncoding = groupDB.GroupPointersEncoding
		groupDB.CopyBasicFieldsToGroup_WOP(&groupAPI.Group_WOP)

		backRepoData.GroupAPIs = append(backRepoData.GroupAPIs, &groupAPI)
	}

	for _, laneDB := range backRepo.BackRepoLane.Map_LaneDBID_LaneDB {

		var laneAPI LaneAPI
		laneAPI.ID = laneDB.ID
		laneAPI.LanePointersEncoding = laneDB.LanePointersEncoding
		laneDB.CopyBasicFieldsToLane_WOP(&laneAPI.Lane_WOP)

		backRepoData.LaneAPIs = append(backRepoData.LaneAPIs, &laneAPI)
	}

	for _, laneuseDB := range backRepo.BackRepoLaneUse.Map_LaneUseDBID_LaneUseDB {

		var laneuseAPI LaneUseAPI
		laneuseAPI.ID = laneuseDB.ID
		laneuseAPI.LaneUsePointersEncoding = laneuseDB.LaneUsePointersEncoding
		laneuseDB.CopyBasicFieldsToLaneUse_WOP(&laneuseAPI.LaneUse_WOP)

		backRepoData.LaneUseAPIs = append(backRepoData.LaneUseAPIs, &laneuseAPI)
	}

	for _, milestoneDB := range backRepo.BackRepoMilestone.Map_MilestoneDBID_MilestoneDB {

		var milestoneAPI MilestoneAPI
		milestoneAPI.ID = milestoneDB.ID
		milestoneAPI.MilestonePointersEncoding = milestoneDB.MilestonePointersEncoding
		milestoneDB.CopyBasicFieldsToMilestone_WOP(&milestoneAPI.Milestone_WOP)

		backRepoData.MilestoneAPIs = append(backRepoData.MilestoneAPIs, &milestoneAPI)
	}

}
