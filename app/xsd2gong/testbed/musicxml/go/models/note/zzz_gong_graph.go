// generated code - do not edit
package note

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
func (accidental *Accidental) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Accidentals[accidental]
	return ok
}

func (accidentalmark *AccidentalMark) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AccidentalMarks[accidentalmark]
	return ok
}

func (arpeggiate *Arpeggiate) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Arpeggiates[arpeggiate]
	return ok
}

func (arrow *Arrow) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Arrows[arrow]
	return ok
}

func (articulations *Articulations) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Articulationss[articulations]
	return ok
}

func (assess *Assess) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Assesss[assess]
	return ok
}

func (backup *Backup) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Backups[backup]
	return ok
}

func (beam *Beam) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Beams[beam]
	return ok
}

func (bend *Bend) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Bends[bend]
	return ok
}

func (breathmark *BreathMark) GongIsStaged(stage *Stage) bool {
	_, ok := stage.BreathMarks[breathmark]
	return ok
}

func (caesura *Caesura) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Caesuras[caesura]
	return ok
}

func (elision *Elision) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Elisions[elision]
	return ok
}

func (emptyline *EmptyLine) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EmptyLines[emptyline]
	return ok
}

func (extend *Extend) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Extends[extend]
	return ok
}

func (figure *Figure) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Figures[figure]
	return ok
}

func (figuredbass *FiguredBass) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FiguredBasss[figuredbass]
	return ok
}

func (forward *Forward) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Forwards[forward]
	return ok
}

func (glissando *Glissando) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Glissandos[glissando]
	return ok
}

func (grace *Grace) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Graces[grace]
	return ok
}

func (hammeronpulloff *HammerOnPullOff) GongIsStaged(stage *Stage) bool {
	_, ok := stage.HammerOnPullOffs[hammeronpulloff]
	return ok
}

func (handbell *Handbell) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Handbells[handbell]
	return ok
}

func (harmonclosed *HarmonClosed) GongIsStaged(stage *Stage) bool {
	_, ok := stage.HarmonCloseds[harmonclosed]
	return ok
}

func (harmonmute *HarmonMute) GongIsStaged(stage *Stage) bool {
	_, ok := stage.HarmonMutes[harmonmute]
	return ok
}

func (harmonic *Harmonic) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Harmonics[harmonic]
	return ok
}

func (heeltoe *HeelToe) GongIsStaged(stage *Stage) bool {
	_, ok := stage.HeelToes[heeltoe]
	return ok
}

func (hole *Hole) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Holes[hole]
	return ok
}

func (holeclosed *HoleClosed) GongIsStaged(stage *Stage) bool {
	_, ok := stage.HoleCloseds[holeclosed]
	return ok
}

func (instrument *Instrument) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Instruments[instrument]
	return ok
}

func (listen *Listen) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Listens[listen]
	return ok
}

func (lyric *Lyric) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Lyrics[lyric]
	return ok
}

func (mordent *Mordent) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Mordents[mordent]
	return ok
}

func (nonarpeggiate *NonArpeggiate) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NonArpeggiates[nonarpeggiate]
	return ok
}

func (notations *Notations) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Notationss[notations]
	return ok
}

func (note *Note) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Notes[note]
	return ok
}

func (notetype *NoteType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteTypes[notetype]
	return ok
}

func (notehead *Notehead) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Noteheads[notehead]
	return ok
}

func (noteheadtext *NoteheadText) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteheadTexts[noteheadtext]
	return ok
}

func (ornaments *Ornaments) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Ornamentss[ornaments]
	return ok
}

func (othernotation *OtherNotation) GongIsStaged(stage *Stage) bool {
	_, ok := stage.OtherNotations[othernotation]
	return ok
}

func (otherplacementtext *OtherPlacementText) GongIsStaged(stage *Stage) bool {
	_, ok := stage.OtherPlacementTexts[otherplacementtext]
	return ok
}

func (pitch *Pitch) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Pitchs[pitch]
	return ok
}

func (placementtext *PlacementText) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PlacementTexts[placementtext]
	return ok
}

func (release *Release) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Releases[release]
	return ok
}

func (rest *Rest) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Rests[rest]
	return ok
}

func (slide *Slide) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Slides[slide]
	return ok
}

func (slur *Slur) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Slurs[slur]
	return ok
}

func (stem *Stem) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Stems[stem]
	return ok
}

func (strongaccent *StrongAccent) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StrongAccents[strongaccent]
	return ok
}

func (tap *Tap) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Taps[tap]
	return ok
}

func (technical *Technical) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Technicals[technical]
	return ok
}

func (textelementdata *TextElementData) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TextElementDatas[textelementdata]
	return ok
}

func (tie *Tie) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Ties[tie]
	return ok
}

func (tied *Tied) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Tieds[tied]
	return ok
}

func (tremolo *Tremolo) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Tremolos[tremolo]
	return ok
}

func (tuplet *Tuplet) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Tuplets[tuplet]
	return ok
}

func (tupletdot *TupletDot) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TupletDots[tupletdot]
	return ok
}

func (tupletnumber *TupletNumber) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TupletNumbers[tupletnumber]
	return ok
}

func (tupletportion *TupletPortion) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TupletPortions[tupletportion]
	return ok
}

func (tuplettype *TupletType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TupletTypes[tuplettype]
	return ok
}

func (unpitched *Unpitched) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Unpitcheds[unpitched]
	return ok
}

func (wait *Wait) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Waits[wait]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (accidental *Accidental) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(accidental) {
		return
	}

	accidental.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (accidentalmark *AccidentalMark) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(accidentalmark) {
		return
	}

	accidentalmark.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (arpeggiate *Arpeggiate) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(arpeggiate) {
		return
	}

	arpeggiate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (arrow *Arrow) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(arrow) {
		return
	}

	arrow.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (articulations *Articulations) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(articulations) {
		return
	}

	articulations.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _strongaccent := range articulations.StrongAccent {
		stage.StageBranch(_strongaccent)
	}
	for _, _emptyline := range articulations.Scoop {
		stage.StageBranch(_emptyline)
	}
	for _, _emptyline := range articulations.Plop {
		stage.StageBranch(_emptyline)
	}
	for _, _emptyline := range articulations.Doit {
		stage.StageBranch(_emptyline)
	}
	for _, _emptyline := range articulations.Falloff {
		stage.StageBranch(_emptyline)
	}
	for _, _breathmark := range articulations.BreathMark {
		stage.StageBranch(_breathmark)
	}
	for _, _caesura := range articulations.Caesura {
		stage.StageBranch(_caesura)
	}
	for _, _otherplacementtext := range articulations.OtherArticulation {
		stage.StageBranch(_otherplacementtext)
	}

}

func (assess *Assess) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(assess) {
		return
	}

	assess.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (backup *Backup) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(backup) {
		return
	}

	backup.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (beam *Beam) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(beam) {
		return
	}

	beam.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bend *Bend) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(bend) {
		return
	}

	bend.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if bend.Release != nil {
		stage.StageBranch(bend.Release)
	}
	if bend.WithBar != nil {
		stage.StageBranch(bend.WithBar)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (breathmark *BreathMark) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(breathmark) {
		return
	}

	breathmark.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (caesura *Caesura) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(caesura) {
		return
	}

	caesura.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (elision *Elision) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(elision) {
		return
	}

	elision.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (emptyline *EmptyLine) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(emptyline) {
		return
	}

	emptyline.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (extend *Extend) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(extend) {
		return
	}

	extend.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (figure *Figure) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(figure) {
		return
	}

	figure.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if figure.Extend != nil {
		stage.StageBranch(figure.Extend)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (figuredbass *FiguredBass) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(figuredbass) {
		return
	}

	figuredbass.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _figure := range figuredbass.Figure {
		stage.StageBranch(_figure)
	}

}

func (forward *Forward) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(forward) {
		return
	}

	forward.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (glissando *Glissando) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(glissando) {
		return
	}

	glissando.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (grace *Grace) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(grace) {
		return
	}

	grace.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (hammeronpulloff *HammerOnPullOff) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(hammeronpulloff) {
		return
	}

	hammeronpulloff.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (handbell *Handbell) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(handbell) {
		return
	}

	handbell.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (harmonclosed *HarmonClosed) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(harmonclosed) {
		return
	}

	harmonclosed.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (harmonmute *HarmonMute) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(harmonmute) {
		return
	}

	harmonmute.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if harmonmute.HarmonClosed != nil {
		stage.StageBranch(harmonmute.HarmonClosed)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (harmonic *Harmonic) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(harmonic) {
		return
	}

	harmonic.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (heeltoe *HeelToe) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(heeltoe) {
		return
	}

	heeltoe.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (hole *Hole) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(hole) {
		return
	}

	hole.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if hole.HoleClosed != nil {
		stage.StageBranch(hole.HoleClosed)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (holeclosed *HoleClosed) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(holeclosed) {
		return
	}

	holeclosed.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (instrument *Instrument) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(instrument) {
		return
	}

	instrument.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (listen *Listen) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(listen) {
		return
	}

	listen.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assess := range listen.Assess {
		stage.StageBranch(_assess)
	}
	for _, _wait := range listen.Wait {
		stage.StageBranch(_wait)
	}

}

func (lyric *Lyric) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(lyric) {
		return
	}

	lyric.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if lyric.Text != nil {
		stage.StageBranch(lyric.Text)
	}
	if lyric.Extend != nil {
		stage.StageBranch(lyric.Extend)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _elision := range lyric.Elision {
		stage.StageBranch(_elision)
	}

}

func (mordent *Mordent) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(mordent) {
		return
	}

	mordent.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (nonarpeggiate *NonArpeggiate) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(nonarpeggiate) {
		return
	}

	nonarpeggiate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notations *Notations) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(notations) {
		return
	}

	notations.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tied := range notations.Tied {
		stage.StageBranch(_tied)
	}
	for _, _slur := range notations.Slur {
		stage.StageBranch(_slur)
	}
	for _, _tuplet := range notations.Tuplet {
		stage.StageBranch(_tuplet)
	}
	for _, _glissando := range notations.Glissando {
		stage.StageBranch(_glissando)
	}
	for _, _slide := range notations.Slide {
		stage.StageBranch(_slide)
	}
	for _, _ornaments := range notations.Ornaments {
		stage.StageBranch(_ornaments)
	}
	for _, _technical := range notations.Technical {
		stage.StageBranch(_technical)
	}
	for _, _articulations := range notations.Articulations {
		stage.StageBranch(_articulations)
	}
	for _, _arpeggiate := range notations.Arpeggiate {
		stage.StageBranch(_arpeggiate)
	}
	for _, _nonarpeggiate := range notations.NonArpeggiate {
		stage.StageBranch(_nonarpeggiate)
	}
	for _, _accidentalmark := range notations.AccidentalMark {
		stage.StageBranch(_accidentalmark)
	}
	for _, _othernotation := range notations.OtherNotation {
		stage.StageBranch(_othernotation)
	}

}

func (note *Note) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(note) {
		return
	}

	note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if note.Grace != nil {
		stage.StageBranch(note.Grace)
	}
	if note.Pitch != nil {
		stage.StageBranch(note.Pitch)
	}
	if note.Unpitched != nil {
		stage.StageBranch(note.Unpitched)
	}
	if note.Rest != nil {
		stage.StageBranch(note.Rest)
	}
	if note.Type != nil {
		stage.StageBranch(note.Type)
	}
	if note.Accidental != nil {
		stage.StageBranch(note.Accidental)
	}
	if note.Stem != nil {
		stage.StageBranch(note.Stem)
	}
	if note.Notehead != nil {
		stage.StageBranch(note.Notehead)
	}
	if note.NoteheadText != nil {
		stage.StageBranch(note.NoteheadText)
	}
	if note.Listen != nil {
		stage.StageBranch(note.Listen)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tie := range note.Tie {
		stage.StageBranch(_tie)
	}
	for _, _instrument := range note.Instrument {
		stage.StageBranch(_instrument)
	}
	for _, _beam := range note.Beam {
		stage.StageBranch(_beam)
	}
	for _, _notations := range note.Notations {
		stage.StageBranch(_notations)
	}
	for _, _lyric := range note.Lyric {
		stage.StageBranch(_lyric)
	}

}

func (notetype *NoteType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(notetype) {
		return
	}

	notetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notehead *Notehead) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(notehead) {
		return
	}

	notehead.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (noteheadtext *NoteheadText) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(noteheadtext) {
		return
	}

	noteheadtext.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ornaments *Ornaments) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ornaments) {
		return
	}

	ornaments.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _mordent := range ornaments.Mordent {
		stage.StageBranch(_mordent)
	}
	for _, _mordent := range ornaments.InvertedMordent {
		stage.StageBranch(_mordent)
	}
	for _, _tremolo := range ornaments.Tremolo {
		stage.StageBranch(_tremolo)
	}
	for _, _otherplacementtext := range ornaments.OtherOrnament {
		stage.StageBranch(_otherplacementtext)
	}
	for _, _accidentalmark := range ornaments.AccidentalMark {
		stage.StageBranch(_accidentalmark)
	}

}

func (othernotation *OtherNotation) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(othernotation) {
		return
	}

	othernotation.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (otherplacementtext *OtherPlacementText) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(otherplacementtext) {
		return
	}

	otherplacementtext.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pitch *Pitch) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(pitch) {
		return
	}

	pitch.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (placementtext *PlacementText) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(placementtext) {
		return
	}

	placementtext.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (release *Release) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(release) {
		return
	}

	release.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rest *Rest) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(rest) {
		return
	}

	rest.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (slide *Slide) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(slide) {
		return
	}

	slide.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (slur *Slur) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(slur) {
		return
	}

	slur.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stem *Stem) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stem) {
		return
	}

	stem.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (strongaccent *StrongAccent) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(strongaccent) {
		return
	}

	strongaccent.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tap *Tap) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tap) {
		return
	}

	tap.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (technical *Technical) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(technical) {
		return
	}

	technical.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _harmonic := range technical.Harmonic {
		stage.StageBranch(_harmonic)
	}
	for _, _placementtext := range technical.Pluck {
		stage.StageBranch(_placementtext)
	}
	for _, _hammeronpulloff := range technical.HammerOn {
		stage.StageBranch(_hammeronpulloff)
	}
	for _, _hammeronpulloff := range technical.PullOff {
		stage.StageBranch(_hammeronpulloff)
	}
	for _, _bend := range technical.Bend {
		stage.StageBranch(_bend)
	}
	for _, _tap := range technical.Tap {
		stage.StageBranch(_tap)
	}
	for _, _heeltoe := range technical.Heel {
		stage.StageBranch(_heeltoe)
	}
	for _, _heeltoe := range technical.Toe {
		stage.StageBranch(_heeltoe)
	}
	for _, _hole := range technical.Hole {
		stage.StageBranch(_hole)
	}
	for _, _arrow := range technical.Arrow {
		stage.StageBranch(_arrow)
	}
	for _, _handbell := range technical.Handbell {
		stage.StageBranch(_handbell)
	}
	for _, _harmonmute := range technical.HarmonMute {
		stage.StageBranch(_harmonmute)
	}
	for _, _otherplacementtext := range technical.OtherTechnical {
		stage.StageBranch(_otherplacementtext)
	}

}

func (textelementdata *TextElementData) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(textelementdata) {
		return
	}

	textelementdata.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tie *Tie) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tie) {
		return
	}

	tie.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tied *Tied) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tied) {
		return
	}

	tied.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tremolo *Tremolo) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tremolo) {
		return
	}

	tremolo.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tuplet *Tuplet) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tuplet) {
		return
	}

	tuplet.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tuplet.TupletActual != nil {
		stage.StageBranch(tuplet.TupletActual)
	}
	if tuplet.TupletNormal != nil {
		stage.StageBranch(tuplet.TupletNormal)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tupletdot *TupletDot) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tupletdot) {
		return
	}

	tupletdot.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tupletnumber *TupletNumber) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tupletnumber) {
		return
	}

	tupletnumber.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tupletportion *TupletPortion) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tupletportion) {
		return
	}

	tupletportion.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tupletportion.TupletNumber != nil {
		stage.StageBranch(tupletportion.TupletNumber)
	}
	if tupletportion.TupletType != nil {
		stage.StageBranch(tupletportion.TupletType)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tupletdot := range tupletportion.TupletDot {
		stage.StageBranch(_tupletdot)
	}

}

