// generated code - do not edit
package direction

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
// Clean garbage collect unstaged instances that are referenced by Bass
func (bass *Bass) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&bass.BassSeparator) || modified
	modified = stage.CleanPointer(&bass.BassStep) || modified
	modified = stage.CleanPointer(&bass.BassAlter) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Degree
func (degree *Degree) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&degree.DegreeValue) || modified
	modified = stage.CleanPointer(&degree.DegreeAlter) || modified
	modified = stage.CleanPointer(&degree.DegreeType) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Direction
func (direction *Direction) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&direction.DirectionType) || modified
	// insertion point per field
	modified = stage.CleanPointer(&direction.Offset) || modified
	modified = stage.CleanPointer(&direction.Sound) || modified
	modified = stage.CleanPointer(&direction.Listening) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DirectionType
func (directiontype *DirectionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&directiontype.Percussion) || modified
	// insertion point per field
	modified = stage.CleanPointer(&directiontype.Wedge) || modified
	modified = stage.CleanPointer(&directiontype.Dashes) || modified
	modified = stage.CleanPointer(&directiontype.Bracket) || modified
	modified = stage.CleanPointer(&directiontype.Pedal) || modified
	modified = stage.CleanPointer(&directiontype.Metronome) || modified
	modified = stage.CleanPointer(&directiontype.OctaveShift) || modified
	modified = stage.CleanPointer(&directiontype.HarpPedals) || modified
	modified = stage.CleanPointer(&directiontype.StringMute) || modified
	modified = stage.CleanPointer(&directiontype.Scordatura) || modified
	modified = stage.CleanPointer(&directiontype.Image) || modified
	modified = stage.CleanPointer(&directiontype.PrincipalVoice) || modified
	modified = stage.CleanPointer(&directiontype.AccordionRegistration) || modified
	modified = stage.CleanPointer(&directiontype.StaffDivide) || modified
	modified = stage.CleanPointer(&directiontype.OtherDirection) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Frame
func (frame *Frame) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&frame.FrameNote) || modified
	// insertion point per field
	modified = stage.CleanPointer(&frame.FirstFret) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FrameNote
func (framenote *FrameNote) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&framenote.Barre) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Grouping
func (grouping *Grouping) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&grouping.Feature) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Harmony
func (harmony *Harmony) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&harmony.Root) || modified
	modified = stage.CleanSlice(&harmony.Numeral) || modified
	modified = stage.CleanSlice(&harmony.Function) || modified
	modified = stage.CleanSlice(&harmony.Kind) || modified
	modified = stage.CleanSlice(&harmony.Inversion) || modified
	modified = stage.CleanSlice(&harmony.Bass) || modified
	modified = stage.CleanSlice(&harmony.Degree) || modified
	// insertion point per field
	modified = stage.CleanPointer(&harmony.Frame) || modified
	modified = stage.CleanPointer(&harmony.Offset) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by HarpPedals
func (harppedals *HarpPedals) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&harppedals.PedalTuning) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by InstrumentChange
func (instrumentchange *InstrumentChange) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&instrumentchange.VirtualInstrument) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Listening
func (listening *Listening) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&listening.Sync) || modified
	modified = stage.CleanSlice(&listening.OtherListening) || modified
	// insertion point per field
	modified = stage.CleanPointer(&listening.Offset) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Metronome
func (metronome *Metronome) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&metronome.BeatUnitTied) || modified
	modified = stage.CleanSlice(&metronome.MetronomeNote) || modified
	// insertion point per field
	modified = stage.CleanPointer(&metronome.PerMinute) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MetronomeNote
func (metronomenote *MetronomeNote) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&metronomenote.MetronomeBeam) || modified
	// insertion point per field
	modified = stage.CleanPointer(&metronomenote.MetronomeTied) || modified
	modified = stage.CleanPointer(&metronomenote.MetronomeTuplet) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Numeral
func (numeral *Numeral) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&numeral.NumeralRoot) || modified
	modified = stage.CleanPointer(&numeral.NumeralAlter) || modified
	modified = stage.CleanPointer(&numeral.NumeralKey) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Percussion
func (percussion *Percussion) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&percussion.Glass) || modified
	modified = stage.CleanPointer(&percussion.Metal) || modified
	modified = stage.CleanPointer(&percussion.Wood) || modified
	modified = stage.CleanPointer(&percussion.Pitched) || modified
	modified = stage.CleanPointer(&percussion.Membrane) || modified
	modified = stage.CleanPointer(&percussion.Effect) || modified
	modified = stage.CleanPointer(&percussion.Timpani) || modified
	modified = stage.CleanPointer(&percussion.Beater) || modified
	modified = stage.CleanPointer(&percussion.Stick) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Print
func (print *Print) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&print.MeasureNumbering) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Root
func (root *Root) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&root.RootStep) || modified
	modified = stage.CleanPointer(&root.RootAlter) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Scordatura
func (scordatura *Scordatura) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&scordatura.Accord) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Sound
func (sound *Sound) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&sound.InstrumentChange) || modified
	// insertion point per field
	modified = stage.CleanPointer(&sound.Swing) || modified
	modified = stage.CleanPointer(&sound.Offset) || modified
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
