// generated code - do not edit
package common

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
	// Compute reverse map for named struct Appearance
	// insertion point per field
	stage.Appearance_LineWidth_reverseMap = make(map[*LineWidth]*Appearance)
	for appearance := range stage.Appearances {
		_ = appearance
		for _, _linewidth := range appearance.LineWidth {
			stage.Appearance_LineWidth_reverseMap[_linewidth] = appearance
		}
	}
	stage.Appearance_NoteSize_reverseMap = make(map[*NoteSize]*Appearance)
	for appearance := range stage.Appearances {
		_ = appearance
		for _, _notesize := range appearance.NoteSize {
			stage.Appearance_NoteSize_reverseMap[_notesize] = appearance
		}
	}
	stage.Appearance_Distance_reverseMap = make(map[*Distance]*Appearance)
	for appearance := range stage.Appearances {
		_ = appearance
		for _, _distance := range appearance.Distance {
			stage.Appearance_Distance_reverseMap[_distance] = appearance
		}
	}
	stage.Appearance_Glyph_reverseMap = make(map[*Glyph]*Appearance)
	for appearance := range stage.Appearances {
		_ = appearance
		for _, _glyph := range appearance.Glyph {
			stage.Appearance_Glyph_reverseMap[_glyph] = appearance
		}
	}
	stage.Appearance_OtherAppearance_reverseMap = make(map[*OtherAppearance]*Appearance)
	for appearance := range stage.Appearances {
		_ = appearance
		for _, _otherappearance := range appearance.OtherAppearance {
			stage.Appearance_OtherAppearance_reverseMap[_otherappearance] = appearance
		}
	}

	// Compute reverse map for named struct Dynamics
	// insertion point per field
	stage.Dynamics_P_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.P {
			stage.Dynamics_P_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Pp_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Pp {
			stage.Dynamics_Pp_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Ppp_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Ppp {
			stage.Dynamics_Ppp_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Pppp_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Pppp {
			stage.Dynamics_Pppp_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Ppppp_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Ppppp {
			stage.Dynamics_Ppppp_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Pppppp_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Pppppp {
			stage.Dynamics_Pppppp_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_F_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.F {
			stage.Dynamics_F_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Ff_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Ff {
			stage.Dynamics_Ff_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Fff_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Fff {
			stage.Dynamics_Fff_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Ffff_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Ffff {
			stage.Dynamics_Ffff_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Fffff_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Fffff {
			stage.Dynamics_Fffff_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Ffffff_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Ffffff {
			stage.Dynamics_Ffffff_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Mp_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Mp {
			stage.Dynamics_Mp_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Mf_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Mf {
			stage.Dynamics_Mf_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Sf_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Sf {
			stage.Dynamics_Sf_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Sfp_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Sfp {
			stage.Dynamics_Sfp_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Sfpp_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Sfpp {
			stage.Dynamics_Sfpp_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Fp_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Fp {
			stage.Dynamics_Fp_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Rf_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Rf {
			stage.Dynamics_Rf_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Rfz_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Rfz {
			stage.Dynamics_Rfz_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Sfz_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Sfz {
			stage.Dynamics_Sfz_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Sffz_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Sffz {
			stage.Dynamics_Sffz_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Fz_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Fz {
			stage.Dynamics_Fz_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_N_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.N {
			stage.Dynamics_N_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Pf_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Pf {
			stage.Dynamics_Pf_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_Sfzp_reverseMap = make(map[*Empty]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _empty := range dynamics.Sfzp {
			stage.Dynamics_Sfzp_reverseMap[_empty] = dynamics
		}
	}
	stage.Dynamics_OtherDynamics_reverseMap = make(map[*OtherText]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _othertext := range dynamics.OtherDynamics {
			stage.Dynamics_OtherDynamics_reverseMap[_othertext] = dynamics
		}
	}

	// Compute reverse map for named struct Encoding
	// insertion point per field
	stage.Encoding_Encoder_reverseMap = make(map[*TypedText]*Encoding)
	for encoding := range stage.Encodings {
		_ = encoding
		for _, _typedtext := range encoding.Encoder {
			stage.Encoding_Encoder_reverseMap[_typedtext] = encoding
		}
	}
	stage.Encoding_Supports_reverseMap = make(map[*Supports]*Encoding)
	for encoding := range stage.Encodings {
		_ = encoding
		for _, _supports := range encoding.Supports {
			stage.Encoding_Supports_reverseMap[_supports] = encoding
		}
	}

	// Compute reverse map for named struct Identification
	// insertion point per field
	stage.Identification_Creator_reverseMap = make(map[*TypedText]*Identification)
	for identification := range stage.Identifications {
		_ = identification
		for _, _typedtext := range identification.Creator {
			stage.Identification_Creator_reverseMap[_typedtext] = identification
		}
	}
	stage.Identification_Rights_reverseMap = make(map[*TypedText]*Identification)
	for identification := range stage.Identifications {
		_ = identification
		for _, _typedtext := range identification.Rights {
			stage.Identification_Rights_reverseMap[_typedtext] = identification
		}
	}
	stage.Identification_Relation_reverseMap = make(map[*TypedText]*Identification)
	for identification := range stage.Identifications {
		_ = identification
		for _, _typedtext := range identification.Relation {
			stage.Identification_Relation_reverseMap[_typedtext] = identification
		}
	}

	// Compute reverse map for named struct Miscellaneous
	// insertion point per field
	stage.Miscellaneous_MiscellaneousField_reverseMap = make(map[*MiscellaneousField]*Miscellaneous)
	for miscellaneous := range stage.Miscellaneouss {
		_ = miscellaneous
		for _, _miscellaneousfield := range miscellaneous.MiscellaneousField {
			stage.Miscellaneous_MiscellaneousField_reverseMap[_miscellaneousfield] = miscellaneous
		}
	}

	// Compute reverse map for named struct NameDisplay
	// insertion point per field
	stage.NameDisplay_DisplayText_reverseMap = make(map[*FormattedText]*NameDisplay)
	for namedisplay := range stage.NameDisplays {
		_ = namedisplay
		for _, _formattedtext := range namedisplay.DisplayText {
			stage.NameDisplay_DisplayText_reverseMap[_formattedtext] = namedisplay
		}
	}
	stage.NameDisplay_AccidentalText_reverseMap = make(map[*AccidentalText]*NameDisplay)
	for namedisplay := range stage.NameDisplays {
		_ = namedisplay
		for _, _accidentaltext := range namedisplay.AccidentalText {
			stage.NameDisplay_AccidentalText_reverseMap[_accidentaltext] = namedisplay
		}
	}

	// Compute reverse map for named struct PageLayout
	// insertion point per field
	stage.PageLayout_PageMargins_reverseMap = make(map[*PageMargins]*PageLayout)
	for pagelayout := range stage.PageLayouts {
		_ = pagelayout
		for _, _pagemargins := range pagelayout.PageMargins {
			stage.PageLayout_PageMargins_reverseMap[_pagemargins] = pagelayout
		}
	}

	// Compute reverse map for named struct Play
	// insertion point per field
	stage.Play_OtherPlay_reverseMap = make(map[*OtherPlay]*Play)
	for play := range stage.Plays {
		_ = play
		for _, _otherplay := range play.OtherPlay {
			stage.Play_OtherPlay_reverseMap[_otherplay] = play
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.AccidentalTexts)

	res = __gong__appendInstances(res, stage.Appearances)

	res = __gong__appendInstances(res, stage.Bookmarks)

	res = __gong__appendInstances(res, stage.Codas)

	res = __gong__appendInstances(res, stage.Distances)

	res = __gong__appendInstances(res, stage.Dynamicss)

	res = __gong__appendInstances(res, stage.Emptys)

	res = __gong__appendInstances(res, stage.EmptyPlacements)

	res = __gong__appendInstances(res, stage.EmptyPlacementSmufls)

	res = __gong__appendInstances(res, stage.EmptyPrintObjectStyleAligns)

	res = __gong__appendInstances(res, stage.EmptyPrintStyles)

	res = __gong__appendInstances(res, stage.EmptyPrintStyleAligns)

	res = __gong__appendInstances(res, stage.EmptyPrintStyleAlignIds)

	res = __gong__appendInstances(res, stage.EmptyTrillSounds)

	res = __gong__appendInstances(res, stage.Encodings)

	res = __gong__appendInstances(res, stage.Fermatas)

	res = __gong__appendInstances(res, stage.Fingerings)

	res = __gong__appendInstances(res, stage.FormattedSymbols)

	res = __gong__appendInstances(res, stage.FormattedSymbolIds)

	res = __gong__appendInstances(res, stage.FormattedTexts)

	res = __gong__appendInstances(res, stage.FormattedTextIds)

	res = __gong__appendInstances(res, stage.Frets)

	res = __gong__appendInstances(res, stage.Glyphs)

	res = __gong__appendInstances(res, stage.HorizontalTurns)

	res = __gong__appendInstances(res, stage.Identifications)

	res = __gong__appendInstances(res, stage.Levels)

	res = __gong__appendInstances(res, stage.LineWidths)

	res = __gong__appendInstances(res, stage.Links)

	res = __gong__appendInstances(res, stage.MeasureLayouts)

	res = __gong__appendInstances(res, stage.MidiDevices)

	res = __gong__appendInstances(res, stage.MidiInstruments)

	res = __gong__appendInstances(res, stage.Miscellaneouss)

	res = __gong__appendInstances(res, stage.MiscellaneousFields)

	res = __gong__appendInstances(res, stage.NameDisplays)

	res = __gong__appendInstances(res, stage.NoteSizes)

	res = __gong__appendInstances(res, stage.OtherAppearances)

	res = __gong__appendInstances(res, stage.OtherPlays)

	res = __gong__appendInstances(res, stage.OtherTexts)

	res = __gong__appendInstances(res, stage.PageLayouts)

	res = __gong__appendInstances(res, stage.PageMarginss)

	res = __gong__appendInstances(res, stage.Plays)

	res = __gong__appendInstances(res, stage.Scalings)

	res = __gong__appendInstances(res, stage.Segnos)

	res = __gong__appendInstances(res, stage.StaffLayouts)

	res = __gong__appendInstances(res, stage.StringTypes)

	res = __gong__appendInstances(res, stage.Supportss)

	res = __gong__appendInstances(res, stage.SystemDividerss)

	res = __gong__appendInstances(res, stage.SystemLayouts)

	res = __gong__appendInstances(res, stage.SystemMarginss)

	res = __gong__appendInstances(res, stage.TypedTexts)

	res = __gong__appendInstances(res, stage.WavyLines)

	return
}

// insertion point per named struct
func (accidentaltext *AccidentalText) GongCopy() GongstructIF {
	newInstance := new(AccidentalText)
	accidentaltext.GongCopyBasicFields(newInstance)
	return newInstance
}

func (appearance *Appearance) GongCopy() GongstructIF {
	newInstance := new(Appearance)
	appearance.GongCopyBasicFields(newInstance)
	return newInstance
}

func (bookmark *Bookmark) GongCopy() GongstructIF {
	newInstance := new(Bookmark)
	bookmark.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coda *Coda) GongCopy() GongstructIF {
	newInstance := new(Coda)
	coda.GongCopyBasicFields(newInstance)
	return newInstance
}

func (distance *Distance) GongCopy() GongstructIF {
	newInstance := new(Distance)
	distance.GongCopyBasicFields(newInstance)
	return newInstance
}

func (dynamics *Dynamics) GongCopy() GongstructIF {
	newInstance := new(Dynamics)
	dynamics.GongCopyBasicFields(newInstance)
	return newInstance
}

func (empty *Empty) GongCopy() GongstructIF {
	newInstance := new(Empty)
	empty.GongCopyBasicFields(newInstance)
	return newInstance
}

func (emptyplacement *EmptyPlacement) GongCopy() GongstructIF {
	newInstance := new(EmptyPlacement)
	emptyplacement.GongCopyBasicFields(newInstance)
	return newInstance
}

func (emptyplacementsmufl *EmptyPlacementSmufl) GongCopy() GongstructIF {
	newInstance := new(EmptyPlacementSmufl)
	emptyplacementsmufl.GongCopyBasicFields(newInstance)
	return newInstance
}

func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) GongCopy() GongstructIF {
	newInstance := new(EmptyPrintObjectStyleAlign)
	emptyprintobjectstylealign.GongCopyBasicFields(newInstance)
	return newInstance
}

func (emptyprintstyle *EmptyPrintStyle) GongCopy() GongstructIF {
	newInstance := new(EmptyPrintStyle)
	emptyprintstyle.GongCopyBasicFields(newInstance)
	return newInstance
}

func (emptyprintstylealign *EmptyPrintStyleAlign) GongCopy() GongstructIF {
	newInstance := new(EmptyPrintStyleAlign)
	emptyprintstylealign.GongCopyBasicFields(newInstance)
	return newInstance
}

func (emptyprintstylealignid *EmptyPrintStyleAlignId) GongCopy() GongstructIF {
	newInstance := new(EmptyPrintStyleAlignId)
	emptyprintstylealignid.GongCopyBasicFields(newInstance)
	return newInstance
}

func (emptytrillsound *EmptyTrillSound) GongCopy() GongstructIF {
	newInstance := new(EmptyTrillSound)
	emptytrillsound.GongCopyBasicFields(newInstance)
	return newInstance
}

func (encoding *Encoding) GongCopy() GongstructIF {
	newInstance := new(Encoding)
	encoding.GongCopyBasicFields(newInstance)
	return newInstance
}

func (fermata *Fermata) GongCopy() GongstructIF {
	newInstance := new(Fermata)
	fermata.GongCopyBasicFields(newInstance)
	return newInstance
}

func (fingering *Fingering) GongCopy() GongstructIF {
	newInstance := new(Fingering)
	fingering.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formattedsymbol *FormattedSymbol) GongCopy() GongstructIF {
	newInstance := new(FormattedSymbol)
	formattedsymbol.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formattedsymbolid *FormattedSymbolId) GongCopy() GongstructIF {
	newInstance := new(FormattedSymbolId)
	formattedsymbolid.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formattedtext *FormattedText) GongCopy() GongstructIF {
	newInstance := new(FormattedText)
	formattedtext.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formattedtextid *FormattedTextId) GongCopy() GongstructIF {
	newInstance := new(FormattedTextId)
	formattedtextid.GongCopyBasicFields(newInstance)
	return newInstance
}

func (fret *Fret) GongCopy() GongstructIF {
	newInstance := new(Fret)
	fret.GongCopyBasicFields(newInstance)
	return newInstance
}

func (glyph *Glyph) GongCopy() GongstructIF {
	newInstance := new(Glyph)
	glyph.GongCopyBasicFields(newInstance)
	return newInstance
}

func (horizontalturn *HorizontalTurn) GongCopy() GongstructIF {
	newInstance := new(HorizontalTurn)
	horizontalturn.GongCopyBasicFields(newInstance)
	return newInstance
}

func (identification *Identification) GongCopy() GongstructIF {
	newInstance := new(Identification)
	identification.GongCopyBasicFields(newInstance)
	return newInstance
}

func (level *Level) GongCopy() GongstructIF {
	newInstance := new(Level)
	level.GongCopyBasicFields(newInstance)
	return newInstance
}

func (linewidth *LineWidth) GongCopy() GongstructIF {
	newInstance := new(LineWidth)
	linewidth.GongCopyBasicFields(newInstance)
	return newInstance
}

func (link *Link) GongCopy() GongstructIF {
	newInstance := new(Link)
	link.GongCopyBasicFields(newInstance)
	return newInstance
}

func (measurelayout *MeasureLayout) GongCopy() GongstructIF {
	newInstance := new(MeasureLayout)
	measurelayout.GongCopyBasicFields(newInstance)
	return newInstance
}

func (mididevice *MidiDevice) GongCopy() GongstructIF {
	newInstance := new(MidiDevice)
	mididevice.GongCopyBasicFields(newInstance)
	return newInstance
}

func (midiinstrument *MidiInstrument) GongCopy() GongstructIF {
	newInstance := new(MidiInstrument)
	midiinstrument.GongCopyBasicFields(newInstance)
	return newInstance
}

func (miscellaneous *Miscellaneous) GongCopy() GongstructIF {
	newInstance := new(Miscellaneous)
	miscellaneous.GongCopyBasicFields(newInstance)
	return newInstance
}

func (miscellaneousfield *MiscellaneousField) GongCopy() GongstructIF {
	newInstance := new(MiscellaneousField)
	miscellaneousfield.GongCopyBasicFields(newInstance)
	return newInstance
}

func (namedisplay *NameDisplay) GongCopy() GongstructIF {
	newInstance := new(NameDisplay)
	namedisplay.GongCopyBasicFields(newInstance)
	return newInstance
}

func (notesize *NoteSize) GongCopy() GongstructIF {
	newInstance := new(NoteSize)
	notesize.GongCopyBasicFields(newInstance)
	return newInstance
}

func (otherappearance *OtherAppearance) GongCopy() GongstructIF {
	newInstance := new(OtherAppearance)
	otherappearance.GongCopyBasicFields(newInstance)
	return newInstance
}

func (otherplay *OtherPlay) GongCopy() GongstructIF {
	newInstance := new(OtherPlay)
	otherplay.GongCopyBasicFields(newInstance)
	return newInstance
}

func (othertext *OtherText) GongCopy() GongstructIF {
	newInstance := new(OtherText)
	othertext.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pagelayout *PageLayout) GongCopy() GongstructIF {
	newInstance := new(PageLayout)
	pagelayout.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pagemargins *PageMargins) GongCopy() GongstructIF {
	newInstance := new(PageMargins)
	pagemargins.GongCopyBasicFields(newInstance)
	return newInstance
}

func (play *Play) GongCopy() GongstructIF {
	newInstance := new(Play)
	play.GongCopyBasicFields(newInstance)
	return newInstance
}

func (scaling *Scaling) GongCopy() GongstructIF {
	newInstance := new(Scaling)
	scaling.GongCopyBasicFields(newInstance)
	return newInstance
}

func (segno *Segno) GongCopy() GongstructIF {
	newInstance := new(Segno)
	segno.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stafflayout *StaffLayout) GongCopy() GongstructIF {
	newInstance := new(StaffLayout)
	stafflayout.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stringtype *StringType) GongCopy() GongstructIF {
	newInstance := new(StringType)
	stringtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (supports *Supports) GongCopy() GongstructIF {
	newInstance := new(Supports)
	supports.GongCopyBasicFields(newInstance)
	return newInstance
}

func (systemdividers *SystemDividers) GongCopy() GongstructIF {
	newInstance := new(SystemDividers)
	systemdividers.GongCopyBasicFields(newInstance)
	return newInstance
}

func (systemlayout *SystemLayout) GongCopy() GongstructIF {
	newInstance := new(SystemLayout)
	systemlayout.GongCopyBasicFields(newInstance)
	return newInstance
}

func (systemmargins *SystemMargins) GongCopy() GongstructIF {
	newInstance := new(SystemMargins)
	systemmargins.GongCopyBasicFields(newInstance)
	return newInstance
}

func (typedtext *TypedText) GongCopy() GongstructIF {
	newInstance := new(TypedText)
	typedtext.GongCopyBasicFields(newInstance)
	return newInstance
}

func (wavyline *WavyLine) GongCopy() GongstructIF {
	newInstance := new(WavyLine)
	wavyline.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (accidentaltext *AccidentalText) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, accidentaltext)
}

func (appearance *Appearance) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, appearance)
}

func (bookmark *Bookmark) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, bookmark)
}

func (coda *Coda) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coda)
}

func (distance *Distance) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, distance)
}

func (dynamics *Dynamics) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, dynamics)
}

func (empty *Empty) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, empty)
}

func (emptyplacement *EmptyPlacement) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, emptyplacement)
}

func (emptyplacementsmufl *EmptyPlacementSmufl) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, emptyplacementsmufl)
}

func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, emptyprintobjectstylealign)
}

func (emptyprintstyle *EmptyPrintStyle) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, emptyprintstyle)
}

func (emptyprintstylealign *EmptyPrintStyleAlign) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, emptyprintstylealign)
}

func (emptyprintstylealignid *EmptyPrintStyleAlignId) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, emptyprintstylealignid)
}

