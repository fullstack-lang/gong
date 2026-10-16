// generated code - do not edit
package note

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
	// Compute reverse map for named struct Articulations
	// insertion point per field
	stage.Articulations_StrongAccent_reverseMap = make(map[*StrongAccent]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _strongaccent := range articulations.StrongAccent {
			stage.Articulations_StrongAccent_reverseMap[_strongaccent] = articulations
		}
	}
	stage.Articulations_Scoop_reverseMap = make(map[*EmptyLine]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _emptyline := range articulations.Scoop {
			stage.Articulations_Scoop_reverseMap[_emptyline] = articulations
		}
	}
	stage.Articulations_Plop_reverseMap = make(map[*EmptyLine]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _emptyline := range articulations.Plop {
			stage.Articulations_Plop_reverseMap[_emptyline] = articulations
		}
	}
	stage.Articulations_Doit_reverseMap = make(map[*EmptyLine]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _emptyline := range articulations.Doit {
			stage.Articulations_Doit_reverseMap[_emptyline] = articulations
		}
	}
	stage.Articulations_Falloff_reverseMap = make(map[*EmptyLine]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _emptyline := range articulations.Falloff {
			stage.Articulations_Falloff_reverseMap[_emptyline] = articulations
		}
	}
	stage.Articulations_BreathMark_reverseMap = make(map[*BreathMark]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _breathmark := range articulations.BreathMark {
			stage.Articulations_BreathMark_reverseMap[_breathmark] = articulations
		}
	}
	stage.Articulations_Caesura_reverseMap = make(map[*Caesura]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _caesura := range articulations.Caesura {
			stage.Articulations_Caesura_reverseMap[_caesura] = articulations
		}
	}
	stage.Articulations_OtherArticulation_reverseMap = make(map[*OtherPlacementText]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _otherplacementtext := range articulations.OtherArticulation {
			stage.Articulations_OtherArticulation_reverseMap[_otherplacementtext] = articulations
		}
	}

	// Compute reverse map for named struct FiguredBass
	// insertion point per field
	stage.FiguredBass_Figure_reverseMap = make(map[*Figure]*FiguredBass)
	for figuredbass := range stage.FiguredBasss {
		_ = figuredbass
		for _, _figure := range figuredbass.Figure {
			stage.FiguredBass_Figure_reverseMap[_figure] = figuredbass
		}
	}

	// Compute reverse map for named struct Listen
	// insertion point per field
	stage.Listen_Assess_reverseMap = make(map[*Assess]*Listen)
	for listen := range stage.Listens {
		_ = listen
		for _, _assess := range listen.Assess {
			stage.Listen_Assess_reverseMap[_assess] = listen
		}
	}
	stage.Listen_Wait_reverseMap = make(map[*Wait]*Listen)
	for listen := range stage.Listens {
		_ = listen
		for _, _wait := range listen.Wait {
			stage.Listen_Wait_reverseMap[_wait] = listen
		}
	}

	// Compute reverse map for named struct Lyric
	// insertion point per field
	stage.Lyric_Elision_reverseMap = make(map[*Elision]*Lyric)
	for lyric := range stage.Lyrics {
		_ = lyric
		for _, _elision := range lyric.Elision {
			stage.Lyric_Elision_reverseMap[_elision] = lyric
		}
	}

	// Compute reverse map for named struct Notations
	// insertion point per field
	stage.Notations_Tied_reverseMap = make(map[*Tied]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _tied := range notations.Tied {
			stage.Notations_Tied_reverseMap[_tied] = notations
		}
	}
	stage.Notations_Slur_reverseMap = make(map[*Slur]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _slur := range notations.Slur {
			stage.Notations_Slur_reverseMap[_slur] = notations
		}
	}
	stage.Notations_Tuplet_reverseMap = make(map[*Tuplet]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _tuplet := range notations.Tuplet {
			stage.Notations_Tuplet_reverseMap[_tuplet] = notations
		}
	}
	stage.Notations_Glissando_reverseMap = make(map[*Glissando]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _glissando := range notations.Glissando {
			stage.Notations_Glissando_reverseMap[_glissando] = notations
		}
	}
	stage.Notations_Slide_reverseMap = make(map[*Slide]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _slide := range notations.Slide {
			stage.Notations_Slide_reverseMap[_slide] = notations
		}
	}
	stage.Notations_Ornaments_reverseMap = make(map[*Ornaments]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _ornaments := range notations.Ornaments {
			stage.Notations_Ornaments_reverseMap[_ornaments] = notations
		}
	}
	stage.Notations_Technical_reverseMap = make(map[*Technical]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _technical := range notations.Technical {
			stage.Notations_Technical_reverseMap[_technical] = notations
		}
	}
	stage.Notations_Articulations_reverseMap = make(map[*Articulations]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _articulations := range notations.Articulations {
			stage.Notations_Articulations_reverseMap[_articulations] = notations
		}
	}
	stage.Notations_Arpeggiate_reverseMap = make(map[*Arpeggiate]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _arpeggiate := range notations.Arpeggiate {
			stage.Notations_Arpeggiate_reverseMap[_arpeggiate] = notations
		}
	}
	stage.Notations_NonArpeggiate_reverseMap = make(map[*NonArpeggiate]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _nonarpeggiate := range notations.NonArpeggiate {
			stage.Notations_NonArpeggiate_reverseMap[_nonarpeggiate] = notations
		}
	}
	stage.Notations_AccidentalMark_reverseMap = make(map[*AccidentalMark]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _accidentalmark := range notations.AccidentalMark {
			stage.Notations_AccidentalMark_reverseMap[_accidentalmark] = notations
		}
	}
	stage.Notations_OtherNotation_reverseMap = make(map[*OtherNotation]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _othernotation := range notations.OtherNotation {
			stage.Notations_OtherNotation_reverseMap[_othernotation] = notations
		}
	}

	// Compute reverse map for named struct Note
	// insertion point per field
	stage.Note_Tie_reverseMap = make(map[*Tie]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _tie := range note.Tie {
			stage.Note_Tie_reverseMap[_tie] = note
		}
	}
	stage.Note_Instrument_reverseMap = make(map[*Instrument]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _instrument := range note.Instrument {
			stage.Note_Instrument_reverseMap[_instrument] = note
		}
	}
	stage.Note_Beam_reverseMap = make(map[*Beam]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _beam := range note.Beam {
			stage.Note_Beam_reverseMap[_beam] = note
		}
	}
	stage.Note_Notations_reverseMap = make(map[*Notations]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _notations := range note.Notations {
			stage.Note_Notations_reverseMap[_notations] = note
		}
	}
	stage.Note_Lyric_reverseMap = make(map[*Lyric]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _lyric := range note.Lyric {
			stage.Note_Lyric_reverseMap[_lyric] = note
		}
	}

	// Compute reverse map for named struct Ornaments
	// insertion point per field
	stage.Ornaments_Mordent_reverseMap = make(map[*Mordent]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _mordent := range ornaments.Mordent {
			stage.Ornaments_Mordent_reverseMap[_mordent] = ornaments
		}
	}
	stage.Ornaments_InvertedMordent_reverseMap = make(map[*Mordent]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _mordent := range ornaments.InvertedMordent {
			stage.Ornaments_InvertedMordent_reverseMap[_mordent] = ornaments
		}
	}
	stage.Ornaments_Tremolo_reverseMap = make(map[*Tremolo]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _tremolo := range ornaments.Tremolo {
			stage.Ornaments_Tremolo_reverseMap[_tremolo] = ornaments
		}
	}
	stage.Ornaments_OtherOrnament_reverseMap = make(map[*OtherPlacementText]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _otherplacementtext := range ornaments.OtherOrnament {
			stage.Ornaments_OtherOrnament_reverseMap[_otherplacementtext] = ornaments
		}
	}
	stage.Ornaments_AccidentalMark_reverseMap = make(map[*AccidentalMark]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _accidentalmark := range ornaments.AccidentalMark {
			stage.Ornaments_AccidentalMark_reverseMap[_accidentalmark] = ornaments
		}
	}

	// Compute reverse map for named struct Technical
	// insertion point per field
	stage.Technical_Harmonic_reverseMap = make(map[*Harmonic]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _harmonic := range technical.Harmonic {
			stage.Technical_Harmonic_reverseMap[_harmonic] = technical
		}
	}
	stage.Technical_Pluck_reverseMap = make(map[*PlacementText]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _placementtext := range technical.Pluck {
			stage.Technical_Pluck_reverseMap[_placementtext] = technical
		}
	}
	stage.Technical_HammerOn_reverseMap = make(map[*HammerOnPullOff]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _hammeronpulloff := range technical.HammerOn {
			stage.Technical_HammerOn_reverseMap[_hammeronpulloff] = technical
		}
	}
	stage.Technical_PullOff_reverseMap = make(map[*HammerOnPullOff]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _hammeronpulloff := range technical.PullOff {
			stage.Technical_PullOff_reverseMap[_hammeronpulloff] = technical
		}
	}
	stage.Technical_Bend_reverseMap = make(map[*Bend]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _bend := range technical.Bend {
			stage.Technical_Bend_reverseMap[_bend] = technical
		}
	}
	stage.Technical_Tap_reverseMap = make(map[*Tap]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _tap := range technical.Tap {
			stage.Technical_Tap_reverseMap[_tap] = technical
		}
	}
	stage.Technical_Heel_reverseMap = make(map[*HeelToe]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _heeltoe := range technical.Heel {
			stage.Technical_Heel_reverseMap[_heeltoe] = technical
		}
	}
	stage.Technical_Toe_reverseMap = make(map[*HeelToe]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _heeltoe := range technical.Toe {
			stage.Technical_Toe_reverseMap[_heeltoe] = technical
		}
	}
	stage.Technical_Hole_reverseMap = make(map[*Hole]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _hole := range technical.Hole {
			stage.Technical_Hole_reverseMap[_hole] = technical
		}
	}
	stage.Technical_Arrow_reverseMap = make(map[*Arrow]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _arrow := range technical.Arrow {
			stage.Technical_Arrow_reverseMap[_arrow] = technical
		}
	}
	stage.Technical_Handbell_reverseMap = make(map[*Handbell]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _handbell := range technical.Handbell {
			stage.Technical_Handbell_reverseMap[_handbell] = technical
		}
	}
	stage.Technical_HarmonMute_reverseMap = make(map[*HarmonMute]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _harmonmute := range technical.HarmonMute {
			stage.Technical_HarmonMute_reverseMap[_harmonmute] = technical
		}
	}
	stage.Technical_OtherTechnical_reverseMap = make(map[*OtherPlacementText]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _otherplacementtext := range technical.OtherTechnical {
			stage.Technical_OtherTechnical_reverseMap[_otherplacementtext] = technical
		}
	}

	// Compute reverse map for named struct TupletPortion
	// insertion point per field
	stage.TupletPortion_TupletDot_reverseMap = make(map[*TupletDot]*TupletPortion)
	for tupletportion := range stage.TupletPortions {
		_ = tupletportion
		for _, _tupletdot := range tupletportion.TupletDot {
			stage.TupletPortion_TupletDot_reverseMap[_tupletdot] = tupletportion
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.Accidentals)

	res = __gong__appendInstances(res, stage.AccidentalMarks)

	res = __gong__appendInstances(res, stage.Arpeggiates)

	res = __gong__appendInstances(res, stage.Arrows)

	res = __gong__appendInstances(res, stage.Articulationss)

	res = __gong__appendInstances(res, stage.Assesss)

	res = __gong__appendInstances(res, stage.Backups)

	res = __gong__appendInstances(res, stage.Beams)

	res = __gong__appendInstances(res, stage.Bends)

	res = __gong__appendInstances(res, stage.BreathMarks)

	res = __gong__appendInstances(res, stage.Caesuras)

	res = __gong__appendInstances(res, stage.Elisions)

	res = __gong__appendInstances(res, stage.EmptyLines)

	res = __gong__appendInstances(res, stage.Extends)

	res = __gong__appendInstances(res, stage.Figures)

	res = __gong__appendInstances(res, stage.FiguredBasss)

	res = __gong__appendInstances(res, stage.Forwards)

	res = __gong__appendInstances(res, stage.Glissandos)

	res = __gong__appendInstances(res, stage.Graces)

	res = __gong__appendInstances(res, stage.HammerOnPullOffs)

	res = __gong__appendInstances(res, stage.Handbells)

	res = __gong__appendInstances(res, stage.HarmonCloseds)

	res = __gong__appendInstances(res, stage.HarmonMutes)

	res = __gong__appendInstances(res, stage.Harmonics)

	res = __gong__appendInstances(res, stage.HeelToes)

	res = __gong__appendInstances(res, stage.Holes)

	res = __gong__appendInstances(res, stage.HoleCloseds)

	res = __gong__appendInstances(res, stage.Instruments)

	res = __gong__appendInstances(res, stage.Listens)

	res = __gong__appendInstances(res, stage.Lyrics)

	res = __gong__appendInstances(res, stage.Mordents)

	res = __gong__appendInstances(res, stage.NonArpeggiates)

	res = __gong__appendInstances(res, stage.Notationss)

	res = __gong__appendInstances(res, stage.Notes)

	res = __gong__appendInstances(res, stage.NoteTypes)

	res = __gong__appendInstances(res, stage.Noteheads)

	res = __gong__appendInstances(res, stage.NoteheadTexts)

	res = __gong__appendInstances(res, stage.Ornamentss)

	res = __gong__appendInstances(res, stage.OtherNotations)

	res = __gong__appendInstances(res, stage.OtherPlacementTexts)

	res = __gong__appendInstances(res, stage.Pitchs)

	res = __gong__appendInstances(res, stage.PlacementTexts)

	res = __gong__appendInstances(res, stage.Releases)

	res = __gong__appendInstances(res, stage.Rests)

	res = __gong__appendInstances(res, stage.Slides)

	res = __gong__appendInstances(res, stage.Slurs)

	res = __gong__appendInstances(res, stage.Stems)

	res = __gong__appendInstances(res, stage.StrongAccents)

	res = __gong__appendInstances(res, stage.Taps)

	res = __gong__appendInstances(res, stage.Technicals)

	res = __gong__appendInstances(res, stage.TextElementDatas)

	res = __gong__appendInstances(res, stage.Ties)

	res = __gong__appendInstances(res, stage.Tieds)

	res = __gong__appendInstances(res, stage.Tremolos)

	res = __gong__appendInstances(res, stage.Tuplets)

	res = __gong__appendInstances(res, stage.TupletDots)

	res = __gong__appendInstances(res, stage.TupletNumbers)

	res = __gong__appendInstances(res, stage.TupletPortions)

	res = __gong__appendInstances(res, stage.TupletTypes)

	res = __gong__appendInstances(res, stage.Unpitcheds)

	res = __gong__appendInstances(res, stage.Waits)

	return
}

