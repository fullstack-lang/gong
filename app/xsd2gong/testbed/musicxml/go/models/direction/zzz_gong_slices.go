// generated code - do not edit
package direction

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Direction
	// insertion point per field
	stage.Direction_DirectionType_reverseMap = make(map[*DirectionType]*Direction)
	for direction := range stage.Directions {
		_ = direction
		for _, _directiontype := range direction.DirectionType {
			stage.Direction_DirectionType_reverseMap[_directiontype] = direction
		}
	}

	// Compute reverse map for named struct DirectionType
	// insertion point per field
	stage.DirectionType_Percussion_reverseMap = make(map[*Percussion]*DirectionType)
	for directiontype := range stage.DirectionTypes {
		_ = directiontype
		for _, _percussion := range directiontype.Percussion {
			stage.DirectionType_Percussion_reverseMap[_percussion] = directiontype
		}
	}

	// Compute reverse map for named struct Frame
	// insertion point per field
	stage.Frame_FrameNote_reverseMap = make(map[*FrameNote]*Frame)
	for frame := range stage.Frames {
		_ = frame
		for _, _framenote := range frame.FrameNote {
			stage.Frame_FrameNote_reverseMap[_framenote] = frame
		}
	}

	// Compute reverse map for named struct Grouping
	// insertion point per field
	stage.Grouping_Feature_reverseMap = make(map[*Feature]*Grouping)
	for grouping := range stage.Groupings {
		_ = grouping
		for _, _feature := range grouping.Feature {
			stage.Grouping_Feature_reverseMap[_feature] = grouping
		}
	}

	// Compute reverse map for named struct Harmony
	// insertion point per field
	stage.Harmony_Root_reverseMap = make(map[*Root]*Harmony)
	for harmony := range stage.Harmonys {
		_ = harmony
		for _, _root := range harmony.Root {
			stage.Harmony_Root_reverseMap[_root] = harmony
		}
	}
	stage.Harmony_Numeral_reverseMap = make(map[*Numeral]*Harmony)
	for harmony := range stage.Harmonys {
		_ = harmony
		for _, _numeral := range harmony.Numeral {
			stage.Harmony_Numeral_reverseMap[_numeral] = harmony
		}
	}
	stage.Harmony_Function_reverseMap = make(map[*StyleText]*Harmony)
	for harmony := range stage.Harmonys {
		_ = harmony
		for _, _styletext := range harmony.Function {
			stage.Harmony_Function_reverseMap[_styletext] = harmony
		}
	}
	stage.Harmony_Kind_reverseMap = make(map[*Kind]*Harmony)
	for harmony := range stage.Harmonys {
		_ = harmony
		for _, _kind := range harmony.Kind {
			stage.Harmony_Kind_reverseMap[_kind] = harmony
		}
	}
	stage.Harmony_Inversion_reverseMap = make(map[*Inversion]*Harmony)
	for harmony := range stage.Harmonys {
		_ = harmony
		for _, _inversion := range harmony.Inversion {
			stage.Harmony_Inversion_reverseMap[_inversion] = harmony
		}
	}
	stage.Harmony_Bass_reverseMap = make(map[*Bass]*Harmony)
	for harmony := range stage.Harmonys {
		_ = harmony
		for _, _bass := range harmony.Bass {
			stage.Harmony_Bass_reverseMap[_bass] = harmony
		}
	}
	stage.Harmony_Degree_reverseMap = make(map[*Degree]*Harmony)
	for harmony := range stage.Harmonys {
		_ = harmony
		for _, _degree := range harmony.Degree {
			stage.Harmony_Degree_reverseMap[_degree] = harmony
		}
	}

	// Compute reverse map for named struct HarpPedals
	// insertion point per field
	stage.HarpPedals_PedalTuning_reverseMap = make(map[*PedalTuning]*HarpPedals)
	for harppedals := range stage.HarpPedalss {
		_ = harppedals
		for _, _pedaltuning := range harppedals.PedalTuning {
			stage.HarpPedals_PedalTuning_reverseMap[_pedaltuning] = harppedals
		}
	}

	// Compute reverse map for named struct Listening
	// insertion point per field
	stage.Listening_Sync_reverseMap = make(map[*Sync]*Listening)
	for listening := range stage.Listenings {
		_ = listening
		for _, _sync := range listening.Sync {
			stage.Listening_Sync_reverseMap[_sync] = listening
		}
	}
	stage.Listening_OtherListening_reverseMap = make(map[*OtherListening]*Listening)
	for listening := range stage.Listenings {
		_ = listening
		for _, _otherlistening := range listening.OtherListening {
			stage.Listening_OtherListening_reverseMap[_otherlistening] = listening
		}
	}

	// Compute reverse map for named struct Metronome
	// insertion point per field
	stage.Metronome_BeatUnitTied_reverseMap = make(map[*BeatUnitTied]*Metronome)
	for metronome := range stage.Metronomes {
		_ = metronome
		for _, _beatunittied := range metronome.BeatUnitTied {
			stage.Metronome_BeatUnitTied_reverseMap[_beatunittied] = metronome
		}
	}
	stage.Metronome_MetronomeNote_reverseMap = make(map[*MetronomeNote]*Metronome)
	for metronome := range stage.Metronomes {
		_ = metronome
		for _, _metronomenote := range metronome.MetronomeNote {
			stage.Metronome_MetronomeNote_reverseMap[_metronomenote] = metronome
		}
	}

	// Compute reverse map for named struct MetronomeNote
	// insertion point per field
	stage.MetronomeNote_MetronomeBeam_reverseMap = make(map[*MetronomeBeam]*MetronomeNote)
	for metronomenote := range stage.MetronomeNotes {
		_ = metronomenote
		for _, _metronomebeam := range metronomenote.MetronomeBeam {
			stage.MetronomeNote_MetronomeBeam_reverseMap[_metronomebeam] = metronomenote
		}
	}

	// Compute reverse map for named struct Scordatura
	// insertion point per field
	stage.Scordatura_Accord_reverseMap = make(map[*Accord]*Scordatura)
	for scordatura := range stage.Scordaturas {
		_ = scordatura
		for _, _accord := range scordatura.Accord {
			stage.Scordatura_Accord_reverseMap[_accord] = scordatura
		}
	}

	// Compute reverse map for named struct Sound
	// insertion point per field
	stage.Sound_InstrumentChange_reverseMap = make(map[*InstrumentChange]*Sound)
	for sound := range stage.Sounds {
		_ = sound
		for _, _instrumentchange := range sound.InstrumentChange {
			stage.Sound_InstrumentChange_reverseMap[_instrumentchange] = sound
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.Accords)

	res = __gong__appendInstances(res, stage.AccordionRegistrations)

	res = __gong__appendInstances(res, stage.Barres)

	res = __gong__appendInstances(res, stage.Basss)

	res = __gong__appendInstances(res, stage.BassSteps)

	res = __gong__appendInstances(res, stage.BeatUnitTieds)

	res = __gong__appendInstances(res, stage.Beaters)

	res = __gong__appendInstances(res, stage.Brackets)

	res = __gong__appendInstances(res, stage.Dashess)

	res = __gong__appendInstances(res, stage.Degrees)

	res = __gong__appendInstances(res, stage.DegreeAlters)

	res = __gong__appendInstances(res, stage.DegreeTypes)

	res = __gong__appendInstances(res, stage.DegreeValues)

	res = __gong__appendInstances(res, stage.Directions)

	res = __gong__appendInstances(res, stage.DirectionTypes)

	res = __gong__appendInstances(res, stage.Effects)

	res = __gong__appendInstances(res, stage.Features)

	res = __gong__appendInstances(res, stage.FirstFrets)

	res = __gong__appendInstances(res, stage.Frames)

	res = __gong__appendInstances(res, stage.FrameNotes)

	res = __gong__appendInstances(res, stage.Glasss)

	res = __gong__appendInstances(res, stage.Groupings)

	res = __gong__appendInstances(res, stage.Harmonys)

	res = __gong__appendInstances(res, stage.HarmonyAlters)

	res = __gong__appendInstances(res, stage.HarpPedalss)

	res = __gong__appendInstances(res, stage.Images)

	res = __gong__appendInstances(res, stage.InstrumentChanges)

	res = __gong__appendInstances(res, stage.Inversions)

	res = __gong__appendInstances(res, stage.Kinds)

	res = __gong__appendInstances(res, stage.Listenings)

	res = __gong__appendInstances(res, stage.MeasureNumberings)

	res = __gong__appendInstances(res, stage.Membranes)

	res = __gong__appendInstances(res, stage.Metals)

	res = __gong__appendInstances(res, stage.Metronomes)

	res = __gong__appendInstances(res, stage.MetronomeBeams)

	res = __gong__appendInstances(res, stage.MetronomeNotes)

	res = __gong__appendInstances(res, stage.MetronomeTieds)

	res = __gong__appendInstances(res, stage.MetronomeTuplets)

	res = __gong__appendInstances(res, stage.Numerals)

	res = __gong__appendInstances(res, stage.NumeralKeys)

	res = __gong__appendInstances(res, stage.NumeralRoots)

	res = __gong__appendInstances(res, stage.OctaveShifts)

	res = __gong__appendInstances(res, stage.Offsets)

	res = __gong__appendInstances(res, stage.OtherDirections)

	res = __gong__appendInstances(res, stage.OtherListenings)

	res = __gong__appendInstances(res, stage.Pedals)

	res = __gong__appendInstances(res, stage.PedalTunings)

	res = __gong__appendInstances(res, stage.PerMinutes)

	res = __gong__appendInstances(res, stage.Percussions)

	res = __gong__appendInstances(res, stage.Pitcheds)

	res = __gong__appendInstances(res, stage.PrincipalVoices)

	res = __gong__appendInstances(res, stage.Prints)

	res = __gong__appendInstances(res, stage.Roots)

	res = __gong__appendInstances(res, stage.RootSteps)

	res = __gong__appendInstances(res, stage.Scordaturas)

	res = __gong__appendInstances(res, stage.Sounds)

	res = __gong__appendInstances(res, stage.StaffDivides)

	res = __gong__appendInstances(res, stage.Sticks)

	res = __gong__appendInstances(res, stage.StringMutes)

	res = __gong__appendInstances(res, stage.StyleTexts)

	res = __gong__appendInstances(res, stage.Swings)

	res = __gong__appendInstances(res, stage.Syncs)

	res = __gong__appendInstances(res, stage.TimeModifications)

	res = __gong__appendInstances(res, stage.Timpanis)

	res = __gong__appendInstances(res, stage.VirtualInstruments)

	res = __gong__appendInstances(res, stage.Wedges)

	res = __gong__appendInstances(res, stage.Woods)

	return
}

// insertion point per named struct
func (accord *Accord) GongCopy() GongstructIF {
	newInstance := new(Accord)
	accord.GongCopyBasicFields(newInstance)
	return newInstance
}

func (accordionregistration *AccordionRegistration) GongCopy() GongstructIF {
	newInstance := new(AccordionRegistration)
	accordionregistration.GongCopyBasicFields(newInstance)
	return newInstance
}

func (barre *Barre) GongCopy() GongstructIF {
	newInstance := new(Barre)
	barre.GongCopyBasicFields(newInstance)
	return newInstance
}

func (bass *Bass) GongCopy() GongstructIF {
	newInstance := new(Bass)
	bass.GongCopyBasicFields(newInstance)
	return newInstance
}

func (bassstep *BassStep) GongCopy() GongstructIF {
	newInstance := new(BassStep)
	bassstep.GongCopyBasicFields(newInstance)
	return newInstance
}

