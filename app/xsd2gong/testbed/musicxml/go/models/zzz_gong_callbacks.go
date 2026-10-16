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
func (musicxmldocument *MusicXMLDocument) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMusicXMLDocumentCreateCallback != nil {
		stage.OnAfterMusicXMLDocumentCreateCallback.OnAfterCreate(stage, musicxmldocument)
	}
}

func (musicxmldocument *MusicXMLDocument) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMusicXMLDocumentUpdateCallback != nil {
		var frontMusicXMLDocument *MusicXMLDocument
		if front != nil {
			frontMusicXMLDocument, _ = front.(*MusicXMLDocument)
		}
		stage.OnAfterMusicXMLDocumentUpdateCallback.OnAfterUpdate(stage, musicxmldocument, frontMusicXMLDocument)
	}
}

func (musicxmldocument *MusicXMLDocument) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMusicXMLDocumentDeleteCallback != nil {
		var frontMusicXMLDocument *MusicXMLDocument
		if front != nil {
			frontMusicXMLDocument, _ = front.(*MusicXMLDocument)
		}
		stage.OnAfterMusicXMLDocumentDeleteCallback.OnAfterDelete(stage, musicxmldocument, frontMusicXMLDocument)
	}
}

