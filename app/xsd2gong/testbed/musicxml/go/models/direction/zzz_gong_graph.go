// generated code - do not edit
package direction

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
func (accord *Accord) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Accords[accord]
	return ok
}

func (accordionregistration *AccordionRegistration) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AccordionRegistrations[accordionregistration]
	return ok
}

func (barre *Barre) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Barres[barre]
	return ok
}

func (bass *Bass) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Basss[bass]
	return ok
}

func (bassstep *BassStep) GongIsStaged(stage *Stage) bool {
	_, ok := stage.BassSteps[bassstep]
	return ok
}

func (beatunittied *BeatUnitTied) GongIsStaged(stage *Stage) bool {
	_, ok := stage.BeatUnitTieds[beatunittied]
	return ok
}

func (beater *Beater) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Beaters[beater]
	return ok
}

func (bracket *Bracket) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Brackets[bracket]
	return ok
}

func (dashes *Dashes) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Dashess[dashes]
	return ok
}

func (degree *Degree) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Degrees[degree]
	return ok
}

func (degreealter *DegreeAlter) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DegreeAlters[degreealter]
	return ok
}

func (degreetype *DegreeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DegreeTypes[degreetype]
	return ok
}

func (degreevalue *DegreeValue) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DegreeValues[degreevalue]
	return ok
}

func (direction *Direction) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Directions[direction]
	return ok
}

func (directiontype *DirectionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DirectionTypes[directiontype]
	return ok
}

func (effect *Effect) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Effects[effect]
	return ok
}

func (feature *Feature) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Features[feature]
	return ok
}

func (firstfret *FirstFret) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FirstFrets[firstfret]
	return ok
}

func (frame *Frame) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Frames[frame]
	return ok
}

func (framenote *FrameNote) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FrameNotes[framenote]
	return ok
}

func (glass *Glass) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Glasss[glass]
	return ok
}

func (grouping *Grouping) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Groupings[grouping]
	return ok
}

func (harmony *Harmony) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Harmonys[harmony]
	return ok
}

func (harmonyalter *HarmonyAlter) GongIsStaged(stage *Stage) bool {
	_, ok := stage.HarmonyAlters[harmonyalter]
	return ok
}

func (harppedals *HarpPedals) GongIsStaged(stage *Stage) bool {
	_, ok := stage.HarpPedalss[harppedals]
	return ok
}

func (image *Image) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Images[image]
	return ok
}

func (instrumentchange *InstrumentChange) GongIsStaged(stage *Stage) bool {
	_, ok := stage.InstrumentChanges[instrumentchange]
	return ok
}

func (inversion *Inversion) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Inversions[inversion]
	return ok
}

func (kind *Kind) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Kinds[kind]
	return ok
}

func (listening *Listening) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Listenings[listening]
	return ok
}

func (measurenumbering *MeasureNumbering) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MeasureNumberings[measurenumbering]
	return ok
}

func (membrane *Membrane) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Membranes[membrane]
	return ok
}

func (metal *Metal) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Metals[metal]
	return ok
}

func (metronome *Metronome) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Metronomes[metronome]
	return ok
}

func (metronomebeam *MetronomeBeam) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MetronomeBeams[metronomebeam]
	return ok
}

func (metronomenote *MetronomeNote) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MetronomeNotes[metronomenote]
	return ok
}

func (metronometied *MetronomeTied) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MetronomeTieds[metronometied]
	return ok
}

func (metronometuplet *MetronomeTuplet) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MetronomeTuplets[metronometuplet]
	return ok
}

func (numeral *Numeral) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Numerals[numeral]
	return ok
}

func (numeralkey *NumeralKey) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NumeralKeys[numeralkey]
	return ok
}

func (numeralroot *NumeralRoot) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NumeralRoots[numeralroot]
	return ok
}

func (octaveshift *OctaveShift) GongIsStaged(stage *Stage) bool {
	_, ok := stage.OctaveShifts[octaveshift]
	return ok
}

func (offset *Offset) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Offsets[offset]
	return ok
}

func (otherdirection *OtherDirection) GongIsStaged(stage *Stage) bool {
	_, ok := stage.OtherDirections[otherdirection]
	return ok
}

func (otherlistening *OtherListening) GongIsStaged(stage *Stage) bool {
	_, ok := stage.OtherListenings[otherlistening]
	return ok
}

func (pedal *Pedal) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Pedals[pedal]
	return ok
}

func (pedaltuning *PedalTuning) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PedalTunings[pedaltuning]
	return ok
}

func (perminute *PerMinute) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PerMinutes[perminute]
	return ok
}

func (percussion *Percussion) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Percussions[percussion]
	return ok
}

func (pitched *Pitched) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Pitcheds[pitched]
	return ok
}

func (principalvoice *PrincipalVoice) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PrincipalVoices[principalvoice]
	return ok
}

func (print *Print) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Prints[print]
	return ok
}

func (root *Root) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Roots[root]
	return ok
}

func (rootstep *RootStep) GongIsStaged(stage *Stage) bool {
	_, ok := stage.RootSteps[rootstep]
	return ok
}

func (scordatura *Scordatura) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Scordaturas[scordatura]
	return ok
}

func (sound *Sound) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Sounds[sound]
	return ok
}

func (staffdivide *StaffDivide) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StaffDivides[staffdivide]
	return ok
}

func (stick *Stick) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Sticks[stick]
	return ok
}

func (stringmute *StringMute) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StringMutes[stringmute]
	return ok
}

func (styletext *StyleText) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StyleTexts[styletext]
	return ok
}

func (swing *Swing) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Swings[swing]
	return ok
}

func (sync *Sync) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Syncs[sync]
	return ok
}

func (timemodification *TimeModification) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TimeModifications[timemodification]
	return ok
}

func (timpani *Timpani) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Timpanis[timpani]
	return ok
}

func (virtualinstrument *VirtualInstrument) GongIsStaged(stage *Stage) bool {
	_, ok := stage.VirtualInstruments[virtualinstrument]
	return ok
}

func (wedge *Wedge) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Wedges[wedge]
	return ok
}