func (beatunittied *BeatUnitTied) GongCopy() GongstructIF {
	newInstance := new(BeatUnitTied)
	beatunittied.GongCopyBasicFields(newInstance)
	return newInstance
}

func (beater *Beater) GongCopy() GongstructIF {
	newInstance := new(Beater)
	beater.GongCopyBasicFields(newInstance)
	return newInstance
}

func (bracket *Bracket) GongCopy() GongstructIF {
	newInstance := new(Bracket)
	bracket.GongCopyBasicFields(newInstance)
	return newInstance
}

func (dashes *Dashes) GongCopy() GongstructIF {
	newInstance := new(Dashes)
	dashes.GongCopyBasicFields(newInstance)
	return newInstance
}

func (degree *Degree) GongCopy() GongstructIF {
	newInstance := new(Degree)
	degree.GongCopyBasicFields(newInstance)
	return newInstance
}

func (degreealter *DegreeAlter) GongCopy() GongstructIF {
	newInstance := new(DegreeAlter)
	degreealter.GongCopyBasicFields(newInstance)
	return newInstance
}

func (degreetype *DegreeType) GongCopy() GongstructIF {
	newInstance := new(DegreeType)
	degreetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (degreevalue *DegreeValue) GongCopy() GongstructIF {
	newInstance := new(DegreeValue)
	degreevalue.GongCopyBasicFields(newInstance)
	return newInstance
}

func (direction *Direction) GongCopy() GongstructIF {
	newInstance := new(Direction)
	direction.GongCopyBasicFields(newInstance)
	return newInstance
}

func (directiontype *DirectionType) GongCopy() GongstructIF {
	newInstance := new(DirectionType)
	directiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (effect *Effect) GongCopy() GongstructIF {
	newInstance := new(Effect)
	effect.GongCopyBasicFields(newInstance)
	return newInstance
}

func (feature *Feature) GongCopy() GongstructIF {
	newInstance := new(Feature)
	feature.GongCopyBasicFields(newInstance)
	return newInstance
}

func (firstfret *FirstFret) GongCopy() GongstructIF {
	newInstance := new(FirstFret)
	firstfret.GongCopyBasicFields(newInstance)
	return newInstance
}

func (frame *Frame) GongCopy() GongstructIF {
	newInstance := new(Frame)
	frame.GongCopyBasicFields(newInstance)
	return newInstance
}

func (framenote *FrameNote) GongCopy() GongstructIF {
	newInstance := new(FrameNote)
	framenote.GongCopyBasicFields(newInstance)
	return newInstance
}

func (glass *Glass) GongCopy() GongstructIF {
	newInstance := new(Glass)
	glass.GongCopyBasicFields(newInstance)
	return newInstance
}

func (grouping *Grouping) GongCopy() GongstructIF {
	newInstance := new(Grouping)
	grouping.GongCopyBasicFields(newInstance)
	return newInstance
}

func (harmony *Harmony) GongCopy() GongstructIF {
	newInstance := new(Harmony)
	harmony.GongCopyBasicFields(newInstance)
	return newInstance
}

func (harmonyalter *HarmonyAlter) GongCopy() GongstructIF {
	newInstance := new(HarmonyAlter)
	harmonyalter.GongCopyBasicFields(newInstance)
	return newInstance
}

func (harppedals *HarpPedals) GongCopy() GongstructIF {
	newInstance := new(HarpPedals)
	harppedals.GongCopyBasicFields(newInstance)
	return newInstance
}

func (image *Image) GongCopy() GongstructIF {
	newInstance := new(Image)
	image.GongCopyBasicFields(newInstance)
	return newInstance
}

func (instrumentchange *InstrumentChange) GongCopy() GongstructIF {
	newInstance := new(InstrumentChange)
	instrumentchange.GongCopyBasicFields(newInstance)
	return newInstance
}

func (inversion *Inversion) GongCopy() GongstructIF {
	newInstance := new(Inversion)
	inversion.GongCopyBasicFields(newInstance)
	return newInstance
}

func (kind *Kind) GongCopy() GongstructIF {
	newInstance := new(Kind)
	kind.GongCopyBasicFields(newInstance)
	return newInstance
}

func (listening *Listening) GongCopy() GongstructIF {
	newInstance := new(Listening)
	listening.GongCopyBasicFields(newInstance)
	return newInstance
}

func (measurenumbering *MeasureNumbering) GongCopy() GongstructIF {
	newInstance := new(MeasureNumbering)
	measurenumbering.GongCopyBasicFields(newInstance)
	return newInstance
}

func (membrane *Membrane) GongCopy() GongstructIF {
	newInstance := new(Membrane)
	membrane.GongCopyBasicFields(newInstance)
	return newInstance
}

func (metal *Metal) GongCopy() GongstructIF {
	newInstance := new(Metal)
	metal.GongCopyBasicFields(newInstance)
	return newInstance
}

func (metronome *Metronome) GongCopy() GongstructIF {
	newInstance := new(Metronome)
	metronome.GongCopyBasicFields(newInstance)
	return newInstance
}

func (metronomebeam *MetronomeBeam) GongCopy() GongstructIF {
	newInstance := new(MetronomeBeam)
	metronomebeam.GongCopyBasicFields(newInstance)
	return newInstance
}

func (metronomenote *MetronomeNote) GongCopy() GongstructIF {
	newInstance := new(MetronomeNote)
	metronomenote.GongCopyBasicFields(newInstance)
	return newInstance
}

func (metronometied *MetronomeTied) GongCopy() GongstructIF {
	newInstance := new(MetronomeTied)
	metronometied.GongCopyBasicFields(newInstance)
	return newInstance
}

func (metronometuplet *MetronomeTuplet) GongCopy() GongstructIF {
	newInstance := new(MetronomeTuplet)
	metronometuplet.GongCopyBasicFields(newInstance)
	return newInstance
}

func (numeral *Numeral) GongCopy() GongstructIF {
	newInstance := new(Numeral)
	numeral.GongCopyBasicFields(newInstance)
	return newInstance
}

func (numeralkey *NumeralKey) GongCopy() GongstructIF {
	newInstance := new(NumeralKey)
	numeralkey.GongCopyBasicFields(newInstance)
	return newInstance
}

func (numeralroot *NumeralRoot) GongCopy() GongstructIF {
	newInstance := new(NumeralRoot)
	numeralroot.GongCopyBasicFields(newInstance)
	return newInstance
}

func (octaveshift *OctaveShift) GongCopy() GongstructIF {
	newInstance := new(OctaveShift)
	octaveshift.GongCopyBasicFields(newInstance)
	return newInstance
}

func (offset *Offset) GongCopy() GongstructIF {
	newInstance := new(Offset)
	offset.GongCopyBasicFields(newInstance)
	return newInstance
}

func (otherdirection *OtherDirection) GongCopy() GongstructIF {
	newInstance := new(OtherDirection)
	otherdirection.GongCopyBasicFields(newInstance)
	return newInstance
}

func (otherlistening *OtherListening) GongCopy() GongstructIF {
	newInstance := new(OtherListening)
	otherlistening.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pedal *Pedal) GongCopy() GongstructIF {
	newInstance := new(Pedal)
	pedal.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pedaltuning *PedalTuning) GongCopy() GongstructIF {
	newInstance := new(PedalTuning)
	pedaltuning.GongCopyBasicFields(newInstance)
	return newInstance
}

func (perminute *PerMinute) GongCopy() GongstructIF {
	newInstance := new(PerMinute)
	perminute.GongCopyBasicFields(newInstance)
	return newInstance
}

func (percussion *Percussion) GongCopy() GongstructIF {
	newInstance := new(Percussion)
	percussion.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pitched *Pitched) GongCopy() GongstructIF {
	newInstance := new(Pitched)
	pitched.GongCopyBasicFields(newInstance)
	return newInstance
}

func (principalvoice *PrincipalVoice) GongCopy() GongstructIF {
	newInstance := new(PrincipalVoice)
	principalvoice.GongCopyBasicFields(newInstance)
	return newInstance
}

func (print *Print) GongCopy() GongstructIF {
	newInstance := new(Print)
	print.GongCopyBasicFields(newInstance)
	return newInstance
}

func (root *Root) GongCopy() GongstructIF {
	newInstance := new(Root)
	root.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rootstep *RootStep) GongCopy() GongstructIF {
	newInstance := new(RootStep)
	rootstep.GongCopyBasicFields(newInstance)
	return newInstance
}

func (scordatura *Scordatura) GongCopy() GongstructIF {
	newInstance := new(Scordatura)
	scordatura.GongCopyBasicFields(newInstance)
	return newInstance
}

func (sound *Sound) GongCopy() GongstructIF {
	newInstance := new(Sound)
	sound.GongCopyBasicFields(newInstance)
	return newInstance
}

func (staffdivide *StaffDivide) GongCopy() GongstructIF {
	newInstance := new(StaffDivide)
	staffdivide.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stick *Stick) GongCopy() GongstructIF {
	newInstance := new(Stick)
	stick.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stringmute *StringMute) GongCopy() GongstructIF {
	newInstance := new(StringMute)
	stringmute.GongCopyBasicFields(newInstance)
	return newInstance
}

func (styletext *StyleText) GongCopy() GongstructIF {
	newInstance := new(StyleText)
	styletext.GongCopyBasicFields(newInstance)
	return newInstance
}

func (swing *Swing) GongCopy() GongstructIF {
	newInstance := new(Swing)
	swing.GongCopyBasicFields(newInstance)
	return newInstance
}

func (sync *Sync) GongCopy() GongstructIF {
	newInstance := new(Sync)
	sync.GongCopyBasicFields(newInstance)
	return newInstance
}

func (timemodification *TimeModification) GongCopy() GongstructIF {
	newInstance := new(TimeModification)
	timemodification.GongCopyBasicFields(newInstance)
	return newInstance
}

func (timpani *Timpani) GongCopy() GongstructIF {
	newInstance := new(Timpani)
	timpani.GongCopyBasicFields(newInstance)
	return newInstance
}

func (virtualinstrument *VirtualInstrument) GongCopy() GongstructIF {
	newInstance := new(VirtualInstrument)
	virtualinstrument.GongCopyBasicFields(newInstance)
	return newInstance
}

func (wedge *Wedge) GongCopy() GongstructIF {
	newInstance := new(Wedge)
	wedge.GongCopyBasicFields(newInstance)
	return newInstance
}

func (wood *Wood) GongCopy() GongstructIF {
	newInstance := new(Wood)
	wood.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (accord *Accord) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, accord)
}

func (accordionregistration *AccordionRegistration) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, accordionregistration)
}

func (barre *Barre) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, barre)
}

func (bass *Bass) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, bass)
}

func (bassstep *BassStep) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, bassstep)
}

func (beatunittied *BeatUnitTied) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, beatunittied)
}

func (beater *Beater) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, beater)
}

func (bracket *Bracket) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, bracket)
}

func (dashes *Dashes) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, dashes)
}

func (degree *Degree) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, degree)
}

func (degreealter *DegreeAlter) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, degreealter)
}

func (degreetype *DegreeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, degreetype)
}

func (degreevalue *DegreeValue) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, degreevalue)
}

func (direction *Direction) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, direction)
}

func (directiontype *DirectionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, directiontype)
}

func (effect *Effect) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, effect)
}

func (feature *Feature) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, feature)
}

func (firstfret *FirstFret) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, firstfret)
}

func (frame *Frame) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, frame)
}

func (framenote *FrameNote) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, framenote)
}

