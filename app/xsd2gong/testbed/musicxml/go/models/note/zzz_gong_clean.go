// generated code - do not edit
package note

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T GongstructPtr](slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if stage.IsStaged(element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T GongstructPtr](element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !stage.IsStaged(*element) {
		*element = zero
		modified = true
		return
	}
	return
}

type GongCleaner interface {
	GongClean(stage *Stage) (modified bool)
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Articulations
func (articulations *Articulations) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&articulations.StrongAccent) || modified
	modified = stage.CleanSlice(&articulations.Scoop) || modified
	modified = stage.CleanSlice(&articulations.Plop) || modified
	modified = stage.CleanSlice(&articulations.Doit) || modified
	modified = stage.CleanSlice(&articulations.Falloff) || modified
	modified = stage.CleanSlice(&articulations.BreathMark) || modified
	modified = stage.CleanSlice(&articulations.Caesura) || modified
	modified = stage.CleanSlice(&articulations.OtherArticulation) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Bend
func (bend *Bend) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&bend.Release) || modified
	modified = stage.CleanPointer(&bend.WithBar) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Figure
func (figure *Figure) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&figure.Extend) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FiguredBass
func (figuredbass *FiguredBass) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&figuredbass.Figure) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by HarmonMute
func (harmonmute *HarmonMute) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&harmonmute.HarmonClosed) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Hole
func (hole *Hole) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&hole.HoleClosed) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Listen
func (listen *Listen) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&listen.Assess) || modified
	modified = stage.CleanSlice(&listen.Wait) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Lyric
func (lyric *Lyric) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&lyric.Elision) || modified
	// insertion point per field
	modified = stage.CleanPointer(&lyric.Text) || modified
	modified = stage.CleanPointer(&lyric.Extend) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Notations
func (notations *Notations) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&notations.Tied) || modified
	modified = stage.CleanSlice(&notations.Slur) || modified
	modified = stage.CleanSlice(&notations.Tuplet) || modified
	modified = stage.CleanSlice(&notations.Glissando) || modified
	modified = stage.CleanSlice(&notations.Slide) || modified
	modified = stage.CleanSlice(&notations.Ornaments) || modified
	modified = stage.CleanSlice(&notations.Technical) || modified
	modified = stage.CleanSlice(&notations.Articulations) || modified
	modified = stage.CleanSlice(&notations.Arpeggiate) || modified
	modified = stage.CleanSlice(&notations.NonArpeggiate) || modified
	modified = stage.CleanSlice(&notations.AccidentalMark) || modified
	modified = stage.CleanSlice(&notations.OtherNotation) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Note
func (note *Note) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&note.Tie) || modified
	modified = stage.CleanSlice(&note.Instrument) || modified
	modified = stage.CleanSlice(&note.Beam) || modified
	modified = stage.CleanSlice(&note.Notations) || modified
	modified = stage.CleanSlice(&note.Lyric) || modified
	// insertion point per field
	modified = stage.CleanPointer(&note.Grace) || modified
	modified = stage.CleanPointer(&note.Pitch) || modified
	modified = stage.CleanPointer(&note.Unpitched) || modified
	modified = stage.CleanPointer(&note.Rest) || modified
	modified = stage.CleanPointer(&note.Type) || modified
	modified = stage.CleanPointer(&note.Accidental) || modified
	modified = stage.CleanPointer(&note.Stem) || modified
	modified = stage.CleanPointer(&note.Notehead) || modified
	modified = stage.CleanPointer(&note.NoteheadText) || modified
	modified = stage.CleanPointer(&note.Listen) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Ornaments
func (ornaments *Ornaments) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ornaments.Mordent) || modified
	modified = stage.CleanSlice(&ornaments.InvertedMordent) || modified
	modified = stage.CleanSlice(&ornaments.Tremolo) || modified
	modified = stage.CleanSlice(&ornaments.OtherOrnament) || modified
	modified = stage.CleanSlice(&ornaments.AccidentalMark) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Technical
func (technical *Technical) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&technical.Harmonic) || modified
	modified = stage.CleanSlice(&technical.Pluck) || modified
	modified = stage.CleanSlice(&technical.HammerOn) || modified
	modified = stage.CleanSlice(&technical.PullOff) || modified
	modified = stage.CleanSlice(&technical.Bend) || modified
	modified = stage.CleanSlice(&technical.Tap) || modified
	modified = stage.CleanSlice(&technical.Heel) || modified
	modified = stage.CleanSlice(&technical.Toe) || modified
	modified = stage.CleanSlice(&technical.Hole) || modified
	modified = stage.CleanSlice(&technical.Arrow) || modified
	modified = stage.CleanSlice(&technical.Handbell) || modified
	modified = stage.CleanSlice(&technical.HarmonMute) || modified
	modified = stage.CleanSlice(&technical.OtherTechnical) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Tuplet
func (tuplet *Tuplet) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&tuplet.TupletActual) || modified
	modified = stage.CleanPointer(&tuplet.TupletNormal) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TupletPortion
func (tupletportion *TupletPortion) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&tupletportion.TupletDot) || modified
	// insertion point per field
	modified = stage.CleanPointer(&tupletportion.TupletNumber) || modified
	modified = stage.CleanPointer(&tupletportion.TupletType) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		if cleaner, ok := any(instance).(GongCleaner); ok {
			modified = cleaner.GongClean(stage) || modified
		}
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