func (emptytrillsound *EmptyTrillSound) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, emptytrillsound)
}

func (encoding *Encoding) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, encoding)
}

func (fermata *Fermata) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, fermata)
}

func (fingering *Fingering) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, fingering)
}

func (formattedsymbol *FormattedSymbol) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formattedsymbol)
}

func (formattedsymbolid *FormattedSymbolId) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formattedsymbolid)
}

func (formattedtext *FormattedText) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formattedtext)
}

func (formattedtextid *FormattedTextId) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formattedtextid)
}

func (fret *Fret) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, fret)
}

func (glyph *Glyph) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, glyph)
}

func (horizontalturn *HorizontalTurn) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, horizontalturn)
}

func (identification *Identification) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, identification)
}

func (level *Level) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, level)
}

func (linewidth *LineWidth) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, linewidth)
}

func (link *Link) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, link)
}

func (measurelayout *MeasureLayout) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, measurelayout)
}

func (mididevice *MidiDevice) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, mididevice)
}

func (midiinstrument *MidiInstrument) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, midiinstrument)
}

func (miscellaneous *Miscellaneous) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, miscellaneous)
}

func (miscellaneousfield *MiscellaneousField) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, miscellaneousfield)
}

func (namedisplay *NameDisplay) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, namedisplay)
}

