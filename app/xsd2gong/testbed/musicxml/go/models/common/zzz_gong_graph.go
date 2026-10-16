// generated code - do not edit
package common

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (accidentaltext *AccidentalText) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AccidentalTexts[accidentaltext]
	return ok
}

func (appearance *Appearance) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Appearances[appearance]
	return ok
}

func (bookmark *Bookmark) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Bookmarks[bookmark]
	return ok
}

func (coda *Coda) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Codas[coda]
	return ok
}

func (distance *Distance) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Distances[distance]
	return ok
}

func (dynamics *Dynamics) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Dynamicss[dynamics]
	return ok
}

func (empty *Empty) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Emptys[empty]
	return ok
}

func (emptyplacement *EmptyPlacement) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EmptyPlacements[emptyplacement]
	return ok
}

func (emptyplacementsmufl *EmptyPlacementSmufl) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EmptyPlacementSmufls[emptyplacementsmufl]
	return ok
}

func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EmptyPrintObjectStyleAligns[emptyprintobjectstylealign]
	return ok
}

func (emptyprintstyle *EmptyPrintStyle) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EmptyPrintStyles[emptyprintstyle]
	return ok
}

func (emptyprintstylealign *EmptyPrintStyleAlign) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EmptyPrintStyleAligns[emptyprintstylealign]
	return ok
}

func (emptyprintstylealignid *EmptyPrintStyleAlignId) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EmptyPrintStyleAlignIds[emptyprintstylealignid]
	return ok
}

func (emptytrillsound *EmptyTrillSound) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EmptyTrillSounds[emptytrillsound]
	return ok
}

func (encoding *Encoding) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Encodings[encoding]
	return ok
}

func (fermata *Fermata) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Fermatas[fermata]
	return ok
}

func (fingering *Fingering) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Fingerings[fingering]
	return ok
}

func (formattedsymbol *FormattedSymbol) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FormattedSymbols[formattedsymbol]
	return ok
}

func (formattedsymbolid *FormattedSymbolId) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FormattedSymbolIds[formattedsymbolid]
	return ok
}

func (formattedtext *FormattedText) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FormattedTexts[formattedtext]
	return ok
}

func (formattedtextid *FormattedTextId) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FormattedTextIds[formattedtextid]
	return ok
}

func (fret *Fret) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Frets[fret]
	return ok
}

func (glyph *Glyph) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Glyphs[glyph]
	return ok
}

func (horizontalturn *HorizontalTurn) GongIsStaged(stage *Stage) bool {
	_, ok := stage.HorizontalTurns[horizontalturn]
	return ok
}

func (identification *Identification) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Identifications[identification]
	return ok
}

func (level *Level) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Levels[level]
	return ok
}

func (linewidth *LineWidth) GongIsStaged(stage *Stage) bool {
	_, ok := stage.LineWidths[linewidth]
	return ok
}

func (link *Link) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Links[link]
	return ok
}

func (measurelayout *MeasureLayout) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MeasureLayouts[measurelayout]
	return ok
}

func (mididevice *MidiDevice) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MidiDevices[mididevice]
	return ok
}

func (midiinstrument *MidiInstrument) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MidiInstruments[midiinstrument]
	return ok
}

func (miscellaneous *Miscellaneous) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Miscellaneouss[miscellaneous]
	return ok
}

func (miscellaneousfield *MiscellaneousField) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MiscellaneousFields[miscellaneousfield]
	return ok
}

func (namedisplay *NameDisplay) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NameDisplays[namedisplay]
	return ok
}

func (notesize *NoteSize) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteSizes[notesize]
	return ok
}

func (otherappearance *OtherAppearance) GongIsStaged(stage *Stage) bool {
	_, ok := stage.OtherAppearances[otherappearance]
	return ok
}

func (otherplay *OtherPlay) GongIsStaged(stage *Stage) bool {
	_, ok := stage.OtherPlays[otherplay]
	return ok
}

func (othertext *OtherText) GongIsStaged(stage *Stage) bool {
	_, ok := stage.OtherTexts[othertext]
	return ok
}

func (pagelayout *PageLayout) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PageLayouts[pagelayout]
	return ok
}

func (pagemargins *PageMargins) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PageMarginss[pagemargins]
	return ok
}

func (play *Play) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Plays[play]
	return ok
}

func (scaling *Scaling) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Scalings[scaling]
	return ok
}

func (segno *Segno) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Segnos[segno]
	return ok
}

func (stafflayout *StaffLayout) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StaffLayouts[stafflayout]
	return ok
}

func (stringtype *StringType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StringTypes[stringtype]
	return ok
}

func (supports *Supports) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Supportss[supports]
	return ok
}

func (systemdividers *SystemDividers) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SystemDividerss[systemdividers]
	return ok
}

func (systemlayout *SystemLayout) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SystemLayouts[systemlayout]
	return ok
}

func (systemmargins *SystemMargins) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SystemMarginss[systemmargins]
	return ok
}

func (typedtext *TypedText) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TypedTexts[typedtext]
	return ok
}

func (wavyline *WavyLine) GongIsStaged(stage *Stage) bool {
	_, ok := stage.WavyLines[wavyline]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (accidentaltext *AccidentalText) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(accidentaltext) {
		return
	}

	accidentaltext.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (appearance *Appearance) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(appearance) {
		return
	}

	appearance.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linewidth := range appearance.LineWidth {
		stage.StageBranch(_linewidth)
	}
	for _, _notesize := range appearance.NoteSize {
		stage.StageBranch(_notesize)
	}
	for _, _distance := range appearance.Distance {
		stage.StageBranch(_distance)
	}
	for _, _glyph := range appearance.Glyph {
		stage.StageBranch(_glyph)
	}
	for _, _otherappearance := range appearance.OtherAppearance {
		stage.StageBranch(_otherappearance)
	}

}