func (wood *Wood) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Woods[wood]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (accord *Accord) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(accord) {
		return
	}

	accord.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (accordionregistration *AccordionRegistration) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(accordionregistration) {
		return
	}

	accordionregistration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (barre *Barre) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(barre) {
		return
	}

	barre.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bass *Bass) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(bass) {
		return
	}

	bass.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if bass.BassSeparator != nil {
		stage.StageBranch(bass.BassSeparator)
	}
	if bass.BassStep != nil {
		stage.StageBranch(bass.BassStep)
	}
	if bass.BassAlter != nil {
		stage.StageBranch(bass.BassAlter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bassstep *BassStep) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(bassstep) {
		return
	}

	bassstep.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (beatunittied *BeatUnitTied) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(beatunittied) {
		return
	}

	beatunittied.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (beater *Beater) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(beater) {
		return
	}

	beater.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bracket *Bracket) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(bracket) {
		return
	}

	bracket.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dashes *Dashes) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(dashes) {
		return
	}

	dashes.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (degree *Degree) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(degree) {
		return
	}

	degree.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if degree.DegreeValue != nil {
		stage.StageBranch(degree.DegreeValue)
	}
	if degree.DegreeAlter != nil {
		stage.StageBranch(degree.DegreeAlter)
	}
	if degree.DegreeType != nil {
		stage.StageBranch(degree.DegreeType)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (degreealter *DegreeAlter) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(degreealter) {
		return
	}

	degreealter.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (degreetype *DegreeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(degreetype) {
		return
	}

	degreetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (degreevalue *DegreeValue) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(degreevalue) {
		return
	}

	degreevalue.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (direction *Direction) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(direction) {
		return
	}

	direction.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if direction.Offset != nil {
		stage.StageBranch(direction.Offset)
	}
	if direction.Sound != nil {
		stage.StageBranch(direction.Sound)
	}
	if direction.Listening != nil {
		stage.StageBranch(direction.Listening)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _directiontype := range direction.DirectionType {
		stage.StageBranch(_directiontype)
	}

}

func (directiontype *DirectionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(directiontype) {
		return
	}

	directiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if directiontype.Wedge != nil {
		stage.StageBranch(directiontype.Wedge)
	}
	if directiontype.Dashes != nil {
		stage.StageBranch(directiontype.Dashes)
	}
	if directiontype.Bracket != nil {
		stage.StageBranch(directiontype.Bracket)
	}
	if directiontype.Pedal != nil {
		stage.StageBranch(directiontype.Pedal)
	}
	if directiontype.Metronome != nil {
		stage.StageBranch(directiontype.Metronome)
	}
	if directiontype.OctaveShift != nil {
		stage.StageBranch(directiontype.OctaveShift)
	}
	if directiontype.HarpPedals != nil {
		stage.StageBranch(directiontype.HarpPedals)
	}
	if directiontype.StringMute != nil {
		stage.StageBranch(directiontype.StringMute)
	}
	if directiontype.Scordatura != nil {
		stage.StageBranch(directiontype.Scordatura)
	}
	if directiontype.Image != nil {
		stage.StageBranch(directiontype.Image)
	}
	if directiontype.PrincipalVoice != nil {
		stage.StageBranch(directiontype.PrincipalVoice)
	}
	if directiontype.AccordionRegistration != nil {
		stage.StageBranch(directiontype.AccordionRegistration)
	}
	if directiontype.StaffDivide != nil {
		stage.StageBranch(directiontype.StaffDivide)
	}
	if directiontype.OtherDirection != nil {
		stage.StageBranch(directiontype.OtherDirection)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _percussion := range directiontype.Percussion {
		stage.StageBranch(_percussion)
	}

}

func (effect *Effect) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(effect) {
		return
	}

	effect.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (feature *Feature) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(feature) {
		return
	}

	feature.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (firstfret *FirstFret) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(firstfret) {
		return
	}

	firstfret.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (frame *Frame) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(frame) {
		return
	}

	frame.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if frame.FirstFret != nil {
		stage.StageBranch(frame.FirstFret)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _framenote := range frame.FrameNote {
		stage.StageBranch(_framenote)
	}

}

func (framenote *FrameNote) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(framenote) {
		return
	}

	framenote.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if framenote.Barre != nil {
		stage.StageBranch(framenote.Barre)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (glass *Glass) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(glass) {
		return
	}

	glass.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (grouping *Grouping) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(grouping) {
		return
	}

	grouping.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _feature := range grouping.Feature {
		stage.StageBranch(_feature)
	}

}

func (harmony *Harmony) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(harmony) {
		return
	}

	harmony.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if harmony.Frame != nil {
		stage.StageBranch(harmony.Frame)
	}
	if harmony.Offset != nil {
		stage.StageBranch(harmony.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _root := range harmony.Root {
		stage.StageBranch(_root)
	}
	for _, _numeral := range harmony.Numeral {
		stage.StageBranch(_numeral)
	}
	for _, _styletext := range harmony.Function {
		stage.StageBranch(_styletext)
	}
	for _, _kind := range harmony.Kind {
		stage.StageBranch(_kind)
	}
	for _, _inversion := range harmony.Inversion {
		stage.StageBranch(_inversion)
	}
	for _, _bass := range harmony.Bass {
		stage.StageBranch(_bass)
	}
	for _, _degree := range harmony.Degree {
		stage.StageBranch(_degree)
	}

}

func (harmonyalter *HarmonyAlter) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(harmonyalter) {
		return
	}

	harmonyalter.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (harppedals *HarpPedals) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(harppedals) {
		return
	}

	harppedals.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _pedaltuning := range harppedals.PedalTuning {
		stage.StageBranch(_pedaltuning)
	}

}

func (image *Image) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(image) {
		return
	}

	image.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (instrumentchange *InstrumentChange) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(instrumentchange) {
		return
	}

	instrumentchange.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if instrumentchange.VirtualInstrument != nil {
		stage.StageBranch(instrumentchange.VirtualInstrument)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (inversion *Inversion) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(inversion) {
		return
	}

	inversion.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (kind *Kind) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(kind) {
		return
	}

	kind.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (listening *Listening) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(listening) {
		return
	}

	listening.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if listening.Offset != nil {
		stage.StageBranch(listening.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sync := range listening.Sync {
		stage.StageBranch(_sync)
	}
	for _, _otherlistening := range listening.OtherListening {
		stage.StageBranch(_otherlistening)
	}

}

func (measurenumbering *MeasureNumbering) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(measurenumbering) {
		return
	}

	measurenumbering.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (membrane *Membrane) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(membrane) {
		return
	}

	membrane.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metal *Metal) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(metal) {
		return
	}

	metal.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metronome *Metronome) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(metronome) {
		return
	}

	metronome.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if metronome.PerMinute != nil {
		stage.StageBranch(metronome.PerMinute)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _beatunittied := range metronome.BeatUnitTied {
		stage.StageBranch(_beatunittied)
	}
	for _, _metronomenote := range metronome.MetronomeNote {
		stage.StageBranch(_metronomenote)
	}

}

func (metronomebeam *MetronomeBeam) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(metronomebeam) {
		return
	}

	metronomebeam.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metronomenote *MetronomeNote) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(metronomenote) {
		return
	}

	metronomenote.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if metronomenote.MetronomeTied != nil {
		stage.StageBranch(metronomenote.MetronomeTied)
	}
	if metronomenote.MetronomeTuplet != nil {
		stage.StageBranch(metronomenote.MetronomeTuplet)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _metronomebeam := range metronomenote.MetronomeBeam {
		stage.StageBranch(_metronomebeam)
	}

}

func (metronometied *MetronomeTied) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(metronometied) {
		return
	}

	metronometied.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metronometuplet *MetronomeTuplet) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(metronometuplet) {
		return
	}

	metronometuplet.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (numeral *Numeral) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(numeral) {
		return
	}

	numeral.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if numeral.NumeralRoot != nil {
		stage.StageBranch(numeral.NumeralRoot)
	}
	if numeral.NumeralAlter != nil {
		stage.StageBranch(numeral.NumeralAlter)
	}
	if numeral.NumeralKey != nil {
		stage.StageBranch(numeral.NumeralKey)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (numeralkey *NumeralKey) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(numeralkey) {
		return
	}

	numeralkey.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (numeralroot *NumeralRoot) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(numeralroot) {
		return
	}

	numeralroot.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (octaveshift *OctaveShift) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(octaveshift) {
		return
	}

	octaveshift.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (offset *Offset) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(offset) {
		return
	}

	offset.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (otherdirection *OtherDirection) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(otherdirection) {
		return
	}

	otherdirection.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (otherlistening *OtherListening) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(otherlistening) {
		return
	}

	otherlistening.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pedal *Pedal) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(pedal) {
		return
	}

	pedal.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pedaltuning *PedalTuning) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(pedaltuning) {
		return
	}

	pedaltuning.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (perminute *PerMinute) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(perminute) {
		return
	}

	perminute.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (percussion *Percussion) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(percussion) {
		return
	}

	percussion.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if percussion.Glass != nil {
		stage.StageBranch(percussion.Glass)
	}
	if percussion.Metal != nil {
		stage.StageBranch(percussion.Metal)
	}
	if percussion.Wood != nil {
		stage.StageBranch(percussion.Wood)
	}
	if percussion.Pitched != nil {
		stage.StageBranch(percussion.Pitched)
	}
	if percussion.Membrane != nil {
		stage.StageBranch(percussion.Membrane)
	}
	if percussion.Effect != nil {
		stage.StageBranch(percussion.Effect)
	}
	if percussion.Timpani != nil {
		stage.StageBranch(percussion.Timpani)
	}
	if percussion.Beater != nil {
		stage.StageBranch(percussion.Beater)
	}
	if percussion.Stick != nil {
		stage.StageBranch(percussion.Stick)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pitched *Pitched) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(pitched) {
		return
	}

	pitched.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (principalvoice *PrincipalVoice) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(principalvoice) {
		return
	}

	principalvoice.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (print *Print) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(print) {
		return
	}

	print.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if print.MeasureNumbering != nil {
		stage.StageBranch(print.MeasureNumbering)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (root *Root) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(root) {
		return
	}

	root.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if root.RootStep != nil {
		stage.StageBranch(root.RootStep)
	}
	if root.RootAlter != nil {
		stage.StageBranch(root.RootAlter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rootstep *RootStep) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(rootstep) {
		return
	}

	rootstep.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (scordatura *Scordatura) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(scordatura) {
		return
	}

	scordatura.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _accord := range scordatura.Accord {
		stage.StageBranch(_accord)
	}

}

func (sound *Sound) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(sound) {
		return
	}

	sound.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if sound.Swing != nil {
		stage.StageBranch(sound.Swing)
	}
	if sound.Offset != nil {
		stage.StageBranch(sound.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrumentchange := range sound.InstrumentChange {
		stage.StageBranch(_instrumentchange)
	}

}

func (staffdivide *StaffDivide) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(staffdivide) {
		return
	}

	staffdivide.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stick *Stick) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stick) {
		return
	}

	stick.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stringmute *StringMute) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stringmute) {
		return
	}

	stringmute.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (styletext *StyleText) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(styletext) {
		return
	}

	styletext.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (swing *Swing) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(swing) {
		return
	}

	swing.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (sync *Sync) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(sync) {
		return
	}

	sync.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (timemodification *TimeModification) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(timemodification) {
		return
	}

	timemodification.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (timpani *Timpani) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(timpani) {
		return
	}

	timpani.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (virtualinstrument *VirtualInstrument) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(virtualinstrument) {
		return
	}

	virtualinstrument.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (wedge *Wedge) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(wedge) {
		return
	}

	wedge.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (wood *Wood) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(wood) {
		return
	}

	wood.Stage(stage)

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
	case *Accord:
		toT := GongCopyBranchAccord(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AccordionRegistration:
		toT := GongCopyBranchAccordionRegistration(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Barre:
		toT := GongCopyBranchBarre(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bass:
		toT := GongCopyBranchBass(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BassStep:
		toT := GongCopyBranchBassStep(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BeatUnitTied:
		toT := GongCopyBranchBeatUnitTied(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Beater:
		toT := GongCopyBranchBeater(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bracket:
		toT := GongCopyBranchBracket(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Dashes:
		toT := GongCopyBranchDashes(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Degree:
		toT := GongCopyBranchDegree(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DegreeAlter:
		toT := GongCopyBranchDegreeAlter(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DegreeType:
		toT := GongCopyBranchDegreeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DegreeValue:
		toT := GongCopyBranchDegreeValue(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Direction:
		toT := GongCopyBranchDirection(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DirectionType:
		toT := GongCopyBranchDirectionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Effect:
		toT := GongCopyBranchEffect(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Feature:
		toT := GongCopyBranchFeature(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FirstFret:
		toT := GongCopyBranchFirstFret(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Frame:
		toT := GongCopyBranchFrame(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FrameNote:
		toT := GongCopyBranchFrameNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Glass:
		toT := GongCopyBranchGlass(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Grouping:
		toT := GongCopyBranchGrouping(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Harmony:
		toT := GongCopyBranchHarmony(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *HarmonyAlter:
		toT := GongCopyBranchHarmonyAlter(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *HarpPedals:
		toT := GongCopyBranchHarpPedals(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Image:
		toT := GongCopyBranchImage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *InstrumentChange:
		toT := GongCopyBranchInstrumentChange(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Inversion:
		toT := GongCopyBranchInversion(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Kind:
		toT := GongCopyBranchKind(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Listening:
		toT := GongCopyBranchListening(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MeasureNumbering:
		toT := GongCopyBranchMeasureNumbering(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Membrane:
		toT := GongCopyBranchMembrane(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Metal:
		toT := GongCopyBranchMetal(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Metronome:
		toT := GongCopyBranchMetronome(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MetronomeBeam:
		toT := GongCopyBranchMetronomeBeam(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MetronomeNote:
		toT := GongCopyBranchMetronomeNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MetronomeTied:
		toT := GongCopyBranchMetronomeTied(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MetronomeTuplet:
		toT := GongCopyBranchMetronomeTuplet(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Numeral:
		toT := GongCopyBranchNumeral(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NumeralKey:
		toT := GongCopyBranchNumeralKey(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NumeralRoot:
		toT := GongCopyBranchNumeralRoot(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *OctaveShift:
		toT := GongCopyBranchOctaveShift(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Offset:
		toT := GongCopyBranchOffset(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *OtherDirection:
		toT := GongCopyBranchOtherDirection(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *OtherListening:
		toT := GongCopyBranchOtherListening(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Pedal:
		toT := GongCopyBranchPedal(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PedalTuning:
		toT := GongCopyBranchPedalTuning(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PerMinute:
		toT := GongCopyBranchPerMinute(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Percussion:
		toT := GongCopyBranchPercussion(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Pitched:
		toT := GongCopyBranchPitched(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PrincipalVoice:
		toT := GongCopyBranchPrincipalVoice(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Print:
		toT := GongCopyBranchPrint(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Root:
		toT := GongCopyBranchRoot(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RootStep:
		toT := GongCopyBranchRootStep(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Scordatura:
		toT := GongCopyBranchScordatura(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Sound:
		toT := GongCopyBranchSound(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StaffDivide:
		toT := GongCopyBranchStaffDivide(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Stick:
		toT := GongCopyBranchStick(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StringMute:
		toT := GongCopyBranchStringMute(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StyleText:
		toT := GongCopyBranchStyleText(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Swing:
		toT := GongCopyBranchSwing(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Sync:
		toT := GongCopyBranchSync(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TimeModification:
		toT := GongCopyBranchTimeModification(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Timpani:
		toT := GongCopyBranchTimpani(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VirtualInstrument:
		toT := GongCopyBranchVirtualInstrument(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Wedge:
		toT := GongCopyBranchWedge(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Wood:
		toT := GongCopyBranchWood(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAccord(mapOrigCopy map[any]any, accordFrom *Accord) (accordTo *Accord) {
	var alreadyCopied bool
	accordTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, accordFrom)
	if alreadyCopied {
		return
	}
	accordFrom.GongCopyBasicFields(accordTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAccordionRegistration(mapOrigCopy map[any]any, accordionregistrationFrom *AccordionRegistration) (accordionregistrationTo *AccordionRegistration) {
	var alreadyCopied bool
	accordionregistrationTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, accordionregistrationFrom)
	if alreadyCopied {
		return
	}
	accordionregistrationFrom.GongCopyBasicFields(accordionregistrationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBarre(mapOrigCopy map[any]any, barreFrom *Barre) (barreTo *Barre) {
	var alreadyCopied bool
	barreTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, barreFrom)
	if alreadyCopied {
		return
	}
	barreFrom.GongCopyBasicFields(barreTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBass(mapOrigCopy map[any]any, bassFrom *Bass) (bassTo *Bass) {
	var alreadyCopied bool
	bassTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, bassFrom)
	if alreadyCopied {
		return
	}
	bassFrom.GongCopyBasicFields(bassTo)

	//insertion point for the staging of instances referenced by pointers
	if bassFrom.BassSeparator != nil {
		bassTo.BassSeparator = GongCopyBranchStyleText(mapOrigCopy, bassFrom.BassSeparator)
	}
	if bassFrom.BassStep != nil {
		bassTo.BassStep = GongCopyBranchBassStep(mapOrigCopy, bassFrom.BassStep)
	}
	if bassFrom.BassAlter != nil {
		bassTo.BassAlter = GongCopyBranchHarmonyAlter(mapOrigCopy, bassFrom.BassAlter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBassStep(mapOrigCopy map[any]any, bassstepFrom *BassStep) (bassstepTo *BassStep) {
	var alreadyCopied bool
	bassstepTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, bassstepFrom)
	if alreadyCopied {
		return
	}
	bassstepFrom.GongCopyBasicFields(bassstepTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBeatUnitTied(mapOrigCopy map[any]any, beatunittiedFrom *BeatUnitTied) (beatunittiedTo *BeatUnitTied) {
	var alreadyCopied bool
	beatunittiedTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, beatunittiedFrom)
	if alreadyCopied {
		return
	}
	beatunittiedFrom.GongCopyBasicFields(beatunittiedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBeater(mapOrigCopy map[any]any, beaterFrom *Beater) (beaterTo *Beater) {
	var alreadyCopied bool
	beaterTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, beaterFrom)
	if alreadyCopied {
		return
	}
	beaterFrom.GongCopyBasicFields(beaterTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBracket(mapOrigCopy map[any]any, bracketFrom *Bracket) (bracketTo *Bracket) {
	var alreadyCopied bool
	bracketTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, bracketFrom)
	if alreadyCopied {
		return
	}
	bracketFrom.GongCopyBasicFields(bracketTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDashes(mapOrigCopy map[any]any, dashesFrom *Dashes) (dashesTo *Dashes) {
	var alreadyCopied bool
	dashesTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, dashesFrom)
	if alreadyCopied {
		return
	}
	dashesFrom.GongCopyBasicFields(dashesTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDegree(mapOrigCopy map[any]any, degreeFrom *Degree) (degreeTo *Degree) {
	var alreadyCopied bool
	degreeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, degreeFrom)
	if alreadyCopied {
		return
	}
	degreeFrom.GongCopyBasicFields(degreeTo)

	//insertion point for the staging of instances referenced by pointers
	if degreeFrom.DegreeValue != nil {
		degreeTo.DegreeValue = GongCopyBranchDegreeValue(mapOrigCopy, degreeFrom.DegreeValue)
	}
	if degreeFrom.DegreeAlter != nil {
		degreeTo.DegreeAlter = GongCopyBranchDegreeAlter(mapOrigCopy, degreeFrom.DegreeAlter)
	}
	if degreeFrom.DegreeType != nil {
		degreeTo.DegreeType = GongCopyBranchDegreeType(mapOrigCopy, degreeFrom.DegreeType)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDegreeAlter(mapOrigCopy map[any]any, degreealterFrom *DegreeAlter) (degreealterTo *DegreeAlter) {
	var alreadyCopied bool
	degreealterTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, degreealterFrom)
	if alreadyCopied {
		return
	}
	degreealterFrom.GongCopyBasicFields(degreealterTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDegreeType(mapOrigCopy map[any]any, degreetypeFrom *DegreeType) (degreetypeTo *DegreeType) {
	var alreadyCopied bool
	degreetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, degreetypeFrom)
	if alreadyCopied {
		return
	}
	degreetypeFrom.GongCopyBasicFields(degreetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDegreeValue(mapOrigCopy map[any]any, degreevalueFrom *DegreeValue) (degreevalueTo *DegreeValue) {
	var alreadyCopied bool
	degreevalueTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, degreevalueFrom)
	if alreadyCopied {
		return
	}
	degreevalueFrom.GongCopyBasicFields(degreevalueTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDirection(mapOrigCopy map[any]any, directionFrom *Direction) (directionTo *Direction) {
	var alreadyCopied bool
	directionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, directionFrom)
	if alreadyCopied {
		return
	}
	directionFrom.GongCopyBasicFields(directionTo)

	//insertion point for the staging of instances referenced by pointers
	if directionFrom.Offset != nil {
		directionTo.Offset = GongCopyBranchOffset(mapOrigCopy, directionFrom.Offset)
	}
	if directionFrom.Sound != nil {
		directionTo.Sound = GongCopyBranchSound(mapOrigCopy, directionFrom.Sound)
	}
	if directionFrom.Listening != nil {
		directionTo.Listening = GongCopyBranchListening(mapOrigCopy, directionFrom.Listening)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _directiontype := range directionFrom.DirectionType {
		directionTo.DirectionType = append(directionTo.DirectionType, GongCopyBranchDirectionType(mapOrigCopy, _directiontype))
	}

	return
}

func GongCopyBranchDirectionType(mapOrigCopy map[any]any, directiontypeFrom *DirectionType) (directiontypeTo *DirectionType) {
	var alreadyCopied bool
	directiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, directiontypeFrom)
	if alreadyCopied {
		return
	}
	directiontypeFrom.GongCopyBasicFields(directiontypeTo)

	//insertion point for the staging of instances referenced by pointers
	if directiontypeFrom.Wedge != nil {
		directiontypeTo.Wedge = GongCopyBranchWedge(mapOrigCopy, directiontypeFrom.Wedge)
	}
	if directiontypeFrom.Dashes != nil {
		directiontypeTo.Dashes = GongCopyBranchDashes(mapOrigCopy, directiontypeFrom.Dashes)
	}
	if directiontypeFrom.Bracket != nil {
		directiontypeTo.Bracket = GongCopyBranchBracket(mapOrigCopy, directiontypeFrom.Bracket)
	}
	if directiontypeFrom.Pedal != nil {
		directiontypeTo.Pedal = GongCopyBranchPedal(mapOrigCopy, directiontypeFrom.Pedal)
	}
	if directiontypeFrom.Metronome != nil {
		directiontypeTo.Metronome = GongCopyBranchMetronome(mapOrigCopy, directiontypeFrom.Metronome)
	}
	if directiontypeFrom.OctaveShift != nil {
		directiontypeTo.OctaveShift = GongCopyBranchOctaveShift(mapOrigCopy, directiontypeFrom.OctaveShift)
	}
	if directiontypeFrom.HarpPedals != nil {
		directiontypeTo.HarpPedals = GongCopyBranchHarpPedals(mapOrigCopy, directiontypeFrom.HarpPedals)
	}
	if directiontypeFrom.StringMute != nil {
		directiontypeTo.StringMute = GongCopyBranchStringMute(mapOrigCopy, directiontypeFrom.StringMute)
	}
	if directiontypeFrom.Scordatura != nil {
		directiontypeTo.Scordatura = GongCopyBranchScordatura(mapOrigCopy, directiontypeFrom.Scordatura)
	}
	if directiontypeFrom.Image != nil {
		directiontypeTo.Image = GongCopyBranchImage(mapOrigCopy, directiontypeFrom.Image)
	}
	if directiontypeFrom.PrincipalVoice != nil {
		directiontypeTo.PrincipalVoice = GongCopyBranchPrincipalVoice(mapOrigCopy, directiontypeFrom.PrincipalVoice)
	}
	if directiontypeFrom.AccordionRegistration != nil {
		directiontypeTo.AccordionRegistration = GongCopyBranchAccordionRegistration(mapOrigCopy, directiontypeFrom.AccordionRegistration)
	}
	if directiontypeFrom.StaffDivide != nil {
		directiontypeTo.StaffDivide = GongCopyBranchStaffDivide(mapOrigCopy, directiontypeFrom.StaffDivide)
	}
	if directiontypeFrom.OtherDirection != nil {
		directiontypeTo.OtherDirection = GongCopyBranchOtherDirection(mapOrigCopy, directiontypeFrom.OtherDirection)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _percussion := range directiontypeFrom.Percussion {
		directiontypeTo.Percussion = append(directiontypeTo.Percussion, GongCopyBranchPercussion(mapOrigCopy, _percussion))
	}

	return
}

func GongCopyBranchEffect(mapOrigCopy map[any]any, effectFrom *Effect) (effectTo *Effect) {
	var alreadyCopied bool
	effectTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, effectFrom)
	if alreadyCopied {
		return
	}
	effectFrom.GongCopyBasicFields(effectTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFeature(mapOrigCopy map[any]any, featureFrom *Feature) (featureTo *Feature) {
	var alreadyCopied bool
	featureTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, featureFrom)
	if alreadyCopied {
		return
	}
	featureFrom.GongCopyBasicFields(featureTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFirstFret(mapOrigCopy map[any]any, firstfretFrom *FirstFret) (firstfretTo *FirstFret) {
	var alreadyCopied bool
	firstfretTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, firstfretFrom)
	if alreadyCopied {
		return
	}
	firstfretFrom.GongCopyBasicFields(firstfretTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFrame(mapOrigCopy map[any]any, frameFrom *Frame) (frameTo *Frame) {
	var alreadyCopied bool
	frameTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, frameFrom)
	if alreadyCopied {
		return
	}
	frameFrom.GongCopyBasicFields(frameTo)

	//insertion point for the staging of instances referenced by pointers
	if frameFrom.FirstFret != nil {
		frameTo.FirstFret = GongCopyBranchFirstFret(mapOrigCopy, frameFrom.FirstFret)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _framenote := range frameFrom.FrameNote {
		frameTo.FrameNote = append(frameTo.FrameNote, GongCopyBranchFrameNote(mapOrigCopy, _framenote))
	}

	return
}

func GongCopyBranchFrameNote(mapOrigCopy map[any]any, framenoteFrom *FrameNote) (framenoteTo *FrameNote) {
	var alreadyCopied bool
	framenoteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, framenoteFrom)
	if alreadyCopied {
		return
	}
	framenoteFrom.GongCopyBasicFields(framenoteTo)

	//insertion point for the staging of instances referenced by pointers
	if framenoteFrom.Barre != nil {
		framenoteTo.Barre = GongCopyBranchBarre(mapOrigCopy, framenoteFrom.Barre)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGlass(mapOrigCopy map[any]any, glassFrom *Glass) (glassTo *Glass) {
	var alreadyCopied bool
	glassTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, glassFrom)
	if alreadyCopied {
		return
	}
	glassFrom.GongCopyBasicFields(glassTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGrouping(mapOrigCopy map[any]any, groupingFrom *Grouping) (groupingTo *Grouping) {
	var alreadyCopied bool
	groupingTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, groupingFrom)
	if alreadyCopied {
		return
	}
	groupingFrom.GongCopyBasicFields(groupingTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _feature := range groupingFrom.Feature {
		groupingTo.Feature = append(groupingTo.Feature, GongCopyBranchFeature(mapOrigCopy, _feature))
	}

	return
}

func GongCopyBranchHarmony(mapOrigCopy map[any]any, harmonyFrom *Harmony) (harmonyTo *Harmony) {
	var alreadyCopied bool
	harmonyTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, harmonyFrom)
	if alreadyCopied {
		return
	}
	harmonyFrom.GongCopyBasicFields(harmonyTo)

	//insertion point for the staging of instances referenced by pointers
	if harmonyFrom.Frame != nil {
		harmonyTo.Frame = GongCopyBranchFrame(mapOrigCopy, harmonyFrom.Frame)
	}
	if harmonyFrom.Offset != nil {
		harmonyTo.Offset = GongCopyBranchOffset(mapOrigCopy, harmonyFrom.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _root := range harmonyFrom.Root {
		harmonyTo.Root = append(harmonyTo.Root, GongCopyBranchRoot(mapOrigCopy, _root))
	}
	for _, _numeral := range harmonyFrom.Numeral {
		harmonyTo.Numeral = append(harmonyTo.Numeral, GongCopyBranchNumeral(mapOrigCopy, _numeral))
	}
	for _, _styletext := range harmonyFrom.Function {
		harmonyTo.Function = append(harmonyTo.Function, GongCopyBranchStyleText(mapOrigCopy, _styletext))
	}
	for _, _kind := range harmonyFrom.Kind {
		harmonyTo.Kind = append(harmonyTo.Kind, GongCopyBranchKind(mapOrigCopy, _kind))
	}
	for _, _inversion := range harmonyFrom.Inversion {
		harmonyTo.Inversion = append(harmonyTo.Inversion, GongCopyBranchInversion(mapOrigCopy, _inversion))
	}
	for _, _bass := range harmonyFrom.Bass {
		harmonyTo.Bass = append(harmonyTo.Bass, GongCopyBranchBass(mapOrigCopy, _bass))
	}
	for _, _degree := range harmonyFrom.Degree {
		harmonyTo.Degree = append(harmonyTo.Degree, GongCopyBranchDegree(mapOrigCopy, _degree))
	}

	return
}

func GongCopyBranchHarmonyAlter(mapOrigCopy map[any]any, harmonyalterFrom *HarmonyAlter) (harmonyalterTo *HarmonyAlter) {
	var alreadyCopied bool
	harmonyalterTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, harmonyalterFrom)
	if alreadyCopied {
		return
	}
	harmonyalterFrom.GongCopyBasicFields(harmonyalterTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchHarpPedals(mapOrigCopy map[any]any, harppedalsFrom *HarpPedals) (harppedalsTo *HarpPedals) {
	var alreadyCopied bool
	harppedalsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, harppedalsFrom)
	if alreadyCopied {
		return
	}
	harppedalsFrom.GongCopyBasicFields(harppedalsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _pedaltuning := range harppedalsFrom.PedalTuning {
		harppedalsTo.PedalTuning = append(harppedalsTo.PedalTuning, GongCopyBranchPedalTuning(mapOrigCopy, _pedaltuning))
	}

	return
}

func GongCopyBranchImage(mapOrigCopy map[any]any, imageFrom *Image) (imageTo *Image) {
	var alreadyCopied bool
	imageTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, imageFrom)
	if alreadyCopied {
		return
	}
	imageFrom.GongCopyBasicFields(imageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchInstrumentChange(mapOrigCopy map[any]any, instrumentchangeFrom *InstrumentChange) (instrumentchangeTo *InstrumentChange) {
	var alreadyCopied bool
	instrumentchangeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, instrumentchangeFrom)
	if alreadyCopied {
		return
	}
	instrumentchangeFrom.GongCopyBasicFields(instrumentchangeTo)

	//insertion point for the staging of instances referenced by pointers
	if instrumentchangeFrom.VirtualInstrument != nil {
		instrumentchangeTo.VirtualInstrument = GongCopyBranchVirtualInstrument(mapOrigCopy, instrumentchangeFrom.VirtualInstrument)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchInversion(mapOrigCopy map[any]any, inversionFrom *Inversion) (inversionTo *Inversion) {
	var alreadyCopied bool
	inversionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, inversionFrom)
	if alreadyCopied {
		return
	}
	inversionFrom.GongCopyBasicFields(inversionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchKind(mapOrigCopy map[any]any, kindFrom *Kind) (kindTo *Kind) {
	var alreadyCopied bool
	kindTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, kindFrom)
	if alreadyCopied {
		return
	}
	kindFrom.GongCopyBasicFields(kindTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchListening(mapOrigCopy map[any]any, listeningFrom *Listening) (listeningTo *Listening) {
	var alreadyCopied bool
	listeningTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, listeningFrom)
	if alreadyCopied {
		return
	}
	listeningFrom.GongCopyBasicFields(listeningTo)

	//insertion point for the staging of instances referenced by pointers
	if listeningFrom.Offset != nil {
		listeningTo.Offset = GongCopyBranchOffset(mapOrigCopy, listeningFrom.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sync := range listeningFrom.Sync {
		listeningTo.Sync = append(listeningTo.Sync, GongCopyBranchSync(mapOrigCopy, _sync))
	}
	for _, _otherlistening := range listeningFrom.OtherListening {
		listeningTo.OtherListening = append(listeningTo.OtherListening, GongCopyBranchOtherListening(mapOrigCopy, _otherlistening))
	}

	return
}

func GongCopyBranchMeasureNumbering(mapOrigCopy map[any]any, measurenumberingFrom *MeasureNumbering) (measurenumberingTo *MeasureNumbering) {
	var alreadyCopied bool
	measurenumberingTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, measurenumberingFrom)
	if alreadyCopied {
		return
	}
	measurenumberingFrom.GongCopyBasicFields(measurenumberingTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMembrane(mapOrigCopy map[any]any, membraneFrom *Membrane) (membraneTo *Membrane) {
	var alreadyCopied bool
	membraneTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, membraneFrom)
	if alreadyCopied {
		return
	}
	membraneFrom.GongCopyBasicFields(membraneTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMetal(mapOrigCopy map[any]any, metalFrom *Metal) (metalTo *Metal) {
	var alreadyCopied bool
	metalTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, metalFrom)
	if alreadyCopied {
		return
	}
	metalFrom.GongCopyBasicFields(metalTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMetronome(mapOrigCopy map[any]any, metronomeFrom *Metronome) (metronomeTo *Metronome) {
	var alreadyCopied bool
	metronomeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, metronomeFrom)
	if alreadyCopied {
		return
	}
	metronomeFrom.GongCopyBasicFields(metronomeTo)

	//insertion point for the staging of instances referenced by pointers
	if metronomeFrom.PerMinute != nil {
		metronomeTo.PerMinute = GongCopyBranchPerMinute(mapOrigCopy, metronomeFrom.PerMinute)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _beatunittied := range metronomeFrom.BeatUnitTied {
		metronomeTo.BeatUnitTied = append(metronomeTo.BeatUnitTied, GongCopyBranchBeatUnitTied(mapOrigCopy, _beatunittied))
	}
	for _, _metronomenote := range metronomeFrom.MetronomeNote {
		metronomeTo.MetronomeNote = append(metronomeTo.MetronomeNote, GongCopyBranchMetronomeNote(mapOrigCopy, _metronomenote))
	}

	return
}

func GongCopyBranchMetronomeBeam(mapOrigCopy map[any]any, metronomebeamFrom *MetronomeBeam) (metronomebeamTo *MetronomeBeam) {
	var alreadyCopied bool
	metronomebeamTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, metronomebeamFrom)
	if alreadyCopied {
		return
	}
	metronomebeamFrom.GongCopyBasicFields(metronomebeamTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMetronomeNote(mapOrigCopy map[any]any, metronomenoteFrom *MetronomeNote) (metronomenoteTo *MetronomeNote) {
	var alreadyCopied bool
	metronomenoteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, metronomenoteFrom)
	if alreadyCopied {
		return
	}
	metronomenoteFrom.GongCopyBasicFields(metronomenoteTo)

	//insertion point for the staging of instances referenced by pointers
	if metronomenoteFrom.MetronomeTied != nil {
		metronomenoteTo.MetronomeTied = GongCopyBranchMetronomeTied(mapOrigCopy, metronomenoteFrom.MetronomeTied)
	}
	if metronomenoteFrom.MetronomeTuplet != nil {
		metronomenoteTo.MetronomeTuplet = GongCopyBranchMetronomeTuplet(mapOrigCopy, metronomenoteFrom.MetronomeTuplet)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _metronomebeam := range metronomenoteFrom.MetronomeBeam {
		metronomenoteTo.MetronomeBeam = append(metronomenoteTo.MetronomeBeam, GongCopyBranchMetronomeBeam(mapOrigCopy, _metronomebeam))
	}

	return
}

func GongCopyBranchMetronomeTied(mapOrigCopy map[any]any, metronometiedFrom *MetronomeTied) (metronometiedTo *MetronomeTied) {
	var alreadyCopied bool
	metronometiedTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, metronometiedFrom)
	if alreadyCopied {
		return
	}
	metronometiedFrom.GongCopyBasicFields(metronometiedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMetronomeTuplet(mapOrigCopy map[any]any, metronometupletFrom *MetronomeTuplet) (metronometupletTo *MetronomeTuplet) {
	var alreadyCopied bool
	metronometupletTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, metronometupletFrom)
	if alreadyCopied {
		return
	}
	metronometupletFrom.GongCopyBasicFields(metronometupletTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNumeral(mapOrigCopy map[any]any, numeralFrom *Numeral) (numeralTo *Numeral) {
	var alreadyCopied bool
	numeralTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, numeralFrom)
	if alreadyCopied {
		return
	}
	numeralFrom.GongCopyBasicFields(numeralTo)

	//insertion point for the staging of instances referenced by pointers
	if numeralFrom.NumeralRoot != nil {
		numeralTo.NumeralRoot = GongCopyBranchNumeralRoot(mapOrigCopy, numeralFrom.NumeralRoot)
	}
	if numeralFrom.NumeralAlter != nil {
		numeralTo.NumeralAlter = GongCopyBranchHarmonyAlter(mapOrigCopy, numeralFrom.NumeralAlter)
	}
	if numeralFrom.NumeralKey != nil {
		numeralTo.NumeralKey = GongCopyBranchNumeralKey(mapOrigCopy, numeralFrom.NumeralKey)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNumeralKey(mapOrigCopy map[any]any, numeralkeyFrom *NumeralKey) (numeralkeyTo *NumeralKey) {
	var alreadyCopied bool
	numeralkeyTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, numeralkeyFrom)
	if alreadyCopied {
		return
	}
	numeralkeyFrom.GongCopyBasicFields(numeralkeyTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNumeralRoot(mapOrigCopy map[any]any, numeralrootFrom *NumeralRoot) (numeralrootTo *NumeralRoot) {
	var alreadyCopied bool
	numeralrootTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, numeralrootFrom)
	if alreadyCopied {
		return
	}
	numeralrootFrom.GongCopyBasicFields(numeralrootTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOctaveShift(mapOrigCopy map[any]any, octaveshiftFrom *OctaveShift) (octaveshiftTo *OctaveShift) {
	var alreadyCopied bool
	octaveshiftTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, octaveshiftFrom)
	if alreadyCopied {
		return
	}
	octaveshiftFrom.GongCopyBasicFields(octaveshiftTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOffset(mapOrigCopy map[any]any, offsetFrom *Offset) (offsetTo *Offset) {
	var alreadyCopied bool
	offsetTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, offsetFrom)
	if alreadyCopied {
		return
	}
	offsetFrom.GongCopyBasicFields(offsetTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOtherDirection(mapOrigCopy map[any]any, otherdirectionFrom *OtherDirection) (otherdirectionTo *OtherDirection) {
	var alreadyCopied bool
	otherdirectionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, otherdirectionFrom)
	if alreadyCopied {
		return
	}
	otherdirectionFrom.GongCopyBasicFields(otherdirectionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOtherListening(mapOrigCopy map[any]any, otherlisteningFrom *OtherListening) (otherlisteningTo *OtherListening) {
	var alreadyCopied bool
	otherlisteningTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, otherlisteningFrom)
	if alreadyCopied {
		return
	}
	otherlisteningFrom.GongCopyBasicFields(otherlisteningTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPedal(mapOrigCopy map[any]any, pedalFrom *Pedal) (pedalTo *Pedal) {
	var alreadyCopied bool
	pedalTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pedalFrom)
	if alreadyCopied {
		return
	}
	pedalFrom.GongCopyBasicFields(pedalTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPedalTuning(mapOrigCopy map[any]any, pedaltuningFrom *PedalTuning) (pedaltuningTo *PedalTuning) {
	var alreadyCopied bool
	pedaltuningTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pedaltuningFrom)
	if alreadyCopied {
		return
	}
	pedaltuningFrom.GongCopyBasicFields(pedaltuningTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPerMinute(mapOrigCopy map[any]any, perminuteFrom *PerMinute) (perminuteTo *PerMinute) {
	var alreadyCopied bool
	perminuteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, perminuteFrom)
	if alreadyCopied {
		return
	}
	perminuteFrom.GongCopyBasicFields(perminuteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPercussion(mapOrigCopy map[any]any, percussionFrom *Percussion) (percussionTo *Percussion) {
	var alreadyCopied bool
	percussionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, percussionFrom)
	if alreadyCopied {
		return
	}
	percussionFrom.GongCopyBasicFields(percussionTo)

	//insertion point for the staging of instances referenced by pointers
	if percussionFrom.Glass != nil {
		percussionTo.Glass = GongCopyBranchGlass(mapOrigCopy, percussionFrom.Glass)
	}
	if percussionFrom.Metal != nil {
		percussionTo.Metal = GongCopyBranchMetal(mapOrigCopy, percussionFrom.Metal)
	}
	if percussionFrom.Wood != nil {
		percussionTo.Wood = GongCopyBranchWood(mapOrigCopy, percussionFrom.Wood)
	}
	if percussionFrom.Pitched != nil {
		percussionTo.Pitched = GongCopyBranchPitched(mapOrigCopy, percussionFrom.Pitched)
	}
	if percussionFrom.Membrane != nil {
		percussionTo.Membrane = GongCopyBranchMembrane(mapOrigCopy, percussionFrom.Membrane)
	}
	if percussionFrom.Effect != nil {
		percussionTo.Effect = GongCopyBranchEffect(mapOrigCopy, percussionFrom.Effect)
	}
	if percussionFrom.Timpani != nil {
		percussionTo.Timpani = GongCopyBranchTimpani(mapOrigCopy, percussionFrom.Timpani)
	}
	if percussionFrom.Beater != nil {
		percussionTo.Beater = GongCopyBranchBeater(mapOrigCopy, percussionFrom.Beater)
	}
	if percussionFrom.Stick != nil {
		percussionTo.Stick = GongCopyBranchStick(mapOrigCopy, percussionFrom.Stick)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPitched(mapOrigCopy map[any]any, pitchedFrom *Pitched) (pitchedTo *Pitched) {
	var alreadyCopied bool
	pitchedTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pitchedFrom)
	if alreadyCopied {
		return
	}
	pitchedFrom.GongCopyBasicFields(pitchedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPrincipalVoice(mapOrigCopy map[any]any, principalvoiceFrom *PrincipalVoice) (principalvoiceTo *PrincipalVoice) {
	var alreadyCopied bool
	principalvoiceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, principalvoiceFrom)
	if alreadyCopied {
		return
	}
	principalvoiceFrom.GongCopyBasicFields(principalvoiceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPrint(mapOrigCopy map[any]any, printFrom *Print) (printTo *Print) {
	var alreadyCopied bool
	printTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, printFrom)
	if alreadyCopied {
		return
	}
	printFrom.GongCopyBasicFields(printTo)

	//insertion point for the staging of instances referenced by pointers
	if printFrom.MeasureNumbering != nil {
		printTo.MeasureNumbering = GongCopyBranchMeasureNumbering(mapOrigCopy, printFrom.MeasureNumbering)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRoot(mapOrigCopy map[any]any, rootFrom *Root) (rootTo *Root) {
	var alreadyCopied bool
	rootTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, rootFrom)
	if alreadyCopied {
		return
	}
	rootFrom.GongCopyBasicFields(rootTo)

	//insertion point for the staging of instances referenced by pointers
	if rootFrom.RootStep != nil {
		rootTo.RootStep = GongCopyBranchRootStep(mapOrigCopy, rootFrom.RootStep)
	}
	if rootFrom.RootAlter != nil {
		rootTo.RootAlter = GongCopyBranchHarmonyAlter(mapOrigCopy, rootFrom.RootAlter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRootStep(mapOrigCopy map[any]any, rootstepFrom *RootStep) (rootstepTo *RootStep) {
	var alreadyCopied bool
	rootstepTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, rootstepFrom)
	if alreadyCopied {
		return
	}
	rootstepFrom.GongCopyBasicFields(rootstepTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchScordatura(mapOrigCopy map[any]any, scordaturaFrom *Scordatura) (scordaturaTo *Scordatura) {
	var alreadyCopied bool
	scordaturaTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, scordaturaFrom)
	if alreadyCopied {
		return
	}
	scordaturaFrom.GongCopyBasicFields(scordaturaTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _accord := range scordaturaFrom.Accord {
		scordaturaTo.Accord = append(scordaturaTo.Accord, GongCopyBranchAccord(mapOrigCopy, _accord))
	}

	return
}

func GongCopyBranchSound(mapOrigCopy map[any]any, soundFrom *Sound) (soundTo *Sound) {
	var alreadyCopied bool
	soundTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, soundFrom)
	if alreadyCopied {
		return
	}
	soundFrom.GongCopyBasicFields(soundTo)

	//insertion point for the staging of instances referenced by pointers
	if soundFrom.Swing != nil {
		soundTo.Swing = GongCopyBranchSwing(mapOrigCopy, soundFrom.Swing)
	}
	if soundFrom.Offset != nil {
		soundTo.Offset = GongCopyBranchOffset(mapOrigCopy, soundFrom.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrumentchange := range soundFrom.InstrumentChange {
		soundTo.InstrumentChange = append(soundTo.InstrumentChange, GongCopyBranchInstrumentChange(mapOrigCopy, _instrumentchange))
	}

	return
}

func GongCopyBranchStaffDivide(mapOrigCopy map[any]any, staffdivideFrom *StaffDivide) (staffdivideTo *StaffDivide) {
	var alreadyCopied bool
	staffdivideTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, staffdivideFrom)
	if alreadyCopied {
		return
	}
	staffdivideFrom.GongCopyBasicFields(staffdivideTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStick(mapOrigCopy map[any]any, stickFrom *Stick) (stickTo *Stick) {
	var alreadyCopied bool
	stickTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stickFrom)
	if alreadyCopied {
		return
	}
	stickFrom.GongCopyBasicFields(stickTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStringMute(mapOrigCopy map[any]any, stringmuteFrom *StringMute) (stringmuteTo *StringMute) {
	var alreadyCopied bool
	stringmuteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stringmuteFrom)
	if alreadyCopied {
		return
	}
	stringmuteFrom.GongCopyBasicFields(stringmuteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStyleText(mapOrigCopy map[any]any, styletextFrom *StyleText) (styletextTo *StyleText) {
	var alreadyCopied bool
	styletextTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, styletextFrom)
	if alreadyCopied {
		return
	}
	styletextFrom.GongCopyBasicFields(styletextTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSwing(mapOrigCopy map[any]any, swingFrom *Swing) (swingTo *Swing) {
	var alreadyCopied bool
	swingTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, swingFrom)
	if alreadyCopied {
		return
	}
	swingFrom.GongCopyBasicFields(swingTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSync(mapOrigCopy map[any]any, syncFrom *Sync) (syncTo *Sync) {
	var alreadyCopied bool
	syncTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, syncFrom)
	if alreadyCopied {
		return
	}
	syncFrom.GongCopyBasicFields(syncTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTimeModification(mapOrigCopy map[any]any, timemodificationFrom *TimeModification) (timemodificationTo *TimeModification) {
	var alreadyCopied bool
	timemodificationTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, timemodificationFrom)
	if alreadyCopied {
		return
	}
	timemodificationFrom.GongCopyBasicFields(timemodificationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTimpani(mapOrigCopy map[any]any, timpaniFrom *Timpani) (timpaniTo *Timpani) {
	var alreadyCopied bool
	timpaniTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, timpaniFrom)
	if alreadyCopied {
		return
	}
	timpaniFrom.GongCopyBasicFields(timpaniTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchVirtualInstrument(mapOrigCopy map[any]any, virtualinstrumentFrom *VirtualInstrument) (virtualinstrumentTo *VirtualInstrument) {
	var alreadyCopied bool
	virtualinstrumentTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, virtualinstrumentFrom)
	if alreadyCopied {
		return
	}
	virtualinstrumentFrom.GongCopyBasicFields(virtualinstrumentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchWedge(mapOrigCopy map[any]any, wedgeFrom *Wedge) (wedgeTo *Wedge) {
	var alreadyCopied bool
	wedgeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, wedgeFrom)
	if alreadyCopied {
		return
	}
	wedgeFrom.GongCopyBasicFields(wedgeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchWood(mapOrigCopy map[any]any, woodFrom *Wood) (woodTo *Wood) {
	var alreadyCopied bool
	woodTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, woodFrom)
	if alreadyCopied {
		return
	}
	woodFrom.GongCopyBasicFields(woodTo)

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
func (accord *Accord) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(accord) {
		return
	}

	accord.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (accordionregistration *AccordionRegistration) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(accordionregistration) {
		return
	}

	accordionregistration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (barre *Barre) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(barre) {
		return
	}

	barre.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bass *Bass) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(bass) {
		return
	}

	bass.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if bass.BassSeparator != nil {
		stage.UnstageBranch(bass.BassSeparator)
	}
	if bass.BassStep != nil {
		stage.UnstageBranch(bass.BassStep)
	}
	if bass.BassAlter != nil {
		stage.UnstageBranch(bass.BassAlter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bassstep *BassStep) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(bassstep) {
		return
	}

	bassstep.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (beatunittied *BeatUnitTied) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(beatunittied) {
		return
	}

	beatunittied.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (beater *Beater) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(beater) {
		return
	}

	beater.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bracket *Bracket) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(bracket) {
		return
	}

	bracket.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dashes *Dashes) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(dashes) {
		return
	}

	dashes.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (degree *Degree) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(degree) {
		return
	}

	degree.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if degree.DegreeValue != nil {
		stage.UnstageBranch(degree.DegreeValue)
	}
	if degree.DegreeAlter != nil {
		stage.UnstageBranch(degree.DegreeAlter)
	}
	if degree.DegreeType != nil {
		stage.UnstageBranch(degree.DegreeType)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (degreealter *DegreeAlter) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(degreealter) {
		return
	}

	degreealter.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (degreetype *DegreeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(degreetype) {
		return
	}

	degreetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (degreevalue *DegreeValue) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(degreevalue) {
		return
	}

	degreevalue.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (direction *Direction) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(direction) {
		return
	}

	direction.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if direction.Offset != nil {
		stage.UnstageBranch(direction.Offset)
	}
	if direction.Sound != nil {
		stage.UnstageBranch(direction.Sound)
	}
	if direction.Listening != nil {
		stage.UnstageBranch(direction.Listening)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _directiontype := range direction.DirectionType {
		stage.UnstageBranch(_directiontype)
	}

}

func (directiontype *DirectionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(directiontype) {
		return
	}

	directiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if directiontype.Wedge != nil {
		stage.UnstageBranch(directiontype.Wedge)
	}
	if directiontype.Dashes != nil {
		stage.UnstageBranch(directiontype.Dashes)
	}
	if directiontype.Bracket != nil {
		stage.UnstageBranch(directiontype.Bracket)
	}
	if directiontype.Pedal != nil {
		stage.UnstageBranch(directiontype.Pedal)
	}
	if directiontype.Metronome != nil {
		stage.UnstageBranch(directiontype.Metronome)
	}
	if directiontype.OctaveShift != nil {
		stage.UnstageBranch(directiontype.OctaveShift)
	}
	if directiontype.HarpPedals != nil {
		stage.UnstageBranch(directiontype.HarpPedals)
	}
	if directiontype.StringMute != nil {
		stage.UnstageBranch(directiontype.StringMute)
	}
	if directiontype.Scordatura != nil {
		stage.UnstageBranch(directiontype.Scordatura)
	}
	if directiontype.Image != nil {
		stage.UnstageBranch(directiontype.Image)
	}
	if directiontype.PrincipalVoice != nil {
		stage.UnstageBranch(directiontype.PrincipalVoice)
	}
	if directiontype.AccordionRegistration != nil {
		stage.UnstageBranch(directiontype.AccordionRegistration)
	}
	if directiontype.StaffDivide != nil {
		stage.UnstageBranch(directiontype.StaffDivide)
	}
	if directiontype.OtherDirection != nil {
		stage.UnstageBranch(directiontype.OtherDirection)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _percussion := range directiontype.Percussion {
		stage.UnstageBranch(_percussion)
	}

}

func (effect *Effect) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(effect) {
		return
	}

	effect.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (feature *Feature) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(feature) {
		return
	}

	feature.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (firstfret *FirstFret) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(firstfret) {
		return
	}

	firstfret.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (frame *Frame) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(frame) {
		return
	}

	frame.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if frame.FirstFret != nil {
		stage.UnstageBranch(frame.FirstFret)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _framenote := range frame.FrameNote {
		stage.UnstageBranch(_framenote)
	}

}

func (framenote *FrameNote) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(framenote) {
		return
	}

	framenote.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if framenote.Barre != nil {
		stage.UnstageBranch(framenote.Barre)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (glass *Glass) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(glass) {
		return
	}

	glass.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (grouping *Grouping) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(grouping) {
		return
	}

	grouping.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _feature := range grouping.Feature {
		stage.UnstageBranch(_feature)
	}

}

func (harmony *Harmony) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(harmony) {
		return
	}

	harmony.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if harmony.Frame != nil {
		stage.UnstageBranch(harmony.Frame)
	}
	if harmony.Offset != nil {
		stage.UnstageBranch(harmony.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _root := range harmony.Root {
		stage.UnstageBranch(_root)
	}
	for _, _numeral := range harmony.Numeral {
		stage.UnstageBranch(_numeral)
	}
	for _, _styletext := range harmony.Function {
		stage.UnstageBranch(_styletext)
	}
	for _, _kind := range harmony.Kind {
		stage.UnstageBranch(_kind)
	}
	for _, _inversion := range harmony.Inversion {
		stage.UnstageBranch(_inversion)
	}
	for _, _bass := range harmony.Bass {
		stage.UnstageBranch(_bass)
	}
	for _, _degree := range harmony.Degree {
		stage.UnstageBranch(_degree)
	}

}

func (harmonyalter *HarmonyAlter) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(harmonyalter) {
		return
	}

	harmonyalter.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (harppedals *HarpPedals) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(harppedals) {
		return
	}

	harppedals.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _pedaltuning := range harppedals.PedalTuning {
		stage.UnstageBranch(_pedaltuning)
	}

}

func (image *Image) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(image) {
		return
	}

	image.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (instrumentchange *InstrumentChange) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(instrumentchange) {
		return
	}

	instrumentchange.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if instrumentchange.VirtualInstrument != nil {
		stage.UnstageBranch(instrumentchange.VirtualInstrument)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (inversion *Inversion) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(inversion) {
		return
	}

	inversion.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (kind *Kind) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(kind) {
		return
	}

	kind.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (listening *Listening) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(listening) {
		return
	}

	listening.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if listening.Offset != nil {
		stage.UnstageBranch(listening.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sync := range listening.Sync {
		stage.UnstageBranch(_sync)
	}
	for _, _otherlistening := range listening.OtherListening {
		stage.UnstageBranch(_otherlistening)
	}

}

func (measurenumbering *MeasureNumbering) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(measurenumbering) {
		return
	}

	measurenumbering.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (membrane *Membrane) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(membrane) {
		return
	}

	membrane.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metal *Metal) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(metal) {
		return
	}

	metal.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metronome *Metronome) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(metronome) {
		return
	}

	metronome.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if metronome.PerMinute != nil {
		stage.UnstageBranch(metronome.PerMinute)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _beatunittied := range metronome.BeatUnitTied {
		stage.UnstageBranch(_beatunittied)
	}
	for _, _metronomenote := range metronome.MetronomeNote {
		stage.UnstageBranch(_metronomenote)
	}

}

func (metronomebeam *MetronomeBeam) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(metronomebeam) {
		return
	}

	metronomebeam.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metronomenote *MetronomeNote) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(metronomenote) {
		return
	}

	metronomenote.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if metronomenote.MetronomeTied != nil {
		stage.UnstageBranch(metronomenote.MetronomeTied)
	}
	if metronomenote.MetronomeTuplet != nil {
		stage.UnstageBranch(metronomenote.MetronomeTuplet)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _metronomebeam := range metronomenote.MetronomeBeam {
		stage.UnstageBranch(_metronomebeam)
	}

}

func (metronometied *MetronomeTied) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(metronometied) {
		return
	}

	metronometied.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metronometuplet *MetronomeTuplet) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(metronometuplet) {
		return
	}

	metronometuplet.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (numeral *Numeral) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(numeral) {
		return
	}

	numeral.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if numeral.NumeralRoot != nil {
		stage.UnstageBranch(numeral.NumeralRoot)
	}
	if numeral.NumeralAlter != nil {
		stage.UnstageBranch(numeral.NumeralAlter)
	}
	if numeral.NumeralKey != nil {
		stage.UnstageBranch(numeral.NumeralKey)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (numeralkey *NumeralKey) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(numeralkey) {
		return
	}

	numeralkey.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (numeralroot *NumeralRoot) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(numeralroot) {
		return
	}

	numeralroot.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (octaveshift *OctaveShift) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(octaveshift) {
		return
	}

	octaveshift.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (offset *Offset) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(offset) {
		return
	}

	offset.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (otherdirection *OtherDirection) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(otherdirection) {
		return
	}

	otherdirection.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (otherlistening *OtherListening) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(otherlistening) {
		return
	}

	otherlistening.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pedal *Pedal) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(pedal) {
		return
	}

	pedal.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pedaltuning *PedalTuning) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(pedaltuning) {
		return
	}

	pedaltuning.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (perminute *PerMinute) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(perminute) {
		return
	}

	perminute.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (percussion *Percussion) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(percussion) {
		return
	}

	percussion.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if percussion.Glass != nil {
		stage.UnstageBranch(percussion.Glass)
	}
	if percussion.Metal != nil {
		stage.UnstageBranch(percussion.Metal)
	}
	if percussion.Wood != nil {
		stage.UnstageBranch(percussion.Wood)
	}
	if percussion.Pitched != nil {
		stage.UnstageBranch(percussion.Pitched)
	}
	if percussion.Membrane != nil {
		stage.UnstageBranch(percussion.Membrane)
	}
	if percussion.Effect != nil {
		stage.UnstageBranch(percussion.Effect)
	}
	if percussion.Timpani != nil {
		stage.UnstageBranch(percussion.Timpani)
	}
	if percussion.Beater != nil {
		stage.UnstageBranch(percussion.Beater)
	}
	if percussion.Stick != nil {
		stage.UnstageBranch(percussion.Stick)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pitched *Pitched) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(pitched) {
		return
	}

	pitched.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (principalvoice *PrincipalVoice) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(principalvoice) {
		return
	}

	principalvoice.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (print *Print) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(print) {
		return
	}

	print.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if print.MeasureNumbering != nil {
		stage.UnstageBranch(print.MeasureNumbering)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (root *Root) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(root) {
		return
	}

	root.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if root.RootStep != nil {
		stage.UnstageBranch(root.RootStep)
	}
	if root.RootAlter != nil {
		stage.UnstageBranch(root.RootAlter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rootstep *RootStep) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(rootstep) {
		return
	}

	rootstep.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (scordatura *Scordatura) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(scordatura) {
		return
	}

	scordatura.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _accord := range scordatura.Accord {
		stage.UnstageBranch(_accord)
	}

}

func (sound *Sound) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(sound) {
		return
	}

	sound.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if sound.Swing != nil {
		stage.UnstageBranch(sound.Swing)
	}
	if sound.Offset != nil {
		stage.UnstageBranch(sound.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrumentchange := range sound.InstrumentChange {
		stage.UnstageBranch(_instrumentchange)
	}

}

func (staffdivide *StaffDivide) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(staffdivide) {
		return
	}

	staffdivide.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stick *Stick) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stick) {
		return
	}

	stick.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stringmute *StringMute) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stringmute) {
		return
	}

	stringmute.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (styletext *StyleText) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(styletext) {
		return
	}

	styletext.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (swing *Swing) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(swing) {
		return
	}

	swing.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (sync *Sync) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(sync) {
		return
	}

	sync.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (timemodification *TimeModification) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(timemodification) {
		return
	}

	timemodification.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (timpani *Timpani) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(timpani) {
		return
	}

	timpani.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (virtualinstrument *VirtualInstrument) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(virtualinstrument) {
		return
	}

	virtualinstrument.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (wedge *Wedge) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(wedge) {
		return
	}

	wedge.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (wood *Wood) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(wood) {
		return
	}

	wood.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *Accord) GongReconstructPointersFromReferences(stage *Stage, instance *Accord) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *AccordionRegistration) GongReconstructPointersFromReferences(stage *Stage, instance *AccordionRegistration) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Barre) GongReconstructPointersFromReferences(stage *Stage, instance *Barre) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Bass) GongReconstructPointersFromReferences(stage *Stage, instance *Bass) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.BassSeparator, stage.StyleTexts_reference, instance.BassSeparator)
	__gong__reconstructPointer(&reference.BassStep, stage.BassSteps_reference, instance.BassStep)
	__gong__reconstructPointer(&reference.BassAlter, stage.HarmonyAlters_reference, instance.BassAlter)
	// insertion point for slice of pointers field
}

func (reference *BassStep) GongReconstructPointersFromReferences(stage *Stage, instance *BassStep) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *BeatUnitTied) GongReconstructPointersFromReferences(stage *Stage, instance *BeatUnitTied) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Beater) GongReconstructPointersFromReferences(stage *Stage, instance *Beater) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Bracket) GongReconstructPointersFromReferences(stage *Stage, instance *Bracket) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Dashes) GongReconstructPointersFromReferences(stage *Stage, instance *Dashes) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Degree) GongReconstructPointersFromReferences(stage *Stage, instance *Degree) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.DegreeValue, stage.DegreeValues_reference, instance.DegreeValue)
	__gong__reconstructPointer(&reference.DegreeAlter, stage.DegreeAlters_reference, instance.DegreeAlter)
	__gong__reconstructPointer(&reference.DegreeType, stage.DegreeTypes_reference, instance.DegreeType)
	// insertion point for slice of pointers field
}

func (reference *DegreeAlter) GongReconstructPointersFromReferences(stage *Stage, instance *DegreeAlter) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *DegreeType) GongReconstructPointersFromReferences(stage *Stage, instance *DegreeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *DegreeValue) GongReconstructPointersFromReferences(stage *Stage, instance *DegreeValue) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Direction) GongReconstructPointersFromReferences(stage *Stage, instance *Direction) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Offset, stage.Offsets_reference, instance.Offset)
	__gong__reconstructPointer(&reference.Sound, stage.Sounds_reference, instance.Sound)
	__gong__reconstructPointer(&reference.Listening, stage.Listenings_reference, instance.Listening)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.DirectionType, stage.DirectionTypes_reference, instance.DirectionType)
}

func (reference *DirectionType) GongReconstructPointersFromReferences(stage *Stage, instance *DirectionType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Wedge, stage.Wedges_reference, instance.Wedge)
	__gong__reconstructPointer(&reference.Dashes, stage.Dashess_reference, instance.Dashes)
	__gong__reconstructPointer(&reference.Bracket, stage.Brackets_reference, instance.Bracket)
	__gong__reconstructPointer(&reference.Pedal, stage.Pedals_reference, instance.Pedal)
	__gong__reconstructPointer(&reference.Metronome, stage.Metronomes_reference, instance.Metronome)
	__gong__reconstructPointer(&reference.OctaveShift, stage.OctaveShifts_reference, instance.OctaveShift)
	__gong__reconstructPointer(&reference.HarpPedals, stage.HarpPedalss_reference, instance.HarpPedals)
	__gong__reconstructPointer(&reference.StringMute, stage.StringMutes_reference, instance.StringMute)
	__gong__reconstructPointer(&reference.Scordatura, stage.Scordaturas_reference, instance.Scordatura)
	__gong__reconstructPointer(&reference.Image, stage.Images_reference, instance.Image)
	__gong__reconstructPointer(&reference.PrincipalVoice, stage.PrincipalVoices_reference, instance.PrincipalVoice)
	__gong__reconstructPointer(&reference.AccordionRegistration, stage.AccordionRegistrations_reference, instance.AccordionRegistration)
	__gong__reconstructPointer(&reference.StaffDivide, stage.StaffDivides_reference, instance.StaffDivide)
	__gong__reconstructPointer(&reference.OtherDirection, stage.OtherDirections_reference, instance.OtherDirection)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Percussion, stage.Percussions_reference, instance.Percussion)
}

func (reference *Effect) GongReconstructPointersFromReferences(stage *Stage, instance *Effect) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Feature) GongReconstructPointersFromReferences(stage *Stage, instance *Feature) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *FirstFret) GongReconstructPointersFromReferences(stage *Stage, instance *FirstFret) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Frame) GongReconstructPointersFromReferences(stage *Stage, instance *Frame) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.FirstFret, stage.FirstFrets_reference, instance.FirstFret)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.FrameNote, stage.FrameNotes_reference, instance.FrameNote)
}

func (reference *FrameNote) GongReconstructPointersFromReferences(stage *Stage, instance *FrameNote) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Barre, stage.Barres_reference, instance.Barre)
	// insertion point for slice of pointers field
}

func (reference *Glass) GongReconstructPointersFromReferences(stage *Stage, instance *Glass) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Grouping) GongReconstructPointersFromReferences(stage *Stage, instance *Grouping) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Feature, stage.Features_reference, instance.Feature)
}

func (reference *Harmony) GongReconstructPointersFromReferences(stage *Stage, instance *Harmony) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Frame, stage.Frames_reference, instance.Frame)
	__gong__reconstructPointer(&reference.Offset, stage.Offsets_reference, instance.Offset)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Root, stage.Roots_reference, instance.Root)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Numeral, stage.Numerals_reference, instance.Numeral)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Function, stage.StyleTexts_reference, instance.Function)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Kind, stage.Kinds_reference, instance.Kind)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Inversion, stage.Inversions_reference, instance.Inversion)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Bass, stage.Basss_reference, instance.Bass)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Degree, stage.Degrees_reference, instance.Degree)
}

func (reference *HarmonyAlter) GongReconstructPointersFromReferences(stage *Stage, instance *HarmonyAlter) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *HarpPedals) GongReconstructPointersFromReferences(stage *Stage, instance *HarpPedals) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.PedalTuning, stage.PedalTunings_reference, instance.PedalTuning)
}

func (reference *Image) GongReconstructPointersFromReferences(stage *Stage, instance *Image) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *InstrumentChange) GongReconstructPointersFromReferences(stage *Stage, instance *InstrumentChange) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.VirtualInstrument, stage.VirtualInstruments_reference, instance.VirtualInstrument)
	// insertion point for slice of pointers field
}

func (reference *Inversion) GongReconstructPointersFromReferences(stage *Stage, instance *Inversion) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Kind) GongReconstructPointersFromReferences(stage *Stage, instance *Kind) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Listening) GongReconstructPointersFromReferences(stage *Stage, instance *Listening) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Offset, stage.Offsets_reference, instance.Offset)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sync, stage.Syncs_reference, instance.Sync)
	__gong__reconstructSliceOfPointersFromReferences(&reference.OtherListening, stage.OtherListenings_reference, instance.OtherListening)
}

func (reference *MeasureNumbering) GongReconstructPointersFromReferences(stage *Stage, instance *MeasureNumbering) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Membrane) GongReconstructPointersFromReferences(stage *Stage, instance *Membrane) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Metal) GongReconstructPointersFromReferences(stage *Stage, instance *Metal) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Metronome) GongReconstructPointersFromReferences(stage *Stage, instance *Metronome) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.PerMinute, stage.PerMinutes_reference, instance.PerMinute)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.BeatUnitTied, stage.BeatUnitTieds_reference, instance.BeatUnitTied)
	__gong__reconstructSliceOfPointersFromReferences(&reference.MetronomeNote, stage.MetronomeNotes_reference, instance.MetronomeNote)
}

func (reference *MetronomeBeam) GongReconstructPointersFromReferences(stage *Stage, instance *MetronomeBeam) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MetronomeNote) GongReconstructPointersFromReferences(stage *Stage, instance *MetronomeNote) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MetronomeTied, stage.MetronomeTieds_reference, instance.MetronomeTied)
	__gong__reconstructPointer(&reference.MetronomeTuplet, stage.MetronomeTuplets_reference, instance.MetronomeTuplet)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.MetronomeBeam, stage.MetronomeBeams_reference, instance.MetronomeBeam)
}

func (reference *MetronomeTied) GongReconstructPointersFromReferences(stage *Stage, instance *MetronomeTied) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MetronomeTuplet) GongReconstructPointersFromReferences(stage *Stage, instance *MetronomeTuplet) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Numeral) GongReconstructPointersFromReferences(stage *Stage, instance *Numeral) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.NumeralRoot, stage.NumeralRoots_reference, instance.NumeralRoot)
	__gong__reconstructPointer(&reference.NumeralAlter, stage.HarmonyAlters_reference, instance.NumeralAlter)
	__gong__reconstructPointer(&reference.NumeralKey, stage.NumeralKeys_reference, instance.NumeralKey)
	// insertion point for slice of pointers field
}

func (reference *NumeralKey) GongReconstructPointersFromReferences(stage *Stage, instance *NumeralKey) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *NumeralRoot) GongReconstructPointersFromReferences(stage *Stage, instance *NumeralRoot) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *OctaveShift) GongReconstructPointersFromReferences(stage *Stage, instance *OctaveShift) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Offset) GongReconstructPointersFromReferences(stage *Stage, instance *Offset) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *OtherDirection) GongReconstructPointersFromReferences(stage *Stage, instance *OtherDirection) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *OtherListening) GongReconstructPointersFromReferences(stage *Stage, instance *OtherListening) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Pedal) GongReconstructPointersFromReferences(stage *Stage, instance *Pedal) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PedalTuning) GongReconstructPointersFromReferences(stage *Stage, instance *PedalTuning) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PerMinute) GongReconstructPointersFromReferences(stage *Stage, instance *PerMinute) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Percussion) GongReconstructPointersFromReferences(stage *Stage, instance *Percussion) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Glass, stage.Glasss_reference, instance.Glass)
	__gong__reconstructPointer(&reference.Metal, stage.Metals_reference, instance.Metal)
	__gong__reconstructPointer(&reference.Wood, stage.Woods_reference, instance.Wood)
	__gong__reconstructPointer(&reference.Pitched, stage.Pitcheds_reference, instance.Pitched)
	__gong__reconstructPointer(&reference.Membrane, stage.Membranes_reference, instance.Membrane)
	__gong__reconstructPointer(&reference.Effect, stage.Effects_reference, instance.Effect)
	__gong__reconstructPointer(&reference.Timpani, stage.Timpanis_reference, instance.Timpani)
	__gong__reconstructPointer(&reference.Beater, stage.Beaters_reference, instance.Beater)
	__gong__reconstructPointer(&reference.Stick, stage.Sticks_reference, instance.Stick)
	// insertion point for slice of pointers field
}

func (reference *Pitched) GongReconstructPointersFromReferences(stage *Stage, instance *Pitched) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PrincipalVoice) GongReconstructPointersFromReferences(stage *Stage, instance *PrincipalVoice) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Print) GongReconstructPointersFromReferences(stage *Stage, instance *Print) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MeasureNumbering, stage.MeasureNumberings_reference, instance.MeasureNumbering)
	// insertion point for slice of pointers field
}

func (reference *Root) GongReconstructPointersFromReferences(stage *Stage, instance *Root) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.RootStep, stage.RootSteps_reference, instance.RootStep)
	__gong__reconstructPointer(&reference.RootAlter, stage.HarmonyAlters_reference, instance.RootAlter)
	// insertion point for slice of pointers field
}

func (reference *RootStep) GongReconstructPointersFromReferences(stage *Stage, instance *RootStep) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Scordatura) GongReconstructPointersFromReferences(stage *Stage, instance *Scordatura) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Accord, stage.Accords_reference, instance.Accord)
}

func (reference *Sound) GongReconstructPointersFromReferences(stage *Stage, instance *Sound) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Swing, stage.Swings_reference, instance.Swing)
	__gong__reconstructPointer(&reference.Offset, stage.Offsets_reference, instance.Offset)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.InstrumentChange, stage.InstrumentChanges_reference, instance.InstrumentChange)
}

func (reference *StaffDivide) GongReconstructPointersFromReferences(stage *Stage, instance *StaffDivide) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Stick) GongReconstructPointersFromReferences(stage *Stage, instance *Stick) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StringMute) GongReconstructPointersFromReferences(stage *Stage, instance *StringMute) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StyleText) GongReconstructPointersFromReferences(stage *Stage, instance *StyleText) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Swing) GongReconstructPointersFromReferences(stage *Stage, instance *Swing) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Sync) GongReconstructPointersFromReferences(stage *Stage, instance *Sync) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TimeModification) GongReconstructPointersFromReferences(stage *Stage, instance *TimeModification) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Timpani) GongReconstructPointersFromReferences(stage *Stage, instance *Timpani) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *VirtualInstrument) GongReconstructPointersFromReferences(stage *Stage, instance *VirtualInstrument) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Wedge) GongReconstructPointersFromReferences(stage *Stage, instance *Wedge) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Wood) GongReconstructPointersFromReferences(stage *Stage, instance *Wood) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *Accord) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *AccordionRegistration) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Barre) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Bass) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.BassSeparator, stage.StyleTexts_instance)
	__gong__reconstructPointerFromInstance(&reference.BassStep, stage.BassSteps_instance)
	__gong__reconstructPointerFromInstance(&reference.BassAlter, stage.HarmonyAlters_instance)
	// insertion point for slice of pointers fields
}

func (reference *BassStep) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *BeatUnitTied) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Beater) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Bracket) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Dashes) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Degree) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.DegreeValue, stage.DegreeValues_instance)
	__gong__reconstructPointerFromInstance(&reference.DegreeAlter, stage.DegreeAlters_instance)
	__gong__reconstructPointerFromInstance(&reference.DegreeType, stage.DegreeTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *DegreeAlter) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *DegreeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *DegreeValue) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Direction) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Offset, stage.Offsets_instance)
	__gong__reconstructPointerFromInstance(&reference.Sound, stage.Sounds_instance)
	__gong__reconstructPointerFromInstance(&reference.Listening, stage.Listenings_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.DirectionType, stage.DirectionTypes_instance)
}

func (reference *DirectionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Wedge, stage.Wedges_instance)
	__gong__reconstructPointerFromInstance(&reference.Dashes, stage.Dashess_instance)
	__gong__reconstructPointerFromInstance(&reference.Bracket, stage.Brackets_instance)
	__gong__reconstructPointerFromInstance(&reference.Pedal, stage.Pedals_instance)
	__gong__reconstructPointerFromInstance(&reference.Metronome, stage.Metronomes_instance)
	__gong__reconstructPointerFromInstance(&reference.OctaveShift, stage.OctaveShifts_instance)
	__gong__reconstructPointerFromInstance(&reference.HarpPedals, stage.HarpPedalss_instance)
	__gong__reconstructPointerFromInstance(&reference.StringMute, stage.StringMutes_instance)
	__gong__reconstructPointerFromInstance(&reference.Scordatura, stage.Scordaturas_instance)
	__gong__reconstructPointerFromInstance(&reference.Image, stage.Images_instance)
	__gong__reconstructPointerFromInstance(&reference.PrincipalVoice, stage.PrincipalVoices_instance)
	__gong__reconstructPointerFromInstance(&reference.AccordionRegistration, stage.AccordionRegistrations_instance)
	__gong__reconstructPointerFromInstance(&reference.StaffDivide, stage.StaffDivides_instance)
	__gong__reconstructPointerFromInstance(&reference.OtherDirection, stage.OtherDirections_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Percussion, stage.Percussions_instance)
}

func (reference *Effect) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Feature) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *FirstFret) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Frame) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.FirstFret, stage.FirstFrets_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.FrameNote, stage.FrameNotes_instance)
}

