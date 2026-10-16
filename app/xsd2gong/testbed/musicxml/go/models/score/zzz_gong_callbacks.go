// generated code - do not edit
package score

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
func (credit *Credit) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCreditCreateCallback != nil {
		stage.OnAfterCreditCreateCallback.OnAfterCreate(stage, credit)
	}
}

func (credit *Credit) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCreditUpdateCallback != nil {
		var frontCredit *Credit
		if front != nil {
			frontCredit, _ = front.(*Credit)
		}
		stage.OnAfterCreditUpdateCallback.OnAfterUpdate(stage, credit, frontCredit)
	}
}

func (credit *Credit) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCreditDeleteCallback != nil {
		var frontCredit *Credit
		if front != nil {
			frontCredit, _ = front.(*Credit)
		}
		stage.OnAfterCreditDeleteCallback.OnAfterDelete(stage, credit, frontCredit)
	}
}

func (defaults *Defaults) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDefaultsCreateCallback != nil {
		stage.OnAfterDefaultsCreateCallback.OnAfterCreate(stage, defaults)
	}
}

func (defaults *Defaults) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDefaultsUpdateCallback != nil {
		var frontDefaults *Defaults
		if front != nil {
			frontDefaults, _ = front.(*Defaults)
		}
		stage.OnAfterDefaultsUpdateCallback.OnAfterUpdate(stage, defaults, frontDefaults)
	}
}

func (defaults *Defaults) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDefaultsDeleteCallback != nil {
		var frontDefaults *Defaults
		if front != nil {
			frontDefaults, _ = front.(*Defaults)
		}
		stage.OnAfterDefaultsDeleteCallback.OnAfterDelete(stage, defaults, frontDefaults)
	}
}

func (emptyfont *EmptyFont) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmptyFontCreateCallback != nil {
		stage.OnAfterEmptyFontCreateCallback.OnAfterCreate(stage, emptyfont)
	}
}

func (emptyfont *EmptyFont) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyFontUpdateCallback != nil {
		var frontEmptyFont *EmptyFont
		if front != nil {
			frontEmptyFont, _ = front.(*EmptyFont)
		}
		stage.OnAfterEmptyFontUpdateCallback.OnAfterUpdate(stage, emptyfont, frontEmptyFont)
	}
}

func (emptyfont *EmptyFont) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyFontDeleteCallback != nil {
		var frontEmptyFont *EmptyFont
		if front != nil {
			frontEmptyFont, _ = front.(*EmptyFont)
		}
		stage.OnAfterEmptyFontDeleteCallback.OnAfterDelete(stage, emptyfont, frontEmptyFont)
	}
}

func (groupbarline *GroupBarline) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroupBarlineCreateCallback != nil {
		stage.OnAfterGroupBarlineCreateCallback.OnAfterCreate(stage, groupbarline)
	}
}

func (groupbarline *GroupBarline) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupBarlineUpdateCallback != nil {
		var frontGroupBarline *GroupBarline
		if front != nil {
			frontGroupBarline, _ = front.(*GroupBarline)
		}
		stage.OnAfterGroupBarlineUpdateCallback.OnAfterUpdate(stage, groupbarline, frontGroupBarline)
	}
}

func (groupbarline *GroupBarline) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupBarlineDeleteCallback != nil {
		var frontGroupBarline *GroupBarline
		if front != nil {
			frontGroupBarline, _ = front.(*GroupBarline)
		}
		stage.OnAfterGroupBarlineDeleteCallback.OnAfterDelete(stage, groupbarline, frontGroupBarline)
	}
}

func (groupname *GroupName) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroupNameCreateCallback != nil {
		stage.OnAfterGroupNameCreateCallback.OnAfterCreate(stage, groupname)
	}
}

func (groupname *GroupName) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupNameUpdateCallback != nil {
		var frontGroupName *GroupName
		if front != nil {
			frontGroupName, _ = front.(*GroupName)
		}
		stage.OnAfterGroupNameUpdateCallback.OnAfterUpdate(stage, groupname, frontGroupName)
	}
}

func (groupname *GroupName) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupNameDeleteCallback != nil {
		var frontGroupName *GroupName
		if front != nil {
			frontGroupName, _ = front.(*GroupName)
		}
		stage.OnAfterGroupNameDeleteCallback.OnAfterDelete(stage, groupname, frontGroupName)
	}
}

