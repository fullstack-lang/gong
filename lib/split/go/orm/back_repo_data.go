// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	AsSplitAPIs []*AsSplitAPI

	AsSplitAreaAPIs []*AsSplitAreaAPI

	ButtonAPIs []*ButtonAPI

	CursorAPIs []*CursorAPI

	DocAPIs []*DocAPI

	FormAPIs []*FormAPI

	LoadAPIs []*LoadAPI

	SliderAPIs []*SliderAPI

	SplitAPIs []*SplitAPI

	SvgAPIs []*SvgAPI

	TableAPIs []*TableAPI

	ToneAPIs []*ToneAPI

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

	for _, cursorDB := range backRepo.BackRepoCursor.Map_CursorDBID_CursorDB {

		var cursorAPI CursorAPI
		cursorAPI.ID = cursorDB.ID
		cursorAPI.CursorPointersEncoding = cursorDB.CursorPointersEncoding
		cursorDB.CopyBasicFieldsToCursor_WOP(&cursorAPI.Cursor_WOP)

		backRepoData.CursorAPIs = append(backRepoData.CursorAPIs, &cursorAPI)
	}

	for _, docDB := range backRepo.BackRepoDoc.Map_DocDBID_DocDB {

		var docAPI DocAPI
		docAPI.ID = docDB.ID
		docAPI.DocPointersEncoding = docDB.DocPointersEncoding
		docDB.CopyBasicFieldsToDoc_WOP(&docAPI.Doc_WOP)

		backRepoData.DocAPIs = append(backRepoData.DocAPIs, &docAPI)
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

	for _, sliderDB := range backRepo.BackRepoSlider.Map_SliderDBID_SliderDB {

		var sliderAPI SliderAPI
		sliderAPI.ID = sliderDB.ID
		sliderAPI.SliderPointersEncoding = sliderDB.SliderPointersEncoding
		sliderDB.CopyBasicFieldsToSlider_WOP(&sliderAPI.Slider_WOP)

		backRepoData.SliderAPIs = append(backRepoData.SliderAPIs, &sliderAPI)
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

	for _, toneDB := range backRepo.BackRepoTone.Map_ToneDBID_ToneDB {

		var toneAPI ToneAPI
		toneAPI.ID = toneDB.ID
		toneAPI.TonePointersEncoding = toneDB.TonePointersEncoding
		toneDB.CopyBasicFieldsToTone_WOP(&toneAPI.Tone_WOP)

		backRepoData.ToneAPIs = append(backRepoData.ToneAPIs, &toneAPI)
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