func (notesize *NoteSize) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notesize)
}

func (otherappearance *OtherAppearance) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, otherappearance)
}

func (otherplay *OtherPlay) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, otherplay)
}

func (othertext *OtherText) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, othertext)
}

func (pagelayout *PageLayout) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pagelayout)
}

func (pagemargins *PageMargins) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pagemargins)
}

func (play *Play) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, play)
}

func (scaling *Scaling) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, scaling)
}

func (segno *Segno) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, segno)
}

func (stafflayout *StaffLayout) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stafflayout)
}

func (stringtype *StringType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stringtype)
}

func (supports *Supports) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, supports)
}

func (systemdividers *SystemDividers) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, systemdividers)
}

func (systemlayout *SystemLayout) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, systemlayout)
}

func (systemmargins *SystemMargins) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, systemmargins)
}

func (typedtext *TypedText) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, typedtext)
}

func (wavyline *WavyLine) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, wavyline)
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
		stage.AccidentalTexts,
		stage.AccidentalText_stagedOrder,
		stage.AccidentalTexts_reference,
		&stage.AccidentalTexts_referenceOrder,
		stage.AccidentalTexts_instance,
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
		stage.Appearances,
		stage.Appearance_stagedOrder,
		stage.Appearances_reference,
		&stage.Appearances_referenceOrder,
		stage.Appearances_instance,
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
		stage.Bookmarks,
		stage.Bookmark_stagedOrder,
		stage.Bookmarks_reference,
		&stage.Bookmarks_referenceOrder,
		stage.Bookmarks_instance,
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
		stage.Codas,
		stage.Coda_stagedOrder,
		stage.Codas_reference,
		&stage.Codas_referenceOrder,
		stage.Codas_instance,
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
		stage.Distances,
		stage.Distance_stagedOrder,
		stage.Distances_reference,
		&stage.Distances_referenceOrder,
		stage.Distances_instance,
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
		stage.Dynamicss,
		stage.Dynamics_stagedOrder,
		stage.Dynamicss_reference,
		&stage.Dynamicss_referenceOrder,
		stage.Dynamicss_instance,
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
		stage.Emptys,
		stage.Empty_stagedOrder,
		stage.Emptys_reference,
		&stage.Emptys_referenceOrder,
		stage.Emptys_instance,
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
		stage.EmptyPlacements,
		stage.EmptyPlacement_stagedOrder,
		stage.EmptyPlacements_reference,
		&stage.EmptyPlacements_referenceOrder,
		stage.EmptyPlacements_instance,
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
		stage.EmptyPlacementSmufls,
		stage.EmptyPlacementSmufl_stagedOrder,
		stage.EmptyPlacementSmufls_reference,
		&stage.EmptyPlacementSmufls_referenceOrder,
		stage.EmptyPlacementSmufls_instance,
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
		stage.EmptyPrintObjectStyleAligns,
		stage.EmptyPrintObjectStyleAlign_stagedOrder,
		stage.EmptyPrintObjectStyleAligns_reference,
		&stage.EmptyPrintObjectStyleAligns_referenceOrder,
		stage.EmptyPrintObjectStyleAligns_instance,
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
		stage.EmptyPrintStyles,
		stage.EmptyPrintStyle_stagedOrder,
		stage.EmptyPrintStyles_reference,
		&stage.EmptyPrintStyles_referenceOrder,
		stage.EmptyPrintStyles_instance,
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
		stage.EmptyPrintStyleAligns,
		stage.EmptyPrintStyleAlign_stagedOrder,
		stage.EmptyPrintStyleAligns_reference,
		&stage.EmptyPrintStyleAligns_referenceOrder,
		stage.EmptyPrintStyleAligns_instance,
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
		stage.EmptyPrintStyleAlignIds,
		stage.EmptyPrintStyleAlignId_stagedOrder,
		stage.EmptyPrintStyleAlignIds_reference,
		&stage.EmptyPrintStyleAlignIds_referenceOrder,
		stage.EmptyPrintStyleAlignIds_instance,
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
		stage.EmptyTrillSounds,
		stage.EmptyTrillSound_stagedOrder,
		stage.EmptyTrillSounds_reference,
		&stage.EmptyTrillSounds_referenceOrder,
		stage.EmptyTrillSounds_instance,
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
		stage.Encodings,
		stage.Encoding_stagedOrder,
		stage.Encodings_reference,
		&stage.Encodings_referenceOrder,
		stage.Encodings_instance,
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
		stage.Fermatas,
		stage.Fermata_stagedOrder,
		stage.Fermatas_reference,
		&stage.Fermatas_referenceOrder,
		stage.Fermatas_instance,
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
		stage.Fingerings,
		stage.Fingering_stagedOrder,
		stage.Fingerings_reference,
		&stage.Fingerings_referenceOrder,
		stage.Fingerings_instance,
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
		stage.FormattedSymbols,
		stage.FormattedSymbol_stagedOrder,
		stage.FormattedSymbols_reference,
		&stage.FormattedSymbols_referenceOrder,
		stage.FormattedSymbols_instance,
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
		stage.FormattedSymbolIds,
		stage.FormattedSymbolId_stagedOrder,
		stage.FormattedSymbolIds_reference,
		&stage.FormattedSymbolIds_referenceOrder,
		stage.FormattedSymbolIds_instance,
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
		stage.FormattedTexts,
		stage.FormattedText_stagedOrder,
		stage.FormattedTexts_reference,
		&stage.FormattedTexts_referenceOrder,
		stage.FormattedTexts_instance,
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
		stage.FormattedTextIds,
		stage.FormattedTextId_stagedOrder,
		stage.FormattedTextIds_reference,
		&stage.FormattedTextIds_referenceOrder,
		stage.FormattedTextIds_instance,
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
		stage.Frets,
		stage.Fret_stagedOrder,
		stage.Frets_reference,
		&stage.Frets_referenceOrder,
		stage.Frets_instance,
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
		stage.Glyphs,
		stage.Glyph_stagedOrder,
		stage.Glyphs_reference,
		&stage.Glyphs_referenceOrder,
		stage.Glyphs_instance,
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
		stage.HorizontalTurns,
		stage.HorizontalTurn_stagedOrder,
		stage.HorizontalTurns_reference,
		&stage.HorizontalTurns_referenceOrder,
		stage.HorizontalTurns_instance,
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
		stage.Identifications,
		stage.Identification_stagedOrder,
		stage.Identifications_reference,
		&stage.Identifications_referenceOrder,
		stage.Identifications_instance,
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
		stage.Levels,
		stage.Level_stagedOrder,
		stage.Levels_reference,
		&stage.Levels_referenceOrder,
		stage.Levels_instance,
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
		stage.LineWidths,
		stage.LineWidth_stagedOrder,
		stage.LineWidths_reference,
		&stage.LineWidths_referenceOrder,
		stage.LineWidths_instance,
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
		stage.Links,
		stage.Link_stagedOrder,
		stage.Links_reference,
		&stage.Links_referenceOrder,
		stage.Links_instance,
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
		stage.MeasureLayouts,
		stage.MeasureLayout_stagedOrder,
		stage.MeasureLayouts_reference,
		&stage.MeasureLayouts_referenceOrder,
		stage.MeasureLayouts_instance,
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
		stage.MidiDevices,
		stage.MidiDevice_stagedOrder,
		stage.MidiDevices_reference,
		&stage.MidiDevices_referenceOrder,
		stage.MidiDevices_instance,
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
		stage.MidiInstruments,
		stage.MidiInstrument_stagedOrder,
		stage.MidiInstruments_reference,
		&stage.MidiInstruments_referenceOrder,
		stage.MidiInstruments_instance,
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
		stage.Miscellaneouss,
		stage.Miscellaneous_stagedOrder,
		stage.Miscellaneouss_reference,
		&stage.Miscellaneouss_referenceOrder,
		stage.Miscellaneouss_instance,
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
		stage.MiscellaneousFields,
		stage.MiscellaneousField_stagedOrder,
		stage.MiscellaneousFields_reference,
		&stage.MiscellaneousFields_referenceOrder,
		stage.MiscellaneousFields_instance,
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
		stage.NameDisplays,
		stage.NameDisplay_stagedOrder,
		stage.NameDisplays_reference,
		&stage.NameDisplays_referenceOrder,
		stage.NameDisplays_instance,
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
		stage.NoteSizes,
		stage.NoteSize_stagedOrder,
		stage.NoteSizes_reference,
		&stage.NoteSizes_referenceOrder,
		stage.NoteSizes_instance,
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
		stage.OtherAppearances,
		stage.OtherAppearance_stagedOrder,
		stage.OtherAppearances_reference,
		&stage.OtherAppearances_referenceOrder,
		stage.OtherAppearances_instance,
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
		stage.OtherPlays,
		stage.OtherPlay_stagedOrder,
		stage.OtherPlays_reference,
		&stage.OtherPlays_referenceOrder,
		stage.OtherPlays_instance,
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
		stage.OtherTexts,
		stage.OtherText_stagedOrder,
		stage.OtherTexts_reference,
		&stage.OtherTexts_referenceOrder,
		stage.OtherTexts_instance,
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
		stage.PageLayouts,
		stage.PageLayout_stagedOrder,
		stage.PageLayouts_reference,
		&stage.PageLayouts_referenceOrder,
		stage.PageLayouts_instance,
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
		stage.PageMarginss,
		stage.PageMargins_stagedOrder,
		stage.PageMarginss_reference,
		&stage.PageMarginss_referenceOrder,
		stage.PageMarginss_instance,
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
		stage.Plays,
		stage.Play_stagedOrder,
		stage.Plays_reference,
		&stage.Plays_referenceOrder,
		stage.Plays_instance,
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
		stage.Scalings,
		stage.Scaling_stagedOrder,
		stage.Scalings_reference,
		&stage.Scalings_referenceOrder,
		stage.Scalings_instance,
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
		stage.Segnos,
		stage.Segno_stagedOrder,
		stage.Segnos_reference,
		&stage.Segnos_referenceOrder,
		stage.Segnos_instance,
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
		stage.StaffLayouts,
		stage.StaffLayout_stagedOrder,
		stage.StaffLayouts_reference,
		&stage.StaffLayouts_referenceOrder,
		stage.StaffLayouts_instance,
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
		stage.StringTypes,
		stage.StringType_stagedOrder,
		stage.StringTypes_reference,
		&stage.StringTypes_referenceOrder,
		stage.StringTypes_instance,
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
		stage.Supportss,
		stage.Supports_stagedOrder,
		stage.Supportss_reference,
		&stage.Supportss_referenceOrder,
		stage.Supportss_instance,
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
		stage.SystemDividerss,
		stage.SystemDividers_stagedOrder,
		stage.SystemDividerss_reference,
		&stage.SystemDividerss_referenceOrder,
		stage.SystemDividerss_instance,
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
		stage.SystemLayouts,
		stage.SystemLayout_stagedOrder,
		stage.SystemLayouts_reference,
		&stage.SystemLayouts_referenceOrder,
		stage.SystemLayouts_instance,
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
		stage.SystemMarginss,
		stage.SystemMargins_stagedOrder,
		stage.SystemMarginss_reference,
		&stage.SystemMarginss_referenceOrder,
		stage.SystemMarginss_instance,
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
		stage.TypedTexts,
		stage.TypedText_stagedOrder,
		stage.TypedTexts_reference,
		&stage.TypedTexts_referenceOrder,
		stage.TypedTexts_instance,
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
		stage.WavyLines,
		stage.WavyLine_stagedOrder,
		stage.WavyLines_reference,
		&stage.WavyLines_referenceOrder,
		stage.WavyLines_instance,
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
	__gong__computeReferencePass1(stage, stage.AccidentalTexts, &stage.AccidentalTexts_reference, &stage.AccidentalTexts_referenceOrder, &stage.AccidentalTexts_instance)

	__gong__computeReferencePass1(stage, stage.Appearances, &stage.Appearances_reference, &stage.Appearances_referenceOrder, &stage.Appearances_instance)

	__gong__computeReferencePass1(stage, stage.Bookmarks, &stage.Bookmarks_reference, &stage.Bookmarks_referenceOrder, &stage.Bookmarks_instance)

	__gong__computeReferencePass1(stage, stage.Codas, &stage.Codas_reference, &stage.Codas_referenceOrder, &stage.Codas_instance)

	__gong__computeReferencePass1(stage, stage.Distances, &stage.Distances_reference, &stage.Distances_referenceOrder, &stage.Distances_instance)

	__gong__computeReferencePass1(stage, stage.Dynamicss, &stage.Dynamicss_reference, &stage.Dynamicss_referenceOrder, &stage.Dynamicss_instance)

	__gong__computeReferencePass1(stage, stage.Emptys, &stage.Emptys_reference, &stage.Emptys_referenceOrder, &stage.Emptys_instance)

	__gong__computeReferencePass1(stage, stage.EmptyPlacements, &stage.EmptyPlacements_reference, &stage.EmptyPlacements_referenceOrder, &stage.EmptyPlacements_instance)

	__gong__computeReferencePass1(stage, stage.EmptyPlacementSmufls, &stage.EmptyPlacementSmufls_reference, &stage.EmptyPlacementSmufls_referenceOrder, &stage.EmptyPlacementSmufls_instance)

	__gong__computeReferencePass1(stage, stage.EmptyPrintObjectStyleAligns, &stage.EmptyPrintObjectStyleAligns_reference, &stage.EmptyPrintObjectStyleAligns_referenceOrder, &stage.EmptyPrintObjectStyleAligns_instance)

	__gong__computeReferencePass1(stage, stage.EmptyPrintStyles, &stage.EmptyPrintStyles_reference, &stage.EmptyPrintStyles_referenceOrder, &stage.EmptyPrintStyles_instance)

	__gong__computeReferencePass1(stage, stage.EmptyPrintStyleAligns, &stage.EmptyPrintStyleAligns_reference, &stage.EmptyPrintStyleAligns_referenceOrder, &stage.EmptyPrintStyleAligns_instance)

	__gong__computeReferencePass1(stage, stage.EmptyPrintStyleAlignIds, &stage.EmptyPrintStyleAlignIds_reference, &stage.EmptyPrintStyleAlignIds_referenceOrder, &stage.EmptyPrintStyleAlignIds_instance)

	__gong__computeReferencePass1(stage, stage.EmptyTrillSounds, &stage.EmptyTrillSounds_reference, &stage.EmptyTrillSounds_referenceOrder, &stage.EmptyTrillSounds_instance)

	__gong__computeReferencePass1(stage, stage.Encodings, &stage.Encodings_reference, &stage.Encodings_referenceOrder, &stage.Encodings_instance)

	__gong__computeReferencePass1(stage, stage.Fermatas, &stage.Fermatas_reference, &stage.Fermatas_referenceOrder, &stage.Fermatas_instance)

	__gong__computeReferencePass1(stage, stage.Fingerings, &stage.Fingerings_reference, &stage.Fingerings_referenceOrder, &stage.Fingerings_instance)

	__gong__computeReferencePass1(stage, stage.FormattedSymbols, &stage.FormattedSymbols_reference, &stage.FormattedSymbols_referenceOrder, &stage.FormattedSymbols_instance)

	__gong__computeReferencePass1(stage, stage.FormattedSymbolIds, &stage.FormattedSymbolIds_reference, &stage.FormattedSymbolIds_referenceOrder, &stage.FormattedSymbolIds_instance)

	__gong__computeReferencePass1(stage, stage.FormattedTexts, &stage.FormattedTexts_reference, &stage.FormattedTexts_referenceOrder, &stage.FormattedTexts_instance)

	__gong__computeReferencePass1(stage, stage.FormattedTextIds, &stage.FormattedTextIds_reference, &stage.FormattedTextIds_referenceOrder, &stage.FormattedTextIds_instance)

	__gong__computeReferencePass1(stage, stage.Frets, &stage.Frets_reference, &stage.Frets_referenceOrder, &stage.Frets_instance)

	__gong__computeReferencePass1(stage, stage.Glyphs, &stage.Glyphs_reference, &stage.Glyphs_referenceOrder, &stage.Glyphs_instance)

	__gong__computeReferencePass1(stage, stage.HorizontalTurns, &stage.HorizontalTurns_reference, &stage.HorizontalTurns_referenceOrder, &stage.HorizontalTurns_instance)

	__gong__computeReferencePass1(stage, stage.Identifications, &stage.Identifications_reference, &stage.Identifications_referenceOrder, &stage.Identifications_instance)

	__gong__computeReferencePass1(stage, stage.Levels, &stage.Levels_reference, &stage.Levels_referenceOrder, &stage.Levels_instance)

	__gong__computeReferencePass1(stage, stage.LineWidths, &stage.LineWidths_reference, &stage.LineWidths_referenceOrder, &stage.LineWidths_instance)

	__gong__computeReferencePass1(stage, stage.Links, &stage.Links_reference, &stage.Links_referenceOrder, &stage.Links_instance)

	__gong__computeReferencePass1(stage, stage.MeasureLayouts, &stage.MeasureLayouts_reference, &stage.MeasureLayouts_referenceOrder, &stage.MeasureLayouts_instance)

	__gong__computeReferencePass1(stage, stage.MidiDevices, &stage.MidiDevices_reference, &stage.MidiDevices_referenceOrder, &stage.MidiDevices_instance)

	__gong__computeReferencePass1(stage, stage.MidiInstruments, &stage.MidiInstruments_reference, &stage.MidiInstruments_referenceOrder, &stage.MidiInstruments_instance)

	__gong__computeReferencePass1(stage, stage.Miscellaneouss, &stage.Miscellaneouss_reference, &stage.Miscellaneouss_referenceOrder, &stage.Miscellaneouss_instance)

	__gong__computeReferencePass1(stage, stage.MiscellaneousFields, &stage.MiscellaneousFields_reference, &stage.MiscellaneousFields_referenceOrder, &stage.MiscellaneousFields_instance)

	__gong__computeReferencePass1(stage, stage.NameDisplays, &stage.NameDisplays_reference, &stage.NameDisplays_referenceOrder, &stage.NameDisplays_instance)

	__gong__computeReferencePass1(stage, stage.NoteSizes, &stage.NoteSizes_reference, &stage.NoteSizes_referenceOrder, &stage.NoteSizes_instance)

	__gong__computeReferencePass1(stage, stage.OtherAppearances, &stage.OtherAppearances_reference, &stage.OtherAppearances_referenceOrder, &stage.OtherAppearances_instance)

	__gong__computeReferencePass1(stage, stage.OtherPlays, &stage.OtherPlays_reference, &stage.OtherPlays_referenceOrder, &stage.OtherPlays_instance)

	__gong__computeReferencePass1(stage, stage.OtherTexts, &stage.OtherTexts_reference, &stage.OtherTexts_referenceOrder, &stage.OtherTexts_instance)

	__gong__computeReferencePass1(stage, stage.PageLayouts, &stage.PageLayouts_reference, &stage.PageLayouts_referenceOrder, &stage.PageLayouts_instance)

	__gong__computeReferencePass1(stage, stage.PageMarginss, &stage.PageMarginss_reference, &stage.PageMarginss_referenceOrder, &stage.PageMarginss_instance)

	__gong__computeReferencePass1(stage, stage.Plays, &stage.Plays_reference, &stage.Plays_referenceOrder, &stage.Plays_instance)

	__gong__computeReferencePass1(stage, stage.Scalings, &stage.Scalings_reference, &stage.Scalings_referenceOrder, &stage.Scalings_instance)

	__gong__computeReferencePass1(stage, stage.Segnos, &stage.Segnos_reference, &stage.Segnos_referenceOrder, &stage.Segnos_instance)

	__gong__computeReferencePass1(stage, stage.StaffLayouts, &stage.StaffLayouts_reference, &stage.StaffLayouts_referenceOrder, &stage.StaffLayouts_instance)

	__gong__computeReferencePass1(stage, stage.StringTypes, &stage.StringTypes_reference, &stage.StringTypes_referenceOrder, &stage.StringTypes_instance)

	__gong__computeReferencePass1(stage, stage.Supportss, &stage.Supportss_reference, &stage.Supportss_referenceOrder, &stage.Supportss_instance)

	__gong__computeReferencePass1(stage, stage.SystemDividerss, &stage.SystemDividerss_reference, &stage.SystemDividerss_referenceOrder, &stage.SystemDividerss_instance)

	__gong__computeReferencePass1(stage, stage.SystemLayouts, &stage.SystemLayouts_reference, &stage.SystemLayouts_referenceOrder, &stage.SystemLayouts_instance)

	__gong__computeReferencePass1(stage, stage.SystemMarginss, &stage.SystemMarginss_reference, &stage.SystemMarginss_referenceOrder, &stage.SystemMarginss_instance)

	__gong__computeReferencePass1(stage, stage.TypedTexts, &stage.TypedTexts_reference, &stage.TypedTexts_referenceOrder, &stage.TypedTexts_instance)

	__gong__computeReferencePass1(stage, stage.WavyLines, &stage.WavyLines_reference, &stage.WavyLines_referenceOrder, &stage.WavyLines_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.AccidentalTexts, stage.AccidentalTexts_reference, stage)

	__gong__computeReferencePass2(stage.Appearances, stage.Appearances_reference, stage)

	__gong__computeReferencePass2(stage.Bookmarks, stage.Bookmarks_reference, stage)

	__gong__computeReferencePass2(stage.Codas, stage.Codas_reference, stage)

	__gong__computeReferencePass2(stage.Distances, stage.Distances_reference, stage)

	__gong__computeReferencePass2(stage.Dynamicss, stage.Dynamicss_reference, stage)

	__gong__computeReferencePass2(stage.Emptys, stage.Emptys_reference, stage)

	__gong__computeReferencePass2(stage.EmptyPlacements, stage.EmptyPlacements_reference, stage)

	__gong__computeReferencePass2(stage.EmptyPlacementSmufls, stage.EmptyPlacementSmufls_reference, stage)

	__gong__computeReferencePass2(stage.EmptyPrintObjectStyleAligns, stage.EmptyPrintObjectStyleAligns_reference, stage)

	__gong__computeReferencePass2(stage.EmptyPrintStyles, stage.EmptyPrintStyles_reference, stage)

	__gong__computeReferencePass2(stage.EmptyPrintStyleAligns, stage.EmptyPrintStyleAligns_reference, stage)

	__gong__computeReferencePass2(stage.EmptyPrintStyleAlignIds, stage.EmptyPrintStyleAlignIds_reference, stage)

	__gong__computeReferencePass2(stage.EmptyTrillSounds, stage.EmptyTrillSounds_reference, stage)

	__gong__computeReferencePass2(stage.Encodings, stage.Encodings_reference, stage)

	__gong__computeReferencePass2(stage.Fermatas, stage.Fermatas_reference, stage)

	__gong__computeReferencePass2(stage.Fingerings, stage.Fingerings_reference, stage)

	__gong__computeReferencePass2(stage.FormattedSymbols, stage.FormattedSymbols_reference, stage)

	__gong__computeReferencePass2(stage.FormattedSymbolIds, stage.FormattedSymbolIds_reference, stage)

	__gong__computeReferencePass2(stage.FormattedTexts, stage.FormattedTexts_reference, stage)

	__gong__computeReferencePass2(stage.FormattedTextIds, stage.FormattedTextIds_reference, stage)

	__gong__computeReferencePass2(stage.Frets, stage.Frets_reference, stage)

	__gong__computeReferencePass2(stage.Glyphs, stage.Glyphs_reference, stage)

	__gong__computeReferencePass2(stage.HorizontalTurns, stage.HorizontalTurns_reference, stage)

	__gong__computeReferencePass2(stage.Identifications, stage.Identifications_reference, stage)

	__gong__computeReferencePass2(stage.Levels, stage.Levels_reference, stage)

	__gong__computeReferencePass2(stage.LineWidths, stage.LineWidths_reference, stage)

	__gong__computeReferencePass2(stage.Links, stage.Links_reference, stage)

	__gong__computeReferencePass2(stage.MeasureLayouts, stage.MeasureLayouts_reference, stage)

	__gong__computeReferencePass2(stage.MidiDevices, stage.MidiDevices_reference, stage)

	__gong__computeReferencePass2(stage.MidiInstruments, stage.MidiInstruments_reference, stage)

	__gong__computeReferencePass2(stage.Miscellaneouss, stage.Miscellaneouss_reference, stage)

	__gong__computeReferencePass2(stage.MiscellaneousFields, stage.MiscellaneousFields_reference, stage)

	__gong__computeReferencePass2(stage.NameDisplays, stage.NameDisplays_reference, stage)

	__gong__computeReferencePass2(stage.NoteSizes, stage.NoteSizes_reference, stage)

	__gong__computeReferencePass2(stage.OtherAppearances, stage.OtherAppearances_reference, stage)

	__gong__computeReferencePass2(stage.OtherPlays, stage.OtherPlays_reference, stage)

	__gong__computeReferencePass2(stage.OtherTexts, stage.OtherTexts_reference, stage)

	__gong__computeReferencePass2(stage.PageLayouts, stage.PageLayouts_reference, stage)

	__gong__computeReferencePass2(stage.PageMarginss, stage.PageMarginss_reference, stage)

	__gong__computeReferencePass2(stage.Plays, stage.Plays_reference, stage)

	__gong__computeReferencePass2(stage.Scalings, stage.Scalings_reference, stage)

	__gong__computeReferencePass2(stage.Segnos, stage.Segnos_reference, stage)

	__gong__computeReferencePass2(stage.StaffLayouts, stage.StaffLayouts_reference, stage)

	__gong__computeReferencePass2(stage.StringTypes, stage.StringTypes_reference, stage)

	__gong__computeReferencePass2(stage.Supportss, stage.Supportss_reference, stage)

	__gong__computeReferencePass2(stage.SystemDividerss, stage.SystemDividerss_reference, stage)

	__gong__computeReferencePass2(stage.SystemLayouts, stage.SystemLayouts_reference, stage)

	__gong__computeReferencePass2(stage.SystemMarginss, stage.SystemMarginss_reference, stage)

	__gong__computeReferencePass2(stage.TypedTexts, stage.TypedTexts_reference, stage)

	__gong__computeReferencePass2(stage.WavyLines, stage.WavyLines_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (accidentaltext *AccidentalText) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AccidentalText_stagedOrder, stage.AccidentalTexts_referenceOrder, accidentaltext, "AccidentalText")
}

