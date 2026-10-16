// generated code by xsd2gong - do not edit
package orchestrator

import (
	"fmt"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/xlink"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/common"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/barline"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/attributes"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/direction"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/note"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/score"
)

var _ = fmt.Errorf

// ValidateSemantics runs semantic validation across all package stages.
func (s *MusicXMLStages) ValidateSemantics() (errs []error) {
	errs = append(errs, xlink.ValidateStageSemantics(s.XlinkStage)...)
	errs = append(errs, common.ValidateStageSemantics(s.CommonStage)...)
	errs = append(errs, barline.ValidateStageSemantics(s.BarlineStage)...)
	errs = append(errs, attributes.ValidateStageSemantics(s.AttributesStage)...)
	errs = append(errs, direction.ValidateStageSemantics(s.DirectionStage)...)
	errs = append(errs, note.ValidateStageSemantics(s.NoteStage)...)
	errs = append(errs, score.ValidateStageSemantics(s.ScoreStage)...)
	return errs
}

// ResolveReferences indexes all xs:ID instances across stages, wires navigable Gong pointers
// for xs:IDREF attributes (e.g. Part.RefScorePart -> *ScorePart), and reports broken references or duplicate IDs.
func (s *MusicXMLStages) ResolveReferences() (errs []error) {
	seenIDs := make(map[string]string)
	_ = seenIDs
	idMap_Player := make(map[string]*score.Player)
	if s.ScoreStage != nil {
		for _, inst := range s.ScoreStage.GetInstancesSorted[*score.Player]() {
			if inst.Id != "" {
				idMap_Player[inst.Id] = inst
			}
		}
	}
	idMap_ScoreInstrument := make(map[string]*score.ScoreInstrument)
	if s.ScoreStage != nil {
		for _, inst := range s.ScoreStage.GetInstancesSorted[*score.ScoreInstrument]() {
			if inst.Id != "" {
				idMap_ScoreInstrument[inst.Id] = inst
			}
		}
	}
	idMap_ScorePart := make(map[string]*score.ScorePart)
	if s.ScoreStage != nil {
		for _, inst := range s.ScoreStage.GetInstancesSorted[*score.ScorePart]() {
			if inst.Id != "" {
				idMap_ScorePart[inst.Id] = inst
			}
		}
	}
	if s.ScoreStage != nil {
		for _, inst := range s.ScoreStage.GetInstancesSorted[*score.InstrumentLink]() {
			if inst.Id != "" {
				if target, ok := idMap_ScoreInstrument[inst.Id]; ok {
					inst.RefScoreInstrument = target
				} else {
					errs = append(errs, fmt.Errorf("InstrumentLink %q: unresolved xs:IDREF id=%q (target ScoreInstrument not found)", inst.Name, "id", inst.Id))
				}
			}
		}
	}
	if s.ScoreStage != nil {
		for _, inst := range s.ScoreStage.GetInstancesSorted[*score.Part]() {
			if inst.Id != "" {
				if target, ok := idMap_ScorePart[inst.Id]; ok {
					inst.RefScorePart = target
				} else {
					errs = append(errs, fmt.Errorf("Part %q: unresolved xs:IDREF id=%q (target ScorePart not found)", inst.Name, "id", inst.Id))
				}
			}
		}
	}
	if s.ScoreStage != nil {
		for _, inst := range s.ScoreStage.GetInstancesSorted[*score.ScoreTimewiseMeasurePart]() {
			if inst.Id != "" {
				if target, ok := idMap_ScorePart[inst.Id]; ok {
					inst.RefScorePart = target
				} else {
					errs = append(errs, fmt.Errorf("ScoreTimewiseMeasurePart %q: unresolved xs:IDREF id=%q (target ScorePart not found)", inst.Name, "id", inst.Id))
				}
			}
		}
	}
	return errs
}

// SyncReferences synchronizes navigable Gong pointers (e.g. Part.RefScorePart) back to their
// raw XML IDREF string attributes prior to XML serialization.
func (s *MusicXMLStages) SyncReferences() {
	if s.ScoreStage != nil {
		for _, inst := range s.ScoreStage.GetInstancesSorted[*score.InstrumentLink]() {
			if inst.RefScoreInstrument != nil && inst.RefScoreInstrument.Id != "" {
				inst.Id = inst.RefScoreInstrument.Id
			}
		}
	}
	if s.ScoreStage != nil {
		for _, inst := range s.ScoreStage.GetInstancesSorted[*score.Part]() {
			if inst.RefScorePart != nil && inst.RefScorePart.Id != "" {
				inst.Id = inst.RefScorePart.Id
			}
		}
	}
	if s.ScoreStage != nil {
		for _, inst := range s.ScoreStage.GetInstancesSorted[*score.ScoreTimewiseMeasurePart]() {
			if inst.RefScorePart != nil && inst.RefScorePart.Id != "" {
				inst.Id = inst.RefScorePart.Id
			}
		}
	}
}
