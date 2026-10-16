// generated code - do not edit
package common

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
func (accidentaltext *AccidentalText) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAccidentalTextCreateCallback != nil {
		stage.OnAfterAccidentalTextCreateCallback.OnAfterCreate(stage, accidentaltext)
	}
}

func (accidentaltext *AccidentalText) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAccidentalTextUpdateCallback != nil {
		var frontAccidentalText *AccidentalText
		if front != nil {
			frontAccidentalText, _ = front.(*AccidentalText)
		}
		stage.OnAfterAccidentalTextUpdateCallback.OnAfterUpdate(stage, accidentaltext, frontAccidentalText)
	}
}

func (accidentaltext *AccidentalText) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAccidentalTextDeleteCallback != nil {
		var frontAccidentalText *AccidentalText
		if front != nil {
			frontAccidentalText, _ = front.(*AccidentalText)
		}
		stage.OnAfterAccidentalTextDeleteCallback.OnAfterDelete(stage, accidentaltext, frontAccidentalText)
	}
}

func (appearance *Appearance) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAppearanceCreateCallback != nil {
		stage.OnAfterAppearanceCreateCallback.OnAfterCreate(stage, appearance)
	}
}

func (appearance *Appearance) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAppearanceUpdateCallback != nil {
		var frontAppearance *Appearance
		if front != nil {
			frontAppearance, _ = front.(*Appearance)
		}
		stage.OnAfterAppearanceUpdateCallback.OnAfterUpdate(stage, appearance, frontAppearance)
	}
}

func (appearance *Appearance) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAppearanceDeleteCallback != nil {
		var frontAppearance *Appearance
		if front != nil {
			frontAppearance, _ = front.(*Appearance)
		}
		stage.OnAfterAppearanceDeleteCallback.OnAfterDelete(stage, appearance, frontAppearance)
	}
}

func (bookmark *Bookmark) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBookmarkCreateCallback != nil {
		stage.OnAfterBookmarkCreateCallback.OnAfterCreate(stage, bookmark)
	}
}

func (bookmark *Bookmark) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBookmarkUpdateCallback != nil {
		var frontBookmark *Bookmark
		if front != nil {
			frontBookmark, _ = front.(*Bookmark)
		}
		stage.OnAfterBookmarkUpdateCallback.OnAfterUpdate(stage, bookmark, frontBookmark)
	}
}

func (bookmark *Bookmark) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBookmarkDeleteCallback != nil {
		var frontBookmark *Bookmark
		if front != nil {
			frontBookmark, _ = front.(*Bookmark)
		}
		stage.OnAfterBookmarkDeleteCallback.OnAfterDelete(stage, bookmark, frontBookmark)
	}
}

func (coda *Coda) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodaCreateCallback != nil {
		stage.OnAfterCodaCreateCallback.OnAfterCreate(stage, coda)
	}
}

func (coda *Coda) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodaUpdateCallback != nil {
		var frontCoda *Coda
		if front != nil {
			frontCoda, _ = front.(*Coda)
		}
		stage.OnAfterCodaUpdateCallback.OnAfterUpdate(stage, coda, frontCoda)
	}
}

func (coda *Coda) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodaDeleteCallback != nil {
		var frontCoda *Coda
		if front != nil {
			frontCoda, _ = front.(*Coda)
		}
		stage.OnAfterCodaDeleteCallback.OnAfterDelete(stage, coda, frontCoda)
	}
}

func (distance *Distance) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDistanceCreateCallback != nil {
		stage.OnAfterDistanceCreateCallback.OnAfterCreate(stage, distance)
	}
}

func (distance *Distance) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDistanceUpdateCallback != nil {
		var frontDistance *Distance
		if front != nil {
			frontDistance, _ = front.(*Distance)
		}
		stage.OnAfterDistanceUpdateCallback.OnAfterUpdate(stage, distance, frontDistance)
	}
}

func (distance *Distance) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDistanceDeleteCallback != nil {
		var frontDistance *Distance
		if front != nil {
			frontDistance, _ = front.(*Distance)
		}
		stage.OnAfterDistanceDeleteCallback.OnAfterDelete(stage, distance, frontDistance)
	}
}

func (dynamics *Dynamics) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDynamicsCreateCallback != nil {
		stage.OnAfterDynamicsCreateCallback.OnAfterCreate(stage, dynamics)
	}
}

func (dynamics *Dynamics) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDynamicsUpdateCallback != nil {
		var frontDynamics *Dynamics
		if front != nil {
			frontDynamics, _ = front.(*Dynamics)
		}
		stage.OnAfterDynamicsUpdateCallback.OnAfterUpdate(stage, dynamics, frontDynamics)
	}
}

func (dynamics *Dynamics) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDynamicsDeleteCallback != nil {
		var frontDynamics *Dynamics
		if front != nil {
			frontDynamics, _ = front.(*Dynamics)
		}
		stage.OnAfterDynamicsDeleteCallback.OnAfterDelete(stage, dynamics, frontDynamics)
	}
}

func (empty *Empty) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmptyCreateCallback != nil {
		stage.OnAfterEmptyCreateCallback.OnAfterCreate(stage, empty)
	}
}

func (empty *Empty) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyUpdateCallback != nil {
		var frontEmpty *Empty
		if front != nil {
			frontEmpty, _ = front.(*Empty)
		}
		stage.OnAfterEmptyUpdateCallback.OnAfterUpdate(stage, empty, frontEmpty)
	}
}

func (empty *Empty) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyDeleteCallback != nil {
		var frontEmpty *Empty
		if front != nil {
			frontEmpty, _ = front.(*Empty)
		}
		stage.OnAfterEmptyDeleteCallback.OnAfterDelete(stage, empty, frontEmpty)
	}
}

