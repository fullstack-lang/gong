package models

import (
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

func TestMarshallAndUnmarshallEmbeddedName(t *testing.T) {
	stage := NewStage("test_marshall")

	g := (&Gstruct{Name: "embedded_name_test"}).Stage(stage)
	stage.Commit()

	if g.Name != "embedded_name_test" {
		t.Fatalf("expected Name to be 'embedded_name_test', got '%s'", g.Name)
	}

	marshalled, err := stage.MarshallToString("github.com/fullstack-lang/gong/test/test1/go/models", "main")
	if err != nil {
		t.Fatalf("MarshallToString failed: %v", err)
	}

	// Verify that the marshalled declaration includes the Name in the struct literal
	expectedDecl := "(&models.Gstruct{Name: `embedded_name_test`}).Stage(stage)"
	if !strings.Contains(marshalled, expectedDecl) {
		t.Fatalf("expected marshalled code to contain %q, but got:\n%s", expectedDecl, marshalled)
	}

	// Verify that unmarshalling that code back restores the stage correctly
	newStage := NewStage("test_unmarshall")
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "stage.go", marshalled, parser.ParseComments)
	if err != nil {
		t.Fatalf("failed to parse marshalled code: %v", err)
	}

	err = ParseAstFileFromAst(newStage, file, fset, false)
	if err != nil {
		t.Fatalf("ParseAstFileFromAst failed: %v", err)
	}

	var found *Gstruct
	for item := range newStage.Gstructs {
		found = item
		break
	}

	if found == nil {
		t.Fatalf("expected to find Gstruct in unmarshalled stage, found none")
	}
	if found.Name != "embedded_name_test" {
		t.Fatalf("expected unmarshalled Gstruct Name to be 'embedded_name_test', got '%s'", found.Name)
	}
}

func TestGetAssociationNameWithPromotedFields(t *testing.T) {
	dstruct := GetAssociationName[Dstruct]()
	if dstruct == nil {
		t.Fatal("GetAssociationName[Dstruct]() returned nil")
	}

	if dstruct.Gstruct == nil {
		t.Fatal("expected dstruct.Gstruct to not be nil")
	}
	if dstruct.Gstruct.Name != "Gstruct" {
		t.Fatalf("expected dstruct.Gstruct.Name to be 'Gstruct', got '%s'", dstruct.Gstruct.Name)
	}

	if len(dstruct.Gstructs) != 1 {
		t.Fatalf("expected 1 Gstruct in dstruct.Gstructs, got %d", len(dstruct.Gstructs))
	}

	if dstruct.Gstructs[0].Name != "Gstructs" {
		t.Fatalf("expected dstruct.Gstructs[0].Name to be 'Gstructs', got '%s'", dstruct.Gstructs[0].Name)
	}
}