func (appearance *Appearance) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Appearance_stagedOrder, stage.Appearances_referenceOrder, appearance, "Appearance")
}

func (bookmark *Bookmark) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Bookmark_stagedOrder, stage.Bookmarks_referenceOrder, bookmark, "Bookmark")
}

func (coda *Coda) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Coda_stagedOrder, stage.Codas_referenceOrder, coda, "Coda")
}

func (distance *Distance) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Distance_stagedOrder, stage.Distances_referenceOrder, distance, "Distance")
}

func (dynamics *Dynamics) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Dynamics_stagedOrder, stage.Dynamicss_referenceOrder, dynamics, "Dynamics")
}

func (empty *Empty) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Empty_stagedOrder, stage.Emptys_referenceOrder, empty, "Empty")
}

func (emptyplacement *EmptyPlacement) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EmptyPlacement_stagedOrder, stage.EmptyPlacements_referenceOrder, emptyplacement, "EmptyPlacement")
}

func (emptyplacementsmufl *EmptyPlacementSmufl) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EmptyPlacementSmufl_stagedOrder, stage.EmptyPlacementSmufls_referenceOrder, emptyplacementsmufl, "EmptyPlacementSmufl")
}

func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EmptyPrintObjectStyleAlign_stagedOrder, stage.EmptyPrintObjectStyleAligns_referenceOrder, emptyprintobjectstylealign, "EmptyPrintObjectStyleAlign")
}