// insertion point per named struct
func (accidental *Accidental) GongCopy() GongstructIF {
	newInstance := new(Accidental)
	accidental.GongCopyBasicFields(newInstance)
	return newInstance
}

func (accidentalmark *AccidentalMark) GongCopy() GongstructIF {
	newInstance := new(AccidentalMark)
	accidentalmark.GongCopyBasicFields(newInstance)
	return newInstance
}

func (arpeggiate *Arpeggiate) GongCopy() GongstructIF {
	newInstance := new(Arpeggiate)
	arpeggiate.GongCopyBasicFields(newInstance)
	return newInstance
}

func (arrow *Arrow) GongCopy() GongstructIF {
	newInstance := new(Arrow)
	arrow.GongCopyBasicFields(newInstance)
	return newInstance
}

func (articulations *Articulations) GongCopy() GongstructIF {
	newInstance := new(Articulations)
	articulations.GongCopyBasicFields(newInstance)
	return newInstance
}

func (assess *Assess) GongCopy() GongstructIF {
	newInstance := new(Assess)
	assess.GongCopyBasicFields(newInstance)
	return newInstance
}

func (backup *Backup) GongCopy() GongstructIF {
	newInstance := new(Backup)
	backup.GongCopyBasicFields(newInstance)
	return newInstance
}

func (beam *Beam) GongCopy() GongstructIF {
	newInstance := new(Beam)
	beam.GongCopyBasicFields(newInstance)
	return newInstance
}

func (bend *Bend) GongCopy() GongstructIF {
	newInstance := new(Bend)
	bend.GongCopyBasicFields(newInstance)
	return newInstance
}

func (breathmark *BreathMark) GongCopy() GongstructIF {
	newInstance := new(BreathMark)
	breathmark.GongCopyBasicFields(newInstance)
	return newInstance
}

func (caesura *Caesura) GongCopy() GongstructIF {
	newInstance := new(Caesura)
	caesura.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elision *Elision) GongCopy() GongstructIF {
	newInstance := new(Elision)
	elision.GongCopyBasicFields(newInstance)
	return newInstance
}

func (emptyline *EmptyLine) GongCopy() GongstructIF {
	newInstance := new(EmptyLine)
	emptyline.GongCopyBasicFields(newInstance)
	return newInstance
}

func (extend *Extend) GongCopy() GongstructIF {
	newInstance := new(Extend)
	extend.GongCopyBasicFields(newInstance)
	return newInstance
}

func (figure *Figure) GongCopy() GongstructIF {
	newInstance := new(Figure)
	figure.GongCopyBasicFields(newInstance)
	return newInstance
}

func (figuredbass *FiguredBass) GongCopy() GongstructIF {
	newInstance := new(FiguredBass)
	figuredbass.GongCopyBasicFields(newInstance)
	return newInstance
}

func (forward *Forward) GongCopy() GongstructIF {
	newInstance := new(Forward)
	forward.GongCopyBasicFields(newInstance)
	return newInstance
}

func (glissando *Glissando) GongCopy() GongstructIF {
	newInstance := new(Glissando)
	glissando.GongCopyBasicFields(newInstance)
	return newInstance
}

func (grace *Grace) GongCopy() GongstructIF {
	newInstance := new(Grace)
	grace.GongCopyBasicFields(newInstance)
	return newInstance
}

func (hammeronpulloff *HammerOnPullOff) GongCopy() GongstructIF {
	newInstance := new(HammerOnPullOff)
	hammeronpulloff.GongCopyBasicFields(newInstance)
	return newInstance
}

func (handbell *Handbell) GongCopy() GongstructIF {
	newInstance := new(Handbell)
	handbell.GongCopyBasicFields(newInstance)
	return newInstance
}

func (harmonclosed *HarmonClosed) GongCopy() GongstructIF {
	newInstance := new(HarmonClosed)
	harmonclosed.GongCopyBasicFields(newInstance)
	return newInstance
}

func (harmonmute *HarmonMute) GongCopy() GongstructIF {
	newInstance := new(HarmonMute)
	harmonmute.GongCopyBasicFields(newInstance)
	return newInstance
}

func (harmonic *Harmonic) GongCopy() GongstructIF {
	newInstance := new(Harmonic)
	harmonic.GongCopyBasicFields(newInstance)
	return newInstance
}

func (heeltoe *HeelToe) GongCopy() GongstructIF {
	newInstance := new(HeelToe)
	heeltoe.GongCopyBasicFields(newInstance)
	return newInstance
}

func (hole *Hole) GongCopy() GongstructIF {
	newInstance := new(Hole)
	hole.GongCopyBasicFields(newInstance)
	return newInstance
}

func (holeclosed *HoleClosed) GongCopy() GongstructIF {
	newInstance := new(HoleClosed)
	holeclosed.GongCopyBasicFields(newInstance)
	return newInstance
}

func (instrument *Instrument) GongCopy() GongstructIF {
	newInstance := new(Instrument)
	instrument.GongCopyBasicFields(newInstance)
	return newInstance
}

func (listen *Listen) GongCopy() GongstructIF {
	newInstance := new(Listen)
	listen.GongCopyBasicFields(newInstance)
	return newInstance
}

func (lyric *Lyric) GongCopy() GongstructIF {
	newInstance := new(Lyric)
	lyric.GongCopyBasicFields(newInstance)
	return newInstance
}

func (mordent *Mordent) GongCopy() GongstructIF {
	newInstance := new(Mordent)
	mordent.GongCopyBasicFields(newInstance)
	return newInstance
}

func (nonarpeggiate *NonArpeggiate) GongCopy() GongstructIF {
	newInstance := new(NonArpeggiate)
	nonarpeggiate.GongCopyBasicFields(newInstance)
	return newInstance
}

func (notations *Notations) GongCopy() GongstructIF {
	newInstance := new(Notations)
	notations.GongCopyBasicFields(newInstance)
	return newInstance
}

func (note *Note) GongCopy() GongstructIF {
	newInstance := new(Note)
	note.GongCopyBasicFields(newInstance)
	return newInstance
}

