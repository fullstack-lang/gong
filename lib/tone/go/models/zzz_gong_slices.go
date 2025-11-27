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

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Freqencys {
		res = append(res, instance)
	}

	for instance := range stage.Notes {
		res = append(res, instance)
	}

	for instance := range stage.Players {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (freqency *Freqency) GongCopy() GongstructIF {
	newInstance := *freqency
	return &newInstance
}

func (note *Note) GongCopy() GongstructIF {
	newInstance := *note
	return &newInstance
}

func (player *Player) GongCopy() GongstructIF {
	newInstance := *player
	return &newInstance
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {
	stage.reference = make(map[GongstructIF]GongstructIF)
	for _, instance := range stage.GetInstances() {
		stage.reference[instance] = instance.GongCopy()
	}
	stage.new = make(map[GongstructIF]struct{})
	stage.modified = make(map[GongstructIF]struct{})
	stage.deleted = make(map[GongstructIF]struct{})
}
