package models

import (
	"os"
	"strings"
	"testing"
)

func TestMusicXMLMultiStageAndRoundtrip(t *testing.T) {
	samples := []string{
		"../../DebuMandSample.musicxml",
		"../../SchbAvMaSample.musicxml",
	}

	for _, samplePath := range samples {
		t.Run(samplePath, func(t *testing.T) {
			xmlBytes, err := os.ReadFile(samplePath)
			if err != nil {
				t.Fatalf("failed reading %s: %v", samplePath, err)
			}

			stage := NewStage("musicxml_test")
			if err := UnmarshallXMLToStage(stage, xmlBytes); err != nil {
				t.Fatalf("UnmarshallXMLToStage failed on %s: %v", samplePath, err)
			}

			scores := stage.GetInstancesSorted[*Score_partwise]()
			if len(scores) != 1 {
				t.Fatalf("expected 1 Score_partwise root instance, got %d", len(scores))
			}

			notes := stage.GetInstancesSorted[*Note]()
			if len(notes) == 0 {
				t.Fatalf("expected staged Note instances from %s, got 0", samplePath)
			}
			t.Logf("%s: staged 1 Score_partwise and %d Note instances", samplePath, len(notes))

			// Verify sub-stages are initialized on Stage
			if stage.XlinkStage == nil || stage.XmlxsdStage == nil {
				t.Fatalf("expected XlinkStage and XmlxsdStage to be initialized on root Stage")
			}

			// Evaluate XSD semantic constraints
			_ = ValidateXSDConstraints(stage)

			// Multi-Stage Go AST Roundtrip
			stage.ComputeReferenceAndOrders()
			goCode, err := stage.MarshallToString("", "main")
			if err != nil {
				t.Fatalf("MarshallToString failed on %s: %v", samplePath, err)
			}
			if strings.Contains(samplePath, "DebuMandSample") {
				_ = os.MkdirAll("../cmd/musicxml/data", 0o755)
				_ = os.WriteFile("../cmd/musicxml/data/stage.go", []byte(goCode), 0o644)
			}

			stage2 := NewStage("musicxml_roundtrip")
			if err := stage2.ParseAstString(goCode, true); err != nil {
				t.Fatalf("ParseAstString failed on %s: %v", samplePath, err)
			}

			notes2 := stage2.GetInstancesSorted[*Note]()
			if len(notes2) != len(notes) {
				t.Fatalf("expected %d Note instances after Go AST roundtrip, got %d", len(notes), len(notes2))
			}

			// XML Export Roundtrip
			exportedXML, err := MarshallStageToXML(stage2)
			if err != nil {
				t.Fatalf("MarshallStageToXML failed on %s: %v", samplePath, err)
			}
			if !strings.Contains(string(exportedXML), "<score-partwise") {
				t.Fatalf("expected exported XML to contain <score-partwise, got:\n%.300s", string(exportedXML))
			}
		})
	}
}

