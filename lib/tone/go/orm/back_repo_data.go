// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	FreqencyAPIs []*FreqencyAPI

	NoteAPIs []*NoteAPI

	PlayerAPIs []*PlayerAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, freqencyDB := range backRepo.BackRepoFreqency.Map_FreqencyDBID_FreqencyDB {

		var freqencyAPI FreqencyAPI
		freqencyAPI.ID = freqencyDB.ID
		freqencyAPI.FreqencyPointersEncoding = freqencyDB.FreqencyPointersEncoding
		freqencyDB.CopyBasicFieldsToFreqency_WOP(&freqencyAPI.Freqency_WOP)

		backRepoData.FreqencyAPIs = append(backRepoData.FreqencyAPIs, &freqencyAPI)
	}

	for _, noteDB := range backRepo.BackRepoNote.Map_NoteDBID_NoteDB {

		var noteAPI NoteAPI
		noteAPI.ID = noteDB.ID
		noteAPI.NotePointersEncoding = noteDB.NotePointersEncoding
		noteDB.CopyBasicFieldsToNote_WOP(&noteAPI.Note_WOP)

		backRepoData.NoteAPIs = append(backRepoData.NoteAPIs, &noteAPI)
	}

	for _, playerDB := range backRepo.BackRepoPlayer.Map_PlayerDBID_PlayerDB {

		var playerAPI PlayerAPI
		playerAPI.ID = playerDB.ID
		playerAPI.PlayerPointersEncoding = playerDB.PlayerPointersEncoding
		playerDB.CopyBasicFieldsToPlayer_WOP(&playerAPI.Player_WOP)

		backRepoData.PlayerAPIs = append(backRepoData.PlayerAPIs, &playerAPI)
	}

}