func (notetype *NoteType) GongCopy() GongstructIF {
	newInstance := new(NoteType)
	notetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (notehead *Notehead) GongCopy() GongstructIF {
	newInstance := new(Notehead)
	notehead.GongCopyBasicFields(newInstance)
	return newInstance
}

func (noteheadtext *NoteheadText) GongCopy() GongstructIF {
	newInstance := new(NoteheadText)
	noteheadtext.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ornaments *Ornaments) GongCopy() GongstructIF {
	newInstance := new(Ornaments)
	ornaments.GongCopyBasicFields(newInstance)
	return newInstance
}

func (othernotation *OtherNotation) GongCopy() GongstructIF {
	newInstance := new(OtherNotation)
	othernotation.GongCopyBasicFields(newInstance)
	return newInstance
}

func (otherplacementtext *OtherPlacementText) GongCopy() GongstructIF {
	newInstance := new(OtherPlacementText)
	otherplacementtext.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pitch *Pitch) GongCopy() GongstructIF {
	newInstance := new(Pitch)
	pitch.GongCopyBasicFields(newInstance)
	return newInstance
}

func (placementtext *PlacementText) GongCopy() GongstructIF {
	newInstance := new(PlacementText)
	placementtext.GongCopyBasicFields(newInstance)
	return newInstance
}

func (release *Release) GongCopy() GongstructIF {
	newInstance := new(Release)
	release.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rest *Rest) GongCopy() GongstructIF {
	newInstance := new(Rest)
	rest.GongCopyBasicFields(newInstance)
	return newInstance
}

func (slide *Slide) GongCopy() GongstructIF {
	newInstance := new(Slide)
	slide.GongCopyBasicFields(newInstance)
	return newInstance
}

func (slur *Slur) GongCopy() GongstructIF {
	newInstance := new(Slur)
	slur.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stem *Stem) GongCopy() GongstructIF {
	newInstance := new(Stem)
	stem.GongCopyBasicFields(newInstance)
	return newInstance
}

func (strongaccent *StrongAccent) GongCopy() GongstructIF {
	newInstance := new(StrongAccent)
	strongaccent.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tap *Tap) GongCopy() GongstructIF {
	newInstance := new(Tap)
	tap.GongCopyBasicFields(newInstance)
	return newInstance
}

func (technical *Technical) GongCopy() GongstructIF {
	newInstance := new(Technical)
	technical.GongCopyBasicFields(newInstance)
	return newInstance
}

func (textelementdata *TextElementData) GongCopy() GongstructIF {
	newInstance := new(TextElementData)
	textelementdata.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tie *Tie) GongCopy() GongstructIF {
	newInstance := new(Tie)
	tie.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tied *Tied) GongCopy() GongstructIF {
	newInstance := new(Tied)
	tied.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tremolo *Tremolo) GongCopy() GongstructIF {
	newInstance := new(Tremolo)
	tremolo.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tuplet *Tuplet) GongCopy() GongstructIF {
	newInstance := new(Tuplet)
	tuplet.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tupletdot *TupletDot) GongCopy() GongstructIF {
	newInstance := new(TupletDot)
	tupletdot.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tupletnumber *TupletNumber) GongCopy() GongstructIF {
	newInstance := new(TupletNumber)
	tupletnumber.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tupletportion *TupletPortion) GongCopy() GongstructIF {
	newInstance := new(TupletPortion)
	tupletportion.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tuplettype *TupletType) GongCopy() GongstructIF {
	newInstance := new(TupletType)
	tuplettype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (unpitched *Unpitched) GongCopy() GongstructIF {
	newInstance := new(Unpitched)
	unpitched.GongCopyBasicFields(newInstance)
	return newInstance
}

func (wait *Wait) GongCopy() GongstructIF {
	newInstance := new(Wait)
	wait.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (accidental *Accidental) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, accidental)
}

func (accidentalmark *AccidentalMark) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, accidentalmark)
}

func (arpeggiate *Arpeggiate) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, arpeggiate)
}

func (arrow *Arrow) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, arrow)
}

func (articulations *Articulations) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, articulations)
}

func (assess *Assess) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, assess)
}

func (backup *Backup) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, backup)
}

func (beam *Beam) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, beam)
}

func (bend *Bend) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, bend)
}

func (breathmark *BreathMark) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, breathmark)
}

func (caesura *Caesura) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, caesura)
}

func (elision *Elision) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elision)
}

func (emptyline *EmptyLine) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, emptyline)
}

func (extend *Extend) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, extend)
}

func (figure *Figure) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, figure)
}

func (figuredbass *FiguredBass) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, figuredbass)
}

func (forward *Forward) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, forward)
}

func (glissando *Glissando) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, glissando)
}

func (grace *Grace) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, grace)
}

func (hammeronpulloff *HammerOnPullOff) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, hammeronpulloff)
}

func (handbell *Handbell) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, handbell)
}

func (harmonclosed *HarmonClosed) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, harmonclosed)
}

func (harmonmute *HarmonMute) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, harmonmute)
}

func (harmonic *Harmonic) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, harmonic)
}

func (heeltoe *HeelToe) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, heeltoe)
}

func (hole *Hole) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, hole)
}

func (holeclosed *HoleClosed) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, holeclosed)
}

func (instrument *Instrument) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, instrument)
}

func (listen *Listen) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, listen)
}

func (lyric *Lyric) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, lyric)
}

func (mordent *Mordent) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, mordent)
}

func (nonarpeggiate *NonArpeggiate) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, nonarpeggiate)
}

func (notations *Notations) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notations)
}

func (note *Note) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, note)
}

func (notetype *NoteType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notetype)
}

func (notehead *Notehead) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notehead)
}

func (noteheadtext *NoteheadText) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, noteheadtext)
}

func (ornaments *Ornaments) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ornaments)
}

func (othernotation *OtherNotation) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, othernotation)
}

func (otherplacementtext *OtherPlacementText) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, otherplacementtext)
}

func (pitch *Pitch) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pitch)
}

func (placementtext *PlacementText) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, placementtext)
}

func (release *Release) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, release)
}

func (rest *Rest) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rest)
}

func (slide *Slide) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, slide)
}

func (slur *Slur) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, slur)
}

func (stem *Stem) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stem)
}

func (strongaccent *StrongAccent) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, strongaccent)
}

func (tap *Tap) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tap)
}

func (technical *Technical) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, technical)
}

func (textelementdata *TextElementData) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textelementdata)
}

func (tie *Tie) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tie)
}

func (tied *Tied) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tied)
}

func (tremolo *Tremolo) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tremolo)
}

func (tuplet *Tuplet) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tuplet)
}

func (tupletdot *TupletDot) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tupletdot)
}

func (tupletnumber *TupletNumber) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tupletnumber)
}

func (tupletportion *TupletPortion) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tupletportion)
}

func (tuplettype *TupletType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tuplettype)
}

func (unpitched *Unpitched) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, unpitched)
}