func (glass *Glass) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, glass)
}

func (grouping *Grouping) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, grouping)
}

func (harmony *Harmony) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, harmony)
}

func (harmonyalter *HarmonyAlter) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, harmonyalter)
}

func (harppedals *HarpPedals) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, harppedals)
}

func (image *Image) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, image)
}

func (instrumentchange *InstrumentChange) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, instrumentchange)
}

func (inversion *Inversion) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, inversion)
}

func (kind *Kind) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, kind)
}

func (listening *Listening) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, listening)
}

func (measurenumbering *MeasureNumbering) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, measurenumbering)
}

func (membrane *Membrane) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, membrane)
}

func (metal *Metal) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, metal)
}

func (metronome *Metronome) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, metronome)
}

func (metronomebeam *MetronomeBeam) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, metronomebeam)
}

func (metronomenote *MetronomeNote) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, metronomenote)
}

func (metronometied *MetronomeTied) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, metronometied)
}

func (metronometuplet *MetronomeTuplet) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, metronometuplet)
}

func (numeral *Numeral) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, numeral)
}

func (numeralkey *NumeralKey) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, numeralkey)
}

func (numeralroot *NumeralRoot) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, numeralroot)
}

func (octaveshift *OctaveShift) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, octaveshift)
}

func (offset *Offset) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, offset)
}

func (otherdirection *OtherDirection) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, otherdirection)
}

func (otherlistening *OtherListening) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, otherlistening)
}

func (pedal *Pedal) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pedal)
}

func (pedaltuning *PedalTuning) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pedaltuning)
}

func (perminute *PerMinute) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, perminute)
}

func (percussion *Percussion) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, percussion)
}

func (pitched *Pitched) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pitched)
}

func (principalvoice *PrincipalVoice) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, principalvoice)
}

func (print *Print) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, print)
}

func (root *Root) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, root)
}

func (rootstep *RootStep) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rootstep)
}

func (scordatura *Scordatura) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, scordatura)
}

func (sound *Sound) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, sound)
}

func (staffdivide *StaffDivide) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, staffdivide)
}

func (stick *Stick) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stick)
}

func (stringmute *StringMute) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stringmute)
}

func (styletext *StyleText) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, styletext)
}

func (swing *Swing) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, swing)
}

func (sync *Sync) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, sync)
}

func (timemodification *TimeModification) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, timemodification)
}

func (timpani *Timpani) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, timpani)
}

func (virtualinstrument *VirtualInstrument) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, virtualinstrument)
}

func (wedge *Wedge) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, wedge)
}

