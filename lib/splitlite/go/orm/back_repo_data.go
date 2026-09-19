// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	AsSplitAPIs []*AsSplitAPI

	AsSplitAreaAPIs []*AsSplitAreaAPI

	ButtonAPIs []*ButtonAPI

	FavIconAPIs []*FavIconAPI

	FormAPIs []*FormAPI

	LoadAPIs []*LoadAPI

	LogoOnTheLeftAPIs []*LogoOnTheLeftAPI

	LogoOnTheRightAPIs []*LogoOnTheRightAPI

	SplitAPIs []*SplitAPI

	SvgAPIs []*SvgAPI

	TableAPIs []*TableAPI

	TitleAPIs []*TitleAPI

	TreeAPIs []*TreeAPI

	ViewAPIs []*ViewAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, assplitDB := range backRepo.BackRepoAsSplit.Map_AsSplitDBID_AsSplitDB {

		var assplitAPI AsSplitAPI
		assplitAPI.ID = assplitDB.ID
		assplitAPI.AsSplitPointersEncoding = assplitDB.AsSplitPointersEncoding
		assplitDB.CopyBasicFieldsToAsSplit_WOP(&assplitAPI.AsSplit_WOP)

		backRepoData.AsSplitAPIs = append(backRepoData.AsSplitAPIs, &assplitAPI)
	}

	for _, assplitareaDB := range backRepo.BackRepoAsSplitArea.Map_AsSplitAreaDBID_AsSplitAreaDB {

		var assplitareaAPI AsSplitAreaAPI
		assplitareaAPI.ID = assplitareaDB.ID
		assplitareaAPI.AsSplitAreaPointersEncoding = assplitareaDB.AsSplitAreaPointersEncoding
		assplitareaDB.CopyBasicFieldsToAsSplitArea_WOP(&assplitareaAPI.AsSplitArea_WOP)

		backRepoData.AsSplitAreaAPIs = append(backRepoData.AsSplitAreaAPIs, &assplitareaAPI)
	}

	for _, buttonDB := range backRepo.BackRepoButton.Map_ButtonDBID_ButtonDB {

		var buttonAPI ButtonAPI
		buttonAPI.ID = buttonDB.ID
		buttonAPI.ButtonPointersEncoding = buttonDB.ButtonPointersEncoding
		buttonDB.CopyBasicFieldsToButton_WOP(&buttonAPI.Button_WOP)

		backRepoData.ButtonAPIs = append(backRepoData.ButtonAPIs, &buttonAPI)
	}

	for _, faviconDB := range backRepo.BackRepoFavIcon.Map_FavIconDBID_FavIconDB {

		var faviconAPI FavIconAPI
		faviconAPI.ID = faviconDB.ID
		faviconAPI.FavIconPointersEncoding = faviconDB.FavIconPointersEncoding
		faviconDB.CopyBasicFieldsToFavIcon_WOP(&faviconAPI.FavIcon_WOP)

		backRepoData.FavIconAPIs = append(backRepoData.FavIconAPIs, &faviconAPI)
	}

	for _, formDB := range backRepo.BackRepoForm.Map_FormDBID_FormDB {

		var formAPI FormAPI
		formAPI.ID = formDB.ID
		formAPI.FormPointersEncoding = formDB.FormPointersEncoding
		formDB.CopyBasicFieldsToForm_WOP(&formAPI.Form_WOP)

		backRepoData.FormAPIs = append(backRepoData.FormAPIs, &formAPI)
	}

	for _, loadDB := range backRepo.BackRepoLoad.Map_LoadDBID_LoadDB {

		var loadAPI LoadAPI
		loadAPI.ID = loadDB.ID
		loadAPI.LoadPointersEncoding = loadDB.LoadPointersEncoding
		loadDB.CopyBasicFieldsToLoad_WOP(&loadAPI.Load_WOP)

		backRepoData.LoadAPIs = append(backRepoData.LoadAPIs, &loadAPI)
	}

	for _, logoontheleftDB := range backRepo.BackRepoLogoOnTheLeft.Map_LogoOnTheLeftDBID_LogoOnTheLeftDB {

		var logoontheleftAPI LogoOnTheLeftAPI
		logoontheleftAPI.ID = logoontheleftDB.ID
		logoontheleftAPI.LogoOnTheLeftPointersEncoding = logoontheleftDB.LogoOnTheLeftPointersEncoding
		logoontheleftDB.CopyBasicFieldsToLogoOnTheLeft_WOP(&logoontheleftAPI.LogoOnTheLeft_WOP)

		backRepoData.LogoOnTheLeftAPIs = append(backRepoData.LogoOnTheLeftAPIs, &logoontheleftAPI)
	}

	for _, logoontherightDB := range backRepo.BackRepoLogoOnTheRight.Map_LogoOnTheRightDBID_LogoOnTheRightDB {

		var logoontherightAPI LogoOnTheRightAPI
		logoontherightAPI.ID = logoontherightDB.ID
		logoontherightAPI.LogoOnTheRightPointersEncoding = logoontherightDB.LogoOnTheRightPointersEncoding
		logoontherightDB.CopyBasicFieldsToLogoOnTheRight_WOP(&logoontherightAPI.LogoOnTheRight_WOP)

		backRepoData.LogoOnTheRightAPIs = append(backRepoData.LogoOnTheRightAPIs, &logoontherightAPI)
	}

	for _, splitDB := range backRepo.BackRepoSplit.Map_SplitDBID_SplitDB {

		var splitAPI SplitAPI
		splitAPI.ID = splitDB.ID
		splitAPI.SplitPointersEncoding = splitDB.SplitPointersEncoding
		splitDB.CopyBasicFieldsToSplit_WOP(&splitAPI.Split_WOP)

		backRepoData.SplitAPIs = append(backRepoData.SplitAPIs, &splitAPI)
	}

	for _, svgDB := range backRepo.BackRepoSvg.Map_SvgDBID_SvgDB {

		var svgAPI SvgAPI
		svgAPI.ID = svgDB.ID
		svgAPI.SvgPointersEncoding = svgDB.SvgPointersEncoding
		svgDB.CopyBasicFieldsToSvg_WOP(&svgAPI.Svg_WOP)

		backRepoData.SvgAPIs = append(backRepoData.SvgAPIs, &svgAPI)
	}

	for _, tableDB := range backRepo.BackRepoTable.Map_TableDBID_TableDB {

		var tableAPI TableAPI
		tableAPI.ID = tableDB.ID
		tableAPI.TablePointersEncoding = tableDB.TablePointersEncoding
		tableDB.CopyBasicFieldsToTable_WOP(&tableAPI.Table_WOP)

		backRepoData.TableAPIs = append(backRepoData.TableAPIs, &tableAPI)
	}

	for _, titleDB := range backRepo.BackRepoTitle.Map_TitleDBID_TitleDB {

		var titleAPI TitleAPI
		titleAPI.ID = titleDB.ID
		titleAPI.TitlePointersEncoding = titleDB.TitlePointersEncoding
		titleDB.CopyBasicFieldsToTitle_WOP(&titleAPI.Title_WOP)

		backRepoData.TitleAPIs = append(backRepoData.TitleAPIs, &titleAPI)
	}

	for _, treeDB := range backRepo.BackRepoTree.Map_TreeDBID_TreeDB {

		var treeAPI TreeAPI
		treeAPI.ID = treeDB.ID
		treeAPI.TreePointersEncoding = treeDB.TreePointersEncoding
		treeDB.CopyBasicFieldsToTree_WOP(&treeAPI.Tree_WOP)

		backRepoData.TreeAPIs = append(backRepoData.TreeAPIs, &treeAPI)
	}

	for _, viewDB := range backRepo.BackRepoView.Map_ViewDBID_ViewDB {

		var viewAPI ViewAPI
		viewAPI.ID = viewDB.ID
		viewAPI.ViewPointersEncoding = viewDB.ViewPointersEncoding
		viewDB.CopyBasicFieldsToView_WOP(&viewAPI.View_WOP)

		backRepoData.ViewAPIs = append(backRepoData.ViewAPIs, &viewAPI)
	}

}