func (emptyprintstyle *EmptyPrintStyle) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EmptyPrintStyle_stagedOrder, stage.EmptyPrintStyles_referenceOrder, emptyprintstyle, "EmptyPrintStyle")
}

func (emptyprintstylealign *EmptyPrintStyleAlign) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EmptyPrintStyleAlign_stagedOrder, stage.EmptyPrintStyleAligns_referenceOrder, emptyprintstylealign, "EmptyPrintStyleAlign")
}

func (emptyprintstylealignid *EmptyPrintStyleAlignId) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EmptyPrintStyleAlignId_stagedOrder, stage.EmptyPrintStyleAlignIds_referenceOrder, emptyprintstylealignid, "EmptyPrintStyleAlignId")
}

func (emptytrillsound *EmptyTrillSound) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EmptyTrillSound_stagedOrder, stage.EmptyTrillSounds_referenceOrder, emptytrillsound, "EmptyTrillSound")
}

func (encoding *Encoding) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Encoding_stagedOrder, stage.Encodings_referenceOrder, encoding, "Encoding")
}

func (fermata *Fermata) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Fermata_stagedOrder, stage.Fermatas_referenceOrder, fermata, "Fermata")
}

func (fingering *Fingering) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Fingering_stagedOrder, stage.Fingerings_referenceOrder, fingering, "Fingering")
}