func (reference *FrameNote) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Barre, stage.Barres_instance)
	// insertion point for slice of pointers fields
}

func (reference *Glass) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Grouping) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Feature, stage.Features_instance)
}

func (reference *Harmony) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Frame, stage.Frames_instance)
	__gong__reconstructPointerFromInstance(&reference.Offset, stage.Offsets_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Root, stage.Roots_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Numeral, stage.Numerals_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Function, stage.StyleTexts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Kind, stage.Kinds_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Inversion, stage.Inversions_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Bass, stage.Basss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Degree, stage.Degrees_instance)
}

func (reference *HarmonyAlter) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *HarpPedals) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.PedalTuning, stage.PedalTunings_instance)
}

func (reference *Image) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *InstrumentChange) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.VirtualInstrument, stage.VirtualInstruments_instance)
	// insertion point for slice of pointers fields
}

func (reference *Inversion) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Kind) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Listening) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Offset, stage.Offsets_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sync, stage.Syncs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.OtherListening, stage.OtherListenings_instance)
}

func (reference *MeasureNumbering) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Membrane) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Metal) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Metronome) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.PerMinute, stage.PerMinutes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.BeatUnitTied, stage.BeatUnitTieds_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.MetronomeNote, stage.MetronomeNotes_instance)
}

func (reference *MetronomeBeam) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MetronomeNote) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MetronomeTied, stage.MetronomeTieds_instance)
	__gong__reconstructPointerFromInstance(&reference.MetronomeTuplet, stage.MetronomeTuplets_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.MetronomeBeam, stage.MetronomeBeams_instance)
}

