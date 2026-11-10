package tests

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"io"
	"path/filepath"
	"strings"
	"testing"

	test_stack "github.com/fullstack-lang/gong/test/test1/go/stack"
	test_static "github.com/fullstack-lang/gong/test/test1/go/static"

	"github.com/fullstack-lang/gong/test/test1/go/models"
)

func TestSerialize(t *testing.T) {

	r := test_static.ServeStaticFiles(false)

	// setup stack
	stack := test_stack.NewTranscientStack(r, "test", false)
	stage := stack.Stage

	bclass1 := (&models.Bstruct{Name: "B1"}).Stage(stage)
	bclass2 := (&models.Bstruct{Name: "B2"}).Stage(stage)

	aclass1 := (&models.Astruct{
		Name:                "A1",
		Floatfield:          10.2,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
	}).Stage(stage)

	aclass1.Anarrayofb = append(aclass1.Anarrayofb, bclass2)
	aclass1.Anarrayofb = append(aclass1.Anarrayofb, bclass1)

	aclass3 := (&models.Astruct{
		Name:                "A3",
		Floatfield:          10.3,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
	}).Stage(stage)

	aclass3.Anarrayofb = append(aclass3.Anarrayofb, bclass2)

	aclass2 := (&models.Astruct{
		Name:                "A2",
		Floatfield:          10.2,
		Booleanfield:        true,
		Anotherbooleanfield: true,
		Associationtob:      bclass1,
	}).Stage(stage)

	aclass2.Anarrayofb = append(aclass2.Anarrayofb, bclass2)

	stage.SerializeStage("serialize_test.xlsx")
	stage.SerializeStage2("serialize_test-withIds.xlsx", true)
}

func readZipCSVs(t *testing.T, zipBytes []byte) map[string][][]string {
	t.Helper()
	zr, err := zip.NewReader(bytes.NewReader(zipBytes), int64(len(zipBytes)))
	if err != nil {
		t.Fatalf("failed to open zip bytes: %v", err)
	}
	res := make(map[string][][]string, len(zr.File))
	for _, f := range zr.File {
		rc, err := f.Open()
		if err != nil {
			t.Fatalf("failed to open zip entry %s: %v", f.Name, err)
		}
		records, err := csv.NewReader(rc).ReadAll()
		_ = rc.Close()
		if err != nil {
			t.Fatalf("failed to parse CSV %s: %v", f.Name, err)
		}
		res[f.Name] = records
	}
	return res
}

