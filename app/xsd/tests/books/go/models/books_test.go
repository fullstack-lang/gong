package models

import (
	"os"
	"strings"
	"testing"
)

func TestBooksXMLAndGoRoundtripAndSemanticRules(t *testing.T) {
	xmlBytes, err := os.ReadFile("../../books.xml")
	if err != nil {
		t.Fatalf("failed reading books.xml: %v", err)
	}

	// 1. XML Drop-in (UnmarshallXMLToStage)
	stage1 := NewStage("books_test_1")
	if err := UnmarshallXMLToStage(stage1, xmlBytes); err != nil {
		t.Fatalf("UnmarshallXMLToStage failed: %v", err)
	}

	books := stage1.GetInstancesSorted[*BookType]()
	if len(books) != 2 {
		t.Fatalf("expected 2 BookType instances, got %d", len(books))
	}
	credits := stage1.GetInstancesSorted[*Credit]()
	if len(credits) != 3 {
		t.Fatalf("expected 3 Credit instances, got %d", len(credits))
	}
	links := stage1.GetInstancesSorted[*Link]()
	if len(links) != 3 {
		t.Fatalf("expected 3 Link instances, got %d", len(links))
	}

	// Verify semantic validation passes on valid books.xml
	errs := ValidateXSDConstraints(stage1)
	if len(errs) != 0 {
		t.Fatalf("expected 0 semantic errors on valid books.xml, got: %v", errs)
	}

	// 2. XML Export & Roundtrip (MarshallStageToXML -> UnmarshallXMLToStage)
	exportedXML, err := MarshallStageToXML(stage1)
	if err != nil {
		t.Fatalf("MarshallStageToXML failed: %v", err)
	}
	stageXMLRoundtrip := NewStage("books_xml_rt")
	if err := UnmarshallXMLToStage(stageXMLRoundtrip, exportedXML); err != nil {
		t.Fatalf("roundtrip UnmarshallXMLToStage failed: %v", err)
	}
	if len(stageXMLRoundtrip.GetInstancesSorted[*BookType]()) != 2 {
		t.Fatalf("expected 2 books after XML roundtrip")
	}

	// 3. Go Multi-Stage File Export & Drop-in Roundtrip (MarshallToString -> ParseAstString)
	stage1.ComputeReferenceAndOrders()
	goCode, err := stage1.MarshallToString("", "main")
	if err != nil {
		t.Fatalf("MarshallToString failed: %v", err)
	}
	_ = os.MkdirAll("../cmd/books/data", 0o755)
	_ = os.WriteFile("../cmd/books/data/stage.go", []byte(goCode), 0o644)
	stageGoRoundtrip := NewStage("books_go_rt")
	if err := stageGoRoundtrip.ParseAstString(goCode, true); err != nil {
		t.Fatalf("ParseAstString failed: %v", err)
	}
	rtBooks := stageGoRoundtrip.GetInstancesSorted[*BookType]()
	if len(rtBooks) != 2 {
		t.Fatalf("expected 2 books after Go AST roundtrip, got %d", len(rtBooks))
	}
	if len(ValidateXSDConstraints(stageGoRoundtrip)) != 0 {
		t.Fatalf("expected 0 semantic errors after Go AST roundtrip")
	}

	// 4. Semantic Constraint Violation Checks
	invalidBook := (&BookType{
		Name: "InvalidBook",
	}).Stage(stageGoRoundtrip)
	invalidBook.Isbn = ""                        // violates use="required"
	invalidBook.Title = "Bad!Title@With#Symbols" // violates pattern="[0-9A-Za-z ]+"
	invalidBook.Year = 999                       // violates minInclusive="1000"

	semErrs := ValidateXSDConstraints(stageGoRoundtrip)
	if len(semErrs) != 3 {
		t.Fatalf("expected 3 semantic errors on invalidBook, got %d: %v", len(semErrs), semErrs)
	}

	joined := strings.Join(semErrs, "\n")
	if !strings.Contains(joined, "required attribute \"isbn\" is missing") {
		t.Errorf("expected required attribute error, got: %s", joined)
	}
	if !strings.Contains(joined, "pattern violation on BookType.Title") {
		t.Errorf("expected pattern violation error, got: %s", joined)
	}
	if !strings.Contains(joined, "minInclusive violation on BookType.Year") {
		t.Errorf("expected minInclusive violation error, got: %s", joined)
	}
}