func (emptyplacement *EmptyPlacement) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmptyPlacementCreateCallback != nil {
		stage.OnAfterEmptyPlacementCreateCallback.OnAfterCreate(stage, emptyplacement)
	}
}

func (emptyplacement *EmptyPlacement) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyPlacementUpdateCallback != nil {
		var frontEmptyPlacement *EmptyPlacement
		if front != nil {
			frontEmptyPlacement, _ = front.(*EmptyPlacement)
		}
		stage.OnAfterEmptyPlacementUpdateCallback.OnAfterUpdate(stage, emptyplacement, frontEmptyPlacement)
	}
}

func (emptyplacement *EmptyPlacement) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyPlacementDeleteCallback != nil {
		var frontEmptyPlacement *EmptyPlacement
		if front != nil {
			frontEmptyPlacement, _ = front.(*EmptyPlacement)
		}
		stage.OnAfterEmptyPlacementDeleteCallback.OnAfterDelete(stage, emptyplacement, frontEmptyPlacement)
	}
}

func (emptyplacementsmufl *EmptyPlacementSmufl) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmptyPlacementSmuflCreateCallback != nil {
		stage.OnAfterEmptyPlacementSmuflCreateCallback.OnAfterCreate(stage, emptyplacementsmufl)
	}
}

func (emptyplacementsmufl *EmptyPlacementSmufl) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyPlacementSmuflUpdateCallback != nil {
		var frontEmptyPlacementSmufl *EmptyPlacementSmufl
		if front != nil {
			frontEmptyPlacementSmufl, _ = front.(*EmptyPlacementSmufl)
		}
		stage.OnAfterEmptyPlacementSmuflUpdateCallback.OnAfterUpdate(stage, emptyplacementsmufl, frontEmptyPlacementSmufl)
	}
}

func (emptyplacementsmufl *EmptyPlacementSmufl) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyPlacementSmuflDeleteCallback != nil {
		var frontEmptyPlacementSmufl *EmptyPlacementSmufl
		if front != nil {
			frontEmptyPlacementSmufl, _ = front.(*EmptyPlacementSmufl)
		}
		stage.OnAfterEmptyPlacementSmuflDeleteCallback.OnAfterDelete(stage, emptyplacementsmufl, frontEmptyPlacementSmufl)
	}
}

func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmptyPrintObjectStyleAlignCreateCallback != nil {
		stage.OnAfterEmptyPrintObjectStyleAlignCreateCallback.OnAfterCreate(stage, emptyprintobjectstylealign)
	}
}

func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyPrintObjectStyleAlignUpdateCallback != nil {
		var frontEmptyPrintObjectStyleAlign *EmptyPrintObjectStyleAlign
		if front != nil {
			frontEmptyPrintObjectStyleAlign, _ = front.(*EmptyPrintObjectStyleAlign)
		}
		stage.OnAfterEmptyPrintObjectStyleAlignUpdateCallback.OnAfterUpdate(stage, emptyprintobjectstylealign, frontEmptyPrintObjectStyleAlign)
	}
}

func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyPrintObjectStyleAlignDeleteCallback != nil {
		var frontEmptyPrintObjectStyleAlign *EmptyPrintObjectStyleAlign
		if front != nil {
			frontEmptyPrintObjectStyleAlign, _ = front.(*EmptyPrintObjectStyleAlign)
		}
		stage.OnAfterEmptyPrintObjectStyleAlignDeleteCallback.OnAfterDelete(stage, emptyprintobjectstylealign, frontEmptyPrintObjectStyleAlign)
	}
}

func (emptyprintstyle *EmptyPrintStyle) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmptyPrintStyleCreateCallback != nil {
		stage.OnAfterEmptyPrintStyleCreateCallback.OnAfterCreate(stage, emptyprintstyle)
	}
}

func (emptyprintstyle *EmptyPrintStyle) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyPrintStyleUpdateCallback != nil {
		var frontEmptyPrintStyle *EmptyPrintStyle
		if front != nil {
			frontEmptyPrintStyle, _ = front.(*EmptyPrintStyle)
		}
		stage.OnAfterEmptyPrintStyleUpdateCallback.OnAfterUpdate(stage, emptyprintstyle, frontEmptyPrintStyle)
	}
}

func (emptyprintstyle *EmptyPrintStyle) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyPrintStyleDeleteCallback != nil {
		var frontEmptyPrintStyle *EmptyPrintStyle
		if front != nil {
			frontEmptyPrintStyle, _ = front.(*EmptyPrintStyle)
		}
		stage.OnAfterEmptyPrintStyleDeleteCallback.OnAfterDelete(stage, emptyprintstyle, frontEmptyPrintStyle)
	}
}

func (emptyprintstylealign *EmptyPrintStyleAlign) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmptyPrintStyleAlignCreateCallback != nil {
		stage.OnAfterEmptyPrintStyleAlignCreateCallback.OnAfterCreate(stage, emptyprintstylealign)
	}
}

func (emptyprintstylealign *EmptyPrintStyleAlign) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyPrintStyleAlignUpdateCallback != nil {
		var frontEmptyPrintStyleAlign *EmptyPrintStyleAlign
		if front != nil {
			frontEmptyPrintStyleAlign, _ = front.(*EmptyPrintStyleAlign)
		}
		stage.OnAfterEmptyPrintStyleAlignUpdateCallback.OnAfterUpdate(stage, emptyprintstylealign, frontEmptyPrintStyleAlign)
	}
}

func (emptyprintstylealign *EmptyPrintStyleAlign) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyPrintStyleAlignDeleteCallback != nil {
		var frontEmptyPrintStyleAlign *EmptyPrintStyleAlign
		if front != nil {
			frontEmptyPrintStyleAlign, _ = front.(*EmptyPrintStyleAlign)
		}
		stage.OnAfterEmptyPrintStyleAlignDeleteCallback.OnAfterDelete(stage, emptyprintstylealign, frontEmptyPrintStyleAlign)
	}
}