func (wait *Wait) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, wait)
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
		stage.Accidentals,
		stage.Accidental_stagedOrder,
		stage.Accidentals_reference,
		&stage.Accidentals_referenceOrder,
		stage.Accidentals_instance,
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
		stage.AccidentalMarks,
		stage.AccidentalMark_stagedOrder,
		stage.AccidentalMarks_reference,
		&stage.AccidentalMarks_referenceOrder,
		stage.AccidentalMarks_instance,
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
		stage.Arpeggiates,
		stage.Arpeggiate_stagedOrder,
		stage.Arpeggiates_reference,
		&stage.Arpeggiates_referenceOrder,
		stage.Arpeggiates_instance,
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
		stage.Arrows,
		stage.Arrow_stagedOrder,
		stage.Arrows_reference,
		&stage.Arrows_referenceOrder,
		stage.Arrows_instance,
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
		stage.Articulationss,
		stage.Articulations_stagedOrder,
		stage.Articulationss_reference,
		&stage.Articulationss_referenceOrder,
		stage.Articulationss_instance,
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
		stage.Assesss,
		stage.Assess_stagedOrder,
		stage.Assesss_reference,
		&stage.Assesss_referenceOrder,
		stage.Assesss_instance,
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
		stage.Backups,
		stage.Backup_stagedOrder,
		stage.Backups_reference,
		&stage.Backups_referenceOrder,
		stage.Backups_instance,
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
		stage.Beams,
		stage.Beam_stagedOrder,
		stage.Beams_reference,
		&stage.Beams_referenceOrder,
		stage.Beams_instance,
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
		stage.Bends,
		stage.Bend_stagedOrder,
		stage.Bends_reference,
		&stage.Bends_referenceOrder,
		stage.Bends_instance,
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
		stage.BreathMarks,
		stage.BreathMark_stagedOrder,
		stage.BreathMarks_reference,
		&stage.BreathMarks_referenceOrder,
		stage.BreathMarks_instance,
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
		stage.Caesuras,
		stage.Caesura_stagedOrder,
		stage.Caesuras_reference,
		&stage.Caesuras_referenceOrder,
		stage.Caesuras_instance,
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
		stage.Elisions,
		stage.Elision_stagedOrder,
		stage.Elisions_reference,
		&stage.Elisions_referenceOrder,
		stage.Elisions_instance,
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
		stage.EmptyLines,
		stage.EmptyLine_stagedOrder,
		stage.EmptyLines_reference,
		&stage.EmptyLines_referenceOrder,
		stage.EmptyLines_instance,
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
		stage.Extends,
		stage.Extend_stagedOrder,
		stage.Extends_reference,
		&stage.Extends_referenceOrder,
		stage.Extends_instance,
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
		stage.Figures,
		stage.Figure_stagedOrder,
		stage.Figures_reference,
		&stage.Figures_referenceOrder,
		stage.Figures_instance,
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
		stage.FiguredBasss,
		stage.FiguredBass_stagedOrder,
		stage.FiguredBasss_reference,
		&stage.FiguredBasss_referenceOrder,
		stage.FiguredBasss_instance,
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
		stage.Forwards,
		stage.Forward_stagedOrder,
		stage.Forwards_reference,
		&stage.Forwards_referenceOrder,
		stage.Forwards_instance,
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
		stage.Glissandos,
		stage.Glissando_stagedOrder,
		stage.Glissandos_reference,
		&stage.Glissandos_referenceOrder,
		stage.Glissandos_instance,
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
		stage.Graces,
		stage.Grace_stagedOrder,
		stage.Graces_reference,
		&stage.Graces_referenceOrder,
		stage.Graces_instance,
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
		stage.HammerOnPullOffs,
		stage.HammerOnPullOff_stagedOrder,
		stage.HammerOnPullOffs_reference,
		&stage.HammerOnPullOffs_referenceOrder,
		stage.HammerOnPullOffs_instance,
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
		stage.Handbells,
		stage.Handbell_stagedOrder,
		stage.Handbells_reference,
		&stage.Handbells_referenceOrder,
		stage.Handbells_instance,
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
		stage.HarmonCloseds,
		stage.HarmonClosed_stagedOrder,
		stage.HarmonCloseds_reference,
		&stage.HarmonCloseds_referenceOrder,
		stage.HarmonCloseds_instance,
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
		stage.HarmonMutes,
		stage.HarmonMute_stagedOrder,
		stage.HarmonMutes_reference,
		&stage.HarmonMutes_referenceOrder,
		stage.HarmonMutes_instance,
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
		stage.Harmonics,
		stage.Harmonic_stagedOrder,
		stage.Harmonics_reference,
		&stage.Harmonics_referenceOrder,
		stage.Harmonics_instance,
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
		stage.HeelToes,
		stage.HeelToe_stagedOrder,
		stage.HeelToes_reference,
		&stage.HeelToes_referenceOrder,
		stage.HeelToes_instance,
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
		stage.Holes,
		stage.Hole_stagedOrder,
		stage.Holes_reference,
		&stage.Holes_referenceOrder,
		stage.Holes_instance,
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
		stage.HoleCloseds,
		stage.HoleClosed_stagedOrder,
		stage.HoleCloseds_reference,
		&stage.HoleCloseds_referenceOrder,
		stage.HoleCloseds_instance,
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
		stage.Instruments,
		stage.Instrument_stagedOrder,
		stage.Instruments_reference,
		&stage.Instruments_referenceOrder,
		stage.Instruments_instance,
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
		stage.Listens,
		stage.Listen_stagedOrder,
		stage.Listens_reference,
		&stage.Listens_referenceOrder,
		stage.Listens_instance,
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
		stage.Lyrics,
		stage.Lyric_stagedOrder,
		stage.Lyrics_reference,
		&stage.Lyrics_referenceOrder,
		stage.Lyrics_instance,
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
		stage.Mordents,
		stage.Mordent_stagedOrder,
		stage.Mordents_reference,
		&stage.Mordents_referenceOrder,
		stage.Mordents_instance,
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
		stage.NonArpeggiates,
		stage.NonArpeggiate_stagedOrder,
		stage.NonArpeggiates_reference,
		&stage.NonArpeggiates_referenceOrder,
		stage.NonArpeggiates_instance,
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
		stage.Notationss,
		stage.Notations_stagedOrder,
		stage.Notationss_reference,
		&stage.Notationss_referenceOrder,
		stage.Notationss_instance,
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
		stage.Notes,
		stage.Note_stagedOrder,
		stage.Notes_reference,
		&stage.Notes_referenceOrder,
		stage.Notes_instance,
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
		stage.NoteTypes,
		stage.NoteType_stagedOrder,
		stage.NoteTypes_reference,
		&stage.NoteTypes_referenceOrder,
		stage.NoteTypes_instance,
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
		stage.Noteheads,
		stage.Notehead_stagedOrder,
		stage.Noteheads_reference,
		&stage.Noteheads_referenceOrder,
		stage.Noteheads_instance,
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
		stage.NoteheadTexts,
		stage.NoteheadText_stagedOrder,
		stage.NoteheadTexts_reference,
		&stage.NoteheadTexts_referenceOrder,
		stage.NoteheadTexts_instance,
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
		stage.Ornamentss,
		stage.Ornaments_stagedOrder,
		stage.Ornamentss_reference,
		&stage.Ornamentss_referenceOrder,
		stage.Ornamentss_instance,
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
		stage.OtherNotations,
		stage.OtherNotation_stagedOrder,
		stage.OtherNotations_reference,
		&stage.OtherNotations_referenceOrder,
		stage.OtherNotations_instance,
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
		stage.OtherPlacementTexts,
		stage.OtherPlacementText_stagedOrder,
		stage.OtherPlacementTexts_reference,
		&stage.OtherPlacementTexts_referenceOrder,
		stage.OtherPlacementTexts_instance,
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
		stage.Pitchs,
		stage.Pitch_stagedOrder,
		stage.Pitchs_reference,
		&stage.Pitchs_referenceOrder,
		stage.Pitchs_instance,
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
		stage.PlacementTexts,
		stage.PlacementText_stagedOrder,
		stage.PlacementTexts_reference,
		&stage.PlacementTexts_referenceOrder,
		stage.PlacementTexts_instance,
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
		stage.Releases,
		stage.Release_stagedOrder,
		stage.Releases_reference,
		&stage.Releases_referenceOrder,
		stage.Releases_instance,
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
		stage.Rests,
		stage.Rest_stagedOrder,
		stage.Rests_reference,
		&stage.Rests_referenceOrder,
		stage.Rests_instance,
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
		stage.Slides,
		stage.Slide_stagedOrder,
		stage.Slides_reference,
		&stage.Slides_referenceOrder,
		stage.Slides_instance,
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
		stage.Slurs,
		stage.Slur_stagedOrder,
		stage.Slurs_reference,
		&stage.Slurs_referenceOrder,
		stage.Slurs_instance,
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
		stage.Stems,
		stage.Stem_stagedOrder,
		stage.Stems_reference,
		&stage.Stems_referenceOrder,
		stage.Stems_instance,
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
		stage.StrongAccents,
		stage.StrongAccent_stagedOrder,
		stage.StrongAccents_reference,
		&stage.StrongAccents_referenceOrder,
		stage.StrongAccents_instance,
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
		stage.Taps,
		stage.Tap_stagedOrder,
		stage.Taps_reference,
		&stage.Taps_referenceOrder,
		stage.Taps_instance,
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
		stage.Technicals,
		stage.Technical_stagedOrder,
		stage.Technicals_reference,
		&stage.Technicals_referenceOrder,
		stage.Technicals_instance,
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
		stage.TextElementDatas,
		stage.TextElementData_stagedOrder,
		stage.TextElementDatas_reference,
		&stage.TextElementDatas_referenceOrder,
		stage.TextElementDatas_instance,
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
		stage.Ties,
		stage.Tie_stagedOrder,
		stage.Ties_reference,
		&stage.Ties_referenceOrder,
		stage.Ties_instance,
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
		stage.Tieds,
		stage.Tied_stagedOrder,
		stage.Tieds_reference,
		&stage.Tieds_referenceOrder,
		stage.Tieds_instance,
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
		stage.Tremolos,
		stage.Tremolo_stagedOrder,
		stage.Tremolos_reference,
		&stage.Tremolos_referenceOrder,
		stage.Tremolos_instance,
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
		stage.Tuplets,
		stage.Tuplet_stagedOrder,
		stage.Tuplets_reference,
		&stage.Tuplets_referenceOrder,
		stage.Tuplets_instance,
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
		stage.TupletDots,
		stage.TupletDot_stagedOrder,
		stage.TupletDots_reference,
		&stage.TupletDots_referenceOrder,
		stage.TupletDots_instance,
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
		stage.TupletNumbers,
		stage.TupletNumber_stagedOrder,
		stage.TupletNumbers_reference,
		&stage.TupletNumbers_referenceOrder,
		stage.TupletNumbers_instance,
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
		stage.TupletPortions,
		stage.TupletPortion_stagedOrder,
		stage.TupletPortions_reference,
		&stage.TupletPortions_referenceOrder,
		stage.TupletPortions_instance,
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
		stage.TupletTypes,
		stage.TupletType_stagedOrder,
		stage.TupletTypes_reference,
		&stage.TupletTypes_referenceOrder,
		stage.TupletTypes_instance,
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
		stage.Unpitcheds,
		stage.Unpitched_stagedOrder,
		stage.Unpitcheds_reference,
		&stage.Unpitcheds_referenceOrder,
		stage.Unpitcheds_instance,
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
		stage.Waits,
		stage.Wait_stagedOrder,
		stage.Waits_reference,
		&stage.Waits_referenceOrder,
		stage.Waits_instance,
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
	__gong__computeReferencePass1(stage, stage.Accidentals, &stage.Accidentals_reference, &stage.Accidentals_referenceOrder, &stage.Accidentals_instance)

	__gong__computeReferencePass1(stage, stage.AccidentalMarks, &stage.AccidentalMarks_reference, &stage.AccidentalMarks_referenceOrder, &stage.AccidentalMarks_instance)

	__gong__computeReferencePass1(stage, stage.Arpeggiates, &stage.Arpeggiates_reference, &stage.Arpeggiates_referenceOrder, &stage.Arpeggiates_instance)

	__gong__computeReferencePass1(stage, stage.Arrows, &stage.Arrows_reference, &stage.Arrows_referenceOrder, &stage.Arrows_instance)

	__gong__computeReferencePass1(stage, stage.Articulationss, &stage.Articulationss_reference, &stage.Articulationss_referenceOrder, &stage.Articulationss_instance)

	__gong__computeReferencePass1(stage, stage.Assesss, &stage.Assesss_reference, &stage.Assesss_referenceOrder, &stage.Assesss_instance)

	__gong__computeReferencePass1(stage, stage.Backups, &stage.Backups_reference, &stage.Backups_referenceOrder, &stage.Backups_instance)

	__gong__computeReferencePass1(stage, stage.Beams, &stage.Beams_reference, &stage.Beams_referenceOrder, &stage.Beams_instance)

	__gong__computeReferencePass1(stage, stage.Bends, &stage.Bends_reference, &stage.Bends_referenceOrder, &stage.Bends_instance)

	__gong__computeReferencePass1(stage, stage.BreathMarks, &stage.BreathMarks_reference, &stage.BreathMarks_referenceOrder, &stage.BreathMarks_instance)

	__gong__computeReferencePass1(stage, stage.Caesuras, &stage.Caesuras_reference, &stage.Caesuras_referenceOrder, &stage.Caesuras_instance)

	__gong__computeReferencePass1(stage, stage.Elisions, &stage.Elisions_reference, &stage.Elisions_referenceOrder, &stage.Elisions_instance)

	__gong__computeReferencePass1(stage, stage.EmptyLines, &stage.EmptyLines_reference, &stage.EmptyLines_referenceOrder, &stage.EmptyLines_instance)

	__gong__computeReferencePass1(stage, stage.Extends, &stage.Extends_reference, &stage.Extends_referenceOrder, &stage.Extends_instance)

	__gong__computeReferencePass1(stage, stage.Figures, &stage.Figures_reference, &stage.Figures_referenceOrder, &stage.Figures_instance)

	__gong__computeReferencePass1(stage, stage.FiguredBasss, &stage.FiguredBasss_reference, &stage.FiguredBasss_referenceOrder, &stage.FiguredBasss_instance)

	__gong__computeReferencePass1(stage, stage.Forwards, &stage.Forwards_reference, &stage.Forwards_referenceOrder, &stage.Forwards_instance)

	__gong__computeReferencePass1(stage, stage.Glissandos, &stage.Glissandos_reference, &stage.Glissandos_referenceOrder, &stage.Glissandos_instance)

	__gong__computeReferencePass1(stage, stage.Graces, &stage.Graces_reference, &stage.Graces_referenceOrder, &stage.Graces_instance)

	__gong__computeReferencePass1(stage, stage.HammerOnPullOffs, &stage.HammerOnPullOffs_reference, &stage.HammerOnPullOffs_referenceOrder, &stage.HammerOnPullOffs_instance)

	__gong__computeReferencePass1(stage, stage.Handbells, &stage.Handbells_reference, &stage.Handbells_referenceOrder, &stage.Handbells_instance)

	__gong__computeReferencePass1(stage, stage.HarmonCloseds, &stage.HarmonCloseds_reference, &stage.HarmonCloseds_referenceOrder, &stage.HarmonCloseds_instance)

	__gong__computeReferencePass1(stage, stage.HarmonMutes, &stage.HarmonMutes_reference, &stage.HarmonMutes_referenceOrder, &stage.HarmonMutes_instance)

	__gong__computeReferencePass1(stage, stage.Harmonics, &stage.Harmonics_reference, &stage.Harmonics_referenceOrder, &stage.Harmonics_instance)

	__gong__computeReferencePass1(stage, stage.HeelToes, &stage.HeelToes_reference, &stage.HeelToes_referenceOrder, &stage.HeelToes_instance)

	__gong__computeReferencePass1(stage, stage.Holes, &stage.Holes_reference, &stage.Holes_referenceOrder, &stage.Holes_instance)

	__gong__computeReferencePass1(stage, stage.HoleCloseds, &stage.HoleCloseds_reference, &stage.HoleCloseds_referenceOrder, &stage.HoleCloseds_instance)

	__gong__computeReferencePass1(stage, stage.Instruments, &stage.Instruments_reference, &stage.Instruments_referenceOrder, &stage.Instruments_instance)

	__gong__computeReferencePass1(stage, stage.Listens, &stage.Listens_reference, &stage.Listens_referenceOrder, &stage.Listens_instance)

	__gong__computeReferencePass1(stage, stage.Lyrics, &stage.Lyrics_reference, &stage.Lyrics_referenceOrder, &stage.Lyrics_instance)

	__gong__computeReferencePass1(stage, stage.Mordents, &stage.Mordents_reference, &stage.Mordents_referenceOrder, &stage.Mordents_instance)

	__gong__computeReferencePass1(stage, stage.NonArpeggiates, &stage.NonArpeggiates_reference, &stage.NonArpeggiates_referenceOrder, &stage.NonArpeggiates_instance)

	__gong__computeReferencePass1(stage, stage.Notationss, &stage.Notationss_reference, &stage.Notationss_referenceOrder, &stage.Notationss_instance)

	__gong__computeReferencePass1(stage, stage.Notes, &stage.Notes_reference, &stage.Notes_referenceOrder, &stage.Notes_instance)

	__gong__computeReferencePass1(stage, stage.NoteTypes, &stage.NoteTypes_reference, &stage.NoteTypes_referenceOrder, &stage.NoteTypes_instance)

	__gong__computeReferencePass1(stage, stage.Noteheads, &stage.Noteheads_reference, &stage.Noteheads_referenceOrder, &stage.Noteheads_instance)

	__gong__computeReferencePass1(stage, stage.NoteheadTexts, &stage.NoteheadTexts_reference, &stage.NoteheadTexts_referenceOrder, &stage.NoteheadTexts_instance)

	__gong__computeReferencePass1(stage, stage.Ornamentss, &stage.Ornamentss_reference, &stage.Ornamentss_referenceOrder, &stage.Ornamentss_instance)

	__gong__computeReferencePass1(stage, stage.OtherNotations, &stage.OtherNotations_reference, &stage.OtherNotations_referenceOrder, &stage.OtherNotations_instance)

	__gong__computeReferencePass1(stage, stage.OtherPlacementTexts, &stage.OtherPlacementTexts_reference, &stage.OtherPlacementTexts_referenceOrder, &stage.OtherPlacementTexts_instance)

	__gong__computeReferencePass1(stage, stage.Pitchs, &stage.Pitchs_reference, &stage.Pitchs_referenceOrder, &stage.Pitchs_instance)

	__gong__computeReferencePass1(stage, stage.PlacementTexts, &stage.PlacementTexts_reference, &stage.PlacementTexts_referenceOrder, &stage.PlacementTexts_instance)

	__gong__computeReferencePass1(stage, stage.Releases, &stage.Releases_reference, &stage.Releases_referenceOrder, &stage.Releases_instance)

	__gong__computeReferencePass1(stage, stage.Rests, &stage.Rests_reference, &stage.Rests_referenceOrder, &stage.Rests_instance)

	__gong__computeReferencePass1(stage, stage.Slides, &stage.Slides_reference, &stage.Slides_referenceOrder, &stage.Slides_instance)

	__gong__computeReferencePass1(stage, stage.Slurs, &stage.Slurs_reference, &stage.Slurs_referenceOrder, &stage.Slurs_instance)

	__gong__computeReferencePass1(stage, stage.Stems, &stage.Stems_reference, &stage.Stems_referenceOrder, &stage.Stems_instance)

	__gong__computeReferencePass1(stage, stage.StrongAccents, &stage.StrongAccents_reference, &stage.StrongAccents_referenceOrder, &stage.StrongAccents_instance)

	__gong__computeReferencePass1(stage, stage.Taps, &stage.Taps_reference, &stage.Taps_referenceOrder, &stage.Taps_instance)

	__gong__computeReferencePass1(stage, stage.Technicals, &stage.Technicals_reference, &stage.Technicals_referenceOrder, &stage.Technicals_instance)

	__gong__computeReferencePass1(stage, stage.TextElementDatas, &stage.TextElementDatas_reference, &stage.TextElementDatas_referenceOrder, &stage.TextElementDatas_instance)

	__gong__computeReferencePass1(stage, stage.Ties, &stage.Ties_reference, &stage.Ties_referenceOrder, &stage.Ties_instance)

	__gong__computeReferencePass1(stage, stage.Tieds, &stage.Tieds_reference, &stage.Tieds_referenceOrder, &stage.Tieds_instance)

	__gong__computeReferencePass1(stage, stage.Tremolos, &stage.Tremolos_reference, &stage.Tremolos_referenceOrder, &stage.Tremolos_instance)

	__gong__computeReferencePass1(stage, stage.Tuplets, &stage.Tuplets_reference, &stage.Tuplets_referenceOrder, &stage.Tuplets_instance)

	__gong__computeReferencePass1(stage, stage.TupletDots, &stage.TupletDots_reference, &stage.TupletDots_referenceOrder, &stage.TupletDots_instance)

	__gong__computeReferencePass1(stage, stage.TupletNumbers, &stage.TupletNumbers_reference, &stage.TupletNumbers_referenceOrder, &stage.TupletNumbers_instance)

	__gong__computeReferencePass1(stage, stage.TupletPortions, &stage.TupletPortions_reference, &stage.TupletPortions_referenceOrder, &stage.TupletPortions_instance)

	__gong__computeReferencePass1(stage, stage.TupletTypes, &stage.TupletTypes_reference, &stage.TupletTypes_referenceOrder, &stage.TupletTypes_instance)

	__gong__computeReferencePass1(stage, stage.Unpitcheds, &stage.Unpitcheds_reference, &stage.Unpitcheds_referenceOrder, &stage.Unpitcheds_instance)

	__gong__computeReferencePass1(stage, stage.Waits, &stage.Waits_reference, &stage.Waits_referenceOrder, &stage.Waits_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.Accidentals, stage.Accidentals_reference, stage)

	__gong__computeReferencePass2(stage.AccidentalMarks, stage.AccidentalMarks_reference, stage)

	__gong__computeReferencePass2(stage.Arpeggiates, stage.Arpeggiates_reference, stage)

	__gong__computeReferencePass2(stage.Arrows, stage.Arrows_reference, stage)

	__gong__computeReferencePass2(stage.Articulationss, stage.Articulationss_reference, stage)

	__gong__computeReferencePass2(stage.Assesss, stage.Assesss_reference, stage)

	__gong__computeReferencePass2(stage.Backups, stage.Backups_reference, stage)

	__gong__computeReferencePass2(stage.Beams, stage.Beams_reference, stage)

	__gong__computeReferencePass2(stage.Bends, stage.Bends_reference, stage)

	__gong__computeReferencePass2(stage.BreathMarks, stage.BreathMarks_reference, stage)

	__gong__computeReferencePass2(stage.Caesuras, stage.Caesuras_reference, stage)

	__gong__computeReferencePass2(stage.Elisions, stage.Elisions_reference, stage)

	__gong__computeReferencePass2(stage.EmptyLines, stage.EmptyLines_reference, stage)

	__gong__computeReferencePass2(stage.Extends, stage.Extends_reference, stage)

	__gong__computeReferencePass2(stage.Figures, stage.Figures_reference, stage)

	__gong__computeReferencePass2(stage.FiguredBasss, stage.FiguredBasss_reference, stage)

	__gong__computeReferencePass2(stage.Forwards, stage.Forwards_reference, stage)

	__gong__computeReferencePass2(stage.Glissandos, stage.Glissandos_reference, stage)

	__gong__computeReferencePass2(stage.Graces, stage.Graces_reference, stage)

	__gong__computeReferencePass2(stage.HammerOnPullOffs, stage.HammerOnPullOffs_reference, stage)

	__gong__computeReferencePass2(stage.Handbells, stage.Handbells_reference, stage)

	__gong__computeReferencePass2(stage.HarmonCloseds, stage.HarmonCloseds_reference, stage)

	__gong__computeReferencePass2(stage.HarmonMutes, stage.HarmonMutes_reference, stage)

	__gong__computeReferencePass2(stage.Harmonics, stage.Harmonics_reference, stage)

	__gong__computeReferencePass2(stage.HeelToes, stage.HeelToes_reference, stage)

	__gong__computeReferencePass2(stage.Holes, stage.Holes_reference, stage)

	__gong__computeReferencePass2(stage.HoleCloseds, stage.HoleCloseds_reference, stage)

	__gong__computeReferencePass2(stage.Instruments, stage.Instruments_reference, stage)

	__gong__computeReferencePass2(stage.Listens, stage.Listens_reference, stage)

	__gong__computeReferencePass2(stage.Lyrics, stage.Lyrics_reference, stage)

	__gong__computeReferencePass2(stage.Mordents, stage.Mordents_reference, stage)

	__gong__computeReferencePass2(stage.NonArpeggiates, stage.NonArpeggiates_reference, stage)

	__gong__computeReferencePass2(stage.Notationss, stage.Notationss_reference, stage)

	__gong__computeReferencePass2(stage.Notes, stage.Notes_reference, stage)

	__gong__computeReferencePass2(stage.NoteTypes, stage.NoteTypes_reference, stage)

	__gong__computeReferencePass2(stage.Noteheads, stage.Noteheads_reference, stage)

	__gong__computeReferencePass2(stage.NoteheadTexts, stage.NoteheadTexts_reference, stage)

	__gong__computeReferencePass2(stage.Ornamentss, stage.Ornamentss_reference, stage)

	__gong__computeReferencePass2(stage.OtherNotations, stage.OtherNotations_reference, stage)

	__gong__computeReferencePass2(stage.OtherPlacementTexts, stage.OtherPlacementTexts_reference, stage)

	__gong__computeReferencePass2(stage.Pitchs, stage.Pitchs_reference, stage)

	__gong__computeReferencePass2(stage.PlacementTexts, stage.PlacementTexts_reference, stage)

	__gong__computeReferencePass2(stage.Releases, stage.Releases_reference, stage)

	__gong__computeReferencePass2(stage.Rests, stage.Rests_reference, stage)

	__gong__computeReferencePass2(stage.Slides, stage.Slides_reference, stage)

	__gong__computeReferencePass2(stage.Slurs, stage.Slurs_reference, stage)

	__gong__computeReferencePass2(stage.Stems, stage.Stems_reference, stage)

	__gong__computeReferencePass2(stage.StrongAccents, stage.StrongAccents_reference, stage)

	__gong__computeReferencePass2(stage.Taps, stage.Taps_reference, stage)

	__gong__computeReferencePass2(stage.Technicals, stage.Technicals_reference, stage)

	__gong__computeReferencePass2(stage.TextElementDatas, stage.TextElementDatas_reference, stage)

	__gong__computeReferencePass2(stage.Ties, stage.Ties_reference, stage)

	__gong__computeReferencePass2(stage.Tieds, stage.Tieds_reference, stage)

	__gong__computeReferencePass2(stage.Tremolos, stage.Tremolos_reference, stage)

	__gong__computeReferencePass2(stage.Tuplets, stage.Tuplets_reference, stage)

	__gong__computeReferencePass2(stage.TupletDots, stage.TupletDots_reference, stage)

	__gong__computeReferencePass2(stage.TupletNumbers, stage.TupletNumbers_reference, stage)

	__gong__computeReferencePass2(stage.TupletPortions, stage.TupletPortions_reference, stage)

	__gong__computeReferencePass2(stage.TupletTypes, stage.TupletTypes_reference, stage)

	__gong__computeReferencePass2(stage.Unpitcheds, stage.Unpitcheds_reference, stage)

	__gong__computeReferencePass2(stage.Waits, stage.Waits_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (accidental *Accidental) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Accidental_stagedOrder, stage.Accidentals_referenceOrder, accidental, "Accidental")
}

