// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Freqency
	// insertion point per field

	// Compute reverse map for named struct Note
	// insertion point per field
	clear(stage.Note_Frequencies_reverseMap)
	stage.Note_Frequencies_reverseMap = make(map[*Freqency]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _freqency := range note.Frequencies {
			stage.Note_Frequencies_reverseMap[_freqency] = note
		}
	}

	// Compute reverse map for named struct Player
	// insertion point per field

}
