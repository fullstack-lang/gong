package models

import (
	"os"
	"strings"
	"testing"

	"github.com/fullstack-lang/gong/app/xsd/tests/complexrules/go/models/refs"
)

func TestComplexRulesMultiStageAndSemantics(t *testing.T) {
	xmlBytes, err := os.ReadFile("../../sample_valid.xml")
	if err != nil {
		t.Fatalf("failed reading sample_valid.xml: %v", err)
	}

	stage := NewStage("test_complex")
	if err := UnmarshallXMLToStage(stage, xmlBytes); err != nil {
		t.Fatalf("UnmarshallXMLToStage failed: %v", err)
	}

	// 1. Verify multi-stage population (root stage + RefsStage)
	catalogs := stage.GetInstancesSorted[*Catalog]()
	if len(catalogs) != 1 {
		t.Fatalf("expected 1 Catalog in root stage, got %d", len(catalogs))
	}
	targetsInRefsStage := stage.RefsStage.GetInstancesSorted[*refs.RefTargetType]()
	if len(targetsInRefsStage) != 2 {
		t.Fatalf("expected 2 RefTargetType instances in stage.RefsStage, got %d", len(targetsInRefsStage))
	}
	items := stage.GetInstancesSorted[*DerivedItemType]()
	if len(items) != 2 {
		t.Fatalf("expected 2 DerivedItemType instances in root stage, got %d", len(items))
	}

	// 2. Verify valid sample passes all XSD semantic constraints
	errs := ValidateXSDConstraints(stage)
	if len(errs) != 0 {
		t.Fatalf("expected 0 semantic errors on sample_valid.xml, got: %v", errs)
	}

	// 3. Test Multi-Stage Go file roundtrip (root + sub-stages in one Go file)
	stage.ComputeReferenceAndOrders()
	goCode, err := stage.MarshallToString("", "main")
	if err != nil {
		t.Fatalf("MarshallToString failed: %v", err)
	}
	_ = os.MkdirAll("../cmd/complexrules/data", 0o755)
	_ = os.WriteFile("../cmd/complexrules/data/stage.go", []byte(goCode), 0o644)
	if !strings.Contains(goCode, "RefsStage") {
		t.Fatalf("expected Multi-Stage Go code to reference RefsStage, got:\n%s", goCode)
	}

	stage2 := NewStage("test_complex_roundtrip")
	if err := stage2.ParseAstString(goCode, true); err != nil {
		t.Fatalf("ParseAstString failed: %v", err)
	}
	if len(stage2.RefsStage.GetInstancesSorted[*refs.RefTargetType]()) != 2 {
		t.Fatalf("expected 2 instances in stage2.RefsStage after Go AST roundtrip, got %d", len(stage2.RefsStage.GetInstancesSorted[*refs.RefTargetType]()))
	}

	// 4. Test XML export from roundtripped stage2
	xmlOut, err := MarshallStageToXML(stage2)
	if err != nil {
		t.Fatalf("MarshallStageToXML failed: %v", err)
	}
	if !strings.Contains(string(xmlOut), "ITEM_01") || !strings.Contains(string(xmlOut), "Primary Target") {
		t.Fatalf("exported XML missing expected content:\n%s", string(xmlOut))
	}

	// 5. Test XSD semantic constraint violations:
	// 5a. whiteSpace="collapse" normalization
	item0 := items[0]
	item0.Code = "  ITEM_01  "
	errs = ValidateXSDConstraints(stage)
	if len(errs) != 0 {
		t.Fatalf("expected whiteSpace collapse to normalize Code without error, got: %v", errs)
	}
	if item0.Code != "ITEM_01" {
		t.Fatalf("expected whiteSpace collapse to normalize Code to 'ITEM_01', got %q", item0.Code)
	}

	// 5b. xs:choice mutual exclusivity violation (both optionA and optionB populated)
	item0.OptionB = 99
	errs = ValidateXSDConstraints(stage)
	assertErrorSubstring(t, errs, "xs:choice exclusivity violation")
	item0.OptionB = 0

	// 5c. xs:IDREF unresolved reference violation
	origRef := item0.TargetRef
	item0.TargetRef = "UNKNOWN-ID"
	errs = ValidateXSDConstraints(stage)
	assertErrorSubstring(t, errs, "xs:IDREF unresolved reference")
	item0.TargetRef = origRef

	// 5d. xs:ID duplicate value violation
	origID := catalogs[0].Target[1].Id
	catalogs[0].Target[1].Id = catalogs[0].Target[0].Id
	errs = ValidateXSDConstraints(stage)
	assertErrorSubstring(t, errs, "xs:ID duplicate value")
	catalogs[0].Target[1].Id = origID

	// 5e. Numeric facet minExclusive / maxExclusive violation on Score (0 < score < 100)
	origScore := item0.Score
	item0.Score = "0"
	errs = ValidateXSDConstraints(stage)
	assertErrorSubstring(t, errs, "minExclusive violation")
	item0.Score = "100"
	errs = ValidateXSDConstraints(stage)
	assertErrorSubstring(t, errs, "maxExclusive violation")
	item0.Score = origScore

	// 5f. Pattern & minLength violation on Code
	origCode := item0.Code
	item0.Code = "ab"
	errs = ValidateXSDConstraints(stage)
	assertErrorSubstring(t, errs, "minLength violation")
	assertErrorSubstring(t, errs, "pattern violation")
	item0.Code = origCode

	// 5g. Required attribute from embedded cross-package AttributeGroup missing
	origCreatedBy := catalogs[0].Target[0].CreatedBy
	catalogs[0].Target[0].CreatedBy = ""
	errs = ValidateXSDConstraints(stage)
	assertErrorSubstring(t, errs, "required attribute \"createdBy\" is missing")
	catalogs[0].Target[0].CreatedBy = origCreatedBy
}

func assertErrorSubstring(t *testing.T, errs []string, sub string) {
	t.Helper()
	for _, e := range errs {
		if strings.Contains(e, sub) {
			return
		}
	}
	t.Fatalf("expected error containing %q, got: %v", sub, errs)
}
