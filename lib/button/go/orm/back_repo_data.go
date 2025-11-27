// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	ButtonAPIs []*ButtonAPI

	ButtonToggleAPIs []*ButtonToggleAPI

	GroupAPIs []*GroupAPI

	GroupToogleAPIs []*GroupToogleAPI

	LayoutAPIs []*LayoutAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, buttonDB := range backRepo.BackRepoButton.Map_ButtonDBID_ButtonDB {

		var buttonAPI ButtonAPI
		buttonAPI.ID = buttonDB.ID
		buttonAPI.ButtonPointersEncoding = buttonDB.ButtonPointersEncoding
		buttonDB.CopyBasicFieldsToButton_WOP(&buttonAPI.Button_WOP)

		backRepoData.ButtonAPIs = append(backRepoData.ButtonAPIs, &buttonAPI)
	}

	for _, buttontoggleDB := range backRepo.BackRepoButtonToggle.Map_ButtonToggleDBID_ButtonToggleDB {

		var buttontoggleAPI ButtonToggleAPI
		buttontoggleAPI.ID = buttontoggleDB.ID
		buttontoggleAPI.ButtonTogglePointersEncoding = buttontoggleDB.ButtonTogglePointersEncoding
		buttontoggleDB.CopyBasicFieldsToButtonToggle_WOP(&buttontoggleAPI.ButtonToggle_WOP)

		backRepoData.ButtonToggleAPIs = append(backRepoData.ButtonToggleAPIs, &buttontoggleAPI)
	}

	for _, groupDB := range backRepo.BackRepoGroup.Map_GroupDBID_GroupDB {

		var groupAPI GroupAPI
		groupAPI.ID = groupDB.ID
		groupAPI.GroupPointersEncoding = groupDB.GroupPointersEncoding
		groupDB.CopyBasicFieldsToGroup_WOP(&groupAPI.Group_WOP)

		backRepoData.GroupAPIs = append(backRepoData.GroupAPIs, &groupAPI)
	}

	for _, grouptoogleDB := range backRepo.BackRepoGroupToogle.Map_GroupToogleDBID_GroupToogleDB {

		var grouptoogleAPI GroupToogleAPI
		grouptoogleAPI.ID = grouptoogleDB.ID
		grouptoogleAPI.GroupTooglePointersEncoding = grouptoogleDB.GroupTooglePointersEncoding
		grouptoogleDB.CopyBasicFieldsToGroupToogle_WOP(&grouptoogleAPI.GroupToogle_WOP)

		backRepoData.GroupToogleAPIs = append(backRepoData.GroupToogleAPIs, &grouptoogleAPI)
	}

	for _, layoutDB := range backRepo.BackRepoLayout.Map_LayoutDBID_LayoutDB {

		var layoutAPI LayoutAPI
		layoutAPI.ID = layoutDB.ID
		layoutAPI.LayoutPointersEncoding = layoutDB.LayoutPointersEncoding
		layoutDB.CopyBasicFieldsToLayout_WOP(&layoutAPI.Layout_WOP)

		backRepoData.LayoutAPIs = append(backRepoData.LayoutAPIs, &layoutAPI)
	}

}