func (emptyprintstylealignid *EmptyPrintStyleAlignId) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmptyPrintStyleAlignIdCreateCallback != nil {
		stage.OnAfterEmptyPrintStyleAlignIdCreateCallback.OnAfterCreate(stage, emptyprintstylealignid)
	}
}

func (emptyprintstylealignid *EmptyPrintStyleAlignId) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyPrintStyleAlignIdUpdateCallback != nil {
		var frontEmptyPrintStyleAlignId *EmptyPrintStyleAlignId
		if front != nil {
			frontEmptyPrintStyleAlignId, _ = front.(*EmptyPrintStyleAlignId)
		}
		stage.OnAfterEmptyPrintStyleAlignIdUpdateCallback.OnAfterUpdate(stage, emptyprintstylealignid, frontEmptyPrintStyleAlignId)
	}
}

func (emptyprintstylealignid *EmptyPrintStyleAlignId) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyPrintStyleAlignIdDeleteCallback != nil {
		var frontEmptyPrintStyleAlignId *EmptyPrintStyleAlignId
		if front != nil {
			frontEmptyPrintStyleAlignId, _ = front.(*EmptyPrintStyleAlignId)
		}
		stage.OnAfterEmptyPrintStyleAlignIdDeleteCallback.OnAfterDelete(stage, emptyprintstylealignid, frontEmptyPrintStyleAlignId)
	}
}

func (emptytrillsound *EmptyTrillSound) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmptyTrillSoundCreateCallback != nil {
		stage.OnAfterEmptyTrillSoundCreateCallback.OnAfterCreate(stage, emptytrillsound)
	}
}

func (emptytrillsound *EmptyTrillSound) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyTrillSoundUpdateCallback != nil {
		var frontEmptyTrillSound *EmptyTrillSound
		if front != nil {
			frontEmptyTrillSound, _ = front.(*EmptyTrillSound)
		}
		stage.OnAfterEmptyTrillSoundUpdateCallback.OnAfterUpdate(stage, emptytrillsound, frontEmptyTrillSound)
	}
}

func (emptytrillsound *EmptyTrillSound) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyTrillSoundDeleteCallback != nil {
		var frontEmptyTrillSound *EmptyTrillSound
		if front != nil {
			frontEmptyTrillSound, _ = front.(*EmptyTrillSound)
		}
		stage.OnAfterEmptyTrillSoundDeleteCallback.OnAfterDelete(stage, emptytrillsound, frontEmptyTrillSound)
	}
}

func (encoding *Encoding) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEncodingCreateCallback != nil {
		stage.OnAfterEncodingCreateCallback.OnAfterCreate(stage, encoding)
	}
}

func (encoding *Encoding) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEncodingUpdateCallback != nil {
		var frontEncoding *Encoding
		if front != nil {
			frontEncoding, _ = front.(*Encoding)
		}
		stage.OnAfterEncodingUpdateCallback.OnAfterUpdate(stage, encoding, frontEncoding)
	}
}

func (encoding *Encoding) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEncodingDeleteCallback != nil {
		var frontEncoding *Encoding
		if front != nil {
			frontEncoding, _ = front.(*Encoding)
		}
		stage.OnAfterEncodingDeleteCallback.OnAfterDelete(stage, encoding, frontEncoding)
	}
}

func (fermata *Fermata) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFermataCreateCallback != nil {
		stage.OnAfterFermataCreateCallback.OnAfterCreate(stage, fermata)
	}
}

func (fermata *Fermata) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFermataUpdateCallback != nil {
		var frontFermata *Fermata
		if front != nil {
			frontFermata, _ = front.(*Fermata)
		}
		stage.OnAfterFermataUpdateCallback.OnAfterUpdate(stage, fermata, frontFermata)
	}
}

func (fermata *Fermata) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFermataDeleteCallback != nil {
		var frontFermata *Fermata
		if front != nil {
			frontFermata, _ = front.(*Fermata)
		}
		stage.OnAfterFermataDeleteCallback.OnAfterDelete(stage, fermata, frontFermata)
	}
}

func (fingering *Fingering) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFingeringCreateCallback != nil {
		stage.OnAfterFingeringCreateCallback.OnAfterCreate(stage, fingering)
	}
}

func (fingering *Fingering) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFingeringUpdateCallback != nil {
		var frontFingering *Fingering
		if front != nil {
			frontFingering, _ = front.(*Fingering)
		}
		stage.OnAfterFingeringUpdateCallback.OnAfterUpdate(stage, fingering, frontFingering)
	}
}

func (fingering *Fingering) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFingeringDeleteCallback != nil {
		var frontFingering *Fingering
		if front != nil {
			frontFingering, _ = front.(*Fingering)
		}
		stage.OnAfterFingeringDeleteCallback.OnAfterDelete(stage, fingering, frontFingering)
	}
}

func (formattedsymbol *FormattedSymbol) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormattedSymbolCreateCallback != nil {
		stage.OnAfterFormattedSymbolCreateCallback.OnAfterCreate(stage, formattedsymbol)
	}
}

func (formattedsymbol *FormattedSymbol) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormattedSymbolUpdateCallback != nil {
		var frontFormattedSymbol *FormattedSymbol
		if front != nil {
			frontFormattedSymbol, _ = front.(*FormattedSymbol)
		}
		stage.OnAfterFormattedSymbolUpdateCallback.OnAfterUpdate(stage, formattedsymbol, frontFormattedSymbol)
	}
}

func (formattedsymbol *FormattedSymbol) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormattedSymbolDeleteCallback != nil {
		var frontFormattedSymbol *FormattedSymbol
		if front != nil {
			frontFormattedSymbol, _ = front.(*FormattedSymbol)
		}
		stage.OnAfterFormattedSymbolDeleteCallback.OnAfterDelete(stage, formattedsymbol, frontFormattedSymbol)
	}
}

