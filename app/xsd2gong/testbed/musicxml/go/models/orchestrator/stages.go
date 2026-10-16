// generated code by xsd2gong - do not edit
package orchestrator

import (
	"fmt"

	models "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/xlink"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/common"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/barline"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/attributes"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/direction"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/note"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/score"
)

// MusicXMLStages coordinates all sub-package Gong stages.
type MusicXMLStages struct {
	StageSet *models.StageSet
	XlinkStage *xlink.Stage
	CommonStage *common.Stage
	BarlineStage *barline.Stage
	AttributesStage *attributes.Stage
	DirectionStage *direction.Stage
	NoteStage *note.Stage
	ScoreStage *score.Stage
}

// NewMusicXMLStages initializes all sub-package stages and binds them to a root StageSet.
func NewMusicXMLStages() *MusicXMLStages {
	ss := models.NewStageSet("musicxml")
	s := &MusicXMLStages{
		StageSet: ss,
		XlinkStage: ss.XlinkStage,
		CommonStage: ss.CommonStage,
		BarlineStage: ss.BarlineStage,
		AttributesStage: ss.AttributesStage,
		DirectionStage: ss.DirectionStage,
		NoteStage: ss.NoteStage,
		ScoreStage: ss.ScoreStage,
	}
	return s
}

// Commit commits all stages in topological dependency order (leaves first, root last).
func (s *MusicXMLStages) Commit() {
	if s.StageSet != nil {
		s.StageSet.Commit()
		return
	}
	if s.XlinkStage != nil {
		s.XlinkStage.Commit()
	}
	if s.CommonStage != nil {
		s.CommonStage.Commit()
	}
	if s.BarlineStage != nil {
		s.BarlineStage.Commit()
	}
	if s.AttributesStage != nil {
		s.AttributesStage.Commit()
	}
	if s.DirectionStage != nil {
		s.DirectionStage.Commit()
	}
	if s.NoteStage != nil {
		s.NoteStage.Commit()
	}
	if s.ScoreStage != nil {
		s.ScoreStage.Commit()
	}
}

// StageSubtree walks the object graph starting from root, assigns deterministic Gong Names,
// and stages each instance into its respective package Stage.
func (s *MusicXMLStages) StageSubtree(root *score.ScorePartwise) {
	counts := make(map[string]int)
	s.stageScore_ScorePartwise(root, counts)
}

func (s *MusicXMLStages) stageXlink_Simple(inst *xlink.Simple, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["xlink.Simple"]++
		inst.Name = fmt.Sprintf("Simple_%d", counts["xlink.Simple"])
	}
	inst.Stage(s.XlinkStage)
}

func (s *MusicXMLStages) stageXlink_Extended(inst *xlink.Extended, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["xlink.Extended"]++
		inst.Name = fmt.Sprintf("Extended_%d", counts["xlink.Extended"])
	}
	inst.Stage(s.XlinkStage)
	for _, child := range inst.Title_1 {
		s.stageXlink_TitleEltType(child, counts)
	}
	for _, child := range inst.Resource {
		s.stageXlink_ResourceType(child, counts)
	}
	for _, child := range inst.Locator {
		s.stageXlink_LocatorType(child, counts)
	}
	for _, child := range inst.Arc {
		s.stageXlink_ArcType(child, counts)
	}
}

func (s *MusicXMLStages) stageXlink_TitleEltType(inst *xlink.TitleEltType, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["xlink.TitleEltType"]++
		inst.Name = fmt.Sprintf("TitleEltType_%d", counts["xlink.TitleEltType"])
	}
	inst.Stage(s.XlinkStage)
}

func (s *MusicXMLStages) stageXlink_ResourceType(inst *xlink.ResourceType, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["xlink.ResourceType"]++
		inst.Name = fmt.Sprintf("ResourceType_%d", counts["xlink.ResourceType"])
	}
	inst.Stage(s.XlinkStage)
}

func (s *MusicXMLStages) stageXlink_LocatorType(inst *xlink.LocatorType, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["xlink.LocatorType"]++
		inst.Name = fmt.Sprintf("LocatorType_%d", counts["xlink.LocatorType"])
	}
	inst.Stage(s.XlinkStage)
	for _, child := range inst.Title_1 {
		s.stageXlink_TitleEltType(child, counts)
	}
}

func (s *MusicXMLStages) stageXlink_ArcType(inst *xlink.ArcType, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["xlink.ArcType"]++
		inst.Name = fmt.Sprintf("ArcType_%d", counts["xlink.ArcType"])
	}
	inst.Stage(s.XlinkStage)
	for _, child := range inst.Title_1 {
		s.stageXlink_TitleEltType(child, counts)
	}
}