func (wood *Wood) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, wood)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	computeCommitsForType(
		stage,
		stage.Accords,
		stage.Accord_stagedOrder,
		stage.Accords_reference,
		&stage.Accords_referenceOrder,
		stage.Accords_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AccordionRegistrations,
		stage.AccordionRegistration_stagedOrder,
		stage.AccordionRegistrations_reference,
		&stage.AccordionRegistrations_referenceOrder,
		stage.AccordionRegistrations_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Barres,
		stage.Barre_stagedOrder,
		stage.Barres_reference,
		&stage.Barres_referenceOrder,
		stage.Barres_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Basss,
		stage.Bass_stagedOrder,
		stage.Basss_reference,
		&stage.Basss_referenceOrder,
		stage.Basss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.BassSteps,
		stage.BassStep_stagedOrder,
		stage.BassSteps_reference,
		&stage.BassSteps_referenceOrder,
		stage.BassSteps_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.BeatUnitTieds,
		stage.BeatUnitTied_stagedOrder,
		stage.BeatUnitTieds_reference,
		&stage.BeatUnitTieds_referenceOrder,
		stage.BeatUnitTieds_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Beaters,
		stage.Beater_stagedOrder,
		stage.Beaters_reference,
		&stage.Beaters_referenceOrder,
		stage.Beaters_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Brackets,
		stage.Bracket_stagedOrder,
		stage.Brackets_reference,
		&stage.Brackets_referenceOrder,
		stage.Brackets_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Dashess,
		stage.Dashes_stagedOrder,
		stage.Dashess_reference,
		&stage.Dashess_referenceOrder,
		stage.Dashess_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Degrees,
		stage.Degree_stagedOrder,
		stage.Degrees_reference,
		&stage.Degrees_referenceOrder,
		stage.Degrees_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DegreeAlters,
		stage.DegreeAlter_stagedOrder,
		stage.DegreeAlters_reference,
		&stage.DegreeAlters_referenceOrder,
		stage.DegreeAlters_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DegreeTypes,
		stage.DegreeType_stagedOrder,
		stage.DegreeTypes_reference,
		&stage.DegreeTypes_referenceOrder,
		stage.DegreeTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DegreeValues,
		stage.DegreeValue_stagedOrder,
		stage.DegreeValues_reference,
		&stage.DegreeValues_referenceOrder,
		stage.DegreeValues_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Directions,
		stage.Direction_stagedOrder,
		stage.Directions_reference,
		&stage.Directions_referenceOrder,
		stage.Directions_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DirectionTypes,
		stage.DirectionType_stagedOrder,
		stage.DirectionTypes_reference,
		&stage.DirectionTypes_referenceOrder,
		stage.DirectionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Effects,
		stage.Effect_stagedOrder,
		stage.Effects_reference,
		&stage.Effects_referenceOrder,
		stage.Effects_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Features,
		stage.Feature_stagedOrder,
		stage.Features_reference,
		&stage.Features_referenceOrder,
		stage.Features_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FirstFrets,
		stage.FirstFret_stagedOrder,
		stage.FirstFrets_reference,
		&stage.FirstFrets_referenceOrder,
		stage.FirstFrets_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Frames,
		stage.Frame_stagedOrder,
		stage.Frames_reference,
		&stage.Frames_referenceOrder,
		stage.Frames_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FrameNotes,
		stage.FrameNote_stagedOrder,
		stage.FrameNotes_reference,
		&stage.FrameNotes_referenceOrder,
		stage.FrameNotes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Glasss,
		stage.Glass_stagedOrder,
		stage.Glasss_reference,
		&stage.Glasss_referenceOrder,
		stage.Glasss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Groupings,
		stage.Grouping_stagedOrder,
		stage.Groupings_reference,
		&stage.Groupings_referenceOrder,
		stage.Groupings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Harmonys,
		stage.Harmony_stagedOrder,
		stage.Harmonys_reference,
		&stage.Harmonys_referenceOrder,
		stage.Harmonys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.HarmonyAlters,
		stage.HarmonyAlter_stagedOrder,
		stage.HarmonyAlters_reference,
		&stage.HarmonyAlters_referenceOrder,
		stage.HarmonyAlters_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.HarpPedalss,
		stage.HarpPedals_stagedOrder,
		stage.HarpPedalss_reference,
		&stage.HarpPedalss_referenceOrder,
		stage.HarpPedalss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Images,
		stage.Image_stagedOrder,
		stage.Images_reference,
		&stage.Images_referenceOrder,
		stage.Images_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.InstrumentChanges,
		stage.InstrumentChange_stagedOrder,
		stage.InstrumentChanges_reference,
		&stage.InstrumentChanges_referenceOrder,
		stage.InstrumentChanges_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Inversions,
		stage.Inversion_stagedOrder,
		stage.Inversions_reference,
		&stage.Inversions_referenceOrder,
		stage.Inversions_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Kinds,
		stage.Kind_stagedOrder,
		stage.Kinds_reference,
		&stage.Kinds_referenceOrder,
		stage.Kinds_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Listenings,
		stage.Listening_stagedOrder,
		stage.Listenings_reference,
		&stage.Listenings_referenceOrder,
		stage.Listenings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MeasureNumberings,
		stage.MeasureNumbering_stagedOrder,
		stage.MeasureNumberings_reference,
		&stage.MeasureNumberings_referenceOrder,
		stage.MeasureNumberings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Membranes,
		stage.Membrane_stagedOrder,
		stage.Membranes_reference,
		&stage.Membranes_referenceOrder,
		stage.Membranes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Metals,
		stage.Metal_stagedOrder,
		stage.Metals_reference,
		&stage.Metals_referenceOrder,
		stage.Metals_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Metronomes,
		stage.Metronome_stagedOrder,
		stage.Metronomes_reference,
		&stage.Metronomes_referenceOrder,
		stage.Metronomes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MetronomeBeams,
		stage.MetronomeBeam_stagedOrder,
		stage.MetronomeBeams_reference,
		&stage.MetronomeBeams_referenceOrder,
		stage.MetronomeBeams_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MetronomeNotes,
		stage.MetronomeNote_stagedOrder,
		stage.MetronomeNotes_reference,
		&stage.MetronomeNotes_referenceOrder,
		stage.MetronomeNotes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MetronomeTieds,
		stage.MetronomeTied_stagedOrder,
		stage.MetronomeTieds_reference,
		&stage.MetronomeTieds_referenceOrder,
		stage.MetronomeTieds_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MetronomeTuplets,
		stage.MetronomeTuplet_stagedOrder,
		stage.MetronomeTuplets_reference,
		&stage.MetronomeTuplets_referenceOrder,
		stage.MetronomeTuplets_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Numerals,
		stage.Numeral_stagedOrder,
		stage.Numerals_reference,
		&stage.Numerals_referenceOrder,
		stage.Numerals_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NumeralKeys,
		stage.NumeralKey_stagedOrder,
		stage.NumeralKeys_reference,
		&stage.NumeralKeys_referenceOrder,
		stage.NumeralKeys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NumeralRoots,
		stage.NumeralRoot_stagedOrder,
		stage.NumeralRoots_reference,
		&stage.NumeralRoots_referenceOrder,
		stage.NumeralRoots_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.OctaveShifts,
		stage.OctaveShift_stagedOrder,
		stage.OctaveShifts_reference,
		&stage.OctaveShifts_referenceOrder,
		stage.OctaveShifts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Offsets,
		stage.Offset_stagedOrder,
		stage.Offsets_reference,
		&stage.Offsets_referenceOrder,
		stage.Offsets_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.OtherDirections,
		stage.OtherDirection_stagedOrder,
		stage.OtherDirections_reference,
		&stage.OtherDirections_referenceOrder,
		stage.OtherDirections_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.OtherListenings,
		stage.OtherListening_stagedOrder,
		stage.OtherListenings_reference,
		&stage.OtherListenings_referenceOrder,
		stage.OtherListenings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Pedals,
		stage.Pedal_stagedOrder,
		stage.Pedals_reference,
		&stage.Pedals_referenceOrder,
		stage.Pedals_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PedalTunings,
		stage.PedalTuning_stagedOrder,
		stage.PedalTunings_reference,
		&stage.PedalTunings_referenceOrder,
		stage.PedalTunings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PerMinutes,
		stage.PerMinute_stagedOrder,
		stage.PerMinutes_reference,
		&stage.PerMinutes_referenceOrder,
		stage.PerMinutes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Percussions,
		stage.Percussion_stagedOrder,
		stage.Percussions_reference,
		&stage.Percussions_referenceOrder,
		stage.Percussions_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Pitcheds,
		stage.Pitched_stagedOrder,
		stage.Pitcheds_reference,
		&stage.Pitcheds_referenceOrder,
		stage.Pitcheds_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PrincipalVoices,
		stage.PrincipalVoice_stagedOrder,
		stage.PrincipalVoices_reference,
		&stage.PrincipalVoices_referenceOrder,
		stage.PrincipalVoices_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Prints,
		stage.Print_stagedOrder,
		stage.Prints_reference,
		&stage.Prints_referenceOrder,
		stage.Prints_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Roots,
		stage.Root_stagedOrder,
		stage.Roots_reference,
		&stage.Roots_referenceOrder,
		stage.Roots_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RootSteps,
		stage.RootStep_stagedOrder,
		stage.RootSteps_reference,
		&stage.RootSteps_referenceOrder,
		stage.RootSteps_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Scordaturas,
		stage.Scordatura_stagedOrder,
		stage.Scordaturas_reference,
		&stage.Scordaturas_referenceOrder,
		stage.Scordaturas_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Sounds,
		stage.Sound_stagedOrder,
		stage.Sounds_reference,
		&stage.Sounds_referenceOrder,
		stage.Sounds_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StaffDivides,
		stage.StaffDivide_stagedOrder,
		stage.StaffDivides_reference,
		&stage.StaffDivides_referenceOrder,
		stage.StaffDivides_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Sticks,
		stage.Stick_stagedOrder,
		stage.Sticks_reference,
		&stage.Sticks_referenceOrder,
		stage.Sticks_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StringMutes,
		stage.StringMute_stagedOrder,
		stage.StringMutes_reference,
		&stage.StringMutes_referenceOrder,
		stage.StringMutes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StyleTexts,
		stage.StyleText_stagedOrder,
		stage.StyleTexts_reference,
		&stage.StyleTexts_referenceOrder,
		stage.StyleTexts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Swings,
		stage.Swing_stagedOrder,
		stage.Swings_reference,
		&stage.Swings_referenceOrder,
		stage.Swings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Syncs,
		stage.Sync_stagedOrder,
		stage.Syncs_reference,
		&stage.Syncs_referenceOrder,
		stage.Syncs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TimeModifications,
		stage.TimeModification_stagedOrder,
		stage.TimeModifications_reference,
		&stage.TimeModifications_referenceOrder,
		stage.TimeModifications_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Timpanis,
		stage.Timpani_stagedOrder,
		stage.Timpanis_reference,
		&stage.Timpanis_referenceOrder,
		stage.Timpanis_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.VirtualInstruments,
		stage.VirtualInstrument_stagedOrder,
		stage.VirtualInstruments_reference,
		&stage.VirtualInstruments_referenceOrder,
		stage.VirtualInstruments_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Wedges,
		stage.Wedge_stagedOrder,
		stage.Wedges_reference,
		&stage.Wedges_referenceOrder,
		stage.Wedges_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Woods,
		stage.Wood_stagedOrder,
		stage.Woods_reference,
		&stage.Woods_referenceOrder,
		stage.Woods_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.AppendForwardCommit(forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.AppendBackwardCommit(backwardCommit)
		stage.SetModified(true)
	} else {
		stage.SetModified(false)
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.Accords, &stage.Accords_reference, &stage.Accords_referenceOrder, &stage.Accords_instance)

	__gong__computeReferencePass1(stage, stage.AccordionRegistrations, &stage.AccordionRegistrations_reference, &stage.AccordionRegistrations_referenceOrder, &stage.AccordionRegistrations_instance)

	__gong__computeReferencePass1(stage, stage.Barres, &stage.Barres_reference, &stage.Barres_referenceOrder, &stage.Barres_instance)

	__gong__computeReferencePass1(stage, stage.Basss, &stage.Basss_reference, &stage.Basss_referenceOrder, &stage.Basss_instance)

	__gong__computeReferencePass1(stage, stage.BassSteps, &stage.BassSteps_reference, &stage.BassSteps_referenceOrder, &stage.BassSteps_instance)

	__gong__computeReferencePass1(stage, stage.BeatUnitTieds, &stage.BeatUnitTieds_reference, &stage.BeatUnitTieds_referenceOrder, &stage.BeatUnitTieds_instance)

	__gong__computeReferencePass1(stage, stage.Beaters, &stage.Beaters_reference, &stage.Beaters_referenceOrder, &stage.Beaters_instance)

	__gong__computeReferencePass1(stage, stage.Brackets, &stage.Brackets_reference, &stage.Brackets_referenceOrder, &stage.Brackets_instance)

	__gong__computeReferencePass1(stage, stage.Dashess, &stage.Dashess_reference, &stage.Dashess_referenceOrder, &stage.Dashess_instance)

	__gong__computeReferencePass1(stage, stage.Degrees, &stage.Degrees_reference, &stage.Degrees_referenceOrder, &stage.Degrees_instance)

	__gong__computeReferencePass1(stage, stage.DegreeAlters, &stage.DegreeAlters_reference, &stage.DegreeAlters_referenceOrder, &stage.DegreeAlters_instance)

	__gong__computeReferencePass1(stage, stage.DegreeTypes, &stage.DegreeTypes_reference, &stage.DegreeTypes_referenceOrder, &stage.DegreeTypes_instance)

	__gong__computeReferencePass1(stage, stage.DegreeValues, &stage.DegreeValues_reference, &stage.DegreeValues_referenceOrder, &stage.DegreeValues_instance)

	__gong__computeReferencePass1(stage, stage.Directions, &stage.Directions_reference, &stage.Directions_referenceOrder, &stage.Directions_instance)

	__gong__computeReferencePass1(stage, stage.DirectionTypes, &stage.DirectionTypes_reference, &stage.DirectionTypes_referenceOrder, &stage.DirectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.Effects, &stage.Effects_reference, &stage.Effects_referenceOrder, &stage.Effects_instance)

	__gong__computeReferencePass1(stage, stage.Features, &stage.Features_reference, &stage.Features_referenceOrder, &stage.Features_instance)

	__gong__computeReferencePass1(stage, stage.FirstFrets, &stage.FirstFrets_reference, &stage.FirstFrets_referenceOrder, &stage.FirstFrets_instance)

	__gong__computeReferencePass1(stage, stage.Frames, &stage.Frames_reference, &stage.Frames_referenceOrder, &stage.Frames_instance)

	__gong__computeReferencePass1(stage, stage.FrameNotes, &stage.FrameNotes_reference, &stage.FrameNotes_referenceOrder, &stage.FrameNotes_instance)

	__gong__computeReferencePass1(stage, stage.Glasss, &stage.Glasss_reference, &stage.Glasss_referenceOrder, &stage.Glasss_instance)

	__gong__computeReferencePass1(stage, stage.Groupings, &stage.Groupings_reference, &stage.Groupings_referenceOrder, &stage.Groupings_instance)

	__gong__computeReferencePass1(stage, stage.Harmonys, &stage.Harmonys_reference, &stage.Harmonys_referenceOrder, &stage.Harmonys_instance)

	__gong__computeReferencePass1(stage, stage.HarmonyAlters, &stage.HarmonyAlters_reference, &stage.HarmonyAlters_referenceOrder, &stage.HarmonyAlters_instance)

	__gong__computeReferencePass1(stage, stage.HarpPedalss, &stage.HarpPedalss_reference, &stage.HarpPedalss_referenceOrder, &stage.HarpPedalss_instance)

	__gong__computeReferencePass1(stage, stage.Images, &stage.Images_reference, &stage.Images_referenceOrder, &stage.Images_instance)

	__gong__computeReferencePass1(stage, stage.InstrumentChanges, &stage.InstrumentChanges_reference, &stage.InstrumentChanges_referenceOrder, &stage.InstrumentChanges_instance)

	__gong__computeReferencePass1(stage, stage.Inversions, &stage.Inversions_reference, &stage.Inversions_referenceOrder, &stage.Inversions_instance)

	__gong__computeReferencePass1(stage, stage.Kinds, &stage.Kinds_reference, &stage.Kinds_referenceOrder, &stage.Kinds_instance)

	__gong__computeReferencePass1(stage, stage.Listenings, &stage.Listenings_reference, &stage.Listenings_referenceOrder, &stage.Listenings_instance)

	__gong__computeReferencePass1(stage, stage.MeasureNumberings, &stage.MeasureNumberings_reference, &stage.MeasureNumberings_referenceOrder, &stage.MeasureNumberings_instance)

	__gong__computeReferencePass1(stage, stage.Membranes, &stage.Membranes_reference, &stage.Membranes_referenceOrder, &stage.Membranes_instance)

	__gong__computeReferencePass1(stage, stage.Metals, &stage.Metals_reference, &stage.Metals_referenceOrder, &stage.Metals_instance)

	__gong__computeReferencePass1(stage, stage.Metronomes, &stage.Metronomes_reference, &stage.Metronomes_referenceOrder, &stage.Metronomes_instance)

	__gong__computeReferencePass1(stage, stage.MetronomeBeams, &stage.MetronomeBeams_reference, &stage.MetronomeBeams_referenceOrder, &stage.MetronomeBeams_instance)

	__gong__computeReferencePass1(stage, stage.MetronomeNotes, &stage.MetronomeNotes_reference, &stage.MetronomeNotes_referenceOrder, &stage.MetronomeNotes_instance)

	__gong__computeReferencePass1(stage, stage.MetronomeTieds, &stage.MetronomeTieds_reference, &stage.MetronomeTieds_referenceOrder, &stage.MetronomeTieds_instance)

	__gong__computeReferencePass1(stage, stage.MetronomeTuplets, &stage.MetronomeTuplets_reference, &stage.MetronomeTuplets_referenceOrder, &stage.MetronomeTuplets_instance)

	__gong__computeReferencePass1(stage, stage.Numerals, &stage.Numerals_reference, &stage.Numerals_referenceOrder, &stage.Numerals_instance)

	__gong__computeReferencePass1(stage, stage.NumeralKeys, &stage.NumeralKeys_reference, &stage.NumeralKeys_referenceOrder, &stage.NumeralKeys_instance)

	__gong__computeReferencePass1(stage, stage.NumeralRoots, &stage.NumeralRoots_reference, &stage.NumeralRoots_referenceOrder, &stage.NumeralRoots_instance)

	__gong__computeReferencePass1(stage, stage.OctaveShifts, &stage.OctaveShifts_reference, &stage.OctaveShifts_referenceOrder, &stage.OctaveShifts_instance)

	__gong__computeReferencePass1(stage, stage.Offsets, &stage.Offsets_reference, &stage.Offsets_referenceOrder, &stage.Offsets_instance)

	__gong__computeReferencePass1(stage, stage.OtherDirections, &stage.OtherDirections_reference, &stage.OtherDirections_referenceOrder, &stage.OtherDirections_instance)

	__gong__computeReferencePass1(stage, stage.OtherListenings, &stage.OtherListenings_reference, &stage.OtherListenings_referenceOrder, &stage.OtherListenings_instance)

	__gong__computeReferencePass1(stage, stage.Pedals, &stage.Pedals_reference, &stage.Pedals_referenceOrder, &stage.Pedals_instance)

	__gong__computeReferencePass1(stage, stage.PedalTunings, &stage.PedalTunings_reference, &stage.PedalTunings_referenceOrder, &stage.PedalTunings_instance)

	__gong__computeReferencePass1(stage, stage.PerMinutes, &stage.PerMinutes_reference, &stage.PerMinutes_referenceOrder, &stage.PerMinutes_instance)

	__gong__computeReferencePass1(stage, stage.Percussions, &stage.Percussions_reference, &stage.Percussions_referenceOrder, &stage.Percussions_instance)

	__gong__computeReferencePass1(stage, stage.Pitcheds, &stage.Pitcheds_reference, &stage.Pitcheds_referenceOrder, &stage.Pitcheds_instance)

	__gong__computeReferencePass1(stage, stage.PrincipalVoices, &stage.PrincipalVoices_reference, &stage.PrincipalVoices_referenceOrder, &stage.PrincipalVoices_instance)

	__gong__computeReferencePass1(stage, stage.Prints, &stage.Prints_reference, &stage.Prints_referenceOrder, &stage.Prints_instance)

	__gong__computeReferencePass1(stage, stage.Roots, &stage.Roots_reference, &stage.Roots_referenceOrder, &stage.Roots_instance)

	__gong__computeReferencePass1(stage, stage.RootSteps, &stage.RootSteps_reference, &stage.RootSteps_referenceOrder, &stage.RootSteps_instance)

	__gong__computeReferencePass1(stage, stage.Scordaturas, &stage.Scordaturas_reference, &stage.Scordaturas_referenceOrder, &stage.Scordaturas_instance)

	__gong__computeReferencePass1(stage, stage.Sounds, &stage.Sounds_reference, &stage.Sounds_referenceOrder, &stage.Sounds_instance)

	__gong__computeReferencePass1(stage, stage.StaffDivides, &stage.StaffDivides_reference, &stage.StaffDivides_referenceOrder, &stage.StaffDivides_instance)

	__gong__computeReferencePass1(stage, stage.Sticks, &stage.Sticks_reference, &stage.Sticks_referenceOrder, &stage.Sticks_instance)

	__gong__computeReferencePass1(stage, stage.StringMutes, &stage.StringMutes_reference, &stage.StringMutes_referenceOrder, &stage.StringMutes_instance)

	__gong__computeReferencePass1(stage, stage.StyleTexts, &stage.StyleTexts_reference, &stage.StyleTexts_referenceOrder, &stage.StyleTexts_instance)

	__gong__computeReferencePass1(stage, stage.Swings, &stage.Swings_reference, &stage.Swings_referenceOrder, &stage.Swings_instance)

	__gong__computeReferencePass1(stage, stage.Syncs, &stage.Syncs_reference, &stage.Syncs_referenceOrder, &stage.Syncs_instance)

	__gong__computeReferencePass1(stage, stage.TimeModifications, &stage.TimeModifications_reference, &stage.TimeModifications_referenceOrder, &stage.TimeModifications_instance)

	__gong__computeReferencePass1(stage, stage.Timpanis, &stage.Timpanis_reference, &stage.Timpanis_referenceOrder, &stage.Timpanis_instance)

	__gong__computeReferencePass1(stage, stage.VirtualInstruments, &stage.VirtualInstruments_reference, &stage.VirtualInstruments_referenceOrder, &stage.VirtualInstruments_instance)

	__gong__computeReferencePass1(stage, stage.Wedges, &stage.Wedges_reference, &stage.Wedges_referenceOrder, &stage.Wedges_instance)

	__gong__computeReferencePass1(stage, stage.Woods, &stage.Woods_reference, &stage.Woods_referenceOrder, &stage.Woods_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.Accords, stage.Accords_reference, stage)

	__gong__computeReferencePass2(stage.AccordionRegistrations, stage.AccordionRegistrations_reference, stage)

	__gong__computeReferencePass2(stage.Barres, stage.Barres_reference, stage)

	__gong__computeReferencePass2(stage.Basss, stage.Basss_reference, stage)

	__gong__computeReferencePass2(stage.BassSteps, stage.BassSteps_reference, stage)

	__gong__computeReferencePass2(stage.BeatUnitTieds, stage.BeatUnitTieds_reference, stage)

	__gong__computeReferencePass2(stage.Beaters, stage.Beaters_reference, stage)

	__gong__computeReferencePass2(stage.Brackets, stage.Brackets_reference, stage)

	__gong__computeReferencePass2(stage.Dashess, stage.Dashess_reference, stage)

	__gong__computeReferencePass2(stage.Degrees, stage.Degrees_reference, stage)

	__gong__computeReferencePass2(stage.DegreeAlters, stage.DegreeAlters_reference, stage)

	__gong__computeReferencePass2(stage.DegreeTypes, stage.DegreeTypes_reference, stage)

	__gong__computeReferencePass2(stage.DegreeValues, stage.DegreeValues_reference, stage)

	__gong__computeReferencePass2(stage.Directions, stage.Directions_reference, stage)

	__gong__computeReferencePass2(stage.DirectionTypes, stage.DirectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.Effects, stage.Effects_reference, stage)

	__gong__computeReferencePass2(stage.Features, stage.Features_reference, stage)

	__gong__computeReferencePass2(stage.FirstFrets, stage.FirstFrets_reference, stage)

	__gong__computeReferencePass2(stage.Frames, stage.Frames_reference, stage)

	__gong__computeReferencePass2(stage.FrameNotes, stage.FrameNotes_reference, stage)

	__gong__computeReferencePass2(stage.Glasss, stage.Glasss_reference, stage)

	__gong__computeReferencePass2(stage.Groupings, stage.Groupings_reference, stage)

	__gong__computeReferencePass2(stage.Harmonys, stage.Harmonys_reference, stage)

	__gong__computeReferencePass2(stage.HarmonyAlters, stage.HarmonyAlters_reference, stage)

	__gong__computeReferencePass2(stage.HarpPedalss, stage.HarpPedalss_reference, stage)

	__gong__computeReferencePass2(stage.Images, stage.Images_reference, stage)

	__gong__computeReferencePass2(stage.InstrumentChanges, stage.InstrumentChanges_reference, stage)

	__gong__computeReferencePass2(stage.Inversions, stage.Inversions_reference, stage)

	__gong__computeReferencePass2(stage.Kinds, stage.Kinds_reference, stage)

	__gong__computeReferencePass2(stage.Listenings, stage.Listenings_reference, stage)

	__gong__computeReferencePass2(stage.MeasureNumberings, stage.MeasureNumberings_reference, stage)

	__gong__computeReferencePass2(stage.Membranes, stage.Membranes_reference, stage)

	__gong__computeReferencePass2(stage.Metals, stage.Metals_reference, stage)

	__gong__computeReferencePass2(stage.Metronomes, stage.Metronomes_reference, stage)

	__gong__computeReferencePass2(stage.MetronomeBeams, stage.MetronomeBeams_reference, stage)

	__gong__computeReferencePass2(stage.MetronomeNotes, stage.MetronomeNotes_reference, stage)

	__gong__computeReferencePass2(stage.MetronomeTieds, stage.MetronomeTieds_reference, stage)

	__gong__computeReferencePass2(stage.MetronomeTuplets, stage.MetronomeTuplets_reference, stage)

	__gong__computeReferencePass2(stage.Numerals, stage.Numerals_reference, stage)

	__gong__computeReferencePass2(stage.NumeralKeys, stage.NumeralKeys_reference, stage)

	__gong__computeReferencePass2(stage.NumeralRoots, stage.NumeralRoots_reference, stage)

	__gong__computeReferencePass2(stage.OctaveShifts, stage.OctaveShifts_reference, stage)

	__gong__computeReferencePass2(stage.Offsets, stage.Offsets_reference, stage)

	__gong__computeReferencePass2(stage.OtherDirections, stage.OtherDirections_reference, stage)

	__gong__computeReferencePass2(stage.OtherListenings, stage.OtherListenings_reference, stage)

	__gong__computeReferencePass2(stage.Pedals, stage.Pedals_reference, stage)

	__gong__computeReferencePass2(stage.PedalTunings, stage.PedalTunings_reference, stage)

	__gong__computeReferencePass2(stage.PerMinutes, stage.PerMinutes_reference, stage)

	__gong__computeReferencePass2(stage.Percussions, stage.Percussions_reference, stage)

	__gong__computeReferencePass2(stage.Pitcheds, stage.Pitcheds_reference, stage)

	__gong__computeReferencePass2(stage.PrincipalVoices, stage.PrincipalVoices_reference, stage)

	__gong__computeReferencePass2(stage.Prints, stage.Prints_reference, stage)

	__gong__computeReferencePass2(stage.Roots, stage.Roots_reference, stage)

	__gong__computeReferencePass2(stage.RootSteps, stage.RootSteps_reference, stage)

	__gong__computeReferencePass2(stage.Scordaturas, stage.Scordaturas_reference, stage)

	__gong__computeReferencePass2(stage.Sounds, stage.Sounds_reference, stage)

	__gong__computeReferencePass2(stage.StaffDivides, stage.StaffDivides_reference, stage)

	__gong__computeReferencePass2(stage.Sticks, stage.Sticks_reference, stage)

	__gong__computeReferencePass2(stage.StringMutes, stage.StringMutes_reference, stage)

	__gong__computeReferencePass2(stage.StyleTexts, stage.StyleTexts_reference, stage)

	__gong__computeReferencePass2(stage.Swings, stage.Swings_reference, stage)

	__gong__computeReferencePass2(stage.Syncs, stage.Syncs_reference, stage)

	__gong__computeReferencePass2(stage.TimeModifications, stage.TimeModifications_reference, stage)

	__gong__computeReferencePass2(stage.Timpanis, stage.Timpanis_reference, stage)

	__gong__computeReferencePass2(stage.VirtualInstruments, stage.VirtualInstruments_reference, stage)

	__gong__computeReferencePass2(stage.Wedges, stage.Wedges_reference, stage)

	__gong__computeReferencePass2(stage.Woods, stage.Woods_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (accord *Accord) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Accord_stagedOrder, stage.Accords_referenceOrder, accord, "Accord")
}