func (groupsymbol *GroupSymbol) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroupSymbolCreateCallback != nil {
		stage.OnAfterGroupSymbolCreateCallback.OnAfterCreate(stage, groupsymbol)
	}
}

func (groupsymbol *GroupSymbol) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupSymbolUpdateCallback != nil {
		var frontGroupSymbol *GroupSymbol
		if front != nil {
			frontGroupSymbol, _ = front.(*GroupSymbol)
		}
		stage.OnAfterGroupSymbolUpdateCallback.OnAfterUpdate(stage, groupsymbol, frontGroupSymbol)
	}
}

func (groupsymbol *GroupSymbol) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupSymbolDeleteCallback != nil {
		var frontGroupSymbol *GroupSymbol
		if front != nil {
			frontGroupSymbol, _ = front.(*GroupSymbol)
		}
		stage.OnAfterGroupSymbolDeleteCallback.OnAfterDelete(stage, groupsymbol, frontGroupSymbol)
	}
}

func (instrumentlink *InstrumentLink) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInstrumentLinkCreateCallback != nil {
		stage.OnAfterInstrumentLinkCreateCallback.OnAfterCreate(stage, instrumentlink)
	}
}

func (instrumentlink *InstrumentLink) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInstrumentLinkUpdateCallback != nil {
		var frontInstrumentLink *InstrumentLink
		if front != nil {
			frontInstrumentLink, _ = front.(*InstrumentLink)
		}
		stage.OnAfterInstrumentLinkUpdateCallback.OnAfterUpdate(stage, instrumentlink, frontInstrumentLink)
	}
}

func (instrumentlink *InstrumentLink) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInstrumentLinkDeleteCallback != nil {
		var frontInstrumentLink *InstrumentLink
		if front != nil {
			frontInstrumentLink, _ = front.(*InstrumentLink)
		}
		stage.OnAfterInstrumentLinkDeleteCallback.OnAfterDelete(stage, instrumentlink, frontInstrumentLink)
	}
}

func (lyricfont *LyricFont) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLyricFontCreateCallback != nil {
		stage.OnAfterLyricFontCreateCallback.OnAfterCreate(stage, lyricfont)
	}
}

func (lyricfont *LyricFont) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLyricFontUpdateCallback != nil {
		var frontLyricFont *LyricFont
		if front != nil {
			frontLyricFont, _ = front.(*LyricFont)
		}
		stage.OnAfterLyricFontUpdateCallback.OnAfterUpdate(stage, lyricfont, frontLyricFont)
	}
}

func (lyricfont *LyricFont) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLyricFontDeleteCallback != nil {
		var frontLyricFont *LyricFont
		if front != nil {
			frontLyricFont, _ = front.(*LyricFont)
		}
		stage.OnAfterLyricFontDeleteCallback.OnAfterDelete(stage, lyricfont, frontLyricFont)
	}
}

func (lyriclanguage *LyricLanguage) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLyricLanguageCreateCallback != nil {
		stage.OnAfterLyricLanguageCreateCallback.OnAfterCreate(stage, lyriclanguage)
	}
}

func (lyriclanguage *LyricLanguage) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLyricLanguageUpdateCallback != nil {
		var frontLyricLanguage *LyricLanguage
		if front != nil {
			frontLyricLanguage, _ = front.(*LyricLanguage)
		}
		stage.OnAfterLyricLanguageUpdateCallback.OnAfterUpdate(stage, lyriclanguage, frontLyricLanguage)
	}
}

func (lyriclanguage *LyricLanguage) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLyricLanguageDeleteCallback != nil {
		var frontLyricLanguage *LyricLanguage
		if front != nil {
			frontLyricLanguage, _ = front.(*LyricLanguage)
		}
		stage.OnAfterLyricLanguageDeleteCallback.OnAfterDelete(stage, lyriclanguage, frontLyricLanguage)
	}
}

func (measure *Measure) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeasureCreateCallback != nil {
		stage.OnAfterMeasureCreateCallback.OnAfterCreate(stage, measure)
	}
}

func (measure *Measure) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureUpdateCallback != nil {
		var frontMeasure *Measure
		if front != nil {
			frontMeasure, _ = front.(*Measure)
		}
		stage.OnAfterMeasureUpdateCallback.OnAfterUpdate(stage, measure, frontMeasure)
	}
}