func (accidentalmark *AccidentalMark) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AccidentalMark_stagedOrder, stage.AccidentalMarks_referenceOrder, accidentalmark, "AccidentalMark")
}

func (arpeggiate *Arpeggiate) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Arpeggiate_stagedOrder, stage.Arpeggiates_referenceOrder, arpeggiate, "Arpeggiate")
}

func (arrow *Arrow) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Arrow_stagedOrder, stage.Arrows_referenceOrder, arrow, "Arrow")
}

func (articulations *Articulations) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Articulations_stagedOrder, stage.Articulationss_referenceOrder, articulations, "Articulations")
}

func (assess *Assess) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Assess_stagedOrder, stage.Assesss_referenceOrder, assess, "Assess")
}

func (backup *Backup) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Backup_stagedOrder, stage.Backups_referenceOrder, backup, "Backup")
}

func (beam *Beam) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Beam_stagedOrder, stage.Beams_referenceOrder, beam, "Beam")
}

func (bend *Bend) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Bend_stagedOrder, stage.Bends_referenceOrder, bend, "Bend")
}

func (breathmark *BreathMark) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.BreathMark_stagedOrder, stage.BreathMarks_referenceOrder, breathmark, "BreathMark")
}

func (caesura *Caesura) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Caesura_stagedOrder, stage.Caesuras_referenceOrder, caesura, "Caesura")
}