func (tuplettype *TupletType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tuplettype) {
		return
	}

	tuplettype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (unpitched *Unpitched) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(unpitched) {
		return
	}

	unpitched.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (wait *Wait) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(wait) {
		return
	}

	wait.Stage(stage)

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
	case *Accidental:
		toT := GongCopyBranchAccidental(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AccidentalMark:
		toT := GongCopyBranchAccidentalMark(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Arpeggiate:
		toT := GongCopyBranchArpeggiate(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Arrow:
		toT := GongCopyBranchArrow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Articulations:
		toT := GongCopyBranchArticulations(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Assess:
		toT := GongCopyBranchAssess(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Backup:
		toT := GongCopyBranchBackup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Beam:
		toT := GongCopyBranchBeam(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bend:
		toT := GongCopyBranchBend(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BreathMark:
		toT := GongCopyBranchBreathMark(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Caesura:
		toT := GongCopyBranchCaesura(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Elision:
		toT := GongCopyBranchElision(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EmptyLine:
		toT := GongCopyBranchEmptyLine(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Extend:
		toT := GongCopyBranchExtend(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Figure:
		toT := GongCopyBranchFigure(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FiguredBass:
		toT := GongCopyBranchFiguredBass(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Forward:
		toT := GongCopyBranchForward(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Glissando:
		toT := GongCopyBranchGlissando(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Grace:
		toT := GongCopyBranchGrace(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *HammerOnPullOff:
		toT := GongCopyBranchHammerOnPullOff(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Handbell:
		toT := GongCopyBranchHandbell(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *HarmonClosed:
		toT := GongCopyBranchHarmonClosed(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *HarmonMute:
		toT := GongCopyBranchHarmonMute(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Harmonic:
		toT := GongCopyBranchHarmonic(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *HeelToe:
		toT := GongCopyBranchHeelToe(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Hole:
		toT := GongCopyBranchHole(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *HoleClosed:
		toT := GongCopyBranchHoleClosed(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Instrument:
		toT := GongCopyBranchInstrument(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Listen:
		toT := GongCopyBranchListen(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Lyric:
		toT := GongCopyBranchLyric(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Mordent:
		toT := GongCopyBranchMordent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NonArpeggiate:
		toT := GongCopyBranchNonArpeggiate(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Notations:
		toT := GongCopyBranchNotations(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note:
		toT := GongCopyBranchNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteType:
		toT := GongCopyBranchNoteType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Notehead:
		toT := GongCopyBranchNotehead(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteheadText:
		toT := GongCopyBranchNoteheadText(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Ornaments:
		toT := GongCopyBranchOrnaments(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *OtherNotation:
		toT := GongCopyBranchOtherNotation(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *OtherPlacementText:
		toT := GongCopyBranchOtherPlacementText(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Pitch:
		toT := GongCopyBranchPitch(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PlacementText:
		toT := GongCopyBranchPlacementText(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Release:
		toT := GongCopyBranchRelease(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Rest:
		toT := GongCopyBranchRest(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Slide:
		toT := GongCopyBranchSlide(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Slur:
		toT := GongCopyBranchSlur(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Stem:
		toT := GongCopyBranchStem(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StrongAccent:
		toT := GongCopyBranchStrongAccent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tap:
		toT := GongCopyBranchTap(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Technical:
		toT := GongCopyBranchTechnical(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TextElementData:
		toT := GongCopyBranchTextElementData(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tie:
		toT := GongCopyBranchTie(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tied:
		toT := GongCopyBranchTied(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tremolo:
		toT := GongCopyBranchTremolo(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tuplet:
		toT := GongCopyBranchTuplet(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TupletDot:
		toT := GongCopyBranchTupletDot(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TupletNumber:
		toT := GongCopyBranchTupletNumber(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TupletPortion:
		toT := GongCopyBranchTupletPortion(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TupletType:
		toT := GongCopyBranchTupletType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Unpitched:
		toT := GongCopyBranchUnpitched(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Wait:
		toT := GongCopyBranchWait(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAccidental(mapOrigCopy map[any]any, accidentalFrom *Accidental) (accidentalTo *Accidental) {
	var alreadyCopied bool
	accidentalTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, accidentalFrom)
	if alreadyCopied {
		return
	}
	accidentalFrom.GongCopyBasicFields(accidentalTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAccidentalMark(mapOrigCopy map[any]any, accidentalmarkFrom *AccidentalMark) (accidentalmarkTo *AccidentalMark) {
	var alreadyCopied bool
	accidentalmarkTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, accidentalmarkFrom)
	if alreadyCopied {
		return
	}
	accidentalmarkFrom.GongCopyBasicFields(accidentalmarkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchArpeggiate(mapOrigCopy map[any]any, arpeggiateFrom *Arpeggiate) (arpeggiateTo *Arpeggiate) {
	var alreadyCopied bool
	arpeggiateTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, arpeggiateFrom)
	if alreadyCopied {
		return
	}
	arpeggiateFrom.GongCopyBasicFields(arpeggiateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchArrow(mapOrigCopy map[any]any, arrowFrom *Arrow) (arrowTo *Arrow) {
	var alreadyCopied bool
	arrowTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, arrowFrom)
	if alreadyCopied {
		return
	}
	arrowFrom.GongCopyBasicFields(arrowTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchArticulations(mapOrigCopy map[any]any, articulationsFrom *Articulations) (articulationsTo *Articulations) {
	var alreadyCopied bool
	articulationsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, articulationsFrom)
	if alreadyCopied {
		return
	}
	articulationsFrom.GongCopyBasicFields(articulationsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _strongaccent := range articulationsFrom.StrongAccent {
		articulationsTo.StrongAccent = append(articulationsTo.StrongAccent, GongCopyBranchStrongAccent(mapOrigCopy, _strongaccent))
	}
	for _, _emptyline := range articulationsFrom.Scoop {
		articulationsTo.Scoop = append(articulationsTo.Scoop, GongCopyBranchEmptyLine(mapOrigCopy, _emptyline))
	}
	for _, _emptyline := range articulationsFrom.Plop {
		articulationsTo.Plop = append(articulationsTo.Plop, GongCopyBranchEmptyLine(mapOrigCopy, _emptyline))
	}
	for _, _emptyline := range articulationsFrom.Doit {
		articulationsTo.Doit = append(articulationsTo.Doit, GongCopyBranchEmptyLine(mapOrigCopy, _emptyline))
	}
	for _, _emptyline := range articulationsFrom.Falloff {
		articulationsTo.Falloff = append(articulationsTo.Falloff, GongCopyBranchEmptyLine(mapOrigCopy, _emptyline))
	}
	for _, _breathmark := range articulationsFrom.BreathMark {
		articulationsTo.BreathMark = append(articulationsTo.BreathMark, GongCopyBranchBreathMark(mapOrigCopy, _breathmark))
	}
	for _, _caesura := range articulationsFrom.Caesura {
		articulationsTo.Caesura = append(articulationsTo.Caesura, GongCopyBranchCaesura(mapOrigCopy, _caesura))
	}
	for _, _otherplacementtext := range articulationsFrom.OtherArticulation {
		articulationsTo.OtherArticulation = append(articulationsTo.OtherArticulation, GongCopyBranchOtherPlacementText(mapOrigCopy, _otherplacementtext))
	}

	return
}

func GongCopyBranchAssess(mapOrigCopy map[any]any, assessFrom *Assess) (assessTo *Assess) {
	var alreadyCopied bool
	assessTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, assessFrom)
	if alreadyCopied {
		return
	}
	assessFrom.GongCopyBasicFields(assessTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBackup(mapOrigCopy map[any]any, backupFrom *Backup) (backupTo *Backup) {
	var alreadyCopied bool
	backupTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, backupFrom)
	if alreadyCopied {
		return
	}
	backupFrom.GongCopyBasicFields(backupTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBeam(mapOrigCopy map[any]any, beamFrom *Beam) (beamTo *Beam) {
	var alreadyCopied bool
	beamTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, beamFrom)
	if alreadyCopied {
		return
	}
	beamFrom.GongCopyBasicFields(beamTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBend(mapOrigCopy map[any]any, bendFrom *Bend) (bendTo *Bend) {
	var alreadyCopied bool
	bendTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, bendFrom)
	if alreadyCopied {
		return
	}
	bendFrom.GongCopyBasicFields(bendTo)

	//insertion point for the staging of instances referenced by pointers
	if bendFrom.Release != nil {
		bendTo.Release = GongCopyBranchRelease(mapOrigCopy, bendFrom.Release)
	}
	if bendFrom.WithBar != nil {
		bendTo.WithBar = GongCopyBranchPlacementText(mapOrigCopy, bendFrom.WithBar)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBreathMark(mapOrigCopy map[any]any, breathmarkFrom *BreathMark) (breathmarkTo *BreathMark) {
	var alreadyCopied bool
	breathmarkTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, breathmarkFrom)
	if alreadyCopied {
		return
	}
	breathmarkFrom.GongCopyBasicFields(breathmarkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCaesura(mapOrigCopy map[any]any, caesuraFrom *Caesura) (caesuraTo *Caesura) {
	var alreadyCopied bool
	caesuraTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, caesuraFrom)
	if alreadyCopied {
		return
	}
	caesuraFrom.GongCopyBasicFields(caesuraTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchElision(mapOrigCopy map[any]any, elisionFrom *Elision) (elisionTo *Elision) {
	var alreadyCopied bool
	elisionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, elisionFrom)
	if alreadyCopied {
		return
	}
	elisionFrom.GongCopyBasicFields(elisionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEmptyLine(mapOrigCopy map[any]any, emptylineFrom *EmptyLine) (emptylineTo *EmptyLine) {
	var alreadyCopied bool
	emptylineTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, emptylineFrom)
	if alreadyCopied {
		return
	}
	emptylineFrom.GongCopyBasicFields(emptylineTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchExtend(mapOrigCopy map[any]any, extendFrom *Extend) (extendTo *Extend) {
	var alreadyCopied bool
	extendTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, extendFrom)
	if alreadyCopied {
		return
	}
	extendFrom.GongCopyBasicFields(extendTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFigure(mapOrigCopy map[any]any, figureFrom *Figure) (figureTo *Figure) {
	var alreadyCopied bool
	figureTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, figureFrom)
	if alreadyCopied {
		return
	}
	figureFrom.GongCopyBasicFields(figureTo)

	//insertion point for the staging of instances referenced by pointers
	if figureFrom.Extend != nil {
		figureTo.Extend = GongCopyBranchExtend(mapOrigCopy, figureFrom.Extend)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFiguredBass(mapOrigCopy map[any]any, figuredbassFrom *FiguredBass) (figuredbassTo *FiguredBass) {
	var alreadyCopied bool
	figuredbassTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, figuredbassFrom)
	if alreadyCopied {
		return
	}
	figuredbassFrom.GongCopyBasicFields(figuredbassTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _figure := range figuredbassFrom.Figure {
		figuredbassTo.Figure = append(figuredbassTo.Figure, GongCopyBranchFigure(mapOrigCopy, _figure))
	}

	return
}

func GongCopyBranchForward(mapOrigCopy map[any]any, forwardFrom *Forward) (forwardTo *Forward) {
	var alreadyCopied bool
	forwardTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, forwardFrom)
	if alreadyCopied {
		return
	}
	forwardFrom.GongCopyBasicFields(forwardTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGlissando(mapOrigCopy map[any]any, glissandoFrom *Glissando) (glissandoTo *Glissando) {
	var alreadyCopied bool
	glissandoTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, glissandoFrom)
	if alreadyCopied {
		return
	}
	glissandoFrom.GongCopyBasicFields(glissandoTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGrace(mapOrigCopy map[any]any, graceFrom *Grace) (graceTo *Grace) {
	var alreadyCopied bool
	graceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, graceFrom)
	if alreadyCopied {
		return
	}
	graceFrom.GongCopyBasicFields(graceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchHammerOnPullOff(mapOrigCopy map[any]any, hammeronpulloffFrom *HammerOnPullOff) (hammeronpulloffTo *HammerOnPullOff) {
	var alreadyCopied bool
	hammeronpulloffTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, hammeronpulloffFrom)
	if alreadyCopied {
		return
	}
	hammeronpulloffFrom.GongCopyBasicFields(hammeronpulloffTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchHandbell(mapOrigCopy map[any]any, handbellFrom *Handbell) (handbellTo *Handbell) {
	var alreadyCopied bool
	handbellTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, handbellFrom)
	if alreadyCopied {
		return
	}
	handbellFrom.GongCopyBasicFields(handbellTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchHarmonClosed(mapOrigCopy map[any]any, harmonclosedFrom *HarmonClosed) (harmonclosedTo *HarmonClosed) {
	var alreadyCopied bool
	harmonclosedTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, harmonclosedFrom)
	if alreadyCopied {
		return
	}
	harmonclosedFrom.GongCopyBasicFields(harmonclosedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchHarmonMute(mapOrigCopy map[any]any, harmonmuteFrom *HarmonMute) (harmonmuteTo *HarmonMute) {
	var alreadyCopied bool
	harmonmuteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, harmonmuteFrom)
	if alreadyCopied {
		return
	}
	harmonmuteFrom.GongCopyBasicFields(harmonmuteTo)

	//insertion point for the staging of instances referenced by pointers
	if harmonmuteFrom.HarmonClosed != nil {
		harmonmuteTo.HarmonClosed = GongCopyBranchHarmonClosed(mapOrigCopy, harmonmuteFrom.HarmonClosed)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchHarmonic(mapOrigCopy map[any]any, harmonicFrom *Harmonic) (harmonicTo *Harmonic) {
	var alreadyCopied bool
	harmonicTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, harmonicFrom)
	if alreadyCopied {
		return
	}
	harmonicFrom.GongCopyBasicFields(harmonicTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchHeelToe(mapOrigCopy map[any]any, heeltoeFrom *HeelToe) (heeltoeTo *HeelToe) {
	var alreadyCopied bool
	heeltoeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, heeltoeFrom)
	if alreadyCopied {
		return
	}
	heeltoeFrom.GongCopyBasicFields(heeltoeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchHole(mapOrigCopy map[any]any, holeFrom *Hole) (holeTo *Hole) {
	var alreadyCopied bool
	holeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, holeFrom)
	if alreadyCopied {
		return
	}
	holeFrom.GongCopyBasicFields(holeTo)

	//insertion point for the staging of instances referenced by pointers
	if holeFrom.HoleClosed != nil {
		holeTo.HoleClosed = GongCopyBranchHoleClosed(mapOrigCopy, holeFrom.HoleClosed)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchHoleClosed(mapOrigCopy map[any]any, holeclosedFrom *HoleClosed) (holeclosedTo *HoleClosed) {
	var alreadyCopied bool
	holeclosedTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, holeclosedFrom)
	if alreadyCopied {
		return
	}
	holeclosedFrom.GongCopyBasicFields(holeclosedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchInstrument(mapOrigCopy map[any]any, instrumentFrom *Instrument) (instrumentTo *Instrument) {
	var alreadyCopied bool
	instrumentTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, instrumentFrom)
	if alreadyCopied {
		return
	}
	instrumentFrom.GongCopyBasicFields(instrumentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchListen(mapOrigCopy map[any]any, listenFrom *Listen) (listenTo *Listen) {
	var alreadyCopied bool
	listenTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, listenFrom)
	if alreadyCopied {
		return
	}
	listenFrom.GongCopyBasicFields(listenTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assess := range listenFrom.Assess {
		listenTo.Assess = append(listenTo.Assess, GongCopyBranchAssess(mapOrigCopy, _assess))
	}
	for _, _wait := range listenFrom.Wait {
		listenTo.Wait = append(listenTo.Wait, GongCopyBranchWait(mapOrigCopy, _wait))
	}

	return
}

func GongCopyBranchLyric(mapOrigCopy map[any]any, lyricFrom *Lyric) (lyricTo *Lyric) {
	var alreadyCopied bool
	lyricTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, lyricFrom)
	if alreadyCopied {
		return
	}
	lyricFrom.GongCopyBasicFields(lyricTo)

	//insertion point for the staging of instances referenced by pointers
	if lyricFrom.Text != nil {
		lyricTo.Text = GongCopyBranchTextElementData(mapOrigCopy, lyricFrom.Text)
	}
	if lyricFrom.Extend != nil {
		lyricTo.Extend = GongCopyBranchExtend(mapOrigCopy, lyricFrom.Extend)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _elision := range lyricFrom.Elision {
		lyricTo.Elision = append(lyricTo.Elision, GongCopyBranchElision(mapOrigCopy, _elision))
	}

	return
}

func GongCopyBranchMordent(mapOrigCopy map[any]any, mordentFrom *Mordent) (mordentTo *Mordent) {
	var alreadyCopied bool
	mordentTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, mordentFrom)
	if alreadyCopied {
		return
	}
	mordentFrom.GongCopyBasicFields(mordentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNonArpeggiate(mapOrigCopy map[any]any, nonarpeggiateFrom *NonArpeggiate) (nonarpeggiateTo *NonArpeggiate) {
	var alreadyCopied bool
	nonarpeggiateTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, nonarpeggiateFrom)
	if alreadyCopied {
		return
	}
	nonarpeggiateFrom.GongCopyBasicFields(nonarpeggiateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNotations(mapOrigCopy map[any]any, notationsFrom *Notations) (notationsTo *Notations) {
	var alreadyCopied bool
	notationsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, notationsFrom)
	if alreadyCopied {
		return
	}
	notationsFrom.GongCopyBasicFields(notationsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tied := range notationsFrom.Tied {
		notationsTo.Tied = append(notationsTo.Tied, GongCopyBranchTied(mapOrigCopy, _tied))
	}
	for _, _slur := range notationsFrom.Slur {
		notationsTo.Slur = append(notationsTo.Slur, GongCopyBranchSlur(mapOrigCopy, _slur))
	}
	for _, _tuplet := range notationsFrom.Tuplet {
		notationsTo.Tuplet = append(notationsTo.Tuplet, GongCopyBranchTuplet(mapOrigCopy, _tuplet))
	}
	for _, _glissando := range notationsFrom.Glissando {
		notationsTo.Glissando = append(notationsTo.Glissando, GongCopyBranchGlissando(mapOrigCopy, _glissando))
	}
	for _, _slide := range notationsFrom.Slide {
		notationsTo.Slide = append(notationsTo.Slide, GongCopyBranchSlide(mapOrigCopy, _slide))
	}
	for _, _ornaments := range notationsFrom.Ornaments {
		notationsTo.Ornaments = append(notationsTo.Ornaments, GongCopyBranchOrnaments(mapOrigCopy, _ornaments))
	}
	for _, _technical := range notationsFrom.Technical {
		notationsTo.Technical = append(notationsTo.Technical, GongCopyBranchTechnical(mapOrigCopy, _technical))
	}
	for _, _articulations := range notationsFrom.Articulations {
		notationsTo.Articulations = append(notationsTo.Articulations, GongCopyBranchArticulations(mapOrigCopy, _articulations))
	}
	for _, _arpeggiate := range notationsFrom.Arpeggiate {
		notationsTo.Arpeggiate = append(notationsTo.Arpeggiate, GongCopyBranchArpeggiate(mapOrigCopy, _arpeggiate))
	}
	for _, _nonarpeggiate := range notationsFrom.NonArpeggiate {
		notationsTo.NonArpeggiate = append(notationsTo.NonArpeggiate, GongCopyBranchNonArpeggiate(mapOrigCopy, _nonarpeggiate))
	}
	for _, _accidentalmark := range notationsFrom.AccidentalMark {
		notationsTo.AccidentalMark = append(notationsTo.AccidentalMark, GongCopyBranchAccidentalMark(mapOrigCopy, _accidentalmark))
	}
	for _, _othernotation := range notationsFrom.OtherNotation {
		notationsTo.OtherNotation = append(notationsTo.OtherNotation, GongCopyBranchOtherNotation(mapOrigCopy, _othernotation))
	}

	return
}

func GongCopyBranchNote(mapOrigCopy map[any]any, noteFrom *Note) (noteTo *Note) {
	var alreadyCopied bool
	noteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteFrom)
	if alreadyCopied {
		return
	}
	noteFrom.GongCopyBasicFields(noteTo)

	//insertion point for the staging of instances referenced by pointers
	if noteFrom.Grace != nil {
		noteTo.Grace = GongCopyBranchGrace(mapOrigCopy, noteFrom.Grace)
	}
	if noteFrom.Pitch != nil {
		noteTo.Pitch = GongCopyBranchPitch(mapOrigCopy, noteFrom.Pitch)
	}
	if noteFrom.Unpitched != nil {
		noteTo.Unpitched = GongCopyBranchUnpitched(mapOrigCopy, noteFrom.Unpitched)
	}
	if noteFrom.Rest != nil {
		noteTo.Rest = GongCopyBranchRest(mapOrigCopy, noteFrom.Rest)
	}
	if noteFrom.Type != nil {
		noteTo.Type = GongCopyBranchNoteType(mapOrigCopy, noteFrom.Type)
	}
	if noteFrom.Accidental != nil {
		noteTo.Accidental = GongCopyBranchAccidental(mapOrigCopy, noteFrom.Accidental)
	}
	if noteFrom.Stem != nil {
		noteTo.Stem = GongCopyBranchStem(mapOrigCopy, noteFrom.Stem)
	}
	if noteFrom.Notehead != nil {
		noteTo.Notehead = GongCopyBranchNotehead(mapOrigCopy, noteFrom.Notehead)
	}
	if noteFrom.NoteheadText != nil {
		noteTo.NoteheadText = GongCopyBranchNoteheadText(mapOrigCopy, noteFrom.NoteheadText)
	}
	if noteFrom.Listen != nil {
		noteTo.Listen = GongCopyBranchListen(mapOrigCopy, noteFrom.Listen)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tie := range noteFrom.Tie {
		noteTo.Tie = append(noteTo.Tie, GongCopyBranchTie(mapOrigCopy, _tie))
	}
	for _, _instrument := range noteFrom.Instrument {
		noteTo.Instrument = append(noteTo.Instrument, GongCopyBranchInstrument(mapOrigCopy, _instrument))
	}
	for _, _beam := range noteFrom.Beam {
		noteTo.Beam = append(noteTo.Beam, GongCopyBranchBeam(mapOrigCopy, _beam))
	}
	for _, _notations := range noteFrom.Notations {
		noteTo.Notations = append(noteTo.Notations, GongCopyBranchNotations(mapOrigCopy, _notations))
	}
	for _, _lyric := range noteFrom.Lyric {
		noteTo.Lyric = append(noteTo.Lyric, GongCopyBranchLyric(mapOrigCopy, _lyric))
	}

	return
}

func GongCopyBranchNoteType(mapOrigCopy map[any]any, notetypeFrom *NoteType) (notetypeTo *NoteType) {
	var alreadyCopied bool
	notetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, notetypeFrom)
	if alreadyCopied {
		return
	}
	notetypeFrom.GongCopyBasicFields(notetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNotehead(mapOrigCopy map[any]any, noteheadFrom *Notehead) (noteheadTo *Notehead) {
	var alreadyCopied bool
	noteheadTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteheadFrom)
	if alreadyCopied {
		return
	}
	noteheadFrom.GongCopyBasicFields(noteheadTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNoteheadText(mapOrigCopy map[any]any, noteheadtextFrom *NoteheadText) (noteheadtextTo *NoteheadText) {
	var alreadyCopied bool
	noteheadtextTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteheadtextFrom)
	if alreadyCopied {
		return
	}
	noteheadtextFrom.GongCopyBasicFields(noteheadtextTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOrnaments(mapOrigCopy map[any]any, ornamentsFrom *Ornaments) (ornamentsTo *Ornaments) {
	var alreadyCopied bool
	ornamentsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ornamentsFrom)
	if alreadyCopied {
		return
	}
	ornamentsFrom.GongCopyBasicFields(ornamentsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _mordent := range ornamentsFrom.Mordent {
		ornamentsTo.Mordent = append(ornamentsTo.Mordent, GongCopyBranchMordent(mapOrigCopy, _mordent))
	}
	for _, _mordent := range ornamentsFrom.InvertedMordent {
		ornamentsTo.InvertedMordent = append(ornamentsTo.InvertedMordent, GongCopyBranchMordent(mapOrigCopy, _mordent))
	}
	for _, _tremolo := range ornamentsFrom.Tremolo {
		ornamentsTo.Tremolo = append(ornamentsTo.Tremolo, GongCopyBranchTremolo(mapOrigCopy, _tremolo))
	}
	for _, _otherplacementtext := range ornamentsFrom.OtherOrnament {
		ornamentsTo.OtherOrnament = append(ornamentsTo.OtherOrnament, GongCopyBranchOtherPlacementText(mapOrigCopy, _otherplacementtext))
	}
	for _, _accidentalmark := range ornamentsFrom.AccidentalMark {
		ornamentsTo.AccidentalMark = append(ornamentsTo.AccidentalMark, GongCopyBranchAccidentalMark(mapOrigCopy, _accidentalmark))
	}

	return
}

func GongCopyBranchOtherNotation(mapOrigCopy map[any]any, othernotationFrom *OtherNotation) (othernotationTo *OtherNotation) {
	var alreadyCopied bool
	othernotationTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, othernotationFrom)
	if alreadyCopied {
		return
	}
	othernotationFrom.GongCopyBasicFields(othernotationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOtherPlacementText(mapOrigCopy map[any]any, otherplacementtextFrom *OtherPlacementText) (otherplacementtextTo *OtherPlacementText) {
	var alreadyCopied bool
	otherplacementtextTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, otherplacementtextFrom)
	if alreadyCopied {
		return
	}
	otherplacementtextFrom.GongCopyBasicFields(otherplacementtextTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPitch(mapOrigCopy map[any]any, pitchFrom *Pitch) (pitchTo *Pitch) {
	var alreadyCopied bool
	pitchTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pitchFrom)
	if alreadyCopied {
		return
	}
	pitchFrom.GongCopyBasicFields(pitchTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPlacementText(mapOrigCopy map[any]any, placementtextFrom *PlacementText) (placementtextTo *PlacementText) {
	var alreadyCopied bool
	placementtextTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, placementtextFrom)
	if alreadyCopied {
		return
	}
	placementtextFrom.GongCopyBasicFields(placementtextTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRelease(mapOrigCopy map[any]any, releaseFrom *Release) (releaseTo *Release) {
	var alreadyCopied bool
	releaseTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, releaseFrom)
	if alreadyCopied {
		return
	}
	releaseFrom.GongCopyBasicFields(releaseTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRest(mapOrigCopy map[any]any, restFrom *Rest) (restTo *Rest) {
	var alreadyCopied bool
	restTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, restFrom)
	if alreadyCopied {
		return
	}
	restFrom.GongCopyBasicFields(restTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSlide(mapOrigCopy map[any]any, slideFrom *Slide) (slideTo *Slide) {
	var alreadyCopied bool
	slideTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, slideFrom)
	if alreadyCopied {
		return
	}
	slideFrom.GongCopyBasicFields(slideTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSlur(mapOrigCopy map[any]any, slurFrom *Slur) (slurTo *Slur) {
	var alreadyCopied bool
	slurTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, slurFrom)
	if alreadyCopied {
		return
	}
	slurFrom.GongCopyBasicFields(slurTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStem(mapOrigCopy map[any]any, stemFrom *Stem) (stemTo *Stem) {
	var alreadyCopied bool
	stemTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stemFrom)
	if alreadyCopied {
		return
	}
	stemFrom.GongCopyBasicFields(stemTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStrongAccent(mapOrigCopy map[any]any, strongaccentFrom *StrongAccent) (strongaccentTo *StrongAccent) {
	var alreadyCopied bool
	strongaccentTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, strongaccentFrom)
	if alreadyCopied {
		return
	}
	strongaccentFrom.GongCopyBasicFields(strongaccentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTap(mapOrigCopy map[any]any, tapFrom *Tap) (tapTo *Tap) {
	var alreadyCopied bool
	tapTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tapFrom)
	if alreadyCopied {
		return
	}
	tapFrom.GongCopyBasicFields(tapTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTechnical(mapOrigCopy map[any]any, technicalFrom *Technical) (technicalTo *Technical) {
	var alreadyCopied bool
	technicalTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, technicalFrom)
	if alreadyCopied {
		return
	}
	technicalFrom.GongCopyBasicFields(technicalTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _harmonic := range technicalFrom.Harmonic {
		technicalTo.Harmonic = append(technicalTo.Harmonic, GongCopyBranchHarmonic(mapOrigCopy, _harmonic))
	}
	for _, _placementtext := range technicalFrom.Pluck {
		technicalTo.Pluck = append(technicalTo.Pluck, GongCopyBranchPlacementText(mapOrigCopy, _placementtext))
	}
	for _, _hammeronpulloff := range technicalFrom.HammerOn {
		technicalTo.HammerOn = append(technicalTo.HammerOn, GongCopyBranchHammerOnPullOff(mapOrigCopy, _hammeronpulloff))
	}
	for _, _hammeronpulloff := range technicalFrom.PullOff {
		technicalTo.PullOff = append(technicalTo.PullOff, GongCopyBranchHammerOnPullOff(mapOrigCopy, _hammeronpulloff))
	}
	for _, _bend := range technicalFrom.Bend {
		technicalTo.Bend = append(technicalTo.Bend, GongCopyBranchBend(mapOrigCopy, _bend))
	}
	for _, _tap := range technicalFrom.Tap {
		technicalTo.Tap = append(technicalTo.Tap, GongCopyBranchTap(mapOrigCopy, _tap))
	}
	for _, _heeltoe := range technicalFrom.Heel {
		technicalTo.Heel = append(technicalTo.Heel, GongCopyBranchHeelToe(mapOrigCopy, _heeltoe))
	}
	for _, _heeltoe := range technicalFrom.Toe {
		technicalTo.Toe = append(technicalTo.Toe, GongCopyBranchHeelToe(mapOrigCopy, _heeltoe))
	}
	for _, _hole := range technicalFrom.Hole {
		technicalTo.Hole = append(technicalTo.Hole, GongCopyBranchHole(mapOrigCopy, _hole))
	}
	for _, _arrow := range technicalFrom.Arrow {
		technicalTo.Arrow = append(technicalTo.Arrow, GongCopyBranchArrow(mapOrigCopy, _arrow))
	}
	for _, _handbell := range technicalFrom.Handbell {
		technicalTo.Handbell = append(technicalTo.Handbell, GongCopyBranchHandbell(mapOrigCopy, _handbell))
	}
	for _, _harmonmute := range technicalFrom.HarmonMute {
		technicalTo.HarmonMute = append(technicalTo.HarmonMute, GongCopyBranchHarmonMute(mapOrigCopy, _harmonmute))
	}
	for _, _otherplacementtext := range technicalFrom.OtherTechnical {
		technicalTo.OtherTechnical = append(technicalTo.OtherTechnical, GongCopyBranchOtherPlacementText(mapOrigCopy, _otherplacementtext))
	}

	return
}

func GongCopyBranchTextElementData(mapOrigCopy map[any]any, textelementdataFrom *TextElementData) (textelementdataTo *TextElementData) {
	var alreadyCopied bool
	textelementdataTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, textelementdataFrom)
	if alreadyCopied {
		return
	}
	textelementdataFrom.GongCopyBasicFields(textelementdataTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTie(mapOrigCopy map[any]any, tieFrom *Tie) (tieTo *Tie) {
	var alreadyCopied bool
	tieTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tieFrom)
	if alreadyCopied {
		return
	}
	tieFrom.GongCopyBasicFields(tieTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTied(mapOrigCopy map[any]any, tiedFrom *Tied) (tiedTo *Tied) {
	var alreadyCopied bool
	tiedTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tiedFrom)
	if alreadyCopied {
		return
	}
	tiedFrom.GongCopyBasicFields(tiedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTremolo(mapOrigCopy map[any]any, tremoloFrom *Tremolo) (tremoloTo *Tremolo) {
	var alreadyCopied bool
	tremoloTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tremoloFrom)
	if alreadyCopied {
		return
	}
	tremoloFrom.GongCopyBasicFields(tremoloTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTuplet(mapOrigCopy map[any]any, tupletFrom *Tuplet) (tupletTo *Tuplet) {
	var alreadyCopied bool
	tupletTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tupletFrom)
	if alreadyCopied {
		return
	}
	tupletFrom.GongCopyBasicFields(tupletTo)

	//insertion point for the staging of instances referenced by pointers
	if tupletFrom.TupletActual != nil {
		tupletTo.TupletActual = GongCopyBranchTupletPortion(mapOrigCopy, tupletFrom.TupletActual)
	}
	if tupletFrom.TupletNormal != nil {
		tupletTo.TupletNormal = GongCopyBranchTupletPortion(mapOrigCopy, tupletFrom.TupletNormal)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTupletDot(mapOrigCopy map[any]any, tupletdotFrom *TupletDot) (tupletdotTo *TupletDot) {
	var alreadyCopied bool
	tupletdotTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tupletdotFrom)
	if alreadyCopied {
		return
	}
	tupletdotFrom.GongCopyBasicFields(tupletdotTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTupletNumber(mapOrigCopy map[any]any, tupletnumberFrom *TupletNumber) (tupletnumberTo *TupletNumber) {
	var alreadyCopied bool
	tupletnumberTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tupletnumberFrom)
	if alreadyCopied {
		return
	}
	tupletnumberFrom.GongCopyBasicFields(tupletnumberTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTupletPortion(mapOrigCopy map[any]any, tupletportionFrom *TupletPortion) (tupletportionTo *TupletPortion) {
	var alreadyCopied bool
	tupletportionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tupletportionFrom)
	if alreadyCopied {
		return
	}
	tupletportionFrom.GongCopyBasicFields(tupletportionTo)

	//insertion point for the staging of instances referenced by pointers
	if tupletportionFrom.TupletNumber != nil {
		tupletportionTo.TupletNumber = GongCopyBranchTupletNumber(mapOrigCopy, tupletportionFrom.TupletNumber)
	}
	if tupletportionFrom.TupletType != nil {
		tupletportionTo.TupletType = GongCopyBranchTupletType(mapOrigCopy, tupletportionFrom.TupletType)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tupletdot := range tupletportionFrom.TupletDot {
		tupletportionTo.TupletDot = append(tupletportionTo.TupletDot, GongCopyBranchTupletDot(mapOrigCopy, _tupletdot))
	}

	return
}

func GongCopyBranchTupletType(mapOrigCopy map[any]any, tuplettypeFrom *TupletType) (tuplettypeTo *TupletType) {
	var alreadyCopied bool
	tuplettypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tuplettypeFrom)
	if alreadyCopied {
		return
	}
	tuplettypeFrom.GongCopyBasicFields(tuplettypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchUnpitched(mapOrigCopy map[any]any, unpitchedFrom *Unpitched) (unpitchedTo *Unpitched) {
	var alreadyCopied bool
	unpitchedTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, unpitchedFrom)
	if alreadyCopied {
		return
	}
	unpitchedFrom.GongCopyBasicFields(unpitchedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchWait(mapOrigCopy map[any]any, waitFrom *Wait) (waitTo *Wait) {
	var alreadyCopied bool
	waitTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, waitFrom)
	if alreadyCopied {
		return
	}
	waitFrom.GongCopyBasicFields(waitTo)

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
func (accidental *Accidental) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(accidental) {
		return
	}

	accidental.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (accidentalmark *AccidentalMark) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(accidentalmark) {
		return
	}

	accidentalmark.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (arpeggiate *Arpeggiate) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(arpeggiate) {
		return
	}

	arpeggiate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (arrow *Arrow) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(arrow) {
		return
	}

	arrow.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (articulations *Articulations) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(articulations) {
		return
	}

	articulations.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _strongaccent := range articulations.StrongAccent {
		stage.UnstageBranch(_strongaccent)
	}
	for _, _emptyline := range articulations.Scoop {
		stage.UnstageBranch(_emptyline)
	}
	for _, _emptyline := range articulations.Plop {
		stage.UnstageBranch(_emptyline)
	}
	for _, _emptyline := range articulations.Doit {
		stage.UnstageBranch(_emptyline)
	}
	for _, _emptyline := range articulations.Falloff {
		stage.UnstageBranch(_emptyline)
	}
	for _, _breathmark := range articulations.BreathMark {
		stage.UnstageBranch(_breathmark)
	}
	for _, _caesura := range articulations.Caesura {
		stage.UnstageBranch(_caesura)
	}
	for _, _otherplacementtext := range articulations.OtherArticulation {
		stage.UnstageBranch(_otherplacementtext)
	}

}

func (assess *Assess) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(assess) {
		return
	}

	assess.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (backup *Backup) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(backup) {
		return
	}

	backup.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (beam *Beam) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(beam) {
		return
	}

	beam.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bend *Bend) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(bend) {
		return
	}

	bend.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if bend.Release != nil {
		stage.UnstageBranch(bend.Release)
	}
	if bend.WithBar != nil {
		stage.UnstageBranch(bend.WithBar)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (breathmark *BreathMark) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(breathmark) {
		return
	}

	breathmark.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (caesura *Caesura) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(caesura) {
		return
	}

	caesura.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (elision *Elision) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(elision) {
		return
	}

	elision.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (emptyline *EmptyLine) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(emptyline) {
		return
	}

	emptyline.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (extend *Extend) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(extend) {
		return
	}

	extend.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (figure *Figure) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(figure) {
		return
	}

	figure.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if figure.Extend != nil {
		stage.UnstageBranch(figure.Extend)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (figuredbass *FiguredBass) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(figuredbass) {
		return
	}

	figuredbass.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _figure := range figuredbass.Figure {
		stage.UnstageBranch(_figure)
	}

}

func (forward *Forward) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(forward) {
		return
	}

	forward.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (glissando *Glissando) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(glissando) {
		return
	}

	glissando.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (grace *Grace) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(grace) {
		return
	}

	grace.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (hammeronpulloff *HammerOnPullOff) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(hammeronpulloff) {
		return
	}

	hammeronpulloff.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (handbell *Handbell) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(handbell) {
		return
	}

	handbell.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (harmonclosed *HarmonClosed) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(harmonclosed) {
		return
	}

	harmonclosed.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (harmonmute *HarmonMute) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(harmonmute) {
		return
	}

	harmonmute.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if harmonmute.HarmonClosed != nil {
		stage.UnstageBranch(harmonmute.HarmonClosed)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (harmonic *Harmonic) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(harmonic) {
		return
	}

	harmonic.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (heeltoe *HeelToe) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(heeltoe) {
		return
	}

	heeltoe.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (hole *Hole) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(hole) {
		return
	}

	hole.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if hole.HoleClosed != nil {
		stage.UnstageBranch(hole.HoleClosed)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (holeclosed *HoleClosed) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(holeclosed) {
		return
	}

	holeclosed.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (instrument *Instrument) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(instrument) {
		return
	}

	instrument.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (listen *Listen) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(listen) {
		return
	}

	listen.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assess := range listen.Assess {
		stage.UnstageBranch(_assess)
	}
	for _, _wait := range listen.Wait {
		stage.UnstageBranch(_wait)
	}

}

func (lyric *Lyric) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(lyric) {
		return
	}

	lyric.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if lyric.Text != nil {
		stage.UnstageBranch(lyric.Text)
	}
	if lyric.Extend != nil {
		stage.UnstageBranch(lyric.Extend)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _elision := range lyric.Elision {
		stage.UnstageBranch(_elision)
	}

}

func (mordent *Mordent) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(mordent) {
		return
	}

	mordent.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (nonarpeggiate *NonArpeggiate) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(nonarpeggiate) {
		return
	}

	nonarpeggiate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notations *Notations) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(notations) {
		return
	}

	notations.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tied := range notations.Tied {
		stage.UnstageBranch(_tied)
	}
	for _, _slur := range notations.Slur {
		stage.UnstageBranch(_slur)
	}
	for _, _tuplet := range notations.Tuplet {
		stage.UnstageBranch(_tuplet)
	}
	for _, _glissando := range notations.Glissando {
		stage.UnstageBranch(_glissando)
	}
	for _, _slide := range notations.Slide {
		stage.UnstageBranch(_slide)
	}
	for _, _ornaments := range notations.Ornaments {
		stage.UnstageBranch(_ornaments)
	}
	for _, _technical := range notations.Technical {
		stage.UnstageBranch(_technical)
	}
	for _, _articulations := range notations.Articulations {
		stage.UnstageBranch(_articulations)
	}
	for _, _arpeggiate := range notations.Arpeggiate {
		stage.UnstageBranch(_arpeggiate)
	}
	for _, _nonarpeggiate := range notations.NonArpeggiate {
		stage.UnstageBranch(_nonarpeggiate)
	}
	for _, _accidentalmark := range notations.AccidentalMark {
		stage.UnstageBranch(_accidentalmark)
	}
	for _, _othernotation := range notations.OtherNotation {
		stage.UnstageBranch(_othernotation)
	}

}

func (note *Note) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(note) {
		return
	}

	note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if note.Grace != nil {
		stage.UnstageBranch(note.Grace)
	}
	if note.Pitch != nil {
		stage.UnstageBranch(note.Pitch)
	}
	if note.Unpitched != nil {
		stage.UnstageBranch(note.Unpitched)
	}
	if note.Rest != nil {
		stage.UnstageBranch(note.Rest)
	}
	if note.Type != nil {
		stage.UnstageBranch(note.Type)
	}
	if note.Accidental != nil {
		stage.UnstageBranch(note.Accidental)
	}
	if note.Stem != nil {
		stage.UnstageBranch(note.Stem)
	}
	if note.Notehead != nil {
		stage.UnstageBranch(note.Notehead)
	}
	if note.NoteheadText != nil {
		stage.UnstageBranch(note.NoteheadText)
	}
	if note.Listen != nil {
		stage.UnstageBranch(note.Listen)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tie := range note.Tie {
		stage.UnstageBranch(_tie)
	}
	for _, _instrument := range note.Instrument {
		stage.UnstageBranch(_instrument)
	}
	for _, _beam := range note.Beam {
		stage.UnstageBranch(_beam)
	}
	for _, _notations := range note.Notations {
		stage.UnstageBranch(_notations)
	}
	for _, _lyric := range note.Lyric {
		stage.UnstageBranch(_lyric)
	}

}

func (notetype *NoteType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(notetype) {
		return
	}

	notetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notehead *Notehead) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(notehead) {
		return
	}

	notehead.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (noteheadtext *NoteheadText) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(noteheadtext) {
		return
	}

	noteheadtext.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ornaments *Ornaments) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ornaments) {
		return
	}

	ornaments.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _mordent := range ornaments.Mordent {
		stage.UnstageBranch(_mordent)
	}
	for _, _mordent := range ornaments.InvertedMordent {
		stage.UnstageBranch(_mordent)
	}
	for _, _tremolo := range ornaments.Tremolo {
		stage.UnstageBranch(_tremolo)
	}
	for _, _otherplacementtext := range ornaments.OtherOrnament {
		stage.UnstageBranch(_otherplacementtext)
	}
	for _, _accidentalmark := range ornaments.AccidentalMark {
		stage.UnstageBranch(_accidentalmark)
	}

}

func (othernotation *OtherNotation) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(othernotation) {
		return
	}

	othernotation.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (otherplacementtext *OtherPlacementText) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(otherplacementtext) {
		return
	}

	otherplacementtext.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pitch *Pitch) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(pitch) {
		return
	}

	pitch.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (placementtext *PlacementText) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(placementtext) {
		return
	}

	placementtext.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (release *Release) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(release) {
		return
	}

	release.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rest *Rest) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(rest) {
		return
	}

	rest.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (slide *Slide) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(slide) {
		return
	}

	slide.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (slur *Slur) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(slur) {
		return
	}

	slur.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stem *Stem) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stem) {
		return
	}

	stem.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (strongaccent *StrongAccent) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(strongaccent) {
		return
	}

	strongaccent.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tap *Tap) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tap) {
		return
	}

	tap.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (technical *Technical) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(technical) {
		return
	}

	technical.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _harmonic := range technical.Harmonic {
		stage.UnstageBranch(_harmonic)
	}
	for _, _placementtext := range technical.Pluck {
		stage.UnstageBranch(_placementtext)
	}
	for _, _hammeronpulloff := range technical.HammerOn {
		stage.UnstageBranch(_hammeronpulloff)
	}
	for _, _hammeronpulloff := range technical.PullOff {
		stage.UnstageBranch(_hammeronpulloff)
	}
	for _, _bend := range technical.Bend {
		stage.UnstageBranch(_bend)
	}
	for _, _tap := range technical.Tap {
		stage.UnstageBranch(_tap)
	}
	for _, _heeltoe := range technical.Heel {
		stage.UnstageBranch(_heeltoe)
	}
	for _, _heeltoe := range technical.Toe {
		stage.UnstageBranch(_heeltoe)
	}
	for _, _hole := range technical.Hole {
		stage.UnstageBranch(_hole)
	}
	for _, _arrow := range technical.Arrow {
		stage.UnstageBranch(_arrow)
	}
	for _, _handbell := range technical.Handbell {
		stage.UnstageBranch(_handbell)
	}
	for _, _harmonmute := range technical.HarmonMute {
		stage.UnstageBranch(_harmonmute)
	}
	for _, _otherplacementtext := range technical.OtherTechnical {
		stage.UnstageBranch(_otherplacementtext)
	}

}

func (textelementdata *TextElementData) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(textelementdata) {
		return
	}

	textelementdata.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tie *Tie) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tie) {
		return
	}

	tie.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tied *Tied) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tied) {
		return
	}

	tied.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tremolo *Tremolo) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tremolo) {
		return
	}

	tremolo.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tuplet *Tuplet) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tuplet) {
		return
	}

	tuplet.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tuplet.TupletActual != nil {
		stage.UnstageBranch(tuplet.TupletActual)
	}
	if tuplet.TupletNormal != nil {
		stage.UnstageBranch(tuplet.TupletNormal)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tupletdot *TupletDot) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tupletdot) {
		return
	}

	tupletdot.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tupletnumber *TupletNumber) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tupletnumber) {
		return
	}

	tupletnumber.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tupletportion *TupletPortion) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tupletportion) {
		return
	}

	tupletportion.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tupletportion.TupletNumber != nil {
		stage.UnstageBranch(tupletportion.TupletNumber)
	}
	if tupletportion.TupletType != nil {
		stage.UnstageBranch(tupletportion.TupletType)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tupletdot := range tupletportion.TupletDot {
		stage.UnstageBranch(_tupletdot)
	}

}

func (tuplettype *TupletType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tuplettype) {
		return
	}

	tuplettype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (unpitched *Unpitched) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(unpitched) {
		return
	}

	unpitched.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (wait *Wait) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(wait) {
		return
	}

	wait.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *Accidental) GongReconstructPointersFromReferences(stage *Stage, instance *Accidental) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *AccidentalMark) GongReconstructPointersFromReferences(stage *Stage, instance *AccidentalMark) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Arpeggiate) GongReconstructPointersFromReferences(stage *Stage, instance *Arpeggiate) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Arrow) GongReconstructPointersFromReferences(stage *Stage, instance *Arrow) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Articulations) GongReconstructPointersFromReferences(stage *Stage, instance *Articulations) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.StrongAccent, stage.StrongAccents_reference, instance.StrongAccent)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Scoop, stage.EmptyLines_reference, instance.Scoop)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Plop, stage.EmptyLines_reference, instance.Plop)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Doit, stage.EmptyLines_reference, instance.Doit)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Falloff, stage.EmptyLines_reference, instance.Falloff)
	__gong__reconstructSliceOfPointersFromReferences(&reference.BreathMark, stage.BreathMarks_reference, instance.BreathMark)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Caesura, stage.Caesuras_reference, instance.Caesura)
	__gong__reconstructSliceOfPointersFromReferences(&reference.OtherArticulation, stage.OtherPlacementTexts_reference, instance.OtherArticulation)
}

func (reference *Assess) GongReconstructPointersFromReferences(stage *Stage, instance *Assess) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Backup) GongReconstructPointersFromReferences(stage *Stage, instance *Backup) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Beam) GongReconstructPointersFromReferences(stage *Stage, instance *Beam) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Bend) GongReconstructPointersFromReferences(stage *Stage, instance *Bend) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Release, stage.Releases_reference, instance.Release)
	__gong__reconstructPointer(&reference.WithBar, stage.PlacementTexts_reference, instance.WithBar)
	// insertion point for slice of pointers field
}

func (reference *BreathMark) GongReconstructPointersFromReferences(stage *Stage, instance *BreathMark) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Caesura) GongReconstructPointersFromReferences(stage *Stage, instance *Caesura) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Elision) GongReconstructPointersFromReferences(stage *Stage, instance *Elision) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EmptyLine) GongReconstructPointersFromReferences(stage *Stage, instance *EmptyLine) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Extend) GongReconstructPointersFromReferences(stage *Stage, instance *Extend) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Figure) GongReconstructPointersFromReferences(stage *Stage, instance *Figure) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Extend, stage.Extends_reference, instance.Extend)
	// insertion point for slice of pointers field
}

func (reference *FiguredBass) GongReconstructPointersFromReferences(stage *Stage, instance *FiguredBass) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Figure, stage.Figures_reference, instance.Figure)
}

func (reference *Forward) GongReconstructPointersFromReferences(stage *Stage, instance *Forward) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Glissando) GongReconstructPointersFromReferences(stage *Stage, instance *Glissando) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Grace) GongReconstructPointersFromReferences(stage *Stage, instance *Grace) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *HammerOnPullOff) GongReconstructPointersFromReferences(stage *Stage, instance *HammerOnPullOff) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Handbell) GongReconstructPointersFromReferences(stage *Stage, instance *Handbell) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *HarmonClosed) GongReconstructPointersFromReferences(stage *Stage, instance *HarmonClosed) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *HarmonMute) GongReconstructPointersFromReferences(stage *Stage, instance *HarmonMute) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.HarmonClosed, stage.HarmonCloseds_reference, instance.HarmonClosed)
	// insertion point for slice of pointers field
}

func (reference *Harmonic) GongReconstructPointersFromReferences(stage *Stage, instance *Harmonic) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *HeelToe) GongReconstructPointersFromReferences(stage *Stage, instance *HeelToe) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Hole) GongReconstructPointersFromReferences(stage *Stage, instance *Hole) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.HoleClosed, stage.HoleCloseds_reference, instance.HoleClosed)
	// insertion point for slice of pointers field
}

func (reference *HoleClosed) GongReconstructPointersFromReferences(stage *Stage, instance *HoleClosed) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Instrument) GongReconstructPointersFromReferences(stage *Stage, instance *Instrument) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Listen) GongReconstructPointersFromReferences(stage *Stage, instance *Listen) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Assess, stage.Assesss_reference, instance.Assess)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Wait, stage.Waits_reference, instance.Wait)
}

func (reference *Lyric) GongReconstructPointersFromReferences(stage *Stage, instance *Lyric) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Text, stage.TextElementDatas_reference, instance.Text)
	__gong__reconstructPointer(&reference.Extend, stage.Extends_reference, instance.Extend)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Elision, stage.Elisions_reference, instance.Elision)
}

func (reference *Mordent) GongReconstructPointersFromReferences(stage *Stage, instance *Mordent) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *NonArpeggiate) GongReconstructPointersFromReferences(stage *Stage, instance *NonArpeggiate) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Notations) GongReconstructPointersFromReferences(stage *Stage, instance *Notations) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Tied, stage.Tieds_reference, instance.Tied)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Slur, stage.Slurs_reference, instance.Slur)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Tuplet, stage.Tuplets_reference, instance.Tuplet)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Glissando, stage.Glissandos_reference, instance.Glissando)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Slide, stage.Slides_reference, instance.Slide)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Ornaments, stage.Ornamentss_reference, instance.Ornaments)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Technical, stage.Technicals_reference, instance.Technical)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Articulations, stage.Articulationss_reference, instance.Articulations)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Arpeggiate, stage.Arpeggiates_reference, instance.Arpeggiate)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NonArpeggiate, stage.NonArpeggiates_reference, instance.NonArpeggiate)
	__gong__reconstructSliceOfPointersFromReferences(&reference.AccidentalMark, stage.AccidentalMarks_reference, instance.AccidentalMark)
	__gong__reconstructSliceOfPointersFromReferences(&reference.OtherNotation, stage.OtherNotations_reference, instance.OtherNotation)
}

func (reference *Note) GongReconstructPointersFromReferences(stage *Stage, instance *Note) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Grace, stage.Graces_reference, instance.Grace)
	__gong__reconstructPointer(&reference.Pitch, stage.Pitchs_reference, instance.Pitch)
	__gong__reconstructPointer(&reference.Unpitched, stage.Unpitcheds_reference, instance.Unpitched)
	__gong__reconstructPointer(&reference.Rest, stage.Rests_reference, instance.Rest)
	__gong__reconstructPointer(&reference.Type, stage.NoteTypes_reference, instance.Type)
	__gong__reconstructPointer(&reference.Accidental, stage.Accidentals_reference, instance.Accidental)
	__gong__reconstructPointer(&reference.Stem, stage.Stems_reference, instance.Stem)
	__gong__reconstructPointer(&reference.Notehead, stage.Noteheads_reference, instance.Notehead)
	__gong__reconstructPointer(&reference.NoteheadText, stage.NoteheadTexts_reference, instance.NoteheadText)
	__gong__reconstructPointer(&reference.Listen, stage.Listens_reference, instance.Listen)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Tie, stage.Ties_reference, instance.Tie)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Instrument, stage.Instruments_reference, instance.Instrument)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Beam, stage.Beams_reference, instance.Beam)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Notations, stage.Notationss_reference, instance.Notations)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Lyric, stage.Lyrics_reference, instance.Lyric)
}

func (reference *NoteType) GongReconstructPointersFromReferences(stage *Stage, instance *NoteType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Notehead) GongReconstructPointersFromReferences(stage *Stage, instance *Notehead) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *NoteheadText) GongReconstructPointersFromReferences(stage *Stage, instance *NoteheadText) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Ornaments) GongReconstructPointersFromReferences(stage *Stage, instance *Ornaments) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Mordent, stage.Mordents_reference, instance.Mordent)
	__gong__reconstructSliceOfPointersFromReferences(&reference.InvertedMordent, stage.Mordents_reference, instance.InvertedMordent)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Tremolo, stage.Tremolos_reference, instance.Tremolo)
	__gong__reconstructSliceOfPointersFromReferences(&reference.OtherOrnament, stage.OtherPlacementTexts_reference, instance.OtherOrnament)
	__gong__reconstructSliceOfPointersFromReferences(&reference.AccidentalMark, stage.AccidentalMarks_reference, instance.AccidentalMark)
}

func (reference *OtherNotation) GongReconstructPointersFromReferences(stage *Stage, instance *OtherNotation) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *OtherPlacementText) GongReconstructPointersFromReferences(stage *Stage, instance *OtherPlacementText) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Pitch) GongReconstructPointersFromReferences(stage *Stage, instance *Pitch) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PlacementText) GongReconstructPointersFromReferences(stage *Stage, instance *PlacementText) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Release) GongReconstructPointersFromReferences(stage *Stage, instance *Release) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Rest) GongReconstructPointersFromReferences(stage *Stage, instance *Rest) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Slide) GongReconstructPointersFromReferences(stage *Stage, instance *Slide) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Slur) GongReconstructPointersFromReferences(stage *Stage, instance *Slur) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Stem) GongReconstructPointersFromReferences(stage *Stage, instance *Stem) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StrongAccent) GongReconstructPointersFromReferences(stage *Stage, instance *StrongAccent) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Tap) GongReconstructPointersFromReferences(stage *Stage, instance *Tap) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Technical) GongReconstructPointersFromReferences(stage *Stage, instance *Technical) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Harmonic, stage.Harmonics_reference, instance.Harmonic)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Pluck, stage.PlacementTexts_reference, instance.Pluck)
	__gong__reconstructSliceOfPointersFromReferences(&reference.HammerOn, stage.HammerOnPullOffs_reference, instance.HammerOn)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PullOff, stage.HammerOnPullOffs_reference, instance.PullOff)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Bend, stage.Bends_reference, instance.Bend)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Tap, stage.Taps_reference, instance.Tap)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Heel, stage.HeelToes_reference, instance.Heel)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Toe, stage.HeelToes_reference, instance.Toe)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Hole, stage.Holes_reference, instance.Hole)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Arrow, stage.Arrows_reference, instance.Arrow)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Handbell, stage.Handbells_reference, instance.Handbell)
	__gong__reconstructSliceOfPointersFromReferences(&reference.HarmonMute, stage.HarmonMutes_reference, instance.HarmonMute)
	__gong__reconstructSliceOfPointersFromReferences(&reference.OtherTechnical, stage.OtherPlacementTexts_reference, instance.OtherTechnical)
}