func (accordionregistration *AccordionRegistration) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AccordionRegistration_stagedOrder, stage.AccordionRegistrations_referenceOrder, accordionregistration, "AccordionRegistration")
}

func (barre *Barre) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Barre_stagedOrder, stage.Barres_referenceOrder, barre, "Barre")
}

func (bass *Bass) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Bass_stagedOrder, stage.Basss_referenceOrder, bass, "Bass")
}

func (bassstep *BassStep) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.BassStep_stagedOrder, stage.BassSteps_referenceOrder, bassstep, "BassStep")
}

func (beatunittied *BeatUnitTied) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.BeatUnitTied_stagedOrder, stage.BeatUnitTieds_referenceOrder, beatunittied, "BeatUnitTied")
}

func (beater *Beater) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Beater_stagedOrder, stage.Beaters_referenceOrder, beater, "Beater")
}

func (bracket *Bracket) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Bracket_stagedOrder, stage.Brackets_referenceOrder, bracket, "Bracket")
}

func (dashes *Dashes) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Dashes_stagedOrder, stage.Dashess_referenceOrder, dashes, "Dashes")
}

func (degree *Degree) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Degree_stagedOrder, stage.Degrees_referenceOrder, degree, "Degree")
}

func (degreealter *DegreeAlter) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DegreeAlter_stagedOrder, stage.DegreeAlters_referenceOrder, degreealter, "DegreeAlter")
}

func (degreetype *DegreeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DegreeType_stagedOrder, stage.DegreeTypes_referenceOrder, degreetype, "DegreeType")
}

func (degreevalue *DegreeValue) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DegreeValue_stagedOrder, stage.DegreeValues_referenceOrder, degreevalue, "DegreeValue")
}

func (direction *Direction) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Direction_stagedOrder, stage.Directions_referenceOrder, direction, "Direction")
}

func (directiontype *DirectionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DirectionType_stagedOrder, stage.DirectionTypes_referenceOrder, directiontype, "DirectionType")
}

func (effect *Effect) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Effect_stagedOrder, stage.Effects_referenceOrder, effect, "Effect")
}

func (feature *Feature) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Feature_stagedOrder, stage.Features_referenceOrder, feature, "Feature")
}

func (firstfret *FirstFret) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FirstFret_stagedOrder, stage.FirstFrets_referenceOrder, firstfret, "FirstFret")
}