func (formattedsymbolid *FormattedSymbolId) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormattedSymbolIdCreateCallback != nil {
		stage.OnAfterFormattedSymbolIdCreateCallback.OnAfterCreate(stage, formattedsymbolid)
	}
}

func (formattedsymbolid *FormattedSymbolId) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormattedSymbolIdUpdateCallback != nil {
		var frontFormattedSymbolId *FormattedSymbolId
		if front != nil {
			frontFormattedSymbolId, _ = front.(*FormattedSymbolId)
		}
		stage.OnAfterFormattedSymbolIdUpdateCallback.OnAfterUpdate(stage, formattedsymbolid, frontFormattedSymbolId)
	}
}

func (formattedsymbolid *FormattedSymbolId) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormattedSymbolIdDeleteCallback != nil {
		var frontFormattedSymbolId *FormattedSymbolId
		if front != nil {
			frontFormattedSymbolId, _ = front.(*FormattedSymbolId)
		}
		stage.OnAfterFormattedSymbolIdDeleteCallback.OnAfterDelete(stage, formattedsymbolid, frontFormattedSymbolId)
	}
}

func (formattedtext *FormattedText) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormattedTextCreateCallback != nil {
		stage.OnAfterFormattedTextCreateCallback.OnAfterCreate(stage, formattedtext)
	}
}

func (formattedtext *FormattedText) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormattedTextUpdateCallback != nil {
		var frontFormattedText *FormattedText
		if front != nil {
			frontFormattedText, _ = front.(*FormattedText)
		}
		stage.OnAfterFormattedTextUpdateCallback.OnAfterUpdate(stage, formattedtext, frontFormattedText)
	}
}

func (formattedtext *FormattedText) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormattedTextDeleteCallback != nil {
		var frontFormattedText *FormattedText
		if front != nil {
			frontFormattedText, _ = front.(*FormattedText)
		}
		stage.OnAfterFormattedTextDeleteCallback.OnAfterDelete(stage, formattedtext, frontFormattedText)
	}
}

func (formattedtextid *FormattedTextId) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormattedTextIdCreateCallback != nil {
		stage.OnAfterFormattedTextIdCreateCallback.OnAfterCreate(stage, formattedtextid)
	}
}

func (formattedtextid *FormattedTextId) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormattedTextIdUpdateCallback != nil {
		var frontFormattedTextId *FormattedTextId
		if front != nil {
			frontFormattedTextId, _ = front.(*FormattedTextId)
		}
		stage.OnAfterFormattedTextIdUpdateCallback.OnAfterUpdate(stage, formattedtextid, frontFormattedTextId)
	}
}

func (formattedtextid *FormattedTextId) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormattedTextIdDeleteCallback != nil {
		var frontFormattedTextId *FormattedTextId
		if front != nil {
			frontFormattedTextId, _ = front.(*FormattedTextId)
		}
		stage.OnAfterFormattedTextIdDeleteCallback.OnAfterDelete(stage, formattedtextid, frontFormattedTextId)
	}
}

func (fret *Fret) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFretCreateCallback != nil {
		stage.OnAfterFretCreateCallback.OnAfterCreate(stage, fret)
	}
}

func (fret *Fret) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFretUpdateCallback != nil {
		var frontFret *Fret
		if front != nil {
			frontFret, _ = front.(*Fret)
		}
		stage.OnAfterFretUpdateCallback.OnAfterUpdate(stage, fret, frontFret)
	}
}

func (fret *Fret) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFretDeleteCallback != nil {
		var frontFret *Fret
		if front != nil {
			frontFret, _ = front.(*Fret)
		}
		stage.OnAfterFretDeleteCallback.OnAfterDelete(stage, fret, frontFret)
	}
}

func (glyph *Glyph) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGlyphCreateCallback != nil {
		stage.OnAfterGlyphCreateCallback.OnAfterCreate(stage, glyph)
	}
}

func (glyph *Glyph) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGlyphUpdateCallback != nil {
		var frontGlyph *Glyph
		if front != nil {
			frontGlyph, _ = front.(*Glyph)
		}
		stage.OnAfterGlyphUpdateCallback.OnAfterUpdate(stage, glyph, frontGlyph)
	}
}

func (glyph *Glyph) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGlyphDeleteCallback != nil {
		var frontGlyph *Glyph
		if front != nil {
			frontGlyph, _ = front.(*Glyph)
		}
		stage.OnAfterGlyphDeleteCallback.OnAfterDelete(stage, glyph, frontGlyph)
	}
}

func (horizontalturn *HorizontalTurn) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHorizontalTurnCreateCallback != nil {
		stage.OnAfterHorizontalTurnCreateCallback.OnAfterCreate(stage, horizontalturn)
	}
}

func (horizontalturn *HorizontalTurn) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHorizontalTurnUpdateCallback != nil {
		var frontHorizontalTurn *HorizontalTurn
		if front != nil {
			frontHorizontalTurn, _ = front.(*HorizontalTurn)
		}
		stage.OnAfterHorizontalTurnUpdateCallback.OnAfterUpdate(stage, horizontalturn, frontHorizontalTurn)
	}
}

func (horizontalturn *HorizontalTurn) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHorizontalTurnDeleteCallback != nil {
		var frontHorizontalTurn *HorizontalTurn
		if front != nil {
			frontHorizontalTurn, _ = front.(*HorizontalTurn)
		}
		stage.OnAfterHorizontalTurnDeleteCallback.OnAfterDelete(stage, horizontalturn, frontHorizontalTurn)
	}
}

func (identification *Identification) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterIdentificationCreateCallback != nil {
		stage.OnAfterIdentificationCreateCallback.OnAfterCreate(stage, identification)
	}
}

