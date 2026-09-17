// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (freqency *Freqency) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFreqencyCreateCallback != nil {
		stage.OnAfterFreqencyCreateCallback.OnAfterCreate(stage, freqency)
	}
}

func (freqency *Freqency) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFreqencyUpdateCallback != nil {
		var frontFreqency *Freqency
		if front != nil {
			frontFreqency, _ = front.(*Freqency)
		}
		stage.OnAfterFreqencyUpdateCallback.OnAfterUpdate(stage, freqency, frontFreqency)
	}
}

func (freqency *Freqency) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFreqencyDeleteCallback != nil {
		var frontFreqency *Freqency
		if front != nil {
			frontFreqency, _ = front.(*Freqency)
		}
		stage.OnAfterFreqencyDeleteCallback.OnAfterDelete(stage, freqency, frontFreqency)
	}
}

func (note *Note) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteCreateCallback != nil {
		stage.OnAfterNoteCreateCallback.OnAfterCreate(stage, note)
	}
}

func (note *Note) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteUpdateCallback != nil {
		var frontNote *Note
		if front != nil {
			frontNote, _ = front.(*Note)
		}
		stage.OnAfterNoteUpdateCallback.OnAfterUpdate(stage, note, frontNote)
	}
}

func (note *Note) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteDeleteCallback != nil {
		var frontNote *Note
		if front != nil {
			frontNote, _ = front.(*Note)
		}
		stage.OnAfterNoteDeleteCallback.OnAfterDelete(stage, note, frontNote)
	}
}

func (player *Player) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPlayerCreateCallback != nil {
		stage.OnAfterPlayerCreateCallback.OnAfterCreate(stage, player)
	}
}

func (player *Player) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlayerUpdateCallback != nil {
		var frontPlayer *Player
		if front != nil {
			frontPlayer, _ = front.(*Player)
		}
		stage.OnAfterPlayerUpdateCallback.OnAfterUpdate(stage, player, frontPlayer)
	}
}

func (player *Player) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlayerDeleteCallback != nil {
		var frontPlayer *Player
		if front != nil {
			frontPlayer, _ = front.(*Player)
		}
		stage.OnAfterPlayerDeleteCallback.OnAfterDelete(stage, player, frontPlayer)
	}
}