func (reference *TextElementData) GongReconstructPointersFromReferences(stage *Stage, instance *TextElementData) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Tie) GongReconstructPointersFromReferences(stage *Stage, instance *Tie) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Tied) GongReconstructPointersFromReferences(stage *Stage, instance *Tied) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Tremolo) GongReconstructPointersFromReferences(stage *Stage, instance *Tremolo) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Tuplet) GongReconstructPointersFromReferences(stage *Stage, instance *Tuplet) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.TupletActual, stage.TupletPortions_reference, instance.TupletActual)
	__gong__reconstructPointer(&reference.TupletNormal, stage.TupletPortions_reference, instance.TupletNormal)
	// insertion point for slice of pointers field
}

func (reference *TupletDot) GongReconstructPointersFromReferences(stage *Stage, instance *TupletDot) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TupletNumber) GongReconstructPointersFromReferences(stage *Stage, instance *TupletNumber) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TupletPortion) GongReconstructPointersFromReferences(stage *Stage, instance *TupletPortion) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.TupletNumber, stage.TupletNumbers_reference, instance.TupletNumber)
	__gong__reconstructPointer(&reference.TupletType, stage.TupletTypes_reference, instance.TupletType)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.TupletDot, stage.TupletDots_reference, instance.TupletDot)
}