func (bookmark *Bookmark) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(bookmark) {
		return
	}

	bookmark.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coda *Coda) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coda) {
		return
	}

	coda.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (distance *Distance) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(distance) {
		return
	}

	distance.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dynamics *Dynamics) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(dynamics) {
		return
	}

	dynamics.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _empty := range dynamics.P {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Pp {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Ppp {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Pppp {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Ppppp {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Pppppp {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.F {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Ff {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Fff {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Ffff {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Fffff {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Ffffff {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Mp {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Mf {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Sf {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Sfp {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Sfpp {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Fp {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Rf {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Rfz {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Sfz {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Sffz {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Fz {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.N {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Pf {
		stage.StageBranch(_empty)
	}
	for _, _empty := range dynamics.Sfzp {
		stage.StageBranch(_empty)
	}
	for _, _othertext := range dynamics.OtherDynamics {
		stage.StageBranch(_othertext)
	}

}

func (empty *Empty) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(empty) {
		return
	}

	empty.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (emptyplacement *EmptyPlacement) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(emptyplacement) {
		return
	}

	emptyplacement.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (emptyplacementsmufl *EmptyPlacementSmufl) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(emptyplacementsmufl) {
		return
	}

	emptyplacementsmufl.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(emptyprintobjectstylealign) {
		return
	}

	emptyprintobjectstylealign.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (emptyprintstyle *EmptyPrintStyle) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(emptyprintstyle) {
		return
	}

	emptyprintstyle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (emptyprintstylealign *EmptyPrintStyleAlign) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(emptyprintstylealign) {
		return
	}

	emptyprintstylealign.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (emptyprintstylealignid *EmptyPrintStyleAlignId) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(emptyprintstylealignid) {
		return
	}

	emptyprintstylealignid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (emptytrillsound *EmptyTrillSound) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(emptytrillsound) {
		return
	}

	emptytrillsound.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (encoding *Encoding) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(encoding) {
		return
	}

	encoding.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _typedtext := range encoding.Encoder {
		stage.StageBranch(_typedtext)
	}
	for _, _supports := range encoding.Supports {
		stage.StageBranch(_supports)
	}

}

func (fermata *Fermata) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(fermata) {
		return
	}

	fermata.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (fingering *Fingering) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(fingering) {
		return
	}

	fingering.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formattedsymbol *FormattedSymbol) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(formattedsymbol) {
		return
	}

	formattedsymbol.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formattedsymbolid *FormattedSymbolId) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(formattedsymbolid) {
		return
	}

	formattedsymbolid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formattedtext *FormattedText) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(formattedtext) {
		return
	}

	formattedtext.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formattedtextid *FormattedTextId) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(formattedtextid) {
		return
	}

	formattedtextid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (fret *Fret) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(fret) {
		return
	}

	fret.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (glyph *Glyph) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(glyph) {
		return
	}

	glyph.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (horizontalturn *HorizontalTurn) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(horizontalturn) {
		return
	}

	horizontalturn.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (identification *Identification) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(identification) {
		return
	}

	identification.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if identification.Encoding != nil {
		stage.StageBranch(identification.Encoding)
	}
	if identification.Miscellaneous != nil {
		stage.StageBranch(identification.Miscellaneous)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _typedtext := range identification.Creator {
		stage.StageBranch(_typedtext)
	}
	for _, _typedtext := range identification.Rights {
		stage.StageBranch(_typedtext)
	}
	for _, _typedtext := range identification.Relation {
		stage.StageBranch(_typedtext)
	}

}

func (level *Level) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(level) {
		return
	}

	level.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (linewidth *LineWidth) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(linewidth) {
		return
	}

	linewidth.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (link *Link) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(link) {
		return
	}

	link.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measurelayout *MeasureLayout) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(measurelayout) {
		return
	}

	measurelayout.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (mididevice *MidiDevice) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(mididevice) {
		return
	}

	mididevice.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (midiinstrument *MidiInstrument) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(midiinstrument) {
		return
	}

	midiinstrument.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (miscellaneous *Miscellaneous) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(miscellaneous) {
		return
	}

	miscellaneous.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _miscellaneousfield := range miscellaneous.MiscellaneousField {
		stage.StageBranch(_miscellaneousfield)
	}

}

func (miscellaneousfield *MiscellaneousField) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(miscellaneousfield) {
		return
	}

	miscellaneousfield.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (namedisplay *NameDisplay) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(namedisplay) {
		return
	}

	namedisplay.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formattedtext := range namedisplay.DisplayText {
		stage.StageBranch(_formattedtext)
	}
	for _, _accidentaltext := range namedisplay.AccidentalText {
		stage.StageBranch(_accidentaltext)
	}

}

func (notesize *NoteSize) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(notesize) {
		return
	}

	notesize.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (otherappearance *OtherAppearance) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(otherappearance) {
		return
	}

	otherappearance.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (otherplay *OtherPlay) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(otherplay) {
		return
	}

	otherplay.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (othertext *OtherText) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(othertext) {
		return
	}

	othertext.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pagelayout *PageLayout) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(pagelayout) {
		return
	}

	pagelayout.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _pagemargins := range pagelayout.PageMargins {
		stage.StageBranch(_pagemargins)
	}

}

func (pagemargins *PageMargins) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(pagemargins) {
		return
	}

	pagemargins.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (play *Play) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(play) {
		return
	}

	play.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _otherplay := range play.OtherPlay {
		stage.StageBranch(_otherplay)
	}

}

func (scaling *Scaling) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(scaling) {
		return
	}

	scaling.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (segno *Segno) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(segno) {
		return
	}

	segno.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stafflayout *StaffLayout) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stafflayout) {
		return
	}

	stafflayout.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stringtype *StringType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stringtype) {
		return
	}

	stringtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (supports *Supports) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(supports) {
		return
	}

	supports.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (systemdividers *SystemDividers) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(systemdividers) {
		return
	}

	systemdividers.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if systemdividers.LeftDivider != nil {
		stage.StageBranch(systemdividers.LeftDivider)
	}
	if systemdividers.RightDivider != nil {
		stage.StageBranch(systemdividers.RightDivider)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (systemlayout *SystemLayout) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(systemlayout) {
		return
	}

	systemlayout.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if systemlayout.SystemMargins != nil {
		stage.StageBranch(systemlayout.SystemMargins)
	}
	if systemlayout.SystemDividers != nil {
		stage.StageBranch(systemlayout.SystemDividers)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (systemmargins *SystemMargins) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(systemmargins) {
		return
	}

	systemmargins.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (typedtext *TypedText) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(typedtext) {
		return
	}

	typedtext.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (wavyline *WavyLine) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(wavyline) {
		return
	}

	wavyline.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *AccidentalText:
		toT := GongCopyBranchAccidentalText(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Appearance:
		toT := GongCopyBranchAppearance(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bookmark:
		toT := GongCopyBranchBookmark(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Coda:
		toT := GongCopyBranchCoda(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Distance:
		toT := GongCopyBranchDistance(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Dynamics:
		toT := GongCopyBranchDynamics(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty:
		toT := GongCopyBranchEmpty(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EmptyPlacement:
		toT := GongCopyBranchEmptyPlacement(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EmptyPlacementSmufl:
		toT := GongCopyBranchEmptyPlacementSmufl(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EmptyPrintObjectStyleAlign:
		toT := GongCopyBranchEmptyPrintObjectStyleAlign(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EmptyPrintStyle:
		toT := GongCopyBranchEmptyPrintStyle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EmptyPrintStyleAlign:
		toT := GongCopyBranchEmptyPrintStyleAlign(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EmptyPrintStyleAlignId:
		toT := GongCopyBranchEmptyPrintStyleAlignId(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EmptyTrillSound:
		toT := GongCopyBranchEmptyTrillSound(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Encoding:
		toT := GongCopyBranchEncoding(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Fermata:
		toT := GongCopyBranchFermata(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Fingering:
		toT := GongCopyBranchFingering(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormattedSymbol:
		toT := GongCopyBranchFormattedSymbol(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormattedSymbolId:
		toT := GongCopyBranchFormattedSymbolId(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormattedText:
		toT := GongCopyBranchFormattedText(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormattedTextId:
		toT := GongCopyBranchFormattedTextId(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Fret:
		toT := GongCopyBranchFret(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Glyph:
		toT := GongCopyBranchGlyph(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *HorizontalTurn:
		toT := GongCopyBranchHorizontalTurn(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Identification:
		toT := GongCopyBranchIdentification(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Level:
		toT := GongCopyBranchLevel(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LineWidth:
		toT := GongCopyBranchLineWidth(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Link:
		toT := GongCopyBranchLink(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MeasureLayout:
		toT := GongCopyBranchMeasureLayout(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MidiDevice:
		toT := GongCopyBranchMidiDevice(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MidiInstrument:
		toT := GongCopyBranchMidiInstrument(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Miscellaneous:
		toT := GongCopyBranchMiscellaneous(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MiscellaneousField:
		toT := GongCopyBranchMiscellaneousField(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NameDisplay:
		toT := GongCopyBranchNameDisplay(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteSize:
		toT := GongCopyBranchNoteSize(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *OtherAppearance:
		toT := GongCopyBranchOtherAppearance(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *OtherPlay:
		toT := GongCopyBranchOtherPlay(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *OtherText:
		toT := GongCopyBranchOtherText(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PageLayout:
		toT := GongCopyBranchPageLayout(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PageMargins:
		toT := GongCopyBranchPageMargins(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Play:
		toT := GongCopyBranchPlay(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Scaling:
		toT := GongCopyBranchScaling(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Segno:
		toT := GongCopyBranchSegno(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StaffLayout:
		toT := GongCopyBranchStaffLayout(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StringType:
		toT := GongCopyBranchStringType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Supports:
		toT := GongCopyBranchSupports(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SystemDividers:
		toT := GongCopyBranchSystemDividers(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SystemLayout:
		toT := GongCopyBranchSystemLayout(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SystemMargins:
		toT := GongCopyBranchSystemMargins(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TypedText:
		toT := GongCopyBranchTypedText(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *WavyLine:
		toT := GongCopyBranchWavyLine(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAccidentalText(mapOrigCopy map[any]any, accidentaltextFrom *AccidentalText) (accidentaltextTo *AccidentalText) {
	var alreadyCopied bool
	accidentaltextTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, accidentaltextFrom)
	if alreadyCopied {
		return
	}
	accidentaltextFrom.GongCopyBasicFields(accidentaltextTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAppearance(mapOrigCopy map[any]any, appearanceFrom *Appearance) (appearanceTo *Appearance) {
	var alreadyCopied bool
	appearanceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, appearanceFrom)
	if alreadyCopied {
		return
	}
	appearanceFrom.GongCopyBasicFields(appearanceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linewidth := range appearanceFrom.LineWidth {
		appearanceTo.LineWidth = append(appearanceTo.LineWidth, GongCopyBranchLineWidth(mapOrigCopy, _linewidth))
	}
	for _, _notesize := range appearanceFrom.NoteSize {
		appearanceTo.NoteSize = append(appearanceTo.NoteSize, GongCopyBranchNoteSize(mapOrigCopy, _notesize))
	}
	for _, _distance := range appearanceFrom.Distance {
		appearanceTo.Distance = append(appearanceTo.Distance, GongCopyBranchDistance(mapOrigCopy, _distance))
	}
	for _, _glyph := range appearanceFrom.Glyph {
		appearanceTo.Glyph = append(appearanceTo.Glyph, GongCopyBranchGlyph(mapOrigCopy, _glyph))
	}
	for _, _otherappearance := range appearanceFrom.OtherAppearance {
		appearanceTo.OtherAppearance = append(appearanceTo.OtherAppearance, GongCopyBranchOtherAppearance(mapOrigCopy, _otherappearance))
	}

	return
}

func GongCopyBranchBookmark(mapOrigCopy map[any]any, bookmarkFrom *Bookmark) (bookmarkTo *Bookmark) {
	var alreadyCopied bool
	bookmarkTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, bookmarkFrom)
	if alreadyCopied {
		return
	}
	bookmarkFrom.GongCopyBasicFields(bookmarkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCoda(mapOrigCopy map[any]any, codaFrom *Coda) (codaTo *Coda) {
	var alreadyCopied bool
	codaTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codaFrom)
	if alreadyCopied {
		return
	}
	codaFrom.GongCopyBasicFields(codaTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDistance(mapOrigCopy map[any]any, distanceFrom *Distance) (distanceTo *Distance) {
	var alreadyCopied bool
	distanceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, distanceFrom)
	if alreadyCopied {
		return
	}
	distanceFrom.GongCopyBasicFields(distanceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDynamics(mapOrigCopy map[any]any, dynamicsFrom *Dynamics) (dynamicsTo *Dynamics) {
	var alreadyCopied bool
	dynamicsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, dynamicsFrom)
	if alreadyCopied {
		return
	}
	dynamicsFrom.GongCopyBasicFields(dynamicsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _empty := range dynamicsFrom.P {
		dynamicsTo.P = append(dynamicsTo.P, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Pp {
		dynamicsTo.Pp = append(dynamicsTo.Pp, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Ppp {
		dynamicsTo.Ppp = append(dynamicsTo.Ppp, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Pppp {
		dynamicsTo.Pppp = append(dynamicsTo.Pppp, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Ppppp {
		dynamicsTo.Ppppp = append(dynamicsTo.Ppppp, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Pppppp {
		dynamicsTo.Pppppp = append(dynamicsTo.Pppppp, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.F {
		dynamicsTo.F = append(dynamicsTo.F, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Ff {
		dynamicsTo.Ff = append(dynamicsTo.Ff, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Fff {
		dynamicsTo.Fff = append(dynamicsTo.Fff, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Ffff {
		dynamicsTo.Ffff = append(dynamicsTo.Ffff, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Fffff {
		dynamicsTo.Fffff = append(dynamicsTo.Fffff, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Ffffff {
		dynamicsTo.Ffffff = append(dynamicsTo.Ffffff, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Mp {
		dynamicsTo.Mp = append(dynamicsTo.Mp, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Mf {
		dynamicsTo.Mf = append(dynamicsTo.Mf, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Sf {
		dynamicsTo.Sf = append(dynamicsTo.Sf, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Sfp {
		dynamicsTo.Sfp = append(dynamicsTo.Sfp, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Sfpp {
		dynamicsTo.Sfpp = append(dynamicsTo.Sfpp, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Fp {
		dynamicsTo.Fp = append(dynamicsTo.Fp, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Rf {
		dynamicsTo.Rf = append(dynamicsTo.Rf, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Rfz {
		dynamicsTo.Rfz = append(dynamicsTo.Rfz, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Sfz {
		dynamicsTo.Sfz = append(dynamicsTo.Sfz, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Sffz {
		dynamicsTo.Sffz = append(dynamicsTo.Sffz, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Fz {
		dynamicsTo.Fz = append(dynamicsTo.Fz, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.N {
		dynamicsTo.N = append(dynamicsTo.N, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Pf {
		dynamicsTo.Pf = append(dynamicsTo.Pf, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _empty := range dynamicsFrom.Sfzp {
		dynamicsTo.Sfzp = append(dynamicsTo.Sfzp, GongCopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _othertext := range dynamicsFrom.OtherDynamics {
		dynamicsTo.OtherDynamics = append(dynamicsTo.OtherDynamics, GongCopyBranchOtherText(mapOrigCopy, _othertext))
	}

	return
}

func GongCopyBranchEmpty(mapOrigCopy map[any]any, emptyFrom *Empty) (emptyTo *Empty) {
	var alreadyCopied bool
	emptyTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, emptyFrom)
	if alreadyCopied {
		return
	}
	emptyFrom.GongCopyBasicFields(emptyTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEmptyPlacement(mapOrigCopy map[any]any, emptyplacementFrom *EmptyPlacement) (emptyplacementTo *EmptyPlacement) {
	var alreadyCopied bool
	emptyplacementTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, emptyplacementFrom)
	if alreadyCopied {
		return
	}
	emptyplacementFrom.GongCopyBasicFields(emptyplacementTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEmptyPlacementSmufl(mapOrigCopy map[any]any, emptyplacementsmuflFrom *EmptyPlacementSmufl) (emptyplacementsmuflTo *EmptyPlacementSmufl) {
	var alreadyCopied bool
	emptyplacementsmuflTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, emptyplacementsmuflFrom)
	if alreadyCopied {
		return
	}
	emptyplacementsmuflFrom.GongCopyBasicFields(emptyplacementsmuflTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEmptyPrintObjectStyleAlign(mapOrigCopy map[any]any, emptyprintobjectstylealignFrom *EmptyPrintObjectStyleAlign) (emptyprintobjectstylealignTo *EmptyPrintObjectStyleAlign) {
	var alreadyCopied bool
	emptyprintobjectstylealignTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, emptyprintobjectstylealignFrom)
	if alreadyCopied {
		return
	}
	emptyprintobjectstylealignFrom.GongCopyBasicFields(emptyprintobjectstylealignTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEmptyPrintStyle(mapOrigCopy map[any]any, emptyprintstyleFrom *EmptyPrintStyle) (emptyprintstyleTo *EmptyPrintStyle) {
	var alreadyCopied bool
	emptyprintstyleTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, emptyprintstyleFrom)
	if alreadyCopied {
		return
	}
	emptyprintstyleFrom.GongCopyBasicFields(emptyprintstyleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEmptyPrintStyleAlign(mapOrigCopy map[any]any, emptyprintstylealignFrom *EmptyPrintStyleAlign) (emptyprintstylealignTo *EmptyPrintStyleAlign) {
	var alreadyCopied bool
	emptyprintstylealignTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, emptyprintstylealignFrom)
	if alreadyCopied {
		return
	}
	emptyprintstylealignFrom.GongCopyBasicFields(emptyprintstylealignTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEmptyPrintStyleAlignId(mapOrigCopy map[any]any, emptyprintstylealignidFrom *EmptyPrintStyleAlignId) (emptyprintstylealignidTo *EmptyPrintStyleAlignId) {
	var alreadyCopied bool
	emptyprintstylealignidTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, emptyprintstylealignidFrom)
	if alreadyCopied {
		return
	}
	emptyprintstylealignidFrom.GongCopyBasicFields(emptyprintstylealignidTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEmptyTrillSound(mapOrigCopy map[any]any, emptytrillsoundFrom *EmptyTrillSound) (emptytrillsoundTo *EmptyTrillSound) {
	var alreadyCopied bool
	emptytrillsoundTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, emptytrillsoundFrom)
	if alreadyCopied {
		return
	}
	emptytrillsoundFrom.GongCopyBasicFields(emptytrillsoundTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEncoding(mapOrigCopy map[any]any, encodingFrom *Encoding) (encodingTo *Encoding) {
	var alreadyCopied bool
	encodingTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, encodingFrom)
	if alreadyCopied {
		return
	}
	encodingFrom.GongCopyBasicFields(encodingTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _typedtext := range encodingFrom.Encoder {
		encodingTo.Encoder = append(encodingTo.Encoder, GongCopyBranchTypedText(mapOrigCopy, _typedtext))
	}
	for _, _supports := range encodingFrom.Supports {
		encodingTo.Supports = append(encodingTo.Supports, GongCopyBranchSupports(mapOrigCopy, _supports))
	}

	return
}

func GongCopyBranchFermata(mapOrigCopy map[any]any, fermataFrom *Fermata) (fermataTo *Fermata) {
	var alreadyCopied bool
	fermataTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, fermataFrom)
	if alreadyCopied {
		return
	}
	fermataFrom.GongCopyBasicFields(fermataTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFingering(mapOrigCopy map[any]any, fingeringFrom *Fingering) (fingeringTo *Fingering) {
	var alreadyCopied bool
	fingeringTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, fingeringFrom)
	if alreadyCopied {
		return
	}
	fingeringFrom.GongCopyBasicFields(fingeringTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormattedSymbol(mapOrigCopy map[any]any, formattedsymbolFrom *FormattedSymbol) (formattedsymbolTo *FormattedSymbol) {
	var alreadyCopied bool
	formattedsymbolTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formattedsymbolFrom)
	if alreadyCopied {
		return
	}
	formattedsymbolFrom.GongCopyBasicFields(formattedsymbolTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormattedSymbolId(mapOrigCopy map[any]any, formattedsymbolidFrom *FormattedSymbolId) (formattedsymbolidTo *FormattedSymbolId) {
	var alreadyCopied bool
	formattedsymbolidTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formattedsymbolidFrom)
	if alreadyCopied {
		return
	}
	formattedsymbolidFrom.GongCopyBasicFields(formattedsymbolidTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormattedText(mapOrigCopy map[any]any, formattedtextFrom *FormattedText) (formattedtextTo *FormattedText) {
	var alreadyCopied bool
	formattedtextTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formattedtextFrom)
	if alreadyCopied {
		return
	}
	formattedtextFrom.GongCopyBasicFields(formattedtextTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormattedTextId(mapOrigCopy map[any]any, formattedtextidFrom *FormattedTextId) (formattedtextidTo *FormattedTextId) {
	var alreadyCopied bool
	formattedtextidTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formattedtextidFrom)
	if alreadyCopied {
		return
	}
	formattedtextidFrom.GongCopyBasicFields(formattedtextidTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFret(mapOrigCopy map[any]any, fretFrom *Fret) (fretTo *Fret) {
	var alreadyCopied bool
	fretTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, fretFrom)
	if alreadyCopied {
		return
	}
	fretFrom.GongCopyBasicFields(fretTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGlyph(mapOrigCopy map[any]any, glyphFrom *Glyph) (glyphTo *Glyph) {
	var alreadyCopied bool
	glyphTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, glyphFrom)
	if alreadyCopied {
		return
	}
	glyphFrom.GongCopyBasicFields(glyphTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchHorizontalTurn(mapOrigCopy map[any]any, horizontalturnFrom *HorizontalTurn) (horizontalturnTo *HorizontalTurn) {
	var alreadyCopied bool
	horizontalturnTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, horizontalturnFrom)
	if alreadyCopied {
		return
	}
	horizontalturnFrom.GongCopyBasicFields(horizontalturnTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchIdentification(mapOrigCopy map[any]any, identificationFrom *Identification) (identificationTo *Identification) {
	var alreadyCopied bool
	identificationTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, identificationFrom)
	if alreadyCopied {
		return
	}
	identificationFrom.GongCopyBasicFields(identificationTo)

	//insertion point for the staging of instances referenced by pointers
	if identificationFrom.Encoding != nil {
		identificationTo.Encoding = GongCopyBranchEncoding(mapOrigCopy, identificationFrom.Encoding)
	}
	if identificationFrom.Miscellaneous != nil {
		identificationTo.Miscellaneous = GongCopyBranchMiscellaneous(mapOrigCopy, identificationFrom.Miscellaneous)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _typedtext := range identificationFrom.Creator {
		identificationTo.Creator = append(identificationTo.Creator, GongCopyBranchTypedText(mapOrigCopy, _typedtext))
	}
	for _, _typedtext := range identificationFrom.Rights {
		identificationTo.Rights = append(identificationTo.Rights, GongCopyBranchTypedText(mapOrigCopy, _typedtext))
	}
	for _, _typedtext := range identificationFrom.Relation {
		identificationTo.Relation = append(identificationTo.Relation, GongCopyBranchTypedText(mapOrigCopy, _typedtext))
	}

	return
}

func GongCopyBranchLevel(mapOrigCopy map[any]any, levelFrom *Level) (levelTo *Level) {
	var alreadyCopied bool
	levelTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, levelFrom)
	if alreadyCopied {
		return
	}
	levelFrom.GongCopyBasicFields(levelTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLineWidth(mapOrigCopy map[any]any, linewidthFrom *LineWidth) (linewidthTo *LineWidth) {
	var alreadyCopied bool
	linewidthTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, linewidthFrom)
	if alreadyCopied {
		return
	}
	linewidthFrom.GongCopyBasicFields(linewidthTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLink(mapOrigCopy map[any]any, linkFrom *Link) (linkTo *Link) {
	var alreadyCopied bool
	linkTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, linkFrom)
	if alreadyCopied {
		return
	}
	linkFrom.GongCopyBasicFields(linkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMeasureLayout(mapOrigCopy map[any]any, measurelayoutFrom *MeasureLayout) (measurelayoutTo *MeasureLayout) {
	var alreadyCopied bool
	measurelayoutTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, measurelayoutFrom)
	if alreadyCopied {
		return
	}
	measurelayoutFrom.GongCopyBasicFields(measurelayoutTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMidiDevice(mapOrigCopy map[any]any, midideviceFrom *MidiDevice) (midideviceTo *MidiDevice) {
	var alreadyCopied bool
	midideviceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, midideviceFrom)
	if alreadyCopied {
		return
	}
	midideviceFrom.GongCopyBasicFields(midideviceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMidiInstrument(mapOrigCopy map[any]any, midiinstrumentFrom *MidiInstrument) (midiinstrumentTo *MidiInstrument) {
	var alreadyCopied bool
	midiinstrumentTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, midiinstrumentFrom)
	if alreadyCopied {
		return
	}
	midiinstrumentFrom.GongCopyBasicFields(midiinstrumentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMiscellaneous(mapOrigCopy map[any]any, miscellaneousFrom *Miscellaneous) (miscellaneousTo *Miscellaneous) {
	var alreadyCopied bool
	miscellaneousTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, miscellaneousFrom)
	if alreadyCopied {
		return
	}
	miscellaneousFrom.GongCopyBasicFields(miscellaneousTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _miscellaneousfield := range miscellaneousFrom.MiscellaneousField {
		miscellaneousTo.MiscellaneousField = append(miscellaneousTo.MiscellaneousField, GongCopyBranchMiscellaneousField(mapOrigCopy, _miscellaneousfield))
	}

	return
}

func GongCopyBranchMiscellaneousField(mapOrigCopy map[any]any, miscellaneousfieldFrom *MiscellaneousField) (miscellaneousfieldTo *MiscellaneousField) {
	var alreadyCopied bool
	miscellaneousfieldTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, miscellaneousfieldFrom)
	if alreadyCopied {
		return
	}
	miscellaneousfieldFrom.GongCopyBasicFields(miscellaneousfieldTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNameDisplay(mapOrigCopy map[any]any, namedisplayFrom *NameDisplay) (namedisplayTo *NameDisplay) {
	var alreadyCopied bool
	namedisplayTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, namedisplayFrom)
	if alreadyCopied {
		return
	}
	namedisplayFrom.GongCopyBasicFields(namedisplayTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formattedtext := range namedisplayFrom.DisplayText {
		namedisplayTo.DisplayText = append(namedisplayTo.DisplayText, GongCopyBranchFormattedText(mapOrigCopy, _formattedtext))
	}
	for _, _accidentaltext := range namedisplayFrom.AccidentalText {
		namedisplayTo.AccidentalText = append(namedisplayTo.AccidentalText, GongCopyBranchAccidentalText(mapOrigCopy, _accidentaltext))
	}

	return
}

func GongCopyBranchNoteSize(mapOrigCopy map[any]any, notesizeFrom *NoteSize) (notesizeTo *NoteSize) {
	var alreadyCopied bool
	notesizeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, notesizeFrom)
	if alreadyCopied {
		return
	}
	notesizeFrom.GongCopyBasicFields(notesizeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOtherAppearance(mapOrigCopy map[any]any, otherappearanceFrom *OtherAppearance) (otherappearanceTo *OtherAppearance) {
	var alreadyCopied bool
	otherappearanceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, otherappearanceFrom)
	if alreadyCopied {
		return
	}
	otherappearanceFrom.GongCopyBasicFields(otherappearanceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOtherPlay(mapOrigCopy map[any]any, otherplayFrom *OtherPlay) (otherplayTo *OtherPlay) {
	var alreadyCopied bool
	otherplayTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, otherplayFrom)
	if alreadyCopied {
		return
	}
	otherplayFrom.GongCopyBasicFields(otherplayTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOtherText(mapOrigCopy map[any]any, othertextFrom *OtherText) (othertextTo *OtherText) {
	var alreadyCopied bool
	othertextTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, othertextFrom)
	if alreadyCopied {
		return
	}
	othertextFrom.GongCopyBasicFields(othertextTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPageLayout(mapOrigCopy map[any]any, pagelayoutFrom *PageLayout) (pagelayoutTo *PageLayout) {
	var alreadyCopied bool
	pagelayoutTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pagelayoutFrom)
	if alreadyCopied {
		return
	}
	pagelayoutFrom.GongCopyBasicFields(pagelayoutTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _pagemargins := range pagelayoutFrom.PageMargins {
		pagelayoutTo.PageMargins = append(pagelayoutTo.PageMargins, GongCopyBranchPageMargins(mapOrigCopy, _pagemargins))
	}

	return
}

func GongCopyBranchPageMargins(mapOrigCopy map[any]any, pagemarginsFrom *PageMargins) (pagemarginsTo *PageMargins) {
	var alreadyCopied bool
	pagemarginsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pagemarginsFrom)
	if alreadyCopied {
		return
	}
	pagemarginsFrom.GongCopyBasicFields(pagemarginsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPlay(mapOrigCopy map[any]any, playFrom *Play) (playTo *Play) {
	var alreadyCopied bool
	playTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, playFrom)
	if alreadyCopied {
		return
	}
	playFrom.GongCopyBasicFields(playTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _otherplay := range playFrom.OtherPlay {
		playTo.OtherPlay = append(playTo.OtherPlay, GongCopyBranchOtherPlay(mapOrigCopy, _otherplay))
	}

	return
}

func GongCopyBranchScaling(mapOrigCopy map[any]any, scalingFrom *Scaling) (scalingTo *Scaling) {
	var alreadyCopied bool
	scalingTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, scalingFrom)
	if alreadyCopied {
		return
	}
	scalingFrom.GongCopyBasicFields(scalingTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSegno(mapOrigCopy map[any]any, segnoFrom *Segno) (segnoTo *Segno) {
	var alreadyCopied bool
	segnoTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, segnoFrom)
	if alreadyCopied {
		return
	}
	segnoFrom.GongCopyBasicFields(segnoTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStaffLayout(mapOrigCopy map[any]any, stafflayoutFrom *StaffLayout) (stafflayoutTo *StaffLayout) {
	var alreadyCopied bool
	stafflayoutTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stafflayoutFrom)
	if alreadyCopied {
		return
	}
	stafflayoutFrom.GongCopyBasicFields(stafflayoutTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStringType(mapOrigCopy map[any]any, stringtypeFrom *StringType) (stringtypeTo *StringType) {
	var alreadyCopied bool
	stringtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stringtypeFrom)
	if alreadyCopied {
		return
	}
	stringtypeFrom.GongCopyBasicFields(stringtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSupports(mapOrigCopy map[any]any, supportsFrom *Supports) (supportsTo *Supports) {
	var alreadyCopied bool
	supportsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, supportsFrom)
	if alreadyCopied {
		return
	}
	supportsFrom.GongCopyBasicFields(supportsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSystemDividers(mapOrigCopy map[any]any, systemdividersFrom *SystemDividers) (systemdividersTo *SystemDividers) {
	var alreadyCopied bool
	systemdividersTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, systemdividersFrom)
	if alreadyCopied {
		return
	}
	systemdividersFrom.GongCopyBasicFields(systemdividersTo)

	//insertion point for the staging of instances referenced by pointers
	if systemdividersFrom.LeftDivider != nil {
		systemdividersTo.LeftDivider = GongCopyBranchEmptyPrintObjectStyleAlign(mapOrigCopy, systemdividersFrom.LeftDivider)
	}
	if systemdividersFrom.RightDivider != nil {
		systemdividersTo.RightDivider = GongCopyBranchEmptyPrintObjectStyleAlign(mapOrigCopy, systemdividersFrom.RightDivider)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSystemLayout(mapOrigCopy map[any]any, systemlayoutFrom *SystemLayout) (systemlayoutTo *SystemLayout) {
	var alreadyCopied bool
	systemlayoutTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, systemlayoutFrom)
	if alreadyCopied {
		return
	}
	systemlayoutFrom.GongCopyBasicFields(systemlayoutTo)

	//insertion point for the staging of instances referenced by pointers
	if systemlayoutFrom.SystemMargins != nil {
		systemlayoutTo.SystemMargins = GongCopyBranchSystemMargins(mapOrigCopy, systemlayoutFrom.SystemMargins)
	}
	if systemlayoutFrom.SystemDividers != nil {
		systemlayoutTo.SystemDividers = GongCopyBranchSystemDividers(mapOrigCopy, systemlayoutFrom.SystemDividers)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSystemMargins(mapOrigCopy map[any]any, systemmarginsFrom *SystemMargins) (systemmarginsTo *SystemMargins) {
	var alreadyCopied bool
	systemmarginsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, systemmarginsFrom)
	if alreadyCopied {
		return
	}
	systemmarginsFrom.GongCopyBasicFields(systemmarginsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTypedText(mapOrigCopy map[any]any, typedtextFrom *TypedText) (typedtextTo *TypedText) {
	var alreadyCopied bool
	typedtextTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, typedtextFrom)
	if alreadyCopied {
		return
	}
	typedtextFrom.GongCopyBasicFields(typedtextTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchWavyLine(mapOrigCopy map[any]any, wavylineFrom *WavyLine) (wavylineTo *WavyLine) {
	var alreadyCopied bool
	wavylineTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, wavylineFrom)
	if alreadyCopied {
		return
	}
	wavylineFrom.GongCopyBasicFields(wavylineTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (accidentaltext *AccidentalText) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(accidentaltext) {
		return
	}

	accidentaltext.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (appearance *Appearance) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(appearance) {
		return
	}

	appearance.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linewidth := range appearance.LineWidth {
		stage.UnstageBranch(_linewidth)
	}
	for _, _notesize := range appearance.NoteSize {
		stage.UnstageBranch(_notesize)
	}
	for _, _distance := range appearance.Distance {
		stage.UnstageBranch(_distance)
	}
	for _, _glyph := range appearance.Glyph {
		stage.UnstageBranch(_glyph)
	}
	for _, _otherappearance := range appearance.OtherAppearance {
		stage.UnstageBranch(_otherappearance)
	}

}

func (bookmark *Bookmark) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(bookmark) {
		return
	}

	bookmark.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coda *Coda) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coda) {
		return
	}

	coda.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (distance *Distance) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(distance) {
		return
	}

	distance.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dynamics *Dynamics) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(dynamics) {
		return
	}

	dynamics.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _empty := range dynamics.P {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Pp {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Ppp {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Pppp {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Ppppp {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Pppppp {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.F {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Ff {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Fff {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Ffff {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Fffff {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Ffffff {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Mp {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Mf {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Sf {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Sfp {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Sfpp {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Fp {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Rf {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Rfz {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Sfz {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Sffz {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Fz {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.N {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Pf {
		stage.UnstageBranch(_empty)
	}
	for _, _empty := range dynamics.Sfzp {
		stage.UnstageBranch(_empty)
	}
	for _, _othertext := range dynamics.OtherDynamics {
		stage.UnstageBranch(_othertext)
	}

}

func (empty *Empty) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(empty) {
		return
	}

	empty.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (emptyplacement *EmptyPlacement) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(emptyplacement) {
		return
	}

	emptyplacement.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (emptyplacementsmufl *EmptyPlacementSmufl) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(emptyplacementsmufl) {
		return
	}

	emptyplacementsmufl.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(emptyprintobjectstylealign) {
		return
	}

	emptyprintobjectstylealign.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (emptyprintstyle *EmptyPrintStyle) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(emptyprintstyle) {
		return
	}

	emptyprintstyle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (emptyprintstylealign *EmptyPrintStyleAlign) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(emptyprintstylealign) {
		return
	}

	emptyprintstylealign.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (emptyprintstylealignid *EmptyPrintStyleAlignId) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(emptyprintstylealignid) {
		return
	}

	emptyprintstylealignid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (emptytrillsound *EmptyTrillSound) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(emptytrillsound) {
		return
	}

	emptytrillsound.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (encoding *Encoding) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(encoding) {
		return
	}

	encoding.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _typedtext := range encoding.Encoder {
		stage.UnstageBranch(_typedtext)
	}
	for _, _supports := range encoding.Supports {
		stage.UnstageBranch(_supports)
	}

}

func (fermata *Fermata) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(fermata) {
		return
	}

	fermata.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (fingering *Fingering) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(fingering) {
		return
	}

	fingering.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formattedsymbol *FormattedSymbol) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(formattedsymbol) {
		return
	}

	formattedsymbol.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formattedsymbolid *FormattedSymbolId) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(formattedsymbolid) {
		return
	}

	formattedsymbolid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formattedtext *FormattedText) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(formattedtext) {
		return
	}

	formattedtext.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formattedtextid *FormattedTextId) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(formattedtextid) {
		return
	}

	formattedtextid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (fret *Fret) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(fret) {
		return
	}

	fret.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (glyph *Glyph) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(glyph) {
		return
	}

	glyph.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (horizontalturn *HorizontalTurn) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(horizontalturn) {
		return
	}

	horizontalturn.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (identification *Identification) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(identification) {
		return
	}

	identification.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if identification.Encoding != nil {
		stage.UnstageBranch(identification.Encoding)
	}
	if identification.Miscellaneous != nil {
		stage.UnstageBranch(identification.Miscellaneous)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _typedtext := range identification.Creator {
		stage.UnstageBranch(_typedtext)
	}
	for _, _typedtext := range identification.Rights {
		stage.UnstageBranch(_typedtext)
	}
	for _, _typedtext := range identification.Relation {
		stage.UnstageBranch(_typedtext)
	}

}

func (level *Level) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(level) {
		return
	}

	level.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (linewidth *LineWidth) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(linewidth) {
		return
	}

	linewidth.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (link *Link) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(link) {
		return
	}

	link.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measurelayout *MeasureLayout) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(measurelayout) {
		return
	}

	measurelayout.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (mididevice *MidiDevice) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(mididevice) {
		return
	}

	mididevice.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (midiinstrument *MidiInstrument) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(midiinstrument) {
		return
	}

	midiinstrument.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (miscellaneous *Miscellaneous) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(miscellaneous) {
		return
	}

	miscellaneous.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _miscellaneousfield := range miscellaneous.MiscellaneousField {
		stage.UnstageBranch(_miscellaneousfield)
	}

}

func (miscellaneousfield *MiscellaneousField) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(miscellaneousfield) {
		return
	}

	miscellaneousfield.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (namedisplay *NameDisplay) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(namedisplay) {
		return
	}

	namedisplay.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formattedtext := range namedisplay.DisplayText {
		stage.UnstageBranch(_formattedtext)
	}
	for _, _accidentaltext := range namedisplay.AccidentalText {
		stage.UnstageBranch(_accidentaltext)
	}

}

func (notesize *NoteSize) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(notesize) {
		return
	}

	notesize.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (otherappearance *OtherAppearance) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(otherappearance) {
		return
	}

	otherappearance.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (otherplay *OtherPlay) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(otherplay) {
		return
	}

	otherplay.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (othertext *OtherText) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(othertext) {
		return
	}

	othertext.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pagelayout *PageLayout) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(pagelayout) {
		return
	}

	pagelayout.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _pagemargins := range pagelayout.PageMargins {
		stage.UnstageBranch(_pagemargins)
	}

}

func (pagemargins *PageMargins) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(pagemargins) {
		return
	}

	pagemargins.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (play *Play) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(play) {
		return
	}

	play.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _otherplay := range play.OtherPlay {
		stage.UnstageBranch(_otherplay)
	}

}

func (scaling *Scaling) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(scaling) {
		return
	}

	scaling.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (segno *Segno) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(segno) {
		return
	}

	segno.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stafflayout *StaffLayout) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stafflayout) {
		return
	}

	stafflayout.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stringtype *StringType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stringtype) {
		return
	}

	stringtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (supports *Supports) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(supports) {
		return
	}

	supports.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (systemdividers *SystemDividers) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(systemdividers) {
		return
	}

	systemdividers.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if systemdividers.LeftDivider != nil {
		stage.UnstageBranch(systemdividers.LeftDivider)
	}
	if systemdividers.RightDivider != nil {
		stage.UnstageBranch(systemdividers.RightDivider)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (systemlayout *SystemLayout) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(systemlayout) {
		return
	}

	systemlayout.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if systemlayout.SystemMargins != nil {
		stage.UnstageBranch(systemlayout.SystemMargins)
	}
	if systemlayout.SystemDividers != nil {
		stage.UnstageBranch(systemlayout.SystemDividers)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (systemmargins *SystemMargins) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(systemmargins) {
		return
	}

	systemmargins.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (typedtext *TypedText) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(typedtext) {
		return
	}

	typedtext.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (wavyline *WavyLine) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(wavyline) {
		return
	}

	wavyline.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *AccidentalText) GongReconstructPointersFromReferences(stage *Stage, instance *AccidentalText) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Appearance) GongReconstructPointersFromReferences(stage *Stage, instance *Appearance) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.LineWidth, stage.LineWidths_reference, instance.LineWidth)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NoteSize, stage.NoteSizes_reference, instance.NoteSize)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Distance, stage.Distances_reference, instance.Distance)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Glyph, stage.Glyphs_reference, instance.Glyph)
	__gong__reconstructSliceOfPointersFromReferences(&reference.OtherAppearance, stage.OtherAppearances_reference, instance.OtherAppearance)
}

func (reference *Bookmark) GongReconstructPointersFromReferences(stage *Stage, instance *Bookmark) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Coda) GongReconstructPointersFromReferences(stage *Stage, instance *Coda) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Distance) GongReconstructPointersFromReferences(stage *Stage, instance *Distance) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Dynamics) GongReconstructPointersFromReferences(stage *Stage, instance *Dynamics) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.P, stage.Emptys_reference, instance.P)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Pp, stage.Emptys_reference, instance.Pp)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Ppp, stage.Emptys_reference, instance.Ppp)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Pppp, stage.Emptys_reference, instance.Pppp)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Ppppp, stage.Emptys_reference, instance.Ppppp)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Pppppp, stage.Emptys_reference, instance.Pppppp)
	__gong__reconstructSliceOfPointersFromReferences(&reference.F, stage.Emptys_reference, instance.F)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Ff, stage.Emptys_reference, instance.Ff)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Fff, stage.Emptys_reference, instance.Fff)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Ffff, stage.Emptys_reference, instance.Ffff)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Fffff, stage.Emptys_reference, instance.Fffff)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Ffffff, stage.Emptys_reference, instance.Ffffff)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Mp, stage.Emptys_reference, instance.Mp)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Mf, stage.Emptys_reference, instance.Mf)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sf, stage.Emptys_reference, instance.Sf)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sfp, stage.Emptys_reference, instance.Sfp)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sfpp, stage.Emptys_reference, instance.Sfpp)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Fp, stage.Emptys_reference, instance.Fp)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Rf, stage.Emptys_reference, instance.Rf)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Rfz, stage.Emptys_reference, instance.Rfz)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sfz, stage.Emptys_reference, instance.Sfz)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sffz, stage.Emptys_reference, instance.Sffz)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Fz, stage.Emptys_reference, instance.Fz)
	__gong__reconstructSliceOfPointersFromReferences(&reference.N, stage.Emptys_reference, instance.N)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Pf, stage.Emptys_reference, instance.Pf)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sfzp, stage.Emptys_reference, instance.Sfzp)
	__gong__reconstructSliceOfPointersFromReferences(&reference.OtherDynamics, stage.OtherTexts_reference, instance.OtherDynamics)
}

func (reference *Empty) GongReconstructPointersFromReferences(stage *Stage, instance *Empty) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EmptyPlacement) GongReconstructPointersFromReferences(stage *Stage, instance *EmptyPlacement) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EmptyPlacementSmufl) GongReconstructPointersFromReferences(stage *Stage, instance *EmptyPlacementSmufl) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EmptyPrintObjectStyleAlign) GongReconstructPointersFromReferences(stage *Stage, instance *EmptyPrintObjectStyleAlign) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EmptyPrintStyle) GongReconstructPointersFromReferences(stage *Stage, instance *EmptyPrintStyle) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EmptyPrintStyleAlign) GongReconstructPointersFromReferences(stage *Stage, instance *EmptyPrintStyleAlign) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EmptyPrintStyleAlignId) GongReconstructPointersFromReferences(stage *Stage, instance *EmptyPrintStyleAlignId) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EmptyTrillSound) GongReconstructPointersFromReferences(stage *Stage, instance *EmptyTrillSound) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Encoding) GongReconstructPointersFromReferences(stage *Stage, instance *Encoding) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Encoder, stage.TypedTexts_reference, instance.Encoder)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Supports, stage.Supportss_reference, instance.Supports)
}

func (reference *Fermata) GongReconstructPointersFromReferences(stage *Stage, instance *Fermata) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Fingering) GongReconstructPointersFromReferences(stage *Stage, instance *Fingering) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *FormattedSymbol) GongReconstructPointersFromReferences(stage *Stage, instance *FormattedSymbol) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *FormattedSymbolId) GongReconstructPointersFromReferences(stage *Stage, instance *FormattedSymbolId) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *FormattedText) GongReconstructPointersFromReferences(stage *Stage, instance *FormattedText) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *FormattedTextId) GongReconstructPointersFromReferences(stage *Stage, instance *FormattedTextId) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Fret) GongReconstructPointersFromReferences(stage *Stage, instance *Fret) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Glyph) GongReconstructPointersFromReferences(stage *Stage, instance *Glyph) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *HorizontalTurn) GongReconstructPointersFromReferences(stage *Stage, instance *HorizontalTurn) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Identification) GongReconstructPointersFromReferences(stage *Stage, instance *Identification) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Encoding, stage.Encodings_reference, instance.Encoding)
	__gong__reconstructPointer(&reference.Miscellaneous, stage.Miscellaneouss_reference, instance.Miscellaneous)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Creator, stage.TypedTexts_reference, instance.Creator)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Rights, stage.TypedTexts_reference, instance.Rights)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Relation, stage.TypedTexts_reference, instance.Relation)
}

func (reference *Level) GongReconstructPointersFromReferences(stage *Stage, instance *Level) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *LineWidth) GongReconstructPointersFromReferences(stage *Stage, instance *LineWidth) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Link) GongReconstructPointersFromReferences(stage *Stage, instance *Link) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MeasureLayout) GongReconstructPointersFromReferences(stage *Stage, instance *MeasureLayout) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MidiDevice) GongReconstructPointersFromReferences(stage *Stage, instance *MidiDevice) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MidiInstrument) GongReconstructPointersFromReferences(stage *Stage, instance *MidiInstrument) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Miscellaneous) GongReconstructPointersFromReferences(stage *Stage, instance *Miscellaneous) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.MiscellaneousField, stage.MiscellaneousFields_reference, instance.MiscellaneousField)
}

func (reference *MiscellaneousField) GongReconstructPointersFromReferences(stage *Stage, instance *MiscellaneousField) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *NameDisplay) GongReconstructPointersFromReferences(stage *Stage, instance *NameDisplay) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.DisplayText, stage.FormattedTexts_reference, instance.DisplayText)
	__gong__reconstructSliceOfPointersFromReferences(&reference.AccidentalText, stage.AccidentalTexts_reference, instance.AccidentalText)
}

func (reference *NoteSize) GongReconstructPointersFromReferences(stage *Stage, instance *NoteSize) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *OtherAppearance) GongReconstructPointersFromReferences(stage *Stage, instance *OtherAppearance) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *OtherPlay) GongReconstructPointersFromReferences(stage *Stage, instance *OtherPlay) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *OtherText) GongReconstructPointersFromReferences(stage *Stage, instance *OtherText) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PageLayout) GongReconstructPointersFromReferences(stage *Stage, instance *PageLayout) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.PageMargins, stage.PageMarginss_reference, instance.PageMargins)
}

func (reference *PageMargins) GongReconstructPointersFromReferences(stage *Stage, instance *PageMargins) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Play) GongReconstructPointersFromReferences(stage *Stage, instance *Play) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.OtherPlay, stage.OtherPlays_reference, instance.OtherPlay)
}

func (reference *Scaling) GongReconstructPointersFromReferences(stage *Stage, instance *Scaling) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Segno) GongReconstructPointersFromReferences(stage *Stage, instance *Segno) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StaffLayout) GongReconstructPointersFromReferences(stage *Stage, instance *StaffLayout) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StringType) GongReconstructPointersFromReferences(stage *Stage, instance *StringType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Supports) GongReconstructPointersFromReferences(stage *Stage, instance *Supports) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *SystemDividers) GongReconstructPointersFromReferences(stage *Stage, instance *SystemDividers) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.LeftDivider, stage.EmptyPrintObjectStyleAligns_reference, instance.LeftDivider)
	__gong__reconstructPointer(&reference.RightDivider, stage.EmptyPrintObjectStyleAligns_reference, instance.RightDivider)
	// insertion point for slice of pointers field
}

func (reference *SystemLayout) GongReconstructPointersFromReferences(stage *Stage, instance *SystemLayout) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.SystemMargins, stage.SystemMarginss_reference, instance.SystemMargins)
	__gong__reconstructPointer(&reference.SystemDividers, stage.SystemDividerss_reference, instance.SystemDividers)
	// insertion point for slice of pointers field
}

func (reference *SystemMargins) GongReconstructPointersFromReferences(stage *Stage, instance *SystemMargins) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TypedText) GongReconstructPointersFromReferences(stage *Stage, instance *TypedText) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *WavyLine) GongReconstructPointersFromReferences(stage *Stage, instance *WavyLine) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *AccidentalText) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Appearance) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.LineWidth, stage.LineWidths_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NoteSize, stage.NoteSizes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Distance, stage.Distances_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Glyph, stage.Glyphs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.OtherAppearance, stage.OtherAppearances_instance)
}

func (reference *Bookmark) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Coda) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Distance) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Dynamics) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.P, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Pp, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Ppp, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Pppp, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Ppppp, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Pppppp, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.F, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Ff, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Fff, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Ffff, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Fffff, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Ffffff, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Mp, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Mf, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sf, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sfp, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sfpp, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Fp, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Rf, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Rfz, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sfz, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sffz, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Fz, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.N, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Pf, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sfzp, stage.Emptys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.OtherDynamics, stage.OtherTexts_instance)
}

func (reference *Empty) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EmptyPlacement) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EmptyPlacementSmufl) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EmptyPrintObjectStyleAlign) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EmptyPrintStyle) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EmptyPrintStyleAlign) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EmptyPrintStyleAlignId) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EmptyTrillSound) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Encoding) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Encoder, stage.TypedTexts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Supports, stage.Supportss_instance)
}

func (reference *Fermata) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Fingering) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *FormattedSymbol) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *FormattedSymbolId) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *FormattedText) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *FormattedTextId) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Fret) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Glyph) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *HorizontalTurn) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Identification) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Encoding, stage.Encodings_instance)
	__gong__reconstructPointerFromInstance(&reference.Miscellaneous, stage.Miscellaneouss_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Creator, stage.TypedTexts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Rights, stage.TypedTexts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Relation, stage.TypedTexts_instance)
}

func (reference *Level) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *LineWidth) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Link) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MeasureLayout) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MidiDevice) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MidiInstrument) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Miscellaneous) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.MiscellaneousField, stage.MiscellaneousFields_instance)
}

func (reference *MiscellaneousField) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *NameDisplay) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.DisplayText, stage.FormattedTexts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.AccidentalText, stage.AccidentalTexts_instance)
}

func (reference *NoteSize) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *OtherAppearance) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *OtherPlay) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *OtherText) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PageLayout) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.PageMargins, stage.PageMarginss_instance)
}

func (reference *PageMargins) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Play) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.OtherPlay, stage.OtherPlays_instance)
}

func (reference *Scaling) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Segno) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StaffLayout) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StringType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Supports) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *SystemDividers) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.LeftDivider, stage.EmptyPrintObjectStyleAligns_instance)
	__gong__reconstructPointerFromInstance(&reference.RightDivider, stage.EmptyPrintObjectStyleAligns_instance)
	// insertion point for slice of pointers fields
}

func (reference *SystemLayout) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.SystemMargins, stage.SystemMarginss_instance)
	__gong__reconstructPointerFromInstance(&reference.SystemDividers, stage.SystemDividerss_instance)
	// insertion point for slice of pointers fields
}

func (reference *SystemMargins) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TypedText) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *WavyLine) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (accidentaltext *AccidentalText) GongDiff(stage *Stage, accidentaltextOther *AccidentalText) (diffs []string) {
	// insertion point for field diffs
	if accidentaltext.Name != accidentaltextOther.Name {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "Name"))
	}
	if accidentaltext.Smufl != accidentaltextOther.Smufl {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "Smufl"))
	}
	if accidentaltext.Lang != accidentaltextOther.Lang {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "Lang"))
	}
	if accidentaltext.Space != accidentaltextOther.Space {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "Space"))
	}
	if accidentaltext.Justify != accidentaltextOther.Justify {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "Justify"))
	}
	if accidentaltext.DefaultX != accidentaltextOther.DefaultX {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "DefaultX"))
	}
	if accidentaltext.DefaultY != accidentaltextOther.DefaultY {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "DefaultY"))
	}
	if accidentaltext.RelativeX != accidentaltextOther.RelativeX {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "RelativeX"))
	}
	if accidentaltext.RelativeY != accidentaltextOther.RelativeY {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "RelativeY"))
	}
	if accidentaltext.FontFamily != accidentaltextOther.FontFamily {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "FontFamily"))
	}
	if accidentaltext.FontStyle != accidentaltextOther.FontStyle {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "FontStyle"))
	}
	if accidentaltext.FontSize != accidentaltextOther.FontSize {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "FontSize"))
	}
	if accidentaltext.FontWeight != accidentaltextOther.FontWeight {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "FontWeight"))
	}
	if accidentaltext.Color != accidentaltextOther.Color {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "Color"))
	}
	if accidentaltext.Halign != accidentaltextOther.Halign {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "Halign"))
	}
	if accidentaltext.Valign != accidentaltextOther.Valign {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "Valign"))
	}
	if accidentaltext.Underline != accidentaltextOther.Underline {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "Underline"))
	}
	if accidentaltext.Overline != accidentaltextOther.Overline {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "Overline"))
	}
	if accidentaltext.LineThrough != accidentaltextOther.LineThrough {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "LineThrough"))
	}
	if accidentaltext.Rotation != accidentaltextOther.Rotation {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "Rotation"))
	}
	if accidentaltext.LetterSpacing != accidentaltextOther.LetterSpacing {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "LetterSpacing"))
	}
	if accidentaltext.LineHeight != accidentaltextOther.LineHeight {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "LineHeight"))
	}
	if accidentaltext.Dir != accidentaltextOther.Dir {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "Dir"))
	}
	if accidentaltext.Enclosure != accidentaltextOther.Enclosure {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "Enclosure"))
	}
	if accidentaltext.Value != accidentaltextOther.Value {
		diffs = append(diffs, accidentaltext.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (appearance *Appearance) GongDiff(stage *Stage, appearanceOther *Appearance) (diffs []string) {
	// insertion point for field diffs
	if appearance.Name != appearanceOther.Name {
		diffs = append(diffs, appearance.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, appearance, "LineWidth", appearanceOther.LineWidth, appearance.LineWidth); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, appearance, "NoteSize", appearanceOther.NoteSize, appearance.NoteSize); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, appearance, "Distance", appearanceOther.Distance, appearance.Distance); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, appearance, "Glyph", appearanceOther.Glyph, appearance.Glyph); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, appearance, "OtherAppearance", appearanceOther.OtherAppearance, appearance.OtherAppearance); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bookmark *Bookmark) GongDiff(stage *Stage, bookmarkOther *Bookmark) (diffs []string) {
	// insertion point for field diffs
	if bookmark.Name != bookmarkOther.Name {
		diffs = append(diffs, bookmark.GongMarshallField(stage, "Name"))
	}
	if bookmark.Id != bookmarkOther.Id {
		diffs = append(diffs, bookmark.GongMarshallField(stage, "Id"))
	}
	if bookmark.NameAttr != bookmarkOther.NameAttr {
		diffs = append(diffs, bookmark.GongMarshallField(stage, "NameAttr"))
	}
	if bookmark.Element != bookmarkOther.Element {
		diffs = append(diffs, bookmark.GongMarshallField(stage, "Element"))
	}
	if bookmark.Position != bookmarkOther.Position {
		diffs = append(diffs, bookmark.GongMarshallField(stage, "Position"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coda *Coda) GongDiff(stage *Stage, codaOther *Coda) (diffs []string) {
	// insertion point for field diffs
	if coda.Name != codaOther.Name {
		diffs = append(diffs, coda.GongMarshallField(stage, "Name"))
	}
	if coda.Smufl != codaOther.Smufl {
		diffs = append(diffs, coda.GongMarshallField(stage, "Smufl"))
	}
	if coda.DefaultX != codaOther.DefaultX {
		diffs = append(diffs, coda.GongMarshallField(stage, "DefaultX"))
	}
	if coda.DefaultY != codaOther.DefaultY {
		diffs = append(diffs, coda.GongMarshallField(stage, "DefaultY"))
	}
	if coda.RelativeX != codaOther.RelativeX {
		diffs = append(diffs, coda.GongMarshallField(stage, "RelativeX"))
	}
	if coda.RelativeY != codaOther.RelativeY {
		diffs = append(diffs, coda.GongMarshallField(stage, "RelativeY"))
	}
	if coda.FontFamily != codaOther.FontFamily {
		diffs = append(diffs, coda.GongMarshallField(stage, "FontFamily"))
	}
	if coda.FontStyle != codaOther.FontStyle {
		diffs = append(diffs, coda.GongMarshallField(stage, "FontStyle"))
	}
	if coda.FontSize != codaOther.FontSize {
		diffs = append(diffs, coda.GongMarshallField(stage, "FontSize"))
	}
	if coda.FontWeight != codaOther.FontWeight {
		diffs = append(diffs, coda.GongMarshallField(stage, "FontWeight"))
	}
	if coda.Color != codaOther.Color {
		diffs = append(diffs, coda.GongMarshallField(stage, "Color"))
	}
	if coda.Halign != codaOther.Halign {
		diffs = append(diffs, coda.GongMarshallField(stage, "Halign"))
	}
	if coda.Valign != codaOther.Valign {
		diffs = append(diffs, coda.GongMarshallField(stage, "Valign"))
	}
	if coda.Id != codaOther.Id {
		diffs = append(diffs, coda.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (distance *Distance) GongDiff(stage *Stage, distanceOther *Distance) (diffs []string) {
	// insertion point for field diffs
	if distance.Name != distanceOther.Name {
		diffs = append(diffs, distance.GongMarshallField(stage, "Name"))
	}
	if distance.Type != distanceOther.Type {
		diffs = append(diffs, distance.GongMarshallField(stage, "Type"))
	}
	if distance.Value != distanceOther.Value {
		diffs = append(diffs, distance.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (dynamics *Dynamics) GongDiff(stage *Stage, dynamicsOther *Dynamics) (diffs []string) {
	// insertion point for field diffs
	if dynamics.Name != dynamicsOther.Name {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Name"))
	}
	if dynamics.DefaultX != dynamicsOther.DefaultX {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "DefaultX"))
	}
	if dynamics.DefaultY != dynamicsOther.DefaultY {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "DefaultY"))
	}
	if dynamics.RelativeX != dynamicsOther.RelativeX {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "RelativeX"))
	}
	if dynamics.RelativeY != dynamicsOther.RelativeY {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "RelativeY"))
	}
	if dynamics.FontFamily != dynamicsOther.FontFamily {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "FontFamily"))
	}
	if dynamics.FontStyle != dynamicsOther.FontStyle {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "FontStyle"))
	}
	if dynamics.FontSize != dynamicsOther.FontSize {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "FontSize"))
	}
	if dynamics.FontWeight != dynamicsOther.FontWeight {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "FontWeight"))
	}
	if dynamics.Color != dynamicsOther.Color {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Color"))
	}
	if dynamics.Halign != dynamicsOther.Halign {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Halign"))
	}
	if dynamics.Valign != dynamicsOther.Valign {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Valign"))
	}
	if dynamics.Placement != dynamicsOther.Placement {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Placement"))
	}
	if dynamics.Underline != dynamicsOther.Underline {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Underline"))
	}
	if dynamics.Overline != dynamicsOther.Overline {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Overline"))
	}
	if dynamics.LineThrough != dynamicsOther.LineThrough {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "LineThrough"))
	}
	if dynamics.Enclosure != dynamicsOther.Enclosure {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Enclosure"))
	}
	if dynamics.Id != dynamicsOther.Id {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "P", dynamicsOther.P, dynamics.P); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Pp", dynamicsOther.Pp, dynamics.Pp); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Ppp", dynamicsOther.Ppp, dynamics.Ppp); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Pppp", dynamicsOther.Pppp, dynamics.Pppp); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Ppppp", dynamicsOther.Ppppp, dynamics.Ppppp); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Pppppp", dynamicsOther.Pppppp, dynamics.Pppppp); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "F", dynamicsOther.F, dynamics.F); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Ff", dynamicsOther.Ff, dynamics.Ff); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Fff", dynamicsOther.Fff, dynamics.Fff); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Ffff", dynamicsOther.Ffff, dynamics.Ffff); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Fffff", dynamicsOther.Fffff, dynamics.Fffff); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Ffffff", dynamicsOther.Ffffff, dynamics.Ffffff); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Mp", dynamicsOther.Mp, dynamics.Mp); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Mf", dynamicsOther.Mf, dynamics.Mf); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Sf", dynamicsOther.Sf, dynamics.Sf); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Sfp", dynamicsOther.Sfp, dynamics.Sfp); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Sfpp", dynamicsOther.Sfpp, dynamics.Sfpp); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Fp", dynamicsOther.Fp, dynamics.Fp); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Rf", dynamicsOther.Rf, dynamics.Rf); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Rfz", dynamicsOther.Rfz, dynamics.Rfz); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Sfz", dynamicsOther.Sfz, dynamics.Sfz); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Sffz", dynamicsOther.Sffz, dynamics.Sffz); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Fz", dynamicsOther.Fz, dynamics.Fz); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "N", dynamicsOther.N, dynamics.N); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Pf", dynamicsOther.Pf, dynamics.Pf); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Sfzp", dynamicsOther.Sfzp, dynamics.Sfzp); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "OtherDynamics", dynamicsOther.OtherDynamics, dynamics.OtherDynamics); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (empty *Empty) GongDiff(stage *Stage, emptyOther *Empty) (diffs []string) {
	// insertion point for field diffs
	if empty.Name != emptyOther.Name {
		diffs = append(diffs, empty.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (emptyplacement *EmptyPlacement) GongDiff(stage *Stage, emptyplacementOther *EmptyPlacement) (diffs []string) {
	// insertion point for field diffs
	if emptyplacement.Name != emptyplacementOther.Name {
		diffs = append(diffs, emptyplacement.GongMarshallField(stage, "Name"))
	}
	if emptyplacement.DefaultX != emptyplacementOther.DefaultX {
		diffs = append(diffs, emptyplacement.GongMarshallField(stage, "DefaultX"))
	}
	if emptyplacement.DefaultY != emptyplacementOther.DefaultY {
		diffs = append(diffs, emptyplacement.GongMarshallField(stage, "DefaultY"))
	}
	if emptyplacement.RelativeX != emptyplacementOther.RelativeX {
		diffs = append(diffs, emptyplacement.GongMarshallField(stage, "RelativeX"))
	}
	if emptyplacement.RelativeY != emptyplacementOther.RelativeY {
		diffs = append(diffs, emptyplacement.GongMarshallField(stage, "RelativeY"))
	}
	if emptyplacement.FontFamily != emptyplacementOther.FontFamily {
		diffs = append(diffs, emptyplacement.GongMarshallField(stage, "FontFamily"))
	}
	if emptyplacement.FontStyle != emptyplacementOther.FontStyle {
		diffs = append(diffs, emptyplacement.GongMarshallField(stage, "FontStyle"))
	}
	if emptyplacement.FontSize != emptyplacementOther.FontSize {
		diffs = append(diffs, emptyplacement.GongMarshallField(stage, "FontSize"))
	}
	if emptyplacement.FontWeight != emptyplacementOther.FontWeight {
		diffs = append(diffs, emptyplacement.GongMarshallField(stage, "FontWeight"))
	}
	if emptyplacement.Color != emptyplacementOther.Color {
		diffs = append(diffs, emptyplacement.GongMarshallField(stage, "Color"))
	}
	if emptyplacement.Placement != emptyplacementOther.Placement {
		diffs = append(diffs, emptyplacement.GongMarshallField(stage, "Placement"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (emptyplacementsmufl *EmptyPlacementSmufl) GongDiff(stage *Stage, emptyplacementsmuflOther *EmptyPlacementSmufl) (diffs []string) {
	// insertion point for field diffs
	if emptyplacementsmufl.Name != emptyplacementsmuflOther.Name {
		diffs = append(diffs, emptyplacementsmufl.GongMarshallField(stage, "Name"))
	}
	if emptyplacementsmufl.DefaultX != emptyplacementsmuflOther.DefaultX {
		diffs = append(diffs, emptyplacementsmufl.GongMarshallField(stage, "DefaultX"))
	}
	if emptyplacementsmufl.DefaultY != emptyplacementsmuflOther.DefaultY {
		diffs = append(diffs, emptyplacementsmufl.GongMarshallField(stage, "DefaultY"))
	}
	if emptyplacementsmufl.RelativeX != emptyplacementsmuflOther.RelativeX {
		diffs = append(diffs, emptyplacementsmufl.GongMarshallField(stage, "RelativeX"))
	}
	if emptyplacementsmufl.RelativeY != emptyplacementsmuflOther.RelativeY {
		diffs = append(diffs, emptyplacementsmufl.GongMarshallField(stage, "RelativeY"))
	}
	if emptyplacementsmufl.FontFamily != emptyplacementsmuflOther.FontFamily {
		diffs = append(diffs, emptyplacementsmufl.GongMarshallField(stage, "FontFamily"))
	}
	if emptyplacementsmufl.FontStyle != emptyplacementsmuflOther.FontStyle {
		diffs = append(diffs, emptyplacementsmufl.GongMarshallField(stage, "FontStyle"))
	}
	if emptyplacementsmufl.FontSize != emptyplacementsmuflOther.FontSize {
		diffs = append(diffs, emptyplacementsmufl.GongMarshallField(stage, "FontSize"))
	}
	if emptyplacementsmufl.FontWeight != emptyplacementsmuflOther.FontWeight {
		diffs = append(diffs, emptyplacementsmufl.GongMarshallField(stage, "FontWeight"))
	}
	if emptyplacementsmufl.Color != emptyplacementsmuflOther.Color {
		diffs = append(diffs, emptyplacementsmufl.GongMarshallField(stage, "Color"))
	}
	if emptyplacementsmufl.Placement != emptyplacementsmuflOther.Placement {
		diffs = append(diffs, emptyplacementsmufl.GongMarshallField(stage, "Placement"))
	}
	if emptyplacementsmufl.Smufl != emptyplacementsmuflOther.Smufl {
		diffs = append(diffs, emptyplacementsmufl.GongMarshallField(stage, "Smufl"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) GongDiff(stage *Stage, emptyprintobjectstylealignOther *EmptyPrintObjectStyleAlign) (diffs []string) {
	// insertion point for field diffs
	if emptyprintobjectstylealign.Name != emptyprintobjectstylealignOther.Name {
		diffs = append(diffs, emptyprintobjectstylealign.GongMarshallField(stage, "Name"))
	}
	if emptyprintobjectstylealign.PrintObject != emptyprintobjectstylealignOther.PrintObject {
		diffs = append(diffs, emptyprintobjectstylealign.GongMarshallField(stage, "PrintObject"))
	}
	if emptyprintobjectstylealign.DefaultX != emptyprintobjectstylealignOther.DefaultX {
		diffs = append(diffs, emptyprintobjectstylealign.GongMarshallField(stage, "DefaultX"))
	}
	if emptyprintobjectstylealign.DefaultY != emptyprintobjectstylealignOther.DefaultY {
		diffs = append(diffs, emptyprintobjectstylealign.GongMarshallField(stage, "DefaultY"))
	}
	if emptyprintobjectstylealign.RelativeX != emptyprintobjectstylealignOther.RelativeX {
		diffs = append(diffs, emptyprintobjectstylealign.GongMarshallField(stage, "RelativeX"))
	}
	if emptyprintobjectstylealign.RelativeY != emptyprintobjectstylealignOther.RelativeY {
		diffs = append(diffs, emptyprintobjectstylealign.GongMarshallField(stage, "RelativeY"))
	}
	if emptyprintobjectstylealign.FontFamily != emptyprintobjectstylealignOther.FontFamily {
		diffs = append(diffs, emptyprintobjectstylealign.GongMarshallField(stage, "FontFamily"))
	}
	if emptyprintobjectstylealign.FontStyle != emptyprintobjectstylealignOther.FontStyle {
		diffs = append(diffs, emptyprintobjectstylealign.GongMarshallField(stage, "FontStyle"))
	}
	if emptyprintobjectstylealign.FontSize != emptyprintobjectstylealignOther.FontSize {
		diffs = append(diffs, emptyprintobjectstylealign.GongMarshallField(stage, "FontSize"))
	}
	if emptyprintobjectstylealign.FontWeight != emptyprintobjectstylealignOther.FontWeight {
		diffs = append(diffs, emptyprintobjectstylealign.GongMarshallField(stage, "FontWeight"))
	}
	if emptyprintobjectstylealign.Color != emptyprintobjectstylealignOther.Color {
		diffs = append(diffs, emptyprintobjectstylealign.GongMarshallField(stage, "Color"))
	}
	if emptyprintobjectstylealign.Halign != emptyprintobjectstylealignOther.Halign {
		diffs = append(diffs, emptyprintobjectstylealign.GongMarshallField(stage, "Halign"))
	}
	if emptyprintobjectstylealign.Valign != emptyprintobjectstylealignOther.Valign {
		diffs = append(diffs, emptyprintobjectstylealign.GongMarshallField(stage, "Valign"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (emptyprintstyle *EmptyPrintStyle) GongDiff(stage *Stage, emptyprintstyleOther *EmptyPrintStyle) (diffs []string) {
	// insertion point for field diffs
	if emptyprintstyle.Name != emptyprintstyleOther.Name {
		diffs = append(diffs, emptyprintstyle.GongMarshallField(stage, "Name"))
	}
	if emptyprintstyle.DefaultX != emptyprintstyleOther.DefaultX {
		diffs = append(diffs, emptyprintstyle.GongMarshallField(stage, "DefaultX"))
	}
	if emptyprintstyle.DefaultY != emptyprintstyleOther.DefaultY {
		diffs = append(diffs, emptyprintstyle.GongMarshallField(stage, "DefaultY"))
	}
	if emptyprintstyle.RelativeX != emptyprintstyleOther.RelativeX {
		diffs = append(diffs, emptyprintstyle.GongMarshallField(stage, "RelativeX"))
	}
	if emptyprintstyle.RelativeY != emptyprintstyleOther.RelativeY {
		diffs = append(diffs, emptyprintstyle.GongMarshallField(stage, "RelativeY"))
	}
	if emptyprintstyle.FontFamily != emptyprintstyleOther.FontFamily {
		diffs = append(diffs, emptyprintstyle.GongMarshallField(stage, "FontFamily"))
	}
	if emptyprintstyle.FontStyle != emptyprintstyleOther.FontStyle {
		diffs = append(diffs, emptyprintstyle.GongMarshallField(stage, "FontStyle"))
	}
	if emptyprintstyle.FontSize != emptyprintstyleOther.FontSize {
		diffs = append(diffs, emptyprintstyle.GongMarshallField(stage, "FontSize"))
	}
	if emptyprintstyle.FontWeight != emptyprintstyleOther.FontWeight {
		diffs = append(diffs, emptyprintstyle.GongMarshallField(stage, "FontWeight"))
	}
	if emptyprintstyle.Color != emptyprintstyleOther.Color {
		diffs = append(diffs, emptyprintstyle.GongMarshallField(stage, "Color"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (emptyprintstylealign *EmptyPrintStyleAlign) GongDiff(stage *Stage, emptyprintstylealignOther *EmptyPrintStyleAlign) (diffs []string) {
	// insertion point for field diffs
	if emptyprintstylealign.Name != emptyprintstylealignOther.Name {
		diffs = append(diffs, emptyprintstylealign.GongMarshallField(stage, "Name"))
	}
	if emptyprintstylealign.DefaultX != emptyprintstylealignOther.DefaultX {
		diffs = append(diffs, emptyprintstylealign.GongMarshallField(stage, "DefaultX"))
	}
	if emptyprintstylealign.DefaultY != emptyprintstylealignOther.DefaultY {
		diffs = append(diffs, emptyprintstylealign.GongMarshallField(stage, "DefaultY"))
	}
	if emptyprintstylealign.RelativeX != emptyprintstylealignOther.RelativeX {
		diffs = append(diffs, emptyprintstylealign.GongMarshallField(stage, "RelativeX"))
	}
	if emptyprintstylealign.RelativeY != emptyprintstylealignOther.RelativeY {
		diffs = append(diffs, emptyprintstylealign.GongMarshallField(stage, "RelativeY"))
	}
	if emptyprintstylealign.FontFamily != emptyprintstylealignOther.FontFamily {
		diffs = append(diffs, emptyprintstylealign.GongMarshallField(stage, "FontFamily"))
	}
	if emptyprintstylealign.FontStyle != emptyprintstylealignOther.FontStyle {
		diffs = append(diffs, emptyprintstylealign.GongMarshallField(stage, "FontStyle"))
	}
	if emptyprintstylealign.FontSize != emptyprintstylealignOther.FontSize {
		diffs = append(diffs, emptyprintstylealign.GongMarshallField(stage, "FontSize"))
	}
	if emptyprintstylealign.FontWeight != emptyprintstylealignOther.FontWeight {
		diffs = append(diffs, emptyprintstylealign.GongMarshallField(stage, "FontWeight"))
	}
	if emptyprintstylealign.Color != emptyprintstylealignOther.Color {
		diffs = append(diffs, emptyprintstylealign.GongMarshallField(stage, "Color"))
	}
	if emptyprintstylealign.Halign != emptyprintstylealignOther.Halign {
		diffs = append(diffs, emptyprintstylealign.GongMarshallField(stage, "Halign"))
	}
	if emptyprintstylealign.Valign != emptyprintstylealignOther.Valign {
		diffs = append(diffs, emptyprintstylealign.GongMarshallField(stage, "Valign"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (emptyprintstylealignid *EmptyPrintStyleAlignId) GongDiff(stage *Stage, emptyprintstylealignidOther *EmptyPrintStyleAlignId) (diffs []string) {
	// insertion point for field diffs
	if emptyprintstylealignid.Name != emptyprintstylealignidOther.Name {
		diffs = append(diffs, emptyprintstylealignid.GongMarshallField(stage, "Name"))
	}
	if emptyprintstylealignid.DefaultX != emptyprintstylealignidOther.DefaultX {
		diffs = append(diffs, emptyprintstylealignid.GongMarshallField(stage, "DefaultX"))
	}
	if emptyprintstylealignid.DefaultY != emptyprintstylealignidOther.DefaultY {
		diffs = append(diffs, emptyprintstylealignid.GongMarshallField(stage, "DefaultY"))
	}
	if emptyprintstylealignid.RelativeX != emptyprintstylealignidOther.RelativeX {
		diffs = append(diffs, emptyprintstylealignid.GongMarshallField(stage, "RelativeX"))
	}
	if emptyprintstylealignid.RelativeY != emptyprintstylealignidOther.RelativeY {
		diffs = append(diffs, emptyprintstylealignid.GongMarshallField(stage, "RelativeY"))
	}
	if emptyprintstylealignid.FontFamily != emptyprintstylealignidOther.FontFamily {
		diffs = append(diffs, emptyprintstylealignid.GongMarshallField(stage, "FontFamily"))
	}
	if emptyprintstylealignid.FontStyle != emptyprintstylealignidOther.FontStyle {
		diffs = append(diffs, emptyprintstylealignid.GongMarshallField(stage, "FontStyle"))
	}
	if emptyprintstylealignid.FontSize != emptyprintstylealignidOther.FontSize {
		diffs = append(diffs, emptyprintstylealignid.GongMarshallField(stage, "FontSize"))
	}
	if emptyprintstylealignid.FontWeight != emptyprintstylealignidOther.FontWeight {
		diffs = append(diffs, emptyprintstylealignid.GongMarshallField(stage, "FontWeight"))
	}
	if emptyprintstylealignid.Color != emptyprintstylealignidOther.Color {
		diffs = append(diffs, emptyprintstylealignid.GongMarshallField(stage, "Color"))
	}
	if emptyprintstylealignid.Halign != emptyprintstylealignidOther.Halign {
		diffs = append(diffs, emptyprintstylealignid.GongMarshallField(stage, "Halign"))
	}
	if emptyprintstylealignid.Valign != emptyprintstylealignidOther.Valign {
		diffs = append(diffs, emptyprintstylealignid.GongMarshallField(stage, "Valign"))
	}
	if emptyprintstylealignid.Id != emptyprintstylealignidOther.Id {
		diffs = append(diffs, emptyprintstylealignid.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (emptytrillsound *EmptyTrillSound) GongDiff(stage *Stage, emptytrillsoundOther *EmptyTrillSound) (diffs []string) {
	// insertion point for field diffs
	if emptytrillsound.Name != emptytrillsoundOther.Name {
		diffs = append(diffs, emptytrillsound.GongMarshallField(stage, "Name"))
	}
	if emptytrillsound.DefaultX != emptytrillsoundOther.DefaultX {
		diffs = append(diffs, emptytrillsound.GongMarshallField(stage, "DefaultX"))
	}
	if emptytrillsound.DefaultY != emptytrillsoundOther.DefaultY {
		diffs = append(diffs, emptytrillsound.GongMarshallField(stage, "DefaultY"))
	}
	if emptytrillsound.RelativeX != emptytrillsoundOther.RelativeX {
		diffs = append(diffs, emptytrillsound.GongMarshallField(stage, "RelativeX"))
	}
	if emptytrillsound.RelativeY != emptytrillsoundOther.RelativeY {
		diffs = append(diffs, emptytrillsound.GongMarshallField(stage, "RelativeY"))
	}
	if emptytrillsound.FontFamily != emptytrillsoundOther.FontFamily {
		diffs = append(diffs, emptytrillsound.GongMarshallField(stage, "FontFamily"))
	}
	if emptytrillsound.FontStyle != emptytrillsoundOther.FontStyle {
		diffs = append(diffs, emptytrillsound.GongMarshallField(stage, "FontStyle"))
	}
	if emptytrillsound.FontSize != emptytrillsoundOther.FontSize {
		diffs = append(diffs, emptytrillsound.GongMarshallField(stage, "FontSize"))
	}
	if emptytrillsound.FontWeight != emptytrillsoundOther.FontWeight {
		diffs = append(diffs, emptytrillsound.GongMarshallField(stage, "FontWeight"))
	}
	if emptytrillsound.Color != emptytrillsoundOther.Color {
		diffs = append(diffs, emptytrillsound.GongMarshallField(stage, "Color"))
	}
	if emptytrillsound.Placement != emptytrillsoundOther.Placement {
		diffs = append(diffs, emptytrillsound.GongMarshallField(stage, "Placement"))
	}
	if emptytrillsound.StartNote != emptytrillsoundOther.StartNote {
		diffs = append(diffs, emptytrillsound.GongMarshallField(stage, "StartNote"))
	}
	if emptytrillsound.TrillStep != emptytrillsoundOther.TrillStep {
		diffs = append(diffs, emptytrillsound.GongMarshallField(stage, "TrillStep"))
	}
	if emptytrillsound.TwoNoteTurn != emptytrillsoundOther.TwoNoteTurn {
		diffs = append(diffs, emptytrillsound.GongMarshallField(stage, "TwoNoteTurn"))
	}
	if emptytrillsound.Accelerate != emptytrillsoundOther.Accelerate {
		diffs = append(diffs, emptytrillsound.GongMarshallField(stage, "Accelerate"))
	}
	if emptytrillsound.Beats != emptytrillsoundOther.Beats {
		diffs = append(diffs, emptytrillsound.GongMarshallField(stage, "Beats"))
	}
	if emptytrillsound.SecondBeat != emptytrillsoundOther.SecondBeat {
		diffs = append(diffs, emptytrillsound.GongMarshallField(stage, "SecondBeat"))
	}
	if emptytrillsound.LastBeat != emptytrillsoundOther.LastBeat {
		diffs = append(diffs, emptytrillsound.GongMarshallField(stage, "LastBeat"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (encoding *Encoding) GongDiff(stage *Stage, encodingOther *Encoding) (diffs []string) {
	// insertion point for field diffs
	if encoding.Name != encodingOther.Name {
		diffs = append(diffs, encoding.GongMarshallField(stage, "Name"))
	}
	if encoding.EncodingDate != encodingOther.EncodingDate {
		diffs = append(diffs, encoding.GongMarshallField(stage, "EncodingDate"))
	}
	if ops := __gong__diffSliceOfPointers(stage, encoding, "Encoder", encodingOther.Encoder, encoding.Encoder); ops != "" {
		diffs = append(diffs, ops)
	}
	if encoding.Software != encodingOther.Software {
		diffs = append(diffs, encoding.GongMarshallField(stage, "Software"))
	}
	if encoding.EncodingDescription != encodingOther.EncodingDescription {
		diffs = append(diffs, encoding.GongMarshallField(stage, "EncodingDescription"))
	}
	if ops := __gong__diffSliceOfPointers(stage, encoding, "Supports", encodingOther.Supports, encoding.Supports); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (fermata *Fermata) GongDiff(stage *Stage, fermataOther *Fermata) (diffs []string) {
	// insertion point for field diffs
	if fermata.Name != fermataOther.Name {
		diffs = append(diffs, fermata.GongMarshallField(stage, "Name"))
	}
	if fermata.Type != fermataOther.Type {
		diffs = append(diffs, fermata.GongMarshallField(stage, "Type"))
	}
	if fermata.DefaultX != fermataOther.DefaultX {
		diffs = append(diffs, fermata.GongMarshallField(stage, "DefaultX"))
	}
	if fermata.DefaultY != fermataOther.DefaultY {
		diffs = append(diffs, fermata.GongMarshallField(stage, "DefaultY"))
	}
	if fermata.RelativeX != fermataOther.RelativeX {
		diffs = append(diffs, fermata.GongMarshallField(stage, "RelativeX"))
	}
	if fermata.RelativeY != fermataOther.RelativeY {
		diffs = append(diffs, fermata.GongMarshallField(stage, "RelativeY"))
	}
	if fermata.FontFamily != fermataOther.FontFamily {
		diffs = append(diffs, fermata.GongMarshallField(stage, "FontFamily"))
	}
	if fermata.FontStyle != fermataOther.FontStyle {
		diffs = append(diffs, fermata.GongMarshallField(stage, "FontStyle"))
	}
	if fermata.FontSize != fermataOther.FontSize {
		diffs = append(diffs, fermata.GongMarshallField(stage, "FontSize"))
	}
	if fermata.FontWeight != fermataOther.FontWeight {
		diffs = append(diffs, fermata.GongMarshallField(stage, "FontWeight"))
	}
	if fermata.Color != fermataOther.Color {
		diffs = append(diffs, fermata.GongMarshallField(stage, "Color"))
	}
	if fermata.Id != fermataOther.Id {
		diffs = append(diffs, fermata.GongMarshallField(stage, "Id"))
	}
	if fermata.Value != fermataOther.Value {
		diffs = append(diffs, fermata.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (fingering *Fingering) GongDiff(stage *Stage, fingeringOther *Fingering) (diffs []string) {
	// insertion point for field diffs
	if fingering.Name != fingeringOther.Name {
		diffs = append(diffs, fingering.GongMarshallField(stage, "Name"))
	}
	if fingering.Substitution != fingeringOther.Substitution {
		diffs = append(diffs, fingering.GongMarshallField(stage, "Substitution"))
	}
	if fingering.Alternate != fingeringOther.Alternate {
		diffs = append(diffs, fingering.GongMarshallField(stage, "Alternate"))
	}
	if fingering.DefaultX != fingeringOther.DefaultX {
		diffs = append(diffs, fingering.GongMarshallField(stage, "DefaultX"))
	}
	if fingering.DefaultY != fingeringOther.DefaultY {
		diffs = append(diffs, fingering.GongMarshallField(stage, "DefaultY"))
	}
	if fingering.RelativeX != fingeringOther.RelativeX {
		diffs = append(diffs, fingering.GongMarshallField(stage, "RelativeX"))
	}
	if fingering.RelativeY != fingeringOther.RelativeY {
		diffs = append(diffs, fingering.GongMarshallField(stage, "RelativeY"))
	}
	if fingering.FontFamily != fingeringOther.FontFamily {
		diffs = append(diffs, fingering.GongMarshallField(stage, "FontFamily"))
	}
	if fingering.FontStyle != fingeringOther.FontStyle {
		diffs = append(diffs, fingering.GongMarshallField(stage, "FontStyle"))
	}
	if fingering.FontSize != fingeringOther.FontSize {
		diffs = append(diffs, fingering.GongMarshallField(stage, "FontSize"))
	}
	if fingering.FontWeight != fingeringOther.FontWeight {
		diffs = append(diffs, fingering.GongMarshallField(stage, "FontWeight"))
	}
	if fingering.Color != fingeringOther.Color {
		diffs = append(diffs, fingering.GongMarshallField(stage, "Color"))
	}
	if fingering.Placement != fingeringOther.Placement {
		diffs = append(diffs, fingering.GongMarshallField(stage, "Placement"))
	}
	if fingering.Value != fingeringOther.Value {
		diffs = append(diffs, fingering.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (formattedsymbol *FormattedSymbol) GongDiff(stage *Stage, formattedsymbolOther *FormattedSymbol) (diffs []string) {
	// insertion point for field diffs
	if formattedsymbol.Name != formattedsymbolOther.Name {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "Name"))
	}
	if formattedsymbol.Justify != formattedsymbolOther.Justify {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "Justify"))
	}
	if formattedsymbol.DefaultX != formattedsymbolOther.DefaultX {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "DefaultX"))
	}
	if formattedsymbol.DefaultY != formattedsymbolOther.DefaultY {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "DefaultY"))
	}
	if formattedsymbol.RelativeX != formattedsymbolOther.RelativeX {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "RelativeX"))
	}
	if formattedsymbol.RelativeY != formattedsymbolOther.RelativeY {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "RelativeY"))
	}
	if formattedsymbol.FontFamily != formattedsymbolOther.FontFamily {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "FontFamily"))
	}
	if formattedsymbol.FontStyle != formattedsymbolOther.FontStyle {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "FontStyle"))
	}
	if formattedsymbol.FontSize != formattedsymbolOther.FontSize {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "FontSize"))
	}
	if formattedsymbol.FontWeight != formattedsymbolOther.FontWeight {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "FontWeight"))
	}
	if formattedsymbol.Color != formattedsymbolOther.Color {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "Color"))
	}
	if formattedsymbol.Halign != formattedsymbolOther.Halign {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "Halign"))
	}
	if formattedsymbol.Valign != formattedsymbolOther.Valign {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "Valign"))
	}
	if formattedsymbol.Underline != formattedsymbolOther.Underline {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "Underline"))
	}
	if formattedsymbol.Overline != formattedsymbolOther.Overline {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "Overline"))
	}
	if formattedsymbol.LineThrough != formattedsymbolOther.LineThrough {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "LineThrough"))
	}
	if formattedsymbol.Rotation != formattedsymbolOther.Rotation {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "Rotation"))
	}
	if formattedsymbol.LetterSpacing != formattedsymbolOther.LetterSpacing {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "LetterSpacing"))
	}
	if formattedsymbol.LineHeight != formattedsymbolOther.LineHeight {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "LineHeight"))
	}
	if formattedsymbol.Dir != formattedsymbolOther.Dir {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "Dir"))
	}
	if formattedsymbol.Enclosure != formattedsymbolOther.Enclosure {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "Enclosure"))
	}
	if formattedsymbol.Value != formattedsymbolOther.Value {
		diffs = append(diffs, formattedsymbol.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (formattedsymbolid *FormattedSymbolId) GongDiff(stage *Stage, formattedsymbolidOther *FormattedSymbolId) (diffs []string) {
	// insertion point for field diffs
	if formattedsymbolid.Name != formattedsymbolidOther.Name {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "Name"))
	}
	if formattedsymbolid.Justify != formattedsymbolidOther.Justify {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "Justify"))
	}
	if formattedsymbolid.DefaultX != formattedsymbolidOther.DefaultX {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "DefaultX"))
	}
	if formattedsymbolid.DefaultY != formattedsymbolidOther.DefaultY {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "DefaultY"))
	}
	if formattedsymbolid.RelativeX != formattedsymbolidOther.RelativeX {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "RelativeX"))
	}
	if formattedsymbolid.RelativeY != formattedsymbolidOther.RelativeY {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "RelativeY"))
	}
	if formattedsymbolid.FontFamily != formattedsymbolidOther.FontFamily {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "FontFamily"))
	}
	if formattedsymbolid.FontStyle != formattedsymbolidOther.FontStyle {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "FontStyle"))
	}
	if formattedsymbolid.FontSize != formattedsymbolidOther.FontSize {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "FontSize"))
	}
	if formattedsymbolid.FontWeight != formattedsymbolidOther.FontWeight {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "FontWeight"))
	}
	if formattedsymbolid.Color != formattedsymbolidOther.Color {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "Color"))
	}
	if formattedsymbolid.Halign != formattedsymbolidOther.Halign {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "Halign"))
	}
	if formattedsymbolid.Valign != formattedsymbolidOther.Valign {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "Valign"))
	}
	if formattedsymbolid.Underline != formattedsymbolidOther.Underline {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "Underline"))
	}
	if formattedsymbolid.Overline != formattedsymbolidOther.Overline {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "Overline"))
	}
	if formattedsymbolid.LineThrough != formattedsymbolidOther.LineThrough {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "LineThrough"))
	}
	if formattedsymbolid.Rotation != formattedsymbolidOther.Rotation {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "Rotation"))
	}
	if formattedsymbolid.LetterSpacing != formattedsymbolidOther.LetterSpacing {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "LetterSpacing"))
	}
	if formattedsymbolid.LineHeight != formattedsymbolidOther.LineHeight {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "LineHeight"))
	}
	if formattedsymbolid.Dir != formattedsymbolidOther.Dir {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "Dir"))
	}
	if formattedsymbolid.Enclosure != formattedsymbolidOther.Enclosure {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "Enclosure"))
	}
	if formattedsymbolid.Id != formattedsymbolidOther.Id {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "Id"))
	}
	if formattedsymbolid.Value != formattedsymbolidOther.Value {
		diffs = append(diffs, formattedsymbolid.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (formattedtext *FormattedText) GongDiff(stage *Stage, formattedtextOther *FormattedText) (diffs []string) {
	// insertion point for field diffs
	if formattedtext.Name != formattedtextOther.Name {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "Name"))
	}
	if formattedtext.Lang != formattedtextOther.Lang {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "Lang"))
	}
	if formattedtext.Space != formattedtextOther.Space {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "Space"))
	}
	if formattedtext.Justify != formattedtextOther.Justify {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "Justify"))
	}
	if formattedtext.DefaultX != formattedtextOther.DefaultX {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "DefaultX"))
	}
	if formattedtext.DefaultY != formattedtextOther.DefaultY {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "DefaultY"))
	}
	if formattedtext.RelativeX != formattedtextOther.RelativeX {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "RelativeX"))
	}
	if formattedtext.RelativeY != formattedtextOther.RelativeY {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "RelativeY"))
	}
	if formattedtext.FontFamily != formattedtextOther.FontFamily {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "FontFamily"))
	}
	if formattedtext.FontStyle != formattedtextOther.FontStyle {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "FontStyle"))
	}
	if formattedtext.FontSize != formattedtextOther.FontSize {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "FontSize"))
	}
	if formattedtext.FontWeight != formattedtextOther.FontWeight {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "FontWeight"))
	}
	if formattedtext.Color != formattedtextOther.Color {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "Color"))
	}
	if formattedtext.Halign != formattedtextOther.Halign {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "Halign"))
	}
	if formattedtext.Valign != formattedtextOther.Valign {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "Valign"))
	}
	if formattedtext.Underline != formattedtextOther.Underline {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "Underline"))
	}
	if formattedtext.Overline != formattedtextOther.Overline {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "Overline"))
	}
	if formattedtext.LineThrough != formattedtextOther.LineThrough {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "LineThrough"))
	}
	if formattedtext.Rotation != formattedtextOther.Rotation {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "Rotation"))
	}
	if formattedtext.LetterSpacing != formattedtextOther.LetterSpacing {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "LetterSpacing"))
	}
	if formattedtext.LineHeight != formattedtextOther.LineHeight {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "LineHeight"))
	}
	if formattedtext.Dir != formattedtextOther.Dir {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "Dir"))
	}
	if formattedtext.Enclosure != formattedtextOther.Enclosure {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "Enclosure"))
	}
	if formattedtext.Value != formattedtextOther.Value {
		diffs = append(diffs, formattedtext.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (formattedtextid *FormattedTextId) GongDiff(stage *Stage, formattedtextidOther *FormattedTextId) (diffs []string) {
	// insertion point for field diffs
	if formattedtextid.Name != formattedtextidOther.Name {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "Name"))
	}
	if formattedtextid.Lang != formattedtextidOther.Lang {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "Lang"))
	}
	if formattedtextid.Space != formattedtextidOther.Space {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "Space"))
	}
	if formattedtextid.Justify != formattedtextidOther.Justify {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "Justify"))
	}
	if formattedtextid.DefaultX != formattedtextidOther.DefaultX {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "DefaultX"))
	}
	if formattedtextid.DefaultY != formattedtextidOther.DefaultY {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "DefaultY"))
	}
	if formattedtextid.RelativeX != formattedtextidOther.RelativeX {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "RelativeX"))
	}
	if formattedtextid.RelativeY != formattedtextidOther.RelativeY {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "RelativeY"))
	}
	if formattedtextid.FontFamily != formattedtextidOther.FontFamily {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "FontFamily"))
	}
	if formattedtextid.FontStyle != formattedtextidOther.FontStyle {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "FontStyle"))
	}
	if formattedtextid.FontSize != formattedtextidOther.FontSize {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "FontSize"))
	}
	if formattedtextid.FontWeight != formattedtextidOther.FontWeight {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "FontWeight"))
	}
	if formattedtextid.Color != formattedtextidOther.Color {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "Color"))
	}
	if formattedtextid.Halign != formattedtextidOther.Halign {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "Halign"))
	}
	if formattedtextid.Valign != formattedtextidOther.Valign {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "Valign"))
	}
	if formattedtextid.Underline != formattedtextidOther.Underline {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "Underline"))
	}
	if formattedtextid.Overline != formattedtextidOther.Overline {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "Overline"))
	}
	if formattedtextid.LineThrough != formattedtextidOther.LineThrough {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "LineThrough"))
	}
	if formattedtextid.Rotation != formattedtextidOther.Rotation {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "Rotation"))
	}
	if formattedtextid.LetterSpacing != formattedtextidOther.LetterSpacing {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "LetterSpacing"))
	}
	if formattedtextid.LineHeight != formattedtextidOther.LineHeight {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "LineHeight"))
	}
	if formattedtextid.Dir != formattedtextidOther.Dir {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "Dir"))
	}
	if formattedtextid.Enclosure != formattedtextidOther.Enclosure {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "Enclosure"))
	}
	if formattedtextid.Id != formattedtextidOther.Id {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "Id"))
	}
	if formattedtextid.Value != formattedtextidOther.Value {
		diffs = append(diffs, formattedtextid.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (fret *Fret) GongDiff(stage *Stage, fretOther *Fret) (diffs []string) {
	// insertion point for field diffs
	if fret.Name != fretOther.Name {
		diffs = append(diffs, fret.GongMarshallField(stage, "Name"))
	}
	if fret.FontFamily != fretOther.FontFamily {
		diffs = append(diffs, fret.GongMarshallField(stage, "FontFamily"))
	}
	if fret.FontStyle != fretOther.FontStyle {
		diffs = append(diffs, fret.GongMarshallField(stage, "FontStyle"))
	}
	if fret.FontSize != fretOther.FontSize {
		diffs = append(diffs, fret.GongMarshallField(stage, "FontSize"))
	}
	if fret.FontWeight != fretOther.FontWeight {
		diffs = append(diffs, fret.GongMarshallField(stage, "FontWeight"))
	}
	if fret.Color != fretOther.Color {
		diffs = append(diffs, fret.GongMarshallField(stage, "Color"))
	}
	if fret.Value != fretOther.Value {
		diffs = append(diffs, fret.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (glyph *Glyph) GongDiff(stage *Stage, glyphOther *Glyph) (diffs []string) {
	// insertion point for field diffs
	if glyph.Name != glyphOther.Name {
		diffs = append(diffs, glyph.GongMarshallField(stage, "Name"))
	}
	if glyph.Type != glyphOther.Type {
		diffs = append(diffs, glyph.GongMarshallField(stage, "Type"))
	}
	if glyph.Value != glyphOther.Value {
		diffs = append(diffs, glyph.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (horizontalturn *HorizontalTurn) GongDiff(stage *Stage, horizontalturnOther *HorizontalTurn) (diffs []string) {
	// insertion point for field diffs
	if horizontalturn.Name != horizontalturnOther.Name {
		diffs = append(diffs, horizontalturn.GongMarshallField(stage, "Name"))
	}
	if horizontalturn.Slash != horizontalturnOther.Slash {
		diffs = append(diffs, horizontalturn.GongMarshallField(stage, "Slash"))
	}
	if horizontalturn.DefaultX != horizontalturnOther.DefaultX {
		diffs = append(diffs, horizontalturn.GongMarshallField(stage, "DefaultX"))
	}
	if horizontalturn.DefaultY != horizontalturnOther.DefaultY {
		diffs = append(diffs, horizontalturn.GongMarshallField(stage, "DefaultY"))
	}
	if horizontalturn.RelativeX != horizontalturnOther.RelativeX {
		diffs = append(diffs, horizontalturn.GongMarshallField(stage, "RelativeX"))
	}
	if horizontalturn.RelativeY != horizontalturnOther.RelativeY {
		diffs = append(diffs, horizontalturn.GongMarshallField(stage, "RelativeY"))
	}
	if horizontalturn.FontFamily != horizontalturnOther.FontFamily {
		diffs = append(diffs, horizontalturn.GongMarshallField(stage, "FontFamily"))
	}
	if horizontalturn.FontStyle != horizontalturnOther.FontStyle {
		diffs = append(diffs, horizontalturn.GongMarshallField(stage, "FontStyle"))
	}
	if horizontalturn.FontSize != horizontalturnOther.FontSize {
		diffs = append(diffs, horizontalturn.GongMarshallField(stage, "FontSize"))
	}
	if horizontalturn.FontWeight != horizontalturnOther.FontWeight {
		diffs = append(diffs, horizontalturn.GongMarshallField(stage, "FontWeight"))
	}
	if horizontalturn.Color != horizontalturnOther.Color {
		diffs = append(diffs, horizontalturn.GongMarshallField(stage, "Color"))
	}
	if horizontalturn.Placement != horizontalturnOther.Placement {
		diffs = append(diffs, horizontalturn.GongMarshallField(stage, "Placement"))
	}
	if horizontalturn.StartNote != horizontalturnOther.StartNote {
		diffs = append(diffs, horizontalturn.GongMarshallField(stage, "StartNote"))
	}
	if horizontalturn.TrillStep != horizontalturnOther.TrillStep {
		diffs = append(diffs, horizontalturn.GongMarshallField(stage, "TrillStep"))
	}
	if horizontalturn.TwoNoteTurn != horizontalturnOther.TwoNoteTurn {
		diffs = append(diffs, horizontalturn.GongMarshallField(stage, "TwoNoteTurn"))
	}
	if horizontalturn.Accelerate != horizontalturnOther.Accelerate {
		diffs = append(diffs, horizontalturn.GongMarshallField(stage, "Accelerate"))
	}
	if horizontalturn.Beats != horizontalturnOther.Beats {
		diffs = append(diffs, horizontalturn.GongMarshallField(stage, "Beats"))
	}
	if horizontalturn.SecondBeat != horizontalturnOther.SecondBeat {
		diffs = append(diffs, horizontalturn.GongMarshallField(stage, "SecondBeat"))
	}
	if horizontalturn.LastBeat != horizontalturnOther.LastBeat {
		diffs = append(diffs, horizontalturn.GongMarshallField(stage, "LastBeat"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (identification *Identification) GongDiff(stage *Stage, identificationOther *Identification) (diffs []string) {
	// insertion point for field diffs
	if identification.Name != identificationOther.Name {
		diffs = append(diffs, identification.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, identification, "Creator", identificationOther.Creator, identification.Creator); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, identification, "Rights", identificationOther.Rights, identification.Rights); ops != "" {
		diffs = append(diffs, ops)
	}
	if identification.Encoding != identificationOther.Encoding {
		diffs = append(diffs, identification.GongMarshallField(stage, "Encoding"))
	}
	if identification.Source != identificationOther.Source {
		diffs = append(diffs, identification.GongMarshallField(stage, "Source"))
	}
	if ops := __gong__diffSliceOfPointers(stage, identification, "Relation", identificationOther.Relation, identification.Relation); ops != "" {
		diffs = append(diffs, ops)
	}
	if identification.Miscellaneous != identificationOther.Miscellaneous {
		diffs = append(diffs, identification.GongMarshallField(stage, "Miscellaneous"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (level *Level) GongDiff(stage *Stage, levelOther *Level) (diffs []string) {
	// insertion point for field diffs
	if level.Name != levelOther.Name {
		diffs = append(diffs, level.GongMarshallField(stage, "Name"))
	}
	if level.Reference != levelOther.Reference {
		diffs = append(diffs, level.GongMarshallField(stage, "Reference"))
	}
	if level.Type != levelOther.Type {
		diffs = append(diffs, level.GongMarshallField(stage, "Type"))
	}
	if level.Parentheses != levelOther.Parentheses {
		diffs = append(diffs, level.GongMarshallField(stage, "Parentheses"))
	}
	if level.Bracket != levelOther.Bracket {
		diffs = append(diffs, level.GongMarshallField(stage, "Bracket"))
	}
	if level.Size != levelOther.Size {
		diffs = append(diffs, level.GongMarshallField(stage, "Size"))
	}
	if level.Value != levelOther.Value {
		diffs = append(diffs, level.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (linewidth *LineWidth) GongDiff(stage *Stage, linewidthOther *LineWidth) (diffs []string) {
	// insertion point for field diffs
	if linewidth.Name != linewidthOther.Name {
		diffs = append(diffs, linewidth.GongMarshallField(stage, "Name"))
	}
	if linewidth.Type != linewidthOther.Type {
		diffs = append(diffs, linewidth.GongMarshallField(stage, "Type"))
	}
	if linewidth.Value != linewidthOther.Value {
		diffs = append(diffs, linewidth.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (link *Link) GongDiff(stage *Stage, linkOther *Link) (diffs []string) {
	// insertion point for field diffs
	if link.Name != linkOther.Name {
		diffs = append(diffs, link.GongMarshallField(stage, "Name"))
	}
	if link.NameAttr != linkOther.NameAttr {
		diffs = append(diffs, link.GongMarshallField(stage, "NameAttr"))
	}
	if link.Href != linkOther.Href {
		diffs = append(diffs, link.GongMarshallField(stage, "Href"))
	}
	if link.Role != linkOther.Role {
		diffs = append(diffs, link.GongMarshallField(stage, "Role"))
	}
	if link.Title != linkOther.Title {
		diffs = append(diffs, link.GongMarshallField(stage, "Title"))
	}
	if link.Element != linkOther.Element {
		diffs = append(diffs, link.GongMarshallField(stage, "Element"))
	}
	if link.Position != linkOther.Position {
		diffs = append(diffs, link.GongMarshallField(stage, "Position"))
	}
	if link.DefaultX != linkOther.DefaultX {
		diffs = append(diffs, link.GongMarshallField(stage, "DefaultX"))
	}
	if link.DefaultY != linkOther.DefaultY {
		diffs = append(diffs, link.GongMarshallField(stage, "DefaultY"))
	}
	if link.RelativeX != linkOther.RelativeX {
		diffs = append(diffs, link.GongMarshallField(stage, "RelativeX"))
	}
	if link.RelativeY != linkOther.RelativeY {
		diffs = append(diffs, link.GongMarshallField(stage, "RelativeY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (measurelayout *MeasureLayout) GongDiff(stage *Stage, measurelayoutOther *MeasureLayout) (diffs []string) {
	// insertion point for field diffs
	if measurelayout.Name != measurelayoutOther.Name {
		diffs = append(diffs, measurelayout.GongMarshallField(stage, "Name"))
	}
	if measurelayout.MeasureDistance != measurelayoutOther.MeasureDistance {
		diffs = append(diffs, measurelayout.GongMarshallField(stage, "MeasureDistance"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (mididevice *MidiDevice) GongDiff(stage *Stage, midideviceOther *MidiDevice) (diffs []string) {
	// insertion point for field diffs
	if mididevice.Name != midideviceOther.Name {
		diffs = append(diffs, mididevice.GongMarshallField(stage, "Name"))
	}
	if mididevice.Port != midideviceOther.Port {
		diffs = append(diffs, mididevice.GongMarshallField(stage, "Port"))
	}
	if mididevice.Id != midideviceOther.Id {
		diffs = append(diffs, mididevice.GongMarshallField(stage, "Id"))
	}
	if mididevice.Value != midideviceOther.Value {
		diffs = append(diffs, mididevice.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (midiinstrument *MidiInstrument) GongDiff(stage *Stage, midiinstrumentOther *MidiInstrument) (diffs []string) {
	// insertion point for field diffs
	if midiinstrument.Name != midiinstrumentOther.Name {
		diffs = append(diffs, midiinstrument.GongMarshallField(stage, "Name"))
	}
	if midiinstrument.Id != midiinstrumentOther.Id {
		diffs = append(diffs, midiinstrument.GongMarshallField(stage, "Id"))
	}
	if midiinstrument.MidiChannel != midiinstrumentOther.MidiChannel {
		diffs = append(diffs, midiinstrument.GongMarshallField(stage, "MidiChannel"))
	}
	if midiinstrument.MidiName != midiinstrumentOther.MidiName {
		diffs = append(diffs, midiinstrument.GongMarshallField(stage, "MidiName"))
	}
	if midiinstrument.MidiBank != midiinstrumentOther.MidiBank {
		diffs = append(diffs, midiinstrument.GongMarshallField(stage, "MidiBank"))
	}
	if midiinstrument.MidiProgram != midiinstrumentOther.MidiProgram {
		diffs = append(diffs, midiinstrument.GongMarshallField(stage, "MidiProgram"))
	}
	if midiinstrument.MidiUnpitched != midiinstrumentOther.MidiUnpitched {
		diffs = append(diffs, midiinstrument.GongMarshallField(stage, "MidiUnpitched"))
	}
	if midiinstrument.Volume != midiinstrumentOther.Volume {
		diffs = append(diffs, midiinstrument.GongMarshallField(stage, "Volume"))
	}
	if midiinstrument.Pan != midiinstrumentOther.Pan {
		diffs = append(diffs, midiinstrument.GongMarshallField(stage, "Pan"))
	}
	if midiinstrument.Elevation != midiinstrumentOther.Elevation {
		diffs = append(diffs, midiinstrument.GongMarshallField(stage, "Elevation"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (miscellaneous *Miscellaneous) GongDiff(stage *Stage, miscellaneousOther *Miscellaneous) (diffs []string) {
	// insertion point for field diffs
	if miscellaneous.Name != miscellaneousOther.Name {
		diffs = append(diffs, miscellaneous.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, miscellaneous, "MiscellaneousField", miscellaneousOther.MiscellaneousField, miscellaneous.MiscellaneousField); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (miscellaneousfield *MiscellaneousField) GongDiff(stage *Stage, miscellaneousfieldOther *MiscellaneousField) (diffs []string) {
	// insertion point for field diffs
	if miscellaneousfield.Name != miscellaneousfieldOther.Name {
		diffs = append(diffs, miscellaneousfield.GongMarshallField(stage, "Name"))
	}
	if miscellaneousfield.NameAttr != miscellaneousfieldOther.NameAttr {
		diffs = append(diffs, miscellaneousfield.GongMarshallField(stage, "NameAttr"))
	}
	if miscellaneousfield.Value != miscellaneousfieldOther.Value {
		diffs = append(diffs, miscellaneousfield.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (namedisplay *NameDisplay) GongDiff(stage *Stage, namedisplayOther *NameDisplay) (diffs []string) {
	// insertion point for field diffs
	if namedisplay.Name != namedisplayOther.Name {
		diffs = append(diffs, namedisplay.GongMarshallField(stage, "Name"))
	}
	if namedisplay.PrintObject != namedisplayOther.PrintObject {
		diffs = append(diffs, namedisplay.GongMarshallField(stage, "PrintObject"))
	}
	if ops := __gong__diffSliceOfPointers(stage, namedisplay, "DisplayText", namedisplayOther.DisplayText, namedisplay.DisplayText); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, namedisplay, "AccidentalText", namedisplayOther.AccidentalText, namedisplay.AccidentalText); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notesize *NoteSize) GongDiff(stage *Stage, notesizeOther *NoteSize) (diffs []string) {
	// insertion point for field diffs
	if notesize.Name != notesizeOther.Name {
		diffs = append(diffs, notesize.GongMarshallField(stage, "Name"))
	}
	if notesize.Type != notesizeOther.Type {
		diffs = append(diffs, notesize.GongMarshallField(stage, "Type"))
	}
	if notesize.Value != notesizeOther.Value {
		diffs = append(diffs, notesize.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (otherappearance *OtherAppearance) GongDiff(stage *Stage, otherappearanceOther *OtherAppearance) (diffs []string) {
	// insertion point for field diffs
	if otherappearance.Name != otherappearanceOther.Name {
		diffs = append(diffs, otherappearance.GongMarshallField(stage, "Name"))
	}
	if otherappearance.Type != otherappearanceOther.Type {
		diffs = append(diffs, otherappearance.GongMarshallField(stage, "Type"))
	}
	if otherappearance.Value != otherappearanceOther.Value {
		diffs = append(diffs, otherappearance.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (otherplay *OtherPlay) GongDiff(stage *Stage, otherplayOther *OtherPlay) (diffs []string) {
	// insertion point for field diffs
	if otherplay.Name != otherplayOther.Name {
		diffs = append(diffs, otherplay.GongMarshallField(stage, "Name"))
	}
	if otherplay.Type != otherplayOther.Type {
		diffs = append(diffs, otherplay.GongMarshallField(stage, "Type"))
	}
	if otherplay.Value != otherplayOther.Value {
		diffs = append(diffs, otherplay.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (othertext *OtherText) GongDiff(stage *Stage, othertextOther *OtherText) (diffs []string) {
	// insertion point for field diffs
	if othertext.Name != othertextOther.Name {
		diffs = append(diffs, othertext.GongMarshallField(stage, "Name"))
	}
	if othertext.Smufl != othertextOther.Smufl {
		diffs = append(diffs, othertext.GongMarshallField(stage, "Smufl"))
	}
	if othertext.Value != othertextOther.Value {
		diffs = append(diffs, othertext.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pagelayout *PageLayout) GongDiff(stage *Stage, pagelayoutOther *PageLayout) (diffs []string) {
	// insertion point for field diffs
	if pagelayout.Name != pagelayoutOther.Name {
		diffs = append(diffs, pagelayout.GongMarshallField(stage, "Name"))
	}
	if pagelayout.PageHeight != pagelayoutOther.PageHeight {
		diffs = append(diffs, pagelayout.GongMarshallField(stage, "PageHeight"))
	}
	if pagelayout.PageWidth != pagelayoutOther.PageWidth {
		diffs = append(diffs, pagelayout.GongMarshallField(stage, "PageWidth"))
	}
	if ops := __gong__diffSliceOfPointers(stage, pagelayout, "PageMargins", pagelayoutOther.PageMargins, pagelayout.PageMargins); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pagemargins *PageMargins) GongDiff(stage *Stage, pagemarginsOther *PageMargins) (diffs []string) {
	// insertion point for field diffs
	if pagemargins.Name != pagemarginsOther.Name {
		diffs = append(diffs, pagemargins.GongMarshallField(stage, "Name"))
	}
	if pagemargins.Type != pagemarginsOther.Type {
		diffs = append(diffs, pagemargins.GongMarshallField(stage, "Type"))
	}
	if pagemargins.LeftMargin != pagemarginsOther.LeftMargin {
		diffs = append(diffs, pagemargins.GongMarshallField(stage, "LeftMargin"))
	}
	if pagemargins.RightMargin != pagemarginsOther.RightMargin {
		diffs = append(diffs, pagemargins.GongMarshallField(stage, "RightMargin"))
	}
	if pagemargins.TopMargin != pagemarginsOther.TopMargin {
		diffs = append(diffs, pagemargins.GongMarshallField(stage, "TopMargin"))
	}
	if pagemargins.BottomMargin != pagemarginsOther.BottomMargin {
		diffs = append(diffs, pagemargins.GongMarshallField(stage, "BottomMargin"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (play *Play) GongDiff(stage *Stage, playOther *Play) (diffs []string) {
	// insertion point for field diffs
	if play.Name != playOther.Name {
		diffs = append(diffs, play.GongMarshallField(stage, "Name"))
	}
	if play.Id != playOther.Id {
		diffs = append(diffs, play.GongMarshallField(stage, "Id"))
	}
	if play.Ipa != playOther.Ipa {
		diffs = append(diffs, play.GongMarshallField(stage, "Ipa"))
	}
	if play.Mute != playOther.Mute {
		diffs = append(diffs, play.GongMarshallField(stage, "Mute"))
	}
	if play.SemiPitched != playOther.SemiPitched {
		diffs = append(diffs, play.GongMarshallField(stage, "SemiPitched"))
	}
	if ops := __gong__diffSliceOfPointers(stage, play, "OtherPlay", playOther.OtherPlay, play.OtherPlay); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (scaling *Scaling) GongDiff(stage *Stage, scalingOther *Scaling) (diffs []string) {
	// insertion point for field diffs
	if scaling.Name != scalingOther.Name {
		diffs = append(diffs, scaling.GongMarshallField(stage, "Name"))
	}
	if scaling.Millimeters != scalingOther.Millimeters {
		diffs = append(diffs, scaling.GongMarshallField(stage, "Millimeters"))
	}
	if scaling.Tenths != scalingOther.Tenths {
		diffs = append(diffs, scaling.GongMarshallField(stage, "Tenths"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (segno *Segno) GongDiff(stage *Stage, segnoOther *Segno) (diffs []string) {
	// insertion point for field diffs
	if segno.Name != segnoOther.Name {
		diffs = append(diffs, segno.GongMarshallField(stage, "Name"))
	}
	if segno.Smufl != segnoOther.Smufl {
		diffs = append(diffs, segno.GongMarshallField(stage, "Smufl"))
	}
	if segno.DefaultX != segnoOther.DefaultX {
		diffs = append(diffs, segno.GongMarshallField(stage, "DefaultX"))
	}
	if segno.DefaultY != segnoOther.DefaultY {
		diffs = append(diffs, segno.GongMarshallField(stage, "DefaultY"))
	}
	if segno.RelativeX != segnoOther.RelativeX {
		diffs = append(diffs, segno.GongMarshallField(stage, "RelativeX"))
	}
	if segno.RelativeY != segnoOther.RelativeY {
		diffs = append(diffs, segno.GongMarshallField(stage, "RelativeY"))
	}
	if segno.FontFamily != segnoOther.FontFamily {
		diffs = append(diffs, segno.GongMarshallField(stage, "FontFamily"))
	}
	if segno.FontStyle != segnoOther.FontStyle {
		diffs = append(diffs, segno.GongMarshallField(stage, "FontStyle"))
	}
	if segno.FontSize != segnoOther.FontSize {
		diffs = append(diffs, segno.GongMarshallField(stage, "FontSize"))
	}
	if segno.FontWeight != segnoOther.FontWeight {
		diffs = append(diffs, segno.GongMarshallField(stage, "FontWeight"))
	}
	if segno.Color != segnoOther.Color {
		diffs = append(diffs, segno.GongMarshallField(stage, "Color"))
	}
	if segno.Halign != segnoOther.Halign {
		diffs = append(diffs, segno.GongMarshallField(stage, "Halign"))
	}
	if segno.Valign != segnoOther.Valign {
		diffs = append(diffs, segno.GongMarshallField(stage, "Valign"))
	}
	if segno.Id != segnoOther.Id {
		diffs = append(diffs, segno.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stafflayout *StaffLayout) GongDiff(stage *Stage, stafflayoutOther *StaffLayout) (diffs []string) {
	// insertion point for field diffs
	if stafflayout.Name != stafflayoutOther.Name {
		diffs = append(diffs, stafflayout.GongMarshallField(stage, "Name"))
	}
	if stafflayout.Number != stafflayoutOther.Number {
		diffs = append(diffs, stafflayout.GongMarshallField(stage, "Number"))
	}
	if stafflayout.StaffDistance != stafflayoutOther.StaffDistance {
		diffs = append(diffs, stafflayout.GongMarshallField(stage, "StaffDistance"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stringtype *StringType) GongDiff(stage *Stage, stringtypeOther *StringType) (diffs []string) {
	// insertion point for field diffs
	if stringtype.Name != stringtypeOther.Name {
		diffs = append(diffs, stringtype.GongMarshallField(stage, "Name"))
	}
	if stringtype.DefaultX != stringtypeOther.DefaultX {
		diffs = append(diffs, stringtype.GongMarshallField(stage, "DefaultX"))
	}
	if stringtype.DefaultY != stringtypeOther.DefaultY {
		diffs = append(diffs, stringtype.GongMarshallField(stage, "DefaultY"))
	}
	if stringtype.RelativeX != stringtypeOther.RelativeX {
		diffs = append(diffs, stringtype.GongMarshallField(stage, "RelativeX"))
	}
	if stringtype.RelativeY != stringtypeOther.RelativeY {
		diffs = append(diffs, stringtype.GongMarshallField(stage, "RelativeY"))
	}
	if stringtype.FontFamily != stringtypeOther.FontFamily {
		diffs = append(diffs, stringtype.GongMarshallField(stage, "FontFamily"))
	}
	if stringtype.FontStyle != stringtypeOther.FontStyle {
		diffs = append(diffs, stringtype.GongMarshallField(stage, "FontStyle"))
	}
	if stringtype.FontSize != stringtypeOther.FontSize {
		diffs = append(diffs, stringtype.GongMarshallField(stage, "FontSize"))
	}
	if stringtype.FontWeight != stringtypeOther.FontWeight {
		diffs = append(diffs, stringtype.GongMarshallField(stage, "FontWeight"))
	}
	if stringtype.Color != stringtypeOther.Color {
		diffs = append(diffs, stringtype.GongMarshallField(stage, "Color"))
	}
	if stringtype.Placement != stringtypeOther.Placement {
		diffs = append(diffs, stringtype.GongMarshallField(stage, "Placement"))
	}
	if stringtype.Value != stringtypeOther.Value {
		diffs = append(diffs, stringtype.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (supports *Supports) GongDiff(stage *Stage, supportsOther *Supports) (diffs []string) {
	// insertion point for field diffs
	if supports.Name != supportsOther.Name {
		diffs = append(diffs, supports.GongMarshallField(stage, "Name"))
	}
	if supports.Type != supportsOther.Type {
		diffs = append(diffs, supports.GongMarshallField(stage, "Type"))
	}
	if supports.Element != supportsOther.Element {
		diffs = append(diffs, supports.GongMarshallField(stage, "Element"))
	}
	if supports.Attribute != supportsOther.Attribute {
		diffs = append(diffs, supports.GongMarshallField(stage, "Attribute"))
	}
	if supports.Value != supportsOther.Value {
		diffs = append(diffs, supports.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (systemdividers *SystemDividers) GongDiff(stage *Stage, systemdividersOther *SystemDividers) (diffs []string) {
	// insertion point for field diffs
	if systemdividers.Name != systemdividersOther.Name {
		diffs = append(diffs, systemdividers.GongMarshallField(stage, "Name"))
	}
	if systemdividers.LeftDivider != systemdividersOther.LeftDivider {
		diffs = append(diffs, systemdividers.GongMarshallField(stage, "LeftDivider"))
	}
	if systemdividers.RightDivider != systemdividersOther.RightDivider {
		diffs = append(diffs, systemdividers.GongMarshallField(stage, "RightDivider"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (systemlayout *SystemLayout) GongDiff(stage *Stage, systemlayoutOther *SystemLayout) (diffs []string) {
	// insertion point for field diffs
	if systemlayout.Name != systemlayoutOther.Name {
		diffs = append(diffs, systemlayout.GongMarshallField(stage, "Name"))
	}
	if systemlayout.SystemMargins != systemlayoutOther.SystemMargins {
		diffs = append(diffs, systemlayout.GongMarshallField(stage, "SystemMargins"))
	}
	if systemlayout.SystemDistance != systemlayoutOther.SystemDistance {
		diffs = append(diffs, systemlayout.GongMarshallField(stage, "SystemDistance"))
	}
	if systemlayout.TopSystemDistance != systemlayoutOther.TopSystemDistance {
		diffs = append(diffs, systemlayout.GongMarshallField(stage, "TopSystemDistance"))
	}
	if systemlayout.SystemDividers != systemlayoutOther.SystemDividers {
		diffs = append(diffs, systemlayout.GongMarshallField(stage, "SystemDividers"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (systemmargins *SystemMargins) GongDiff(stage *Stage, systemmarginsOther *SystemMargins) (diffs []string) {
	// insertion point for field diffs
	if systemmargins.Name != systemmarginsOther.Name {
		diffs = append(diffs, systemmargins.GongMarshallField(stage, "Name"))
	}
	if systemmargins.LeftMargin != systemmarginsOther.LeftMargin {
		diffs = append(diffs, systemmargins.GongMarshallField(stage, "LeftMargin"))
	}
	if systemmargins.RightMargin != systemmarginsOther.RightMargin {
		diffs = append(diffs, systemmargins.GongMarshallField(stage, "RightMargin"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (typedtext *TypedText) GongDiff(stage *Stage, typedtextOther *TypedText) (diffs []string) {
	// insertion point for field diffs
	if typedtext.Name != typedtextOther.Name {
		diffs = append(diffs, typedtext.GongMarshallField(stage, "Name"))
	}
	if typedtext.Type != typedtextOther.Type {
		diffs = append(diffs, typedtext.GongMarshallField(stage, "Type"))
	}
	if typedtext.Value != typedtextOther.Value {
		diffs = append(diffs, typedtext.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (wavyline *WavyLine) GongDiff(stage *Stage, wavylineOther *WavyLine) (diffs []string) {
	// insertion point for field diffs
	if wavyline.Name != wavylineOther.Name {
		diffs = append(diffs, wavyline.GongMarshallField(stage, "Name"))
	}
	if wavyline.Type != wavylineOther.Type {
		diffs = append(diffs, wavyline.GongMarshallField(stage, "Type"))
	}
	if wavyline.Number != wavylineOther.Number {
		diffs = append(diffs, wavyline.GongMarshallField(stage, "Number"))
	}
	if wavyline.Smufl != wavylineOther.Smufl {
		diffs = append(diffs, wavyline.GongMarshallField(stage, "Smufl"))
	}
	if wavyline.DefaultX != wavylineOther.DefaultX {
		diffs = append(diffs, wavyline.GongMarshallField(stage, "DefaultX"))
	}
	if wavyline.DefaultY != wavylineOther.DefaultY {
		diffs = append(diffs, wavyline.GongMarshallField(stage, "DefaultY"))
	}
	if wavyline.RelativeX != wavylineOther.RelativeX {
		diffs = append(diffs, wavyline.GongMarshallField(stage, "RelativeX"))
	}
	if wavyline.RelativeY != wavylineOther.RelativeY {
		diffs = append(diffs, wavyline.GongMarshallField(stage, "RelativeY"))
	}
	if wavyline.Placement != wavylineOther.Placement {
		diffs = append(diffs, wavyline.GongMarshallField(stage, "Placement"))
	}
	if wavyline.Color != wavylineOther.Color {
		diffs = append(diffs, wavyline.GongMarshallField(stage, "Color"))
	}
	if wavyline.StartNote != wavylineOther.StartNote {
		diffs = append(diffs, wavyline.GongMarshallField(stage, "StartNote"))
	}
	if wavyline.TrillStep != wavylineOther.TrillStep {
		diffs = append(diffs, wavyline.GongMarshallField(stage, "TrillStep"))
	}
	if wavyline.TwoNoteTurn != wavylineOther.TwoNoteTurn {
		diffs = append(diffs, wavyline.GongMarshallField(stage, "TwoNoteTurn"))
	}
	if wavyline.Accelerate != wavylineOther.Accelerate {
		diffs = append(diffs, wavyline.GongMarshallField(stage, "Accelerate"))
	}
	if wavyline.Beats != wavylineOther.Beats {
		diffs = append(diffs, wavyline.GongMarshallField(stage, "Beats"))
	}
	if wavyline.SecondBeat != wavylineOther.SecondBeat {
		diffs = append(diffs, wavyline.GongMarshallField(stage, "SecondBeat"))
	}
	if wavyline.LastBeat != wavylineOther.LastBeat {
		diffs = append(diffs, wavyline.GongMarshallField(stage, "LastBeat"))
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