func (elision *Elision) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Elision_stagedOrder, stage.Elisions_referenceOrder, elision, "Elision")
}

func (emptyline *EmptyLine) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EmptyLine_stagedOrder, stage.EmptyLines_referenceOrder, emptyline, "EmptyLine")
}

func (extend *Extend) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Extend_stagedOrder, stage.Extends_referenceOrder, extend, "Extend")
}

func (figure *Figure) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Figure_stagedOrder, stage.Figures_referenceOrder, figure, "Figure")
}

func (figuredbass *FiguredBass) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FiguredBass_stagedOrder, stage.FiguredBasss_referenceOrder, figuredbass, "FiguredBass")
}

func (forward *Forward) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Forward_stagedOrder, stage.Forwards_referenceOrder, forward, "Forward")
}

func (glissando *Glissando) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Glissando_stagedOrder, stage.Glissandos_referenceOrder, glissando, "Glissando")
}

func (grace *Grace) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Grace_stagedOrder, stage.Graces_referenceOrder, grace, "Grace")
}

func (hammeronpulloff *HammerOnPullOff) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HammerOnPullOff_stagedOrder, stage.HammerOnPullOffs_referenceOrder, hammeronpulloff, "HammerOnPullOff")
}

func (handbell *Handbell) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Handbell_stagedOrder, stage.Handbells_referenceOrder, handbell, "Handbell")
}

func (harmonclosed *HarmonClosed) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HarmonClosed_stagedOrder, stage.HarmonCloseds_referenceOrder, harmonclosed, "HarmonClosed")
}

func (harmonmute *HarmonMute) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HarmonMute_stagedOrder, stage.HarmonMutes_referenceOrder, harmonmute, "HarmonMute")
}

func (harmonic *Harmonic) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Harmonic_stagedOrder, stage.Harmonics_referenceOrder, harmonic, "Harmonic")
}

func (heeltoe *HeelToe) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HeelToe_stagedOrder, stage.HeelToes_referenceOrder, heeltoe, "HeelToe")
}

func (hole *Hole) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Hole_stagedOrder, stage.Holes_referenceOrder, hole, "Hole")
}

func (holeclosed *HoleClosed) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HoleClosed_stagedOrder, stage.HoleCloseds_referenceOrder, holeclosed, "HoleClosed")
}

func (instrument *Instrument) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Instrument_stagedOrder, stage.Instruments_referenceOrder, instrument, "Instrument")
}

func (listen *Listen) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Listen_stagedOrder, stage.Listens_referenceOrder, listen, "Listen")
}

func (lyric *Lyric) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Lyric_stagedOrder, stage.Lyrics_referenceOrder, lyric, "Lyric")
}

func (mordent *Mordent) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Mordent_stagedOrder, stage.Mordents_referenceOrder, mordent, "Mordent")
}

func (nonarpeggiate *NonArpeggiate) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NonArpeggiate_stagedOrder, stage.NonArpeggiates_referenceOrder, nonarpeggiate, "NonArpeggiate")
}

func (notations *Notations) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Notations_stagedOrder, stage.Notationss_referenceOrder, notations, "Notations")
}

func (note *Note) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Note_stagedOrder, stage.Notes_referenceOrder, note, "Note")
}

func (notetype *NoteType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteType_stagedOrder, stage.NoteTypes_referenceOrder, notetype, "NoteType")
}

func (notehead *Notehead) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Notehead_stagedOrder, stage.Noteheads_referenceOrder, notehead, "Notehead")
}

func (noteheadtext *NoteheadText) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteheadText_stagedOrder, stage.NoteheadTexts_referenceOrder, noteheadtext, "NoteheadText")
}

func (ornaments *Ornaments) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Ornaments_stagedOrder, stage.Ornamentss_referenceOrder, ornaments, "Ornaments")
}

func (othernotation *OtherNotation) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OtherNotation_stagedOrder, stage.OtherNotations_referenceOrder, othernotation, "OtherNotation")
}

func (otherplacementtext *OtherPlacementText) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OtherPlacementText_stagedOrder, stage.OtherPlacementTexts_referenceOrder, otherplacementtext, "OtherPlacementText")
}

func (pitch *Pitch) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Pitch_stagedOrder, stage.Pitchs_referenceOrder, pitch, "Pitch")
}

func (placementtext *PlacementText) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PlacementText_stagedOrder, stage.PlacementTexts_referenceOrder, placementtext, "PlacementText")
}

func (release *Release) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Release_stagedOrder, stage.Releases_referenceOrder, release, "Release")
}

func (rest *Rest) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Rest_stagedOrder, stage.Rests_referenceOrder, rest, "Rest")
}

func (slide *Slide) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Slide_stagedOrder, stage.Slides_referenceOrder, slide, "Slide")
}

func (slur *Slur) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Slur_stagedOrder, stage.Slurs_referenceOrder, slur, "Slur")
}

func (stem *Stem) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Stem_stagedOrder, stage.Stems_referenceOrder, stem, "Stem")
}

func (strongaccent *StrongAccent) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StrongAccent_stagedOrder, stage.StrongAccents_referenceOrder, strongaccent, "StrongAccent")
}

func (tap *Tap) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Tap_stagedOrder, stage.Taps_referenceOrder, tap, "Tap")
}

func (technical *Technical) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Technical_stagedOrder, stage.Technicals_referenceOrder, technical, "Technical")
}

func (textelementdata *TextElementData) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextElementData_stagedOrder, stage.TextElementDatas_referenceOrder, textelementdata, "TextElementData")
}

func (tie *Tie) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Tie_stagedOrder, stage.Ties_referenceOrder, tie, "Tie")
}

func (tied *Tied) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Tied_stagedOrder, stage.Tieds_referenceOrder, tied, "Tied")
}

func (tremolo *Tremolo) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Tremolo_stagedOrder, stage.Tremolos_referenceOrder, tremolo, "Tremolo")
}

func (tuplet *Tuplet) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Tuplet_stagedOrder, stage.Tuplets_referenceOrder, tuplet, "Tuplet")
}

func (tupletdot *TupletDot) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TupletDot_stagedOrder, stage.TupletDots_referenceOrder, tupletdot, "TupletDot")
}

func (tupletnumber *TupletNumber) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TupletNumber_stagedOrder, stage.TupletNumbers_referenceOrder, tupletnumber, "TupletNumber")
}

func (tupletportion *TupletPortion) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TupletPortion_stagedOrder, stage.TupletPortions_referenceOrder, tupletportion, "TupletPortion")
}

func (tuplettype *TupletType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TupletType_stagedOrder, stage.TupletTypes_referenceOrder, tuplettype, "TupletType")
}

func (unpitched *Unpitched) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Unpitched_stagedOrder, stage.Unpitcheds_referenceOrder, unpitched, "Unpitched")
}