func (formattedsymbol *FormattedSymbol) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FormattedSymbol_stagedOrder, stage.FormattedSymbols_referenceOrder, formattedsymbol, "FormattedSymbol")
}

func (formattedsymbolid *FormattedSymbolId) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FormattedSymbolId_stagedOrder, stage.FormattedSymbolIds_referenceOrder, formattedsymbolid, "FormattedSymbolId")
}

func (formattedtext *FormattedText) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FormattedText_stagedOrder, stage.FormattedTexts_referenceOrder, formattedtext, "FormattedText")
}

func (formattedtextid *FormattedTextId) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FormattedTextId_stagedOrder, stage.FormattedTextIds_referenceOrder, formattedtextid, "FormattedTextId")
}

func (fret *Fret) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Fret_stagedOrder, stage.Frets_referenceOrder, fret, "Fret")
}

func (glyph *Glyph) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Glyph_stagedOrder, stage.Glyphs_referenceOrder, glyph, "Glyph")
}

func (horizontalturn *HorizontalTurn) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HorizontalTurn_stagedOrder, stage.HorizontalTurns_referenceOrder, horizontalturn, "HorizontalTurn")
}

func (identification *Identification) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Identification_stagedOrder, stage.Identifications_referenceOrder, identification, "Identification")
}

func (level *Level) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Level_stagedOrder, stage.Levels_referenceOrder, level, "Level")
}

func (linewidth *LineWidth) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LineWidth_stagedOrder, stage.LineWidths_referenceOrder, linewidth, "LineWidth")
}

func (link *Link) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Link_stagedOrder, stage.Links_referenceOrder, link, "Link")
}

func (measurelayout *MeasureLayout) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MeasureLayout_stagedOrder, stage.MeasureLayouts_referenceOrder, measurelayout, "MeasureLayout")
}

func (mididevice *MidiDevice) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MidiDevice_stagedOrder, stage.MidiDevices_referenceOrder, mididevice, "MidiDevice")
}

func (midiinstrument *MidiInstrument) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MidiInstrument_stagedOrder, stage.MidiInstruments_referenceOrder, midiinstrument, "MidiInstrument")
}

func (miscellaneous *Miscellaneous) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Miscellaneous_stagedOrder, stage.Miscellaneouss_referenceOrder, miscellaneous, "Miscellaneous")
}

func (miscellaneousfield *MiscellaneousField) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MiscellaneousField_stagedOrder, stage.MiscellaneousFields_referenceOrder, miscellaneousfield, "MiscellaneousField")
}

func (namedisplay *NameDisplay) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NameDisplay_stagedOrder, stage.NameDisplays_referenceOrder, namedisplay, "NameDisplay")
}

func (notesize *NoteSize) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteSize_stagedOrder, stage.NoteSizes_referenceOrder, notesize, "NoteSize")
}

func (otherappearance *OtherAppearance) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OtherAppearance_stagedOrder, stage.OtherAppearances_referenceOrder, otherappearance, "OtherAppearance")
}

func (otherplay *OtherPlay) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OtherPlay_stagedOrder, stage.OtherPlays_referenceOrder, otherplay, "OtherPlay")
}

func (othertext *OtherText) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OtherText_stagedOrder, stage.OtherTexts_referenceOrder, othertext, "OtherText")
}

func (pagelayout *PageLayout) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PageLayout_stagedOrder, stage.PageLayouts_referenceOrder, pagelayout, "PageLayout")
}

func (pagemargins *PageMargins) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PageMargins_stagedOrder, stage.PageMarginss_referenceOrder, pagemargins, "PageMargins")
}

func (play *Play) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Play_stagedOrder, stage.Plays_referenceOrder, play, "Play")
}

func (scaling *Scaling) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Scaling_stagedOrder, stage.Scalings_referenceOrder, scaling, "Scaling")
}

func (segno *Segno) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Segno_stagedOrder, stage.Segnos_referenceOrder, segno, "Segno")
}

func (stafflayout *StaffLayout) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StaffLayout_stagedOrder, stage.StaffLayouts_referenceOrder, stafflayout, "StaffLayout")
}

func (stringtype *StringType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StringType_stagedOrder, stage.StringTypes_referenceOrder, stringtype, "StringType")
}

func (supports *Supports) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Supports_stagedOrder, stage.Supportss_referenceOrder, supports, "Supports")
}

func (systemdividers *SystemDividers) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SystemDividers_stagedOrder, stage.SystemDividerss_referenceOrder, systemdividers, "SystemDividers")
}

func (systemlayout *SystemLayout) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SystemLayout_stagedOrder, stage.SystemLayouts_referenceOrder, systemlayout, "SystemLayout")
}

func (systemmargins *SystemMargins) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SystemMargins_stagedOrder, stage.SystemMarginss_referenceOrder, systemmargins, "SystemMargins")
}

func (typedtext *TypedText) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TypedText_stagedOrder, stage.TypedTexts_referenceOrder, typedtext, "TypedText")
}