func (frame *Frame) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Frame_stagedOrder, stage.Frames_referenceOrder, frame, "Frame")
}

func (framenote *FrameNote) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FrameNote_stagedOrder, stage.FrameNotes_referenceOrder, framenote, "FrameNote")
}

func (glass *Glass) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Glass_stagedOrder, stage.Glasss_referenceOrder, glass, "Glass")
}

func (grouping *Grouping) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Grouping_stagedOrder, stage.Groupings_referenceOrder, grouping, "Grouping")
}

func (harmony *Harmony) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Harmony_stagedOrder, stage.Harmonys_referenceOrder, harmony, "Harmony")
}

func (harmonyalter *HarmonyAlter) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HarmonyAlter_stagedOrder, stage.HarmonyAlters_referenceOrder, harmonyalter, "HarmonyAlter")
}

func (harppedals *HarpPedals) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HarpPedals_stagedOrder, stage.HarpPedalss_referenceOrder, harppedals, "HarpPedals")
}

func (image *Image) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Image_stagedOrder, stage.Images_referenceOrder, image, "Image")
}

func (instrumentchange *InstrumentChange) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.InstrumentChange_stagedOrder, stage.InstrumentChanges_referenceOrder, instrumentchange, "InstrumentChange")
}

func (inversion *Inversion) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Inversion_stagedOrder, stage.Inversions_referenceOrder, inversion, "Inversion")
}

func (kind *Kind) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Kind_stagedOrder, stage.Kinds_referenceOrder, kind, "Kind")
}

func (listening *Listening) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Listening_stagedOrder, stage.Listenings_referenceOrder, listening, "Listening")
}

func (measurenumbering *MeasureNumbering) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MeasureNumbering_stagedOrder, stage.MeasureNumberings_referenceOrder, measurenumbering, "MeasureNumbering")
}

func (membrane *Membrane) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Membrane_stagedOrder, stage.Membranes_referenceOrder, membrane, "Membrane")
}

func (metal *Metal) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Metal_stagedOrder, stage.Metals_referenceOrder, metal, "Metal")
}

func (metronome *Metronome) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Metronome_stagedOrder, stage.Metronomes_referenceOrder, metronome, "Metronome")
}

func (metronomebeam *MetronomeBeam) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MetronomeBeam_stagedOrder, stage.MetronomeBeams_referenceOrder, metronomebeam, "MetronomeBeam")
}

func (metronomenote *MetronomeNote) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MetronomeNote_stagedOrder, stage.MetronomeNotes_referenceOrder, metronomenote, "MetronomeNote")
}

func (metronometied *MetronomeTied) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MetronomeTied_stagedOrder, stage.MetronomeTieds_referenceOrder, metronometied, "MetronomeTied")
}

func (metronometuplet *MetronomeTuplet) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MetronomeTuplet_stagedOrder, stage.MetronomeTuplets_referenceOrder, metronometuplet, "MetronomeTuplet")
}

func (numeral *Numeral) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Numeral_stagedOrder, stage.Numerals_referenceOrder, numeral, "Numeral")
}

func (numeralkey *NumeralKey) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NumeralKey_stagedOrder, stage.NumeralKeys_referenceOrder, numeralkey, "NumeralKey")
}

func (numeralroot *NumeralRoot) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NumeralRoot_stagedOrder, stage.NumeralRoots_referenceOrder, numeralroot, "NumeralRoot")
}

func (octaveshift *OctaveShift) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OctaveShift_stagedOrder, stage.OctaveShifts_referenceOrder, octaveshift, "OctaveShift")
}

func (offset *Offset) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Offset_stagedOrder, stage.Offsets_referenceOrder, offset, "Offset")
}

func (otherdirection *OtherDirection) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OtherDirection_stagedOrder, stage.OtherDirections_referenceOrder, otherdirection, "OtherDirection")
}

func (otherlistening *OtherListening) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OtherListening_stagedOrder, stage.OtherListenings_referenceOrder, otherlistening, "OtherListening")
}

func (pedal *Pedal) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Pedal_stagedOrder, stage.Pedals_referenceOrder, pedal, "Pedal")
}

func (pedaltuning *PedalTuning) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PedalTuning_stagedOrder, stage.PedalTunings_referenceOrder, pedaltuning, "PedalTuning")
}

func (perminute *PerMinute) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PerMinute_stagedOrder, stage.PerMinutes_referenceOrder, perminute, "PerMinute")
}

func (percussion *Percussion) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Percussion_stagedOrder, stage.Percussions_referenceOrder, percussion, "Percussion")
}

func (pitched *Pitched) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Pitched_stagedOrder, stage.Pitcheds_referenceOrder, pitched, "Pitched")
}

func (principalvoice *PrincipalVoice) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PrincipalVoice_stagedOrder, stage.PrincipalVoices_referenceOrder, principalvoice, "PrincipalVoice")
}

func (print *Print) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Print_stagedOrder, stage.Prints_referenceOrder, print, "Print")
}

func (root *Root) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Root_stagedOrder, stage.Roots_referenceOrder, root, "Root")
}

func (rootstep *RootStep) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RootStep_stagedOrder, stage.RootSteps_referenceOrder, rootstep, "RootStep")
}

func (scordatura *Scordatura) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Scordatura_stagedOrder, stage.Scordaturas_referenceOrder, scordatura, "Scordatura")
}

func (sound *Sound) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Sound_stagedOrder, stage.Sounds_referenceOrder, sound, "Sound")
}

func (staffdivide *StaffDivide) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StaffDivide_stagedOrder, stage.StaffDivides_referenceOrder, staffdivide, "StaffDivide")
}

func (stick *Stick) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Stick_stagedOrder, stage.Sticks_referenceOrder, stick, "Stick")
}

func (stringmute *StringMute) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StringMute_stagedOrder, stage.StringMutes_referenceOrder, stringmute, "StringMute")
}

func (styletext *StyleText) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StyleText_stagedOrder, stage.StyleTexts_referenceOrder, styletext, "StyleText")
}

func (swing *Swing) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Swing_stagedOrder, stage.Swings_referenceOrder, swing, "Swing")
}

func (sync *Sync) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Sync_stagedOrder, stage.Syncs_referenceOrder, sync, "Sync")
}

func (timemodification *TimeModification) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TimeModification_stagedOrder, stage.TimeModifications_referenceOrder, timemodification, "TimeModification")
}

func (timpani *Timpani) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Timpani_stagedOrder, stage.Timpanis_referenceOrder, timpani, "Timpani")
}

func (virtualinstrument *VirtualInstrument) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VirtualInstrument_stagedOrder, stage.VirtualInstruments_referenceOrder, virtualinstrument, "VirtualInstrument")
}

func (wedge *Wedge) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Wedge_stagedOrder, stage.Wedges_referenceOrder, wedge, "Wedge")
}

