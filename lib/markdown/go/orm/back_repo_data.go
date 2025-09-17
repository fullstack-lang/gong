// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	ContentAPIs []*ContentAPI

	JpgImageAPIs []*JpgImageAPI

	PngImageAPIs []*PngImageAPI

	SvgImageAPIs []*SvgImageAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, contentDB := range backRepo.BackRepoContent.Map_ContentDBID_ContentDB {

		var contentAPI ContentAPI
		contentAPI.ID = contentDB.ID
		contentAPI.ContentPointersEncoding = contentDB.ContentPointersEncoding
		contentDB.CopyBasicFieldsToContent_WOP(&contentAPI.Content_WOP)

		backRepoData.ContentAPIs = append(backRepoData.ContentAPIs, &contentAPI)
	}

	for _, jpgimageDB := range backRepo.BackRepoJpgImage.Map_JpgImageDBID_JpgImageDB {

		var jpgimageAPI JpgImageAPI
		jpgimageAPI.ID = jpgimageDB.ID
		jpgimageAPI.JpgImagePointersEncoding = jpgimageDB.JpgImagePointersEncoding
		jpgimageDB.CopyBasicFieldsToJpgImage_WOP(&jpgimageAPI.JpgImage_WOP)

		backRepoData.JpgImageAPIs = append(backRepoData.JpgImageAPIs, &jpgimageAPI)
	}

	for _, pngimageDB := range backRepo.BackRepoPngImage.Map_PngImageDBID_PngImageDB {

		var pngimageAPI PngImageAPI
		pngimageAPI.ID = pngimageDB.ID
		pngimageAPI.PngImagePointersEncoding = pngimageDB.PngImagePointersEncoding
		pngimageDB.CopyBasicFieldsToPngImage_WOP(&pngimageAPI.PngImage_WOP)

		backRepoData.PngImageAPIs = append(backRepoData.PngImageAPIs, &pngimageAPI)
	}

	for _, svgimageDB := range backRepo.BackRepoSvgImage.Map_SvgImageDBID_SvgImageDB {

		var svgimageAPI SvgImageAPI
		svgimageAPI.ID = svgimageDB.ID
		svgimageAPI.SvgImagePointersEncoding = svgimageDB.SvgImagePointersEncoding
		svgimageDB.CopyBasicFieldsToSvgImage_WOP(&svgimageAPI.SvgImage_WOP)

		backRepoData.SvgImageAPIs = append(backRepoData.SvgImageAPIs, &svgimageAPI)
	}

}