func (identification *Identification) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterIdentificationUpdateCallback != nil {
		var frontIdentification *Identification
		if front != nil {
			frontIdentification, _ = front.(*Identification)
		}
		stage.OnAfterIdentificationUpdateCallback.OnAfterUpdate(stage, identification, frontIdentification)
	}
}

func (identification *Identification) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterIdentificationDeleteCallback != nil {
		var frontIdentification *Identification
		if front != nil {
			frontIdentification, _ = front.(*Identification)
		}
		stage.OnAfterIdentificationDeleteCallback.OnAfterDelete(stage, identification, frontIdentification)
	}
}

func (level *Level) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLevelCreateCallback != nil {
		stage.OnAfterLevelCreateCallback.OnAfterCreate(stage, level)
	}
}

func (level *Level) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLevelUpdateCallback != nil {
		var frontLevel *Level
		if front != nil {
			frontLevel, _ = front.(*Level)
		}
		stage.OnAfterLevelUpdateCallback.OnAfterUpdate(stage, level, frontLevel)
	}
}

func (level *Level) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLevelDeleteCallback != nil {
		var frontLevel *Level
		if front != nil {
			frontLevel, _ = front.(*Level)
		}
		stage.OnAfterLevelDeleteCallback.OnAfterDelete(stage, level, frontLevel)
	}
}

func (linewidth *LineWidth) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLineWidthCreateCallback != nil {
		stage.OnAfterLineWidthCreateCallback.OnAfterCreate(stage, linewidth)
	}
}

func (linewidth *LineWidth) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLineWidthUpdateCallback != nil {
		var frontLineWidth *LineWidth
		if front != nil {
			frontLineWidth, _ = front.(*LineWidth)
		}
		stage.OnAfterLineWidthUpdateCallback.OnAfterUpdate(stage, linewidth, frontLineWidth)
	}
}

func (linewidth *LineWidth) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLineWidthDeleteCallback != nil {
		var frontLineWidth *LineWidth
		if front != nil {
			frontLineWidth, _ = front.(*LineWidth)
		}
		stage.OnAfterLineWidthDeleteCallback.OnAfterDelete(stage, linewidth, frontLineWidth)
	}
}

func (link *Link) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLinkCreateCallback != nil {
		stage.OnAfterLinkCreateCallback.OnAfterCreate(stage, link)
	}
}

func (link *Link) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinkUpdateCallback != nil {
		var frontLink *Link
		if front != nil {
			frontLink, _ = front.(*Link)
		}
		stage.OnAfterLinkUpdateCallback.OnAfterUpdate(stage, link, frontLink)
	}
}

func (link *Link) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinkDeleteCallback != nil {
		var frontLink *Link
		if front != nil {
			frontLink, _ = front.(*Link)
		}
		stage.OnAfterLinkDeleteCallback.OnAfterDelete(stage, link, frontLink)
	}
}

func (measurelayout *MeasureLayout) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeasureLayoutCreateCallback != nil {
		stage.OnAfterMeasureLayoutCreateCallback.OnAfterCreate(stage, measurelayout)
	}
}

func (measurelayout *MeasureLayout) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureLayoutUpdateCallback != nil {
		var frontMeasureLayout *MeasureLayout
		if front != nil {
			frontMeasureLayout, _ = front.(*MeasureLayout)
		}
		stage.OnAfterMeasureLayoutUpdateCallback.OnAfterUpdate(stage, measurelayout, frontMeasureLayout)
	}
}

func (measurelayout *MeasureLayout) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureLayoutDeleteCallback != nil {
		var frontMeasureLayout *MeasureLayout
		if front != nil {
			frontMeasureLayout, _ = front.(*MeasureLayout)
		}
		stage.OnAfterMeasureLayoutDeleteCallback.OnAfterDelete(stage, measurelayout, frontMeasureLayout)
	}
}

func (mididevice *MidiDevice) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMidiDeviceCreateCallback != nil {
		stage.OnAfterMidiDeviceCreateCallback.OnAfterCreate(stage, mididevice)
	}
}

func (mididevice *MidiDevice) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMidiDeviceUpdateCallback != nil {
		var frontMidiDevice *MidiDevice
		if front != nil {
			frontMidiDevice, _ = front.(*MidiDevice)
		}
		stage.OnAfterMidiDeviceUpdateCallback.OnAfterUpdate(stage, mididevice, frontMidiDevice)
	}
}

func (mididevice *MidiDevice) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMidiDeviceDeleteCallback != nil {
		var frontMidiDevice *MidiDevice
		if front != nil {
			frontMidiDevice, _ = front.(*MidiDevice)
		}
		stage.OnAfterMidiDeviceDeleteCallback.OnAfterDelete(stage, mididevice, frontMidiDevice)
	}
}

func (midiinstrument *MidiInstrument) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMidiInstrumentCreateCallback != nil {
		stage.OnAfterMidiInstrumentCreateCallback.OnAfterCreate(stage, midiinstrument)
	}
}

func (midiinstrument *MidiInstrument) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMidiInstrumentUpdateCallback != nil {
		var frontMidiInstrument *MidiInstrument
		if front != nil {
			frontMidiInstrument, _ = front.(*MidiInstrument)
		}
		stage.OnAfterMidiInstrumentUpdateCallback.OnAfterUpdate(stage, midiinstrument, frontMidiInstrument)
	}
}

func (midiinstrument *MidiInstrument) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMidiInstrumentDeleteCallback != nil {
		var frontMidiInstrument *MidiInstrument
		if front != nil {
			frontMidiInstrument, _ = front.(*MidiInstrument)
		}
		stage.OnAfterMidiInstrumentDeleteCallback.OnAfterDelete(stage, midiinstrument, frontMidiInstrument)
	}
}

