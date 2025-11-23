// generated code - do not edit
package models

// Clean computes the reverse map, for all intances, for all clean to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) Clean() {
	// insertion point per named struct
	// Compute reverse map for named struct Freqency
	for freqency := range stage.Freqencys {
		_ = freqency
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Note
	for note := range stage.Notes {
		_ = note
		// insertion point per field
		var _Frequencies []*Freqency
		for _, _freqency := range note.Frequencies {
			if IsStaged(stage, _freqency) {
			 	_Frequencies = append(_Frequencies, _freqency)
			}
		}
		note.Frequencies = _Frequencies
		// insertion point per field
	}

	// Compute reverse map for named struct Player
	for player := range stage.Players {
		_ = player
		// insertion point per field
		// insertion point per field
	}

}