func (s *MusicXMLStages) stageCommon_AccidentalText(inst *common.AccidentalText, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.AccidentalText"]++
		inst.Name = fmt.Sprintf("AccidentalText_%d", counts["common.AccidentalText"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_Coda(inst *common.Coda, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.Coda"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Coda_%s_%d", inst.Id, counts["common.Coda"])
		} else {
			inst.Name = fmt.Sprintf("Coda_%d", counts["common.Coda"])
		}
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_Dynamics(inst *common.Dynamics, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.Dynamics"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Dynamics_%s_%d", inst.Id, counts["common.Dynamics"])
		} else {
			inst.Name = fmt.Sprintf("Dynamics_%d", counts["common.Dynamics"])
		}
	}
	inst.Stage(s.CommonStage)
	for _, child := range inst.P {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Pp {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Ppp {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Pppp {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Ppppp {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Pppppp {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.F {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Ff {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Fff {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Ffff {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Fffff {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Ffffff {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Mp {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Mf {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Sf {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Sfp {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Sfpp {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Fp {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Rf {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Rfz {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Sfz {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Sffz {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Fz {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.N {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Pf {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.Sfzp {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.OtherDynamics {
		s.stageCommon_OtherText(child, counts)
	}
}

func (s *MusicXMLStages) stageCommon_Empty(inst *common.Empty, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.Empty"]++
		inst.Name = fmt.Sprintf("Empty_%d", counts["common.Empty"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_EmptyPlacement(inst *common.EmptyPlacement, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.EmptyPlacement"]++
		inst.Name = fmt.Sprintf("EmptyPlacement_%d", counts["common.EmptyPlacement"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_EmptyPlacementSmufl(inst *common.EmptyPlacementSmufl, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.EmptyPlacementSmufl"]++
		inst.Name = fmt.Sprintf("EmptyPlacementSmufl_%d", counts["common.EmptyPlacementSmufl"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_EmptyPrintStyle(inst *common.EmptyPrintStyle, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.EmptyPrintStyle"]++
		inst.Name = fmt.Sprintf("EmptyPrintStyle_%d", counts["common.EmptyPrintStyle"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_EmptyPrintStyleAlign(inst *common.EmptyPrintStyleAlign, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.EmptyPrintStyleAlign"]++
		inst.Name = fmt.Sprintf("EmptyPrintStyleAlign_%d", counts["common.EmptyPrintStyleAlign"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_EmptyPrintStyleAlignId(inst *common.EmptyPrintStyleAlignId, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.EmptyPrintStyleAlignId"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("EmptyPrintStyleAlignId_%s_%d", inst.Id, counts["common.EmptyPrintStyleAlignId"])
		} else {
			inst.Name = fmt.Sprintf("EmptyPrintStyleAlignId_%d", counts["common.EmptyPrintStyleAlignId"])
		}
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_EmptyPrintObjectStyleAlign(inst *common.EmptyPrintObjectStyleAlign, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.EmptyPrintObjectStyleAlign"]++
		inst.Name = fmt.Sprintf("EmptyPrintObjectStyleAlign_%d", counts["common.EmptyPrintObjectStyleAlign"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_EmptyTrillSound(inst *common.EmptyTrillSound, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.EmptyTrillSound"]++
		inst.Name = fmt.Sprintf("EmptyTrillSound_%d", counts["common.EmptyTrillSound"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_HorizontalTurn(inst *common.HorizontalTurn, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.HorizontalTurn"]++
		inst.Name = fmt.Sprintf("HorizontalTurn_%d", counts["common.HorizontalTurn"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_Fermata(inst *common.Fermata, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.Fermata"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Fermata_%s_%d", inst.Id, counts["common.Fermata"])
		} else {
			inst.Name = fmt.Sprintf("Fermata_%d", counts["common.Fermata"])
		}
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_Fingering(inst *common.Fingering, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.Fingering"]++
		inst.Name = fmt.Sprintf("Fingering_%d", counts["common.Fingering"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_FormattedSymbol(inst *common.FormattedSymbol, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.FormattedSymbol"]++
		inst.Name = fmt.Sprintf("FormattedSymbol_%d", counts["common.FormattedSymbol"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_FormattedSymbolId(inst *common.FormattedSymbolId, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.FormattedSymbolId"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("FormattedSymbolId_%s_%d", inst.Id, counts["common.FormattedSymbolId"])
		} else {
			inst.Name = fmt.Sprintf("FormattedSymbolId_%d", counts["common.FormattedSymbolId"])
		}
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_FormattedText(inst *common.FormattedText, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.FormattedText"]++
		inst.Name = fmt.Sprintf("FormattedText_%d", counts["common.FormattedText"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_FormattedTextId(inst *common.FormattedTextId, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.FormattedTextId"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("FormattedTextId_%s_%d", inst.Id, counts["common.FormattedTextId"])
		} else {
			inst.Name = fmt.Sprintf("FormattedTextId_%d", counts["common.FormattedTextId"])
		}
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_Fret(inst *common.Fret, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.Fret"]++
		inst.Name = fmt.Sprintf("Fret_%d", counts["common.Fret"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_Level(inst *common.Level, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.Level"]++
		inst.Name = fmt.Sprintf("Level_%d", counts["common.Level"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_MidiDevice(inst *common.MidiDevice, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.MidiDevice"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("MidiDevice_%s_%d", inst.Id, counts["common.MidiDevice"])
		} else {
			inst.Name = fmt.Sprintf("MidiDevice_%d", counts["common.MidiDevice"])
		}
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_MidiInstrument(inst *common.MidiInstrument, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.MidiInstrument"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("MidiInstrument_%s_%d", inst.Id, counts["common.MidiInstrument"])
		} else {
			inst.Name = fmt.Sprintf("MidiInstrument_%d", counts["common.MidiInstrument"])
		}
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_NameDisplay(inst *common.NameDisplay, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.NameDisplay"]++
		inst.Name = fmt.Sprintf("NameDisplay_%d", counts["common.NameDisplay"])
	}
	inst.Stage(s.CommonStage)
	for _, child := range inst.DisplayText {
		s.stageCommon_FormattedText(child, counts)
	}
	for _, child := range inst.AccidentalText {
		s.stageCommon_AccidentalText(child, counts)
	}
}

func (s *MusicXMLStages) stageCommon_OtherPlay(inst *common.OtherPlay, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.OtherPlay"]++
		inst.Name = fmt.Sprintf("OtherPlay_%d", counts["common.OtherPlay"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_Play(inst *common.Play, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.Play"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Play_%s_%d", inst.Id, counts["common.Play"])
		} else {
			inst.Name = fmt.Sprintf("Play_%d", counts["common.Play"])
		}
	}
	inst.Stage(s.CommonStage)
	for _, child := range inst.OtherPlay {
		s.stageCommon_OtherPlay(child, counts)
	}
}

func (s *MusicXMLStages) stageCommon_Segno(inst *common.Segno, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.Segno"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Segno_%s_%d", inst.Id, counts["common.Segno"])
		} else {
			inst.Name = fmt.Sprintf("Segno_%d", counts["common.Segno"])
		}
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_StringType(inst *common.StringType, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.StringType"]++
		inst.Name = fmt.Sprintf("StringType_%d", counts["common.StringType"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_TypedText(inst *common.TypedText, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.TypedText"]++
		inst.Name = fmt.Sprintf("TypedText_%d", counts["common.TypedText"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_WavyLine(inst *common.WavyLine, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.WavyLine"]++
		inst.Name = fmt.Sprintf("WavyLine_%d", counts["common.WavyLine"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_Encoding(inst *common.Encoding, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.Encoding"]++
		inst.Name = fmt.Sprintf("Encoding_%d", counts["common.Encoding"])
	}
	inst.Stage(s.CommonStage)
	for _, child := range inst.Encoder {
		s.stageCommon_TypedText(child, counts)
	}
	for _, child := range inst.Supports {
		s.stageCommon_Supports(child, counts)
	}
}

func (s *MusicXMLStages) stageCommon_Identification(inst *common.Identification, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.Identification"]++
		inst.Name = fmt.Sprintf("Identification_%d", counts["common.Identification"])
	}
	inst.Stage(s.CommonStage)
	for _, child := range inst.Creator {
		s.stageCommon_TypedText(child, counts)
	}
	for _, child := range inst.Rights {
		s.stageCommon_TypedText(child, counts)
	}
	s.stageCommon_Encoding(inst.Encoding, counts)
	for _, child := range inst.Relation {
		s.stageCommon_TypedText(child, counts)
	}
	s.stageCommon_Miscellaneous(inst.Miscellaneous, counts)
}

func (s *MusicXMLStages) stageCommon_Miscellaneous(inst *common.Miscellaneous, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.Miscellaneous"]++
		inst.Name = fmt.Sprintf("Miscellaneous_%d", counts["common.Miscellaneous"])
	}
	inst.Stage(s.CommonStage)
	for _, child := range inst.MiscellaneousField {
		s.stageCommon_MiscellaneousField(child, counts)
	}
}

func (s *MusicXMLStages) stageCommon_MiscellaneousField(inst *common.MiscellaneousField, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.MiscellaneousField"]++
		inst.Name = fmt.Sprintf("MiscellaneousField_%d", counts["common.MiscellaneousField"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_Supports(inst *common.Supports, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.Supports"]++
		inst.Name = fmt.Sprintf("Supports_%d", counts["common.Supports"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_Appearance(inst *common.Appearance, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.Appearance"]++
		inst.Name = fmt.Sprintf("Appearance_%d", counts["common.Appearance"])
	}
	inst.Stage(s.CommonStage)
	for _, child := range inst.LineWidth {
		s.stageCommon_LineWidth(child, counts)
	}
	for _, child := range inst.NoteSize {
		s.stageCommon_NoteSize(child, counts)
	}
	for _, child := range inst.Distance {
		s.stageCommon_Distance(child, counts)
	}
	for _, child := range inst.Glyph {
		s.stageCommon_Glyph(child, counts)
	}
	for _, child := range inst.OtherAppearance {
		s.stageCommon_OtherAppearance(child, counts)
	}
}

func (s *MusicXMLStages) stageCommon_Distance(inst *common.Distance, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.Distance"]++
		inst.Name = fmt.Sprintf("Distance_%d", counts["common.Distance"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_Glyph(inst *common.Glyph, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.Glyph"]++
		inst.Name = fmt.Sprintf("Glyph_%d", counts["common.Glyph"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_LineWidth(inst *common.LineWidth, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.LineWidth"]++
		inst.Name = fmt.Sprintf("LineWidth_%d", counts["common.LineWidth"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_MeasureLayout(inst *common.MeasureLayout, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.MeasureLayout"]++
		inst.Name = fmt.Sprintf("MeasureLayout_%d", counts["common.MeasureLayout"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_NoteSize(inst *common.NoteSize, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.NoteSize"]++
		inst.Name = fmt.Sprintf("NoteSize_%d", counts["common.NoteSize"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_OtherAppearance(inst *common.OtherAppearance, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.OtherAppearance"]++
		inst.Name = fmt.Sprintf("OtherAppearance_%d", counts["common.OtherAppearance"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_PageLayout(inst *common.PageLayout, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.PageLayout"]++
		inst.Name = fmt.Sprintf("PageLayout_%d", counts["common.PageLayout"])
	}
	inst.Stage(s.CommonStage)
	for _, child := range inst.PageMargins {
		s.stageCommon_PageMargins(child, counts)
	}
}

func (s *MusicXMLStages) stageCommon_PageMargins(inst *common.PageMargins, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.PageMargins"]++
		inst.Name = fmt.Sprintf("PageMargins_%d", counts["common.PageMargins"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_Scaling(inst *common.Scaling, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.Scaling"]++
		inst.Name = fmt.Sprintf("Scaling_%d", counts["common.Scaling"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_StaffLayout(inst *common.StaffLayout, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.StaffLayout"]++
		inst.Name = fmt.Sprintf("StaffLayout_%d", counts["common.StaffLayout"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_SystemDividers(inst *common.SystemDividers, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.SystemDividers"]++
		inst.Name = fmt.Sprintf("SystemDividers_%d", counts["common.SystemDividers"])
	}
	inst.Stage(s.CommonStage)
	s.stageCommon_EmptyPrintObjectStyleAlign(inst.LeftDivider, counts)
	s.stageCommon_EmptyPrintObjectStyleAlign(inst.RightDivider, counts)
}

func (s *MusicXMLStages) stageCommon_SystemLayout(inst *common.SystemLayout, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.SystemLayout"]++
		inst.Name = fmt.Sprintf("SystemLayout_%d", counts["common.SystemLayout"])
	}
	inst.Stage(s.CommonStage)
	s.stageCommon_SystemMargins(inst.SystemMargins, counts)
	s.stageCommon_SystemDividers(inst.SystemDividers, counts)
}

func (s *MusicXMLStages) stageCommon_SystemMargins(inst *common.SystemMargins, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.SystemMargins"]++
		inst.Name = fmt.Sprintf("SystemMargins_%d", counts["common.SystemMargins"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_Bookmark(inst *common.Bookmark, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.Bookmark"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Bookmark_%s_%d", inst.Id, counts["common.Bookmark"])
		} else {
			inst.Name = fmt.Sprintf("Bookmark_%d", counts["common.Bookmark"])
		}
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_Link(inst *common.Link, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.Link"]++
		inst.Name = fmt.Sprintf("Link_%d", counts["common.Link"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageCommon_OtherText(inst *common.OtherText, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["common.OtherText"]++
		inst.Name = fmt.Sprintf("OtherText_%d", counts["common.OtherText"])
	}
	inst.Stage(s.CommonStage)
}

func (s *MusicXMLStages) stageBarline_BarStyleColor(inst *barline.BarStyleColor, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["barline.BarStyleColor"]++
		inst.Name = fmt.Sprintf("BarStyleColor_%d", counts["barline.BarStyleColor"])
	}
	inst.Stage(s.BarlineStage)
}

func (s *MusicXMLStages) stageBarline_Barline(inst *barline.Barline, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["barline.Barline"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Barline_%s_%d", inst.Id, counts["barline.Barline"])
		} else {
			inst.Name = fmt.Sprintf("Barline_%d", counts["barline.Barline"])
		}
	}
	inst.Stage(s.BarlineStage)
	s.stageBarline_BarStyleColor(inst.BarStyle, counts)
	s.stageCommon_FormattedText(inst.Footnote, counts)
	s.stageCommon_Level(inst.Level, counts)
	s.stageCommon_WavyLine(inst.WavyLine, counts)
	s.stageCommon_Segno(inst.Segno_1, counts)
	s.stageCommon_Coda(inst.Coda_1, counts)
	for _, child := range inst.Fermata {
		s.stageCommon_Fermata(child, counts)
	}
	s.stageBarline_Ending(inst.Ending, counts)
	s.stageBarline_Repeat(inst.Repeat, counts)
}

func (s *MusicXMLStages) stageBarline_Ending(inst *barline.Ending, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["barline.Ending"]++
		if inst.Number != "" {
			inst.Name = fmt.Sprintf("Ending_%s_%d", inst.Number, counts["barline.Ending"])
		} else {
			inst.Name = fmt.Sprintf("Ending_%d", counts["barline.Ending"])
		}
	}
	inst.Stage(s.BarlineStage)
}

func (s *MusicXMLStages) stageBarline_Repeat(inst *barline.Repeat, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["barline.Repeat"]++
		inst.Name = fmt.Sprintf("Repeat_%d", counts["barline.Repeat"])
	}
	inst.Stage(s.BarlineStage)
}

func (s *MusicXMLStages) stageAttributes_Attributes(inst *attributes.Attributes, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.Attributes"]++
		inst.Name = fmt.Sprintf("Attributes_%d", counts["attributes.Attributes"])
	}
	inst.Stage(s.AttributesStage)
	s.stageCommon_FormattedText(inst.Footnote, counts)
	s.stageCommon_Level(inst.Level, counts)
	for _, child := range inst.Key {
		s.stageAttributes_Key(child, counts)
	}
	for _, child := range inst.Time {
		s.stageAttributes_Time(child, counts)
	}
	s.stageAttributes_PartSymbol(inst.PartSymbol, counts)
	for _, child := range inst.Clef {
		s.stageAttributes_Clef(child, counts)
	}
	for _, child := range inst.StaffDetails {
		s.stageAttributes_StaffDetails(child, counts)
	}
	for _, child := range inst.Transpose {
		s.stageAttributes_Transpose(child, counts)
	}
	for _, child := range inst.ForPart {
		s.stageAttributes_ForPart(child, counts)
	}
	for _, child := range inst.Directive {
		s.stageAttributes_AttributesDirective(child, counts)
	}
	for _, child := range inst.MeasureStyle {
		s.stageAttributes_MeasureStyle(child, counts)
	}
}

func (s *MusicXMLStages) stageAttributes_BeatRepeat(inst *attributes.BeatRepeat, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.BeatRepeat"]++
		inst.Name = fmt.Sprintf("BeatRepeat_%d", counts["attributes.BeatRepeat"])
	}
	inst.Stage(s.AttributesStage)
	for _, child := range inst.SlashDot {
		s.stageCommon_Empty(child, counts)
	}
}

func (s *MusicXMLStages) stageAttributes_Cancel(inst *attributes.Cancel, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.Cancel"]++
		inst.Name = fmt.Sprintf("Cancel_%d", counts["attributes.Cancel"])
	}
	inst.Stage(s.AttributesStage)
}

func (s *MusicXMLStages) stageAttributes_Clef(inst *attributes.Clef, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.Clef"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Clef_%s_%d", inst.Id, counts["attributes.Clef"])
		} else {
			inst.Name = fmt.Sprintf("Clef_%d", counts["attributes.Clef"])
		}
	}
	inst.Stage(s.AttributesStage)
}

func (s *MusicXMLStages) stageAttributes_Double(inst *attributes.Double, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.Double"]++
		inst.Name = fmt.Sprintf("Double_%d", counts["attributes.Double"])
	}
	inst.Stage(s.AttributesStage)
}

func (s *MusicXMLStages) stageAttributes_ForPart(inst *attributes.ForPart, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.ForPart"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("ForPart_%s_%d", inst.Id, counts["attributes.ForPart"])
		} else {
			inst.Name = fmt.Sprintf("ForPart_%d", counts["attributes.ForPart"])
		}
	}
	inst.Stage(s.AttributesStage)
	s.stageAttributes_PartClef(inst.PartClef, counts)
	s.stageAttributes_PartTranspose(inst.PartTranspose, counts)
}

func (s *MusicXMLStages) stageAttributes_Interchangeable(inst *attributes.Interchangeable, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.Interchangeable"]++
		inst.Name = fmt.Sprintf("Interchangeable_%d", counts["attributes.Interchangeable"])
	}
	inst.Stage(s.AttributesStage)
}

func (s *MusicXMLStages) stageAttributes_Key(inst *attributes.Key, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.Key"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Key_%s_%d", inst.Id, counts["attributes.Key"])
		} else {
			inst.Name = fmt.Sprintf("Key_%d", counts["attributes.Key"])
		}
	}
	inst.Stage(s.AttributesStage)
	s.stageAttributes_Cancel(inst.Cancel, counts)
	for _, child := range inst.KeyAccidental {
		s.stageAttributes_KeyAccidental(child, counts)
	}
	for _, child := range inst.KeyOctave {
		s.stageAttributes_KeyOctave(child, counts)
	}
}

func (s *MusicXMLStages) stageAttributes_KeyAccidental(inst *attributes.KeyAccidental, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.KeyAccidental"]++
		inst.Name = fmt.Sprintf("KeyAccidental_%d", counts["attributes.KeyAccidental"])
	}
	inst.Stage(s.AttributesStage)
}

func (s *MusicXMLStages) stageAttributes_KeyOctave(inst *attributes.KeyOctave, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.KeyOctave"]++
		inst.Name = fmt.Sprintf("KeyOctave_%d", counts["attributes.KeyOctave"])
	}
	inst.Stage(s.AttributesStage)
}

func (s *MusicXMLStages) stageAttributes_LineDetail(inst *attributes.LineDetail, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.LineDetail"]++
		inst.Name = fmt.Sprintf("LineDetail_%d", counts["attributes.LineDetail"])
	}
	inst.Stage(s.AttributesStage)
}

func (s *MusicXMLStages) stageAttributes_MeasureRepeat(inst *attributes.MeasureRepeat, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.MeasureRepeat"]++
		inst.Name = fmt.Sprintf("MeasureRepeat_%d", counts["attributes.MeasureRepeat"])
	}
	inst.Stage(s.AttributesStage)
}

func (s *MusicXMLStages) stageAttributes_MeasureStyle(inst *attributes.MeasureStyle, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.MeasureStyle"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("MeasureStyle_%s_%d", inst.Id, counts["attributes.MeasureStyle"])
		} else {
			inst.Name = fmt.Sprintf("MeasureStyle_%d", counts["attributes.MeasureStyle"])
		}
	}
	inst.Stage(s.AttributesStage)
	s.stageAttributes_MultipleRest(inst.MultipleRest, counts)
	s.stageAttributes_MeasureRepeat(inst.MeasureRepeat, counts)
	s.stageAttributes_BeatRepeat(inst.BeatRepeat, counts)
	s.stageAttributes_Slash(inst.Slash, counts)
}

func (s *MusicXMLStages) stageAttributes_MultipleRest(inst *attributes.MultipleRest, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.MultipleRest"]++
		inst.Name = fmt.Sprintf("MultipleRest_%d", counts["attributes.MultipleRest"])
	}
	inst.Stage(s.AttributesStage)
}

func (s *MusicXMLStages) stageAttributes_PartClef(inst *attributes.PartClef, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.PartClef"]++
		inst.Name = fmt.Sprintf("PartClef_%d", counts["attributes.PartClef"])
	}
	inst.Stage(s.AttributesStage)
}

func (s *MusicXMLStages) stageAttributes_PartSymbol(inst *attributes.PartSymbol, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.PartSymbol"]++
		inst.Name = fmt.Sprintf("PartSymbol_%d", counts["attributes.PartSymbol"])
	}
	inst.Stage(s.AttributesStage)
}

func (s *MusicXMLStages) stageAttributes_PartTranspose(inst *attributes.PartTranspose, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.PartTranspose"]++
		inst.Name = fmt.Sprintf("PartTranspose_%d", counts["attributes.PartTranspose"])
	}
	inst.Stage(s.AttributesStage)
	s.stageAttributes_Double(inst.Double, counts)
}

func (s *MusicXMLStages) stageAttributes_Slash(inst *attributes.Slash, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.Slash"]++
		inst.Name = fmt.Sprintf("Slash_%d", counts["attributes.Slash"])
	}
	inst.Stage(s.AttributesStage)
	for _, child := range inst.SlashDot {
		s.stageCommon_Empty(child, counts)
	}
}

func (s *MusicXMLStages) stageAttributes_StaffDetails(inst *attributes.StaffDetails, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.StaffDetails"]++
		inst.Name = fmt.Sprintf("StaffDetails_%d", counts["attributes.StaffDetails"])
	}
	inst.Stage(s.AttributesStage)
	for _, child := range inst.LineDetail {
		s.stageAttributes_LineDetail(child, counts)
	}
	for _, child := range inst.StaffTuning {
		s.stageAttributes_StaffTuning(child, counts)
	}
	s.stageAttributes_StaffSize(inst.StaffSize, counts)
}

func (s *MusicXMLStages) stageAttributes_StaffSize(inst *attributes.StaffSize, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.StaffSize"]++
		inst.Name = fmt.Sprintf("StaffSize_%d", counts["attributes.StaffSize"])
	}
	inst.Stage(s.AttributesStage)
}

func (s *MusicXMLStages) stageAttributes_StaffTuning(inst *attributes.StaffTuning, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.StaffTuning"]++
		inst.Name = fmt.Sprintf("StaffTuning_%d", counts["attributes.StaffTuning"])
	}
	inst.Stage(s.AttributesStage)
}

func (s *MusicXMLStages) stageAttributes_Time(inst *attributes.Time, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.Time"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Time_%s_%d", inst.Id, counts["attributes.Time"])
		} else {
			inst.Name = fmt.Sprintf("Time_%d", counts["attributes.Time"])
		}
	}
	inst.Stage(s.AttributesStage)
	s.stageAttributes_Interchangeable(inst.Interchangeable, counts)
}

func (s *MusicXMLStages) stageAttributes_Transpose(inst *attributes.Transpose, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.Transpose"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Transpose_%s_%d", inst.Id, counts["attributes.Transpose"])
		} else {
			inst.Name = fmt.Sprintf("Transpose_%d", counts["attributes.Transpose"])
		}
	}
	inst.Stage(s.AttributesStage)
	s.stageAttributes_Double(inst.Double, counts)
}

func (s *MusicXMLStages) stageAttributes_AttributesDirective(inst *attributes.AttributesDirective, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["attributes.AttributesDirective"]++
		inst.Name = fmt.Sprintf("AttributesDirective_%d", counts["attributes.AttributesDirective"])
	}
	inst.Stage(s.AttributesStage)
}

func (s *MusicXMLStages) stageDirection_Accord(inst *direction.Accord, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Accord"]++
		inst.Name = fmt.Sprintf("Accord_%d", counts["direction.Accord"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_AccordionRegistration(inst *direction.AccordionRegistration, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.AccordionRegistration"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("AccordionRegistration_%s_%d", inst.Id, counts["direction.AccordionRegistration"])
		} else {
			inst.Name = fmt.Sprintf("AccordionRegistration_%d", counts["direction.AccordionRegistration"])
		}
	}
	inst.Stage(s.DirectionStage)
	s.stageCommon_Empty(inst.AccordionHigh, counts)
	s.stageCommon_Empty(inst.AccordionLow, counts)
}

func (s *MusicXMLStages) stageDirection_Barre(inst *direction.Barre, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Barre"]++
		inst.Name = fmt.Sprintf("Barre_%d", counts["direction.Barre"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Bass(inst *direction.Bass, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Bass"]++
		inst.Name = fmt.Sprintf("Bass_%d", counts["direction.Bass"])
	}
	inst.Stage(s.DirectionStage)
	s.stageDirection_StyleText(inst.BassSeparator, counts)
	s.stageDirection_BassStep(inst.BassStep, counts)
	s.stageDirection_HarmonyAlter(inst.BassAlter, counts)
}

func (s *MusicXMLStages) stageDirection_HarmonyAlter(inst *direction.HarmonyAlter, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.HarmonyAlter"]++
		inst.Name = fmt.Sprintf("HarmonyAlter_%d", counts["direction.HarmonyAlter"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_BassStep(inst *direction.BassStep, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.BassStep"]++
		inst.Name = fmt.Sprintf("BassStep_%d", counts["direction.BassStep"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Beater(inst *direction.Beater, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Beater"]++
		inst.Name = fmt.Sprintf("Beater_%d", counts["direction.Beater"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_BeatUnitTied(inst *direction.BeatUnitTied, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.BeatUnitTied"]++
		inst.Name = fmt.Sprintf("BeatUnitTied_%d", counts["direction.BeatUnitTied"])
	}
	inst.Stage(s.DirectionStage)
	for _, child := range inst.BeatUnitDot {
		s.stageCommon_Empty(child, counts)
	}
}

func (s *MusicXMLStages) stageDirection_Bracket(inst *direction.Bracket, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Bracket"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Bracket_%s_%d", inst.Id, counts["direction.Bracket"])
		} else {
			inst.Name = fmt.Sprintf("Bracket_%d", counts["direction.Bracket"])
		}
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Dashes(inst *direction.Dashes, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Dashes"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Dashes_%s_%d", inst.Id, counts["direction.Dashes"])
		} else {
			inst.Name = fmt.Sprintf("Dashes_%d", counts["direction.Dashes"])
		}
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Degree(inst *direction.Degree, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Degree"]++
		inst.Name = fmt.Sprintf("Degree_%d", counts["direction.Degree"])
	}
	inst.Stage(s.DirectionStage)
	s.stageDirection_DegreeValue(inst.DegreeValue, counts)
	s.stageDirection_DegreeAlter(inst.DegreeAlter, counts)
	s.stageDirection_DegreeType(inst.DegreeType, counts)
}

func (s *MusicXMLStages) stageDirection_DegreeAlter(inst *direction.DegreeAlter, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.DegreeAlter"]++
		inst.Name = fmt.Sprintf("DegreeAlter_%d", counts["direction.DegreeAlter"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_DegreeType(inst *direction.DegreeType, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.DegreeType"]++
		inst.Name = fmt.Sprintf("DegreeType_%d", counts["direction.DegreeType"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_DegreeValue(inst *direction.DegreeValue, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.DegreeValue"]++
		inst.Name = fmt.Sprintf("DegreeValue_%d", counts["direction.DegreeValue"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Direction(inst *direction.Direction, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Direction"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Direction_%s_%d", inst.Id, counts["direction.Direction"])
		} else {
			inst.Name = fmt.Sprintf("Direction_%d", counts["direction.Direction"])
		}
	}
	inst.Stage(s.DirectionStage)
	for _, child := range inst.DirectionType {
		s.stageDirection_DirectionType(child, counts)
	}
	s.stageDirection_Offset(inst.Offset, counts)
	s.stageCommon_FormattedText(inst.Footnote, counts)
	s.stageCommon_Level(inst.Level, counts)
	s.stageDirection_Sound(inst.Sound, counts)
	s.stageDirection_Listening(inst.Listening, counts)
}

func (s *MusicXMLStages) stageDirection_DirectionType(inst *direction.DirectionType, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.DirectionType"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("DirectionType_%s_%d", inst.Id, counts["direction.DirectionType"])
		} else {
			inst.Name = fmt.Sprintf("DirectionType_%d", counts["direction.DirectionType"])
		}
	}
	inst.Stage(s.DirectionStage)
	for _, child := range inst.Rehearsal {
		s.stageCommon_FormattedTextId(child, counts)
	}
	for _, child := range inst.Segno {
		s.stageCommon_Segno(child, counts)
	}
	for _, child := range inst.Coda {
		s.stageCommon_Coda(child, counts)
	}
	for _, child := range inst.Words {
		s.stageCommon_FormattedTextId(child, counts)
	}
	for _, child := range inst.Symbol {
		s.stageCommon_FormattedSymbolId(child, counts)
	}
	s.stageDirection_Wedge(inst.Wedge, counts)
	for _, child := range inst.Dynamics {
		s.stageCommon_Dynamics(child, counts)
	}
	s.stageDirection_Dashes(inst.Dashes, counts)
	s.stageDirection_Bracket(inst.Bracket, counts)
	s.stageDirection_Pedal(inst.Pedal, counts)
	s.stageDirection_Metronome(inst.Metronome, counts)
	s.stageDirection_OctaveShift(inst.OctaveShift, counts)
	s.stageDirection_HarpPedals(inst.HarpPedals, counts)
	s.stageCommon_EmptyPrintStyleAlignId(inst.Damp, counts)
	s.stageCommon_EmptyPrintStyleAlignId(inst.DampAll, counts)
	s.stageCommon_EmptyPrintStyleAlignId(inst.Eyeglasses, counts)
	s.stageDirection_StringMute(inst.StringMute, counts)
	s.stageDirection_Scordatura(inst.Scordatura, counts)
	s.stageDirection_Image(inst.Image, counts)
	s.stageDirection_PrincipalVoice(inst.PrincipalVoice, counts)
	for _, child := range inst.Percussion {
		s.stageDirection_Percussion(child, counts)
	}
	s.stageDirection_AccordionRegistration(inst.AccordionRegistration, counts)
	s.stageDirection_StaffDivide(inst.StaffDivide, counts)
	s.stageDirection_OtherDirection(inst.OtherDirection, counts)
}

func (s *MusicXMLStages) stageDirection_Effect(inst *direction.Effect, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Effect"]++
		inst.Name = fmt.Sprintf("Effect_%d", counts["direction.Effect"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Feature(inst *direction.Feature, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Feature"]++
		inst.Name = fmt.Sprintf("Feature_%d", counts["direction.Feature"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_FirstFret(inst *direction.FirstFret, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.FirstFret"]++
		inst.Name = fmt.Sprintf("FirstFret_%d", counts["direction.FirstFret"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Frame(inst *direction.Frame, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Frame"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Frame_%s_%d", inst.Id, counts["direction.Frame"])
		} else {
			inst.Name = fmt.Sprintf("Frame_%d", counts["direction.Frame"])
		}
	}
	inst.Stage(s.DirectionStage)
	s.stageDirection_FirstFret(inst.FirstFret, counts)
	for _, child := range inst.FrameNote {
		s.stageDirection_FrameNote(child, counts)
	}
}

func (s *MusicXMLStages) stageDirection_FrameNote(inst *direction.FrameNote, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.FrameNote"]++
		inst.Name = fmt.Sprintf("FrameNote_%d", counts["direction.FrameNote"])
	}
	inst.Stage(s.DirectionStage)
	s.stageCommon_StringType(inst.String, counts)
	s.stageCommon_Fret(inst.Fret, counts)
	s.stageCommon_Fingering(inst.Fingering, counts)
	s.stageDirection_Barre(inst.Barre, counts)
}

func (s *MusicXMLStages) stageDirection_Glass(inst *direction.Glass, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Glass"]++
		inst.Name = fmt.Sprintf("Glass_%d", counts["direction.Glass"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Grouping(inst *direction.Grouping, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Grouping"]++
		if inst.Number != "" {
			inst.Name = fmt.Sprintf("Grouping_%s_%d", inst.Number, counts["direction.Grouping"])
		} else {
			inst.Name = fmt.Sprintf("Grouping_%d", counts["direction.Grouping"])
		}
	}
	inst.Stage(s.DirectionStage)
	for _, child := range inst.Feature {
		s.stageDirection_Feature(child, counts)
	}
}

func (s *MusicXMLStages) stageDirection_Harmony(inst *direction.Harmony, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Harmony"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Harmony_%s_%d", inst.Id, counts["direction.Harmony"])
		} else {
			inst.Name = fmt.Sprintf("Harmony_%d", counts["direction.Harmony"])
		}
	}
	inst.Stage(s.DirectionStage)
	for _, child := range inst.Root {
		s.stageDirection_Root(child, counts)
	}
	for _, child := range inst.Numeral {
		s.stageDirection_Numeral(child, counts)
	}
	for _, child := range inst.Function {
		s.stageDirection_StyleText(child, counts)
	}
	for _, child := range inst.Kind {
		s.stageDirection_Kind(child, counts)
	}
	for _, child := range inst.Inversion {
		s.stageDirection_Inversion(child, counts)
	}
	for _, child := range inst.Bass {
		s.stageDirection_Bass(child, counts)
	}
	for _, child := range inst.Degree {
		s.stageDirection_Degree(child, counts)
	}
	s.stageDirection_Frame(inst.Frame, counts)
	s.stageDirection_Offset(inst.Offset, counts)
	s.stageCommon_FormattedText(inst.Footnote, counts)
	s.stageCommon_Level(inst.Level, counts)
}

func (s *MusicXMLStages) stageDirection_HarpPedals(inst *direction.HarpPedals, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.HarpPedals"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("HarpPedals_%s_%d", inst.Id, counts["direction.HarpPedals"])
		} else {
			inst.Name = fmt.Sprintf("HarpPedals_%d", counts["direction.HarpPedals"])
		}
	}
	inst.Stage(s.DirectionStage)
	for _, child := range inst.PedalTuning {
		s.stageDirection_PedalTuning(child, counts)
	}
}

func (s *MusicXMLStages) stageDirection_Image(inst *direction.Image, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Image"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Image_%s_%d", inst.Id, counts["direction.Image"])
		} else {
			inst.Name = fmt.Sprintf("Image_%d", counts["direction.Image"])
		}
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_InstrumentChange(inst *direction.InstrumentChange, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.InstrumentChange"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("InstrumentChange_%s_%d", inst.Id, counts["direction.InstrumentChange"])
		} else {
			inst.Name = fmt.Sprintf("InstrumentChange_%d", counts["direction.InstrumentChange"])
		}
	}
	inst.Stage(s.DirectionStage)
	s.stageCommon_Empty(inst.Solo, counts)
	s.stageDirection_VirtualInstrument(inst.VirtualInstrument, counts)
}

func (s *MusicXMLStages) stageDirection_Inversion(inst *direction.Inversion, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Inversion"]++
		inst.Name = fmt.Sprintf("Inversion_%d", counts["direction.Inversion"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Kind(inst *direction.Kind, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Kind"]++
		inst.Name = fmt.Sprintf("Kind_%d", counts["direction.Kind"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Listening(inst *direction.Listening, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Listening"]++
		inst.Name = fmt.Sprintf("Listening_%d", counts["direction.Listening"])
	}
	inst.Stage(s.DirectionStage)
	for _, child := range inst.Sync {
		s.stageDirection_Sync(child, counts)
	}
	for _, child := range inst.OtherListening {
		s.stageDirection_OtherListening(child, counts)
	}
	s.stageDirection_Offset(inst.Offset, counts)
}

func (s *MusicXMLStages) stageDirection_MeasureNumbering(inst *direction.MeasureNumbering, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.MeasureNumbering"]++
		inst.Name = fmt.Sprintf("MeasureNumbering_%d", counts["direction.MeasureNumbering"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Membrane(inst *direction.Membrane, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Membrane"]++
		inst.Name = fmt.Sprintf("Membrane_%d", counts["direction.Membrane"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Metal(inst *direction.Metal, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Metal"]++
		inst.Name = fmt.Sprintf("Metal_%d", counts["direction.Metal"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Metronome(inst *direction.Metronome, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Metronome"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Metronome_%s_%d", inst.Id, counts["direction.Metronome"])
		} else {
			inst.Name = fmt.Sprintf("Metronome_%d", counts["direction.Metronome"])
		}
	}
	inst.Stage(s.DirectionStage)
	for _, child := range inst.BeatUnitDot {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.BeatUnitTied {
		s.stageDirection_BeatUnitTied(child, counts)
	}
	s.stageDirection_PerMinute(inst.PerMinute, counts)
	s.stageCommon_Empty(inst.MetronomeArrows, counts)
	for _, child := range inst.MetronomeNote {
		s.stageDirection_MetronomeNote(child, counts)
	}
}

func (s *MusicXMLStages) stageDirection_MetronomeBeam(inst *direction.MetronomeBeam, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.MetronomeBeam"]++
		inst.Name = fmt.Sprintf("MetronomeBeam_%d", counts["direction.MetronomeBeam"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_MetronomeNote(inst *direction.MetronomeNote, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.MetronomeNote"]++
		inst.Name = fmt.Sprintf("MetronomeNote_%d", counts["direction.MetronomeNote"])
	}
	inst.Stage(s.DirectionStage)
	for _, child := range inst.MetronomeDot {
		s.stageCommon_Empty(child, counts)
	}
	for _, child := range inst.MetronomeBeam {
		s.stageDirection_MetronomeBeam(child, counts)
	}
	s.stageDirection_MetronomeTied(inst.MetronomeTied, counts)
	s.stageDirection_MetronomeTuplet(inst.MetronomeTuplet, counts)
}

func (s *MusicXMLStages) stageDirection_MetronomeTied(inst *direction.MetronomeTied, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.MetronomeTied"]++
		inst.Name = fmt.Sprintf("MetronomeTied_%d", counts["direction.MetronomeTied"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_MetronomeTuplet(inst *direction.MetronomeTuplet, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.MetronomeTuplet"]++
		inst.Name = fmt.Sprintf("MetronomeTuplet_%d", counts["direction.MetronomeTuplet"])
	}
	inst.Stage(s.DirectionStage)
	for _, child := range inst.NormalDot {
		s.stageCommon_Empty(child, counts)
	}
}

func (s *MusicXMLStages) stageDirection_Numeral(inst *direction.Numeral, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Numeral"]++
		inst.Name = fmt.Sprintf("Numeral_%d", counts["direction.Numeral"])
	}
	inst.Stage(s.DirectionStage)
	s.stageDirection_NumeralRoot(inst.NumeralRoot, counts)
	s.stageDirection_HarmonyAlter(inst.NumeralAlter, counts)
	s.stageDirection_NumeralKey(inst.NumeralKey, counts)
}

func (s *MusicXMLStages) stageDirection_NumeralKey(inst *direction.NumeralKey, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.NumeralKey"]++
		inst.Name = fmt.Sprintf("NumeralKey_%d", counts["direction.NumeralKey"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_NumeralRoot(inst *direction.NumeralRoot, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.NumeralRoot"]++
		inst.Name = fmt.Sprintf("NumeralRoot_%d", counts["direction.NumeralRoot"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_OctaveShift(inst *direction.OctaveShift, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.OctaveShift"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("OctaveShift_%s_%d", inst.Id, counts["direction.OctaveShift"])
		} else {
			inst.Name = fmt.Sprintf("OctaveShift_%d", counts["direction.OctaveShift"])
		}
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Offset(inst *direction.Offset, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Offset"]++
		inst.Name = fmt.Sprintf("Offset_%d", counts["direction.Offset"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_OtherDirection(inst *direction.OtherDirection, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.OtherDirection"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("OtherDirection_%s_%d", inst.Id, counts["direction.OtherDirection"])
		} else {
			inst.Name = fmt.Sprintf("OtherDirection_%d", counts["direction.OtherDirection"])
		}
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_OtherListening(inst *direction.OtherListening, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.OtherListening"]++
		inst.Name = fmt.Sprintf("OtherListening_%d", counts["direction.OtherListening"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Pedal(inst *direction.Pedal, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Pedal"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Pedal_%s_%d", inst.Id, counts["direction.Pedal"])
		} else {
			inst.Name = fmt.Sprintf("Pedal_%d", counts["direction.Pedal"])
		}
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_PedalTuning(inst *direction.PedalTuning, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.PedalTuning"]++
		inst.Name = fmt.Sprintf("PedalTuning_%d", counts["direction.PedalTuning"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_PerMinute(inst *direction.PerMinute, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.PerMinute"]++
		inst.Name = fmt.Sprintf("PerMinute_%d", counts["direction.PerMinute"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Percussion(inst *direction.Percussion, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Percussion"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Percussion_%s_%d", inst.Id, counts["direction.Percussion"])
		} else {
			inst.Name = fmt.Sprintf("Percussion_%d", counts["direction.Percussion"])
		}
	}
	inst.Stage(s.DirectionStage)
	s.stageDirection_Glass(inst.Glass, counts)
	s.stageDirection_Metal(inst.Metal, counts)
	s.stageDirection_Wood(inst.Wood, counts)
	s.stageDirection_Pitched(inst.Pitched, counts)
	s.stageDirection_Membrane(inst.Membrane, counts)
	s.stageDirection_Effect(inst.Effect, counts)
	s.stageDirection_Timpani(inst.Timpani, counts)
	s.stageDirection_Beater(inst.Beater, counts)
	s.stageDirection_Stick(inst.Stick, counts)
	s.stageCommon_OtherText(inst.OtherPercussion, counts)
}

func (s *MusicXMLStages) stageDirection_Pitched(inst *direction.Pitched, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Pitched"]++
		inst.Name = fmt.Sprintf("Pitched_%d", counts["direction.Pitched"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_PrincipalVoice(inst *direction.PrincipalVoice, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.PrincipalVoice"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("PrincipalVoice_%s_%d", inst.Id, counts["direction.PrincipalVoice"])
		} else {
			inst.Name = fmt.Sprintf("PrincipalVoice_%d", counts["direction.PrincipalVoice"])
		}
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Print(inst *direction.Print, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Print"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Print_%s_%d", inst.Id, counts["direction.Print"])
		} else {
			inst.Name = fmt.Sprintf("Print_%d", counts["direction.Print"])
		}
	}
	inst.Stage(s.DirectionStage)
	s.stageCommon_PageLayout(inst.PageLayout, counts)
	s.stageCommon_SystemLayout(inst.SystemLayout, counts)
	for _, child := range inst.StaffLayout {
		s.stageCommon_StaffLayout(child, counts)
	}
	s.stageCommon_MeasureLayout(inst.MeasureLayout, counts)
	s.stageDirection_MeasureNumbering(inst.MeasureNumbering, counts)
	s.stageCommon_NameDisplay(inst.PartNameDisplay, counts)
	s.stageCommon_NameDisplay(inst.PartAbbreviationDisplay, counts)
}

func (s *MusicXMLStages) stageDirection_Root(inst *direction.Root, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Root"]++
		inst.Name = fmt.Sprintf("Root_%d", counts["direction.Root"])
	}
	inst.Stage(s.DirectionStage)
	s.stageDirection_RootStep(inst.RootStep, counts)
	s.stageDirection_HarmonyAlter(inst.RootAlter, counts)
}

func (s *MusicXMLStages) stageDirection_RootStep(inst *direction.RootStep, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.RootStep"]++
		inst.Name = fmt.Sprintf("RootStep_%d", counts["direction.RootStep"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Scordatura(inst *direction.Scordatura, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Scordatura"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Scordatura_%s_%d", inst.Id, counts["direction.Scordatura"])
		} else {
			inst.Name = fmt.Sprintf("Scordatura_%d", counts["direction.Scordatura"])
		}
	}
	inst.Stage(s.DirectionStage)
	for _, child := range inst.Accord {
		s.stageDirection_Accord(child, counts)
	}
}

func (s *MusicXMLStages) stageDirection_Sound(inst *direction.Sound, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Sound"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Sound_%s_%d", inst.Id, counts["direction.Sound"])
		} else {
			inst.Name = fmt.Sprintf("Sound_%d", counts["direction.Sound"])
		}
	}
	inst.Stage(s.DirectionStage)
	for _, child := range inst.InstrumentChange {
		s.stageDirection_InstrumentChange(child, counts)
	}
	for _, child := range inst.MidiDevice {
		s.stageCommon_MidiDevice(child, counts)
	}
	for _, child := range inst.MidiInstrument {
		s.stageCommon_MidiInstrument(child, counts)
	}
	for _, child := range inst.Play {
		s.stageCommon_Play(child, counts)
	}
	s.stageDirection_Swing(inst.Swing, counts)
	s.stageDirection_Offset(inst.Offset, counts)
}

func (s *MusicXMLStages) stageDirection_StaffDivide(inst *direction.StaffDivide, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.StaffDivide"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("StaffDivide_%s_%d", inst.Id, counts["direction.StaffDivide"])
		} else {
			inst.Name = fmt.Sprintf("StaffDivide_%d", counts["direction.StaffDivide"])
		}
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Stick(inst *direction.Stick, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Stick"]++
		inst.Name = fmt.Sprintf("Stick_%d", counts["direction.Stick"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_StringMute(inst *direction.StringMute, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.StringMute"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("StringMute_%s_%d", inst.Id, counts["direction.StringMute"])
		} else {
			inst.Name = fmt.Sprintf("StringMute_%d", counts["direction.StringMute"])
		}
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Swing(inst *direction.Swing, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Swing"]++
		inst.Name = fmt.Sprintf("Swing_%d", counts["direction.Swing"])
	}
	inst.Stage(s.DirectionStage)
	s.stageCommon_Empty(inst.Straight, counts)
}

func (s *MusicXMLStages) stageDirection_Sync(inst *direction.Sync, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Sync"]++
		inst.Name = fmt.Sprintf("Sync_%d", counts["direction.Sync"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Timpani(inst *direction.Timpani, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Timpani"]++
		inst.Name = fmt.Sprintf("Timpani_%d", counts["direction.Timpani"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Wedge(inst *direction.Wedge, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Wedge"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Wedge_%s_%d", inst.Id, counts["direction.Wedge"])
		} else {
			inst.Name = fmt.Sprintf("Wedge_%d", counts["direction.Wedge"])
		}
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_Wood(inst *direction.Wood, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.Wood"]++
		inst.Name = fmt.Sprintf("Wood_%d", counts["direction.Wood"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_StyleText(inst *direction.StyleText, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.StyleText"]++
		inst.Name = fmt.Sprintf("StyleText_%d", counts["direction.StyleText"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageDirection_TimeModification(inst *direction.TimeModification, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.TimeModification"]++
		inst.Name = fmt.Sprintf("TimeModification_%d", counts["direction.TimeModification"])
	}
	inst.Stage(s.DirectionStage)
	for _, child := range inst.NormalDot {
		s.stageCommon_Empty(child, counts)
	}
}

func (s *MusicXMLStages) stageDirection_VirtualInstrument(inst *direction.VirtualInstrument, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["direction.VirtualInstrument"]++
		inst.Name = fmt.Sprintf("VirtualInstrument_%d", counts["direction.VirtualInstrument"])
	}
	inst.Stage(s.DirectionStage)
}

func (s *MusicXMLStages) stageNote_Accidental(inst *note.Accidental, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Accidental"]++
		inst.Name = fmt.Sprintf("Accidental_%d", counts["note.Accidental"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_AccidentalMark(inst *note.AccidentalMark, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.AccidentalMark"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("AccidentalMark_%s_%d", inst.Id, counts["note.AccidentalMark"])
		} else {
			inst.Name = fmt.Sprintf("AccidentalMark_%d", counts["note.AccidentalMark"])
		}
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Arpeggiate(inst *note.Arpeggiate, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Arpeggiate"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Arpeggiate_%s_%d", inst.Id, counts["note.Arpeggiate"])
		} else {
			inst.Name = fmt.Sprintf("Arpeggiate_%d", counts["note.Arpeggiate"])
		}
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Articulations(inst *note.Articulations, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Articulations"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Articulations_%s_%d", inst.Id, counts["note.Articulations"])
		} else {
			inst.Name = fmt.Sprintf("Articulations_%d", counts["note.Articulations"])
		}
	}
	inst.Stage(s.NoteStage)
	for _, child := range inst.Accent {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.StrongAccent {
		s.stageNote_StrongAccent(child, counts)
	}
	for _, child := range inst.Staccato {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.Tenuto {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.DetachedLegato {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.Staccatissimo {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.Spiccato {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.Scoop {
		s.stageNote_EmptyLine(child, counts)
	}
	for _, child := range inst.Plop {
		s.stageNote_EmptyLine(child, counts)
	}
	for _, child := range inst.Doit {
		s.stageNote_EmptyLine(child, counts)
	}
	for _, child := range inst.Falloff {
		s.stageNote_EmptyLine(child, counts)
	}
	for _, child := range inst.BreathMark {
		s.stageNote_BreathMark(child, counts)
	}
	for _, child := range inst.Caesura {
		s.stageNote_Caesura(child, counts)
	}
	for _, child := range inst.Stress {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.Unstress {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.SoftAccent {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.OtherArticulation {
		s.stageNote_OtherPlacementText(child, counts)
	}
}

func (s *MusicXMLStages) stageNote_Arrow(inst *note.Arrow, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Arrow"]++
		inst.Name = fmt.Sprintf("Arrow_%d", counts["note.Arrow"])
	}
	inst.Stage(s.NoteStage)
	s.stageCommon_Empty(inst.Arrowhead, counts)
}

func (s *MusicXMLStages) stageNote_Assess(inst *note.Assess, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Assess"]++
		inst.Name = fmt.Sprintf("Assess_%d", counts["note.Assess"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Backup(inst *note.Backup, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Backup"]++
		inst.Name = fmt.Sprintf("Backup_%d", counts["note.Backup"])
	}
	inst.Stage(s.NoteStage)
	s.stageCommon_FormattedText(inst.Footnote, counts)
	s.stageCommon_Level(inst.Level, counts)
}

func (s *MusicXMLStages) stageNote_Beam(inst *note.Beam, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Beam"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Beam_%s_%d", inst.Id, counts["note.Beam"])
		} else {
			inst.Name = fmt.Sprintf("Beam_%d", counts["note.Beam"])
		}
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Bend(inst *note.Bend, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Bend"]++
		inst.Name = fmt.Sprintf("Bend_%d", counts["note.Bend"])
	}
	inst.Stage(s.NoteStage)
	s.stageCommon_Empty(inst.PreBend, counts)
	s.stageNote_Release(inst.Release, counts)
	s.stageNote_PlacementText(inst.WithBar, counts)
}

func (s *MusicXMLStages) stageNote_BreathMark(inst *note.BreathMark, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.BreathMark"]++
		inst.Name = fmt.Sprintf("BreathMark_%d", counts["note.BreathMark"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Caesura(inst *note.Caesura, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Caesura"]++
		inst.Name = fmt.Sprintf("Caesura_%d", counts["note.Caesura"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Elision(inst *note.Elision, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Elision"]++
		inst.Name = fmt.Sprintf("Elision_%d", counts["note.Elision"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_EmptyLine(inst *note.EmptyLine, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.EmptyLine"]++
		inst.Name = fmt.Sprintf("EmptyLine_%d", counts["note.EmptyLine"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Extend(inst *note.Extend, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Extend"]++
		inst.Name = fmt.Sprintf("Extend_%d", counts["note.Extend"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Figure(inst *note.Figure, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Figure"]++
		inst.Name = fmt.Sprintf("Figure_%d", counts["note.Figure"])
	}
	inst.Stage(s.NoteStage)
	s.stageDirection_StyleText(inst.Prefix, counts)
	s.stageDirection_StyleText(inst.FigureNumber, counts)
	s.stageDirection_StyleText(inst.Suffix, counts)
	s.stageNote_Extend(inst.Extend, counts)
	s.stageCommon_FormattedText(inst.Footnote, counts)
	s.stageCommon_Level(inst.Level, counts)
}

func (s *MusicXMLStages) stageNote_FiguredBass(inst *note.FiguredBass, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.FiguredBass"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("FiguredBass_%s_%d", inst.Id, counts["note.FiguredBass"])
		} else {
			inst.Name = fmt.Sprintf("FiguredBass_%d", counts["note.FiguredBass"])
		}
	}
	inst.Stage(s.NoteStage)
	for _, child := range inst.Figure {
		s.stageNote_Figure(child, counts)
	}
	s.stageCommon_FormattedText(inst.Footnote, counts)
	s.stageCommon_Level(inst.Level, counts)
}

func (s *MusicXMLStages) stageNote_Forward(inst *note.Forward, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Forward"]++
		inst.Name = fmt.Sprintf("Forward_%d", counts["note.Forward"])
	}
	inst.Stage(s.NoteStage)
	s.stageCommon_FormattedText(inst.Footnote, counts)
	s.stageCommon_Level(inst.Level, counts)
}

func (s *MusicXMLStages) stageNote_Glissando(inst *note.Glissando, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Glissando"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Glissando_%s_%d", inst.Id, counts["note.Glissando"])
		} else {
			inst.Name = fmt.Sprintf("Glissando_%d", counts["note.Glissando"])
		}
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Grace(inst *note.Grace, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Grace"]++
		inst.Name = fmt.Sprintf("Grace_%d", counts["note.Grace"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_HammerOnPullOff(inst *note.HammerOnPullOff, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.HammerOnPullOff"]++
		inst.Name = fmt.Sprintf("HammerOnPullOff_%d", counts["note.HammerOnPullOff"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Handbell(inst *note.Handbell, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Handbell"]++
		inst.Name = fmt.Sprintf("Handbell_%d", counts["note.Handbell"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_HarmonClosed(inst *note.HarmonClosed, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.HarmonClosed"]++
		inst.Name = fmt.Sprintf("HarmonClosed_%d", counts["note.HarmonClosed"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_HarmonMute(inst *note.HarmonMute, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.HarmonMute"]++
		inst.Name = fmt.Sprintf("HarmonMute_%d", counts["note.HarmonMute"])
	}
	inst.Stage(s.NoteStage)
	s.stageNote_HarmonClosed(inst.HarmonClosed, counts)
}

func (s *MusicXMLStages) stageNote_Harmonic(inst *note.Harmonic, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Harmonic"]++
		inst.Name = fmt.Sprintf("Harmonic_%d", counts["note.Harmonic"])
	}
	inst.Stage(s.NoteStage)
	s.stageCommon_Empty(inst.Natural, counts)
	s.stageCommon_Empty(inst.Artificial, counts)
	s.stageCommon_Empty(inst.BasePitch, counts)
	s.stageCommon_Empty(inst.TouchingPitch, counts)
	s.stageCommon_Empty(inst.SoundingPitch, counts)
}

func (s *MusicXMLStages) stageNote_HeelToe(inst *note.HeelToe, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.HeelToe"]++
		inst.Name = fmt.Sprintf("HeelToe_%d", counts["note.HeelToe"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Hole(inst *note.Hole, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Hole"]++
		inst.Name = fmt.Sprintf("Hole_%d", counts["note.Hole"])
	}
	inst.Stage(s.NoteStage)
	s.stageNote_HoleClosed(inst.HoleClosed, counts)
}

func (s *MusicXMLStages) stageNote_HoleClosed(inst *note.HoleClosed, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.HoleClosed"]++
		inst.Name = fmt.Sprintf("HoleClosed_%d", counts["note.HoleClosed"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Instrument(inst *note.Instrument, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Instrument"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Instrument_%s_%d", inst.Id, counts["note.Instrument"])
		} else {
			inst.Name = fmt.Sprintf("Instrument_%d", counts["note.Instrument"])
		}
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Listen(inst *note.Listen, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Listen"]++
		inst.Name = fmt.Sprintf("Listen_%d", counts["note.Listen"])
	}
	inst.Stage(s.NoteStage)
	for _, child := range inst.Assess {
		s.stageNote_Assess(child, counts)
	}
	for _, child := range inst.Wait {
		s.stageNote_Wait(child, counts)
	}
	for _, child := range inst.OtherListen {
		s.stageDirection_OtherListening(child, counts)
	}
}

func (s *MusicXMLStages) stageNote_Lyric(inst *note.Lyric, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Lyric"]++
		if inst.Number != "" {
			inst.Name = fmt.Sprintf("Lyric_%s_%d", inst.Number, counts["note.Lyric"])
		} else {
			inst.Name = fmt.Sprintf("Lyric_%d", counts["note.Lyric"])
		}
	}
	inst.Stage(s.NoteStage)
	s.stageNote_TextElementData(inst.Text, counts)
	for _, child := range inst.Elision {
		s.stageNote_Elision(child, counts)
	}
	s.stageNote_Extend(inst.Extend, counts)
	s.stageCommon_Empty(inst.Laughing, counts)
	s.stageCommon_Empty(inst.Humming, counts)
	s.stageCommon_Empty(inst.EndLine, counts)
	s.stageCommon_Empty(inst.EndParagraph, counts)
	s.stageCommon_FormattedText(inst.Footnote, counts)
	s.stageCommon_Level(inst.Level, counts)
}

func (s *MusicXMLStages) stageNote_Mordent(inst *note.Mordent, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Mordent"]++
		inst.Name = fmt.Sprintf("Mordent_%d", counts["note.Mordent"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_NonArpeggiate(inst *note.NonArpeggiate, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.NonArpeggiate"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("NonArpeggiate_%s_%d", inst.Id, counts["note.NonArpeggiate"])
		} else {
			inst.Name = fmt.Sprintf("NonArpeggiate_%d", counts["note.NonArpeggiate"])
		}
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Notations(inst *note.Notations, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Notations"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Notations_%s_%d", inst.Id, counts["note.Notations"])
		} else {
			inst.Name = fmt.Sprintf("Notations_%d", counts["note.Notations"])
		}
	}
	inst.Stage(s.NoteStage)
	s.stageCommon_FormattedText(inst.Footnote, counts)
	s.stageCommon_Level(inst.Level, counts)
	for _, child := range inst.Tied {
		s.stageNote_Tied(child, counts)
	}
	for _, child := range inst.Slur {
		s.stageNote_Slur(child, counts)
	}
	for _, child := range inst.Tuplet {
		s.stageNote_Tuplet(child, counts)
	}
	for _, child := range inst.Glissando {
		s.stageNote_Glissando(child, counts)
	}
	for _, child := range inst.Slide {
		s.stageNote_Slide(child, counts)
	}
	for _, child := range inst.Ornaments {
		s.stageNote_Ornaments(child, counts)
	}
	for _, child := range inst.Technical {
		s.stageNote_Technical(child, counts)
	}
	for _, child := range inst.Articulations {
		s.stageNote_Articulations(child, counts)
	}
	for _, child := range inst.Dynamics {
		s.stageCommon_Dynamics(child, counts)
	}
	for _, child := range inst.Fermata {
		s.stageCommon_Fermata(child, counts)
	}
	for _, child := range inst.Arpeggiate {
		s.stageNote_Arpeggiate(child, counts)
	}
	for _, child := range inst.NonArpeggiate {
		s.stageNote_NonArpeggiate(child, counts)
	}
	for _, child := range inst.AccidentalMark {
		s.stageNote_AccidentalMark(child, counts)
	}
	for _, child := range inst.OtherNotation {
		s.stageNote_OtherNotation(child, counts)
	}
}

func (s *MusicXMLStages) stageNote_Note(inst *note.Note, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Note"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Note_%s_%d", inst.Id, counts["note.Note"])
		} else {
			inst.Name = fmt.Sprintf("Note_%d", counts["note.Note"])
		}
	}
	inst.Stage(s.NoteStage)
	s.stageNote_Grace(inst.Grace, counts)
	s.stageCommon_Empty(inst.Chord, counts)
	s.stageNote_Pitch(inst.Pitch, counts)
	s.stageNote_Unpitched(inst.Unpitched, counts)
	s.stageNote_Rest(inst.Rest, counts)
	for _, child := range inst.Tie {
		s.stageNote_Tie(child, counts)
	}
	s.stageCommon_Empty(inst.Cue, counts)
	for _, child := range inst.Instrument {
		s.stageNote_Instrument(child, counts)
	}
	s.stageCommon_FormattedText(inst.Footnote, counts)
	s.stageCommon_Level(inst.Level, counts)
	s.stageNote_NoteType(inst.Type, counts)
	for _, child := range inst.Dot {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	s.stageNote_Accidental(inst.Accidental, counts)
	s.stageDirection_TimeModification(inst.TimeModification, counts)
	s.stageNote_Stem(inst.Stem, counts)
	s.stageNote_Notehead(inst.Notehead, counts)
	s.stageNote_NoteheadText(inst.NoteheadText, counts)
	for _, child := range inst.Beam {
		s.stageNote_Beam(child, counts)
	}
	for _, child := range inst.Notations {
		s.stageNote_Notations(child, counts)
	}
	for _, child := range inst.Lyric {
		s.stageNote_Lyric(child, counts)
	}
	s.stageCommon_Play(inst.Play, counts)
	s.stageNote_Listen(inst.Listen, counts)
}

func (s *MusicXMLStages) stageNote_NoteType(inst *note.NoteType, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.NoteType"]++
		inst.Name = fmt.Sprintf("NoteType_%d", counts["note.NoteType"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Notehead(inst *note.Notehead, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Notehead"]++
		inst.Name = fmt.Sprintf("Notehead_%d", counts["note.Notehead"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_NoteheadText(inst *note.NoteheadText, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.NoteheadText"]++
		inst.Name = fmt.Sprintf("NoteheadText_%d", counts["note.NoteheadText"])
	}
	inst.Stage(s.NoteStage)
	for _, child := range inst.DisplayText {
		s.stageCommon_FormattedText(child, counts)
	}
	for _, child := range inst.AccidentalText {
		s.stageCommon_AccidentalText(child, counts)
	}
}

func (s *MusicXMLStages) stageNote_Ornaments(inst *note.Ornaments, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Ornaments"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Ornaments_%s_%d", inst.Id, counts["note.Ornaments"])
		} else {
			inst.Name = fmt.Sprintf("Ornaments_%d", counts["note.Ornaments"])
		}
	}
	inst.Stage(s.NoteStage)
	for _, child := range inst.TrillMark {
		s.stageCommon_EmptyTrillSound(child, counts)
	}
	for _, child := range inst.Turn {
		s.stageCommon_HorizontalTurn(child, counts)
	}
	for _, child := range inst.DelayedTurn {
		s.stageCommon_HorizontalTurn(child, counts)
	}
	for _, child := range inst.InvertedTurn {
		s.stageCommon_HorizontalTurn(child, counts)
	}
	for _, child := range inst.DelayedInvertedTurn {
		s.stageCommon_HorizontalTurn(child, counts)
	}
	for _, child := range inst.VerticalTurn {
		s.stageCommon_EmptyTrillSound(child, counts)
	}
	for _, child := range inst.InvertedVerticalTurn {
		s.stageCommon_EmptyTrillSound(child, counts)
	}
	for _, child := range inst.Shake {
		s.stageCommon_EmptyTrillSound(child, counts)
	}
	for _, child := range inst.WavyLine {
		s.stageCommon_WavyLine(child, counts)
	}
	for _, child := range inst.Mordent {
		s.stageNote_Mordent(child, counts)
	}
	for _, child := range inst.InvertedMordent {
		s.stageNote_Mordent(child, counts)
	}
	for _, child := range inst.Schleifer {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.Tremolo {
		s.stageNote_Tremolo(child, counts)
	}
	for _, child := range inst.Haydn {
		s.stageCommon_EmptyTrillSound(child, counts)
	}
	for _, child := range inst.OtherOrnament {
		s.stageNote_OtherPlacementText(child, counts)
	}
	for _, child := range inst.AccidentalMark {
		s.stageNote_AccidentalMark(child, counts)
	}
}

func (s *MusicXMLStages) stageNote_OtherNotation(inst *note.OtherNotation, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.OtherNotation"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("OtherNotation_%s_%d", inst.Id, counts["note.OtherNotation"])
		} else {
			inst.Name = fmt.Sprintf("OtherNotation_%d", counts["note.OtherNotation"])
		}
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_OtherPlacementText(inst *note.OtherPlacementText, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.OtherPlacementText"]++
		inst.Name = fmt.Sprintf("OtherPlacementText_%d", counts["note.OtherPlacementText"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Pitch(inst *note.Pitch, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Pitch"]++
		inst.Name = fmt.Sprintf("Pitch_%d", counts["note.Pitch"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_PlacementText(inst *note.PlacementText, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.PlacementText"]++
		inst.Name = fmt.Sprintf("PlacementText_%d", counts["note.PlacementText"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Release(inst *note.Release, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Release"]++
		inst.Name = fmt.Sprintf("Release_%d", counts["note.Release"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Rest(inst *note.Rest, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Rest"]++
		inst.Name = fmt.Sprintf("Rest_%d", counts["note.Rest"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Slide(inst *note.Slide, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Slide"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Slide_%s_%d", inst.Id, counts["note.Slide"])
		} else {
			inst.Name = fmt.Sprintf("Slide_%d", counts["note.Slide"])
		}
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Slur(inst *note.Slur, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Slur"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Slur_%s_%d", inst.Id, counts["note.Slur"])
		} else {
			inst.Name = fmt.Sprintf("Slur_%d", counts["note.Slur"])
		}
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Stem(inst *note.Stem, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Stem"]++
		inst.Name = fmt.Sprintf("Stem_%d", counts["note.Stem"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_StrongAccent(inst *note.StrongAccent, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.StrongAccent"]++
		inst.Name = fmt.Sprintf("StrongAccent_%d", counts["note.StrongAccent"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Tap(inst *note.Tap, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Tap"]++
		inst.Name = fmt.Sprintf("Tap_%d", counts["note.Tap"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Technical(inst *note.Technical, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Technical"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Technical_%s_%d", inst.Id, counts["note.Technical"])
		} else {
			inst.Name = fmt.Sprintf("Technical_%d", counts["note.Technical"])
		}
	}
	inst.Stage(s.NoteStage)
	for _, child := range inst.UpBow {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.DownBow {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.Harmonic {
		s.stageNote_Harmonic(child, counts)
	}
	for _, child := range inst.OpenString {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.ThumbPosition {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.Fingering {
		s.stageCommon_Fingering(child, counts)
	}
	for _, child := range inst.Pluck {
		s.stageNote_PlacementText(child, counts)
	}
	for _, child := range inst.DoubleTongue {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.TripleTongue {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.Stopped {
		s.stageCommon_EmptyPlacementSmufl(child, counts)
	}
	for _, child := range inst.SnapPizzicato {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.Fret {
		s.stageCommon_Fret(child, counts)
	}
	for _, child := range inst.String {
		s.stageCommon_StringType(child, counts)
	}
	for _, child := range inst.HammerOn {
		s.stageNote_HammerOnPullOff(child, counts)
	}
	for _, child := range inst.PullOff {
		s.stageNote_HammerOnPullOff(child, counts)
	}
	for _, child := range inst.Bend {
		s.stageNote_Bend(child, counts)
	}
	for _, child := range inst.Tap {
		s.stageNote_Tap(child, counts)
	}
	for _, child := range inst.Heel {
		s.stageNote_HeelToe(child, counts)
	}
	for _, child := range inst.Toe {
		s.stageNote_HeelToe(child, counts)
	}
	for _, child := range inst.Fingernails {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.Hole {
		s.stageNote_Hole(child, counts)
	}
	for _, child := range inst.Arrow {
		s.stageNote_Arrow(child, counts)
	}
	for _, child := range inst.Handbell {
		s.stageNote_Handbell(child, counts)
	}
	for _, child := range inst.BrassBend {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.Flip {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.Smear {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.Open {
		s.stageCommon_EmptyPlacementSmufl(child, counts)
	}
	for _, child := range inst.HalfMuted {
		s.stageCommon_EmptyPlacementSmufl(child, counts)
	}
	for _, child := range inst.HarmonMute {
		s.stageNote_HarmonMute(child, counts)
	}
	for _, child := range inst.Golpe {
		s.stageCommon_EmptyPlacement(child, counts)
	}
	for _, child := range inst.OtherTechnical {
		s.stageNote_OtherPlacementText(child, counts)
	}
}

func (s *MusicXMLStages) stageNote_TextElementData(inst *note.TextElementData, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.TextElementData"]++
		inst.Name = fmt.Sprintf("TextElementData_%d", counts["note.TextElementData"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Tie(inst *note.Tie, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Tie"]++
		inst.Name = fmt.Sprintf("Tie_%d", counts["note.Tie"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Tied(inst *note.Tied, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Tied"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Tied_%s_%d", inst.Id, counts["note.Tied"])
		} else {
			inst.Name = fmt.Sprintf("Tied_%d", counts["note.Tied"])
		}
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Tremolo(inst *note.Tremolo, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Tremolo"]++
		inst.Name = fmt.Sprintf("Tremolo_%d", counts["note.Tremolo"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Tuplet(inst *note.Tuplet, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Tuplet"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Tuplet_%s_%d", inst.Id, counts["note.Tuplet"])
		} else {
			inst.Name = fmt.Sprintf("Tuplet_%d", counts["note.Tuplet"])
		}
	}
	inst.Stage(s.NoteStage)
	s.stageNote_TupletPortion(inst.TupletActual, counts)
	s.stageNote_TupletPortion(inst.TupletNormal, counts)
}

func (s *MusicXMLStages) stageNote_TupletDot(inst *note.TupletDot, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.TupletDot"]++
		inst.Name = fmt.Sprintf("TupletDot_%d", counts["note.TupletDot"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_TupletNumber(inst *note.TupletNumber, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.TupletNumber"]++
		inst.Name = fmt.Sprintf("TupletNumber_%d", counts["note.TupletNumber"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_TupletPortion(inst *note.TupletPortion, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.TupletPortion"]++
		inst.Name = fmt.Sprintf("TupletPortion_%d", counts["note.TupletPortion"])
	}
	inst.Stage(s.NoteStage)
	s.stageNote_TupletNumber(inst.TupletNumber, counts)
	s.stageNote_TupletType(inst.TupletType, counts)
	for _, child := range inst.TupletDot {
		s.stageNote_TupletDot(child, counts)
	}
}

func (s *MusicXMLStages) stageNote_TupletType(inst *note.TupletType, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.TupletType"]++
		inst.Name = fmt.Sprintf("TupletType_%d", counts["note.TupletType"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Unpitched(inst *note.Unpitched, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Unpitched"]++
		inst.Name = fmt.Sprintf("Unpitched_%d", counts["note.Unpitched"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageNote_Wait(inst *note.Wait, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["note.Wait"]++
		inst.Name = fmt.Sprintf("Wait_%d", counts["note.Wait"])
	}
	inst.Stage(s.NoteStage)
}

func (s *MusicXMLStages) stageScore_Credit(inst *score.Credit, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.Credit"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Credit_%s_%d", inst.Id, counts["score.Credit"])
		} else {
			inst.Name = fmt.Sprintf("Credit_%d", counts["score.Credit"])
		}
	}
	inst.Stage(s.ScoreStage)
	for _, child := range inst.Link {
		s.stageCommon_Link(child, counts)
	}
	for _, child := range inst.Bookmark {
		s.stageCommon_Bookmark(child, counts)
	}
	s.stageDirection_Image(inst.CreditImage, counts)
	for _, child := range inst.CreditWords {
		s.stageCommon_FormattedTextId(child, counts)
	}
	for _, child := range inst.CreditSymbol {
		s.stageCommon_FormattedSymbolId(child, counts)
	}
}

func (s *MusicXMLStages) stageScore_Defaults(inst *score.Defaults, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.Defaults"]++
		inst.Name = fmt.Sprintf("Defaults_%d", counts["score.Defaults"])
	}
	inst.Stage(s.ScoreStage)
	s.stageCommon_Scaling(inst.Scaling, counts)
	s.stageCommon_Empty(inst.ConcertScore, counts)
	s.stageCommon_PageLayout(inst.PageLayout, counts)
	s.stageCommon_SystemLayout(inst.SystemLayout, counts)
	for _, child := range inst.StaffLayout {
		s.stageCommon_StaffLayout(child, counts)
	}
	s.stageCommon_Appearance(inst.Appearance, counts)
	s.stageScore_EmptyFont(inst.MusicFont, counts)
	s.stageScore_EmptyFont(inst.WordFont, counts)
	for _, child := range inst.LyricFont {
		s.stageScore_LyricFont(child, counts)
	}
	for _, child := range inst.LyricLanguage {
		s.stageScore_LyricLanguage(child, counts)
	}
}

func (s *MusicXMLStages) stageScore_EmptyFont(inst *score.EmptyFont, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.EmptyFont"]++
		inst.Name = fmt.Sprintf("EmptyFont_%d", counts["score.EmptyFont"])
	}
	inst.Stage(s.ScoreStage)
}

func (s *MusicXMLStages) stageScore_GroupBarline(inst *score.GroupBarline, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.GroupBarline"]++
		inst.Name = fmt.Sprintf("GroupBarline_%d", counts["score.GroupBarline"])
	}
	inst.Stage(s.ScoreStage)
}

func (s *MusicXMLStages) stageScore_GroupName(inst *score.GroupName, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.GroupName"]++
		inst.Name = fmt.Sprintf("GroupName_%d", counts["score.GroupName"])
	}
	inst.Stage(s.ScoreStage)
}

func (s *MusicXMLStages) stageScore_GroupSymbol(inst *score.GroupSymbol, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.GroupSymbol"]++
		inst.Name = fmt.Sprintf("GroupSymbol_%d", counts["score.GroupSymbol"])
	}
	inst.Stage(s.ScoreStage)
}

func (s *MusicXMLStages) stageScore_InstrumentLink(inst *score.InstrumentLink, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.InstrumentLink"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("InstrumentLink_%s_%d", inst.Id, counts["score.InstrumentLink"])
		} else {
			inst.Name = fmt.Sprintf("InstrumentLink_%d", counts["score.InstrumentLink"])
		}
	}
	inst.Stage(s.ScoreStage)
}

func (s *MusicXMLStages) stageScore_LyricFont(inst *score.LyricFont, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.LyricFont"]++
		if inst.Number != "" {
			inst.Name = fmt.Sprintf("LyricFont_%s_%d", inst.Number, counts["score.LyricFont"])
		} else {
			inst.Name = fmt.Sprintf("LyricFont_%d", counts["score.LyricFont"])
		}
	}
	inst.Stage(s.ScoreStage)
}

func (s *MusicXMLStages) stageScore_LyricLanguage(inst *score.LyricLanguage, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.LyricLanguage"]++
		if inst.Number != "" {
			inst.Name = fmt.Sprintf("LyricLanguage_%s_%d", inst.Number, counts["score.LyricLanguage"])
		} else {
			inst.Name = fmt.Sprintf("LyricLanguage_%d", counts["score.LyricLanguage"])
		}
	}
	inst.Stage(s.ScoreStage)
}

func (s *MusicXMLStages) stageScore_Opus(inst *score.Opus, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.Opus"]++
		inst.Name = fmt.Sprintf("Opus_%d", counts["score.Opus"])
	}
	inst.Stage(s.ScoreStage)
}

func (s *MusicXMLStages) stageScore_PartGroup(inst *score.PartGroup, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.PartGroup"]++
		if inst.Number != "" {
			inst.Name = fmt.Sprintf("PartGroup_%s_%d", inst.Number, counts["score.PartGroup"])
		} else {
			inst.Name = fmt.Sprintf("PartGroup_%d", counts["score.PartGroup"])
		}
	}
	inst.Stage(s.ScoreStage)
	s.stageScore_GroupName(inst.GroupName, counts)
	s.stageCommon_NameDisplay(inst.GroupNameDisplay, counts)
	s.stageScore_GroupName(inst.GroupAbbreviation, counts)
	s.stageCommon_NameDisplay(inst.GroupAbbreviationDisplay, counts)
	s.stageScore_GroupSymbol(inst.GroupSymbol, counts)
	s.stageScore_GroupBarline(inst.GroupBarline, counts)
	s.stageCommon_Empty(inst.GroupTime, counts)
	s.stageCommon_FormattedText(inst.Footnote, counts)
	s.stageCommon_Level(inst.Level, counts)
}

func (s *MusicXMLStages) stageScore_PartLink(inst *score.PartLink, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.PartLink"]++
		inst.Name = fmt.Sprintf("PartLink_%d", counts["score.PartLink"])
	}
	inst.Stage(s.ScoreStage)
	for _, child := range inst.InstrumentLink {
		s.stageScore_InstrumentLink(child, counts)
	}
}

func (s *MusicXMLStages) stageScore_PartList(inst *score.PartList, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.PartList"]++
		inst.Name = fmt.Sprintf("PartList_%d", counts["score.PartList"])
	}
	inst.Stage(s.ScoreStage)
	for _, child := range inst.PartGroup {
		s.stageScore_PartGroup(child, counts)
	}
	for _, child := range inst.ScorePart {
		s.stageScore_ScorePart(child, counts)
	}
}

func (s *MusicXMLStages) stageScore_PartName(inst *score.PartName, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.PartName"]++
		inst.Name = fmt.Sprintf("PartName_%d", counts["score.PartName"])
	}
	inst.Stage(s.ScoreStage)
}

func (s *MusicXMLStages) stageScore_Player(inst *score.Player, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.Player"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Player_%s_%d", inst.Id, counts["score.Player"])
		} else {
			inst.Name = fmt.Sprintf("Player_%d", counts["score.Player"])
		}
	}
	inst.Stage(s.ScoreStage)
}

func (s *MusicXMLStages) stageScore_ScoreInstrument(inst *score.ScoreInstrument, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.ScoreInstrument"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("ScoreInstrument_%s_%d", inst.Id, counts["score.ScoreInstrument"])
		} else {
			inst.Name = fmt.Sprintf("ScoreInstrument_%d", counts["score.ScoreInstrument"])
		}
	}
	inst.Stage(s.ScoreStage)
	s.stageCommon_Empty(inst.Solo, counts)
	s.stageDirection_VirtualInstrument(inst.VirtualInstrument, counts)
}

func (s *MusicXMLStages) stageScore_ScorePart(inst *score.ScorePart, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.ScorePart"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("ScorePart_%s_%d", inst.Id, counts["score.ScorePart"])
		} else {
			inst.Name = fmt.Sprintf("ScorePart_%d", counts["score.ScorePart"])
		}
	}
	inst.Stage(s.ScoreStage)
	s.stageCommon_Identification(inst.Identification, counts)
	for _, child := range inst.PartLink {
		s.stageScore_PartLink(child, counts)
	}
	s.stageScore_PartName(inst.PartName, counts)
	s.stageCommon_NameDisplay(inst.PartNameDisplay, counts)
	s.stageScore_PartName(inst.PartAbbreviation, counts)
	s.stageCommon_NameDisplay(inst.PartAbbreviationDisplay, counts)
	for _, child := range inst.ScoreInstrument {
		s.stageScore_ScoreInstrument(child, counts)
	}
	for _, child := range inst.Player {
		s.stageScore_Player(child, counts)
	}
	for _, child := range inst.MidiDevice {
		s.stageCommon_MidiDevice(child, counts)
	}
	for _, child := range inst.MidiInstrument {
		s.stageCommon_MidiInstrument(child, counts)
	}
}

func (s *MusicXMLStages) stageScore_Work(inst *score.Work, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.Work"]++
		inst.Name = fmt.Sprintf("Work_%d", counts["score.Work"])
	}
	inst.Stage(s.ScoreStage)
	s.stageScore_Opus(inst.Opus, counts)
}

func (s *MusicXMLStages) stageScore_ScorePartwise(inst *score.ScorePartwise, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.ScorePartwise"]++
		inst.Name = fmt.Sprintf("ScorePartwise_%d", counts["score.ScorePartwise"])
	}
	inst.Stage(s.ScoreStage)
	s.stageScore_Work(inst.Work, counts)
	s.stageCommon_Identification(inst.Identification, counts)
	s.stageScore_Defaults(inst.Defaults, counts)
	for _, child := range inst.Credit {
		s.stageScore_Credit(child, counts)
	}
	s.stageScore_PartList(inst.PartList, counts)
	for _, child := range inst.Part {
		s.stageScore_Part(child, counts)
	}
}

func (s *MusicXMLStages) stageScore_ScoreTimewise(inst *score.ScoreTimewise, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.ScoreTimewise"]++
		inst.Name = fmt.Sprintf("ScoreTimewise_%d", counts["score.ScoreTimewise"])
	}
	inst.Stage(s.ScoreStage)
	s.stageScore_Work(inst.Work, counts)
	s.stageCommon_Identification(inst.Identification, counts)
	s.stageScore_Defaults(inst.Defaults, counts)
	for _, child := range inst.Credit {
		s.stageScore_Credit(child, counts)
	}
	s.stageScore_PartList(inst.PartList, counts)
	for _, child := range inst.Measure {
		s.stageScore_ScoreTimewiseMeasure(child, counts)
	}
}

func (s *MusicXMLStages) stageScore_Part(inst *score.Part, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.Part"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("Part_%s_%d", inst.Id, counts["score.Part"])
		} else {
			inst.Name = fmt.Sprintf("Part_%d", counts["score.Part"])
		}
	}
	inst.Stage(s.ScoreStage)
	for _, child := range inst.Measure {
		s.stageScore_Measure(child, counts)
	}
}

func (s *MusicXMLStages) stageScore_Measure(inst *score.Measure, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.Measure"]++
		if inst.Number != "" {
			inst.Name = fmt.Sprintf("Measure_%s_%d", inst.Number, counts["score.Measure"])
		} else {
			inst.Name = fmt.Sprintf("Measure_%d", counts["score.Measure"])
		}
	}
	inst.Stage(s.ScoreStage)
	for _, child := range inst.MusicData {
		s.stageScore_MeasureMusicDataChoice(child, counts)
	}
}

func (s *MusicXMLStages) stageScore_MeasureMusicDataChoice(inst *score.MeasureMusicDataChoice, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.MeasureMusicDataChoice"]++
		inst.Name = fmt.Sprintf("MeasureMusicDataChoice_%d", counts["score.MeasureMusicDataChoice"])
	}
	inst.Stage(s.ScoreStage)
	s.stageNote_Note(inst.Note, counts)
	s.stageNote_Backup(inst.Backup, counts)
	s.stageNote_Forward(inst.Forward, counts)
	s.stageDirection_Direction(inst.Direction, counts)
	s.stageAttributes_Attributes(inst.Attributes, counts)
	s.stageDirection_Harmony(inst.Harmony, counts)
	s.stageNote_FiguredBass(inst.FiguredBass, counts)
	s.stageDirection_Print(inst.Print, counts)
	s.stageDirection_Sound(inst.Sound, counts)
	s.stageDirection_Listening(inst.Listening, counts)
	s.stageBarline_Barline(inst.Barline, counts)
	s.stageDirection_Grouping(inst.Grouping, counts)
	s.stageCommon_Link(inst.Link, counts)
	s.stageCommon_Bookmark(inst.Bookmark, counts)
}

func (s *MusicXMLStages) stageScore_ScoreTimewiseMeasure(inst *score.ScoreTimewiseMeasure, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.ScoreTimewiseMeasure"]++
		if inst.Number != "" {
			inst.Name = fmt.Sprintf("ScoreTimewiseMeasure_%s_%d", inst.Number, counts["score.ScoreTimewiseMeasure"])
		} else {
			inst.Name = fmt.Sprintf("ScoreTimewiseMeasure_%d", counts["score.ScoreTimewiseMeasure"])
		}
	}
	inst.Stage(s.ScoreStage)
	for _, child := range inst.Part {
		s.stageScore_ScoreTimewiseMeasurePart(child, counts)
	}
}

func (s *MusicXMLStages) stageScore_ScoreTimewiseMeasurePart(inst *score.ScoreTimewiseMeasurePart, counts map[string]int) {
	if inst == nil {
		return
	}
	if inst.Name == "" {
		counts["score.ScoreTimewiseMeasurePart"]++
		if inst.Id != "" {
			inst.Name = fmt.Sprintf("ScoreTimewiseMeasurePart_%s_%d", inst.Id, counts["score.ScoreTimewiseMeasurePart"])
		} else {
			inst.Name = fmt.Sprintf("ScoreTimewiseMeasurePart_%d", counts["score.ScoreTimewiseMeasurePart"])
		}
	}
	inst.Stage(s.ScoreStage)
	for _, child := range inst.Note {
		s.stageNote_Note(child, counts)
	}
	for _, child := range inst.Backup {
		s.stageNote_Backup(child, counts)
	}
	for _, child := range inst.Forward {
		s.stageNote_Forward(child, counts)
	}
	for _, child := range inst.Direction {
		s.stageDirection_Direction(child, counts)
	}
	for _, child := range inst.Attributes {
		s.stageAttributes_Attributes(child, counts)
	}
	for _, child := range inst.Harmony {
		s.stageDirection_Harmony(child, counts)
	}
	for _, child := range inst.FiguredBass {
		s.stageNote_FiguredBass(child, counts)
	}
	for _, child := range inst.Print {
		s.stageDirection_Print(child, counts)
	}
	for _, child := range inst.Sound {
		s.stageDirection_Sound(child, counts)
	}
	for _, child := range inst.Listening {
		s.stageDirection_Listening(child, counts)
	}
	for _, child := range inst.Barline {
		s.stageBarline_Barline(child, counts)
	}
	for _, child := range inst.Grouping {
		s.stageDirection_Grouping(child, counts)
	}
	for _, child := range inst.Link {
		s.stageCommon_Link(child, counts)
	}
	for _, child := range inst.Bookmark {
		s.stageCommon_Bookmark(child, counts)
	}
}

