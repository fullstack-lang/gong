package models

import (
	"os"
	"strings"
	"testing"
)

func TestReqIFMultiStageAndRoundtrip(t *testing.T) {
	xmlBytes, err := os.ReadFile("../../samples/sample.reqif")
	if err != nil {
		t.Fatalf("failed reading sample.reqif: %v", err)
	}

	stage := NewStage("reqif_test")
	if err := UnmarshallXMLToStage(stage, xmlBytes); err != nil {
		t.Fatalf("UnmarshallXMLToStage failed: %v", err)
	}

	roots := stage.GetInstancesSorted[*REQ_IF]()
	if len(roots) != 1 {
		t.Fatalf("expected 1 REQ_IF root instance, got %d", len(roots))
	}

	stage.ComputeReferenceAndOrders()
	goCode, err := stage.MarshallToString("", "main")
	if err != nil {
		t.Fatalf("MarshallToString failed: %v", err)
	}
	_ = os.MkdirAll("../cmd/reqif/data", 0o755)
	_ = os.WriteFile("../cmd/reqif/data/stage.go", []byte(goCode), 0o644)

	stage2 := NewStage("reqif_roundtrip")
	if err := stage2.ParseAstString(goCode, true); err != nil {
		t.Fatalf("ParseAstString failed: %v", err)
	}

	exportedXML, err := MarshallStageToXML(stage2)
	if err != nil {
		t.Fatalf("MarshallStageToXML failed: %v", err)
	}
	if !strings.Contains(string(exportedXML), "REQ-IF") {
		t.Fatalf("expected exported XML to contain REQ-IF, got:\n%.300s", string(exportedXML))
	}
}