func (measure *Measure) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureDeleteCallback != nil {
		var frontMeasure *Measure
		if front != nil {
			frontMeasure, _ = front.(*Measure)
		}
		stage.OnAfterMeasureDeleteCallback.OnAfterDelete(stage, measure, frontMeasure)
	}
}

func (measuremusicdatachoice *MeasureMusicDataChoice) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeasureMusicDataChoiceCreateCallback != nil {
		stage.OnAfterMeasureMusicDataChoiceCreateCallback.OnAfterCreate(stage, measuremusicdatachoice)
	}
}

func (measuremusicdatachoice *MeasureMusicDataChoice) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureMusicDataChoiceUpdateCallback != nil {
		var frontMeasureMusicDataChoice *MeasureMusicDataChoice
		if front != nil {
			frontMeasureMusicDataChoice, _ = front.(*MeasureMusicDataChoice)
		}
		stage.OnAfterMeasureMusicDataChoiceUpdateCallback.OnAfterUpdate(stage, measuremusicdatachoice, frontMeasureMusicDataChoice)
	}
}

func (measuremusicdatachoice *MeasureMusicDataChoice) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureMusicDataChoiceDeleteCallback != nil {
		var frontMeasureMusicDataChoice *MeasureMusicDataChoice
		if front != nil {
			frontMeasureMusicDataChoice, _ = front.(*MeasureMusicDataChoice)
		}
		stage.OnAfterMeasureMusicDataChoiceDeleteCallback.OnAfterDelete(stage, measuremusicdatachoice, frontMeasureMusicDataChoice)
	}
}

func (opus *Opus) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOpusCreateCallback != nil {
		stage.OnAfterOpusCreateCallback.OnAfterCreate(stage, opus)
	}
}

func (opus *Opus) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOpusUpdateCallback != nil {
		var frontOpus *Opus
		if front != nil {
			frontOpus, _ = front.(*Opus)
		}
		stage.OnAfterOpusUpdateCallback.OnAfterUpdate(stage, opus, frontOpus)
	}
}

func (opus *Opus) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOpusDeleteCallback != nil {
		var frontOpus *Opus
		if front != nil {
			frontOpus, _ = front.(*Opus)
		}
		stage.OnAfterOpusDeleteCallback.OnAfterDelete(stage, opus, frontOpus)
	}
}

func (part *Part) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartCreateCallback != nil {
		stage.OnAfterPartCreateCallback.OnAfterCreate(stage, part)
	}
}

func (part *Part) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartUpdateCallback != nil {
		var frontPart *Part
		if front != nil {
			frontPart, _ = front.(*Part)
		}
		stage.OnAfterPartUpdateCallback.OnAfterUpdate(stage, part, frontPart)
	}
}

func (part *Part) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartDeleteCallback != nil {
		var frontPart *Part
		if front != nil {
			frontPart, _ = front.(*Part)
		}
		stage.OnAfterPartDeleteCallback.OnAfterDelete(stage, part, frontPart)
	}
}

func (partgroup *PartGroup) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartGroupCreateCallback != nil {
		stage.OnAfterPartGroupCreateCallback.OnAfterCreate(stage, partgroup)
	}
}

func (partgroup *PartGroup) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartGroupUpdateCallback != nil {
		var frontPartGroup *PartGroup
		if front != nil {
			frontPartGroup, _ = front.(*PartGroup)
		}
		stage.OnAfterPartGroupUpdateCallback.OnAfterUpdate(stage, partgroup, frontPartGroup)
	}
}

func (partgroup *PartGroup) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartGroupDeleteCallback != nil {
		var frontPartGroup *PartGroup
		if front != nil {
			frontPartGroup, _ = front.(*PartGroup)
		}
		stage.OnAfterPartGroupDeleteCallback.OnAfterDelete(stage, partgroup, frontPartGroup)
	}
}

func (partlink *PartLink) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartLinkCreateCallback != nil {
		stage.OnAfterPartLinkCreateCallback.OnAfterCreate(stage, partlink)
	}
}

func (partlink *PartLink) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartLinkUpdateCallback != nil {
		var frontPartLink *PartLink
		if front != nil {
			frontPartLink, _ = front.(*PartLink)
		}
		stage.OnAfterPartLinkUpdateCallback.OnAfterUpdate(stage, partlink, frontPartLink)
	}
}