func (wavyline *WavyLine) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.WavyLine_stagedOrder, stage.WavyLines_referenceOrder, wavyline, "WavyLine")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (accidentaltext *AccidentalText) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(accidentaltext, accidentaltext.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (accidentaltext *AccidentalText) GongGetReferenceIdentifier(stage *Stage) string {
	return accidentaltext.GongGetIdentifier(stage)
}

func (appearance *Appearance) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(appearance, appearance.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (appearance *Appearance) GongGetReferenceIdentifier(stage *Stage) string {
	return appearance.GongGetIdentifier(stage)
}

func (bookmark *Bookmark) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(bookmark, bookmark.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bookmark *Bookmark) GongGetReferenceIdentifier(stage *Stage) string {
	return bookmark.GongGetIdentifier(stage)
}

func (coda *Coda) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coda, coda.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coda *Coda) GongGetReferenceIdentifier(stage *Stage) string {
	return coda.GongGetIdentifier(stage)
}

func (distance *Distance) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(distance, distance.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (distance *Distance) GongGetReferenceIdentifier(stage *Stage) string {
	return distance.GongGetIdentifier(stage)
}

func (dynamics *Dynamics) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(dynamics, dynamics.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dynamics *Dynamics) GongGetReferenceIdentifier(stage *Stage) string {
	return dynamics.GongGetIdentifier(stage)
}

func (empty *Empty) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(empty, empty.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty *Empty) GongGetReferenceIdentifier(stage *Stage) string {
	return empty.GongGetIdentifier(stage)
}

func (emptyplacement *EmptyPlacement) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(emptyplacement, emptyplacement.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (emptyplacement *EmptyPlacement) GongGetReferenceIdentifier(stage *Stage) string {
	return emptyplacement.GongGetIdentifier(stage)
}

func (emptyplacementsmufl *EmptyPlacementSmufl) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(emptyplacementsmufl, emptyplacementsmufl.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (emptyplacementsmufl *EmptyPlacementSmufl) GongGetReferenceIdentifier(stage *Stage) string {
	return emptyplacementsmufl.GongGetIdentifier(stage)
}

func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(emptyprintobjectstylealign, emptyprintobjectstylealign.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) GongGetReferenceIdentifier(stage *Stage) string {
	return emptyprintobjectstylealign.GongGetIdentifier(stage)
}

func (emptyprintstyle *EmptyPrintStyle) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(emptyprintstyle, emptyprintstyle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (emptyprintstyle *EmptyPrintStyle) GongGetReferenceIdentifier(stage *Stage) string {
	return emptyprintstyle.GongGetIdentifier(stage)
}

func (emptyprintstylealign *EmptyPrintStyleAlign) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(emptyprintstylealign, emptyprintstylealign.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (emptyprintstylealign *EmptyPrintStyleAlign) GongGetReferenceIdentifier(stage *Stage) string {
	return emptyprintstylealign.GongGetIdentifier(stage)
}

func (emptyprintstylealignid *EmptyPrintStyleAlignId) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(emptyprintstylealignid, emptyprintstylealignid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (emptyprintstylealignid *EmptyPrintStyleAlignId) GongGetReferenceIdentifier(stage *Stage) string {
	return emptyprintstylealignid.GongGetIdentifier(stage)
}

func (emptytrillsound *EmptyTrillSound) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(emptytrillsound, emptytrillsound.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (emptytrillsound *EmptyTrillSound) GongGetReferenceIdentifier(stage *Stage) string {
	return emptytrillsound.GongGetIdentifier(stage)
}

func (encoding *Encoding) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(encoding, encoding.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (encoding *Encoding) GongGetReferenceIdentifier(stage *Stage) string {
	return encoding.GongGetIdentifier(stage)
}

func (fermata *Fermata) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(fermata, fermata.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (fermata *Fermata) GongGetReferenceIdentifier(stage *Stage) string {
	return fermata.GongGetIdentifier(stage)
}

func (fingering *Fingering) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(fingering, fingering.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (fingering *Fingering) GongGetReferenceIdentifier(stage *Stage) string {
	return fingering.GongGetIdentifier(stage)
}

func (formattedsymbol *FormattedSymbol) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formattedsymbol, formattedsymbol.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formattedsymbol *FormattedSymbol) GongGetReferenceIdentifier(stage *Stage) string {
	return formattedsymbol.GongGetIdentifier(stage)
}

func (formattedsymbolid *FormattedSymbolId) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formattedsymbolid, formattedsymbolid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formattedsymbolid *FormattedSymbolId) GongGetReferenceIdentifier(stage *Stage) string {
	return formattedsymbolid.GongGetIdentifier(stage)
}

func (formattedtext *FormattedText) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formattedtext, formattedtext.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formattedtext *FormattedText) GongGetReferenceIdentifier(stage *Stage) string {
	return formattedtext.GongGetIdentifier(stage)
}

func (formattedtextid *FormattedTextId) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formattedtextid, formattedtextid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formattedtextid *FormattedTextId) GongGetReferenceIdentifier(stage *Stage) string {
	return formattedtextid.GongGetIdentifier(stage)
}

func (fret *Fret) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(fret, fret.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (fret *Fret) GongGetReferenceIdentifier(stage *Stage) string {
	return fret.GongGetIdentifier(stage)
}

func (glyph *Glyph) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(glyph, glyph.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (glyph *Glyph) GongGetReferenceIdentifier(stage *Stage) string {
	return glyph.GongGetIdentifier(stage)
}

func (horizontalturn *HorizontalTurn) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(horizontalturn, horizontalturn.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (horizontalturn *HorizontalTurn) GongGetReferenceIdentifier(stage *Stage) string {
	return horizontalturn.GongGetIdentifier(stage)
}

func (identification *Identification) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(identification, identification.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (identification *Identification) GongGetReferenceIdentifier(stage *Stage) string {
	return identification.GongGetIdentifier(stage)
}

func (level *Level) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(level, level.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (level *Level) GongGetReferenceIdentifier(stage *Stage) string {
	return level.GongGetIdentifier(stage)
}

func (linewidth *LineWidth) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(linewidth, linewidth.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (linewidth *LineWidth) GongGetReferenceIdentifier(stage *Stage) string {
	return linewidth.GongGetIdentifier(stage)
}

func (link *Link) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(link, link.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (link *Link) GongGetReferenceIdentifier(stage *Stage) string {
	return link.GongGetIdentifier(stage)
}

func (measurelayout *MeasureLayout) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(measurelayout, measurelayout.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (measurelayout *MeasureLayout) GongGetReferenceIdentifier(stage *Stage) string {
	return measurelayout.GongGetIdentifier(stage)
}

func (mididevice *MidiDevice) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(mididevice, mididevice.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (mididevice *MidiDevice) GongGetReferenceIdentifier(stage *Stage) string {
	return mididevice.GongGetIdentifier(stage)
}

func (midiinstrument *MidiInstrument) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(midiinstrument, midiinstrument.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (midiinstrument *MidiInstrument) GongGetReferenceIdentifier(stage *Stage) string {
	return midiinstrument.GongGetIdentifier(stage)
}

func (miscellaneous *Miscellaneous) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(miscellaneous, miscellaneous.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (miscellaneous *Miscellaneous) GongGetReferenceIdentifier(stage *Stage) string {
	return miscellaneous.GongGetIdentifier(stage)
}

func (miscellaneousfield *MiscellaneousField) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(miscellaneousfield, miscellaneousfield.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (miscellaneousfield *MiscellaneousField) GongGetReferenceIdentifier(stage *Stage) string {
	return miscellaneousfield.GongGetIdentifier(stage)
}

func (namedisplay *NameDisplay) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(namedisplay, namedisplay.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (namedisplay *NameDisplay) GongGetReferenceIdentifier(stage *Stage) string {
	return namedisplay.GongGetIdentifier(stage)
}

func (notesize *NoteSize) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(notesize, notesize.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notesize *NoteSize) GongGetReferenceIdentifier(stage *Stage) string {
	return notesize.GongGetIdentifier(stage)
}

func (otherappearance *OtherAppearance) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(otherappearance, otherappearance.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (otherappearance *OtherAppearance) GongGetReferenceIdentifier(stage *Stage) string {
	return otherappearance.GongGetIdentifier(stage)
}

func (otherplay *OtherPlay) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(otherplay, otherplay.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (otherplay *OtherPlay) GongGetReferenceIdentifier(stage *Stage) string {
	return otherplay.GongGetIdentifier(stage)
}

func (othertext *OtherText) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(othertext, othertext.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (othertext *OtherText) GongGetReferenceIdentifier(stage *Stage) string {
	return othertext.GongGetIdentifier(stage)
}

func (pagelayout *PageLayout) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pagelayout, pagelayout.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pagelayout *PageLayout) GongGetReferenceIdentifier(stage *Stage) string {
	return pagelayout.GongGetIdentifier(stage)
}

func (pagemargins *PageMargins) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pagemargins, pagemargins.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pagemargins *PageMargins) GongGetReferenceIdentifier(stage *Stage) string {
	return pagemargins.GongGetIdentifier(stage)
}

func (play *Play) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(play, play.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (play *Play) GongGetReferenceIdentifier(stage *Stage) string {
	return play.GongGetIdentifier(stage)
}

func (scaling *Scaling) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(scaling, scaling.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (scaling *Scaling) GongGetReferenceIdentifier(stage *Stage) string {
	return scaling.GongGetIdentifier(stage)
}

func (segno *Segno) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(segno, segno.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (segno *Segno) GongGetReferenceIdentifier(stage *Stage) string {
	return segno.GongGetIdentifier(stage)
}

func (stafflayout *StaffLayout) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stafflayout, stafflayout.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stafflayout *StaffLayout) GongGetReferenceIdentifier(stage *Stage) string {
	return stafflayout.GongGetIdentifier(stage)
}

func (stringtype *StringType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stringtype, stringtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stringtype *StringType) GongGetReferenceIdentifier(stage *Stage) string {
	return stringtype.GongGetIdentifier(stage)
}

func (supports *Supports) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(supports, supports.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (supports *Supports) GongGetReferenceIdentifier(stage *Stage) string {
	return supports.GongGetIdentifier(stage)
}

func (systemdividers *SystemDividers) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(systemdividers, systemdividers.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (systemdividers *SystemDividers) GongGetReferenceIdentifier(stage *Stage) string {
	return systemdividers.GongGetIdentifier(stage)
}

func (systemlayout *SystemLayout) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(systemlayout, systemlayout.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (systemlayout *SystemLayout) GongGetReferenceIdentifier(stage *Stage) string {
	return systemlayout.GongGetIdentifier(stage)
}

func (systemmargins *SystemMargins) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(systemmargins, systemmargins.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (systemmargins *SystemMargins) GongGetReferenceIdentifier(stage *Stage) string {
	return systemmargins.GongGetIdentifier(stage)
}

func (typedtext *TypedText) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(typedtext, typedtext.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (typedtext *TypedText) GongGetReferenceIdentifier(stage *Stage) string {
	return typedtext.GongGetIdentifier(stage)
}

func (wavyline *WavyLine) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(wavyline, wavyline.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (wavyline *WavyLine) GongGetReferenceIdentifier(stage *Stage) string {
	return wavyline.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (accidentaltext *AccidentalText) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(accidentaltext.GongGetIdentifier(stage), "AccidentalText", accidentaltext.Name)
}

func (appearance *Appearance) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(appearance.GongGetIdentifier(stage), "Appearance", appearance.Name)
}

func (bookmark *Bookmark) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(bookmark.GongGetIdentifier(stage), "Bookmark", bookmark.Name)
}

func (coda *Coda) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coda.GongGetIdentifier(stage), "Coda", coda.Name)
}

func (distance *Distance) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(distance.GongGetIdentifier(stage), "Distance", distance.Name)
}

func (dynamics *Dynamics) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(dynamics.GongGetIdentifier(stage), "Dynamics", dynamics.Name)
}

func (empty *Empty) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(empty.GongGetIdentifier(stage), "Empty", empty.Name)
}

func (emptyplacement *EmptyPlacement) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(emptyplacement.GongGetIdentifier(stage), "EmptyPlacement", emptyplacement.Name)
}

func (emptyplacementsmufl *EmptyPlacementSmufl) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(emptyplacementsmufl.GongGetIdentifier(stage), "EmptyPlacementSmufl", emptyplacementsmufl.Name)
}

func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(emptyprintobjectstylealign.GongGetIdentifier(stage), "EmptyPrintObjectStyleAlign", emptyprintobjectstylealign.Name)
}

func (emptyprintstyle *EmptyPrintStyle) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(emptyprintstyle.GongGetIdentifier(stage), "EmptyPrintStyle", emptyprintstyle.Name)
}

func (emptyprintstylealign *EmptyPrintStyleAlign) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(emptyprintstylealign.GongGetIdentifier(stage), "EmptyPrintStyleAlign", emptyprintstylealign.Name)
}

func (emptyprintstylealignid *EmptyPrintStyleAlignId) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(emptyprintstylealignid.GongGetIdentifier(stage), "EmptyPrintStyleAlignId", emptyprintstylealignid.Name)
}

func (emptytrillsound *EmptyTrillSound) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(emptytrillsound.GongGetIdentifier(stage), "EmptyTrillSound", emptytrillsound.Name)
}

func (encoding *Encoding) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(encoding.GongGetIdentifier(stage), "Encoding", encoding.Name)
}

func (fermata *Fermata) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(fermata.GongGetIdentifier(stage), "Fermata", fermata.Name)
}

func (fingering *Fingering) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(fingering.GongGetIdentifier(stage), "Fingering", fingering.Name)
}

func (formattedsymbol *FormattedSymbol) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formattedsymbol.GongGetIdentifier(stage), "FormattedSymbol", formattedsymbol.Name)
}

func (formattedsymbolid *FormattedSymbolId) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formattedsymbolid.GongGetIdentifier(stage), "FormattedSymbolId", formattedsymbolid.Name)
}

func (formattedtext *FormattedText) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formattedtext.GongGetIdentifier(stage), "FormattedText", formattedtext.Name)
}

func (formattedtextid *FormattedTextId) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formattedtextid.GongGetIdentifier(stage), "FormattedTextId", formattedtextid.Name)
}

func (fret *Fret) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(fret.GongGetIdentifier(stage), "Fret", fret.Name)
}

func (glyph *Glyph) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(glyph.GongGetIdentifier(stage), "Glyph", glyph.Name)
}

func (horizontalturn *HorizontalTurn) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(horizontalturn.GongGetIdentifier(stage), "HorizontalTurn", horizontalturn.Name)
}

func (identification *Identification) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(identification.GongGetIdentifier(stage), "Identification", identification.Name)
}