func (wood *Wood) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Wood_stagedOrder, stage.Woods_referenceOrder, wood, "Wood")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (accord *Accord) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(accord, accord.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (accord *Accord) GongGetReferenceIdentifier(stage *Stage) string {
	return accord.GongGetIdentifier(stage)
}

func (accordionregistration *AccordionRegistration) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(accordionregistration, accordionregistration.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (accordionregistration *AccordionRegistration) GongGetReferenceIdentifier(stage *Stage) string {
	return accordionregistration.GongGetIdentifier(stage)
}

func (barre *Barre) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(barre, barre.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (barre *Barre) GongGetReferenceIdentifier(stage *Stage) string {
	return barre.GongGetIdentifier(stage)
}

func (bass *Bass) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(bass, bass.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bass *Bass) GongGetReferenceIdentifier(stage *Stage) string {
	return bass.GongGetIdentifier(stage)
}

func (bassstep *BassStep) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(bassstep, bassstep.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bassstep *BassStep) GongGetReferenceIdentifier(stage *Stage) string {
	return bassstep.GongGetIdentifier(stage)
}

func (beatunittied *BeatUnitTied) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(beatunittied, beatunittied.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (beatunittied *BeatUnitTied) GongGetReferenceIdentifier(stage *Stage) string {
	return beatunittied.GongGetIdentifier(stage)
}

func (beater *Beater) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(beater, beater.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (beater *Beater) GongGetReferenceIdentifier(stage *Stage) string {
	return beater.GongGetIdentifier(stage)
}

func (bracket *Bracket) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(bracket, bracket.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bracket *Bracket) GongGetReferenceIdentifier(stage *Stage) string {
	return bracket.GongGetIdentifier(stage)
}

func (dashes *Dashes) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(dashes, dashes.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dashes *Dashes) GongGetReferenceIdentifier(stage *Stage) string {
	return dashes.GongGetIdentifier(stage)
}

func (degree *Degree) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(degree, degree.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (degree *Degree) GongGetReferenceIdentifier(stage *Stage) string {
	return degree.GongGetIdentifier(stage)
}

func (degreealter *DegreeAlter) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(degreealter, degreealter.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (degreealter *DegreeAlter) GongGetReferenceIdentifier(stage *Stage) string {
	return degreealter.GongGetIdentifier(stage)
}

func (degreetype *DegreeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(degreetype, degreetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (degreetype *DegreeType) GongGetReferenceIdentifier(stage *Stage) string {
	return degreetype.GongGetIdentifier(stage)
}

func (degreevalue *DegreeValue) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(degreevalue, degreevalue.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (degreevalue *DegreeValue) GongGetReferenceIdentifier(stage *Stage) string {
	return degreevalue.GongGetIdentifier(stage)
}

func (direction *Direction) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(direction, direction.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (direction *Direction) GongGetReferenceIdentifier(stage *Stage) string {
	return direction.GongGetIdentifier(stage)
}

func (directiontype *DirectionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(directiontype, directiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (directiontype *DirectionType) GongGetReferenceIdentifier(stage *Stage) string {
	return directiontype.GongGetIdentifier(stage)
}

func (effect *Effect) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(effect, effect.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (effect *Effect) GongGetReferenceIdentifier(stage *Stage) string {
	return effect.GongGetIdentifier(stage)
}

func (feature *Feature) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(feature, feature.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (feature *Feature) GongGetReferenceIdentifier(stage *Stage) string {
	return feature.GongGetIdentifier(stage)
}

func (firstfret *FirstFret) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(firstfret, firstfret.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (firstfret *FirstFret) GongGetReferenceIdentifier(stage *Stage) string {
	return firstfret.GongGetIdentifier(stage)
}

func (frame *Frame) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(frame, frame.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (frame *Frame) GongGetReferenceIdentifier(stage *Stage) string {
	return frame.GongGetIdentifier(stage)
}

func (framenote *FrameNote) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(framenote, framenote.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (framenote *FrameNote) GongGetReferenceIdentifier(stage *Stage) string {
	return framenote.GongGetIdentifier(stage)
}

func (glass *Glass) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(glass, glass.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (glass *Glass) GongGetReferenceIdentifier(stage *Stage) string {
	return glass.GongGetIdentifier(stage)
}

func (grouping *Grouping) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(grouping, grouping.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (grouping *Grouping) GongGetReferenceIdentifier(stage *Stage) string {
	return grouping.GongGetIdentifier(stage)
}

func (harmony *Harmony) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(harmony, harmony.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (harmony *Harmony) GongGetReferenceIdentifier(stage *Stage) string {
	return harmony.GongGetIdentifier(stage)
}

func (harmonyalter *HarmonyAlter) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(harmonyalter, harmonyalter.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (harmonyalter *HarmonyAlter) GongGetReferenceIdentifier(stage *Stage) string {
	return harmonyalter.GongGetIdentifier(stage)
}

func (harppedals *HarpPedals) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(harppedals, harppedals.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (harppedals *HarpPedals) GongGetReferenceIdentifier(stage *Stage) string {
	return harppedals.GongGetIdentifier(stage)
}

func (image *Image) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(image, image.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (image *Image) GongGetReferenceIdentifier(stage *Stage) string {
	return image.GongGetIdentifier(stage)
}

func (instrumentchange *InstrumentChange) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(instrumentchange, instrumentchange.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (instrumentchange *InstrumentChange) GongGetReferenceIdentifier(stage *Stage) string {
	return instrumentchange.GongGetIdentifier(stage)
}

func (inversion *Inversion) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(inversion, inversion.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (inversion *Inversion) GongGetReferenceIdentifier(stage *Stage) string {
	return inversion.GongGetIdentifier(stage)
}

func (kind *Kind) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(kind, kind.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (kind *Kind) GongGetReferenceIdentifier(stage *Stage) string {
	return kind.GongGetIdentifier(stage)
}

func (listening *Listening) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(listening, listening.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (listening *Listening) GongGetReferenceIdentifier(stage *Stage) string {
	return listening.GongGetIdentifier(stage)
}

func (measurenumbering *MeasureNumbering) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(measurenumbering, measurenumbering.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (measurenumbering *MeasureNumbering) GongGetReferenceIdentifier(stage *Stage) string {
	return measurenumbering.GongGetIdentifier(stage)
}

func (membrane *Membrane) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(membrane, membrane.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (membrane *Membrane) GongGetReferenceIdentifier(stage *Stage) string {
	return membrane.GongGetIdentifier(stage)
}

func (metal *Metal) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(metal, metal.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metal *Metal) GongGetReferenceIdentifier(stage *Stage) string {
	return metal.GongGetIdentifier(stage)
}

func (metronome *Metronome) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(metronome, metronome.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metronome *Metronome) GongGetReferenceIdentifier(stage *Stage) string {
	return metronome.GongGetIdentifier(stage)
}

func (metronomebeam *MetronomeBeam) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(metronomebeam, metronomebeam.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metronomebeam *MetronomeBeam) GongGetReferenceIdentifier(stage *Stage) string {
	return metronomebeam.GongGetIdentifier(stage)
}

func (metronomenote *MetronomeNote) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(metronomenote, metronomenote.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metronomenote *MetronomeNote) GongGetReferenceIdentifier(stage *Stage) string {
	return metronomenote.GongGetIdentifier(stage)
}

func (metronometied *MetronomeTied) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(metronometied, metronometied.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metronometied *MetronomeTied) GongGetReferenceIdentifier(stage *Stage) string {
	return metronometied.GongGetIdentifier(stage)
}

func (metronometuplet *MetronomeTuplet) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(metronometuplet, metronometuplet.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metronometuplet *MetronomeTuplet) GongGetReferenceIdentifier(stage *Stage) string {
	return metronometuplet.GongGetIdentifier(stage)
}

func (numeral *Numeral) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(numeral, numeral.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (numeral *Numeral) GongGetReferenceIdentifier(stage *Stage) string {
	return numeral.GongGetIdentifier(stage)
}

func (numeralkey *NumeralKey) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(numeralkey, numeralkey.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (numeralkey *NumeralKey) GongGetReferenceIdentifier(stage *Stage) string {
	return numeralkey.GongGetIdentifier(stage)
}

func (numeralroot *NumeralRoot) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(numeralroot, numeralroot.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (numeralroot *NumeralRoot) GongGetReferenceIdentifier(stage *Stage) string {
	return numeralroot.GongGetIdentifier(stage)
}

func (octaveshift *OctaveShift) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(octaveshift, octaveshift.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (octaveshift *OctaveShift) GongGetReferenceIdentifier(stage *Stage) string {
	return octaveshift.GongGetIdentifier(stage)
}

func (offset *Offset) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(offset, offset.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (offset *Offset) GongGetReferenceIdentifier(stage *Stage) string {
	return offset.GongGetIdentifier(stage)
}

func (otherdirection *OtherDirection) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(otherdirection, otherdirection.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (otherdirection *OtherDirection) GongGetReferenceIdentifier(stage *Stage) string {
	return otherdirection.GongGetIdentifier(stage)
}

func (otherlistening *OtherListening) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(otherlistening, otherlistening.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (otherlistening *OtherListening) GongGetReferenceIdentifier(stage *Stage) string {
	return otherlistening.GongGetIdentifier(stage)
}

func (pedal *Pedal) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pedal, pedal.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pedal *Pedal) GongGetReferenceIdentifier(stage *Stage) string {
	return pedal.GongGetIdentifier(stage)
}

func (pedaltuning *PedalTuning) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pedaltuning, pedaltuning.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pedaltuning *PedalTuning) GongGetReferenceIdentifier(stage *Stage) string {
	return pedaltuning.GongGetIdentifier(stage)
}

func (perminute *PerMinute) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(perminute, perminute.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (perminute *PerMinute) GongGetReferenceIdentifier(stage *Stage) string {
	return perminute.GongGetIdentifier(stage)
}

func (percussion *Percussion) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(percussion, percussion.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (percussion *Percussion) GongGetReferenceIdentifier(stage *Stage) string {
	return percussion.GongGetIdentifier(stage)
}

func (pitched *Pitched) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pitched, pitched.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pitched *Pitched) GongGetReferenceIdentifier(stage *Stage) string {
	return pitched.GongGetIdentifier(stage)
}

func (principalvoice *PrincipalVoice) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(principalvoice, principalvoice.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (principalvoice *PrincipalVoice) GongGetReferenceIdentifier(stage *Stage) string {
	return principalvoice.GongGetIdentifier(stage)
}

func (print *Print) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(print, print.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (print *Print) GongGetReferenceIdentifier(stage *Stage) string {
	return print.GongGetIdentifier(stage)
}

func (root *Root) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(root, root.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (root *Root) GongGetReferenceIdentifier(stage *Stage) string {
	return root.GongGetIdentifier(stage)
}

func (rootstep *RootStep) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rootstep, rootstep.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rootstep *RootStep) GongGetReferenceIdentifier(stage *Stage) string {
	return rootstep.GongGetIdentifier(stage)
}

func (scordatura *Scordatura) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(scordatura, scordatura.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (scordatura *Scordatura) GongGetReferenceIdentifier(stage *Stage) string {
	return scordatura.GongGetIdentifier(stage)
}

func (sound *Sound) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(sound, sound.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (sound *Sound) GongGetReferenceIdentifier(stage *Stage) string {
	return sound.GongGetIdentifier(stage)
}

func (staffdivide *StaffDivide) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(staffdivide, staffdivide.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (staffdivide *StaffDivide) GongGetReferenceIdentifier(stage *Stage) string {
	return staffdivide.GongGetIdentifier(stage)
}

func (stick *Stick) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stick, stick.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stick *Stick) GongGetReferenceIdentifier(stage *Stage) string {
	return stick.GongGetIdentifier(stage)
}

func (stringmute *StringMute) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stringmute, stringmute.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stringmute *StringMute) GongGetReferenceIdentifier(stage *Stage) string {
	return stringmute.GongGetIdentifier(stage)
}

func (styletext *StyleText) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(styletext, styletext.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (styletext *StyleText) GongGetReferenceIdentifier(stage *Stage) string {
	return styletext.GongGetIdentifier(stage)
}

func (swing *Swing) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(swing, swing.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (swing *Swing) GongGetReferenceIdentifier(stage *Stage) string {
	return swing.GongGetIdentifier(stage)
}

func (sync *Sync) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(sync, sync.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (sync *Sync) GongGetReferenceIdentifier(stage *Stage) string {
	return sync.GongGetIdentifier(stage)
}

func (timemodification *TimeModification) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(timemodification, timemodification.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (timemodification *TimeModification) GongGetReferenceIdentifier(stage *Stage) string {
	return timemodification.GongGetIdentifier(stage)
}

func (timpani *Timpani) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(timpani, timpani.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (timpani *Timpani) GongGetReferenceIdentifier(stage *Stage) string {
	return timpani.GongGetIdentifier(stage)
}

func (virtualinstrument *VirtualInstrument) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(virtualinstrument, virtualinstrument.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (virtualinstrument *VirtualInstrument) GongGetReferenceIdentifier(stage *Stage) string {
	return virtualinstrument.GongGetIdentifier(stage)
}

func (wedge *Wedge) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(wedge, wedge.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (wedge *Wedge) GongGetReferenceIdentifier(stage *Stage) string {
	return wedge.GongGetIdentifier(stage)
}

func (wood *Wood) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(wood, wood.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (wood *Wood) GongGetReferenceIdentifier(stage *Stage) string {
	return wood.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (accord *Accord) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(accord.GongGetIdentifier(stage), "Accord", accord.Name)
}

func (accordionregistration *AccordionRegistration) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(accordionregistration.GongGetIdentifier(stage), "AccordionRegistration", accordionregistration.Name)
}

func (barre *Barre) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(barre.GongGetIdentifier(stage), "Barre", barre.Name)
}

func (bass *Bass) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(bass.GongGetIdentifier(stage), "Bass", bass.Name)
}

func (bassstep *BassStep) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(bassstep.GongGetIdentifier(stage), "BassStep", bassstep.Name)
}

func (beatunittied *BeatUnitTied) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(beatunittied.GongGetIdentifier(stage), "BeatUnitTied", beatunittied.Name)
}

func (beater *Beater) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(beater.GongGetIdentifier(stage), "Beater", beater.Name)
}

func (bracket *Bracket) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(bracket.GongGetIdentifier(stage), "Bracket", bracket.Name)
}

func (dashes *Dashes) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(dashes.GongGetIdentifier(stage), "Dashes", dashes.Name)
}

func (degree *Degree) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(degree.GongGetIdentifier(stage), "Degree", degree.Name)
}

func (degreealter *DegreeAlter) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(degreealter.GongGetIdentifier(stage), "DegreeAlter", degreealter.Name)
}

func (degreetype *DegreeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(degreetype.GongGetIdentifier(stage), "DegreeType", degreetype.Name)
}

func (degreevalue *DegreeValue) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(degreevalue.GongGetIdentifier(stage), "DegreeValue", degreevalue.Name)
}

func (direction *Direction) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(direction.GongGetIdentifier(stage), "Direction", direction.Name)
}

func (directiontype *DirectionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(directiontype.GongGetIdentifier(stage), "DirectionType", directiontype.Name)
}

func (effect *Effect) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(effect.GongGetIdentifier(stage), "Effect", effect.Name)
}

func (feature *Feature) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(feature.GongGetIdentifier(stage), "Feature", feature.Name)
}

func (firstfret *FirstFret) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(firstfret.GongGetIdentifier(stage), "FirstFret", firstfret.Name)
}

func (frame *Frame) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(frame.GongGetIdentifier(stage), "Frame", frame.Name)
}

func (framenote *FrameNote) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(framenote.GongGetIdentifier(stage), "FrameNote", framenote.Name)
}

func (glass *Glass) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(glass.GongGetIdentifier(stage), "Glass", glass.Name)
}

func (grouping *Grouping) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(grouping.GongGetIdentifier(stage), "Grouping", grouping.Name)
}

func (harmony *Harmony) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(harmony.GongGetIdentifier(stage), "Harmony", harmony.Name)
}

func (harmonyalter *HarmonyAlter) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(harmonyalter.GongGetIdentifier(stage), "HarmonyAlter", harmonyalter.Name)
}

func (harppedals *HarpPedals) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(harppedals.GongGetIdentifier(stage), "HarpPedals", harppedals.Name)
}

func (image *Image) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(image.GongGetIdentifier(stage), "Image", image.Name)
}

func (instrumentchange *InstrumentChange) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(instrumentchange.GongGetIdentifier(stage), "InstrumentChange", instrumentchange.Name)
}

func (inversion *Inversion) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(inversion.GongGetIdentifier(stage), "Inversion", inversion.Name)
}

func (kind *Kind) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(kind.GongGetIdentifier(stage), "Kind", kind.Name)
}

func (listening *Listening) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(listening.GongGetIdentifier(stage), "Listening", listening.Name)
}

func (measurenumbering *MeasureNumbering) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(measurenumbering.GongGetIdentifier(stage), "MeasureNumbering", measurenumbering.Name)
}

func (membrane *Membrane) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(membrane.GongGetIdentifier(stage), "Membrane", membrane.Name)
}

func (metal *Metal) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(metal.GongGetIdentifier(stage), "Metal", metal.Name)
}

func (metronome *Metronome) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(metronome.GongGetIdentifier(stage), "Metronome", metronome.Name)
}

func (metronomebeam *MetronomeBeam) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(metronomebeam.GongGetIdentifier(stage), "MetronomeBeam", metronomebeam.Name)
}

func (metronomenote *MetronomeNote) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(metronomenote.GongGetIdentifier(stage), "MetronomeNote", metronomenote.Name)
}

func (metronometied *MetronomeTied) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(metronometied.GongGetIdentifier(stage), "MetronomeTied", metronometied.Name)
}

func (metronometuplet *MetronomeTuplet) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(metronometuplet.GongGetIdentifier(stage), "MetronomeTuplet", metronometuplet.Name)
}

func (numeral *Numeral) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(numeral.GongGetIdentifier(stage), "Numeral", numeral.Name)
}

func (numeralkey *NumeralKey) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(numeralkey.GongGetIdentifier(stage), "NumeralKey", numeralkey.Name)
}

func (numeralroot *NumeralRoot) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(numeralroot.GongGetIdentifier(stage), "NumeralRoot", numeralroot.Name)
}

func (octaveshift *OctaveShift) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(octaveshift.GongGetIdentifier(stage), "OctaveShift", octaveshift.Name)
}

func (offset *Offset) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(offset.GongGetIdentifier(stage), "Offset", offset.Name)
}

func (otherdirection *OtherDirection) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(otherdirection.GongGetIdentifier(stage), "OtherDirection", otherdirection.Name)
}

func (otherlistening *OtherListening) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(otherlistening.GongGetIdentifier(stage), "OtherListening", otherlistening.Name)
}

func (pedal *Pedal) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pedal.GongGetIdentifier(stage), "Pedal", pedal.Name)
}

func (pedaltuning *PedalTuning) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pedaltuning.GongGetIdentifier(stage), "PedalTuning", pedaltuning.Name)
}

func (perminute *PerMinute) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(perminute.GongGetIdentifier(stage), "PerMinute", perminute.Name)
}

