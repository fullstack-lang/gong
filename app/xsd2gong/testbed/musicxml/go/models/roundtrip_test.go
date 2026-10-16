package models_test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/attributes"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/note"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/orchestrator"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/score"
)

func TestCanonicalMusicXMLRoundTrip(t *testing.T) {
	samples := []string{
		"DebuMandSample.musicxml",
		"MozaChloSample.musicxml",
		"SchbAvMaSample.musicxml",
		"phylotaxy.musicxml",
	}

	for _, sample := range samples {
		t.Run(sample, func(t *testing.T) {
			xmlPath := filepath.Join("..", "..", sample)

			// 1. Unmarshal into multi-package Gong stages, resolve IDREFs, and validate semantics
			stages, root, errs, err := orchestrator.UnmarshalScorePartwiseFile(xmlPath)
			if err != nil {
				t.Fatalf("UnmarshalScorePartwiseFile(%s) failed: %v", sample, err)
			}
			if root == nil {
				t.Fatalf("expected non-nil ScorePartwise root for %s", sample)
			}
			if len(errs) > 0 {
				var msgs []string
				for _, e := range errs {
					msgs = append(msgs, e.Error())
				}
				t.Fatalf("expected 0 semantic/IDREF errors on canonical score %s, got %d:\n%s",
					sample, len(errs), strings.Join(msgs, "\n"))
			}

			// 2. Verify cross-package stages are populated
			if len(stages.ScoreStage.ScorePartwises) != 1 {
				t.Errorf("expected 1 ScorePartwise in ScoreStage, got %d", len(stages.ScoreStage.ScorePartwises))
			}
			if len(stages.ScoreStage.Parts) == 0 {
				t.Errorf("expected >0 Parts in ScoreStage for %s", sample)
			}
			if len(stages.ScoreStage.Measures) == 0 {
				t.Errorf("expected >0 Measures in ScoreStage for %s", sample)
			}
			if len(stages.NoteStage.Notes) == 0 {
				t.Errorf("expected >0 Notes in NoteStage for %s", sample)
			}
			if len(stages.AttributesStage.GetInstancesSorted[*attributes.Attributes]()) == 0 {
				t.Errorf("expected >0 Attributes in AttributesStage for %s", sample)
			}

			// 3. Verify xs:ID / xs:IDREF pointer resolution (Part.RefScorePart -> *ScorePart)
			for _, p := range root.Part {
				if p.Id == "" {
					t.Errorf("expected non-empty Part.Id")
				}
				if p.RefScorePart == nil {
					t.Errorf("expected Part %q RefScorePart pointer to be resolved, got nil", p.Id)
				} else if p.RefScorePart.Id != p.Id {
					t.Errorf("expected Part.RefScorePart.Id == %q, got %q", p.Id, p.RefScorePart.Id)
				}
			}

			// 4. Verify StageSet Commit succeeds across all 7 packages in topological order
			stages.Commit()

			// 5. Verify lossless XML round-trip serialization & measure timeline order preservation
			marshalledBytes, err := orchestrator.MarshalScorePartwise(stages, root)
			if err != nil {
				t.Fatalf("MarshalScorePartwise(%s) failed: %v", sample, err)
			}
			if len(marshalledBytes) == 0 {
				t.Fatalf("expected non-empty marshalled XML for %s", sample)
			}

			// Re-unmarshal the serialized XML and verify structural & musical equivalence
			stages2, root2, errs2, err := orchestrator.UnmarshalScorePartwise(marshalledBytes)
			if err != nil {
				t.Fatalf("re-unmarshalling round-trip XML for %s failed: %v", sample, err)
			}
			if len(errs2) > 0 {
				t.Errorf("expected 0 semantic errors on round-tripped XML for %s, got %v", sample, errs2)
			}
			if len(root2.Part) != len(root.Part) {
				t.Fatalf("part count mismatch after round-trip: got %d, want %d", len(root2.Part), len(root.Part))
			}
			for i := range root.Part {
				p1, p2 := root.Part[i], root2.Part[i]
				if p1.Id != p2.Id {
					t.Errorf("part[%d] Id mismatch: got %q, want %q", i, p2.Id, p1.Id)
				}
				if len(p1.Measure) != len(p2.Measure) {
					t.Fatalf("part[%d] measure count mismatch: got %d, want %d", i, len(p2.Measure), len(p1.Measure))
				}
				for j := range p1.Measure {
					m1, m2 := p1.Measure[j], p2.Measure[j]
					if m1.Number != m2.Number {
						t.Errorf("part[%d] measure[%d] Number mismatch: got %q, want %q", i, j, m2.Number, m1.Number)
					}
					if len(m1.MusicData) != len(m2.MusicData) {
						t.Fatalf("part[%d] measure[%d] (%s) MusicData timeline length mismatch: got %d, want %d",
							i, j, m1.Number, len(m2.MusicData), len(m1.MusicData))
					}
					// Verify exact timeline sequence of choice items inside each measure
					for k := range m1.MusicData {
						c1, c2 := m1.MusicData[k], m2.MusicData[k]
						if (c1.Note != nil) != (c2.Note != nil) ||
							(c1.Backup != nil) != (c2.Backup != nil) ||
							(c1.Forward != nil) != (c2.Forward != nil) ||
							(c1.Direction != nil) != (c2.Direction != nil) ||
							(c1.Attributes != nil) != (c2.Attributes != nil) ||
							(c1.Harmony != nil) != (c2.Harmony != nil) ||
							(c1.Barline != nil) != (c2.Barline != nil) ||
							(c1.Print != nil) != (c2.Print != nil) ||
							(c1.Sound != nil) != (c2.Sound != nil) {
							t.Fatalf("part[%d] measure[%d] item[%d] timeline element kind mismatch after round-trip", i, j, k)
						}
					}
				}
			}

			// Verify second marshalling is byte-for-byte deterministic
			marshalledBytes2, err := orchestrator.MarshalScorePartwise(stages2, root2)
			if err != nil {
				t.Fatalf("second MarshalScorePartwise(%s) failed: %v", sample, err)
			}
			if string(marshalledBytes) != string(marshalledBytes2) {
				t.Errorf("expected deterministic idempotent XML output on second marshal for %s", sample)
			}
		})
	}
}