func (level *Level) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(level.GongGetIdentifier(stage), "Level", level.Name)
}

func (linewidth *LineWidth) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(linewidth.GongGetIdentifier(stage), "LineWidth", linewidth.Name)
}

func (link *Link) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(link.GongGetIdentifier(stage), "Link", link.Name)
}

func (measurelayout *MeasureLayout) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(measurelayout.GongGetIdentifier(stage), "MeasureLayout", measurelayout.Name)
}

func (mididevice *MidiDevice) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(mididevice.GongGetIdentifier(stage), "MidiDevice", mididevice.Name)
}

func (midiinstrument *MidiInstrument) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(midiinstrument.GongGetIdentifier(stage), "MidiInstrument", midiinstrument.Name)
}

func (miscellaneous *Miscellaneous) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(miscellaneous.GongGetIdentifier(stage), "Miscellaneous", miscellaneous.Name)
}

func (miscellaneousfield *MiscellaneousField) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(miscellaneousfield.GongGetIdentifier(stage), "MiscellaneousField", miscellaneousfield.Name)
}

func (namedisplay *NameDisplay) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(namedisplay.GongGetIdentifier(stage), "NameDisplay", namedisplay.Name)
}

func (notesize *NoteSize) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notesize.GongGetIdentifier(stage), "NoteSize", notesize.Name)
}

func (otherappearance *OtherAppearance) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(otherappearance.GongGetIdentifier(stage), "OtherAppearance", otherappearance.Name)
}

func (otherplay *OtherPlay) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(otherplay.GongGetIdentifier(stage), "OtherPlay", otherplay.Name)
}

func (othertext *OtherText) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(othertext.GongGetIdentifier(stage), "OtherText", othertext.Name)
}

func (pagelayout *PageLayout) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pagelayout.GongGetIdentifier(stage), "PageLayout", pagelayout.Name)
}

func (pagemargins *PageMargins) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pagemargins.GongGetIdentifier(stage), "PageMargins", pagemargins.Name)
}

func (play *Play) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(play.GongGetIdentifier(stage), "Play", play.Name)
}

func (scaling *Scaling) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(scaling.GongGetIdentifier(stage), "Scaling", scaling.Name)
}

func (segno *Segno) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(segno.GongGetIdentifier(stage), "Segno", segno.Name)
}

func (stafflayout *StaffLayout) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stafflayout.GongGetIdentifier(stage), "StaffLayout", stafflayout.Name)
}

func (stringtype *StringType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stringtype.GongGetIdentifier(stage), "StringType", stringtype.Name)
}

func (supports *Supports) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(supports.GongGetIdentifier(stage), "Supports", supports.Name)
}

func (systemdividers *SystemDividers) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(systemdividers.GongGetIdentifier(stage), "SystemDividers", systemdividers.Name)
}

func (systemlayout *SystemLayout) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(systemlayout.GongGetIdentifier(stage), "SystemLayout", systemlayout.Name)
}

func (systemmargins *SystemMargins) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(systemmargins.GongGetIdentifier(stage), "SystemMargins", systemmargins.Name)
}

func (typedtext *TypedText) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(typedtext.GongGetIdentifier(stage), "TypedText", typedtext.Name)
}

func (wavyline *WavyLine) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(wavyline.GongGetIdentifier(stage), "WavyLine", wavyline.Name)
}

// insertion point for unstaging
func (accidentaltext *AccidentalText) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(accidentaltext.GongGetReferenceIdentifier(stage))
}

func (appearance *Appearance) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(appearance.GongGetReferenceIdentifier(stage))
}

func (bookmark *Bookmark) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(bookmark.GongGetReferenceIdentifier(stage))
}

func (coda *Coda) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coda.GongGetReferenceIdentifier(stage))
}

func (distance *Distance) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(distance.GongGetReferenceIdentifier(stage))
}

func (dynamics *Dynamics) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(dynamics.GongGetReferenceIdentifier(stage))
}

func (empty *Empty) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(empty.GongGetReferenceIdentifier(stage))
}

func (emptyplacement *EmptyPlacement) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(emptyplacement.GongGetReferenceIdentifier(stage))
}

func (emptyplacementsmufl *EmptyPlacementSmufl) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(emptyplacementsmufl.GongGetReferenceIdentifier(stage))
}

func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(emptyprintobjectstylealign.GongGetReferenceIdentifier(stage))
}

func (emptyprintstyle *EmptyPrintStyle) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(emptyprintstyle.GongGetReferenceIdentifier(stage))
}

func (emptyprintstylealign *EmptyPrintStyleAlign) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(emptyprintstylealign.GongGetReferenceIdentifier(stage))
}

func (emptyprintstylealignid *EmptyPrintStyleAlignId) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(emptyprintstylealignid.GongGetReferenceIdentifier(stage))
}

func (emptytrillsound *EmptyTrillSound) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(emptytrillsound.GongGetReferenceIdentifier(stage))
}

func (encoding *Encoding) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(encoding.GongGetReferenceIdentifier(stage))
}

func (fermata *Fermata) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(fermata.GongGetReferenceIdentifier(stage))
}

func (fingering *Fingering) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(fingering.GongGetReferenceIdentifier(stage))
}

func (formattedsymbol *FormattedSymbol) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formattedsymbol.GongGetReferenceIdentifier(stage))
}

func (formattedsymbolid *FormattedSymbolId) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formattedsymbolid.GongGetReferenceIdentifier(stage))
}

func (formattedtext *FormattedText) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formattedtext.GongGetReferenceIdentifier(stage))
}

func (formattedtextid *FormattedTextId) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formattedtextid.GongGetReferenceIdentifier(stage))
}

func (fret *Fret) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(fret.GongGetReferenceIdentifier(stage))
}

func (glyph *Glyph) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(glyph.GongGetReferenceIdentifier(stage))
}

func (horizontalturn *HorizontalTurn) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(horizontalturn.GongGetReferenceIdentifier(stage))
}

func (identification *Identification) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(identification.GongGetReferenceIdentifier(stage))
}

func (level *Level) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(level.GongGetReferenceIdentifier(stage))
}

func (linewidth *LineWidth) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(linewidth.GongGetReferenceIdentifier(stage))
}

func (link *Link) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(link.GongGetReferenceIdentifier(stage))
}

func (measurelayout *MeasureLayout) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(measurelayout.GongGetReferenceIdentifier(stage))
}

func (mididevice *MidiDevice) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(mididevice.GongGetReferenceIdentifier(stage))
}

func (midiinstrument *MidiInstrument) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(midiinstrument.GongGetReferenceIdentifier(stage))
}

func (miscellaneous *Miscellaneous) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(miscellaneous.GongGetReferenceIdentifier(stage))
}

func (miscellaneousfield *MiscellaneousField) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(miscellaneousfield.GongGetReferenceIdentifier(stage))
}

func (namedisplay *NameDisplay) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(namedisplay.GongGetReferenceIdentifier(stage))
}

func (notesize *NoteSize) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notesize.GongGetReferenceIdentifier(stage))
}

func (otherappearance *OtherAppearance) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(otherappearance.GongGetReferenceIdentifier(stage))
}

func (otherplay *OtherPlay) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(otherplay.GongGetReferenceIdentifier(stage))
}

func (othertext *OtherText) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(othertext.GongGetReferenceIdentifier(stage))
}

func (pagelayout *PageLayout) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pagelayout.GongGetReferenceIdentifier(stage))
}

func (pagemargins *PageMargins) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pagemargins.GongGetReferenceIdentifier(stage))
}

func (play *Play) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(play.GongGetReferenceIdentifier(stage))
}

func (scaling *Scaling) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(scaling.GongGetReferenceIdentifier(stage))
}

func (segno *Segno) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(segno.GongGetReferenceIdentifier(stage))
}

func (stafflayout *StaffLayout) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stafflayout.GongGetReferenceIdentifier(stage))
}

func (stringtype *StringType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stringtype.GongGetReferenceIdentifier(stage))
}

func (supports *Supports) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(supports.GongGetReferenceIdentifier(stage))
}

func (systemdividers *SystemDividers) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(systemdividers.GongGetReferenceIdentifier(stage))
}

func (systemlayout *SystemLayout) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(systemlayout.GongGetReferenceIdentifier(stage))
}

func (systemmargins *SystemMargins) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(systemmargins.GongGetReferenceIdentifier(stage))
}

func (typedtext *TypedText) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(typedtext.GongGetReferenceIdentifier(stage))
}

func (wavyline *WavyLine) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(wavyline.GongGetReferenceIdentifier(stage))
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