func (wait *Wait) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Wait_stagedOrder, stage.Waits_referenceOrder, wait, "Wait")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (accidental *Accidental) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(accidental, accidental.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (accidental *Accidental) GongGetReferenceIdentifier(stage *Stage) string {
	return accidental.GongGetIdentifier(stage)
}

func (accidentalmark *AccidentalMark) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(accidentalmark, accidentalmark.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (accidentalmark *AccidentalMark) GongGetReferenceIdentifier(stage *Stage) string {
	return accidentalmark.GongGetIdentifier(stage)
}

func (arpeggiate *Arpeggiate) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(arpeggiate, arpeggiate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (arpeggiate *Arpeggiate) GongGetReferenceIdentifier(stage *Stage) string {
	return arpeggiate.GongGetIdentifier(stage)
}

func (arrow *Arrow) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(arrow, arrow.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (arrow *Arrow) GongGetReferenceIdentifier(stage *Stage) string {
	return arrow.GongGetIdentifier(stage)
}

func (articulations *Articulations) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(articulations, articulations.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (articulations *Articulations) GongGetReferenceIdentifier(stage *Stage) string {
	return articulations.GongGetIdentifier(stage)
}

func (assess *Assess) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(assess, assess.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (assess *Assess) GongGetReferenceIdentifier(stage *Stage) string {
	return assess.GongGetIdentifier(stage)
}

func (backup *Backup) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(backup, backup.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (backup *Backup) GongGetReferenceIdentifier(stage *Stage) string {
	return backup.GongGetIdentifier(stage)
}

func (beam *Beam) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(beam, beam.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (beam *Beam) GongGetReferenceIdentifier(stage *Stage) string {
	return beam.GongGetIdentifier(stage)
}

func (bend *Bend) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(bend, bend.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bend *Bend) GongGetReferenceIdentifier(stage *Stage) string {
	return bend.GongGetIdentifier(stage)
}

func (breathmark *BreathMark) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(breathmark, breathmark.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (breathmark *BreathMark) GongGetReferenceIdentifier(stage *Stage) string {
	return breathmark.GongGetIdentifier(stage)
}

func (caesura *Caesura) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(caesura, caesura.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (caesura *Caesura) GongGetReferenceIdentifier(stage *Stage) string {
	return caesura.GongGetIdentifier(stage)
}

func (elision *Elision) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elision, elision.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elision *Elision) GongGetReferenceIdentifier(stage *Stage) string {
	return elision.GongGetIdentifier(stage)
}

func (emptyline *EmptyLine) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(emptyline, emptyline.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (emptyline *EmptyLine) GongGetReferenceIdentifier(stage *Stage) string {
	return emptyline.GongGetIdentifier(stage)
}

func (extend *Extend) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(extend, extend.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (extend *Extend) GongGetReferenceIdentifier(stage *Stage) string {
	return extend.GongGetIdentifier(stage)
}

func (figure *Figure) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(figure, figure.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (figure *Figure) GongGetReferenceIdentifier(stage *Stage) string {
	return figure.GongGetIdentifier(stage)
}

func (figuredbass *FiguredBass) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(figuredbass, figuredbass.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (figuredbass *FiguredBass) GongGetReferenceIdentifier(stage *Stage) string {
	return figuredbass.GongGetIdentifier(stage)
}

func (forward *Forward) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(forward, forward.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (forward *Forward) GongGetReferenceIdentifier(stage *Stage) string {
	return forward.GongGetIdentifier(stage)
}

func (glissando *Glissando) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(glissando, glissando.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (glissando *Glissando) GongGetReferenceIdentifier(stage *Stage) string {
	return glissando.GongGetIdentifier(stage)
}

func (grace *Grace) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(grace, grace.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (grace *Grace) GongGetReferenceIdentifier(stage *Stage) string {
	return grace.GongGetIdentifier(stage)
}

func (hammeronpulloff *HammerOnPullOff) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(hammeronpulloff, hammeronpulloff.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (hammeronpulloff *HammerOnPullOff) GongGetReferenceIdentifier(stage *Stage) string {
	return hammeronpulloff.GongGetIdentifier(stage)
}

func (handbell *Handbell) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(handbell, handbell.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (handbell *Handbell) GongGetReferenceIdentifier(stage *Stage) string {
	return handbell.GongGetIdentifier(stage)
}

func (harmonclosed *HarmonClosed) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(harmonclosed, harmonclosed.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (harmonclosed *HarmonClosed) GongGetReferenceIdentifier(stage *Stage) string {
	return harmonclosed.GongGetIdentifier(stage)
}

func (harmonmute *HarmonMute) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(harmonmute, harmonmute.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (harmonmute *HarmonMute) GongGetReferenceIdentifier(stage *Stage) string {
	return harmonmute.GongGetIdentifier(stage)
}

func (harmonic *Harmonic) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(harmonic, harmonic.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (harmonic *Harmonic) GongGetReferenceIdentifier(stage *Stage) string {
	return harmonic.GongGetIdentifier(stage)
}

func (heeltoe *HeelToe) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(heeltoe, heeltoe.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (heeltoe *HeelToe) GongGetReferenceIdentifier(stage *Stage) string {
	return heeltoe.GongGetIdentifier(stage)
}

func (hole *Hole) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(hole, hole.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (hole *Hole) GongGetReferenceIdentifier(stage *Stage) string {
	return hole.GongGetIdentifier(stage)
}

func (holeclosed *HoleClosed) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(holeclosed, holeclosed.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (holeclosed *HoleClosed) GongGetReferenceIdentifier(stage *Stage) string {
	return holeclosed.GongGetIdentifier(stage)
}

func (instrument *Instrument) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(instrument, instrument.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (instrument *Instrument) GongGetReferenceIdentifier(stage *Stage) string {
	return instrument.GongGetIdentifier(stage)
}

func (listen *Listen) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(listen, listen.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (listen *Listen) GongGetReferenceIdentifier(stage *Stage) string {
	return listen.GongGetIdentifier(stage)
}

func (lyric *Lyric) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(lyric, lyric.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (lyric *Lyric) GongGetReferenceIdentifier(stage *Stage) string {
	return lyric.GongGetIdentifier(stage)
}

func (mordent *Mordent) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(mordent, mordent.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (mordent *Mordent) GongGetReferenceIdentifier(stage *Stage) string {
	return mordent.GongGetIdentifier(stage)
}

func (nonarpeggiate *NonArpeggiate) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(nonarpeggiate, nonarpeggiate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (nonarpeggiate *NonArpeggiate) GongGetReferenceIdentifier(stage *Stage) string {
	return nonarpeggiate.GongGetIdentifier(stage)
}

func (notations *Notations) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(notations, notations.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notations *Notations) GongGetReferenceIdentifier(stage *Stage) string {
	return notations.GongGetIdentifier(stage)
}

func (note *Note) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(note, note.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (note *Note) GongGetReferenceIdentifier(stage *Stage) string {
	return note.GongGetIdentifier(stage)
}

func (notetype *NoteType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(notetype, notetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notetype *NoteType) GongGetReferenceIdentifier(stage *Stage) string {
	return notetype.GongGetIdentifier(stage)
}

func (notehead *Notehead) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(notehead, notehead.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notehead *Notehead) GongGetReferenceIdentifier(stage *Stage) string {
	return notehead.GongGetIdentifier(stage)
}

func (noteheadtext *NoteheadText) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(noteheadtext, noteheadtext.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteheadtext *NoteheadText) GongGetReferenceIdentifier(stage *Stage) string {
	return noteheadtext.GongGetIdentifier(stage)
}

func (ornaments *Ornaments) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ornaments, ornaments.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ornaments *Ornaments) GongGetReferenceIdentifier(stage *Stage) string {
	return ornaments.GongGetIdentifier(stage)
}

func (othernotation *OtherNotation) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(othernotation, othernotation.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (othernotation *OtherNotation) GongGetReferenceIdentifier(stage *Stage) string {
	return othernotation.GongGetIdentifier(stage)
}

func (otherplacementtext *OtherPlacementText) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(otherplacementtext, otherplacementtext.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (otherplacementtext *OtherPlacementText) GongGetReferenceIdentifier(stage *Stage) string {
	return otherplacementtext.GongGetIdentifier(stage)
}

func (pitch *Pitch) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pitch, pitch.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pitch *Pitch) GongGetReferenceIdentifier(stage *Stage) string {
	return pitch.GongGetIdentifier(stage)
}

func (placementtext *PlacementText) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(placementtext, placementtext.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (placementtext *PlacementText) GongGetReferenceIdentifier(stage *Stage) string {
	return placementtext.GongGetIdentifier(stage)
}

func (release *Release) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(release, release.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (release *Release) GongGetReferenceIdentifier(stage *Stage) string {
	return release.GongGetIdentifier(stage)
}

func (rest *Rest) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rest, rest.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rest *Rest) GongGetReferenceIdentifier(stage *Stage) string {
	return rest.GongGetIdentifier(stage)
}

func (slide *Slide) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(slide, slide.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (slide *Slide) GongGetReferenceIdentifier(stage *Stage) string {
	return slide.GongGetIdentifier(stage)
}

func (slur *Slur) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(slur, slur.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (slur *Slur) GongGetReferenceIdentifier(stage *Stage) string {
	return slur.GongGetIdentifier(stage)
}

func (stem *Stem) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stem, stem.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stem *Stem) GongGetReferenceIdentifier(stage *Stage) string {
	return stem.GongGetIdentifier(stage)
}

func (strongaccent *StrongAccent) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(strongaccent, strongaccent.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (strongaccent *StrongAccent) GongGetReferenceIdentifier(stage *Stage) string {
	return strongaccent.GongGetIdentifier(stage)
}

func (tap *Tap) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tap, tap.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tap *Tap) GongGetReferenceIdentifier(stage *Stage) string {
	return tap.GongGetIdentifier(stage)
}

func (technical *Technical) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(technical, technical.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (technical *Technical) GongGetReferenceIdentifier(stage *Stage) string {
	return technical.GongGetIdentifier(stage)
}

func (textelementdata *TextElementData) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(textelementdata, textelementdata.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (textelementdata *TextElementData) GongGetReferenceIdentifier(stage *Stage) string {
	return textelementdata.GongGetIdentifier(stage)
}

func (tie *Tie) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tie, tie.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tie *Tie) GongGetReferenceIdentifier(stage *Stage) string {
	return tie.GongGetIdentifier(stage)
}

func (tied *Tied) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tied, tied.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tied *Tied) GongGetReferenceIdentifier(stage *Stage) string {
	return tied.GongGetIdentifier(stage)
}

func (tremolo *Tremolo) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tremolo, tremolo.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tremolo *Tremolo) GongGetReferenceIdentifier(stage *Stage) string {
	return tremolo.GongGetIdentifier(stage)
}

func (tuplet *Tuplet) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tuplet, tuplet.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tuplet *Tuplet) GongGetReferenceIdentifier(stage *Stage) string {
	return tuplet.GongGetIdentifier(stage)
}

func (tupletdot *TupletDot) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tupletdot, tupletdot.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tupletdot *TupletDot) GongGetReferenceIdentifier(stage *Stage) string {
	return tupletdot.GongGetIdentifier(stage)
}

func (tupletnumber *TupletNumber) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tupletnumber, tupletnumber.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tupletnumber *TupletNumber) GongGetReferenceIdentifier(stage *Stage) string {
	return tupletnumber.GongGetIdentifier(stage)
}

func (tupletportion *TupletPortion) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tupletportion, tupletportion.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tupletportion *TupletPortion) GongGetReferenceIdentifier(stage *Stage) string {
	return tupletportion.GongGetIdentifier(stage)
}

func (tuplettype *TupletType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tuplettype, tuplettype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tuplettype *TupletType) GongGetReferenceIdentifier(stage *Stage) string {
	return tuplettype.GongGetIdentifier(stage)
}

func (unpitched *Unpitched) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(unpitched, unpitched.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (unpitched *Unpitched) GongGetReferenceIdentifier(stage *Stage) string {
	return unpitched.GongGetIdentifier(stage)
}

func (wait *Wait) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(wait, wait.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (wait *Wait) GongGetReferenceIdentifier(stage *Stage) string {
	return wait.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (accidental *Accidental) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(accidental.GongGetIdentifier(stage), "Accidental", accidental.Name)
}

func (accidentalmark *AccidentalMark) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(accidentalmark.GongGetIdentifier(stage), "AccidentalMark", accidentalmark.Name)
}

func (arpeggiate *Arpeggiate) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(arpeggiate.GongGetIdentifier(stage), "Arpeggiate", arpeggiate.Name)
}

func (arrow *Arrow) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(arrow.GongGetIdentifier(stage), "Arrow", arrow.Name)
}

func (articulations *Articulations) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(articulations.GongGetIdentifier(stage), "Articulations", articulations.Name)
}

func (assess *Assess) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(assess.GongGetIdentifier(stage), "Assess", assess.Name)
}

func (backup *Backup) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(backup.GongGetIdentifier(stage), "Backup", backup.Name)
}

func (beam *Beam) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(beam.GongGetIdentifier(stage), "Beam", beam.Name)
}

func (bend *Bend) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(bend.GongGetIdentifier(stage), "Bend", bend.Name)
}

func (breathmark *BreathMark) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(breathmark.GongGetIdentifier(stage), "BreathMark", breathmark.Name)
}

func (caesura *Caesura) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(caesura.GongGetIdentifier(stage), "Caesura", caesura.Name)
}

func (elision *Elision) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elision.GongGetIdentifier(stage), "Elision", elision.Name)
}

func (emptyline *EmptyLine) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(emptyline.GongGetIdentifier(stage), "EmptyLine", emptyline.Name)
}

func (extend *Extend) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(extend.GongGetIdentifier(stage), "Extend", extend.Name)
}

func (figure *Figure) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(figure.GongGetIdentifier(stage), "Figure", figure.Name)
}

func (figuredbass *FiguredBass) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(figuredbass.GongGetIdentifier(stage), "FiguredBass", figuredbass.Name)
}

func (forward *Forward) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(forward.GongGetIdentifier(stage), "Forward", forward.Name)
}

func (glissando *Glissando) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(glissando.GongGetIdentifier(stage), "Glissando", glissando.Name)
}

func (grace *Grace) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(grace.GongGetIdentifier(stage), "Grace", grace.Name)
}

func (hammeronpulloff *HammerOnPullOff) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(hammeronpulloff.GongGetIdentifier(stage), "HammerOnPullOff", hammeronpulloff.Name)
}

func (handbell *Handbell) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(handbell.GongGetIdentifier(stage), "Handbell", handbell.Name)
}

func (harmonclosed *HarmonClosed) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(harmonclosed.GongGetIdentifier(stage), "HarmonClosed", harmonclosed.Name)
}

func (harmonmute *HarmonMute) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(harmonmute.GongGetIdentifier(stage), "HarmonMute", harmonmute.Name)
}

func (harmonic *Harmonic) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(harmonic.GongGetIdentifier(stage), "Harmonic", harmonic.Name)
}

func (heeltoe *HeelToe) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(heeltoe.GongGetIdentifier(stage), "HeelToe", heeltoe.Name)
}

func (hole *Hole) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(hole.GongGetIdentifier(stage), "Hole", hole.Name)
}

func (holeclosed *HoleClosed) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(holeclosed.GongGetIdentifier(stage), "HoleClosed", holeclosed.Name)
}

func (instrument *Instrument) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(instrument.GongGetIdentifier(stage), "Instrument", instrument.Name)
}

func (listen *Listen) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(listen.GongGetIdentifier(stage), "Listen", listen.Name)
}

func (lyric *Lyric) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(lyric.GongGetIdentifier(stage), "Lyric", lyric.Name)
}