func (reference *TupletType) GongReconstructPointersFromReferences(stage *Stage, instance *TupletType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Unpitched) GongReconstructPointersFromReferences(stage *Stage, instance *Unpitched) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Wait) GongReconstructPointersFromReferences(stage *Stage, instance *Wait) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *Accidental) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *AccidentalMark) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Arpeggiate) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Arrow) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Articulations) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.StrongAccent, stage.StrongAccents_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Scoop, stage.EmptyLines_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Plop, stage.EmptyLines_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Doit, stage.EmptyLines_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Falloff, stage.EmptyLines_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.BreathMark, stage.BreathMarks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Caesura, stage.Caesuras_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.OtherArticulation, stage.OtherPlacementTexts_instance)
}

func (reference *Assess) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Backup) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Beam) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Bend) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Release, stage.Releases_instance)
	__gong__reconstructPointerFromInstance(&reference.WithBar, stage.PlacementTexts_instance)
	// insertion point for slice of pointers fields
}

func (reference *BreathMark) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Caesura) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Elision) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EmptyLine) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Extend) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Figure) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Extend, stage.Extends_instance)
	// insertion point for slice of pointers fields
}

func (reference *FiguredBass) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Figure, stage.Figures_instance)
}

func (reference *Forward) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Glissando) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Grace) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *HammerOnPullOff) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Handbell) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *HarmonClosed) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *HarmonMute) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.HarmonClosed, stage.HarmonCloseds_instance)
	// insertion point for slice of pointers fields
}

func (reference *Harmonic) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *HeelToe) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Hole) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.HoleClosed, stage.HoleCloseds_instance)
	// insertion point for slice of pointers fields
}

func (reference *HoleClosed) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Instrument) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Listen) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Assess, stage.Assesss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Wait, stage.Waits_instance)
}

func (reference *Lyric) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Text, stage.TextElementDatas_instance)
	__gong__reconstructPointerFromInstance(&reference.Extend, stage.Extends_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Elision, stage.Elisions_instance)
}