func (miscellaneous *Miscellaneous) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMiscellaneousCreateCallback != nil {
		stage.OnAfterMiscellaneousCreateCallback.OnAfterCreate(stage, miscellaneous)
	}
}

func (miscellaneous *Miscellaneous) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMiscellaneousUpdateCallback != nil {
		var frontMiscellaneous *Miscellaneous
		if front != nil {
			frontMiscellaneous, _ = front.(*Miscellaneous)
		}
		stage.OnAfterMiscellaneousUpdateCallback.OnAfterUpdate(stage, miscellaneous, frontMiscellaneous)
	}
}

func (miscellaneous *Miscellaneous) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMiscellaneousDeleteCallback != nil {
		var frontMiscellaneous *Miscellaneous
		if front != nil {
			frontMiscellaneous, _ = front.(*Miscellaneous)
		}
		stage.OnAfterMiscellaneousDeleteCallback.OnAfterDelete(stage, miscellaneous, frontMiscellaneous)
	}
}

func (miscellaneousfield *MiscellaneousField) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMiscellaneousFieldCreateCallback != nil {
		stage.OnAfterMiscellaneousFieldCreateCallback.OnAfterCreate(stage, miscellaneousfield)
	}
}

func (miscellaneousfield *MiscellaneousField) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMiscellaneousFieldUpdateCallback != nil {
		var frontMiscellaneousField *MiscellaneousField
		if front != nil {
			frontMiscellaneousField, _ = front.(*MiscellaneousField)
		}
		stage.OnAfterMiscellaneousFieldUpdateCallback.OnAfterUpdate(stage, miscellaneousfield, frontMiscellaneousField)
	}
}

func (miscellaneousfield *MiscellaneousField) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMiscellaneousFieldDeleteCallback != nil {
		var frontMiscellaneousField *MiscellaneousField
		if front != nil {
			frontMiscellaneousField, _ = front.(*MiscellaneousField)
		}
		stage.OnAfterMiscellaneousFieldDeleteCallback.OnAfterDelete(stage, miscellaneousfield, frontMiscellaneousField)
	}
}

func (namedisplay *NameDisplay) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNameDisplayCreateCallback != nil {
		stage.OnAfterNameDisplayCreateCallback.OnAfterCreate(stage, namedisplay)
	}
}

func (namedisplay *NameDisplay) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNameDisplayUpdateCallback != nil {
		var frontNameDisplay *NameDisplay
		if front != nil {
			frontNameDisplay, _ = front.(*NameDisplay)
		}
		stage.OnAfterNameDisplayUpdateCallback.OnAfterUpdate(stage, namedisplay, frontNameDisplay)
	}
}

func (namedisplay *NameDisplay) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNameDisplayDeleteCallback != nil {
		var frontNameDisplay *NameDisplay
		if front != nil {
			frontNameDisplay, _ = front.(*NameDisplay)
		}
		stage.OnAfterNameDisplayDeleteCallback.OnAfterDelete(stage, namedisplay, frontNameDisplay)
	}
}

func (notesize *NoteSize) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteSizeCreateCallback != nil {
		stage.OnAfterNoteSizeCreateCallback.OnAfterCreate(stage, notesize)
	}
}

func (notesize *NoteSize) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteSizeUpdateCallback != nil {
		var frontNoteSize *NoteSize
		if front != nil {
			frontNoteSize, _ = front.(*NoteSize)
		}
		stage.OnAfterNoteSizeUpdateCallback.OnAfterUpdate(stage, notesize, frontNoteSize)
	}
}

func (notesize *NoteSize) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteSizeDeleteCallback != nil {
		var frontNoteSize *NoteSize
		if front != nil {
			frontNoteSize, _ = front.(*NoteSize)
		}
		stage.OnAfterNoteSizeDeleteCallback.OnAfterDelete(stage, notesize, frontNoteSize)
	}
}

func (otherappearance *OtherAppearance) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOtherAppearanceCreateCallback != nil {
		stage.OnAfterOtherAppearanceCreateCallback.OnAfterCreate(stage, otherappearance)
	}
}

func (otherappearance *OtherAppearance) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOtherAppearanceUpdateCallback != nil {
		var frontOtherAppearance *OtherAppearance
		if front != nil {
			frontOtherAppearance, _ = front.(*OtherAppearance)
		}
		stage.OnAfterOtherAppearanceUpdateCallback.OnAfterUpdate(stage, otherappearance, frontOtherAppearance)
	}
}

func (otherappearance *OtherAppearance) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOtherAppearanceDeleteCallback != nil {
		var frontOtherAppearance *OtherAppearance
		if front != nil {
			frontOtherAppearance, _ = front.(*OtherAppearance)
		}
		stage.OnAfterOtherAppearanceDeleteCallback.OnAfterDelete(stage, otherappearance, frontOtherAppearance)
	}
}

func (otherplay *OtherPlay) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOtherPlayCreateCallback != nil {
		stage.OnAfterOtherPlayCreateCallback.OnAfterCreate(stage, otherplay)
	}
}

func (otherplay *OtherPlay) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOtherPlayUpdateCallback != nil {
		var frontOtherPlay *OtherPlay
		if front != nil {
			frontOtherPlay, _ = front.(*OtherPlay)
		}
		stage.OnAfterOtherPlayUpdateCallback.OnAfterUpdate(stage, otherplay, frontOtherPlay)
	}
}

func (otherplay *OtherPlay) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOtherPlayDeleteCallback != nil {
		var frontOtherPlay *OtherPlay
		if front != nil {
			frontOtherPlay, _ = front.(*OtherPlay)
		}
		stage.OnAfterOtherPlayDeleteCallback.OnAfterDelete(stage, otherplay, frontOtherPlay)
	}
}

func (othertext *OtherText) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOtherTextCreateCallback != nil {
		stage.OnAfterOtherTextCreateCallback.OnAfterCreate(stage, othertext)
	}
}