func TestSerializeCSV(t *testing.T) {
	r := test_static.ServeStaticFiles(false)
	stack := test_stack.NewTranscientStack(r, "test_csv", false)
	stage := stack.Stage

	b1 := (&models.Bstruct{Name: "B1"}).Stage(stage)
	b2 := (&models.Bstruct{Name: "B2"}).Stage(stage)

	// > 32,000 character payload (exceeds Excel's 32,767 cell limit) with commas, quotes, and newlines
	largeText := strings.Repeat("line1, \"quoted\"; value\n", 2000) // 46,000 chars
	if len(largeText) <= 32767 {
		t.Fatalf("expected largeText > 32767 chars, got %d", len(largeText))
	}

	a1 := (&models.Astruct{
		Name:           "A1",
		Floatfield:     10.2,
		Booleanfield:   true,
		Associationtob: b1,
		Anarrayofb:     []*models.Bstruct{b2, b1},
		TextArea:       largeText,
	}).Stage(stage)
	_ = a1

	a2 := (&models.Astruct{
		Name:           "A2",
		Floatfield:     20.5,
		Associationtob: b2,
	}).Stage(stage)
	_ = a2

	stage.ComputeReferenceAndOrders()

	tmpDir := t.TempDir()
	csvZipNoIDs := filepath.Join(tmpDir, "serialize_test.zip")
	csvZipWithIDs := filepath.Join(tmpDir, "serialize_test_with_ids.zip")

	stage.SerializeStageCSV(csvZipNoIDs)
	stage.SerializeStageCSV2(csvZipWithIDs, true)

	// 1. Verify file output without IDs
	zrNoIDs, err := zip.OpenReader(csvZipNoIDs)
	if err != nil {
		t.Fatalf("failed to open %s: %v", csvZipNoIDs, err)
	}
	defer zrNoIDs.Close()

	var bufNoIDs bytes.Buffer
	for _, f := range zrNoIDs.File {
		if f.Name == "Astruct.csv" {
			rc, err := f.Open()
			if err != nil {
				t.Fatal(err)
			}
			_, err = io.Copy(&bufNoIDs, rc)
			_ = rc.Close()
			if err != nil {
				t.Fatal(err)
			}
		}
	}
	astructNoIDsRecords, err := csv.NewReader(&bufNoIDs).ReadAll()
	if err != nil {
		t.Fatalf("failed to read Astruct.csv (no IDs): %v", err)
	}
	if len(astructNoIDsRecords) != 3 {
		t.Fatalf("expected 3 rows (header + 2 instances), got %d", len(astructNoIDsRecords))
	}
	headersNoIDs := astructNoIDsRecords[0]
	colIdx := make(map[string]int, len(headersNoIDs))
	for i, h := range headersNoIDs {
		colIdx[h] = i
	}
	if got := astructNoIDsRecords[1][colIdx["Name"]]; got != "A1" {
		t.Errorf("expected row 1 Name=A1, got %q", got)
	}
	if got := astructNoIDsRecords[1][colIdx["TextArea"]]; got != largeText {
		t.Errorf("largeText mismatch: got len %d, want len %d", len(got), len(largeText))
	}
	if got := astructNoIDsRecords[1][colIdx["Anarrayofb"]]; got != "B2\nB1" {
		t.Errorf("expected Anarrayofb=%q, got %q", "B2\nB1", got)
	}

	// 2. Verify SerializeStageCSVAsBytes with IDs (and full >31-char struct name entry)
	zipBytesWithIDs, err := stage.SerializeStageCSVAsBytes(true)
	if err != nil {
		t.Fatalf("SerializeStageCSVAsBytes(true) failed: %v", err)
	}
	csvMap := readZipCSVs(t, zipBytesWithIDs)

	// Empty struct with >31 chars in name should be present with header row
	longStructCSV, ok := csvMap["F0123456789012345678901234567890.csv"]
	if !ok {
		t.Fatalf("expected F0123456789012345678901234567890.csv in zip")
	}
	if len(longStructCSV) != 1 {
		t.Errorf("expected 1 header row for empty Gongstruct, got %d", len(longStructCSV))
	}

	astructWithIDs := csvMap["Astruct.csv"]
	if len(astructWithIDs) != 3 {
		t.Fatalf("expected 3 rows in Astruct.csv with IDs, got %d", len(astructWithIDs))
	}
	headersWithIDs := astructWithIDs[0]
	if len(headersWithIDs) != 2*len(headersNoIDs) {
		t.Fatalf("expected %d columns with IDs, got %d", 2*len(headersNoIDs), len(headersWithIDs))
	}
	idColIdx := make(map[string]int, len(headersWithIDs))
	for i, h := range headersWithIDs {
		idColIdx[h] = i
	}
	if headersWithIDs[1] != "ID" {
		t.Errorf("expected second column header 'ID', got %q", headersWithIDs[1])
	}
	if astructWithIDs[1][1] == "" {
		t.Errorf("expected non-empty UUID for A1")
	}
	if got := astructWithIDs[1][idColIdx["Associationtob:Bstruct:ID"]]; got != "deadbeef-dead-4bef-a555-000000000000" {
		t.Errorf("unexpected Associationtob:Bstruct:ID: %q", got)
	}
	if got := astructWithIDs[1][idColIdx["Anarrayofb:Bstruct:IDs"]]; got != "deadbeef-dead-4bef-a555-000000000000;deadbeef-dead-4bef-a555-000000000000" {
		t.Errorf("unexpected Anarrayofb:Bstruct:IDs: %q", got)
	}
	if got := astructWithIDs[1][idColIdx["TextArea"]]; got != largeText {
		t.Errorf("with IDs: largeText mismatch: got len %d, want len %d", len(got), len(largeText))
	}

	// 3. Verify per-struct SerializeCSVPointer2 helper
	var singleCSV bytes.Buffer
	if err := stage.SerializeCSVPointer2[*models.Astruct](&singleCSV, true); err != nil {
		t.Fatalf("SerializeCSVPointer2 failed: %v", err)
	}
	singleRecords, err := csv.NewReader(&singleCSV).ReadAll()
	if err != nil {
		t.Fatalf("failed to parse single struct CSV: %v", err)
	}
	if len(singleRecords) != 3 {
		t.Errorf("expected 3 rows from SerializeCSVPointer2, got %d", len(singleRecords))
	}
}