func (reference *Mordent) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *NonArpeggiate) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Notations) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Tied, stage.Tieds_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Slur, stage.Slurs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Tuplet, stage.Tuplets_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Glissando, stage.Glissandos_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Slide, stage.Slides_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Ornaments, stage.Ornamentss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Technical, stage.Technicals_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Articulations, stage.Articulationss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Arpeggiate, stage.Arpeggiates_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NonArpeggiate, stage.NonArpeggiates_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.AccidentalMark, stage.AccidentalMarks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.OtherNotation, stage.OtherNotations_instance)
}

func (reference *Note) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Grace, stage.Graces_instance)
	__gong__reconstructPointerFromInstance(&reference.Pitch, stage.Pitchs_instance)
	__gong__reconstructPointerFromInstance(&reference.Unpitched, stage.Unpitcheds_instance)
	__gong__reconstructPointerFromInstance(&reference.Rest, stage.Rests_instance)
	__gong__reconstructPointerFromInstance(&reference.Type, stage.NoteTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.Accidental, stage.Accidentals_instance)
	__gong__reconstructPointerFromInstance(&reference.Stem, stage.Stems_instance)
	__gong__reconstructPointerFromInstance(&reference.Notehead, stage.Noteheads_instance)
	__gong__reconstructPointerFromInstance(&reference.NoteheadText, stage.NoteheadTexts_instance)
	__gong__reconstructPointerFromInstance(&reference.Listen, stage.Listens_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Tie, stage.Ties_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Instrument, stage.Instruments_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Beam, stage.Beams_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Notations, stage.Notationss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Lyric, stage.Lyrics_instance)
}

func (reference *NoteType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Notehead) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *NoteheadText) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Ornaments) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Mordent, stage.Mordents_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.InvertedMordent, stage.Mordents_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Tremolo, stage.Tremolos_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.OtherOrnament, stage.OtherPlacementTexts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.AccidentalMark, stage.AccidentalMarks_instance)
}

func (reference *OtherNotation) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *OtherPlacementText) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Pitch) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PlacementText) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Release) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Rest) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Slide) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Slur) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Stem) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StrongAccent) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Tap) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Technical) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Harmonic, stage.Harmonics_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Pluck, stage.PlacementTexts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.HammerOn, stage.HammerOnPullOffs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PullOff, stage.HammerOnPullOffs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Bend, stage.Bends_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Tap, stage.Taps_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Heel, stage.HeelToes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Toe, stage.HeelToes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Hole, stage.Holes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Arrow, stage.Arrows_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Handbell, stage.Handbells_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.HarmonMute, stage.HarmonMutes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.OtherTechnical, stage.OtherPlacementTexts_instance)
}

func (reference *TextElementData) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Tie) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Tied) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Tremolo) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Tuplet) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.TupletActual, stage.TupletPortions_instance)
	__gong__reconstructPointerFromInstance(&reference.TupletNormal, stage.TupletPortions_instance)
	// insertion point for slice of pointers fields
}

func (reference *TupletDot) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TupletNumber) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TupletPortion) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.TupletNumber, stage.TupletNumbers_instance)
	__gong__reconstructPointerFromInstance(&reference.TupletType, stage.TupletTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.TupletDot, stage.TupletDots_instance)
}

