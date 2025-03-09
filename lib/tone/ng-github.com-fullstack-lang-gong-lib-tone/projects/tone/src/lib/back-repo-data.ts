// generated code - do not edit

//insertion point for imports
import { FreqencyAPI } from './freqency-api'

import { NoteAPI } from './note-api'

import { PlayerAPI } from './player-api'


export class BackRepoData {
	// insertion point for declarations
	FreqencyAPIs = new Array<FreqencyAPI>()

	NoteAPIs = new Array<NoteAPI>()

	PlayerAPIs = new Array<PlayerAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.FreqencyAPIs = data?.FreqencyAPIs || [];

		this.NoteAPIs = data?.NoteAPIs || [];

		this.PlayerAPIs = data?.PlayerAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}