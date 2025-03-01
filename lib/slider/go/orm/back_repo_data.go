// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	CheckboxAPIs []*CheckboxAPI

	GroupAPIs []*GroupAPI

	LayoutAPIs []*LayoutAPI

	SliderAPIs []*SliderAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, checkboxDB := range backRepo.BackRepoCheckbox.Map_CheckboxDBID_CheckboxDB {

		var checkboxAPI CheckboxAPI
		checkboxAPI.ID = checkboxDB.ID
		checkboxAPI.CheckboxPointersEncoding = checkboxDB.CheckboxPointersEncoding
		checkboxDB.CopyBasicFieldsToCheckbox_WOP(&checkboxAPI.Checkbox_WOP)

		backRepoData.CheckboxAPIs = append(backRepoData.CheckboxAPIs, &checkboxAPI)
	}

	for _, groupDB := range backRepo.BackRepoGroup.Map_GroupDBID_GroupDB {

		var groupAPI GroupAPI
		groupAPI.ID = groupDB.ID
		groupAPI.GroupPointersEncoding = groupDB.GroupPointersEncoding
		groupDB.CopyBasicFieldsToGroup_WOP(&groupAPI.Group_WOP)

		backRepoData.GroupAPIs = append(backRepoData.GroupAPIs, &groupAPI)
	}

	for _, layoutDB := range backRepo.BackRepoLayout.Map_LayoutDBID_LayoutDB {

		var layoutAPI LayoutAPI
		layoutAPI.ID = layoutDB.ID
		layoutAPI.LayoutPointersEncoding = layoutDB.LayoutPointersEncoding
		layoutDB.CopyBasicFieldsToLayout_WOP(&layoutAPI.Layout_WOP)

		backRepoData.LayoutAPIs = append(backRepoData.LayoutAPIs, &layoutAPI)
	}

	for _, sliderDB := range backRepo.BackRepoSlider.Map_SliderDBID_SliderDB {

		var sliderAPI SliderAPI
		sliderAPI.ID = sliderDB.ID
		sliderAPI.SliderPointersEncoding = sliderDB.SliderPointersEncoding
		sliderDB.CopyBasicFieldsToSlider_WOP(&sliderAPI.Slider_WOP)

		backRepoData.SliderAPIs = append(backRepoData.SliderAPIs, &sliderAPI)
	}

}