func (othertext *OtherText) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOtherTextUpdateCallback != nil {
		var frontOtherText *OtherText
		if front != nil {
			frontOtherText, _ = front.(*OtherText)
		}
		stage.OnAfterOtherTextUpdateCallback.OnAfterUpdate(stage, othertext, frontOtherText)
	}
}

func (othertext *OtherText) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOtherTextDeleteCallback != nil {
		var frontOtherText *OtherText
		if front != nil {
			frontOtherText, _ = front.(*OtherText)
		}
		stage.OnAfterOtherTextDeleteCallback.OnAfterDelete(stage, othertext, frontOtherText)
	}
}

func (pagelayout *PageLayout) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPageLayoutCreateCallback != nil {
		stage.OnAfterPageLayoutCreateCallback.OnAfterCreate(stage, pagelayout)
	}
}

func (pagelayout *PageLayout) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPageLayoutUpdateCallback != nil {
		var frontPageLayout *PageLayout
		if front != nil {
			frontPageLayout, _ = front.(*PageLayout)
		}
		stage.OnAfterPageLayoutUpdateCallback.OnAfterUpdate(stage, pagelayout, frontPageLayout)
	}
}

func (pagelayout *PageLayout) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPageLayoutDeleteCallback != nil {
		var frontPageLayout *PageLayout
		if front != nil {
			frontPageLayout, _ = front.(*PageLayout)
		}
		stage.OnAfterPageLayoutDeleteCallback.OnAfterDelete(stage, pagelayout, frontPageLayout)
	}
}

func (pagemargins *PageMargins) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPageMarginsCreateCallback != nil {
		stage.OnAfterPageMarginsCreateCallback.OnAfterCreate(stage, pagemargins)
	}
}

func (pagemargins *PageMargins) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPageMarginsUpdateCallback != nil {
		var frontPageMargins *PageMargins
		if front != nil {
			frontPageMargins, _ = front.(*PageMargins)
		}
		stage.OnAfterPageMarginsUpdateCallback.OnAfterUpdate(stage, pagemargins, frontPageMargins)
	}
}

func (pagemargins *PageMargins) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPageMarginsDeleteCallback != nil {
		var frontPageMargins *PageMargins
		if front != nil {
			frontPageMargins, _ = front.(*PageMargins)
		}
		stage.OnAfterPageMarginsDeleteCallback.OnAfterDelete(stage, pagemargins, frontPageMargins)
	}
}

func (play *Play) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPlayCreateCallback != nil {
		stage.OnAfterPlayCreateCallback.OnAfterCreate(stage, play)
	}
}

func (play *Play) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlayUpdateCallback != nil {
		var frontPlay *Play
		if front != nil {
			frontPlay, _ = front.(*Play)
		}
		stage.OnAfterPlayUpdateCallback.OnAfterUpdate(stage, play, frontPlay)
	}
}

func (play *Play) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlayDeleteCallback != nil {
		var frontPlay *Play
		if front != nil {
			frontPlay, _ = front.(*Play)
		}
		stage.OnAfterPlayDeleteCallback.OnAfterDelete(stage, play, frontPlay)
	}
}

func (scaling *Scaling) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterScalingCreateCallback != nil {
		stage.OnAfterScalingCreateCallback.OnAfterCreate(stage, scaling)
	}
}

func (scaling *Scaling) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScalingUpdateCallback != nil {
		var frontScaling *Scaling
		if front != nil {
			frontScaling, _ = front.(*Scaling)
		}
		stage.OnAfterScalingUpdateCallback.OnAfterUpdate(stage, scaling, frontScaling)
	}
}

func (scaling *Scaling) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScalingDeleteCallback != nil {
		var frontScaling *Scaling
		if front != nil {
			frontScaling, _ = front.(*Scaling)
		}
		stage.OnAfterScalingDeleteCallback.OnAfterDelete(stage, scaling, frontScaling)
	}
}

func (segno *Segno) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSegnoCreateCallback != nil {
		stage.OnAfterSegnoCreateCallback.OnAfterCreate(stage, segno)
	}
}

func (segno *Segno) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSegnoUpdateCallback != nil {
		var frontSegno *Segno
		if front != nil {
			frontSegno, _ = front.(*Segno)
		}
		stage.OnAfterSegnoUpdateCallback.OnAfterUpdate(stage, segno, frontSegno)
	}
}

func (segno *Segno) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSegnoDeleteCallback != nil {
		var frontSegno *Segno
		if front != nil {
			frontSegno, _ = front.(*Segno)
		}
		stage.OnAfterSegnoDeleteCallback.OnAfterDelete(stage, segno, frontSegno)
	}
}

func (stafflayout *StaffLayout) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStaffLayoutCreateCallback != nil {
		stage.OnAfterStaffLayoutCreateCallback.OnAfterCreate(stage, stafflayout)
	}
}

func (stafflayout *StaffLayout) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaffLayoutUpdateCallback != nil {
		var frontStaffLayout *StaffLayout
		if front != nil {
			frontStaffLayout, _ = front.(*StaffLayout)
		}
		stage.OnAfterStaffLayoutUpdateCallback.OnAfterUpdate(stage, stafflayout, frontStaffLayout)
	}
}

func (stafflayout *StaffLayout) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaffLayoutDeleteCallback != nil {
		var frontStaffLayout *StaffLayout
		if front != nil {
			frontStaffLayout, _ = front.(*StaffLayout)
		}
		stage.OnAfterStaffLayoutDeleteCallback.OnAfterDelete(stage, stafflayout, frontStaffLayout)
	}
}

func (stringtype *StringType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStringTypeCreateCallback != nil {
		stage.OnAfterStringTypeCreateCallback.OnAfterCreate(stage, stringtype)
	}
}