func (mordent *Mordent) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(mordent.GongGetIdentifier(stage), "Mordent", mordent.Name)
}

func (nonarpeggiate *NonArpeggiate) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(nonarpeggiate.GongGetIdentifier(stage), "NonArpeggiate", nonarpeggiate.Name)
}

func (notations *Notations) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notations.GongGetIdentifier(stage), "Notations", notations.Name)
}

func (note *Note) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(note.GongGetIdentifier(stage), "Note", note.Name)
}

func (notetype *NoteType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notetype.GongGetIdentifier(stage), "NoteType", notetype.Name)
}

func (notehead *Notehead) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notehead.GongGetIdentifier(stage), "Notehead", notehead.Name)
}

func (noteheadtext *NoteheadText) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(noteheadtext.GongGetIdentifier(stage), "NoteheadText", noteheadtext.Name)
}

func (ornaments *Ornaments) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ornaments.GongGetIdentifier(stage), "Ornaments", ornaments.Name)
}

func (othernotation *OtherNotation) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(othernotation.GongGetIdentifier(stage), "OtherNotation", othernotation.Name)
}

func (otherplacementtext *OtherPlacementText) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(otherplacementtext.GongGetIdentifier(stage), "OtherPlacementText", otherplacementtext.Name)
}

func (pitch *Pitch) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pitch.GongGetIdentifier(stage), "Pitch", pitch.Name)
}

func (placementtext *PlacementText) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(placementtext.GongGetIdentifier(stage), "PlacementText", placementtext.Name)
}

func (release *Release) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(release.GongGetIdentifier(stage), "Release", release.Name)
}

func (rest *Rest) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rest.GongGetIdentifier(stage), "Rest", rest.Name)
}

func (slide *Slide) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(slide.GongGetIdentifier(stage), "Slide", slide.Name)
}

func (slur *Slur) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(slur.GongGetIdentifier(stage), "Slur", slur.Name)
}

func (stem *Stem) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stem.GongGetIdentifier(stage), "Stem", stem.Name)
}

func (strongaccent *StrongAccent) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(strongaccent.GongGetIdentifier(stage), "StrongAccent", strongaccent.Name)
}

func (tap *Tap) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tap.GongGetIdentifier(stage), "Tap", tap.Name)
}

func (technical *Technical) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(technical.GongGetIdentifier(stage), "Technical", technical.Name)
}

func (textelementdata *TextElementData) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textelementdata.GongGetIdentifier(stage), "TextElementData", textelementdata.Name)
}

func (tie *Tie) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tie.GongGetIdentifier(stage), "Tie", tie.Name)
}

func (tied *Tied) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tied.GongGetIdentifier(stage), "Tied", tied.Name)
}

func (tremolo *Tremolo) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tremolo.GongGetIdentifier(stage), "Tremolo", tremolo.Name)
}

func (tuplet *Tuplet) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tuplet.GongGetIdentifier(stage), "Tuplet", tuplet.Name)
}

func (tupletdot *TupletDot) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tupletdot.GongGetIdentifier(stage), "TupletDot", tupletdot.Name)
}

func (tupletnumber *TupletNumber) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tupletnumber.GongGetIdentifier(stage), "TupletNumber", tupletnumber.Name)
}

func (tupletportion *TupletPortion) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tupletportion.GongGetIdentifier(stage), "TupletPortion", tupletportion.Name)
}

func (tuplettype *TupletType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tuplettype.GongGetIdentifier(stage), "TupletType", tuplettype.Name)
}

func (unpitched *Unpitched) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(unpitched.GongGetIdentifier(stage), "Unpitched", unpitched.Name)
}

func (wait *Wait) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(wait.GongGetIdentifier(stage), "Wait", wait.Name)
}

// insertion point for unstaging
func (accidental *Accidental) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(accidental.GongGetReferenceIdentifier(stage))
}

func (accidentalmark *AccidentalMark) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(accidentalmark.GongGetReferenceIdentifier(stage))
}

func (arpeggiate *Arpeggiate) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(arpeggiate.GongGetReferenceIdentifier(stage))
}

func (arrow *Arrow) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(arrow.GongGetReferenceIdentifier(stage))
}

func (articulations *Articulations) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(articulations.GongGetReferenceIdentifier(stage))
}

func (assess *Assess) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(assess.GongGetReferenceIdentifier(stage))
}

func (backup *Backup) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(backup.GongGetReferenceIdentifier(stage))
}

func (beam *Beam) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(beam.GongGetReferenceIdentifier(stage))
}

func (bend *Bend) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(bend.GongGetReferenceIdentifier(stage))
}

func (breathmark *BreathMark) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(breathmark.GongGetReferenceIdentifier(stage))
}

func (caesura *Caesura) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(caesura.GongGetReferenceIdentifier(stage))
}

func (elision *Elision) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elision.GongGetReferenceIdentifier(stage))
}

func (emptyline *EmptyLine) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(emptyline.GongGetReferenceIdentifier(stage))
}

func (extend *Extend) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(extend.GongGetReferenceIdentifier(stage))
}

func (figure *Figure) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(figure.GongGetReferenceIdentifier(stage))
}

func (figuredbass *FiguredBass) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(figuredbass.GongGetReferenceIdentifier(stage))
}

func (forward *Forward) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(forward.GongGetReferenceIdentifier(stage))
}

func (glissando *Glissando) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(glissando.GongGetReferenceIdentifier(stage))
}

func (grace *Grace) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(grace.GongGetReferenceIdentifier(stage))
}

func (hammeronpulloff *HammerOnPullOff) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(hammeronpulloff.GongGetReferenceIdentifier(stage))
}

func (handbell *Handbell) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(handbell.GongGetReferenceIdentifier(stage))
}

func (harmonclosed *HarmonClosed) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(harmonclosed.GongGetReferenceIdentifier(stage))
}

func (harmonmute *HarmonMute) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(harmonmute.GongGetReferenceIdentifier(stage))
}

func (harmonic *Harmonic) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(harmonic.GongGetReferenceIdentifier(stage))
}

func (heeltoe *HeelToe) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(heeltoe.GongGetReferenceIdentifier(stage))
}

func (hole *Hole) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(hole.GongGetReferenceIdentifier(stage))
}

func (holeclosed *HoleClosed) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(holeclosed.GongGetReferenceIdentifier(stage))
}

func (instrument *Instrument) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(instrument.GongGetReferenceIdentifier(stage))
}

func (listen *Listen) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(listen.GongGetReferenceIdentifier(stage))
}

func (lyric *Lyric) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(lyric.GongGetReferenceIdentifier(stage))
}

func (mordent *Mordent) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(mordent.GongGetReferenceIdentifier(stage))
}

func (nonarpeggiate *NonArpeggiate) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(nonarpeggiate.GongGetReferenceIdentifier(stage))
}

func (notations *Notations) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notations.GongGetReferenceIdentifier(stage))
}

func (note *Note) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(note.GongGetReferenceIdentifier(stage))
}

func (notetype *NoteType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notetype.GongGetReferenceIdentifier(stage))
}

func (notehead *Notehead) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notehead.GongGetReferenceIdentifier(stage))
}

func (noteheadtext *NoteheadText) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(noteheadtext.GongGetReferenceIdentifier(stage))
}

func (ornaments *Ornaments) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ornaments.GongGetReferenceIdentifier(stage))
}

func (othernotation *OtherNotation) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(othernotation.GongGetReferenceIdentifier(stage))
}

func (otherplacementtext *OtherPlacementText) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(otherplacementtext.GongGetReferenceIdentifier(stage))
}

func (pitch *Pitch) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pitch.GongGetReferenceIdentifier(stage))
}

func (placementtext *PlacementText) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(placementtext.GongGetReferenceIdentifier(stage))
}

func (release *Release) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(release.GongGetReferenceIdentifier(stage))
}

func (rest *Rest) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rest.GongGetReferenceIdentifier(stage))
}

func (slide *Slide) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(slide.GongGetReferenceIdentifier(stage))
}

func (slur *Slur) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(slur.GongGetReferenceIdentifier(stage))
}

func (stem *Stem) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stem.GongGetReferenceIdentifier(stage))
}

func (strongaccent *StrongAccent) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(strongaccent.GongGetReferenceIdentifier(stage))
}

func (tap *Tap) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tap.GongGetReferenceIdentifier(stage))
}

func (technical *Technical) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(technical.GongGetReferenceIdentifier(stage))
}

func (textelementdata *TextElementData) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textelementdata.GongGetReferenceIdentifier(stage))
}

func (tie *Tie) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tie.GongGetReferenceIdentifier(stage))
}

func (tied *Tied) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tied.GongGetReferenceIdentifier(stage))
}

func (tremolo *Tremolo) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tremolo.GongGetReferenceIdentifier(stage))
}

func (tuplet *Tuplet) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tuplet.GongGetReferenceIdentifier(stage))
}

func (tupletdot *TupletDot) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tupletdot.GongGetReferenceIdentifier(stage))
}

func (tupletnumber *TupletNumber) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tupletnumber.GongGetReferenceIdentifier(stage))
}

func (tupletportion *TupletPortion) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tupletportion.GongGetReferenceIdentifier(stage))
}

func (tuplettype *TupletType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tuplettype.GongGetReferenceIdentifier(stage))
}

func (unpitched *Unpitched) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(unpitched.GongGetReferenceIdentifier(stage))
}

func (wait *Wait) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(wait.GongGetReferenceIdentifier(stage))
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