func (reference *TupletType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Unpitched) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Wait) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (accidental *Accidental) GongDiff(stage *Stage, accidentalOther *Accidental) (diffs []string) {
	// insertion point for field diffs
	if accidental.Name != accidentalOther.Name {
		diffs = append(diffs, accidental.GongMarshallField(stage, "Name"))
	}
	if accidental.Smufl != accidentalOther.Smufl {
		diffs = append(diffs, accidental.GongMarshallField(stage, "Smufl"))
	}
	if accidental.DefaultX != accidentalOther.DefaultX {
		diffs = append(diffs, accidental.GongMarshallField(stage, "DefaultX"))
	}
	if accidental.DefaultY != accidentalOther.DefaultY {
		diffs = append(diffs, accidental.GongMarshallField(stage, "DefaultY"))
	}
	if accidental.RelativeX != accidentalOther.RelativeX {
		diffs = append(diffs, accidental.GongMarshallField(stage, "RelativeX"))
	}
	if accidental.RelativeY != accidentalOther.RelativeY {
		diffs = append(diffs, accidental.GongMarshallField(stage, "RelativeY"))
	}
	if accidental.FontFamily != accidentalOther.FontFamily {
		diffs = append(diffs, accidental.GongMarshallField(stage, "FontFamily"))
	}
	if accidental.FontSize != accidentalOther.FontSize {
		diffs = append(diffs, accidental.GongMarshallField(stage, "FontSize"))
	}
	if accidental.Color != accidentalOther.Color {
		diffs = append(diffs, accidental.GongMarshallField(stage, "Color"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (accidentalmark *AccidentalMark) GongDiff(stage *Stage, accidentalmarkOther *AccidentalMark) (diffs []string) {
	// insertion point for field diffs
	if accidentalmark.Name != accidentalmarkOther.Name {
		diffs = append(diffs, accidentalmark.GongMarshallField(stage, "Name"))
	}
	if accidentalmark.Smufl != accidentalmarkOther.Smufl {
		diffs = append(diffs, accidentalmark.GongMarshallField(stage, "Smufl"))
	}
	if accidentalmark.DefaultX != accidentalmarkOther.DefaultX {
		diffs = append(diffs, accidentalmark.GongMarshallField(stage, "DefaultX"))
	}
	if accidentalmark.DefaultY != accidentalmarkOther.DefaultY {
		diffs = append(diffs, accidentalmark.GongMarshallField(stage, "DefaultY"))
	}
	if accidentalmark.RelativeX != accidentalmarkOther.RelativeX {
		diffs = append(diffs, accidentalmark.GongMarshallField(stage, "RelativeX"))
	}
	if accidentalmark.RelativeY != accidentalmarkOther.RelativeY {
		diffs = append(diffs, accidentalmark.GongMarshallField(stage, "RelativeY"))
	}
	if accidentalmark.FontFamily != accidentalmarkOther.FontFamily {
		diffs = append(diffs, accidentalmark.GongMarshallField(stage, "FontFamily"))
	}
	if accidentalmark.FontSize != accidentalmarkOther.FontSize {
		diffs = append(diffs, accidentalmark.GongMarshallField(stage, "FontSize"))
	}
	if accidentalmark.Color != accidentalmarkOther.Color {
		diffs = append(diffs, accidentalmark.GongMarshallField(stage, "Color"))
	}
	if accidentalmark.Id != accidentalmarkOther.Id {
		diffs = append(diffs, accidentalmark.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (arpeggiate *Arpeggiate) GongDiff(stage *Stage, arpeggiateOther *Arpeggiate) (diffs []string) {
	// insertion point for field diffs
	if arpeggiate.Name != arpeggiateOther.Name {
		diffs = append(diffs, arpeggiate.GongMarshallField(stage, "Name"))
	}
	if arpeggiate.Number != arpeggiateOther.Number {
		diffs = append(diffs, arpeggiate.GongMarshallField(stage, "Number"))
	}
	if arpeggiate.DefaultX != arpeggiateOther.DefaultX {
		diffs = append(diffs, arpeggiate.GongMarshallField(stage, "DefaultX"))
	}
	if arpeggiate.DefaultY != arpeggiateOther.DefaultY {
		diffs = append(diffs, arpeggiate.GongMarshallField(stage, "DefaultY"))
	}
	if arpeggiate.RelativeX != arpeggiateOther.RelativeX {
		diffs = append(diffs, arpeggiate.GongMarshallField(stage, "RelativeX"))
	}
	if arpeggiate.RelativeY != arpeggiateOther.RelativeY {
		diffs = append(diffs, arpeggiate.GongMarshallField(stage, "RelativeY"))
	}
	if arpeggiate.Color != arpeggiateOther.Color {
		diffs = append(diffs, arpeggiate.GongMarshallField(stage, "Color"))
	}
	if arpeggiate.Id != arpeggiateOther.Id {
		diffs = append(diffs, arpeggiate.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (arrow *Arrow) GongDiff(stage *Stage, arrowOther *Arrow) (diffs []string) {
	// insertion point for field diffs
	if arrow.Name != arrowOther.Name {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Name"))
	}
	if arrow.DefaultX != arrowOther.DefaultX {
		diffs = append(diffs, arrow.GongMarshallField(stage, "DefaultX"))
	}
	if arrow.DefaultY != arrowOther.DefaultY {
		diffs = append(diffs, arrow.GongMarshallField(stage, "DefaultY"))
	}
	if arrow.RelativeX != arrowOther.RelativeX {
		diffs = append(diffs, arrow.GongMarshallField(stage, "RelativeX"))
	}
	if arrow.RelativeY != arrowOther.RelativeY {
		diffs = append(diffs, arrow.GongMarshallField(stage, "RelativeY"))
	}
	if arrow.FontFamily != arrowOther.FontFamily {
		diffs = append(diffs, arrow.GongMarshallField(stage, "FontFamily"))
	}
	if arrow.FontSize != arrowOther.FontSize {
		diffs = append(diffs, arrow.GongMarshallField(stage, "FontSize"))
	}
	if arrow.Color != arrowOther.Color {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Color"))
	}
	if arrow.Smufl != arrowOther.Smufl {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Smufl"))
	}
	if arrow.ArrowDirection != arrowOther.ArrowDirection {
		diffs = append(diffs, arrow.GongMarshallField(stage, "ArrowDirection"))
	}
	if arrow.ArrowStyle != arrowOther.ArrowStyle {
		diffs = append(diffs, arrow.GongMarshallField(stage, "ArrowStyle"))
	}
	if arrow.CircularArrow != arrowOther.CircularArrow {
		diffs = append(diffs, arrow.GongMarshallField(stage, "CircularArrow"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (articulations *Articulations) GongDiff(stage *Stage, articulationsOther *Articulations) (diffs []string) {
	// insertion point for field diffs
	if articulations.Name != articulationsOther.Name {
		diffs = append(diffs, articulations.GongMarshallField(stage, "Name"))
	}
	if articulations.Id != articulationsOther.Id {
		diffs = append(diffs, articulations.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "StrongAccent", articulationsOther.StrongAccent, articulations.StrongAccent); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Scoop", articulationsOther.Scoop, articulations.Scoop); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Plop", articulationsOther.Plop, articulations.Plop); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Doit", articulationsOther.Doit, articulations.Doit); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Falloff", articulationsOther.Falloff, articulations.Falloff); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "BreathMark", articulationsOther.BreathMark, articulations.BreathMark); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Caesura", articulationsOther.Caesura, articulations.Caesura); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "OtherArticulation", articulationsOther.OtherArticulation, articulations.OtherArticulation); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (assess *Assess) GongDiff(stage *Stage, assessOther *Assess) (diffs []string) {
	// insertion point for field diffs
	if assess.Name != assessOther.Name {
		diffs = append(diffs, assess.GongMarshallField(stage, "Name"))
	}
	if assess.Player != assessOther.Player {
		diffs = append(diffs, assess.GongMarshallField(stage, "Player"))
	}
	if assess.TimeOnly != assessOther.TimeOnly {
		diffs = append(diffs, assess.GongMarshallField(stage, "TimeOnly"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (backup *Backup) GongDiff(stage *Stage, backupOther *Backup) (diffs []string) {
	// insertion point for field diffs
	if backup.Name != backupOther.Name {
		diffs = append(diffs, backup.GongMarshallField(stage, "Name"))
	}
	if backup.Duration != backupOther.Duration {
		diffs = append(diffs, backup.GongMarshallField(stage, "Duration"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (beam *Beam) GongDiff(stage *Stage, beamOther *Beam) (diffs []string) {
	// insertion point for field diffs
	if beam.Name != beamOther.Name {
		diffs = append(diffs, beam.GongMarshallField(stage, "Name"))
	}
	if beam.Number != beamOther.Number {
		diffs = append(diffs, beam.GongMarshallField(stage, "Number"))
	}
	if beam.Fan != beamOther.Fan {
		diffs = append(diffs, beam.GongMarshallField(stage, "Fan"))
	}
	if beam.Color != beamOther.Color {
		diffs = append(diffs, beam.GongMarshallField(stage, "Color"))
	}
	if beam.Id != beamOther.Id {
		diffs = append(diffs, beam.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bend *Bend) GongDiff(stage *Stage, bendOther *Bend) (diffs []string) {
	// insertion point for field diffs
	if bend.Name != bendOther.Name {
		diffs = append(diffs, bend.GongMarshallField(stage, "Name"))
	}
	if bend.Shape != bendOther.Shape {
		diffs = append(diffs, bend.GongMarshallField(stage, "Shape"))
	}
	if bend.DefaultX != bendOther.DefaultX {
		diffs = append(diffs, bend.GongMarshallField(stage, "DefaultX"))
	}
	if bend.DefaultY != bendOther.DefaultY {
		diffs = append(diffs, bend.GongMarshallField(stage, "DefaultY"))
	}
	if bend.RelativeX != bendOther.RelativeX {
		diffs = append(diffs, bend.GongMarshallField(stage, "RelativeX"))
	}
	if bend.RelativeY != bendOther.RelativeY {
		diffs = append(diffs, bend.GongMarshallField(stage, "RelativeY"))
	}
	if bend.FontFamily != bendOther.FontFamily {
		diffs = append(diffs, bend.GongMarshallField(stage, "FontFamily"))
	}
	if bend.FontSize != bendOther.FontSize {
		diffs = append(diffs, bend.GongMarshallField(stage, "FontSize"))
	}
	if bend.Color != bendOther.Color {
		diffs = append(diffs, bend.GongMarshallField(stage, "Color"))
	}
	if bend.Beats != bendOther.Beats {
		diffs = append(diffs, bend.GongMarshallField(stage, "Beats"))
	}
	if bend.FirstBeat != bendOther.FirstBeat {
		diffs = append(diffs, bend.GongMarshallField(stage, "FirstBeat"))
	}
	if bend.LastBeat != bendOther.LastBeat {
		diffs = append(diffs, bend.GongMarshallField(stage, "LastBeat"))
	}
	if bend.BendAlter != bendOther.BendAlter {
		diffs = append(diffs, bend.GongMarshallField(stage, "BendAlter"))
	}
	if bend.Release != bendOther.Release {
		diffs = append(diffs, bend.GongMarshallField(stage, "Release"))
	}
	if bend.WithBar != bendOther.WithBar {
		diffs = append(diffs, bend.GongMarshallField(stage, "WithBar"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (breathmark *BreathMark) GongDiff(stage *Stage, breathmarkOther *BreathMark) (diffs []string) {
	// insertion point for field diffs
	if breathmark.Name != breathmarkOther.Name {
		diffs = append(diffs, breathmark.GongMarshallField(stage, "Name"))
	}
	if breathmark.DefaultX != breathmarkOther.DefaultX {
		diffs = append(diffs, breathmark.GongMarshallField(stage, "DefaultX"))
	}
	if breathmark.DefaultY != breathmarkOther.DefaultY {
		diffs = append(diffs, breathmark.GongMarshallField(stage, "DefaultY"))
	}
	if breathmark.RelativeX != breathmarkOther.RelativeX {
		diffs = append(diffs, breathmark.GongMarshallField(stage, "RelativeX"))
	}
	if breathmark.RelativeY != breathmarkOther.RelativeY {
		diffs = append(diffs, breathmark.GongMarshallField(stage, "RelativeY"))
	}
	if breathmark.FontFamily != breathmarkOther.FontFamily {
		diffs = append(diffs, breathmark.GongMarshallField(stage, "FontFamily"))
	}
	if breathmark.FontSize != breathmarkOther.FontSize {
		diffs = append(diffs, breathmark.GongMarshallField(stage, "FontSize"))
	}
	if breathmark.Color != breathmarkOther.Color {
		diffs = append(diffs, breathmark.GongMarshallField(stage, "Color"))
	}
	if breathmark.Value != breathmarkOther.Value {
		diffs = append(diffs, breathmark.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (caesura *Caesura) GongDiff(stage *Stage, caesuraOther *Caesura) (diffs []string) {
	// insertion point for field diffs
	if caesura.Name != caesuraOther.Name {
		diffs = append(diffs, caesura.GongMarshallField(stage, "Name"))
	}
	if caesura.DefaultX != caesuraOther.DefaultX {
		diffs = append(diffs, caesura.GongMarshallField(stage, "DefaultX"))
	}
	if caesura.DefaultY != caesuraOther.DefaultY {
		diffs = append(diffs, caesura.GongMarshallField(stage, "DefaultY"))
	}
	if caesura.RelativeX != caesuraOther.RelativeX {
		diffs = append(diffs, caesura.GongMarshallField(stage, "RelativeX"))
	}
	if caesura.RelativeY != caesuraOther.RelativeY {
		diffs = append(diffs, caesura.GongMarshallField(stage, "RelativeY"))
	}
	if caesura.FontFamily != caesuraOther.FontFamily {
		diffs = append(diffs, caesura.GongMarshallField(stage, "FontFamily"))
	}
	if caesura.FontSize != caesuraOther.FontSize {
		diffs = append(diffs, caesura.GongMarshallField(stage, "FontSize"))
	}
	if caesura.Color != caesuraOther.Color {
		diffs = append(diffs, caesura.GongMarshallField(stage, "Color"))
	}
	if caesura.Value != caesuraOther.Value {
		diffs = append(diffs, caesura.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (elision *Elision) GongDiff(stage *Stage, elisionOther *Elision) (diffs []string) {
	// insertion point for field diffs
	if elision.Name != elisionOther.Name {
		diffs = append(diffs, elision.GongMarshallField(stage, "Name"))
	}
	if elision.Smufl != elisionOther.Smufl {
		diffs = append(diffs, elision.GongMarshallField(stage, "Smufl"))
	}
	if elision.FontFamily != elisionOther.FontFamily {
		diffs = append(diffs, elision.GongMarshallField(stage, "FontFamily"))
	}
	if elision.FontSize != elisionOther.FontSize {
		diffs = append(diffs, elision.GongMarshallField(stage, "FontSize"))
	}
	if elision.Color != elisionOther.Color {
		diffs = append(diffs, elision.GongMarshallField(stage, "Color"))
	}
	if elision.Value != elisionOther.Value {
		diffs = append(diffs, elision.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (emptyline *EmptyLine) GongDiff(stage *Stage, emptylineOther *EmptyLine) (diffs []string) {
	// insertion point for field diffs
	if emptyline.Name != emptylineOther.Name {
		diffs = append(diffs, emptyline.GongMarshallField(stage, "Name"))
	}
	if emptyline.DashLength != emptylineOther.DashLength {
		diffs = append(diffs, emptyline.GongMarshallField(stage, "DashLength"))
	}
	if emptyline.SpaceLength != emptylineOther.SpaceLength {
		diffs = append(diffs, emptyline.GongMarshallField(stage, "SpaceLength"))
	}
	if emptyline.DefaultX != emptylineOther.DefaultX {
		diffs = append(diffs, emptyline.GongMarshallField(stage, "DefaultX"))
	}
	if emptyline.DefaultY != emptylineOther.DefaultY {
		diffs = append(diffs, emptyline.GongMarshallField(stage, "DefaultY"))
	}
	if emptyline.RelativeX != emptylineOther.RelativeX {
		diffs = append(diffs, emptyline.GongMarshallField(stage, "RelativeX"))
	}
	if emptyline.RelativeY != emptylineOther.RelativeY {
		diffs = append(diffs, emptyline.GongMarshallField(stage, "RelativeY"))
	}
	if emptyline.FontFamily != emptylineOther.FontFamily {
		diffs = append(diffs, emptyline.GongMarshallField(stage, "FontFamily"))
	}
	if emptyline.FontSize != emptylineOther.FontSize {
		diffs = append(diffs, emptyline.GongMarshallField(stage, "FontSize"))
	}
	if emptyline.Color != emptylineOther.Color {
		diffs = append(diffs, emptyline.GongMarshallField(stage, "Color"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (extend *Extend) GongDiff(stage *Stage, extendOther *Extend) (diffs []string) {
	// insertion point for field diffs
	if extend.Name != extendOther.Name {
		diffs = append(diffs, extend.GongMarshallField(stage, "Name"))
	}
	if extend.DefaultX != extendOther.DefaultX {
		diffs = append(diffs, extend.GongMarshallField(stage, "DefaultX"))
	}
	if extend.DefaultY != extendOther.DefaultY {
		diffs = append(diffs, extend.GongMarshallField(stage, "DefaultY"))
	}
	if extend.RelativeX != extendOther.RelativeX {
		diffs = append(diffs, extend.GongMarshallField(stage, "RelativeX"))
	}
	if extend.RelativeY != extendOther.RelativeY {
		diffs = append(diffs, extend.GongMarshallField(stage, "RelativeY"))
	}
	if extend.Color != extendOther.Color {
		diffs = append(diffs, extend.GongMarshallField(stage, "Color"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (figure *Figure) GongDiff(stage *Stage, figureOther *Figure) (diffs []string) {
	// insertion point for field diffs
	if figure.Name != figureOther.Name {
		diffs = append(diffs, figure.GongMarshallField(stage, "Name"))
	}
	if figure.Extend != figureOther.Extend {
		diffs = append(diffs, figure.GongMarshallField(stage, "Extend"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (figuredbass *FiguredBass) GongDiff(stage *Stage, figuredbassOther *FiguredBass) (diffs []string) {
	// insertion point for field diffs
	if figuredbass.Name != figuredbassOther.Name {
		diffs = append(diffs, figuredbass.GongMarshallField(stage, "Name"))
	}
	if figuredbass.DefaultX != figuredbassOther.DefaultX {
		diffs = append(diffs, figuredbass.GongMarshallField(stage, "DefaultX"))
	}
	if figuredbass.DefaultY != figuredbassOther.DefaultY {
		diffs = append(diffs, figuredbass.GongMarshallField(stage, "DefaultY"))
	}
	if figuredbass.RelativeX != figuredbassOther.RelativeX {
		diffs = append(diffs, figuredbass.GongMarshallField(stage, "RelativeX"))
	}
	if figuredbass.RelativeY != figuredbassOther.RelativeY {
		diffs = append(diffs, figuredbass.GongMarshallField(stage, "RelativeY"))
	}
	if figuredbass.FontFamily != figuredbassOther.FontFamily {
		diffs = append(diffs, figuredbass.GongMarshallField(stage, "FontFamily"))
	}
	if figuredbass.FontSize != figuredbassOther.FontSize {
		diffs = append(diffs, figuredbass.GongMarshallField(stage, "FontSize"))
	}
	if figuredbass.Color != figuredbassOther.Color {
		diffs = append(diffs, figuredbass.GongMarshallField(stage, "Color"))
	}
	if figuredbass.Id != figuredbassOther.Id {
		diffs = append(diffs, figuredbass.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, figuredbass, "Figure", figuredbassOther.Figure, figuredbass.Figure); ops != "" {
		diffs = append(diffs, ops)
	}
	if figuredbass.Duration != figuredbassOther.Duration {
		diffs = append(diffs, figuredbass.GongMarshallField(stage, "Duration"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (forward *Forward) GongDiff(stage *Stage, forwardOther *Forward) (diffs []string) {
	// insertion point for field diffs
	if forward.Name != forwardOther.Name {
		diffs = append(diffs, forward.GongMarshallField(stage, "Name"))
	}
	if forward.Duration != forwardOther.Duration {
		diffs = append(diffs, forward.GongMarshallField(stage, "Duration"))
	}
	if forward.Voice != forwardOther.Voice {
		diffs = append(diffs, forward.GongMarshallField(stage, "Voice"))
	}
	if forward.Staff != forwardOther.Staff {
		diffs = append(diffs, forward.GongMarshallField(stage, "Staff"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (glissando *Glissando) GongDiff(stage *Stage, glissandoOther *Glissando) (diffs []string) {
	// insertion point for field diffs
	if glissando.Name != glissandoOther.Name {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Name"))
	}
	if glissando.Number != glissandoOther.Number {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Number"))
	}
	if glissando.DashLength != glissandoOther.DashLength {
		diffs = append(diffs, glissando.GongMarshallField(stage, "DashLength"))
	}
	if glissando.SpaceLength != glissandoOther.SpaceLength {
		diffs = append(diffs, glissando.GongMarshallField(stage, "SpaceLength"))
	}
	if glissando.DefaultX != glissandoOther.DefaultX {
		diffs = append(diffs, glissando.GongMarshallField(stage, "DefaultX"))
	}
	if glissando.DefaultY != glissandoOther.DefaultY {
		diffs = append(diffs, glissando.GongMarshallField(stage, "DefaultY"))
	}
	if glissando.RelativeX != glissandoOther.RelativeX {
		diffs = append(diffs, glissando.GongMarshallField(stage, "RelativeX"))
	}
	if glissando.RelativeY != glissandoOther.RelativeY {
		diffs = append(diffs, glissando.GongMarshallField(stage, "RelativeY"))
	}
	if glissando.FontFamily != glissandoOther.FontFamily {
		diffs = append(diffs, glissando.GongMarshallField(stage, "FontFamily"))
	}
	if glissando.FontSize != glissandoOther.FontSize {
		diffs = append(diffs, glissando.GongMarshallField(stage, "FontSize"))
	}
	if glissando.Color != glissandoOther.Color {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Color"))
	}
	if glissando.Id != glissandoOther.Id {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Id"))
	}
	if glissando.Value != glissandoOther.Value {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (grace *Grace) GongDiff(stage *Stage, graceOther *Grace) (diffs []string) {
	// insertion point for field diffs
	if grace.Name != graceOther.Name {
		diffs = append(diffs, grace.GongMarshallField(stage, "Name"))
	}
	if grace.StealTimePrevious != graceOther.StealTimePrevious {
		diffs = append(diffs, grace.GongMarshallField(stage, "StealTimePrevious"))
	}
	if grace.StealTimeFollowing != graceOther.StealTimeFollowing {
		diffs = append(diffs, grace.GongMarshallField(stage, "StealTimeFollowing"))
	}
	if grace.MakeTime != graceOther.MakeTime {
		diffs = append(diffs, grace.GongMarshallField(stage, "MakeTime"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (hammeronpulloff *HammerOnPullOff) GongDiff(stage *Stage, hammeronpulloffOther *HammerOnPullOff) (diffs []string) {
	// insertion point for field diffs
	if hammeronpulloff.Name != hammeronpulloffOther.Name {
		diffs = append(diffs, hammeronpulloff.GongMarshallField(stage, "Name"))
	}
	if hammeronpulloff.Number != hammeronpulloffOther.Number {
		diffs = append(diffs, hammeronpulloff.GongMarshallField(stage, "Number"))
	}
	if hammeronpulloff.DefaultX != hammeronpulloffOther.DefaultX {
		diffs = append(diffs, hammeronpulloff.GongMarshallField(stage, "DefaultX"))
	}
	if hammeronpulloff.DefaultY != hammeronpulloffOther.DefaultY {
		diffs = append(diffs, hammeronpulloff.GongMarshallField(stage, "DefaultY"))
	}
	if hammeronpulloff.RelativeX != hammeronpulloffOther.RelativeX {
		diffs = append(diffs, hammeronpulloff.GongMarshallField(stage, "RelativeX"))
	}
	if hammeronpulloff.RelativeY != hammeronpulloffOther.RelativeY {
		diffs = append(diffs, hammeronpulloff.GongMarshallField(stage, "RelativeY"))
	}
	if hammeronpulloff.FontFamily != hammeronpulloffOther.FontFamily {
		diffs = append(diffs, hammeronpulloff.GongMarshallField(stage, "FontFamily"))
	}
	if hammeronpulloff.FontSize != hammeronpulloffOther.FontSize {
		diffs = append(diffs, hammeronpulloff.GongMarshallField(stage, "FontSize"))
	}
	if hammeronpulloff.Color != hammeronpulloffOther.Color {
		diffs = append(diffs, hammeronpulloff.GongMarshallField(stage, "Color"))
	}
	if hammeronpulloff.Value != hammeronpulloffOther.Value {
		diffs = append(diffs, hammeronpulloff.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (handbell *Handbell) GongDiff(stage *Stage, handbellOther *Handbell) (diffs []string) {
	// insertion point for field diffs
	if handbell.Name != handbellOther.Name {
		diffs = append(diffs, handbell.GongMarshallField(stage, "Name"))
	}
	if handbell.DefaultX != handbellOther.DefaultX {
		diffs = append(diffs, handbell.GongMarshallField(stage, "DefaultX"))
	}
	if handbell.DefaultY != handbellOther.DefaultY {
		diffs = append(diffs, handbell.GongMarshallField(stage, "DefaultY"))
	}
	if handbell.RelativeX != handbellOther.RelativeX {
		diffs = append(diffs, handbell.GongMarshallField(stage, "RelativeX"))
	}
	if handbell.RelativeY != handbellOther.RelativeY {
		diffs = append(diffs, handbell.GongMarshallField(stage, "RelativeY"))
	}
	if handbell.FontFamily != handbellOther.FontFamily {
		diffs = append(diffs, handbell.GongMarshallField(stage, "FontFamily"))
	}
	if handbell.FontSize != handbellOther.FontSize {
		diffs = append(diffs, handbell.GongMarshallField(stage, "FontSize"))
	}
	if handbell.Color != handbellOther.Color {
		diffs = append(diffs, handbell.GongMarshallField(stage, "Color"))
	}
	if handbell.Value != handbellOther.Value {
		diffs = append(diffs, handbell.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (harmonclosed *HarmonClosed) GongDiff(stage *Stage, harmonclosedOther *HarmonClosed) (diffs []string) {
	// insertion point for field diffs
	if harmonclosed.Name != harmonclosedOther.Name {
		diffs = append(diffs, harmonclosed.GongMarshallField(stage, "Name"))
	}
	if harmonclosed.Location != harmonclosedOther.Location {
		diffs = append(diffs, harmonclosed.GongMarshallField(stage, "Location"))
	}
	if harmonclosed.Value != harmonclosedOther.Value {
		diffs = append(diffs, harmonclosed.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (harmonmute *HarmonMute) GongDiff(stage *Stage, harmonmuteOther *HarmonMute) (diffs []string) {
	// insertion point for field diffs
	if harmonmute.Name != harmonmuteOther.Name {
		diffs = append(diffs, harmonmute.GongMarshallField(stage, "Name"))
	}
	if harmonmute.DefaultX != harmonmuteOther.DefaultX {
		diffs = append(diffs, harmonmute.GongMarshallField(stage, "DefaultX"))
	}
	if harmonmute.DefaultY != harmonmuteOther.DefaultY {
		diffs = append(diffs, harmonmute.GongMarshallField(stage, "DefaultY"))
	}
	if harmonmute.RelativeX != harmonmuteOther.RelativeX {
		diffs = append(diffs, harmonmute.GongMarshallField(stage, "RelativeX"))
	}
	if harmonmute.RelativeY != harmonmuteOther.RelativeY {
		diffs = append(diffs, harmonmute.GongMarshallField(stage, "RelativeY"))
	}
	if harmonmute.FontFamily != harmonmuteOther.FontFamily {
		diffs = append(diffs, harmonmute.GongMarshallField(stage, "FontFamily"))
	}
	if harmonmute.FontSize != harmonmuteOther.FontSize {
		diffs = append(diffs, harmonmute.GongMarshallField(stage, "FontSize"))
	}
	if harmonmute.Color != harmonmuteOther.Color {
		diffs = append(diffs, harmonmute.GongMarshallField(stage, "Color"))
	}
	if harmonmute.HarmonClosed != harmonmuteOther.HarmonClosed {
		diffs = append(diffs, harmonmute.GongMarshallField(stage, "HarmonClosed"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (harmonic *Harmonic) GongDiff(stage *Stage, harmonicOther *Harmonic) (diffs []string) {
	// insertion point for field diffs
	if harmonic.Name != harmonicOther.Name {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "Name"))
	}
	if harmonic.DefaultX != harmonicOther.DefaultX {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "DefaultX"))
	}
	if harmonic.DefaultY != harmonicOther.DefaultY {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "DefaultY"))
	}
	if harmonic.RelativeX != harmonicOther.RelativeX {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "RelativeX"))
	}
	if harmonic.RelativeY != harmonicOther.RelativeY {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "RelativeY"))
	}
	if harmonic.FontFamily != harmonicOther.FontFamily {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "FontFamily"))
	}
	if harmonic.FontSize != harmonicOther.FontSize {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "FontSize"))
	}
	if harmonic.Color != harmonicOther.Color {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "Color"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (heeltoe *HeelToe) GongDiff(stage *Stage, heeltoeOther *HeelToe) (diffs []string) {
	// insertion point for field diffs
	if heeltoe.Name != heeltoeOther.Name {
		diffs = append(diffs, heeltoe.GongMarshallField(stage, "Name"))
	}
	if heeltoe.DefaultX != heeltoeOther.DefaultX {
		diffs = append(diffs, heeltoe.GongMarshallField(stage, "DefaultX"))
	}
	if heeltoe.DefaultY != heeltoeOther.DefaultY {
		diffs = append(diffs, heeltoe.GongMarshallField(stage, "DefaultY"))
	}
	if heeltoe.RelativeX != heeltoeOther.RelativeX {
		diffs = append(diffs, heeltoe.GongMarshallField(stage, "RelativeX"))
	}
	if heeltoe.RelativeY != heeltoeOther.RelativeY {
		diffs = append(diffs, heeltoe.GongMarshallField(stage, "RelativeY"))
	}
	if heeltoe.FontFamily != heeltoeOther.FontFamily {
		diffs = append(diffs, heeltoe.GongMarshallField(stage, "FontFamily"))
	}
	if heeltoe.FontSize != heeltoeOther.FontSize {
		diffs = append(diffs, heeltoe.GongMarshallField(stage, "FontSize"))
	}
	if heeltoe.Color != heeltoeOther.Color {
		diffs = append(diffs, heeltoe.GongMarshallField(stage, "Color"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (hole *Hole) GongDiff(stage *Stage, holeOther *Hole) (diffs []string) {
	// insertion point for field diffs
	if hole.Name != holeOther.Name {
		diffs = append(diffs, hole.GongMarshallField(stage, "Name"))
	}
	if hole.DefaultX != holeOther.DefaultX {
		diffs = append(diffs, hole.GongMarshallField(stage, "DefaultX"))
	}
	if hole.DefaultY != holeOther.DefaultY {
		diffs = append(diffs, hole.GongMarshallField(stage, "DefaultY"))
	}
	if hole.RelativeX != holeOther.RelativeX {
		diffs = append(diffs, hole.GongMarshallField(stage, "RelativeX"))
	}
	if hole.RelativeY != holeOther.RelativeY {
		diffs = append(diffs, hole.GongMarshallField(stage, "RelativeY"))
	}
	if hole.FontFamily != holeOther.FontFamily {
		diffs = append(diffs, hole.GongMarshallField(stage, "FontFamily"))
	}
	if hole.FontSize != holeOther.FontSize {
		diffs = append(diffs, hole.GongMarshallField(stage, "FontSize"))
	}
	if hole.Color != holeOther.Color {
		diffs = append(diffs, hole.GongMarshallField(stage, "Color"))
	}
	if hole.HoleType != holeOther.HoleType {
		diffs = append(diffs, hole.GongMarshallField(stage, "HoleType"))
	}
	if hole.HoleClosed != holeOther.HoleClosed {
		diffs = append(diffs, hole.GongMarshallField(stage, "HoleClosed"))
	}
	if hole.HoleShape != holeOther.HoleShape {
		diffs = append(diffs, hole.GongMarshallField(stage, "HoleShape"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (holeclosed *HoleClosed) GongDiff(stage *Stage, holeclosedOther *HoleClosed) (diffs []string) {
	// insertion point for field diffs
	if holeclosed.Name != holeclosedOther.Name {
		diffs = append(diffs, holeclosed.GongMarshallField(stage, "Name"))
	}
	if holeclosed.Location != holeclosedOther.Location {
		diffs = append(diffs, holeclosed.GongMarshallField(stage, "Location"))
	}
	if holeclosed.Value != holeclosedOther.Value {
		diffs = append(diffs, holeclosed.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (instrument *Instrument) GongDiff(stage *Stage, instrumentOther *Instrument) (diffs []string) {
	// insertion point for field diffs
	if instrument.Name != instrumentOther.Name {
		diffs = append(diffs, instrument.GongMarshallField(stage, "Name"))
	}
	if instrument.Id != instrumentOther.Id {
		diffs = append(diffs, instrument.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (listen *Listen) GongDiff(stage *Stage, listenOther *Listen) (diffs []string) {
	// insertion point for field diffs
	if listen.Name != listenOther.Name {
		diffs = append(diffs, listen.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, listen, "Assess", listenOther.Assess, listen.Assess); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, listen, "Wait", listenOther.Wait, listen.Wait); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (lyric *Lyric) GongDiff(stage *Stage, lyricOther *Lyric) (diffs []string) {
	// insertion point for field diffs
	if lyric.Name != lyricOther.Name {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Name"))
	}
	if lyric.Number != lyricOther.Number {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Number"))
	}
	if lyric.NameAttr != lyricOther.NameAttr {
		diffs = append(diffs, lyric.GongMarshallField(stage, "NameAttr"))
	}
	if lyric.TimeOnly != lyricOther.TimeOnly {
		diffs = append(diffs, lyric.GongMarshallField(stage, "TimeOnly"))
	}
	if lyric.DefaultX != lyricOther.DefaultX {
		diffs = append(diffs, lyric.GongMarshallField(stage, "DefaultX"))
	}
	if lyric.DefaultY != lyricOther.DefaultY {
		diffs = append(diffs, lyric.GongMarshallField(stage, "DefaultY"))
	}
	if lyric.RelativeX != lyricOther.RelativeX {
		diffs = append(diffs, lyric.GongMarshallField(stage, "RelativeX"))
	}
	if lyric.RelativeY != lyricOther.RelativeY {
		diffs = append(diffs, lyric.GongMarshallField(stage, "RelativeY"))
	}
	if lyric.Color != lyricOther.Color {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Color"))
	}
	if lyric.Id != lyricOther.Id {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Id"))
	}
	if lyric.Syllabic != lyricOther.Syllabic {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Syllabic"))
	}
	if lyric.Text != lyricOther.Text {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Text"))
	}
	if ops := __gong__diffSliceOfPointers(stage, lyric, "Elision", lyricOther.Elision, lyric.Elision); ops != "" {
		diffs = append(diffs, ops)
	}
	if lyric.Extend != lyricOther.Extend {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Extend"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (mordent *Mordent) GongDiff(stage *Stage, mordentOther *Mordent) (diffs []string) {
	// insertion point for field diffs
	if mordent.Name != mordentOther.Name {
		diffs = append(diffs, mordent.GongMarshallField(stage, "Name"))
	}
	if mordent.DefaultX != mordentOther.DefaultX {
		diffs = append(diffs, mordent.GongMarshallField(stage, "DefaultX"))
	}
	if mordent.DefaultY != mordentOther.DefaultY {
		diffs = append(diffs, mordent.GongMarshallField(stage, "DefaultY"))
	}
	if mordent.RelativeX != mordentOther.RelativeX {
		diffs = append(diffs, mordent.GongMarshallField(stage, "RelativeX"))
	}
	if mordent.RelativeY != mordentOther.RelativeY {
		diffs = append(diffs, mordent.GongMarshallField(stage, "RelativeY"))
	}
	if mordent.FontFamily != mordentOther.FontFamily {
		diffs = append(diffs, mordent.GongMarshallField(stage, "FontFamily"))
	}
	if mordent.FontSize != mordentOther.FontSize {
		diffs = append(diffs, mordent.GongMarshallField(stage, "FontSize"))
	}
	if mordent.Color != mordentOther.Color {
		diffs = append(diffs, mordent.GongMarshallField(stage, "Color"))
	}
	if mordent.Beats != mordentOther.Beats {
		diffs = append(diffs, mordent.GongMarshallField(stage, "Beats"))
	}
	if mordent.SecondBeat != mordentOther.SecondBeat {
		diffs = append(diffs, mordent.GongMarshallField(stage, "SecondBeat"))
	}
	if mordent.LastBeat != mordentOther.LastBeat {
		diffs = append(diffs, mordent.GongMarshallField(stage, "LastBeat"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (nonarpeggiate *NonArpeggiate) GongDiff(stage *Stage, nonarpeggiateOther *NonArpeggiate) (diffs []string) {
	// insertion point for field diffs
	if nonarpeggiate.Name != nonarpeggiateOther.Name {
		diffs = append(diffs, nonarpeggiate.GongMarshallField(stage, "Name"))
	}
	if nonarpeggiate.Number != nonarpeggiateOther.Number {
		diffs = append(diffs, nonarpeggiate.GongMarshallField(stage, "Number"))
	}
	if nonarpeggiate.DefaultX != nonarpeggiateOther.DefaultX {
		diffs = append(diffs, nonarpeggiate.GongMarshallField(stage, "DefaultX"))
	}
	if nonarpeggiate.DefaultY != nonarpeggiateOther.DefaultY {
		diffs = append(diffs, nonarpeggiate.GongMarshallField(stage, "DefaultY"))
	}
	if nonarpeggiate.RelativeX != nonarpeggiateOther.RelativeX {
		diffs = append(diffs, nonarpeggiate.GongMarshallField(stage, "RelativeX"))
	}
	if nonarpeggiate.RelativeY != nonarpeggiateOther.RelativeY {
		diffs = append(diffs, nonarpeggiate.GongMarshallField(stage, "RelativeY"))
	}
	if nonarpeggiate.Color != nonarpeggiateOther.Color {
		diffs = append(diffs, nonarpeggiate.GongMarshallField(stage, "Color"))
	}
	if nonarpeggiate.Id != nonarpeggiateOther.Id {
		diffs = append(diffs, nonarpeggiate.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notations *Notations) GongDiff(stage *Stage, notationsOther *Notations) (diffs []string) {
	// insertion point for field diffs
	if notations.Name != notationsOther.Name {
		diffs = append(diffs, notations.GongMarshallField(stage, "Name"))
	}
	if notations.Id != notationsOther.Id {
		diffs = append(diffs, notations.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Tied", notationsOther.Tied, notations.Tied); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Slur", notationsOther.Slur, notations.Slur); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Tuplet", notationsOther.Tuplet, notations.Tuplet); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Glissando", notationsOther.Glissando, notations.Glissando); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Slide", notationsOther.Slide, notations.Slide); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Ornaments", notationsOther.Ornaments, notations.Ornaments); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Technical", notationsOther.Technical, notations.Technical); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Articulations", notationsOther.Articulations, notations.Articulations); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Arpeggiate", notationsOther.Arpeggiate, notations.Arpeggiate); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "NonArpeggiate", notationsOther.NonArpeggiate, notations.NonArpeggiate); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "AccidentalMark", notationsOther.AccidentalMark, notations.AccidentalMark); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "OtherNotation", notationsOther.OtherNotation, notations.OtherNotation); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (note *Note) GongDiff(stage *Stage, noteOther *Note) (diffs []string) {
	// insertion point for field diffs
	if note.Name != noteOther.Name {
		diffs = append(diffs, note.GongMarshallField(stage, "Name"))
	}
	if note.Dynamics != noteOther.Dynamics {
		diffs = append(diffs, note.GongMarshallField(stage, "Dynamics"))
	}
	if note.EndDynamics != noteOther.EndDynamics {
		diffs = append(diffs, note.GongMarshallField(stage, "EndDynamics"))
	}
	if note.Attack != noteOther.Attack {
		diffs = append(diffs, note.GongMarshallField(stage, "Attack"))
	}
	if note.Release != noteOther.Release {
		diffs = append(diffs, note.GongMarshallField(stage, "Release"))
	}
	if note.TimeOnly != noteOther.TimeOnly {
		diffs = append(diffs, note.GongMarshallField(stage, "TimeOnly"))
	}
	if note.DefaultX != noteOther.DefaultX {
		diffs = append(diffs, note.GongMarshallField(stage, "DefaultX"))
	}
	if note.DefaultY != noteOther.DefaultY {
		diffs = append(diffs, note.GongMarshallField(stage, "DefaultY"))
	}
	if note.RelativeX != noteOther.RelativeX {
		diffs = append(diffs, note.GongMarshallField(stage, "RelativeX"))
	}
	if note.RelativeY != noteOther.RelativeY {
		diffs = append(diffs, note.GongMarshallField(stage, "RelativeY"))
	}
	if note.FontFamily != noteOther.FontFamily {
		diffs = append(diffs, note.GongMarshallField(stage, "FontFamily"))
	}
	if note.FontSize != noteOther.FontSize {
		diffs = append(diffs, note.GongMarshallField(stage, "FontSize"))
	}
	if note.Color != noteOther.Color {
		diffs = append(diffs, note.GongMarshallField(stage, "Color"))
	}
	if note.Id != noteOther.Id {
		diffs = append(diffs, note.GongMarshallField(stage, "Id"))
	}
	if note.Grace != noteOther.Grace {
		diffs = append(diffs, note.GongMarshallField(stage, "Grace"))
	}
	if note.Pitch != noteOther.Pitch {
		diffs = append(diffs, note.GongMarshallField(stage, "Pitch"))
	}
	if note.Unpitched != noteOther.Unpitched {
		diffs = append(diffs, note.GongMarshallField(stage, "Unpitched"))
	}
	if note.Rest != noteOther.Rest {
		diffs = append(diffs, note.GongMarshallField(stage, "Rest"))
	}
	if ops := __gong__diffSliceOfPointers(stage, note, "Tie", noteOther.Tie, note.Tie); ops != "" {
		diffs = append(diffs, ops)
	}
	if note.Duration != noteOther.Duration {
		diffs = append(diffs, note.GongMarshallField(stage, "Duration"))
	}
	if ops := __gong__diffSliceOfPointers(stage, note, "Instrument", noteOther.Instrument, note.Instrument); ops != "" {
		diffs = append(diffs, ops)
	}
	if note.Voice != noteOther.Voice {
		diffs = append(diffs, note.GongMarshallField(stage, "Voice"))
	}
	if note.Type != noteOther.Type {
		diffs = append(diffs, note.GongMarshallField(stage, "Type"))
	}
	if note.Accidental != noteOther.Accidental {
		diffs = append(diffs, note.GongMarshallField(stage, "Accidental"))
	}
	if note.Stem != noteOther.Stem {
		diffs = append(diffs, note.GongMarshallField(stage, "Stem"))
	}
	if note.Notehead != noteOther.Notehead {
		diffs = append(diffs, note.GongMarshallField(stage, "Notehead"))
	}
	if note.NoteheadText != noteOther.NoteheadText {
		diffs = append(diffs, note.GongMarshallField(stage, "NoteheadText"))
	}
	if note.Staff != noteOther.Staff {
		diffs = append(diffs, note.GongMarshallField(stage, "Staff"))
	}
	if ops := __gong__diffSliceOfPointers(stage, note, "Beam", noteOther.Beam, note.Beam); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, note, "Notations", noteOther.Notations, note.Notations); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, note, "Lyric", noteOther.Lyric, note.Lyric); ops != "" {
		diffs = append(diffs, ops)
	}
	if note.Listen != noteOther.Listen {
		diffs = append(diffs, note.GongMarshallField(stage, "Listen"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notetype *NoteType) GongDiff(stage *Stage, notetypeOther *NoteType) (diffs []string) {
	// insertion point for field diffs
	if notetype.Name != notetypeOther.Name {
		diffs = append(diffs, notetype.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notehead *Notehead) GongDiff(stage *Stage, noteheadOther *Notehead) (diffs []string) {
	// insertion point for field diffs
	if notehead.Name != noteheadOther.Name {
		diffs = append(diffs, notehead.GongMarshallField(stage, "Name"))
	}
	if notehead.FontFamily != noteheadOther.FontFamily {
		diffs = append(diffs, notehead.GongMarshallField(stage, "FontFamily"))
	}
	if notehead.FontSize != noteheadOther.FontSize {
		diffs = append(diffs, notehead.GongMarshallField(stage, "FontSize"))
	}
	if notehead.Color != noteheadOther.Color {
		diffs = append(diffs, notehead.GongMarshallField(stage, "Color"))
	}
	if notehead.Smufl != noteheadOther.Smufl {
		diffs = append(diffs, notehead.GongMarshallField(stage, "Smufl"))
	}
	if notehead.Value != noteheadOther.Value {
		diffs = append(diffs, notehead.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (noteheadtext *NoteheadText) GongDiff(stage *Stage, noteheadtextOther *NoteheadText) (diffs []string) {
	// insertion point for field diffs
	if noteheadtext.Name != noteheadtextOther.Name {
		diffs = append(diffs, noteheadtext.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ornaments *Ornaments) GongDiff(stage *Stage, ornamentsOther *Ornaments) (diffs []string) {
	// insertion point for field diffs
	if ornaments.Name != ornamentsOther.Name {
		diffs = append(diffs, ornaments.GongMarshallField(stage, "Name"))
	}
	if ornaments.Id != ornamentsOther.Id {
		diffs = append(diffs, ornaments.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "Mordent", ornamentsOther.Mordent, ornaments.Mordent); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "InvertedMordent", ornamentsOther.InvertedMordent, ornaments.InvertedMordent); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "Tremolo", ornamentsOther.Tremolo, ornaments.Tremolo); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "OtherOrnament", ornamentsOther.OtherOrnament, ornaments.OtherOrnament); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "AccidentalMark", ornamentsOther.AccidentalMark, ornaments.AccidentalMark); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (othernotation *OtherNotation) GongDiff(stage *Stage, othernotationOther *OtherNotation) (diffs []string) {
	// insertion point for field diffs
	if othernotation.Name != othernotationOther.Name {
		diffs = append(diffs, othernotation.GongMarshallField(stage, "Name"))
	}
	if othernotation.Number != othernotationOther.Number {
		diffs = append(diffs, othernotation.GongMarshallField(stage, "Number"))
	}
	if othernotation.DefaultX != othernotationOther.DefaultX {
		diffs = append(diffs, othernotation.GongMarshallField(stage, "DefaultX"))
	}
	if othernotation.DefaultY != othernotationOther.DefaultY {
		diffs = append(diffs, othernotation.GongMarshallField(stage, "DefaultY"))
	}
	if othernotation.RelativeX != othernotationOther.RelativeX {
		diffs = append(diffs, othernotation.GongMarshallField(stage, "RelativeX"))
	}
	if othernotation.RelativeY != othernotationOther.RelativeY {
		diffs = append(diffs, othernotation.GongMarshallField(stage, "RelativeY"))
	}
	if othernotation.FontFamily != othernotationOther.FontFamily {
		diffs = append(diffs, othernotation.GongMarshallField(stage, "FontFamily"))
	}
	if othernotation.FontSize != othernotationOther.FontSize {
		diffs = append(diffs, othernotation.GongMarshallField(stage, "FontSize"))
	}
	if othernotation.Color != othernotationOther.Color {
		diffs = append(diffs, othernotation.GongMarshallField(stage, "Color"))
	}
	if othernotation.Smufl != othernotationOther.Smufl {
		diffs = append(diffs, othernotation.GongMarshallField(stage, "Smufl"))
	}
	if othernotation.Id != othernotationOther.Id {
		diffs = append(diffs, othernotation.GongMarshallField(stage, "Id"))
	}
	if othernotation.Value != othernotationOther.Value {
		diffs = append(diffs, othernotation.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (otherplacementtext *OtherPlacementText) GongDiff(stage *Stage, otherplacementtextOther *OtherPlacementText) (diffs []string) {
	// insertion point for field diffs
	if otherplacementtext.Name != otherplacementtextOther.Name {
		diffs = append(diffs, otherplacementtext.GongMarshallField(stage, "Name"))
	}
	if otherplacementtext.DefaultX != otherplacementtextOther.DefaultX {
		diffs = append(diffs, otherplacementtext.GongMarshallField(stage, "DefaultX"))
	}
	if otherplacementtext.DefaultY != otherplacementtextOther.DefaultY {
		diffs = append(diffs, otherplacementtext.GongMarshallField(stage, "DefaultY"))
	}
	if otherplacementtext.RelativeX != otherplacementtextOther.RelativeX {
		diffs = append(diffs, otherplacementtext.GongMarshallField(stage, "RelativeX"))
	}
	if otherplacementtext.RelativeY != otherplacementtextOther.RelativeY {
		diffs = append(diffs, otherplacementtext.GongMarshallField(stage, "RelativeY"))
	}
	if otherplacementtext.FontFamily != otherplacementtextOther.FontFamily {
		diffs = append(diffs, otherplacementtext.GongMarshallField(stage, "FontFamily"))
	}
	if otherplacementtext.FontSize != otherplacementtextOther.FontSize {
		diffs = append(diffs, otherplacementtext.GongMarshallField(stage, "FontSize"))
	}
	if otherplacementtext.Color != otherplacementtextOther.Color {
		diffs = append(diffs, otherplacementtext.GongMarshallField(stage, "Color"))
	}
	if otherplacementtext.Smufl != otherplacementtextOther.Smufl {
		diffs = append(diffs, otherplacementtext.GongMarshallField(stage, "Smufl"))
	}
	if otherplacementtext.Value != otherplacementtextOther.Value {
		diffs = append(diffs, otherplacementtext.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pitch *Pitch) GongDiff(stage *Stage, pitchOther *Pitch) (diffs []string) {
	// insertion point for field diffs
	if pitch.Name != pitchOther.Name {
		diffs = append(diffs, pitch.GongMarshallField(stage, "Name"))
	}
	if pitch.Alter != pitchOther.Alter {
		diffs = append(diffs, pitch.GongMarshallField(stage, "Alter"))
	}
	if pitch.Octave != pitchOther.Octave {
		diffs = append(diffs, pitch.GongMarshallField(stage, "Octave"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (placementtext *PlacementText) GongDiff(stage *Stage, placementtextOther *PlacementText) (diffs []string) {
	// insertion point for field diffs
	if placementtext.Name != placementtextOther.Name {
		diffs = append(diffs, placementtext.GongMarshallField(stage, "Name"))
	}
	if placementtext.DefaultX != placementtextOther.DefaultX {
		diffs = append(diffs, placementtext.GongMarshallField(stage, "DefaultX"))
	}
	if placementtext.DefaultY != placementtextOther.DefaultY {
		diffs = append(diffs, placementtext.GongMarshallField(stage, "DefaultY"))
	}
	if placementtext.RelativeX != placementtextOther.RelativeX {
		diffs = append(diffs, placementtext.GongMarshallField(stage, "RelativeX"))
	}
	if placementtext.RelativeY != placementtextOther.RelativeY {
		diffs = append(diffs, placementtext.GongMarshallField(stage, "RelativeY"))
	}
	if placementtext.FontFamily != placementtextOther.FontFamily {
		diffs = append(diffs, placementtext.GongMarshallField(stage, "FontFamily"))
	}
	if placementtext.FontSize != placementtextOther.FontSize {
		diffs = append(diffs, placementtext.GongMarshallField(stage, "FontSize"))
	}
	if placementtext.Color != placementtextOther.Color {
		diffs = append(diffs, placementtext.GongMarshallField(stage, "Color"))
	}
	if placementtext.Value != placementtextOther.Value {
		diffs = append(diffs, placementtext.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (release *Release) GongDiff(stage *Stage, releaseOther *Release) (diffs []string) {
	// insertion point for field diffs
	if release.Name != releaseOther.Name {
		diffs = append(diffs, release.GongMarshallField(stage, "Name"))
	}
	if release.Offset != releaseOther.Offset {
		diffs = append(diffs, release.GongMarshallField(stage, "Offset"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (rest *Rest) GongDiff(stage *Stage, restOther *Rest) (diffs []string) {
	// insertion point for field diffs
	if rest.Name != restOther.Name {
		diffs = append(diffs, rest.GongMarshallField(stage, "Name"))
	}
	if rest.DisplayOctave != restOther.DisplayOctave {
		diffs = append(diffs, rest.GongMarshallField(stage, "DisplayOctave"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (slide *Slide) GongDiff(stage *Stage, slideOther *Slide) (diffs []string) {
	// insertion point for field diffs
	if slide.Name != slideOther.Name {
		diffs = append(diffs, slide.GongMarshallField(stage, "Name"))
	}
	if slide.Number != slideOther.Number {
		diffs = append(diffs, slide.GongMarshallField(stage, "Number"))
	}
	if slide.DashLength != slideOther.DashLength {
		diffs = append(diffs, slide.GongMarshallField(stage, "DashLength"))
	}
	if slide.SpaceLength != slideOther.SpaceLength {
		diffs = append(diffs, slide.GongMarshallField(stage, "SpaceLength"))
	}
	if slide.DefaultX != slideOther.DefaultX {
		diffs = append(diffs, slide.GongMarshallField(stage, "DefaultX"))
	}
	if slide.DefaultY != slideOther.DefaultY {
		diffs = append(diffs, slide.GongMarshallField(stage, "DefaultY"))
	}
	if slide.RelativeX != slideOther.RelativeX {
		diffs = append(diffs, slide.GongMarshallField(stage, "RelativeX"))
	}
	if slide.RelativeY != slideOther.RelativeY {
		diffs = append(diffs, slide.GongMarshallField(stage, "RelativeY"))
	}
	if slide.FontFamily != slideOther.FontFamily {
		diffs = append(diffs, slide.GongMarshallField(stage, "FontFamily"))
	}
	if slide.FontSize != slideOther.FontSize {
		diffs = append(diffs, slide.GongMarshallField(stage, "FontSize"))
	}
	if slide.Color != slideOther.Color {
		diffs = append(diffs, slide.GongMarshallField(stage, "Color"))
	}
	if slide.Beats != slideOther.Beats {
		diffs = append(diffs, slide.GongMarshallField(stage, "Beats"))
	}
	if slide.FirstBeat != slideOther.FirstBeat {
		diffs = append(diffs, slide.GongMarshallField(stage, "FirstBeat"))
	}
	if slide.LastBeat != slideOther.LastBeat {
		diffs = append(diffs, slide.GongMarshallField(stage, "LastBeat"))
	}
	if slide.Id != slideOther.Id {
		diffs = append(diffs, slide.GongMarshallField(stage, "Id"))
	}
	if slide.Value != slideOther.Value {
		diffs = append(diffs, slide.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (slur *Slur) GongDiff(stage *Stage, slurOther *Slur) (diffs []string) {
	// insertion point for field diffs
	if slur.Name != slurOther.Name {
		diffs = append(diffs, slur.GongMarshallField(stage, "Name"))
	}
	if slur.Number != slurOther.Number {
		diffs = append(diffs, slur.GongMarshallField(stage, "Number"))
	}
	if slur.DashLength != slurOther.DashLength {
		diffs = append(diffs, slur.GongMarshallField(stage, "DashLength"))
	}
	if slur.SpaceLength != slurOther.SpaceLength {
		diffs = append(diffs, slur.GongMarshallField(stage, "SpaceLength"))
	}
	if slur.DefaultX != slurOther.DefaultX {
		diffs = append(diffs, slur.GongMarshallField(stage, "DefaultX"))
	}
	if slur.DefaultY != slurOther.DefaultY {
		diffs = append(diffs, slur.GongMarshallField(stage, "DefaultY"))
	}
	if slur.RelativeX != slurOther.RelativeX {
		diffs = append(diffs, slur.GongMarshallField(stage, "RelativeX"))
	}
	if slur.RelativeY != slurOther.RelativeY {
		diffs = append(diffs, slur.GongMarshallField(stage, "RelativeY"))
	}
	if slur.BezierX != slurOther.BezierX {
		diffs = append(diffs, slur.GongMarshallField(stage, "BezierX"))
	}
	if slur.BezierY != slurOther.BezierY {
		diffs = append(diffs, slur.GongMarshallField(stage, "BezierY"))
	}
	if slur.BezierX2 != slurOther.BezierX2 {
		diffs = append(diffs, slur.GongMarshallField(stage, "BezierX2"))
	}
	if slur.BezierY2 != slurOther.BezierY2 {
		diffs = append(diffs, slur.GongMarshallField(stage, "BezierY2"))
	}
	if slur.BezierOffset != slurOther.BezierOffset {
		diffs = append(diffs, slur.GongMarshallField(stage, "BezierOffset"))
	}
	if slur.BezierOffset2 != slurOther.BezierOffset2 {
		diffs = append(diffs, slur.GongMarshallField(stage, "BezierOffset2"))
	}
	if slur.Color != slurOther.Color {
		diffs = append(diffs, slur.GongMarshallField(stage, "Color"))
	}
	if slur.Id != slurOther.Id {
		diffs = append(diffs, slur.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stem *Stem) GongDiff(stage *Stage, stemOther *Stem) (diffs []string) {
	// insertion point for field diffs
	if stem.Name != stemOther.Name {
		diffs = append(diffs, stem.GongMarshallField(stage, "Name"))
	}
	if stem.DefaultX != stemOther.DefaultX {
		diffs = append(diffs, stem.GongMarshallField(stage, "DefaultX"))
	}
	if stem.DefaultY != stemOther.DefaultY {
		diffs = append(diffs, stem.GongMarshallField(stage, "DefaultY"))
	}
	if stem.RelativeX != stemOther.RelativeX {
		diffs = append(diffs, stem.GongMarshallField(stage, "RelativeX"))
	}
	if stem.RelativeY != stemOther.RelativeY {
		diffs = append(diffs, stem.GongMarshallField(stage, "RelativeY"))
	}
	if stem.Color != stemOther.Color {
		diffs = append(diffs, stem.GongMarshallField(stage, "Color"))
	}
	if stem.Value != stemOther.Value {
		diffs = append(diffs, stem.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (strongaccent *StrongAccent) GongDiff(stage *Stage, strongaccentOther *StrongAccent) (diffs []string) {
	// insertion point for field diffs
	if strongaccent.Name != strongaccentOther.Name {
		diffs = append(diffs, strongaccent.GongMarshallField(stage, "Name"))
	}
	if strongaccent.DefaultX != strongaccentOther.DefaultX {
		diffs = append(diffs, strongaccent.GongMarshallField(stage, "DefaultX"))
	}
	if strongaccent.DefaultY != strongaccentOther.DefaultY {
		diffs = append(diffs, strongaccent.GongMarshallField(stage, "DefaultY"))
	}
	if strongaccent.RelativeX != strongaccentOther.RelativeX {
		diffs = append(diffs, strongaccent.GongMarshallField(stage, "RelativeX"))
	}
	if strongaccent.RelativeY != strongaccentOther.RelativeY {
		diffs = append(diffs, strongaccent.GongMarshallField(stage, "RelativeY"))
	}
	if strongaccent.FontFamily != strongaccentOther.FontFamily {
		diffs = append(diffs, strongaccent.GongMarshallField(stage, "FontFamily"))
	}
	if strongaccent.FontSize != strongaccentOther.FontSize {
		diffs = append(diffs, strongaccent.GongMarshallField(stage, "FontSize"))
	}
	if strongaccent.Color != strongaccentOther.Color {
		diffs = append(diffs, strongaccent.GongMarshallField(stage, "Color"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tap *Tap) GongDiff(stage *Stage, tapOther *Tap) (diffs []string) {
	// insertion point for field diffs
	if tap.Name != tapOther.Name {
		diffs = append(diffs, tap.GongMarshallField(stage, "Name"))
	}
	if tap.Hand != tapOther.Hand {
		diffs = append(diffs, tap.GongMarshallField(stage, "Hand"))
	}
	if tap.DefaultX != tapOther.DefaultX {
		diffs = append(diffs, tap.GongMarshallField(stage, "DefaultX"))
	}
	if tap.DefaultY != tapOther.DefaultY {
		diffs = append(diffs, tap.GongMarshallField(stage, "DefaultY"))
	}
	if tap.RelativeX != tapOther.RelativeX {
		diffs = append(diffs, tap.GongMarshallField(stage, "RelativeX"))
	}
	if tap.RelativeY != tapOther.RelativeY {
		diffs = append(diffs, tap.GongMarshallField(stage, "RelativeY"))
	}
	if tap.FontFamily != tapOther.FontFamily {
		diffs = append(diffs, tap.GongMarshallField(stage, "FontFamily"))
	}
	if tap.FontSize != tapOther.FontSize {
		diffs = append(diffs, tap.GongMarshallField(stage, "FontSize"))
	}
	if tap.Color != tapOther.Color {
		diffs = append(diffs, tap.GongMarshallField(stage, "Color"))
	}
	if tap.Value != tapOther.Value {
		diffs = append(diffs, tap.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (technical *Technical) GongDiff(stage *Stage, technicalOther *Technical) (diffs []string) {
	// insertion point for field diffs
	if technical.Name != technicalOther.Name {
		diffs = append(diffs, technical.GongMarshallField(stage, "Name"))
	}
	if technical.Id != technicalOther.Id {
		diffs = append(diffs, technical.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Harmonic", technicalOther.Harmonic, technical.Harmonic); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Pluck", technicalOther.Pluck, technical.Pluck); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "HammerOn", technicalOther.HammerOn, technical.HammerOn); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "PullOff", technicalOther.PullOff, technical.PullOff); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Bend", technicalOther.Bend, technical.Bend); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Tap", technicalOther.Tap, technical.Tap); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Heel", technicalOther.Heel, technical.Heel); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Toe", technicalOther.Toe, technical.Toe); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Hole", technicalOther.Hole, technical.Hole); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Arrow", technicalOther.Arrow, technical.Arrow); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Handbell", technicalOther.Handbell, technical.Handbell); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "HarmonMute", technicalOther.HarmonMute, technical.HarmonMute); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "OtherTechnical", technicalOther.OtherTechnical, technical.OtherTechnical); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (textelementdata *TextElementData) GongDiff(stage *Stage, textelementdataOther *TextElementData) (diffs []string) {
	// insertion point for field diffs
	if textelementdata.Name != textelementdataOther.Name {
		diffs = append(diffs, textelementdata.GongMarshallField(stage, "Name"))
	}
	if textelementdata.Lang != textelementdataOther.Lang {
		diffs = append(diffs, textelementdata.GongMarshallField(stage, "Lang"))
	}
	if textelementdata.FontFamily != textelementdataOther.FontFamily {
		diffs = append(diffs, textelementdata.GongMarshallField(stage, "FontFamily"))
	}
	if textelementdata.FontSize != textelementdataOther.FontSize {
		diffs = append(diffs, textelementdata.GongMarshallField(stage, "FontSize"))
	}
	if textelementdata.Color != textelementdataOther.Color {
		diffs = append(diffs, textelementdata.GongMarshallField(stage, "Color"))
	}
	if textelementdata.Underline != textelementdataOther.Underline {
		diffs = append(diffs, textelementdata.GongMarshallField(stage, "Underline"))
	}
	if textelementdata.Overline != textelementdataOther.Overline {
		diffs = append(diffs, textelementdata.GongMarshallField(stage, "Overline"))
	}
	if textelementdata.LineThrough != textelementdataOther.LineThrough {
		diffs = append(diffs, textelementdata.GongMarshallField(stage, "LineThrough"))
	}
	if textelementdata.Rotation != textelementdataOther.Rotation {
		diffs = append(diffs, textelementdata.GongMarshallField(stage, "Rotation"))
	}
	if textelementdata.LetterSpacing != textelementdataOther.LetterSpacing {
		diffs = append(diffs, textelementdata.GongMarshallField(stage, "LetterSpacing"))
	}
	if textelementdata.Value != textelementdataOther.Value {
		diffs = append(diffs, textelementdata.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tie *Tie) GongDiff(stage *Stage, tieOther *Tie) (diffs []string) {
	// insertion point for field diffs
	if tie.Name != tieOther.Name {
		diffs = append(diffs, tie.GongMarshallField(stage, "Name"))
	}
	if tie.TimeOnly != tieOther.TimeOnly {
		diffs = append(diffs, tie.GongMarshallField(stage, "TimeOnly"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tied *Tied) GongDiff(stage *Stage, tiedOther *Tied) (diffs []string) {
	// insertion point for field diffs
	if tied.Name != tiedOther.Name {
		diffs = append(diffs, tied.GongMarshallField(stage, "Name"))
	}
	if tied.Number != tiedOther.Number {
		diffs = append(diffs, tied.GongMarshallField(stage, "Number"))
	}
	if tied.DashLength != tiedOther.DashLength {
		diffs = append(diffs, tied.GongMarshallField(stage, "DashLength"))
	}
	if tied.SpaceLength != tiedOther.SpaceLength {
		diffs = append(diffs, tied.GongMarshallField(stage, "SpaceLength"))
	}
	if tied.DefaultX != tiedOther.DefaultX {
		diffs = append(diffs, tied.GongMarshallField(stage, "DefaultX"))
	}
	if tied.DefaultY != tiedOther.DefaultY {
		diffs = append(diffs, tied.GongMarshallField(stage, "DefaultY"))
	}
	if tied.RelativeX != tiedOther.RelativeX {
		diffs = append(diffs, tied.GongMarshallField(stage, "RelativeX"))
	}
	if tied.RelativeY != tiedOther.RelativeY {
		diffs = append(diffs, tied.GongMarshallField(stage, "RelativeY"))
	}
	if tied.BezierX != tiedOther.BezierX {
		diffs = append(diffs, tied.GongMarshallField(stage, "BezierX"))
	}
	if tied.BezierY != tiedOther.BezierY {
		diffs = append(diffs, tied.GongMarshallField(stage, "BezierY"))
	}
	if tied.BezierX2 != tiedOther.BezierX2 {
		diffs = append(diffs, tied.GongMarshallField(stage, "BezierX2"))
	}
	if tied.BezierY2 != tiedOther.BezierY2 {
		diffs = append(diffs, tied.GongMarshallField(stage, "BezierY2"))
	}
	if tied.BezierOffset != tiedOther.BezierOffset {
		diffs = append(diffs, tied.GongMarshallField(stage, "BezierOffset"))
	}
	if tied.BezierOffset2 != tiedOther.BezierOffset2 {
		diffs = append(diffs, tied.GongMarshallField(stage, "BezierOffset2"))
	}
	if tied.Color != tiedOther.Color {
		diffs = append(diffs, tied.GongMarshallField(stage, "Color"))
	}
	if tied.Id != tiedOther.Id {
		diffs = append(diffs, tied.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tremolo *Tremolo) GongDiff(stage *Stage, tremoloOther *Tremolo) (diffs []string) {
	// insertion point for field diffs
	if tremolo.Name != tremoloOther.Name {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "Name"))
	}
	if tremolo.DefaultX != tremoloOther.DefaultX {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "DefaultX"))
	}
	if tremolo.DefaultY != tremoloOther.DefaultY {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "DefaultY"))
	}
	if tremolo.RelativeX != tremoloOther.RelativeX {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "RelativeX"))
	}
	if tremolo.RelativeY != tremoloOther.RelativeY {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "RelativeY"))
	}
	if tremolo.FontFamily != tremoloOther.FontFamily {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "FontFamily"))
	}
	if tremolo.FontSize != tremoloOther.FontSize {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "FontSize"))
	}
	if tremolo.Color != tremoloOther.Color {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "Color"))
	}
	if tremolo.Smufl != tremoloOther.Smufl {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "Smufl"))
	}
	if tremolo.Value != tremoloOther.Value {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tuplet *Tuplet) GongDiff(stage *Stage, tupletOther *Tuplet) (diffs []string) {
	// insertion point for field diffs
	if tuplet.Name != tupletOther.Name {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "Name"))
	}
	if tuplet.Number != tupletOther.Number {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "Number"))
	}
	if tuplet.DefaultX != tupletOther.DefaultX {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "DefaultX"))
	}
	if tuplet.DefaultY != tupletOther.DefaultY {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "DefaultY"))
	}
	if tuplet.RelativeX != tupletOther.RelativeX {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "RelativeX"))
	}
	if tuplet.RelativeY != tupletOther.RelativeY {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "RelativeY"))
	}
	if tuplet.Id != tupletOther.Id {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "Id"))
	}
	if tuplet.TupletActual != tupletOther.TupletActual {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "TupletActual"))
	}
	if tuplet.TupletNormal != tupletOther.TupletNormal {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "TupletNormal"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tupletdot *TupletDot) GongDiff(stage *Stage, tupletdotOther *TupletDot) (diffs []string) {
	// insertion point for field diffs
	if tupletdot.Name != tupletdotOther.Name {
		diffs = append(diffs, tupletdot.GongMarshallField(stage, "Name"))
	}
	if tupletdot.FontFamily != tupletdotOther.FontFamily {
		diffs = append(diffs, tupletdot.GongMarshallField(stage, "FontFamily"))
	}
	if tupletdot.FontSize != tupletdotOther.FontSize {
		diffs = append(diffs, tupletdot.GongMarshallField(stage, "FontSize"))
	}
	if tupletdot.Color != tupletdotOther.Color {
		diffs = append(diffs, tupletdot.GongMarshallField(stage, "Color"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tupletnumber *TupletNumber) GongDiff(stage *Stage, tupletnumberOther *TupletNumber) (diffs []string) {
	// insertion point for field diffs
	if tupletnumber.Name != tupletnumberOther.Name {
		diffs = append(diffs, tupletnumber.GongMarshallField(stage, "Name"))
	}
	if tupletnumber.FontFamily != tupletnumberOther.FontFamily {
		diffs = append(diffs, tupletnumber.GongMarshallField(stage, "FontFamily"))
	}
	if tupletnumber.FontSize != tupletnumberOther.FontSize {
		diffs = append(diffs, tupletnumber.GongMarshallField(stage, "FontSize"))
	}
	if tupletnumber.Color != tupletnumberOther.Color {
		diffs = append(diffs, tupletnumber.GongMarshallField(stage, "Color"))
	}
	if tupletnumber.Value != tupletnumberOther.Value {
		diffs = append(diffs, tupletnumber.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tupletportion *TupletPortion) GongDiff(stage *Stage, tupletportionOther *TupletPortion) (diffs []string) {
	// insertion point for field diffs
	if tupletportion.Name != tupletportionOther.Name {
		diffs = append(diffs, tupletportion.GongMarshallField(stage, "Name"))
	}
	if tupletportion.TupletNumber != tupletportionOther.TupletNumber {
		diffs = append(diffs, tupletportion.GongMarshallField(stage, "TupletNumber"))
	}
	if tupletportion.TupletType != tupletportionOther.TupletType {
		diffs = append(diffs, tupletportion.GongMarshallField(stage, "TupletType"))
	}
	if ops := __gong__diffSliceOfPointers(stage, tupletportion, "TupletDot", tupletportionOther.TupletDot, tupletportion.TupletDot); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tuplettype *TupletType) GongDiff(stage *Stage, tuplettypeOther *TupletType) (diffs []string) {
	// insertion point for field diffs
	if tuplettype.Name != tuplettypeOther.Name {
		diffs = append(diffs, tuplettype.GongMarshallField(stage, "Name"))
	}
	if tuplettype.FontFamily != tuplettypeOther.FontFamily {
		diffs = append(diffs, tuplettype.GongMarshallField(stage, "FontFamily"))
	}
	if tuplettype.FontSize != tuplettypeOther.FontSize {
		diffs = append(diffs, tuplettype.GongMarshallField(stage, "FontSize"))
	}
	if tuplettype.Color != tuplettypeOther.Color {
		diffs = append(diffs, tuplettype.GongMarshallField(stage, "Color"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (unpitched *Unpitched) GongDiff(stage *Stage, unpitchedOther *Unpitched) (diffs []string) {
	// insertion point for field diffs
	if unpitched.Name != unpitchedOther.Name {
		diffs = append(diffs, unpitched.GongMarshallField(stage, "Name"))
	}
	if unpitched.DisplayOctave != unpitchedOther.DisplayOctave {
		diffs = append(diffs, unpitched.GongMarshallField(stage, "DisplayOctave"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (wait *Wait) GongDiff(stage *Stage, waitOther *Wait) (diffs []string) {
	// insertion point for field diffs
	if wait.Name != waitOther.Name {
		diffs = append(diffs, wait.GongMarshallField(stage, "Name"))
	}
	if wait.Player != waitOther.Player {
		diffs = append(diffs, wait.GongMarshallField(stage, "Player"))
	}
	if wait.TimeOnly != waitOther.TimeOnly {
		diffs = append(diffs, wait.GongMarshallField(stage, "TimeOnly"))
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
