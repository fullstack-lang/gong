// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	ChapterAPIs []*ChapterAPI

	ContentAPIs []*ContentAPI

	PageAPIs []*PageAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, chapterDB := range backRepo.BackRepoChapter.Map_ChapterDBID_ChapterDB {

		var chapterAPI ChapterAPI
		chapterAPI.ID = chapterDB.ID
		chapterAPI.ChapterPointersEncoding = chapterDB.ChapterPointersEncoding
		chapterDB.CopyBasicFieldsToChapter_WOP(&chapterAPI.Chapter_WOP)

		backRepoData.ChapterAPIs = append(backRepoData.ChapterAPIs, &chapterAPI)
	}

	for _, contentDB := range backRepo.BackRepoContent.Map_ContentDBID_ContentDB {

		var contentAPI ContentAPI
		contentAPI.ID = contentDB.ID
		contentAPI.ContentPointersEncoding = contentDB.ContentPointersEncoding
		contentDB.CopyBasicFieldsToContent_WOP(&contentAPI.Content_WOP)

		backRepoData.ContentAPIs = append(backRepoData.ContentAPIs, &contentAPI)
	}

	for _, pageDB := range backRepo.BackRepoPage.Map_PageDBID_PageDB {

		var pageAPI PageAPI
		pageAPI.ID = pageDB.ID
		pageAPI.PagePointersEncoding = pageDB.PagePointersEncoding
		pageDB.CopyBasicFieldsToPage_WOP(&pageAPI.Page_WOP)

		backRepoData.PageAPIs = append(backRepoData.PageAPIs, &pageAPI)
	}

}