func TestSemanticValidationAndIDREFSync(t *testing.T) {
	stages := orchestrator.NewMusicXMLStages()

	sp := (&score.ScorePart{
		Name: "ScorePart_P1",
		Id:   "P1",
	}).Stage(stages.ScoreStage)

	part := (&score.Part{
		Name: "Part_P1",
		Id:   "P1",
	}).Stage(stages.ScoreStage)

	// 1. Resolve references and verify Part.RefScorePart points to sp
	if errs := stages.ResolveReferences(); len(errs) != 0 {
		t.Fatalf("unexpected ResolveReferences errors: %v", errs)
	}
	if part.RefScorePart != sp {
		t.Fatalf("expected part.RefScorePart == sp, got %v", part.RefScorePart)
	}

	// 2. Mutate sp.Id and verify SyncReferences updates part.Id
	sp.Id = "P_RENAMED"
	stages.SyncReferences()
	if part.Id != "P_RENAMED" {
		t.Errorf("expected SyncReferences to update part.Id to P_RENAMED, got %q", part.Id)
	}

	// 3. Test broken IDREF detection
	part.Id = "NON_EXISTENT_PART"
	if errs := stages.ResolveReferences(); len(errs) == 0 {
		t.Errorf("expected ResolveReferences to report error for broken IDREF NON_EXISTENT_PART")
	}
	part.Id = "P_RENAMED"

	// 4. Test XSD numeric restriction facet validation (KeyOctave.Value in [0, 9]) and regex pattern (LineDetail.Color)
	invalidOctave := (&attributes.KeyOctave{
		Name:  "InvalidOctave",
		Value: 12, // violates XSD octave restriction [0, 9]
	}).Stage(stages.AttributesStage)

	semErrs := stages.ValidateSemantics()
	if len(semErrs) == 0 {
		t.Errorf("expected ValidateSemantics to report error for KeyOctave.Value=12")
	}
	invalidOctave.Value = 4 // valid octave
	if errs := invalidOctave.ValidateSemantics(); len(errs) != 0 {
		t.Errorf("expected KeyOctave.Value=4 to be valid, got %v", errs)
	}

	invalidLineDetail := (&attributes.LineDetail{
		Name:  "InvalidLineDetail",
		Color: "NOT_A_HEX_COLOR", // violates XSD color regex pattern #[0-9A-F]{6}([0-9A-F]{2})?
	}).Stage(stages.AttributesStage)
	if errs := invalidLineDetail.ValidateSemantics(); len(errs) == 0 {
		t.Errorf("expected ValidateSemantics to report regex pattern error for invalid Color")
	}
	invalidLineDetail.Color = "#FF00AA"
	if errs := invalidLineDetail.ValidateSemantics(); len(errs) != 0 {
		t.Errorf("expected valid hex color #FF00AA to pass ValidateSemantics, got %v", errs)
	}

	// 5. Test xs:choice XOR exclusivity on MeasureMusicDataChoice
	invalidChoice := (&score.MeasureMusicDataChoice{
		Name:   "InvalidChoice",
		Note:   &note.Note{Name: "N1"},
		Backup: &note.Backup{Name: "B1"},
	}).Stage(stages.ScoreStage)

	choiceErrs := invalidChoice.ValidateSemantics()
	if len(choiceErrs) == 0 {
		t.Errorf("expected ValidateSemantics to report xs:choice XOR violation when both Note and Backup are set")
	}
}
