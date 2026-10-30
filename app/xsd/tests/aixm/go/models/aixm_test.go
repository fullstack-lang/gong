package models

import (
	"os"
	"strings"
	"testing"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/basictypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gml"
)

func TestAIXMMultiStageAndRoundtrip(t *testing.T) {
	xmlBytes, err := os.ReadFile("../../sample_aixm_basic_message.xml")
	if err != nil {
		t.Fatalf("failed reading sample_aixm_basic_message.xml: %v", err)
	}

	stage := NewStage("aixm_test")
	if err := UnmarshallXMLToStage(stage, xmlBytes); err != nil {
		t.Fatalf("UnmarshallXMLToStage failed: %v", err)
	}

	// 1. Verify root AIXMBasicMessage in root stage
	msgs := stage.GetInstancesSorted[*AIXMBasicMessage]()
	if len(msgs) != 1 {
		t.Fatalf("expected 1 AIXMBasicMessage in root stage, got %d", len(msgs))
	}
	if len(msgs[0].HasMember) != 2 {
		t.Fatalf("expected 2 HasMember elements on AIXMBasicMessage, got %d", len(msgs[0].HasMember))
	}

	// 2. Verify cross-package staging into AixmfeaturesStage, AixmdatatypesStage, GmlStage, BasictypesStage
	airports := stage.AixmfeaturesStage.GetInstancesSorted[*aixmfeatures.AirportHeliportType]()
	if len(airports) != 1 {
		t.Fatalf("expected 1 AirportHeliportType in stage.AixmfeaturesStage, got %d", len(airports))
	}
	runways := stage.AixmfeaturesStage.GetInstancesSorted[*aixmfeatures.RunwayType]()
	if len(runways) != 1 {
		t.Fatalf("expected 1 RunwayType in stage.AixmfeaturesStage, got %d", len(runways))
	}
	icaoCodes := stage.AixmdatatypesStage.GetInstancesSorted[*aixmdatatypes.CodeICAOType]()
	if len(icaoCodes) != 1 || icaoCodes[0].EnclosedText != "LFPG" {
		t.Fatalf("expected 1 CodeICAOType ('LFPG') in stage.AixmdatatypesStage, got %v", icaoCodes)
	}
	identifiers := stage.BasictypesStage.GetInstancesSorted[*basictypes.CodeWithAuthorityType]()
	if len(identifiers) != 3 {
		t.Fatalf("expected 3 CodeWithAuthorityType instances in stage.BasictypesStage, got %d", len(identifiers))
	}
	descriptions := stage.GmlStage.GetInstancesSorted[*gml.StringOrRefType]()
	if len(descriptions) != 1 {
		t.Fatalf("expected 1 StringOrRefType in stage.GmlStage, got %d", len(descriptions))
	}

	// 3. Test Multi-Stage Go file roundtrip across all 16 stages
	stage.ComputeReferenceAndOrders()
	goCode, err := stage.MarshallToString("", "main")
	if err != nil {
		t.Fatalf("MarshallToString failed: %v", err)
	}
	_ = os.MkdirAll("../cmd/aixm/data", 0o755)
	_ = os.WriteFile("../cmd/aixm/data/stage.go", []byte(goCode), 0o644)
	if !strings.Contains(goCode, "AixmfeaturesStage") || !strings.Contains(goCode, "AixmdatatypesStage") {
		t.Fatalf("expected Multi-Stage Go code to reference AixmfeaturesStage and AixmdatatypesStage")
	}

	stage2 := NewStage("aixm_roundtrip")
	if err := stage2.ParseAstString(goCode, true); err != nil {
		t.Fatalf("ParseAstString failed: %v", err)
	}

	airports2 := stage2.AixmfeaturesStage.GetInstancesSorted[*aixmfeatures.AirportHeliportType]()
	if len(airports2) != 1 {
		t.Fatalf("expected 1 AirportHeliportType in stage2.AixmfeaturesStage after Go roundtrip, got %d", len(airports2))
	}
	icaoCodes2 := stage2.AixmdatatypesStage.GetInstancesSorted[*aixmdatatypes.CodeICAOType]()
	if len(icaoCodes2) != 1 || icaoCodes2[0].EnclosedText != "LFPG" {
		t.Fatalf("expected CodeICAOType 'LFPG' in stage2.AixmdatatypesStage after Go roundtrip, got %v", icaoCodes2)
	}

	// 4. Test XML export from roundtripped stage2
	exportedXML, err := MarshallStageToXML(stage2)
	if err != nil {
		t.Fatalf("MarshallStageToXML failed: %v", err)
	}
	if !strings.Contains(string(exportedXML), "LFPG") || !strings.Contains(string(exportedXML), "09L/27R") {
		t.Fatalf("expected exported XML to contain 'LFPG' and '09L/27R', got:\n%s", string(exportedXML))
	}
}