func (percussion *Percussion) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(percussion.GongGetIdentifier(stage), "Percussion", percussion.Name)
}

func (pitched *Pitched) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pitched.GongGetIdentifier(stage), "Pitched", pitched.Name)
}

func (principalvoice *PrincipalVoice) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(principalvoice.GongGetIdentifier(stage), "PrincipalVoice", principalvoice.Name)
}

func (print *Print) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(print.GongGetIdentifier(stage), "Print", print.Name)
}

func (root *Root) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(root.GongGetIdentifier(stage), "Root", root.Name)
}

func (rootstep *RootStep) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rootstep.GongGetIdentifier(stage), "RootStep", rootstep.Name)
}

func (scordatura *Scordatura) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(scordatura.GongGetIdentifier(stage), "Scordatura", scordatura.Name)
}

func (sound *Sound) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(sound.GongGetIdentifier(stage), "Sound", sound.Name)
}

func (staffdivide *StaffDivide) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(staffdivide.GongGetIdentifier(stage), "StaffDivide", staffdivide.Name)
}

func (stick *Stick) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stick.GongGetIdentifier(stage), "Stick", stick.Name)
}

func (stringmute *StringMute) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stringmute.GongGetIdentifier(stage), "StringMute", stringmute.Name)
}

func (styletext *StyleText) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(styletext.GongGetIdentifier(stage), "StyleText", styletext.Name)
}

func (swing *Swing) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(swing.GongGetIdentifier(stage), "Swing", swing.Name)
}

func (sync *Sync) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(sync.GongGetIdentifier(stage), "Sync", sync.Name)
}

func (timemodification *TimeModification) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(timemodification.GongGetIdentifier(stage), "TimeModification", timemodification.Name)
}

func (timpani *Timpani) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(timpani.GongGetIdentifier(stage), "Timpani", timpani.Name)
}

func (virtualinstrument *VirtualInstrument) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(virtualinstrument.GongGetIdentifier(stage), "VirtualInstrument", virtualinstrument.Name)
}

func (wedge *Wedge) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(wedge.GongGetIdentifier(stage), "Wedge", wedge.Name)
}

func (wood *Wood) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(wood.GongGetIdentifier(stage), "Wood", wood.Name)
}

// insertion point for unstaging
func (accord *Accord) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(accord.GongGetReferenceIdentifier(stage))
}

func (accordionregistration *AccordionRegistration) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(accordionregistration.GongGetReferenceIdentifier(stage))
}

func (barre *Barre) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(barre.GongGetReferenceIdentifier(stage))
}

func (bass *Bass) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(bass.GongGetReferenceIdentifier(stage))
}

func (bassstep *BassStep) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(bassstep.GongGetReferenceIdentifier(stage))
}

func (beatunittied *BeatUnitTied) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(beatunittied.GongGetReferenceIdentifier(stage))
}

func (beater *Beater) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(beater.GongGetReferenceIdentifier(stage))
}

func (bracket *Bracket) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(bracket.GongGetReferenceIdentifier(stage))
}

func (dashes *Dashes) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(dashes.GongGetReferenceIdentifier(stage))
}

func (degree *Degree) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(degree.GongGetReferenceIdentifier(stage))
}

func (degreealter *DegreeAlter) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(degreealter.GongGetReferenceIdentifier(stage))
}

func (degreetype *DegreeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(degreetype.GongGetReferenceIdentifier(stage))
}

func (degreevalue *DegreeValue) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(degreevalue.GongGetReferenceIdentifier(stage))
}

func (direction *Direction) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(direction.GongGetReferenceIdentifier(stage))
}

func (directiontype *DirectionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(directiontype.GongGetReferenceIdentifier(stage))
}

func (effect *Effect) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(effect.GongGetReferenceIdentifier(stage))
}

func (feature *Feature) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(feature.GongGetReferenceIdentifier(stage))
}

func (firstfret *FirstFret) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(firstfret.GongGetReferenceIdentifier(stage))
}

func (frame *Frame) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(frame.GongGetReferenceIdentifier(stage))
}

func (framenote *FrameNote) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(framenote.GongGetReferenceIdentifier(stage))
}

func (glass *Glass) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(glass.GongGetReferenceIdentifier(stage))
}

func (grouping *Grouping) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(grouping.GongGetReferenceIdentifier(stage))
}

func (harmony *Harmony) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(harmony.GongGetReferenceIdentifier(stage))
}

func (harmonyalter *HarmonyAlter) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(harmonyalter.GongGetReferenceIdentifier(stage))
}

func (harppedals *HarpPedals) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(harppedals.GongGetReferenceIdentifier(stage))
}

func (image *Image) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(image.GongGetReferenceIdentifier(stage))
}

func (instrumentchange *InstrumentChange) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(instrumentchange.GongGetReferenceIdentifier(stage))
}

func (inversion *Inversion) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(inversion.GongGetReferenceIdentifier(stage))
}

func (kind *Kind) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(kind.GongGetReferenceIdentifier(stage))
}

func (listening *Listening) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(listening.GongGetReferenceIdentifier(stage))
}

func (measurenumbering *MeasureNumbering) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(measurenumbering.GongGetReferenceIdentifier(stage))
}

func (membrane *Membrane) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(membrane.GongGetReferenceIdentifier(stage))
}

func (metal *Metal) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(metal.GongGetReferenceIdentifier(stage))
}

func (metronome *Metronome) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(metronome.GongGetReferenceIdentifier(stage))
}

func (metronomebeam *MetronomeBeam) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(metronomebeam.GongGetReferenceIdentifier(stage))
}

func (metronomenote *MetronomeNote) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(metronomenote.GongGetReferenceIdentifier(stage))
}

func (metronometied *MetronomeTied) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(metronometied.GongGetReferenceIdentifier(stage))
}

func (metronometuplet *MetronomeTuplet) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(metronometuplet.GongGetReferenceIdentifier(stage))
}

func (numeral *Numeral) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(numeral.GongGetReferenceIdentifier(stage))
}

func (numeralkey *NumeralKey) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(numeralkey.GongGetReferenceIdentifier(stage))
}

func (numeralroot *NumeralRoot) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(numeralroot.GongGetReferenceIdentifier(stage))
}

func (octaveshift *OctaveShift) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(octaveshift.GongGetReferenceIdentifier(stage))
}

func (offset *Offset) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(offset.GongGetReferenceIdentifier(stage))
}

func (otherdirection *OtherDirection) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(otherdirection.GongGetReferenceIdentifier(stage))
}

func (otherlistening *OtherListening) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(otherlistening.GongGetReferenceIdentifier(stage))
}

func (pedal *Pedal) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pedal.GongGetReferenceIdentifier(stage))
}

func (pedaltuning *PedalTuning) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pedaltuning.GongGetReferenceIdentifier(stage))
}

func (perminute *PerMinute) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(perminute.GongGetReferenceIdentifier(stage))
}

func (percussion *Percussion) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(percussion.GongGetReferenceIdentifier(stage))
}

func (pitched *Pitched) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pitched.GongGetReferenceIdentifier(stage))
}

func (principalvoice *PrincipalVoice) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(principalvoice.GongGetReferenceIdentifier(stage))
}

func (print *Print) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(print.GongGetReferenceIdentifier(stage))
}

func (root *Root) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(root.GongGetReferenceIdentifier(stage))
}

func (rootstep *RootStep) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rootstep.GongGetReferenceIdentifier(stage))
}

func (scordatura *Scordatura) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(scordatura.GongGetReferenceIdentifier(stage))
}

func (sound *Sound) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(sound.GongGetReferenceIdentifier(stage))
}

func (staffdivide *StaffDivide) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(staffdivide.GongGetReferenceIdentifier(stage))
}

func (stick *Stick) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stick.GongGetReferenceIdentifier(stage))
}

func (stringmute *StringMute) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stringmute.GongGetReferenceIdentifier(stage))
}

func (styletext *StyleText) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(styletext.GongGetReferenceIdentifier(stage))
}

func (swing *Swing) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(swing.GongGetReferenceIdentifier(stage))
}

func (sync *Sync) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(sync.GongGetReferenceIdentifier(stage))
}

func (timemodification *TimeModification) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(timemodification.GongGetReferenceIdentifier(stage))
}

func (timpani *Timpani) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(timpani.GongGetReferenceIdentifier(stage))
}

func (virtualinstrument *VirtualInstrument) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(virtualinstrument.GongGetReferenceIdentifier(stage))
}

func (wedge *Wedge) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(wedge.GongGetReferenceIdentifier(stage))
}

func (wood *Wood) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(wood.GongGetReferenceIdentifier(stage))
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