func (partlink *PartLink) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartLinkDeleteCallback != nil {
		var frontPartLink *PartLink
		if front != nil {
			frontPartLink, _ = front.(*PartLink)
		}
		stage.OnAfterPartLinkDeleteCallback.OnAfterDelete(stage, partlink, frontPartLink)
	}
}

func (partlist *PartList) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartListCreateCallback != nil {
		stage.OnAfterPartListCreateCallback.OnAfterCreate(stage, partlist)
	}
}

func (partlist *PartList) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartListUpdateCallback != nil {
		var frontPartList *PartList
		if front != nil {
			frontPartList, _ = front.(*PartList)
		}
		stage.OnAfterPartListUpdateCallback.OnAfterUpdate(stage, partlist, frontPartList)
	}
}

func (partlist *PartList) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartListDeleteCallback != nil {
		var frontPartList *PartList
		if front != nil {
			frontPartList, _ = front.(*PartList)
		}
		stage.OnAfterPartListDeleteCallback.OnAfterDelete(stage, partlist, frontPartList)
	}
}

func (partname *PartName) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartNameCreateCallback != nil {
		stage.OnAfterPartNameCreateCallback.OnAfterCreate(stage, partname)
	}
}

func (partname *PartName) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartNameUpdateCallback != nil {
		var frontPartName *PartName
		if front != nil {
			frontPartName, _ = front.(*PartName)
		}
		stage.OnAfterPartNameUpdateCallback.OnAfterUpdate(stage, partname, frontPartName)
	}
}

func (partname *PartName) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartNameDeleteCallback != nil {
		var frontPartName *PartName
		if front != nil {
			frontPartName, _ = front.(*PartName)
		}
		stage.OnAfterPartNameDeleteCallback.OnAfterDelete(stage, partname, frontPartName)
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

func (scoreinstrument *ScoreInstrument) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterScoreInstrumentCreateCallback != nil {
		stage.OnAfterScoreInstrumentCreateCallback.OnAfterCreate(stage, scoreinstrument)
	}
}

func (scoreinstrument *ScoreInstrument) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScoreInstrumentUpdateCallback != nil {
		var frontScoreInstrument *ScoreInstrument
		if front != nil {
			frontScoreInstrument, _ = front.(*ScoreInstrument)
		}
		stage.OnAfterScoreInstrumentUpdateCallback.OnAfterUpdate(stage, scoreinstrument, frontScoreInstrument)
	}
}

func (scoreinstrument *ScoreInstrument) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScoreInstrumentDeleteCallback != nil {
		var frontScoreInstrument *ScoreInstrument
		if front != nil {
			frontScoreInstrument, _ = front.(*ScoreInstrument)
		}
		stage.OnAfterScoreInstrumentDeleteCallback.OnAfterDelete(stage, scoreinstrument, frontScoreInstrument)
	}
}

func (scorepart *ScorePart) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterScorePartCreateCallback != nil {
		stage.OnAfterScorePartCreateCallback.OnAfterCreate(stage, scorepart)
	}
}

func (scorepart *ScorePart) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScorePartUpdateCallback != nil {
		var frontScorePart *ScorePart
		if front != nil {
			frontScorePart, _ = front.(*ScorePart)
		}
		stage.OnAfterScorePartUpdateCallback.OnAfterUpdate(stage, scorepart, frontScorePart)
	}
}

func (scorepart *ScorePart) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScorePartDeleteCallback != nil {
		var frontScorePart *ScorePart
		if front != nil {
			frontScorePart, _ = front.(*ScorePart)
		}
		stage.OnAfterScorePartDeleteCallback.OnAfterDelete(stage, scorepart, frontScorePart)
	}
}

func (scorepartwise *ScorePartwise) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterScorePartwiseCreateCallback != nil {
		stage.OnAfterScorePartwiseCreateCallback.OnAfterCreate(stage, scorepartwise)
	}
}

func (scorepartwise *ScorePartwise) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScorePartwiseUpdateCallback != nil {
		var frontScorePartwise *ScorePartwise
		if front != nil {
			frontScorePartwise, _ = front.(*ScorePartwise)
		}
		stage.OnAfterScorePartwiseUpdateCallback.OnAfterUpdate(stage, scorepartwise, frontScorePartwise)
	}
}

func (scorepartwise *ScorePartwise) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScorePartwiseDeleteCallback != nil {
		var frontScorePartwise *ScorePartwise
		if front != nil {
			frontScorePartwise, _ = front.(*ScorePartwise)
		}
		stage.OnAfterScorePartwiseDeleteCallback.OnAfterDelete(stage, scorepartwise, frontScorePartwise)
	}
}

