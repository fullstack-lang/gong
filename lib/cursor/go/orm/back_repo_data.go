// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	CursorAPIs []*CursorAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, cursorDB := range backRepo.BackRepoCursor.Map_CursorDBID_CursorDB {

		var cursorAPI CursorAPI
		cursorAPI.ID = cursorDB.ID
		cursorAPI.CursorPointersEncoding = cursorDB.CursorPointersEncoding
		cursorDB.CopyBasicFieldsToCursor_WOP(&cursorAPI.Cursor_WOP)

		backRepoData.CursorAPIs = append(backRepoData.CursorAPIs, &cursorAPI)
	}

}