func (reference *MetronomeTied) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MetronomeTuplet) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Numeral) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.NumeralRoot, stage.NumeralRoots_instance)
	__gong__reconstructPointerFromInstance(&reference.NumeralAlter, stage.HarmonyAlters_instance)
	__gong__reconstructPointerFromInstance(&reference.NumeralKey, stage.NumeralKeys_instance)
	// insertion point for slice of pointers fields
}

func (reference *NumeralKey) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *NumeralRoot) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *OctaveShift) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Offset) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *OtherDirection) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *OtherListening) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Pedal) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PedalTuning) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PerMinute) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Percussion) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Glass, stage.Glasss_instance)
	__gong__reconstructPointerFromInstance(&reference.Metal, stage.Metals_instance)
	__gong__reconstructPointerFromInstance(&reference.Wood, stage.Woods_instance)
	__gong__reconstructPointerFromInstance(&reference.Pitched, stage.Pitcheds_instance)
	__gong__reconstructPointerFromInstance(&reference.Membrane, stage.Membranes_instance)
	__gong__reconstructPointerFromInstance(&reference.Effect, stage.Effects_instance)
	__gong__reconstructPointerFromInstance(&reference.Timpani, stage.Timpanis_instance)
	__gong__reconstructPointerFromInstance(&reference.Beater, stage.Beaters_instance)
	__gong__reconstructPointerFromInstance(&reference.Stick, stage.Sticks_instance)
	// insertion point for slice of pointers fields
}

