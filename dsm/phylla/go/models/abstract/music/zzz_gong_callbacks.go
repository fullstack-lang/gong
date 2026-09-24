// generated code - do not edit
package music

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
func (musicabstract *MusicAbstract) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMusicAbstractCreateCallback != nil {
		stage.OnAfterMusicAbstractCreateCallback.OnAfterCreate(stage, musicabstract)
	}
}

func (musicabstract *MusicAbstract) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMusicAbstractUpdateCallback != nil {
		var frontMusicAbstract *MusicAbstract
		if front != nil {
			frontMusicAbstract, _ = front.(*MusicAbstract)
		}
		stage.OnAfterMusicAbstractUpdateCallback.OnAfterUpdate(stage, musicabstract, frontMusicAbstract)
	}
}

func (musicabstract *MusicAbstract) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMusicAbstractDeleteCallback != nil {
		var frontMusicAbstract *MusicAbstract
		if front != nil {
			frontMusicAbstract, _ = front.(*MusicAbstract)
		}
		stage.OnAfterMusicAbstractDeleteCallback.OnAfterDelete(stage, musicabstract, frontMusicAbstract)
	}
}