func (scoretimewise *ScoreTimewise) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterScoreTimewiseCreateCallback != nil {
		stage.OnAfterScoreTimewiseCreateCallback.OnAfterCreate(stage, scoretimewise)
	}
}

func (scoretimewise *ScoreTimewise) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScoreTimewiseUpdateCallback != nil {
		var frontScoreTimewise *ScoreTimewise
		if front != nil {
			frontScoreTimewise, _ = front.(*ScoreTimewise)
		}
		stage.OnAfterScoreTimewiseUpdateCallback.OnAfterUpdate(stage, scoretimewise, frontScoreTimewise)
	}
}

func (scoretimewise *ScoreTimewise) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScoreTimewiseDeleteCallback != nil {
		var frontScoreTimewise *ScoreTimewise
		if front != nil {
			frontScoreTimewise, _ = front.(*ScoreTimewise)
		}
		stage.OnAfterScoreTimewiseDeleteCallback.OnAfterDelete(stage, scoretimewise, frontScoreTimewise)
	}
}

func (scoretimewisemeasure *ScoreTimewiseMeasure) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterScoreTimewiseMeasureCreateCallback != nil {
		stage.OnAfterScoreTimewiseMeasureCreateCallback.OnAfterCreate(stage, scoretimewisemeasure)
	}
}

func (scoretimewisemeasure *ScoreTimewiseMeasure) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScoreTimewiseMeasureUpdateCallback != nil {
		var frontScoreTimewiseMeasure *ScoreTimewiseMeasure
		if front != nil {
			frontScoreTimewiseMeasure, _ = front.(*ScoreTimewiseMeasure)
		}
		stage.OnAfterScoreTimewiseMeasureUpdateCallback.OnAfterUpdate(stage, scoretimewisemeasure, frontScoreTimewiseMeasure)
	}
}

func (scoretimewisemeasure *ScoreTimewiseMeasure) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScoreTimewiseMeasureDeleteCallback != nil {
		var frontScoreTimewiseMeasure *ScoreTimewiseMeasure
		if front != nil {
			frontScoreTimewiseMeasure, _ = front.(*ScoreTimewiseMeasure)
		}
		stage.OnAfterScoreTimewiseMeasureDeleteCallback.OnAfterDelete(stage, scoretimewisemeasure, frontScoreTimewiseMeasure)
	}
}

func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterScoreTimewiseMeasurePartCreateCallback != nil {
		stage.OnAfterScoreTimewiseMeasurePartCreateCallback.OnAfterCreate(stage, scoretimewisemeasurepart)
	}
}

func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScoreTimewiseMeasurePartUpdateCallback != nil {
		var frontScoreTimewiseMeasurePart *ScoreTimewiseMeasurePart
		if front != nil {
			frontScoreTimewiseMeasurePart, _ = front.(*ScoreTimewiseMeasurePart)
		}
		stage.OnAfterScoreTimewiseMeasurePartUpdateCallback.OnAfterUpdate(stage, scoretimewisemeasurepart, frontScoreTimewiseMeasurePart)
	}
}

func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScoreTimewiseMeasurePartDeleteCallback != nil {
		var frontScoreTimewiseMeasurePart *ScoreTimewiseMeasurePart
		if front != nil {
			frontScoreTimewiseMeasurePart, _ = front.(*ScoreTimewiseMeasurePart)
		}
		stage.OnAfterScoreTimewiseMeasurePartDeleteCallback.OnAfterDelete(stage, scoretimewisemeasurepart, frontScoreTimewiseMeasurePart)
	}
}

func (work *Work) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterWorkCreateCallback != nil {
		stage.OnAfterWorkCreateCallback.OnAfterCreate(stage, work)
	}
}

func (work *Work) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWorkUpdateCallback != nil {
		var frontWork *Work
		if front != nil {
			frontWork, _ = front.(*Work)
		}
		stage.OnAfterWorkUpdateCallback.OnAfterUpdate(stage, work, frontWork)
	}
}

func (work *Work) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWorkDeleteCallback != nil {
		var frontWork *Work
		if front != nil {
			frontWork, _ = front.(*Work)
		}
		stage.OnAfterWorkDeleteCallback.OnAfterDelete(stage, work, frontWork)
	}
}