func (reference *Pitched) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PrincipalVoice) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Print) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MeasureNumbering, stage.MeasureNumberings_instance)
	// insertion point for slice of pointers fields
}

func (reference *Root) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.RootStep, stage.RootSteps_instance)
	__gong__reconstructPointerFromInstance(&reference.RootAlter, stage.HarmonyAlters_instance)
	// insertion point for slice of pointers fields
}

func (reference *RootStep) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Scordatura) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Accord, stage.Accords_instance)
}

func (reference *Sound) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Swing, stage.Swings_instance)
	__gong__reconstructPointerFromInstance(&reference.Offset, stage.Offsets_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.InstrumentChange, stage.InstrumentChanges_instance)
}

func (reference *StaffDivide) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Stick) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StringMute) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StyleText) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Swing) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Sync) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TimeModification) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Timpani) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *VirtualInstrument) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Wedge) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Wood) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (accord *Accord) GongDiff(stage *Stage, accordOther *Accord) (diffs []string) {
	// insertion point for field diffs
	if accord.Name != accordOther.Name {
		diffs = append(diffs, accord.GongMarshallField(stage, "Name"))
	}
	if accord.String != accordOther.String {
		diffs = append(diffs, accord.GongMarshallField(stage, "String"))
	}
	if accord.TuningAlter != accordOther.TuningAlter {
		diffs = append(diffs, accord.GongMarshallField(stage, "TuningAlter"))
	}
	if accord.TuningOctave != accordOther.TuningOctave {
		diffs = append(diffs, accord.GongMarshallField(stage, "TuningOctave"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (accordionregistration *AccordionRegistration) GongDiff(stage *Stage, accordionregistrationOther *AccordionRegistration) (diffs []string) {
	// insertion point for field diffs
	if accordionregistration.Name != accordionregistrationOther.Name {
		diffs = append(diffs, accordionregistration.GongMarshallField(stage, "Name"))
	}
	if accordionregistration.DefaultX != accordionregistrationOther.DefaultX {
		diffs = append(diffs, accordionregistration.GongMarshallField(stage, "DefaultX"))
	}
	if accordionregistration.DefaultY != accordionregistrationOther.DefaultY {
		diffs = append(diffs, accordionregistration.GongMarshallField(stage, "DefaultY"))
	}
	if accordionregistration.RelativeX != accordionregistrationOther.RelativeX {
		diffs = append(diffs, accordionregistration.GongMarshallField(stage, "RelativeX"))
	}
	if accordionregistration.RelativeY != accordionregistrationOther.RelativeY {
		diffs = append(diffs, accordionregistration.GongMarshallField(stage, "RelativeY"))
	}
	if accordionregistration.FontFamily != accordionregistrationOther.FontFamily {
		diffs = append(diffs, accordionregistration.GongMarshallField(stage, "FontFamily"))
	}
	if accordionregistration.FontSize != accordionregistrationOther.FontSize {
		diffs = append(diffs, accordionregistration.GongMarshallField(stage, "FontSize"))
	}
	if accordionregistration.Color != accordionregistrationOther.Color {
		diffs = append(diffs, accordionregistration.GongMarshallField(stage, "Color"))
	}
	if accordionregistration.Id != accordionregistrationOther.Id {
		diffs = append(diffs, accordionregistration.GongMarshallField(stage, "Id"))
	}
	if accordionregistration.AccordionMiddle != accordionregistrationOther.AccordionMiddle {
		diffs = append(diffs, accordionregistration.GongMarshallField(stage, "AccordionMiddle"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (barre *Barre) GongDiff(stage *Stage, barreOther *Barre) (diffs []string) {
	// insertion point for field diffs
	if barre.Name != barreOther.Name {
		diffs = append(diffs, barre.GongMarshallField(stage, "Name"))
	}
	if barre.Color != barreOther.Color {
		diffs = append(diffs, barre.GongMarshallField(stage, "Color"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bass *Bass) GongDiff(stage *Stage, bassOther *Bass) (diffs []string) {
	// insertion point for field diffs
	if bass.Name != bassOther.Name {
		diffs = append(diffs, bass.GongMarshallField(stage, "Name"))
	}
	if bass.Arrangement != bassOther.Arrangement {
		diffs = append(diffs, bass.GongMarshallField(stage, "Arrangement"))
	}
	if bass.BassSeparator != bassOther.BassSeparator {
		diffs = append(diffs, bass.GongMarshallField(stage, "BassSeparator"))
	}
	if bass.BassStep != bassOther.BassStep {
		diffs = append(diffs, bass.GongMarshallField(stage, "BassStep"))
	}
	if bass.BassAlter != bassOther.BassAlter {
		diffs = append(diffs, bass.GongMarshallField(stage, "BassAlter"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bassstep *BassStep) GongDiff(stage *Stage, bassstepOther *BassStep) (diffs []string) {
	// insertion point for field diffs
	if bassstep.Name != bassstepOther.Name {
		diffs = append(diffs, bassstep.GongMarshallField(stage, "Name"))
	}
	if bassstep.Text != bassstepOther.Text {
		diffs = append(diffs, bassstep.GongMarshallField(stage, "Text"))
	}
	if bassstep.DefaultX != bassstepOther.DefaultX {
		diffs = append(diffs, bassstep.GongMarshallField(stage, "DefaultX"))
	}
	if bassstep.DefaultY != bassstepOther.DefaultY {
		diffs = append(diffs, bassstep.GongMarshallField(stage, "DefaultY"))
	}
	if bassstep.RelativeX != bassstepOther.RelativeX {
		diffs = append(diffs, bassstep.GongMarshallField(stage, "RelativeX"))
	}
	if bassstep.RelativeY != bassstepOther.RelativeY {
		diffs = append(diffs, bassstep.GongMarshallField(stage, "RelativeY"))
	}
	if bassstep.FontFamily != bassstepOther.FontFamily {
		diffs = append(diffs, bassstep.GongMarshallField(stage, "FontFamily"))
	}
	if bassstep.FontSize != bassstepOther.FontSize {
		diffs = append(diffs, bassstep.GongMarshallField(stage, "FontSize"))
	}
	if bassstep.Color != bassstepOther.Color {
		diffs = append(diffs, bassstep.GongMarshallField(stage, "Color"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (beatunittied *BeatUnitTied) GongDiff(stage *Stage, beatunittiedOther *BeatUnitTied) (diffs []string) {
	// insertion point for field diffs
	if beatunittied.Name != beatunittiedOther.Name {
		diffs = append(diffs, beatunittied.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (beater *Beater) GongDiff(stage *Stage, beaterOther *Beater) (diffs []string) {
	// insertion point for field diffs
	if beater.Name != beaterOther.Name {
		diffs = append(diffs, beater.GongMarshallField(stage, "Name"))
	}
	if beater.Tip != beaterOther.Tip {
		diffs = append(diffs, beater.GongMarshallField(stage, "Tip"))
	}
	if beater.Value != beaterOther.Value {
		diffs = append(diffs, beater.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bracket *Bracket) GongDiff(stage *Stage, bracketOther *Bracket) (diffs []string) {
	// insertion point for field diffs
	if bracket.Name != bracketOther.Name {
		diffs = append(diffs, bracket.GongMarshallField(stage, "Name"))
	}
	if bracket.Number != bracketOther.Number {
		diffs = append(diffs, bracket.GongMarshallField(stage, "Number"))
	}
	if bracket.LineEnd != bracketOther.LineEnd {
		diffs = append(diffs, bracket.GongMarshallField(stage, "LineEnd"))
	}
	if bracket.EndLength != bracketOther.EndLength {
		diffs = append(diffs, bracket.GongMarshallField(stage, "EndLength"))
	}
	if bracket.DashLength != bracketOther.DashLength {
		diffs = append(diffs, bracket.GongMarshallField(stage, "DashLength"))
	}
	if bracket.SpaceLength != bracketOther.SpaceLength {
		diffs = append(diffs, bracket.GongMarshallField(stage, "SpaceLength"))
	}
	if bracket.DefaultX != bracketOther.DefaultX {
		diffs = append(diffs, bracket.GongMarshallField(stage, "DefaultX"))
	}
	if bracket.DefaultY != bracketOther.DefaultY {
		diffs = append(diffs, bracket.GongMarshallField(stage, "DefaultY"))
	}
	if bracket.RelativeX != bracketOther.RelativeX {
		diffs = append(diffs, bracket.GongMarshallField(stage, "RelativeX"))
	}
	if bracket.RelativeY != bracketOther.RelativeY {
		diffs = append(diffs, bracket.GongMarshallField(stage, "RelativeY"))
	}
	if bracket.Color != bracketOther.Color {
		diffs = append(diffs, bracket.GongMarshallField(stage, "Color"))
	}
	if bracket.Id != bracketOther.Id {
		diffs = append(diffs, bracket.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (dashes *Dashes) GongDiff(stage *Stage, dashesOther *Dashes) (diffs []string) {
	// insertion point for field diffs
	if dashes.Name != dashesOther.Name {
		diffs = append(diffs, dashes.GongMarshallField(stage, "Name"))
	}
	if dashes.Number != dashesOther.Number {
		diffs = append(diffs, dashes.GongMarshallField(stage, "Number"))
	}
	if dashes.DashLength != dashesOther.DashLength {
		diffs = append(diffs, dashes.GongMarshallField(stage, "DashLength"))
	}
	if dashes.SpaceLength != dashesOther.SpaceLength {
		diffs = append(diffs, dashes.GongMarshallField(stage, "SpaceLength"))
	}
	if dashes.DefaultX != dashesOther.DefaultX {
		diffs = append(diffs, dashes.GongMarshallField(stage, "DefaultX"))
	}
	if dashes.DefaultY != dashesOther.DefaultY {
		diffs = append(diffs, dashes.GongMarshallField(stage, "DefaultY"))
	}
	if dashes.RelativeX != dashesOther.RelativeX {
		diffs = append(diffs, dashes.GongMarshallField(stage, "RelativeX"))
	}
	if dashes.RelativeY != dashesOther.RelativeY {
		diffs = append(diffs, dashes.GongMarshallField(stage, "RelativeY"))
	}
	if dashes.Color != dashesOther.Color {
		diffs = append(diffs, dashes.GongMarshallField(stage, "Color"))
	}
	if dashes.Id != dashesOther.Id {
		diffs = append(diffs, dashes.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (degree *Degree) GongDiff(stage *Stage, degreeOther *Degree) (diffs []string) {
	// insertion point for field diffs
	if degree.Name != degreeOther.Name {
		diffs = append(diffs, degree.GongMarshallField(stage, "Name"))
	}
	if degree.DegreeValue != degreeOther.DegreeValue {
		diffs = append(diffs, degree.GongMarshallField(stage, "DegreeValue"))
	}
	if degree.DegreeAlter != degreeOther.DegreeAlter {
		diffs = append(diffs, degree.GongMarshallField(stage, "DegreeAlter"))
	}
	if degree.DegreeType != degreeOther.DegreeType {
		diffs = append(diffs, degree.GongMarshallField(stage, "DegreeType"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (degreealter *DegreeAlter) GongDiff(stage *Stage, degreealterOther *DegreeAlter) (diffs []string) {
	// insertion point for field diffs
	if degreealter.Name != degreealterOther.Name {
		diffs = append(diffs, degreealter.GongMarshallField(stage, "Name"))
	}
	if degreealter.DefaultX != degreealterOther.DefaultX {
		diffs = append(diffs, degreealter.GongMarshallField(stage, "DefaultX"))
	}
	if degreealter.DefaultY != degreealterOther.DefaultY {
		diffs = append(diffs, degreealter.GongMarshallField(stage, "DefaultY"))
	}
	if degreealter.RelativeX != degreealterOther.RelativeX {
		diffs = append(diffs, degreealter.GongMarshallField(stage, "RelativeX"))
	}
	if degreealter.RelativeY != degreealterOther.RelativeY {
		diffs = append(diffs, degreealter.GongMarshallField(stage, "RelativeY"))
	}
	if degreealter.FontFamily != degreealterOther.FontFamily {
		diffs = append(diffs, degreealter.GongMarshallField(stage, "FontFamily"))
	}
	if degreealter.FontSize != degreealterOther.FontSize {
		diffs = append(diffs, degreealter.GongMarshallField(stage, "FontSize"))
	}
	if degreealter.Color != degreealterOther.Color {
		diffs = append(diffs, degreealter.GongMarshallField(stage, "Color"))
	}
	if degreealter.Value != degreealterOther.Value {
		diffs = append(diffs, degreealter.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (degreetype *DegreeType) GongDiff(stage *Stage, degreetypeOther *DegreeType) (diffs []string) {
	// insertion point for field diffs
	if degreetype.Name != degreetypeOther.Name {
		diffs = append(diffs, degreetype.GongMarshallField(stage, "Name"))
	}
	if degreetype.Text != degreetypeOther.Text {
		diffs = append(diffs, degreetype.GongMarshallField(stage, "Text"))
	}
	if degreetype.DefaultX != degreetypeOther.DefaultX {
		diffs = append(diffs, degreetype.GongMarshallField(stage, "DefaultX"))
	}
	if degreetype.DefaultY != degreetypeOther.DefaultY {
		diffs = append(diffs, degreetype.GongMarshallField(stage, "DefaultY"))
	}
	if degreetype.RelativeX != degreetypeOther.RelativeX {
		diffs = append(diffs, degreetype.GongMarshallField(stage, "RelativeX"))
	}
	if degreetype.RelativeY != degreetypeOther.RelativeY {
		diffs = append(diffs, degreetype.GongMarshallField(stage, "RelativeY"))
	}
	if degreetype.FontFamily != degreetypeOther.FontFamily {
		diffs = append(diffs, degreetype.GongMarshallField(stage, "FontFamily"))
	}
	if degreetype.FontSize != degreetypeOther.FontSize {
		diffs = append(diffs, degreetype.GongMarshallField(stage, "FontSize"))
	}
	if degreetype.Color != degreetypeOther.Color {
		diffs = append(diffs, degreetype.GongMarshallField(stage, "Color"))
	}
	if degreetype.Value != degreetypeOther.Value {
		diffs = append(diffs, degreetype.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (degreevalue *DegreeValue) GongDiff(stage *Stage, degreevalueOther *DegreeValue) (diffs []string) {
	// insertion point for field diffs
	if degreevalue.Name != degreevalueOther.Name {
		diffs = append(diffs, degreevalue.GongMarshallField(stage, "Name"))
	}
	if degreevalue.Symbol != degreevalueOther.Symbol {
		diffs = append(diffs, degreevalue.GongMarshallField(stage, "Symbol"))
	}
	if degreevalue.Text != degreevalueOther.Text {
		diffs = append(diffs, degreevalue.GongMarshallField(stage, "Text"))
	}
	if degreevalue.DefaultX != degreevalueOther.DefaultX {
		diffs = append(diffs, degreevalue.GongMarshallField(stage, "DefaultX"))
	}
	if degreevalue.DefaultY != degreevalueOther.DefaultY {
		diffs = append(diffs, degreevalue.GongMarshallField(stage, "DefaultY"))
	}
	if degreevalue.RelativeX != degreevalueOther.RelativeX {
		diffs = append(diffs, degreevalue.GongMarshallField(stage, "RelativeX"))
	}
	if degreevalue.RelativeY != degreevalueOther.RelativeY {
		diffs = append(diffs, degreevalue.GongMarshallField(stage, "RelativeY"))
	}
	if degreevalue.FontFamily != degreevalueOther.FontFamily {
		diffs = append(diffs, degreevalue.GongMarshallField(stage, "FontFamily"))
	}
	if degreevalue.FontSize != degreevalueOther.FontSize {
		diffs = append(diffs, degreevalue.GongMarshallField(stage, "FontSize"))
	}
	if degreevalue.Color != degreevalueOther.Color {
		diffs = append(diffs, degreevalue.GongMarshallField(stage, "Color"))
	}
	if degreevalue.Value != degreevalueOther.Value {
		diffs = append(diffs, degreevalue.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (direction *Direction) GongDiff(stage *Stage, directionOther *Direction) (diffs []string) {
	// insertion point for field diffs
	if direction.Name != directionOther.Name {
		diffs = append(diffs, direction.GongMarshallField(stage, "Name"))
	}
	if direction.Id != directionOther.Id {
		diffs = append(diffs, direction.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, direction, "DirectionType", directionOther.DirectionType, direction.DirectionType); ops != "" {
		diffs = append(diffs, ops)
	}
	if direction.Offset != directionOther.Offset {
		diffs = append(diffs, direction.GongMarshallField(stage, "Offset"))
	}
	if direction.Voice != directionOther.Voice {
		diffs = append(diffs, direction.GongMarshallField(stage, "Voice"))
	}
	if direction.Staff != directionOther.Staff {
		diffs = append(diffs, direction.GongMarshallField(stage, "Staff"))
	}
	if direction.Sound != directionOther.Sound {
		diffs = append(diffs, direction.GongMarshallField(stage, "Sound"))
	}
	if direction.Listening != directionOther.Listening {
		diffs = append(diffs, direction.GongMarshallField(stage, "Listening"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (directiontype *DirectionType) GongDiff(stage *Stage, directiontypeOther *DirectionType) (diffs []string) {
	// insertion point for field diffs
	if directiontype.Name != directiontypeOther.Name {
		diffs = append(diffs, directiontype.GongMarshallField(stage, "Name"))
	}
	if directiontype.Id != directiontypeOther.Id {
		diffs = append(diffs, directiontype.GongMarshallField(stage, "Id"))
	}
	if directiontype.Wedge != directiontypeOther.Wedge {
		diffs = append(diffs, directiontype.GongMarshallField(stage, "Wedge"))
	}
	if directiontype.Dashes != directiontypeOther.Dashes {
		diffs = append(diffs, directiontype.GongMarshallField(stage, "Dashes"))
	}
	if directiontype.Bracket != directiontypeOther.Bracket {
		diffs = append(diffs, directiontype.GongMarshallField(stage, "Bracket"))
	}
	if directiontype.Pedal != directiontypeOther.Pedal {
		diffs = append(diffs, directiontype.GongMarshallField(stage, "Pedal"))
	}
	if directiontype.Metronome != directiontypeOther.Metronome {
		diffs = append(diffs, directiontype.GongMarshallField(stage, "Metronome"))
	}
	if directiontype.OctaveShift != directiontypeOther.OctaveShift {
		diffs = append(diffs, directiontype.GongMarshallField(stage, "OctaveShift"))
	}
	if directiontype.HarpPedals != directiontypeOther.HarpPedals {
		diffs = append(diffs, directiontype.GongMarshallField(stage, "HarpPedals"))
	}
	if directiontype.StringMute != directiontypeOther.StringMute {
		diffs = append(diffs, directiontype.GongMarshallField(stage, "StringMute"))
	}
	if directiontype.Scordatura != directiontypeOther.Scordatura {
		diffs = append(diffs, directiontype.GongMarshallField(stage, "Scordatura"))
	}
	if directiontype.Image != directiontypeOther.Image {
		diffs = append(diffs, directiontype.GongMarshallField(stage, "Image"))
	}
	if directiontype.PrincipalVoice != directiontypeOther.PrincipalVoice {
		diffs = append(diffs, directiontype.GongMarshallField(stage, "PrincipalVoice"))
	}
	if ops := __gong__diffSliceOfPointers(stage, directiontype, "Percussion", directiontypeOther.Percussion, directiontype.Percussion); ops != "" {
		diffs = append(diffs, ops)
	}
	if directiontype.AccordionRegistration != directiontypeOther.AccordionRegistration {
		diffs = append(diffs, directiontype.GongMarshallField(stage, "AccordionRegistration"))
	}
	if directiontype.StaffDivide != directiontypeOther.StaffDivide {
		diffs = append(diffs, directiontype.GongMarshallField(stage, "StaffDivide"))
	}
	if directiontype.OtherDirection != directiontypeOther.OtherDirection {
		diffs = append(diffs, directiontype.GongMarshallField(stage, "OtherDirection"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (effect *Effect) GongDiff(stage *Stage, effectOther *Effect) (diffs []string) {
	// insertion point for field diffs
	if effect.Name != effectOther.Name {
		diffs = append(diffs, effect.GongMarshallField(stage, "Name"))
	}
	if effect.Smufl != effectOther.Smufl {
		diffs = append(diffs, effect.GongMarshallField(stage, "Smufl"))
	}
	if effect.Value != effectOther.Value {
		diffs = append(diffs, effect.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (feature *Feature) GongDiff(stage *Stage, featureOther *Feature) (diffs []string) {
	// insertion point for field diffs
	if feature.Name != featureOther.Name {
		diffs = append(diffs, feature.GongMarshallField(stage, "Name"))
	}
	if feature.Type != featureOther.Type {
		diffs = append(diffs, feature.GongMarshallField(stage, "Type"))
	}
	if feature.Value != featureOther.Value {
		diffs = append(diffs, feature.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (firstfret *FirstFret) GongDiff(stage *Stage, firstfretOther *FirstFret) (diffs []string) {
	// insertion point for field diffs
	if firstfret.Name != firstfretOther.Name {
		diffs = append(diffs, firstfret.GongMarshallField(stage, "Name"))
	}
	if firstfret.Text != firstfretOther.Text {
		diffs = append(diffs, firstfret.GongMarshallField(stage, "Text"))
	}
	if firstfret.Value != firstfretOther.Value {
		diffs = append(diffs, firstfret.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (frame *Frame) GongDiff(stage *Stage, frameOther *Frame) (diffs []string) {
	// insertion point for field diffs
	if frame.Name != frameOther.Name {
		diffs = append(diffs, frame.GongMarshallField(stage, "Name"))
	}
	if frame.Height != frameOther.Height {
		diffs = append(diffs, frame.GongMarshallField(stage, "Height"))
	}
	if frame.Width != frameOther.Width {
		diffs = append(diffs, frame.GongMarshallField(stage, "Width"))
	}
	if frame.Unplayed != frameOther.Unplayed {
		diffs = append(diffs, frame.GongMarshallField(stage, "Unplayed"))
	}
	if frame.DefaultX != frameOther.DefaultX {
		diffs = append(diffs, frame.GongMarshallField(stage, "DefaultX"))
	}
	if frame.DefaultY != frameOther.DefaultY {
		diffs = append(diffs, frame.GongMarshallField(stage, "DefaultY"))
	}
	if frame.RelativeX != frameOther.RelativeX {
		diffs = append(diffs, frame.GongMarshallField(stage, "RelativeX"))
	}
	if frame.RelativeY != frameOther.RelativeY {
		diffs = append(diffs, frame.GongMarshallField(stage, "RelativeY"))
	}
	if frame.Color != frameOther.Color {
		diffs = append(diffs, frame.GongMarshallField(stage, "Color"))
	}
	if frame.Id != frameOther.Id {
		diffs = append(diffs, frame.GongMarshallField(stage, "Id"))
	}
	if frame.FrameStrings != frameOther.FrameStrings {
		diffs = append(diffs, frame.GongMarshallField(stage, "FrameStrings"))
	}
	if frame.FrameFrets != frameOther.FrameFrets {
		diffs = append(diffs, frame.GongMarshallField(stage, "FrameFrets"))
	}
	if frame.FirstFret != frameOther.FirstFret {
		diffs = append(diffs, frame.GongMarshallField(stage, "FirstFret"))
	}
	if ops := __gong__diffSliceOfPointers(stage, frame, "FrameNote", frameOther.FrameNote, frame.FrameNote); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (framenote *FrameNote) GongDiff(stage *Stage, framenoteOther *FrameNote) (diffs []string) {
	// insertion point for field diffs
	if framenote.Name != framenoteOther.Name {
		diffs = append(diffs, framenote.GongMarshallField(stage, "Name"))
	}
	if framenote.Barre != framenoteOther.Barre {
		diffs = append(diffs, framenote.GongMarshallField(stage, "Barre"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (glass *Glass) GongDiff(stage *Stage, glassOther *Glass) (diffs []string) {
	// insertion point for field diffs
	if glass.Name != glassOther.Name {
		diffs = append(diffs, glass.GongMarshallField(stage, "Name"))
	}
	if glass.Smufl != glassOther.Smufl {
		diffs = append(diffs, glass.GongMarshallField(stage, "Smufl"))
	}
	if glass.Value != glassOther.Value {
		diffs = append(diffs, glass.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (grouping *Grouping) GongDiff(stage *Stage, groupingOther *Grouping) (diffs []string) {
	// insertion point for field diffs
	if grouping.Name != groupingOther.Name {
		diffs = append(diffs, grouping.GongMarshallField(stage, "Name"))
	}
	if grouping.Number != groupingOther.Number {
		diffs = append(diffs, grouping.GongMarshallField(stage, "Number"))
	}
	if grouping.MemberOf != groupingOther.MemberOf {
		diffs = append(diffs, grouping.GongMarshallField(stage, "MemberOf"))
	}
	if grouping.Id != groupingOther.Id {
		diffs = append(diffs, grouping.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, grouping, "Feature", groupingOther.Feature, grouping.Feature); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (harmony *Harmony) GongDiff(stage *Stage, harmonyOther *Harmony) (diffs []string) {
	// insertion point for field diffs
	if harmony.Name != harmonyOther.Name {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Name"))
	}
	if harmony.Type != harmonyOther.Type {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Type"))
	}
	if harmony.Arrangement != harmonyOther.Arrangement {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Arrangement"))
	}
	if harmony.DefaultX != harmonyOther.DefaultX {
		diffs = append(diffs, harmony.GongMarshallField(stage, "DefaultX"))
	}
	if harmony.DefaultY != harmonyOther.DefaultY {
		diffs = append(diffs, harmony.GongMarshallField(stage, "DefaultY"))
	}
	if harmony.RelativeX != harmonyOther.RelativeX {
		diffs = append(diffs, harmony.GongMarshallField(stage, "RelativeX"))
	}
	if harmony.RelativeY != harmonyOther.RelativeY {
		diffs = append(diffs, harmony.GongMarshallField(stage, "RelativeY"))
	}
	if harmony.FontFamily != harmonyOther.FontFamily {
		diffs = append(diffs, harmony.GongMarshallField(stage, "FontFamily"))
	}
	if harmony.FontSize != harmonyOther.FontSize {
		diffs = append(diffs, harmony.GongMarshallField(stage, "FontSize"))
	}
	if harmony.Color != harmonyOther.Color {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Color"))
	}
	if harmony.Id != harmonyOther.Id {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, harmony, "Root", harmonyOther.Root, harmony.Root); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, harmony, "Numeral", harmonyOther.Numeral, harmony.Numeral); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, harmony, "Function", harmonyOther.Function, harmony.Function); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, harmony, "Kind", harmonyOther.Kind, harmony.Kind); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, harmony, "Inversion", harmonyOther.Inversion, harmony.Inversion); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, harmony, "Bass", harmonyOther.Bass, harmony.Bass); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, harmony, "Degree", harmonyOther.Degree, harmony.Degree); ops != "" {
		diffs = append(diffs, ops)
	}
	if harmony.Frame != harmonyOther.Frame {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Frame"))
	}
	if harmony.Offset != harmonyOther.Offset {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Offset"))
	}
	if harmony.Staff != harmonyOther.Staff {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Staff"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (harmonyalter *HarmonyAlter) GongDiff(stage *Stage, harmonyalterOther *HarmonyAlter) (diffs []string) {
	// insertion point for field diffs
	if harmonyalter.Name != harmonyalterOther.Name {
		diffs = append(diffs, harmonyalter.GongMarshallField(stage, "Name"))
	}
	if harmonyalter.DefaultX != harmonyalterOther.DefaultX {
		diffs = append(diffs, harmonyalter.GongMarshallField(stage, "DefaultX"))
	}
	if harmonyalter.DefaultY != harmonyalterOther.DefaultY {
		diffs = append(diffs, harmonyalter.GongMarshallField(stage, "DefaultY"))
	}
	if harmonyalter.RelativeX != harmonyalterOther.RelativeX {
		diffs = append(diffs, harmonyalter.GongMarshallField(stage, "RelativeX"))
	}
	if harmonyalter.RelativeY != harmonyalterOther.RelativeY {
		diffs = append(diffs, harmonyalter.GongMarshallField(stage, "RelativeY"))
	}
	if harmonyalter.FontFamily != harmonyalterOther.FontFamily {
		diffs = append(diffs, harmonyalter.GongMarshallField(stage, "FontFamily"))
	}
	if harmonyalter.FontSize != harmonyalterOther.FontSize {
		diffs = append(diffs, harmonyalter.GongMarshallField(stage, "FontSize"))
	}
	if harmonyalter.Color != harmonyalterOther.Color {
		diffs = append(diffs, harmonyalter.GongMarshallField(stage, "Color"))
	}
	if harmonyalter.Value != harmonyalterOther.Value {
		diffs = append(diffs, harmonyalter.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (harppedals *HarpPedals) GongDiff(stage *Stage, harppedalsOther *HarpPedals) (diffs []string) {
	// insertion point for field diffs
	if harppedals.Name != harppedalsOther.Name {
		diffs = append(diffs, harppedals.GongMarshallField(stage, "Name"))
	}
	if harppedals.DefaultX != harppedalsOther.DefaultX {
		diffs = append(diffs, harppedals.GongMarshallField(stage, "DefaultX"))
	}
	if harppedals.DefaultY != harppedalsOther.DefaultY {
		diffs = append(diffs, harppedals.GongMarshallField(stage, "DefaultY"))
	}
	if harppedals.RelativeX != harppedalsOther.RelativeX {
		diffs = append(diffs, harppedals.GongMarshallField(stage, "RelativeX"))
	}
	if harppedals.RelativeY != harppedalsOther.RelativeY {
		diffs = append(diffs, harppedals.GongMarshallField(stage, "RelativeY"))
	}
	if harppedals.FontFamily != harppedalsOther.FontFamily {
		diffs = append(diffs, harppedals.GongMarshallField(stage, "FontFamily"))
	}
	if harppedals.FontSize != harppedalsOther.FontSize {
		diffs = append(diffs, harppedals.GongMarshallField(stage, "FontSize"))
	}
	if harppedals.Color != harppedalsOther.Color {
		diffs = append(diffs, harppedals.GongMarshallField(stage, "Color"))
	}
	if harppedals.Id != harppedalsOther.Id {
		diffs = append(diffs, harppedals.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, harppedals, "PedalTuning", harppedalsOther.PedalTuning, harppedals.PedalTuning); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (image *Image) GongDiff(stage *Stage, imageOther *Image) (diffs []string) {
	// insertion point for field diffs
	if image.Name != imageOther.Name {
		diffs = append(diffs, image.GongMarshallField(stage, "Name"))
	}
	if image.Source != imageOther.Source {
		diffs = append(diffs, image.GongMarshallField(stage, "Source"))
	}
	if image.Type != imageOther.Type {
		diffs = append(diffs, image.GongMarshallField(stage, "Type"))
	}
	if image.Height != imageOther.Height {
		diffs = append(diffs, image.GongMarshallField(stage, "Height"))
	}
	if image.Width != imageOther.Width {
		diffs = append(diffs, image.GongMarshallField(stage, "Width"))
	}
	if image.DefaultX != imageOther.DefaultX {
		diffs = append(diffs, image.GongMarshallField(stage, "DefaultX"))
	}
	if image.DefaultY != imageOther.DefaultY {
		diffs = append(diffs, image.GongMarshallField(stage, "DefaultY"))
	}
	if image.RelativeX != imageOther.RelativeX {
		diffs = append(diffs, image.GongMarshallField(stage, "RelativeX"))
	}
	if image.RelativeY != imageOther.RelativeY {
		diffs = append(diffs, image.GongMarshallField(stage, "RelativeY"))
	}
	if image.Id != imageOther.Id {
		diffs = append(diffs, image.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (instrumentchange *InstrumentChange) GongDiff(stage *Stage, instrumentchangeOther *InstrumentChange) (diffs []string) {
	// insertion point for field diffs
	if instrumentchange.Name != instrumentchangeOther.Name {
		diffs = append(diffs, instrumentchange.GongMarshallField(stage, "Name"))
	}
	if instrumentchange.Id != instrumentchangeOther.Id {
		diffs = append(diffs, instrumentchange.GongMarshallField(stage, "Id"))
	}
	if instrumentchange.InstrumentSound != instrumentchangeOther.InstrumentSound {
		diffs = append(diffs, instrumentchange.GongMarshallField(stage, "InstrumentSound"))
	}
	if instrumentchange.Ensemble != instrumentchangeOther.Ensemble {
		diffs = append(diffs, instrumentchange.GongMarshallField(stage, "Ensemble"))
	}
	if instrumentchange.VirtualInstrument != instrumentchangeOther.VirtualInstrument {
		diffs = append(diffs, instrumentchange.GongMarshallField(stage, "VirtualInstrument"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (inversion *Inversion) GongDiff(stage *Stage, inversionOther *Inversion) (diffs []string) {
	// insertion point for field diffs
	if inversion.Name != inversionOther.Name {
		diffs = append(diffs, inversion.GongMarshallField(stage, "Name"))
	}
	if inversion.Text != inversionOther.Text {
		diffs = append(diffs, inversion.GongMarshallField(stage, "Text"))
	}
	if inversion.DefaultX != inversionOther.DefaultX {
		diffs = append(diffs, inversion.GongMarshallField(stage, "DefaultX"))
	}
	if inversion.DefaultY != inversionOther.DefaultY {
		diffs = append(diffs, inversion.GongMarshallField(stage, "DefaultY"))
	}
	if inversion.RelativeX != inversionOther.RelativeX {
		diffs = append(diffs, inversion.GongMarshallField(stage, "RelativeX"))
	}
	if inversion.RelativeY != inversionOther.RelativeY {
		diffs = append(diffs, inversion.GongMarshallField(stage, "RelativeY"))
	}
	if inversion.FontFamily != inversionOther.FontFamily {
		diffs = append(diffs, inversion.GongMarshallField(stage, "FontFamily"))
	}
	if inversion.FontSize != inversionOther.FontSize {
		diffs = append(diffs, inversion.GongMarshallField(stage, "FontSize"))
	}
	if inversion.Color != inversionOther.Color {
		diffs = append(diffs, inversion.GongMarshallField(stage, "Color"))
	}
	if inversion.Value != inversionOther.Value {
		diffs = append(diffs, inversion.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (kind *Kind) GongDiff(stage *Stage, kindOther *Kind) (diffs []string) {
	// insertion point for field diffs
	if kind.Name != kindOther.Name {
		diffs = append(diffs, kind.GongMarshallField(stage, "Name"))
	}
	if kind.Text != kindOther.Text {
		diffs = append(diffs, kind.GongMarshallField(stage, "Text"))
	}
	if kind.DefaultX != kindOther.DefaultX {
		diffs = append(diffs, kind.GongMarshallField(stage, "DefaultX"))
	}
	if kind.DefaultY != kindOther.DefaultY {
		diffs = append(diffs, kind.GongMarshallField(stage, "DefaultY"))
	}
	if kind.RelativeX != kindOther.RelativeX {
		diffs = append(diffs, kind.GongMarshallField(stage, "RelativeX"))
	}
	if kind.RelativeY != kindOther.RelativeY {
		diffs = append(diffs, kind.GongMarshallField(stage, "RelativeY"))
	}
	if kind.FontFamily != kindOther.FontFamily {
		diffs = append(diffs, kind.GongMarshallField(stage, "FontFamily"))
	}
	if kind.FontSize != kindOther.FontSize {
		diffs = append(diffs, kind.GongMarshallField(stage, "FontSize"))
	}
	if kind.Color != kindOther.Color {
		diffs = append(diffs, kind.GongMarshallField(stage, "Color"))
	}
	if kind.Value != kindOther.Value {
		diffs = append(diffs, kind.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (listening *Listening) GongDiff(stage *Stage, listeningOther *Listening) (diffs []string) {
	// insertion point for field diffs
	if listening.Name != listeningOther.Name {
		diffs = append(diffs, listening.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, listening, "Sync", listeningOther.Sync, listening.Sync); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, listening, "OtherListening", listeningOther.OtherListening, listening.OtherListening); ops != "" {
		diffs = append(diffs, ops)
	}
	if listening.Offset != listeningOther.Offset {
		diffs = append(diffs, listening.GongMarshallField(stage, "Offset"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (measurenumbering *MeasureNumbering) GongDiff(stage *Stage, measurenumberingOther *MeasureNumbering) (diffs []string) {
	// insertion point for field diffs
	if measurenumbering.Name != measurenumberingOther.Name {
		diffs = append(diffs, measurenumbering.GongMarshallField(stage, "Name"))
	}
	if measurenumbering.Staff != measurenumberingOther.Staff {
		diffs = append(diffs, measurenumbering.GongMarshallField(stage, "Staff"))
	}
	if measurenumbering.DefaultX != measurenumberingOther.DefaultX {
		diffs = append(diffs, measurenumbering.GongMarshallField(stage, "DefaultX"))
	}
	if measurenumbering.DefaultY != measurenumberingOther.DefaultY {
		diffs = append(diffs, measurenumbering.GongMarshallField(stage, "DefaultY"))
	}
	if measurenumbering.RelativeX != measurenumberingOther.RelativeX {
		diffs = append(diffs, measurenumbering.GongMarshallField(stage, "RelativeX"))
	}
	if measurenumbering.RelativeY != measurenumberingOther.RelativeY {
		diffs = append(diffs, measurenumbering.GongMarshallField(stage, "RelativeY"))
	}
	if measurenumbering.FontFamily != measurenumberingOther.FontFamily {
		diffs = append(diffs, measurenumbering.GongMarshallField(stage, "FontFamily"))
	}
	if measurenumbering.FontSize != measurenumberingOther.FontSize {
		diffs = append(diffs, measurenumbering.GongMarshallField(stage, "FontSize"))
	}
	if measurenumbering.Color != measurenumberingOther.Color {
		diffs = append(diffs, measurenumbering.GongMarshallField(stage, "Color"))
	}
	if measurenumbering.Value != measurenumberingOther.Value {
		diffs = append(diffs, measurenumbering.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (membrane *Membrane) GongDiff(stage *Stage, membraneOther *Membrane) (diffs []string) {
	// insertion point for field diffs
	if membrane.Name != membraneOther.Name {
		diffs = append(diffs, membrane.GongMarshallField(stage, "Name"))
	}
	if membrane.Smufl != membraneOther.Smufl {
		diffs = append(diffs, membrane.GongMarshallField(stage, "Smufl"))
	}
	if membrane.Value != membraneOther.Value {
		diffs = append(diffs, membrane.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (metal *Metal) GongDiff(stage *Stage, metalOther *Metal) (diffs []string) {
	// insertion point for field diffs
	if metal.Name != metalOther.Name {
		diffs = append(diffs, metal.GongMarshallField(stage, "Name"))
	}
	if metal.Smufl != metalOther.Smufl {
		diffs = append(diffs, metal.GongMarshallField(stage, "Smufl"))
	}
	if metal.Value != metalOther.Value {
		diffs = append(diffs, metal.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (metronome *Metronome) GongDiff(stage *Stage, metronomeOther *Metronome) (diffs []string) {
	// insertion point for field diffs
	if metronome.Name != metronomeOther.Name {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Name"))
	}
	if metronome.DefaultX != metronomeOther.DefaultX {
		diffs = append(diffs, metronome.GongMarshallField(stage, "DefaultX"))
	}
	if metronome.DefaultY != metronomeOther.DefaultY {
		diffs = append(diffs, metronome.GongMarshallField(stage, "DefaultY"))
	}
	if metronome.RelativeX != metronomeOther.RelativeX {
		diffs = append(diffs, metronome.GongMarshallField(stage, "RelativeX"))
	}
	if metronome.RelativeY != metronomeOther.RelativeY {
		diffs = append(diffs, metronome.GongMarshallField(stage, "RelativeY"))
	}
	if metronome.FontFamily != metronomeOther.FontFamily {
		diffs = append(diffs, metronome.GongMarshallField(stage, "FontFamily"))
	}
	if metronome.FontSize != metronomeOther.FontSize {
		diffs = append(diffs, metronome.GongMarshallField(stage, "FontSize"))
	}
	if metronome.Color != metronomeOther.Color {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Color"))
	}
	if metronome.Id != metronomeOther.Id {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, metronome, "BeatUnitTied", metronomeOther.BeatUnitTied, metronome.BeatUnitTied); ops != "" {
		diffs = append(diffs, ops)
	}
	if metronome.PerMinute != metronomeOther.PerMinute {
		diffs = append(diffs, metronome.GongMarshallField(stage, "PerMinute"))
	}
	if ops := __gong__diffSliceOfPointers(stage, metronome, "MetronomeNote", metronomeOther.MetronomeNote, metronome.MetronomeNote); ops != "" {
		diffs = append(diffs, ops)
	}
	if metronome.MetronomeRelation != metronomeOther.MetronomeRelation {
		diffs = append(diffs, metronome.GongMarshallField(stage, "MetronomeRelation"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (metronomebeam *MetronomeBeam) GongDiff(stage *Stage, metronomebeamOther *MetronomeBeam) (diffs []string) {
	// insertion point for field diffs
	if metronomebeam.Name != metronomebeamOther.Name {
		diffs = append(diffs, metronomebeam.GongMarshallField(stage, "Name"))
	}
	if metronomebeam.Number != metronomebeamOther.Number {
		diffs = append(diffs, metronomebeam.GongMarshallField(stage, "Number"))
	}
	if metronomebeam.Value != metronomebeamOther.Value {
		diffs = append(diffs, metronomebeam.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (metronomenote *MetronomeNote) GongDiff(stage *Stage, metronomenoteOther *MetronomeNote) (diffs []string) {
	// insertion point for field diffs
	if metronomenote.Name != metronomenoteOther.Name {
		diffs = append(diffs, metronomenote.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, metronomenote, "MetronomeBeam", metronomenoteOther.MetronomeBeam, metronomenote.MetronomeBeam); ops != "" {
		diffs = append(diffs, ops)
	}
	if metronomenote.MetronomeTied != metronomenoteOther.MetronomeTied {
		diffs = append(diffs, metronomenote.GongMarshallField(stage, "MetronomeTied"))
	}
	if metronomenote.MetronomeTuplet != metronomenoteOther.MetronomeTuplet {
		diffs = append(diffs, metronomenote.GongMarshallField(stage, "MetronomeTuplet"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (metronometied *MetronomeTied) GongDiff(stage *Stage, metronometiedOther *MetronomeTied) (diffs []string) {
	// insertion point for field diffs
	if metronometied.Name != metronometiedOther.Name {
		diffs = append(diffs, metronometied.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (metronometuplet *MetronomeTuplet) GongDiff(stage *Stage, metronometupletOther *MetronomeTuplet) (diffs []string) {
	// insertion point for field diffs
	if metronometuplet.Name != metronometupletOther.Name {
		diffs = append(diffs, metronometuplet.GongMarshallField(stage, "Name"))
	}
	if metronometuplet.ShowNumber != metronometupletOther.ShowNumber {
		diffs = append(diffs, metronometuplet.GongMarshallField(stage, "ShowNumber"))
	}
	if metronometuplet.ActualNotes != metronometupletOther.ActualNotes {
		diffs = append(diffs, metronometuplet.GongMarshallField(stage, "ActualNotes"))
	}
	if metronometuplet.NormalNotes != metronometupletOther.NormalNotes {
		diffs = append(diffs, metronometuplet.GongMarshallField(stage, "NormalNotes"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (numeral *Numeral) GongDiff(stage *Stage, numeralOther *Numeral) (diffs []string) {
	// insertion point for field diffs
	if numeral.Name != numeralOther.Name {
		diffs = append(diffs, numeral.GongMarshallField(stage, "Name"))
	}
	if numeral.NumeralRoot != numeralOther.NumeralRoot {
		diffs = append(diffs, numeral.GongMarshallField(stage, "NumeralRoot"))
	}
	if numeral.NumeralAlter != numeralOther.NumeralAlter {
		diffs = append(diffs, numeral.GongMarshallField(stage, "NumeralAlter"))
	}
	if numeral.NumeralKey != numeralOther.NumeralKey {
		diffs = append(diffs, numeral.GongMarshallField(stage, "NumeralKey"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (numeralkey *NumeralKey) GongDiff(stage *Stage, numeralkeyOther *NumeralKey) (diffs []string) {
	// insertion point for field diffs
	if numeralkey.Name != numeralkeyOther.Name {
		diffs = append(diffs, numeralkey.GongMarshallField(stage, "Name"))
	}
	if numeralkey.NumeralFifths != numeralkeyOther.NumeralFifths {
		diffs = append(diffs, numeralkey.GongMarshallField(stage, "NumeralFifths"))
	}
	if numeralkey.NumeralMode != numeralkeyOther.NumeralMode {
		diffs = append(diffs, numeralkey.GongMarshallField(stage, "NumeralMode"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (numeralroot *NumeralRoot) GongDiff(stage *Stage, numeralrootOther *NumeralRoot) (diffs []string) {
	// insertion point for field diffs
	if numeralroot.Name != numeralrootOther.Name {
		diffs = append(diffs, numeralroot.GongMarshallField(stage, "Name"))
	}
	if numeralroot.Text != numeralrootOther.Text {
		diffs = append(diffs, numeralroot.GongMarshallField(stage, "Text"))
	}
	if numeralroot.DefaultX != numeralrootOther.DefaultX {
		diffs = append(diffs, numeralroot.GongMarshallField(stage, "DefaultX"))
	}
	if numeralroot.DefaultY != numeralrootOther.DefaultY {
		diffs = append(diffs, numeralroot.GongMarshallField(stage, "DefaultY"))
	}
	if numeralroot.RelativeX != numeralrootOther.RelativeX {
		diffs = append(diffs, numeralroot.GongMarshallField(stage, "RelativeX"))
	}
	if numeralroot.RelativeY != numeralrootOther.RelativeY {
		diffs = append(diffs, numeralroot.GongMarshallField(stage, "RelativeY"))
	}
	if numeralroot.FontFamily != numeralrootOther.FontFamily {
		diffs = append(diffs, numeralroot.GongMarshallField(stage, "FontFamily"))
	}
	if numeralroot.FontSize != numeralrootOther.FontSize {
		diffs = append(diffs, numeralroot.GongMarshallField(stage, "FontSize"))
	}
	if numeralroot.Color != numeralrootOther.Color {
		diffs = append(diffs, numeralroot.GongMarshallField(stage, "Color"))
	}
	if numeralroot.Value != numeralrootOther.Value {
		diffs = append(diffs, numeralroot.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (octaveshift *OctaveShift) GongDiff(stage *Stage, octaveshiftOther *OctaveShift) (diffs []string) {
	// insertion point for field diffs
	if octaveshift.Name != octaveshiftOther.Name {
		diffs = append(diffs, octaveshift.GongMarshallField(stage, "Name"))
	}
	if octaveshift.Type != octaveshiftOther.Type {
		diffs = append(diffs, octaveshift.GongMarshallField(stage, "Type"))
	}
	if octaveshift.Number != octaveshiftOther.Number {
		diffs = append(diffs, octaveshift.GongMarshallField(stage, "Number"))
	}
	if octaveshift.Size != octaveshiftOther.Size {
		diffs = append(diffs, octaveshift.GongMarshallField(stage, "Size"))
	}
	if octaveshift.DashLength != octaveshiftOther.DashLength {
		diffs = append(diffs, octaveshift.GongMarshallField(stage, "DashLength"))
	}
	if octaveshift.SpaceLength != octaveshiftOther.SpaceLength {
		diffs = append(diffs, octaveshift.GongMarshallField(stage, "SpaceLength"))
	}
	if octaveshift.DefaultX != octaveshiftOther.DefaultX {
		diffs = append(diffs, octaveshift.GongMarshallField(stage, "DefaultX"))
	}
	if octaveshift.DefaultY != octaveshiftOther.DefaultY {
		diffs = append(diffs, octaveshift.GongMarshallField(stage, "DefaultY"))
	}
	if octaveshift.RelativeX != octaveshiftOther.RelativeX {
		diffs = append(diffs, octaveshift.GongMarshallField(stage, "RelativeX"))
	}
	if octaveshift.RelativeY != octaveshiftOther.RelativeY {
		diffs = append(diffs, octaveshift.GongMarshallField(stage, "RelativeY"))
	}
	if octaveshift.FontFamily != octaveshiftOther.FontFamily {
		diffs = append(diffs, octaveshift.GongMarshallField(stage, "FontFamily"))
	}
	if octaveshift.FontSize != octaveshiftOther.FontSize {
		diffs = append(diffs, octaveshift.GongMarshallField(stage, "FontSize"))
	}
	if octaveshift.Color != octaveshiftOther.Color {
		diffs = append(diffs, octaveshift.GongMarshallField(stage, "Color"))
	}
	if octaveshift.Id != octaveshiftOther.Id {
		diffs = append(diffs, octaveshift.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (offset *Offset) GongDiff(stage *Stage, offsetOther *Offset) (diffs []string) {
	// insertion point for field diffs
	if offset.Name != offsetOther.Name {
		diffs = append(diffs, offset.GongMarshallField(stage, "Name"))
	}
	if offset.Value != offsetOther.Value {
		diffs = append(diffs, offset.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (otherdirection *OtherDirection) GongDiff(stage *Stage, otherdirectionOther *OtherDirection) (diffs []string) {
	// insertion point for field diffs
	if otherdirection.Name != otherdirectionOther.Name {
		diffs = append(diffs, otherdirection.GongMarshallField(stage, "Name"))
	}
	if otherdirection.DefaultX != otherdirectionOther.DefaultX {
		diffs = append(diffs, otherdirection.GongMarshallField(stage, "DefaultX"))
	}
	if otherdirection.DefaultY != otherdirectionOther.DefaultY {
		diffs = append(diffs, otherdirection.GongMarshallField(stage, "DefaultY"))
	}
	if otherdirection.RelativeX != otherdirectionOther.RelativeX {
		diffs = append(diffs, otherdirection.GongMarshallField(stage, "RelativeX"))
	}
	if otherdirection.RelativeY != otherdirectionOther.RelativeY {
		diffs = append(diffs, otherdirection.GongMarshallField(stage, "RelativeY"))
	}
	if otherdirection.FontFamily != otherdirectionOther.FontFamily {
		diffs = append(diffs, otherdirection.GongMarshallField(stage, "FontFamily"))
	}
	if otherdirection.FontSize != otherdirectionOther.FontSize {
		diffs = append(diffs, otherdirection.GongMarshallField(stage, "FontSize"))
	}
	if otherdirection.Color != otherdirectionOther.Color {
		diffs = append(diffs, otherdirection.GongMarshallField(stage, "Color"))
	}
	if otherdirection.Smufl != otherdirectionOther.Smufl {
		diffs = append(diffs, otherdirection.GongMarshallField(stage, "Smufl"))
	}
	if otherdirection.Id != otherdirectionOther.Id {
		diffs = append(diffs, otherdirection.GongMarshallField(stage, "Id"))
	}
	if otherdirection.Value != otherdirectionOther.Value {
		diffs = append(diffs, otherdirection.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (otherlistening *OtherListening) GongDiff(stage *Stage, otherlisteningOther *OtherListening) (diffs []string) {
	// insertion point for field diffs
	if otherlistening.Name != otherlisteningOther.Name {
		diffs = append(diffs, otherlistening.GongMarshallField(stage, "Name"))
	}
	if otherlistening.Type != otherlisteningOther.Type {
		diffs = append(diffs, otherlistening.GongMarshallField(stage, "Type"))
	}
	if otherlistening.Player != otherlisteningOther.Player {
		diffs = append(diffs, otherlistening.GongMarshallField(stage, "Player"))
	}
	if otherlistening.TimeOnly != otherlisteningOther.TimeOnly {
		diffs = append(diffs, otherlistening.GongMarshallField(stage, "TimeOnly"))
	}
	if otherlistening.Value != otherlisteningOther.Value {
		diffs = append(diffs, otherlistening.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pedal *Pedal) GongDiff(stage *Stage, pedalOther *Pedal) (diffs []string) {
	// insertion point for field diffs
	if pedal.Name != pedalOther.Name {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Name"))
	}
	if pedal.Type != pedalOther.Type {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Type"))
	}
	if pedal.Number != pedalOther.Number {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Number"))
	}
	if pedal.DefaultX != pedalOther.DefaultX {
		diffs = append(diffs, pedal.GongMarshallField(stage, "DefaultX"))
	}
	if pedal.DefaultY != pedalOther.DefaultY {
		diffs = append(diffs, pedal.GongMarshallField(stage, "DefaultY"))
	}
	if pedal.RelativeX != pedalOther.RelativeX {
		diffs = append(diffs, pedal.GongMarshallField(stage, "RelativeX"))
	}
	if pedal.RelativeY != pedalOther.RelativeY {
		diffs = append(diffs, pedal.GongMarshallField(stage, "RelativeY"))
	}
	if pedal.FontFamily != pedalOther.FontFamily {
		diffs = append(diffs, pedal.GongMarshallField(stage, "FontFamily"))
	}
	if pedal.FontSize != pedalOther.FontSize {
		diffs = append(diffs, pedal.GongMarshallField(stage, "FontSize"))
	}
	if pedal.Color != pedalOther.Color {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Color"))
	}
	if pedal.Id != pedalOther.Id {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pedaltuning *PedalTuning) GongDiff(stage *Stage, pedaltuningOther *PedalTuning) (diffs []string) {
	// insertion point for field diffs
	if pedaltuning.Name != pedaltuningOther.Name {
		diffs = append(diffs, pedaltuning.GongMarshallField(stage, "Name"))
	}
	if pedaltuning.PedalAlter != pedaltuningOther.PedalAlter {
		diffs = append(diffs, pedaltuning.GongMarshallField(stage, "PedalAlter"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (perminute *PerMinute) GongDiff(stage *Stage, perminuteOther *PerMinute) (diffs []string) {
	// insertion point for field diffs
	if perminute.Name != perminuteOther.Name {
		diffs = append(diffs, perminute.GongMarshallField(stage, "Name"))
	}
	if perminute.FontFamily != perminuteOther.FontFamily {
		diffs = append(diffs, perminute.GongMarshallField(stage, "FontFamily"))
	}
	if perminute.FontSize != perminuteOther.FontSize {
		diffs = append(diffs, perminute.GongMarshallField(stage, "FontSize"))
	}
	if perminute.Value != perminuteOther.Value {
		diffs = append(diffs, perminute.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (percussion *Percussion) GongDiff(stage *Stage, percussionOther *Percussion) (diffs []string) {
	// insertion point for field diffs
	if percussion.Name != percussionOther.Name {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Name"))
	}
	if percussion.DefaultX != percussionOther.DefaultX {
		diffs = append(diffs, percussion.GongMarshallField(stage, "DefaultX"))
	}
	if percussion.DefaultY != percussionOther.DefaultY {
		diffs = append(diffs, percussion.GongMarshallField(stage, "DefaultY"))
	}
	if percussion.RelativeX != percussionOther.RelativeX {
		diffs = append(diffs, percussion.GongMarshallField(stage, "RelativeX"))
	}
	if percussion.RelativeY != percussionOther.RelativeY {
		diffs = append(diffs, percussion.GongMarshallField(stage, "RelativeY"))
	}
	if percussion.FontFamily != percussionOther.FontFamily {
		diffs = append(diffs, percussion.GongMarshallField(stage, "FontFamily"))
	}
	if percussion.FontSize != percussionOther.FontSize {
		diffs = append(diffs, percussion.GongMarshallField(stage, "FontSize"))
	}
	if percussion.Color != percussionOther.Color {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Color"))
	}
	if percussion.Id != percussionOther.Id {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Id"))
	}
	if percussion.Glass != percussionOther.Glass {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Glass"))
	}
	if percussion.Metal != percussionOther.Metal {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Metal"))
	}
	if percussion.Wood != percussionOther.Wood {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Wood"))
	}
	if percussion.Pitched != percussionOther.Pitched {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Pitched"))
	}
	if percussion.Membrane != percussionOther.Membrane {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Membrane"))
	}
	if percussion.Effect != percussionOther.Effect {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Effect"))
	}
	if percussion.Timpani != percussionOther.Timpani {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Timpani"))
	}
	if percussion.Beater != percussionOther.Beater {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Beater"))
	}
	if percussion.Stick != percussionOther.Stick {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Stick"))
	}
	if percussion.StickLocation != percussionOther.StickLocation {
		diffs = append(diffs, percussion.GongMarshallField(stage, "StickLocation"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pitched *Pitched) GongDiff(stage *Stage, pitchedOther *Pitched) (diffs []string) {
	// insertion point for field diffs
	if pitched.Name != pitchedOther.Name {
		diffs = append(diffs, pitched.GongMarshallField(stage, "Name"))
	}
	if pitched.Smufl != pitchedOther.Smufl {
		diffs = append(diffs, pitched.GongMarshallField(stage, "Smufl"))
	}
	if pitched.Value != pitchedOther.Value {
		diffs = append(diffs, pitched.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (principalvoice *PrincipalVoice) GongDiff(stage *Stage, principalvoiceOther *PrincipalVoice) (diffs []string) {
	// insertion point for field diffs
	if principalvoice.Name != principalvoiceOther.Name {
		diffs = append(diffs, principalvoice.GongMarshallField(stage, "Name"))
	}
	if principalvoice.Symbol != principalvoiceOther.Symbol {
		diffs = append(diffs, principalvoice.GongMarshallField(stage, "Symbol"))
	}
	if principalvoice.DefaultX != principalvoiceOther.DefaultX {
		diffs = append(diffs, principalvoice.GongMarshallField(stage, "DefaultX"))
	}
	if principalvoice.DefaultY != principalvoiceOther.DefaultY {
		diffs = append(diffs, principalvoice.GongMarshallField(stage, "DefaultY"))
	}
	if principalvoice.RelativeX != principalvoiceOther.RelativeX {
		diffs = append(diffs, principalvoice.GongMarshallField(stage, "RelativeX"))
	}
	if principalvoice.RelativeY != principalvoiceOther.RelativeY {
		diffs = append(diffs, principalvoice.GongMarshallField(stage, "RelativeY"))
	}
	if principalvoice.FontFamily != principalvoiceOther.FontFamily {
		diffs = append(diffs, principalvoice.GongMarshallField(stage, "FontFamily"))
	}
	if principalvoice.FontSize != principalvoiceOther.FontSize {
		diffs = append(diffs, principalvoice.GongMarshallField(stage, "FontSize"))
	}
	if principalvoice.Color != principalvoiceOther.Color {
		diffs = append(diffs, principalvoice.GongMarshallField(stage, "Color"))
	}
	if principalvoice.Id != principalvoiceOther.Id {
		diffs = append(diffs, principalvoice.GongMarshallField(stage, "Id"))
	}
	if principalvoice.Value != principalvoiceOther.Value {
		diffs = append(diffs, principalvoice.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (print *Print) GongDiff(stage *Stage, printOther *Print) (diffs []string) {
	// insertion point for field diffs
	if print.Name != printOther.Name {
		diffs = append(diffs, print.GongMarshallField(stage, "Name"))
	}
	if print.StaffSpacing != printOther.StaffSpacing {
		diffs = append(diffs, print.GongMarshallField(stage, "StaffSpacing"))
	}
	if print.BlankPage != printOther.BlankPage {
		diffs = append(diffs, print.GongMarshallField(stage, "BlankPage"))
	}
	if print.PageNumber != printOther.PageNumber {
		diffs = append(diffs, print.GongMarshallField(stage, "PageNumber"))
	}
	if print.Id != printOther.Id {
		diffs = append(diffs, print.GongMarshallField(stage, "Id"))
	}
	if print.MeasureNumbering != printOther.MeasureNumbering {
		diffs = append(diffs, print.GongMarshallField(stage, "MeasureNumbering"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (root *Root) GongDiff(stage *Stage, rootOther *Root) (diffs []string) {
	// insertion point for field diffs
	if root.Name != rootOther.Name {
		diffs = append(diffs, root.GongMarshallField(stage, "Name"))
	}
	if root.RootStep != rootOther.RootStep {
		diffs = append(diffs, root.GongMarshallField(stage, "RootStep"))
	}
	if root.RootAlter != rootOther.RootAlter {
		diffs = append(diffs, root.GongMarshallField(stage, "RootAlter"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (rootstep *RootStep) GongDiff(stage *Stage, rootstepOther *RootStep) (diffs []string) {
	// insertion point for field diffs
	if rootstep.Name != rootstepOther.Name {
		diffs = append(diffs, rootstep.GongMarshallField(stage, "Name"))
	}
	if rootstep.Text != rootstepOther.Text {
		diffs = append(diffs, rootstep.GongMarshallField(stage, "Text"))
	}
	if rootstep.DefaultX != rootstepOther.DefaultX {
		diffs = append(diffs, rootstep.GongMarshallField(stage, "DefaultX"))
	}
	if rootstep.DefaultY != rootstepOther.DefaultY {
		diffs = append(diffs, rootstep.GongMarshallField(stage, "DefaultY"))
	}
	if rootstep.RelativeX != rootstepOther.RelativeX {
		diffs = append(diffs, rootstep.GongMarshallField(stage, "RelativeX"))
	}
	if rootstep.RelativeY != rootstepOther.RelativeY {
		diffs = append(diffs, rootstep.GongMarshallField(stage, "RelativeY"))
	}
	if rootstep.FontFamily != rootstepOther.FontFamily {
		diffs = append(diffs, rootstep.GongMarshallField(stage, "FontFamily"))
	}
	if rootstep.FontSize != rootstepOther.FontSize {
		diffs = append(diffs, rootstep.GongMarshallField(stage, "FontSize"))
	}
	if rootstep.Color != rootstepOther.Color {
		diffs = append(diffs, rootstep.GongMarshallField(stage, "Color"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (scordatura *Scordatura) GongDiff(stage *Stage, scordaturaOther *Scordatura) (diffs []string) {
	// insertion point for field diffs
	if scordatura.Name != scordaturaOther.Name {
		diffs = append(diffs, scordatura.GongMarshallField(stage, "Name"))
	}
	if scordatura.Id != scordaturaOther.Id {
		diffs = append(diffs, scordatura.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, scordatura, "Accord", scordaturaOther.Accord, scordatura.Accord); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (sound *Sound) GongDiff(stage *Stage, soundOther *Sound) (diffs []string) {
	// insertion point for field diffs
	if sound.Name != soundOther.Name {
		diffs = append(diffs, sound.GongMarshallField(stage, "Name"))
	}
	if sound.Tempo != soundOther.Tempo {
		diffs = append(diffs, sound.GongMarshallField(stage, "Tempo"))
	}
	if sound.Dynamics != soundOther.Dynamics {
		diffs = append(diffs, sound.GongMarshallField(stage, "Dynamics"))
	}
	if sound.Segno != soundOther.Segno {
		diffs = append(diffs, sound.GongMarshallField(stage, "Segno"))
	}
	if sound.Dalsegno != soundOther.Dalsegno {
		diffs = append(diffs, sound.GongMarshallField(stage, "Dalsegno"))
	}
	if sound.Coda != soundOther.Coda {
		diffs = append(diffs, sound.GongMarshallField(stage, "Coda"))
	}
	if sound.Tocoda != soundOther.Tocoda {
		diffs = append(diffs, sound.GongMarshallField(stage, "Tocoda"))
	}
	if sound.Divisions != soundOther.Divisions {
		diffs = append(diffs, sound.GongMarshallField(stage, "Divisions"))
	}
	if sound.Fine != soundOther.Fine {
		diffs = append(diffs, sound.GongMarshallField(stage, "Fine"))
	}
	if sound.TimeOnly != soundOther.TimeOnly {
		diffs = append(diffs, sound.GongMarshallField(stage, "TimeOnly"))
	}
	if sound.Pan != soundOther.Pan {
		diffs = append(diffs, sound.GongMarshallField(stage, "Pan"))
	}
	if sound.Elevation != soundOther.Elevation {
		diffs = append(diffs, sound.GongMarshallField(stage, "Elevation"))
	}
	if sound.DamperPedal != soundOther.DamperPedal {
		diffs = append(diffs, sound.GongMarshallField(stage, "DamperPedal"))
	}
	if sound.SoftPedal != soundOther.SoftPedal {
		diffs = append(diffs, sound.GongMarshallField(stage, "SoftPedal"))
	}
	if sound.SostenutoPedal != soundOther.SostenutoPedal {
		diffs = append(diffs, sound.GongMarshallField(stage, "SostenutoPedal"))
	}
	if sound.Id != soundOther.Id {
		diffs = append(diffs, sound.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, sound, "InstrumentChange", soundOther.InstrumentChange, sound.InstrumentChange); ops != "" {
		diffs = append(diffs, ops)
	}
	if sound.Swing != soundOther.Swing {
		diffs = append(diffs, sound.GongMarshallField(stage, "Swing"))
	}
	if sound.Offset != soundOther.Offset {
		diffs = append(diffs, sound.GongMarshallField(stage, "Offset"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (staffdivide *StaffDivide) GongDiff(stage *Stage, staffdivideOther *StaffDivide) (diffs []string) {
	// insertion point for field diffs
	if staffdivide.Name != staffdivideOther.Name {
		diffs = append(diffs, staffdivide.GongMarshallField(stage, "Name"))
	}
	if staffdivide.Type != staffdivideOther.Type {
		diffs = append(diffs, staffdivide.GongMarshallField(stage, "Type"))
	}
	if staffdivide.DefaultX != staffdivideOther.DefaultX {
		diffs = append(diffs, staffdivide.GongMarshallField(stage, "DefaultX"))
	}
	if staffdivide.DefaultY != staffdivideOther.DefaultY {
		diffs = append(diffs, staffdivide.GongMarshallField(stage, "DefaultY"))
	}
	if staffdivide.RelativeX != staffdivideOther.RelativeX {
		diffs = append(diffs, staffdivide.GongMarshallField(stage, "RelativeX"))
	}
	if staffdivide.RelativeY != staffdivideOther.RelativeY {
		diffs = append(diffs, staffdivide.GongMarshallField(stage, "RelativeY"))
	}
	if staffdivide.FontFamily != staffdivideOther.FontFamily {
		diffs = append(diffs, staffdivide.GongMarshallField(stage, "FontFamily"))
	}
	if staffdivide.FontSize != staffdivideOther.FontSize {
		diffs = append(diffs, staffdivide.GongMarshallField(stage, "FontSize"))
	}
	if staffdivide.Color != staffdivideOther.Color {
		diffs = append(diffs, staffdivide.GongMarshallField(stage, "Color"))
	}
	if staffdivide.Id != staffdivideOther.Id {
		diffs = append(diffs, staffdivide.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stick *Stick) GongDiff(stage *Stage, stickOther *Stick) (diffs []string) {
	// insertion point for field diffs
	if stick.Name != stickOther.Name {
		diffs = append(diffs, stick.GongMarshallField(stage, "Name"))
	}
	if stick.Tip != stickOther.Tip {
		diffs = append(diffs, stick.GongMarshallField(stage, "Tip"))
	}
	if stick.StickType != stickOther.StickType {
		diffs = append(diffs, stick.GongMarshallField(stage, "StickType"))
	}
	if stick.StickMaterial != stickOther.StickMaterial {
		diffs = append(diffs, stick.GongMarshallField(stage, "StickMaterial"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stringmute *StringMute) GongDiff(stage *Stage, stringmuteOther *StringMute) (diffs []string) {
	// insertion point for field diffs
	if stringmute.Name != stringmuteOther.Name {
		diffs = append(diffs, stringmute.GongMarshallField(stage, "Name"))
	}
	if stringmute.Type != stringmuteOther.Type {
		diffs = append(diffs, stringmute.GongMarshallField(stage, "Type"))
	}
	if stringmute.DefaultX != stringmuteOther.DefaultX {
		diffs = append(diffs, stringmute.GongMarshallField(stage, "DefaultX"))
	}
	if stringmute.DefaultY != stringmuteOther.DefaultY {
		diffs = append(diffs, stringmute.GongMarshallField(stage, "DefaultY"))
	}
	if stringmute.RelativeX != stringmuteOther.RelativeX {
		diffs = append(diffs, stringmute.GongMarshallField(stage, "RelativeX"))
	}
	if stringmute.RelativeY != stringmuteOther.RelativeY {
		diffs = append(diffs, stringmute.GongMarshallField(stage, "RelativeY"))
	}
	if stringmute.FontFamily != stringmuteOther.FontFamily {
		diffs = append(diffs, stringmute.GongMarshallField(stage, "FontFamily"))
	}
	if stringmute.FontSize != stringmuteOther.FontSize {
		diffs = append(diffs, stringmute.GongMarshallField(stage, "FontSize"))
	}
	if stringmute.Color != stringmuteOther.Color {
		diffs = append(diffs, stringmute.GongMarshallField(stage, "Color"))
	}
	if stringmute.Id != stringmuteOther.Id {
		diffs = append(diffs, stringmute.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (styletext *StyleText) GongDiff(stage *Stage, styletextOther *StyleText) (diffs []string) {
	// insertion point for field diffs
	if styletext.Name != styletextOther.Name {
		diffs = append(diffs, styletext.GongMarshallField(stage, "Name"))
	}
	if styletext.DefaultX != styletextOther.DefaultX {
		diffs = append(diffs, styletext.GongMarshallField(stage, "DefaultX"))
	}
	if styletext.DefaultY != styletextOther.DefaultY {
		diffs = append(diffs, styletext.GongMarshallField(stage, "DefaultY"))
	}
	if styletext.RelativeX != styletextOther.RelativeX {
		diffs = append(diffs, styletext.GongMarshallField(stage, "RelativeX"))
	}
	if styletext.RelativeY != styletextOther.RelativeY {
		diffs = append(diffs, styletext.GongMarshallField(stage, "RelativeY"))
	}
	if styletext.FontFamily != styletextOther.FontFamily {
		diffs = append(diffs, styletext.GongMarshallField(stage, "FontFamily"))
	}
	if styletext.FontSize != styletextOther.FontSize {
		diffs = append(diffs, styletext.GongMarshallField(stage, "FontSize"))
	}
	if styletext.Color != styletextOther.Color {
		diffs = append(diffs, styletext.GongMarshallField(stage, "Color"))
	}
	if styletext.Value != styletextOther.Value {
		diffs = append(diffs, styletext.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (swing *Swing) GongDiff(stage *Stage, swingOther *Swing) (diffs []string) {
	// insertion point for field diffs
	if swing.Name != swingOther.Name {
		diffs = append(diffs, swing.GongMarshallField(stage, "Name"))
	}
	if swing.First != swingOther.First {
		diffs = append(diffs, swing.GongMarshallField(stage, "First"))
	}
	if swing.Second != swingOther.Second {
		diffs = append(diffs, swing.GongMarshallField(stage, "Second"))
	}
	if swing.SwingStyle != swingOther.SwingStyle {
		diffs = append(diffs, swing.GongMarshallField(stage, "SwingStyle"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (sync *Sync) GongDiff(stage *Stage, syncOther *Sync) (diffs []string) {
	// insertion point for field diffs
	if sync.Name != syncOther.Name {
		diffs = append(diffs, sync.GongMarshallField(stage, "Name"))
	}
	if sync.Type != syncOther.Type {
		diffs = append(diffs, sync.GongMarshallField(stage, "Type"))
	}
	if sync.Latency != syncOther.Latency {
		diffs = append(diffs, sync.GongMarshallField(stage, "Latency"))
	}
	if sync.Player != syncOther.Player {
		diffs = append(diffs, sync.GongMarshallField(stage, "Player"))
	}
	if sync.TimeOnly != syncOther.TimeOnly {
		diffs = append(diffs, sync.GongMarshallField(stage, "TimeOnly"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (timemodification *TimeModification) GongDiff(stage *Stage, timemodificationOther *TimeModification) (diffs []string) {
	// insertion point for field diffs
	if timemodification.Name != timemodificationOther.Name {
		diffs = append(diffs, timemodification.GongMarshallField(stage, "Name"))
	}
	if timemodification.ActualNotes != timemodificationOther.ActualNotes {
		diffs = append(diffs, timemodification.GongMarshallField(stage, "ActualNotes"))
	}
	if timemodification.NormalNotes != timemodificationOther.NormalNotes {
		diffs = append(diffs, timemodification.GongMarshallField(stage, "NormalNotes"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (timpani *Timpani) GongDiff(stage *Stage, timpaniOther *Timpani) (diffs []string) {
	// insertion point for field diffs
	if timpani.Name != timpaniOther.Name {
		diffs = append(diffs, timpani.GongMarshallField(stage, "Name"))
	}
	if timpani.Smufl != timpaniOther.Smufl {
		diffs = append(diffs, timpani.GongMarshallField(stage, "Smufl"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (virtualinstrument *VirtualInstrument) GongDiff(stage *Stage, virtualinstrumentOther *VirtualInstrument) (diffs []string) {
	// insertion point for field diffs
	if virtualinstrument.Name != virtualinstrumentOther.Name {
		diffs = append(diffs, virtualinstrument.GongMarshallField(stage, "Name"))
	}
	if virtualinstrument.VirtualLibrary != virtualinstrumentOther.VirtualLibrary {
		diffs = append(diffs, virtualinstrument.GongMarshallField(stage, "VirtualLibrary"))
	}
	if virtualinstrument.VirtualName != virtualinstrumentOther.VirtualName {
		diffs = append(diffs, virtualinstrument.GongMarshallField(stage, "VirtualName"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (wedge *Wedge) GongDiff(stage *Stage, wedgeOther *Wedge) (diffs []string) {
	// insertion point for field diffs
	if wedge.Name != wedgeOther.Name {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Name"))
	}
	if wedge.Type != wedgeOther.Type {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Type"))
	}
	if wedge.Number != wedgeOther.Number {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Number"))
	}
	if wedge.Spread != wedgeOther.Spread {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Spread"))
	}
	if wedge.DashLength != wedgeOther.DashLength {
		diffs = append(diffs, wedge.GongMarshallField(stage, "DashLength"))
	}
	if wedge.SpaceLength != wedgeOther.SpaceLength {
		diffs = append(diffs, wedge.GongMarshallField(stage, "SpaceLength"))
	}
	if wedge.DefaultX != wedgeOther.DefaultX {
		diffs = append(diffs, wedge.GongMarshallField(stage, "DefaultX"))
	}
	if wedge.DefaultY != wedgeOther.DefaultY {
		diffs = append(diffs, wedge.GongMarshallField(stage, "DefaultY"))
	}
	if wedge.RelativeX != wedgeOther.RelativeX {
		diffs = append(diffs, wedge.GongMarshallField(stage, "RelativeX"))
	}
	if wedge.RelativeY != wedgeOther.RelativeY {
		diffs = append(diffs, wedge.GongMarshallField(stage, "RelativeY"))
	}
	if wedge.Color != wedgeOther.Color {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Color"))
	}
	if wedge.Id != wedgeOther.Id {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (wood *Wood) GongDiff(stage *Stage, woodOther *Wood) (diffs []string) {
	// insertion point for field diffs
	if wood.Name != woodOther.Name {
		diffs = append(diffs, wood.GongMarshallField(stage, "Name"))
	}
	if wood.Smufl != woodOther.Smufl {
		diffs = append(diffs, wood.GongMarshallField(stage, "Smufl"))
	}
	if wood.Value != woodOther.Value {
		diffs = append(diffs, wood.GongMarshallField(stage, "Value"))
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
