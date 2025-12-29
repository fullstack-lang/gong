// generated code - do not edit
package models

import (
	"strings"
	"time"
)

var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

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

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	// insertion point per named struct
	var freqencys_newInstances []*Freqency
	var freqencys_deletedInstances []*Freqency

	// parse all staged instances and check if they have a reference
	for freqency := range stage.Freqencys {
		if ref, ok := stage.Freqencys_reference[freqency]; !ok {
			freqencys_newInstances = append(freqencys_newInstances, freqency)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Freqency "+freqency.Name,
				)
			}
		} else {
			diffs := freqency.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Freqency \""+freqency.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for freqency := range stage.Freqencys_reference {
		if _, ok := stage.Freqencys[freqency]; !ok {
			freqencys_deletedInstances = append(freqencys_deletedInstances, freqency)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Freqency "+freqency.Name,
				)
			}
		}
	}

	lenNewInstances += len(freqencys_newInstances)
	lenDeletedInstances += len(freqencys_deletedInstances)
	var notes_newInstances []*Note
	var notes_deletedInstances []*Note

	// parse all staged instances and check if they have a reference
	for note := range stage.Notes {
		if ref, ok := stage.Notes_reference[note]; !ok {
			notes_newInstances = append(notes_newInstances, note)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Note "+note.Name,
				)
			}
		} else {
			diffs := note.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Note \""+note.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for note := range stage.Notes_reference {
		if _, ok := stage.Notes[note]; !ok {
			notes_deletedInstances = append(notes_deletedInstances, note)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Note "+note.Name,
				)
			}
		}
	}

	lenNewInstances += len(notes_newInstances)
	lenDeletedInstances += len(notes_deletedInstances)
	var players_newInstances []*Player
	var players_deletedInstances []*Player

	// parse all staged instances and check if they have a reference
	for player := range stage.Players {
		if ref, ok := stage.Players_reference[player]; !ok {
			players_newInstances = append(players_newInstances, player)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Player "+player.Name,
				)
			}
		} else {
			diffs := player.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Player \""+player.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for player := range stage.Players_reference {
		if _, ok := stage.Players[player]; !ok {
			players_deletedInstances = append(players_deletedInstances, player)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Player "+player.Name,
				)
			}
		}
	}

	lenNewInstances += len(players_newInstances)
	lenDeletedInstances += len(players_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		// if stage.GetProbeIF() != nil {
		// 	stage.GetProbeIF().CommitNotificationTable()
		// }
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Freqencys_reference = make(map[*Freqency]*Freqency)
	for instance := range stage.Freqencys {
		stage.Freqencys_reference[instance] = instance.GongCopy().(*Freqency)
	}

	stage.Notes_reference = make(map[*Note]*Note)
	for instance := range stage.Notes {
		stage.Notes_reference[instance] = instance.GongCopy().(*Note)
	}

	stage.Players_reference = make(map[*Player]*Player)
	for instance := range stage.Players {
		stage.Players_reference[instance] = instance.GongCopy().(*Player)
	}

}