func (stringtype *StringType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStringTypeUpdateCallback != nil {
		var frontStringType *StringType
		if front != nil {
			frontStringType, _ = front.(*StringType)
		}
		stage.OnAfterStringTypeUpdateCallback.OnAfterUpdate(stage, stringtype, frontStringType)
	}
}

func (stringtype *StringType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStringTypeDeleteCallback != nil {
		var frontStringType *StringType
		if front != nil {
			frontStringType, _ = front.(*StringType)
		}
		stage.OnAfterStringTypeDeleteCallback.OnAfterDelete(stage, stringtype, frontStringType)
	}
}

func (supports *Supports) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSupportsCreateCallback != nil {
		stage.OnAfterSupportsCreateCallback.OnAfterCreate(stage, supports)
	}
}

func (supports *Supports) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSupportsUpdateCallback != nil {
		var frontSupports *Supports
		if front != nil {
			frontSupports, _ = front.(*Supports)
		}
		stage.OnAfterSupportsUpdateCallback.OnAfterUpdate(stage, supports, frontSupports)
	}
}

func (supports *Supports) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSupportsDeleteCallback != nil {
		var frontSupports *Supports
		if front != nil {
			frontSupports, _ = front.(*Supports)
		}
		stage.OnAfterSupportsDeleteCallback.OnAfterDelete(stage, supports, frontSupports)
	}
}

func (systemdividers *SystemDividers) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSystemDividersCreateCallback != nil {
		stage.OnAfterSystemDividersCreateCallback.OnAfterCreate(stage, systemdividers)
	}
}

func (systemdividers *SystemDividers) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSystemDividersUpdateCallback != nil {
		var frontSystemDividers *SystemDividers
		if front != nil {
			frontSystemDividers, _ = front.(*SystemDividers)
		}
		stage.OnAfterSystemDividersUpdateCallback.OnAfterUpdate(stage, systemdividers, frontSystemDividers)
	}
}

func (systemdividers *SystemDividers) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSystemDividersDeleteCallback != nil {
		var frontSystemDividers *SystemDividers
		if front != nil {
			frontSystemDividers, _ = front.(*SystemDividers)
		}
		stage.OnAfterSystemDividersDeleteCallback.OnAfterDelete(stage, systemdividers, frontSystemDividers)
	}
}

func (systemlayout *SystemLayout) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSystemLayoutCreateCallback != nil {
		stage.OnAfterSystemLayoutCreateCallback.OnAfterCreate(stage, systemlayout)
	}
}

func (systemlayout *SystemLayout) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSystemLayoutUpdateCallback != nil {
		var frontSystemLayout *SystemLayout
		if front != nil {
			frontSystemLayout, _ = front.(*SystemLayout)
		}
		stage.OnAfterSystemLayoutUpdateCallback.OnAfterUpdate(stage, systemlayout, frontSystemLayout)
	}
}

func (systemlayout *SystemLayout) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSystemLayoutDeleteCallback != nil {
		var frontSystemLayout *SystemLayout
		if front != nil {
			frontSystemLayout, _ = front.(*SystemLayout)
		}
		stage.OnAfterSystemLayoutDeleteCallback.OnAfterDelete(stage, systemlayout, frontSystemLayout)
	}
}

func (systemmargins *SystemMargins) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSystemMarginsCreateCallback != nil {
		stage.OnAfterSystemMarginsCreateCallback.OnAfterCreate(stage, systemmargins)
	}
}

func (systemmargins *SystemMargins) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSystemMarginsUpdateCallback != nil {
		var frontSystemMargins *SystemMargins
		if front != nil {
			frontSystemMargins, _ = front.(*SystemMargins)
		}
		stage.OnAfterSystemMarginsUpdateCallback.OnAfterUpdate(stage, systemmargins, frontSystemMargins)
	}
}

func (systemmargins *SystemMargins) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSystemMarginsDeleteCallback != nil {
		var frontSystemMargins *SystemMargins
		if front != nil {
			frontSystemMargins, _ = front.(*SystemMargins)
		}
		stage.OnAfterSystemMarginsDeleteCallback.OnAfterDelete(stage, systemmargins, frontSystemMargins)
	}
}

func (typedtext *TypedText) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTypedTextCreateCallback != nil {
		stage.OnAfterTypedTextCreateCallback.OnAfterCreate(stage, typedtext)
	}
}

func (typedtext *TypedText) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTypedTextUpdateCallback != nil {
		var frontTypedText *TypedText
		if front != nil {
			frontTypedText, _ = front.(*TypedText)
		}
		stage.OnAfterTypedTextUpdateCallback.OnAfterUpdate(stage, typedtext, frontTypedText)
	}
}

func (typedtext *TypedText) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTypedTextDeleteCallback != nil {
		var frontTypedText *TypedText
		if front != nil {
			frontTypedText, _ = front.(*TypedText)
		}
		stage.OnAfterTypedTextDeleteCallback.OnAfterDelete(stage, typedtext, frontTypedText)
	}
}

func (wavyline *WavyLine) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterWavyLineCreateCallback != nil {
		stage.OnAfterWavyLineCreateCallback.OnAfterCreate(stage, wavyline)
	}
}

func (wavyline *WavyLine) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWavyLineUpdateCallback != nil {
		var frontWavyLine *WavyLine
		if front != nil {
			frontWavyLine, _ = front.(*WavyLine)
		}
		stage.OnAfterWavyLineUpdateCallback.OnAfterUpdate(stage, wavyline, frontWavyLine)
	}
}

func (wavyline *WavyLine) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWavyLineDeleteCallback != nil {
		var frontWavyLine *WavyLine
		if front != nil {
			frontWavyLine, _ = front.(*WavyLine)
		}
		stage.OnAfterWavyLineDeleteCallback.OnAfterDelete(stage, wavyline, frontWavyLine)
	}
}

